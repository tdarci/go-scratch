package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tdarci/prj-999/repos"

	"github.com/tdarci/prj-999/models"

	"github.com/tdarci/prj-999/config"
)

const (
	usgsURL = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/1.0_hour.geojson"
)

type QuakePoller struct {
	*config.Config
	repo *repos.Quake
}

func NewQuakePoller(cfg *config.Config) *QuakePoller {

	return &QuakePoller{
		Config: cfg,
		repo:   repos.NewQuake(),
	}
}

// Run starts the poller polling from USGS at the given duration.
func (q *QuakePoller) Run(interval time.Duration) {

	ticker := time.NewTicker(interval)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := q.poll()
				if err != nil {
					q.Logger().Fatalf("Error while polling! %s", err)
				}
			}
		}
	}()

}
func (q *QuakePoller) poll() error {

	q.Logger().Println("==== POLLING =====")

	data, err := q.fetch()
	if err != nil {
		return err
	}

	for _, quake := range data.Features {
		q.Logger().Printf("Received Quake: %s", quake.ID)
		if existing := q.repo.Get(quake.ID); existing == nil {
			// this is a NEW quake
			q.Logger().Printf("Found NEW Quake: %s", quake.ID)

			// save it
			q.repo.Add(&quake)

			// notify our endpoint
			q.notify(&quake)
		}
	}

	// store earthquakes
	// (see which are new)

	return nil
}

func (q *QuakePoller) notify(quake *models.Earthquake) {
	// TODO: implement!
}

func (q *QuakePoller) fetch() (*models.USGSResponse, error) {

	resp, err := http.Get(usgsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	usgsData := models.USGSResponse{}
	err = json.Unmarshal(b, &usgsData)
	if err != nil {
		return nil, err
	}

	return &usgsData, nil
}
