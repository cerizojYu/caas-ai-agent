package clustermanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"caas-ai-agent.ccsp.zte.com.cn/internal/consts"
	"caas-ai-agent.ccsp.zte.com.cn/internal/types"
	"resty.dev/v3"
)

type ClusterAgentClient struct {
	Address string
}

func NewClusterAgentClient(addr string) *ClusterAgentClient {
	return &ClusterAgentClient{Address: addr}
}

func (c *ClusterAgentClient) EventList(namespace, filter string) ([]string, error) {
	url := fmt.Sprintf("%s/%s?namespace=%s&filter=%s", c.Address, consts.ClusterAgentEventsURI, namespace, filter)
	resp, err := resty.New().R().Get(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("list events failed: %s", err)
	}
	eventsResp := types.EventListResp{}
	err = json.Unmarshal(resp.Bytes(), &eventsResp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal events response failed: %s", err)
	}
	return eventsResp.Events, nil
}
