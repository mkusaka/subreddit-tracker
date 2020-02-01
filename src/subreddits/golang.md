# golang
## [1][Python (Flask) web developer interested in learning more about GoLang!](https://www.reddit.com/r/golang/comments/ex40nx/python_flask_web_developer_interested_in_learning/)
- url: https://www.reddit.com/r/golang/comments/ex40nx/python_flask_web_developer_interested_in_learning/
---
Hey GoLangeers! :)
I am interested in learning about web development in Go, I already know the basics of Go and want to learn more about building web apps.
What resources do y'all recommend?
## [2][Why use go at small scale when you have the convenience of rails and flask/django?](https://www.reddit.com/r/golang/comments/ex18cx/why_use_go_at_small_scale_when_you_have_the/)
- url: https://www.reddit.com/r/golang/comments/ex18cx/why_use_go_at_small_scale_when_you_have_the/
---
Sorry if this post is inappropriate for the sub, but I've really been struggling to digest the reasoning by the verbosity of go web development (mainly in HTTP template rendering/value insertion and static content rendering). Is there a reason why one would recommend go at a smaller scale over an MVC framework/scripting language? Thanks again.
## [3][Thoughts about real time communication from SPA to and from Go client over websocket](https://www.reddit.com/r/golang/comments/ewzkc1/thoughts_about_real_time_communication_from_spa/)
- url: https://www.reddit.com/r/golang/comments/ewzkc1/thoughts_about_real_time_communication_from_spa/
---
I've been working for about 1 1/2 years on and off on a multi-forum application  that is written in Angular on the client side and Go on the server side.

In Angular it uses ngrx/data which requires a simple data structure and RESTful communication.

On the Go side it's postgres for the time being. It doesn't use NOTIFY/LISTEN because of the 8MB limit but my own little invention that is essentially a channel that sends updates via websocket to the clients.

The message is simple
```JSON
{
"table": "table-name",
"action:" ["UPDATE","INSERT","DELETE"],
"data": {"id":"entity_id", "the":"payload"}
}
```

This way the client is always up to date. The sent message gets filtered on the client and the appropriate cached entities are updated, inserted or deleted.

On initial application load (client side) the client requests all data from the server. At some point I'll have to limit this to let's say the last 2 years and the rest should then be lazily loaded on demand.

The problem: imagine 10000 communities with 10 categories each with 100 forums each with 1000s of threads and 1000000s of posts in them and all emitting changes.
What happens in community A when I'm in community B isn't relevant for me right now.
Only when I enter community A should changes in community A and its related entities be transmitted. 
I could send location updates every time I navigate to a new community, and that works, however I have to go through a non-trivial navigation through the entity hierarchy, yes I could attach the communityID to every entity but that seems non-elegant.

There has to be a better way. A better pattern.

I read about MQTT, pubsub etc.

If I would use MQTT or some other kind of pubsub protocol or "technology", since messages are stored on the broker and the clients are receiving messages depending on last time active, that could also be a "bad way to go"TM.

If I switch to purely SSR (server side rendered, with Go) I lose the real time feature and make myself vulnerable to spambots, which currently doesn't seem to be able to deal with SPA.

How does, for instance, firebase handle this problem? Or is there an already established path, I'm sure I can't be the only one with that issue.
## [4][Go: How Are Loops Translated to Assembly?](https://www.reddit.com/r/golang/comments/ex5z4o/go_how_are_loops_translated_to_assembly/)
- url: https://medium.com/a-journey-with-go/go-how-are-loops-translated-to-assembly-835b985309b3
---

## [5][0 to Go, building a CLI for Postwoman](https://www.reddit.com/r/golang/comments/ex5sl9/0_to_go_building_a_cli_for_postwoman/)
- url: https://www.reddit.com/r/golang/comments/ex5sl9/0_to_go_building_a_cli_for_postwoman/
---
Wrote a blog post on Building a CLI

https://blog.athulcyriac.co/pwcli
## [6][Framework for automatic firmware reverse-engineering written in golang](https://www.reddit.com/r/golang/comments/ex4wfv/framework_for_automatic_firmware/)
- url: https://github.com/9elements/autorev
---

## [7][GoWest Conference tickets](https://www.reddit.com/r/golang/comments/ewx8ke/gowest_conference_tickets/)
- url: https://www.reddit.com/r/golang/comments/ewx8ke/gowest_conference_tickets/
---
We are happy to announce that tickets for the first ever GoWest Conference are officially on sale. This one day conference will be held on May 8th in Sandy, Utah. We are also offering 2 pre-conference workshops the day before the conference on May 7th. The workshop tickets are sold separately. Check out the link to purchase your tickets now. We have a limited number of early bird tickets so hurry up and get yours! The site with additional information about speakers and content will be up shortly. [https://ti.to/go-west-conf/2020](https://ti.to/go-west-conf/2020)
## [8][Next steps for pkg.go.dev](https://www.reddit.com/r/golang/comments/ewpw5e/next_steps_for_pkggodev/)
- url: https://blog.golang.org/pkg.go.dev-2020
---

## [9][GitHub - el10savio/goCircularRingBuffer: A bounded circular ring buffer queue implemented in Go](https://www.reddit.com/r/golang/comments/ex58hb/github_el10saviogocircularringbuffer_a_bounded/)
- url: https://github.com/el10savio/goCircularRingBuffer
---

## [10][what is the good thing that makes my minimal web site?](https://www.reddit.com/r/golang/comments/ex4twt/what_is_the_good_thing_that_makes_my_minimal_web/)
- url: https://www.reddit.com/r/golang/comments/ex4twt/what_is_the_good_thing_that_makes_my_minimal_web/
---
At first, this is my first post at reddit. and i'm not good at english so you are hard to understant my word☺️.

in short, i want to make my blog(will upload my pet 'bori's photo, and some text). so i studied wordpress. but some days ago, i knew if i use go framework, i can make blog like WP. some docu said go use effective memory, and if i use that my page will be fast than WP. plus, i don't need to concern about WP version &amp; it's plugins version. 

Anyway So i want to use Go language. i don't know well about Go. how can i make system structure(framework, db, and something...).

now i read Go guide, but if you have some advise please reply without hesitation.

thanks~
