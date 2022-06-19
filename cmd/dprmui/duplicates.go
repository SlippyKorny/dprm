package main

import (
	"math/rand"
	"time"

	"github.com/lxn/walk"
)

type DuplicatesModel struct {
	walk.SortedReflectTableModelBase
	items []*Duplicate
}

func (m *DuplicatesModel) Items() interface{} {
	return m.items
}

type Duplicate struct {
	Index  int
	Name   string
	Format string
	Size   float64
}

func NewEmptyDuplicatesModel() *DuplicatesModel {
	now := time.Now()

	rand.Seed(now.UnixNano())

	m := &DuplicatesModel{items: make([]*Duplicate, 1)}

	for i := range m.items {
		m.items[i] = &Duplicate{
			Index:  i,
			Name:   "",
			Format: "",
			Size:   0,
		}
	}

	return m
}
