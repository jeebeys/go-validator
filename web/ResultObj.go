package web

type ResultObj struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Page    int64         `json:"page"`
	Size    int64         `json:"size"`
	Items   []interface{} `json:"items"`
	Pages   int64         `json:"pages"`
	Total   int64         `json:"total"`
}
