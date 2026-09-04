package models

import (
	"fmt"

	"mini-game-library/internal/apperror"
)

type Genre string

const (
	GenreAction          Genre = "ACTION"
	GenreAdventure       Genre = "ADVENTURE"
	GenreActionAdventure Genre = "ACTION_ADVENTURE"
	GenreShooter         Genre = "SHOOTER"
	GenreFighting        Genre = "FIGHTING"
	GenreStealth         Genre = "STEALTH"
	GenreRPG             Genre = "RPG"
	GenreActionRPG       Genre = "ACTION_RPG"
	GenreJRPG            Genre = "JRPG"
	GenreMMORPG          Genre = "MMORPG"
	GenreSimulation      Genre = "SIMULATION"
	GenreSports          Genre = "SPORTS"
	GenreRacing          Genre = "RACING"
	GenreStrategy        Genre = "STRATEGY"
	GenrePuzzle          Genre = "PUZZLE"
	GenrePlatformer      Genre = "PLATFORMER"
	GenreHorror          Genre = "HORROR"
	GenreSurvival        Genre = "SURVIVAL"
	GenreSandbox         Genre = "SANDBOX"
	GenreRoguelike       Genre = "ROGUELIKE"
	GenreMetroidvania    Genre = "METROIDVANIA"
	GenreMOBA            Genre = "MOBA"
	GenreCasual          Genre = "CASUAL"
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
