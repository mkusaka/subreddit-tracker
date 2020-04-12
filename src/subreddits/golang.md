# golang
## [1][3mux: Terminal multiplexer inspired by i3, in Go](https://www.reddit.com/r/golang/comments/fzttbv/3mux_terminal_multiplexer_inspired_by_i3_in_go/)
- url: https://github.com/aaronjanse/3mux
---

## [2][What ORM are you use?](https://www.reddit.com/r/golang/comments/fzs8p1/what_orm_are_you_use/)
- url: https://www.reddit.com/r/golang/comments/fzs8p1/what_orm_are_you_use/
---
I have used gorm for a project but I dont like it so much.  
I found pop from Buffalo has anyone expierence with this orm?

Or can recommend me another?
## [3][HFD – a design language for laser cutters written in Go](https://www.reddit.com/r/golang/comments/fztwf3/hfd_a_design_language_for_laser_cutters_written/)
- url: https://github.com/dustismo/heavyfishdesign/
---

## [4][Statically compiling Go programs](https://www.reddit.com/r/golang/comments/fz8piz/statically_compiling_go_programs/)
- url: https://www.arp242.net/static-go.html
---

## [5][A chain replicated KV store using Kubernetes, Golang &amp; Redis](https://www.reddit.com/r/golang/comments/fzwtff/a_chain_replicated_kv_store_using_kubernetes/)
- url: https://github.com/el10savio/Kube-Chain-Replication
---

## [6][Managing Shared Components in Go Microservices with Fx](https://www.reddit.com/r/golang/comments/fzvs3f/managing_shared_components_in_go_microservices/)
- url: https://pmihaylov.com/shared-components-go-microservices/
---

## [7][How to approximately convert an Unicode character to ASCII?](https://www.reddit.com/r/golang/comments/fzr5gm/how_to_approximately_convert_an_unicode_character/)
- url: https://www.reddit.com/r/golang/comments/fzr5gm/how_to_approximately_convert_an_unicode_character/
---
Say I have s string like: "I love eating Pâté"

See the word Pâté has several non-ASCII characters, and I want to convert it to the closest approximation of ACSII ones, like â -&gt; a and é -&gt; e.

In python, there is the unicodedata module to do it, though I need in Go. So I was wondering if there is some package to do it in Go.
## [8][chasquid (SMTP server in Go) v1.3 released](https://www.reddit.com/r/golang/comments/fzv52k/chasquid_smtp_server_in_go_v13_released/)
- url: https://www.reddit.com/r/golang/comments/fzv52k/chasquid_smtp_server_in_go_v13_released/
---
Hi!

[chasquid](https://blitiri.com.ar/p/chasquid/) v1.3 has been released.

It is an open source SMTP server written in Go, with a focus on simplicity, security, and ease of operation.

The main changes in this release are:

 - Improved handling of DNS temporary errors.                                                                                                                       
 - Documentation updates (use of SRS when forwarding, Dovecot  troubleshooting, Arch installation instructions).                                                                                                                                   
 - Miscellaneous test improvements and cleanups.                                                                                                                    


If you have any questions, feel free to ask here, on the [mailing list](https://groups.google.com/forum/#!forum/chasquid), or IRC (*#chasquid* on Freenode)!

[Git repository](https://blitiri.com.ar/git/r/chasquid/) / [Github mirror](https://github.com/albertito/chasquid)\
[Installation guide](https://blitiri.com.ar/p/chasquid/docs/install/) / [Debian+Dovecot how-to](https://blitiri.com.ar/p/chasquid/docs/howto/) / [Docker how-to](https://blitiri.com.ar/p/chasquid/docker/)
## [9][A GO library to fetch data with GraphQL instead of struggling with ORM or complex SQL. Automatic compiling from GraphQL to SQL.](https://www.reddit.com/r/golang/comments/fzj179/a_go_library_to_fetch_data_with_graphql_instead/)
- url: https://www.reddit.com/r/golang/comments/fzj179/a_go_library_to_fetch_data_with_graphql_instead/
---
Super Graph is now a GO library that you can use in your own GO code. Write your queries, inserts, updates in GraphQL instead of struggling with ORM's and complex SQL. No matter how complex your query, it will always be converted into a single SQL query.

Tons of very useful features like full-text search, efficient cursor based pagination, deep nested queries and mutations, treat array columns as foreign keys, treat json / json-b columns as separate tables and much more.

You don't have to know the structure of your database or deal with complex joins and json mangling just to efficiently fetch related data. Just use GraphQL to ask for what you want in the structure you want it in. 

https://github.com/dosco/super-graph

https://github.com/dosco/super-graph/issues/26

The API is pretty simple too.

```golang
package main

import (
  "database/sql"
  "fmt"
  "time"
  "github.com/dosco/super-graph/core"
  _ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
  db, err := sql.Open("pgx", "postgres://postgrs:@localhost:5432/example_db")
  if err != nil {
    log.Fatalf(err)
  }

  conf, err := core.ReadInConfig("./config/dev.yml")
  if err != nil {
    log.Fatalf(err)
  }

  sg, err = core.NewSuperGraph(conf, db)
  if err != nil {
    log.Fatalf(err)
  }

  query := `
    query {
      posts {
      id
      title
    }
  }`

  res, err := sg.GraphQL(context.Background(), query, nil)
  if err != nil {
    log.Fatalf(err)
  }

  fmt.Println(string(res.Data))
}
```

And your GraphQL query can be as complex as needed.

```graphql
query  {
  thread {
    slug
    title
    published
    createdAt : created_at
    totalVotes : cached_votes_total
    totalPosts : cached_posts_total
    vote : thread_vote(where: { user_id: { eq: $user_id } }) {
      created_at
    }
    topics {
      slug
      name
    }
    author : me {
      slug
    }
    posts(first: 1, order_by: { score: desc }) {
      slug
      body
      published
      createdAt : created_at
      totalVotes : cached_votes_total
      totalComments : cached_comments_total
      vote {
        created_at
      }
      author : user {
        slug 
        firstName : first_name
        lastName : last_name
      }
    }
    posts_cursor
  }
}
```
## [10][How to show error in frontend using “html/template”](https://www.reddit.com/r/golang/comments/fzluck/how_to_show_error_in_frontend_using_htmltemplate/)
- url: https://www.reddit.com/r/golang/comments/fzluck/how_to_show_error_in_frontend_using_htmltemplate/
---
I am a building a frontend that has an HTML form using “html/template” package, so in form element, I have action="/" attribute where I am doing a post request and saving the form data. My saving in DB is working perfectly.  
I want to know how to show errors?  
I am using validator/v10 package to do validation, I am able to get errors in JSON response but not understanding how to render that on the front end  
[Here is the full question](https://forum.golangbridge.org/t/how-to-show-error-in-frontend-using-html-template/18351)
