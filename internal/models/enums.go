package models

import (
	"fmt"
	"mini-game-library/internal/apperror"
)

type Genre string

const (
	GenreAction          Genre = "Action"
	GenreAdventure       Genre = "Adventure"
	GenreActionAdventure Genre = "Action-Adventure"
	GenreShooter         Genre = "Shooter"
	GenreFighting        Genre = "Fighting"
	GenreStealth         Genre = "Stealth"
	GenreRPG             Genre = "RPG"
	GenreActionRPG       Genre = "Action RPG"
	GenreJRPG            Genre = "JRPG"
	GenreMMORPG          Genre = "MMORPG"
	GenreSimulation      Genre = "Simulation"
	GenreSports          Genre = "Sports"
	GenreRacing          Genre = "Racing"
	GenreStrategy        Genre = "Strategy"
	GenrePuzzle          Genre = "Puzzle"
	GenrePlatformer      Genre = "Platformer"
	GenreHorror          Genre = "Horror"
	GenreSurvival        Genre = "Survival"
	GenreSandbox         Genre = "Sandbox"
	GenreRoguelike       Genre = "Roguelike"
	GenreMetroidvania    Genre = "Metroidvania"
	GenreMOBA            Genre = "MOBA"
	GenreCasual          Genre = "Casual"
)

var validGenres = map[Genre]struct{}{
	GenreAction:          {},
	GenreAdventure:       {},
	GenreActionAdventure: {},
	GenreShooter:         {},
	GenreFighting:        {},
	GenreStealth:         {},
	GenreRPG:             {},
	GenreActionRPG:       {},
	GenreJRPG:            {},
	GenreMMORPG:          {},
	GenreSimulation:      {},
	GenreSports:          {},
	GenreRacing:          {},
	GenreStrategy:        {},
	GenrePuzzle:          {},
	GenrePlatformer:      {},
	GenreHorror:          {},
	GenreSurvival:        {},
	GenreSandbox:         {},
	GenreRoguelike:       {},
	GenreMetroidvania:    {},
	GenreMOBA:            {},
	GenreCasual:          {},
}

func (g Genre) IsValid() bool {
	_, ok := validGenres[g]
	return ok
}

func (g Genre) Validate() error {
	if !g.IsValid() {
		return fmt.Errorf("%w: %q", apperror.ErrInvalidGenre, string(g))
	}
	return nil
}

func AllGenres() []Genre {
	genres := make([]Genre, 0, len(validGenres))
	for g := range validGenres {
		genres = append(genres, g)
	}
	return genres
}
