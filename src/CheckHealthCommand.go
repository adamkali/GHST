package src

import (
	"fmt"
	"os/exec"
)

func RunCheckHealth() {
	fmt.Println("Checking health...")
    tasks:= []GhostTask{
        {
            Name: "Checking surrealdb connection",
            State: "🤔",
        },
        {
            Name: "Checking tailwindcss installation",
            State: "🤔",
        },
        {
            Name: "Checking golang installation",
            State: "🤔",
        },
    }
    for _, task := range tasks {
        task.Progress()
    }

    go func() {
        cmd := exec.Command("go", "version")
        err := cmd.Run()
        if err != nil {
            tasks[2].Fail(err)
        } else {
            tasks[2].Complete()
        }
    }()
    go func() {
        cmd := exec.Command("tailwindcss", "help")
        err := cmd.Run()
        if err != nil {
            tasks[1].Fail(err)
        } else {
            tasks[1].Complete()
        }
    }()
    go func() {
        cmd := exec.Command("surreal", "version")
        err := cmd.Run()
        if err != nil {
            tasks[0].State = "🛑"
            tasks[0].Name = "Could not connect, if you are using docker, or a cloud servicte url it is possible to regard this message"
            tasks[0].Progress()
        } else {
            tasks[0].Complete()
        }
    }()

    // wait for all tasks to complete 
    for _, task := range tasks {
        for task.State == "🤔" {
            continue
        }
    }
}
