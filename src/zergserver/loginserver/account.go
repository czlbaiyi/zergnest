package loginserver

type Account struct {
	Id          uint64 //自增
	AccountType int    //帐号类型，0：本地uuid登录 mac地址等，1：本地账号密码，2：平台登陆方式
	ClientType  string //客户端类型：“IOS”，“ANDROID”或者“OTHER”
	AccountId   string //帐号id
	AccountPwd  string //帐号密码
	Uuid        string //密码 客户端唯一标示(IOS为UUID，ANDROID为IMEI)
	CreateTime  int64  //创建时间
}
