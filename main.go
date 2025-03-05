package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	BotToken  string  `yaml:"bot_token"`
	WhiteList []int64 `yaml:"whitelist"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:  config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		if contains(config.WhiteList, c.Sender().ID) {
			command := c.Text()
			if c.Text() == "/start" {
				c.Send("It's ok!")
			}
			output, err := executeCommand(command)
			if err != nil {
				return c.Send("Error: \n" + err.Error())
			}
			return c.Send(output)
		}
		return c.Send("Authorization failed!")
	})

	bot.Start()
}

func executeCommand(command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func contains(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
