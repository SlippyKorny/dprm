package main

import (
	"flag"
	"fmt"
	"os"
)

// args contains runtime arguments for the drpm command
type args struct {
	method    string
	remove    bool
	recursive bool
	csv       bool
	verbose   bool
	directory string
}

// loadArgs creates the arguments struct with the passed arguments
func loadArgs() args {
	var a args

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `dprm version 0.0.3
Copyright (C) 2021 by Kornel Domeradzki
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

	flag.StringVar(&a.method, "method", "hashes", "")
	flag.BoolVar(&a.remove, "remove", false, "")
	flag.BoolVar(&a.recursive, "recursive", false, "")
	flag.BoolVar(&a.csv, "csv", false, "")
	flag.BoolVar(&a.verbose, "verbose", false, "")
	flag.Parse()

	if flag.NArg() == 0 {
		// flag.Usage()
		fmt.Printf("No target directory specified\n")
		os.Exit(1)
	}

	a.directory = flag.Args()[0] // directory

	return a
}

func main() {
	// Load the arguments
	a := loadArgs()

	// Select the output style
	var style string
	if a.csv {
		style = "csv"
	} else if a.verbose {
		style = "normal"
	} else {
		style = "none"
	}

	// Remove duplicates if the remove flag was selected
	var s string
	if a.method == "hashes" {
		s = GetHashDupStr(a.directory, a.recursive, a.remove, style)
	} else if a.method == "perceptual" {
		s = GetPerceptualDupStr(a.directory, a.recursive, a.remove, style)
	} else {
		fmt.Printf("No such method as '%s'\n", a.method)
		os.Exit(2)
	}

	fmt.Println(s)
}
