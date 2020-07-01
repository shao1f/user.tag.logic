package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shao1f/user.tag.logic/model"
	"log"
)

var (
	useTagTableName = "user_tag"
	entityTableName = "entity_tag"
)

type MySqlDB struct {
	db *gorm.DB
}

func NewMysqlDB() *MySqlDB {
	// init db instance
	db, err := gorm.Open("mysql", "root:Syf3344521jsy@tcp(rm-bp1kr713l8l4px1814o.mysql.rds.aliyuncs.com:3306)/user_tag?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		log.Fatalf("init mysql error,err %v", err)
	}
	return &MySqlDB{db: db}
}

func (m *MySqlDB) TagQuery(name string) (*model.UserTag, error) {
	tag := &model.UserTag{}
	db := m.db.Where("name = ?", name).First(tag)
	if db.Error != nil && !db.RecordNotFound() {
		return nil, db.Error
	}
	return tag, nil
}

func (m *MySqlDB) TagInsert(name string) (int, error) {
	tag := model.UserTag{
		TagName: name,
	}
	db := m.db.Create(&tag)
	if db.Error != nil {
		return 0, db.Error
	}
	return tag.TagID, nil
}
