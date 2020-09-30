# golang
## [1][Show us your static site generator!](https://www.reddit.com/r/golang/comments/j2gbgi/show_us_your_static_site_generator/)
- url: https://www.reddit.com/r/golang/comments/j2gbgi/show_us_your_static_site_generator/
---
Inspired by a recent post, I thought it could be fun to show each other the static site generators we've written in Go.

They've become the new graphing library, or config parser, meaning that a lot of developers seem to want to write one (especially _just_ before they start blogging again) and put it out into the world.

I think the reasons that so many exist are manyfold, but I see some trends:

* The drive to create something and put it out into the world (who knows, maybe you'll be internet-famous!)
* They're within a reasonably small scope, at least at first
* Everyone with a static website could use one
* Not-invented-here-syndrome, because Hugo is missing _that one feature_ you need, and therefore Hugo suxxxxx.

I also think Go lends itself beautifully to writing static site generators, because it has templating built into the stdlib, compiles down to a single binary, and it's easy to write something that's fast (so you can create comparison tables!).

If you think this is fun:
1) Link to your static site generator.
2) Post a few sentences on why you wrote it.

Happy posting!
## [2][Resetting/Re-Using timeout context](https://www.reddit.com/r/golang/comments/j2c7v7/resettingreusing_timeout_context/)
- url: https://play.golang.org/p/6YDCw30J6ok
---

