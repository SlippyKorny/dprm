package dprm

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"runtime"
	"sync"

	"github.com/vitali-fedulov/images"
)

// genContentHashes generates hashes based on the raw content of the given files
func genContentHashes(paths []string) ([][32]byte, error) {
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
			dat, er := ioutil.ReadFile(paths[i])
			if er != nil {
				fmt.Fprintf(os.Stderr, "'%s': %s\n", paths[i], er.Error())
				err = er
				return
			} else if err != nil {
				return
			}

			hashes[i] = sha256.Sum256(dat)
		}

		statuses[index] = true
	}

	// Execute the hash generation concurrently
	i := 0
	for ; i < threads-1; i++ {
		go hashGenFunc(i, i*interval, (i+1)*interval)
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
func genAvgColourHashes(paths []string) ([][]float32, []image.Point, error) {
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
	hashes := make([][]float32, len(paths))
	sizes := make([]image.Point, len(paths))

	// Declare the concurrent closure
	hashGenFunc := func(index, start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			img, er := images.Open(paths[i])
			if er != nil {
				fmt.Fprintf(os.Stderr, "'%s': %s\n", paths[i], er.Error())
				err = er
				return
			} else if err != nil {
				return
			}

			hashes[i], sizes[i] = images.Hash(img)
		}

		statuses[index] = true
	}

	// Execute the hash generation concurrently
	i := 0
	for ; i < threads-1; i++ {
		go hashGenFunc(i, i*interval, (i+1)*interval)
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

	return hashes, sizes, err
}
