package main

import (
	"fmt"
	"strings"
)

var StateSingleton UIState

func init() {
	StateSingleton = UIState{
		Format:    "csv",
		Method:    "",
		Directory: "",
		Recursive: false,
		Remove:    false,
	}
}

type UIState struct {
	Format    string
	Method    string
	Directory string
	Recursive bool
	Remove    bool
}

func (state UIState) IsValid() (bool, string) {
	var sb strings.Builder
	if state.Method != "perceptual" && state.Method != "hashes" {
		sb.WriteString("Method needs to either 'perceptual' or 'hashes'. Please select one!\n")
	}
	if state.Directory == "" {
		sb.WriteString("Directory was not selected!\n")
	}

	fmt.Println(sb.String())
	return sb.Len() == 0, sb.String()
}
