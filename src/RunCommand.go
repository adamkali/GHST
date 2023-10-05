package src

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func RunSuperThread(config string) {
    chanDone := make(chan bool);
    chanExit := make(chan error);
    chanOutput := make(chan string); 

    // spawn a waitgroup 
    wg := sync.WaitGroup{};
    wg.Add(1);

    // spawn a goroutine for each command 
    go runGoThread(chanDone, chanExit, chanOutput); 
    go runTailwindcssThread(config, chanDone, chanExit, chanOutput); 

    // spawn a goroutine to listen for output 
    go func() { 
        for {
            select {
            case output := <- chanOutput:
                fmt.Println(output);
            case <- chanDone:
                wg.Done();
                return;
            }
        }
    }()

    // spawn a goroutine to listen for errors
    go func() {
        for {
            select {
            case err := <- chanExit:
                ErrOut(err);
                wg.Done();
                return;
            }
        }
    }()

    // wait for all goroutines to finish 
    wg.Wait();
    return;
}

//func RunSuperThreadWithGhostConfigPath(path string) {
//
//}


func runTailwindcssThread(config string, chanDone chan bool, chanExit chan error, chanOutput chan string) {
    ghostConfig, err := LoadGhostConfig();
    if err != nil { 
        err = fmt.Errorf("ghost encountered an error loading the ghost.yaml file: %s", err.Error());
        chanExit <- err;
        return;
    }
    // check if the tailwindcss input file exists
    if _, err := os.Stat(ghostConfig.TailwindCSS.Input); os.IsNotExist(err) {
        err = fmt.Errorf("ghost encountered an error loading the tailwindcss input file: %s", err.Error());
        chanExit <- err;
        return;
    }
    // check if the tailwindcss output file exists 
    if _, err := os.Stat(ghostConfig.TailwindCSS.Output); os.IsNotExist(err) {
        err = fmt.Errorf("ghost encountered an error loading the tailwindcss output file: %s", err.Error());
        chanExit <- err;
        return;
    }
    // run the tailwindcss command 
    cmd := exec.Command(
        "tailwindcss",
        "-i",
        ghostConfig.TailwindCSS.Input,
        "-o",
        ghostConfig.TailwindCSS.Output,
        "--watch",
    );
    output, err := cmd.CombinedOutput();
    if err != nil { 
        err = fmt.Errorf("ghost encountered an error running the program: %s", err.Error());
        chanExit <- err;
        return;
    }
    go func() {
        for {
            select {
            case chanOutput <- string(output):
            case <- chanDone:
                return;
            }
        }
    }()
}


func runGoThread(chanDone chan bool, chanExit chan error, chanOutput chan string) {
    cmd := exec.Command("go", "run", "main.go");
    output, err := cmd.CombinedOutput();
    if err != nil { 
        err = fmt.Errorf("ghost encountered an error running the program: %s", err.Error());
        chanExit <- err;
        return;
    }
    go func() { 
        for {
            select {
            case chanOutput <- string(output):
            case <- chanDone:
                return;
            }
        }
    }()
}
