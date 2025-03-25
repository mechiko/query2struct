package output

type ProductionProductsMarks struct {
    Id	int64	`db:"id"` 
    IdProductionProducts	int	`db:"id_production_products"` 
    IdProductionProductsBoxes	int	`db:"id_production_products_boxes"` 
    Mark	string	`db:"mark"` 
}

type ProductionProductsMarksSlice []*ProductionProductsMarks

func (r *ProductionProductsMarks) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_products_marks (id_production_products,id_production_products_boxes,mark) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdProductionProducts,r.IdProductionProductsBoxes,r.Mark)
	return err
}

func (r *ProductionProductsMarks) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_products_marks SET id_production_products=?,id_production_products_boxes=?,mark=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionProducts,r.IdProductionProductsBoxes,r.Mark,r.Id)
	return err
}

func (r *ProductionProductsMarks) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_products_marks WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionProductsMarks) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_products_marks WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionProductsMarks) All(db *sqlx.DB) (out ProductionProductsMarksSlice, err error) {
	out = make(ProductionProductsMarksSlice, 0)
	sql := `select * from production_products_marks order by id;`
	err = db.Select(&out, sql)
	return out, err
}
