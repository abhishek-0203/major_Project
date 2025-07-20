package client

type Client struct {
	ClientID     string   `json:"client_id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Skills       []string `json:"skills"`
	Experience   int      `json:"experience"`
	LinkedIn     string   `json:"linkedin"`
	Available    string   `json:"available"`
	Company      string   `json:"company"`
	ContactEmail string   `json:"contact_email"`
	Address      string   `json:"address"`
	Projects     []string `json:"projects"`
}
