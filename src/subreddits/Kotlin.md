# Kotlin
## [1][Kotlin Illustrated Guide - Conditionals: When and If](https://www.reddit.com/r/Kotlin/comments/geizkn/kotlin_illustrated_guide_conditionals_when_and_if/)
- url: https://typealias.com/start/kotlin-conditionals/
---

## [2][Ok, I found something weird](https://www.reddit.com/r/Kotlin/comments/geeto3/ok_i_found_something_weird/)
- url: https://www.reddit.com/r/Kotlin/comments/geeto3/ok_i_found_something_weird/
---
I was just playing around with enum classes, and I noticed that you can declare a method inside a enum (please note, not overriding an abstract method, just a adding a regular method) 

sample:

    enum class MyEnum {
        GREEN {
            fun action(): String {
                return "I'm green"
            }
        },
        RED {
            fun anotherAction(): String {
                return "I'm red"
            }
        }
    }

Kotlin has no problem compiling this code, and decompiled bytecode looks okay:

    static final class GREEN extends MyEnum {
          @NotNull
          public final String action() {
             return "I'm green";
          }
    
          GREEN(String $enum_name_or_ordinal$0, int $enum_name_or_ordinal$1) {
             super((DefaultConstructorMarker)null);
          }
       }

However, if I try to call the method:

    MyEnum.GREEN.action() 

I get: \`Unresolved reference: action\`

&amp;#x200B;

Is this a bug in compiler? Should it not compile the these functions in the enum entry? Or there is another way to call the \`action\` function?
## [3][Swift vs. Kotlin: The Similarities and Differences You Should Know](https://www.reddit.com/r/Kotlin/comments/gek1gq/swift_vs_kotlin_the_similarities_and_differences/)
- url: https://medium.com/better-programming/swift-vs-kotlin-the-similarities-and-differences-you-should-know-b2f1be201888?source=friends_link&amp;sk=7e9202d4766b9d72a3a5fc44aebbd134
---

## [4][HBPE – Histogram Based Percentile Estimator (New Kotlin Library)](https://www.reddit.com/r/Kotlin/comments/gegri6/hbpe_histogram_based_percentile_estimator_new/)
- url: https://www.reddit.com/r/Kotlin/comments/gegri6/hbpe_histogram_based_percentile_estimator_new/
---
[https://github.com/IBM/HBPE](https://github.com/IBM/HBPE)
## [5][Machine Learning in Kotlin](https://www.reddit.com/r/Kotlin/comments/gejbo7/machine_learning_in_kotlin/)
- url: http://haifengl.github.io/api/kotlin/smile-kotlin/index.html
---

## [6][The Kotlin Programming Language - Free Course](https://www.reddit.com/r/Kotlin/comments/gdzied/the_kotlin_programming_language_free_course/)
- url: https://caster.io/courses/kotlin-programming-language
---

## [7][Kotlin chat server or community for quick questions?](https://www.reddit.com/r/Kotlin/comments/geifq6/kotlin_chat_server_or_community_for_quick/)
- url: https://www.reddit.com/r/Kotlin/comments/geifq6/kotlin_chat_server_or_community_for_quick/
---
Hi All,

I'm learning Kotlin and have a background in C#. I was wondering if there is an IRC channel / Slack / Discord or some other \*active\* community where I can ask questions? I have a background in C# and sometimes Kotlin's syntax confuses me a little.
## [8][Is it possible to use Kotlin/Native (multiplatform) with a Flutter UI in place of SwiftUI and Jetpack Compose?](https://www.reddit.com/r/Kotlin/comments/ge9cut/is_it_possible_to_use_kotlinnative_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/ge9cut/is_it_possible_to_use_kotlinnative_multiplatform/
---

## [9][TornadoFX - How to distribute an app](https://www.reddit.com/r/Kotlin/comments/ge26qn/tornadofx_how_to_distribute_an_app/)
- url: https://www.reddit.com/r/Kotlin/comments/ge26qn/tornadofx_how_to_distribute_an_app/
---
I'm a beginner. What are some ways to distribute this application that I've written? I know that I can just create a `.jar` file but is there anyway to make an installer or something like that? And how would I go about integrating that into my gradle build?

I am using grade/kotlin/IntelliJ IDEA
## [10][(Advice) What after Head First Kotlin?](https://www.reddit.com/r/Kotlin/comments/gdvfzo/advice_what_after_head_first_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gdvfzo/advice_what_after_head_first_kotlin/
---
I just finished the book ' Head First Kotlin ' and have no prior java experience. However I do have some python and C++ experience.

*Question*  :  I am looking to get started in Android dev. Should I learn Java or just brush up on the abstraction Kotlin provides? Also please include some resources, if you have any, for the same.

Ps: books preferably

Thank you for your time.
