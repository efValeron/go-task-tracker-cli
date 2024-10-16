package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"example.com/task-tracker-cli/constants"
	"example.com/task-tracker-cli/types"
	"example.com/task-tracker-cli/utils"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func Update(id uuid.UUID, newDescription string) (error bool) {
	utils.DebugLog(color.RedString("Updating task: ") + "\"%v\"\n")

	var tasksJSON types.TasksArrayFormat

	// Check if file *not* exists
	if _, err := os.Stat(constants.FILE_NAME); errors.Is(err, os.ErrNotExist) {
		fmt.Println(color.RedString("Tasks file tasks.json does not exist!"))
		return true
	} else {
		// Open file
		file, fileOpenError := os.ReadFile(constants.FILE_NAME)
		utils.CheckErr(fileOpenError)

		// Parse file
		jsonParseError := json.Unmarshal([]byte(string(file)), &tasksJSON)
		utils.CheckErr(jsonParseError)
	}

	// Find and update task
	for i, task := range tasksJSON.Tasks {
		if task.ID == id {
			tasksJSON.Tasks[i].Descripton = newDescription
			tasksJSON.Tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}

	utils.DebugLog(color.BlueString("Struct with updated task: ")+"%+v\n", tasksJSON.Tasks)

	// Convert changes
	resultJSON, jsonConvertingError := json.Marshal(tasksJSON)
	utils.CheckErr(jsonConvertingError)
	utils.DebugLog(color.BlueString("Converted JSON: ")+"%v\n", string(resultJSON))

	// Write changes
	writingError := os.WriteFile(constants.FILE_NAME, []byte(resultJSON), 0777)
	utils.CheckErr(writingError)

	fmt.Println(color.GreenString("Task updated successfully!"))
	return
}
