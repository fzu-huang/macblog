package ctl

import (
	//	"errors"
	"github.com/fzu-huang/macblog/dbutil"
	"github.com/fzu-huang/macblog/model"
	//"html/template"
)
func Submitcontent(blogname, writername, content string,tagid int) (bool, string) {

	return dbutil.DBsubmitcontent(blogname, writername, content,tagid)
}

func Updatecontent(blogid,blogname,  content string,tagid int) (bool, string) {

	return dbutil.DBupdatecontent(blogid ,blogname,  content,tagid)
}

func Getallcontent(page int) (bool, string, []model.OnViewBlog) {
	//	ok, err, blogs := dbutil.DBgetallcontentbylist(page)
	//	if err != nil {
	//		return err, nil
	//	}
	//	if blogs == nil {
	//		return errors.New("find no blogs.", nil)
	//	}

	return dbutil.DBgetallcontentbylist(page)
}

func Getcontentbyid(blogid string) (bool, string, model.OnDetailBlog) {
	//	ok, err, blogs := dbutil.DBgetallcontentbylist(page)
	//	if err != nil {
	//		return err, nil
	//	}
	//	if blogs == nil {
	//		return errors.New("find no blogs.", nil)
	//	}

	return dbutil.DBgetcontentbyid(blogid)
}

func GetcontentByTagName(page int , tagname string)(bool, string, []model.OnViewBlog) {
	return dbutil.DBgetallcontentbytagnameandlist(tagname,page)	
}

func GetcontentByTagid(page int , tagid int)(bool, string, []model.OnViewBlog) {
	return dbutil.DBgetallcontentbytagandlist(tagid,page)	
}