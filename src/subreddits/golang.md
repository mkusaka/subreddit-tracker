# golang
## [1][Why Discord is switching from Go to Rust](https://www.reddit.com/r/golang/comments/eywx4q/why_discord_is_switching_from_go_to_rust/)
- url: https://blog.discordapp.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f
---

## [2][godoc.org will be closed owing to legal reason](https://www.reddit.com/r/golang/comments/ez7m26/godocorg_will_be_closed_owing_to_legal_reason/)
- url: https://groups.google.com/d/msg/golang-dev/mfiPCtJ1BGU/qtCrqlrEEwAJ
---

## [3][Continuous Profiling Go applications running in Kubernetes](https://www.reddit.com/r/golang/comments/ez8n40/continuous_profiling_go_applications_running_in/)
- url: https://gianarb.it/blog/continuous-profiling-go-apps-in-kubernetes
---

## [4][Go gRPC Cheap Ping](https://www.reddit.com/r/golang/comments/ez6j41/go_grpc_cheap_ping/)
- url: https://medium.com/@kainlite_32799/go-grpc-cheap-ping-ed7a629f05de
---

## [5][Hey gophers any suggestions like this to become employable with golang ?](https://www.reddit.com/r/golang/comments/ez8dvl/hey_gophers_any_suggestions_like_this_to_become/)
- url: https://i.redd.it/s0f5czaff3f41.jpg
---

## [6][Seeking new custodian for godoc.org](https://www.reddit.com/r/golang/comments/ez1wc9/seeking_new_custodian_for_godocorg/)
- url: https://groups.google.com/forum/#!topic/golang-nuts/OA-7lUbZJMY
---

## [7][How to search for (possibly multiple) nodes in a binary tree where all parents match criteria?](https://www.reddit.com/r/golang/comments/ez7qli/how_to_search_for_possibly_multiple_nodes_in_a/)
- url: https://www.reddit.com/r/golang/comments/ez7qli/how_to_search_for_possibly_multiple_nodes_in_a/
---
Hello, I'm writing a path tracer and to optimize it, I started implementing BVHs. These are binary trees that encapsulate objects in the scene. For example, let's say I have a model that consists of 100 triangles and I have 3 levels of BVH. The 1st level encapsulates whole model, the 2nd level consists of two branches, of 50 tris each, 3rd level has 4 branches of 25. Now I don't need to test for ray-triangle intersection for all 100 triangles, but I firstly test for intersection with 1st level of BVH. If it doesn't intersect, the ray won't intersect with any triangle in the model. If it does, I want to test with intersection with left and right branch. Now there's an issue: the ray can intersect with both branches, so I want to intersect it further until it goes to the last level. Let's say a ray intersects with three boxes at the last level of BVH. I want to get information about all the triangles in these boxes, not only one of them. I tried to implement it in level order tree traversal, like so:

`func levelTravelsal(tree *BVH, level int, r Ray, tMin, tMax float64) []Triangle {`  
 `temp := []Triangle{}`  
 `if tree == nil {`  
   `return nil`  
 `}`  
 `if tree.last {`  
   `if tree.bounds.hit(r, tMin, tMax) {`  
`return tree.tris`  
   `}`  
   `return nil`  
 `} else if level &gt; 1 {`  
   `if tree.left.bounds.hit(r, tMin, tMax) {`  
`temp = levelTravelsal(tree.left, level-1, r, tMin, tMax)`  
   `}`  
   `if tree.right.bounds.hit(r, tMin, tMax) {`  
`tr := levelTravelsal(tree.right, level-1, r, tMin, tMax)`  
`for i := 0; i &lt; len(tr); i++ {`  
`temp = append(temp, tr[i])`  
`}`  
   `}`  
 `}`  
 `return temp`  
`}`

But I get triangles from only one box. Here's the BVH data structure:

`type BVH struct {`  
`left, right *BVH`  
`tris        []Triangle`  
`bounds      AABB`  
`last        bool`  
`}`

Where left, right are branches, tris are triangles encapsulated in the BVH, bounds is the bounding box of the BVH and last is boolean that tells if it's the last level's BVH. So how could I implement this?
## [8][HTTP router performance](https://www.reddit.com/r/golang/comments/ez7cfg/http_router_performance/)
- url: https://www.reddit.com/r/golang/comments/ez7cfg/http_router_performance/
---
Hi,

I'm looking to make my own HTTP router so I can better learn GO and as a practice project, and who knows maybe I'll also use it in the future.

I've looked at how the http routers are done in frameworks like Koa, Express, Django, Flask, some PHP frameworks, and some GO ones and from what I noticed in non GO programming languages they mostly use regex to parse the routes and do simple for each route, does it match? using regex while in GO I've seen people mostly using a Radix tree (the famous [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) for example).

Is regex slow in Go? Is radix tree the best option? I'm just curious and trying to understand the logic why other frameworks are ok with regex based routes and in Go we mostly see radix being used.
## [9][Problem with go-guru](https://www.reddit.com/r/golang/comments/ez2yad/problem_with_goguru/)
- url: https://www.reddit.com/r/golang/comments/ez2yad/problem_with_goguru/
---
Hello!

so, I have issues with go-guru since moving to `go mod`, when trying to locate types defined in the current package I'm working it fails with the following:

```
guru: no object for identifier
```

I tried with an `interface{}` and with a `struct`. some times it works for some symbols and some times it does not for other set of symbols, and is annoying.
## [10][Golang homework interview challenge for Storj - Looking for some feedback](https://www.reddit.com/r/golang/comments/eyphsm/golang_homework_interview_challenge_for_storj/)
- url: https://www.reddit.com/r/golang/comments/eyphsm/golang_homework_interview_challenge_for_storj/
---
I've posted this question on code review in case you want some programming internet points [https://codereview.stackexchange.com/questions/236641/golang-homework-interview-challenge-for-storj](https://codereview.stackexchange.com/questions/236641/golang-homework-interview-challenge-for-storj) 

Basically I submitted this solution [https://github.com/Samyoul/storj-file-sender](https://github.com/Samyoul/storj-file-sender) in response to a code challenge from Storj as part of their interview process. Storj is the only company that I've interviewed with that has not given any kind of feedback at all. So perhaps you guys could give me some pointers.

Thank you and let the brutalities commence :D
