package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	version = "0.1"
)

type Options struct {
	Port    int    `short:"p" long:"port" description:"Specify the port" default:"8080"`
	Create  string `short:"n" long:"new" description:"Create a new project" value-name:"project name"`
	Version bool   `short:"v" long:"version" description:"Display version" default:"false"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	fmt.Println("Port Number ", options.Port)
	fmt.Println("Create ", options.Create)
	if options.Version {
		fmt.Println("Version", version)
	}

}
