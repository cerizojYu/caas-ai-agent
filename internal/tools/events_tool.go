package tools

import (
	"fmt"

	clustermanager "caas-ai-agent.ccsp.zte.com.cn/internal/cluster-manager"
)

const EventsListToolName = `EventsListTool`

const EventsListToolDescription = `
Use this tool for listing  events in specify k8s cluster.
	example:
		在1.2.3.4这个集群中，列出default命名空间中的所有事件
	then Action Input is: 1.2.3.4, default
`

const EventsListToolParam = `{"type":"object","properties":{"agentIP":{"type":"string","description":"集群代理的IP地址"},"ns":{"type":"string","description":"命名空间"}},"required":["agentIP","ns"]}`

func EventsListTool(agentIP string, ns string) string {
	events, err := clustermanager.ListEvents(agentIP, ns, "")
	if err != nil {
		return fmt.Sprintf("服务器返回错误: %s", err)
	}
	return fmt.Sprintf("事件列表: %v", events)
}
