package ap

type ActivityType string

const (
	ACTIVITY_NIL             ActivityType = ""
	ACTIVITY_ACCEPT          ActivityType = "Accept"
	ACTIVITY_ADD             ActivityType = "Add"
	ACTIVITY_ANNOUNCE        ActivityType = "Announce"
	ACTIVITY_ARRIVE          ActivityType = "Arrive"
	ACTIVITY_BLOCK           ActivityType = "Block"
	ACTIVITY_CREATE          ActivityType = "Create"
	ACTIVITY_DELETE          ActivityType = "Delete"
	ACTIVITY_DISLIKE         ActivityType = "Dislike"
	ACTIVITY_FLAG            ActivityType = "Flag"
	ACTIVITY_FOLLOW          ActivityType = "Follow"
	ACTIVITY_IGNORE          ActivityType = "Ignore"
	ACTIVITY_INVITE          ActivityType = "Invite"
	ACTIVITY_JOIN            ActivityType = "Join"
	ACTIVITY_LEAVE           ActivityType = "Leave"
	ACTIVITY_LIKE            ActivityType = "Like"
	ACTIVITY_LISTEN          ActivityType = "Listen"
	ACTIVITY_MOVE            ActivityType = "Move"
	ACTIVITY_OFFER           ActivityType = "Offer"
	ACTIVITY_QUESTION        ActivityType = "Question"
	ACTIVITY_REJECT          ActivityType = "Reject"
	ACTIVITY_READ            ActivityType = "Read"
	ACTIVITY_REMOVE          ActivityType = "Remove"
	ACTIVITY_TENTATIVEREJECT ActivityType = "TentativeReject"
	ACTIVITY_TENTATIVEACCEPT ActivityType = "TentativeAccept"
	ACTIVITY_TRAVEL          ActivityType = "Travel"
	ACTIVITY_UNDO            ActivityType = "Undo"
	ACTIVITY_UPDATE          ActivityType = "Update"
	ACTIVITY_VIEW            ActivityType = "View"
)

func NewActivity(t ActivityType) (activity activity) {
	activity.Type = string(t)

	return
}

type intransitiveActivityTrait struct {
	Actor          *ObjectLink  `json:"actor,omitempty"`
	Target         *ObjectLink  `json:"target,omitempty"`
	Result         *ObjectLink  `json:"result,omitempty"`
	Origin         *ObjectLink  `json:"origin,omitempty"`
	Instrument     *ObjectLink  `json:"instrument,omitempty"`
}

type activityTrait struct {
	Object *anyObject   `json:"object,omitempty"`
}

type activity struct {
	idTrait
	baseTrait
	intransitiveActivityTrait
	activityTrait
}
