# golang
## [1][Any good beginner books in Go that is updated covering from basics to intermediate concepts?](https://www.reddit.com/r/golang/comments/hneda0/any_good_beginner_books_in_go_that_is_updated/)
- url: https://www.reddit.com/r/golang/comments/hneda0/any_good_beginner_books_in_go_that_is_updated/
---
Very new to Go language. I know basics like conditionals,arrays,slice,loops,etc. But I see there are lots of terminologies like Buffered I/o, bufio, using command lines, flags, etc. I am finding it hard to understand all these things. Is there any beginner books,resources or any video tutorials/courses to understand these things from Scratch?
## [2][[can-haz-password] A Library for Generating Random, Rule Based Passwords](https://www.reddit.com/r/golang/comments/hne7bf/canhazpassword_a_library_for_generating_random/)
- url: https://github.com/kloeckner-i/can-haz-password
---

## [3][[PROJECT] Odin - The Programmable, Observable and Distributed Job Scheduler](https://www.reddit.com/r/golang/comments/hmy4mb/project_odin_the_programmable_observable_and/)
- url: https://www.reddit.com/r/golang/comments/hmy4mb/project_odin_the_programmable_observable_and/
---
TLDR; Odin is a programmable, observable and distributed job orchestration  system which allows for the scheduling, management and unattended  background execution of individual user created tasks on Linux based  systems. 

I worked on this as a final year project in college, and consulted this subreddit a lot during that time! I built Odin to change the way in which one can run and manage scheduled jobs, and now I'm looking for help to manifest it's potential!
 
