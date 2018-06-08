package ap

import (
	"time"
)

type tombstoneTrait struct {
	FormerType *baseObject `json:"formerType,omitempty"`
	Deleted    *time.Time  `json:"deleted,omitempty"`
}

type tombstone struct {
	*baseObject
	tombstoneTrait
}

func NewTombstone() (t tombstone) {
	t.Type = string(OBJECT_TOMBSTONE)

	return
}

func (t tombstone) ToObjectLink() (ol *ObjectLink) {
	ol = t.baseObject.ToObjectLink()
	ol.tombstoneTrait = t.tombstoneTrait

	return
}

func (t tombstone) ToAnyObject() (ao *anyObject) {
	ao = t.baseObject.ToAnyObject()
	ao.tombstoneTrait = t.tombstoneTrait

	return
}