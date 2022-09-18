/**
 * main.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

package main

import (
	"time"

	"github.com/Furkan9015/twitchbot"
)

func main() {

	myBot := twitchbot.BasicBot{
		Channel:     "sedomegalul",
		MsgRate:     time.Duration(20/30) * time.Millisecond,
		Name:        "cumhurbaskanl",
		Port:        "6667",
		PrivatePath: "./private/oauth.json",
		Server:      "irc.chat.twitch.tv",
	}
	myBot.Start()
}
