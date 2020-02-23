# golang
## [1][Gopher made by my niece](https://www.reddit.com/r/golang/comments/f876c6/gopher_made_by_my_niece/)
- url: https://i.redd.it/plyj3iek3ni41.jpg
---

## [2][Crochet Gopher](https://www.reddit.com/r/golang/comments/f7yxss/crochet_gopher/)
- url: https://i.redd.it/3q0plnxujji41.jpg
---

## [3][The Zen of Go | Dave Cheney](https://www.reddit.com/r/golang/comments/f8804m/the_zen_of_go_dave_cheney/)
- url: https://dave.cheney.net/2020/02/23/the-zen-of-go
---

## [4][Gopher on the Go](https://www.reddit.com/r/golang/comments/f83p0q/gopher_on_the_go/)
- url: https://i.redd.it/n4tx6rkvbli41.jpg
---

## [5][Security in Go - Building a Port Scanner - Video Tutorial](https://www.reddit.com/r/golang/comments/f7vcoe/security_in_go_building_a_port_scanner_video/)
- url: https://www.youtube.com/watch?v=H-PWrdkowVA
---

## [6][Go-flavored Pascal: The old meets the modern in a tiny self-hosting compiler](https://www.reddit.com/r/golang/comments/f7zkp8/goflavored_pascal_the_old_meets_the_modern_in_a/)
- url: https://www.reddit.com/r/golang/comments/f7zkp8/goflavored_pascal_the_old_meets_the_modern_in_a/
---
By replacing the heavyweight Delphi-style OOP with a much simpler method/interface model inspired by Go, I have written an extremely compact (\~10000 lines) self-hosting [Pascal compiler](https://github.com/vtereshkov/xdpw) for Windows. It can be viewed as an implementation of Russ Cox's thought:

&gt;If I could export one feature of Go into other languages, it would be interfaces.

The compiler directly emits native x86 code and doesn't require any external assembler or linker. It can be easily embedded into larger software systems and used for educational purposes, e.g., as a playground for language design amateurs.
## [7][Design guidelines and principles for RESTful services](https://www.reddit.com/r/golang/comments/f89gd2/design_guidelines_and_principles_for_restful/)
- url: https://www.reddit.com/r/golang/comments/f89gd2/design_guidelines_and_principles_for_restful/
---
Hi Everyone,  
First post here though I've been lurking for a while.   
Started tinkering with Golang at work a while back and wrote some CLI tools but now I've been tasked with writing a RESTful-based mircoservice. I've decided to write in Golang. I'll mention up front that this microservice will be used only in the CI/CD pipeline to aid in automated tests and will never go into production.

I'm coming from a Python background and any prior REST services that I've written in the past were either in Python or NodeJS/Express. I chosen Golang because I had the freedom of choice, wanted something fast (hence why not in Python) and new (plus I really dislike Javascript).

That said, can anyone offer me some advice or resources how best to structure Goland projects, particularly RESTful APIs/Services. I've googled it but seem to get differing opinions.

The service itself will initially have 3/4 endpoints but this might grow so I'd rather get it correct from the get-go rather than having to refactor it massively later.

Also, any suggestions as to when to use Goroutines rather than direct function calls? I've been having fun with Golang and seems I want to put most things into a goroutine but maybe its overkill. The service in question will perform FS read/write operations so I figure (perhaps incorrectly) that goroutines are the way to go.
## [8][Testing Tips for Beginner Gophers](https://www.reddit.com/r/golang/comments/f848g9/testing_tips_for_beginner_gophers/)
- url: https://kinbiko.com/go/testing/beginners/
---

## [9][why i shouldn't use frame work and use the go lang library instead?](https://www.reddit.com/r/golang/comments/f88nhh/why_i_shouldnt_use_frame_work_and_use_the_go_lang/)
- url: https://www.reddit.com/r/golang/comments/f88nhh/why_i_shouldnt_use_frame_work_and_use_the_go_lang/
---
will i decide to take go more professional and use it in upcoming projects but i don't know why go community keep saying do not use frameworks and use the libraries  even like gorilla and  gin

i don't understand the concept behind this decision most of the time frame works help you to develop projects faster

note i'm taking about web development 
## [10][Trying to use /internal in local project?](https://www.reddit.com/r/golang/comments/f88dm1/trying_to_use_internal_in_local_project/)
- url: https://www.reddit.com/r/golang/comments/f88dm1/trying_to_use_internal_in_local_project/
---
Hi, I'd like to adopt the directory structure for my project as recommend on [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

In my project root's go.mod I have added:

`replace` [`github.com/user/x/internal/db`](https://github.com/user/x/internal/db) `=&gt; ../x/internal/db`

as my code is just running locally for now.

When trying to go run one test package from cmd/test/ I get the following error:

`cmd/test/main.go:12:2: use of internal package` [`github.com/user/x/internal/db`](https://github.com/user/x/internal/db) `not allowed`

&amp;#x200B;

Moving everything from *local* into *pkg* and updating paths works fine. But the code I've got in pkg is not really meant to be used as future public package.

Is it possible to use the internal approach just locally?  I am trying to be a good gopher here ;-)

&amp;#x200B;

([https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) ....?)
