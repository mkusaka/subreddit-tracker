# Kotlin
## [1][Understand Kotlin Collection Function Past Tense](https://www.reddit.com/r/Kotlin/comments/gpk2w7/understand_kotlin_collection_function_past_tense/)
- url: https://medium.com/@elye.project/understand-kotlin-collection-function-past-tense-59f592af9436?source=friends_link&amp;sk=e22ccd272ebd28ef6f417e7b455b5b4f
---

## [2][How to secure Ktor web app with Keycloak Jetty 9.x Adapters?](https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/)
- url: https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/
---
Hi all 

I would like to secure the Ktor webapp, that is run on top of Jetty server with [Keycloak Jetty 9.x Adapters][1]. 

Ktor provides a [hook][2] for Jetty server initialization and maybe it is the right place in integrate the Keycloak Jetty 9.x Adapters.  

How to integrate Keycloak Jetty 9.x Adapters into Ktor app?  

Thanks


  [1]: https://www.keycloak.org/docs/latest/securing_apps/index.html#_jetty9_adapter
  [2]: https://ktor.io/servers/configuration.html#jetty
## [3][Harmony - A process-safe SharedPreference library](https://www.reddit.com/r/Kotlin/comments/gpbmci/harmony_a_processsafe_sharedpreference_library/)
- url: /r/android_devs/comments/gpbkyp/harmony_a_processsafe_sharedpreference_library/
---

## [4][Kotlin native performance compared to JVM - Opinions wanted](https://www.reddit.com/r/Kotlin/comments/gpdnt7/kotlin_native_performance_compared_to_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/gpdnt7/kotlin_native_performance_compared_to_jvm/
---
So my experience with Kotlin native has been less than ideal, bit of a rant (sorry!)

Essentially, after 3 years of Kotlin, went into the native field, thinking i would get better performance or memory usage for large scale applications.  


Although the latter is true, especially with millions of objects, the same cannot be said about the former.

First example. Java's File class vs native code using C functions such as fread(). I did some benchmarks (Loading a 55MB text file and load all the lines into an array. JVM does this in 10ms, where native code does it in 2500ms. How is it that much slower!?

Second example, loading loads of classes. JVM does the whole thing about 1000x faster, when converting lines in arrays to objects using a parser. Admittedly, under memory pressure, Kotlin native does perform more consistently, where JVM will pause to do a GC sweep. But still, total performance difference is about 1000x!

Has anyone else had similar experiences? - Or do you reckon i am doing something wrong here?
## [5][Trying to understand something about Kotlin's Collections interface and Java](https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/)
- url: https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/
---
I'm using Data Binding on Android and I'm trying to understand why I have to import `java.util.List` into my data binding layouts even tho the rest of the code is using Kotlin:

 [https://imgur.com/a/jFiIfy9](https://imgur.com/a/jFiIfy9) 

I tried replacing this for `kotlin.collections.List` but the code doesn't compile. Online I read something about that the Kotlin collection interfaces "don't actually exist". Can someone shed some light on this?
## [6][Spring Framework returns 404 for some URLs even when there are controllers for them.](https://www.reddit.com/r/Kotlin/comments/gpgcsh/spring_framework_returns_404_for_some_urls_even/)
- url: https://www.reddit.com/r/Kotlin/comments/gpgcsh/spring_framework_returns_404_for_some_urls_even/
---
I two controllers for my site written in Spring Boot and Kotlin, `IndexController` and `AboutController`. Both with a single function with a `@GetMapping` annotation to `/` and `/about`, respectively. However, the index page works find but the about page returns a 404 error. Even when I move the about page controller method to the IndexController, it still doesn't work.

In another test project, I initially couldn't even get the index page to route correctly, but that problem randomly fixed itself for some reason.

What could be going on, and how do I fix it?
## [7][What are the best resources to learn Kotlin?](https://www.reddit.com/r/Kotlin/comments/gp5hul/what_are_the_best_resources_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gp5hul/what_are_the_best_resources_to_learn_kotlin/
---
Hi,

I'm looking for something that would let be up and running with Kotlin ASAP. I'm comfortable with Java (11) and Typescript and also done some things with Scala, so I feel that most of this language will feel familiar to me.

Anything that you would recommend? Thank you
## [8][I want to learn Android development with Kotlin.](https://www.reddit.com/r/Kotlin/comments/gpbler/i_want_to_learn_android_development_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gpbler/i_want_to_learn_android_development_with_kotlin/
---
I want to learn Android development using Kotlin. But there's a big issue. My laptop has only 2GB of RAM and i can't afford a new laptop right now. What should I do?
## [9][Why LiveData is the best solution (yet) for UI](https://www.reddit.com/r/Kotlin/comments/gp8atk/why_livedata_is_the_best_solution_yet_for_ui/)
- url: https://www.coroutinedispatcher.com/2020/05/why-livedata-is-best-solution-yet-for-ui.html
---

## [10][30 days of Kotlin - Seminar 2 - Decoding Kotlin: The Modern Way To Build on Android](https://www.reddit.com/r/Kotlin/comments/gp044m/30_days_of_kotlin_seminar_2_decoding_kotlin_the/)
- url: /r/androiddev/comments/gp038k/30_days_of_kotlin_seminar_2_decoding_kotlin_the/
---

