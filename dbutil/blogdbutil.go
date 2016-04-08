package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/fzu-huang/macblog/model"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"time"
)

func DBsubmitcontent(blogname, writername, content string, tagid int) (bool, string) {
	blogname = template.HTMLEscapeString(blogname)
	writername = template.HTMLEscapeString(writername)
	submittime := time.Now().Format("2006-01-02 15:04:05")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("insert into blog (blogname,writername,content,submittime,updatetime,tagid) values (?,?,?,?,?,?)")
	checkErr(err)
	result, err := stmt.Exec(blogname, writername, content, submittime, submittime,tagid)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBupdatecontent(blogid string, blogname, content string,tagid int) (bool, string) {
	blogname = template.HTMLEscapeString(blogname)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("update  blog set blogname=?, content=?, tagid=? where id=?")
	checkErr(err)
	result, err := stmt.Exec(blogname, content,tagid, blogid)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBdeletecontent(blogid string) (bool, string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete  blog where id=?")
	checkErr(err)
	result, err := stmt.Exec(blogid)
	checkErr(err)
	if err != nil {
		return false, err.Error()
	}
	fmt.Println(result)
	return true, ""
}

func DBgetcontentbyid(blogid string) (bool, string, model.OnDetailBlog) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("select blogname, writername, content, submittime, updatetime, id,itemtag.tagname from blog LEFT JOIN itemtag on itemtag.tagid = blog.tagid where id=?")
	checkErr(err)
	raws, err := stmt.Query(blogid)
	checkErr(err)
	result := model.OnDetailBlog{}
	var id, bname, wname, content, stime, utime,tagname string
	defer raws.Close()
	if raws.Next() {
		err = raws.Scan(&bname, &wname, &content, &stime, &utime, &id, &tagname)
		checkErr(err)
		result.BlogName = bname
		result.Content = template.HTML(content)
		result.WriterName = wname
		result.BlogId = id
		result.Submittime = stime
		result.Updatetime = utime
		result.Tag = tagname
		return true, "", result
	}
	return false, "find no result match id: " + blogid, result
}

func DBgetallcontentbylist(page int) (bool, string, []model.OnViewBlog) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")

	checkErr(err)
	defer db.Close()
	countper := 9
	limitindex := (page - 1) * countper
	stmt, err := db.Prepare("select id, blogname, writername, content, submittime,itemtag.tagname from blog LEFT JOIN itemtag on itemtag.tagid = blog.tagid order by submittime desc limit ?,?")
	checkErr(err)
	raws, err := stmt.Query(limitindex, countper)
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}
	var id, blogname, writername, content, submittime,tagname string

	//var viewblogs  []model.OnViewBlog{}
	viewblogs := make([]model.OnViewBlog, countper)
	defer raws.Close()
	for i := 0; i < 9; i++ {
		if raws.Next() {
			err = raws.Scan(&id, &blogname, &writername, &content, &submittime, &tagname)
			checkErr(err)
			viewblogs[i] = model.OnViewBlog{
				BlogId:     id,
				BlogName:   blogname,
				WriterName: writername,
				Content:    template.HTML(content),
				Submittime: submittime,
				Tag:tagname,
			}
		} else {
			if i == 0 {
				return true, "find no blog.", nil
			}
			viewblogs = viewblogs[:i]
			break
		}
	}
	return true, "", viewblogs
}

func DBgetallcontentbytagandlist(tagid int, page int) (bool, string, []model.OnViewBlog) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")

	checkErr(err)
	defer db.Close()
	countper := 9
	limitindex := (page - 1) * countper
	stmt, err := db.Prepare("select b.id, b.blogname, b.writername, b.content, b.submittime,t.tagname from blog b, itemtag t where b.tagid = ? and t.tagid = ? order by b.submittime desc limit ?,?")
	checkErr(err)
	raws, err := stmt.Query(tagid,tagid,limitindex, countper)
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}
	var id, blogname, writername, content, submittime,tagname string

	//var viewblogs  []model.OnViewBlog{}
	viewblogs := make([]model.OnViewBlog, countper)
	defer raws.Close()
	for i := 0; i < 9; i++ {
		if raws.Next() {
			err = raws.Scan(&id, &blogname, &writername, &content, &submittime,&tagname)
			checkErr(err)
			viewblogs[i] = model.OnViewBlog{
				BlogId:     id,
				BlogName:   blogname,
				WriterName: writername,
				Content:    template.HTML(content),
				Submittime: submittime,
				Tag: tagname,
			}
		} else {
			if i == 0 {
				return true, "find no blog.", nil
			}
			viewblogs = viewblogs[:i]
			break
		}
	}
	return true, "", viewblogs
}

func DBgetallcontentbytagnameandlist(tagname string, page int) (bool, string, []model.OnViewBlog) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")

	checkErr(err)
	defer db.Close()
	countper := 9
	limitindex := (page - 1) * countper
	stmt, err := db.Prepare("select b.id, b.blogname, b.writername, b.content, b.submittime from blog b, itemtag t where t.tagname = ? and t.tagid = b.tagid order by b.submittime desc limit ?,?")
	checkErr(err)
	raws, err := stmt.Query(tagname,limitindex, countper)
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}
	var id, blogname, writername, content, submittime string

	//var viewblogs  []model.OnViewBlog{}
	viewblogs := make([]model.OnViewBlog, countper)
	defer raws.Close()
	for i := 0; i < 9; i++ {
		if raws.Next() {
			err = raws.Scan(&id, &blogname, &writername, &content, &submittime)
			checkErr(err)
			viewblogs[i] = model.OnViewBlog{
				BlogId:     id,
				BlogName:   blogname,
				WriterName: writername,
				Content:    template.HTML(content),
				Submittime: submittime,
				Tag: tagname,
			}
		} else {
			if i == 0 {
				return true, "find no blog.", nil
			}
			viewblogs = viewblogs[:i]
			break
		}
	}
	return true, "", viewblogs
}