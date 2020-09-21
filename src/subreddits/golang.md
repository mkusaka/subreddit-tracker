# golang
## [1][A load testing tool with a real-time analyzer, written in Go](https://www.reddit.com/r/golang/comments/iwdx5v/a_load_testing_tool_with_a_realtime_analyzer/)
- url: https://i.redd.it/bf5f5nb42bo51.gif
---

## [2][k6 v0.28.0 is out!](https://www.reddit.com/r/golang/comments/iwzy8u/k6_v0280_is_out/)
- url: https://www.reddit.com/r/golang/comments/iwzy8u/k6_v0280_is_out/
---
We are happy to announce that k6 v0.28.0 is released with new features and improvements. k6 is a modern open-source performance and load-testing tool, written in Go and scriptable in JavaScript. 🎊🎉🥳

- k6 Cloud execution logs
- Pushing logs to Loki
- Optional port to host mappings
- Support for specifying data types to InfluxDB fields
- Support for automatic gzip-ing of the CSV output result
- UX improvements
- Lots of bugfixes

You can read more about it in the [release notes](https://github.com/loadimpact/k6/releases/tag/v0.28.0). We'd be happy to hear your feedback about it.
## [3][Code Help: Go idomatic way of stopping a go routine](https://www.reddit.com/r/golang/comments/iwt1ih/code_help_go_idomatic_way_of_stopping_a_go_routine/)
- url: https://www.reddit.com/r/golang/comments/iwt1ih/code_help_go_idomatic_way_of_stopping_a_go_routine/
---
I've read no less than 5 different articles this evening about using Contexts in Go, as I think that's what I need to solve my issue, but I'd appreciate some review of what I'm trying to do.

Consider the pseudocode below. The app is a pretty simple, producer-consumer. The main has a go func (A) that just does a fetch-compare-sleep-loop: Fetch data from a source, compare to previous fetch, if different data put on channel, sleep, repeat. If same data as before, sleep, repeat. (ie: polling)

The other part of main (B) simply waits for new data on the channel and processes it using myFunc(). This is not a go func, so it waits for the func to finish before starting the wait-loop over.

Let's say that while myFunc is processing data, new data comes in to the channel. I need to cancel/abort the current processing run and start myFunc over with the new data.

I believe this is where a Context comes in use. Here's what I'm thinking: Create a Context.WithCancel() at (B), change myFunc to be 'go myFunc', passing the ctx and data to myFunc, and hold the .Cancel() variable inside B's for-loop. If I get new data on the channel, call .Cancel, create a new ctx.WithCancel and call myFunc again.

If that's sounds good (baring any minor tweaks), inside myFunc, do I just have periodic checks to determine if the ctx has been canceled by the for-select (B)? Worded differently, how do I determine inside myFunc() if it's been canceled?

    func main() {
     myChan := make(chan string, 1)
     // A
     go func(c chan&lt;- string) {
      for {
        // fetch data
        // compare previous; if same, continue/sleep/loop
        // if different, put on channel
      }
     }
    
     // B
     for {
      select &lt;-chan:
        myFunc(data)
     }
    }
    
    // C
    func myFunc(data) {
      // process data
    }
## [4][Appreciation post for spf13's Cobra and Viper packages, that was possibly the best experience I've had writing a CLI in a loooong time.](https://www.reddit.com/r/golang/comments/iwjmub/appreciation_post_for_spf13s_cobra_and_viper/)
- url: https://www.reddit.com/r/golang/comments/iwjmub/appreciation_post_for_spf13s_cobra_and_viper/
---
Took some time over the weekend to write a command line client for privnote.com and decided to use cobra/viper for configuration management. Most people know about it I'm sure but it pays to show your appreciation to open source projects sometimes.
## [5][Tool for managing parameters in AWS SSM Parameter Store - Written in Go!](https://www.reddit.com/r/golang/comments/iwyp03/tool_for_managing_parameters_in_aws_ssm_parameter/)
- url: https://www.reddit.com/r/golang/comments/iwyp03/tool_for_managing_parameters_in_aws_ssm_parameter/
---
[https://github.com/kevinglasson/goss](https://github.com/kevinglasson/goss)

I developed a tool to help managing config for various environments in AWS SSM Parameter Store. I thought some people here might be interested as there seems to be a few golang tools in the devops / cloud space.

I've used a few 'major' go packages i.e. cobra to create the tool. Looking for any general coding / go advice and any other thought about the tool / it's usefulness and ideas to extend it!

Check out the repo for a little demo and some explanation of the commands!
## [6][This is a JavaScript bundler and minifier. It packages up JavaScript and TypeScript code for distribution on the web.](https://www.reddit.com/r/golang/comments/iw9kfv/this_is_a_javascript_bundler_and_minifier_it/)
- url: https://github.com/evanw/esbuild
---

## [7][Is NATS on par/better than say MQTT over RabbitMQ?](https://www.reddit.com/r/golang/comments/iwjjzk/is_nats_on_parbetter_than_say_mqtt_over_rabbitmq/)
- url: https://www.reddit.com/r/golang/comments/iwjjzk/is_nats_on_parbetter_than_say_mqtt_over_rabbitmq/
---
As the topic asks.. I just heard of NATS.. and little pissed I hadn't heard of it before. I guess I didn't do enough due diligence to discover what this was. I have a microservices setup using RabbitMQ, and using MQTT for the messages. It works very well.. very fast.. but just read about NATS and was starting to wonder.. if that may be a better way to go. Largely my "pub/sub" could be swapped easily enough.. I didn't quite build it generic enough but it wouldn't take much work as I use a common send() and pub/sub functions in all my services. I am using the Eclipse Pahao library.

Has anyone experience with both and found that NATS is easier/faster/better for basic simple messaging between services than using MQTT and a typical message bus like RabbitMQ? 

I'll be reading a bit more into it now, but thought I would ask the community what their experience/thoughts are on the two approaches.

Thank you.
## [8][A security-focused MUX?](https://www.reddit.com/r/golang/comments/iwr0wc/a_securityfocused_mux/)
- url: https://www.reddit.com/r/golang/comments/iwr0wc/a_securityfocused_mux/
---
Howdy Y'all,

I've made a mux, then realized what I want might already exist.
In NodeJS, there's express.  You can do stuff like this:
`    app.get('/myroute/', (req, res)=&gt;{...});`

I like this style of writing, where I can just think in terms of the route and the handler, but I wanted to be able to use it in more secure contexts.

So, I kind of made an explicit route handling mux like so:


    myRoute := webserver.ServerMux{
        Options: webserver.MuxOptions{RequireAuthorization: false, AllowedFailuresPerIP: 3},
        Method:  "PUT",
        Route:   "/v1/Login/",
        RouteHandler: func(res http.ResponseWriter, req *http.Request) bool {...return SUCCESS}}


That route gives an unauthenticated user 3 tries to login, or it'll begin an escalating block on them.
You then put these routes into a slice, you then can serve the slice of routes with options on the entire slice ( for example, which user groups are allowed ).  

It handles a lot of things auto-magically:

1. Routes can return a boolean indicating if they "worked".  Success and failures are recorded to a postgres db, and it supports SIEMs by emitting syslog style events over TCP.  

2. It inherently supports both Authorization ( aka Authorization headers --&gt; hash verification ) and RBAC ( aka, you can mark certain routes as allowed by certain users )

Before I keep developing this thing, does it already exist?  I already checked Gorilla, and it has far fewer security features ( though it's a great MUX. )

If it doesn't exist, would anyone join me as a maintainer ( I'm thinking of open sourcing it -- but only if there's at least 1 other active collaborator who will actually work on it, too.  )

Thanks!
## [9][Typesafe database access for Go](https://www.reddit.com/r/golang/comments/iw02hp/typesafe_database_access_for_go/)
- url: https://github.com/prisma/prisma-client-go
---

## [10][YoMo is an open-source project for building your own IoT edge computing applications. With YoMo, you can speed up the development of microservices-based applications, and your industrial IoT platform will take full advantage of the low latency and high bandwidth brought by 5G.](https://www.reddit.com/r/golang/comments/iw1b3z/yomo_is_an_opensource_project_for_building_your/)
- url: https://github.com/yomorun/yomo
---

