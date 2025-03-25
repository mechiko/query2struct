package output

type WriteOffProductsBoxes struct {
    Id	int64	`db:"id"` 
    IdWriteOffProducts	int	`db:"id_write_off_products"` 
    IdWriteOffProductsPallets	int	`db:"id_write_off_products_pallets"` 
    BoxNumber	string	`db:"box_number"` 
}

type WriteOffProductsBoxesSlice []*WriteOffProductsBoxes

func (r *WriteOffProductsBoxes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO write_off_products_boxes (id_write_off_products,id_write_off_products_pallets,box_number) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.IdWriteOffProductsPallets,r.BoxNumber)
	return err
}

func (r *WriteOffProductsBoxes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE write_off_products_boxes SET id_write_off_products=?,id_write_off_products_pallets=?,box_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.IdWriteOffProductsPallets,r.BoxNumber,r.Id)
	return err
}

func (r *WriteOffProductsBoxes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE write_off_products_boxes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *WriteOffProductsBoxes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from write_off_products_boxes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *WriteOffProductsBoxes) All(db *sqlx.DB) (out WriteOffProductsBoxesSlice, err error) {
	out = make(WriteOffProductsBoxesSlice, 0)
	sql := `select * from write_off_products_boxes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
