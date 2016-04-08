package filter

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/fzu-huang/macblog/model"
	"github.com/macaron-contrib/session"
)

func OnlineFilter() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		sessionID := ctx.GetCookie("MacaronSession")
		if sess.Get(sessionID) == nil {
			fmt.Println("need to login")
			ctx.Data[`LogStatus`] = false
		} else {
			sessionval := sess.Get(sessionID).(model.UserStatus)
			if sessionval.Online {
				fmt.Println("USER:", sessionval.Name, " on line!")
				ctx.Data[`UserName`] = sessionval.Name
				ctx.Data[`Authority`] = sessionval.Authority
				ctx.Data[`LogStatus`] = sessionval.Online
				//sess.Set(sessionID, sessionval)
			}
		}
	}
}
