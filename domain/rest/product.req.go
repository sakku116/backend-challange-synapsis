package rest

type GetProductListReq struct {
	Category  string `json:"category"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}
