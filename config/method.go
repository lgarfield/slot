package config

import (
	"errors"
	"os"
	"reflect"
	"time"
)

// get the format time -- YYYY-MM-DD HH:ii:ss
func SelfTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
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
func ValidRequest(dir, path string) (err error) {
	dirName, err := os.Getwd()
	if err != nil {
		return
	}

	fileName := GetFileAbsolutePath(dirName, dir, path)

	if flag := PathExists(fileName); flag == false {
		return errors.New("file not exist.")
	}

	return nil
}

// Call
func Call(name interface{}, params ...interface{}) (result []interface{}, err error) {
	f := reflect.ValueOf(name)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of param is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	out := f.Call(in)
	if len(out) > 0 {
		//prepare out paras
		result = make([]interface{}, 0)
		for _, v := range out {
			if "error" == v.Type().String() {
				if nil != v.Interface() {
					err = v.Interface().(error)
					return
				}
			} else {
				result = append(result, v.Interface())
			}
		}
	}

	return
}
