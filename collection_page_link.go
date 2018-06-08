package ap

type CollectionPageLink struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	linkTrait

	collectionTrait
	unorderedCollectionTrait

	unorderedCollectionPageTraits
}

func (cpl CollectionPageLink) IsLink() (bool) {
	if cpl.Type == string(LINK_LINK) {
		return true
	}

	return false
}

func (cpl CollectionPageLink) GetLink() (l link) {
	l.idTrait = cpl.idTrait
	l.nameTrait = cpl.nameTrait
	l.mediaTypeTrait = cpl.mediaTypeTrait
	l.previewTrait = cpl.previewTrait

	l.linkTrait = cpl.linkTrait

	return
}

func (cpl CollectionPageLink) IsCollectionPage() (bool) {
	return !cpl.IsLink()
}

func (cpl CollectionPageLink) GetCollectionPage() (cp CollectionPage) {
	cp.idTrait = cpl.idTrait
	cp.nameTrait = cpl.nameTrait
	cp.mediaTypeTrait = cpl.mediaTypeTrait
	cp.previewTrait = cpl.previewTrait

	cp.collectionTrait  = cpl.collectionTrait
	cp.unorderedCollectionTrait  = cpl.unorderedCollectionTrait

	cp.unorderedCollectionPageTraits  = cpl.unorderedCollectionPageTraits

	return
}