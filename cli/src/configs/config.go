package configs

import (
	"github.com/beto0607/goconfiglib"
)

type TodoIntegration struct {
	Enabled   bool
	UseApiKey bool
	ApiKey    string
}

type TodoTuiConfig struct {
	GoogleKeep  TodoIntegration
	GoogleTasks TodoIntegration
}

func LoadConfig() (*TodoTuiConfig, error) {
	configs, err := goconfiglib.LoadConfigs("todo-tui/todo-tui.ini", goconfiglib.Settings{UseXDGConfigHome: true})
	if err != nil {
		return nil, err
	}
	appConfig := TodoTuiConfig{
		GoogleKeep:  TodoIntegration{Enabled: false},
		GoogleTasks: TodoIntegration{Enabled: false},
	}

	for k := range configs.Root.Subsections {
		subSection := configs.Root.Subsections[k]
		if subSection.Name == "google-keep" {
			appConfig.GoogleKeep.Enabled = subSection.Values["enabled"] != "false"
			appConfig.GoogleKeep.ApiKey = subSection.Values["apiKey"]
			appConfig.GoogleKeep.UseApiKey = subSection.Values["useApiKey"] == "true"
		} else if subSection.Name == "google-tasks" {
			appConfig.GoogleTasks.Enabled = subSection.Values["enabled"] != "false"
			appConfig.GoogleTasks.ApiKey = subSection.Values["apiKey"]
			appConfig.GoogleTasks.UseApiKey = subSection.Values["useApiKey"] == "true"
		}
	}

	return &appConfig, nil
}
