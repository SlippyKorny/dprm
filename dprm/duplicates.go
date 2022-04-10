package dprm

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/vitali-fedulov/images"
)

type TraverseData struct {
	Path      string
	Remove    bool
	Recursive bool
	Format    string
	Ignored   [][]string
}

// Run runs the commandline utility. It accepts a pointer to arguments so that other applications can incorporate this
// tools functionality. If the pointer is nil then it reads the arguments from command line.
func Run(format, method, directory string, recursive, remove bool) string {
	data := TraverseData{
		Path:      directory,
		Format:    format,
		Recursive: recursive,
		Remove:    remove,
	}

	// Load ignored item list
	lines, err := loadIgnoreFileList()
	if err != nil {
		return err.Error()
	}
	data.Ignored = lines

	// Remove duplicates if the remove flag was selected
	var s string
	if method == "hashes" {
		s = GetHashDupStr(data)
	} else if method == "perceptual" {
		s = GetPerceptualDupStr(data)
	} else {
		fmt.Printf("No such method as '%s'\n", method)
		os.Exit(2)
	}

	return s
}

// GetHashDupStr searches for duplicates with the content hash comparison method, prepares the output
// containing all of the duplicates and if the remove flag is set to true it also removes the duplicates.
func GetHashDupStr(data TraverseData) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extrFilenames(data.Path, data.Recursive)
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
	d := findDupsByte(f, data.Ignored, h)

	// Format the output (delete duplicates if delete flag is on) and return it
	if data.Format == "normal" {
		return dupOutputTerm(data.Remove, d)
	} else if data.Format == "csv" {
		return dupOutputCSV(data.Remove, d)
	} else {
		return ""
	}
}

// GetPerceptualDupStr searches for duplicates with the perceptual image comparison method, prepares
// the output containing all of the duplicates and if the remove flag is set to true it also removes
// the duplicates.
func GetPerceptualDupStr(data TraverseData) string {
	// Extract the names of all files that are being taken into consideration
	f, err := extrFilenames(data.Path, data.Recursive)
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
	d := findDupsFloat32(f, data.Ignored, h, s)

	// Format the output (delete duplicates if delete flag is on) and return it
	if data.Format == "normal" {
		return dupOutputTerm(data.Remove, d)
	} else if data.Format == "csv" {
		return dupOutputCSV(data.Remove, d)
	} else {
		return ""
	}
}

// findDupsByte finds duplicate files with the use of byte hashes.
func findDupsByte(files []string, ignored [][]string, hashes [][32]byte) map[string][]string {
	dups := make(map[string][]string)
	fLen := len(files)

	for i := 0; i < fLen-1; i++ {
		// If hashes[i] is already in map either as key or value then skip this
		if isInMap(dups, files[i]) {
			continue
		}

		for j := i + 1; j < fLen; j++ {
			if shouldBeIgnored(files[i], files[j], ignored) {
				continue
			}

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

// findDupsFloat32 finds duplicate files with the use of float32 hashes.
func findDupsFloat32(files []string, ignored [][]string, hashes [][]float32, sizes []image.Point) map[string][]string {
	dups := make(map[string][]string)
	fLen := len(files)

	for i := 0; i < fLen-1; i++ {
		// If hashes[i] is already in map either as key or value then skip this
		if isInMap(dups, files[i]) {
			continue
		}

		for j := i + 1; j < fLen; j++ {
			if shouldBeIgnored(files[i], files[j], ignored) {
				continue
			}

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

// dupOutputTerm creates a string with all the found duplicates in a format for terminal output and if necessary it also
// deletes them.
func dupOutputTerm(rm bool, d map[string][]string) string {
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

// fmtDupOutputCSV creates a string with all the found duplicates in a csv format and if necessary it also deletes them.
func dupOutputCSV(rm bool, d map[string][]string) string {
	var sb strings.Builder
	if len(d) == 0 {
		sb.WriteString("")
	} else {
		sb.WriteString("original,duplicates\n")

		for k, v := range d {
			// sb.WriteString(fmt.Sprintf("File %s is the same as:\n", k))
			sb.WriteString(k)

			for i := 0; i < len(v); i++ {
				sb.WriteString("," + v[i])

				// Delete duplicates if flag is true
				if rm {
					_ = os.Remove(v[i])
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// shouldBeIgnored checks whether the f1 and f2 are to be ignored when looking for duplicates.
func shouldBeIgnored(f1, f2 string, ignored [][]string) bool {
	for i := 0; i < len(ignored); i++ {
		f1Found, f2Found := false, false
		for j := 0; j < len(ignored[i]); j++ {
			if f1 == ignored[i][j] {
				f1Found = true
			} else if f2 == ignored[i][j] {
				f2Found = true
			}
		}
		if f1Found && f2Found {
			return true
		}
	}
	return false
}
