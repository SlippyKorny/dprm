package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"

	"github.com/vitali-fedulov/images"
)

// genContentHashes generates hashes based on the raw content of the given files
func genContentHashes(paths []string, verbose bool) ([][32]byte, error) {
	hashes := make([][32]byte, len(paths))
	pathLen := len(paths)

	for i := 0; i < pathLen; i++ {
		dat, err := ioutil.ReadFile(paths[i])
		if err != nil {
			return nil, err
		}

		hashes[i] = sha256.Sum256(dat)
	}

	return hashes, nil
}

// genContentHashes generates hashes based on a slice of average color values of an image at the
// position of white pixels of a mask. One average value corresponds to one mask
func genAvgColourHashes(paths []string, verbose bool) ([][]float32, error) {
	hashes := make([][]float32, len(paths))
	pathLen := len(paths)

	for i := 0; i < pathLen; i++ {
		// Output the currently generated hash
		if verbose {
			fmt.Printf("\r\rGenerating hash for '%s'", paths[i])
		}

		img, err := images.Open(paths[i])
		if err != nil {
			return nil, err
		}

		hashes[i], _ = images.Hash(img)
	}

	fmt.Println()
	return hashes, nil
}
