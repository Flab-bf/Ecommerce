package dao

import (
	"ecommerce/model"
	"errors"
	"gorm.io/gorm"
)

func IsPraise(uid int, info *[]model.Comment) (*[]model.Comment, error) {
	for index, pc := range *info {
		result := DB.Model(&model.Praise{}).Select("is_praised").
			Where("user_id=? AND post_id=? AND parent_id is NULL", uid, pc.PostId).
			First(&(*info)[index].IsPraised)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		for id, pct := range (*info)[index].Reply {
			result = DB.Model(&model.Praise{}).Select("is_praised").
				Where("user_id=? AND post_id=? AND parent_id=?", uid, pct.PostId, pc.PostId).
				First(&(*info)[index].Reply[id].IsPraised)
			if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, result.Error
			}
		}
	}
	return info, nil
}

func GetProductComment(pid int) (*[]model.Comment, error) {
	var info []model.Comment
	result := DB.Model(&model.Comment{}).Where("product_id=? AND  parent_id IS NULL", pid).Find(&info)
	if result.Error != nil {
		return nil, result.Error
	}
	return &info, nil
}

func GetReply(info *[]model.Comment) (*[]model.Comment, error) {
	for index, comment := range *info {
		result := DB.Model(&model.Comment{}).Where("parent_id=?", comment.PostId).Find(&(*info)[index].Reply)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return info, nil
}

func CommentGetUmsg(cmt *model.Comment) (*model.Comment, error) {
	result := DB.Model(&model.UserMassage{}).Select("nick_name").Where("uid=?", cmt.UserId).First(&cmt.NickName)
	if result.Error != nil {
		return nil, result.Error
	}
	result = DB.Model(model.UserMassage{}).Select("avatar").Where("uid=?", cmt.UserId).First(&cmt.Avatar)
	if result.Error != nil {
		return nil, result.Error
	}
	return cmt, nil
}
func Comment(cmt *model.Comment) (int64, error) {

	result := DB.Model(&model.Comment{}).Omit("parent_id").Create(cmt)
	if result.Error != nil {
		return 0, result.Error
	}
	return cmt.PostId, nil
}

func Reply(cmt *model.Comment) (int64, error) {
	result := DB.Model(&model.Comment{}).Create(cmt)
	if result.Error != nil {
		return 0, result.Error
	}
	return cmt.PostId, nil
}

func Delete(cid int) error {
	result := DB.Model(&model.Comment{}).Where("post_id=?", cid).Delete(&model.Comment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Update(cmt *model.Comment) error {
	result := DB.Model(&model.Comment{}).Where("post_id=?", cmt.PostId).Updates(cmt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Praise(pid int64, ipd int, uid int) error {
	var prs model.Praise
	result := DB.Model(&model.Comment{}).Select("parent_id,product_id").
		Where("post_id=?", pid).First(&prs)
	if result.Error != nil {
		return result.Error
	}
	prs.UserId = uid
	prs.PostId = pid
	prs.IsPraised = ipd
	var count int64
	result = DB.Model(&model.Praise{}).Where("post_id=?", pid).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count == 0 {
		if prs.ParentId == 0 {
			result = DB.Model(&model.Praise{}).Omit("parent_id").Create(&prs)
		} else {
			result = DB.Model(&model.Praise{}).Create(&prs)
		}
	} else {
		if prs.ParentId == 0 {
			result = DB.Model(&model.Praise{}).Omit("parent_id").Where("post_id", pid).Updates(&prs)
		} else {
			result = DB.Model(&model.Praise{}).Where("post_id", pid).Updates(&prs)
		}
	}
	if result.Error != nil {
		return result.Error
	}
	if ipd == 1 {
		rsut := DB.Model(&model.Comment{}).Where("post_id=?", prs.PostId).Update("praise_count", gorm.Expr("praise_count+?", 1))
		if result.Error != nil || rsut.Error != nil {
			return errors.New("error update")
		}
	}
	return nil
}
