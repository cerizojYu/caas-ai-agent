package clusteragent

import (
	"context"
	"fmt"
	"time"

	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
	"caas-ai-agent.ccsp.zte.com.cn/pkg/utils"
)

// 上报心跳间隔
const heartbeatInterval = time.Minute

// 开启死循环，定期上报心跳
func StartHeartbeatLoop(ctx context.Context) {
	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():

		case <-ticker.C:
			err := reportHeartbeat(ctx)
			if err != nil {

			}
		}
	}
}

func reportHeartbeat(ctx context.Context) error {
	req := types.HeartBeatReq{}
	// 获取node相关信息
	ni, err := collectNodeInfo()
	if err != nil {
		return fmt.Errorf("collect node info failed: %s", err)
	}
	req.Node = ni
	return nil
}

// collectNodeInfo 采集节点信息
// 目前先只采集IP信息
func collectNodeInfo() (types.NodeInfo, error) {
	ni := types.NodeInfo{}
	nodeIP, err := utils.GuessNodeIP()
	if err != nil {
		return ni, err
	}
	ni.IP = nodeIP
	return ni, nil
}
