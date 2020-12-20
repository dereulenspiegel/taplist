package filestore

import (
	"fmt"
	"sync"

	"github.com/dereulenspiegel/taplist/graph/model"
	"github.com/jinzhu/copier"
)

type Store struct {
	lock *sync.Mutex

	kegerator model.Kegerator
}

func New(kegeratorName string, numberOfTaps int) *Store {
	taps := make([]*model.Tap, numberOfTaps)

	for i := 0; i < numberOfTaps; i++ {
		taps[i] = &model.Tap{
			Number: i + 1,
			ID:     fmt.Sprintf("%d", i),
			Empty:  true,
		}
	}
	return &Store{
		kegerator: model.Kegerator{
			Name: &kegeratorName,
			Taps: taps,
		},
		lock: &sync.Mutex{},
	}
}

func (s *Store) Taps() ([]*model.Tap, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	taps := []*model.Tap{}
	copier.Copy(&taps, &s.kegerator.Taps)
	return taps, nil
}

func (s *Store) Kegerator() (*model.Kegerator, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	kegerator := &model.Kegerator{}
	copier.Copy(&kegerator, &s.kegerator)
	return kegerator, nil
}
