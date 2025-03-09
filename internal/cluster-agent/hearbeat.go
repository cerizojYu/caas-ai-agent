package clusteragent

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
	clusterhelper "caas-ai-agent.ccsp.zte.com.cn/pkg/cluster-helper"
	"caas-ai-agent.ccsp.zte.com.cn/pkg/utils"
	"resty.dev/v3"
)

// 上报心跳间隔
const heartbeatInterval = time.Minute

// 开启死循环，定期上报心跳
func StartHeartbeatLoop(ctx context.Context, heartbeatServer string, listenPort int, ch *clusterhelper.ClusterHelper) {
	if heartbeatServer == "" {
		slog.Info("heartbeat server is empty, skip heartbeat")
		return
	}

	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("receive context done, stop heartbeat")
			return
		case <-ticker.C:
			// 采集节点和集群信息
			ai, ci, err := collect(ch)
			if err != nil {
				slog.Warn("collect info failed", "error", err)
				continue
			}
			// 上报心跳
			req := types.HeartBeatReq{}
			req.Agent = ai
			req.Cluster = ci
			err = reportHeartbeat(ctx, heartbeatServer, req)
			if err != nil {
				slog.Warn("report heartbeat failed", "error", err)
			}
		}
	}
}

// collect 采集节点和集群信息
func collect(ch *clusterhelper.ClusterHelper) (types.AgentInfo, types.ClusterInfo, error) {
	ai, err := collectAgentInfo()
	if err != nil {
		return ai, types.ClusterInfo{}, err
	}
	ci, err := collectClusterInfo(ch)
	return ai, ci, err
}

// collectAgentInfo 采集节点信息
// 目前先只采集IP信息
func collectAgentInfo() (types.AgentInfo, error) {
	ai := types.AgentInfo{}
	nodeIP, err := utils.GuessNodeIP()
	if err != nil {
		return ai, err
	}
	ai.IP = nodeIP
	return ai, nil
}

// collectClusterInfo 采集集群信息
func collectClusterInfo(ch *clusterhelper.ClusterHelper) (types.ClusterInfo, error) {
	ci := types.ClusterInfo{}
	sv, err := ch.ClusterServerVersion()
	if err != nil {
		return ci, err
	}
	ci.ServerVersion = sv.String()
	return ci, nil
}

// reportHeartbeat 上报心跳
func reportHeartbeat(ctx context.Context, heartbeatServer string, req types.HeartBeatReq) error {
	resp, err := resty.New().R().SetBody(req).SetContext(ctx).Post(heartbeatServer)
	if err != nil {
		return fmt.Errorf("send heartbeat to server failed", "server", heartbeatServer, "req", req, "resp", resp, "error", err)
	}
	slog.Debug("report heartbeat success", "server", heartbeatServer, "req", req, "resp", resp)
	return nil
}
