package interfaces

type TempFile struct {
	Type *string `json:"Type"`
	Url  string  `json:"Url"`
	Name string  `json:"Name"`
}
