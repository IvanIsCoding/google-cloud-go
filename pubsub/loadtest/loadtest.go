package loadtest

import (
	"bytes"
	"errors"
	"log"
	"strconv"
	"sync/atomic"
	"time"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes"

	"cloud.google.com/go/pubsub"
	pb "cloud.google.com/go/pubsub/loadtest/pb"
)

type serverConfig struct {
	topic     *pubsub.Topic
	msgData   []byte
	batchSize int32
}

type Server struct {
	ID string

	cfg    atomic.Value
	seqNum int32
}

func (l *Server) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Println("received start")
	c, err := pubsub.NewClient(ctx, req.Project)
	if err != nil {
		return nil, err
	}
	dur, err := ptypes.Duration(req.PublishBatchDuration)
	if err != nil {
		return nil, err
	}
	l.init(c, req.Topic, req.MessageSize, req.PublishBatchSize, dur)
	log.Println("started")
	return &pb.StartResponse{}, nil
}

func (l *Server) init(c *pubsub.Client, topicName string, msgSize, batchSize int32, batchDur time.Duration) {
	topic := c.Topic(topicName)
	topic.PublishSettings = pubsub.PublishSettings{
		DelayThreshold:    batchDur,
		CountThreshold:    950,
		ByteThreshold:     9500000,
		BufferedByteLimit: 1e9,
	}

	l.cfg.Store(serverConfig{
		topic:     topic,
		msgData:   bytes.Repeat([]byte{'A'}, int(msgSize)),
		batchSize: batchSize,
	})
}

func (l *Server) Execute(ctx context.Context, _ *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	latencies, err := l.publishBatch()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	return &pb.ExecuteResponse{Latencies: latencies}, nil
}

func (l *Server) publishBatch() ([]int64, error) {
	var cfg serverConfig
	if c, ok := l.cfg.Load().(serverConfig); ok {
		cfg = c
	} else {
		return nil, errors.New("config not loaded")
	}

	start := time.Now()
	index := int64(cfg.batchSize)
	latencies := make([]int64, cfg.batchSize)
	startStr := strconv.FormatInt(start.UnixNano()/1e6, 10)
	seqNum := atomic.AddInt32(&l.seqNum, cfg.batchSize) - cfg.batchSize

	var rs []*pubsub.PublishResult
	for i := int32(0); i < cfg.batchSize; i++ {
		r := cfg.topic.Publish(context.TODO(), &pubsub.Message{
			Data: cfg.msgData,
			Attributes: map[string]string{
				"sendTime":       startStr,
				"clientId":       l.ID,
				"sequenceNumber": strconv.Itoa(int(seqNum + i)),
			},
		})
		rs = append(rs, r)
	}
	for _, r := range rs {
		_, err := r.Get(context.Background())
		if err != nil {
			return nil, err
		}
		i := atomic.AddInt64(&index, -1)
		// TODO(jba,pongad): fix latencies
		// Later values will be skewed by earlier ones, since we wait for the
		// results in order. (On the other hand, it may not matter much, since
		// messages are added to bundles in order and bundles get sent more or
		// less in order.) If we want more accurate values, we can either start
		// a goroutine for each result (similar to the original code using a
		// callback), or call reflect.Select with the Ready channels of the
		// results.
		latencies[i] = time.Since(start).Nanoseconds() / 1e6
	}
	return latencies, nil
}
