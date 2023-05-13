package def

const VERSION = "Pre-Alpha build, v0.0.13 07092021-1212a"
const LICENSE = "GOMud-Server " + VERSION + "\n" +
	"COPYRIGHT 2020-2021 Carl Frank Otto III (carlotto81@gmail.com)\n" +
	"License: Mozilla Public License 2.0\n" +
	"This information must remain unmodified, fully intact and shown to end-users.\n" +
	"Source: https://github.com/Distortions81/gomud\n" +
	"\n"
const DEFAULT_PORT = ":7777"
const DEFAULT_PORT_SSL = ":7778"

// SSL
const SSL_PEM = "server.pem"
const SSL_KEY = "server.key"

// Maximums
const MAX_SECTORS = 10000
const MAX_USERS = 1000
const MAX_DESC = 950
const MAX_MLE = 100

const MAX_INPUT_LENGTH = 4096
const MAX_OUTPUT_LENGTH = 65536
const MAX_INPUT_LINES = 10
const MAX_CMATCH_SEARCH = 10000
const MAX_COMMANDS_PER_LINE = 15

// Timers
const ROUND_LENGTH_uS = 250000
const ROUND_REST_MS = 3
const CONNECT_THROTTLE_MS = 500
const WELCOME_TIMEOUT_S = 30

const PLAYER_BACKGROUND_uS = 5000000  //5s
const SECTOR_BACKGROUND_uS = 15000000 //15s

// Player/sector defaults
const PLAYER_START_SECTOR = 1
const PLAYER_START_ROOM = 1

const PFILE_VERSION = "0.0.1"
const SECTOR_VERSION = "0.0.1"
const HELPS_VERSION = "0.0.1"

const PASSWORD_HASH_COST = 10
const MAX_PLAYER_NAME_LENGTH = 25
const MIN_PLAYER_NAME_LENGTH = 2
const STRING_UNKNOWN = "unknown"

/*Dir & File*/
const DATA_DIR = "data/"
const PLAYER_DIR = "players/"
const SECTOR_DIR = "sectors/"
const PSECTOR_DIR = "psectors/"
const TEXTS_DIR = "texts/"

const SECTOR_PREFIX = "sec-"
const FILE_SUFFIX = ".dat"

const GREET_FILE = "greet.txt"
const AUREVOIR_FILE = "aurevoir.txt"
const NEWS_FILE = "news.txt"
const HELPS_FILE = "help.txt"

/*Server mode*/
const SERVER_RUNNING = 0
const SERVER_BOOTING = 1
const SERVER_CLOSING = 2
const SERVER_CLOSED = 3
const SERVER_PAUSED = 4
const SERVER_PRIVATE = 5

/*Connection State*/
const CON_STATE_DISCONNECTED = -3
const CON_STATE_DISCONNECTING = -2
const CON_STATE_RELOG = -1

const CON_STATE_WELCOME = 0
const CON_STATE_PASSWORD = 100

const CON_STATE_NEWS = 200
const CON_STATE_RECONNECT_CONFIRM = 300
const CON_STATE_PLAYING = 1000

// New Users
const CON_STATE_NEW_LOGIN = 400
const CON_STATE_NEW_LOGIN_CONFIRM = 500
const CON_STATE_NEW_PASSWORD = 600
const CON_STATE_NEW_PASSWORD_CONFIRM = 700

/*Player States*/
const PLAYER_UNLINKED = -1
const PLAYER_ALIVE = 0
const PLAYER_SIT = 100
const PLAYER_REST = 200
const PLAYER_SLEEP = 300
const PLAYER_STUNNED = 400
const PLAYER_DEAD = 1000

/*Errors*/
const ERROR_NONFATAL = false
const ERROR_FATAL = true

/*Player Type*/
const PLAYER_TYPE_NEW = 0
const PLAYER_TYPE_NORMAL = 100
const PLAYER_TYPE_VETERAN = 200
const PLAYER_TYPE_TRUSTED = 300

const PLAYER_TYPE_BUILDER = 700
const PLAYER_TYPE_MODERATOR = 800
const PLAYER_TYPE_ADMIN = 900
const PLAYER_TYPE_OWNER = 1000

