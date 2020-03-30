# golang
## [1][Collection of Technical Interview Problems Solved In Go](https://www.reddit.com/r/golang/comments/frdm4u/collection_of_technical_interview_problems_solved/)
- url: https://github.com/shomali11/go-interview
---

## [2][Mokku - A clipboard-based mocking framework for Go that gets out of your way](https://www.reddit.com/r/golang/comments/frq068/mokku_a_clipboardbased_mocking_framework_for_go/)
- url: https://github.com/kinbiko/mokku
---

## [3][Retrieve Location of macOS Device from Go](https://www.reddit.com/r/golang/comments/frn6qt/retrieve_location_of_macos_device_from_go/)
- url: https://vladimir.varank.in/notes/2020/03/go-osx-core-location/
---

## [4][Molecule – Streaming, zero-allocation protobuf decoding in Go](https://www.reddit.com/r/golang/comments/fr6c8n/molecule_streaming_zeroallocation_protobuf/)
- url: https://github.com/richardartoul/molecule
---

## [5][How does one express tagged unions in Go?](https://www.reddit.com/r/golang/comments/frqg5h/how_does_one_express_tagged_unions_in_go/)
- url: https://www.reddit.com/r/golang/comments/frqg5h/how_does_one_express_tagged_unions_in_go/
---
I want to represent a state called slope, the state can be any of positive, negative or zero. As far as I came to understand, Go doesn't support unions [yet](https://github.com/golang/go/issues/19412). What is the popular pattern for emulating unions?

In Rust I would do something like 

```
enum Slope {
  positive, zero, negative
}
```
## [6][My program is slower when it comes to time spent on single operation when using goroutines](https://www.reddit.com/r/golang/comments/frqerw/my_program_is_slower_when_it_comes_to_time_spent/)
- url: https://www.reddit.com/r/golang/comments/frqerw/my_program_is_slower_when_it_comes_to_time_spent/
---
Hello, I'm writing a path tracer in Golang, which is basically a Monte Carlo program for rendering photorealistic 3D scenes. Using Monte Carlo method, it implements random sampling. When I render each pixel, it doesn't need information about other pixels or about other samples of itself. Here's my sampling strategy:

