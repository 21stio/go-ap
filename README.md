## go-ap

A couple of structs and abstractions that allow you to easily construct, encode, decode and access valid ActivityPub objects.

###### Limitations

- can't encode/decode Links that are represented as a slice of strings
- can't render `null` as all field are tagged with ``json:"omitempty"``

#### Installation

`go get github.com/21stio/go-ap`

####Usage

```golang
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

println(string(encoded))

```

```json
{
   "@context": [
      "https://www.w3.org/ns/activitystreams",
      "https://w3id.org/security/v1"
   ],
   "id": "https://mastodon.social/users/kiriska",
   "inbox": {
      "type": "Link",
      "href": "https://mastodon.social/users/kiriska/inbox"
   },
   "following": {
      "type": "Link",
      "href": "https://mastodon.social/users/kiriska/following"
   },
   "followers": {
      "type": "Link",
      "href": "https://mastodon.social/users/kiriska/followers"
   },
   "endpoints": {
      "sharedInbox": {
         "type": "OrderedCollection",
         "orderedItems": [
            {
               "type": "Link",
               "href": "https://mastodon.social/inbox"
            },
            {
               "type": "Link",
               "href": "https://mastodon.social/super/inbox"
            }
         ]
      }
   },
   "preferredUsername": "kiriska",
   "name": "Kiri",
   "summary": "\u003cp\u003eArtist, writer, blogger. Eldritch beast. Brush pen enthusiast. Perpetual motion machine.\u003c/p\u003e\u003cp\u003ekiriska@gmail.com\u003cbr /\u003e\u003ca href=\"http://kiriska.com\" rel=\"nofollow noopener\" target=\"_blank\"\u003e\u003cspan class=\"invisible\"\u003ehttp://\u003c/span\u003e\u003cspan class=\"\"\u003ekiriska.com\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e",
   "icon": {
      "type": "Link",
      "mediaType": "image/png",
      "href": "https://files.mastodon.social/accounts/headers/000/014/402/original/31c305f052363eaf.jpg"
   }
}
```