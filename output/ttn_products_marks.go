package output

type TtnProductsMarks struct {
    Id	int64	`db:"id"` 
    IdTtnProducts	int	`db:"id_ttn_products"` 
    IdTtnProductsBoxes	int	`db:"id_ttn_products_boxes"` 
    Mark	string	`db:"mark"` 
    Archive	int	`db:"archive"` 
}

type TtnProductsMarksSlice []*TtnProductsMarks

func (r *TtnProductsMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_products_marks (id_ttn_products,id_ttn_products_boxes,mark,archive) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.IdTtnProducts,r.IdTtnProductsBoxes,r.Mark,r.Archive)
	return err
}

func (r *TtnProductsMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_products_marks SET id_ttn_products=?,id_ttn_products_boxes=?,mark=?,archive=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnProducts,r.IdTtnProductsBoxes,r.Mark,r.Archive,r.Id)
	return err
}

func (r *TtnProductsMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_products_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnProductsMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_products_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnProductsMarks) All(db *sqlx.DB) (out TtnProductsMarksSlice, err error) {
	out = make(TtnProductsMarksSlice, 0)
	sql := `select * from ttn_products_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
