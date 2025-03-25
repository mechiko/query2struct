package output

type SqliteSequence struct {
    Name	string	`db:"name"` 
    Seq	string	`db:"seq"` 
}

type SqliteSequenceSlice []*SqliteSequence

func (r *SqliteSequence) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO sqlite_sequence (name,seq) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.Name,r.Seq)
	return err
}

func (r *SqliteSequence) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE sqlite_sequence SET name=?,seq=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.Name,r.Seq,)
	return err
}

func (r *SqliteSequence) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE sqlite_sequence WHERE =?;`
	_, err = db.Exec(sql, )
	return err
}

func (r *SqliteSequence) Find(db *sqlx.DB) (err error) {
	if  == 0 {
		return fmt.Errorf(" = 0")
	}
	sql := `select * from sqlite_sequence WHERE =?;`
	err = db.Get(r, sql, )
	return err
}

func (r *SqliteSequence) All(db *sqlx.DB) (out SqliteSequenceSlice, err error) {
	out = make(SqliteSequenceSlice, 0)
	sql := `select * from sqlite_sequence order by ;`
	err = db.Select(&out, sql)
	return out, err
}
