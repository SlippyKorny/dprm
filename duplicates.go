package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// GetDupStr returns the output of the duplicate image search and removes duplicates if remove is
// true
func GetDupStr(path string, recursive bool, remove bool) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extrFilenames(path, recursive)
	if err != nil {
		return err.Error()
	}

	// Generate and save all of the hashes in pair with the corresponding files
	var h [][32]byte
	h, err = genHashes(f)
	if err != nil {
		return err.Error()
	}

	// Find the duplicates and store their hashes and names
	d := findDups(f, h)

	// Format the output (delete duplicates if delete flag is on) and return it
	var sb strings.Builder
	if len(d) == 0 {
		sb.WriteString("no duplicates found")
	} else {
		sb.WriteString(fmt.Sprintf("found %d duplicate instances:\n", len(d)))

		for k, v := range d {
			sb.WriteString(fmt.Sprintf("\t\"%s\" is the same as: ", k))

			sb.WriteString(fmt.Sprintf("%s ", v[0]))
			// Delete duplicates if flag is true
			if remove {
				err := os.Remove(v[0])
				if err != nil {
					sb.WriteString("(failed to delete that file)")
				} else {
					sb.WriteString("(successfully deleted)")
				}
			}

			for i := 1; i < len(v); i++ {
				sb.WriteString(fmt.Sprintf(", %s ", v[i]))

				// Delete duplicates if flag is true
				if remove {
					err := os.Remove(v[i])
					if err != nil {
						sb.WriteString("(failed to delete that file)")
					} else {
						sb.WriteString("(successfully deleted)")
					}
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// findDups finds duplicate files in the given files and hashes arrays
func findDups(files []string, hashes [][32]byte) map[string][]string {
	dups := make(map[string][]string)
	fLen := len(files)

	for i := 0; i < fLen-1; i++ {
		// If hashes[i] is already in map either as key or value then skip this
		if isInMap(dups, files[i]) {
			continue
		}

		for j := i + 1; j < fLen; j++ {
			// If hashes are exactly the same (conversion to slice from array)
			if bytes.Compare(hashes[i][:], hashes[j][:]) == 0 {
				// If it's a new entry then save like this
				if _, k := dups[files[i]]; !k {
					dups[files[i]] = make([]string, 1)
					dups[files[i]][0] = files[j]
				} else { // otherwise like this
					dups[files[i]] = append(dups[files[i]], files[j])
				}
			}
		}
	}

	return dups
}
