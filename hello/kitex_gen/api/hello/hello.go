// Code generated by Kitex v0.8.0. DO NOT EDIT.

package hello

import (
	"context"
	api "github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

var helloServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*api.Hello)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo": kitex.NewMethodInfo(echoHandler, newHelloEchoArgs, newHelloEchoResult, false),
		"add":  kitex.NewMethodInfo(addHandler, newHelloAddArgs, newHelloAddResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "api",
		"ServiceFilePath": `hello.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloEchoArgs)
	realResult := result.(*api.HelloEchoResult)
	success, err := handler.(api.Hello).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloEchoArgs() interface{} {
	return api.NewHelloEchoArgs()
}

func newHelloEchoResult() interface{} {
	return api.NewHelloEchoResult()
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloAddArgs)
	realResult := result.(*api.HelloAddResult)
	success, err := handler.(api.Hello).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloAddArgs() interface{} {
	return api.NewHelloAddArgs()
}

func newHelloAddResult() interface{} {
	return api.NewHelloAddResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.HelloEchoArgs
	_args.Req = req
	var _result api.HelloEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Add(ctx context.Context, req *api.AddRequest) (r *api.AddResponse, err error) {
	var _args api.HelloAddArgs
	_args.Req = req
	var _result api.HelloAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
