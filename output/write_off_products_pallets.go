package output

type WriteOffProductsPallets struct {
    Id	int64	`db:"id"` 
    IdWriteOffProducts	int	`db:"id_write_off_products"` 
    PalletNumber	string	`db:"pallet_number"` 
}

type WriteOffProductsPalletsSlice []*WriteOffProductsPallets

func (r *WriteOffProductsPallets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO write_off_products_pallets (id_write_off_products,pallet_number) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.PalletNumber)
	return err
}

func (r *WriteOffProductsPallets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE write_off_products_pallets SET id_write_off_products=?,pallet_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdWriteOffProducts,r.PalletNumber,r.Id)
	return err
}

func (r *WriteOffProductsPallets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE write_off_products_pallets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *WriteOffProductsPallets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from write_off_products_pallets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *WriteOffProductsPallets) All(db *sqlx.DB) (out WriteOffProductsPalletsSlice, err error) {
	out = make(WriteOffProductsPalletsSlice, 0)
	sql := `select * from write_off_products_pallets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
