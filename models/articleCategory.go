package models

import (
	"time"
)

type ArticleCategory struct {
	Id           int       `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	CategoryName string    `json:"category_name" gorm:"default 'NULL' CHAR(20) 'category_name'"`
	CreateTime   time.Time `json:"create_time" gorm:"default 'NULL' TIMESTAMP 'create_time'"`
	UpdateTime   time.Time `json:"update_time" gorm:"default 'NULL' TIMESTAMP 'update_time'"`
}

/*
文章分类表
CREATE TABLE `article_categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`category_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
