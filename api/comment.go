package api

import (
	"context"
	"ecommerce/model"
	"ecommerce/service"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func Comment(c context.Context, ctx *app.RequestContext) {
	var cmt model.Comment
	err := ctx.Bind(&cmt)
	if err != nil {
		ctx.JSON(400, utils.ErrorResponse(10002, "参数错误"))
		return
	}
	if len(cmt.Content) > 20000 {
		ctx.JSON(403, utils.ErrorResponse(40001, "字数超限"))
		return
	}
	id, err := service.Comment(&cmt, ctx)
	if err != nil || id == 0 {
		ctx.JSON(403, utils.ErrorResponse(40001, "评论失败"))
		return
	}
	ctx.JSON(200, utils.SuccessResponse(id))
}

func Reply(c context.Context, ctx *app.RequestContext) {

}

func GetComment(c context.Context, ctx *app.RequestContext) {
	pidString := ctx.Param("product_id")
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		ctx.JSON(500, utils.ErrorResponse(10002, "意外错误"))
		return
	}
	info, err := service.GetProductComment(pid, ctx)
	if err != nil {
		ctx.JSON(404, utils.ErrorResponse(40003, "评论查找失败"))
		return
	}
	ctx.JSON(200, utils.SuccessResponse(info))
}

func DeleteComment(c context.Context, ctx *app.RequestContext) {
	cid := ctx.Param("comment_id")
	cidInt, err := strconv.Atoi(cid)
	if err != nil {
		ctx.JSON(500, utils.ErrorResponse(10002, "意外错误"))
		return
	}
	err = service.DeleteComment(cidInt)
	if err != nil {
		ctx.JSON(40004, "删除失败")
		return
	}
	ctx.JSON(200, utils.SuccessResponse(10000))
}

func UpdateComment(c context.Context, ctx *app.RequestContext) {

}

func PraiseOrNot(c context.Context, ctx *app.RequestContext) {}
