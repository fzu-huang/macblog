package action

import(
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	"github.com/fzu-huang/macblog/model"
)

func GetAllTags(ctx *macaron.Context){
	ret, reason, tags := ctl.GetAllTags()
	if !ret {
		ctx.Error(403, reason)
		return
	}
	data ,err := json.Marshal(tags)
	if err!= nil {
		ctx.Error(403, err.Error())
		return
	}
	//fmt.Println(tags)
	ctx.Write(data)
	return
}

func CreateTag (ctx*macaron.Context){
	defer ctx.Req.Request.Body.Close()
	data, err := ioutil.ReadAll(ctx.Req.Request.Body)
	if err!= nil {
		ctx.Error(403, err.Error())
		return
	}
	var t model.OnViewTag
	err  = json.Unmarshal(data, &t)
	if err!= nil {
		ctx.Error(403, err.Error())
		return
	}
	fmt.Println(t)
	ret,reason := ctl.SubmitTag(t.TagName, t.TagDesc)
	if ret {
		ctx.WriteHeader(201)
		return
	} else{
//		if strings.Contains(reason, "DUPLICATE") ||  strings.Contains(reason, "duplicate") {
//			ctx.Error(403, "类目名已存在，"+err.Error())
//			return
//		}
		ctx.Error(403, "创建失败！"+reason)
		return
	}
}
