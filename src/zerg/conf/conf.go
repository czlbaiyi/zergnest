package conf

type ServerUrl struct {
	Ip   string
	Port string
}

func (s ServerUrl) GetServerUrl() string {
	var url string = s.Ip
	url += ":"
	url += s.Port
	return url
}

func (s ServerUrl) GetServerSeparatorPort() string {
	var url string = ":"
	url += s.Port
	return url
}

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

	LoginServerUrls []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10086"}}
	DataServerUrls  []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10186"}}
	StateServerUrls []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10286"}}
	GameServerUrls  []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10386"}}
	RankServerUrls  []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10486"}}
	MailServerUrls  []ServerUrl = []ServerUrl{ServerUrl{"127.0.0.1", "10586"}}
)