You can check out the project [here](https://github.com/theycallmemac/odin). More info about the project can be found below!

I've set up [this Discord Server](https://discord.gg/gFr2Yq) to concentrate the development process as we get closer to v2.0.0! If you are interested in helping out in anyway feel free to join!

The primary objective of such a system is to provide users/teams a  shared platform for jobs that allows individual members to package their  code for periodic execution, providing a set of metrics and variable  insights which will in turn lend transparency and understanding into the  execution of all system run jobs. Odin aims to do this by changing the  way in which we approach scheduling and managing jobs.

Job schedulers by definition are supposed to eliminate toil, a kind  of work tied to running a service which is manual, repetitive and most  importantly - automatable. Classically, job schedulers are ad-hoc  systems that treat it’s jobs as code to be executed, specifying the  parameters of what is to be executed, and when it is to be executed.  This presents a problem for those implementing the best practices of  DevOps. DevOps is something to be practiced in the tools a team uses,  and when traditional job schedulers fail they introduce a new level of  toil in debugging what went wrong.

Odin treats it’s jobs as code to be managed before and after  execution. While caring about what is to be executed and when it will be  executed, Odin is equally concerned with the expected behavior of your  job, which is to be described entirely by the user’s code. This  observability can be achieved through a web facing user interface which  displays job logs and metrics. All of this will be gathered through the  use of Odin libraries (written in Go, Python, Bash and Node.js) and will help  infer the internal state of jobs. For teams, this means Odin can  directly help diagnose where the problems are and get to the root cause  of any interruptions. Debugging, but faster!

Again, if there's any interest in helping out or you have any suggestions please feel free to comment! Appreciate any stars also!
## [4][Keeping Your Modules Compatible](https://www.reddit.com/r/golang/comments/hmz28j/keeping_your_modules_compatible/)
- url: https://blog.golang.org/module-compatibility
---

## [5][How do you guys manage sql queries?](https://www.reddit.com/r/golang/comments/hn96kt/how_do_you_guys_manage_sql_queries/)
- url: https://www.reddit.com/r/golang/comments/hn96kt/how_do_you_guys_manage_sql_queries/
---
Hi everyone, i'm new to golang and still learning about it. Previously when I was developing on Laravel or Play Framework, I was used to using an ORM library. I've decided not to use any ORM library in golang.

`stmt := "select id, name, email, created from users where id = $1"`

`row := u.DB.QueryRow(stmt, id)`

`user := &amp;models.User{}`

`err := row.Scan(&amp;user.ID, &amp;user.Name, &amp;user.Email, &amp;user.Created)`

`if err != nil {`

`return nil, err`

`}`

I've have a few repeating queries in several locations (as shown above). It is fine until it gets too repeating, for example, if I needed a user in another table, I would have to copy the above query and repeats it again. I normally run the all queries in transaction. How to manage the repeating sql query in several different location of your code? If I were to change a field in the user table, I would have to add the changes to the several different locations of the code. which is bad as I will miss out one of them by accident one day.
## [6][convert source code to png image](https://www.reddit.com/r/golang/comments/hn78wd/convert_source_code_to_png_image/)
- url: https://www.reddit.com/r/golang/comments/hn78wd/convert_source_code_to_png_image/
---
https://i.redd.it/qg7e40rl8j951.gif

&amp;#x200B;

&amp;#x200B;

[https://github.com/skanehira/code2img](https://github.com/skanehira/code2img)
## [7][Build highly-scalable realtime apps with Super Graph GraphQL subscriptions](https://www.reddit.com/r/golang/comments/hmyrij/build_highlyscalable_realtime_apps_with_super/)
- url: https://www.reddit.com/r/golang/comments/hmyrij/build_highlyscalable_realtime_apps_with_super/
---
Super Graph is a GO library and standalone service that automagically compiles GraphQL into efficient SQL. Its goal is to make it fast and easy to work with databases. No more struggling with ORMs or managing handwritten SQL.

We've recently added what I feel hands-down is the coolest feature ever, GraphQL subscriptions. Subscribe to a query and instant updated results as soon as data related to the query changes. It's also designed to be highly scalable where only a single query is used for tens of thousands of subscriptions. A 100K plus subscriptions would only result in under 10 queries on your database. 

Take this example below even if 100K subscribers want to watch for comments on an equally large number of varied posts this would still only result in about 20 queries on the database every few seconds which even for the simplest of databases is a breeze to handle. 

```
subscription query {
   comments(where: { post_id: $post_id }) {
     id
     body
   }
}
```

https://github.com/dosco/super-graph
## [8][Looking for feedback on Go program - parses AWS credentials file and helps you set one as $AWS_PROFILE](https://www.reddit.com/r/golang/comments/hn6iu2/looking_for_feedback_on_go_program_parses_aws/)
- url: https://www.reddit.com/r/golang/comments/hn6iu2/looking_for_feedback_on_go_program_parses_aws/
---
Hey,

As someone who is just starting to learn Go and also has a job that requires dozens of AWS profiles in my \~/.aws/credentials file, I wrote this little program today. It parses through the credentials file and lets you pick an AWS profile and will copy the command to set that as your active profile to the clipboard. I tried to have it set the environment variable directly but learned that is not possible ([https://stackoverflow.com/questions/24938877/how-to-set-environment-variables-that-last-in-go](https://stackoverflow.com/questions/24938877/how-to-set-environment-variables-that-last-in-go)).

I'm definitely in the "just try and make it work" phase so looking for any feedback on best practices.

The program: [https://gist.github.com/gingimli/b97026c7b35b68b744861b37fdaacd01](https://gist.github.com/gingimli/b97026c7b35b68b744861b37fdaacd01)

Thanks!
## [9][Why Go’s Error Handling is Awesome](https://www.reddit.com/r/golang/comments/hmnhkz/why_gos_error_handling_is_awesome/)
- url: https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html
---

## [10][How to optimize this Golang job queue processor](https://www.reddit.com/r/golang/comments/hndhl8/how_to_optimize_this_golang_job_queue_processor/)
- url: https://www.reddit.com/r/golang/comments/hndhl8/how_to_optimize_this_golang_job_queue_processor/
---
I have written an engine in Go which receives successful payment information from 5 - 6 sources (web payment, offline payment, multiple mobile payment vendors) and then writes them to a final mysql table according to some business rules.

Before writing into the final db, some preprocessing has to happen, so I allow all the requests to land in an SQLITE database, which I use as a temporary job queue. A Go process checks sqlite every 300 ms, processes jobs in a sequential manner (FIFO) and writes a payment id to it which is generated after successful insertion in final mysql table. Multiple services from multiple vendors just leave their payment data there, get a job id in return and leave all further operations to my go job processor.

A new scenario has cropped up where the requesting call needs payment id in response immediately, it totally screws up my pub/sub kind of architecture. One solution (which is very inefficient) is to ask that service to write data to sqlite, sleep for a second and then check sqlite again if that item has been processed and get its payment id.

It works for now but it's gonna pose a problem when volume of calls grow, looking for better suggestions. Thanks in advance.
