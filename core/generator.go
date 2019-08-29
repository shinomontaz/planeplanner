package main

import (
	"fmt"
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

func NewRandomJobs(max int, bg *types.Brigade) []*types.Job {
	randNumJobs := rand.Intn(max)
	list := make([]*types.Job, randNumJobs)
	for i := 0; i < randNumJobs; i++ {
		list = append(list, &types.Job{
			Id:      bg.Id + i + 1,
			Brigade: bg,
			Time:    3000 + rand.Intn(9000),
			Count:   rand.Intn(len(bg.Team)),
		})
	}

	for _, job := range list {
		if rand.Float64() < 0.4 {
			randJob := list[rand.Intn(len(list))]
			if randJob.Id != job.Id {
				job.Parent = randJob
			}
		}
	}

	return list
}

func NewBrigade(id int, maxTeam int) *types.Brigade {
	brigadeLen := 1 + rand.Intn(maxTeam-1)
	team := make([]*types.Employee, 0)

	for i := 0; i < brigadeLen; i++ {
		team = append(team, NewEmployee(id*100+i))
	}

	return &types.Brigade{
		Id:   id,
		Name: fmt.Sprintf("%d:%d", id, rand.Int()),
		Team: team,
	}
}

func NewEmployee(id int) *types.Employee {
	jobStart := rand.Intn(24 * 2)
	jobEnd := jobStart + 8*2

	if jobEnd > 24*2 {
		jobEnd = 24 * 2
	}

	return &types.Employee{
		Id:    id,
		Start: time.Unix(int64(jobStart*30*60), 0), // начало работы с шагом по полчаса
		End:   time.Unix(int64(jobEnd*30*60), 0),
	}
}
