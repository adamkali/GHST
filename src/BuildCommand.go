package src

import (
	"fmt"
	"os"
	"os/exec"
)

func minifyCSS(
    config string,
    outputChannel chan string,
    errorChannel chan error,
    doneChannel chan bool,
) {
	// load the ghost config
	ghostConfig, err := Ghost(config)
	if err != nil {
		errorChannel <- err
        doneChannel <- true
		return
	}

	// minify the css
	cmd := exec.Command(
        "tailwindcss",
        "-i",
        ghostConfig.TailwindCSS.Input,
        "-o",
        ghostConfig.TailwindCSS.Output,
        "--watch",
    )
    output, err := cmd.CombinedOutput()
    if err != nil {
        errorChannel <- err
        doneChannel <- true
        return
    }
    outputChannel <- string(output)
}

func runGoBuild(
    config string,
    outputChannel chan string,
    errorChannel chan error,
    doneChannel chan bool,
) {
    ghostConfig, err := Ghost(config) 
    if err != nil {
        errorChannel <- err
        doneChannel <- true
        return
    }

    cmd := exec.Command("go", "build", "-o", ghostConfig.Name)
    output, err := cmd.CombinedOutput()
    if err != nil {
        errorChannel <- err
        doneChannel <- true
        return
    }
    outputChannel <- string(output)
    doneChannel <- true
}

func BuildSuperThread(config string) {
    // create the channels
    outputChannel := make(chan string)
    errorChannel := make(chan error)
    doneChannel := make(chan bool)

    // run the threads
    go minifyCSS(config, outputChannel, errorChannel, doneChannel)
    go runGoBuild(config, outputChannel, errorChannel, doneChannel)

    // wait for the threads to finish
    // using the done channel 
    for { 
        select {
        case output := <-outputChannel:
            fmt.Println(output)
        case err := <-errorChannel:
            ErrOut(err)
        case <-doneChannel:
            fmt.Println("Build finished.")
            os.Exit(0)
        }
    }
}



