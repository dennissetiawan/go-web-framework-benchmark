package logic

import (
	"context"

	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/svc"
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{Message: "Hello World"}, nil
}
