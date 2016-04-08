package action

import (
	"encoding/json"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	"strconv"
)

func Commentinblog(ctx *macaron.Context) {
	blogid := ctx.Params("blogid")
	pagestr := ctx.Req.FormValue("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		ctx.Error(403, `need a  int param: page`)
		return
	}
	ret, reason, result := ctl.GetallcommentByBlogId(blogid, page)

	if ret {
		if len(result) > 0 {
			data, err := json.Marshal(result)
			if err != nil {
				ctx.Error(403, "get data failed...")
			}
			ctx.Write(data)
			return
		}
	} else {
		ctx.Error(403, reason)
	}
}
