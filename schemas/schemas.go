package schemas

import frz "github.com/razshare/frizzante"

var Sql = frz.SqlCreate()

type Account struct {
	Id        string `sql:"varchar(36) not null"`
	Password  string `sql:"varchar(255) not null"`
	CreatedAt int64  `sql:"int not null"`
	UpdatedAt int64  `sql:"int not null"`
}

type Article struct {
	Id        string `sql:"varchar(36) not null"`
	CreatedAt int64  `sql:"int not null"`
	Account   Account
}

type Comment struct {
	Id        string `sql:"varchar(36) not null"`
	CreatedAt int64  `sql:"int not null"`
	Account   Account
}

type ArticleContent struct {
	Id        string `sql:"varchar(36) not null"`
	CreatedAt int64  `sql:"int not null"`
	Article   Article
}

type CommentContent struct {
	Id        string `sql:"varchar(36) not null"`
	CreatedAt int64  `sql:"int not null"`
	Comment   Comment
}
