package models


// 文章表
type ArticleList struct {
	Id           int    `json:"id" gorm:"not null pk autoincr comment('id') INT(5) 'id'"`
	Title        string `json:"title" gorm:"not null comment('文章标题') VARCHAR(255) 'title'"`
	CreateTime   int64  `json:"create_time" gorm:"default NULL comment('发布时间') BIGINT(20) 'create_time'"`
	UpdateTime   int64  `json:"update_time" gorm:"default NULL comment('修改时间') BIGINT(20) 'update_time'"`
	Status       bool   `json:"status" gorm:"not null comment('状态 0:未发布  1:发布')  'status'"`
	//Status       bool   `json:"status"`
	Md           string `json:"md" gorm:"not null comment('markdown文本') LONGTEXT 'md'"`
	Html         string `json:"html" gorm:"not null comment('markdown编译后的html') LONGTEXT 'html'"`
	CoverAddress string `json:"cover_address" gorm:"default 'NULL' comment('封面地址') VARCHAR(255) 'cover_address'"`
	Author       string `json:"author" gorm:"not null comment('作者') CHAR(20) 'author'"`
	Top          int    `json:"top" gorm:"not null comment('文章置顶') INT(3) 'top'"`
	CategoryId   int    `json:"category_id" gorm:"not null comment('分类id') INT(3) 'category_id'"`
	Summary      string `json:"summary" gorm:"comment('文章摘要') TEXT 'summary'"`
	Views        int    `json:"views" gorm:"default NULL comment('文章浏览量') INT(10) 'views'"`
	//CoverPreview string `json:"cover_preview" gorm:"default 'NULL' comment('封面预览地址') VARCHAR(255) 'cover_preview'"`
}
