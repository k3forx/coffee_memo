package coffee_bean

type CreateRequest struct {
	Name        string `json:"name"`
	FarmName    string `json:"farmName"`
	Country     string `json:"country"`
	RoastDegree string `json:"roasted_degree"`
	RoastedAt   string `json:"roasted_at"`
}
