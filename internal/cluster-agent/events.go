package clusteragent

import (
	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
	clusterhelper "caas-ai-agent.ccsp.zte.com.cn/pkg/cluster-helper"
	"github.com/gin-gonic/gin"
)

func eventsListHandler(c *gin.Context) {
	ctx := c.Request.Context()
	ns := c.Query("namespace")
	filter := c.Query("filter")

	ch, err := clusterhelper.New("")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	events, err := ch.ListEvents(ctx, ns, filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := types.EventListResp{Events: events}
	c.JSON(200, resp)

}
