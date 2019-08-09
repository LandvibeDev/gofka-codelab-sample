package message

import (
	"encoding/json"
	"github.com/LandvibeDev/gofka-codelab-sample/model"
)

type Message interface {
	GetMessage() []byte
}

type LogMessage struct {
	User model.User `json:"user" form:"user" query:"user"`
	Msg  string     `json:"msg" form:"msg" query:"msg"`
}

func (l LogMessage) GetMessage() []byte {
	msg, _ := json.Marshal(l)
	return msg
}

func Parse(msg []byte) (Message, error) {
	log := new(LogMessage)
	err := json.Unmarshal(msg, log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
