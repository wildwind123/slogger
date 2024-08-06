package slogger

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/go-faster/errors"
)

type Vector struct {
	Client   *http.Client
	Url      string
	User     string
	Password string
}

func (v *Vector) Write(p []byte) (n int, err error) {
	err = v.sendLog(p)
	if err != nil {
		fmt.Printf("cant send log \n err = %v \n -------log----- \n %s  \n ----- end log ------", err, string(p))
		return 0, errors.Wrap(err, "cant send log")
	}
	return
}

func (v *Vector) sendLog(p []byte) error {
	method := "POST"

	payload := bytes.NewReader(p)

	req, err := http.NewRequest(method, v.Url, payload)

	if err != nil {
		return errors.Wrap(err, "cant NewRequest")
	}
	req.Header.Add("Content-Type", "application/json")

	req.SetBasicAuth(v.User, v.Password)

	res, err := v.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "cant request")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "cant read body")
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("wrong status code %d, body = %s", res.StatusCode, string(body)))
	}

	return nil
}
