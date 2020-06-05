# api

## 1、数据库连接
```数据库连接```
```markdown
dsn := "root:root@tcp(127.0.0.1:3306)/spiders?charset=utf8mb4&parseTime=True&loc=Local"
数据库名：spiders
```

## 2、登录

```请求接口```

```markdown
http://127.0.0.1:8081/api/user/login
```

```响应数据```

```markdown
# 登录成功
{
	"code":1000,
	"message":"success"
}
# 登录失败
{
	"code":1001,
	"message":"failed"
}
```

## 3、文章列表展示

```响应数据```

```json
[{
	"article_id": "1",
	"article_class": "博客全站导航",
	"article_class_url": "https://www.cnblogs.com/songzhixue/",
	"article_detail_url": "https://www.cnblogs.com/songzhixue/p/10717975.html",
	"article_title": "GIT版本控制工具",
	"article_release_time": "2020-05-30 14-22-09",
	"article_author": "Lz12Code",
	"article_focus_url": "https://images.cnblogs.com/cnblogs_com/songzhixue/1487435/o_1.jpg",
	"article_content": "Git 是一个开源的分布式版本控制系统，用于敏捷高效地处理任何或小或大的项目。Git 是 Linus Torvalds 为了帮助管理 Linux 内核开发而开发的一个开放源码的版本控制软件。Git 与常用的版本控制工具 CVS, Subversion 等不同，它采用了分布式版本库的方式，不必服务器端软件支持。",
	"article_views": "9999",
	"article_comments": "8888",
	"article_tags": [{
			"tag_id": "1",
			"tag": "GO",
			"tag_class_url": "https://github.com/Lz12Code/go-blog"
		},
		{
			"tag_id": "2",
			"tag": "Python",
			"tag_class_url": "https://www.cnblogs.com/songzhixue/"
		},
		{
			"tag_id": "3",
			"tag": "Linux",
			"tag_class_url": "https://www.cnblogs.com/songzhixue/p/11145760.html"
		}
	]
}, ]
```

## 4、轮播图

```响应数据```

```json
[  {
    "carousel_src":
      "http://www.yyyweb.com/demo/slice-box/images/2.jpg",
    "carousel_id": "1",
    "carousel_caption": "时间就是金钱",
    "carousel_href": "https://mp.weixin.qq.com/s/CiUGkKNKYsQyVPtsRv6z4Q"
  },
  {
    "carousel_src":
      "http://www.yyyweb.com/demo/slice-box/images/4.jpg",
    "carousel_id": "2",
    "carousel_caption": "风光壁纸-缤纷植物图",
    "carousel_href": "https://mp.weixin.qq.com/s/CiUGkKNKYsQyVPtsRv6z4Q"
  },
  {
    "carousel_src":
      "http://www.yyyweb.com/demo/slice-box/images/7.jpg",
    "carousel_id": "3",
    "carousel_caption": "做个有良知的人",
    "carousel_href": "https://mp.weixin.qq.com/s/CiUGkKNKYsQyVPtsRv6z4Q"
  }
]
```



