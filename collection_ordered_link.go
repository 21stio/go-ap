package ap

type OrderedCollectionLink struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	linkTrait

	collectionTrait
	orderedCollectionTrait
}

func (ocpl OrderedCollectionLink) IsLink() (bool) {
	if ocpl.Type == string(LINK_LINK) {
		return true
	}

	return false
}

func (ocpl OrderedCollectionLink) GetLink() (l link) {
	l.idTrait = ocpl.idTrait
	l.nameTrait = ocpl.nameTrait
	l.mediaTypeTrait = ocpl.mediaTypeTrait
	l.previewTrait = ocpl.previewTrait

	l.linkTrait = ocpl.linkTrait

	return
}

func (ocpl OrderedCollectionLink) IsOrderedCollection() (bool) {
	return !ocpl.IsLink()
}

func (ocpl OrderedCollectionLink) GetOrderedCollection() (oc OrderedCollection) {
	oc.idTrait = ocpl.idTrait
	oc.nameTrait = ocpl.nameTrait
	oc.mediaTypeTrait = ocpl.mediaTypeTrait
	oc.previewTrait = ocpl.previewTrait

	oc.collectionTrait  = ocpl.collectionTrait
	oc.orderedCollectionTrait  = ocpl.orderedCollectionTrait

	return
}
