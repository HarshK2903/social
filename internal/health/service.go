package health

import "time"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type Status struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *Service) Check() Status {
	return Status{
		Status:    "ok",
		Message:   "social API is running",
		Timestamp: time.Now(),
	}
}
