package output

type DeclarationPersons struct {
    Id	int64	`db:"id"` 
    Type	string	`db:"type"` 
    Surname	string	`db:"surname"` 
    Name	string	`db:"name"` 
    MiddleName	string	`db:"middle_name"` 
}

type DeclarationPersonsSlice []*DeclarationPersons

func (r *DeclarationPersons) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_persons (type,surname,name,middle_name) 
	VALUES(?,?,?,?);`
	_, err = db.Exec(sql, r.Type,r.Surname,r.Name,r.MiddleName)
	return err
}

func (r *DeclarationPersons) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_persons SET type=?,surname=?,name=?,middle_name=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.Type,r.Surname,r.Name,r.MiddleName,r.Id)
	return err
}

func (r *DeclarationPersons) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_persons WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *DeclarationPersons) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from declaration_persons WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *DeclarationPersons) All(db *sqlx.DB) (out DeclarationPersonsSlice, err error) {
	out = make(DeclarationPersonsSlice, 0)
	sql := `select * from declaration_persons order by id;`
	err = db.Select(&out, sql)
	return out, err
}
