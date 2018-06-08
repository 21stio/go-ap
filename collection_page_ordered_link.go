package ap

type OrderedCollectionPageLink struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	linkTrait

	collectionTrait
	orderedCollectionTrait

	orderedCollectionPageTraits
}

func (ocpl OrderedCollectionPageLink) IsLink() (bool) {
	if ocpl.Type == string(LINK_LINK) {
		return true
	}

	return false
}

func (ocpl OrderedCollectionPageLink) GetLink() (l link) {
	l.idTrait = ocpl.idTrait
	l.nameTrait = ocpl.nameTrait
	l.mediaTypeTrait = ocpl.mediaTypeTrait
	l.previewTrait = ocpl.previewTrait

	l.linkTrait = ocpl.linkTrait

	return
}

func (ocpl OrderedCollectionPageLink) IsOrderedCollectionPage() (bool) {
	return !ocpl.IsLink()
}

func (ocpl OrderedCollectionPageLink) GetOrderedCollectionPage() (ocp OrderedCollectionPage) {
	ocp.idTrait = ocpl.idTrait
	ocp.nameTrait = ocpl.nameTrait
	ocp.mediaTypeTrait = ocpl.mediaTypeTrait
	ocp.previewTrait = ocpl.previewTrait

	ocp.collectionTrait  = ocpl.collectionTrait
	ocp.orderedCollectionTrait  = ocpl.orderedCollectionTrait

	ocp.orderedCollectionPageTraits  = ocpl.orderedCollectionPageTraits

	return
}