/*OLC */
const OLC_NONE = 0
const OLC_ROOM = 100
const OLC_OBJECT = 200
const OLC_RESET = 300
const OLC_MOBILE = 400
const OLC_QUEST = 500
const OLC_SECTOR = 600
const OLC_EXITS = 700

const SETTING_TYPE_BOOL = 0
const SETTING_TYPE_INT = 100
const SETTING_TYPE_STRING = 200
const SETTING_TYPE_INDEX = 300

const LINESEPA = "-------------------------------------------------------------------------------\r\n"
const LINESEPB = "_______________________________________________________________________________\r\n"

/*Mle Editor*/
const MLE_ADD = 100
const MLE_REMOVE = 200
const MLE_INSERT = 300
const MLE_REPLACE = 400

/*Objects*/
const OBJ_TYPE_NORMAL = 0

/*Wear locations*/
const OBJ_WEAR_HEAD = 100
const OBJ_WEAR_FACE = 200
const OBJ_WEAR_LEYE = 300
const OBJ_WEAR_REYE = 400
const OBJ_WEAR_EYES = 500
const OBJ_WEAR_LEAR = 600
const OBJ_WEAR_REAR = 700
const OBJ_WEAR_EARS = 800
const OBJ_WEAR_NECK = 900
const OBJ_WEAR_LSHOULDER = 1000
const OBJ_WEAR_RSHOULDER = 1100
const OBJ_WEAR_SHOULDERS = 1200
const OBJ_WEAR_CHEST = 1300
const OBJ_WEAR_BODY = 1400
const OBJ_WEAR_BACK = 1500
const OBJ_WEAR_LWIELD = 1600
const OBJ_WEAR_RWIELD = 1700
const OBJ_WEAR_TWOHAND = 1800
const OBJ_WEAR_LBICEP = 1900
const OBJ_WEAR_RBICEP = 2000
const OBJ_WEAR_BICEPS = 2100
const OBJ_WEAR_LFOREARM = 2200
const OBJ_WEAR_RFOREARM = 2300
const OBJ_WEAR_FOREARMS = 2400
const OBJ_WEAR_LELBOW = 2500
const OBJ_WEAR_RELBOW = 2600
const OBJ_WEAR_ELBOWS = 2700
const OBJ_WEAR_LWRIST = 2800
const OBJ_WEAR_RWRIST = 2900
const OBJ_WEAR_WRISTS = 3000
const OBJ_WEAR_LHAND = 3100
const OBJ_WEAR_RHAND = 3200
const OBJ_WEAR_HANDS = 3300
const OBJ_WEAR_LLITTLE_FINGER = 3400
const OBJ_WEAR_LRING_FINGER = 3500
const OBJ_WEAR_LMIDDLE_FINGER = 3600
const OBJ_WEAR_LINDEX_FINGER = 3700
const OBJ_WEAR_LTHUMB = 3800
const OBJ_WEAR_RLITTLE_FINGER = 3900
const OBJ_WEAR_RRING_FINGER = 4000
const OBJ_WEAR_RMIDDLE_FINGER = 4100
const OBJ_WEAR_RINDEX_FINGER = 4200
const OBJ_WEAR_RTHUMB = 4300
const OBJ_WEAR_LEFT_FINGERS = 4400
const OBJ_WEAR_RIGHT_FINGERS = 4500
const OBJ_WEAR_BELLY = 4600
const OBJ_WEAR_WAIST = 4700
const OBJ_WEAR_LCALF = 4800
const OBJ_WEAR_LSHIN = 4900
const OBJ_WEAR_RCALF = 5000
const OBJ_WEAR_RSHIN = 5100
const OBJ_WEAR_CALVES = 5200
const OBJ_WEAR_SHINS = 5300
const OBJ_WEAR_RANKLE = 5400
const OBJ_WEAR_LANKLE = 5500
const OBJ_WEAR_ANKLES = 5600
const OBJ_WEAR_LFOOT = 5700
const OBJ_WEAR_RFOOT = 5800
const OBJ_WEAR_FEET = 5900
const OBJ_WEAR_FLOATING = 6000
