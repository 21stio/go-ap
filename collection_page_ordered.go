package ap

type orderedCollectionPageTraits struct {
	PartOf *OrderedCollectionLink     `json:"partOf,omitempty"`
	Next   *OrderedCollectionPageLink `json:"next,omitempty"`
	Prev   *OrderedCollectionPageLink `json:"prev,omitempty"`
}

type OrderedCollectionPage struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	collectionTrait
	orderedCollectionTrait

	orderedCollectionPageTraits
}

func NewOrderedCollectionPage() (cp OrderedCollectionPage) {
	cp.Type = string(COLLECTIONPAGE_ORDERED)

	return
}
