# Amazing time table generator
## How to use
- Download the artifact from the [GitHub actions](https://github.com/Liko05/amazingTimetable/actions)
- Two ways to use this program:
    - Run the program with the command line arguments
    - Run the program without the command line arguments and follow the instructions
- List of available command line arguments can be found by running the program with the argument `-h`
````
  -d        Enable debug level logging
  -g int    The number of generators (default 3)
  -h        Show help
  -p int    The time between progress updates in seconds (default 10)
  -r int    The number of graders (default 3)
  -t int    The time limit in seconds (default 180)
  ````
## How to build
- Install [Go](https://golang.org/)
- Clone this repository
- Run `go build` in the root of the repository
## How to docs
- Install godoc by running `go get golang.org/x/tools/cmd/godoc` or `go install golang.org/x/tools/cmd/godoc`
- Documentation can be seen by running `godoc -http=:6060` in the root of the repository and then going to [localhost:6060](http://localhost:6060/)
- Note that in order to see the documentation you need to change package name to something else than `main`. [More info](https://stackoverflow.com/questions/21778556/what-steps-are-needed-to-document-package-main-in-godoc)
## How to test
- Run `go test` in the root of the repository
## External libraries
- [logrus](https://github.com/sirupsen/logrus) for logging

## More info
- Most of the logger code is from [here](https://stackoverflow.com/questions/48971780/how-to-change-the-format-of-log-output-in-logrus)
