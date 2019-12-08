package games

import (
	"errors"
	"joueur/base"
)

type GameNamespace interface {
	Name() string
	Version() string
	PlayerName() string
	CreateAI() (base.AI, *base.AIImpl)
	CreateDeltaMerge(base.DeltaMergeImpl) base.DeltaMerge
	CreateGame() (base.Game, *base.GameImpl)
	CreateGameObject(string) (base.GameObject, *base.GameObjectImpl, error)
	OrderAI(base.AI, string, []interface{}) (interface{}, error)
}

var gamesNamespaceTypes = make(map[string]GameNamespace)

func Register(gameName string, gameNamespace GameNamespace) {
	gamesNamespaceTypes[gameName] = gameNamespace
}

func Get(gameName string) (GameNamespace, error) {
	if gameNamespace, found := gamesNamespaceTypes[gameName]; found {
		return gameNamespace, nil
	}

	return nil, errors.New("Cannot get and create namespace for game " + gameName)
}
