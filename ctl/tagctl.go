package ctl

import (
	//	"errors"
	"github.com/fzu-huang/macblog/dbutil"
	"github.com/fzu-huang/macblog/model"
	//"html/template"
)

func GetAllTags()(bool, string, []model.OnViewTag){
	return dbutil.DBgetalltags()
}

func SubmitTag(tagname,tagdesc string)(bool, string){
	return dbutil.DBaddtag(tagname,tagdesc)
}

func GetTagByName(tagname string )(bool, string, model.OnViewTag){
	return dbutil.DBgetagbyname(tagname)
}