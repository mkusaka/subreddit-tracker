# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Manual Memory Management in Go with jemalloc](https://www.reddit.com/r/golang/comments/jobgq9/manual_memory_management_in_go_with_jemalloc/)
- url: https://dgraph.io/blog/post/manual-memory-management-golang-jemalloc/
---

## [3][AppSec for Go powered by dynamic run-time instrumentation](https://www.reddit.com/r/golang/comments/jofuza/appsec_for_go_powered_by_dynamic_runtime/)
- url: https://blog.sqreen.com/dynamic-instrumentation-go/
---

## [4][Golang project structure &amp; Go vs other languages for REST and web applications](https://www.reddit.com/r/golang/comments/jogis7/golang_project_structure_go_vs_other_languages/)
- url: https://www.reddit.com/r/golang/comments/jogis7/golang_project_structure_go_vs_other_languages/
---
Hello everyone,

Hope this ok to post here, i'm not sure of anywhere else to post it.

TLDR; is go a good option for creating a REST API and also a web application or do other langages fit better such as PHP, Java etc. if ytou know of any good resources doing this I'd appreciate a link to it so i can read through.

TLDR2; my project structure is one of the things putting me off go, how else can i improve this?

**Background:**

I'm relativly new to go and have only ever used iit for small personal projects. I am currently working on my main project for the final year of my computer science degree. However i'm having a hard time choosing which language to used for my backend server development. Over the years i have mostly used PHP for this (again all small personal projects).

I have currently written the start of my API in go and am kind of torn on weather i should carry on, my project structure seems a little messy, and there seems to be alot of boiler plate and repeated code i can't break out (such as unmarshling requests into objects).

**Project structure:**

my project structure currently looks something like below. Each database file handles the CRUD function of that type for all entities (i'm using GORM). each handler class contains the handler function for that particular entity, and the rest service creates the MUX routes for this. along with the middleware. 

    project/
    ├── database/
    │   ├── DatabaseConnection.go
    │   ├── DatabaseCreate.go
    │   ├── DatabaseUpdate.go
    │   └── ...
    ├── rest/
    |   ├── Handlers/
    |   |   ├── Entity1Handler.go
    |   |   ├── Entity2Handler.go
    |   |   ├── ...
    │   ├── RestService.go
    

&amp;#x200B;

Thanks for reading

Happy coding :)
## [5][How do you determine the maximum number of concurrent requests for your Application?](https://www.reddit.com/r/golang/comments/jo1pka/how_do_you_determine_the_maximum_number_of/)
- url: https://www.reddit.com/r/golang/comments/jo1pka/how_do_you_determine_the_maximum_number_of/
---
For example when deploying to Cloud Run, I'm asked what levels of concurrency I want. What ways should I test these levels?
## [6][Hello,what would you recommend for someone trying to learn go?](https://www.reddit.com/r/golang/comments/jofua3/hellowhat_would_you_recommend_for_someone_trying/)
- url: https://www.reddit.com/r/golang/comments/jofua3/hellowhat_would_you_recommend_for_someone_trying/
---

## [7][Unit testing help](https://www.reddit.com/r/golang/comments/joc970/unit_testing_help/)
- url: https://www.reddit.com/r/golang/comments/joc970/unit_testing_help/
---
I've been looking for a good way to test a function that may contain multiple database operations like  \`MultipleOperationsOnUserFunc\`

    // Serivce
    type UserService struct {
        db SomeDB
    }
    
    func (s *UserService) MultipleOperationsOnUserFunc(id uuid.UUID) error {
        tx := db.Begin()
        user, _ := Get(tx, id)
        user.Name = "new name"
        _ := Update(tx, user)
        tx.Commit()
        ...
    }

Example Db operations:

    func Get(tx *sql.Tx, userID uuid.UUID) (user models.User, error) {...
    func Update(tx *sql.Tx, user models.User) error ...

For extra context, this is a big monolith with MANY operations.

Should I make one giant struct / interface for all database related functions and the mock it or is there a better way. (this is how I usually do it but i keep seeing examples like [https://github.com/Azure-Samples/azure-cosmos-db-cassandra-go-getting-started/blob/main/operations/crud.go#L28](https://github.com/Azure-Samples/azure-cosmos-db-cassandra-go-getting-started/blob/main/operations/crud.go#L28) which makes me wonder if i'm missing something)
## [8][Combining DDD, CQRS, and Clean Architecture by refactoring a Go project](https://www.reddit.com/r/golang/comments/join3w/combining_ddd_cqrs_and_clean_architecture_by/)
- url: https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/
---

## [9][Part time golang jobs](https://www.reddit.com/r/golang/comments/joi2qs/part_time_golang_jobs/)
- url: https://www.reddit.com/r/golang/comments/joi2qs/part_time_golang_jobs/
---
Where do ypu suggest to look for part time programming jobs? I checked upwork and I see that most of offers pay $3-$10 / hour.
Is it possible to find 10-15 hours a week job and to get $30+/hour?
## [10][Has anyone tried flamingo commerce?](https://www.reddit.com/r/golang/comments/jo98wi/has_anyone_tried_flamingo_commerce/)
- url: https://www.reddit.com/r/golang/comments/jo98wi/has_anyone_tried_flamingo_commerce/
---
I am trying to find a simple ecommerce solution in Go, but so far did not find much at all. There was a QoR (seems old and clunky). Today I searched again and stumbled upon flamingo commerce. It has detailed docs but I am not sure how easy to work with it as I am fairly new to Go. Is there anyone using it? What is your experience? I can't believe there are so few ecommerce solution using Go :(
## [11][Creating Your VERY OWN Command in Go [Why It's Good For You] Getting Started](https://www.reddit.com/r/golang/comments/jo2of1/creating_your_very_own_command_in_go_why_its_good/)
- url: https://www.youtube.com/watch?v=yl0phDUrnwc
---

