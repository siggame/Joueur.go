package jungle

import "joueur/base"

// Game is a 7x9 board game with pieces.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Fen is forsyth-Edwards Notation (fen), a notation that
	// describes the game board state.
	Fen() string

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// History is the array of [known] moves that have occurred in the
	// game, in a format. The first element is the first move, with
	// the last element being the most recent.
	History() []string

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string
}
