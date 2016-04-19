package ctl

import(
	"github.com/fzu-huang/macblog/dbutil"
	"github.com/fzu-huang/macblog/model"
)

func Getalldatecounts()(bool,string, []model.DateBlogCount){
	return dbutil.DBgetalldate()
}

func Getallblogsbydate(year,month string,page int)(bool, string, []model.OnViewBlog) {
	return dbutil.DBgetallcontentbydateandlist(year,month,page)	
}