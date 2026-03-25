package client

type LinkResponseDTO struct {
	ID      int64    `json:"id"`
	URL     string   `json:"url"`
	Tags    []string `json:"tags"`
	Filters []string `json:"filters"`
}

type ApiErrorResponseDTO struct {
	Description      string   `json:"description"`
	Code             string   `json:"code"`
	ExceptionName    string   `json:"exceptionName"`
	ExceptionMessage string   `json:"exceptionMessage"`
	Stacktrace       []string `json:"stacktrace"`
}

type AddLinkRequestDTO struct {
	Link    string   `json:"link"`
	Tags    []string `json:"tags"`
	Filters []string `json:"filters"`
}

type ListLinksResponseDTO struct {
	Links []LinkResponseDTO `json:"links"`
	Size  int32             `json:"size"`
}

type RemoveLinkRequestDTO struct {
	Link string `json:"link"`
}