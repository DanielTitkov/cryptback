package domain

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Challenge string    `json:"challenge"`
	KeyLength int8      `json:"keyLength"`
}

type Key struct {
	GameID  uuid.UUID `json:"gameId"`
	Key     string    `json:"key"`
	Success bool      `json:"success"`
}

type User struct {
	ID        uuid.UUID
	IPAddress string
}

// system types
type (
	SystemSummary struct {
		ID            int
		Users         int
		Tests         int
		FinishedTakes int
		Responses     int
		CreateTime    time.Time
	}
	Event struct {
		Name      string
		StartTime time.Time
		EndTime   time.Time
		Elapsed   time.Duration
	}
)

func (k *Key) CheckKey(candidate string) (string, error) {
	return "", nil
}
