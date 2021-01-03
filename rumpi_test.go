package main

import (
	"example/watch-api/internal"
	"example/watch-api/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var (
	// to testing this please open https://crudcrud.com then change hash variable from that page
	mockHashVar   = "310dbd6dcef840459afcabe4712b8b3c"
	realNotifyUrl = fmt.Sprintf("https://crudcrud.com/api/%s/notifyChangesUserData", mockHashVar)
	RealWatchUrl  = fmt.Sprintf("https://crudcrud.com/api/%s/users", mockHashVar)
)

func simulateChangeData(suffix string) {
	utils.Post(
		[]byte(fmt.Sprintf(`{"email": "user_%s"}`, suffix)),
		fmt.Sprintf("https://crudcrud.com/api/%s/users", mockHashVar),
	)
}

func TestUtilsGet(t *testing.T) {
	var statusCode, resp, error = utils.Get("https://google.com")
	assert.Equal(t, statusCode, 200)
	assert.Equal(t, error, nil)
	assert.Greater(t, len(resp), 0)
}

func TestJobWithError(t *testing.T) {
	notifyUrl := "http://localhost:8090/notify"
	watchUrl := "http://localhost:8090/watch"
	var d = internal.NewJob(watchUrl, notifyUrl, 4, true)
	d.WatchResource()
	assert.Equal(t, d.Verbose, true)
	assert.Equal(t, d.IntervalInSecond, uint64(4))
	assert.Equal(t, d.NotifyApi, notifyUrl)
	assert.Equal(t, d.WatchApi, watchUrl)
	assert.Equal(t, d.LastHash, "19451ee3be8ae701d07293d90baac591aee5542c")
	assert.Equal(t, d.Data.StatusCode, 0)
	assert.NotEqual(t, d.Data.Error, nil)
	assert.Equal(t, len(d.Data.Data), 0)
}

func TestJobWithSuccess(t *testing.T) {
	// to testing this please open https://crudcrud.com then change hash variable from that page
	simulateChangeData(time.Now().String())
	var d = internal.NewJob(RealWatchUrl, realNotifyUrl, 1, true)
	d.WatchResource()
	var lastHash = d.LastHash
	assert.Equal(t, d.Verbose, true)
	assert.Equal(t, d.IntervalInSecond, uint64(1))
	assert.Equal(t, d.NotifyApi, realNotifyUrl)
	assert.Equal(t, d.WatchApi, RealWatchUrl)
	assert.NotEqual(t, d.LastHash, "")
	assert.Equal(t, d.Data.StatusCode, 200)
	assert.Equal(t, d.Data.Error, nil)
	assert.Greater(t, len(d.Data.Data), 0)

	// now simulate changes data
	simulateChangeData(time.Now().String())
	d.WatchResource()
	assert.NotEqual(t, d.LastHash, lastHash)
}

func TestMultipleJobWithSuccess(t *testing.T) {
	// to testing this please open https://crudcrud.com then change hash variable from that page
	var d = internal.NewJob(RealWatchUrl, realNotifyUrl, 1, true)
	for i := 0; i <= 5; i++ {
		simulateChangeData(strconv.Itoa(i))
		d.WatchResource()
		var lastHash = d.LastHash
		assert.Equal(t, d.Verbose, true)
		assert.Equal(t, d.IntervalInSecond, uint64(1))
		assert.Equal(t, d.NotifyApi, realNotifyUrl)
		assert.Equal(t, d.WatchApi, RealWatchUrl)
		assert.NotEqual(t, d.LastHash, "")
		assert.Equal(t, d.Data.StatusCode, 200)
		assert.Equal(t, d.Data.Error, nil)
		assert.Greater(t, len(d.Data.Data), 0)

		// now simulate changes data
		simulateChangeData(strconv.Itoa(i))
		d.WatchResource()
		assert.NotEqual(t, d.LastHash, lastHash)
	}
}
