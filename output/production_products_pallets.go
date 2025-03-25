package output

type ProductionProductsPallets struct {
    Id	int64	`db:"id"` 
    IdProductionProducts	int	`db:"id_production_products"` 
    PalletNumber	string	`db:"pallet_number"` 
}

type ProductionProductsPalletsSlice []*ProductionProductsPallets

func (r *ProductionProductsPallets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_products_pallets (id_production_products,pallet_number) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdProductionProducts,r.PalletNumber)
	return err
}

func (r *ProductionProductsPallets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_products_pallets SET id_production_products=?,pallet_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionProducts,r.PalletNumber,r.Id)
	return err
}

func (r *ProductionProductsPallets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_products_pallets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionProductsPallets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_products_pallets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionProductsPallets) All(db *sqlx.DB) (out ProductionProductsPalletsSlice, err error) {
	out = make(ProductionProductsPalletsSlice, 0)
	sql := `select * from production_products_pallets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
