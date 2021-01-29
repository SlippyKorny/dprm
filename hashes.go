package main

import (
	"crypto/sha256"
	"io/ioutil"
)

// genHashes generates hashes for the contents of the files that are pointed by the passed paths
func genHashes(paths []string) ([][32]byte, error) {
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
