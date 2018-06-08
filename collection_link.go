package ap

type CollectionLink struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait

	baseTrait
	linkTrait

	collectionTrait
	orderedCollectionTrait
	unorderedCollectionTrait
}

func (cl CollectionLink) IsLink() (bool) {
	if cl.Type == string(LINK_LINK) {
		return true
	}

	return false
}

func (cl CollectionLink) GetLink() (l link) {
	l.idTrait = cl.idTrait
	l.nameTrait = cl.nameTrait
	l.mediaTypeTrait = cl.mediaTypeTrait
	l.previewTrait = cl.previewTrait

	l.linkTrait = cl.linkTrait

	return
}

func (cl CollectionLink) IsAnyCollection() (bool) {
	return !cl.IsLink()
}

func (cl CollectionLink) GetAnyCollection() (ac anyCollection) {
	ac.idTrait = cl.idTrait
	ac.nameTrait = cl.nameTrait
	ac.mediaTypeTrait = cl.mediaTypeTrait
	ac.previewTrait = cl.previewTrait

	ac.baseTrait = cl.baseTrait
	ac.collectionTrait = cl.collectionTrait
	ac.orderedCollectionTrait = cl.orderedCollectionTrait
	ac.unorderedCollectionTrait = cl.unorderedCollectionTrait

	return
}
