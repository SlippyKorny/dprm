package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TheSlipper/dprm/dprm"
)

// args contains runtime arguments for the drpm command
type Args struct {
	Method    string
	Remove    bool
	Recursive bool
	CSV       bool
	Verbose   bool
	Directory string
}

// loadArgs creates the arguments struct with the passed arguments
func loadArgs() Args {
	var a Args

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `dprm version 0.0.3
Made by Kornel Domeradzki
Source: http://github.com/TheSlipper/dprm

dprm is a simple commandline hash based duplicate image search and removal tool.

dprm comes with ABSOLUTELY NO WARRANTY.  This is free software, and you
are welcome to redistribute it under certain conditions.  See the GNU General
Public Licence for details.

Usage: dprm [OPTION...] [DIRECTORY]

--method string
		Specifies the method with which the duplicates are searched for.
		Available methods are 'hashes' (default) and 'perceptual'.
--remove
		If set to true will remove the duplicates autonomously.
--recursive
		If set to true will recursively traverse the folder tree.
--csv
		If set to true will output the duplicate images in csv format.
--verbose
		Verbosity of the command's execution. If remove argument is not
		set to true then the program will set verbose to true.
--help
		Prints out this help section.
`)
	}

	flag.StringVar(&a.Method, "method", "hashes", "")
	flag.BoolVar(&a.Remove, "remove", false, "")
	flag.BoolVar(&a.Recursive, "recursive", false, "")
	flag.BoolVar(&a.CSV, "csv", false, "")
	flag.BoolVar(&a.Verbose, "verbose", false, "")
	flag.Parse()

	if flag.NArg() == 0 {
		// flag.Usage()
		fmt.Printf("No target directory specified\n")
		os.Exit(1)
	}

	// TODO: This should first get the arguments, check the length
	// and throw an error if directory not provided
	a.Directory = flag.Args()[0] // directory

	return a
}

func main() {
	a := loadArgs()

	var format string
	if a.CSV {
		format = "csv"
	} else {
		format = "normal"
	}

	fmt.Println(dprm.Run(format, a.Method, a.Directory, a.Recursive, a.Remove))
}
