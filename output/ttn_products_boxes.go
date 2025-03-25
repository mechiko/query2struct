package output

type TtnProductsBoxes struct {
    Id	int64	`db:"id"` 
    IdTtnProducts	int	`db:"id_ttn_products"` 
    IdTtnProductsPallets	int	`db:"id_ttn_products_pallets"` 
    BoxNumber	string	`db:"box_number"` 
}

type TtnProductsBoxesSlice []*TtnProductsBoxes

func (r *TtnProductsBoxes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_products_boxes (id_ttn_products,id_ttn_products_pallets,box_number) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdTtnProducts,r.IdTtnProductsPallets,r.BoxNumber)
	return err
}

func (r *TtnProductsBoxes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_products_boxes SET id_ttn_products=?,id_ttn_products_pallets=?,box_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnProducts,r.IdTtnProductsPallets,r.BoxNumber,r.Id)
	return err
}

func (r *TtnProductsBoxes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_products_boxes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnProductsBoxes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_products_boxes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnProductsBoxes) All(db *sqlx.DB) (out TtnProductsBoxesSlice, err error) {
	out = make(TtnProductsBoxesSlice, 0)
	sql := `select * from ttn_products_boxes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
