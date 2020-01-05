# golang
## [1][Building a BitTorrent client from the ground up in Go](https://www.reddit.com/r/golang/comments/ek4dp3/building_a_bittorrent_client_from_the_ground_up/)
- url: https://blog.jse.li/posts/torrent/
---

## [2][Just wrote up a cookbook for IO in Go](https://www.reddit.com/r/golang/comments/ek8nu9/just_wrote_up_a_cookbook_for_io_in_go/)
- url: https://www.reddit.com/r/golang/comments/ek8nu9/just_wrote_up_a_cookbook_for_io_in_go/
---
https://github.com/jesseduffield/notes/wiki/Golang-IO-Cookbook

I wrote this mostly for myself to consolidate the stuff I learnt when changing [Horcrux](https://github.com/jesseduffield/horcrux) to stream file data rather than load it all into memory. Maybe it can be of use to other people trying to get their heads around this stuff :)

Please let me know if there's any corrections/changes you think I should make. And thanks again for all those who gave feedback on the original post about Horcrux!
## [3][New to Go, not to programming: tutorial frustration](https://www.reddit.com/r/golang/comments/ejtm8n/new_to_go_not_to_programming_tutorial_frustration/)
- url: https://www.reddit.com/r/golang/comments/ejtm8n/new_to_go_not_to_programming_tutorial_frustration/
---
I've been programming for years in multiple languages, and I'm frustrated with the Go tutorials. They all teach the same thing: syntax. After learning the pieces (loops, slices, structs, etc.) in the end I'm left without knowing how to write an application. 

I have yet to see anything in the tutorials that is more than a single file,  package main,  func main. If I search for packages, I see nothing but importing packages into that single file. 

Am I missing something? Where can I find info on creating multi file applications? Code organization? I really want to delve into this but the tutorials are too basic to apply to useful functionality.

Am I not getting far enough? Writing switch statements and for loops is not very satisfying.

EDIT: Thank you all for the many responses! There looks to be many great resources and suggestions here. Upvote for the community!
## [4][How to extend a struct dynamically?](https://www.reddit.com/r/golang/comments/ekc11n/how_to_extend_a_struct_dynamically/)
- url: https://www.reddit.com/r/golang/comments/ekc11n/how_to_extend_a_struct_dynamically/
---
I have a struct St1. In some places I use it as is -- as an array of St1. In other place I need to be able to dynamically attach extra data to each struct, **akin to adding a few additional fields \*dynamically\* to each St1.**

&amp;#x200B;

What's a way to go?
## [5][What "Build a &lt;X&gt; With Go" deserves it's own online course](https://www.reddit.com/r/golang/comments/ekb54y/what_build_a_x_with_go_deserves_its_own_online/)
- url: https://www.reddit.com/r/golang/comments/ekb54y/what_build_a_x_with_go_deserves_its_own_online/
---
I'm a fanatic Go programmer and have quite some experience in teaching others. I believe the best way of teaching is to build something concrete! So what do you think would be the best topic of such an online course? 

First thing that comes to mind is the creation of an API or Web App but maybe that is already covered in-depth by others? 

I'm curious what you think!
## [6][How do I handle interdependent modules?](https://www.reddit.com/r/golang/comments/ekai39/how_do_i_handle_interdependent_modules/)
- url: https://www.reddit.com/r/golang/comments/ekai39/how_do_i_handle_interdependent_modules/
---
For example lets say we have an online chatroom. Users can join a room and post messages there and the messages are delivered to other users in the room via some realtime delivery mechanism.

So we have a `roomService` struct.

    type roomService struct {
        room      messagerooms.RoomRepository 
        message   messagerooms.MessageRepository 
        publisher pubsub.Service 
    }

Here `roomService` is dependent on the `pubsub` service because every time a new message is posted, that message is sent to the pubsub system for distributing accross the room members via maybe an websocket connection.

My Problem here is that if an user tries to subscribe to the message streams for a room, before subscribing the user to the stream, the `pubsub` service needs to check if that user is actually part of the room. And thus the `pubsub` service becomes dependent on the `roomService`  for deciding if the user is part of the room, which is creating a circular dependency.

So how to handle situations like this?
## [7][sshtargate - Host SSH portals to applications](https://www.reddit.com/r/golang/comments/ek3uez/sshtargate_host_ssh_portals_to_applications/)
- url: https://git.sr.ht/~tslocum/sshtargate
---

## [8][Guys, using gorilla mux, a router and it's subrouter, are they different entities or are they connected?](https://www.reddit.com/r/golang/comments/ek24zu/guys_using_gorilla_mux_a_router_and_its_subrouter/)
- url: https://www.reddit.com/r/golang/comments/ek24zu/guys_using_gorilla_mux_a_router_and_its_subrouter/
---
Does a middlware attached to the parent router affect the subrouter?
Thank you
## [9][How to Find All Packages that implements io.Reader](https://www.reddit.com/r/golang/comments/ek7k6c/how_to_find_all_packages_that_implements_ioreader/)
- url: https://www.reddit.com/r/golang/comments/ek7k6c/how_to_find_all_packages_that_implements_ioreader/
---
Disclaimer: I am a noob, so I apologize in advance if the terminology is incorrect.

The [csv.NewReader\(\)](https://golang.org/pkg/encoding/csv/#NewReader) from  re
quires an io.Reader as an input. And [os.File](https://golang.org/pkg/os/#File.Read)
has this same method so os.File type implements the io.Reader.

My question is how to find all the packages in the standard library that implements the io.Reader. If the input needs an io.Reader, then it is not a problem. It is those methods like File.Read where the method is the same as the io.Reader.
## [10][return (ok bool) instead of error?](https://www.reddit.com/r/golang/comments/ek2nib/return_ok_bool_instead_of_error/)
- url: https://www.reddit.com/r/golang/comments/ek2nib/return_ok_bool_instead_of_error/
---
I am writing a function that is supposed to take a string, and return the first rune along with a new string where that character has been chopped off, e.g., `"hello" -&gt; ('h', "ello")`. However, I want to explicitly handle the case of the string being empty by returning an `error`. But as I think about it, this error shouldn't carry any information other than whether or not it exists; it shouldn't need to have the `Error()` function defined. So it feels like using the `error` interface is a bit heavy. I see that the language also commonly uses `(ok bool)` as a return type to indicate success or failure, but that seems limited to language features like type assertions. My third option is just to return a `nil` string that the caller has to check for, but that doesnt seem idiomatic either.

Would it be good practice to return `(rune, string, ok bool)`? Or should I still just return some instance of `error`? I come from Rust, and in this case I would have returned an `Option&lt;(char, &amp;str)&gt;`, which seems more analogous to a bool than an error.

Thanks for any reply!

Edit: Thanks to numerous responses, especially from u/ruertar, it seems like `error` is still the best option here. Thank you very much for all the discussion! It has been very interesting to see everyone's opinions :)
