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
    created_at int         not null,
    account_id varchar(36) not null,
    primary key (id),
    foreign key (account_id) references user_account (id)
);

create table if not exists article_content
(
    id         varchar(36) not null,
    created_at int         not null,
    article_id varchar(36) not null,
    title      varchar(36) not null,
    content    text        not null,
    primary key (id),
    foreign key (article_id) references article (id)
);

create table if not exists comment
(
    id         varchar(36) not null,
    created_at int         not null,
    account_id varchar(36) not null,
    article_id varchar(36) not null,
    primary key (id),
    foreign key (account_id) references user_account (id),
    foreign key (article_id) references article (id)
);

create table if not exists comment_content
(
    id         varchar(36) not null,
    created_at int         not null,
    comment_id varchar(36) not null,
    content    text        not null,
    primary key (id),
    foreign key (comment_id) references comment (id)
);