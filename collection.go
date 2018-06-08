package ap

type unorderedCollectionTrait struct {
	Current *CollectionPageLink `json:"current,omitempty"`
	First   *CollectionPageLink `json:"first,omitempty"`
	Last    *CollectionPageLink `json:"last,omitempty"`
	Items   []*ObjectLink       `json:"items,omitempty"`
}

type Collection struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	collectionTrait
	unorderedCollectionTrait
}

func NewCollection(items []*ObjectLink) (c Collection) {
	c.Items = items

	return
}
