package main

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	magicArrTypes = [4]string{"image/jpeg", "image/png", "image/gif", "image/gif"}
	magicArrVals  = [4]string{"\xff\xd8\xff", "\x89PNG\r\n\x1a\n", "GIF87a", "GIF89a"}
)

// fIsIMG checks whether the file in the given relative path is a jpeg, png or gif file
func fIsIMG(path string) (int, error) {
	// TODO: Optimize by reading just a bit of the file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return -1, err
	}

	fStr := string(f)
	for i := 0; i < len(magicArrVals); i++ {
		if strings.HasPrefix(fStr, magicArrVals[i]) {
			return i, nil
		}
	}

	return -1, nil
}

// extrFilenames extracts all of the file names in a directory recursively
func extrFilenames(path string, recursive bool) ([]string, error) {
	var files []string

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fs, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, file := range fs {
		// If is directory then recursively call this function
		isdir := file.IsDir()
		if isdir && recursive {
			dir := path + file.Name()
			if !strings.HasSuffix(dir, "/") {
				dir = dir + "/"
			}

			fsr, err := extrFilenames(dir, recursive)
			if err != nil {
				return nil, err
			}

			files = append(files, fsr...)
			continue
		} else if isdir {
			continue
		}

		// Check if the file is an image and if it is then add it to the files slice
		ffn := path + file.Name()
		c, err := fIsIMG(ffn)
		if c != -1 && err == nil {
			files = append(files, ffn)
		} else if err != nil {
			return nil, err
		}
	}

	return files, nil
}

// rmFiles removes duplicate files and returns the list of removed files
func rmFiles(files []string) (rmd []string, err error) {
	// TODO: Implement this

	return nil, nil
}
