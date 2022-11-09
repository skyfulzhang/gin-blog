package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(num int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	data := make([]rune, num)
	for index := range data {
		data[index] = letters[rand.Intn(len(letters))]
	}
	return string(data)
}

func EncryMd5(str string) string {
	// 实例化md5,返回一个 hash.Hash 对象
	m := md5.New()
	// 加密的内容,转成byte
	m.Write([]byte(str))
	// 对内容进行校验计算
	uintData := m.Sum(nil)
	// 转化成字符串
	strData := hex.EncodeToString(uintData)
	return strData
}

func GetRequestData(context *gin.Context) string {
	if context.ContentType() == "multipart/form-data" {
		_ = context.Request.ParseMultipartForm(1024)
		formData, _ := json.Marshal(context.Request.Form)
		return string(formData)
	}
	if context.ContentType() == "application/x-www-form-urlencoded" {
		_ = context.Request.ParseForm()
		formData, _ := json.Marshal(context.Request.Form)
		return string(formData)
	}
	byteData, _ := context.GetRawData()
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteData))
	return string(byteData)
}

var incrNum uint64

// NewTraceID New trace id
func NewTraceID() string {
	return fmt.Sprintf("trace-id-%d-%s-%d",
		os.Getpid(),
		time.Now().Format("2006.01.02.15.04.05.999"),
		atomic.AddUint64(&incrNum, 1))
}

// NewUUID Create uuid
func NewUUID() string {
	return uuid.New().String()
}
