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

func Add(description string) (error bool) {
	utils.DebugLog(color.BlueString("Adding task: ")+"\"%v\"\n", description)

	var tasksJSON types.TasksArrayFormat

	// Check if file *not* exists
	if _, err := os.Stat(constants.FILE_NAME); errors.Is(err, os.ErrNotExist) {
		utils.DebugLog(color.YellowString("File %v not found! ")+color.BlueString("Creating new one...\n"), constants.FILE_NAME)

		file, fileCreateError := os.Create(constants.FILE_NAME)
		utils.CheckErr(fileCreateError)

		file.Close()
	} else {
		// Open file
		file, fileOpenError := os.ReadFile(constants.FILE_NAME)
		utils.CheckErr(fileOpenError)

		// Parse file
		jsonParseError := json.Unmarshal([]byte(string(file)), &tasksJSON)
		utils.CheckErr(jsonParseError)
	}

	// Create new task
	id := uuid.New()

	tasksJSON.Tasks = append([]types.TaskJSONFormat{{
		Descripton: description,
		ID:         id,
		Status:     types.TODO,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}}, tasksJSON.Tasks...)

	utils.DebugLog(color.BlueString("Struct with created task: ")+"%+v\n", tasksJSON.Tasks)

	// Convert changes
	resultJSON, jsonConvertingError := json.Marshal("adddad")
	utils.CheckErr(jsonConvertingError)
	utils.DebugLog(color.BlueString("Converted JSON: ")+"%v\n", string(resultJSON))

	// Write changes
	writingError := os.WriteFile(constants.FILE_NAME, []byte(resultJSON), 0777)
	utils.CheckErr(writingError)

	fmt.Printf(color.GreenString("Task added successfully! ")+"ID: %v\n", id)
	return
}
