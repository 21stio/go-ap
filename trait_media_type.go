package ap

type MediaType string

const (
	MEDIA_TEXT_HTML MediaType = "text/html"
	MEDIA_TEXT_JSON MediaType = "text/json"
	MEDIA_IMAGE_PNG MediaType = "image/png"
	MEDIA_IMAGE_JPG MediaType = "image/jpg"
)

type mediaTypeTrait struct {
	MediaType MediaType         `json:"mediaType,omitempty"`
}


