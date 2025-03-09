package clusteragent

import (
	"context"
	"fmt"

	"caas-ai-agent.ccsp.zte.com.cn/internal/consts"
	"github.com/gin-gonic/gin"
)

func StartServer(ctx context.Context, listenPort int) error {
	r := gin.Default()
	r.GET(consts.ClusterAgentEventsURI, eventsListHandler)
	return r.Run(fmt.Sprintf(":%d", listenPort))
}
