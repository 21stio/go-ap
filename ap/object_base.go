package ap

import (
	"log"
	"fmt"
	"time"
)

var hasTraitType = map[ObjectType]bool{OBJECT_PLACE: true, OBJECT_RELATIONSHIP: true, OBJECT_TOMBSTONE: true}

type baseTrait struct {
	Attachment   []*ObjectLink `json:"attachment,omitempty"`
	AttributedTo *ObjectLink   `json:"attributedTo,omitempty"`
	Audience     *ObjectLink   `json:"audience,omitempty"`
	Content      string        `json:"content,omitempty"`
	Context      *ObjectLink   `json:"context,omitempty"`

	EndTime   *time.Time    `json:"endTime,omitempty"`
	Generator *ObjectLink   `json:"generator,omitempty"`
	Icon      string        `json:"icon,omitempty"`
	Image     *baseObject   `json:"image,omitempty"`
	InReplyTo *ObjectLink   `json:"inReplyTo,omitempty"`
	Location  *PlaceLinkBag `json:"location,omitempty"`

	Published  *time.Time         `json:"published,omitempty"`
	Replies    *orderedCollection `json:"replies,omitempty"`
	StartTime  *time.Time         `json:"startTime,omitempty"`
	Summary    string             `json:"summary,omitempty"`
	SummaryMap map[string]string  `json:"summaryMap,omitempty"`
	Tag        *orderedCollection `json:"tag,omitempty"`
	Updated    *time.Time         `json:"updated,omitempty"`
	Url        *link              `json:"url,omitempty"`
	To         []*ObjectLink      `json:"to,omitempty"`
	Bto        []*ObjectLink      `json:"bto,omitempty"`
	Cc         []*ObjectLink      `json:"cc,omitempty"`
	Bcc        []*ObjectLink      `json:"bcc,omitempty"`
	Duration   string             `json:"duration,omitempty"`
}

type baseObject struct {
	idTrait
	nameTrait
	mediaTypeTrait
	previewTrait

	baseTrait
}

func NewBaseObject(t ObjectType) (o baseObject) {
	_, ok := hasTraitType[t]
	if ok {
		log.Print(fmt.Sprintf("WARN: you used NewBaseObject() to create an baseObject of the extended type: %v, it's recommended to use New%v()", t, t))
	}

	o.Type = string(t)

	return
}

func (bo baseObject) ToObjectLink() (ol *ObjectLink) {
	ol = &ObjectLink{}
	ol.idTrait = bo.idTrait
	ol.nameTrait = bo.nameTrait
	ol.mediaTypeTrait = bo.mediaTypeTrait
	ol.previewTrait = bo.previewTrait

	ol.baseTrait = bo.baseTrait

	return
}

func (bo baseObject) ToAnyObject() (any *anyObject) {
	any = &anyObject{}
	any.idTrait = bo.idTrait
	any.nameTrait = bo.nameTrait
	any.mediaTypeTrait = bo.mediaTypeTrait
	any.previewTrait = bo.previewTrait

	any.baseTrait = bo.baseTrait

	return
}
