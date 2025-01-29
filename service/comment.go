package service

import (
	"ecommerce/dao"
	"ecommerce/model"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func GetProductComment(pid int, ctx *app.RequestContext) (*[]model.Comment, error) {
	uid, _ := ctx.Get("uid")
	uidInt := uid.(int)
	info, err := dao.GetProductComment(pid)
	if err != nil {
		return nil, err
	}
	info, err = dao.GetReply(info)
	if err != nil {
		return nil, err
	}
	info, err = dao.IsPraise(uidInt, info)
	if err != nil {
		return nil, err
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
	cmt, err := dao.CommentGetUmsg(cmt)
	if err != nil {
		return 0, err
	}
	id, err := dao.Comment(cmt)
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
	cmt, err := dao.CommentGetUmsg(cmt)
	if err != nil {
		return 0, err
	}
	id, err := dao.Reply(cmt)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func DeleteComment(cid int) error {
	return dao.Delete(cid)
}

func UpdateComment(cmt *model.Comment) error {
	cmt.PublishTime = time.Now().Format("2006-01-02 15:04:05")
	return dao.Update(cmt)
}

func IsPraised(pid int64, ipd int, uid int) error {
	return dao.Praise(pid, ipd, uid)
}
