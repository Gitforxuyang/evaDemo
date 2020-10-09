
package handler

import (
	"context"
	"evaDemo/proto/evaDemo"
	"fmt"
	"github.com/Gitforxuyang/eva/util/utils"
)

type HandlerService struct {

}

func (m *HandlerService) Hello(ctx context.Context,req *evaDemo.HelloReq) (*evaDemo.Nil, error) {
	fmt.Println(utils.StructToJson(req))
	return &evaDemo.Nil{}, nil
}

func (m *HandlerService) Ping(context.Context, *evaDemo.Nil) (*evaDemo.Nil, error) {
	return &evaDemo.Nil{}, nil
}

