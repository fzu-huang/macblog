package dbutil

import (
	"strings"
	"database/sql"
	"fmt"
	"github.com/fzu-huang/macblog/model"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	//"time"
)

func DBaddtag(tagname, tagdesc string) (bool, string){
	tagname = template.HTMLEscapeString(tagname)
	tagdesc = template.HTMLEscapeString(tagdesc)
	
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("insert into itemtag (tagname,tagdescribe) values (?,?)")
	if err != nil &&strings.Contains(err.Error(), "Duplicate"){
		return false, err.Error()
	}
	checkErr(err)
	result, err := stmt.Exec(tagname, tagdesc)
	if err != nil &&strings.Contains(err.Error(), "Duplicate"){
		return false, err.Error()
	}
	checkErr(err)
	fmt.Println(result)
	return true, ""
}

func DBgetalltags() (bool, string, []model.OnViewTag) {
	db, err := sql.Open("mysql", DBURL)

	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("select * from itemtag")
	checkErr(err)
	raws, err := stmt.Query()
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}
	var tagid, tagname, tagdesc string
	//var viewblogs  []model.OnViewBlog{}
	viewtags := []model.OnViewTag{}
	defer raws.Close()
	for  {
		if raws.Next() {
			err = raws.Scan(&tagid, &tagname, &tagdesc)
			checkErr(err)
			tmpviewtag := model.OnViewTag{
				TagId:     tagid,
				TagName:   tagname,
				TagDesc: tagdesc,
			}
			viewtags = append(viewtags, tmpviewtag)
		} else{
			break
		}
	}
	//fmt.Println(len(viewtags))
	if len(viewtags) == 0 {
		return true, "find no tags" , nil
	}
	//fmt.Println(viewtags)
	return true, "", viewtags
}


func DBgetagbyname(tagname string )(bool, string, model.OnViewTag) {
	db, err := sql.Open("mysql", DBURL)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("select * from itemtag where tagname = ?")
	checkErr(err)
	raws, err := stmt.Query(tagname)
	checkErr(err)
	viewtag := model.OnViewTag{}
	if err != nil {
		return false, err.Error(), viewtag
	}

	defer raws.Close()
	for  {
		if raws.Next() {
			err = raws.Scan(&viewtag.TagId, &viewtag.TagName, &viewtag.TagDesc)
			checkErr(err)
		} else{
			break
		}
	}
	return true, "", viewtag
}