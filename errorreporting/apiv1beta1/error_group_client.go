// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package errorreporting

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	errorreportingpb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newErrorGroupClientHook clientHook

// ErrorGroupCallOptions contains the retry settings for each method of ErrorGroupClient.
type ErrorGroupCallOptions struct {
	GetGroup    []gax.CallOption
	UpdateGroup []gax.CallOption
}

func defaultErrorGroupGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("clouderrorreporting.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("clouderrorreporting.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultErrorGroupCallOptions() *ErrorGroupCallOptions {
	return &ErrorGroupCallOptions{
		GetGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultErrorGroupRESTCallOptions() *ErrorGroupCallOptions {
	return &ErrorGroupCallOptions{
		GetGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		UpdateGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
	}
}

// internalErrorGroupClient is an interface that defines the methods available from Error Reporting API.
type internalErrorGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetGroup(context.Context, *errorreportingpb.GetGroupRequest, ...gax.CallOption) (*errorreportingpb.ErrorGroup, error)
	UpdateGroup(context.Context, *errorreportingpb.UpdateGroupRequest, ...gax.CallOption) (*errorreportingpb.ErrorGroup, error)
}

// ErrorGroupClient is a client for interacting with Error Reporting API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for retrieving and updating individual error groups.
type ErrorGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalErrorGroupClient

	// The call options for this service.
	CallOptions *ErrorGroupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ErrorGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ErrorGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ErrorGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetGroup get the specified group.
func (c *ErrorGroupClient) GetGroup(ctx context.Context, req *errorreportingpb.GetGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	return c.internalClient.GetGroup(ctx, req, opts...)
}

// UpdateGroup replace the data for the specified group.
// Fails if the group does not exist.
func (c *ErrorGroupClient) UpdateGroup(ctx context.Context, req *errorreportingpb.UpdateGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	return c.internalClient.UpdateGroup(ctx, req, opts...)
}

// errorGroupGRPCClient is a client for interacting with Error Reporting API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type errorGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ErrorGroupClient
	CallOptions **ErrorGroupCallOptions

	// The gRPC API client.
	errorGroupClient errorreportingpb.ErrorGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewErrorGroupClient creates a new error group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for retrieving and updating individual error groups.
func NewErrorGroupClient(ctx context.Context, opts ...option.ClientOption) (*ErrorGroupClient, error) {
	clientOpts := defaultErrorGroupGRPCClientOptions()
	if newErrorGroupClientHook != nil {
		hookOpts, err := newErrorGroupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ErrorGroupClient{CallOptions: defaultErrorGroupCallOptions()}

	c := &errorGroupGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		errorGroupClient: errorreportingpb.NewErrorGroupServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *errorGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *errorGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *errorGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type errorGroupRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ErrorGroupClient
	CallOptions **ErrorGroupCallOptions
}

// NewErrorGroupRESTClient creates a new error group service rest client.
//
// Service for retrieving and updating individual error groups.
func NewErrorGroupRESTClient(ctx context.Context, opts ...option.ClientOption) (*ErrorGroupClient, error) {
	clientOpts := append(defaultErrorGroupRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultErrorGroupRESTCallOptions()
	c := &errorGroupRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ErrorGroupClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultErrorGroupRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://clouderrorreporting.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://clouderrorreporting.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *errorGroupRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *errorGroupRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *errorGroupRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *errorGroupGRPCClient) GetGroup(ctx context.Context, req *errorreportingpb.GetGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group_name", url.QueryEscape(req.GetGroupName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGroup[0:len((*c.CallOptions).GetGroup):len((*c.CallOptions).GetGroup)], opts...)
	var resp *errorreportingpb.ErrorGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorGroupClient.GetGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *errorGroupGRPCClient) UpdateGroup(ctx context.Context, req *errorreportingpb.UpdateGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group.name", url.QueryEscape(req.GetGroup().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateGroup[0:len((*c.CallOptions).UpdateGroup):len((*c.CallOptions).UpdateGroup)], opts...)
	var resp *errorreportingpb.ErrorGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorGroupClient.UpdateGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetGroup get the specified group.
func (c *errorGroupRESTClient) GetGroup(ctx context.Context, req *errorreportingpb.GetGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetGroupName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group_name", url.QueryEscape(req.GetGroupName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetGroup[0:len((*c.CallOptions).GetGroup):len((*c.CallOptions).GetGroup)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &errorreportingpb.ErrorGroup{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// UpdateGroup replace the data for the specified group.
// Fails if the group does not exist.
func (c *errorGroupRESTClient) UpdateGroup(ctx context.Context, req *errorreportingpb.UpdateGroupRequest, opts ...gax.CallOption) (*errorreportingpb.ErrorGroup, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetGroup()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetGroup().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group.name", url.QueryEscape(req.GetGroup().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).UpdateGroup[0:len((*c.CallOptions).UpdateGroup):len((*c.CallOptions).UpdateGroup)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &errorreportingpb.ErrorGroup{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
