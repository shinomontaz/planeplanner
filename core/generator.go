package main

import (
	"planeplanner/types"
	"time"
)

func NewDay() *types.Day {
	d := &types.Day{}
	for i := 0; i < 24*60/5; i++ {
		d.Quants[i] = struct{}{}
	}

	return d
}

func NewRandomPlane(day time.Time) *types.Plane {
	p := &types.Plane{}

	return p
}
