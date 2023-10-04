package main;

import (
    "flag"
    "fmt"
    "os"

    ghst "github.com/adamkali/ghost/src"
);

const (
    VERSION = "0.0.1"
    GHSTASCII = `
        ,.,
     /"    "\
    /  O   O \
 \""          ""/    
  '"_          /    
    |        ./
  ,"______,-"`
);

func main() {
    var noconfirm bool;

    var projectName string;
    var projectPath string;

    var modelName string;
    var modelPathInput string;


    flag.BoolVar(&noconfirm, "C", false, "no confirmation");
    flag.BoolVar(&noconfirm, "noconfirm", false, "no confirmation");

    // make a new flag set 
    newFlagSet := flag.NewFlagSet("new", flag.ExitOnError);    
    modelFlagSet := flag.NewFlagSet("model", flag.ExitOnError);

    // new subcommand
    newFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION);
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m");
        fmt.Println("\nCreate a new project using ghost");
        fmt.Println("\nUsage: ghost new <project name> [options]");
        newFlagSet.PrintDefaults();
    }

    //newFlagSet.StringVar(&projectName, "n", "", "project name");  
    //newFlagSet.StringVar(&projectName, "name", "", "project name");
    newFlagSet.StringVar(&projectPath, "p", "", "project path");
    newFlagSet.StringVar(&projectPath, "path", "", "project path");

    // model subcommand
    modelFlagSet.Usage = func() {
        fmt.Println("ghost: " + VERSION);
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m");
        fmt.Println("\nCreate a new model using ghost");
        fmt.Println("\nUsage: ghost model <model name> [options]");
        modelFlagSet.PrintDefaults();
    }

    flag.Usage = func() {
        fmt.Println("ghost: " + VERSION);
        fmt.Println("\033[0;34m" + GHSTASCII + "\033[0m");
        fmt.Println("\nGhost is a simple CLI tool for creating GHST projects and managing them.");
        fmt.Println("Usage: ghost <command> [options]");
        fmt.Println("Commands:");
        fmt.Println("\tnew\t\tCreate a new project");
        fmt.Println("\tmodel\t\tCreate a new model");
        fmt.Println("Options:");
        flag.PrintDefaults();
    }

    modelFlagSet.StringVar(&modelName, "n", "", "model name");
    modelFlagSet.StringVar(&modelName, "name", "", "model name");
    modelFlagSet.StringVar(&modelPathInput, "p", "", "model path");
    modelFlagSet.StringVar(&modelPathInput, "path", "", "model path");
    
    if len(os.Args) < 2 {
        flag.Usage();
        os.Exit(1);
    }


    switch os.Args[1] {
    case "new": 
        // the command will be of the form
        // ghost new <project name> -p <project path>
        if len(os.Args) < 3 {
            newFlagSet.Usage();
            os.Exit(1);
        }
        projectName = os.Args[2];
        newFlagSet.Parse(os.Args[3:]);
        if projectName == "" {
            // in red 
            fmt.Println("\033[0;31mProject name is required\033[0m");
            newFlagSet.Usage();
            os.Exit(1);
        }
        fmt.Println("Creating a new GHST project: " + projectName);
        if projectPath != "" {
            fmt.Println("Project path: " + projectPath);
            ghst.CreateProject(projectName, projectPath, noconfirm);
        } else {
            fmt.Println("Project path used current directory");
            ghst.CreateProjectCurrentDir(projectName, noconfirm);
        }

    case "model":
        modelFlagSet.Parse(os.Args[2:]);
        if modelName == "" {
            fmt.Println("\033[0;31mModel name is required\033[0m");
            modelFlagSet.Usage();
            os.Exit(1);
        }
    }
}
