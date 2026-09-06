package dto

import (
	"mini-game-library/internal/models"
	"time"

	"github.com/google/uuid"
)

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Timestamp string        `json:"timestamp"`
	Status    int           `json:"status"`
	Error     string        `json:"error"`
	Message   string        `json:"message"`
	Path      string        `json:"path"`
	Details   []ErrorDetail `json:"details,omitempty"`
}

type RegisterResponse struct {
	Id uuid.UUID `json:"id"`
}

func NewRegisterResponse(user *models.User) RegisterResponse {
	return RegisterResponse{
		Id: user.Id,
	}
}

type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type GameResponse struct {
	Id          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Genre       models.Genre `json:"genre"`
	ReleaseYear int          `json:"release_year"`
}

func NewGameResponse(game *models.Game) GameResponse {
	return GameResponse{
		Id:          game.Id,
		Title:       game.Title,
		Description: game.Description,
		Genre:       game.Genre,
		ReleaseYear: game.ReleaseYear,
	}
}

type MetaResponse struct {
	Total int `json:"total"`
}

type GamesResponse struct {
	Data []GameResponse `json:"data"`
	Meta MetaResponse   `json:"meta"`
}

func NewGamesResponse(games []*models.Game, total int) GamesResponse {
	var data []GameResponse
	for _, game := range games {
		data = append(data, NewGameResponse(game))
	}
	if data == nil {
		data = make([]GameResponse, 0)
	}
	return GamesResponse{
		Data: data,
		Meta: MetaResponse{
			Total: total,
		},
	}
}
