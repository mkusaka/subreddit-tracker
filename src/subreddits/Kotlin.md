# Kotlin
## [1][A sample app built with Kotlin, Coroutines, Flow, Dagger 2, Architecture Components, MVVM, Room, Retrofit, Moshi, Material Components and other Jetpack libraries.](https://www.reddit.com/r/Kotlin/comments/fm91pb/a_sample_app_built_with_kotlin_coroutines_flow/)
- url: https://github.com/PatilShreyas/Foodium
---

## [2][KotlinConf 2020 has been postponed until 2021](https://www.reddit.com/r/Kotlin/comments/fluber/kotlinconf_2020_has_been_postponed_until_2021/)
- url: https://kotlinconf.com/?2020Postponed
---

## [3][Why does Kotlin have Unit when Nothing? could have done the same job?](https://www.reddit.com/r/Kotlin/comments/fm3hke/why_does_kotlin_have_unit_when_nothing_could_have/)
- url: https://www.reddit.com/r/Kotlin/comments/fm3hke/why_does_kotlin_have_unit_when_nothing_could_have/
---
I understand that `Unit` exists to be a lack of a useful value, i.e., what you return from a function that just performs a side effect. I'm wondering why the language designers decided that `Unit` was necessary when it seems to me that `Nothing?` would have worked just as well, all without introducing a new type.

    // Instead of this:
    fun doSomething() {  // implicit Unit return type
        doStuff()
        // implicitly return Unit
    }
    
    // Why not this?
    fun doSomething() {  // implicit Nothing? return type
        doStuff()
        // implicitly return null
    }
## [4][Are there any examples of websites built with kotlin/kotlinJS/react?](https://www.reddit.com/r/Kotlin/comments/fm3qa3/are_there_any_examples_of_websites_built_with/)
- url: https://www.reddit.com/r/Kotlin/comments/fm3qa3/are_there_any_examples_of_websites_built_with/
---
I know that react supports kotlin now.  There are a number of videos showing web applications being built with kotlin-react.  But are there any full websites made with kotlin that you can point me to?

Thanks in advance!
## [5][Help needed with YAML parsing and accessing specific keys(?)](https://www.reddit.com/r/Kotlin/comments/fm2jat/help_needed_with_yaml_parsing_and_accessing/)
- url: https://www.reddit.com/r/Kotlin/comments/fm2jat/help_needed_with_yaml_parsing_and_accessing/
---
Now I'm sure that there's something out there for this, but I haven't found it yet. I need a YAML parsing library that lets me access the keys(?) (I dunno the terminology, here) in a file, using the dot notation (like how we do in Spring Boot).

For example, say I have a YAML file like this . . .

```
some-property: C:\some\folder\path

root:
  foo: sample text
  bar: 1234
  baz: some@random.text
```

. . . and let's say that I wanna access ```"root.foo"``` . . . **is there a way to do so _without_ having to create nested maps or creating DTOs or (basically) data classes?**
## [6][TornadoFX Help](https://www.reddit.com/r/Kotlin/comments/fm0acd/tornadofx_help/)
- url: https://www.reddit.com/r/Kotlin/comments/fm0acd/tornadofx_help/
---
Hey all, I am trying to set up a tornadoFX project. I have openjdk 8, and I need to get javafx in there somewhere, preferably through gradle. Whats the best way to do this? Thanks!
## [7][I want to do addition function. What's wrong here ?](https://www.reddit.com/r/Kotlin/comments/flxdu8/i_want_to_do_addition_function_whats_wrong_here/)
- url: https://www.reddit.com/r/Kotlin/comments/flxdu8/i_want_to_do_addition_function_whats_wrong_here/
---
`fun addition(x,y){`  
`var z = x+y`  
`print(z)`  
`}`  


`fun main() {`  
 `addition(5,5)`  
`}`
## [8][Beginners question: How do I include an external library?](https://www.reddit.com/r/Kotlin/comments/flwd1u/beginners_question_how_do_i_include_an_external/)
- url: https://www.reddit.com/r/Kotlin/comments/flwd1u/beginners_question_how_do_i_include_an_external/
---
I'm trying to add the Exposed library to a project, but there seems to be no way to add an external library to a Kotlin project. 

I've added this to my gradle.build:
dependencies {
  compile("org.jetbrains.exposed", "exposed-core", "0.22.1")
  compile("org.jetbrains.exposed", "exposed-dao", "0.22.1")
  compile("org.jetbrains.exposed", "exposed-jdbc", "0.22.1")
}

and this to my class:
import org.jetbrains.exposed.sql.*
import org.jetbrains.exposed.sql.transactions.transaction

But it just refuses to see the Exposed library. What am I missing?
## [9][Announcing GraphQLize - an open-source JVM library for developing GraphQL API instantly from MySQL &amp; Postgres](https://www.reddit.com/r/Kotlin/comments/flw2sz/announcing_graphqlize_an_opensource_jvm_library/)
- url: https://www.graphqlize.org/blog/announcing-graphqlize-alpha
---

## [10][Unit Testing Delays, Errors &amp; Retries with Kotlin Flows](https://www.reddit.com/r/Kotlin/comments/fll2a7/unit_testing_delays_errors_retries_with_kotlin/)
- url: https://medium.com/@heyitsmohit/unit-testing-delays-errors-retries-with-kotlin-flows-77ce00d0c2f3
---

