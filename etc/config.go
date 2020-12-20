package etc

import (
	"github.com/astaxie/beego/config"
)

var Cfg config.Configer

func LoadConfig(path string) (err error) {
	if path == "" {
		path = "./config.json"
	}
	Cfg, err = config.NewConfig("json", path)
	return
}
