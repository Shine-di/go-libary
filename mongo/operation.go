package mongo

type Mdb string

type MCollection string

func (e Mdb) parse() string {
	return string(e)
}

func (e MCollection) parse() string {
	return string(e)
}

func Insert(db Mdb, c MCollection, docs ...interface{}) error {
	s := session.Clone()
	defer s.Close()

	return s.DB(db.parse()).C(c.parse()).Insert(docs...)
}

func FindOne(db Mdb, c MCollection, query interface{}, result interface{}) error {
	s := session.Clone()
	defer s.Close()
	if err := s.DB(db.parse()).C(c.parse()).Find(query).One(result); err != nil {
		return err
	}
	return nil
}

func FindList(db Mdb, c MCollection, query interface{}, selector interface{}, sort string, result interface{}, offset, limit int) error {
	s := session.Clone()
	defer s.Close()

	return s.DB(db.parse()).C(c.parse()).Find(query).Sort(sort).Select(selector).Skip(offset).Limit(limit).All(result)
}

func Update(db Mdb, c MCollection, selector interface{}, update interface{}) error {
	s := session.Clone()
	defer s.Close()

	return s.DB(db.parse()).C(c.parse()).Update(selector, update)
}
