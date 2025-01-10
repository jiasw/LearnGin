package dtos

type APIPageResponse struct {
	ApiResponse `json:"api_response"`
	Page        int `json:"page"`
	Size        int `json:"size"`
	Total       int `json:"total"`
}

func SuccessPageResponse(data interface{}, page, size, total int) *APIPageResponse {
	return &APIPageResponse{
		ApiResponse: ApiResponse{
			Code:    0,
			Message: "success",
			Data:    data,
		},
		Page:  page,
		Size:  size,
		Total: total,
	}
}
