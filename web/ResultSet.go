package web

type ResultSet []interface{}

func (r ResultSet) Append(items ...interface{}) ResultSet {
	return append(r, items...)
}
