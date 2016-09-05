package conf

var (
	LenStackBuf = 4096

	// log
	LogLevel string
	LogPath  string

	// console
	ConsolePort   int
	ConsolePrompt string = "Leaf# "
	ProfilePath   string

	PendingWriteNum int
	MACHINE_ID      uint16 = 1

	LoginServerPort    string   = ":10086"
	DataServerPort     string   = ":10186"
	StateServerPort    string   = ":10286"
	GameServerPort     string   = ":10386"
	MailServerPort     string   = ":10486"
	RankServerPort     string   = ":10586"
	LoginServerAddList []string = []string{"127.0.0.1"}
	DataServerAddList  []string = []string{"127.0.0.1"}
	StateServerAddList []string = []string{"127.0.0.1"}
	GameServerAddList  []string = []string{"127.0.0.1"}
	RankServerAddList  []string = []string{"127.0.0.1"}
	MailServerAddList  []string = []string{"127.0.0.1"}
)
