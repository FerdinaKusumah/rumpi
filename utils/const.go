package utils

import (
	"github.com/apsdehal/go-logger"
	"os"
	"time"
)

const (
	TimeOutDuration = 10 * time.Second
	TimeFormat      = "2006-01-02 15:04:05"
)

// logger config
var (
	LogError, _ = logger.New("[ERROR]", 4, os.Stdout)
	LogInfo, _ = logger.New("[INFO]", 6, os.Stdout)
)