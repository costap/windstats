package service

import (
	"log"
	"time"

	"github.com/costap/windstats/internal/app"
)

// STOP ...
const STOP = 1

// WindStatsService ...
type WindStatsService struct {
	metc          *app.MetClient
	dbc           *app.DBClient
	conf          *app.Config
	lastTried     time.Time
	lastSucceeded time.Time
	running       bool
	c             chan (int)
}

// NewWindStatsService ...
func NewWindStatsService(metc *app.MetClient, dbc *app.DBClient, conf *app.Config) *WindStatsService {
	return &WindStatsService{
		metc:          metc,
		dbc:           dbc,
		conf:          conf,
		lastTried:     time.Unix(0, 0),
		lastSucceeded: time.Unix(0, 0),
		running:       false,
		c:             make(chan int),
	}
}

// Run ...
func (s *WindStatsService) Run() {
	if s.running {
		return
	}
	log.Println("Running...")
	s.running = true
	for {
		select {
		case <-s.c:
			s.running = false
			log.Println("Stopped.")
			return
		default:
			go s.GetData()
		}
		log.Printf("Sleeping for %v seconds...\n", s.conf.RefreshRateSecs)
		time.Sleep(time.Duration(s.conf.RefreshRateSecs) * time.Second)
	}
}

// Stop ...
func (s *WindStatsService) Stop() {
	s.c <- STOP
}

// GetData ...
func (s *WindStatsService) GetData() {
	log.Println("Start get data ... ")
	s.lastTried = time.Now()

	si, err := s.metc.GetSystemInfo()
	if err != nil {
		log.Printf("Error getting system info: %v\n", err)
		return
	}
	log.Printf("System Info: %v\n", si)
	cs, err := s.metc.GetConnectionStatus()
	if err != nil {
		log.Printf("Error getting connections status: %v\n", err)
		return
	}
	if !cs.IsConnected {
		log.Printf("Not connected: %v\n", cs)
		return
	}
	ms, err := s.metc.GetMeasurement()
	if err != nil {
		log.Printf("Error getting measurements: %v\n", err)
		return
	}

	for _, m := range ms {
		if err := s.dbc.Insert("wind",
			map[string]string{"sourceId": "1"},
			map[string]interface{}{"speed": m.Speed, "direction": m.Direction}); err != nil {
			log.Printf("Error writing data %v : %v\n", m, err)
		}
	}

	s.lastSucceeded = time.Now()
	log.Println("Get data done.")
}

// Healthy ...
func (s *WindStatsService) Healthy() bool {
	return !s.running || time.Now().Sub(s.lastTried) <= time.Duration(s.conf.RefreshRateSecs*2)*time.Second
}
