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