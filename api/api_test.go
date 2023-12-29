package api

import (
	jikan "github.com/darenliang/jikan-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var day string = time.Now().Weekday().String()

func TestGetDataOk(t *testing.T) {
    filter := jikan.ScheduleFilter(day)
    _, err := GetData(filter)
    assert.Nil(t, err)
}

func TestGetDataError(t *testing.T) {
    filter := jikan.ScheduleFilter("invalid")
    _, err := GetData(filter)
    assert.NotNil(t, err)
}
