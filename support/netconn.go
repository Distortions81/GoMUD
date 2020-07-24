package support

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"../def"
	"../glob"
)

func AutoResolveAddress(con *glob.ConnectionData) {

	addr := con.Desc.RemoteAddr().String()
	addrp := strings.Split(addr, ":")
	addrLen := len(addrp)
	if addrLen > 0 {
		addr = addrp[0]
	}
	con.Address = addr
}

func NewDescriptor(desc net.Conn, ssl bool) {

	if desc == nil {
		return
	}

	/*--- LOCK ---*/
	glob.ConnectionListLock.Lock()
	defer glob.ConnectionListLock.Unlock()
	/*--- LOCK ---*/

	for x := 1; x <= glob.ConnectionListEnd; x++ {

		if glob.ConnectionList[x].Valid == true {
			continue
		} else {
			newConnection := glob.ConnectionData{
				Name:         def.STRING_UNKNOWN,
				Desc:         desc,
				Address:      "",
				SSL:          ssl,
				State:        def.CON_STATE_WELCOME,
				ConnectedFor: time.Now(),
				IdleTime:     time.Now(),
				BytesOut:     0,
				BytesIn:      0,
				Player:       nil,
				Valid:        true}
			glob.ConnectionList[x] = newConnection
			AutoResolveAddress(&newConnection)

			buf := fmt.Sprintf("Recycling connection #%d.", x)
			log.Println(buf)

			go ReadConnection(&glob.ConnectionList[x])
			return
		}
	}

	/*Generate new descriptor data*/
	if glob.ConnectionListEnd >= def.MAX_USERS-1 {
		log.Println("Create ConnectionData: MAX_USERS REACHED!")
		desc.Write([]byte("Sorry, something has gone wrong (MAX_DESCRIPTORS)!\r\nGoodbye!\r\n"))
		desc.Close()
		return
	}

	/*Create*/
	glob.ConnectionListEnd++
	newConnection := glob.ConnectionData{
		Name:         def.STRING_UNKNOWN,
		Desc:         desc,
		Address:      "",
		SSL:          ssl,
		State:        def.CON_STATE_WELCOME,
		ConnectedFor: time.Now(),
		IdleTime:     time.Now(),
		BytesOut:     0,
		BytesIn:      0,
		Player:       nil,
		Valid:        true}
	AutoResolveAddress(&newConnection)
	glob.ConnectionList[glob.ConnectionListEnd] = newConnection
	buf := fmt.Sprintf("Created new connection #%d.", glob.ConnectionListEnd)
	log.Println(buf)

	go ReadConnection(&glob.ConnectionList[glob.ConnectionListEnd])
	return

}

func ReadConnection(con *glob.ConnectionData) {

	if con == nil || !con.Valid {
		return
	}

	/*--- LOCK ---*/
	glob.ConnectionListLock.Lock()
	/*Create reader*/
	reader := bufio.NewReader(con.Desc)
	/*Create reader*/
	glob.ConnectionListLock.Unlock()
	/*--- UNLOCK ---*/

	for {

		input, err := reader.ReadString('\n')
		bIn := len(input)

		con.BytesIn += bIn
		trackBytesIn(con)

		/*Connection died*/
		if err != nil {
			glob.ConnectionListLock.Lock()
			DescWriteError(con, err)
			glob.ConnectionListLock.Unlock()
			return
		}

		<-glob.Round
		time.Sleep(def.SYNC_DELAY * time.Millisecond)
		go HandleReadConnection(con, input)
	}
}

func HandleReadConnection(con *glob.ConnectionData, input string) {
	/*--- LOCK ---*/
	glob.ConnectionListLock.Lock()
	/*--- LOCK ---*/

	/*Handles all user input*/
	interpretInput(con, input)

	/*Handle players marked for disconnection*/
	/*Doing this at the end is cleaner*/
	if con.State == def.CON_STATE_DISCONNECTING {
		WriteToDesc(con, "\r\n\r\n *** Goodbye! ***")

		con.State = def.CON_STATE_DISCONNECTED

		/*Invalidate player's connection*/
		if con.Player != nil && con.Player.Valid &&
			con.Player.Connection != nil {
			con.Player.Connection.Valid = false
		}

		con.Desc.Close()
		con.State = def.CON_STATE_DISCONNECTED
		con.Player = nil
		con.Valid = false
	}

	/*--- UNLOCK ---*/
	glob.ConnectionListLock.Unlock()
	/*--- UNLOCK ---*/
}

