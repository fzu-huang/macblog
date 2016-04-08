package action

import (
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
)

func Blogdetail(ctx *macaron.Context) {
	blogid := ctx.Params("blogid")
	ret, reason, result := ctl.Getcontentbyid(blogid)
	if ret {
		ctx.Data["OnDetailBlog"] = result
		ctx.Data["Blogid"] = blogid
	} else {
		ctx.Error(403, reason)
		return
	}
	ctx.HTML(200, "blog")
}

func ListBlogByTagname(ctx *macaron.Context){
	tagname := ctx.Params("tagname")
	if tagname == "" {
		BlogIndex(ctx)
		return
	}
	
	ret ,reason, tagresult :=ctl.GetTagByName(tagname)
	if ret {
		ctx.Data["Tagdesc"] = tagresult.TagDesc
	} else {
		ctx.Error(403, reason)
		return
	}
	ret, reason, result := ctl.GetcontentByTagName(1,tagname)
	if ret {
		ctx.Data["OnViewBlogs"] = result
		ctx.Data["Tagname"] = tagname
	} else {
		ctx.Error(403, reason)
		return
	}
	
	ctx.HTML(200, "blogindex")
}