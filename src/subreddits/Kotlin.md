# Kotlin
## [1][Kotlin generics in multiple platforms.](https://www.reddit.com/r/Kotlin/comments/gvtjbp/kotlin_generics_in_multiple_platforms/)
- url: https://www.reddit.com/r/Kotlin/comments/gvtjbp/kotlin_generics_in_multiple_platforms/
---
Greetings to all. I will acknowledge beforehand that I could have just missed reading about this. As I know already when working with Kotlin/JVM that generic types are erased after compilation. To circumvent this and perform some tasks like checking types at runtime, we use a reified keyword. My question is if this is the same behavior when working with Kotlin/JS or Kotlin/Native?
## [2][Android MVI with Kotlin Coroutines &amp; Flow](https://www.reddit.com/r/Kotlin/comments/gv8vv7/android_mvi_with_kotlin_coroutines_flow/)
- url: https://quickbirdstudios.com/blog/android-mvi-kotlin-coroutines-flow/?utm_source=reddit.com
---

## [3][What name is the best for a Kotlin Multiplatform MVI framework?](https://www.reddit.com/r/Kotlin/comments/gvf255/what_name_is_the_best_for_a_kotlin_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/gvf255/what_name_is_the_best_for_a_kotlin_multiplatform/
---
I was told that the name "MVIKotlin" violates [Kotlin Foundation guidelines](https://kotlinlang.org/foundation/guidelines.html). Please help me find another name for the project.

Project page: [https://github.com/arkivanov/MVIKotlin](https://github.com/arkivanov/MVIKotlin)

Thanks!

[View Poll](https://www.reddit.com/poll/gvf255)
## [4][API design question for those who lean pretty hard into the FP camp.](https://www.reddit.com/r/Kotlin/comments/gvbymb/api_design_question_for_those_who_lean_pretty/)
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
## [5][[Style] Free functions or object methods for business logic?](https://www.reddit.com/r/Kotlin/comments/gvbnk6/style_free_functions_or_object_methods_for/)
- url: https://www.reddit.com/r/Kotlin/comments/gvbnk6/style_free_functions_or_object_methods_for/
---
I think top-level functions are great, but maybe only for fundamental or very generic things. I think for business logic I'm leaning more toward the way I used to do things in Scala/Java:

    object FooService {
        fun processThing(thing: Thing): ProcessedThing
    }

Maybe just leave the free functions for entrypoints to DSLs and `mapOf()` constructor-like functions. I just like the discoverability of "namespacing" my business logic services, I guess...

Opinions?
## [6]["JSON conversion problem: Parameter specified as non-null is null" is thrown when trying to reference one JDL entity from another entity in JHipster Kotlin.](https://www.reddit.com/r/Kotlin/comments/gvi84q/json_conversion_problem_parameter_specified_as/)
- url: https://www.reddit.com/r/Kotlin/comments/gvi84q/json_conversion_problem_parameter_specified_as/
---
I'm trying to build a website in JDL Kotlin, but am having a problem with the JDL generated entities and edit forms. I have a JDL entity called `Provider`, an entity called `Plan` that needs to reference which provider it belongs to, and a third entity called `PerDayPricing` that references both `Provider` and `Plan`, all as per the following JDL configuration: 

    entity Provider {
        // [...]
    }

    entity Plan {
        // [...]
    }

    entity PerDayPricing {
        // [...]
    }

    relationship OneToMany {
        Provider to Plan
    }

    relationship OneToMany {
        Provider to PerDayPricing
        Plan to PerDayPricing
    }
    
    service all with serviceImpl

However, when I try to create or set the Provider field on a Plan item, it spits out the following error:

    WARN [PID] --- [ XNIO-1 task-10] .m.m.a.ExceptionHandlerExceptionResolver : Resolved [org.springframework.http.converter.HttpMessageConversionException: JSON conversion problem: Parameter specified as non-null is null: method com.example.project.domain.Plan.setPerDayPricings, parameter &lt;set-?&gt;; nested exception is com.fasterxml.jackson.databind.JsonMappingException: Parameter specified as non-null is null: method com.example.project.domain.Plan.setPerDayPricings, parameter &lt;set-?&gt; (through reference chain: com.example.project.domain.Plan["perDayPricings"])]

`PerDayPricing` is referenced here even though I didn't change any of its items. When I do try to set the Provider field on a PerDayPricing item, it gives the following error:

    WARN [PID] --- [ XNIO-1 task-32] .m.m.a.ExceptionHandlerExceptionResolver : Resolved [org.springframework.http.converter.HttpMessageConversionException: JSON conversion problem: Parameter specified as non-null is null: method com.example.project.domain.Provider.setPlans, parameter &lt;set-?&gt;; nested exception is com.fasterxml.jackson.databind.JsonMappingException: Parameter specified as non-null is null: method com.example.project.domain.Provider.setPlans, parameter &lt;set-?&gt; (through reference chain: com.example.project.domain.PerDayPricing["provider"]-&gt;com.example.project.domain.Provider["plans"])]

I actually have no idea what's going on here as I don't have a lot of experience with JHipster. I simply imported the JDL source file and let JHipster create files based on its configuration, without changing any of them, and this is already happening. References to the method names `setPlans` and `setPerDayPricings` don't even exist in the codebase, so I'm assuming they are being generated by Kotlin in the background?

Does anyone know what's going on and how I can fix it?
## [7][Kotlin react router example using the new gradle plugin](https://www.reddit.com/r/Kotlin/comments/gur5n8/kotlin_react_router_example_using_the_new_gradle/)
- url: https://www.reddit.com/r/Kotlin/comments/gur5n8/kotlin_react_router_example_using_the_new_gradle/
---
Is there any examples of using kotlin routers using the new kotlin js. I have seen examples with npm style. Please help
## [8][For people who are into listening to music while coding...](https://www.reddit.com/r/Kotlin/comments/gvay6h/for_people_who_are_into_listening_to_music_while/)
- url: https://www.reddit.com/r/Kotlin/comments/gvay6h/for_people_who_are_into_listening_to_music_while/
---
Post here what you're into!

Curious about what you listen to.
## [9][10 reasons to try Kotlin for Android development](https://www.reddit.com/r/Kotlin/comments/gv2l8c/10_reasons_to_try_kotlin_for_android_development/)
- url: https://haizrulamrie.com/watch?v=LEi1ecigDFE
---

## [10][1st KorGE Gamejam - 100% Kotlin - 600 EUR prizepool](https://www.reddit.com/r/Kotlin/comments/gu4wdd/1st_korge_gamejam_100_kotlin_600_eur_prizepool/)
- url: https://www.reddit.com/r/Kotlin/comments/gu4wdd/1st_korge_gamejam_100_kotlin_600_eur_prizepool/
---
We are happy to announce the first KorGE GameJAM. KorGE is a 100% Kotlin game engine and we put so much work and effort in it, that we are really happy to be able to hold our first game jam. We were able to collect *600 EUR of prizes* and a *1 year all products license* from Jetbrains.  


You can get all infos here: [https://blog.korge.org/2020/05/korge-game-jam.html](https://blog.korge.org/2020/05/korge-game-jam.html).  
Feel free to contact us with any questions. Here or in our [slack](https://slack.korge.org).

https://preview.redd.it/5dn96wgsf5251.png?width=640&amp;format=png&amp;auto=webp&amp;s=f9b7796fc095ff08a98c2481ced80ae65865b42e
