package client

import (
	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
)

const (
	UnknownError api.ErrorCode = iota
	UserNotFound
	UserAlreadyBanned
	UserAlreadyUnbanned
	YouAreBanned
)

var codeMessages = map[api.ErrorCode]string{
	UnknownError:        "Unknown Error",
	UserNotFound:        "User not found",
	UserAlreadyBanned:   "User already banned",
	UserAlreadyUnbanned: "User is already unbanned",
	YouAreBanned:        "You've been banned from the game. Please contact us if you feel this is a mistake.",
}

func CodeText(code int) string {
	if text, ok := codeMessages[code]; ok {
		return text
	}

	logrus.Error("unsupported code, defaulting to unknown")
	return codeMessages[UnknownError]
}
