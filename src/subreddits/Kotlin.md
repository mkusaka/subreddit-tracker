# Kotlin
## [1][Kotlin Quarkus OAuth2 and security with Keycloak - Piotr's TechBlog](https://www.reddit.com/r/Kotlin/comments/itr232/kotlin_quarkus_oauth2_and_security_with_keycloak/)
- url: https://piotrminkowski.com/2020/09/16/quarkus-oauth2-and-security-with-keycloak/
---

## [2][Discover Zoe, the new Kafka CLI written in Kotlin and the story behind it : )](https://www.reddit.com/r/Kotlin/comments/itvh82/discover_zoe_the_new_kafka_cli_written_in_kotlin/)
- url: https://medium.com/@adevinta_techblog/zoe-the-kafka-cli-for-humans-3e01584d0d3f
---

## [3][Newbie Question: IDEA not catching all compile errors](https://www.reddit.com/r/Kotlin/comments/itfydp/newbie_question_idea_not_catching_all_compile/)
- url: https://www.reddit.com/r/Kotlin/comments/itfydp/newbie_question_idea_not_catching_all_compile/
---
I'm new to IDEA/Kotlin but not jetbrains.  I'm curious if this is how its  supposed to work or do I need to configure something.

Here's a test:

```kotlin
open class SomeSuper {
    val someprop = ""
}

class SomeSub: SomeSuper() {
    override val someprop = "someval"
}
```

```shell
$ kotlinc play.kt &amp;&amp; kotlin PlayKt

play.kt:44:5: error: 'someprop' in 'SomeSuper' is final and cannot be overridden
    override val someprop = "someval
```

Shouldn't the IDE catch this as I'm typing it?  I'm sure I'm overlooking something. Nothing is shown in "problems window".
## [4][Simple Kotlin Null Check for Multiple Mutable Variables](https://www.reddit.com/r/Kotlin/comments/it6ghq/simple_kotlin_null_check_for_multiple_mutable/)
- url: https://medium.com/mobile-app-development-publication/simple-kotlin-null-check-for-multiple-mutable-variables-b095f7ac9bf1?source=friends_link&amp;sk=c182bb5d89f7d10985392d8e36f4c228
---

## [5][Should I learn Dart/Flutter or Kotlin?](https://www.reddit.com/r/Kotlin/comments/it4lpf/should_i_learn_dartflutter_or_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/it4lpf/should_i_learn_dartflutter_or_kotlin/
---

## [6][Kotlin, Hibernate, MySQL Tutorial needed](https://www.reddit.com/r/Kotlin/comments/itbrf0/kotlin_hibernate_mysql_tutorial_needed/)
- url: https://www.reddit.com/r/Kotlin/comments/itbrf0/kotlin_hibernate_mysql_tutorial_needed/
---
Hey guys,

I need to set up a database connection for the project I'm working on, but I don't really know the tools I'm supposed to use, and apparently I'm not competent enough to find a tutorial that I can get to work. So I'm once again asking you for support.

I'm looking for a Hibernate Tutorial in Kotlin (not Java, found plenty of those, they didn't work). One that includes MySQL instad of the often used H2 embedded database. I'd prefer to stay clear of build tools like Maven since I don't know them either, but since it seems impossible to find any tutorial without one of those lately, I'd prefer to use Maven over SpringBoot or Gradle.

Please help!
## [7][shellin- a Kotlin library for launching processes](https://www.reddit.com/r/Kotlin/comments/isul3h/shellin_a_kotlin_library_for_launching_processes/)
- url: https://github.com/ScottPeterJohnson/shellin
---

## [8][Kotlin 1.4 Online Event – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/islbad/kotlin_14_online_event_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/09/kotlin-1-4-online-event/
---

## [9][What I Learnt from Benchmarking Http4k, Ktor (Kotlin) and Actix v2, v3 (Rust) Microservices](https://www.reddit.com/r/Kotlin/comments/isgvm6/what_i_learnt_from_benchmarking_http4k_ktor/)
- url: https://matej.laitl.cz/bench-rust-kotlin-microservices/
---

## [10][Where do you host your server kotlin app?](https://www.reddit.com/r/Kotlin/comments/isoc7f/where_do_you_host_your_server_kotlin_app/)
- url: https://www.reddit.com/r/Kotlin/comments/isoc7f/where_do_you_host_your_server_kotlin_app/
---
Hi all, the quick question is as the title says :) 

A bit more of context: 

I am planning on using Ktor for implementing a REST API to be consumed from a mobile native app, I may some basic web functionality, but the main project ia the REST API.

i did a quick search, and the obvious options are Heroku and Google Engine.

I also read that there is a way of deploying into AWS Lambda functions.

I expect to host some database (not necessarily SQL) and some static content (images and so on) but not much user generated content, so as today I won't have a big need on space/storage.

I am a developer, with some little knowledge of devops, and I would rather keep that way :) (do not want to periodically spend hours a week configuring or maintaining a server infrastructure) 

Can you share your scenarios? What type of app do you run, where do you host? how cheap/expensive is? And would you recommend it? 

It will be appreciated :) 

Thanks!
