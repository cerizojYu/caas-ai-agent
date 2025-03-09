package types

type EventListResp struct {
	Namespace string   `json:"namespace"`
	Filter    string   `json:"filter"`
	Events    []string `json:"events"`
}
