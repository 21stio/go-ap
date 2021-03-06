package ap

type ObjectType string

const (
	OBJECT_NIL          ObjectType = ""
	OBJECT_ARTICLE      ObjectType = "Article"
	OBJECT_AUDIO        ObjectType = "Audio"
	OBJECT_DOCUMENT     ObjectType = "Document"
	OBJECT_EVENT        ObjectType = "Event"
	OBJECT_IMAGE        ObjectType = "Image"
	OBJECT_NOTE         ObjectType = "Note"
	OBJECT_PAGE         ObjectType = "Page"
	OBJECT_PLACE        ObjectType = "Place"
	OBJECT_PROFILE      ObjectType = "Profile"
	OBJECT_RELATIONSHIP ObjectType = "Relationship"
	OBJECT_TOMBSTONE    ObjectType = "Tombstone"
	OBJECT_VIDEO        ObjectType = "Video"
)

type anyObject struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait

	baseTrait
	placeTrait
	relationshipTrait
	tombstoneTrait
}

func (any anyObject) IsPlace() (bool) {
	return any.Type == string(OBJECT_PLACE)
}

func (any anyObject) IsRelationship() (bool) {
	return any.Type == string(OBJECT_RELATIONSHIP)
}

func (any anyObject) IsTombstone() (bool) {
	return any.Type == string(OBJECT_TOMBSTONE)
}

func (any anyObject) IsBaseObject() (bool) {
	return !any.IsPlace() && any.IsRelationship() && any.IsTombstone()
}

func (any anyObject) GetBaseObject() (bo baseObject) {
	bo.idTrait = any.idTrait
	bo.nameTrait = any.nameTrait
	bo.mediaTypeTrait = any.mediaTypeTrait
	bo.previewTrait = any.previewTrait

	bo.baseTrait = any.baseTrait

	return
}

func (any anyObject) GetPlace() (p place) {
	p.idTrait = any.idTrait
	p.nameTrait = any.nameTrait
	p.mediaTypeTrait = any.mediaTypeTrait
	p.previewTrait = any.previewTrait

	p.baseTrait = any.baseTrait

	return
}

func (any anyObject) GetTombstone() (t tombstone) {
	t.idTrait = any.idTrait
	t.nameTrait = any.nameTrait
	t.mediaTypeTrait = any.mediaTypeTrait
	t.previewTrait = any.previewTrait

	t.baseTrait = any.baseTrait
	t.tombstoneTrait = any.tombstoneTrait

	return
}

func (any anyObject) GetRelationship() (r relationship) {
	r.idTrait = any.idTrait
	r.nameTrait = any.nameTrait
	r.mediaTypeTrait = any.mediaTypeTrait
	r.previewTrait = any.previewTrait

	r.baseTrait = any.baseTrait
	r.relationshipTrait = any.relationshipTrait

	return
}
