package interfaces

const (
	Video  = "Video"
	Folder = "Folder"
)

type Meida struct {
	OnlineMedia bool
	MediaType   string
	Url         string
	Name        string
}
