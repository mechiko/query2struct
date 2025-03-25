package output

type ProductionResourcesPallets struct {
    Id	int64	`db:"id"` 
    IdProductionResources	int	`db:"id_production_resources"` 
    PalletNumber	string	`db:"pallet_number"` 
}

type ProductionResourcesPalletsSlice []*ProductionResourcesPallets

func (r *ProductionResourcesPallets) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_resources_pallets (id_production_resources,pallet_number) 
	VALUES(?,?);`
	_, err = db.Exec(sql, r.IdProductionResources,r.PalletNumber)
	return err
}

func (r *ProductionResourcesPallets) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_resources_pallets SET id_production_resources=?,pallet_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionResources,r.PalletNumber,r.Id)
	return err
}

func (r *ProductionResourcesPallets) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_resources_pallets WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionResourcesPallets) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_resources_pallets WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionResourcesPallets) All(db *sqlx.DB) (out ProductionResourcesPalletsSlice, err error) {
	out = make(ProductionResourcesPalletsSlice, 0)
	sql := `select * from production_resources_pallets order by id;`
	err = db.Select(&out, sql)
	return out, err
}
