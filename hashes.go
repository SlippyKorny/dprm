package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"sync"

	"github.com/vitali-fedulov/images"
)

// genContentHashes generates hashes based on the raw content of the given files
func genContentHashes(paths []string, verbose bool) ([][32]byte, error) {
	// Set up concurrency vars
	var wg sync.WaitGroup
	var err error
	threads := runtime.NumCPU()
	if threads > len(paths) {
		threads = len(paths)
	}
	wg.Add(threads)

	statuses := make([]bool, threads)
	interval := len(paths) / threads

	// Declare the hash slice
	hashes := make([][32]byte, len(paths))

	// Declare the concurrent closure
	hashGenFunc := func(index, start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			dat, err := ioutil.ReadFile(paths[i])
			if err != nil {
				return
			}

			hashes[i] = sha256.Sum256(dat)
		}

		statuses[index] = true
	}

	// Execute the hash generation concurrently
	i := 0
	for ; i < threads-1; i++ {
		go hashGenFunc(i, i*interval, i+1*interval)
	}
	go hashGenFunc(i, i*interval, len(paths))

	// Await the execution and check if the hashes were generated correctly
	wg.Wait()
	for i = 0; i < len(statuses); i++ {
		if statuses[i] == false {
			err = errors.New("an unexpected error occurred while generating hashes")
			break
		}
	}

	return hashes, err
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
