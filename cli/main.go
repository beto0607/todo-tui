package main

import (
	"fmt"
	"todo-tui/src/configs"
	"todo-tui/src/integrations"
)

func main() {
	appConfigs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GKeep enabled %t\n", appConfigs.GoogleKeep.Enabled)
	fmt.Printf("GTask enabled %t\n", appConfigs.GoogleTasks.Enabled)

	if appConfigs.GoogleTasks.Enabled {
		tasksManager, err := integrations.NewGoogleTasksManager(appConfigs)
		if err != nil {
			panic(err)
		}

		tasksManager.LoadTasks()
	}
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }
}
