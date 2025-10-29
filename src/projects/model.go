package projects

type Project struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Requirements     []string `json:"requirements"`
	Budget           float64  `json:"budget"`
	Timeline         string   `json:"timeline"`
	Status           string   `json:"status"`
	ClientID         string   `json:"clientId,omitempty"`
	CreatedAt        string   `json:"createdAt,omitempty"`
	ProjectID        string   `json:"id,omitempty"`
	ProjectQuotation int      `json:"project_quotation_price,omitempty"`
	ContactEmail     string   `json:"contact_email,omitempty"`
	RepositoryLink   string   `json:"repository_link,omitempty"`
	Notes            string   `json:"notes,omitempty"`
}
