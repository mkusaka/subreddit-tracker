# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Stream and share your terminal without installing anything, server written in Go](https://www.reddit.com/r/golang/comments/i3tuxe/stream_and_share_your_terminal_without_installing/)
- url: https://github.com/miguelmota/streamhut
---

## [4][Does a module's name have to resolve to its location if it is shared?](https://www.reddit.com/r/golang/comments/i44fxy/does_a_modules_name_have_to_resolve_to_its/)
- url: https://www.reddit.com/r/golang/comments/i44fxy/does_a_modules_name_have_to_resolve_to_its/
---
If you name your module `github.com/myusername/mymodule` and publish it on GitHub in that location, if/when other people do `go mod download` after referencing the module in their package I'm assuming it uses that URL to fetch the module.

But if you name your module `example.com/mymodule`, and that doesn't actually resolve to the git repository of `mymodule`, does that break things if somebody was to `go mod download` (or build or test)?

In other words, is the module name used as the url for others to retrieve it from?
## [5][Seaworthy - A CLI to verify Kubernetes resource health](https://www.reddit.com/r/golang/comments/i3zzeo/seaworthy_a_cli_to_verify_kubernetes_resource/)
- url: https://www.reddit.com/r/golang/comments/i3zzeo/seaworthy_a_cli_to_verify_kubernetes_resource/
---
[https://github.com/cakehappens/seaworthy](https://github.com/cakehappens/seaworthy?ts=4)  


I started working on this CLI because I wanted to encapsulate the feature that I've found in [ArgoCD](https://argoproj.github.io/argo-cd/) as well as [Spinnaker](https://spinnaker.io/) that enables those tools to deploy resources to k8s and verify the health of the resources they deploy.  


The goal here is to democratize and enable simple workflows such as  


`kubectl apply -f ./manifests` 

`seaworthy verify -f ./manifests --timeout 5m`  


This would enable folks, such as myself, to make better use of other workflow tools, such as GitHub actions.  


This is just a simple example, but I plan on full support for a variety of ways to pass in resource information:  


`seaworthy verify deployments`

`seaworthy verify deployment web`

`jsonnet ./main.jsonnet | seaworthy verify --input-format json -f -`

`tanka show . | seaworth verify -f -`

No official release yet, I just made the asciinema video this evening, but I figured I'd post here and get some feedback.
## [6][Olric v0.3.0-beta.1 is out: Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.](https://www.reddit.com/r/golang/comments/i3i6gi/olric_v030beta1_is_out_distributed_cache_and/)
- url: https://github.com/buraksezer/olric/releases/tag/v0.3.0-beta.1
---

## [7][Notify.is - my first Go project](https://www.reddit.com/r/golang/comments/i3jsyv/notifyis_my_first_go_project/)
- url: https://www.reddit.com/r/golang/comments/i3jsyv/notifyis_my_first_go_project/
---
Hi there,

For the past month or so I've been working on my first Go project, a service called [Notify.is](https://notify.is) which notifies you when your favourite username on Instagram, Twitter or GitHub becomes available.

The frontend uses the React framework Next.js and the backend uses Go to deal with RESTful API routes for signing up and deleting data.

The Go code that checks with the mentioned services is deployed in a Docker container on Google Cloud Run.

I've had a lot of fun building it and would love any feedback on the design, code or anything else. Also, trying out the service would be helpful for me finding any bugs I have not yet found.

The frontend and backend are deployed using Vercel and the GitHub repository can be found here:

[https://github.com/oliverproud/notify.is](https://github.com/oliverproud/notify.is)

The Go code that is deployed in a Docker container on Google Cloud Run, you can find the GitHub repository for that here:

[https://github.com/oliverproud/notify.is-gcloud](https://github.com/oliverproud/notify.is-gcloud)

The website:

[https://notify.is](https://notify.is)

Let me know what you think!
## [8][Automatic deploy of debian on ESXi host](https://www.reddit.com/r/golang/comments/i412v9/automatic_deploy_of_debian_on_esxi_host/)
- url: https://www.reddit.com/r/golang/comments/i412v9/automatic_deploy_of_debian_on_esxi_host/
---
hi guys, here's bootp/ansible/golang tool that via command line allows you to deploy a VM

feel free to contribute

[https://github.com/lucabodd/ESXi-vm-deploy](https://github.com/lucabodd/ESXi-vm-deploy)
## [9][Golang Developers Can Make the Shift to Brighter Business Future (Build Nex-gen Enterprise solution like these Big Companies)](https://www.reddit.com/r/golang/comments/i440fx/golang_developers_can_make_the_shift_to_brighter/)
- url: https://www.bacancytechnology.com/blog/golang-for-brighter-business-future
---

## [10][Database Resolver - advanced read/write supports for GORM V2](https://www.reddit.com/r/golang/comments/i3kons/database_resolver_advanced_readwrite_supports_for/)
- url: http://v2.gorm.io/docs/dbresolver.html
---

## [11][What is the best design of database connection?](https://www.reddit.com/r/golang/comments/i3mj2w/what_is_the_best_design_of_database_connection/)
- url: https://www.reddit.com/r/golang/comments/i3mj2w/what_is_the_best_design_of_database_connection/
---
Hi everyone.There are two ways in my current usage.

1. I used the gin and gorm ,set into a context with a middlewares like

```go

// middleware

db :=NewDB(datasourcename..)

c.set("db",db)

// Usage

db := c.Value("db").(*gorm.DB)

userRepo := repository.NewUserRepository(db)

userServ := service.NewUserService(userRepo)

userHandler := handles.NewUserHandler(userServ)

r.Post("/user/create/",userHandler.Create)

```

2.  use a private variable in db package then use a struct variable call Get function

```go

// db.go

var conn *gorm.DB

type DB struct()

func (d *DB) Get() *gorm.DB

// usage

db := &amp;db.DB{}

userRepo := repository.NewUserRepository(db.Get())

userServ := service.NewUserService(userRepo)

userHandler := handles.NewUserHandler(userServ)

r.Post("/user/create/",userHandler.Create)

```

I tired to set the db into the request context. but it seems like the first way.And I need to combine the context.Context with gin.Context,then send ctx in ervery router function NewRepostiory function..

```go

ctx :=context.WithContext(context.Background(),"db",db)

c.Request = c.Request.WithContext(ctx)

```

How do you design the database connection? ? I don't have too much develop experience. Thanks for share !
## [12][Switching to pgx.. when to use connection pool, and do i pass context from my http handler?](https://www.reddit.com/r/golang/comments/i3vb9z/switching_to_pgx_when_to_use_connection_pool_and/)
- url: https://www.reddit.com/r/golang/comments/i3vb9z/switching_to_pgx_when_to_use_connection_pool_and/
---
So I was using the db.sql and pg driver for accessing database. I was using the var route with a global db, and I am not clear if that is a good way to go in terms of handling database connections concurrently from different http endpoint handlers that are running in their own threads (go funcs).

Started digging in a bit more and it seems the better way to go is using pgx. pgx however unlike the pg driver, seems to indicate the connection you get is not thread safe, and to use the connection pool for that. So looking into it, I am attempting to create a struct with methods that contains a pointer to the created pool, and make that accessible to then grab a connection to make a db request. Is that the right approach.. and will it scale to say, 100s of requests a second (or more?). 

Also, the example code I found indicates it requires a context. I can use context.Background(), but as each http handler has a context associated with it, does it make sense to pass that one in and use that, since in my service starter code I am also using a context to handle a graceful shutdown though I am not sure that that context is in any way related to the one each request has associated with it. Or.. would it make sense to pass the context in my service starter that handles graceful shutdown, to each connection so that if somehow the db cancels.. it ends up using the context created with the service starter to trigger a graceful shutdown?

Sorry..little confused on how all these contexts are used in this manner.
