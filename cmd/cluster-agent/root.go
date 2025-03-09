package main

import (
	"fmt"
	"os"

	clusteragent "caas-ai-agent.ccsp.zte.com.cn/internal/cluster-agent"
	"github.com/spf13/cobra"
)

var (
	serverAddr string
	addr       string
	kubeconfig string
)

func init() {
	rootCmd.Flags().StringVarP(&addr, "addr", "a", ":12673", "HTTP server addr")
	rootCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", "", "Path to kubeconfig file (default is ~/.kube/config)")
}

var rootCmd = &cobra.Command{
	Use:   "cluster-agent",
	Short: "Agent for k8s cluster",
	Long:  `This tool runs on the bastion host or control node of k8s, exposing a set of RESTful APIs to provide resource operations for the k8s cluster.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		// 上报心跳

		// 启动server，提供k8s操作接口
		return clusteragent.StartServer(ctx, addr)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("start cluster agent error:", err)
		os.Exit(1)
	}
}
