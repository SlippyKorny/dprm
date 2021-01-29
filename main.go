package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// args contains runtime arguments for the drpm command
type args struct {
	verbose   bool
	remove    bool
	print     bool
	recursive bool
	directory string
}

// loadArgs creates the arguments struct with the passed arguments
func loadArgs() (a args, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
dprm version 0.0.1
Copyright (C) 2021 by Kornel Domeradzki
Source: http://github.com/TheSlipper/dprm

dprm is a simple commandline hash based duplicate image search and removal tool.

dprm comes with ABSOLUTELY NO WARRANTY.  This is free software, and you
are welcome to redistribute it under certain conditions.  See the GNU General Public Licence
for details.

Usage: dprm [-dir STRING] [-print | -remove] [-recursive] [-verbose]

-dir string
	defines the directory of operation
-print
		prints out the duplicates
-remove
		if set to true will remove the duplicates autonomously  
-recursive
		if set to true will recursively traverse the folder tree
-verbose
		verbosity of the command's execution
-help
		prints out this help section`)
	}

	flag.BoolVar(&a.verbose, "verbose", false, "verbosity of the command's execution")
	flag.BoolVar(&a.remove, "remove", false, "if set to true will remove the duplicates autonomously")
	flag.BoolVar(&a.print, "print", false, "prints out the duplicates")
	flag.BoolVar(&a.recursive, "recursive", false, "if set to true will recursively traverse the folder tree")
	flag.StringVar(&a.directory, "dir", "", "defines the directory of operation")
	flag.Parse()

	if !a.print && !a.remove {
		err = errors.New("neither print nor remove flag was selected - no action to perform")
	} else if a.print && a.remove {
		err = errors.New("exclusive actions print and remove were selected")
	} else if a.directory == "" {
		err = errors.New("no path selected")
	}

	return
}

func main() {
	// Load the arguments
	a, err := loadArgs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Remove duplicates if the remove flag was selected
	s := GetDupStr(a.directory, a.recursive, a.remove)
	if !a.remove || a.verbose {
		fmt.Println(s)
	}
}
