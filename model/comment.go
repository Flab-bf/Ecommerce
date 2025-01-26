package model

type Comment struct {
	PostId      int64  `gorm:"primaryKey;type:bigint" json:"postId"`
	ParentId    int64  `gorm:"type:bigint" json:"parentId"`
	PublishTime string `gorm:"type:varchar(50)" json:"publishTime"`
	Content     string `gorm:"type:text" json:"content"`
	UserId      int    `gorm:"type:int" json:"userId"`
	Avatar      string `gorm:"type:varchar(255)" json:"avatar"`
	NickName    string `gorm:"type:varchar(255)" json:"nickName"`
	PraiseCount int    `gorm:"type:int" json:"praiseCount"`
	IsPraised   int    `gorm:"type:int" json:"isPraised"` //0:未处理;1:点赞;2:点踩
	ProductId   int    `gorm:"type:int" json:"productId"`
}

type Praise struct {
	PostId    int64 `gorm:"primaryKey;type:bigint" json:"postId"`
	ProductId int   `gorm:"type:int" json:"productId"`
	UserId    int   `gorm:"type:int" json:"userId"`
	ParentId  int64 `gorm:"type:bigint" json:"parentId"`
	IsPraised int   `gorm:"type:int" json:"isPraised"` //0:未处理;1:点赞;2:点踩
}
