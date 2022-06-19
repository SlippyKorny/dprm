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
					StateSingleton.Directory = dir
				},
			},
			CheckBox{
				Name:             "recursive",
				Text:             "Recursive traversal",
				Checked:          false,
				OnCheckedChanged: func() { StateSingleton.Recursive = !StateSingleton.Recursive },
			},
			CheckBox{
				Name:    "removal",
				Text:    "Remove duplicates automatically",
				Checked: false,
				OnCheckedChanged: func() {
					dialog.Message("This will irreversibely delete your files without warning! If you are not certain, uncheck \"removal\" checkbox!").Info()
					StateSingleton.Recursive = !StateSingleton.Recursive
				},
			},
			RadioButtonGroup{
				DataMember: "Method",
				Buttons: []RadioButton{
					RadioButton{
						Name:      "method-hashes",
						Text:      "SHA-256 Hashes",
						Value:     "hashes",
						OnClicked: func() { StateSingleton.Method = "hashes" },
					},
					RadioButton{
						Name:      "method-perceptual",
						Text:      "Perceptual image similarity",
						Value:     "perceptual",
						OnClicked: func() { StateSingleton.Method = "perceptual" },
					},
				},
			},
			PushButton{
				Text: "Find duplicates",
				OnClicked: func() {
					str := dprm.Run(StateSingleton.Format, StateSingleton.Method,
						StateSingleton.Directory, StateSingleton.Recursive, StateSingleton.Remove)
					fmt.Println(str)
				},
			},
			// TableView{
			// 	Name:             "tableView", // Name is needed for settings persistence
			// 	AlternatingRowBG: true,
			// 	ColumnsOrderable: true,
			// 	Columns: []TableViewColumn{
			// 		// Name is needed for settings persistence
			// 		{Name: "#", DataMember: "Index"}, // Use DataMember, if names differ
			// 		{Name: "Bar"},
			// 		{Name: "Baz", Format: "%.2f", Alignment: AlignFar},
			// 		{Name: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
			// 	},
			// 	Model: NewFooModel(),
			// },
		},
	}
}

// type FooModel struct {
// 	walk.SortedReflectTableModelBase
// 	items []*Foo
// }

// func (m *FooModel) Items() interface{} {
// 	return m.items
// }

// type Foo struct {
// 	Index int
// 	Bar   string
// 	Baz   float64
// 	Quux  time.Time
// }

// func NewFooModel() *FooModel {
// 	now := time.Now()

// 	rand.Seed(now.UnixNano())

// 	m := &FooModel{items: make([]*Foo, 1000)}

// 	for i := range m.items {
// 		m.items[i] = &Foo{
// 			Index: i,
// 			Bar:   strings.Repeat("*", rand.Intn(5)+1),
// 			Baz:   rand.Float64() * 1000,
// 			Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
// 		}
// 	}

// 	return m
// }
