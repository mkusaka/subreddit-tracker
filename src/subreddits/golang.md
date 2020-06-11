# golang
## [1][Exploring container package in Go (list, ring and heap)](https://www.reddit.com/r/golang/comments/h0xjc2/exploring_container_package_in_go_list_ring_and/)
- url: https://therebelsource.com/blog/exploring-container-package-in-go-list-ring-and-heap/9zTBiMaaYg
---

## [2][Google authentication with GoLang and Goth](https://www.reddit.com/r/golang/comments/h0yrhk/google_authentication_with_golang_and_goth/)
- url: https://www.loginradius.com/engineering/blog/google-authentication-with-golang-and-goth/
---

## [3][An introduction to Go for non-Go developers](https://www.reddit.com/r/golang/comments/h0jc27/an_introduction_to_go_for_nongo_developers/)
- url: https://benhoyt.com/writings/go-intro/
---

## [4][Telebot V2.3 released: The most feature-complete bot framework for Telegram in Go](https://www.reddit.com/r/golang/comments/h0yz9o/telebot_v23_released_the_most_featurecomplete_bot/)
- url: https://github.com/tucnak/telebot/releases/tag/v2.3
---

## [5][UtahFS is an encrypted storage system that provides a user-friendly FUSE drive backed by cloud storage](https://www.reddit.com/r/golang/comments/h0szzl/utahfs_is_an_encrypted_storage_system_that/)
- url: https://github.com/cloudflare/utahfs
---

## [6][[P] spaGO: a work-in-progress yet ready-to-use library for ML/NLP entirely in Go](https://www.reddit.com/r/golang/comments/h0yys5/p_spago_a_workinprogress_yet_readytouse_library/)
- url: https://www.reddit.com/r/golang/comments/h0yys5/p_spago_a_workinprogress_yet_readytouse_library/
---
From today, Gophers can exploit the potential of state-of-the-art NLP technologies without moving to other languages. I am excited to share some results from my personal "NLP Odyssey": spaGO is a work-in-progress yet ready-to-use library for ML/NLP, written entirely in Go. You can create and train your own neural model, or import a pre-trained one from Hugging Face and FlairNLP. The deploy is just a single executable for your server (no heavy dependencies, no virtual environments).

[spaGO HTTP API for Question-Answering](https://preview.redd.it/irz8bu945a451.png?width=1811&amp;format=png&amp;auto=webp&amp;s=6a3f6efb4c010876ab478a1582812a9d6ae4d5cc)

Please check it out and let me know what you think!

The code is aviable at: [https://github.com/nlpodyssey/spago](https://github.com/nlpodyssey/spago)

If you like the project, please star the repository to show your support!

P.S. I warn you that for now, spaGO doesn't use GPU and there are a ton of optimizations to make it more performant. As a minimum, we can start with something!
## [7][Stream processing similar to Kafka](https://www.reddit.com/r/golang/comments/h0xtya/stream_processing_similar_to_kafka/)
- url: https://www.reddit.com/r/golang/comments/h0xtya/stream_processing_similar_to_kafka/
---
Hi.

Are there any stream processing systems similar to Kafka written in Golang? Why go, because I want the consumer/client to be first class citizen. I want an alternative for Kafka because of the annoying Zookeeper dependency. One requirements: If the client/consumer stops I want the ability to pick up where it left off.

How does nsq and nats stack up?
## [8][Can anyone translate relation sql to MongoDB?](https://www.reddit.com/r/golang/comments/h0ywdk/can_anyone_translate_relation_sql_to_mongodb/)
- url: https://www.reddit.com/r/golang/comments/h0ywdk/can_anyone_translate_relation_sql_to_mongodb/
---
Can anyone give me example in MongoDB to createOne and getOne for relationships

Let's say user has many companies 

Companies has many files 

How to translate it to createOne and getOne ?

Should I create collection company, user, and file for those ? 

Thank you
## [9][Testing and Validating Input in Go](https://www.reddit.com/r/golang/comments/h0z9zg/testing_and_validating_input_in_go/)
- url: https://hackwild.com/article/go-input-validation-and-testing/
---

## [10][Avoiding nested function (using pixel)](https://www.reddit.com/r/golang/comments/h0z7x5/avoiding_nested_function_using_pixel/)
- url: https://www.reddit.com/r/golang/comments/h0z7x5/avoiding_nested_function_using_pixel/
---
Hi I'm using the pixel library (in which basically you define a run function and it's your "real" main fuction the main jsute does : pixel.Run(run) and I feel that it's not a good idea to define any pixel object outside of the run (thats probabely on this point that I'm wrong...) ): 
I'm really new in programming and with pixel I've created a "shoot" function which lets my character shoot, but... I've first created the function inside of the run func... before discovering that go does not let us create nested func... 
So I defined it outside the run function but... I need some variable (such as my main window or things like that ) which are defined inside the run function (so on a local scope... so basically I can't access them)... could you please tell me a way to go through this or tell me if I'm not thinking my code the right way... I've never coded a game before (and almost never coded) 
Thanks you ! &lt;3
