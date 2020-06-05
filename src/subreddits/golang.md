# golang
## [1][why does json.unmarshall change the value of uint64 to float64 ?](https://www.reddit.com/r/golang/comments/gx1xdv/why_does_jsonunmarshall_change_the_value_of/)
- url: https://www.reddit.com/r/golang/comments/gx1xdv/why_does_jsonunmarshall_change_the_value_of/
---
i have a json file that contains some json data. The values of the keys are in UINT64,when i decode this in my map using

    var objmap map[string]interface{}
    err = json.Unmarshal(jsondata, &amp;objmap)

the values of the key come out as float64

example my json would be 

    {
       "someval": {
          "first": 1228224835,
       }
    }

but when i unmarshall i get  first :1.228224835e+09

Edit : https://stackoverflow.com/questions/55436628/json-decoded-value-is-treated-as-float64-instead-of-int/55436758 seems to be the answer

Edit 2 : with the solution, so basically what i did was 

    data, err := os.Open("./file.json")
    check(err)
    d := json.NewDecoder(data)
    d.UseNumber()
    err = d.Decode(&amp;objmap)

this seems to work https://stackoverflow.com/questions/22343083/json-unmarshaling-with-long-numbers-gives-floating-point-number
## [2][Handling multidomain requests with simple host switch](https://www.reddit.com/r/golang/comments/gx0t6e/handling_multidomain_requests_with_simple_host/)
- url: https://rafallorenz.com/go/go-multidomain-host-switch/
---

## [3][The Go compiler needs to be smarter](https://www.reddit.com/r/golang/comments/gwlrms/the_go_compiler_needs_to_be_smarter/)
- url: https://lemire.me/blog/2020/06/04/the-go-compiler-needs-to-be-smarter/
---

