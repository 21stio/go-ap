package ap

type ActorType ObjectType

const (
	ACTOR_APPLICATION  ActorType = "Application"
	ACTOR_GROUP        ActorType = "Group"
	ACTOR_ORGANIZATION ActorType = "Organization"
	ACTOR_PERSON       ActorType = "Person"
	ACTOR_SERVICE      ActorType = "Service"
)

type Endpoint struct {
	ProxyUrl string
	OauthAuthorizationEndpoint string
	OauthTokenEndpoint string
	ProvideClientKey string
	SignClientKey string
	SharedInbox orderedCollection
}

type actorTraits struct {
	Inbox orderedCollection `json:"inbox,omitempty"`
	Outbox orderedCollection `json:"outbox,omitempty"`
	Following *orderedCollection `json:"following,omitempty"`
	Followers *orderedCollection `json:"followers,omitempty"`
	Liked *orderedCollection `json:"type,omitempty"`
	Streams []orderedCollection `json:"streams,omitempty"`
	Endpoints []Endpoint

	PrefferedUsername string
	Name string
	Summary string
}

type Actor struct {
	idTrait
	actorTraits
}

func NewActor(t ActorType) (o baseObject) {
	o = NewBaseObject(ObjectType(t))

	return
}
