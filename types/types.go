package types

type Agent struct {
	ID    int    `storm:"id,increment"`
	Token string `json:"token"`
	Name  string `json:"name"`
}
