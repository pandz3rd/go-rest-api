package request

type InsertCategoryReq struct {
	Name string `json:"name" validate:"required"`
}
