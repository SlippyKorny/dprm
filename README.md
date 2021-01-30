# dprm

dprm (duplicate remover) is a concurrent, simple and lightweight commandline utility for finding and removing duplicate images in a given directory. The program is still in early stage of development however it already provides a very useful duplication detection and removal functionalities in a non-bloated form.

## Quick usage guide

I have created this utility as I could not find a lightweight, terminal based utility for finding and removing duplicate images from my server storage. As of the moment two distinct duplicate detection methods were introduced. Each one of them is meant to be used for different use cases so it is highly recommended for all users to read this small guide before using the program.

### Content based hash comparison

This method uses SHA256 cryptographic hash function to generate a 32 byte hash for each of the images. Best use case for this method is finding exact matches of images.

### Perceptual image similarity

This method uses an implementation of a perceptual image similarity algorithm made by [Vitali Fedulov](https://github.com/vitali-fedulov/images). This method works great with:
1. Finding similar images of different compression rate
2. Finding images with the same features and little differences

**WARNING:** This method may detect images with and without captions as the same duplicates which may not be desired in some use cases.

## Usage

```
dprm version 0.0.2
Copyright (C) 2021 by Kornel Domeradzki
Source: http://github.com/TheSlipper/dprm

dprm is a simple commandline hash based duplicate image search and removal tool.

dprm comes with ABSOLUTELY NO WARRANTY.  This is free software, and you
are welcome to redistribute it under certain conditions.  See the GNU General
Public Licence for details.

Usage: dprm [OPTION...] [DIRECTORY]

--method string
		specifies the method with which the duplicates are searched for.
		Available methods are 'hashes' (default) and 'perceptual'
--remove
		if set to true will remove the duplicates autonomously
--recursive
		if set to true will recursively traverse the folder tree
--verbose
		verbosity of the command's execution. If remove argument is not
		set to true then the program will set verbose to true.
--help
		prints out this help section
```

## Setup and Compilation

In the [release section](https://github.com/TheSlipper/dprm/releases) you will find the builds for x86 Linux, Windows and macOS systems. All you need to do is download it and add the utility to your path. If you need to use the utility on a different architecture then you will need to compile it from source with the instructions available bellow.   

In order to set up this utility you need to have go installed. The process varies for every operating system therefore that part is not included in the instructions. Once you install it, run the commands listed below:

```
go get github.com/TheSlipper/dprm
cd $GOPATH/src/github.com/TheSlipper/dprm/
go install
```
You can check if your installation was successful by running `dprm --help`.

## Licence

dprm comes with ABSOLUTELY NO WARRANTY.  This is free software, and you
are welcome to redistribute it under certain conditions.  See the GNU General Public Licence
for details.
