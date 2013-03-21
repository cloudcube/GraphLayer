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
