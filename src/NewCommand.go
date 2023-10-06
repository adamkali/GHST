package src

import (
	"fmt"
	"io"
	"io/ioutil"
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
    STEPTOTAL = 10
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

type ghostAsset struct { 
    Name string 
    URL string
}

var ghostAssets = []ghostAsset{ 
    { Name: "gin", URL: "https://raw.githubusercontent.com/gin-gonic/logo/master/color.svg", },
    { Name: "surrealdb", URL: "https://surrealdb.com/static/img/assets/logo/logo-3fccfc517c1fa85d61441f736f7bb6ac.svg", },
    { Name: "tailwindcss", URL: "https://tailwindcss.com/_next/static/media/tailwindcss-mark.3c5441fc7a190fb1800d4a5c7f07ba4b1345a9c8.svg", },
    // oops i cant find one for htmx
}

func ShowStep(stepMessage string) {
    step += 1
    fmt.Println("( " + strconv.Itoa(step) + "/" + strconv.Itoa(STEPTOTAL) +  ")\t" + stepMessage + "\n")
}

func ErrOut(err error) {
    fmt.Println("🛑\t\033[0;31m" + err.Error() + "\033[0m")
    os.Exit(1)
}

func CreateProject(projectName string, projectPath string, noconfirm bool) {
    var err error 
    var signal string
    Project.ProjectName = projectName
    ghostTask := GhostTask{
        Name: "Creating project: " + projectName,
        State: "📁",
    }

    ghostTask.Progress()
    if _, err = os.Stat(projectPath); os.IsNotExist(err) {
        ErrOut(err)
    } else {
        os.Chdir(projectPath)
    }
    ghostTask.Name = "Created project: " + projectName + " at " + projectPath
    ghostTask.Progress()

    signal, err = bootstrapGoProject(projectPath, noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🍸"
        ghostTask.Progress()
    }

    bootsrapStaticFolder()
    ghostTask.Name = "Created the /static folder"
    ghostTask.State = "📁"
    ghostTask.Progress()

    signal, err = getHTMX()
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📦"
        ghostTask.Progress()
    }

    signal, err = getTailwindCSS(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "💨"
        ghostTask.Progress()
    }
    
    signal, err = bootstrapSurrealDB(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🌌"
        ghostTask.Progress()
    }

    signal, err = bootsrapMainFile(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📃"
        ghostTask.Progress()
    }

    signal, err = bootstrapViewsFolder(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🔭"
        ghostTask.Progress()
    }

    signal, err = bootstrapRoutesAndModelsFolders(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📁"
        ghostTask.Progress()
    }
    
    signal, err = createGhostYamlFile()
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "👻"
        ghostTask.Progress()
    }

    ghostTask.Name = "🎉\t\033[0;32mProject created successfully\033[0m"
    ghostTask.State = "🍾"
    ghostTask.Progress()
}

func CreateProjectCurrentDir(projectName string, noconfirm bool) {
    fmt.Println("Creating project: " + projectName + " in the current directory")

    var err error 
    var signal string 
    Project.ProjectName = projectName 
    ghostTask := GhostTask{
        Name: "Creating project: ./" + projectName,
        State: "📁",
    }
    ghostTask.Progress()

    signal, err = bootstrapGoProject(projectName, noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🍸"
        ghostTask.Progress()
    }

    bootsrapStaticFolder()
    ghostTask.Name = "Created the /static folder"
    ghostTask.State = "📁"
    ghostTask.Progress()

    signal, err = getHTMX()
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📦"
        ghostTask.Progress()
    }

    signal, err = getTailwindCSS(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "💨"
        ghostTask.Progress()
    }
    
    signal, err = bootstrapSurrealDB(noconfirm)
    if err != nil { 
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🌌"
        ghostTask.Progress()
    }

    signal, err = bootsrapMainFile(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📃"
        ghostTask.Progress()
    }

    signal, err = bootstrapViewsFolder(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "🔭"
        ghostTask.Progress()
    }

    signal, err = bootstrapRoutesAndModelsFolders(noconfirm)
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "📁"
        ghostTask.Progress()
    }

    signal, err = createGhostYamlFile()
    if err != nil {
        ErrOut(err)
    } else {
        ghostTask.Name = signal
        ghostTask.State = "👻"
        ghostTask.Progress()
    }

    ghostTask.Name = "🎉\t\033[0;32mProject created successfully\033[0m"
    ghostTask.State = "🍾"
    ghostTask.Progress()
}

