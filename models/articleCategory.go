package models


// 定义模型类
type ArticleCategory struct {
	Id           int    `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	CategoryName string `json:"category_name" gorm:"default 'NULL' CHAR(20) 'category_name'"`
	CreateTime   int64  `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime   int64  `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}

/*
文章分类表
CREATE TABLE `article_categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` char(20) CHARACTER SET utf8,
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4;
*/
