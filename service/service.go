package service

import (
	"context"
	"github.com/shao1f/user.tag.logic/dao"
	"github.com/shao1f/user.tag.logic/model"
	"github.com/shao1f/user.tag.logic/util"
)

type Service struct {
	// dao: db
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New(),
	}
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}

func (s *Service) AddTag(ctx context.Context, name string) (int, error) {
	queryTagID, err := s.dao.TagGet(ctx, name)
	if err != nil {
		return 0, util.ErrInsertTagError
	}
	if queryTagID != 0 {
		return queryTagID, nil
	}
	// 未查询到name对应的tag，进行插入
	tagID, err := s.dao.TagAdd(ctx, name)
	if err != nil {
		return 0, util.ErrInsertTagError
	}
	esTag := &model.UserTag{
		TagID:   tagID,
		TagName: name,
	}
	go s.dao.UploadToEs(ctx, esTag)
	return tagID, nil
}

func (s *Service) SearchTag(ctx context.Context, key string) ([]*model.UserTag, error) {
	return s.dao.SearchFromEs(ctx, key)
}
