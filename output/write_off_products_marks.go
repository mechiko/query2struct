package output

type WriteOffProductsMarks struct {
    Id	int64	`db:"id"` 
    IdWriteOffProducts	int	`db:"id_write_off_products"` 
    IdWriteOffProductsBoxes	int	`db:"id_write_off_products_boxes"` 
    Mark	string	`db:"mark"` 
}

type WriteOffProductsMarksSlice []*WriteOffProductsMarks

func (r *WriteOffProductsMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO write_off_products_marks (id_write_off_products,id_write_off_products_boxes,mark) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.IdWriteOffProductsBoxes,r.Mark)
	return err
}

func (r *WriteOffProductsMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE write_off_products_marks SET id_write_off_products=?,id_write_off_products_boxes=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.IdWriteOffProductsBoxes,r.Mark,r.Id)
	return err
}

func (r *WriteOffProductsMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE write_off_products_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *WriteOffProductsMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from write_off_products_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *WriteOffProductsMarks) All(db *sqlx.DB) (out WriteOffProductsMarksSlice, err error) {
	out = make(WriteOffProductsMarksSlice, 0)
	sql := `select * from write_off_products_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