## [3][sq: A type-safe SQL query builder and struct mapper](https://www.reddit.com/r/golang/comments/j2irku/sq_a_typesafe_sql_query_builder_and_struct_mapper/)
- url: https://www.reddit.com/r/golang/comments/j2irku/sq_a_typesafe_sql_query_builder_and_struct_mapper/
---
Repo link: [https://github.com/bokwoon95/go-structured-query](https://github.com/bokwoon95/go-structured-query)

Hi r/golang, sq is a type-safe query builder and struct mapper that is designed as an alternative to [sqlx](https://github.com/jmoiron/sqlx). Here is what it brings over sqlx:

- **sqlx's string-based struct annotations are type-unsafe and susceptible to typos.** sq uses types definitions generated from your database to ensure that whatever columns you are mapping to your struct fields actually exist in the table, and that the data type of your column matches the data type of the Go variable you are mapping it to.
    - sq does this with a callback mapper function, and so does not use any reflect-based struct annotation for struct mapping at all. You define the mapper function, and sq calls your mapper function.

- **sqlx only automates the mapping of columns to struct fields. You still have to SELECT the columns yourself.** If your SELECTed columns are out of sync with sqlx's struct annotations, good luck. On the other hand sq's mapper functions serve the purpose of both SELECT-ing the columns *and* mapping them back to the struct fields.
    - This means you can declare a mapper function on a User struct and never have to write the boilerplate of `SELECT user.user_id, user.name, user.email... etc` ever again because the mapper function does it for you.

- **sqlx does not handle nested structs of the same type: any nested struct type must exist at most once inside a parent struct in order for StructScan to work.** sq's mapper function does not have such a limitation because it lets you map columns to anything. Do you want to map columns to two different structs? You can. Do you want to map columns to ordinary Go variables without needing to instantiate a struct? You can.
    - If you are interested in how nested struct mapping works, I recently gave a [presentation on sq](https://www.youtube.com/watch?v=6DL7LU4ki2U&amp;feature=youtu.be&amp;t=4027) which involves a nested struct example.

- **sqlx falls back on using sql.NullXXX to do NULL handling, which pollutes your domain structs.** If you do not want your domain structs to be peppered with sql.NullXXX, you have to create intermediate structs to handle the possible NULL values, then transfer them over to the domain struct. 
    - sq offers the convenience of implicitly scanning NULL values as zero values *while* still offering a way to check for the actual NULLness of a column.

sq is a query builder. How does it differ from other query builders e.g. [squirrel](https://github.com/Masterminds/squirrel)? It offers you type-safety (due to code-generated table definitions) and also provides struct mapping facilities. Most query builders out there just build query strings and hand off the responsibility of scanning results back to you. It also provides the most SQL features I've seen in any query builder (check out the list of features [here](https://bokwoon95.github.io/sq/#query-building) and [here](https://bokwoon95.github.io/sq/#sql-expressions)).

[Example queries](https://github.com/bokwoon95/go-structured-query#examples).

I have just updated the version to v1.0. Do check if out and let me know what you think.
## [4][ent + gqlgen = &lt;3](https://www.reddit.com/r/golang/comments/j2ljgh/ent_gqlgen_3/)
- url: https://www.reddit.com/r/golang/comments/j2ljgh/ent_gqlgen_3/
---
Hey guys,
A few weeks ago, I mentioned here that we'll OSS our GraphQL integration for [ent](https://github.com/facebook/ent). So, as promised, here's the link to the docs [entgo.io/docs/graphql](https://entgo.io/docs/graphql/), and to the [entgql extension](https://github.com/facebookincubator/ent-contrib) in GitHub. 


What exactly we currently provide in this integration: 

- [Node API](https://entgo.io/docs/graphql/#node-api) support. 

- [Pagination support](https://entgo.io/docs/graphql/#pagination) - Follows the *Relay Cursor Connections Spec*.

- [Connection Ordering](https://entgo.io/docs/graphql/#connection-ordering).

- [Fields Collection](https://entgo.io/docs/graphql/#fields-collection) - (*This one is super cool!*) Auto eager-load relations/edges that exist in the GraphQL request.

- [Transactional Mutations](https://entgo.io/docs/graphql/#transactional-mutations) - A handler that executes each GraphQL mutation in a transaction and either commit or rollback the transaction based on the response. 

- And additional small features that makes the binding between the 2 libraries much friendly. See the docs for more details.
 
We are actively working to improve the integration between the 2 libraries and we expect to OSS additional options in the future. If you're using gqlgen and ent, please give it a look/try and let me know what do you think :)
## [5][Wrote an article/tutorial on Concurrency. Beginner Friendly!](https://www.reddit.com/r/golang/comments/j1xij2/wrote_an_articletutorial_on_concurrency_beginner/)
- url: https://medium.com/@yashaswi_nayak/go-a-tale-of-concurrency-a-beginners-guide-b8976b26feb
---

## [6][What are the most useful functions / methods etc. to know when learning Go as a second language?](https://www.reddit.com/r/golang/comments/j29uph/what_are_the_most_useful_functions_methods_etc_to/)
- url: https://www.reddit.com/r/golang/comments/j29uph/what_are_the_most_useful_functions_methods_etc_to/
---
In the past few days, I decided I want to learn how to write my backend in Go instead of node (just for fun). Coming from js, picking up the basics of arrays, slices, structs etc. hasn’t been hard, so i’m not looking for ULTRA-ULTRA basic stuff. 


I have run into some useful functions, such as strings.TrimSpace() and fmt.Sprint(x)


I can see myself using these often. 


What are some of the most common funcs I should be aware of? The standard library looks robust, and it’s a lot to go through at first!
## [7][A question around social movements and the Golang community](https://www.reddit.com/r/golang/comments/j2llih/a_question_around_social_movements_and_the_golang/)
- url: https://www.reddit.com/r/golang/comments/j2llih/a_question_around_social_movements_and_the_golang/
---
Throwaway because I don't know how to word an innocent question in the current social/political environment without someone labelling or  attacking me.

I went to check the Go standard library and noticed a banner at the top of the screen. At first I thought it was another "accept cookies" overlay, but instead it is [promoting a current social movement](https://imgur.com/a/bMJD2FE). I wanted to know if anyone else finds this odd, or disagrees with banners like this?

I want to be deliberately anonymous about whether I agree or disagree with \_this particular social cause\_, because my question is an abstract one about \_any social cause\_ being promoted this way. I live in an English speaking country, and it's \*not\* the United States. The Golang community is touted as global, and this feels like US-centric issue. There are social movements I support strongly, both in my country and globally, but I do not expect them to be advertised in a place like software language documentation.

I mostly agree with the idea of [The Go Code of Conduct](https://golang.org/conduct), and I want to work in a community that is welcoming and accepting of people. But just like I personally prefer separation of state and religion, I think I prefer separation of software engineering technical discussion and social cause(s).

&amp;#x200B;

Can I ask this banner is removed without it being seen as a statement that I am against this cause?
## [8][Empty slice vs nil slice in GoLang](https://www.reddit.com/r/golang/comments/j2lkzm/empty_slice_vs_nil_slice_in_golang/)
- url: https://www.pixelstech.net/article/1539870875-Empty-slice-vs-nil-slice-in-GoLang
---

## [9][A new Kafka Go client library](https://www.reddit.com/r/golang/comments/j28vw5/a_new_kafka_go_client_library/)
- url: /r/apachekafka/comments/j28vos/a_new_kafka_go_client_library/
---

## [10][Billion-Dollar Mistake in Go?](https://www.reddit.com/r/golang/comments/j2jm62/billiondollar_mistake_in_go/)
- url: https://hackernoon.com/billion-dollar-mistake-in-go-ll1s3tkc?source=rss
---

