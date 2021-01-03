package internal

import (
	"encoding/json"
	"github.com/FerdinaKusumah/rumpi/model"
	"github.com/FerdinaKusumah/rumpi/utils"
	"github.com/jasonlvhit/gocron"
	"sync"
	"time"
)

type Job struct {
	sync.Mutex
	WatchApi         string
	NotifyApi        string
	LastHash         string
	IntervalInSecond uint64
	Data             *model.ResponseData
	LastCheck        time.Time
	Verbose          bool
}

func NewJob(watchUrl, postUrl string, intervalInSecond uint64, verbose bool) *Job {
	return &Job{
		Mutex:            sync.Mutex{},
		WatchApi:         watchUrl,
		NotifyApi:        postUrl,
		Verbose:          verbose,
		LastHash:         "",
		IntervalInSecond: intervalInSecond,
		Data:             &model.ResponseData{},
		LastCheck:        time.Time{},
	}
}

func (j *Job) WatchResource() {
	j.Lock()
	defer j.Unlock()
	var result = new(model.ResponseData)
	if result.StatusCode, result.Data, result.Error = utils.Get(j.WatchApi); result.Error != nil {
		utils.LogError.ErrorF(`Error Watch %s`, j.WatchApi)
	}

	// get current time
	var currentTime = time.Now().UTC()
	var newHash = utils.HashData(result)
	j.LastCheck = currentTime
	j.Data = result

	if j.Verbose == true {
		utils.LogInfo.InfoF(`Watching %s with status %d every %d second`, j.WatchApi, result.StatusCode, j.IntervalInSecond)
	}

	// if resource is down or any changes is resource then notify to others
	if j.LastHash == "" {
		j.LastHash = newHash
	}

	if !utils.IsStatusCodeIsOk(result.StatusCode) || newHash != j.LastHash {
		// prepare message
		var message = utils.NewMessage(j.WatchApi, currentTime, result)
		res, _ := json.Marshal(message)
		if err := utils.Post(res, j.NotifyApi); err != nil {
			utils.LogError.ErrorF(`Error Notify %s`, j.NotifyApi)
		}
		j.LastHash = newHash
	}
}

// RunWatch is main function to run watcher
func RunWatch(conf *Option) {
	var s = gocron.NewScheduler()
	for _, p := range conf.ListApi {
		var d = NewJob(p.WatchApi, p.NotifyApi, uint64(p.Interval), p.Verbose)
		s.Every(d.IntervalInSecond).Seconds().Do(d.WatchResource)
	}
	<-s.Start()
}
