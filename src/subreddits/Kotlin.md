# Kotlin
## [1][Judge0 IDE adds Kotlin support](https://www.reddit.com/r/Kotlin/comments/gk03pm/judge0_ide_adds_kotlin_support/)
- url: https://ide.judge0.com/?YZyR
---

## [2][Less clunky way of declaring byte literals?](https://www.reddit.com/r/Kotlin/comments/gjt5iv/less_clunky_way_of_declaring_byte_literals/)
- url: https://www.reddit.com/r/Kotlin/comments/gjt5iv/less_clunky_way_of_declaring_byte_literals/
---
I work with a decent amount of byte data, and I'm writing some example byte arrays for unit tests. The default way of doing this in Kotlin is pretty cumbersome:

    val array = byteArrayOf(
        0xDE.toByte(), 0xAD.toByte(), 0xBE.toByte(), 0xEF.toByte()
    )

Is there any more elegant way of doing this? In general I love Kotlin, but this byte declaration syntax and the use of infix functions instead of bitshift operators have been something of a downgrade from Java. Or am I missing something? I've thought of trying to make some kind of convenience function, though I'm not sure if it's any better:

    fun b(value: Int): Byte {
        return value.toByte()
    }

Which then lets me do:

    val array = byteArrayOf(
        b(0xDE), b(0xAD), b(0xBE), b(0xEF)
    )

Am I missing something here or are byte literals just kinda clunky in Kotlin?
## [3][Code a (mini) Twitter clone in Kotlin using the Hexagon Toolkit](https://www.reddit.com/r/Kotlin/comments/gjq174/code_a_mini_twitter_clone_in_kotlin_using_the/)
- url: https://twitter.com/hexagon_kt/status/1260973355362107392
---

## [4][IntelliJ crashes when type inferencing](https://www.reddit.com/r/Kotlin/comments/gjycqg/intellij_crashes_when_type_inferencing/)
- url: https://www.reddit.com/r/Kotlin/comments/gjycqg/intellij_crashes_when_type_inferencing/
---
As of the two most recent updates to IntelliJ, my somewhat complicated codebase causes it to crash on a regular basis.  Any time there is sufficiently complex type inferencing going on, the IDE hangs and either takes over a minute to report an error, or just freezes.  I have filed bugs with JetBrains, but they are of course slow to respond.

Has anyone else experienced these issues?  Are there any good alternatives to IntelliJ for Kotlin development?

Honestly with this going on for a couple of months, it is becoming unusable.  I hate to say it but if this is the way forward and there are no alternatives, my team may have to slowly migrate away from Kotlin.  The language is great, but I wish it had better support in vscode etc.
## [5][Armeria - a Java RPC framework supports gRPC Kotlin](https://www.reddit.com/r/Kotlin/comments/gjkf38/armeria_a_java_rpc_framework_supports_grpc_kotlin/)
- url: https://twitter.com/armeria_project/status/1260858017601351680?s=21
---

## [6][Why can't I run fun main without vararg args: String](https://www.reddit.com/r/Kotlin/comments/gjquww/why_cant_i_run_fun_main_without_vararg_args_string/)
- url: https://www.reddit.com/r/Kotlin/comments/gjquww/why_cant_i_run_fun_main_without_vararg_args_string/
---
On one computer's intellij I couldn't run the main function without something like

    fun main(vararg args: String){
    }
But it works perfectly fine on another computer using the same IDE, i.e.


    fun main(){
    }

I searched but didn't find satisfactory explanation. Can anyone tell me why?
## [7][Kotlin Heroes 4: the Next Round is Almost Here](https://www.reddit.com/r/Kotlin/comments/gjr6ev/kotlin_heroes_4_the_next_round_is_almost_here/)
- url: https://blog.jetbrains.com/kotlin/2020/05/kotlin-heroes-4-the-next-round-is-almost-here/
---

## [8][Looking for a good resources to learn about kotlin concurrency model and coroutines](https://www.reddit.com/r/Kotlin/comments/gjcdxt/looking_for_a_good_resources_to_learn_about/)
- url: https://www.reddit.com/r/Kotlin/comments/gjcdxt/looking_for_a_good_resources_to_learn_about/
---
I am an experienced developer that is now approaching kotlin, I'd like to really understand coroutines, can you recommend some book/guide/paper
## [9][Accessible tutorial on APIs](https://www.reddit.com/r/Kotlin/comments/gj5tl5/accessible_tutorial_on_apis/)
- url: https://www.reddit.com/r/Kotlin/comments/gj5tl5/accessible_tutorial_on_apis/
---
Hi, I recently started learning Kotlin and Android Development. Already developed some basic applications, now I want to consume data from APIs and show them on the app. Do you recommend any beginner tutorial on how to connect to APIs?
## [10][Project ideas to build portfolio (not android)](https://www.reddit.com/r/Kotlin/comments/gixuca/project_ideas_to_build_portfolio_not_android/)
- url: https://www.reddit.com/r/Kotlin/comments/gixuca/project_ideas_to_build_portfolio_not_android/
---
Hello peeps,

I’m new to programming and kotlin and I was told that I will get better if I build something on my own. I did something similar with JavaScript from an online Udemy course and I made a simple website. Do you have any suggestions for projects that are not android? 

Thank you :)
