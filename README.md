# dprm

dprm (duplicate remover) is a simple and lightweight commandline utility for finding and removing duplicate images in a given directory.


## About the project

I have created this utility as I could not find a lightweight, terminal based utility for finding and removing duplicate images from my storage. This project is still in early stage of development and therefore some of the features might not work correctly or are not yet implemented. As of the moment the only implemented image comparison method is using hashes derived from image content which makes the program detect only the exact same images. In other words those two images below will be detected as different images despite only having a varying compression rate:

<img align="left" src="https://raw.githubusercontent.com/TheSlipper/dprm/main/img/compr_1.jpg?token=AGZOOL7WY5VRSANH34NEJO3AB3O7Y" width="45%">
<img align="right" src="https://raw.githubusercontent.com/TheSlipper/dprm/main/img/compr_2.jpg?token=AGZOOL42BOL2GQXUUCZYLOTAB3PAO" width="45%">
<div style="clear: both;"></div>

(Art source: [click](https://twitter.com/lezon_re/status/1352567928109993984?s=20))

I will be however trying to find a simple solution for detecting duplicate images of varying compression levels in the future.

## Usage

```
dprm [-dir STRING] [-print | -remove] [-recursive] [-verbose]

-dir string
	defines the directory of operation
-print
		prints out the duplicates
-remove
		if set to true will remove the duplicates autonomously  
-recursive
		if set to true will recursively traverse the folder tree
-verbose
		verbosity of the command's execution
-help
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
