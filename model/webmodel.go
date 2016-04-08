package model

type LogUser struct {
	Name string `form:"name" binding:"Required"`
	Pwd  string `form:"passwd" binding:"Required"`
}

type UserStatus struct {
	Name      string
	Sex       string
	Authority string
	Online    bool
}

type RegistUser struct {
	Name     string
	Sex      string
	Password string
	Email    string
}

type OnViewBlog struct {
	BlogId     string
	BlogName   string
	WriterName string
	Content    interface{}
	Submittime string
	Tag string //tagname
}

type OnDetailBlog struct {
	BlogName   string
	WriterName string
	Content    interface{}
	Submittime string
	Updatetime string
	BlogId     string
	Tag string //tagname
	//类别，热度等
}

type OnDetailComment struct {
	CommentId     string
	CommenterName string
	BlogName      string
	SuperId       string
	Content       interface{}
	CmtTime       string
}

type OnViewTag struct {
	TagId	string	`json:"tagid,omitempty"`
	TagName string	`json:"tagname"`
	TagDesc string	`json:"tagdesc"`
}
