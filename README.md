# Amazing time table generator
## How to use
- Download the artifact from the [GitHub actions](https://github.com/Liko05/amazingTimetable/actions) or build it yourself
- Two ways to use this program:
    - Run the program with the command line arguments
    - Run the program without the command line arguments and follow the instructions
- List of available command line arguments can be found by running the program with the argument `-h` `-help` or just inputting invalid arguments
````
  -d    Enable debug level logging
  -h    Show help
  -p int The time between progress updates in seconds (default 10)
  -t int The time limit in seconds (default 180)
  -w int Number of generator - grader pairs (default = logicalProcessors / 2)
  ````
## How to build
- Install [Go](https://golang.org/)
- Clone this repository
- Run `go build` in the root of the repository
## How to docs
- Install godoc by running `go get golang.org/x/tools/cmd/godoc` or `go install golang.org/x/tools/cmd/godoc`
- Documentation can be seen by running `godoc -http=:6060` in the root of the repository and then going to [localhost:6060/pkg/amazingTimetable](http://localhost:6060/pkg/amazingTimetable/)
## How to test
- Run `go test` in the root of the repository
- To see coverage of each package change directory to the package directory and run `go test -cover`
## External libraries
- [logrus](https://github.com/sirupsen/logrus) for logging
- [murmur3](https://github.com/spaolacci/murmur3) for hashing purposes
- [table](https://github.com/rodaine/table) for printing pretty tables
## More info
- Most of the logger code is from [here](https://stackoverflow.com/questions/48971780/how-to-change-the-format-of-log-output-in-logrus)
## Implementation notes
- The program is split into 3 parts:
    - The generator
    - The grader
    - The worker
- Grader is being initialized together with the generator as a worker pair, so they can communicate via exclusive channel
- Map containing the hashes is being initialized with the size of 200_000_000 because resizing the map is expensive and in most cases GC gets overwhelmed and system starts to run out of memory
- Number mappings of subjects, teachers and classes is defined in [ImplNotes.md](ImplNotes.md)