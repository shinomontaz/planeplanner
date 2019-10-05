package main

import (
	"planeplanner/types"

	"github.com/shinomontaz/ga"
)

type Schedule struct {
	jobs     []*types.Job
	planes   []*types.Plane
	brigades []*types.Brigade
	List     map[*types.Plane]map[*types.Job]*types.Employee //
	//	Table map[Employee]map[time.Time]*Job // расписание с точки зреня сотрудников
}

func algorithm(planes []*types.Plane, jobs []*types.Job, brigades []*types.Brigade) *Schedule {
	res := &Schedule{}

	return res
}

type ScheduleFactory struct {
	jobs     []*types.Job
	planes   []*types.Plane
	brigades []*types.Brigade
}

func (sf *ScheduleFactory) Create() ga.Individual {
	// собираем случайную особь так, чтобы уложиться в лимиты в конфиге
	//	indexes := rand.Perm(len(sf.jobs))
	schedule := &Schedule{
		jobs:     sf.jobs,
		planes:   sf.planes,
		brigades: sf.brigades,
	}

	return schedule
}

func (s *Schedule) Clone() ga.Individual {
	return s
}

func (s *Schedule) Crossover(p ga.Individual) ga.Individual {
	return s
}

func (s *Schedule) Educate() {

}

func (s *Schedule) Fitness() float64 {
	return 0.0
}

func (s *Schedule) Mutate() ga.Individual {
	return s
}
