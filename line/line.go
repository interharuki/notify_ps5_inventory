package line

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Line struct {
	msg string
}

const URL = "https://notify-api.line.me/api/notify"
const accessToken = "ACCESS_TOKEN"

func NewLine(msg string) *Line {
	return &Line{
		msg: msg,
	}
}

func (l *Line) Notify() error {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Errorf("error %v", err)
		return err
	}

	c := &http.Client{}
	form := url.Values{}
	form.Add("message", l.msg)
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		fmt.Errorf("error %v", err)
		return err

	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		fmt.Errorf("error %v", err)
		return err

	}
	return nil
}
