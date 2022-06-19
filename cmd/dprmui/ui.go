package main

import (
	"fmt"
	"os"

	"github.com/TheSlipper/dprm/dprm"
	. "github.com/lxn/walk/declarative"
	"github.com/sqweek/dialog"
)

func BuildMainWindow() MainWindow {
	// walk.NewFlowLayout().LayoutBase
	return MainWindow{
		Title:  "dprm - duplicate removal tool",
		Size:   Size{400, 600},
		Layout: Grid{Columns: 4},
		Children: []Widget{
			PushButton{
				Text: "Select Directory",
				OnClicked: func() {
					dir, err := dialog.Directory().Title("Select target directory").Browse()
					if err != nil && err.Error() != "Cancelled" {
						dialog.Message("error while selecting directory: " + err.Error()).Error()
						return
					}
					dir = dir + string(os.PathSeparator)
					StateSingleton.Form.Directory = dir
				},
			},
			CheckBox{
				Name:             "recursive",
				Text:             "Recursive traversal",
				Checked:          false,
				OnCheckedChanged: func() { StateSingleton.Form.Recursive = !StateSingleton.Form.Recursive },
			},
			CheckBox{
				Name:    "removal",
				Text:    "Remove duplicates automatically",
				Checked: false,
				OnCheckedChanged: func() {
					dialog.Message("This will irreversibely delete your files without warning! If you are not certain, uncheck \"removal\" checkbox!").Info()
					StateSingleton.Form.Recursive = !StateSingleton.Form.Recursive
				},
			},
			RadioButtonGroup{
				DataMember: "Method",
				Buttons: []RadioButton{
					RadioButton{
						Name:      "method-hashes",
						Text:      "SHA-256 Hashes",
						Value:     "hashes",
						OnClicked: func() { StateSingleton.Form.Method = "hashes" },
					},
					RadioButton{
						Name:      "method-perceptual",
						Text:      "Perceptual image similarity",
						Value:     "perceptual",
						OnClicked: func() { StateSingleton.Form.Method = "perceptual" },
					},
				},
			},
			PushButton{
				Text: "Find duplicates",
				OnClicked: func() {
					if res, msg := StateSingleton.Form.IsValid(); !res {
						dialog.Message(msg).Error()
						return
					}

					str := dprm.Run(StateSingleton.Form.Format, StateSingleton.Form.Method,
						StateSingleton.Form.Directory, StateSingleton.Form.Recursive, StateSingleton.Form.Remove)
					fmt.Println(str)
					err := StateSingleton.UI.duplicatesTable.SetModel(NewEmptyDuplicatesModel())
					if err != nil {
						dialog.Message(err.Error()).Error()
					}
				},
			},
			TableView{
				AssignTo:         &StateSingleton.UI.duplicatesTable,
				Name:             "duplicate-view", // Name is needed for settings persistence
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				Columns: []TableViewColumn{
					// Name is needed for settings persistence
					{Name: "#", DataMember: "Index"}, // Use DataMember, if names differ
					{DataMember: "Name", Name: "Name"},
					{DataMember: "Format", Name: "Format"},
					{DataMember: "Size", Name: "Size", Format: "%.2fMB"},
				},
				Model: NewEmptyDuplicatesModel(),
			},
		},
	}
}
