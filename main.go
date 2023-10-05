package main;

import (
    "flag"
    "fmt"
    "os"

    ghst "github.com/adamkali/ghost/src"
);

const (
    VERSION = "0.1.0"
    GHSTASCII = `
       _,.,
     /'    '\
    / ()  () \
 \""          ""/    
  '-_    o     /    
    |        ./
  ,;______,-'` 
)

func main() {
    var noconfirm bool
    var config string

    var projectName string
    var projectPath string
    var modelName string
    var modelPathInput string
    var help bool


    // make a new flag set 
    newFlagSet := flag.NewFlagSet("new", flag.ExitOnError)
    modelFlagSet := flag.NewFlagSet("model", flag.ExitOnError)
    twInitFlagSet := flag.NewFlagSet("tw-init", flag.ExitOnError)

    // new subcommand
    newFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION)
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nCreate a new project using ghost")
        fmt.Println("\nUsage: ghost new <project name> [options]")
        newFlagSet.PrintDefaults()
    }

    //newFlagSet.StringVar(&projectName, "n", "", "project name");  
    //newFlagSet.StringVar(&projectName, "name", "", "project name");
    newFlagSet.BoolVar(&help, "h", false, "help")
    newFlagSet.BoolVar(&help, "help", false, "help")
    newFlagSet.StringVar(&projectPath, "p", "", "project path")
    newFlagSet.StringVar(&projectPath, "path", "", "project path")
    newFlagSet.BoolVar(&noconfirm, "y", false, "no confirmation")
    newFlagSet.BoolVar(&noconfirm, "noconfirm", false, "no confirmation")

    // model subcommand
    modelFlagSet.StringVar(&modelName, "n", "", "model name")
    modelFlagSet.StringVar(&modelName, "name", "", "model name")
    modelFlagSet.StringVar(&modelPathInput, "p", "", "model path")
    modelFlagSet.StringVar(&modelPathInput, "path", "", "model path")
    modelFlagSet.BoolVar(&noconfirm, "y", false, "no confirmation")
    modelFlagSet.BoolVar(&noconfirm, "noconfirm", false, "no confirmation")
    modelFlagSet.BoolVar(&help, "h", false, "help")
    modelFlagSet.BoolVar(&help, "help", false, "help")
    modelFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION)
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nCreate a new model using ghost")
        fmt.Println("\nUsage: ghost model <model name> [options]")
        modelFlagSet.PrintDefaults()
    }

    // tw-init subcommand 
    twInitFlagSet.BoolVar(&help, "h", false, "help")
    twInitFlagSet.BoolVar(&help, "help", false, "help")
    twInitFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION)
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nInitialize tailwindcss using ghost")
        fmt.Println("Uses tailwindcss standalone executable")
        fmt.Println("run `ghost checkhealth` to check if tailwindcss is installed")
        fmt.Println("\nUsage: ghost tw-init [options]")
        twInitFlagSet.PrintDefaults()
    }

    // run subcommand 
    runFlagSet := flag.NewFlagSet("run", flag.ExitOnError)
    runFlagSet.StringVar(&config, "c", "", "config file")
    runFlagSet.StringVar(&config, "config", "", "config file")
    runFlagSet.BoolVar(&help, "h", false, "help")
    runFlagSet.BoolVar(&help, "help", false, "help")
    runFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION);
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nRun the current project using ghost")
        fmt.Println("\nUsage: ghost run [options]")
        runFlagSet.PrintDefaults()
    }

    // build subcommand 
    buildFlagSet := flag.NewFlagSet("build", flag.ExitOnError)
    buildFlagSet.StringVar(&config, "c", "", "config file")
    buildFlagSet.StringVar(&config, "config", "", "config file")
    buildFlagSet.BoolVar(&help, "h", false, "help")
    buildFlagSet.BoolVar(&help, "help", false, "help")
    buildFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION)
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nBuild the current project using ghost")
        fmt.Println("\nUsage: ghost build [options]")
        buildFlagSet.PrintDefaults()
    }

    flag.Usage = func() {
        fmt.Println("ghost: " + VERSION)
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m")
        fmt.Println("\nGhost is a simple CLI tool for creating GHST projects and managing them.")
        fmt.Println("Usage: ghost <command> [options]")
        fmt.Println("Commands:")
        fmt.Println("\tnew\t\tCreate a new project")
        fmt.Println("\tmodel\t\tCreate a new model")
        fmt.Println("Options:")
        flag.PrintDefaults()
    }

    
    if len(os.Args) < 2 {
        flag.Usage()
        os.Exit(1)
    }


    switch os.Args[1] {
    case "new": 
        if len(os.Args) < 3 {
            newFlagSet.Usage()
            os.Exit(1)
        }
        projectName = os.Args[2]
        newFlagSet.Parse(os.Args[3:])
        if help {
            newFlagSet.Usage()
            os.Exit(0)
        }
        if projectName == "" {
            // in red 
            fmt.Println("\033[0;31mProject name is required\033[0m")
            newFlagSet.Usage()
            os.Exit(1)
        }
        fmt.Println("Creating a new GHST project: " + projectName)
        if projectPath != "" {
            fmt.Println("Project path: " + projectPath)
            ghst.CreateProject(projectName, projectPath, noconfirm)
        } else {
            fmt.Println("Project path used current directory")
            ghst.CreateProjectCurrentDir(projectName, noconfirm)
        }
        break;
    case "model":
        modelFlagSet.Parse(os.Args[2:]);
        if help {
            modelFlagSet.Usage()
            os.Exit(0)
        }
        if modelName == "" {
            fmt.Println("\033[0;31mModel name is required\033[0m")
            modelFlagSet.Usage()
            os.Exit(1)
        }
        panic("Not implemented yet")
    case "run":
        runFlagSet.Parse(os.Args[2:])
        if help {
            runFlagSet.Usage()
            os.Exit(0)
        }
        if config == "" {
            ghst.RunSuperThread("")
        } else {
            ghst.RunSuperThread(config)
        }
        break;
    case "build": 
        if help { 
            buildFlagSet.Usage()
            os.Exit(0)
        }
        buildFlagSet.Parse(os.Args[2:])
        if config == "" {
            ghst.BuildSuperThread("")
        } else {
            ghst.BuildSuperThread(config)
        }
        break;
    case "tw-init":
        ghst.RunTWInit()
        break;
    case "version":
        fmt.Println("ghost: " + VERSION)
        break;
    case "help":
        flag.Usage()
        break;
    case "checkhealth":
        ghst.RunCheckHealth()
        break;
    default:
        flag.Usage();
        os.Exit(1);
    }
    os.Exit(0);
}
