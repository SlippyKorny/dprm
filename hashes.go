package main

import (
	"crypto/sha256"
	"io/ioutil"
)

// genContentHashes generates hashes based on the raw content of the given files
func genContentHashes(paths []string) ([][32]byte, error) {
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
func genAvgColourHashes(paths []string) ([][]float32, error) {

}
