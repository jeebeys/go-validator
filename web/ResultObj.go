package web

import "math"

var (
	SUCCESS = ResultObj{}.Data("success", true)
	FAILURE = ResultObj{}.Data("success", false)
)

type ResultObj map[string]interface {
}

func (r ResultObj) Message(message string) ResultObj {
	r["message"] = message
	return r
}

func (r ResultObj) Data(key string, val interface{}) ResultObj {
	r[key] = val
	return r
}

func (r ResultObj) Pagination(items []interface{}, total, page, size int) ResultObj {
	r["items"] = items
	r["total"] = total
	r["pages"] = int64(math.Ceil(float64(total) / float64(size) * 1.0))
	r["size"] = size
	r["page"] = page
	return r
}

func (r ResultObj) GetMessage() string {
	return r["message"].(string)
}

func (r ResultObj) GetData(key string) interface{} {
	return r[key]
}

func (r ResultObj) Ptr() *ResultObj {
	return &r
}
