# golang
## [1][GoLand 2020.1 EAP](https://www.reddit.com/r/golang/comments/f02u4p/goland_20201_eap/)
- url: https://blog.jetbrains.com/go/2020/02/06/welcome-to-the-goland-2020-1-eap/
---

## [2][GitHub - K-Phoen/grabana: User-friendly library for building Grafana dashboards](https://www.reddit.com/r/golang/comments/ezygm4/github_kphoengrabana_userfriendly_library_for/)
- url: https://github.com/K-Phoen/grabana
---

## [3][Web-server hexagonal architecture](https://www.reddit.com/r/golang/comments/f0ae4y/webserver_hexagonal_architecture/)
- url: https://www.reddit.com/r/golang/comments/f0ae4y/webserver_hexagonal_architecture/
---
I create this template for a web-server.  
It is created with fasthttp, fasthttp router, postgres and jsoniter.

It is based on the hexogonal architecture from [Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/).

&amp;#x200B;

Take a look at the repo and can you give some pointers what can be better or what you like.

The repo: [REPO](https://github.com/ChielTimmermans/go-hexagonal-web-server)
## [4][I move documentation from GitHub wiki to GitHub pages, let me know what you think.](https://www.reddit.com/r/golang/comments/f0a465/i_move_documentation_from_github_wiki_to_github/)
- url: http://rafallorenz.com/gorouter/
---

## [5][Including Go submodules part of a shared repository](https://www.reddit.com/r/golang/comments/f0a33z/including_go_submodules_part_of_a_shared/)
- url: https://blog.cloudbear.nl/private-go-submodules-using-go-import-meta-tags/
---

## [6][Convert a 492 lines shell entrypoint to a Go program? I did it!](https://www.reddit.com/r/golang/comments/f03bv5/convert_a_492_lines_shell_entrypoint_to_a_go/)
- url: https://www.reddit.com/r/golang/comments/f03bv5/convert_a_492_lines_shell_entrypoint_to_a_go/
---
Hello all!

I just wanted to share my gopher adventures!

I am the author of the VPN client Docker image `qmcgaw/private-internet-access` to connect to Private Internet Access servers, which turned out to be quite popular with today 2.1M Docker pulls in two years.

But its core functionality came from a [mutant /bin/sh entrypoint](https://github.com/qdm12/private-internet-access-docker/blob/60a69f316bfc4c19a4db29c1440d99dbc593ccda/entrypoint.sh) of **492 lines**.

Being a gopher, that could not continue. I finally merged my Go branch, packing all the features essentially in Go code. It is now [much more code](https://github.com/qdm12/private-internet-access-docker/tree/master/internal), but it feels much nicer :)

This subreddit also helped me merge stdout streams from different programs, so a big thank you to people who helped me!

PS: It's not fully unit tested yet, and I don't use iptables-go yet either, shame on me for now :)
## [7][[GORM]How to mix a Has One relationship and a Has Many on the same table](https://www.reddit.com/r/golang/comments/f09elf/gormhow_to_mix_a_has_one_relationship_and_a_has/)
- url: https://www.reddit.com/r/golang/comments/f09elf/gormhow_to_mix_a_has_one_relationship_and_a_has/
---
Starting from a basic `Has Many` relationship 

```Go
type Owner struct {
	ID         uint      `gorm:"primary_key"`
	Owned      []Owned   `gorm:"foreignkey:OwnerId"`
}

typ Owned struct{
	ID         uint      `gorm:"primary_key"`
	OwnerId    uint
}
```

Witch gave the following SQL database


```
Owner
ID 
1
2

Owned 
ID | OwnerId
1  |       1
2  |       1
3  |       1
4  |       2
5  |       2
```

I would like to have a column `FavoriteOwnedId` in the owner table to reference my favorite owned row. Which should look like that:

```
Owner
ID | FavoriteOwnedId
1  |               2
2  |               5

Owned 
ID | OwnerId
1  |       1
2  |       1
3  |       1
4  |       2
5  |       2
```

I imagine something like the following would allow me to get it 

```go 
owner := &amp;Owner
db.Where("id = ?", someId).
   Preload("FavoriteOwned").
   Preload("Owned").
   First(&amp;owner)
```

But I cannot find a way to write the necessary gorm tag to make it work. I manage to write the `Has One` or the `Has Many` but not writing them simultaneously.

Any idea ?
## [8][how do i call a postgres function from go?](https://www.reddit.com/r/golang/comments/f08uql/how_do_i_call_a_postgres_function_from_go/)
- url: https://www.reddit.com/r/golang/comments/f08uql/how_do_i_call_a_postgres_function_from_go/
---
im trying to call a postgres function in go, heres what i have so far:

 

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+  
 "password=%s dbname=%s sslmode=disable",  
        host, port, user, password, dbname)  
 db, err := sql.Open("postgres", psqlInfo)  
 if err != nil {  
 panic(err)  
    }  
 tx, err := db.Begin()  
    tx.Query("CALL data.create\_table();")  
 defer db.Close()  
 err = db.Ping()  
 if err != nil {  
 panic(err)  


if i call the function in postgres itself it works, go also doesnt return any errors or anything, but for some reason the function obviously doesnt run.

what am i doing wrong?
## [9][Rose – A Website Compiler Allowing the Use of Go Alongside HTML](https://www.reddit.com/r/golang/comments/ezvwdv/rose_a_website_compiler_allowing_the_use_of_go/)
- url: https://gitlab.com/DevelopmentDuck/rose
---

## [10][Github - 7days-golang - Let's finish a web framework or a distributed cache within a week](https://www.reddit.com/r/golang/comments/f04sky/github_7daysgolang_lets_finish_a_web_framework_or/)
- url: https://www.reddit.com/r/golang/comments/f04sky/github_7daysgolang_lets_finish_a_web_framework_or/
---
What can you write in 7 days? A gin-like web framework? A distributed cache like groupcache? Or a simple Python interpreter? Hope this repo can give you the answer.

[https://github.com/geektutu/7days-golang](https://github.com/geektutu/7days-golang)

## Web Framework - Gee

[Gee](https://geektutu.com/post/gee.html) is a [gin](https://github.com/gin-gonic/gin)\-like framework

* Day 1 - http.Handler Interface Basic [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day1-http-base)
* Day 2 - Design a Flexiable Context [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day2-context)
* Day 3 - Router with Tire-Tree Algorithm [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day3-router)
* Day 4 - Group Control [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day4-group)
* Day 5 - Middleware Mechanism [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day5-middleware)
* Day 6 - Embeded Template Support [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day6-template)
* Day 7 - Panic Recover &amp; Make it Robust [Code](https://github.com/geektutu/7days-golang/blob/master/gee-web/day7-panic-recover)

## Distributed Cache - Geecache

Geecache is a [groupcache](https://github.com/golang/groupcache)\-like distributed cache

* Day 1 - LRU (Least Recently Used) Caching Strategy [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day1-lru)
* Day 2 - Single Machine Concurrent Cache [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day2-single-node)
* Day 3 - Launch a HTTP Server [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day3-http-server)
* Day 4 - Consistent Hash Algorithm [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day4-consistent-hash)
* Day 5 - Communication between Distributed Nodes [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day5-multi-nodes)
* Day 6 - Cache Breakdown &amp; Single Flight | [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day6-single-flight)
* Day 7 - Use Protobuf as RPC Data Exchange Type | [Code](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day7-proto-buf)
