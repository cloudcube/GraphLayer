package graphdb

import (
	"testing"
)

func TestDial(t *testing.T) {
	settingFile := "databaseConfig.json"
	_, err := Dial(settingFile)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDialWithParam(t *testing.T) {
	url := "http://localhost:7474/db/data"
	userName := "test"
	password := "123456"
	_, err := DialWithParam(url, userName, password)
	if err != nil {
		t.Error(err)
	}
}
