# golang
## [1][10 Top Places to Learn Go Programming](https://www.reddit.com/r/golang/comments/evm28e/10_top_places_to_learn_go_programming/)
- url: https://www.agiratech.com/learn-go-programming-tutorials/
---

## [2][Proposals for Go 1.15 - The Go Blog](https://www.reddit.com/r/golang/comments/eva9qo/proposals_for_go_115_the_go_blog/)
- url: https://blog.golang.org/go1.15-proposals
---

## [3][reading a large number of csv files, with large size, then sorting it Golang](https://www.reddit.com/r/golang/comments/evl398/reading_a_large_number_of_csv_files_with_large/)
- url: https://www.reddit.com/r/golang/comments/evl398/reading_a_large_number_of_csv_files_with_large/
---
I have a large number of csv files, the first column is an id, second is the price, how I can sort all of this by price? I can read all file, then I can sort one file, after process all files I can use merge sort for merge all data. but can I solve it by streaming read(line by line)? How to do the right thing?
## [4][How to differentiate dev and prod environment?](https://www.reddit.com/r/golang/comments/evlj8n/how_to_differentiate_dev_and_prod_environment/)
- url: https://www.reddit.com/r/golang/comments/evlj8n/how_to_differentiate_dev_and_prod_environment/
---
I am writing integration test for a microservice and while running tests I want my application to use a different database and not the production one. I am thinking of having an env variable APP\_ENV to differentiate between prod and dev environment.  What are the other possible ways I can achieve this? I would really appreciate to know how others are doing it, their use cases and anything that would be useful. Thank you.
## [5][Tutorial for mongo go driver with gin gonic](https://www.reddit.com/r/golang/comments/evlfx8/tutorial_for_mongo_go_driver_with_gin_gonic/)
- url: https://www.reddit.com/r/golang/comments/evlfx8/tutorial_for_mongo_go_driver_with_gin_gonic/
---
Gin claims to be one of the fastest web frameworks out there. I have been struggling to find a good tutorial which make a rest API using gin gonic and MongoDB, but using the official MongoDB go driver which has proper structuring of the application. Thanks for the help.
## [6][ThreatBite: IP / e-mail reputation checking tool](https://www.reddit.com/r/golang/comments/evk8j7/threatbite_ip_email_reputation_checking_tool/)
- url: https://github.com/optimatiq/threatbite
---

## [7][Functional options on steroids](https://www.reddit.com/r/golang/comments/ev863i/functional_options_on_steroids/)
- url: https://sagikazarmark.hu/blog/functional-options-on-steroids/
---

## [8][[security] Go 1.13.7 and Go 1.12.16 are released](https://www.reddit.com/r/golang/comments/ev893f/security_go_1137_and_go_11216_are_released/)
- url: https://groups.google.com/forum/m/#!topic/golang-announce/Hsw4mHYc470
---

## [9][A GraphQL-Based Web App with Go, React and MongoDB](https://www.reddit.com/r/golang/comments/ev9z0m/a_graphqlbased_web_app_with_go_react_and_mongodb/)
- url: https://github.com/Shpota/skmz
---

## [10][External API protocol choice](https://www.reddit.com/r/golang/comments/evl4k2/external_api_protocol_choice/)
- url: https://www.reddit.com/r/golang/comments/evl4k2/external_api_protocol_choice/
---
Hi all, sorry not entirely Golang-specific but the work will be done in Go, so library support matters.

We're creating an external API. The client will be IoT platforms - servers that are connecting to our API to send &amp; receive events:

- changes to devices sent from device server to the API
- actions from the API backend sent to device server)

I'd rather avoid the overhead associated with HTTP, especially with the clients sending device updates. A webhooks approach is an option however they add some complexity on the client, requiring an HTTP server their for us to make calls against, and to setup the webhooks.

Any specific, Go-friendly solutions you would steer us towards? Other options we are considering: MQTT, gRPC.

gRPC seems alright but I'm not a fan of protobufs too much for an external API, since the client would need the .proto files and to use protobuf, which is not the nicest workflow. gRPC with a JSON codec instead of Protobuf codec is another option that might be a good solution, and it is also curl-friendly.
