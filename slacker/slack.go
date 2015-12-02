package slacker

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Slack is struct
type Slack struct {
	C *Config
}

// Message is post payload
type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

var cli = &http.Client{Timeout: time.Duration(5) * time.Second}

// Post is post message to slack
func (s *Slack) Post(msg Message) (err error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)

	url := s.C.URL
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}

	res, err := cli.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return errors.New("post.Failed")
	}

	return
}
