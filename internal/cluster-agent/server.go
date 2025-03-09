package clusteragent

import (
	"context"

	"github.com/gin-gonic/gin"
)

func StartServer(ctx context.Context, addr string) error {
	r := gin.Default()
	return r.Run(addr)
}
