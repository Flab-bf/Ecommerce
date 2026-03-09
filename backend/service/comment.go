package service

import (
	dao2 "ecommerce/backend/dao"
	"ecommerce/backend/model"
	"ecommerce/backend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
	"time"
)

func GetProductComment(pid int, ctx *app.RequestContext) (*[]model.Comment, error) {
	info, err := dao2.GetProductComment(pid)
	if err != nil {
		return nil, err
	}
	info, err = dao2.GetReply(info)
	if err != nil {
		return nil, err
	}
	//检验token是否合法登录，获取点踩信息
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader != "" {

		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		}

		myClaims, er := utils.ParseAccessToken(authHeader)
		if er == nil && myClaims.Uid != 0 {
			ok, e := dao2.IsLegalUser(myClaims.Uid)
			if ok && e == nil {
				info, err = dao2.IsPraise(myClaims.Uid, info)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return info, nil
}

func Comment(cmt *model.Comment, ctx *app.RequestContext) (int64, error) {
	uid, _ := ctx.Get("uid")
	uidInt := uid.(int)
	cmt.UserId = uidInt
	cmt.PraiseCount = 0
	now := time.Now()
	cmt.PublishTime = now.Format("2006-01-02 15:04:05")
	cmt, err := dao2.CommentGetUmsg(cmt)
	if err != nil {
		return 0, err
	}
	id, err := dao2.Comment(cmt)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func Reply(cmt *model.Comment, ctx *app.RequestContext) (int64, error) {
	uid, _ := ctx.Get("uid")
	uidInt := uid.(int)
	cmt.UserId = uidInt
	cmt.PraiseCount = 0
	cmt.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	cmt, err := dao2.CommentGetUmsg(cmt)
	if err != nil {
		return 0, err
	}
	id, err := dao2.Reply(cmt)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func DeleteComment(cid int) error {
	return dao2.Delete(cid)
}

func UpdateComment(cmt *model.Comment) error {
	cmt.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	return dao2.Update(cmt)
}

func IsPraised(pid int64, ipd int, uid int) error {
	return dao2.Praise(pid, ipd, uid)
}
