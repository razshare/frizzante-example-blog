-- name: VerifyAccount :one
select id
from user_account
where id = ?
  and password = ?
limit 1;

-- name: AddAccount :exec
insert into user_account(id,
                         display_name,
                         password,
                         created_at,
                         updated_at)
values (?, ?, ?, ?, ?);

-- name: ChangeAccount :exec
update user_account
set display_name = ?,
    password     = ?,
    updated_at   = ?
where id = ?;

-- name: AddArticle :exec
insert into article(id,
                    title,
                    created_at,
                    account_id)
values (?, ?, ?, ?);

-- name: AddArticleContent :exec
insert into article_content(id,
                            created_at,
                            article_id,
                            content)
values (?, ?, ?, ?);

-- name: FindArticleContent :one
select content
from article_content
where article_id = ?
order by created_at desc
limit 1;

-- name: RemoveArticle :exec
delete
from article
where id = ?;

-- name: AddComment :exec
insert into comment(id,
                    created_at,
                    account_id,
                    article_id)
values (?, ?, ?, ?);

-- name: AddCommentContent :exec
insert into comment_content(id,
                            created_at,
                            comment_id,
                            content)
values (?, ?, ?, ?);

-- name: FindCommentContent :one
select content
from comment_content
where id = ?
order by created_at desc
limit 1;

-- name: RemoveComment :exec
delete
from comment
where id = ?;

-- name: FindArticles :many
select article.id,
       article.title,
       article.created_at,
       article.account_id,
       article_content.content
from article
inner join article_content on article.id = article_content.article_id
order by article.created_at desc
limit ? offset ?;

-- name: FindCommentsByArticleId :many
select id,
       created_at,
       account_id,
       article_id
from comment
where article_id = ?
limit ? offset ?;

-- name: FindAccounts :many
select id,
       display_name,
       created_at,
       updated_at
from user_account
limit ? offset ?;

-- name: FindAccountById :one
select id,
       display_name,
       created_at,
       updated_at
from user_account
where id = ?;
