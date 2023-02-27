package logic

import (
	"context"

	"new-greet/internal/svc"
	"new-greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailUserLogic {
	return &DetailUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailUserLogic) DetailUser(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	logx.Info("Here!")
	var r *types.Response
	r = &types.Response{}
	r.Message = "HORSE"
	return r, nil
}
