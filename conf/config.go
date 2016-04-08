package conf

import (
	"bytes"
	"log"
	"os"
)

var PORT int = 4001

var ConfigJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

func init() {
	file, err := os.Open(`public/conf/config.json`)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)

	ConfigJson = buf.Bytes()
}
