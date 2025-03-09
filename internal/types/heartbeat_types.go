package types

import "net"

type NodeInfo struct {
	IP net.IP `json:"ip"`
}

type ClusterInfo struct {
}

type HeartBeatReq struct {
	Node    NodeInfo    `json:"node"`
	Cluster ClusterInfo `json:"cluster"`
}
