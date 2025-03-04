package schemas

type User struct {
	Id        string `sql:"varchar(36) not null primary key"`
	Password  string `sql:"varchar(255) not null"`
	CreatedAt int64  `sql:"int not null"`
	UpdatedAt int64  `sql:"int not null"`
}
