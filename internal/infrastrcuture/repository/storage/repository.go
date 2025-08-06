package storage

import "time"

type Map struct {
	detailsMap map[uint64]Details
}

func NewMapRepository() *Map {
	return &Map{
		detailsMap: map[uint64]Details{
			1: {ID: 1, Description: "Sample description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			2: {ID: 2, Description: "Another description", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
}

type Slice struct {
}

func NewSliceRepository() *Slice {
	return &Slice{}

}
