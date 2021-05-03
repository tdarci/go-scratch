package repos

import "github.com/tdarci/prj-999/models"

type Quake struct {
	quakes map[string]*models.Earthquake
}

func NewQuake() *Quake {
	return &Quake{
		quakes: make(map[string]*models.Earthquake, 0),
	}
}

func (q *Quake) Get(id string) *models.Earthquake {
	if quake, found := q.quakes[id]; !found {
		return nil
	} else {
		return quake
	}
}

func (q *Quake) Add(quake *models.Earthquake) {
	q.quakes[quake.ID] = quake
}
