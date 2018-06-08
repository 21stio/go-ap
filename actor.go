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
	ProxyUrl                   string                 `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint string                 `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string                 `json:"oauthTokenEndpoint,omitempty"`
	ProvideClientKey           string                 `json:"provideClientKey,omitempty"`
	SignClientKey              string                 `json:"signClientKey,omitempty"`
	SharedInbox                *OrderedCollectionLink `json:"sharedInbox,omitempty"`
}

type actorTrait struct {
	Inbox     *OrderedCollectionLink `json:"inbox,omitempty"`
	Outbox    *OrderedCollectionLink `json:"outbox,omitempty"`
	Following *CollectionLink        `json:"following,omitempty"`
	Followers *CollectionLink        `json:"followers,omitempty"`
	Liked     *CollectionLink        `json:"type,omitempty"`
	Streams   []CollectionLink       `json:"streams,omitempty"`
	Endpoints Endpoint               `json:"endpoints,omitempty"`

	PrefferedUsername string `json:"preferredUsername,omitempty"`
	Name              string `json:"name,omitempty"`
	Summary           string `json:"summary,omitempty"`

	Icon *ObjectLink `json:"icon,omitempty"`
}

type Actor struct {
	idTrait
	actorTrait
}

func NewActor(t ActorType) (a Actor) {
	a.Type = string(t)

	return
}
