package logic

import (
	"context"

	"cluster/internal/svc"
	"cluster/types/cluster"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClusterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterLogic {
	return &ClusterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterLogic) Cluster(in *cluster.ClusterRequest) (*cluster.ClusterResponse, error) {
	// todo: add your logic here and delete this line

	return &cluster.ClusterResponse{}, nil
}
