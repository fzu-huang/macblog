// macblog project main.go
package main

import (
	"github.com/Unknwon/macaron"
	. "github.com/fzu-huang/macblog/action"
	"github.com/fzu-huang/macblog/filter"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/i18n"
	"github.com/macaron-contrib/pongo2"

	"github.com/fzu-huang/macblog/model"
	"github.com/macaron-contrib/session"

	"github.com/golang/glog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	
	"github.com/fzu-huang/macblog/conf"
)

func main() {

	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)

	go func() {
		glog.Infoln(http.ListenAndServe("localhost:6060", nil))
	}()

	m := macaron.Classic()

	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Gziper())
	m.Use(macaron.Static("public/"))
//	m.Use(macaron.Renderer(macaron.RenderOptions{
//		// 模板文件目录，默认为 "templates"
//		Directory: `templates`,
//	}))
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl"},
	}))
	m.Use(i18n.I18n(i18n.Options{
		Directory: "public/conf/locale",
		Langs:     []string{"en-US", "zh-CN"},
		Names:     []string{"English", "简体中文"},
	}))
	m.SetDefaultCookieSecret("huangyang")
	m.Use(session.Sessioner(session.Options{
		//Provider:       "mysql",
		//	ProviderConfig: "root:root@tcp(localhost:3306)/bbs?charset=utf8",
		CookieName:     "MacaronSession",
		CookieLifeTime: 3600,
	}))

	//m.Use(cache.Cacher())

	//加入自定义的过滤，如用户登录判断
	m.Use(filter.OnlineFilter())

	m.Group("", func() {
		m.Get(`/`, BlogIndex)
		m.Get(`/blog`, BlogIndex)
		m.Get("/edit", Edit)
		m.Get("/edit/:blogid", Edit)
		m.Get(`/login`, Log)
		m.Get(`/logout`, Logout)
		m.Post(`/edit/updatecontent/:blogid`, Updatecontent)
		m.Post(`/edit/submitcontent`, Submitcontent)
		m.Post(`/login`, binding.Bind(model.LogUser{}), Login)
		m.Post(`/register`, binding.Bind(model.RegistUser{}), Register)
		m.Post(`/edit/upload`, Uploadfile)
		m.Get(`/ueditor/go/controller`, Controller)
		m.Post(`/ueditor/go/controller`, Uploadfile)
		m.Get(`/blog/:blogid`, Blogdetail)
		m.Get(`/getmoreblogs`, Getmoreblogs)
		m.Get(`/getcomments/:blogid`, Commentinblog)
		m.Get(`/datecounts`,GetDateCounts)
		m.Get(`/date/:year/:month`,ListBlogByDate)
		
		m.Get(`/tags`,GetAllTags)
		m.Post(`/tags/create`,CreateTag)
		m.Get(`/blog/tag/:tagname`,ListBlogByTagname)
	})
	m.Run(conf.PORT)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	go func() {
		sig := <-c

		// todo: stop event
		glog.Infof("application stop...%v", sig)
		pprof.StopCPUProfile()
		os.Exit(0)
	}()
}

func applicationStop() {

}
