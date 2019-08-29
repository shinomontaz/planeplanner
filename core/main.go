package main

import (
	"fmt"
	"planeplanner/types"
	"time"
)

func main() {

	numBrigades := 10
	numPlanes := 50
	maxTeam := 10
	maxJobs := 20

	currentDay, err := time.Parse("01-02-2006", time.Now().Format("01-02-2006"))
	if err != nil {
		panic(err)
	}

	brigades := make([]*types.Brigade, numBrigades)
	jobs := make([]*types.Job, numBrigades)
	planes := make([]*types.Plane, numPlanes)
	for i := 0; i < numBrigades; i++ {
		brigade := NewBrigade(i, maxTeam)
		jobs = append(jobs, NewRandomJobs(maxJobs, brigade)...)

		brigades = append(brigades, brigade)
	}

	for i := 0; i < numBrigades; i++ {
		brigades = append(brigades, NewBrigade(i, maxTeam))
	}

	for i := 0; i < numPlanes; i++ {
		planes = append(planes, NewRandomPlane(currentDay, jobs))
	}

	fmt.Println(brigades)
	fmt.Println(jobs)
	fmt.Println(planes)
}
