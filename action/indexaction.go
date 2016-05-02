package action

import (
	"fmt"
	"encoding/json"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	//	"github.com/fzu-huang/macblog/model"
	//"github.com/macaron-contrib/session"
	"strconv"
)

func BlogIndex(ctx *macaron.Context) {
	
	keyword := ctx.Req.FormValue("key")
	

	//should be   GetAllContentAbout 获取文章简要
	getblogret, err, blogs := ctl.Getallcontent(1,keyword)
	if getblogret {
		ctx.Data["OnViewBlogs"] = blogs
	} else {
		ctx.Error(403, err)
		return
	}
	fmt.Println(blogs)
	ctx.Data["Tagname"] = ""
	ctx.HTML(200, "blogindex")
}

func Getmoreblogs(ctx *macaron.Context) {
	page := ctx.Req.FormValue("page")
	pageint, err := strconv.Atoi(page)
	if err != nil {
		ctx.Error(403, "page arg wrong!")
	}
	
	tagname :=ctx.Req.FormValue("tagname")
	keyword := ctx.Req.FormValue("key")
	//date := ctx.Req.FormValue("date")
	var data []byte
	if tagname != "" {
		ret, reason, result := ctl.GetcontentByTagName(pageint,tagname)
		if ret == false {
			ctx.Error(403, reason)
			return
		}
		data, err = json.Marshal(result)
	} else{
		ret, reason, result := ctl.Getallcontent(pageint,keyword)
		if ret == false {
			ctx.Error(403, reason)
			return
		}
		fmt.Println(result)
		data, err = json.Marshal(result)
	}

	
	if err != nil {
		ctx.Error(403, err.Error())
		return
	}

	ctx.Write(data)
	return
}

func GetDateCounts(ctx *macaron.Context) {
	ret,reason,result := ctl.Getalldatecounts()
	if !ret {
		ctx.Error(403, reason)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		ctx.Error(403, err.Error())
		return
	}

	ctx.Write(data)
	return
}