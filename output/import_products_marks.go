package output

type ImportProductsMarks struct {
    Id	int64	`db:"id"` 
    IdImportProducts	int	`db:"id_import_products"` 
    IdImportProductsBoxes	int	`db:"id_import_products_boxes"` 
    Mark	string	`db:"mark"` 
}

type ImportProductsMarksSlice []*ImportProductsMarks

func (r *ImportProductsMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_products_marks (id_import_products,id_import_products_boxes,mark) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdImportProducts,r.IdImportProductsBoxes,r.Mark)
	return err
}

func (r *ImportProductsMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_products_marks SET id_import_products=?,id_import_products_boxes=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportProducts,r.IdImportProductsBoxes,r.Mark,r.Id)
	return err
}

func (r *ImportProductsMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_products_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportProductsMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_products_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportProductsMarks) All(db *sqlx.DB) (out ImportProductsMarksSlice, err error) {
	out = make(ImportProductsMarksSlice, 0)
	sql := `select * from import_products_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
