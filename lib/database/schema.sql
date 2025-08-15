create table if not exists user_account
(
    id           varchar(36)  not null,
    display_name varchar(255) not null unique,
    password     varchar(255) not null,
    created_at   int          not null,
    updated_at   int          not null,
    primary key (id)
);

create table if not exists article
(
    id         varchar(36) not null,
    title      varchar(36) not null,
    content    text        not null,
    created_at int         not null,
    account_id varchar(36) not null
);