package service

import "github.com/shao1f/user.tag.logic/dao"

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
