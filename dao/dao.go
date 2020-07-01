package dao

import (
	"context"
	"log"

	"github.com/shao1f/user.tag.logic/model"
)

type Dao struct {
	mysql *MySqlDB
	es    *Elastic
}

func New() *Dao {
	db := NewMysqlDB()
	es := NewElastic()
	return &Dao{
		db,
		es,
	}
}

func (d *Dao) Close() error {
	return nil
}

func (d *Dao) TagGet(ctx context.Context, name string) (int, error) {
	tag, err := d.mysql.TagQuery(name)
	if err != nil {
		return 0, err
	}
	return tag.TagID, nil
}

func (d *Dao) TagAdd(ctx context.Context, name string) (int, error) {
	return d.mysql.TagInsert(name)
}

func (d *Dao) UploadToEs(ctx context.Context, tag *model.UserTag) {
	d.es.UploadToES(ctx, tag)
}

func (d *Dao) SearchFromEs(ctx context.Context, key string) ([]*model.UserTag, error) {
	tags, err := d.es.SearchFromES(ctx, key)
	if err != nil {
		log.Print("search err", err)
		return nil, err
	}
	return tags, nil
}
