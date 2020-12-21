package filestore

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/dereulenspiegel/taplist/graph/model"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type Store struct {
	lock         *sync.Mutex
	dataFilePath string

	kegerator model.Kegerator
}

func newTap(i int) *model.Tap {
	return &model.Tap{
		Number: i + 1,
		ID:     fmt.Sprintf("tap:%d", i),
		Empty:  true,
		Name:   fmt.Sprintf("Tap %d", i+1),
	}
}

func New(dataFilePath, kegeratorName string, numberOfTaps int) *Store {

	s := &Store{
		dataFilePath: dataFilePath,
		lock:         &sync.Mutex{},
	}
	if err := s.load(); err != nil {
		logrus.WithField("dataPath", s.dataFilePath).WithError(err).Fatal("Failed to load persisted state")
	}

	s.kegerator.Name = &kegeratorName
	tapDiff := numberOfTaps - len(s.kegerator.Taps)
	if tapDiff > 0 {
		for i := 0; i < tapDiff; i++ {
			tapNum := i + len(s.kegerator.Taps)
			s.kegerator.Taps = append(s.kegerator.Taps, newTap(tapNum))
		}
	}

	if tapDiff < 0 {
		logrus.WithFields(logrus.Fields{
			"oldTapCount": len(s.kegerator.Taps),
			"newTapCount": numberOfTaps,
		}).Warn("Number of taps was decreased, removing all tap data, because I can't determine which taps are left")
		s.kegerator.Taps = make([]*model.Tap, numberOfTaps)
		for i := 0; i < numberOfTaps; i++ {
			s.kegerator.Taps[i] = newTap(i)
		}
	}

	return s
}

func (s *Store) load() error {
	dataFile, err := os.Open(s.dataFilePath)
	if err != nil && os.IsNotExist(err) {
		logrus.WithField("dataPath", s.dataFilePath).Info("Data file not found, starting with persisted data")
		return nil
	} else if err != nil {
		return fmt.Errorf("Failed to open data file %s: %w", s.dataFilePath, err)
	}
	defer dataFile.Close()
	if err := json.NewDecoder(dataFile).Decode(&s.kegerator); err != nil {
		return fmt.Errorf("Failed to unmarshal data file %s: %w", s.dataFilePath, err)
	}
	return nil
}

func (s *Store) save() error {
	dataFile, err := os.Create(s.dataFilePath)
	if err != nil {
		return fmt.Errorf("Failed to open data file %s: %w", s.dataFilePath, err)
	}
	defer dataFile.Close()

	if err := json.NewEncoder(dataFile).Encode(s.kegerator); err != nil {
		return fmt.Errorf("Failed to marshal state to %s: %w", s.dataFilePath, err)
	}
	if err := dataFile.Sync(); err != nil {
		return fmt.Errorf("Failed to sync file %s: %w", s.dataFilePath, err)
	}
	return nil
}

func (s *Store) Taps() ([]*model.Tap, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	taps := []*model.Tap{}
	copier.Copy(&taps, &s.kegerator.Taps)
	return taps, nil
}

func (s *Store) Tap(id string) (*model.Tap, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var st *model.Tap
	for _, t := range s.kegerator.Taps {
		if t.ID == id {
			st = t
			break
		}
	}
	if st != nil {
		tap := model.Tap{}
		copier.Copy(&tap, &st)
		return &tap, nil
	}
	return nil, fmt.Errorf("Tap %s not found", id)
}

func (s *Store) UpdateTap(tap *model.Tap) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	updated := false
	for i, t := range s.kegerator.Taps {
		if t.ID == tap.ID {
			s.kegerator.Taps[i] = tap
			updated = true
			break
		}
	}
	if updated {
		if err := s.save(); err != nil {
			return fmt.Errorf("Failed to update tap: %w", err)
		}
	} else {
		return fmt.Errorf("Tap %s not found", tap.ID)
	}
	return nil
}

func (s *Store) Kegerator() (*model.Kegerator, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	kegerator := &model.Kegerator{}
	copier.Copy(&kegerator, &s.kegerator)
	return kegerator, nil
}
