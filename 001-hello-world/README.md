# Part01 - code structure

This example will cover how to set up the development environment, write the program, and run it.

## Information

- his basic program demonstrates importing packages, declaring functions, and basic output. 
- From here, you can explore more complex computational concepts and Go's vast ecosystem.


## Set Up Your Go Environment

Before you write your Go code, you need to ensure your environment is set up properly:
1. **Download and Install Go:**
   - Visit the official Go downloads page: [Go Downloads](https://go.dev/dl/)
   - Choose the correct package for your operating system (Windows, macOS, or Linux), download it, and follow the installation instructions.
2. **Verify Installation:**
   - Open a command line interface (CLI) and type `go version`. This command should print the installed version of Go, confirming that Go is properly installed, for example:
   ```text
     go version go1.22.3 darwin/arm64
   ```

## Write the Go Program
1. **Create a New Directory for Your Project:**
   - Open the CLI and navigate to a place where you want to keep your Go projects.
   - Create a new directory named "hello-world" and navigate into it:
      - `mkdir hello-world`
      - `cd hello-world`
2. **Create a New Go File:**
   - Within the `hello-world` directory, create a file named `main.go`. 
   - You can do this using a text editor or from the command line with `touch main.go`.

3. **Write Your Go Code:**
   - Open `main.go` in a text editor.
   - Enter the following Go code:
   ```go
   // Define the package name, which is main for executables
   package main

   // Import the fmt package which contains functions for formatting text, including printing to the console
   import "fmt"

   // Define the main function, which is the entry point of the application
   func main() {
   // Call the Println function of the fmt package to print text to the console
   fmt.Println("Hello, World!")
   }
   ```
   - Save the file.

   ### Detailed Description of the Code
   - `package main`: Every Go file starts with a package declaration. Packages are Go's way of organizing and reusing code. The main package is special because it's used to build executables.
   - `import "fmt"`: Importing allows you to use code from other packages. The `fmt` package (short for format) provides I/O (Input/Output) functions.
   - `func main()`: The `main` function is the entry point of the app, where the execution begins. Note that Go executes the `main` function in the `main` package.
   - `fmt.Println("Hello, World!")`: `Println` is a function in the `fmt` package that prints text to the standard output followed by a newline. Here it's used to output **"Hello, World!"**.

## Run Your Program
   - Go back to your CLI, make sure you are in the "hello-world" directory.
   - Run the following command:
   ```bash
      go run main.go
   ```
   - You should see:
   ```text
   Hello World!
   ```
   
## Congratulations!