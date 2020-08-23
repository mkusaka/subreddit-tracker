# Kotlin
## [1][A Chip-8 emulator written in Kotlin](https://www.reddit.com/r/Kotlin/comments/ievuul/a_chip8_emulator_written_in_kotlin/)
- url: https://github.com/cbeust/chip-8
---

## [2][My solution for mandatory parameters in Kotlin DSL](https://www.reddit.com/r/Kotlin/comments/if2hgj/my_solution_for_mandatory_parameters_in_kotlin_dsl/)
- url: https://www.reddit.com/r/Kotlin/comments/if2hgj/my_solution_for_mandatory_parameters_in_kotlin_dsl/
---
Hey all,

I've been looking for a clean solution to mandatory parameters in Kotlin DSL.

Ended up writing a linter for it, tell me what you guys think

[https://blog.kotlin-academy.com/kotlin-dsl-know-your-limits-2deaef1bab66](https://blog.kotlin-academy.com/kotlin-dsl-know-your-limits-2deaef1bab66)
## [3][SQLDelight release with Kotlin Native 1.4.0](https://www.reddit.com/r/Kotlin/comments/ieanwm/sqldelight_release_with_kotlin_native_140/)
- url: https://github.com/cashapp/sqldelight/releases/tag/1.4.1
---

## [4][How to get some Game Design principles and perhaps historical context?](https://www.reddit.com/r/Kotlin/comments/ierwah/how_to_get_some_game_design_principles_and/)
- url: https://www.reddit.com/r/Kotlin/comments/ierwah/how_to_get_some_game_design_principles_and/
---
I am a pac-man generation programmer in a F.E.A.R. (First Encounter Assault Recon - said to use advanced AI)  world. 

As I am learning the KorGE Game Library I am continually running across new terms like H.U.D.  (Heads Up Display, you knew that right?) and 'Roguelike'.   Thing is I'm designing a game while I have spent very little time playing them.  

Would somebody here care to recommend a book, a online course, a forum, or someway to bootstrap myself into the world of games?   (As developer, with commercialization in mind, not as a player.)
## [5][JVM Performance and Engineering](https://www.reddit.com/r/Kotlin/comments/ieilwp/jvm_performance_and_engineering/)
- url: https://blog.gceasy.io/2015/07/17/jvm-performance-engineering-troubleshooting/
---

## [6][ktfmt, a Kotlin code formatter based on google-java-format, now with a Plugin for IntelliJ](https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/)
- url: https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/
---
[https://github.com/facebookincubator/ktfmt/](https://github.com/facebookincubator/ktfmt/)

This came up in a couple of archived posts here, so I'm excited to say that `ktfmt` now has an IntelliJ plugin (as well as a Spotless plugin).

Also, a lot has changed in `ktfmt` since it was announced ~8 months ago - give it a try! we're listening for feedback over at https://github.com/facebookincubator/ktfmt/issues :)
## [7][To initialize a new variable that is only used once, or just use the code instead of the variable?](https://www.reddit.com/r/Kotlin/comments/idcvaz/to_initialize_a_new_variable_that_is_only_used/)
- url: https://www.reddit.com/r/Kotlin/comments/idcvaz/to_initialize_a_new_variable_that_is_only_used/
---
I was wondering what the best practice is.

Case example:

    fun foo() {
        for (x in y) {
            val bar = [some super long series of calls involving x and y]
            if (bar fulfills some condition) 
                doThing()
        }
    }

VS

    fun foo() {
        for (x in y) {
            if ([some super long series of calls involving x and y] fulfills some condition) 
                doThing()
        }
    }

I like the first one because to me it's more readable. I can assign a meaningful name to those series of calls and later I won't have to think much to get what it means.

But... is it wrong? What if I expect to run that loop many many times. Is creating variables all those times a bad thing to do? Especially when it's not needed because I only use that variable in one place.
## [8][Need Help with Ktor](https://www.reddit.com/r/Kotlin/comments/ided9u/need_help_with_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/ided9u/need_help_with_ktor/
---
Hi guys I'm a new kotlin developer and was looking into ktor. I couldn't find any videos or tutorials(other than its doc) on how to use it. I'm used to react but I've never used anything like ktor and wanted to know how to work with HTML and css on ktor and how hard it is. Can I use HTML css directly with it??
## [9][User Activity Recognition || Activity Transition API || Kotlin](https://www.reddit.com/r/Kotlin/comments/idcxpp/user_activity_recognition_activity_transition_api/)
- url: https://youtu.be/2Em5ep_KYIY
---

## [10][Video game engine](https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/)
- url: https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/
---
Hello I am learning to program with kotlin, my idea is to create video games for android. I wanted to ask you if you know of any game engine for programming that supports kotlin.

Note: Sorry for my English I speak Spanish.
