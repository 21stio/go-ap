package ap

type unorderedCollectionPageTraits struct {
	PartOf *Collection
	Next   *CollectionPage
	Prev   *CollectionPage
}

type CollectionPage struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	collectionTrait
	unorderedCollectionTrait

	unorderedCollectionPageTraits
}

func NewCollectionPage() (cp CollectionPage) {
	cp.Type = string(COLLECTIONPAGE_COLLECTIONPAGE)

	return
}