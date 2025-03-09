package main

import (
	"fmt"
	"log/slog"
	"os"

	clusteragent "caas-ai-agent.ccsp.zte.com.cn/internal/cluster-agent"
	clusterhelper "caas-ai-agent.ccsp.zte.com.cn/pkg/cluster-helper"
	"github.com/spf13/cobra"
)

var (
	heartbeatServer string
	listenPort      int
	kubeconfig      string
	logLevel        string
)

func init() {
	rootCmd.Flags().StringVarP(&heartbeatServer, "heartbeat-server", "s", "", "Heartbeat server addr")
	rootCmd.Flags().IntVarP(&listenPort, "port", "p", 12673, "HTTP server port for listening")
	rootCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "c", "", "Path to kubeconfig file (default is ~/.kube/config)")
	rootCmd.Flags().StringVarP(&logLevel, "log-level", "l", "info", "Log level: debug, info, warn, error")

	// 设置日志级别
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

var rootCmd = &cobra.Command{
	Use:   "cluster-agent",
	Short: "Agent for k8s cluster",
	Long:  `This tool runs on the bastion host or control node of k8s, exposing a set of RESTful APIs to provide resource operations for the k8s cluster.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// 创建集群client
		ch, err := clusterhelper.New(kubeconfig)
		if err != nil {
			return fmt.Errorf("create cluster helper failed: %s", err)
		}

		// 上报心跳
		go clusteragent.StartHeartbeatLoop(ctx, heartbeatServer, listenPort, ch)
		// 启动server，提供k8s操作接口
		return clusteragent.StartServer(ctx, listenPort)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("start cluster agent failed:", err)
		os.Exit(1)
	}
}
