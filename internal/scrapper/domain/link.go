package domain 


type Link struct {
	ID int64
	URL string 
	Tags []string
	Filters []string
}


type TrackedLink struct {
	ChatID  int64
	ID      int64
	URL     string
	Tags    []string
	Filters []string
}