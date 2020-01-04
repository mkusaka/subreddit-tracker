# golang
## [1][Going all in with Golang from PHP. Am I too ambitious ?](https://www.reddit.com/r/golang/comments/ejjppg/going_all_in_with_golang_from_php_am_i_too/)
- url: https://www.reddit.com/r/golang/comments/ejjppg/going_all_in_with_golang_from_php_am_i_too/
---
Sorry couldn't really construct the title well but this is not meant to be clickbaity. Our current SAAS product is PHP based and we are a 100% PHP team at the moment. I am in love with Golang and being the founder of the company, I decided to rewrite in Go due to the following reasons:

\- Better performance. We sell a platform and performance is very important.

\- Ease of deployment. Go Binaries are awesome.

\- I want to find more technical partners who can promote. Using Golang, I can just give them the binary to self host wherever.

\- Our current PHP stack is too bloated and since our platform can be customized for each client, the frameworks and bloat get in our way. It will be easier to have a slimmed down version and then customize on top for each client. Go's simplicity has won me over.

\- We need to remove/simplify some features that are not really needed or are coded in a complex way. 

\- Better suited to build REST APIs. I don't hate PHP (it has paid the bills) but I really like how simple Golang is with its http library

\- Attract better developers ? This one has been so controversial. I want to attract better talent. Does that help ? Or I am stupid.

&amp;#x200B;

The cons:

\- No one in our dev. team knows Golang. I  finally broke the news to my CTO and he was like WTF. He is very good at what he does and comes from a .NET background. So he can pick this up in 2-3 weeks BUT our team is mostly junior devs and they will need to learn Golang.

\- Existing customers on the PHP platform will need to be supported forever. We cannot retire the legacy platform that quickly due to the complex functions built for clients.

Am I committing business suicide ?

&amp;#x200B;

EDIT: Wow, so many great responses. I do appreciate everyone and I will think hard before putting our team at risk. It seems like the consensus is to not do a full re-write but consider chipping away slowly and may be build a few smaller components in Go as microservice if at all.
## [2][Go: g0, Special Goroutine](https://www.reddit.com/r/golang/comments/ejjglc/go_g0_special_goroutine/)
- url: https://medium.com/a-journey-with-go/go-g0-special-goroutine-8c778c6704d8
---

## [3][Im at a point in learning go where I just don't understand it anymore (Dependency Injection) what is the wisest move to do?](https://www.reddit.com/r/golang/comments/ejn44m/im_at_a_point_in_learning_go_where_i_just_dont/)
- url: https://www.reddit.com/r/golang/comments/ejn44m/im_at_a_point_in_learning_go_where_i_just_dont/
---
Would reading the previous chapters in the tutorial I follow work or is it just a matter of time and excercise?
## [4][Using local development modules without pushing to Github everytime](https://www.reddit.com/r/golang/comments/ejsgl0/using_local_development_modules_without_pushing/)
- url: https://www.reddit.com/r/golang/comments/ejsgl0/using_local_development_modules_without_pushing/
---
So I am trying to learn the new module system where I put my codebase outside of GOPATH and I am trying to figure out if my development flow is incorrect or not.

E.g. I developed this package/library: https://github.com/psykidellic/gomath. I checkout the repo $HOME/projects/gomath and start hacking on it.

Now, I want to use this library in another of my side project $HOME/projects/"gomathtest". Now if I push changes from gomath to github and use it in my new project 

    import (
        "github.com/psykidellic/gomath/calc"
    )

everything works.

But this means, during development I have to constantly push to github. Is there a way I can use the library gomath in gomathtest with the same import but then use the code from my local projects folder during development?

In my final ci/cd, obviously I want to do the correct way where code is actually imported from github.
## [5][koanf: Light weight extensible library for reading config (file, S3 etc.) in Go applications. Built in support for JSON, TOML, YAML, env, command line.](https://www.reddit.com/r/golang/comments/ejjyz8/koanf_light_weight_extensible_library_for_reading/)
- url: https://github.com/knadh/koanf
---

## [6][Chaos-mesh: A Chaos Engineering Platform for Kubernetes](https://www.reddit.com/r/golang/comments/ejbgmn/chaosmesh_a_chaos_engineering_platform_for/)
- url: https://github.com/pingcap/chaos-mesh
---

## [7][Where to place project source code?](https://www.reddit.com/r/golang/comments/ejkbx0/where_to_place_project_source_code/)
- url: https://www.reddit.com/r/golang/comments/ejkbx0/where_to_place_project_source_code/
---
Hi, newbie Gopher here. I'm starting to learn Go for fun and I'm a bit confused about the relationship between the GOPATH directory and where my project source code should live. From the Go wiki on GitHub, I see that:

_The GOPATH environment variable is used to specify directories outside of $GOROOT that contain the source for Go projects and their binaries._

Does this include my own projects? Am I breaking a convention if my Go source lives in some other directory? I'm using Go `1.13.5`
## [8][Blackhat go book](https://www.reddit.com/r/golang/comments/ejjtbe/blackhat_go_book/)
- url: https://www.reddit.com/r/golang/comments/ejjtbe/blackhat_go_book/
---
Hey, 

I'm learning/writing go for 2 months or so , and I saw this book . It's nearly out and a preview is available for few months . 


Any recommendations? 


Cheers
## [9][opencensus with jaeger, client and server example](https://www.reddit.com/r/golang/comments/ejoe3a/opencensus_with_jaeger_client_and_server_example/)
- url: https://github.com/juanpabloaj/opencensus-go-example
---

## [10][Proper way to structure code? (for testing)](https://www.reddit.com/r/golang/comments/ejo9er/proper_way_to_structure_code_for_testing/)
- url: https://www.reddit.com/r/golang/comments/ejo9er/proper_way_to_structure_code_for_testing/
---
I have a service that uses various external dependencies that im currently refactoring. (Rest, DB, Object Storage). I want to be able to test the core code, while stubbing/mocking/etc the external dependencies whilst having clean, well structured code.

Based on various strategies I've read, I've created a package for each dependency (Data, UserService, Storage) while having interfaces for required functionality in the core code. (Data.StoreImageMetadata(imageID int, data ImageData), UserService.GetUserData(userID int), Storage.Upload(data \[\]byte))

I have some concerns with this. 

1. It feels too object oriented which seems to be an anti-pattern in go
2. For the database package. I want to support transactions across various functions. Im not sure how to support this without using sql.TX in my business logic and passing it around to my data functions.

I have some thought on #2. The methods in the data package accept a context for other unrelated purposes. I was considering taking advantage of that. I am weary of using context, since it's meant for request scoped data and seems to be a hack rather than an answer most of the time.

I would provide a startTx(ctx context.Context) function in the Data interface that creates a transaction and store it into the context in the database code. This allows the business logic to not require an import to sql. Commit/Rollback would look the same. Then any other database functions would check the context for a TX and if present execute statements on that TX instead of as standalone statements.

All of this stems from wanting to properly structure the code, and making testing easier.

Any thoughts are appreciated.
