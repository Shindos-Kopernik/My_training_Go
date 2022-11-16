package apperror

type AppError struct {
	Err              error  `json:"-"` // игнорируем исходную ошибку
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	Code             string `json:"code,omitempty"`
}

func ()  {
	
}