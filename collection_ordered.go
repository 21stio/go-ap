package ap

type orderedCollectionTrait struct {
	Current    *OrderedCollectionPageLink `json:"current,omitempty"`
	First      *OrderedCollectionPageLink `json:"first,omitempty"`
	Last       *OrderedCollectionPageLink `json:"last,omitempty"`
	OrderedItems []*ObjectLink `json:"orderedItems,omitempty"`
}

type OrderedCollection struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	collectionTrait
	orderedCollectionTrait
}

func NewOrderedCollection() (oc OrderedCollection) {
	oc.Type = string(COLLECTION_ORDERED)

	return
}

func (oc OrderedCollection) ToOrderedCollectionLink() (ocl *OrderedCollectionLink) {
	ocl = &OrderedCollectionLink{}

	ocl.idTrait = oc.idTrait
	ocl.nameTrait = oc.nameTrait
	ocl.mediaTypeTrait = oc.mediaTypeTrait
	ocl.previewTrait = oc.previewTrait
	ocl.baseTrait = oc.baseTrait

	ocl.collectionTrait = oc.collectionTrait
	ocl.orderedCollectionTrait = oc.orderedCollectionTrait

	return
}