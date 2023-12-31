// Code generated by trpc-go/trpc-go-cmdline v2.0.13. DO NOT EDIT.
// source: isp_discern.proto

package tme_isp

import (
	"context"
	"fmt"

	_ "git.tencent.com/trpc-go/trpc-go"
	"git.tencent.com/trpc-go/trpc-go/client"
	"git.tencent.com/trpc-go/trpc-go/codec"
	_ "git.tencent.com/trpc-go/trpc-go/http"
	"git.tencent.com/trpc-go/trpc-go/server"
)

/* ************************************ Service Definition ************************************ */

// IspDiscernService defines service
type IspDiscernService interface {
	GetIpInfo(ctx context.Context, req *IpInfoReq) (*IpInfoRsp, error)
}

func IspDiscernService_GetIpInfo_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &IpInfoReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(IspDiscernService).GetIpInfo(ctx, reqbody.(*IpInfoReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// IspDiscernServer_ServiceDesc descriptor for server.RegisterService
var IspDiscernServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.tme_songbank.isp_discern.IspDiscern",
	HandlerType: ((*IspDiscernService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.tme_songbank.isp_discern.IspDiscern/GetIpInfo",
			Func: IspDiscernService_GetIpInfo_Handler,
		},
	},
}

// RegisterIspDiscernService register service
func RegisterIspDiscernService(s server.Service, svr IspDiscernService) {
	if err := s.Register(&IspDiscernServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("IspDiscern register error:%v", err))
	}
}

/* ************************************ Client Definition ************************************ */

// IspDiscernClientProxy defines service client proxy
type IspDiscernClientProxy interface {
	GetIpInfo(ctx context.Context, req *IpInfoReq, opts ...client.Option) (rsp *IpInfoRsp, err error)
}

type IspDiscernClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewIspDiscernClientProxy = func(opts ...client.Option) IspDiscernClientProxy {
	return &IspDiscernClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *IspDiscernClientProxyImpl) GetIpInfo(ctx context.Context, req *IpInfoReq, opts ...client.Option) (*IpInfoRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.tme_songbank.isp_discern.IspDiscern/GetIpInfo")
	msg.WithCalleeServiceName(IspDiscernServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("tme_songbank")
	msg.WithCalleeServer("isp_discern")
	msg.WithCalleeService("IspDiscern")
	msg.WithCalleeMethod("GetIpInfo")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &IpInfoRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}
