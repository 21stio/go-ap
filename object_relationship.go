package ap

type relationshipTrait struct {
	Subject      *ObjectLink `json:"subject,omitempty"`
	Object       *ObjectLink `json:"baseObject,omitempty"`
	Relationship *ObjectLink `json:"relationship,omitempty"`
}

type relationship struct {
	*baseObject
	relationshipTrait
}

func NewRelationship() (r relationship) {
	r.Type = string(OBJECT_RELATIONSHIP)

	return
}

func (r relationship) ToObjectLinkBag() (ol *ObjectLink) {
	ol = r.baseObject.ToObjectLink()
	ol.relationshipTrait = r.relationshipTrait

	return
}

func (r relationship) ToAnyObject() (ao *anyObject) {
	ao = r.baseObject.ToAnyObject()
	ao.relationshipTrait = r.relationshipTrait

	return
}