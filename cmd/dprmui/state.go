package main

import (
	"fmt"
	"strings"

	"github.com/lxn/walk"
)

var StateSingleton State

func init() {

	StateSingleton = State{
		Form: FormState{
			Format:    "csv",
			Method:    "",
			Directory: "",
			Recursive: false,
			Remove:    false,
		},
	}
}

type State struct {
	Form FormState
	UI   UIState
}

type FormState struct {
	Format    string
	Method    string
	Directory string
	Recursive bool
	Remove    bool
}

func (state FormState) IsValid() (bool, string) {
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

type UIState struct {
	duplicatesTable *walk.TableView
}