func bootstrapGoProject(projectName string, noconfirm bool) (string, error) {
    var err error
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
    _, err = cmd.CombinedOutput()
    if err != nil {
        err = fmt.Errorf("\033[0;31mError creating go.mod\033[0m\nInternal Error -> " + err.Error())
        return "", err
    }

    cmd = exec.Command("go", "get", "-u", "github.com/gin-gonic/gin")
    _, err = cmd.CombinedOutput()
    if err != nil { 
        err = fmt.Errorf("\033[0;31mError installing Gin\033[0m\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mGin installed\033[0m", nil
}

func bootsrapStaticFolder() (string, error) {
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
    err = getAssets("static/images")
    if err != nil { 
        err = fmt.Errorf("Error downloading assets\nInternal Error -> " + err.Error()) 
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

    Project.Htmx = fmt.Sprintf("<script src=\"/%s\"></script>", HTMXNAME)
    return "\033[0;32mHTMX downloaded\033[0m", nil
}

func getTailwindCSS(noconfirm bool) (string, error) {
    // create the tailwind.config.js file 
    file, err := os.Create("./tailwind.config.js")
    if err != nil {
        err = fmt.Errorf("Error creating tailwind.config.js\nInternal Error -> " + err.Error())
        return "", err
    }
    defer file.Close()
    _, err = file.WriteString(`
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ './src/**/*.html', ],
  theme: {
    extend: {},
  },
  plugins: [],
}`)
    if err != nil {
        err = fmt.Errorf("Error writing to tailwind.config.js\nInternal Error -> " + err.Error())
        return "", err
    }
    
    // create the input.css file 
    file, err = os.Create(INPUTCSSNAME)
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

    return "\033[0;32mTailwindCSS installed\033[0m", nil
}

func bootstrapSurrealDB(noconfirm bool) (string, error) {
    cmd := exec.Command("go", "get", "-u", "github.com/surrealdb/surrealdb.go")
    _, err := cmd.CombinedOutput()
    if err != nil {
        err = fmt.Errorf("Error bootstraping SurrealDB\nInternal Error -> " + err.Error())
        return "", err
    }
    Project.SurrealDB = true
    return "\033[0;32mSurrealDB installed\033[0m", nil
}

func bootsrapMainFile(noconfirm bool) (string, error) {
    cmd := exec.Command("go", "get", "-u", "github.com/adamkali/ghost_utils")
    _, err := cmd.CombinedOutput()
    if err != nil {
        err = fmt.Errorf("Error bootstraping Ghost Utils\nInternal Error -> " + err.Error())
        return "", err
    }
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
	"fmt"

	ghostutils "github.com/adamkali/ghost-utils/pkg/ghost_utils"
	"github.com/gin-gonic/gin"
)

func main() { 
    /* GHOST-UTILS */
    // load ghost config from the root of the project
    conf, err := ghostutils.New()
    if err != nil {
        panic(err)
    }

    // create a new gin engine
    r := gin.Default()

    // do basic setup with your config
    db, err := conf.Setup(r)
    if err != nil {
        panic(err)
    }
   
    // create a new basic route
    // it is suggested to make a controller with 
    // *surrealdb.DB as a field and use it in the
    // controller methods
    var basicRoute ghostutils.BasicRoute
    index := basicRoute.New("/", db)
    index.RG().GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })

    // run the server on your config's port :)
    r.Run(fmt.Sprintf(":%d", conf.Port))
}
    `)
    if err != nil {
        err = fmt.Errorf("Error writing to main.go\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32m main.go bootstrapped\033[0m", nil
}

func bootstrapViewsFolder(noconfirm bool) (string, error) {
    err := os.Mkdir("src/views", 0755)
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
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    ` + Project.Css + 
    `
    ` + Project.Htmx +
    `
</head>
<body class="flex flex-1 bg-slate-700 text-fuchsia-500">
    <div class="flex flex-col justify-center items-center">
        <div class="flex flex-row justify-evenly items-center">
            <img src="/static/images/gin.svg" alt="gin" class="w-16 h-16">
            <img src="/static/images/surrealdb.svg" alt="ghost" class="w-16 h-16">
            <img src="/static/images/tailwindcss.svg" alt="ghost" class="w-16 h-16">
        </div>
        <span class="text-4xl text-blue-500">` + Project.ProjectName + `</span>
        <div class="flex flex-row text-xl text-fuchsia-500">
            <span class="mr-2">👻</span>
            <span class="mr-2">To edit this page, go to <b class"rounded bg-slate-900 text-slate-200">src/views/index.html</b></span>
            <span class="mr-2">👻</span>
        </div>
    </div>
</body>
</html>
    `)
    if err != nil {
        err = fmt.Errorf("Error writing to src/views/index.html\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mviews folder bootstrapped\033[0m", nil
}

func bootstrapRoutesAndModelsFolders(noconfirm bool) (string, error) {
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
    return "\033[0;32mroutes and models folders bootstrapped\033[0m", nil
}

func getAssets(pathToStaticImgDir string) error {
    // use the list of asset urls to download the assets 
    for _, asset := range ghostAssets {
        err := getAsset(asset.Name, asset.URL, pathToStaticImgDir)
        if err != nil {
            return err
        }
    }
    return nil
}

func getAsset(name string, url string, pathToStaticImgDir string) error {

    req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	res, err := http.DefaultClient.Do(req)

    if err != nil {
        err = fmt.Errorf("Error making the request to " + url + "\nInternal Error -> " + err.Error())
        return err
    }

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        err = fmt.Errorf("Error reading the response body\nInternal Error -> " + err.Error())
        return err
    }
    file, err := os.Create(pathToStaticImgDir + "/" + name + ".svg")
    if err != nil { 
        err = fmt.Errorf("Error creating the file at " + pathToStaticImgDir + "/" + name + ".svg" + "\nInternal Error -> " + err.Error())
        return err 
    }
    defer file.Close()
    if _, err = file.Write(body); err != nil {
        err = fmt.Errorf("Error writing to the file at " + pathToStaticImgDir + "/" + name + ".svg" + "\nInternal Error -> " + err.Error())
        return err
    }

    return nil
}

