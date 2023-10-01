package response

type Member struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Identity int    `json:"identity_number"`
}
