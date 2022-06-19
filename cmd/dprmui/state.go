package main

var StateSingleton uiState

func init() {
	StateSingleton = uiState{
		Format:    "csv",
		Method:    "hashes",
		Directory: ".",
		Recursive: false,
		Remove:    false,
	}
}

type uiState struct {
	Format    string
	Method    string
	Directory string
	Recursive bool
	Remove    bool
}