* split samples evenly for each CPU core (eg. if I have 128 samples and 16 cores, then each core will render 8 samples
* each sample works like this:
   * for each row in a column:
      * shoot a ray towards a scene for each pixel in a row

Everything works well, I can see 100% CPU utilization on all cores, but when I compare time spent for single pixel sample I can see significantly better results if I sample without the first step above (so without goroutines). In one example scene I had average 20.677µs for single pixel sample with goroutines and 12.138µs without goroutines, so calculating single sample gets about 70% slower, which is a lot. Here's example code, I removed unnecessary parts and simplified it:

    cpus := runtime.NumCPU()
    runtime.GOMAXPROCS(cpus)
    
    buf := make([][]Color, cpus)
    canvas := make([]Color, vsize*hsize)
    
    for i := 0; i &lt; cpus; i++ {
    	buf[i] = make([]Color, vsize*hsize)
    }
    
    ch := make(chan int, cpus)
    
    samplesCPU := samples / cpus
    
    for i := 0; i &lt; cpus; i++ {
    	go func(i int) {
    		for s := 0; s &lt; samplesCPU; s++ {
    			source := rand.NewSource(time.Now().UnixNano())
    			generator := rand.New(source) // without using custom generator for each sample, it gets very slow, because of mutex lock
    			for y := vsize - 1; y &gt;= 0; y-- {
    				for x := 0; x &lt; hsize; x++ {
    					// calculate sample here
    
    					buf[i][y*hsize+x] = buf[i][y*hsize+x].Add(col) // add pixel value to temporary buffer, it'll get all averaged later
    				}
    			}
    		}
    		ch &lt;- 1
    	}(i)
    }
    
    for i := 0; i &lt; cpus; i++ {
    	&lt;-ch
    }
    close(ch)

I noticed another weird thing: when I use goroutines, but set my program to render on one core, I can see in the Task Manager that my program doesn't use only one core on 100%, but it uses multiple cores on low %. What am I missing here?
## [7][KubeAct - A Helm chart for hosting your own runner on Kubernetes to run jobs in your GitHub Actions workflows. ☸️ 🚀](https://www.reddit.com/r/golang/comments/frpag4/kubeact_a_helm_chart_for_hosting_your_own_runner/)
- url: https://www.reddit.com/r/golang/comments/frpag4/kubeact_a_helm_chart_for_hosting_your_own_runner/
---
Hey guys,  
Today I have completed an [**#opensource**](https://www.linkedin.com/feed/hashtag/?keywords=opensource&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848) project that I've built especially [**#devops**](https://www.linkedin.com/feed/hashtag/?keywords=devops&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848) enthusiasts &amp; I call it [**#kubeact**](https://www.linkedin.com/feed/hashtag/?keywords=kubeact&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848) .  
It is a package for [**#Kubernetes**](https://www.linkedin.com/feed/hashtag/?keywords=kubernetes&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848) ([**#K8s**](https://www.linkedin.com/feed/hashtag/?keywords=k8s&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848)) to deploy and host your own runners on the cluster to run jobs in your [**#githubactions**](https://www.linkedin.com/feed/hashtag/?keywords=githubactions&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848) workflows. It makes it easy to automate all your software workflows, now with world-class CI/CD. Build, test, and deploy your code right from [**#GitHub**](https://www.linkedin.com/feed/hashtag/?keywords=github&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6650351024202702848).  
Checkout : [https://github.com/itsksaurabh/kubeact](https://github.com/itsksaurabh/kubeact)

&amp;#x200B;

Feel free to review and contribute 🍻

[https:\/\/github.com\/itsksaurabh\/kubeact](https://preview.redd.it/pigrrip3ssp41.jpg?width=847&amp;format=pjpg&amp;auto=webp&amp;s=a810e3c237240168348d352b26502d057d01d7fc)
## [8][Call rust from golang](https://www.reddit.com/r/golang/comments/freknn/call_rust_from_golang/)
- url: https://www.reddit.com/r/golang/comments/freknn/call_rust_from_golang/
---
I've recently found this blog https://blog.filippo.io/rustgo/ Have anyone tried to implement this? Does really works?
## [9][Access a member in an anonymous C structure from Go](https://www.reddit.com/r/golang/comments/frnq02/access_a_member_in_an_anonymous_c_structure_from/)
- url: https://www.reddit.com/r/golang/comments/frnq02/access_a_member_in_an_anonymous_c_structure_from/
---
Hi,

I need to access a member ("imm_data") of an anonymous union in a C structure:

	struct ibv_send_wr {
			...
	        union {
	                __be32                  imm_data;
	                uint32_t                invalidate_rkey;
	        };

However, when compiling this with GCCGO it cannot find the reference:

	var sendWr C.struct_ibv_send_wr
    if imm &gt; 0 {
			...
            sendWr.imm_data = C.uint32_t(imm)
    }


```
error: reference to undefined field or method ‘imm_data’
```

For background, this is from the Infiniband kernel driver and this member was previously not in a union - and that worked.

I'm not a Go programmer, unfortunately.

Can anyone help?

Many thanks!
## [10][Question about pointers and references from a Golang learner](https://www.reddit.com/r/golang/comments/frne8v/question_about_pointers_and_references_from_a/)
- url: https://www.reddit.com/r/golang/comments/frne8v/question_about_pointers_and_references_from_a/
---
This is my first language that has pointers, so forgive me if this is a silly question.

I was looking at the `http` package yesterday and I saw this snippet of code: 

```
client := &amp;http.Client{
   CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

On the first line, why do they write `client := &amp;http.Client` instead of just `client := http.Client`? What practical differences does this mean?

`client` now holds the reference to the `http.Client` struct instead of the struct itself. But when you call `client.Get`, it gets dereferenced anyway. Why not just use the `client` variable for the struct itself instead of the pointer to it?

Also, please tell me if any of my terminology is wrong!
