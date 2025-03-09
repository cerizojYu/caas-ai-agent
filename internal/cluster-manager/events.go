package clustermanager

import "fmt"

// listEvents lists events in the cluster
func ListEvents(agentIP string, namespace string, filter string) ([]string, error) {
	agentAddr, err := discoveryAgentAddr(agentIP)
	if err != nil {
		return nil, fmt.Errorf("discovery agent address failed: %s", err)
	}
	agentCli := NewClusterAgentClient(agentAddr)
	events, err := agentCli.EventList(namespace, filter)
	if err != nil {
		return nil, fmt.Errorf("list events failed: %s", err)
	}
	return events, nil
}
