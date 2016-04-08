package action

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/conf"
	"io"
	"os"
	"path"
	"strings"
)

func Uploadfile(ctx *macaron.Context) {

	switch ctx.Req.FormValue("action") {
	case "uploadimage":
		uploadimage(ctx)
	case "uploadscrawl":
		uploadscrawl(ctx)
	case "":
		ctx.Resp.Write([]byte(`wrong request! check action `))
	}

}

func uploadscrawl(ctx *macaron.Context) {
	fmt.Println(ctx.Req.RequestURI)

	//	file, header, err := ctx.Req.FormFile("serverparam")
	//	if err != nil {
	//		panic(err)
	//	}
	inputfilebyte := []byte(ctx.Req.FormValue("upfile"))
	fmt.Println(ctx.Req.FormValue("upfile"))
	inputfilestream := bytes.NewReader(inputfilebyte)

	base64is := base64.NewDecoder(base64.StdEncoding, inputfilestream)
	filename := uuid.NewUUID().String() + `.png`

	err := os.MkdirAll(path.Join("public", "upload", "scrawl"), 0775)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(path.Join("public", "upload", "scrawl", filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, base64is)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/upload/scrawl/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": filename,                                   //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	ctx.Resp.Write(b)
}

func uploadimage(ctx *macaron.Context) {
	file, header, err := ctx.Req.FormFile("upfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)

	err = os.MkdirAll(path.Join("public", "upload", "image"), 0775)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(path.Join("public", "upload", "image", filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/upload/image/%s", filename), //保存后的文件路径
		"title":    "",                                        //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                           //原始文件名
		"state":    "SUCCESS",                                 //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	ctx.Resp.Write(b)
}

func Controller(ctx *macaron.Context) {
	ctx.Resp.Write(conf.ConfigJson)
	return
}
