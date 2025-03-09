package clusterhelper

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (ch *ClusterHelper) ListEvents(ctx context.Context, namespace string, filter string) ([]string, error) {
	events, err := ch.clientset.CoreV1().Events(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var eventList []string
	for _, event := range events.Items {
		eventList = append(eventList, event.Message)
	}
	return eventList, nil
}
