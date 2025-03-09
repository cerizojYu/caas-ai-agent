package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"caas-ai-agent.ccsp.zte.com.cn/internal/ai"
	clustermanager "caas-ai-agent.ccsp.zte.com.cn/internal/cluster-manager"
	"caas-ai-agent.ccsp.zte.com.cn/internal/prompt"
	"caas-ai-agent.ccsp.zte.com.cn/internal/tools"
	"github.com/spf13/cobra"
)

var (
	listenPort int
)

var rootCmd = &cobra.Command{
	Use:   "manager",
	Short: "manager for caas ai agent",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// 启动cluster manager
		go clustermanager.StartServer(ctx, listenPort)
		time.Sleep(time.Minute)

		query := ""

		eventsListTool := tools.EventsListToolName + ":" + tools.EventsListToolDescription + "\nparam: \n" + tools.EventsListToolParam
		toolsL := make([]string, 0)
		toolsL = append(toolsL, eventsListTool)

		tool_names := make([]string, 0)
		tool_names = append(tool_names, tools.EventsListToolName)

		prompt := fmt.Sprintf(prompt.Template, toolsL, tool_names, query)
		fmt.Println("prompt: ", prompt)
		i := 1
		for {
			first_response := ai.NormalChat(ai.MessageStore.ToMessage())
			fmt.Printf("========第%d轮回答========\n", i)
			fmt.Println(first_response)
			regexPattern := regexp.MustCompile(`Final Answer:\s*(.*)`)
			finalAnswer := regexPattern.FindStringSubmatch(first_response.Content)
			if len(finalAnswer) > 1 {
				fmt.Println("========最终 GPT 回复========")
				fmt.Println(first_response.Content)
				break
			}

			ai.MessageStore.AddForAssistant(first_response)

			regexAction := regexp.MustCompile(`Action:\s*(.*?)[.\n]`)
			regexActionInput := regexp.MustCompile(`Action Input:\s*(.*?)[.\n]`)

			action := regexAction.FindStringSubmatch(first_response.Content)
			actionInput := regexActionInput.FindStringSubmatch(first_response.Content)
			if len(action) > 1 && len(actionInput) > 1 {
				i++
				result := ""
				//需要调用工具
				if action[1] == tools.EventsListToolName {
					fmt.Println("calls tools.EventsListToolNam")
					result = tools.EventsListTool(actionInput[1], actionInput[2])
				}
				fmt.Println("========函数返回结果========")
				fmt.Println(result)

				Observation := "Observation: " + result
				prompt = first_response.Content + Observation
				fmt.Printf("========第%d轮的prompt========\n", i)
				fmt.Println(prompt)
				ai.MessageStore.AddForUser(prompt)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().IntVarP(&listenPort, "port", "p", 12672, "HTTP server port for listening")
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("start caas ai manager failed:", err)
		os.Exit(1)
	}
}
