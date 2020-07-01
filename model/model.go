package model

import (
	"bytes"
	"encoding/json"
)

type BaseResponse struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type UserTag struct {
	TagID   int    `gorm:"column:id" json:"tag_id"`
	TagName string `gorm:"column:name" json:"tag_name"`
}

func (t *UserTag) MustToJSON() string {
	bs, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

type EntityTag struct {
	LinkID   int `gorm:"column:id" json:"-"`
	EntityID int `gorm:"column:entity_id" json:"entity_id"`
	TagID    int `gorm:"column:tag_id" json:"tag_id"`
}

type AddTagReq struct {
	Name string `json:"name"`
}

type AddTagResp struct {
	BaseResponse
	Data UserTag `json:"data"`
}

type NormalMap map[string]interface{}

func (this *NormalMap) MustToJSONBytesBuffer() *bytes.Buffer {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(this); err != nil {
		panic(err)
	}
	return &buf
}

type SearchTagResp struct {
	BaseResponse
	Data []*UserTag `json:"data"`
}
