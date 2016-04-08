package ctl

import (
	//	"github.com/go-sql-driver/mysql"

	"github.com/fzu-huang/macblog/dbutil"
	"github.com/fzu-huang/macblog/model"
)

func LogCheck(loguser model.LogUser) (bool, model.UserStatus) {
	return dbutil.DBLogCheck(loguser)
}

func RegCheck(reguser model.RegistUser) (bool, string) {
	return dbutil.DBRegCheck(reguser)
}
