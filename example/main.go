package main

import (
	"github.com/21stio/go-ap"
	"encoding/json"
	"log"
)

func main() {
	encoded, err := encodeActor()
	if err != nil {
		log.Fatal(err)
	}

	err = decodeActor(encoded)
}

func encodeExample5() (encoded []byte, err error) {
	actorLink := ap.NewLink("https://social.example/alyssa/")

	toLink := ap.NewLink("https://chatty.example/ben/")

	object := ap.NewBaseObject(ap.OBJECT_NOTE)
	object.Content = "Say, did you finish reading that book I lent you?"

	activity := ap.NewActivity(ap.ACTIVITY_CREATE)
	activity.Id = "https://social.example/alyssa/posts/a29a6843-9feb-4c74-a7f7-081b9c9201d3"
	activity.Actor = actorLink.ToObjectLink()
	activity.To = ap.ToObjectLinks(toLink)
	activity.Object = object.ToAnyObject()

	encoded, err = json.MarshalIndent(activity, "", "   ")
	if err != nil {
		return
	}

	println("ENCODED:\n")
	println(string(encoded))

	return
}

func encodeOutbox() (encoded []byte, err error) {
	oc := ap.NewOrderedCollection()
	oc.Context = []string{
		"https://www.w3.org/ns/activitystreams",
		"https://w3id.org/security/v1",
	}
	oc.Id = "https://mastodon.social/users/kiriska/outbox"
	oc.TotalItems = 412

	firstLink := ap.NewLink("https://mastodon.social/users/kiriska/outbox?page=true")
	oc.First = firstLink.ToOrderedCollectionPageLink()

	lastLink := ap.NewLink("https://mastodon.social/users/kiriska/outbox?min_id=0&page=true")
	oc.Last = lastLink.ToOrderedCollectionPageLink()

	encoded, err = json.MarshalIndent(oc, "", "   ")
	if err != nil {
		return
	}

	println("ENCODED:\n")
	println(string(encoded))

	return
}

func encodeActor() (encoded []byte, err error) {
	a := ap.NewActor(ap.ACTOR_PERSON)
	a.Context = []string{
		"https://www.w3.org/ns/activitystreams",
		"https://w3id.org/security/v1",
	}
	a.Id = "https://mastodon.social/users/kiriska"
	a.Following = ap.NewLink("https://mastodon.social/users/kiriska/following").ToCollectionLink()
	a.Followers = ap.NewLink("https://mastodon.social/users/kiriska/followers").ToCollectionLink()
	a.Inbox = ap.NewLink("https://mastodon.social/users/kiriska/inbox").ToOrderedCollectionLink()
	a.PrefferedUsername = "kiriska"
	a.Name = "Kiri"
	a.Summary = "<p>Artist, writer, blogger. Eldritch beast. Brush pen enthusiast. Perpetual motion machine.</p><p>kiriska@gmail.com<br /><a href=\"http://kiriska.com\" rel=\"nofollow noopener\" target=\"_blank\"><span class=\"invisible\">http://</span><span class=\"\">kiriska.com</span><span class=\"invisible\"></span></a></p>"

	sharedInbox := ap.NewOrderedCollection()
	sharedInbox.OrderedItems = ap.ToObjectLinks(
		ap.NewLink("https://mastodon.social/inbox"),
		ap.NewLink("https://mastodon.social/super/inbox"),
	)

	a.Endpoints = ap.Endpoint{
		SharedInbox: sharedInbox.ToOrderedCollectionLink(),
	}

	icon := ap.NewLink("https://files.mastodon.social/accounts/headers/000/014/402/original/31c305f052363eaf.jpg")
	icon.MediaType = ap.MEDIA_IMAGE_PNG

	a.Icon = icon.ToObjectLink()

	encoded, err = json.MarshalIndent(a, "", "   ")
	if err != nil {
		return
	}

	println("ENCODED:\n")
	println(string(encoded))

	return
}

func decodeActor(encoded []byte) (err error) {
	a := ap.Actor{}

	err = json.Unmarshal(encoded, &a)
	if err != nil {
		return
	}

	println("\nDECODED:\n")

	if a.Icon.IsLink() {
		iconLink := a.Icon.GetLink()
		println(iconLink.MediaType)
		println(iconLink.Href)
	}

	switch a.Endpoints.SharedInbox.Type {
	case string(ap.LINK_LINK):
		println(a.Endpoints.SharedInbox.GetLink().Href)
	case string(ap.COLLECTION_ORDERED):
		oc := a.Endpoints.SharedInbox.GetOrderedCollection()

		for _, item := range oc.OrderedItems {
			println(item.GetLink().Href)
		}
	}

	return
}

func encodeFollowing() (encoded []byte, err error) {
	oc := ap.NewOrderedCollection()
	oc.Context = []string{
		"https://www.w3.org/ns/activitystreams",
		"https://w3id.org/security/v1",
	}
	oc.Id = "https://mastodon.social/users/kiriska/following"
	oc.TotalItems = 29

	firstLink := ap.NewLink("https://mastodon.social/users/kiriska/following?page=2")
	oc.First = firstLink.ToOrderedCollectionPageLink()

	encoded, err = json.MarshalIndent(oc, "", "   ")
	if err != nil {
		return
	}

	println("ENCODED:\n")
	println(string(encoded))

	return
}

func encodeFollowingPage() (encoded []byte, err error) {
	oc := ap.NewOrderedCollectionPage()
	oc.Context = []string{
		"https://www.w3.org/ns/activitystreams",
		"https://w3id.org/security/v1",
	}
	oc.Id = "https://mastodon.social/users/kiriska/following?page=1"
	oc.TotalItems = 29

	nextLink := ap.NewLink("https://mastodon.social/users/kiriska/following?page=2")
	oc.Next = nextLink.ToOrderedCollectionPageLink()

	partOfLink := ap.NewLink("https://mastodon.social/users/kiriska/following")
	oc.PartOf = partOfLink.ToOrderedCollectionLink()

	oc.OrderedItems = ap.ToObjectLinks(
		ap.NewLink("https://mastodon.social/users/aainsleyy"),
		ap.NewLink("https://mastodon.art/users/eecks"),
		ap.NewLink("https://glitch.social/users/fish"),
		ap.NewLink("https://mastodon.art/users/elefluff"),
		ap.NewLink("https://mastodon.social/users/cutewitchirl"),
		ap.NewLink("https://mastodon.social/users/jemmasalume"),
		ap.NewLink("https://mastodon.art/users/ichi"),
		ap.NewLink("https://mastodon.social/users/binglinhu"),
		ap.NewLink("https://mastodon.social/users/zynchilada"),
		ap.NewLink("https://mastodon.art/users/shazzbaa"),
		ap.NewLink("https://monsterpit.net/users/gatorbites"),
		ap.NewLink("https://mastodon.social/users/locirodraws"),
	)

	encoded, err = json.MarshalIndent(oc, "", "   ")
	if err != nil {
		return
	}

	println("ENCODED:\n")
	println(string(encoded))

	return
}
