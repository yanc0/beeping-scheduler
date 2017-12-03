package scheduler

import (
	"github.com/yanc0/beeping/httpcheck"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"

	beeping "github.com/yanc0/beeping/httpcheck"
)

// Job represents a check to be done
type Job struct {
	ID string
	Check    beeping.Check
	Interval time.Duration
	NextRun  time.Time
}

func (j *Job) GenNextRun() {
	j.NextRun = j.NextRun.Add(j.Interval)
}

func (j *Job) Do(url string) {
	body, _ := json.Marshal(j.Check)
	resp := &httpcheck.Response{}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Unmarshal(ret, resp)
	fmt.Println(resp.HTTPStatus)
}