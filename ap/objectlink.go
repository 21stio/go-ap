package ap

type IToObjectLink interface {
	ToObjectLink() (*ObjectLink)
}

func ToObjectLinks(tobls ...IToObjectLink) (obls []*ObjectLink) {
	for _, tobl := range tobls {
		obls = append(obls, tobl.ToObjectLink())
	}

	return
}

type ObjectLink struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait

	linkTrait

	baseTrait
	placeTrait
	relationshipTrait
	tombstoneTrait
}

func (ol ObjectLink) IsLink() (bool) {
	if ol.Type == string(LINK_LINK) || ol.Type == string(OBJECT_NIL) {
		return true
	}

	return false
}

func (ol ObjectLink) GetLink() (l link) {
	l.idTrait = ol.idTrait
	l.nameTrait = ol.nameTrait
	l.mediaTypeTrait = ol.mediaTypeTrait
	l.previewTrait = ol.previewTrait
	l.linkTrait = ol.linkTrait

	return
}

func (ol ObjectLink) IsAnyObject() (bool) {
	return !ol.IsLink()
}

func (ol ObjectLink) GetAnyObject() (ob anyObject) {
	ob.idTrait = ol.idTrait
	ob.nameTrait = ol.nameTrait
	ob.mediaTypeTrait = ol.mediaTypeTrait
	ob.previewTrait = ol.previewTrait

	ob.baseTrait = ol.baseTrait
	ob.relationshipTrait = ol.relationshipTrait
	ob.placeTrait = ol.placeTrait
	ob.tombstoneTrait = ol.tombstoneTrait

	return
}

func (ol ObjectLink) IsBaseObject() (bool) {
	return ol.GetAnyObject().IsBaseObject()
}

func (ol ObjectLink) GetBaseObject() (baseObject) {
	return ol.GetAnyObject().GetBaseObject()
}

func (ol ObjectLink) IsPlace() (bool) {
	return ol.GetAnyObject().IsPlace()
}

func (ol ObjectLink) GetPlace() (place) {
	return ol.GetAnyObject().GetPlace()
}

func (ol ObjectLink) IsRelationship() (bool) {
	return ol.GetAnyObject().IsRelationship()
}

func (ol ObjectLink) GetRelationship() (relationship) {
	return ol.GetAnyObject().GetRelationship()
}

func (ol ObjectLink) IsTombstone() (bool) {
	return ol.GetAnyObject().IsTombstone()
}

func (ol ObjectLink) GetTombstone() (tombstone) {
	return ol.GetAnyObject().GetTombstone()
}