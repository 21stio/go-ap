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
