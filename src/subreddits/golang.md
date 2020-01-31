# golang
## [1][How to optimize PDFs using Go](https://www.reddit.com/r/golang/comments/ewmauy/how_to_optimize_pdfs_using_go/)
- url: https://www.reddit.com/r/golang/comments/ewmauy/how_to_optimize_pdfs_using_go/
---
You can optimize PDFs using Go language. 

The linked article explains how you can compress PDFs using a library based on Go. You can even create new PDFs or edit old ones but you may need to have a basic understanding of the PDF structure to get it working efficiently. 

  
 [https://unidoc.io/news/compressing-and-optimizing-pdfs-golang](https://unidoc.io/news/compressing-and-optimizing-pdfs-golang)
## [2][Noise v1.1.2: Fearless, decentralized p2p networking in Go.](https://www.reddit.com/r/golang/comments/ewm9wd/noise_v112_fearless_decentralized_p2p_networking/)
- url: https://medium.com/perlin-network/noise-v1-1-2-fearless-decentralized-p2p-networking-in-go-bf3afdd77230
---

## [3][A list of Go GUI projects](https://www.reddit.com/r/golang/comments/ew7odi/a_list_of_go_gui_projects/)
- url: https://github.com/go-graphics/go-gui-projects
---

## [4][Micro v2.0.0 (gRPC by default)](https://www.reddit.com/r/golang/comments/ewmpz9/micro_v200_grpc_by_default/)
- url: https://micro.mu/blog/2020/01/30/micro-v2.html
---

## [5][vim-go v1.22](https://www.reddit.com/r/golang/comments/ew975j/vimgo_v122/)
- url: https://github.com/fatih/vim-go/releases/tag/v1.22
---

## [6][What are some decent feature tools compatible with Go?](https://www.reddit.com/r/golang/comments/ewo7y9/what_are_some_decent_feature_tools_compatible/)
- url: https://www.reddit.com/r/golang/comments/ewo7y9/what_are_some_decent_feature_tools_compatible/
---


   Hello.  I'm trying to introduce feature flags to my team, as I've pitched a pretty large migration and my strategy is to make some use of feature flags in this migration.  Without getting too specific we currently have an event driven architecture with listeners, a service layer (for APIs), a message bus and some workers/adapters.  I'm dealing with an extremely young team so many of the engineers on my staff do not know what feature flags are.  

   I'm looking for a feature flag library that probably isn't going to cost much (though money isn't an issue).   That's very lightweight and easy to maintain.   I've looked at this

https://github.com/vsco/dcdr

And it looks promising.  But it doesn't appear to be actively maintained.   I would love for some of you to weigh in
## [7][go-mod-upgrade - Update outdated Go dependencies interactively](https://www.reddit.com/r/golang/comments/ewnzwo/gomodupgrade_update_outdated_go_dependencies/)
- url: https://www.reddit.com/r/golang/comments/ewnzwo/gomodupgrade_update_outdated_go_dependencies/
---
Hi,

I created a small tool to update outdated Go dependencies interactively.

You can find the source code as well as instructions on how to install and use it here:

[https://github.com/oligot/go-mod-upgrade](https://github.com/oligot/go-mod-upgrade)
## [8][[Code Review] Omdb api wrapper](https://www.reddit.com/r/golang/comments/ewnfge/code_review_omdb_api_wrapper/)
- url: https://www.reddit.com/r/golang/comments/ewnfge/code_review_omdb_api_wrapper/
---
Hi,

Can someone give some feedback about my OMDB api wrapper. 

[https://github.com/matad2k/GOmdb/tree/master/goomdb](https://github.com/matad2k/GOmdb/tree/master/goomdb)
## [9][dynamic REST routing in Go real world examples](https://www.reddit.com/r/golang/comments/ewn3yq/dynamic_rest_routing_in_go_real_world_examples/)
- url: https://www.reddit.com/r/golang/comments/ewn3yq/dynamic_rest_routing_in_go_real_world_examples/
---
Hey folks,

do you know any nice real world examples for dynamic rest routing done in Go? I'm reading through blogs n stuff, but its always that basic examples of a product shop or something that might not cover problems real applications are facing.
## [10][Can Two TestMain run in parallel?](https://www.reddit.com/r/golang/comments/ewn09l/can_two_testmain_run_in_parallel/)
- url: https://www.reddit.com/r/golang/comments/ewn09l/can_two_testmain_run_in_parallel/
---
I have written two integration tests, one for db layer and other for service layer and I have two TestMain to do the db setup and tear down. Also, I am using the same db (mongo) in both the tests. My tests run fine locally but fails during the pipeline run with the error:

`\`error[(DatabaseDropPending) Cannot create collection 5ca5a2c479358e67a64f927b_testdb.users - database is in the process of being dropped.]\``

I believe the two TestMain are running in parallel and while one of them is dropping the db, the other one is trying to create indexes on that db. Can anyone please help?
