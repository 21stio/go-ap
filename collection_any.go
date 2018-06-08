package ap

type CollectionType string

const (
	COLLECTION_NIL        ObjectType = ""
	COLLECTION_ORDERED    ObjectType = "OrderedCollection"
	COLLECTION_COLLECTION ObjectType = "Collection"
)

type anyCollection struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait
	baseTrait

	collectionTrait
	orderedCollectionTrait
	unorderedCollectionTrait
}

func (any anyCollection) IsOrderedCollection() (bool) {
	return any.Type == string(COLLECTION_ORDERED)
}

func (any anyCollection) IsCollection() (bool) {
	return any.Type == string(COLLECTION_COLLECTION)
}

func (any anyCollection) GetOrderedCollection() (oc OrderedCollection) {
	oc.idTrait = any.idTrait
	oc.nameTrait = any.nameTrait
	oc.mediaTypeTrait = any.mediaTypeTrait
	oc.previewTrait = any.previewTrait
	oc.baseTrait = any.baseTrait

	oc.collectionTrait = any.collectionTrait
	oc.orderedCollectionTrait = any.orderedCollectionTrait

	return
}

func (any anyCollection) GetCollection() (c Collection) {
	c.idTrait = any.idTrait
	c.nameTrait = any.nameTrait
	c.mediaTypeTrait = any.mediaTypeTrait
	c.previewTrait = any.previewTrait
	c.baseTrait = any.baseTrait

	c.collectionTrait = any.collectionTrait
	c.unorderedCollectionTrait = any.unorderedCollectionTrait

	return
}


