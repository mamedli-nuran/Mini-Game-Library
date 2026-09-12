package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type CreateGameRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}

type UpdateGameRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Genre       *string `json:"genre,omitempty"`
	ReleaseYear *int    `json:"release_year,omitempty"`
}

type AddToLibraryRequest struct {
	GameId string `json:"game_id"`
	Status string `json:"status"`
}

type UpdateLibraryStatusRequest struct {
	Status string `json:"status"`
}
