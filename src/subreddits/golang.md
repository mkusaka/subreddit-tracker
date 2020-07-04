# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][What is the most idiomatic way of getting input from the terminal?](https://www.reddit.com/r/golang/comments/hkz1im/what_is_the_most_idiomatic_way_of_getting_input/)
- url: https://www.reddit.com/r/golang/comments/hkz1im/what_is_the_most_idiomatic_way_of_getting_input/
---
I have been spoilt by the simple `input()` function in Python. 

What is the most simple way of getting input from an external source? I have made several google searches and everyone often seems to recommend using different methods.

There is a `bufio.NewReader`. There is also a `bufio.NewScanner` , and then there's also a `fmt.Scanln`...
## [3][An opinionated guideline for structuring web applications/services](https://www.reddit.com/r/golang/comments/hl2fnw/an_opinionated_guideline_for_structuring_web/)
- url: https://www.reddit.com/r/golang/comments/hl2fnw/an_opinionated_guideline_for_structuring_web/
---
hello all, I made an opinionated guideline for creating/structuring a web application. I’ve seen a few out there some felt too complex, and some seemed confusing. I tried using a sample (non-functional app) app to explain how to make use of the structure. And explain why things are setup the way they are. Hope you find it useful! [https://github.com/bnkamalesh/goapp](https://github.com/bnkamalesh/goapp)
## [4][Ways to avoid an import cycle when keeping interface definition and implementation separate ?](https://www.reddit.com/r/golang/comments/hkzbou/ways_to_avoid_an_import_cycle_when_keeping/)
- url: https://www.reddit.com/r/golang/comments/hkzbou/ways_to_avoid_an_import_cycle_when_keeping/
---
Hi everyone!

I've been writing a go application in which I've two packages: a "db" package that interacts with database to execute queries and a "X" package that imports "db" and uses its methods. "db" package defines an interface and I've kept the implementation of this interface in the same package, which is "db" itself:

    package db
    
    type Store interface {
        GetUserByID(ctx context.Context, userID string) (user, error)
        GetUserByName(ctx context.Context, name string) (user, error)
    }

&amp;#x200B;

    package db
    
    type postgresClient struct {}
    
    func New() Store {
        return postgresClient{}
    }
    
    func (pg postgresClient) GetUserByID(ctx context.Context, userID string) ...
    func (pg postgresClient) GetUserByName(ctx context.Context, name string) ...

&amp;#x200B;

    package X
    
    type x struct {
        store db.Store
    }
    
    func NewX() x {
        dbInstance := db.New()
        
        return x{store: dbInstance}
    }

I've found that interfaces should be defined where they are being used. But If I had defined the "Store" interface inside the "X" package then I'd have to import it inside "db" package and that would create an import cycle. In that case, I would require a third package to initialize a "postgresClient" instance and pass it to the "New" function of package "X". But, With my current approach, I just need two packages.

It would be very helpful to know how would you guys structure the above application, keeping the interface definition and implementation in separate packages. I understand that there's no absolute solution to this and it totally depends on the use case, application etc, but getting some ideas would definitely help me.

Note: This might be a very silly question but please bear with me, I'm still learning. Thanks for your patience and efforts.

EDIT: I forgot to mention that the "user" type currently lives in the "db" package but according to the business logic it makes more sense to keep it inside package "X" but just to avoid import cycle, I kept it inside "db" package.  So, to add more complexity to my question, in addition to having interfaces defined inside package "X", I would also want to keep the "user" type inside "X".
## [5][Canvas API built with Gio](https://www.reddit.com/r/golang/comments/hkzz12/canvas_api_built_with_gio/)
- url: https://github.com/ajstarks/giocanvas
---

## [6][Finding the correct PATH](https://www.reddit.com/r/golang/comments/hl0jfa/finding_the_correct_path/)
- url: https://blog.haroldadmin.com/finding-right-path/
---

## [7][Started learning Go today. Is my understanding correct that GOPATH is no longer required in favour of Go modules?](https://www.reddit.com/r/golang/comments/hkishj/started_learning_go_today_is_my_understanding/)
- url: https://www.reddit.com/r/golang/comments/hkishj/started_learning_go_today_is_my_understanding/
---
My [Go Programming Language](https://www.gopl.io/) book mentions the use of GOPATH as a means to store Go code related to source files, packages and a bin directory. I'm slightly confused about this, does that mean that the only place I can write go code would be in the $GOPATH ?

Also, [go w/tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go#go-modules) states that GOPATH will be deprecated soon. So should I just forget about this concept and move on with Go modules?

Truth be told, I haven't researched into the benefits of Go modules too much as of yet as I'm just trying to set up my local Go environment before I start coding. It seems that I should just forget about GOPATH and just get cracking with code. Would anyone be able to ease my mind off this issue?

Much appreciated :)
## [8][An introduction to concurrency in Go (as a beginner)](https://www.reddit.com/r/golang/comments/hkn4tj/an_introduction_to_concurrency_in_go_as_a_beginner/)
- url: https://therebelsource.com/blog/an-introduction-to-concurrency-in-go-as-a-beginner/wBHuGLZRt0l
---

## [9][will Go generics remove the need of the interface{} wildcard type?](https://www.reddit.com/r/golang/comments/hl2402/will_go_generics_remove_the_need_of_the_interface/)
- url: https://www.reddit.com/r/golang/comments/hl2402/will_go_generics_remove_the_need_of_the_interface/
---
Basically a generic type T without any constraints is the interface{} type? So will this replace the use of interface{} in go code? Will generics break compatibility or best practices and when will they arrive in the language?
## [10][Github-like calendar heatmap in plain Go](https://www.reddit.com/r/golang/comments/hkgcz9/githublike_calendar_heatmap_in_plain_go/)
- url: https://github.com/nikolaydubina/calendarheatmap
---

## [11][Module and package management](https://www.reddit.com/r/golang/comments/hkzqz3/module_and_package_management/)
- url: https://www.reddit.com/r/golang/comments/hkzqz3/module_and_package_management/
---
Hi everyone, I'm new in go and I have a simple question. When I'm running `go get "some package"` it is just changing my go.mod file and adding the new package with the version, but actually it is not downloading to my $GOPATH/src/. I have run  the `go get` command in the $GOPATH and then it has downloaded. Is it the right way how it should work?
