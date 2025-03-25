package output

type Requests struct {
    Id	int64	`db:"id"` 
    CreateDate	string	`db:"create_date"` 
    Type	string	`db:"type"` 
    Name	string	`db:"name"` 
    Value	string	`db:"value"` 
    Version	string	`db:"version"` 
    Status	string	`db:"status"` 
    ReplyId	string	`db:"reply_id"` 
    Archive	int	`db:"archive"` 
    XmlQuery	string	`db:"xml_query"` 
    XmlReply	string	`db:"xml_reply"` 
}

type RequestsSlice []*Requests

func (r *Requests) Insert(db *sqlx.DB) (err error) {
	sql := `INSERT INTO requests (create_date,type,name,value,version,status,reply_id,archive,xml_query,xml_reply) 
	VALUES(?,?,?,?,?,?,?,?,?,?);`
	_, err = db.Exec(sql, r.CreateDate,r.Type,r.Name,r.Value,r.Version,r.Status,r.ReplyId,r.Archive,r.XmlQuery,r.XmlReply)
	return err
}

func (r *Requests) Update(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return r.Insert(db)
	}
	sql := `UPDATE requests SET create_date=?,type=?,name=?,value=?,version=?,status=?,reply_id=?,archive=?,xml_query=?,xml_reply=?
	WHERE id=?;`
	_, err = db.Exec(sql, r.CreateDate,r.Type,r.Name,r.Value,r.Version,r.Status,r.ReplyId,r.Archive,r.XmlQuery,r.XmlReply,r.Id)
	return err
}

func (r *Requests) Delete(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("id = 0")
	}
	sql := `DELETE requests WHERE id=?;`
	_, err = db.Exec(sql, r.Id)
	return err
}

func (r *Requests) Find(db *sqlx.DB) (err error) {
	if r.Id == 0 {
		return fmt.Errorf("r.Id = 0")
	}
	sql := `select * from requests WHERE id=?;`
	err = db.Get(r, sql, r.Id)
	return err
}

func (r *Requests) All(db *sqlx.DB) (out RequestsSlice, err error) {
	out = make(RequestsSlice, 0)
	sql := `select * from requests order by id;`
	err = db.Select(&out, sql)
	return out, err
}
