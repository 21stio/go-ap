package ap

type LinkType string

const (
	LINK_LINK LinkType = "Link"
	LINK_MENTION LinkType = "Mention"
)

type linkTrait struct {
	Href     string   `json:"href,omitempty"`
	Rel      []string `json:"rel,omitempty"`
	Hreflang string   `json:"hreflang,omitempty"`
	Height   int64    `json:"height,omitempty"`
	Width    int64    `json:"width,omitempty"`
}

type link struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait

	linkTrait
}

func (l link) ToObjectLink() (ol ObjectLink) {
	ol = ObjectLink{}
	ol.idTrait = l.idTrait
	ol.nameTrait = l.nameTrait
	ol.mediaTypeTrait = l.mediaTypeTrait
	ol.previewTrait = l.previewTrait

	ol.linkTrait = l.linkTrait

	return
}