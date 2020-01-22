# golang
## [1][Go's Tooling is an Undervalued Technology](https://www.reddit.com/r/golang/comments/es621w/gos_tooling_is_an_undervalued_technology/)
- url: https://nullprogram.com/blog/2020/01/21/
---

## [2][Hiring Go Engineers at Coder](https://www.reddit.com/r/golang/comments/erz90i/hiring_go_engineers_at_coder/)
- url: https://www.reddit.com/r/golang/comments/erz90i/hiring_go_engineers_at_coder/
---
I'm a senior engineer at https://coder.com.

We're a small startup in Austin, TX looking to scale our Go team with solid engineers. The positions involve developing and maintaining the Go microservices that serve our development platform on kubernetes. There are more details on precisely what the company does on the site.

We want someone by our downtown Austin office but are open to remote for the first few months until we can get you relocated. If you’re really good, full time remote is also possible. US &amp; Canada preferred.

We are big on open source. Check out some of our projects at https://github.com/cdr or mine at https://github.com/nhooyr.

We want solid engineers and will pay accordingly (90k - 200k USD).

The general hiring process looks like:

1. Technical interview directly with me
    - I won't quiz you on useless trivia
    - Most of my questions will be about general programming, Go, docker, k8s etc.
2. I'll give you a 2 hour take home for which you'll be compensated with a $200 amazon gift card.
3. You have a interview with our CTO about your goals and a touch of tech
4. You have a general call with a different senior engineer
5. One last call to discuss your compensation and then the offer gets formally sent out

You’ll get 3 weeks PTO standard (you can always take more unpaid) plus 10 company holidays yearly, solid work life balance (we don’t work more than 40 hours a week) and whatever desk setup you want. We have a very talented team full of Go lovers and I’m looking forward to seeing it grow.

Please email me at [anmol@coder.com](mailto:anmol@coder.com) with your resume.
## [3][Reduce debugging time while programming. Use static and stack-trace analysis to determine which func call causes the error.](https://www.reddit.com/r/golang/comments/esb2i4/reduce_debugging_time_while_programming_use/)
- url: https://github.com/snwfdhmp/errlog
---

## [4][Load balancing goroutines in Go](https://www.reddit.com/r/golang/comments/es9wuq/load_balancing_goroutines_in_go/)
- url: https://medium.com/@addityasingh/load-balancing-goroutines-in-go-57e0896c7f86?source=friends_link&amp;sk=6e5b32bb42717b2b0577a78753bfa285
---

## [5][golang systemd reload configuration instruction](https://www.reddit.com/r/golang/comments/es9lpn/golang_systemd_reload_configuration_instruction/)
- url: https://www.reddit.com/r/golang/comments/es9lpn/golang_systemd_reload_configuration_instruction/
---
Does anyone know how to catch a systemd "reload" instruction in a golang service? 

I'm looking at the [go-systemd libraries](https://github.com/coreos/go-systemd) but I don't see a method of responding to a D-bus "Reload" instruction?
## [6][gzip: An out-of-the-box, also customizable gzip middleware for Gin and net/http](https://www.reddit.com/r/golang/comments/es97qf/gzip_an_outofthebox_also_customizable_gzip/)
- url: https://www.reddit.com/r/golang/comments/es97qf/gzip_an_outofthebox_also_customizable_gzip/
---
https://github.com/nanmu42/gzip

# Look and Feel

## Gin

```go
import github.com/nanmu42/gzip

func main() {
	g := gin.Default()

        // use default settings
	g.Use(gzip.DefaultHandler().Gin)

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "hello",
			"data": fmt.Sprintf("l%sng!", strings.Repeat("o", 1000)),
		})
	})

	log.Println(g.Run(fmt.Sprintf(":%d", 3000)))
}
```

## net/http

```go
import github.com/nanmu42/gzip

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeString(w, fmt.Sprintf("This content is compressed: l%sng!", strings.Repeat("o", 1000)))
	})

    // wrap http.Handler using default settings
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", 3001), gzip.DefaultHandler().WrapHandler(mux)))
}

func writeString(w http.ResponseWriter, payload string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	_, _ = io.WriteString(w, payload+"\n")
}
```

# Customizable Handler

Handler can be customized during initialization:

```go
import github.com/nanmu42/gzip

handler := gzip.NewHandler(gzip.Config{
    // gzip compression level to use
	CompressionLevel: 6,
    // minimum content length to trigger gzip, the unit is in byte.
	MinContentLength: 256,
    // RequestFilter decide whether or not to compress response judging by request.
    // Filters are applied in the sequence here.
	RequestFilter: []RequestFilter{
	    NewCommonRequestFilter(),
	    DefaultExtensionFilter(),
	},
    // ResponseHeaderFilter decide whether or not to compress response
    // judging by request
	ResponseHeaderFilter: []ResponseHeaderFilter{
		NewSkipCompressedFilter(),
		DefaultContentTypeFilter(),
	},
})
```

`RequestFilter` and `ResponseHeaderFilter` are interfaces.
You may define one that specially suits your need.
## [7][Is there something like npmtrends.com for Go?](https://www.reddit.com/r/golang/comments/es91qy/is_there_something_like_npmtrendscom_for_go/)
- url: https://www.reddit.com/r/golang/comments/es91qy/is_there_something_like_npmtrendscom_for_go/
---
npmtrends is great to check traction, growth base of a node module. Is there anything similar in Go land?
## [8][cabify/gotoprom: Type-safe Prometheus metrics builder library for golang](https://www.reddit.com/r/golang/comments/ers1df/cabifygotoprom_typesafe_prometheus_metrics/)
- url: https://github.com/cabify/gotoprom
---

## [9][Interface in library instead of consumer](https://www.reddit.com/r/golang/comments/es615w/interface_in_library_instead_of_consumer/)
- url: https://www.reddit.com/r/golang/comments/es615w/interface_in_library_instead_of_consumer/
---
I was doing a code review for a package some of my fellow devs were making and found an interface that was defined in out that had every method from the package's main struct in it. I naturally commented that it's not being used by anything and so should be removed. They countered with it being for the consumers to easily mock and even included the mock in this package. We got into a debate and their argument was generally that it provides more by including a mock than it takes away - being less declarative in its usage elsewhere, and that generally for now all consumers will likely use all methods. So we went with leaning it in for now and monitoring changes.

I know that the style guide says interfaces go in consumers but I couldn't really find _why_ that's the case.

Anyone able to shed light on it?
## [10][Building a global services network using Go, QUIC and Micro](https://www.reddit.com/r/golang/comments/ertnaa/building_a_global_services_network_using_go_quic/)
- url: https://micro.mu/blog/2019/12/05/building-a-microservices-network.html
---

