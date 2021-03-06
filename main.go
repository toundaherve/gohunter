package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	greating()
	section("Environment", detectEnvironment)
	section("OS signals", reportSignals)
}

func greating() {
	fmt.Println("Welcome to GoHunter\n")
}

func detectEnvironment() {
	environment := make(map[string]string)

	environment["Operating system"] = runtime.GOOS
	environment["Architecture"] = runtime.GOARCH
	environment["Go version"] = runtime.Version()
	environment["Root directory"] = runtime.GOROOT()
	environment["Max threads"] = os.Getenv("GOMAXPROCS")
	environment["Num of cores"] = fmt.Sprint(runtime.NumCPU())

	for env, value := range environment {
		fmt.Printf("%v : %v\n", env, value)
	}
}

type fn func()

func section(title string, fn fn) {
	fmt.Println(title)
	fmt.Println("-------------------------------")
	fn()
	fmt.Print("\n")
}
