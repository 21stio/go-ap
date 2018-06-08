package main

import (
	"github.com/21stio/go-ap/ap"
	"encoding/json"
	"log"
)

func main() {
	err := func() (err error) {
		actor := ap.NewBaseObject(ap.ObjectType(ap.ACTOR_PERSON))
		actor.Id = "Alyssa"

		to := ap.NewActor(ap.ACTOR_PERSON)
		to.Name = "Ben"

		object := ap.NewBaseObject(ap.OBJECT_NOTE)
		object.Content = "Say, did you finish reading that book I lent you?"

		activity := ap.NewActivity(ap.CASE_CONTENT_CREATE)
		activity.Id = "https://social.example/alyssa/posts/a29a6843-9feb-4c74-a7f7-081b9c9201d3"
		activity.Actor = actor.ToObjectLink()
		activity.To = ap.ToObjectLinks(to)
		activity.Object = object.ToAnyObject()

		j, err := json.MarshalIndent(activity, "", "   ")
		if err != nil {
			return
		}

		activity, err = ap.NewActivityFromJson(j)
		if err != nil {
			return
		}

		if activity.Actor.IsLink() == false {
			//actor := activity.Actor.ToObject()
		}

		println(string(j))

		return
	}()
	if err != nil {
		log.Fatal(err)
	}

}
