package loginserver

var playersData map[string]string = make(map[string]string)

func LoginOnServer(accout string, password string) (bool, string) {
	if accout == "" {
		return false, "accout must not empty"
	}

	if password == "" {
		return false, "password must not empty"
	}

	if playerData, ok := playersData[accout]; !ok {
		playersData[accout] = password
		return true, ""
	} else if playerData == password {
		return true, ""
	} else {
		return false, "password errer"
	}
}
