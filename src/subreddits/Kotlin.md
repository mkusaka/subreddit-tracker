# Kotlin
## [1][Kotlin 1.4-M2 Released](https://www.reddit.com/r/Kotlin/comments/gwfyn3/kotlin_14m2_released/)
- url: https://blog.jetbrains.com/kotlin/2020/06/kotlin-1-4-m2-released
---

## [2][Koge the new 100 % pure Kotlin OpenGL Game Engine](https://www.reddit.com/r/Kotlin/comments/gw259e/koge_the_new_100_pure_kotlin_opengl_game_engine/)
- url: https://www.reddit.com/r/Kotlin/comments/gw259e/koge_the_new_100_pure_kotlin_opengl_game_engine/
---
Koge (Kotlin OpenGL Game Engine) is a 2D game framework developed in Kotlin that works in Windows, Linux and Mac OS X. It is very simple and intuitive to use. For more information check [Koge on github](https://github.com/KogeLabs/Koge/blob/master/README.md)
## [3][Kotlin if statements and exception handling](https://www.reddit.com/r/Kotlin/comments/gwcusx/kotlin_if_statements_and_exception_handling/)
- url: https://www.reddit.com/r/Kotlin/comments/gwcusx/kotlin_if_statements_and_exception_handling/
---
A coworker and I noticed an interesting behaviour today, where the `get` operator on a `List`, if used inside an `if` statement, would never result in a crash due to an `OutOfBoundsException`, even if that index was definitely out of bounds. The statement will just return false instead.

This makes sense, in my opinion, but was wondering if anyone has any insight or documentation they can point me to regarding this, just out of interest and intrigue. I couldn't find anything with a quick google search.
## [4][Is Kotlin enough similar to Java to let me study concepts from java books and apply them in a kt app easily?](https://www.reddit.com/r/Kotlin/comments/gvy4ec/is_kotlin_enough_similar_to_java_to_let_me_study/)
- url: https://www.reddit.com/r/Kotlin/comments/gvy4ec/is_kotlin_enough_similar_to_java_to_let_me_study/
---
I'm a java programmer. I'm curious about Kotlin and I'd like to give it a try, if I could transfer my java knowledge easily. Especially the things regarding network programming and parallel programming.  
Any opinion/suggestions?
## [5][Kotlin generics in multiple platforms.](https://www.reddit.com/r/Kotlin/comments/gvtjbp/kotlin_generics_in_multiple_platforms/)
- url: https://www.reddit.com/r/Kotlin/comments/gvtjbp/kotlin_generics_in_multiple_platforms/
---
Greetings to all. I will acknowledge beforehand that I could have just missed reading about this. As I know already when working with Kotlin/JVM that generic types are erased after compilation. To circumvent this and perform some tasks like checking types at runtime, we use a reified keyword. My question is if this is the same behavior when working with Kotlin/JS or Kotlin/Native?
## [6][rant this language is driving me mad](https://www.reddit.com/r/Kotlin/comments/gwfmsy/rant_this_language_is_driving_me_mad/)
- url: https://www.reddit.com/r/Kotlin/comments/gwfmsy/rant_this_language_is_driving_me_mad/
---
i have to use a var called prevVal in a loop but the first time through the loop it doesn't make sense to have a prevVal since we are going through it for the first time and it hasn't been assigned until we go through the first iteration

 So i googled and it said use lateinit and then check prevVal.isInitialized  .. ofc the compiler won't let me do that, it doesn't know anything about isInitialized on a reference type (thanks dickheads .. you don't allow lateinit on primitive vals so what use is this lateinit anyway? .. retards)

fine i say, i'll just use a boolean flag (which is hacky and ugly af) to check if it's the first iteration.. except ofc now it complains that prevVal is never intialized. NO MOTHERFUCKER IT ISNT but if the compiler wasn't written by a retard it'd realise it has been initialized by the ONLY other branch. 

WHAT A FUCKED UP LANGUAGE

i will also add this was my 3rd attempt at a solution, the 2nd one was to check if  the first element of my results arraylist is not null.. which the kotlin List docs say first()  return null if there isn't one (THIS IS A LIE, i get a java.lang.exception saying no such element exists) .. so the docs are not only useles they're wrong as well
## [7][Android MVI with Kotlin Coroutines &amp; Flow](https://www.reddit.com/r/Kotlin/comments/gv8vv7/android_mvi_with_kotlin_coroutines_flow/)
- url: https://quickbirdstudios.com/blog/android-mvi-kotlin-coroutines-flow/?utm_source=reddit.com
---

## [8][What name is the best for a Kotlin Multiplatform MVI framework?](https://www.reddit.com/r/Kotlin/comments/gvf255/what_name_is_the_best_for_a_kotlin_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/gvf255/what_name_is_the_best_for_a_kotlin_multiplatform/
---
I was told that the name "MVIKotlin" violates [Kotlin Foundation guidelines](https://kotlinlang.org/foundation/guidelines.html). Please help me find another name for the project.

Project page: [https://github.com/arkivanov/MVIKotlin](https://github.com/arkivanov/MVIKotlin)

Thanks!

[View Poll](https://www.reddit.com/poll/gvf255)
## [9][API design question for those who lean pretty hard into the FP camp.](https://www.reddit.com/r/Kotlin/comments/gvbymb/api_design_question_for_those_who_lean_pretty/)
- url: https://www.reddit.com/r/Kotlin/comments/gvbymb/api_design_question_for_those_who_lean_pretty/
---
The tl;dr is that I have some business workflow that uses several of my `Repository` interfaces to fetch a bunch of data to massage into some result. I'm trying to decide if I should have a stateful `Service` class that has all of the repos as properties or if I should just have a function (or object method) that take the repos as parameters.

Choice #1:

    class FooService(private val barRepo: BarRepository, private val bazRepo: BazRepository) {
        fun analyzeStuff(id: UUID, timestamp: Instant): Result&lt;Data&gt; {
            TODO()
        }
    }

Choice #2:

    object FooService {
        fun analyzeStuff(id: UUID, timestamp: Instant, barRepo: BarRespository, bazRepo: BazRepository): Result&lt;Data&gt; {
            TODO()
        }
    }

The repositories are used in a read-only manner (for now) for this particular service. The problem I have with choice #1 is that I've grown weary of the deep layering of typical Java backend projects I've worked on: Controller -&gt; Service -&gt; Repository -&gt; Database. I like the idea of flattening it a bit by basically putting the business logic (Service) *next to* the persistence logic (Repository) instead of the Service *owning* the Repository. I also like that choice #2 is more explicit about what method is using what repo code.

One thing that I very much don't like about choice #2 is that this function is in no way pure, but it kinda looks like a pure function (not being attached to a class, which I use as a conventional litmus test for whether something is supposed to be pure or not). I can always call the object method once, fiddle with the DB from my ssh terminal, and then call it again and get different results. So this is not pure.

One more issue with choice #1 is that some of the methods on the `Service` **will** be pure. But you can't really tell from the signatures which ones are pure transformations and which involve persistence, since all methods will have access to the repos.

Thoughts?
## [10][[Style] Free functions or object methods for business logic?](https://www.reddit.com/r/Kotlin/comments/gvbnk6/style_free_functions_or_object_methods_for/)
- url: https://www.reddit.com/r/Kotlin/comments/gvbnk6/style_free_functions_or_object_methods_for/
---
I think top-level functions are great, but maybe only for fundamental or very generic things. I think for business logic I'm leaning more toward the way I used to do things in Scala/Java:

    object FooService {
        fun processThing(thing: Thing): ProcessedThing
    }

Maybe just leave the free functions for entrypoints to DSLs and `mapOf()` constructor-like functions. I just like the discoverability of "namespacing" my business logic services, I guess...

Opinions?
