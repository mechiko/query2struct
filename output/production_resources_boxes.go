package output

type ProductionResourcesBoxes struct {
    Id	int64	`db:"id"` 
    IdProductionResources	int	`db:"id_production_resources"` 
    IdProductionResourcesPallets	int	`db:"id_production_resources_pallets"` 
    BoxNumber	string	`db:"box_number"` 
}

type ProductionResourcesBoxesSlice []*ProductionResourcesBoxes

func (r *ProductionResourcesBoxes) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO production_resources_boxes (id_production_resources,id_production_resources_pallets,box_number) 
	VALUES(?,?,?);`
	_, err = db.Exec(sql, r.IdProductionResources,r.IdProductionResourcesPallets,r.BoxNumber)
	return err
}

func (r *ProductionResourcesBoxes) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE production_resources_boxes SET id_production_resources=?,id_production_resources_pallets=?,box_number=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.IdProductionResources,r.IdProductionResourcesPallets,r.BoxNumber,r.Id)
	return err
}

func (r *ProductionResourcesBoxes) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE production_resources_boxes WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *ProductionResourcesBoxes) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from production_resources_boxes WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *ProductionResourcesBoxes) All(db *sqlx.DB) (out ProductionResourcesBoxesSlice, err error) {
	out = make(ProductionResourcesBoxesSlice, 0)
	sql := `select * from production_resources_boxes order by id;`
	err = db.Select(&out, sql)
	return out, err
}
