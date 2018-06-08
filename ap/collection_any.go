package ap

func init() {
	//log.Print("asdds")
	//a := CollectionPage{}
	//a.Id = "a"
}

type collectionTraits struct {
	TotalItems int64              `json:"totalItems,omitempty"`
	Current    collectionPageLink `json:"current,omitempty"`
	First      collectionPageLink `json:"first,omitempty"`
	Last       collectionPageLink `json:"last,omitempty"`
	Items      []*ObjectLink      `json:"items,omitempty"`
}

type orderedCollectionTraits struct {
	OrderedItems []*ObjectLink `json:"orderedItems,omitempty"`
}

type orderedCollection struct {
	baseObject
	collectionTraits
	orderedCollectionTraits
}

func NewOrderedCollection(orderedItems []*ObjectLink) (c orderedCollection) {
	c.OrderedItems = orderedItems

	return
}

func NewCollection(items []*ObjectLink) (c orderedCollection) {
	c.Items = items

	return
}

type collectionPageTraits struct {
	PartOf *orderedCollection
	Next   *CollectionPage
	Prev   *CollectionPage
}

type CollectionPage struct {
	*orderedCollection
	collectionPageTraits
}

type CollectionLink struct {
	link
	collectionTraits
}

type collectionPageLink struct {
	*CollectionPage
	*link
}
