# Kotlin
## [1][What's the best general purpose array class in Kotlin?](https://www.reddit.com/r/Kotlin/comments/hbqxne/whats_the_best_general_purpose_array_class_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hbqxne/whats_the_best_general_purpose_array_class_in/
---
When I took an intro to Java class, if we need to a basic array object, the majority of time we used an `ArrayList`. It was never really explained, and they said we could any suitable class we wanted in our assignments, but most of the code given to us used `ArrayList`.

Obviously, I can still use `ArrayList` in Kotlin though the JVM, but is there a different class that fulfils a general purpose array?
## [2][A Kotlin/JS Boids simulation using three.js](https://www.reddit.com/r/Kotlin/comments/hbolen/a_kotlinjs_boids_simulation_using_threejs/)
- url: https://www.reddit.com/r/Kotlin/comments/hbolen/a_kotlinjs_boids_simulation_using_threejs/
---
[https://github.com/liorgonnen/boids.kt](https://github.com/liorgonnen/boids.kt)

[Kotlin in the Browser! 🤘🏼](https://preview.redd.it/lcapa6pdvq551.jpg?width=1280&amp;format=pjpg&amp;auto=webp&amp;s=3a313722d022a4dbd70c7e2c814ee00540a87a8a)

After many years of Android development, I'm betting on the Web as the best platform for indie developers. 

I've always preferred strongly typed languages, and I absolutely love Kotlin. I wanted to start exploring the Kotlin developer experience for the browser.

Using Dukat to generate the Kotlin externals for three.js wasn't working out of the box for three.js, so I ended up modifying the tool to support this use-case. Other than that, things are a pretty smooth (and fun!) ride.

Why Boids?

Just seemed like a fun and satisfying use-case :)
## [3][[Kord] - Discord Kotlin Library](https://www.reddit.com/r/Kotlin/comments/hbaeh8/kord_discord_kotlin_library/)
- url: https://www.reddit.com/r/Kotlin/comments/hbaeh8/kord_discord_kotlin_library/
---
Kord was created as an answer to the frustrations of writing Discord bots with other JVM libraries, which either use thread-blocking code or verbose and scope restrictive reactive systems. We believe an API written from the ground up in Kotlin with coroutines can give you the best of both worlds: The conciseness of imperative code with the concurrency of reactive code.

Aside from coroutines, we also wanted to give the user full access to lower-level APIs. Sometimes you have to do some unconventional things, and we want to allow you to do those in a safe and supported way.

* Discord Gateway
* Discord Rest API
* High-level abstraction and caching

Voice and Multiplatform support coming in future releases.

you can find all the things you need to get started at [GitHub](https://github.com/kordlib/kord) and our [Discord](https://discord.gg/mpDQm5N)  


&amp;#x200B;
## [4][Kotlin gives me kotlin.unit instead of correct string](https://www.reddit.com/r/Kotlin/comments/hbig5w/kotlin_gives_me_kotlinunit_instead_of_correct/)
- url: https://www.reddit.com/r/Kotlin/comments/hbig5w/kotlin_gives_me_kotlinunit_instead_of_correct/
---
I found a weird issue for my following code. The value is 3, isDone is false, so I should get "3 scans", but it prints as "3 kotlin.unit". What is the reason of it? The code was auto suggested by Android Studio.

`string += if(isDone) {`  
 `"scanned"`  
`} else {`  
 `if(value == null || value == 0) {`  
 `"scan"`  
 `} else {`  
 `"scans"`  
 `}`  
`}`

Dear mod, please remove this post if it is not allowed to ask programming question.
## [5][My first Kotlin Multiplatform app in the app store](https://www.reddit.com/r/Kotlin/comments/haz2vj/my_first_kotlin_multiplatform_app_in_the_app_store/)
- url: https://www.reddit.com/r/Kotlin/comments/haz2vj/my_first_kotlin_multiplatform_app_in_the_app_store/
---
My side project Packrat, built with Kotlin Multiplatform in order to learn a bit about Kotlin is now live in the app store. https://apps.apple.com/us/app/packrat/id1490422067?ls=1. I started with iOS as I knew it better but a lot of my app already works on Android too. Overall I really enjoyed starting to learn Kotlin as a language, at first it seemed similar to Swift (which I love) but as I dived further in I realised this was only superficial, I found it to be it's own unique and productive language with some very cool design choices. 

Using Swift and Kotlin together worked really nicely and I didn't feel any issues with context switching (apart from the occasional fun/func annoyance). I believe Kotlin Multiplatform has a bright future and I hope to use it for a client one day soon.
## [6][How to explain implementation in Kotlin](https://www.reddit.com/r/Kotlin/comments/hbe3u5/how_to_explain_implementation_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hbe3u5/how_to_explain_implementation_in_kotlin/
---
What is the meaning of this code line? 

 fun  a.b(y: Int) = lazy { u.v(y) }
## [7][Status and future of TornadoFX?](https://www.reddit.com/r/Kotlin/comments/hasb8y/status_and_future_of_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/hasb8y/status_and_future_of_tornadofx/
---
Hey Guys, 

can you help me understand, what the status and the future of TornadoFX is? On the github page it still says that it is not (yet) compatible with java 9/10. I see that the releases are quite far apart.

Don't get me wrong, i love TornadoFX. I think the way stuff is done is very cool and I have used it in a couple of projects already. 

I am just starting a new project and wanna use TornadoFX with this too, but i need to make sure it is the right lib to use for Kotlin and JavaFX. I just wanna ask if it is still being actively developed, evolving and future proof.

Any comments and hints would be appreciated! :)
## [8][Published my first app to the play store today 😁](https://www.reddit.com/r/Kotlin/comments/haoxom/published_my_first_app_to_the_play_store_today/)
- url: https://www.reddit.com/r/Kotlin/comments/haoxom/published_my_first_app_to_the_play_store_today/
---

I started learning kotlin for Android app development a few weeks ago and I’m loving it


Published my first android app to the play store using Kotlin. Yes I know it’s a ***VERY basic app*** However I’m learning new stuff everyday.

The ***primary focus*** of making this simple app was just to understand the fundamentals of kotlin,  how to use android studio and how to implement ads using admob.

Check it out if you want to see what a terrible app this is lol. The main thing is that I’m learning anyway.

https://play.google.com/store/apps/details?id=com.herd.whattodo
## [9][Zoe: a new CLI tool for Apache Kafka written in Kotlin](https://www.reddit.com/r/Kotlin/comments/hau0b2/zoe_a_new_cli_tool_for_apache_kafka_written_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hau0b2/zoe_a_new_cli_tool_for_apache_kafka_written_in/
---
Hi!

Within Adevinta, we are heavy users of Apache Kafka. We are also Kotlin Lovers! Recently, we open sourced a tool called Zoe that makes interacting with Kafka much easier. This tool is written in Kotlin.

* Checkout [the repository here](https://github.com/adevinta/zoe) where you will find a screen cast that demo the tool.
* And [the documentation](https://adevinta.github.io/zoe/)

Any feedback is welcome : )
## [10][Elide - JSON:API or GraphQL web service starting from a JPA annotated data model (by Yahoo!)](https://www.reddit.com/r/Kotlin/comments/haq2ht/elide_jsonapi_or_graphql_web_service_starting/)
- url: https://blog.graphqleditor.com/elide-opinionated-apis/
---

