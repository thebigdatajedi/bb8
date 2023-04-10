package main

import (
	"fmt"
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
		count := i + 1
		log.Println("Index completed in creation of root level Python modules: " + strconv.Itoa(count) + " out of " + strconv.Itoa(totalNumberOfRootLevelDirectoriesToBeCreated))

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
		count := i + 1
		log.Println(strconv.Itoa(count) + " out of " + strconv.Itoa(totalNumberOfRootLevelFiles) + " completed.")
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
		_, err = infoPlistFile.WriteString(types.Field(i).Name + "\n" + values.Field(i).String() + "\n")
	}

	defer infoPlistFile.Close()

} //end of CreateInfoPlistFile

// Context::
// Lots of comments is usually a sign of code smell.
// There are exceptions to every rule, there are exceptions to this rule too.
// The steps in copyResource use a concept of a Reader and a Writer that I've not had to reason about before.
// In experiences with other languages the source resource and the destination resources were both strings.
// If you look at this code I had to do the os.Open so that fileToCopy would have a Reader that io.Copy could use.
// I also had to do the os.Create so that dst would have a Writer that io.Copy could use.
// I'd like to keep this context text and comments in the function for a little while,
// while I gel a little more with this concept.
func copyResource(sourceFile string, targetDir string) {
	//take the string reference of fileToCopy.icsn and get *File type that has  Reader that
	//has a reader that io.Copy can use for the copy operation.
	fileToCopy, err := os.Open(sourceFile)
	StandardErrHandler(err)

	//take the target dir to Create a *File object that has a Writer that io.Copy
	//can use for the copy operation.
	dst, err := os.Create(filepath.Join(targetDir, filepath.Base(sourceFile)))
	StandardErrHandler(err)

	//do the actual copy operation with fileToCopy's reader and dst's writer.
	_, err = io.Copy(dst, fileToCopy)
	StandardErrHandler(err)
} //end of copyResource

func CreateAppBundleStructure(appNameParam string, projectRootDirParam string, newBinaryNameParam string,
	bb8IcnsResourceLocationSourceParam string, pListcontentParam InfoPListFileContent) {
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
			sourceFile := bb8IcnsResourceLocationSourceParam
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
	cmd.Stdin = strings.NewReader("bb8.go")
	var out strings.Builder
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error in building the binary:%v\n" + err.Error())
	}
	StandardErrHandler(err)
} //BuildApp

func Usage() {
	fmt.Println("USAGE::")
	fmt.Println("1. bb8 - with no parameters passed reports usage.")
	fmt.Println("CREATE FastAPI project::")
	fmt.Println("1. cd into directory where you want to create your FastAPI project and type::")
	fmt.Println("2. bb8 create_fapi_project <project_name> - 'create fapi project'.")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("ORGANIZE existing FastAPI project::")
	fmt.Println("1. cd into already existing <project_name> dir and type::")
	fmt.Println("2. bb8 organize_fapi_project - 'organize fap project'.")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	fmt.Println("BUNDLE APP for macOS::")
	fmt.Println("1. bb8 bundle")
	fmt.Println("2. Respond to the questions asked by the program.")
	fmt.Println("3. The program will create a .app bundle in the current directory.")
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

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
	//Shopping List::
	//todo [*]Parameterize the Info.plist file content
	//todo [*]rename bundleApp to bundleSelf
	//todo [*] add bundleApp to bundle other applications that have a UI.
	//todo [*] change usage to print to screen and not log to screen.
	//todo [] add flags to the command line tool and remove the action label style of passing parameters.
	//todo [*] make usage easier to read and use.
	//todo [*] message the user better when a cli action has completed.
	//todo [] configure bb8 so that logging goes to screen and to a log file all the time and make that configuration easy to change via flag and config file.
	var appName string
	var projectRootDir string
	var newBinaryName string
	var appNameDotApp string
	var bb8IcnsName string
	var bb8IcnsBundleLocationTarget string
	var bb8IcnsResourceLocationSourceDirectory string
	var bb8IcnsResourceLocationSource string
	var bundleIDPrefix string
	var bundleID string
	var bundleInfoDictionaryVersion string
	var bundlePackageType string
	var bundleShortVersionString string
	var bundleVersion string

	numOfParameters := len(os.Args)
	if numOfParameters > 1 {
		if os.Args[1] == "organize_fapi_project" {
			CreateRootDirLevelFiles()
			CreateRootLevelPythonModules()
		} else if os.Args[1] == "create_fapi_project" {
			if true { //this is checking if projectName is not empty
				CreateProjectRootDir(os.Args[2])
			}
		} else if os.Args[1] == "bundle" {
			//gather data for buildApp function
			fmt.Println("Enter your appName: ")
			fmt.Scanln(&appName)
			fmt.Println("Enter your projectRootDir: ")
			fmt.Scanln(&projectRootDir)
			fmt.Println("Enter your newBinaryName: ")
			fmt.Scanln(&newBinaryName)
			fmt.Println("Enter your bb8IcnsName: ")
			fmt.Scanln(&bb8IcnsName)
			fmt.Println("Enter your bundleIDPrefix: ")
			fmt.Scanln(&bundleIDPrefix)
			fmt.Println("Enter your bb8IcnsResourceLocationSourceDirectory: ")
			fmt.Scanln(&bb8IcnsResourceLocationSourceDirectory)

			fmt.Println("appName: ", appName)
			fmt.Println("projectRootDir: ", projectRootDir)
			fmt.Println("newBinaryName: ", newBinaryName)
			fmt.Println("bb8IcnsName: ", bb8IcnsName)
			fmt.Println("bundleIDPrefix: ", bundleIDPrefix)
			fmt.Println("bb8IcnsResourceLocationSourceDirectory: ", bb8IcnsResourceLocationSourceDirectory)

			//dynamic Info.plist file content - from parameters passed to the command line
			bundleID = bundleIDPrefix + "." + appName
			appNameDotApp = appName + ".app"
			bb8IcnsBundleLocationTarget = appNameDotApp + "/Contents/Resources/" + bb8IcnsName
			bb8IcnsResourceLocationSource = bb8IcnsResourceLocationSourceDirectory + "/" + bb8IcnsName

			//static Info.plist file content
			bundleInfoDictionaryVersion = "6.0"
			bundlePackageType = "APPL"
			bundleShortVersionString = "1.0"
			bundleVersion = "1"

			infoPListFileContent := InfoPListFileContent{
				CFBundleDisplayName:           appName,
				CFBundleExecutable:            newBinaryName,
				CFBundleIconFile:              bb8IcnsBundleLocationTarget,
				CFBundleIdentifier:            bundleID,
				CFBundleInfoDictionaryVersion: bundleInfoDictionaryVersion,
				CFBundlePackageType:           bundlePackageType,
				CFBundleShortVersionString:    bundleShortVersionString,
				CFBundleVersion:               bundleVersion,
			}

			CreateAppBundleStructure(appNameDotApp, projectRootDir, newBinaryName, bb8IcnsResourceLocationSource, infoPListFileContent)
			log.Println("App bundle created successfully!!!")
		}

	} else {
		Usage()
	}
} //end of main
