package clustermanager

import (
	"context"
	"fmt"

	"caas-ai-agent.ccsp.zte.com.cn/internal/consts"
	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
	"github.com/gin-gonic/gin"
)

func StartServer(ctx context.Context, listenPort int) error {
	r := gin.Default()
	r.POST(consts.ClusterManagerHeartbeatURI, heartbeatHandler)
	return r.Run(fmt.Sprintf(":%d", listenPort))
}

func heartbeatHandler(c *gin.Context) {
	req := types.HeartBeatReq{}
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		// 无效请求
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	// 处理心跳请求
	cacheNodeInfo(req.Agent.IP.String(), req)
}
