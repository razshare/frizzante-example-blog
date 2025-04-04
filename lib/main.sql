drop database if exists `forum`;

create database `forum`;

use `forum`;

create table if not exists
    `Account`
(
    `AccountId`   varchar(36)  not null,
    `DisplayName` varchar(255) not null unique,
    `Password`    varchar(255) not null,
    `CreatedAt`   int          not null,
    `UpdatedAt`   int          not null,
    primary key (`AccountId`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

create table if not exists
    `Article`
(
    `ArticleId` varchar(36) not null,
    `CreatedAt` int         not null,
    `AccountId` varchar(36) not null,
    primary key (`ArticleId`),
    foreign key (`AccountId`) references `Account` (`AccountId`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

create table if not exists
    `ArticleContent`
(
    `ArticleContentId` varchar(36) not null,
    `CreatedAt`        int         not null,
    `ArticleId`        varchar(36) not null,
    `Content`          text        not null,
    primary key (`ArticleContentId`),
    foreign key (`ArticleId`) references `Article` (`ArticleId`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

create table if not exists
    `Comment`
(
    `CommentId` varchar(36) not null,
    `CreatedAt` int         not null,
    `AccountId` varchar(36) not null,
    `ArticleId` varchar(36) not null,
    primary key (`CommentId`),
    foreign key (`AccountId`) references `Account` (`AccountId`),
    foreign key (`ArticleId`) references `Article` (`ArticleId`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;

create table if not exists
    `CommentContent`
(
    `CommentContentId` varchar(36) not null,
    `CreatedAt`        int         not null,
    `CommentId`        varchar(36) not null,
    `Content`          text        not null,
    primary key (`CommentContentId`),
    foreign key (`CommentId`) references `Comment` (`CommentId`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_0900_ai_ci;