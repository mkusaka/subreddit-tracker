# Kotlin
## [1][Is there a way to make a data class open?](https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/)
- url: https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/
---
I'm trying to use the ModelMapper library with data classes. To build an explicit type map it needs to proxy my class, and data classes cannot be extended and therefore cannot be proxied. Is there a way around this? Some compiler setting? I've tried all-open but I can't seem to get it to work.

Thanks.
## [2][How to run ktor application?](https://www.reddit.com/r/Kotlin/comments/f47a4r/how_to_run_ktor_application/)
- url: https://stackoverflow.com/questions/60233921/how-to-run-ktor-application/60237120#60237120
---

## [3][Fun tip of the day, is your val in a class? Any descendants can make it mutable without reflection :)](https://www.reddit.com/r/Kotlin/comments/f3lyfo/fun_tip_of_the_day_is_your_val_in_a_class_any/)
- url: https://i.redd.it/29j056k34tg41.png
---

## [4][Contextual Overriding with Kotlin](https://www.reddit.com/r/Kotlin/comments/f3il6z/contextual_overriding_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/f3il6z/contextual_overriding_with_kotlin/
---
What do you think of this? Useful somehow?

https://link.medium.com/tqB7FBlz33
## [5][Framework for RESTfull API](https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/)
- url: https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/
---
For our college project we decided to use Kotlin, it meets our criteria and we wanted to learn a new language. We decided to make an API where our clients could connect to. As I did some research I came across some potential frameworks like: Ktor, Spring, Javalin, Spring boot and some more. It was unclear which framework would be best for our usecase. I think that Ktor and Spring would be a bit overkill since it also comes with web client "builder", or is using these framework a recommend approach? In the tutorial section of the Kotlin website there was a Web development tutorial which was using Spring Boot, is this the most recommended approach?  


My question is, which (micro) frameworks for Kotlin do you know/recommend for using a RESTfull API where different clients could use it? And why you recommend these frameworks. Thanks a lot :)
## [6][Creating a simple search server using Kotlin, Ktor, and my Kotlin client for Elasticsearch. This article is part of my effort to create a nice manual for my Kotlin client for Elasticsearch and shows off how easy it is to use co-routines, bulk indexing, and the ES search DSL from Kotlin.](https://www.reddit.com/r/Kotlin/comments/f3acxq/creating_a_simple_search_server_using_kotlin_ktor/)
- url: https://github.com/jillesvangurp/es-kotlin-wrapper-client/blob/master/manual/recipe-search-engine.md
---

## [7][Where can I find some challenges in kotlin that don't require algorithms?](https://www.reddit.com/r/Kotlin/comments/f3c0bo/where_can_i_find_some_challenges_in_kotlin_that/)
- url: https://www.reddit.com/r/Kotlin/comments/f3c0bo/where_can_i_find_some_challenges_in_kotlin_that/
---
Title basically. 

&amp;#x200B;

I was following this [https://beginnersbook.com/2019/03/kotlin-class-and-objects-oop/](https://beginnersbook.com/2019/03/kotlin-class-and-objects-oop/) tutorial about OOP.
## [8][Exploring Kotlin IR](https://www.reddit.com/r/Kotlin/comments/f29qrn/exploring_kotlin_ir/)
- url: https://medium.com/@bnorm/exploring-kotlin-ir-bed8df167c23
---

## [9][JVM target 13](https://www.reddit.com/r/Kotlin/comments/f2dh0y/jvm_target_13/)
- url: https://www.reddit.com/r/Kotlin/comments/f2dh0y/jvm_target_13/
---
Hi all

How to change the target compiler to JDK 13? At the moment, it is defined as follow:
```
plugins {
    kotlin("jvm") version "1.3.61"
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
}

tasks {
    compileKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
}
``` 

The following version is installed on my computer:
```
java -version
openjdk version "13.0.2" 2020-01-14
OpenJDK Runtime Environment Zulu13.29+9-CA (build 13.0.2+6-MTS)
OpenJDK 64-Bit Server VM Zulu13.29+9-CA (build 13.0.2+6-MTS, mixed mode, sharing)
```

Thanks
## [10][The state of Kotlin for microservices](https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/)
- url: https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/
---
Hi all 

I need your advices and would like to know, if it is the right choice to do microservices in Kotlin.
Of course, Golang would be the better choice, but I do not like it because of boilerplate codes.
Rust seems to be interesting, but libraries are not stable yet.
One point, that I concern about it, is the big image size of JVM. But I think with GraalVM is good to go.

The question is, should I use Kotlin to build my microservices or Kotlin is only suitable for frontend programming? What is your opinion? 

Thanks
