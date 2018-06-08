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

func NewLink(href string) (l link) {
	l.Type = string(LINK_LINK)
	l.Href = href

	return
}

func (l link) ToObjectLink() (ol *ObjectLink) {
	ol = &ObjectLink{}

	ol.idTrait = l.idTrait
	ol.nameTrait = l.nameTrait
	ol.mediaTypeTrait = l.mediaTypeTrait
	ol.previewTrait = l.previewTrait

	ol.linkTrait = l.linkTrait

	return
}

func (l link) ToCollectionPageLink() (cpl *CollectionPageLink) {
	cpl = &CollectionPageLink{}

	cpl.idTrait = l.idTrait
	cpl.nameTrait = l.nameTrait
	cpl.mediaTypeTrait = l.mediaTypeTrait
	cpl.previewTrait = l.previewTrait

	cpl.linkTrait = l.linkTrait

	return
}

func (l link) ToOrderedCollectionPageLink() (ocpl *OrderedCollectionPageLink) {
	ocpl = &OrderedCollectionPageLink{}

	ocpl.idTrait = l.idTrait
	ocpl.nameTrait = l.nameTrait
	ocpl.mediaTypeTrait = l.mediaTypeTrait
	ocpl.previewTrait = l.previewTrait

	ocpl.linkTrait = l.linkTrait

	return
}

func (l link) ToOrderedCollectionLink() (ocl *OrderedCollectionLink) {
	ocl = &OrderedCollectionLink{}

	ocl.idTrait = l.idTrait
	ocl.nameTrait = l.nameTrait
	ocl.mediaTypeTrait = l.mediaTypeTrait
	ocl.previewTrait = l.previewTrait

	ocl.linkTrait = l.linkTrait

	return
}

func (l link) ToCollectionLink() (cl *CollectionLink) {
	cl = &CollectionLink{}

	cl.idTrait = l.idTrait
	cl.nameTrait = l.nameTrait
	cl.mediaTypeTrait = l.mediaTypeTrait
	cl.previewTrait = l.previewTrait

	cl.linkTrait = l.linkTrait

	return
}