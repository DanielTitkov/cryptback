package api

import "github.com/google/uuid"

type GetGameResponse struct {
	ID        uuid.UUID `json:"id"`
	Challenge string    `json:"challenge"`
	KeyLength int8      `json:"keyLength"`
}

type PostKeyRequest struct {
	GameID uuid.UUID `json:"gameId"`
	Key    string    `json:"key"`
}

type PostKeyResponse struct {
	GameID  uuid.UUID `json:"gameId"`
	Mask    string    `json:"mask"`
	Success bool      `json:"success"`
}
