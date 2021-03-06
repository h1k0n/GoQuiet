package gqclient

import (
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"time"
)

type stateManager interface {
	ParseConfig(string) error
	SetAESKey(string)
}

// State stores global variables
type State struct {
	SS_LOCAL_HOST  string
	SS_LOCAL_PORT  string
	SS_REMOTE_HOST string
	SS_REMOTE_PORT string
	Now            func() time.Time
	Opaque         int
	Key            string
	TicketTimeHint int
	AESKey         []byte
	ServerName     string
	Browser        string
}

// ParseConfig parses the config file into a State variable
func (sta *State) ParseConfig(configPath string) error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &sta)
	if err != nil {
		return err
	}
	return nil
}

// SetAESKey calculates the SHA256 of the string key
func (sta *State) SetAESKey() {
	h := sha256.New()
	h.Write([]byte(sta.Key))
	sta.AESKey = h.Sum(nil)
}
