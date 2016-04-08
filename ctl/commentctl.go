package ctl

import (
	"github.com/fzu-huang/macblog/dbutil"
	"github.com/fzu-huang/macblog/model"
)

func Submitcomment(commenterid, blogid, superid, content string) (bool, string) {

	return dbutil.DBsubmitcomment(commenterid, blogid, superid, content)
}

func Deletecomment(commentid string) (bool, string) {
	return dbutil.DBdeletecomment(commentid)
}
func GetallcommentByBlogId(blogid string, page int) (bool, string, []model.OnDetailComment) {

	ret, reason, result := dbutil.DBgetallcommentbylist(blogid, page)
	length := len(result)
	for i := 0; i < length; i++ {
		if result[i].SuperId != "-1" {
			getsupercomret, _, getsupercomresult := Getcommentbyid(result[i].SuperId)
			if getsupercomret {
				result = append(result, getsupercomresult)
			}
		}
	}
	return ret, reason, result
}

func Getcommentbyid(commentid string) (bool, string, model.OnDetailComment) {

	return dbutil.DBgetcommentbyid(commentid)
}
