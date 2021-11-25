package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"

	"gopkg.in/yaml.v2"
)

type config struct {
	Guild    uint64 `yaml:"guild"`
	Channel  uint64 `yaml:"channel"`
	Keywords string `yaml:"keywords"`
	Groups   string `yaml:"groups"`
}

func (c *config) getConf() *config {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func action(eventinfo HMC_Bot.MessagePrivate) {
	if eventinfo.UserID == eventinfo.SelfID {
		return
	}
	var c config
	conf := c.getConf()
	keywords_slice := strings.Split(conf.Keywords, "|")
	key_send := false
	for _, value := range keywords_slice {
		//fmt.Printf("index:,value:%d\n", value)
		if strings.Contains(eventinfo.Message, value) {
			key_send = true
		}
	}
	if key_send {
		form_msg_byte, form_msg_err := ioutil.ReadFile("conf/form_msg.txt")
		if form_msg_err != nil {
			fmt.Print(form_msg_err)
		}
		form_msg := string(form_msg_byte) // convert content to a 'string'
		form_msg = strings.ReplaceAll(form_msg, "[form_msg]", eventinfo.Message)
		form_msg = strings.ReplaceAll(form_msg, "[form_nick]", eventinfo.Sender.Nickname)
		form_msg = strings.ReplaceAll(form_msg, "[form_qq]", strconv.FormatUint(uint64(eventinfo.UserID), 10))
		form_msg = strings.ReplaceAll(form_msg, "[form_time]", time.Now().Format("2006-01-02 15:04:05"))

		thank_byte, thank_err := ioutil.ReadFile("conf/thank.txt")
		if thank_err != nil {
			fmt.Print(thank_err)
		}
		HMC_Bot.SendGuildMsg(form_msg, uint64(conf.Guild), uint64(conf.Channel))
		HMC_Bot.SendPrivateMsg(string(thank_byte), eventinfo.UserID)

		//转发群
		groups_slice := strings.Split(conf.Groups, "|")
		for _, value := range groups_slice {
			group_id, _ := strconv.Atoi(value)
			HMC_Bot.SendGroupMsg(form_msg, int64(group_id))
		}
	}

}
func main() {
	// 向事件池里添加函数
	HMC_Bot.Listeners.OnPrivateMsg = append(HMC_Bot.Listeners.OnPrivateMsg, action)
	//初始化一个Bot 并更改配置
	Bot := HMC_Bot.NewBot()
	Bot.Config = HMC_Bot.Config{
		Loglvl:   HMC_Bot.LOGGER_LEVEL_INFO,
		Host:     "0.0.0.0:6700",
		MasterQQ: 779019185,
		Path:     "/",
	}
	Bot.Run()
}
