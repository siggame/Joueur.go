package ${lowercase_first(game_name)}
<%include file='functions.noCreer' />
import "joueur/base"

type AI struct {
	base.BaseAI

	// The reference to the Game instance this AI is playing.
	game Game

	// The reference to the Player this AI controls in the Game.
	player Player
}

func (ai AI) getPlayerName() string {
${merge(
	'// ', 'getName',
	'	return "' + game_name + ' Go Player"'
)}
}

// This is called once the game starts and your AI knows its playerID and game.
// You can initialize your AI here.
func (ai AI) start() {
${merge(
	'// ', 'start',
	'	// pass'
)}
}

// This is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai AI) gameUpdated()  {
${merge(
	'// ', 'game-updated',
	'	// pass'
)}
}

// This is called when the game ends, you can clean up your data and dump files here if need be.
//
// @param won True means you won, false means you lost.
// @param reason The human readable string explaining why you won or lost.
func (ai AI) ended(won bool, reason string) {
${merge(
	'// ', 'ended',
	'	// pass'
)}
}

// Chess specific AI actions
% for function_name in ai['function_names']:
<% function_params = ai['functions'][function_name]%>
${shared['go']['function_top'](function_name, function_params, 'AI')}
${merge(
	'// ', function_name,
	"""// Put your game logic here for {}
	return{}
""".format(
		function_name,
		(' ' + shared['go']['default_value'](function_params['returns']['type'])) if function_params['returns'] else ''
	)
)}
}
% endfor