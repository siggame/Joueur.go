package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Stumped.
type PlayerImpl struct {
	GameObjectImpl

	beaversImpl              []stumped.Beaver
	branchesToBuildLodgeImpl int64
	clientTypeImpl           string
	lodgesImpl               []stumped.Tile
	lostImpl                 bool
	nameImpl                 string
	opponentImpl             stumped.Player
	reasonLostImpl           string
	reasonWonImpl            string
	timeRemainingImpl        float64
	wonImpl                  bool
}

// Beavers returns the array of Beavers owned by this Player.
func (playerImpl *PlayerImpl) Beavers() []stumped.Beaver {
	return playerImpl.beaversImpl
}

// BranchesToBuildLodge returns how many branches are required to build a
// lodge for this Player.
func (playerImpl *PlayerImpl) BranchesToBuildLodge() int64 {
	return playerImpl.branchesToBuildLodgeImpl
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Lodges returns an array of Tiles that contain lodges owned by this
// player.
func (playerImpl *PlayerImpl) Lodges() []stumped.Tile {
	return playerImpl.lodgesImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() stumped.Player {
	return playerImpl.opponentImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.beaversImpl = []stumped.Beaver{}
	playerImpl.branchesToBuildLodgeImpl = 0
	playerImpl.clientTypeImpl = ""
	playerImpl.lodgesImpl = []stumped.Tile{}
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.wonImpl = true
}

// DeltaMerge merges the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := playerImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	stumpedDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stumped.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "beavers":
		playerImpl.beaversImpl = stumpedDeltaMerge.ArrayOfBeaver(&playerImpl.beaversImpl, delta)
		return true, nil
	case "branchesToBuildLodge":
		playerImpl.branchesToBuildLodgeImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "lodges":
		playerImpl.lodgesImpl = stumpedDeltaMerge.ArrayOfTile(&playerImpl.lodgesImpl, delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = stumpedDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = stumpedDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = stumpedDeltaMerge.Float(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = stumpedDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
