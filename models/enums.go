package models

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
