package utils

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     = &Config{}
)

type Config struct {
	Mysql *Mysql
}

type Mysql struct {
	Username string
	Password string
	Iphost   string
	Dbname   string
}

func init() {
	flag.StringVar(&confPath, "conf", "/home/mlcc/dev/GOPATH/src/mg1/app01/conf/conf.toml", "-conf path")
}

func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
