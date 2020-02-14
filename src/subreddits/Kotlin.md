# Kotlin
## [1][Fun tip of the day, is your val in a class? Any descendants can make it mutable without reflection :)](https://www.reddit.com/r/Kotlin/comments/f3lyfo/fun_tip_of_the_day_is_your_val_in_a_class_any/)
- url: https://i.redd.it/29j056k34tg41.png
---

## [2][Contextual Overriding with Kotlin](https://www.reddit.com/r/Kotlin/comments/f3il6z/contextual_overriding_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/f3il6z/contextual_overriding_with_kotlin/
---
What do you think of this? Useful somehow?

https://link.medium.com/tqB7FBlz33
## [3][Framework for RESTfull API](https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/)
- url: https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/
---
For our college project we decided to use Kotlin, it meets our criteria and we wanted to learn a new language. We decided to make an API where our clients could connect to. As I did some research I came across some potential frameworks like: Ktor, Spring, Javalin, Spring boot and some more. It was unclear which framework would be best for our usecase. I think that Ktor and Spring would be a bit overkill since it also comes with web client "builder", or is using these framework a recommend approach? In the tutorial section of the Kotlin website there was a Web development tutorial which was using Spring Boot, is this the most recommended approach?  


My question is, which (micro) frameworks for Kotlin do you know/recommend for using a RESTfull API where different clients could use it? And why you recommend these frameworks. Thanks a lot :)
## [4][Creating a simple search server using Kotlin, Ktor, and my Kotlin client for Elasticsearch. This article is part of my effort to create a nice manual for my Kotlin client for Elasticsearch and shows off how easy it is to use co-routines, bulk indexing, and the ES search DSL from Kotlin.](https://www.reddit.com/r/Kotlin/comments/f3acxq/creating_a_simple_search_server_using_kotlin_ktor/)
- url: https://github.com/jillesvangurp/es-kotlin-wrapper-client/blob/master/manual/recipe-search-engine.md
---

## [5][Where can I find some challenges in kotlin that don't require algorithms?](https://www.reddit.com/r/Kotlin/comments/f3c0bo/where_can_i_find_some_challenges_in_kotlin_that/)
- url: https://www.reddit.com/r/Kotlin/comments/f3c0bo/where_can_i_find_some_challenges_in_kotlin_that/
---
Title basically. 

&amp;#x200B;

I was following this [https://beginnersbook.com/2019/03/kotlin-class-and-objects-oop/](https://beginnersbook.com/2019/03/kotlin-class-and-objects-oop/) tutorial about OOP.
## [6][Exploring Kotlin IR](https://www.reddit.com/r/Kotlin/comments/f29qrn/exploring_kotlin_ir/)
- url: https://medium.com/@bnorm/exploring-kotlin-ir-bed8df167c23
---

## [7][JVM target 13](https://www.reddit.com/r/Kotlin/comments/f2dh0y/jvm_target_13/)
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
## [8][The state of Kotlin for microservices](https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/)
- url: https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/
---
Hi all 

I need your advices and would like to know, if it is the right choice to do microservices in Kotlin.
Of course, Golang would be the better choice, but I do not like it because of boilerplate codes.
Rust seems to be interesting, but libraries are not stable yet.
One point, that I concern about it, is the big image size of JVM. But I think with GraalVM is good to go.

The question is, should I use Kotlin to build my microservices or Kotlin is only suitable for frontend programming? What is your opinion? 

Thanks
## [9][Kotlin Multiplatform interface implementation in JavaScript?](https://www.reddit.com/r/Kotlin/comments/f2bhka/kotlin_multiplatform_interface_implementation_in/)
- url: https://www.reddit.com/r/Kotlin/comments/f2bhka/kotlin_multiplatform_interface_implementation_in/
---
Suppose I have a Kotlin Multiplatform project with a method, which takes an interface as an argument like this.

`package com.example`  


`interface MyInterface {`  
 `operator fun invoke()`  
`}`  


`class MyClass {`  
 `fun execute(client: MyInterface) {`  
`client()`  
`}`  
`}`

&amp;#x200B;

In Java I could just call the execute() method like this:

`void test() {`  
`MyClass myClass = new MyClass();`  
`myClass.execute(() -&gt; {`  
`//Do stuff`  
`});`  
`}`

Now in my Multiplatform project I have generated the JS file, but how do I call it in JavaScript, since JS doesn't have interfaces? I am trying to do this in Node.js.
## [10][Getting Started with Kotlin on iOS, Part 1](https://www.reddit.com/r/Kotlin/comments/f1uhwh/getting_started_with_kotlin_on_ios_part_1/)
- url: https://benasher.co/kotlin-ios-getting-started/
---

