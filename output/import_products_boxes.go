package output

type ImportProductsBoxes struct {
    Id	int64	`db:"id"` 
    IdImportProducts	int	`db:"id_import_products"` 
    IdImportProductsPallets	int	`db:"id_import_products_pallets"` 
    BoxNumber	string	`db:"box_number"` 
}

type ImportProductsBoxesSlice []*ImportProductsBoxes

func (r *ImportProductsBoxes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_products_boxes (id_import_products,id_import_products_pallets,box_number) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdImportProducts,r.IdImportProductsPallets,r.BoxNumber)
	return err
}

func (r *ImportProductsBoxes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_products_boxes SET id_import_products=?,id_import_products_pallets=?,box_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportProducts,r.IdImportProductsPallets,r.BoxNumber,r.Id)
	return err
}

func (r *ImportProductsBoxes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_products_boxes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportProductsBoxes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_products_boxes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportProductsBoxes) All(db *sqlx.DB) (out ImportProductsBoxesSlice, err error) {
	out = make(ImportProductsBoxesSlice, 0)
	sql := `select * from import_products_boxes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
