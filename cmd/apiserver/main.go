package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/PandaGoL/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file") // для ввода пути конфига из консоли. флаг (1 аргумент), дефолтное значение (2 аргумент), хелп (3 аргумент)
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig() //создаем переменную для конфига. Данная конструкция нуждна, что бы не палить то, что другим видеть не нужно
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
