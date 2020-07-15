# golang
## [1][Go generics will use square brackets [] not parenthesis ()](https://www.reddit.com/r/golang/comments/hrce8e/go_generics_will_use_square_brackets_not/)
- url: https://groups.google.com/forum/#!topic/golang-nuts/7t-Q2vt60J8
---

## [2][performance cost of dockerizing a Go app](https://www.reddit.com/r/golang/comments/hr5yzx/performance_cost_of_dockerizing_a_go_app/)
- url: https://www.reddit.com/r/golang/comments/hr5yzx/performance_cost_of_dockerizing_a_go_app/
---
In [this](https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9) Go+Docker tutorial, one the the [comments](https://medium.com/p/a27141f36ba9/responses/show) says:

&gt;Docker is great for development, but the performance gains you get from Golang, you lose with Docker, you may as well go php or python, since the whole point of golang is a single distributable binary. build in Docker, deploy to bare metal is my advice.

I can't help but to think that this is inaccurate. It seems like the runtime performance cost of dockerizing would be independent of what gets run inside it. You would still incur that cost regardless of the app's language.

Furthermore, the performance cost is practically negligible, according to [this SO answer](https://stackoverflow.com/questions/21889053/what-is-the-runtime-performance-cost-of-a-docker-container#answer-26149994):

&gt;An excellent 2014 IBM research paper “An Updated Performance Comparison of Virtual Machines and Linux Containers” by Felter et al. provides a comparison between bare metal, KVM, and Docker containers. The general result is: **Docker is nearly identical to native performance** and faster than KVM in every category.

Lastly, containerization is so common practice, and I've never heard/read about not containerizing due to performance concerns.

As a fairly inexperienced backend dev, I was hoping to get a few thoughts on this from more experienced devs. Thanks
## [3][Go 1.14.5 is just released](https://www.reddit.com/r/golang/comments/hr82o2/go_1145_is_just_released/)
- url: https://www.reddit.com/r/golang/comments/hr82o2/go_1145_is_just_released/
---
[https://golang.org/dl/](https://golang.org/dl/)

[https://golang.org/doc/go1.14](https://golang.org/doc/go1.14)

[https://golang.org/doc/devel/release.html#go1.14](https://golang.org/doc/devel/release.html#go1.14)
## [4][Build a REST API with Go, PostgreSQL, and Test-Driven Development](https://www.reddit.com/r/golang/comments/hr7euq/build_a_rest_api_with_go_postgresql_and/)
- url: https://medium.com/@juancurti.it/go-tutorial-tdd-with-go-and-postgresql-part-ii-489c929f02c9
---

## [5][Golang Templates](https://www.reddit.com/r/golang/comments/hrl187/golang_templates/)
- url: https://www.reddit.com/r/golang/comments/hrl187/golang_templates/
---
 

    &lt;div class="horzontialList" id="List"&gt;{{range .}} {{if ne .Symbol " "}} &lt;span&gt;{{.Symbol}}&lt;/span&gt; {{end}} {{end}}
    var div=document.getElementById('List');
var point = div.getElementsByTagName('span');

console.log(point[0].innerHTML)

Actual Output: Empty stringExcepted Output: Y How to resolve this 
## [6][Virtual Meetup on Generics Proposal Tomorrow](https://www.reddit.com/r/golang/comments/hr955o/virtual_meetup_on_generics_proposal_tomorrow/)
- url: https://www.reddit.com/r/golang/comments/hr955o/virtual_meetup_on_generics_proposal_tomorrow/
---
Pacific Northwest Go is hosting Bill Kennedy tomorrow night at 6pm US Pacific Time (UTC-7). He will review the new generics proposal. All are welcome to join and participate and we'd love to see you there.

Register here:  [https://www.crowdcast.io/e/virtual-pnw-go-meetup--/register](https://www.crowdcast.io/e/virtual-pnw-go-meetup--/register)
## [7][Looking for an idea to make a solution to a problem](https://www.reddit.com/r/golang/comments/hremvv/looking_for_an_idea_to_make_a_solution_to_a/)
- url: https://www.reddit.com/r/golang/comments/hremvv/looking_for_an_idea_to_make_a_solution_to_a/
---
I'm not very good at Go but I am faced with a challenge I would like to work on and get better at. I'm hoping someone can provide some knowledge on how to go about my issue.

Scenario :

I have a Redis DB that I get data from every hour.  There can be different number of data, for example 00:00 may give 5 sets of data while 00:01 may give 80. The data comes in the form of syntax :

`0x03948f44fee48548cf5b11aa26d76acf8430ff7c232095ce4975791ac222fa25da:6701:0x4fee48548cf5b11:0x0000000000000ce8`

It's the same syntax of data every time.

The data must be given to TCP connections but with criteria.

For example data can't be given to no more than 3 TCP Clients that are currently online. Meaning if one disconnections, the data can be given to a new or existing TCP client that doesn't already have data.

Something like if there three TCP Clients that has the same data, mark that data so that it can't be given out anymore times unless a TCP Client disconnects and the number of TCP Clients that have the data is no longer == 3.

TCP Clients use the data for their needs then sends True to the TCP server eventually. When TCP Clients that have the same data sends True, then all TCP Clients that have the data is given a new Data that doesn't have the true flag sent to it and also meets the criteria that data can't be given to anymore than 3 TCP clients.

That's it. 

Any ideas will be greatly appreciated.
## [8][Adrian - An image glitcher](https://www.reddit.com/r/golang/comments/hrd58e/adrian_an_image_glitcher/)
- url: https://github.com/manoloesparta/adrian
---

## [9][Hugo 0.74.0 released: Adds Native JS Bundler, Open API Support and Inline Partials](https://www.reddit.com/r/golang/comments/hqx78m/hugo_0740_released_adds_native_js_bundler_open/)
- url: https://gohugo.io/news/0.74.0-relnotes/
---

## [10][Help Wanted: Implement seek support in FLAC library](https://www.reddit.com/r/golang/comments/hrbyil/help_wanted_implement_seek_support_in_flac_library/)
- url: https://github.com/mewkiz/flac/issues/16
---

