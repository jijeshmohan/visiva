package main

import (
	"fmt"
	"os"

	"./server"
	"github.com/jessevdk/go-flags"
)

const (
	version = "0.1"
)

type Options struct {
	Port    int    `short:"p" long:"port" description:"Specify the port" default:"8080"`
	Create  string `short:"n" long:"new" description:"Create a new project" value-name:"project name"`
	Version bool   `short:"v" long:"version" description:"Display version" default:"false"`
	Start   bool   `short:"s" long:"start" description:"Start the server" default:"false"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	if options.Version {
		fmt.Println("Version", version)
		os.Exit(0)
	}

	if options.Start {
		server.StartServer(options.Port)
	} else {
		fmt.Println("For print usage use -h.")
	}
}
