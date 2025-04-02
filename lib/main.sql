DROP DATABASE IF EXISTS `forum`;

CREATE DATABASE `forum`;

USE `forum`;

CREATE TABLE IF NOT EXISTS
    `Account`
(
    `Id`        varchar(36)  NOT NULL,
    `Password`  varchar(255) NOT NULL,
    `CreatedAt` int          NOT NULL,
    `UpdatedAt` int          NOT NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS
    `Article`
(
    `Id`        varchar(36) NOT NULL,
    `CreatedAt` int         NOT NULL,
    `Account`   varchar(36) NOT NULL,
    PRIMARY KEY (`Id`),
    KEY `Account` (`Account`),
    CONSTRAINT `Article_ibfk_1` FOREIGN KEY (`Account`) REFERENCES `Account` (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS
    `ArticleContent`
(
    `Id`        varchar(36) NOT NULL,
    `CreatedAt` int         NOT NULL,
    `Article`   varchar(36) NOT NULL,
    PRIMARY KEY (`Id`),
    KEY `Article` (`Article`),
    CONSTRAINT `ArticleContent_ibfk_1` FOREIGN KEY (`Article`) REFERENCES `Article` (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS
    `Comment`
(
    `Id`        varchar(36) NOT NULL,
    `CreatedAt` int         NOT NULL,
    `Account`   varchar(36) NOT NULL,
    PRIMARY KEY (`Id`),
    KEY `Account` (`Account`),
    CONSTRAINT `Comment_ibfk_1` FOREIGN KEY (`Account`) REFERENCES `Account` (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS
    `CommentContent`
(
    `Id`        varchar(36) NOT NULL,
    `CreatedAt` int         NOT NULL,
    `Comment`   varchar(36) NOT NULL,
    PRIMARY KEY (`Id`),
    KEY `Comment` (`Comment`),
    CONSTRAINT `CommentContent_ibfk_1` FOREIGN KEY (`Comment`) REFERENCES `Comment` (`Id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;