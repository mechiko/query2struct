package output

type ImportProductsPallets struct {
    Id	int64	`db:"id"` 
    IdImportProducts	int	`db:"id_import_products"` 
    PalletNumber	string	`db:"pallet_number"` 
}

type ImportProductsPalletsSlice []*ImportProductsPallets

func (r *ImportProductsPallets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO import_products_pallets (id_import_products,pallet_number) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdImportProducts,r.PalletNumber)
	return err
}

func (r *ImportProductsPallets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE import_products_pallets SET id_import_products=?,pallet_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdImportProducts,r.PalletNumber,r.Id)
	return err
}

func (r *ImportProductsPallets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE import_products_pallets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ImportProductsPallets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from import_products_pallets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ImportProductsPallets) All(db *sqlx.DB) (out ImportProductsPalletsSlice, err error) {
	out = make(ImportProductsPalletsSlice, 0)
	sql := `select * from import_products_pallets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
