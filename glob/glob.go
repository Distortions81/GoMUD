package glob

import (
	"net"
	"os"
	"time"

	"gomud/def"

	"github.com/sasha-s/go-deadlock"
)

var SignalHandle chan os.Signal
var BootTime time.Time

/*Descriptor counting*/
var OpenDesc int
var OpenDescLock deadlock.Mutex

/*Files desc locks*/
var PlayerFileLock deadlock.Mutex
var SectorsFileLock deadlock.Mutex

/*Listeners, server state*/
var ServerState = def.SERVER_RUNNING
var ServerListener *net.TCPListener
var ServerListenerSSL net.Listener

/*Log desc, round channel*/
var MudLog *os.File
var Round <-chan struct{}

/*Main Game Data*/
var ConnectionListEnd int
var ConnectionList [def.MAX_USERS + 1]ConnectionData
var ConnectionListLock deadlock.Mutex

var PlayerListEnd int
var PlayerList [def.MAX_USERS + 1]*PlayerData

var SectorsListEnd int
var SectorsList [def.MAX_SECTORS]SectorData

var HelpSystem HelpMain

var QuickHelp string
var WizHelp string

// Texts
var Greeting string
var AuRevoir string
var News string

// Autosave
var PlayerBackgroundPos int
var SectorBackgroundPos int
var NumPlayers int

/*Performance stats*/
var MaxRun time.Duration
var MinRun time.Duration
var MedRun time.Duration
var PerfStats string
var PerLock deadlock.Mutex
