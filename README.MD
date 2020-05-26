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




