package sql

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

var Sql = frz.SqlCreate()

func find(query string, props ...any) (next func(dest ...any) bool, close func()) {
	return frz.SqlFind(Sql, query, props...)
}

func run(query string, props ...any) *sql.Result {
	return frz.SqlExecute(Sql, query, props...)
}

// VerifyAccount verifies that the combination of id and password are exists.
func VerifyAccount(id string, password string) bool {
	fetch, closeFetch := find(
		"select AccountId from `Account` where AccountId = ? and Password = ? limit 1",
		id, password,
	)
	defer closeFetch()
	return fetch(&id)
}

// AddAccount adds an account.
func AddAccount(id string, displayName string, password string) {
	now := time.Now().Unix()
	run(
		"insert into `Account`(AccountId,DisplayName,Password,CreatedAt,UpdatedAt) values(?,?,?,?,?)",
		id, displayName, password, now, now,
	)
}

// ChangeAccount changes properties of an account.
func ChangeAccount(
	accountId string,
	displayName string,
	password string,
) {
	updatedAt := time.Now().Unix()
	run(
		"update `Account` set DisplayName = ?, Password = ?, UpdatedAt = ? where AccountId = ?",
		displayName, password, updatedAt, accountId,
	)
}

// AddArticle adds an article and returns its id.
func AddArticle(accountId string) string {
	uuidLocal, uuidError := uuid.NewV4()
	if nil != uuidError {
		lib.Failure(uuidError)
		return ""
	}

	articleId := uuidLocal.String()
	createdAt := time.Now().Unix()
	run(
		"insert into `Article`(ArticleId,CreatedAt,Account) values(?,?,?)",
		articleId, createdAt, accountId,
	)

	return articleId
}

// AddArticleContent adds content to an article.
func AddArticleContent(articleId string, content string) string {
	id, idError := uuid.NewV4()
	if nil != idError {
		lib.Failure(idError)
		return ""
	}

	articleContentId := id.String()
	createdAt := time.Now().Unix()
	run(
		"insert into `ArticleContent`(ArticleContentId,CreatedAt,Account,Content) values(?,?,?,?)",
		articleContentId, createdAt, articleId, content,
	)

	return articleContentId
}

// FindArticleContent finds the content of an article.
func FindArticleContent(articleId string) (content string) {
	fetch, closeFetch := find(
		"select `Content` from `ArticleContent` where `ArticleId` = ? order by CreatedAt desc limit 1",
		articleId,
	)
	defer closeFetch()
	fetch(content)
	return
}

// RemoveArticle removes an article.
func RemoveArticle(articleId string) {
	run(
		"delete from `Article`  where ArticleId = ?",
		articleId,
	)
}

// AddComment adds a comment to an article and returns its id.
func AddComment(accountId string, articleId string) string {
	uuidLocal, uuidError := uuid.NewV4()
	if nil != uuidError {
		lib.Failure(uuidError)
		return ""
	}

	commentId := uuidLocal.String()
	createdAt := time.Now().Unix()
	run(
		"insert into `Comment`(CommentId,CreatedAt,AccountId) values(?,?,?)",
		commentId, createdAt, accountId,
	)

	return commentId
}

// AddCommentContent adds content to a comment.
func AddCommentContent(commentId string, articleId string, content string) string {
	id, idError := uuid.NewV4()
	if nil != idError {
		lib.Failure(idError)
		return ""
	}

	commentContentId := id.String()
	createdAt := time.Now().Unix()
	run(
		"insert into `CommentContent`(CommentContentId,CreatedAt,CommentId,AccountId,Content) values(?,?,?,?)",
		commentContentId, createdAt, commentId, articleId, content,
	)

	return commentContentId
}

// FindCommentContent finds the content of a comment.
func FindCommentContent(commentId string) (content string) {
	fetch, closeFetch := find(
		"select `Content` from `CommentContent` where `CommentContentId` = ? order by CreatedAt desc limit 1",
		commentId,
	)
	defer closeFetch()
	fetch(content)
	return
}

// RemoveComment removes a comment.
func RemoveComment(commentId string) {
	run(
		"delete from `Comment`  where CommentId = ?",
		commentId,
	)
}
