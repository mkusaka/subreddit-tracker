# golang
## [1][spluggy: Static Plugins for Go](https://www.reddit.com/r/golang/comments/j9oia1/spluggy_static_plugins_for_go/)
- url: https://www.reddit.com/r/golang/comments/j9oia1/spluggy_static_plugins_for_go/
---
Everybody loves [**plugins**](https://en.wikipedia.org/wiki/Plug-in_(computing)). Well, not everybody, but they can offer many advantages. Arguably the biggest ones are modularization of your program and making it easy to extend it. I'm working on an open source project in Go and needed these benefits, so I opted for a plugin architecture. But how?

Go offers a [**way to create and load plugins**](https://golang.org/pkg/plugin/). Unfortunately it imposes a lot of [**technical limitations**](https://www.reddit.com/r/golang/comments/b6h8qq/is_anyone_actually_using_go_plugins/ejkxd2k/?utm_source=reddit&amp;utm_medium=web2x&amp;context=3). Yet even disregarding those, your functionality is spread over several binary files. I decided to create a small tool to solve this: [***spluggy***](https://github.com/codomatech/spluggy).

spluggy is a [**code generation**](https://blog.golang.org/generate) tool which helps you define static plugins (i.e. at compile time.) These plugins are then built into a single executable. It provides a way to have modularity and extensibility without managing multiple binaries.

spluggy is written with the ~~lazy~~ busy developer in mind. You don't need to maintain an explicit list of your plugins. You only have to define them as sub-packages within the same package. As long as they expose a common function, spluggy will discover your plugins and expose them to external packages.

spluggy is available on [**Github**](https://github.com/codomatech/spluggy). The repository has an [**example**](https://github.com/codomatech/spluggy/tree/main/example) demonstrating how to use spluggy in your project.
## [2][I made a small app to play spooky sounds for the Halloween season](https://www.reddit.com/r/golang/comments/j9f651/i_made_a_small_app_to_play_spooky_sounds_for_the/)
- url: https://github.com/fagnercarvalho/go-spooky-sounds
---

## [3][Use of Go for Distributed Computing](https://www.reddit.com/r/golang/comments/j97ct8/use_of_go_for_distributed_computing/)
- url: https://www.reddit.com/r/golang/comments/j97ct8/use_of_go_for_distributed_computing/
---
Hi all, I"m a newbie trying to better understand between go and distributed computing. Can you answer the below questions I have?

1. Noticed a lot of distributed computing/distributed systems class taught in go. Why is this?
1. How is Go actually used to manage cluster computing/HPC systems in practice? Examples would be helpful
1. Traditional HPC systems are build on C/C++ with openMP, MPI, etc. Can Go do the same things?
1. What are the pros/cons of building a modern distributed compute/HPC system today in Go vs C++? 


Thanks all.
## [4][Most efficient probabilistic datastructure: Bloom vs Cuckoo filters](https://www.reddit.com/r/golang/comments/j9qsb8/most_efficient_probabilistic_datastructure_bloom/)
- url: https://panmari.github.io/2020/10/09/probabilistic-filter-golang.html
---

## [5][Interface Semantics &amp; Polymorphism: Leveraging the interface to write code whose behavior can change depending on the data being passed through](https://www.reddit.com/r/golang/comments/j9a54c/interface_semantics_polymorphism_leveraging_the/)
- url: https://youtu.be/wLgpY-wzy2o
---

## [6][brightness - Backlight utility](https://www.reddit.com/r/golang/comments/j9h5xq/brightness_backlight_utility/)
- url: https://gitlab.com/tslocum/brightness
---

## [7][gin-brotli: A gin-gonic middleware for brotli compression](https://www.reddit.com/r/golang/comments/j9e8lv/ginbrotli_a_gingonic_middleware_for_brotli/)
- url: https://github.com/anargu/gin-brotli
---

## [8][Separating Data Store from the Domain](https://www.reddit.com/r/golang/comments/j9g79h/separating_data_store_from_the_domain/)
- url: https://www.reddit.com/r/golang/comments/j9g79h/separating_data_store_from_the_domain/
---
Hey Gophers,

I'm working on a personal project and am just working on figuring out how I'd like to structure my code. I'm running into a bit of a snag and wanted to see how you all might handle this.

I would like to use the "repository" pattern you see presented in Domain Driven Design (and used in plenty of other places). The idea is simple, to have a struct that acts as a repository and handles all interactions with whatever kind of data store I like. This way I avoid cluttering the domain with mapping tables to structs and executing queries.

Simple enough, you might end up with an interface that looks like this:

```go
package inventory

import "context"

type Repository interface {
	GetProduct(ctx context.Context, sku string) (Product, error)
	AddProduct(ctx context.Context, product Product, quantity int) (error)
	AddTransaction(ctx context.Context, product Product, quantity int) (error)
}
```

Pretty nice. Maybe the factory function that creates this repository takes in your database connection and the resulting struct uses that DB connection to implement the "GetProduct". This is an especially nice pattern because it makes mocking the "Repository" interface dead simple.

The rub I'm having here is this. Imagine that we wanted to have an atomic transaction that updated two tables. For instance you need to call the `AddProduct` function, and then the `AddTransaction` function above. We would want to begin a transaction, execute our updates, and commit if everything looks good. That transaction would need to be in the "model" layer, allowing the model to decide whether it makes sense to commit or not.

Your model layer function might look something like this:

```go
package inventory

type Inventory struct {
	repo *Repository
}

func (i *Inventory) AddInventory(ctx context.Context, product Product, quantity int) error {
	err := repo.AddProduct(ctx, product, quantity)
	if err != nil {
		return err
	}
	err := repo.AddTransaction(ctx, product, quantity)
	if err != nil {
		return err // Uh oh! A Problem happened here, we need to roll back adding the product!
	}
	return nil
}
```

The code above is no good. If we want to roll back our `AddProduct`, we simply don't have that option. I've thought about some possible solutions but they all seem... messy. The main advantage I'm looking for here is to be able to abstract away my repository so that I can cleanly design and test my domain model.

I'd love to hear if you've come up with an interesting solution in your personal or professional projects.
## [9][Am I wasting my time?](https://www.reddit.com/r/golang/comments/j9b6en/am_i_wasting_my_time/)
- url: https://www.reddit.com/r/golang/comments/j9b6en/am_i_wasting_my_time/
---
For a school semester project, I've decided to use Go as my backend language for a web application I have to build (including database and everything you could want from a web app).

Since I didn't know where to start, I enrolled in this course: [\*click\*](https://www.udemy.com/course/go-programming-language). It teaches me server-side programming with Go, without using any third-party libraries. This course is 20h in video-material alone. After this, I would still need to learn how to do the frontend side of things. 

**Is there any more efficient way I could be learning this?**

I'm really not sure whether this is the best course for what I'm looking for or whether there's a better way to go about this. Any suggestions/thoughts?
## [10][What are some of the best Golang video tutorials out there?](https://www.reddit.com/r/golang/comments/j9gnju/what_are_some_of_the_best_golang_video_tutorials/)
- url: https://www.reddit.com/r/golang/comments/j9gnju/what_are_some_of_the_best_golang_video_tutorials/
---

