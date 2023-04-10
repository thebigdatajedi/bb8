# bb8  - a FastAPI CLI

## The Origin Story

bb8 is the name I’ve given to the tiny project I’m doing of creating a CLI for FastAPI and doing it all in Go.  Yes, Golang! Well, to allow me to think about other things other than just Python and its ecosystem and because I anticipate that Python will be a big part of my life now, no matter where I go and what I do, and for who I do it, I needed a way to mitigate burnout on such a great set of tools and ecosystem.  So, I decided to write a CLI for FastAPI with as little functionality on it as I can, just the very essential, and the criteria are if I have to do a task for FastAPI manually more than three times or if I think I will have to do a task for a FastAPI project more than three times. Yes, If I think, what I mean by that is that many times I’ve done the same task for the same group of obligations way more than three times, but I lose focus on the number of times, and then it turns out I pierced the rule of three times for a task, and then I automate it. I suspect that many, many times, I’ve done a particular task manually closer to one hundred times.  If I’ve done something for FastAPI three or more times, it goes in the CLI.  I originally called the code base fapi but will change the name to bb8 because it seems more appropriate. 

I did some Google searches the closest thing I could find that would generate a project from a template but has everything, and the kitchen sink was::

[Project Generation - Template - FastAPI](https://fastapi.tiangolo.com/project-generation/)

Reading this Project Generation - Template page convinced me I needed to start a FastAPI CLI that would give me the bare minimum. 

I needed a very minimalistic set of features when starting a FastAPI project. The FastAPI offering for Project Generation was incompatible with something with fewer features that worked quickly with minimal configuration. Also, that would initially not provide such a huge security attack surface.  The offering by IntelliJ, combined with what bb8 can do, seems to be a perfect starting point for me today, 04-09-2023. This doesn’t mean it will always be perfect. Enhancements for me to have less dependency on IntelliJ templating will come. As the FastAPI framework expands and more boilerplate tasks will inevitably have to be dealt with, there will be more to come.  More to come!

## Building the project.

- It’s straightforward once you have Go installed on your system.
- Go to the command line, **cd to the project root dir.**  Type::

```go
go build bb8.go
```

- When building for a macOS app, the -o means you pass in the path as a parameter to build with the -o flag.  So, you don’t have to copy the main.go and build it in the macOS dir; go build will take care of putting the binary in the MacOS dir and renaming it for you.

```go
go build -o bb8.app/Contents/MacOS/bb8 bb8.go 
//Note the go file can be called anything, usually, 
//you will see it called main.go in our project; it's called bb8.go. 
```

- `go build [-o output]`
- The -o flag forces the build to write the resulting executable or artifact/deployable to the named output file or directory. If the named output is an existing directory or ends with a slash or backslash, then any resulting executables will be written to that directory.
    - In the build command above, you can see that main.go will build as an **artifact** called YourGoBinary inside the MacOS dir.

Source::

[go command - cmd/go - Go Packages](https://pkg.go.dev/cmd/go)

- Usage::
    - Once you have a binary called bb8, all you have to do is to get its usage is type in the command line::
    
    ```go
    bb8 <enter key>
    ```
    
    - At the time of this writing, 04-09-2023, this is what usage reports.
    
    ```go
    USAGE::
    1. bb8 - with no parameters passed reports usage.
    CREATE FastAPI project::
    1. cd into directory where you want to create your FastAPI project and type::
    2. bb8 create_fapi_project <project_name> - 'create fapi project'.
    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    ORGANIZE existing FastAPI project::
    1. cd into already existing <project_name> dir and type::
    2. bb8 organize_fapi_project - 'organize fap project'.
    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    BUNDLE APP for macOS::
    1. bb8 bundle
    2. Respond to the questions asked by the program.
    3. The program will create a .app bundle in the current directory.
    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    ```
    
    - bb8 was built with the ORGANIZE existing FastAPI projects in mind.
    - IntelliJ can create FastAPI projects with needful facilities like Poetry and installing uvicorn out of the box for the project, but there were missing things, and because I was spiking a lot, I didn’t want to add those things manually anymore.  So, I created the bb8 CLI to help me.
    - The CREATE FastAPI project does the same steps as the ORGANIZE but creates its own project root dir. So, this version of a project doesn’t have all the cool things IntelliJ offers for a FastAPI project out of the box.  It doesn’t have Poetry or uvicorn out of the box, but it has a tidy, organized project to which all those things and even counterparts can be added.  It creates a simple project folder structure that works as a pseudo-to-do list of FastAPI components to be created and how they should be organized. It opportunistically works as a reminder for a layered architecture in FastAPI (layered architecture as opposed to a microservices architecture or event-driven microservices architecture).
    - The BUNDLE APP allows users to use any Go application and bundle it as <applicationName>.app with Go, and it will build the <applicationName>.app bundle in the Go project root dir. - The BUNDLE APP feature came about in an interesting way.  I wanted to give bb8 a little branding, so I created a bb8.icns placeholder icon for bb8; I say placeholder because I’m not an artist, but I felt it was better than the default icon of the terminal icon that Go gives the binary once it’s compiled.  I researched how to add the icon to the app, and it kept stating that I needed to create an <applicationName>.app bundle and it explained how to add an icon.icns file to the bundle in a way that the app could use, and it would gain a more branded appearance.
    - All the app features except the BUNDLE APP work.
        - Currently, these are the issues::
            - The icon. icns is not appearing when the app is bundled. Just a blank icon appears.
            - As a command line interface app, the app doesn’t run when put in a folder in the path.
            - So the original reason for wanting to bundle the app, which was the branding via having a unique icon, did not work.
            - The app doesn’t work when placed in /usr/local/bin or /Applications.
            - I think that command line apps like CLI may not be intended to be bundled that way in macOS, although I read on StackOverflow that it could run, and it was even possible to debug command line applications in the bundle. I don’t know; I haven’t attempted it.  I haven’t gotten that far.  I will leave it in its current state, where it can create an inert bundle.
            - The bundling feature can work with bundling any app from any environment as long as the ecosystem for that language or framework provides the elements needed to create the bundle.
            - Go provided all the elements, but I couldn’t get it to run, so I have to work on debugging this problem.