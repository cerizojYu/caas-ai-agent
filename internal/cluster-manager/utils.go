package clustermanager

import "fmt"

func discoveryAgentAddr(agentIP string) (string, error) {
	ni, ok := getNodeInfo(agentIP)
	if !ok {
		return "", fmt.Errorf("agent %s not found", agentIP)
	}
	return fmt.Sprintf("%s:%d", ni.Agent.IP, ni.Agent.Port), nil
}
