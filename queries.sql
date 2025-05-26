-- name: SqlVerifyAccount :one
select id
from user_account
where id = ?
  and Password = ?
limit 1;

-- name: SqlAddAccount :exec
insert into user_account(id,
                         display_name,
                         password,
                         created_at,
                         updated_at)
values (?, ?, ?, ?, ?);

-- name: SqlChangeAccount :exec
update user_account
set display_name = ?,
    password     = ?,
    updated_at   = ?
where id = ?;

-- name: SqlAddArticle :exec
insert into article(id,
                    created_at,
                    account_id)
values (?, ?, ?);

-- name: SqlAddArticleContent :exec
insert into article_content(id,
                            created_at,
                            article_id,
                            title,
                            content)
values (?, ?, ?, ?, ?);

-- name: SqlFindArticleContent :one
select content
from article_content
where article_id = ?
order by created_at desc
limit 1;

-- name: SqlRemoveArticle :exec
delete
from article
where id = ?;

-- name: SqlAddComment :exec
insert into comment(id,
                    created_at,
                    account_id,
                    article_id)
values (?, ?, ?, ?);

-- name: SqlAddCommentContent :exec
insert into comment_content(id,
                            created_at,
                            comment_id,
                            content)
values (?, ?, ?, ?);

-- name: SqlFindCommentContent :one
select content
from comment_content
where id = ?
order by created_at desc
limit 1;

-- name: SqlRemoveComment :exec
delete
from comment
where id = ?;

-- name: SqlFindArticles :many
select id,
       created_at,
       account_id
from article
limit ?, ?;

-- name: SqlFindCommentsByArticleId :many
select id,
       created_at,
       account_id,
       article_id
from comment
where article_id = ?
limit ?, ?;

-- name: SqlFindAccounts :many
select id,
       display_name,
       created_at,
       updated_at
from user_account
limit ?, ?;

-- name: SqlFindAccountById :one
select id,
       display_name,
       created_at,
       updated_at
from user_account
where id = ?;
