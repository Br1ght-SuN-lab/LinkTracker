package client


type LinkResponseDTO struct {
	ID int64         `yaml:"id"`
	URL string       `yaml:"url"`
	Tags []string    `yaml:"tags"`
	Filters []string `yaml:"filters"`
}

type ApiErrorResponseDTO struct {
	Description string      `yaml:"description"`
	Code string             `yaml:"code"`
	ExceptionName string    `yaml:"exceptionName"`
	ExceptionMessage string `yaml:"exceptionMessage"`
	Stacktrace []string     `yaml:"stacktrace"`
}

type AddLinkRequestDTO struct {
	Link string      `yaml:"link"`
	Tags []string    `yaml:"tags"`
	Filters []string `yaml:"filters"`
}

type ListLinksResponseDTO struct {
	Links []LinkResponseDTO `yaml:"links"`
	Size int32     `yaml:"size"`
}

type RemoveLinkRequestDTO struct {
	Link string `yaml:"link"`
}