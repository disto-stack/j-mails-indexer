package types

type ApiResponse struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`

	// Data added with builder method
	Data ZincsearchApiResponse `json:"data"`
}

func (a *ApiResponse) WithData(data ZincsearchApiResponse) {
	a.Data = data
}
