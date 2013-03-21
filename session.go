package graphdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type databaseConfig struct {
	url      string
	username string
	password string
}

func Dial(settingFile string) (*Session, error) {
	var (
		dbSetting databaseConfig
		session   *Session
	)
	session = new(Session)
	f, err := ioutil.ReadFile(settingFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(f, &dbSetting)
	if len(dbSetting.url) < 1 {
		dbSetting.url = "http://127.0.0.1:7474/db/data"
	}
	session.URL = dbSetting.url
	if len(dbSetting.username) < 0 {
		return nil, errors.New("UserName can't match,please check your config file!")
	}
	session.Username = dbSetting.username
	if len(dbSetting.password) < 0 {
		return nil, errors.New("Password can't match,please check your config file!")
	}
	session.Password = dbSetting.password
	_, err = session.Send(dbSetting.url, "")
	return session, err
}

func (session *Session) Send(url string, data string) (string, error) {
	var (
		resp *http.Response // http response
		buf  bytes.Buffer   // contains http response body
		err  error
	)
	if len(url) < 1 {
		url = session.URL + "node" // default path
	}
	// fmt.Println(url)
	client := new(http.Client)
	switch strings.ToLower(session.Method) { // which http method
	case "delete":
		req, e := http.NewRequest("DELETE", url, nil)
		if e != nil {
			err = e
			break
		}
		session.setAuth(*req)
		resp, err = client.Do(req)
	case "post":
		body := strings.NewReader(data)
		req, e := http.NewRequest("POST", url, body)
		if e != nil {
			err = e
			break
		}
		req.Header.Set("Content-Type", "application/json")
		session.setAuth(*req)
		resp, err = client.Do(req)
	case "put":
		body := strings.NewReader(data)
		req, e := http.NewRequest("PUT", url, body)
		if e != nil {
			err = e
			break
		}
		req.Header.Set("Content-Type", "application/json")
		session.setAuth(*req)
		resp, err = client.Do(req)
	case "get":
		fallthrough
	default:
		req, e := http.NewRequest("GET", url, nil)
		if e != nil {
			err = e
			break
		}
		session.setAuth(*req)
		resp, err = client.Do(req)
		// fmt.Print("Response content: ")
		// fmt.Println(resp)

	}
	if err != nil {
		return "", err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	session.StatusCode = resp.StatusCode // the calling method should do more inspection with chkStatusCode() method and determine if the operation was successful or not.
	return buf.String(), nil
}

// sets Basic HTTP Auth
func (session *Session) setAuth(req http.Request) {
	if len(session.Username) > 0 || len(session.Password) > 0 {
		req.SetBasicAuth(session.Username, session.Password)
	}
}
