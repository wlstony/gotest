package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/toml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

type Config struct{
	Hosts map[string]Host `json:"hosts"`
}

func main() {
	err := config.LoadFile("/opt/case/gotest/gomicro/demo/config/conf.json")
	fmt.Println("LoadFile err:", err)

	enc := toml.NewEncoder()

	// 通过编码器加载 toml 文件
	err = config.Load(file.NewSource(
		file.WithPath("/opt/case/gotest/gomicro/demo/config/conf.json"),
		source.WithEncoder(enc),
	))
	fmt.Println("config.Load err:", err)
	conf := config.Map()
	fmt.Println(conf)
	var structConf Config
	config.Scan(&structConf)
	fmt.Println("structConf:", structConf)
	cget := config.Get("hosts", "database")
	var host Host
	cget.Scan(&host)
	fmt.Println(host.Address, host.Port)

}