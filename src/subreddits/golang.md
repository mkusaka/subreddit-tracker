# golang
## [1][Rate limit http middleware with few algorithms (Sliding Window, Leaky Bucket)](https://www.reddit.com/r/golang/comments/g308cp/rate_limit_http_middleware_with_few_algorithms/)
- url: https://www.reddit.com/r/golang/comments/g308cp/rate_limit_http_middleware_with_few_algorithms/
---
Hi, if you need rate limit [library](https://github.com/Shareed2k/go_limiter) or middleware for rate limit with algorithms like (Sliding Window, Leaky Bucket) , redis as store, you welcome to try test and give me a feedback... thanks ;)

[fiber framework middleware](https://github.com/Shareed2k/fiber_limiter)

[echo framework middleware](https://github.com/Shareed2k/echo_limiter)

[http middleware](https://github.com/Shareed2k/http_limiter)
## [2][Search and copy lyrics from the terminal (First go program)](https://www.reddit.com/r/golang/comments/g2vh9j/search_and_copy_lyrics_from_the_terminal_first_go/)
- url: https://github.com/asvvvad/cply
---

## [3][How I write an opensource cloud gaming service with WebRTC and Golang](https://www.reddit.com/r/golang/comments/g2e88k/how_i_write_an_opensource_cloud_gaming_service/)
- url: https://webrtchacks.com/open-source-cloud-gaming-with-webrtc/
---

## [4][Highly concurrent access to map?](https://www.reddit.com/r/golang/comments/g31vx8/highly_concurrent_access_to_map/)
- url: https://www.reddit.com/r/golang/comments/g31vx8/highly_concurrent_access_to_map/
---
Hi, i'm trying to build an efficient port scanner in Go.

Full code: [https://gist.github.com/hazcod/00af72341a5347e90673785ff02efd0d](https://gist.github.com/hazcod/00af72341a5347e90673785ff02efd0d)

However, how I would aggregate returns from a lot of goroutines? Use blocking channels to feed results back to the function?

    	ps.lock = semaphore.NewWeighted(maxDesc)
    
    	wg := sync.WaitGroup{}
    	defer wg.Wait()
    
    	results := map[uint]bool{}
    
    	for port := firstPort; port &lt;= lastPort; port++ {
    		if err := ps.lock.Acquire(context.TODO(), 1); err != nil {
    			return errors.Wrap(err, "could not acquire lock")
    		}
    
    		wg.Add(1)
    		go func(port uint) {
    			defer ps.lock.Release(1)
    			defer wg.Done()
    			fmt.Println(ScanPort(address, port, timeout))
    			results[port] = ScanPort(address, port, timeout)
    		}(port)
    	}
## [5][I want to make my own webrtc gateway sorts, in golang, is there anything like article or guide around it?](https://www.reddit.com/r/golang/comments/g30xde/i_want_to_make_my_own_webrtc_gateway_sorts_in/)
- url: https://www.reddit.com/r/golang/comments/g30xde/i_want_to_make_my_own_webrtc_gateway_sorts_in/
---
I do know about janus but i was wondering if there any more articles or guide around it.
## [6][Is there prominent pros in cobra over urfave/cli ?](https://www.reddit.com/r/golang/comments/g2v2yy/is_there_prominent_pros_in_cobra_over_urfavecli/)
- url: https://www.reddit.com/r/golang/comments/g2v2yy/is_there_prominent_pros_in_cobra_over_urfavecli/
---
I'm on situation that have to select one of two.

1. [https://github.com/spf13/cobra](https://github.com/spf13/cobra) 
2. [https://github.com/urfave/cli](https://github.com/urfave/cli)

In my short-term eyes, both project is alive.
## [7][How hard is this language for beginners from a scale to 1 to 10](https://www.reddit.com/r/golang/comments/g304f1/how_hard_is_this_language_for_beginners_from_a/)
- url: https://www.reddit.com/r/golang/comments/g304f1/how_hard_is_this_language_for_beginners_from_a/
---
apologies if i asked the same question but i think no one responded, so i am here yet again.

also i never programmed (well technically i watched some tutorials about some programming languages but meh)

so is this language hard? 

Thanks for anwsering!.
## [8][Tag search suggestions for my project](https://www.reddit.com/r/golang/comments/g2zyjc/tag_search_suggestions_for_my_project/)
- url: https://www.reddit.com/r/golang/comments/g2zyjc/tag_search_suggestions_for_my_project/
---
Hello guys, I am working on a project where I need to search the results based on user-defined tags. I believe the tags will be too much, thus causing performance issues when I search from a tag in my BoltDB.   
I would love to hear any suggestions from the community as to how should I implement the same. Its a CLI tool.
## [9][Generate GraphQL API from a postgres database](https://www.reddit.com/r/golang/comments/g2zhp3/generate_graphql_api_from_a_postgres_database/)
- url: https://blog.graphqleditor.com/graphqlize-instant-graphql-api-from-postgresql-mysql/
---

## [10][SDL2: How to make my program quit after 5 seconds and keep the application responding?](https://www.reddit.com/r/golang/comments/g2zbmw/sdl2_how_to_make_my_program_quit_after_5_seconds/)
- url: https://www.reddit.com/r/golang/comments/g2zbmw/sdl2_how_to_make_my_program_quit_after_5_seconds/
---
I am currently learning SDL2 binding for Go and Goroutines (I am a beginner in those two).

I would like my app to display a screen with ANYTHING, and close when either I click the close button of the window or the window has been displayed for 5 seconds. 

Let's say my code looks like this:

    package main
    
    import (
    	"fmt"
    
    	"github.com/veandco/go-sdl2/sdl"
    )
    
    const windowWidth = 800
    const windowHeight = 600
    
    func main() {
    	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
    		panic(err)
    	}
    	defer sdl.Quit()
    
    	window, err := sdl.CreateWindow("Testing SLD2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    		int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	defer window.Destroy()
    
    	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	defer renderer.Destroy()
    
    	running := true
    	for running {
    		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent()             {
    			switch event.(type) {
    			case *sdl.QuitEvent:
    				println("Quit")
    				running = false
    				break
    			}
    		}
    	}
    }

What I have managed to find out:

\- I could rewrite the last paragraph to wait for 5 seconds like this:

    	running := true
    	for running {
    		select {
    		case &lt;-time.Tick(5 * time.Second):
    			println("Quit")
    			running = false
    			break
    		}
    	}

But then my application is showing as "not responding" before it shuts down. It's not something that's happening in the tutorials I have checked, not sure if it's something related to the current version of the SDL2 binding module (github.com/veandco/go-sdl2 v0.4.1) or my OS (Mac OS X Catalina 10.15.2). My SDL binary version is 2.0.12.

\- To avoid the issue mentioned above, I would need to have a channel that polls the events all the time. However, it looks like the package does not allow to poll events in the goroutines.

Please help, I would really appreciate a piece of working code.
