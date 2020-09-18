# golang
## [1][golang-dev : Implementing Generics](https://www.reddit.com/r/golang/comments/iv4idd/golangdev_implementing_generics/)
- url: https://groups.google.com/forum/#!topic/golang-dev/OcW0ATRS4oM
---

## [2][Decoding JSON, the... a way](https://www.reddit.com/r/golang/comments/iv2zpd/decoding_json_the_a_way/)
- url: https://www.reddit.com/r/golang/comments/iv2zpd/decoding_json_the_a_way/
---
 [https://twitter.com/caarlos0/status/1306447785831673867?s=19](https://twitter.com/caarlos0/status/1306447785831673867?s=19)
## [3][What is the most efficient algorithm for an HTTP router?](https://www.reddit.com/r/golang/comments/iuvhxd/what_is_the_most_efficient_algorithm_for_an_http/)
- url: https://www.reddit.com/r/golang/comments/iuvhxd/what_is_the_most_efficient_algorithm_for_an_http/
---
Yes I know, there are TONS os really good HTTP routers for Go. But this is for a custom project and I would like to implement the router as an exercise.

Isn't really an HTTP framework or something like that. It's a sort of hooks to be executed on given URLs and that's why a need a router to do that.

Memory usage is not a problem because the number of endpoints is limited, usually less than 100. I would like to have the most performance possible at the router.
## [4][Opting for a distributed Pub/Sub system VS in-process comm. Pub/Sub VS neither](https://www.reddit.com/r/golang/comments/iv2zgc/opting_for_a_distributed_pubsub_system_vs/)
- url: https://www.reddit.com/r/golang/comments/iv2zgc/opting_for_a_distributed_pubsub_system_vs/
---
Hello,I am working on bringing up a server from scratch, for a tiny startup that will have a decent, yet not a crazy-massive scale in the beginning. I have doubts about a certain part of the architecture, and would VERY much appreciate some helpful insights and pieces of advice.

I have two services : one for mobile Pushes (FCM), and the other for Emails (SendGrid at the moment).

When I got to implementing it and using the services (for example upon user registration sending a welcome email), it seemed to me that the wisest thing in terms of best practices, fault-tolerance, scalability and performance would be to have a Pub/Sub system, like RabbitMQ.

So instead of :

    // registration endpoint's controller : 
     user := service.SignUp(...)  
     email_service.SendEmail(user.id)  

push\_service.SendPush(user.id) ...

I would have something like this I guess :

    // registration endpoint's controller : 
     user := service.SignUp(...) 
     publisher.Publish(Type.EMAIL, user.id) 
     publisher.Publish(Type.PUSH, user.id) ... 

Then I stumbled upon some technical dilemmas, that first started from things like "are Email and Push different exchanges? or simply two different Queues with different Topics, to one Exchange ?",and then proceeded some other complexities, like how to manage the different errors and reconnections that can (and probably will) happen unexpectedly sometime.

Which led me to the following question - why not simply leverage Go's capabilities and implement an in-process solution as depicted here : [https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/](https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/)  using Goroutines &amp; Channels.

Drawbacks :

I lose the fault-tolerance, and the second one is perhaps the problematic scalability if we have to add more instances in the future as scale grows bigger (but I'm not sure about that?)

Advantages :

Much simpler to maintain and less error-prone (because it doesn't have all the edge cases that a real pub/sub system like rabbit can have, like unexpected different errors, etc.)

**However** :

1. Since at the moment we will have one server, is there really a point for a pub/sub system like rabbit? or is it better to prepare for large scale from the start ?
2. Since the tasks are not really time consuming (the Consumers/Subscribers would only need to query the DB to get some users' details (e.g - user's fcm token)), is there really a point in having a pub/sub AT ALL?

**Ultimately**, my main goal here is to learn and understand - what are the questions that have to be answered when making a choice between those 3 options :

1. distributed pub/sub system (be it RabbitMQ, Google Pub/Sub, Kafka, etc.)
2. in-process pub/sub system - like here: [https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/](https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/)
3. simply running the code in the endpoint's controller, be it in a blocking way, or simply spawning a goroutine for it "on-the-fly"
## [5][Awesome Go Repos](https://www.reddit.com/r/golang/comments/iv5mbd/awesome_go_repos/)
- url: https://aggregatedawesome.com/go
---

## [6][What is The Best GraphQL Library for Go?](https://www.reddit.com/r/golang/comments/iv1xx1/what_is_the_best_graphql_library_for_go/)
- url: https://www.reddit.com/r/golang/comments/iv1xx1/what_is_the_best_graphql_library_for_go/
---
What i want:

* Type safe
* Simple API
* Easy to convert endpoint to serve NonHTTP/WebSocket request eg: AMQP, Kafka etc
## [7][Experimental tool for major version updates](https://www.reddit.com/r/golang/comments/iuovci/experimental_tool_for_major_version_updates/)
- url: https://github.com/icholy/gomajor
---

## [8][Datatable : opensource go in-memory table](https://www.reddit.com/r/golang/comments/iufaui/datatable_opensource_go_inmemory_table/)
- url: https://www.reddit.com/r/golang/comments/iufaui/datatable_opensource_go_inmemory_table/
---
Hi all,

Datatable is a Go package to manipulate tabular data, like an excel spreadsheet.

Select, sort, join, agregate, import from csv...See features and howto on github.

[https://github.com/datasweet/datatable](https://github.com/datasweet/datatable)

Please Try it and give us feedback.
## [9][How we built a Lucene-inspired parser in Go](https://www.reddit.com/r/golang/comments/iv1q42/how_we_built_a_luceneinspired_parser_in_go/)
- url: https://www.mailgun.com/blog/how-we-built-a-lucene-inspired-parser-in-go
---

## [10][Is there a frontend you recommend for using with Go?](https://www.reddit.com/r/golang/comments/iuygz1/is_there_a_frontend_you_recommend_for_using_with/)
- url: https://www.reddit.com/r/golang/comments/iuygz1/is_there_a_frontend_you_recommend_for_using_with/
---
Wondering if anything in particular sticks out for being efficient or very friendly with Go?

I have experience with React, and Django. Was wondering if I should learn something else for the frontend or if I can use React or Django and there's not a performance problem?
