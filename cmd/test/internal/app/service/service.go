package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Daysleft() int64 {
	d := time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	return int64(dur.Hours()) / 24
}
