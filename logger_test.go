package logger

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"
)

type testMessage struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (t *testMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

func TestInfo(t *testing.T) {
	Info(
		"started_in", (time.Millisecond * 20).String(),
		"enabled", true,
		errors.New("first"),
		"started in: %s", time.Millisecond*25,
	)

	Info(
		"started_in", time.Millisecond*20,
		"enabled", true,
		errors.New("first"),
		"started",
	)

	Info(
		"started_in", time.Millisecond*20,
		"enabled", true,
		"started",
	)

	jsonBytes, _ := (&testMessage{
		Id:   "150",
		Name: "MNO",
	}).MarshalJSON()

	Info(
		os.ErrClosed,
		&testMessage{
			Id:   "100",
			Name: "XYZ",
		},
		"started_in", time.Millisecond*20,
		"enabled", true,
		"payload", &testMessage{
			Id:   "101",
			Name: "ABC",
		},
		JSON(jsonBytes),
		jsonBytes,
		"event", JSON(jsonBytes),
		"started",
	)
}
