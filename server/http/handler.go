package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shao1f/user.tag.logic/model"
	"github.com/shao1f/user.tag.logic/util"
)

func tagAdd(c *gin.Context) {
	tagAddReq := &model.AddTagReq{}
	if err := c.BindJSON(tagAddReq); err != nil {
		result := util.Result(util.ErrInternalError)
		c.JSON(result.ErrCode, result)
		return
	}
	if tagAddReq.Name == "" {
		result := util.Result(util.ErrParamError)
		c.JSON(result.ErrCode, result)
		return
	}
	id, err := svc.AddTag(c.Request.Context(), tagAddReq.Name)
	if err != nil {
		result := util.Result(err)
		c.JSON(result.ErrCode, result)
		return
	}

	succ := util.Result(nil)
	resp := &model.AddTagResp{
		BaseResponse: succ,
		Data: model.UserTag{
			TagID:   id,
			TagName: tagAddReq.Name,
		},
	}
	c.JSON(succ.ErrCode, resp)
}

func tagSearch(c *gin.Context) {
	key := c.Query("key")
	// key := c.Request.URL.Query().Get("key")
	if key == "" {
		result := util.Result(util.ErrParamError)
		c.JSON(result.ErrCode, result)
		return
	}
	tags, err := svc.SearchTag(c.Request.Context(), key)
	if err != nil {
		result := util.Result(err)
		c.JSON(result.ErrCode, result)
		return
	}
	succ := util.Result(util.ErrIsNil)
	resp := model.SearchTagResp{
		BaseResponse: succ,
		Data:         tags,
	}
	c.JSON(succ.ErrCode, resp)
}

func entityLink(c *gin.Context) {

}

func entitySearch(c *gin.Context) {

}
