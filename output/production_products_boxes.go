package output

type ProductionProductsBoxes struct {
    Id	int64	`db:"id"` 
    IdProductionProducts	int	`db:"id_production_products"` 
    IdProductionProductsPallets	int	`db:"id_production_products_pallets"` 
    BoxNumber	string	`db:"box_number"` 
}

type ProductionProductsBoxesSlice []*ProductionProductsBoxes

func (r *ProductionProductsBoxes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_products_boxes (id_production_products,id_production_products_pallets,box_number) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdProductionProducts,r.IdProductionProductsPallets,r.BoxNumber)
	return err
}

func (r *ProductionProductsBoxes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_products_boxes SET id_production_products=?,id_production_products_pallets=?,box_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionProducts,r.IdProductionProductsPallets,r.BoxNumber,r.Id)
	return err
}

func (r *ProductionProductsBoxes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_products_boxes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionProductsBoxes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_products_boxes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionProductsBoxes) All(db *sqlx.DB) (out ProductionProductsBoxesSlice, err error) {
	out = make(ProductionProductsBoxesSlice, 0)
	sql := `select * from production_products_boxes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
