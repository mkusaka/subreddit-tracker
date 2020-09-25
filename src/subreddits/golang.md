# golang
## [1][I just finished and released v1.0 of my programming language, written entirely in Go! I have no idea how to properly write a language but I gave it my best shot](https://www.reddit.com/r/golang/comments/izgz5w/i_just_finished_and_released_v10_of_my/)
- url: https://github.com/odddollar/Leafscript
---

## [2][Run protoc with docker](https://www.reddit.com/r/golang/comments/izbx4b/run_protoc_with_docker/)
- url: https://www.reddit.com/r/golang/comments/izbx4b/run_protoc_with_docker/
---
I just want to share my personal project for runing protoc with docker.

[https://github.com/haunt98/docker-protoc](https://github.com/haunt98/docker-protoc)

The problem when I generate protobuf and grpc in golang is grpc version seems to be different between OS. And man, version conflict is so hard. So I wrap everything should need to run protoc inside docker.
## [3][Create Secure Clients and Servers in Go using HTTPS](https://www.reddit.com/r/golang/comments/iz5otz/create_secure_clients_and_servers_in_go_using/)
- url: https://www.reddit.com/r/golang/comments/iz5otz/create_secure_clients_and_servers_in_go_using/
---
[https://youngkin.github.io/post/gohttpsclientserver/](https://youngkin.github.io/post/gohttpsclientserver/)

Spent some time trying to add HTTPS to a service and turned my notes into a blog post. Feedback is welcome
## [4][gomponents: declarative view components in Go](https://www.reddit.com/r/golang/comments/izh5dm/gomponents_declarative_view_components_in_go/)
- url: https://www.maragu.dk/blog/gomponents-declarative-view-components-in-go/
---

## [5][gin-gonic/gin metrics exporter for Prometheus.](https://www.reddit.com/r/golang/comments/izfdnt/gingonicgin_metrics_exporter_for_prometheus/)
- url: https://www.reddit.com/r/golang/comments/izfdnt/gingonicgin_metrics_exporter_for_prometheus/
---
Hello everyone, i just want to share my personal project, that add metrics for Prometheus when run a gin http-server

Repo address: [https://github.com/penglongli/gin-metrics](https://github.com/penglongli/gin-metrics)

# Intoduction

gin-metrics defines some metrics for gin http-server. There have easy way to use it.

Below is the detailed description for every metric.

|*Metric*|*Type*|*Description*|
|:-|:-|:-|
|gin\_request\_total|Counter|all the server received request num.|
|gin\_request\_uv|Counter|all the server received ip num.|
|gin\_uri\_request\_total|Counter|all the server received request num with every uri.|
|gin\_request\_body\_total|Counter|the server received request body size, unit byte.|
|gin\_response\_body\_total|Counter|the server send response body size, unit byte.|
|gin\_request\_duration|Histogram|the time server took to handle the request.|
|gin\_slow\_request\_total|Counter|the server handled slow requests counter, t=%d.|

# Usage

Your can see some metrics across [*http://localhost:8080/metrics*](http://localhost:8080/metrics)

    package main
    
    import (
    	"github.com/gin-gonic/gin"
    	
    	"github.com/penglongli/gin-metrics/ginmetrics"
    )
    
    func main() {
    	r := gin.Default()
    
    	// get global Monitor object
    	m := ginmetrics.GetMonitor()
    
    	// +optional set metric path, default /debug/metrics
    	m.SetMetricPath("/metrics")
    	// +optional set slow time, default 5s
    	m.SetSlowTime(10)
    	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
    	// used to p95, p99
    	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
    
    	// set middleware for gin
    	m.Use(r)
    
    	r.GET("/product/:id", func(ctx *gin.Context) {
    			"productId": ctx.Param("id"),
    		})
    	})
    	_ = r.Run()
    }

# Custom Metric

You can use it to custom your own metric, contain ***Gauge/Counter/Histogram/Summary*** type.

If you met some problems, you can new an ISSUE for me.
## [6][Why does json.Marshal return error type?](https://www.reddit.com/r/golang/comments/izjbej/why_does_jsonmarshal_return_error_type/)
- url: https://www.reddit.com/r/golang/comments/izjbej/why_does_jsonmarshal_return_error_type/
---
json.Marshal is used in a lot of places, including in error handling for HTTP endpoints. Any extraneous error values here lead to wasteful boilerplate.

Is there a specific reason that json.Marshal features an error return type?

I get why json.Unmarshal could fail. An error type makes sense there.

But say a Go entity is logically unable to JSON encode (it has circular references). Would the encoder truly expend the effort to detect such cases? What kinds of real world errors does json.Marshal return? Some native Go types like chan that are unencodable? Out of memory errors?

Note that many unencodable situations could theoretically be detected at compile time, rather than lingering for runtime.
## [7][[help wanted] - MongoDB Filter](https://www.reddit.com/r/golang/comments/izg5su/help_wanted_mongodb_filter/)
- url: https://www.reddit.com/r/golang/comments/izg5su/help_wanted_mongodb_filter/
---
I have this starting point: [https://mongoplayground.net/p/Wue-aotf4cw](https://mongoplayground.net/p/Wue-aotf4cw) and I am trying to implement the specified filter with Go.: [https://play.golang.org/p/tco0rOADi7q](https://play.golang.org/p/tco0rOADi7q) ...but unfortunately it does not work...  
What am I doing wrong?
## [8][Live programming a Terraform Provider #3 - Writing a better API client (starts at 5pm est)](https://www.reddit.com/r/golang/comments/iz5nst/live_programming_a_terraform_provider_3_writing_a/)
- url: https://www.twitch.tv/bobbytables
---

## [9][TamaGo – bare metal Go for ARM SoCs](https://www.reddit.com/r/golang/comments/iyv2zd/tamago_bare_metal_go_for_arm_socs/)
- url: https://github.com/f-secure-foundry/tamago
---

## [10][Looks more like a Gopher to me](https://www.reddit.com/r/golang/comments/iz9ikc/looks_more_like_a_gopher_to_me/)
- url: https://i.imgur.com/ZDx5XoV.jpg
---

