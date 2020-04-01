# golang
## [1][Help with concurrent linear search](https://www.reddit.com/r/golang/comments/fsxbhj/help_with_concurrent_linear_search/)
- url: https://www.reddit.com/r/golang/comments/fsxbhj/help_with_concurrent_linear_search/
---
I've been trying to write a simple concurrent linear search function in Go but haven't been too successful. I'm not sure if my understanding of channels and go routines is poor (I'm assuming it to be true) but here's my code -

    func simpleLinSearch(slice []int, num int, c chan int, startIndex int) {
        for index, item := range slice {
            if item == num {
                c &lt;- index + startIndex
                return
            }
        }
        c &lt;- -1
    }
    
    func concurrentLinSearch(slice []int, num int) {
        c := make(chan int)
        go simpleLinSearch(slice[:len(slice)/2], num, c, 0)
        go simpleLinSearch(slice[len(slice)/2:], num, c, len(slice)/2)
        x, y := &lt;-c, &lt;-c
        if x == -1 {
            return y
        } else {
            return x
        }
    }

A description of what I'm trying to achieve - the `concurrentLinSearch` function takes in a slice and a number to search for. It creates a channel. It splits the slice into two equal halves and calls `simpleLinSearch` that simply goes through each element of the slice and tries to write the index to the channel if found otherwise write -1 to it. I call this for both halves of the slice with appropriate starting indicies (because the second half of the slice would start from 0 but the elements in the original slice would have the index as `len(slice) / 2`). Then I read from the channel into x and y.

If x is -1 then it means the element wasn't in the half that was given to the function that wrote to the channel and vice versa.

It runs perfectly fine, `concurrentLinSearch` returns the correct index for a given slice and a given number to check in it but when I tried printing out the comparisions, I found that it runs linearly in both cases - not concurrently.

What am I doing wrong? Do the reads from the channel to `x` and `y` block the other goroutine so as to not read from the channel at the same time? I tried reading about `sync.WaitGroup` but couldn't go too far as I kept encountering deadlocks.
## [2][Let's make a tiny chess engine in Go](https://www.reddit.com/r/golang/comments/fsknxh/lets_make_a_tiny_chess_engine_in_go/)
- url: https://zserge.com/posts/carnatus/
---

## [3][Using Go for my Distributed Systems class](https://www.reddit.com/r/golang/comments/fsvlrc/using_go_for_my_distributed_systems_class/)
- url: https://www.reddit.com/r/golang/comments/fsvlrc/using_go_for_my_distributed_systems_class/
---
It is my final quarter of university and in my Distributed Systems class we will be implementing a RESTful API across multiple Docker containers, and the professor gave us the choice of programming language. I'm heckin' stoked cause my group is down to clown and write the assignment in Go.
## [4][Are you an employed software engineering using Go that is entirely self taught? What was your path to success and major milestones you remember from your first day of using Go to getting your offer letter?](https://www.reddit.com/r/golang/comments/fsegsq/are_you_an_employed_software_engineering_using_go/)
- url: https://www.reddit.com/r/golang/comments/fsegsq/are_you_an_employed_software_engineering_using_go/
---
I'm trying to learn from other people's experience and make a structured timeline with projects, resources, and milestones on my way to becoming a software engineer.
## [5][cancelling blocking read from stdin](https://www.reddit.com/r/golang/comments/fsxkqr/cancelling_blocking_read_from_stdin/)
- url: https://www.reddit.com/r/golang/comments/fsxkqr/cancelling_blocking_read_from_stdin/
---
Hey,

I'm processing stdin reading lines, but at some point, I want the user to be able to interrupt the process. This means stop channels etc but my goroutine that scans stdin will block until there is more data to be read, which might be an unreasonable amount of time in the future..

The way I've "solved" this feels wrong:

    lines := make(chan string)
    go func() { 
      s := bufio.NewScanner(os.Stdin)
      s.Split(bufio.ScanLines)
      for s.Scan() {
        lines &lt;- s.Text()
      }
    }()
    // ....
    go func() {
      for {
        select {
        case &lt;- stop:
          return
        case line := &lt;- lines:
         // process line
        }
      }
    }()

I think if I `close(stop)` then this will stop my second goroutine as I want, but then nothing is receiving on `lines` anymore, so `lines &lt;- s.Text()` will block indefinitely.

I can't come up with a better solution as I/O read is always exposed as synchronous, even if under the covers it isn't.

Should I just ignore this blocking goroutine? Is there a better solution? maybe if I `close(lines)` then I could handle the panic that I think will happen because `lines &lt;- s.Text()` is still trying to write to it..

Thanks!
## [6][A little help for a beginner :&gt; [type system]](https://www.reddit.com/r/golang/comments/fszo8l/a_little_help_for_a_beginner_type_system/)
- url: https://www.reddit.com/r/golang/comments/fszo8l/a_little_help_for_a_beginner_type_system/
---
Hey, could you help me understand type system in Go? Why does \*widget.Label started to be not compatible with fyne.CanvasObject.   
Thanks very much for the help!  
PS Do gophers have their char room?

https://preview.redd.it/nluov05si7q41.png?width=773&amp;format=png&amp;auto=webp&amp;s=8af869263e34dab07b624176d858c12a50b70779
## [7][Running Golang on the browser with WebAssembly and TinyGo](https://www.reddit.com/r/golang/comments/fszeix/running_golang_on_the_browser_with_webassembly/)
- url: https://marianogappa.github.io/software/2020/04/01/webassembly-tinygo-cheesse/
---

## [8][Go: How Does a Goroutine Start and Exit?](https://www.reddit.com/r/golang/comments/fsyl42/go_how_does_a_goroutine_start_and_exit/)
- url: https://medium.com/a-journey-with-go/go-how-does-a-goroutine-start-and-exit-2b3303890452
---

## [9][Why don't more people use Golang for scientific computing?](https://www.reddit.com/r/golang/comments/fsfqg0/why_dont_more_people_use_golang_for_scientific/)
- url: https://www.reddit.com/r/golang/comments/fsfqg0/why_dont_more_people_use_golang_for_scientific/
---
Because of its concurrency, wouldn't Go be a good candidate for lots of (especially) bioinformatics workflows where you're just running the same functions against lots of inputs in many cases?  Why haven't people in the scientific community taken to Golang?
## [10][Virtual Go Meetup - Come learn about WebRTC and how you can build sub-second decentralized real-time communications software!](https://www.reddit.com/r/golang/comments/fs6lrc/virtual_go_meetup_come_learn_about_webrtc_and_how/)
- url: https://www.meetup.com/golang/events/269676725/
---

