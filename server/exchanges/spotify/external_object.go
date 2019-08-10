package spotify

// External most spotify objects have these 5 things that help you find it
type External struct {
	ExternalURLs map[string]string `json:"external_urls"`
	HREF         string            `json:"href"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// ExternalID most full objects also need this field
type ExternalID map[string]string
