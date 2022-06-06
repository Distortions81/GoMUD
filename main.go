package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gomud/def"
	"gomud/glob"
	"gomud/mlog"
	"gomud/support"
)

func setupListenerSSL() {
	//openssl ecparam -genkey -name prime256v1 -out server.key
	//openssl req -new -x509 -key server.key -out server.pem -days 3650
	cert, err := tls.LoadX509KeyPair(def.DATA_DIR+def.SSL_PEM, def.DATA_DIR+def.SSL_KEY)
	if err != nil {
		log.Print("Error loading SSL certificate, SSL port not opened.")
		log.Print("How to make cert: (put in data directory)")
		log.Print("openssl ecparam -genkey -name prime256v1 -out server.key")
		log.Print("openssl req -new -x509 -key server.key -out server.pem -days 3650")
		return
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	/*Open Listener*/
	listener, err := tls.Listen("tcp4", def.DEFAULT_PORT_SSL, tlsCfg)
	glob.ServerListenerSSL = listener
	support.CheckError("setupSSL: tls.listen", err, def.ERROR_FATAL)

	/*Print Connection*/
	buf := fmt.Sprintf("SSL listener online at: %s", def.DEFAULT_PORT_SSL)
	mlog.Write(buf)
}

func setupListener() {
	/*Find Network*/
	addr, err := net.ResolveTCPAddr("tcp4", def.DEFAULT_PORT)
	support.CheckError("main: resolveTCP", err, def.ERROR_FATAL)

	/*Open Listener*/
	listener, err := net.ListenTCP("tcp4", addr)
	glob.ServerListener = listener
	support.CheckError("main: ListenTCP", err, def.ERROR_FATAL)

	/*Print Connection*/
	buf := fmt.Sprintf("TCP listener online at: %s", addr.String())
	mlog.Write(buf)
}

func WaitNewConnectionSSL() {

	if glob.ServerListenerSSL != nil {

		for glob.ServerState == def.SERVER_RUNNING {

			time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)
			desc, err := glob.ServerListenerSSL.Accept()
			support.AddNetDesc()
			time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)

			/* If there is a connection flood, sleep listeners */
			if err != nil || support.CheckNetDesc() {
				time.Sleep(5 * time.Second)
				desc.Close()
				support.RemoveNetDesc()
			} else {

				_, err = desc.Write([]byte(
					def.LICENSE + glob.Greeting +
						"(SSL Encryption Enabled!)\n(Type NEW to create character) Name:"))
				time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)
				support.NewDescriptor(desc, true)
			}

		}

		glob.ServerListenerSSL.Close()
	}
}

func WaitNewConnection() {

	for glob.ServerState == def.SERVER_RUNNING {

		time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)
		desc, err := glob.ServerListener.Accept()
		support.AddNetDesc()
		time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)

		/* If there is a connection flood, sleep listeners */
		if err != nil || support.CheckNetDesc() {
			time.Sleep(5 * time.Second)
			desc.Close()
			support.RemoveNetDesc()
		} else {

			_, err = desc.Write([]byte(
				def.LICENSE + glob.Greeting +
					"(ENCRYPTION NOT ENABLED!)\n(Type NEW to create character) Name:"))

			time.Sleep(def.CONNECT_THROTTLE_MS * time.Millisecond)
			support.NewDescriptor(desc, false)
		}
	}

	glob.ServerListener.Close()
}

func main() {

	var err error
	t := time.Now()
	glob.MaxRun = time.Nanosecond
	glob.MinRun = time.Second
	glob.MedRun = time.Nanosecond
	glob.BootTime = t

	logName := fmt.Sprintf("log/%v-%v-%v.log", t.Day(), t.Month(), t.Year())
	err = os.Mkdir("log", os.ModePerm)
	if err != nil {
		fmt.Println("Unable to make log dir.")
	}
	glob.MudLog, err = os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to open log file!")
		os.Exit(1)
		return
	}

	support.ReadHelps()
	support.CreateShortcuts()
	support.MakeQuickHelp()
	support.MakeWizHelp()
	support.ReadSectorList()
	support.ReadTextFiles()

	setupListener()
	setupListenerSSL()
	go WaitNewConnection()
	go WaitNewConnectionSSL()

	/*Process connections*/
	mainLoop()

	//After starting loops, wait here for process signals
	glob.SignalHandle = make(chan os.Signal, 1)

	signal.Notify(glob.SignalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-glob.SignalHandle

	support.WriteToAll("Server is shutting down!")
	ServerClose()
}

