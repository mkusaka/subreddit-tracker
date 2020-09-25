# Kotlin
## [1][Any Kotlin/Google tech speakers?](https://www.reddit.com/r/Kotlin/comments/izgwwb/any_kotlingoogle_tech_speakers/)
- url: https://www.reddit.com/r/Kotlin/comments/izgwwb/any_kotlingoogle_tech_speakers/
---
\#gdg #devfest #google  
We know there are great speakers among us, technology guru’s, leaders of the pack… Yeah! That’s you we are speaking about!

We want to offer you to take part in GDG Devfest POLAND 2020 as a speaker and share your precious knowledge with our attendees.  
Kotlin is already well settled in Android, well it's a primary language there already, so if you would like to share your story please join us, and submit your talk at [https://sessionize.com/devfest-poland-2020-with-gdg/](https://sessionize.com/devfest-poland-2020-with-gdg/)  
The conference will be held online, but it's still a great opportunity for speakers and attendees to meet other awesome people :)
## [2][Are there any free courses that combined learning Kotlin and Java](https://www.reddit.com/r/Kotlin/comments/izhq2y/are_there_any_free_courses_that_combined_learning/)
- url: https://www.reddit.com/r/Kotlin/comments/izhq2y/are_there_any_free_courses_that_combined_learning/
---
I learned Kotlin first, and now learning Java to help me with Kotlin, and feel a course that combines the two or teaches you them simultaneously first java then what kotklin took and improved upon in Java would help me pick up Java quicker
## [3][The Dark Secrets of Fast Compilation for Kotlin – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/iyyiaf/the_dark_secrets_of_fast_compilation_for_kotlin/)
- url: https://blog.jetbrains.com/kotlin/2020/09/the-dark-secrets-of-fast-compilation-for-kotlin/
---

## [4][Ktor Design Process](https://www.reddit.com/r/Kotlin/comments/izivwv/ktor_design_process/)
- url: https://blog.jetbrains.com/ktor/2020/09/24/ktor-design-process/
---

## [5][When- matching a bad practice?](https://www.reddit.com/r/Kotlin/comments/izg7yr/when_matching_a_bad_practice/)
- url: https://www.reddit.com/r/Kotlin/comments/izg7yr/when_matching_a_bad_practice/
---
I was wondering, because I use `when`\-matching extensively for checking enums and sealed classes, if it isn't just an extension of `instanceof` checking in Java and as such is considered a bad practice?

Sealed classes allow the compiler to warn you if the match is not exhaustive, but this is not the case for enums.
## [6][Announcing a painless Kotlin/Multiplatform NoSQL embedded database](https://www.reddit.com/r/Kotlin/comments/iziu2q/announcing_a_painless_kotlinmultiplatform_nosql/)
- url: https://medium.com/kodein-koders/announcing-a-painless-kotlin-multiplatform-nosql-embedded-database-30fed677549c
---

## [7][Dependency injection and test api application](https://www.reddit.com/r/Kotlin/comments/izb6kw/dependency_injection_and_test_api_application/)
- url: https://www.reddit.com/r/Kotlin/comments/izb6kw/dependency_injection_and_test_api_application/
---
Hello everyone!

First time using kotlin and I'm trying to write a simple api server on it (looking at ktor or javalin for this) 
Coming from a c# background, the first thing I tried to configure was DI, and I found out about guice and koin. On c# apps it is quite common to bind some services lifetimes to the duration of an HTTP request (such as domain services, Db sessions, etc). However all the examples I've seen for both javalin and ktor seem to have either static controllers (meaning any dependency required by using koin's inject is bound to a singleton lifetime) or eagerly resolved controllers/routes which also cause services to be either singleton or eagerly instantiated.

My question is: is it not a good practice to somehow bind controllers services and Db "instances" to the lifetime of an htto request, instead of making them singletons/statics?


EDIT: I meant rest api in the title, I'm dumb
## [8][Hello, i need some help with RelativeLayout](https://www.reddit.com/r/Kotlin/comments/iz90fm/hello_i_need_some_help_with_relativelayout/)
- url: https://www.reddit.com/r/Kotlin/comments/iz90fm/hello_i_need_some_help_with_relativelayout/
---
Hello, i´m new, my name is Diego, i´m starting learning Kotlin, but i have some troubles with RelativeLayout, i can´t move the elements and i don´t know why, please help me.

[This is the video of my code](https://reddit.com/link/iz90fm/video/d0whpukhq6p51/player)
## [9][[Codewars] Implement a simple for-comprehension](https://www.reddit.com/r/Kotlin/comments/iz3pc0/codewars_implement_a_simple_forcomprehension/)
- url: https://www.reddit.com/r/Kotlin/comments/iz3pc0/codewars_implement_a_simple_forcomprehension/
---
I came across this problem and I honestly have no idea how to do it!
https://www.codewars.com/kata/5a6f71185084d76d2000001b/kotlin

I know it's probably a faux pas to discuss the solution publicly, but if you know how you'd do it, would you at least shoot me a private message?

I can see the "shape" of a solution would be something like this:

    object OptionalSyntax {
        fun &lt;T&gt; bind(o: Optional&lt;T&gt;): T = TODO()
        fun &lt;T&gt; yield(() -&gt; T): Complete&lt;T&gt; = TODO()

        class Complete&lt;T&gt; private constructor(val t: T)
    }
    
    fun &lt;T&gt; `for`(OptionalSyntax.() -&gt; Complete&lt;T&gt;): Optional&lt;T&gt; = TODO()

But I have very little clue how I'd implement something like that. I suppose you could throw a privately defined Exception in `bind` and then catch it in `for`... But can we do it without exceptions?

Can we use suspend functions to do something here? Would that let us avoid throwing exceptions?

I know Kotlin Arrow has the Fx stuff, but that involves their higher-kinded types, etc, and is very hard to skim the code to see what they're actually doing...
## [10][Construct object through parent sealed class, and child private constructor](https://www.reddit.com/r/Kotlin/comments/iyw1mo/construct_object_through_parent_sealed_class_and/)
- url: https://www.reddit.com/r/Kotlin/comments/iyw1mo/construct_object_through_parent_sealed_class_and/
---
Hi,

I want to implement some Scala inspired `Try` monad

I wanted to know if it is possible to make child class (`Success` and `Failure`) not instantiate by themselves (no public constructor), but only from the `Try.Companion.invoke` method. Here's my current code with public constructor for both child, when I make them private I got an error in `invoke` method (which I understand). I found nothing in the Kotlin documentation about this kind of thing and I don't think it is possible, so I post here just in case

    sealed class Try&lt;T&gt; {
        companion object {
            operator fun &lt;T&gt;invoke(elem: () -&gt; T) =
                try {
                    Success(elem())
                } catch (e: Exception) {
                    Failure&lt;T&gt;(e)
                }
        }
    }
    
    data class Success&lt;T&gt;(val value: T) : Try&lt;T&gt;()
    data class Failure&lt;T&gt;(val e: Exception) : Try&lt;T&gt;()

Thanks in advance :D
