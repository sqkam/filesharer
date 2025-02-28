// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.19.6
// source: file/v1/file.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFileDownloadByAddrHttp = "/api.file.v1.File/DownloadByAddrHttp"
const OperationFileDownloadDirByAddrHttp = "/api.file.v1.File/DownloadDirByAddrHttp"
const OperationFileGetDetailByAddrHttp = "/api.file.v1.File/GetDetailByAddrHttp"
const OperationFileListByAddrHttp = "/api.file.v1.File/ListByAddrHttp"
const OperationFileListNode = "/api.file.v1.File/ListNode"

type FileHTTPServer interface {
	DownloadByAddrHttp(context.Context, *DownloadByAddrRequest) (*DownloadByAddrReply, error)
	DownloadDirByAddrHttp(context.Context, *DownloadDirByAddrRequest) (*DownloadDirByAddrReply, error)
	GetDetailByAddrHttp(context.Context, *GetDetailByAddrRequest) (*GetDetailByAddrReply, error)
	ListByAddrHttp(context.Context, *ListByAddrRequest) (*ListByAddrReply, error)
	ListNode(context.Context, *ListNodeRequest) (*ListNodeReply, error)
}

func RegisterFileHTTPServer(s *http.Server, srv FileHTTPServer) {
	r := s.Route("/")
	r.POST("v1/file/list", _File_ListByAddrHttp0_HTTP_Handler(srv))
	r.POST("v1/file/downloadfile", _File_DownloadByAddrHttp0_HTTP_Handler(srv))
	r.POST("v1/file/downloaddir", _File_DownloadDirByAddrHttp0_HTTP_Handler(srv))
	r.POST("v1/node/list", _File_ListNode0_HTTP_Handler(srv))
	r.POST("v1/file/detail", _File_GetDetailByAddrHttp0_HTTP_Handler(srv))
}

func _File_ListByAddrHttp0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListByAddrRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileListByAddrHttp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListByAddrHttp(ctx, req.(*ListByAddrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListByAddrReply)
		return ctx.Result(200, reply)
	}
}

func _File_DownloadByAddrHttp0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadByAddrRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileDownloadByAddrHttp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadByAddrHttp(ctx, req.(*DownloadByAddrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadByAddrReply)
		return ctx.Result(200, reply)
	}
}

func _File_DownloadDirByAddrHttp0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadDirByAddrRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileDownloadDirByAddrHttp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadDirByAddrHttp(ctx, req.(*DownloadDirByAddrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadDirByAddrReply)
		return ctx.Result(200, reply)
	}
}

func _File_ListNode0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListNodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileListNode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNode(ctx, req.(*ListNodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListNodeReply)
		return ctx.Result(200, reply)
	}
}

func _File_GetDetailByAddrHttp0_HTTP_Handler(srv FileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDetailByAddrRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFileGetDetailByAddrHttp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDetailByAddrHttp(ctx, req.(*GetDetailByAddrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDetailByAddrReply)
		return ctx.Result(200, reply)
	}
}

type FileHTTPClient interface {
	DownloadByAddrHttp(ctx context.Context, req *DownloadByAddrRequest, opts ...http.CallOption) (rsp *DownloadByAddrReply, err error)
	DownloadDirByAddrHttp(ctx context.Context, req *DownloadDirByAddrRequest, opts ...http.CallOption) (rsp *DownloadDirByAddrReply, err error)
	GetDetailByAddrHttp(ctx context.Context, req *GetDetailByAddrRequest, opts ...http.CallOption) (rsp *GetDetailByAddrReply, err error)
	ListByAddrHttp(ctx context.Context, req *ListByAddrRequest, opts ...http.CallOption) (rsp *ListByAddrReply, err error)
	ListNode(ctx context.Context, req *ListNodeRequest, opts ...http.CallOption) (rsp *ListNodeReply, err error)
}

type FileHTTPClientImpl struct {
	cc *http.Client
}

func NewFileHTTPClient(client *http.Client) FileHTTPClient {
	return &FileHTTPClientImpl{client}
}

func (c *FileHTTPClientImpl) DownloadByAddrHttp(ctx context.Context, in *DownloadByAddrRequest, opts ...http.CallOption) (*DownloadByAddrReply, error) {
	var out DownloadByAddrReply
	pattern := "v1/file/downloadfile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileDownloadByAddrHttp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) DownloadDirByAddrHttp(ctx context.Context, in *DownloadDirByAddrRequest, opts ...http.CallOption) (*DownloadDirByAddrReply, error) {
	var out DownloadDirByAddrReply
	pattern := "v1/file/downloaddir"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileDownloadDirByAddrHttp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) GetDetailByAddrHttp(ctx context.Context, in *GetDetailByAddrRequest, opts ...http.CallOption) (*GetDetailByAddrReply, error) {
	var out GetDetailByAddrReply
	pattern := "v1/file/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileGetDetailByAddrHttp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) ListByAddrHttp(ctx context.Context, in *ListByAddrRequest, opts ...http.CallOption) (*ListByAddrReply, error) {
	var out ListByAddrReply
	pattern := "v1/file/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileListByAddrHttp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FileHTTPClientImpl) ListNode(ctx context.Context, in *ListNodeRequest, opts ...http.CallOption) (*ListNodeReply, error) {
	var out ListNodeReply
	pattern := "v1/node/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFileListNode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
