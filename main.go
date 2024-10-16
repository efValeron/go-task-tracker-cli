package main

import (
	"fmt"
	"os"

	"example.com/task-tracker-cli/functions"
	"example.com/task-tracker-cli/types"
	"example.com/task-tracker-cli/utils"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(color.RedString("Oops! An error occured! ") + "Please try again. Add " + color.CyanString("debug") + " and try again to see more info about the error.")
			utils.DebugLog(color.BlueString("Error message: ")+"%v\n", err)
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println(color.RedString("No command provided! ") + "Type " + color.CyanString("help") + " to see available commands.")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		if len(args) == 0 {
			fmt.Println(color.RedString("No description provided!"))
			return
		}

		description := args[0]
		error := functions.Add(description)

		if error {
			fmt.Println("An error occured! Please try again.")
			return
		}
	case "list":
		var filter types.StatusNames

		if len(args) > 0 {
			filter = types.StatusNames(args[0])
		}

		error := functions.List(filter)

		if error {
			fmt.Println("An error occured! Please try again.")
			return
		}
	case "delete":
		id, err := uuid.Parse(args[0])

		if err != nil {
			fmt.Println(color.RedString("Invalid UUID!"))
			return
		}

		error := functions.Delete(id)

		if error {
			fmt.Println("An error occured! Please try again.")
			return
		}
	case "update":
		if len(args) < 2 {
			fmt.Println(color.RedString("Not all arguments provided! ") + "Provide: " + color.CyanString("<id> <description>"))
			return
		}
		id, err := uuid.Parse(args[0])

		if err != nil {
			fmt.Println(color.RedString("Invalid UUID!"))
			return
		}

		error := functions.Update(id, args[1])

		if error {
			fmt.Println("An error occured! Please try again.")
			return
		}
	case "change-status":
		if len(args) < 2 {
			fmt.Println(color.RedString("Not all arguments provided! ") + "Provide: " + color.CyanString("<id> <status>"))
			return
		}

		id, err := uuid.Parse(args[0])

		if err != nil {
			fmt.Println(color.RedString("Invalid UUID!"))
			return
		}

		status := types.StatusNames(args[1])

		error := functions.ChangeStatus(id, status)

		if error {
			fmt.Println("An error occured! Please try again.")
			return
		}
	case "help":
		fmt.Println()
		fmt.Println(color.MagentaString("Task Tracker CLI"))
		fmt.Println("A simple task tracker CLI written in Go.")
		fmt.Println()
		fmt.Println(color.YellowString("If you modify the ") + color.GreenString("tasks.json") + color.YellowString(" file directly, the program may behave unexpectedly!"))
		fmt.Println(color.YellowString("Make sure the command window is wide enough to display the tasks list properly."))
		fmt.Println()
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println()
		fmt.Println(" Task Management:")
		fmt.Println(color.CyanString("  add <description>") + " - Add a new task.")
		fmt.Println(color.CyanString("  list") + " - List all tasks.")
		fmt.Println(color.CyanString("  list <todo | in-progress | done>") + " - List tasks by status.")
		fmt.Println(color.CyanString("  delete <index>") + " - Delete a task.")
		fmt.Println(color.CyanString("  update <index> <description>") + " - Update a task description.")
		fmt.Println(color.CyanString("  change-status <index> <todo | in-progress | done>") + " - Change the status of a task.")
		fmt.Println()
		fmt.Println(" Debugging:")
		fmt.Println(color.CyanString("  add \"Go is great!\" debug") + " - Add debug option for additional logging.")
		fmt.Println()
		fmt.Println(" General:")
		fmt.Println(color.CyanString("  help") + " - Show this help message.")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println(color.CyanString(" add \"Buy groceries\"") + " - Adds a task 'Buy groceries.'")
		fmt.Println(color.CyanString(" list in-progress") + " - Lists all tasks that are in progress.")
		fmt.Println(color.CyanString(" delete 1") + " - Deletes the task with index 1.")
		fmt.Println(color.CyanString(" update 2 \"Finish report\"") + " - Updates task 2 with the description 'Finish report.'")
		fmt.Println(color.CyanString(" change-status 2 done") + " - Marks task 2 as done.")
		fmt.Println()
		fmt.Println()
		fmt.Println(color.MagentaString("Version: ") + color.YellowString("1.0.0"))
		fmt.Println(color.MagentaString("Author: ") + color.YellowString("efvaleron"))
		fmt.Println(color.MagentaString("GitHub: ") + color.YellowString("https://github.com/efValeron/go-task-tracker-cli"))
		fmt.Println(color.MagentaString("Inspiration: ") + color.YellowString("https://roadmap.sh/projects/task-tracker"))
		fmt.Println()
	default:
		fmt.Println(color.RedString("Command not found! ") + "Type " + color.CyanString("help") + " to see available commands.")
		return
	}
}
