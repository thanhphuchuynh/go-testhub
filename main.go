package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func runTest(testName string) (string, error) {
	cmd := exec.Command("go", "test", "-run", testName)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func runTestsInFolder(folderPath string) (string, error) {
	// Change the current working directory to the specified folder
	err := os.Chdir(folderPath)
	if err != nil {
		return "", fmt.Errorf("failed to change directory: %v", err)
	}

	// Run the tests in the current directory (which is now the specified folder)
	cmd := exec.Command("go", "test", "./...")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	return out.String(), err
}

func listTestCases(packagePath string) ([]string, error) {

	err := os.Chdir(packagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to change directory: %v", err)
	}

	// Command to list all test cases
	cmd := exec.Command("go", "test", `./...`, `-list=.`)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the command
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to list test cases: %v", err)
	}

	// Parse the output to extract test case names
	lines := strings.Split(out.String(), "\n")
	var testCases []string
	for _, line := range lines {
		if strings.HasPrefix(line, "Test") {
			testCases = append(testCases, line)
		}
	}

	return testCases, nil
}

func main() {
	// Create a Gin router
	r := gin.Default()

	// Serve static files from the "build" directory (where React files are located)
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// Serve the main index.html file
	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.GET("/api/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Catch-all route to serve index.html for React Router
	r.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	// Start the server on port 8080 (or any other port)
	r.Run(":8080")
}
