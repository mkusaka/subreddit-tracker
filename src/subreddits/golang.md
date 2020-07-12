# golang
## [1][Go Channel Use Cases](https://www.reddit.com/r/golang/comments/hpnlc4/go_channel_use_cases/)
- url: https://go101.org/article/channel-use-cases.html
---

## [2][Going Driverless (Oscar Söderlund, GopherCon EU 2020)](https://www.reddit.com/r/golang/comments/hpt2c4/going_driverless_oscar_söderlund_gophercon_eu_2020/)
- url: https://www.youtube.com/watch?v=IbggJHJUv0U
---

## [3][Gentle: code-first GraphQL server](https://www.reddit.com/r/golang/comments/hpfb1o/gentle_codefirst_graphql_server/)
- url: https://www.reddit.com/r/golang/comments/hpfb1o/gentle_codefirst_graphql_server/
---
I'm creating a code-first GraphQL server (Gentle) in Go which inspired by [nexus](https://www.nexusjs.org).

## Why?

with project like [Graphile](https://graphile.org) or [Hasura](https://hasura.io/) you'll get a very nice developer experience because all you need to do is defining your DB schema and your GraphQL API will be ready in no time. they might works very well for small projects (which all the logic can be defined in DB) but things can get ugly for bigger projects:

* DB logics get complicated
* Authentication
* Authorization (for example maintaining RLS in PostgreSQL can become nightmare)
* performance
* deployment
* migration
* you still have to write codes for custom queries and mutations you might end up with lots of duplicate codes (for example you might have to re-define types in TS or for db query results, etc)
* only support PostgreSQL

recently I've learned about nexusjs which solves some of mentioned problems (thanks to Prisma) which seems very nice so I decided to create something similar in Go by generating some CRUD code from [Ent](https://entgo.io/) models and feed it gqlgen or [Thunder](https://github.com/samsarahq/thunder). long story short after  struggles and facing weird bugs (specially with Thunder) I've decided to create a code-first GraphQL server (at this moment it uses gqlgen executor).

## Code Example

your GraphQL server code will going to look like this:

```go
package graph

import (
        "context"

	"github.com/sijad/gentle"
)

// gentle will looks for Query, Mutation and Subscription named types
type Query struct {}

// only ctx and args are reserved names, other arguments will treated as dependencies
func (q *Query) Hello(args struct{Name: string}, ctx context.Context) string {
  return "Hello " + args.Name
}

type Mutation struct {}

type SaveNameInput struct {
  Name string
}

type SaveNameArgs struct {
  Input SaveNameInput
}

type SaveNamePayload struct {
  name string
}

func (q *Query) SaveName(db sql.DB, args SaveNameArgs, myOtherDependency MyDependency) (*SaveNamePayload, error) {
  if err := db.SaveName(args.Input.Name); err != nil {
    return nil, err
  }
  return &amp;SaveNamePayload{args.Input.Name}
}
```

## Status

Gentle is in early stages and it's not ready, I'd love to get some feedback to make sure it'll be usable and moving to a right path. please feel free contribute.
## [4][What feature of Go is used very often by experienced programmers, but not so much by a newbie?](https://www.reddit.com/r/golang/comments/hp4mk3/what_feature_of_go_is_used_very_often_by/)
- url: https://www.reddit.com/r/golang/comments/hp4mk3/what_feature_of_go_is_used_very_often_by/
---

## [5][Tidy Cobra command option storage](https://www.reddit.com/r/golang/comments/hpj60f/tidy_cobra_command_option_storage/)
- url: https://www.reddit.com/r/golang/comments/hpj60f/tidy_cobra_command_option_storage/
---
A simple technique for keeping Cobra command option storage close to the Run/RunE command handler 

[http://blog.carpediem.org/posts/tidy-cobra/](http://blog.carpediem.org/posts/tidy-cobra/)
## [6][Difference between generics and parametric polymorphism?](https://www.reddit.com/r/golang/comments/hpcpsw/difference_between_generics_and_parametric/)
- url: https://www.reddit.com/r/golang/comments/hpcpsw/difference_between_generics_and_parametric/
---
In the [Type Parameters Draft Design](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md), the authors refer to their design as introducing a "form of parametric polymorphism".

Shortly after, they point out the following:

&gt; As the term *generic* is widely used in the Go community, we will use it below as a shorthand to mean a function or type that takes type parameters. Don't confuse the term generic as used in this design with the same term in other languages like C++, C#, Java, or Rust; they have similarities but are not the same.

After researching this for a while, I came to the conclusion that the terms "generics" and "parametric polymorphism" actually identify the same thing.

However the quote above shows that there is indeed a difference. I'm not familiar enough with those other languages to tell what that difference is.

Can someone explain it in further detail?
## [7][image to gotk3 drawing area](https://www.reddit.com/r/golang/comments/hpryrc/image_to_gotk3_drawing_area/)
- url: https://www.reddit.com/r/golang/comments/hpryrc/image_to_gotk3_drawing_area/
---
Looking for some examples of displaying golfing image data to a gotk3 DrawingArea

Not clear how to set individual pixels of the drawing area. 

thanks, srini
## [8][Command-line flag with environment variables](https://www.reddit.com/r/golang/comments/hpls5z/commandline_flag_with_environment_variables/)
- url: https://www.reddit.com/r/golang/comments/hpls5z/commandline_flag_with_environment_variables/
---
A basic package adding environment variables option to flags. Flags can be set with command-line flags or environment variable.  


```none
$ go run main.go -name "coulson" -level 7

$ NAME=coulson LEVEL=7 go run main.go
``` 


 [https://github.com/kahlys/flagenv](https://github.com/kahlys/flagenv)
## [9][Help with Websockets (Gorilla)](https://www.reddit.com/r/golang/comments/hph33m/help_with_websockets_gorilla/)
- url: https://www.reddit.com/r/golang/comments/hph33m/help_with_websockets_gorilla/
---
I have an app that needs to provide notifications to certain users in real time.  Currently the app is polling the server every ten seconds, but I am trying implement websockets.

I am able to open websockets and verify the user by passing in a JWT through the opened connection.  My problem is with how to store these connections.

I've learned that I can't store websockets in a database, so I'm not sure what the best approach is.  I see that a lot of people making chat applications use redis Pub Sub, but I don't think the pub sub architecture is actually what I need (but I'm open to being corrected here).  Dealing with channels on redis seems like a bunch of extra stuff that I don't really need/want to deal with.

In my backend I already have all the code that sorts out which users need to get the notifications.  So I just pass the userId to my webSocketNotification function which writes the message.  But then I need to get the correct websocket connection for each user.  Can I simply store my user id and websocket connection in redis?  I think I can just make a slice that stores structs with userId and the websocket connection but I believe that will run into race conditions, and I don't think it is scaleable.

Any advice on what to do here would be very much appreciated.
## [10][goreleaser is failing on a version tag I can't find](https://www.reddit.com/r/golang/comments/hpou5a/goreleaser_is_failing_on_a_version_tag_i_cant_find/)
- url: https://www.reddit.com/r/golang/comments/hpou5a/goreleaser_is_failing_on_a_version_tag_i_cant_find/
---
I know this is pilot error, but can't figure out what's happening. When I run `goreleaser init` I get this error:

`release failed after 20.13s error=failed to build for darwin_amd64: go: finding module for package github.com/tomcam/mb/pkg/app`

`go: downloading github.com/tomcam/mb v0.1.9`

`main.go:4:2: module github.com/tomcam/mb@latest found (v0.1.9), but does not contain package` [`github.com/tomcam/mb/pkg/app`](https://github.com/tomcam/mb/pkg/app)

I don't know where the v0.1.9 build tag is coming from. Can someone tell me where it finds that information?
