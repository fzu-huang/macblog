package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/fzu-huang/macblog/model"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"time"
)

func DBsubmitcomment(commenterid, blogid, superid, content string) (bool, string) {

	cmttime := time.Now().Format("2006-01-02 15:04:05")
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("insert into comment (commenterid,blogid,superid,content,cmttime) values (?,?,?,?,?)")
	checkErr(err)
	result, err := stmt.Exec(commenterid, blogid, superid, content, cmttime)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBdeletecomment(commentid string) (bool, string) {
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete  comment where id=?")
	checkErr(err)
	result, err := stmt.Exec(commentid)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBgetcommentbyid(commentid string) (bool, string, model.OnDetailComment) {
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("select comment.id, comment.superid, comment.content, comment.cmttime, usermsg.name, blog.blogname from comment left join usermsg on comment.commenterid=usermsg.id left join blog on comment.blogid=blog.id  where comment.id=?")
	checkErr(err)
	raws, err := stmt.Query(commentid)
	checkErr(err)
	result := model.OnDetailComment{}
	var id, superid, content, cmttime, commentername, blogname string
	defer raws.Close()
	if raws.Next() {
		err = raws.Scan(&id, &superid, &content, &cmttime, &commentername, &blogname)
		checkErr(err)
		result.CommentId = id
		result.CommenterName = commentername
		result.BlogName = blogname
		result.SuperId = superid
		result.Content = template.HTML(content)
		result.CmtTime = cmttime
		return true, "", result
	}
	return false, "find no result match id: " + commentid, result
}

func DBgetallcommentbylist(blogid string, page int) (bool, string, []model.OnDetailComment) {
	db, err := sql.Open("mysql", "DBURL")

	checkErr(err)
	defer db.Close()
	countper := 9
	limitindex := (page - 1) * countper
	stmt, err := db.Prepare("select comment.id, comment.superid, comment.content, comment.cmttime, usermsg.name, blog.blogname from comment left join usermsg on comment.commenterid=usermsg.id left join blog on comment.blogid=blog.id where comment.blogid=? order by comment.cmttime desc limit ?,?")
	checkErr(err)
	raws, err := stmt.Query(blogid, limitindex, countper)
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}
	var id, superid, content, cmttime, commentername, blogname string

	//var viewcomments  []model.OnDetailComment{}
	viewcomments := make([]model.OnDetailComment, countper)
	defer raws.Close()
	for i := 0; i < 9; i++ {
		if raws.Next() {
			err = raws.Scan(&id, &superid, &content, &cmttime, &commentername, &blogname)
			checkErr(err)
			viewcomments[i] = model.OnDetailComment{
				CommentId:     id,
				CommenterName: commentername,
				BlogName:      blogname,
				SuperId:       superid,
				Content:       template.HTML(content),
				CmtTime:       cmttime,
			}
		} else {
			if i == 0 {
				return true, "find no blog.", nil
			}
			viewcomments = viewcomments[:i]
			break
		}
	}
	return true, "", viewcomments
}
