package support

import (
	"gomud/def"
	"gomud/glob"
)

/*Command list, types/levels in constants.go*/
var CommandList = []glob.Command{

	/*allow short, short, name, function, type, quick-help*/

	/*Moderator*/
	{AS: true, Short: "wiz", Name: "wizhelp", Cmd: CmdWizHelp, Type: def.PLAYER_TYPE_NEW,
		Help: "Help for builders/moderators"},
	{AS: true, Short: "", Name: "stats", Cmd: CmdStats, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "See bandwidth usage"},
	{AS: false, Short: "", Name: "reloadText", Cmd: CmdReloadText, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Reload text files, greeting, aurevoir, news, etc."},
	{AS: false, Short: "", Name: "reloadPlayer", Cmd: CmdReloadPlayer, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Reload a player that is currently logged in."},
	{AS: false, Short: "", Name: "playerType", Cmd: CmdPlayerType, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Set a player's type"},
	{AS: false, Short: "", Name: "SavePlayers", Cmd: CmdSavePlayers, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Save all players"},
	{AS: false, Short: "", Name: "SaveHelps", Cmd: CmdWriteHelps, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Save help file"},
	{AS: false, Short: "", Name: "reloadHelp", Cmd: CmdReloadHelpst, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Reload help file."},
	{AS: false, Short: "", Name: "perfStats", Cmd: CmdPerfStat, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Show performance stats."},
	{AS: false, Short: "", Name: "shutdown", Cmd: CmdShutdown, Type: def.PLAYER_TYPE_MODERATOR,
		Help: "Shutdown the server."},

	/*Builder*/
	{AS: false, Short: "", Name: "asave", Cmd: CmdAsave, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Save game areas"},
	{AS: false, Short: "", Name: "olc", Cmd: CmdOLC, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Edit sectors, rooms, objs, etc (WIP)."},
	{AS: false, Short: "", Name: "dig", Cmd: CmdDig, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Create a new room, to the <exit name>"},
	{AS: false, Short: "", Name: "goto", Cmd: CmdGoto, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Goto <location>"},
	{AS: false, Short: "slist", Name: "sectorlist", Cmd: CmdSectorList, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Shows list of sectors."},
	{AS: false, Short: "rlist", Name: "roomlist", Cmd: CmdRoomList, Type: def.PLAYER_TYPE_BUILDER,
		Help: "Shows list of rooms."},

	/*Shortcuts*/
	{AS: true, Short: "n", Name: "north", Cmd: CmdNorth, Type: def.PLAYER_TYPE_NEW,
		Help: "Go north"},
	{AS: true, Short: "s", Name: "south", Cmd: CmdSouth, Type: def.PLAYER_TYPE_NEW,
		Help: "Go south"},
	{AS: true, Short: "e", Name: "east", Cmd: CmdEast, Type: def.PLAYER_TYPE_NEW,
		Help: "Go east"},
	{AS: true, Short: "w", Name: "west", Cmd: CmdWest, Type: def.PLAYER_TYPE_NEW,
		Help: "Go west"},
	{AS: true, Short: "u", Name: "up", Cmd: CmdUp, Type: def.PLAYER_TYPE_NEW,
		Help: "Go up"},
	{AS: true, Short: "d", Name: "down", Cmd: CmdDown, Type: def.PLAYER_TYPE_NEW,
		Help: "Go down"},

	/*Player*/
	{AS: true, Short: "", Name: "recall", Cmd: CmdRecall, Type: def.PLAYER_TYPE_NEW,
		Help: "transport home, to set: recall set"},
	{AS: true, Short: "", Name: "emote", Cmd: CmdEmote, Type: def.PLAYER_TYPE_NEW,
		Help: "emote tests... -> SomePlayer tests..."},
	{AS: true, Short: "", Name: "commands", Cmd: CmdCommands, Type: def.PLAYER_TYPE_NEW,
		Help: "You are here"},
	{AS: true, Short: "", Name: "who", Cmd: CmdWho, Type: def.PLAYER_TYPE_NEW,
		Help: "See who is online"},
	{AS: true, Short: "", Name: "look", Cmd: CmdLook, Type: def.PLAYER_TYPE_NEW,
		Help: "Look around the room"},
	{AS: true, Short: "", Name: "go", Cmd: CmdGo, Type: def.PLAYER_TYPE_NEW,
		Help: "Move around! go <exit name>"},
	{AS: false, Short: "", Name: "alias", Cmd: CmdAlias, Type: def.PLAYER_TYPE_NEW,
		Help: "alias add <shortcut> <output> (incomplete)"},
	{AS: true, Short: "", Name: "say", Cmd: CmdSay, Type: def.PLAYER_TYPE_NEW,
		Help: "Talk to other people in the room"},
	{AS: true, Short: "", Name: "chat", Cmd: CmdChat, Type: def.PLAYER_TYPE_NEW,
		Help: "Talk to other people across the world"},
	{AS: true, Short: "", Name: "save", Cmd: CmdSave, Type: def.PLAYER_TYPE_NEW,
		Help: "Save your character's progress. (autosave is on)"},
	{AS: false, Short: "", Name: "quit", Cmd: CmdQuit, Type: def.PLAYER_TYPE_NEW,
		Help: "Quit the game"},
	{AS: false, Short: "", Name: "relogin", Cmd: CmdRelog, Type: def.PLAYER_TYPE_NEW,
		Help: "Go back to login screen."},
	{AS: true, Short: "", Name: "config", Cmd: CmdConfig, Type: def.PLAYER_TYPE_NEW,
		Help: "Configure settings"},
	{AS: true, Short: "", Name: "news", Cmd: CmdNews, Type: def.PLAYER_TYPE_NEW,
		Help: "See news"},
	{AS: true, Short: "", Name: "editor", Cmd: MleEditor, Type: def.PLAYER_TYPE_NEW,
		Help: "Text editor."},
	{AS: true, Short: "", Name: "help", Cmd: CmdGetHelps, Type: def.PLAYER_TYPE_NEW,
		Help: "Help system"},
}