func createGhostYamlFile() (string, error) {
    // create the ghost.yaml file 
    file, err := os.Create("ghost.yaml")
    if err != nil {
        err = fmt.Errorf("Error creating ghost.yaml\nInternal Error -> " + err.Error())
        return "", err
    }
    defer file.Close()

    // write to the ghost.yaml file 
    _, err = file.WriteString(
`
# This is the configuration file for a ghost project 
ghost: 
  name: ` + Project.ProjectName + `
  version: 0.0.1
  description: A GHST project made with ghost
  port: 8080
  surrealdb:
    # the following is url for the SurrealDB database 
    # ensure that if you are using a remote database, that you 
    # have the correct url.
    surrealdb-url: ws://localhost:8000/rpc
    # the following are the credentials for the SurrealDB database 
    # that will be used for this project
    surrealdb-username: CHANGE_ME
    surrealdb-password: CHANGE_ME
    surrealdb-database: CHANGE_ME
    surrealdb-collection: CHANGE_ME
  tailwindcss: 
    input: ./input.css
    output: ./static/styles/output.css
`)
    if err != nil {
        err = fmt.Errorf("Error writing to ghost.yaml\nInternal Error -> " + err.Error())
        return "", err
    }

    return "\033[0;32mghost.yaml created\033[0m", nil
}
