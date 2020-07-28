# Kotlin
## [1][Zephyria - Kotlin RPG](https://www.reddit.com/r/Kotlin/comments/hzbupy/zephyria_kotlin_rpg/)
- url: https://www.reddit.com/r/Kotlin/comments/hzbupy/zephyria_kotlin_rpg/
---
Hi,

This is a post to make those interested in Kotlin and games aware of the Zephyria RPG. Here's a 4-minute gameplay demo: [https://www.youtube.com/watch?v=LIV8UJAokA0](https://www.youtube.com/watch?v=LIV8UJAokA0)

The game is originally a testbed (sandboxed development area so to speak) for the FXGL engine. As it currently stands, the game can be compiled and run. Whilst some actions, such as movement, are fully implemented, other actions, such as combat, might be glitchy, whereas talking to NPCs will outright crash the game.

The quality of code also varies. The newly added code uses latest FXGL features and is concise. The leftover legacy code is messy in places with a few dozen "TODO"s.

Contributions *are* welcome, however further refactoring and code cleanup are needed from my side to make the repository contributor-friendly. Once it is ready, there will be a further post to call for contributions.

I hope this is a useful addition to the Kotlin game development community.

Thank you!

Game sources: [https://github.com/AlmasB/Zephyria](https://github.com/AlmasB/Zephyria) 

Engine sources: [https://github.com/AlmasB/FXGL](https://github.com/AlmasB/FXGL)

[Screenshot](https://preview.redd.it/vlcl5ziqlkd51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=08737f1fd9b3f92bd7577f126c64437fe0703786)
## [2][Kotlin Cheat Sheet](https://www.reddit.com/r/Kotlin/comments/hzdk56/kotlin_cheat_sheet/)
- url: https://i.redd.it/w4warklofid51.jpg
---

## [3][Kotlin 1.4.0-RC Released – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/hysqs8/kotlin_140rc_released_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-rc-released/
---

## [4][What's the best way to launch a small coroutine?](https://www.reddit.com/r/Kotlin/comments/hzd4u3/whats_the_best_way_to_launch_a_small_coroutine/)
- url: https://www.reddit.com/r/Kotlin/comments/hzd4u3/whats_the_best_way_to_launch_a_small_coroutine/
---
I think this is more related to Kotlin than Android, that's why I post here.

What's the best/fastest way to launch a small thread from a RecyclerView's Adapter that has no direct access to the ViewModel (to get the viewModelScope)? The coroutine is mostly used for choreographing animations.

`MainScopre().lauch { }` or should I pass the ViewModel somehow indirectly? 

Sorry in advance if it's a stupid question.
## [5][Ktor application using hexagonal architecture](https://www.reddit.com/r/Kotlin/comments/hz4d7a/ktor_application_using_hexagonal_architecture/)
- url: https://www.reddit.com/r/Kotlin/comments/hz4d7a/ktor_application_using_hexagonal_architecture/
---
Hi everyone!

I recently searched for some examples of backend apps using ktor to benchmark with spring framework.

I wrote this application to make an article in which I am working in my spare time with a comparison between the two frameworks and the experience of working with a framework that brings out the best in language.  


if you want to take a look in a backend application using kafka, ktor, exposed and other stuffs, check it out:

[https://github.com/jgabrielfreitas/ktor-hexagonal-benchmark](https://github.com/jgabrielfreitas/ktor-hexagonal-benchmark)

&amp;#x200B;

*feel free to give a feedback! :)*
## [6][After 3 years of Kotlin, I finally found an understandable explanation of "in" and "out"](https://www.reddit.com/r/Kotlin/comments/hyo565/after_3_years_of_kotlin_i_finally_found_an/)
- url: https://www.reddit.com/r/Kotlin/comments/hyo565/after_3_years_of_kotlin_i_finally_found_an/
---
Stumbled upon this article after yet again trying to understand "in" and "out" and how it correlates to Java. Might come in handy for someone else!  [https://kotlin.christmas/2019/22](https://kotlin.christmas/2019/22)
## [7][GitHub - JetBrains/skiko: Kotlin MPP bindings to Skia](https://www.reddit.com/r/Kotlin/comments/hywxuc/github_jetbrainsskiko_kotlin_mpp_bindings_to_skia/)
- url: https://github.com/JetBrains/skiko
---

## [8][What a beginner should do?](https://www.reddit.com/r/Kotlin/comments/hyuo37/what_a_beginner_should_do/)
- url: https://www.reddit.com/r/Kotlin/comments/hyuo37/what_a_beginner_should_do/
---
Hello! I am a complete beginner in coding. I woke up one day and said to myself: I shall learn a programming language.

It seems to me though that each documentation or course I find is for those that already know what they are doing. 

This text, for example, taken from the [official Kotlin introduction](https://play.kotlinlang.org/byExample/01_introduction/02_Functions):  


"A simple function that takes a parameter of type 

    String

and returns 

    Unit

(i.e., no return value). "

&amp;#x200B;

I have no idea what is a parameter of type String or why it returns Unit or why are they there. 

The code stills make no sense to me and I don't know where I should look up to understand those things or where I should start. Something that says "this is here because it does this and together they do this". I'm completely lost...

Perhaps I'm starting with the wrong language?
## [9][Is there a model-mapper equivalent native to Kotlin?](https://www.reddit.com/r/Kotlin/comments/hz10tx/is_there_a_modelmapper_equivalent_native_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hz10tx/is_there_a_modelmapper_equivalent_native_to_kotlin/
---
So I like the model-mapper library in Java. It's a library that, if given two Java classes whose fields have the same names, will copy the values from one to another. It's useful for converting between JPA entity and DTO classes. In my current scenario, I have DTOs that are similar but not identical to my entities. It would be nice to do something like this:

    modelMapper.map(entity, Dto::class.java).copy(moreFields = data)

Some parts of model-mapper work in Kotlin, but it's very inconsistent. I was wondering if there is some native Kotlin construct that could do this, especially with immutable Data classes?
## [10][How to implement Retrofit in an Android app](https://www.reddit.com/r/Kotlin/comments/hyzl66/how_to_implement_retrofit_in_an_android_app/)
- url: https://youtu.be/cJ4_06AOwJE
---