func ServerClose() {
	glob.ConnectionListLock.Lock()
	//glob.ServerState = def.SERVER_CLOSING

	/*Save everything*/
	for x := 1; x <= glob.ConnectionListEnd; x++ {
		con := glob.ConnectionList[x]
		if con.Player != nil && con.Player.Valid {
			support.WritePlayer(con.Player, false)
			if con.Desc != nil {
				con.Desc.Write([]byte("Your character has been saved!\r\n"))
			}
		}
	}
	support.WriteSectorList()
	support.WriteToAll("All sectors saved!")
	support.WriteToAll("")
	time.Sleep(1 * time.Second)
	support.WriteToAll(glob.AuRevoir)

	glob.ConnectionListLock.Unlock()
}

/*TODO: If performance becomes an issue, sleep once per round */
func mainLoop() {

	rand.Seed(time.Now().UTC().UnixNano())

	glob.MinRun = time.Duration(time.Second)

	/* Player rounds */
	go func() {
		numPlayerLast := glob.ConnectionListEnd
		sleepFor := time.Duration(def.ROUND_LENGTH_uS)
		tickNum := 0

		for glob.ServerState == def.SERVER_RUNNING {
			tickNum++

			glob.ConnectionListLock.Lock() /*--- LOCK ---*/

			/*Handle 0 players*/
			if numPlayerLast <= 0 {
				sleepFor = time.Duration(def.ROUND_LENGTH_uS) * time.Microsecond
				time.Sleep(sleepFor)
			} else {
				sleepFor = time.Duration(def.ROUND_LENGTH_uS/numPlayerLast) * time.Microsecond
			}

			cEnd := glob.ConnectionListEnd
			glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/

			tempCount := 0
			for x := 0; x <= cEnd; x++ {

				glob.ConnectionListLock.Lock() /*--- LOCK ---*/
				if glob.ConnectionList[x].Valid {
					start := time.Now()
					tempCount++

					/*Check for stale connections*/
					if time.Since(glob.ConnectionList[x].ConnectedFor) > (def.WELCOME_TIMEOUT_S*time.Second) &&
						glob.ConnectionList[x].State <= def.CON_STATE_WELCOME {
						glob.ConnectionList[x].Valid = false
						glob.ConnectionList[x].Desc.Close()
						support.RemoveNetDesc()
					}

					support.ReadPlayerInputBuffer(&glob.ConnectionList[x])

					glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/
					end := time.Now()
					spent := end.Sub(start) /*Round sleep, slice per player*/
					time.Sleep(sleepFor - spent)

					glob.PerLock.Lock() //PERF-LOCK
					avrTime := support.MovingExpAvg(float64(spent), float64(10*time.Nanosecond), 1.0, 1000.0)

					if spent > glob.MaxRun {
						glob.MaxRun = spent
					} else if spent < glob.MinRun {
						glob.MinRun = spent
					}

					if tickNum%10 == 0 {
						glob.PerfStats = fmt.Sprintf("Ran for %10v, Max: %10v, Min: %10v, Avr: %10v, Slept: %10v",
							spent.String(), glob.MaxRun, glob.MinRun, time.Duration(avrTime).String(), sleepFor-spent)
					}

					glob.PerLock.Unlock() //PERF-UNLOCK

					glob.ConnectionListLock.Lock() /*--- LOCK ---*/
				}
				glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/
			}

			glob.ConnectionListLock.Lock() /*--- LOCK ---*/
			numPlayerLast = tempCount
			glob.NumPlayers = tempCount

			/*Tick down connection list end if not used*/
			if glob.ConnectionListEnd > 0 &&
				glob.ConnectionList[glob.ConnectionListEnd].Valid == false {
				glob.ConnectionListEnd--
			}

			glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/

			time.Sleep(def.ROUND_REST_MS) /*Limit max CPU, and allow background to run*/

		}
	}()

	/*Player background tasks*/
	go func() {
		for glob.ServerState == def.SERVER_RUNNING {

			glob.ConnectionListLock.Lock() /*--- LOCK ---*/
			numPlayers := glob.NumPlayers
			glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/

			/*Delay based on number of characters*/
			if numPlayers > 1 {
				sleepFor := def.PLAYER_BACKGROUND_uS / (numPlayers)
				time.Sleep(time.Duration(sleepFor) * time.Microsecond)
				//fmt.Println(fmt.Sprintf("Player autosave slept for %v uS.", sleepFor))
			} else {
				time.Sleep(time.Duration(def.PLAYER_BACKGROUND_uS) * time.Microsecond)
			}

			glob.ConnectionListLock.Lock() /*--- LOCK ---*/

			/*Cycle players, skipping anyone not playing*/
			if glob.PlayerBackgroundPos <= glob.PlayerListEnd {
				glob.PlayerBackgroundPos++

				for glob.PlayerBackgroundPos < glob.PlayerListEnd &&
					glob.PlayerList[glob.PlayerBackgroundPos].Valid == false &&
					glob.PlayerList[glob.PlayerBackgroundPos].Connection != nil &&
					glob.PlayerList[glob.PlayerBackgroundPos].Connection.Valid {
					glob.PlayerBackgroundPos++
				}

			} else {
				glob.PlayerBackgroundPos = 1
			}

			/*Autosave players, check if valid*/
			if glob.PlayerList[glob.PlayerBackgroundPos] != nil &&
				glob.PlayerList[glob.PlayerBackgroundPos].Valid {

				/*Marked dirty, or save requested*/
				if glob.PlayerList[glob.PlayerBackgroundPos].Dirty ||
					glob.PlayerList[glob.PlayerBackgroundPos].ReqSave {

					/*Write*/
					support.WritePlayer(glob.PlayerList[glob.PlayerBackgroundPos], true)

					/*If requested, notify player*/
					if glob.PlayerList[glob.PlayerBackgroundPos].ReqSave {
						support.WriteToPlayer(glob.PlayerList[glob.PlayerBackgroundPos], "Character saved.")
						glob.PlayerList[glob.PlayerBackgroundPos].ReqSave = false
					}

				}
			}

			/*Remove players that haven't been connected for some time*/
			player := glob.PlayerList[glob.PlayerBackgroundPos]
			if player != nil &&
				player.Valid &&
				player.Location.RoomLink != nil {

				if player.Connection.Valid == false &&
					player.UnlinkedTime.IsZero() == false &&
					time.Since(player.UnlinkedTime) > (2*time.Minute) {

					player.UnlinkedTime = time.Time{}
					support.WritePlayer(player, true)
					support.WriteToRoom(player, fmt.Sprintf("%s fades into nothing...", player.Name))
					support.RemovePlayer(player)
				}
			}

			glob.ConnectionListLock.Unlock() /*--- UNLOCK ---*/
		}
	}()

	/*Sector background tasks*/
	go func() {
		for glob.ServerState == def.SERVER_RUNNING {

			if glob.SectorsListEnd > 2 {
				sleepFor := def.SECTOR_BACKGROUND_uS / (glob.SectorsListEnd - 1)
				time.Sleep(time.Duration(sleepFor) * time.Microsecond)
				//fmt.Println(fmt.Sprintf("Sector autosave slept for %v uS,", sleepFor))
			} else {
				time.Sleep(def.SECTOR_BACKGROUND_uS * time.Microsecond)
			}

			/*Cycle sectors*/
			if glob.SectorBackgroundPos < glob.SectorsListEnd {
				glob.SectorBackgroundPos++
			} else {
				glob.SectorBackgroundPos = 0
			}

			/*--- LOCK ---*/
			glob.ConnectionListLock.Lock()
			/*--- LOCK ---*/

			/*Autosave sectors*/
			if glob.SectorsList[glob.SectorBackgroundPos].Dirty &&
				glob.SectorsList[glob.SectorBackgroundPos].Valid &&
				glob.SectorsList[glob.SectorBackgroundPos].Name != "" {
				support.WriteSector(&glob.SectorsList[glob.SectorBackgroundPos])
			}

			/*--- UNLOCK ---*/
			glob.ConnectionListLock.Unlock()
			/*--- UNLOCK ---*/
		}
	}()
}
