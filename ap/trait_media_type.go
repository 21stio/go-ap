package ap

type MediaType string

const (
	MEDIA_HTML MediaType = "text/html"
	MEDIA_JSON MediaType = "text/json"
)

type mediaTypeTrait struct {
	MediaType MediaType         `json:"mediaType,omitempty"`
}


