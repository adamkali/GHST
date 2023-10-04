package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

const (
    HTMXNAME = "static/scripts/htmx.org.js"
    TAILWINDCSSNAME = "static/styles/output.css"
    INPUTCSSNAME = "input.css"
    TAILWINDCSSCDNNAME = "<script src='https://cdn.tailwindcss.com'></script>"
    STEPTOTAL = 7
)

var (
    step = 0
    Project = Setup{}
)

type Setup struct {
    ProjectName string 
    Css string
    Htmx string
    TailwindCSS bool
    SurrealDB bool
}


func ShowStep(stepMessage string) {
    step += 1
    fmt.Println("( " + strconv.Itoa(step) + "/" + strconv.Itoa(STEPTOTAL) +  ")\t" + stepMessage)
}

func ErrOut(err error) {
    fmt.Println("🛑\t\t\033[0;31m" + err.Error() + "\033[0m")
    os.Exit(1)
}

func CreateProject(projectName string, projectPath string, noconfirm bool) {
    var err error 
    var signal string
    Project.ProjectName = projectName

    if _, err = os.Stat(projectPath); os.IsNotExist(err) {
        ErrOut(err)
    } else {
        os.Chdir(projectPath)
    }
    ShowStep("Created project: " + projectName + " at " + projectPath)

    signal, err = bootstrapGoProject(projectPath, noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    bootsrapStaticFolder()
    ShowStep("Created the /static folder 📁")

    signal, err = getHTMX()
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = getTailwindCSS(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ShowStep(signal)
    }
    
    signal, err = bootstrapSurrealDB(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootsrapMainFile(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootstrapViewsFolder(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootstrapRoutesAndModelsFolders(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    fmt.Println("🎉\t\t\033[0;32mProject created successfully\033[0m")
}

func CreateProjectCurrentDir(projectName string, noconfirm bool) {
    fmt.Println("Creating project: " + projectName + " at current directory")

    var err error 
    var signal string 
    Project.ProjectName = projectName 

    ShowStep("Created project: ./" + projectName)

    signal, err = bootstrapGoProject(".", noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    bootsrapStaticFolder()
    ShowStep("Created the /static folder 📁")

    signal, err = getHTMX()
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = getTailwindCSS(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ShowStep(signal)
    }
    
    signal, err = bootstrapSurrealDB(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootsrapMainFile(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootstrapViewsFolder(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    signal, err = bootstrapRoutesAndModelsFolders(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ShowStep(signal)
    }

    fmt.Println("🎉\t\t\033[0;32mProject created successfully\033[0m")
}

func bootstrapGoProject(projectName string, noconfirm bool) (string, error) {
    var err error
    
    fmt.Println("Creating a new Golang project")
    err = os.Mkdir(projectName, 0755)
    if err != nil {
        err = fmt.Errorf("Error creating project directory\nInternal Error -> " + err.Error())
        return "", err 
    }
    err = os.Chdir(projectName)
    if err != nil { 
        err = fmt.Errorf("Error changing directory to project directory\nInternal Error -> " + err.Error())
        return "", err 
    }
    err = os.Mkdir("src", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating src directory\nInternal Error -> " + err.Error())
        return "", err 
    }

    cmd := exec.Command("go", "version")
    _, err = cmd.CombinedOutput()
    if err != nil {
        err = fmt.Errorf("\033[0;31mGolang is not installed\033[0m\nInternal Error ->" + err.Error())
        return "", err
    }
    cmd = exec.Command("go", "mod", "init", projectName)
    output, err := cmd.CombinedOutput()
    fmt.Println(string(output))
    if err != nil {
        err = fmt.Errorf("\033[0;31mError creating go.mod\033[0m\nInternal Error -> " + err.Error())
        return "", err
    }

    cmd = exec.Command("go", "get", "-u", "github.com/gin-gonic/gin")
    fmt.Println("Installing \033[0;36mGin\033[0m of the GHST framework")
    fmt.Println("Continue? (y/n)")
    var input string 
    if noconfirm {
        input = "y"
    } else {
        fmt.Scanln(&input)
    }
    if input != "y" && input != "Y" && input != "" {
        fmt.Println("Aborting")
        fmt.Println("Removing project directory")
        os.Chdir("..") 
        os.RemoveAll(projectName) 
        os.Exit(0)
    }    

    output, err = cmd.CombinedOutput()
    fmt.Println(string(output)) 
    if err != nil { 
        err = fmt.Errorf("\033[0;31mError installing Gin\033[0m\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mGin 🍸 installed\033[0m", nil
}

func bootsrapStaticFolder() (string, error) {
    // crate a static folder 
    fmt.Println("Creating a static folder 📁")

    err := os.Mkdir("static", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating static folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("static/styles", 0755)
    if err != nil {
        err = fmt.Errorf("Error creating static/styles folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("static/scripts", 0755)
    if err != nil {
        err = fmt.Errorf("Error creating static/scripts folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("static/images", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating static/images folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("static/fonts", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating static/fonts folder\nInternal Error -> " + err.Error()) 
        return "", err
    }
    err = os.Mkdir("static/media", 0755)
    if err != nil {
        err = fmt.Errorf("Error creating static/media folder\nInternal Error -> " + err.Error())
        return "", err
    }
    return "\033[0;32mStatic folder 📁 created\033[0m", nil
}

func getHTMX() (string, error) {
    fmt.Println("Downloading \033[0;35mHTMX\033[0m of the GHST framework")
    url := "https://unpkg.com/htmx.org@1.9.6" 

    response, err := http.Get(url)
    if err != nil { 
        err = fmt.Errorf("Error downloading HTMX\nInternal Error -> " + err.Error())
        return "", err
    }
    defer response.Body.Close() 

    if response.StatusCode != http.StatusOK {
        err = fmt.Errorf("Error downloading HTMX\nStatus Code -> " + strconv.Itoa(response.StatusCode))
        return "", err
    }

    // create the file 
    file, err := os.Create(HTMXNAME)
    if err != nil { 
        fmt.Println("\033[0;31mError downloading HTMX\033[0m\nInternal Error -> " + err.Error())
        err = fmt.Errorf("Error downloading HTMX\nInternal Error -> " + err.Error())
        return "", err
    }
    defer file.Close() 
    _, err = io.Copy(file, response.Body)
    if err != nil {
        err := fmt.Errorf("Error downloading HTMX\nInternal Error -> " + err.Error())
        return "", err
    }

    Project.Htmx = "<script src=" + HTMXNAME + "></script>"
    return "\033[0;32mHTMX 🖼️ downloaded\033[0m", nil
}

func getTailwindCSS(noconfirm bool) (string, error) {
    fmt.Println("Bootsraping \033[0;34mTailwindCSS\033[0m of the GHST framework")
    fmt.Println("Continue? (y/n)") 
    fmt.Println("Tailwind CSS is a utility-first CSS framework for rapidly building custom user interfaces.")
    fmt.Println("https://tailwindcss.com/")
    fmt.Println("You may choose to skip this step and use your own CSS framework")

    var input string
    if noconfirm {
        input = "y"
    } else {
        fmt.Scanln(&input)
    }

    if input != "y" && input != "Y" && input != "" {
        return "GHST did not install TailwindCSS", nil
    }

    fmt.Println("Setting up\033[0;34mTailwindCSS\033[0m of the GHST framework")
    fmt.Println("GHST uses the standalone CLI of TailwindCSS")
    fmt.Println("https://tailwindcss.com/blog/standalone-cli")
    cmd := exec.Command("tailwindcss", "init")
    _, err := cmd.CombinedOutput()
    if err != nil { 
        err = fmt.Errorf("Error bootstraping TailwindCSS\nInternal Error -> " + err.Error()) 
        return "", err
    }
    
    // create the input.css file 
    file, err := os.Create(INPUTCSSNAME)
    if err != nil { 
        err = fmt.Errorf("Error creating input.css\nInternal Error -> " + err.Error()) 
        return "", err
    }
    defer file.Close()
    
    // write to the input.css file 
    _, err = file.WriteString(`
@tailwind base;
@tailwind components;
@tailwind utilities;
    `)
    if err != nil {
        err = fmt.Errorf("Error writing to input.css\nInternal Error -> " + err.Error()) 
        return "", err
    }

    // create the output.css  
    file, err = os.Create(TAILWINDCSSNAME) 
    if err != nil { 
        err = fmt.Errorf("Error creating output.css\nInternal Error -> " + err.Error()) 
        return "", err
    }
    defer file.Close() 

    Project.Css = "<link rel=\"stylesheet\" href=\"" + TAILWINDCSSNAME + "\">"

    return "\033[0;32mTailwindCSS 💨 installed\033[0m", nil
}

func bootstrapSurrealDB(noconfirm bool) (string, error) {
    // make surralDB pink 
    fmt.Println("Bootsraping \033[0;35mSurrealDB\033[0m of the GHST framework")
    fmt.Println("SurrealDB is a Futuristic Database for the Web")
    fmt.Println("Learn more at: https://surrealdb.com/")
    fmt.Println("Continue? (y/n)")
    var input string 
    if noconfirm { 
        input = "y"
    } else {
        fmt.Scanln(&input)
    }
    if input != "y" && input != "Y" && input != "" { 
        return "GHST did not install SurrealDB", nil
    }

    fmt.Println("Installing SurrealDB")
    cmd := exec.Command("go", "get", "-u", "github.com/surrealdb/surrealdb")
    output, err := cmd.CombinedOutput()
    if err != nil {
        err = fmt.Errorf("Error installing SurrealDB\nInternal Error -> " + err.Error())
        return "", err
    }
    fmt.Println(string(output))
    
    fmt.Println("SurrealDB is a database, so it needs to be managed.")
    fmt.Println("GHST does not manage this in this version.")
    fmt.Println("Please follow the instructions at https://surrealdb.com/docs/introduction/start to manage SurrealDB")
    fmt.Println("Please also refer to the Golang SDK at https://surrealdb.com/docs/integration/sdks/golang")
    Project.SurrealDB = true

    return "\033[0;32mSurrealDB 🌌 installed\033[0m", nil
}

func bootsrapMainFile(noconfirm bool) (string, error) {
    fmt.Println("Bootsraping the main.go file")
    fmt.Println("Continue? (y/n)")
    var input string 
    if noconfirm {
        input = "y"
    } else {
        fmt.Scanln(&input)
    }
    if input != "y" && input != "Y" && input != "" {
        return "GHST did not bootstrap the main.go file", nil
    }

    // create the main.go file 
    file, err := os.Create("main.go")
    if err != nil {
        err = fmt.Errorf("Error creating main.go\nInternal Error -> " + err.Error())
        return "", err
    }
    defer file.Close() 

    // write to the main.go file 
    _, err = file.WriteString(`
package main 

import ( 
    "github.com/gin-gonic/gin" 
    "github.com/surrealdb/surrealdb.go"
)

func main() { 
    /* GIN */
    r := gin.Default()
    r.LoadHTMLGlob("src/views/*")
    r.Static("/static", "./static")

    /* SURREALDB */
    db, err := surrealdb.New(ws://localhost:8000/rpc)
    if err != nil { 
        panic(err)
    }
    if _, err = db.SignIn(map[string]interface{}{
        "username": "CHANGE_ME",
        "password": "CHANGE_ME",
    }); err != nil {
        panic(err)
    }
    if _, err = db.Use("CHANGE_ME", "CHANGE_ME"); err != nil {
        panic(err)
    }

    /* ROUTES */
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })
    r.Run(":8080")
}
    `)
    if err != nil {
        err = fmt.Errorf("Error writing to main.go\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mmain.go 📜 bootstrapped\033[0m", nil
}

func bootstrapViewsFolder(noconfirm bool) (string, error) {
    fmt.Println("Bootsraping the views folder") 
    fmt.Println("Bootsraping the src/views/index.html file")
    err := os.Mkdir("src", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating src folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("src/views", 0755)
    if err != nil { 
        err = fmt.Errorf("Error creating src/views folder\nInternal Error -> " + err.Error())
        return "", err
    }

    // create the index.html file 
    file, err := os.Create("src/views/index.html")
    if err != nil { 
        err = fmt.Errorf("Error creating src/views/index.html\nInternal Error -> " + err.Error())
        return "", err
    }
    defer file.Close()

    // write to the index.html 
    _, err = file.WriteString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>` + Project.ProjectName + `</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">\n` + Project.Css + 
    "\n" + Project.Htmx +
    `
</head>
<body class="flex flex-1 bg-slate-700 text-fuchsia-500">
    <h1 class="text-blue-500">` + Project.ProjectName + ` Created by <b>GHOST</b></h1>
    <p>GHST: GIN + HTMX + SURREALDB + TAILWINDCSS</p>
    <br>
    <p>Change this page at src/views/index.html</p>
</body>
</html>
    `)
    if err != nil {
        err = fmt.Errorf("Error writing to src/views/index.html\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mviews folder 📁 bootstrapped\033[0m", nil
}

func bootstrapRoutesAndModelsFolders(noconfirm bool) (string, error) {
    fmt.Println("Bootsraping the routes and models folders")
    err := os.Mkdir("src/routes", 0755)
    if err != nil {
        err = fmt.Errorf("Error creating src/routes folder\nInternal Error -> " + err.Error())
        return "", err
    }
    err = os.Mkdir("src/models", 0755)
    if err != nil {
        err = fmt.Errorf("Error creating src/models folder\nInternal Error -> " + err.Error())
        return "", err
    }
    return "Created src/routes and src/models", nil
}
