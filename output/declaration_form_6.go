package output

type DeclarationForm6 struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    Type	string	`db:"type"` 
    ReportingQuarter	string	`db:"reporting_quarter"` 
    ReportingYear	string	`db:"reporting_year"` 
    CorrectionNumber	string	`db:"correction_number"` 
    DirectorSurname	string	`db:"director_surname"` 
    DirectorName	string	`db:"director_name"` 
    DirectorMiddleName	string	`db:"director_middle_name"` 
    AccountantSurname	string	`db:"accountant_surname"` 
    AccountantName	string	`db:"accountant_name"` 
    AccountantMiddleName	string	`db:"accountant_middle_name"` 
    Guid	string	`db:"guid"` 
}

type DeclarationForm6Slice []*DeclarationForm6

func (r *DeclarationForm6) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO declaration_form_6 (create_date,type,reporting_quarter,reporting_year,correction_number,director_surname,director_name,director_middle_name,accountant_surname,accountant_name,accountant_middle_name,guid) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.Type,r.ReportingQuarter,r.ReportingYear,r.CorrectionNumber,r.DirectorSurname,r.DirectorName,r.DirectorMiddleName,r.AccountantSurname,r.AccountantName,r.AccountantMiddleName,r.Guid)
	return err
}

func (r *DeclarationForm6) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE declaration_form_6 SET create_date=?,type=?,reporting_quarter=?,reporting_year=?,correction_number=?,director_surname=?,director_name=?,director_middle_name=?,accountant_surname=?,accountant_name=?,accountant_middle_name=?,guid=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.Type,r.ReportingQuarter,r.ReportingYear,r.CorrectionNumber,r.DirectorSurname,r.DirectorName,r.DirectorMiddleName,r.AccountantSurname,r.AccountantName,r.AccountantMiddleName,r.Guid,r.Id)
	return err
}

func (r *DeclarationForm6) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE declaration_form_6 WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *DeclarationForm6) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from declaration_form_6 WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *DeclarationForm6) All(db *sqlx.DB) (out DeclarationForm6Slice, err error) {
	out = make(DeclarationForm6Slice, 0)
	sql := `select * from declaration_form_6 order by id;`
	err = db.Select(&out, sql)
	return out, err
}
