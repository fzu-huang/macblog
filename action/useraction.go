package action

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/ctl"
	"github.com/fzu-huang/macblog/model"
	"github.com/macaron-contrib/session"
)
func Log(ctx *macaron.Context){
	ctx.HTML(200, "userlog")
}


func Login(ctx *macaron.Context, loguser model.LogUser, sess session.Store) {

	logret, onlineuser := ctl.LogCheck(loguser)
	if logret {
		fmt.Println("USER:", loguser.Name, " logged in!")
		sess.Set(ctx.GetCookie("MacaronSession"), onlineuser)
		fmt.Println(ctx.GetCookie("MacaronSession"), ":", sess.Get(ctx.GetCookie("MacaronSession")))
		if ctx.Req.FormValue("remember") == "1" {
			//fmt.Println("remember = 1")
		}
		from := ctx.Req.FormValue("urlfrom")
		fmt.Println("from:", from)
		ctx.Resp.Write([]byte(`登陆成功`))
		//ctx.Redirect(from)
	} else {
		//ctx.Resp.Write([]byte("用户名或密码错误！！"))
		ctx.Error(404, "用户名或密码错误！！")
	}
}

func Logout(ctx *macaron.Context, sess session.Store) {
	sessionid := ctx.GetCookie("MacaronSession")
	if sess.Get(sessionid) == nil {
		fmt.Println("no session,don't need to logout")
		ctx.Resp.Write([]byte(`未登录`))
		return
	} else {
		sessionval := sess.Get(sessionid).(model.UserStatus)
		if sessionval.Online {
			err := sess.Delete(sessionid)
			if err != nil {
				ctx.Error(404, "注销出错")
				return
			}
			fmt.Println("USER:", sessionid, " logout")
			ctx.Data["LogStatus"] = false
			ctx.Resp.Write([]byte(`已注销`))
			return
		}
		fmt.Println("no session,don't need to logout")
		ctx.Resp.Write([]byte(`未登录`))
		return
	}
}

func Register(ctx *macaron.Context, regUser model.RegistUser, sess session.Store) {
	sessionid := ctx.GetCookie("MacaronSession")
	if sess.Get(sessionid) != nil {
		err := sess.Delete(sessionid)
		if err != nil {
			fmt.Println("注销出错！")
		}
	}
	regret, regresult := ctl.RegCheck(regUser)
	if !regret {
		ctx.Error(403, regresult)
		return
	}
	ctx.Resp.Write([]byte(`注册成功`))
	ctx.Redirect(`/`)
}
