package mlog

import (
	"fmt"
	"log"
	"time"

	"../def"
	"../glob"
)

func Write(line string) {
	t := time.Now()
	date := fmt.Sprintf("%02d-%02d-%04d_%02d-%02d-%02d", t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute(), t.Second())

	//Async write
	go func(line string, date string) {
		buf := fmt.Sprintf("%s: %s\n", date, line)

		go writeToMods(line)
		log.Print(line)
		glob.MudLog.WriteString(buf)
	}(line, date)
}

func writeToMods(text string) {
	if text == "" {
		return
	}

	for x := 1; x <= glob.PlayerListEnd; x++ {
		player := glob.PlayerList[x]

		if player != nil && player.Valid && player.Connection.Valid {
			if player.Connection.State == def.CON_STATE_PLAYING && player.PlayerType >= def.PLAYER_TYPE_BUILDER {
				message := fmt.Sprintf("[LOG] %s\r\n", text)
				player.Connection.Desc.Write([]byte(message))
			}
		}
	}
}
