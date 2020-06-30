package dao

type Dao struct {
	mysql *MySqlDB
}

func New() *Dao {
	db := NewMysqlDB()
	return &Dao{
		db,
	}
}

func (d *Dao) Close() error {
	return nil
}