func trackBytesOut(con *glob.ConnectionData) {

	player := con.Player

	if player == nil || !player.Valid || con == nil || !con.Valid {
		return
	}
	player.BytesOut[con.Address] += (con.BytesOut - con.BytesOutRecorded)
	con.BytesOutRecorded = con.BytesOut
}

func trackBytesIn(con *glob.ConnectionData) {

	player := con.Player

	if player == nil || !player.Valid || con == nil || !con.Valid {
		return
	}
	player.BytesIn[con.Address] += (con.BytesIn - con.BytesInRecorded)
	con.BytesInRecorded = con.BytesIn
}

func DescWriteError(c *glob.ConnectionData, err error) {
	if err != nil {
		CheckError("DescWriteError", err, def.ERROR_NONFATAL)

		if c != nil {
			if c.Valid && c.Name != def.STRING_UNKNOWN && c.State == def.CON_STATE_PLAYING {
				if c.Player != nil && c.Player.Valid {
					buf := fmt.Sprintf("%s lost their network connection.", c.Player.Name)
					c.Player.UnlinkedTime = time.Now()
					c.Player.Connection.Valid = false

					WriteToRoom(c.Player, buf)
				}
			} else {
				buf := fmt.Sprintf("%s disconnected.", c.Address)
				log.Println(buf)
			}

			c.State = def.CON_STATE_DISCONNECTED
			c.Desc.Close()
			c.Valid = false
			c.Player = nil
		}
	}
}

func WriteToDesc(c *glob.ConnectionData, text string) {

	if c == nil || !c.Valid {
		return
	}
	text, overflow := TruncateString(text, def.MAX_OUTPUT_LENGTH)
	if overflow {
		cstring := " Name: " + c.Name + ", Addr:" + c.Address
		log.Println("WriteToDesc: string too large, Truncated!" + cstring)
	}

	message := fmt.Sprintf("%s\r\n", text)
	bytes, err := c.Desc.Write([]byte(message))
	c.BytesOut += bytes
	trackBytesOut(c)

	DescWriteError(c, err)
}

func WriteToPlayer(player *glob.PlayerData, text string) {

	if player == nil || !player.Valid || player.Connection == nil || !player.Connection.Valid {
		return
	}

	message := fmt.Sprintf("%s\r\n", text)
	bytes, err := player.Connection.Desc.Write([]byte(message))
	player.Connection.BytesOut += bytes
	trackBytesOut(player.Connection)

	DescWriteError(player.Connection, err)
}

func WriteToAll(text string) {
	if text == "" {
		return
	}

	for x := 0; x <= glob.ConnectionListEnd; x++ {
		var con *glob.ConnectionData
		con = &glob.ConnectionList[x]
		if con.Valid == false {
			continue
		}
		if con.State == def.CON_STATE_PLAYING {
			message := fmt.Sprintf("[Global] %s\r\n", text)
			bytes, err := con.Desc.Write([]byte(message))
			con.BytesOut += bytes
			trackBytesOut(con)

			DescWriteError(con, err)
		}
	}
	log.Println(text)
}

func WriteToOthers(player *glob.PlayerData, text string) {
	if player == nil || !player.Valid || text == "" {
		return
	}

	pc := player.Connection

	for x := 0; x <= glob.ConnectionListEnd; x++ {
		var con *glob.ConnectionData
		con = &glob.ConnectionList[x]
		if con.Valid == false {
			continue
		}
		if con.Desc != pc.Desc && con.State == def.CON_STATE_PLAYING {
			message := fmt.Sprintf("[Global] %s\r\n", text)
			bytes, err := con.Desc.Write([]byte(message))
			con.BytesOut += bytes
			trackBytesOut(con)

			DescWriteError(con, err)
		}
	}
	log.Println(text)
}

func WriteToRoom(player *glob.PlayerData, text string) {
	if player == nil || !player.Valid || text == "" {
		return
	}

	if player.RoomLink != nil {
		for _, target := range player.RoomLink.Players {
			if target != nil && target != player {
				WriteToPlayer(target, "[Room] "+text)
			}
		}
	} else {
		buf := fmt.Sprintf("WriteToRoom: %v is not in a room.", player.Name)
		log.Println(buf)
	}
}
