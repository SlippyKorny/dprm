package main

import (
	"fmt"
	"strings"
)

// GetDupStr returns the output of the duplicate image search
func GetDupStr(path string, recursive bool) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extFilenames()
	if err != nil {
		return err.Error()
	}

	// Generate and save all of the hashes in pair with the corresponding files
	var h []string
	h, err = genHashes(f)
	if err != nil {
		return err.Error()
	}

	// Find the duplicates and store their hashes and names
	d := findDups(f, h)

	// Format the output and return it
	var sb strings.Builder
	if len(d) == 0 {
		sb.WriteString("no duplicates found")
	} else {
		sb.WriteString(fmt.Sprintf("found %d duplicate images:\n", len(d)))
		for k, v := range d {
			var conc strings.Builder
			for vv := range v {
				conc.WriteString(fmt.Sprintf("%s ", vv))
			}
			sb.WriteString(fmt.Sprintf("%s is same as: %s", k, conc.String()))
		}
	}

	return sb.String()
}

// RmDup removes the duplicate files and returns a formatted output of their names
func RmDup(path string, recursive bool) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extFilenames()
	if err != nil {
		return err.Error()
	}

	// Generate and save all of the hashes in pair with the corresponding files
	var h []string
	h, err = genHashes()
	if err != nil {
		return err.Error()
	}

	// Find the duplicates and store their hashes and names

	// Remove duplicates and save the names of the removed files

	// Format the output and return it

	return ""
}

// genHashes generates hashes for the contents of the files that are pointed by the passed paths
func genHashes(paths []string) (hashes []string, err error) {
	return
}

// findDups finds duplicate files in the given files and hashes arrays
func findDups(files, hashes []string) (d map[string][]string) {
	return nil
}
