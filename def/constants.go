package def

const VERSION = "v0.0.5 7-22-2020-836a"
const DEFAULT_PORT = ":7777"
const MAX_SECTORS = 10000
const MAX_USERS = 10000
const PLAYER_START_ROOM = 0
const PLAYER_START_SECTOR = 0
const PFILE_VERSION = "0.0.1"
const SECTOR_VERSION = "0.0.1"

const PASSWORD_HASH_COST = 10
const MAX_PLAYER_NAME_LENGTH = 25
const MIN_PLAYER_NAME_LENGTH = 2
const STRING_UNKNOWN = "unknown"

/*Dir & File*/
const DATA_DIR = "data/"
const PLAYER_DIR = "players/"
const SECTOR_DIR = "sectors/"
const SECTOR_PREFIX = "sec-"
const FILE_SUFFIX = ".dat"

const GREET_FILE = "greet.txt"
const NEWS_FILE = "news.txt"
const PFILE_MAXARGS = 10000

const SERVER_RUNNING = 0
const SERVER_BOOTING = 1
const SERVER_CLOSING = 2
const SERVER_CLOSED = 3
const SERVER_PAUSED = 4
const SERVER_PRIVATE = 5

/*Connection State*/
const CON_STATE_DISCONNECTED = -2
const CON_STATE_DISCONNECTING = -1

const CON_STATE_WELCOME = 0
const CON_STATE_PASSWORD = 100

const CON_STATE_NEWS = 200
const CON_STATE_RECONNECT = 300
const CON_STATE_PLAYING = 1000

/*New Users*/
const CON_STATE_NEW_LOGIN = 300
const CON_STATE_NEW_LOGIN_CONFIRM = 400
const CON_STATE_NEW_PASSWORD = 500
const CON_STATE_NEW_PASSWORD_CONFIRM = 600

/*Player States*/
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
