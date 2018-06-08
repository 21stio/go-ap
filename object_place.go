package ap

type UnitType string

const (
	CM     UnitType = "cm"
	FEET   UnitType = "feet"
	INCHES UnitType = "inches"
	KM     UnitType = "km"
	M      UnitType = "m"
	MILES  UnitType = "miles"
)

type placeTrait struct {
	Accuracy  float64  `json:"accuracy,omitempty"`
	Altitude  float64  `json:"altitude,omitempty"`
	Latitude  float64  `json:"latitude,omitempty"`
	Longitude float64  `json:"longitude,omitempty"`
	Radius    float64  `json:"radius,omitempty"`
	Units     UnitType `json:"units,omitempty"`
}

type place struct {
	*baseObject
	placeTrait
}

func NewPlace() (place place) {
	place.Type = string(OBJECT_PLACE)

	return
}

func (p place) ToObjectLink() (ol *ObjectLink) {
	ol = p.baseObject.ToObjectLink()
	ol.placeTrait = p.placeTrait

	return
}

func (p place) ToAnyObject() (ao *anyObject) {
	ao = p.baseObject.ToAnyObject()
	ao.placeTrait = p.placeTrait

	return
}