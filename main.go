package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func CreateProjectRootDir(dirName string) {
	//Create root directory for project
	err := os.Mkdir(dirName, 0750)
	if err != nil && !os.IsExist(err) {
		StandardErrHandler(err)
	}
	//cd into the directory created
	err = os.Chdir(dirName)
	if err != nil {
		StandardErrHandler(err)
	}
	CreateRootDirLevelFiles()
	CreateRootLevelPythonModules()
} //end of CreateProjectRootDir
func CreateRootLevelPythonModules() {
	//Create root level Python modules -database, models, routes & tests for now
	rootLevelDirectoriesToBeCreated := []string{"database", "models", "routes", "test"}
	var totalNumberOfRootLevelDirectoriesToBeCreated = len(rootLevelDirectoriesToBeCreated)
	for i, dirName := range rootLevelDirectoriesToBeCreated {
		err := os.Mkdir(dirName, 0750)
		if err != nil && !os.IsExist(err) {
			StandardErrHandler(err)
		}

		//Make those directories just created into modules by creating a dunder init file in them.
		rootLevelPythonModule, err := os.Create(dirName + "/__init__.py")
		StandardErrHandler(err)

		log.Println("Python module created: " + rootLevelPythonModule.Name())
		log.Println("Index completed in creation of root level Python modules: " + strconv.Itoa(i) + " out of " + strconv.Itoa(totalNumberOfRootLevelDirectoriesToBeCreated))

	}
} //end of CreateRootLevelPythonModules

func CreateRootDirLevelFiles() {
	//creating main.py, Dockerfile in the directory created by IntelliJ for now...
	listOfRootLevelFilesToBeCreated := []string{"main.py", "Dockerfile"}
	var totalNumberOfRootLevelFiles = len(listOfRootLevelFilesToBeCreated)

	for i, s := range listOfRootLevelFilesToBeCreated {
		var rootLevelFile, _ = os.Create(s)
		err := rootLevelFile.Close()
		StandardErrHandler(err)
		log.Println("File created: " + rootLevelFile.Name())
		log.Println(strconv.Itoa(i) + " out of " + strconv.Itoa(totalNumberOfRootLevelFiles) + " completed.")
	}
} //end of CreateRootDirLevelFiles

func CreateInfoPlistFile(pListcontent InfoPListFileContent) {

	//Create Info.plist file
	infoPlistFile, err := os.Create("Info.plist")
	StandardErrHandler(err)

	//Write to Info.plist file
	var docType = "<.DOCTYPE plist PUBLIC '-//Apple//DTD PLIST 1.0//EN' " +
		"'http://www.apple.com/DTDs/PropertyList-1.0.dtd'>"
	_, err = infoPlistFile.WriteString(docType)
	StandardErrHandler(err)

	values := reflect.ValueOf(pListcontent)
	types := values.Type()

	for i := 0; values.NumField() > i; i++ {
		//_, err = infoPlistFile.WriteString("<key>" + i + "</key>" + "\n" + "<string>" + s + "</string>")
		_, err = infoPlistFile.WriteString(types.Field(i).Name + "\n" + values.Field(i).String() + "\n")
	}

	defer infoPlistFile.Close()

} //end of CreateInfoPlistFile

func copyResource(sourceFile string, targetDir string) {
	//take the string reference of icon.icsn and get *File type that has  Reader that
	//has a reader that io.Copy can use for the copy operation.
	icon, err := os.Open(sourceFile)
	StandardErrHandler(err)

	//take the target dir to Create a *File object that has a Writer that io.Copy
	//can use for the copy operation.
	dst, err := os.Create(filepath.Join(targetDir, filepath.Base(sourceFile)))
	StandardErrHandler(err)

	//do the actual copy operation with icon's reader and dst's writer.
	_, err = io.Copy(dst, icon)
	StandardErrHandler(err)
} //end of copyResource

func CreateAppBundleStructure(appNameParam string, projectRootDirParam string, newBinaryNameParam string, pListcontentParam InfoPListFileContent) {
	appBundleStructure := []string{appNameParam, "Contents", "MacOS", "Resources"}
	for _, s := range appBundleStructure {
		err := os.Mkdir(s, 0750)
		if err != nil && !os.IsExist(err) {
			StandardErrHandler(err)
		}
		err = os.Chdir(s)
		if err != nil && !os.IsExist(err) {
			StandardErrHandler(err)
		}

		if s == "Contents" {
			CreateInfoPlistFile(pListcontentParam)
		}

		if s == "MacOS" {
			log.Println("Now in MacOS dir")
		}

		if s == "Resources" {
			targetDir := "."
			sourceFile := "/Users/gabe.cruz/wrk_cool/bb8/resources/bb8.icns"
			copyResource(sourceFile, targetDir)
		}
	} //appBundleStructure for loop
	//the file structure is built
	//cd back to project root dir
	err := os.Chdir(projectRootDirParam)
	StandardErrHandler(err)

	BuildApp(appNameParam, newBinaryNameParam, err)
} //CreateAppBundleStructure

