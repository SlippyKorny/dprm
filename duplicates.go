package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/vitali-fedulov/images"
)

// GetHashDupStr searches for duplicates with the content hash comparison method, prepares the output
// containing all of the duplicates and if the remove flag is set to true it also removes the duplicates
func GetHashDupStr(path string, recursive bool, remove bool, verbose bool) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extrFilenames(path, recursive)
	if err != nil {
		return err.Error()
	}

	// Generate and save all of the hashes in pair with the corresponding files
	var h [][32]byte
	h, err = genContentHashes(f)
	if err != nil {
		return err.Error()
	}

	// Find the duplicates and store their hashes and names
	d := findDupsByte(f, h)

	// Format the output (delete duplicates if delete flag is on) and return it
	return fmtDupOutput(remove, d)
}

// GetPerceptualDupStr searches for duplicates with the perceptual image comparison method, prepares
// the output containing all of the duplicates and if the remove flag is set to true it also removes
// the duplicates
func GetPerceptualDupStr(path string, recursive bool, remove bool, verbose bool) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extrFilenames(path, recursive)
	if err != nil {
		return err.Error()
	}

	// Generates average color value hashes for each one of them
	var h [][]float32
	var s []image.Point
	h, s, err = genAvgColourHashes(f)
	if err != nil {
		return err.Error()
	}

	// Find the duplicates and store their hashes and names
	d := findDupsFloat32(f, h, s)

	// Format the output (delete duplicates if delete flag is on) and return it
	return fmtDupOutput(remove, d)
}

// findDupsByte finds duplicate files with the use of byte hashes
func findDupsByte(files []string, hashes [][32]byte) map[string][]string {
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

// findDupsFloat32 finds duplicate files with the use of float32 hashes
func findDupsFloat32(files []string, hashes [][]float32, sizes []image.Point) map[string][]string {
	dups := make(map[string][]string)
	fLen := len(files)

	for i := 0; i < fLen-1; i++ {
		// If hashes[i] is already in map either as key or value then skip this
		if isInMap(dups, files[i]) {
			continue
		}

		for j := i + 1; j < fLen; j++ {
			// If hashes are exactly the same
			if images.Similar(hashes[i], hashes[j], sizes[i], sizes[j]) {
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

// fmtDupOutput formats a string with all the found duplicates and if necessary also deletes them
func fmtDupOutput(rm bool, d map[string][]string) string {
	var sb strings.Builder
	if len(d) == 0 {
		sb.WriteString("no duplicates found")
	} else {
		sb.WriteString(fmt.Sprintf("found %d duplicate instances:\n", len(d)))

		for k, v := range d {
			sb.WriteString(fmt.Sprintf("File %s is the same as:\n", k))

			for i := 0; i < len(v); i++ {
				sb.WriteString(fmt.Sprintf("\t%s", v[i]))

				// Delete duplicates if flag is true
				if rm {
					err := os.Remove(v[i])
					if err != nil {
						sb.WriteString("(failed to delete that file)")
					} else {
						sb.WriteString("(successfully deleted)")
					}
				}

				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
