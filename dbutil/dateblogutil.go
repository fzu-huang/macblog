package dbutil

//import(
//	"database/sql"
//	"fmt"
//	"github.com/fzu-huang/macblog/model"
//	_ "github.com/go-sql-driver/mysql"
//)


//func DBgetallcontentbydateandlist(date string, page int) (bool, string, []model.OnViewBlog) {
//	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bbs?charset=utf8")

//	checkErr(err)
//	defer db.Close()
//	countper := 9
//	limitindex := (page - 1) * countper
//	stmt, err := db.Prepare("select b.id, b.blogname, b.writername, b.content, b.submittime from blog b, itemtag t where t.tagname = ? and t.tagid = b.tagid order by b.submittime desc limit ?,?")
//	checkErr(err)
//	raws, err := stmt.Query(tagname,limitindex, countper)
//	checkErr(err)
//	if err != nil {
//		return false, err.Error(), nil
//	}
//	var id, blogname, writername, content, submittime string

//	//var viewblogs  []model.OnViewBlog{}
//	viewblogs := make([]model.OnViewBlog, countper)
//	defer raws.Close()
//	for i := 0; i < 9; i++ {
//		if raws.Next() {
//			err = raws.Scan(&id, &blogname, &writername, &content, &submittime)
//			checkErr(err)
//			viewblogs[i] = model.OnViewBlog{
//				BlogId:     id,
//				BlogName:   blogname,
//				WriterName: writername,
//				Content:    template.HTML(content),
//				Submittime: submittime,
//				Tag: tagname,
//			}
//		} else {
//			if i == 0 {
//				return true, "find no blog.", nil
//			}
//			viewblogs = viewblogs[:i]
//			break
//		}
//	}
//	return true, "", viewblogs
//}