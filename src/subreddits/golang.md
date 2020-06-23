# golang
## [1][A Go package for drawing basic graphics in a terminal - Release v0.0.2](https://www.reddit.com/r/golang/comments/he9czv/a_go_package_for_drawing_basic_graphics_in_a/)
- url: https://github.com/daoleno/tgraph/releases/tag/v0.0.2
---

## [2][Test driving Go 2 Go Generics](https://www.reddit.com/r/golang/comments/he0h31/test_driving_go_2_go_generics/)
- url: https://www.youtube.com/watch?v=O4V-s9YntNk
---

## [3][Co-ordination between multiple services/modules in a single monolith](https://www.reddit.com/r/golang/comments/hedlbz/coordination_between_multiple_servicesmodules_in/)
- url: https://www.reddit.com/r/golang/comments/hedlbz/coordination_between_multiple_servicesmodules_in/
---
Hi, I am developing a medium scale monolith in Go. Which has 4 to 5 modules or services.

Suppose they are called: 

A, B, C, D

&amp;#x200B;

Each service is separated by individual packages and struct based API abstraction. They need to call/interact with each other. Now, what is the best way to design the architecture so that each can call each other's API easily?

One solution is to include all the service reference in each other's struct. Like service A's struct will have references for B, C, D so that A can call B. C and D's API. And service B will have references for A, C, D. So on and so forth.

Is this elegant? Is there any design pattern for this?
## [4][go-funk 0.7 released: a modern utility library which provides generic helpers (map, find, contains, keys, ...)](https://www.reddit.com/r/golang/comments/hea9co/gofunk_07_released_a_modern_utility_library_which/)
- url: https://github.com/thoas/go-funk/releases/tag/v0.7.0
---

## [5][How will generics change the official libs?](https://www.reddit.com/r/golang/comments/heeazm/how_will_generics_change_the_official_libs/)
- url: https://www.reddit.com/r/golang/comments/heeazm/how_will_generics_change_the_official_libs/
---
Go was designed without generics. Such need would be mitigated by using `interface{}` or a map workaround. Further: I've read many people defending the idea that Go didn't need generics.

Now, generics are coming to Go. So, basically, all that `interface{}` type assertion to mimic generics will become obsolete. But every Go library that exists to this point uses it.

So we will keep using them as they are, with `interface{}` casts, or for the official Go libs (at least) we'll have newer versions under different names?

...Or maybe new official libraries?

...Or different method names in the same libraries?

I read everyone talking about the generics syntax, but no one talking about the **consequences** of its inception. Am I missing something?
## [6][Ptracing a long-running process hangs](https://www.reddit.com/r/golang/comments/hecol7/ptracing_a_longrunning_process_hangs/)
- url: https://www.reddit.com/r/golang/comments/hecol7/ptracing_a_longrunning_process_hangs/
---
Hey everyone :)

I'm a novice Gopher so apologies if this is something obvious.

I'm trying to ptrace a long-running process, but for some reason the tracing hangs at `wait4` call (in code below, it's the call inside the for loop). The output seems to be of variable length before the process hangs. If the variable `len` is changed to something smaller, such as `9`, the code completes as expected.

I've done similar implementation in C and didn't face any issues there. I'd appreciate insights or tips how to debug this.


Code to reproduce the issue:
```
import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    len := "9999999"
    cmd := exec.Command("openssl", "rand", "-hex", len)
    cmd.SysProcAttr = &amp;syscall.SysProcAttr{Ptrace: true}
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    cmd.Start()
    pid, _ := syscall.Wait4(-1, nil, syscall.WALL, nil)

    for {
        syscall.PtraceSyscall(pid, 0)
        _, err := syscall.Wait4(-1, nil, syscall.WALL, nil) 

        if err != nil {
            fmt.Println(err)
            break
        }
    }
}
```

I'm running Go version `go1.14.4 linux/amd64`
## [7][Performance improvements in precise code intel](https://www.reddit.com/r/golang/comments/hdsrjr/performance_improvements_in_precise_code_intel/)
- url: https://about.sourcegraph.com/blog/performance-improvements-in-precise-code-intel
---

## [8][facing error ": cannot execute binary file: Exec format error "](https://www.reddit.com/r/golang/comments/heamp6/facing_error_cannot_execute_binary_file_exec/)
- url: https://www.reddit.com/r/golang/comments/heamp6/facing_error_cannot_execute_binary_file_exec/
---
I'm trying to follow instructions mentioned in [https://blog.realkinetic.com/load-testing-with-locust-part-1-174040afdf23](https://blog.realkinetic.com/load-testing-with-locust-part-1-174040afdf23)

While running ./example\_server I'm facing error: bash: ./example\_server: cannot execute binary file: Exec format error. on my ubuntu 18.04 system

I have checked the version as well and this is as follows:

sonali@sonali-Latitude-3490:\~/locust\_k8s$ uname -m

x86\_64

sonali@sonali-Latitude-3490:\~/locust\_k8s$ file example\_server

example\_server: Mach-O 64-bit x86\_64 executable, flags:&lt;NOUNDEFS&gt;

can someone please help.
## [9][How to enable Live Reload on VSCode with Go?](https://www.reddit.com/r/golang/comments/hea34e/how_to_enable_live_reload_on_vscode_with_go/)
- url: https://www.reddit.com/r/golang/comments/hea34e/how_to_enable_live_reload_on_vscode_with_go/
---
I have been learning golang for a while now and I work on VSCode, which works great. I use it, mainly for creating web applications.  
But whenever I have to make some changes in my code, I need to restart the server and then go to the browser to see the changes.   
It's very cumbersome and Live Server extension, doesn't work with it for some reason.  
In JS development, I just had to reload the browser window after updating the code.  


Any suggestions/tips on how to resolve this?
## [10][A Simple State Machine Framework in Go.](https://www.reddit.com/r/golang/comments/he9xqo/a_simple_state_machine_framework_in_go/)
- url: https://venilnoronha.io/a-simple-state-machine-framework-in-go
---

