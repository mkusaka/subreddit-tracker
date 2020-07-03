# Kotlin
## [1][The Best Kotlin online courses and tutorials for beginners to learn Kotlin programming language in 2020](https://www.reddit.com/r/Kotlin/comments/hkhq9j/the_best_kotlin_online_courses_and_tutorials_for/)
- url: https://www.reddit.com/r/Kotlin/comments/hkhq9j/the_best_kotlin_online_courses_and_tutorials_for/
---
Found an amazing list of all the top-rated [Kotlin courses](https://blog.coursesity.com/best-kotlin-tutorials?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=redditPost&amp;utm_term=best-kotlin-js) of all time.  

Many of these courses are very helpful to Kotlin for the beginners.
## [2][Integrating Kotlin JS with Typescript](https://www.reddit.com/r/Kotlin/comments/hka8l5/integrating_kotlin_js_with_typescript/)
- url: https://www.reddit.com/r/Kotlin/comments/hka8l5/integrating_kotlin_js_with_typescript/
---
I recently joined a project and they have a development system typed from backend to frontend:

1. A project with DTOs and service interfaces
2. A backend that uses those DTOs, and implements those service interfaces to an HTTP API
3. JSweet converts those DTOs to TypeScript types, and creates HTTP clients for the service interfaces

As a result, you can re-use the same `User` class in the backend or frontend, fully typed everywhere.

It's nice but a bit kludgey: Jsweet is stuck on Java 8, and the service code is a bit of a mystery. Kotlin seems like it would be the perfect use case here, and I've seen a few projects with Kotlin JVM backends and Kotlin JS clients all integrated together.

However, I'd like to hear from others that have done something like this. Ideally, I'd keep something similar, with my services' interfaces implemented with Kotlin on the backend, and also as Kotlin JS clients for the frontend, and data classes for use on both.

But I think I'd lose type support - one thing we really appreciate is the shared DTOs and their TypeScript definitions. I see there is a TS generator that can do this (https://github.com/ntrrgc/ts-generator) -is anyone using it today? Or are there other pitfalls we should be considering? Would appreciate any advice from others who've done something similar.

edit: I just found this: https://youtrack.jetbrains.com/issue/KT-16604 - apparently Kotlin 1.4-M1 will generate d.ts files when using the new compiler, which should help a lot.
## [3][Any more pleasant way to type this?](https://www.reddit.com/r/Kotlin/comments/hkfjxw/any_more_pleasant_way_to_type_this/)
- url: https://www.reddit.com/r/Kotlin/comments/hkfjxw/any_more_pleasant_way_to_type_this/
---
    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?,
    savedInstanceState: Bundle?): View? = 
        context?.run { SomeView(this).also { layout = it } } 

&amp;#x200B;

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?,
    savedInstanceState: Bundle?): View? =
        context?.let { layout = SomeView(it); layout }

&amp;#x200B;
## [4][Zoe: Discover the new release of the command line tool for Kafka written in Kotlin (with a new Katacoda environment to try it out from the browser)](https://www.reddit.com/r/Kotlin/comments/hki8r1/zoe_discover_the_new_release_of_the_command_line/)
- url: https://www.reddit.com/r/Kotlin/comments/hki8r1/zoe_discover_the_new_release_of_the_command_line/
---
[https://github.com/adevinta/zoe/releases/tag/v0.24.0](https://github.com/adevinta/zoe/releases/tag/v0.24.0)
## [5][Kotlin and IntelliJ Idea](https://www.reddit.com/r/Kotlin/comments/hkhbkz/kotlin_and_intellij_idea/)
- url: https://www.reddit.com/r/Kotlin/comments/hkhbkz/kotlin_and_intellij_idea/
---
Hi everyone !

I have a problem with Intellij Idea and Kotlin. In fact, Kotlin code compilation works, but IDEA says that Kotlin is not configured and throws "unresolved reference" for "println" or "TODO".

I tried everything I could find on the internet, like invalidating the cache, trying to reconfigure Kotlin, but nothing works. 

I'm am trying to make a native app, and the informations about native Kotlin apps are pretty rare, so maybe it is linked.

If you have other ideas, tell them, right now, I'm blocked...

Thanks,
## [6][Kotlin scripting API request](https://www.reddit.com/r/Kotlin/comments/hjv2fn/kotlin_scripting_api_request/)
- url: https://www.reddit.com/r/Kotlin/comments/hjv2fn/kotlin_scripting_api_request/
---
Hi, I'm definetly not new to Kotlin, but new to scripting.

The script I want to make should fetch data from an API and save it inside a text file. The file part is doable, but I struggle with the API part.

Do I need to use a library and if so, which one? I know Retrofit and OkHTTP from Android development. And how do I integrate that into a script?

&amp;#x200B;

Thanks in advance!
## [7][Is there something like JSHint and JSLint for Kotlin?](https://www.reddit.com/r/Kotlin/comments/hk198b/is_there_something_like_jshint_and_jslint_for/)
- url: https://www.reddit.com/r/Kotlin/comments/hk198b/is_there_something_like_jshint_and_jslint_for/
---

## [8][NoCopy compiler plugin for Kotlin.](https://www.reddit.com/r/Kotlin/comments/hjoyxx/nocopy_compiler_plugin_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hjoyxx/nocopy_compiler_plugin_for_kotlin/
---
Just released NoCopy, a Kotlin compiler plugin that enables using data classes as value-based classes by moderating usage of their \`copy\` method.

[https://github.com/AhmedMourad0/no-copy](https://github.com/AhmedMourad0/no-copy)
## [9][KVision 3.11.0 is released (with support for Onsen UI web components)](https://www.reddit.com/r/Kotlin/comments/hjfe5y/kvision_3110_is_released_with_support_for_onsen/)
- url: https://www.reddit.com/r/Kotlin/comments/hjfe5y/kvision_3110_is_released_with_support_for_onsen/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.11.0, after a few weeks of effort to fully implement [Onsen UI](https://onsen.io/) library bindings. Now you can create hybrid and mobile web applications, with natively designed web components, that automatically choose the look and feel based on the platform on which they are running. All with fully type-safe, consistent Kotlin API and readable DSL builders.

There is a [new example application](https://github.com/rjaros/kvision-examples/tree/master/onsenui-kitchensink), which is almost one to one rewrite of an official [Vue Kitchen Sink example](https://github.com/OnsenUI/vue-onsenui-kitchensink). You can see live demo for both [iOS](https://rjaros.github.io/kvision-examples/onsenui-kitchensink/?platform=ios) and [Android](https://rjaros.github.io/kvision-examples/onsenui-kitchensink/?platform=android) design.

Other highlights of this release:

* extended functionality of the Maps module (contributed by Jörg Rade)
* upgraded dependencies (Javalin 3.9.0, Jooby 2.8.8, Spring Boot 2.3.1, ReduxKotlin 0.5.1, Tabulator 4.7.1, Chart.js 2.9.3 and others)

For more details see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.11.0).

As always any feedback is welcomed :-)
## [10][With the receiver in scope](https://www.reddit.com/r/Kotlin/comments/hjfy4s/with_the_receiver_in_scope/)
- url: https://medium.com/@elizarov/with-the-receiver-in-scope-7b52bdcca6e9
---

