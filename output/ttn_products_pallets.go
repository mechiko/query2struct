package output

type TtnProductsPallets struct {
    Id	int64	`db:"id"` 
    IdTtnProducts	int	`db:"id_ttn_products"` 
    PalletNumber	string	`db:"pallet_number"` 
}

type TtnProductsPalletsSlice []*TtnProductsPallets

func (r *TtnProductsPallets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO ttn_products_pallets (id_ttn_products,pallet_number) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdTtnProducts,r.PalletNumber)
	return err
}

func (r *TtnProductsPallets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE ttn_products_pallets SET id_ttn_products=?,pallet_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdTtnProducts,r.PalletNumber,r.Id)
	return err
}

func (r *TtnProductsPallets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE ttn_products_pallets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *TtnProductsPallets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from ttn_products_pallets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *TtnProductsPallets) All(db *sqlx.DB) (out TtnProductsPalletsSlice, err error) {
	out = make(TtnProductsPalletsSlice, 0)
	sql := `select * from ttn_products_pallets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