## [4][Fictional 80s computer](https://www.reddit.com/r/golang/comments/gwpl6m/fictional_80s_computer/)
- url: https://www.reddit.com/r/golang/comments/gwpl6m/fictional_80s_computer/
---
My project [benji4000](https://github.com/uzudil/benji4000) is a fictional 80s gaming computer. It's programmed in an interpreted, js-like [language](https://github.com/uzudil/benji4000/wiki/LanguageFeatures). The project is just starting out and comes with a couple of games. Let me know what you think!

https://preview.redd.it/5a9l4o9jfy251.png?width=1504&amp;format=png&amp;auto=webp&amp;s=87fdf00fd60ec16ac4c459ea60e18285ee9f3f52
## [5][Junior golang developer interview](https://www.reddit.com/r/golang/comments/gwxctz/junior_golang_developer_interview/)
- url: https://www.reddit.com/r/golang/comments/gwxctz/junior_golang_developer_interview/
---
I'm fairly new to golang, coming from Java and c++ mainly. I've been learning golang for a couple of months.
 I have to crack an junior dev interview that's in about a week.
How would you prepare for a junior golang dev position? I'm looking for commonly asked interview questions and other interview protocols or tips. Though, anything would be helpful.
The answers to this post will surely be helpful to other devs too.
## [6][How I solved Jepsen with OpenCensus Distributed Tracing: A personal journey](https://www.reddit.com/r/golang/comments/gwgaf6/how_i_solved_jepsen_with_opencensus_distributed/)
- url: https://dgraph.io/blog/post/solving-jepsen-with-opencensus/
---

## [7][gql: A new way of doing GraphQL in Go](https://www.reddit.com/r/golang/comments/gwj0b4/gql_a_new_way_of_doing_graphql_in_go/)
- url: https://www.reddit.com/r/golang/comments/gwj0b4/gql_a_new_way_of_doing_graphql_in_go/
---
In the last 2 years, I learned more and more on GraphQL in Go because of my work and enthusiasm in programming, so I started contributing to graphql-go, worked on extensions (middleware is a better word for it) for schemas, so a middleware for tracing could be added to the schema. But I don’t like that there’s way too much boilerplate code required to get started and in the last few months, I couldn’t fail to notice that there’s less and less activity in the repository, which results in no new features.

Last year, as a project to learn more about parsers, I built one for GraphQL. It turned out great, there are so many things in that topic and can be interesting (might write a post on that topic later). I found it useful, so I can only recommend writing a parser for a fictional language, configuration files, or for templating, you can learn a ton.

Since I already had the core implementation of the GraphQL language, I decided to do more. I started working on a GraphQL server, implementing the type system, validation, and execution.

Early this year, I decided that I want to have something, I want to build an open-source project, one that I care about, I start and use. After a little time, I had to realize that there’s no better project than a GraphQL package for Go, targeting features missing from existing packages, or ones that could be done differently. So I started coding, revisited my parser, my existing implementation, and did a huge refactor to make it more efficient…

The [github.com/rigglo/gql](https://github.com/rigglo/gql) package is still in a work-in-progress solution, but already supports the most common features, like

* Custom scalars
* Built-in playground
* Subscriptions: using the new [github.com/rigglo/gqlws](https://github.com/rigglo/gqlws) package, you can create one endpoint to support regular request and others with WebSockets for subscriptions
* Field directives: can be used to check if the requester has permissions for a given field, or to solve other cases
* Extensions: tools can be developed for tracing or to resolve custom issues
* Validation: validates all the queries before execution as the specification describes
* Concurrent execution

And many more are coming…

* Checks for query complexity
* Apollo File Upload
* Apollo Federation support
* Custom validation for inputs
* Custom rules-based introspection
* Converting structs into GraphQL types
* Parse inputs into structs

In the last few weeks, I started implementing it into one of the projects I’m working on, so there’s continuous feedback coming from there, but it’s not enough. If you’re planning on doing GraphQL in Go, about to start a project, or just curious about this, I’d like you to try, test the [gql](https://github.com/rigglo/gql) package, contribute, or provide feedback for me on how I can improve it, make it better. I have so many ideas for this project, solutions I’d like to bring to life, so with your help, with your feedback and contribution, we could shape it to be an awesome GraphQL package for Go, a solution for many issues, with not much boilerplate, and great learning curve.

I'm happy to answer any questions!

&amp;#x200B;

Original post: [https://medium.com/rigglo/a-new-way-of-doing-graphql-in-go-e909b8eb83f5](https://medium.com/rigglo/a-new-way-of-doing-graphql-in-go-e909b8eb83f5)

Repository: [https://github.com/rigglo/gql](https://github.com/rigglo/gql)

Me: [https://github.com/Fontinalis](https://github.com/Fontinalis)
## [8][Best Go email server for Web App sending emails?](https://www.reddit.com/r/golang/comments/gx0eru/best_go_email_server_for_web_app_sending_emails/)
- url: https://www.reddit.com/r/golang/comments/gx0eru/best_go_email_server_for_web_app_sending_emails/
---
Is there a small Go project suitable for sending verification emails for a small web app?
## [9][Call for Speakers - GopherCon Turkey (25th July)](https://www.reddit.com/r/golang/comments/gwpflq/call_for_speakers_gophercon_turkey_25th_july/)
- url: https://www.reddit.com/r/golang/comments/gwpflq/call_for_speakers_gophercon_turkey_25th_july/
---
We're proud to announce that we are organizing Turkey's first Go programming language conference. GopherCon Turkey is planned to bring together the Go community and share experiences. The conference will be online and will host 17 speakers in 2 parallel sessions, one in Turkish and the other in English on the 25th of July. Registration is free.

Call for Papers will be ended on June, 22nd.

Details: [http://gophercon.ist/](http://gophercon.ist/)

https://preview.redd.it/o3s0ai07dy251.png?width=1200&amp;format=png&amp;auto=webp&amp;s=29233538069a3c512981550b49c88ef4eebfcd12
## [10][[help] [gotk3] Need help in removing signal/function from gtk button](https://www.reddit.com/r/golang/comments/gwz9hl/help_gotk3_need_help_in_removing_signalfunction/)
- url: https://www.reddit.com/r/golang/comments/gwz9hl/help_gotk3_need_help_in_removing_signalfunction/
---
I need help in removing the signal/function or whatever i connect to the button

in python it can be done by
button.disconnect_by_func(func)

how to done the same in go
