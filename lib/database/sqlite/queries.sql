-- name: VerifyAccount :one
select id from user_account where id = ? and password = ? limit 1;

-- name: AddAccount :exec
insert into user_account(id, display_name, password, created_at, updated_at) values (?, ?, ?, ?, ?);

-- name: AddArticle :exec
insert into article(id, title, content, created_at, account_id) values (?, ?, ?, ?, ?);

-- name: RemoveArticle :exec
delete from article where id = ?;

-- name: FindArticles :many
select id, title, content, created_at, account_id from article limit ? offset ?;

-- name: FindAccountById :one
select id, display_name, created_at, updated_at from user_account where id = ?;
