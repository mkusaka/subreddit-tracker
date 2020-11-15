# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Oh my `gotch` - Most comprehensive Pytorch C++ API Go binding for deep learning in Go](https://www.reddit.com/r/golang/comments/judpna/oh_my_gotch_most_comprehensive_pytorch_c_api_go/)
- url: https://www.reddit.com/r/golang/comments/judpna/oh_my_gotch_most_comprehensive_pytorch_c_api_go/
---
[https://github.com/sugarme/gotch](https://github.com/sugarme/gotch)

We are happy to share a new toolkit for developing deep learning in Go - [gotch](https://github.com/sugarme/gotch). `gotch` is a Pytorch C++ API Go binding featuring with:

* Comprehensive Pytorch tensor APIs (\~ 1404)
* Fully featured Pytorch dynamic graph computation
* JIT interface to run model trained/saved using PyTorch Python API
* Load pretrained Pytorch models and run inference
* Pure Go APIs to build and train neural network models with both CPU and GPU support
* Most recent image models
* NLP Language models - [Transformer](https://github.com/sugarme/transformer) in separate package built with GoTch and [pure Go Tokenizer](https://github.com/sugarme/tokenizer).
## [3][Featherweight (Generic) Go @ OOPSLA 2020 by Robert Griesemer, Philip Wadler et. al.](https://www.reddit.com/r/golang/comments/ju82tt/featherweight_generic_go_oopsla_2020_by_robert/)
- url: https://www.youtube.com/watch?v=62xlcsJ0AUs
---

## [4][What is the current (1.15) best practice to work on interdependent modules?](https://www.reddit.com/r/golang/comments/juizr6/what_is_the_current_115_best_practice_to_work_on/)
- url: https://www.reddit.com/r/golang/comments/juizr6/what_is_the_current_115_best_practice_to_work_on/
---
With `GOPATH` on its way out, what is the current best practice when working on multiple modules that form a dependency tree?

For simplicity, let's assume two modules:

example.com/library/go.mod:

    module example.com/library
    
    go 1.15

example.com/project/go.mod:

    module example.com/project
    
    go 1.15
    
    require (
        example.com/library v0.0.1
    )

When working on these two modules concurrently, is there a way to have changes in library show up in project **without**

* adding and removing `replace` directives to `project/go.mod` all the time
* having to commit and update the library project on single-line changes
* crazy symlink hackery with vendoring
* falling back to `GOPATH` mode
* putting everything in one monorepo-like monster module

?

I'm pretty sure I read the relevant articles and blogs, and maybe I'm missing something very obvious. But I just cannot find a way to make this dev process work properly.

Any hints?
## [5][GORM - grandchild foreign key](https://www.reddit.com/r/golang/comments/jukcqg/gorm_grandchild_foreign_key/)
- url: https://www.reddit.com/r/golang/comments/jukcqg/gorm_grandchild_foreign_key/
---
Hello fellow go developers,

I'm currently creating a rest application in golang, and am using GORM for my orm. So far i successfully have everything implemented however i'm now wanting to add some stuff in order to make further queries easier.

This will be adding "grandchildren" foreign keys. I can't see anything in the documentation about this but effectivly what i'm wanting is the following: 

    type Map struct{
    	Id int `gorm:"primaryKey"`
    	Buildings []Building `gorm:"references:Id"`
    }
    
    type Building struct{
    	Id int	`gorm:"primaryKey"`
    	MapId int
    	Floors []Floor `gorm:"references:Id"`
    }
    
    type Floor struct{
    	Id              int `gorm:"primaryKey"`
    	BuildingId      int
            MapId           int
    }

From reading documentation I can't seem to find a sane way of doing this, if anyone can link to some documentation or an example that would be great, please note: I don't want to hold an instance of the map within the floor just the ID.

&amp;#x200B;

Thanks for your help.

Coffee-to-code
## [6][mtojek/gdriver : Download large files from Google Drive (API v3)](https://www.reddit.com/r/golang/comments/ju1brr/mtojekgdriver_download_large_files_from_google/)
- url: https://github.com/mtojek/gdriver
---

## [7][ANN: Apache H2 Database pure-go SQL Driver](https://www.reddit.com/r/golang/comments/jukaeq/ann_apache_h2_database_purego_sql_driver/)
- url: https://github.com/jmrobles/h2go
---

## [8][Article on how bad practices can be good with example from Go's standard library](https://www.reddit.com/r/golang/comments/ju4u87/article_on_how_bad_practices_can_be_good_with/)
- url: https://pmihaylov.com/good-and-bad-practices/
---

## [9][v1.4 of sessionup - HTTP sessions now with dynamic metadata support](https://www.reddit.com/r/golang/comments/ju8f78/v14_of_sessionup_http_sessions_now_with_dynamic/)
- url: https://github.com/swithek/sessionup
---

## [10][What's "wrong" with my Go code?](https://www.reddit.com/r/golang/comments/juexn6/whats_wrong_with_my_go_code/)
- url: https://www.reddit.com/r/golang/comments/juexn6/whats_wrong_with_my_go_code/
---
Hello!

I'm teaching myself Go and I don't know anyone else who knows Go. So far I like the language. The concurrency is great.

I've written a simple program called `mx-counter` that takes in a list of email addresses and outputs the count of each mail server.

https://github.com/tom-on-the-internet/mx-counter

Example:

Given

```
tom@tomontheinternet.com
jane@yahoo.ca
joe@gmail.com
fred@youtube.com
ingrid@slack.com
rory@microsoft.com
```

Outputs

```
google.com 4
yahoodns.net 1
outlook.com 1
```

It would mean the world to me if you could **point out some things I'm doing wrong, or places I could improve**. If I get enough feedback, I'll turn it into a blog post and/or video.

I also have a few questions. It would be incredible is you could answer any or all of them.

1. I wrote a function to get the unique values from a slice. Is this really something I should write? Should I be importing it from another package? In the JavaScript world, I would probably import a package for this. Not because I can't write this myself, but because it's not core to my project.

2. How should I test this? I've read that I should only test my public functions. But that would mean testing a single function that makes network requests. And if a domain changed their MX records, my tests would fail.

3. Should any one my functions be public? They aren't consumed by another package.

4. Should I split this into multiple files?

5. Am I approaching any of this in the way a seasoned Go developer would?

Thanks so much!

Tom
## [11][Go's Recurring Security Problem](https://www.reddit.com/r/golang/comments/jtmlve/gos_recurring_security_problem/)
- url: https://medium.com/tempus-ex/gos-recurring-security-problem-2b5339f19216
---

