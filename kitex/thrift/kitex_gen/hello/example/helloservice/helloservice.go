// Code generated by Kitex v0.7.3. DO NOT EDIT.

package helloservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	example "kitex-examples/kitex/thrift/kitex_gen/hello/example"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

var helloServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*example.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{
		"HelloMethod": kitex.NewMethodInfo(helloMethodHandler, newHelloServiceHelloMethodArgs, newHelloServiceHelloMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "example",
		"ServiceFilePath": `idl/hello.proto`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func helloMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*example.HelloServiceHelloMethodArgs)
	realResult := result.(*example.HelloServiceHelloMethodResult)
	success, err := handler.(example.HelloService).HelloMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloServiceHelloMethodArgs() interface{} {
	return example.NewHelloServiceHelloMethodArgs()
}

func newHelloServiceHelloMethodResult() interface{} {
	return example.NewHelloServiceHelloMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) HelloMethod(ctx context.Context, request *example.HelloReq) (r *example.HelloResp, err error) {
	var _args example.HelloServiceHelloMethodArgs
	_args.Request = request
	var _result example.HelloServiceHelloMethodResult
	if err = p.c.Call(ctx, "HelloMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
