package config

import (
	"errors"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

// get the format time -- YYYY-MM-DD HH:ii:ss
func SelfTime() string {
	currentTime := time.Now()
	t := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, time.UTC)

	return t.String()
}

// get the custom IP address
func SelfIp(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}

// judgment path is or not exist
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// get file absolute path
func GetFileAbsolutePath(dir, currentDir, path string) string {
	separator := string(os.PathSeparator)
	return dir + separator + currentDir + separator + path + ".go"
}

// verify request is valid or not.
func ValidRequest(path string) (err error) {
	dirName, err := os.Getwd()
	if err != nil {
		return
	}

	fileName := GetFileAbsolutePath(dirName, "client", path)

	if flag := PathExists(fileName); flag == false {
		return errors.New("file not exist.")
	}

	return nil
}

// Call
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of param is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}
