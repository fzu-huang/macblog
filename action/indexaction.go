package action

import (
	"encoding/json"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	//	"github.com/fzu-huang/macblog/model"
	//"github.com/macaron-contrib/session"
	"strconv"
)

func BlogIndex(ctx *macaron.Context) {

	//should be   GetAllContentAbout 获取文章简要
	getblogret, err, blogs := ctl.Getallcontent(1)
	if getblogret {
		ctx.Data["OnViewBlogs"] = blogs
	} else {
		ctx.Error(403, err)
		return
	}
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
		ret, reason, result := ctl.Getallcontent(pageint)
		if ret == false {
			ctx.Error(403, reason)
			return
		}
		data, err = json.Marshal(result)
	}

	
	if err != nil {
		ctx.Error(403, err.Error())
		return
	}

	ctx.Write(data)
	return
}

