package types

import "net"

type AgentInfo struct {
	IP   net.IP `json:"ip"`
	Port int    `json:"port"`
}

type ClusterInfo struct {
	ServerVersion string `json:"server_version"`
}

type HeartBeatReq struct {
	Agent   AgentInfo   `json:"agent"`
	Cluster ClusterInfo `json:"cluster"`
}

type NodeInfo = HeartBeatReq
