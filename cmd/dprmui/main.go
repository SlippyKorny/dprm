// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	type Mode struct {
		Name  string
		Value ImageViewMode
	}
	var widgets []Widget

	widgets = append(widgets,
		Label{
			Text: "ImageViewModeShrink",
		},
		ImageView{
			Image:  "../img/image.jpg",
			Margin: 10,
			Mode:   ImageViewModeShrink,
		},
	)

	MainWindow{
		Title:    "Walk ImageView Example",
		Size:     Size{400, 600},
		Layout:   Grid{Columns: 2},
		Children: widgets,
	}.Run()
}
