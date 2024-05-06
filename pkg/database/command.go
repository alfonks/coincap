package database

import "gorm.io/gorm"

func wrapGromDB(gormDB *gorm.DB) DBItf {
	return &db{
		conn: gormDB,
	}
}

func (d *db) Exec(sql string, values ...interface{}) DBItf {
	return wrapGromDB(d.conn.Exec(sql, values...))
}

func (d *db) Raw(sql string, values ...interface{}) DBItf {
	return wrapGromDB(d.conn.Raw(sql, values...))
}

func (d *db) Scan(dest interface{}) DBItf {
	return wrapGromDB(d.conn.Scan(dest))
}

func (d *db) Begin() DBItf {
	return wrapGromDB(d.conn.Begin())
}

func (d *db) Commit() DBItf {
	return wrapGromDB(d.conn.Commit())
}

func (d *db) Rollback() DBItf {
	return wrapGromDB(d.conn.Rollback())
}

func (d *db) Error() error {
	return d.conn.Error
}
