package user_coffee_bean

type CreateRequest struct {
	Name        string `json:"name"`
	FarmName    string `json:"farmName"`
	Country     string `json:"country"`
	RoastDegree string `json:"roastDegree"`
	RoastedAt   string `json:"roastedAt"`
}

type EditByIDRequest struct {
	Name        string `json:"name"`
	FarmName    string `json:"farmName"`
	Country     string `json:"country"`
	RoastDegree string `json:"roastDegree"`
	RoastedAt   string `json:"roastedAt"`
}
