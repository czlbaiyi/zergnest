package loginserver

var playersData map[string]string = make(map[string]string)

func addPlayer(accout string, password string) (bool, string) {
	playersData[accout] = password
	return true, ""
}
