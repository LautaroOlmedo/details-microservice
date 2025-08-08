### How to Run

**All command must be executed in the root folder**

run
* go mod tidy: Add the missing dependencies: Search your code (.go files) and add to the file go.mod all modules that are needed, but not yet listed. Remove unused dependencies: Remove from the go.mod file any modules that are no longer used in your code, making sure the file is accurate and up to date.

After that you can run:
* make build: This command compiles the Go source code into an executable file.
* make run: This command executes the application.
* make test: This command runs the project's unit tests.
* make clean: This command removes the build artifacts to clean the project directory.