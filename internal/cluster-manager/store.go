package clustermanager

import (
	"sync"

	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
)

var mtx sync.Mutex

var clusterStore = make(map[string]types.NodeInfo)

func cacheNodeInfo(agentIP string, nodeInfo types.NodeInfo) {
	mtx.Lock()
	defer mtx.Unlock()

	clusterStore[agentIP] = nodeInfo
}

func getNodeInfo(agentIP string) (types.NodeInfo, bool) {
	mtx.Lock()
	defer mtx.Unlock()

	nodeInfo, ok := clusterStore[agentIP]
	return nodeInfo, ok
}
