package main

import (
	"math/rand"
	"planeplanner/types"
	"time"
)

func NewRandomPlane(day time.Time, availableJobs []*types.Job) *types.Plane {
	p := &types.Plane{}

	randTime := rand.Intn(24 * 60 / 5) // c точностью до 5 минут определяем

	p.Arrival = day.Add(time.Duration(randTime*60) * time.Second)

	randJobNum := rand.Intn(len(availableJobs))
	randJobPerm := rand.Perm(len(availableJobs))

	for i := 0; i < randJobNum; i++ {
		p.Jobs = append(p.Jobs, availableJobs[randJobPerm[i]])
	}

	return p
}
