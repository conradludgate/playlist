package spotify

type External struct {
	ExternalURLs map[string]string `json:"external_urls"`
	HREF         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type ExternalID map[string]string
