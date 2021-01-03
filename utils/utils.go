package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"example/watch-api/model"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Get is method to doing http get request
func Get(url string) (int, []byte, error) {
	return fasthttp.GetTimeout(nil, url, TimeOutDuration)
}

// Post is method to doing http post request to notify.go
func Post(data []byte, url string) error {
	var req = fasthttp.AcquireRequest()
	req.Header.SetContentType("application/json")
	req.SetBody(data)
	req.Header.SetMethodBytes([]byte("POST"))
	req.SetRequestURIBytes([]byte(url))
	return fasthttp.Do(req, nil)
}

// HashToEncodeString is method to convert hash data to encode string
func HashToEncodeString(data []byte) string {
	return hex.EncodeToString(data)
}

// HashData is method to make signature in every request action
func HashData(data *model.ResponseData) string {
	h := sha1.New()
	res, _ := json.Marshal(data)
	h.Write(res)
	return HashToEncodeString(h.Sum(nil))
}

// IsStatusCodeIsOk is method to check if response status bigger than 500
// which is error in server
func IsStatusCodeIsOk(statusCode int) bool {
	return statusCode < http.StatusInternalServerError
}

// MessageCode is method to translate status code to string message
// Http status message refer to wikipedia
// https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
func MessageCode(statusCode int) string {
	return http.StatusText(statusCode)
}

// NewMessage is method to determine what message should build based on response
func NewMessage(url string, lastCheck time.Time, response *model.ResponseData) *model.PostModel {
	var d = new(model.PostModel)
	d.Url = url
	d.DateTime = lastCheck.Format(TimeFormat)
	d.Message = MessageCode(response.StatusCode)
	// prepared data payload
	d.Status = new(model.NotifyData)
	d.Status.Error = response.Error
	d.Status.Data = string(response.Data)
	d.Status.StatusCode = response.StatusCode
	return d
}

// LoadConfigFromFile is method to load config from files
func LoadConfigFromFile(path string) []byte {
	var (
		conf []byte
		err  error
	)
	if conf, err = ioutil.ReadFile(path); err != nil {
		LogInfo.InfoF(`Unable to load file config from path: %s, err : %s`, path, err)
		os.Exit(1)
	}
	return conf
}

// LoadConfigFromUrl is method to load config from url
func LoadConfigFromUrl(url string) []byte {
	var (
		conf []byte
		err  error
	)
	if _, conf, err = Get(url); err != nil {
		LogInfo.InfoF(`Unable to load file config from path: %s, err : %s`, url, err)
		os.Exit(1)
	}
	return conf
}

// ParseConfig is method to parse config from byte to list config
func ParseConfig(data []byte) []*model.ConfigFile {
	var res []*model.ConfigFile
	if err := json.Unmarshal(data, &res); err != nil {
		LogInfo.InfoF(`Unable parse config: %s`, err)
		os.Exit(1)
	}
	return res
}

// LoadConfig is config to determine what source need to load and load it to json file
func LoadConfig(srcFile string) []*model.ConfigFile {
	if strings.Contains(srcFile, "http") || strings.Contains(srcFile, "https") {
		return ParseConfig(LoadConfigFromUrl(srcFile))
	}
	return ParseConfig(LoadConfigFromFile(srcFile))
}
