package dbutil

import(
	"database/sql"
	"github.com/fzu-huang/macblog/model"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
)


func DBgetallcontentbydateandlist(year,month string, page int) (bool, string, []model.OnViewBlog) {
	if year == ""{
		return DBgetallcontentbylist(page)
	}
	
	db, err := sql.Open("mysql", DBURL)

	checkErr(err)
	defer db.Close()
	countper := 9
	limitindex := (page - 1) * countper
	
	queryurl := ""
	var raws *sql.Rows
	if month ==""{
		queryurl = "select b.id, b.blogname, b.writername, b.content, b.submittime from blog b where b.year=? order by b.submittime desc limit ?,?"
		stmt , err:= db.Prepare(queryurl)
		checkErr(err)
		raws, err = stmt.Query(year,limitindex,countper)
	}else{
		queryurl = "select b.id, b.blogname, b.writername, b.content, b.submittime from blog b where b.year=? and b.month=? order by b.submittime desc limit ?,?"
		stmt , err:= db.Prepare(queryurl)
		checkErr(err)
		raws, err = stmt.Query(year,month,limitindex,countper)
	}
	
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

func DBgetalldate()(bool, string, []model.DateBlogCount){
	db, err := sql.Open("mysql", DBURL)

	checkErr(err)
	defer db.Close()

	queryurl := "select year,month,count(*) from blog group by year,month ORDER BY blog.submittime desc"
		stmt , err:= db.Prepare(queryurl)
		checkErr(err)
		raws, err := stmt.Query()
	checkErr(err)
	if err != nil {
		return false, err.Error(), nil
	}

	datecounts := []model.DateBlogCount{}
	var datecount model.DateBlogCount
	defer raws.Close()
	for  {
		if raws.Next() {
			err = raws.Scan(&datecount.Year, &datecount.Month, &datecount.Count)
			checkErr(err)
			datecounts = append(datecounts, datecount)
		} else{
			break
		}
	}
	
	return true, "", datecounts
}