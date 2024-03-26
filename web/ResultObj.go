package web

type ResultObj struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Page    int64         `json:"page,omitempty"`
	Size    int64         `json:"size,omitempty"`
	Items   []interface{} `json:"items,omitempty"`
	Pages   int64         `json:"pages,omitempty"`
	Total   int64         `json:"total,omitempty"`
}
