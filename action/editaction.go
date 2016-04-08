package action

import (
	"strconv"
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	"github.com/fzu-huang/macblog/model"
	"github.com/macaron-contrib/session"
	"html/template"
)

func Edit(ctx *macaron.Context, sess session.Store) {
	sessionid := ctx.GetCookie("MacaronSession")
	fmt.Println(sess.Get(sessionid))
	ctx.Data["LogStatus"] = false
	if sess.Get(sessionid) == nil {
		fmt.Println("no session, need to login")
		ctx.HTML(200, "userlog")
		return
	} else {
		sessionval := sess.Get(sessionid).(model.UserStatus)
		if sessionval.Online {
			fmt.Println("USER:", sessionid, " on line!")
			ctx.Data["LogStatus"] = true
			blogid := ctx.Params("blogid")
			if blogid != ""{
				ctx.Data["blogid"] = blogid
				
				ret, _, sourceblog := ctl.Getcontentbyid(blogid)
				if ret{
					ctx.Data["findresult"] =true
					ctx.Data["sourceblog"] = sourceblog
				} else {
					ctx.Data["findresult"] =false
				}	
			}
			ctx.HTML(200, "edit")
			return
		} else {
			ctx.HTML(200, "userlog")
			return
		}
	}
}

func Submitcontent(ctx *macaron.Context) {

	ctx.Req.ParseForm()
	//tagidstr := ctx.Params("tagid")//从url中获取tagid
	tagidstr := ctx.Req.Request.FormValue("tagid")
	tagid,err  := strconv.Atoi(tagidstr)
	if err != nil {
		ctx.Error(403, err.Error())
		return
	}
	blogname := ctx.Req.FormValue("blogname")

	writername := ctx.Data[`UserName`].(string)
	content := ctx.Req.Request.FormValue("content")
	//fmt.Println("blogname:", blogname, "content:", content)

	ret, reason := ctl.Submitcontent(blogname, writername, content,tagid)
	if ret {
		ctx.Data["blogname"] = blogname
		ctx.Data["writername"] = writername
		//content = template.HTML(content)

		ctx.Data["content"] = template.HTML(content)
		ctx.HTML(200, "success")
	}

	ctx.Error(403, reason)
	return
}

func Updatecontent(ctx *macaron.Context) {
	blogid := ctx.Params("blogid")
	if blogid == "" {
		fmt.Println("blogid:",blogid)
		ctx.Error(403, "请检查blogid")
		return
	}
	
	ctx.Req.ParseForm()
	//tagidstr := ctx.Params("tagid")//从url中获取tagid
	tagidstr := ctx.Req.Request.FormValue("tagid")
	tagid,err  := strconv.Atoi(tagidstr)
	if err != nil {
		fmt.Println(err)
		ctx.Error(403, err.Error())
		return
	}
	blogname := ctx.Req.FormValue("blogname")

	writername := ctx.Data[`UserName`].(string)
	content := ctx.Req.Request.FormValue("content")
	fmt.Println("blogname:", blogname, "content:", content)

	ret, reason := ctl.Updatecontent(blogid,blogname,  content,tagid)
	if ret {
		ctx.Data["blogname"] = blogname
		ctx.Data["writername"] = writername
		//content = template.HTML(content)
		//ctx.Data["tagname"] = tagidstr
		ctx.Data["content"] = template.HTML(content)
		ctx.HTML(200, "success")
		return
	}

	ctx.Error(403, reason)
	return
}
