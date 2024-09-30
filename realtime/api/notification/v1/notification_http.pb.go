// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: realtime/api/notification/v1/notification.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNotificationServiceDeleteNotification = "/notification.v1.NotificationService/DeleteNotification"
const OperationNotificationServiceGetNotification = "/notification.v1.NotificationService/GetNotification"
const OperationNotificationServiceListNotification = "/notification.v1.NotificationService/ListNotification"
const OperationNotificationServiceReadNotification = "/notification.v1.NotificationService/ReadNotification"

type NotificationServiceHTTPServer interface {
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*DeleteNotificationReply, error)
	GetNotification(context.Context, *GetNotificationRequest) (*Notification, error)
	ListNotification(context.Context, *ListNotificationRequest) (*ListNotificationReply, error)
	// ReadNotificationReadNotification set notification as read
	ReadNotification(context.Context, *ReadNotificationRequest) (*emptypb.Empty, error)
}

func RegisterNotificationServiceHTTPServer(s *http.Server, srv NotificationServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/realtime/notification/list", _NotificationService_ListNotification0_HTTP_Handler(srv))
	r.GET("/v1/realtime/notifications", _NotificationService_ListNotification1_HTTP_Handler(srv))
	r.GET("/v1/realtime/notification/{id}", _NotificationService_GetNotification0_HTTP_Handler(srv))
	r.PUT("/v1/realtime/notification/{id}/read", _NotificationService_ReadNotification0_HTTP_Handler(srv))
	r.DELETE("/v1/realtime/notification/{id}", _NotificationService_DeleteNotification0_HTTP_Handler(srv))
}

func _NotificationService_ListNotification0_HTTP_Handler(srv NotificationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListNotificationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationServiceListNotification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNotification(ctx, req.(*ListNotificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListNotificationReply)
		return ctx.Result(200, reply)
	}
}

func _NotificationService_ListNotification1_HTTP_Handler(srv NotificationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListNotificationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationServiceListNotification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNotification(ctx, req.(*ListNotificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListNotificationReply)
		return ctx.Result(200, reply)
	}
}

func _NotificationService_GetNotification0_HTTP_Handler(srv NotificationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNotificationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationServiceGetNotification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNotification(ctx, req.(*GetNotificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Notification)
		return ctx.Result(200, reply)
	}
}

func _NotificationService_ReadNotification0_HTTP_Handler(srv NotificationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReadNotificationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationServiceReadNotification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReadNotification(ctx, req.(*ReadNotificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _NotificationService_DeleteNotification0_HTTP_Handler(srv NotificationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteNotificationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationServiceDeleteNotification)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNotification(ctx, req.(*DeleteNotificationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteNotificationReply)
		return ctx.Result(200, reply)
	}
}

type NotificationServiceHTTPClient interface {
	DeleteNotification(ctx context.Context, req *DeleteNotificationRequest, opts ...http.CallOption) (rsp *DeleteNotificationReply, err error)
	GetNotification(ctx context.Context, req *GetNotificationRequest, opts ...http.CallOption) (rsp *Notification, err error)
	ListNotification(ctx context.Context, req *ListNotificationRequest, opts ...http.CallOption) (rsp *ListNotificationReply, err error)
	ReadNotification(ctx context.Context, req *ReadNotificationRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type NotificationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewNotificationServiceHTTPClient(client *http.Client) NotificationServiceHTTPClient {
	return &NotificationServiceHTTPClientImpl{client}
}

func (c *NotificationServiceHTTPClientImpl) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...http.CallOption) (*DeleteNotificationReply, error) {
	var out DeleteNotificationReply
	pattern := "/v1/realtime/notification/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationServiceDeleteNotification))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationServiceHTTPClientImpl) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...http.CallOption) (*Notification, error) {
	var out Notification
	pattern := "/v1/realtime/notification/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationServiceGetNotification))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationServiceHTTPClientImpl) ListNotification(ctx context.Context, in *ListNotificationRequest, opts ...http.CallOption) (*ListNotificationReply, error) {
	var out ListNotificationReply
	pattern := "/v1/realtime/notifications"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationServiceListNotification))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotificationServiceHTTPClientImpl) ReadNotification(ctx context.Context, in *ReadNotificationRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/realtime/notification/{id}/read"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotificationServiceReadNotification))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}