package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/task-tracker-cli/constants"
	"example.com/task-tracker-cli/types"
	"example.com/task-tracker-cli/utils"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (error bool) {
	utils.DebugLog(color.RedString("Deleting task: ") + "\"%v\"\n")

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

	// Find and delete task
	lenBeforeDeleting := len(tasksJSON.Tasks)

	for i, task := range tasksJSON.Tasks {
		if task.ID == id {
			tasksJSON.Tasks = append(tasksJSON.Tasks[:i], tasksJSON.Tasks[i+1:]...)
			break
		}
	}

	if len(tasksJSON.Tasks) == lenBeforeDeleting {
		fmt.Printf(color.RedString("Task with id ")+color.CyanString("%v")+color.RedString(" not found!\n"), id)
		return
	}

	utils.DebugLog(color.BlueString("Struct without deleted task: ")+"%+v\n", tasksJSON.Tasks)

	// Convert changes
	resultJSON, jsonConvertingError := json.Marshal(tasksJSON)
	utils.CheckErr(jsonConvertingError)
	utils.DebugLog(color.BlueString("Converted JSON: ")+"%v\n", string(resultJSON))

	// Write changes
	writingError := os.WriteFile(constants.FILE_NAME, []byte(resultJSON), 0777)
	utils.CheckErr(writingError)

	fmt.Println(color.GreenString("Task deleted successfully!"))
	return
}
