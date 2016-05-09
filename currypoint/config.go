package currypoint

import "github.com/newtechfellas/CurryPoint/util"

type Config struct {
	Log struct {
		Level string `json:level`
		File  string `json:file`
	} `json:log`
	MySql struct {
		ConnectionString string
	} `json:mysql`
	ApiPort  string `json:apiPort`
	AdminKey string `json:adminKey`
	Twilio   struct {
		AccountId   string `json:accountId`
		AuthToken   string `json:authToken`
		PhoneNumber string `json:phoneNumber`
	} `json:twilio`
	Plivo struct {
		AuthId      string `json:authId`
		AuthToken   string `json:authToken`
		PhoneNumber string `json:phoneNumber`
	} `json:plivo`
}

func (c Config) String() string {
	return util.Jsonify(c)
}

//Global config
var GlobalConfig Config
