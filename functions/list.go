package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"errors"

	"example.com/task-tracker-cli/constants"
	"example.com/task-tracker-cli/types"
	"example.com/task-tracker-cli/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func List(filter types.StatusNames) (error bool) {
	// Check if file *not* exists
	if _, err := os.Stat(constants.FILE_NAME); errors.Is(err, os.ErrNotExist) {
		fmt.Println(color.RedString("Task file does not exist!"))
		return true
	}

	// Check if filter is valid
	filter = types.StatusNames(strings.ToUpper(string(filter)))

	switch filter {
	case types.TODO, types.IN_PROGRESS, types.DONE, "DEBUG", "":
		if filter != "DEBUG" && filter != "" {
			utils.DebugLog(color.BlueString("Searching tasks with filter: ")+color.CyanString("%v\n"), filter)
		}
	default:
		fmt.Println(color.RedString("Invalid filter!"))
		return
	}

	// Read file
	file, fileOpenError := os.ReadFile(constants.FILE_NAME)
	utils.CheckErr(fileOpenError)

	var tasksJSON types.TasksArrayFormat
	jsonParseError := json.Unmarshal([]byte(string(file)), &tasksJSON)
	utils.CheckErr(jsonParseError)

	utils.DebugLog(color.BlueString("Parsed JSON: ")+"%v\n", tasksJSON)

	// Print tasks
	rows := make([]table.Row, len(tasksJSON.Tasks))
	for i, task := range tasksJSON.Tasks {
		if filter != "" && task.Status != filter {
			continue
		}

		status := string(task.Status)
		switch status {
		case string(types.TODO):
			status = color.BlueString(status)
		case string(types.IN_PROGRESS):
			status = color.YellowString(status)
		case string(types.DONE):
			status = color.GreenString(status)
		}

		rows[i] = table.Row{task.Descripton, status, task.CreatedAt, task.UpdatedAt, task.ID}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Description", "Status", "Created At", "Updated At", "ID"})
	t.AppendRows(rows)
	t.Render()

	return
}
