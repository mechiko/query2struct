package output

type RestsApVolume struct {
    Id	int64	`db:"id"` 
    ProductInformF2RegId	string	`db:"product_inform_f2_reg_id"` 
    ProductTotalMlVolume	string	`db:"product_total_ml_volume"` 
    ProductCurrentMlVolume	string	`db:"product_current_ml_volume"` 
    ProductIsOpen	int	`db:"product_is_open"` 
}

type RestsApVolumeSlice []*RestsApVolume

func (r *RestsApVolume) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO rests_ap_volume (product_inform_f2_reg_id,product_total_ml_volume,product_current_ml_volume,product_is_open) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.ProductInformF2RegId,r.ProductTotalMlVolume,r.ProductCurrentMlVolume,r.ProductIsOpen)
	return err
}

func (r *RestsApVolume) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE rests_ap_volume SET product_inform_f2_reg_id=?,product_total_ml_volume=?,product_current_ml_volume=?,product_is_open=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.ProductInformF2RegId,r.ProductTotalMlVolume,r.ProductCurrentMlVolume,r.ProductIsOpen,r.Id)
	return err
}

func (r *RestsApVolume) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE rests_ap_volume WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *RestsApVolume) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from rests_ap_volume WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *RestsApVolume) All(db *sqlx.DB) (out RestsApVolumeSlice, err error) {
	out = make(RestsApVolumeSlice, 0)
	sql := `select * from rests_ap_volume order by id;`
	err = db.Select(&out, sql)
	return out, err
}
