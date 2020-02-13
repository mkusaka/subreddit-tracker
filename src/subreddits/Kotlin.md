# Kotlin
## [1][Framework for RESTfull API](https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/)
- url: https://www.reddit.com/r/Kotlin/comments/f37yl2/framework_for_restfull_api/
---
For our college project we decided to use Kotlin, it meets our criteria and we wanted to learn a new language. We decided to make an API where our clients could connect to. As I did some research I came across some potential frameworks like: Ktor, Spring, Javalin, Spring boot and some more. It was unclear which framework would be best for our usecase. I think that Ktor and Spring would be a bit overkill since it also comes with web client "builder", or is using these framework a recommend approach? In the tutorial section of the Kotlin website there was a Web development tutorial which was using Spring Boot, is this the most recommended approach?  


My question is, which (micro) frameworks for Kotlin do you know/recommend for using a RESTfull API where different clients could use it? And why you recommend these frameworks. Thanks a lot :)
## [2][Exploring Kotlin IR](https://www.reddit.com/r/Kotlin/comments/f29qrn/exploring_kotlin_ir/)
- url: https://medium.com/@bnorm/exploring-kotlin-ir-bed8df167c23
---

## [3][JVM target 13](https://www.reddit.com/r/Kotlin/comments/f2dh0y/jvm_target_13/)
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
## [4][The state of Kotlin for microservices](https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/)
- url: https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/
---
Hi all 

I need your advices and would like to know, if it is the right choice to do microservices in Kotlin.
Of course, Golang would be the better choice, but I do not like it because of boilerplate codes.
Rust seems to be interesting, but libraries are not stable yet.
One point, that I concern about it, is the big image size of JVM. But I think with GraalVM is good to go.

The question is, should I use Kotlin to build my microservices or Kotlin is only suitable for frontend programming? What is your opinion? 

Thanks
## [5][Kotlin Multiplatform interface implementation in JavaScript?](https://www.reddit.com/r/Kotlin/comments/f2bhka/kotlin_multiplatform_interface_implementation_in/)
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
## [6][Getting Started with Kotlin on iOS, Part 1](https://www.reddit.com/r/Kotlin/comments/f1uhwh/getting_started_with_kotlin_on_ios_part_1/)
- url: https://benasher.co/kotlin-ios-getting-started/
---

## [7][does anyone know if kotlin/native already support curses/ncurses on linux?](https://www.reddit.com/r/Kotlin/comments/f1pmwt/does_anyone_know_if_kotlinnative_already_support/)
- url: https://www.reddit.com/r/Kotlin/comments/f1pmwt/does_anyone_know_if_kotlinnative_already_support/
---

## [8][Accelerate Your Kotlin Multiplatform Evaluation with KaMP Kit](https://www.reddit.com/r/Kotlin/comments/f1p6nq/accelerate_your_kotlin_multiplatform_evaluation/)
- url: https://blog.jetbrains.com/kotlin/2020/02/accelerate-your-kotlin-multiplatform-evaluation-with-kamp-kit/
---

## [9][Run code with "delay" in a browser](https://www.reddit.com/r/Kotlin/comments/f1c31k/run_code_with_delay_in_a_browser/)
- url: https://www.reddit.com/r/Kotlin/comments/f1c31k/run_code_with_delay_in_a_browser/
---
Hi everyone! I was looking at code examples here [https://kotlinlang.org/docs/tutorials/coroutines/coroutines-basic-jvm.html](https://kotlinlang.org/docs/tutorials/coroutines/coroutines-basic-jvm.html)

and you can notice that if we'll run this code we won't see delays between printing lines, because the code sends to the server and executes there, that's why we just receive the result.

https://preview.redd.it/581tl4k9pxf41.png?width=752&amp;format=png&amp;auto=webp&amp;s=5a517c804ff7bfe810ad57b905ca16651148c4f2

Is it possible to execute such codes somehow so we'll be able to see delays? (For example, after pushing Run button there appears a window just like on the screen, but "Start", "Hello" and "Stop" appear with the 1 and 2 seconds intervals) Maybe we can convert kotlin to JavaScript and run it in the browser? I'm only studying and not a professional so I'm not even sure if it's possible. I tried to google this problem but didn't succeed. It would be great if you can give me advice or suggestion about how this problem can possibly be solved or where to find the information.
## [10][abstract fun](https://www.reddit.com/r/Kotlin/comments/f1cait/abstract_fun/)
- url: https://www.reddit.com/r/Kotlin/comments/f1cait/abstract_fun/
---
What is the difference between using this:

`abstract val sleepDatabaseDao: SleepDatabaseDao`

vs using this:

`abstract fun chapterDao(): ChapterDao`

Cant I in both cases as the methods inside the DAO interface?
