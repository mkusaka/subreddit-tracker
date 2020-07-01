# Kotlin
## [1]["?" Over "!!" always?](https://www.reddit.com/r/Kotlin/comments/hj6qqj/over_always/)
- url: https://www.reddit.com/r/Kotlin/comments/hj6qqj/over_always/
---
Since we dont want an NPE and ? guarantees null safety.I can't think of a scenario I would use !! over ?
Why would I use !! and get null
## [2][Best Kotlin framework for Kafka &amp; Kafka Streams?](https://www.reddit.com/r/Kotlin/comments/hj5l0y/best_kotlin_framework_for_kafka_kafka_streams/)
- url: https://www.reddit.com/r/Kotlin/comments/hj5l0y/best_kotlin_framework_for_kafka_kafka_streams/
---
Hello, my team is looking into using Kotlin as we all really like the syntax &amp; semantics of the language. We've also seen that Spring-Boot and Kotlin work really well together. However, we've come across a few issues and we're wondering if there is another Kotlin framework that has great support for Kafka &amp; Kafka Streams and possibly great support on VS code? We're open to using IntelliJ if need be, however, our team of 12 really enjoy using VS Code and would like to stick to it.  


To clarify, the issues we're struggling with were to get live reload to work out the box with gradle &amp; maven, as well as no support with Kotlin + Spring-Boot on VS Code. 
## [3][CMV: The SCREAMING_SNAKE_CASE is a C legacy that we should have already abandoned](https://www.reddit.com/r/Kotlin/comments/hj0fn7/cmv_the_screaming_snake_case_is_a_c_legacy_that/)
- url: https://www.reddit.com/r/Kotlin/comments/hj0fn7/cmv_the_screaming_snake_case_is_a_c_legacy_that/
---
Looking by the perspective that we use a language to program stuff, and the language is used to describe the programs, express their intents and highlight points of attention, I think the SCREAMING\_SNAKE\_CASE convention is useful for C/C++ programs because it's important to differentiate macros from symbols, but has no value for constants in modern languages. It draws attention to something that *won't* change and so is very unlikely to cause problems. On the other hand, things like global mutable variables (sounds redundant, but just to be explicit I'm talking about `var` not `val` in Kotlin) are dangerous and could use this convention to draw the attention of readers.

Some languages don't use this convention. Examples are C# (PascalCase) and Swift (camelCase). I think the official Kotlin convention should also drop it. One may argue that conventions don't need to be followed, but I think it's important, for example, for library authors to use the same convention.

Am I missing something here? I get really annoyed when I open some code and the first thing that caughts my attention are constants, for... nothing?
## [4][QLDB at Amazon](https://www.reddit.com/r/Kotlin/comments/hiu57x/qldb_at_amazon/)
- url: https://talkingkotlin.com/qldb/
---

## [5][Kotlin Illustrated Guide - Classes &amp; Objects](https://www.reddit.com/r/Kotlin/comments/him27h/kotlin_illustrated_guide_classes_objects/)
- url: https://typealias.com/start/kotlin-classes-and-objects/
---

## [6][Type Proofs and FP for the Kotlin Type System](https://www.reddit.com/r/Kotlin/comments/hisagk/type_proofs_and_fp_for_the_kotlin_type_system/)
- url: https://www.youtube.com/watch?v=ETn_6LmMjho
---

## [7][MVIKotlin 2.0.0-rc1 is out with many new things!](https://www.reddit.com/r/Kotlin/comments/hiv3at/mvikotlin_200rc1_is_out_with_many_new_things/)
- url: https://www.reddit.com/r/Kotlin/comments/hiv3at/mvikotlin_200rc1_is_out_with_many_new_things/
---
What's new:

* Support `macOS` target
* Added `StateKeeperProvider.retainStore(...)` extension functions
* MacOS time travel client app for iOS and macOS debugging
* Time travel export/import in IDEA plugin (first steps towards remote debugging)
* Sample JS app
* Time travel UI for sample JS app
* Unit and integration tests for sample shared code

Some cool demo videos:

* [Debugging iOS application using MVIKotlin time travel client app](https://youtu.be/rj6GwA2ZQkk)
* [Export/import time travel data in Android with MVIKotlin IDEA plugin](https://youtu.be/SIxfSgBkHS0)

Project GitHub: https://arkivanov.github.io/MVIKotlin
## [8][Static code analysis tools for Kotlin](https://www.reddit.com/r/Kotlin/comments/hirwnt/static_code_analysis_tools_for_kotlin/)
- url: https://www.rockandnull.com/static-code-analysis-android/
---

## [9][A help for a quick function](https://www.reddit.com/r/Kotlin/comments/hiygnc/a_help_for_a_quick_function/)
- url: https://www.reddit.com/r/Kotlin/comments/hiygnc/a_help_for_a_quick_function/
---
Hey guys,

Recently I really got really involved in learning Kotlin and just to exercise a bit I'm trying to write some exercises I was doing while learning Python in Kotlin. This one I found particularly tricky and I was wondering If someone can help me to decode the problem. I leave the code below, first the Python code, then the code I've written in Kotlin.

I always get to the point in which a `No get method providing array access` problem comes in, I really don't understand this. And sorry if this is trivial for you guys, I've just started learning!

Thanks!

Python

    def frequencies (example_string):
        map = {}
        for key in example_string:
            if key not in map.keys:
                map[key] = 1
            else:
                map[key] += 1
        print(map)
    
    &gt;&gt;&gt; frequencies("hehooo")
    {"h":2,"e":1,"o":3}

Kotlin

    fun frequencies (example_string: String) {
        var map_result = mutableMapOf&lt;String,Int&gt;()
        example_string.forEach { string : Char -&gt;
            map_result.keys[string] += 1
        }
        for (key in map_result) {
            println("${map_result.keys[key]} : ${map_result.values[key]}")
        }
    }
## [10][Ktor to your heart’s content: easy mobile backends in Kotlin](https://www.reddit.com/r/Kotlin/comments/hisx7k/ktor_to_your_hearts_content_easy_mobile_backends/)
- url: https://www.youtube.com/watch?v=p8RA-3t0jGA
---

