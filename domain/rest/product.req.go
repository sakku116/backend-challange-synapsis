package rest

type GetProductListReq struct {
	Category  string `form:"category" default:""`
	Search    string `form:"search" default:""` // using regex
	Page      int    `form:"page" default:"1"`
	Limit     int    `form:"limit" default:"10"`
	SortBy    string `form:"sort_by" default:"created_at"`
	SortOrder string `form:"sort_order" default:"desc"`
}