func BuildApp(appNameParam string, newBinaryNameParam string, err error) {
	//now run the command line tool to create the binary
	newBinaryNameTargetPath := filepath.Join(appNameParam, "Contents", "MacOS", newBinaryNameParam)
	cmd := exec.Command("go", "build", "-o", newBinaryNameTargetPath)
	cmd.Stdin = strings.NewReader("main.go")
	var out strings.Builder
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error in building the binary:%v\n" + err.Error())
	}
	StandardErrHandler(err)
} //BuildApp

func Usage() {
	log.Println("USAGE::")
	log.Println("cd into directory where you want to create your project and type::")
	log.Println("fapi oi <project_name> the oi parameter stands for 'create project from outside in'.")
	log.Println("or cd into already existing <project_name> dir and type:: fapi io which stands for " +
		"'create project from inside out'.")
	log.Println("No parameters passed reports Usage.")
	return
} //end of Usage

func StandardErrHandler(e error) {
	if e != nil {
		panic(e)
	}
} //end of StandardErrHandler

type InfoPListFileContent struct {
	CFBundleDisplayName           string
	CFBundleExecutable            string
	CFBundleIconFile              string
	CFBundleIdentifier            string
	CFBundleInfoDictionaryVersion string
	CFBundlePackageType           string
	CFBundleShortVersionString    string
	CFBundleVersion               string
} //end of InfoPListFileContent struct

func main() {
	var appName string
	var projectRootDir string
	var newBinaryName string
	var appNameDotApp string

	infoPListFileContent := InfoPListFileContent{
		CFBundleDisplayName:           "MyApp",
		CFBundleExecutable:            "Binary",
		CFBundleIconFile:              "icon.icns",
		CFBundleIdentifier:            "com.mycompany.myapp",
		CFBundleInfoDictionaryVersion: "6.0",
		CFBundlePackageType:           "APPL",
		CFBundleShortVersionString:    "1.0",
		CFBundleVersion:               "1",
	}

	//parameter checking engine
	numOfParameters := len(os.Args)
	if numOfParameters > 0 {
		if os.Args[1] == "create_fapi_project_io" {
			CreateRootDirLevelFiles()
			CreateRootLevelPythonModules()
		} else if os.Args[1] == "create_fapi_project_oi" {
			if true {
				CreateProjectRootDir(os.Args[2])
			}
		} else if os.Args[1] == "bundleApp" {
			//Create the app bundle
			if numOfParameters == 4 { //For this code to run the 4 parameters in args are:
				//1. "bundleApp" value
				//2. "appName" variable
				//3. "projectRootDir" variable
				//4. "newBinaryName" variable
				if true { //if bundleApp request is not empty
					if true { //if appName is not empty
						appName = os.Args[2]
					}
					if true { //if projectRootDir is not empty
						projectRootDir = os.Args[3]
					}
					if true { //if newBinaryName is not empty
						newBinaryName = os.Args[4]
					}

					appNameDotApp = appName + ".app"

					CreateAppBundleStructure(appNameDotApp, projectRootDir, newBinaryName, infoPListFileContent)
					//shopping list:
					//[*]create a function that will take appName, projectRootDir, and newBinaryName
					//[*] run from the command line the following command::
					//go build -o ${appName}.app/Contents/MacOS/${YourGoBinary} main.go
					//Additonal Research Notes::
					//This guy is saying to bootstrap the project root dir by passing a
					//a project root dir and then he shares some source that makes it look like he's
					//using something called a flag and it looks like the way it can be used is to
					//ask the user for parameters in the command line.
					//https://stackoverflow.com/questions/47531760/how-to-get-the-root-path-of-the-project
					//Copilot Notes::
					//https://stackoverflow.com/questions/28322997/how-to-pass-arguments-to-go-program
					//The top one is the one I found the second one is one that Copilot suggested.
				}
			}
		} else {
			Usage()
		}
	}
} //end of main
