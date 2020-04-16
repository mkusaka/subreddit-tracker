# Kotlin
## [1][Question about nullable generics](https://www.reddit.com/r/Kotlin/comments/g2e4mh/question_about_nullable_generics/)
- url: https://www.reddit.com/r/Kotlin/comments/g2e4mh/question_about_nullable_generics/
---
Let's say I have a class with a single type parameter, T.

In this class, I can use T? when I want a value of type T to be nullable.

But T could already be a nullable type. 

My question is, how do I do the opposite, i.e. safely represent that a member function that returns type T is NOT nullable even if T happens to be nullable?

Or if that is confusing let me phrase it in another way: given intended construction like MyClass&lt;MyType?&gt;(), how could I write a function inside MyClass that returns a (MyType) instead of a (MyType?)

Thanks for your help
## [2][Multithreading in Kotlin Multiplatform Apps](https://www.reddit.com/r/Kotlin/comments/g258dd/multithreading_in_kotlin_multiplatform_apps/)
- url: https://helw.net/2020/04/16/multithreading-in-kotlin-multiplatform-apps/
---

## [3][Unit Testing Coroutines on Android](https://www.reddit.com/r/Kotlin/comments/g2bcuu/unit_testing_coroutines_on_android/)
- url: https://www.ericdecanini.com/2020/04/06/unit-testing-coroutines-on-android/
---

## [4][Simple Multiplatform RPC with Kotlin Serialization](https://www.reddit.com/r/Kotlin/comments/g1tipe/simple_multiplatform_rpc_with_kotlin_serialization/)
- url: https://jakewharton.com/simple-multiplatform-rpc-with-kotlin-serialization/
---

## [5][Trying to position Elements in TornadoFX](https://www.reddit.com/r/Kotlin/comments/g29r1s/trying_to_position_elements_in_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/g29r1s/trying_to_position_elements_in_tornadofx/
---
I was working on a project using TornadoFX in order to try to create a dashboard for an app that I was writing. In the process of writing this, I wanted to know how to put more layers on top of a borderpane. I was using a borderpane in order to have separate panes for the different cardinal directions however I wanted to know if it was possible to have multiple smaller panes in the borderpane.[ Image of desired view.](https://i.stack.imgur.com/Ze6NK.jpg) So when examining the image, the red rectangle is where I want to add a few buttons and a drop-down menu. However, I do not know how to add another pane on top of the borderpane in order to not only hold the title but also the "toolbar" if you will. Below, you can see my current attempt at trying to do so.

    class MainView : View("Dashboard") {
        override val root = borderpane {
            top = hbox {
                val title = label("Dashboard") {
                    addClass(Styles.title)
    
                }
            }
    
            val toolbar = stackpane {
                label("Toolbar") {
                    addClass(Styles.title)
                }
            }
            (top as HBox).alignment = Pos.TOP_CENTER
            (toolbar as StackPane).alignment = Pos.CENTER_RIGHT

This code only shows the title and the top portion of the borderpane but I can show the whole class if need be. Please help if possible. Thank you so much in advance!
## [6][Deploying a KTor site to a VPS](https://www.reddit.com/r/Kotlin/comments/g245ns/deploying_a_ktor_site_to_a_vps/)
- url: https://www.reddit.com/r/Kotlin/comments/g245ns/deploying_a_ktor_site_to_a_vps/
---
I'm an android developer, and I like the idea of creating my own APIs in the same language as my code, and KTor looks great for that. I can build my api on my local machine, but I'm not seeing any guidance on deploying that to a VPS so I can use it remotely.  Admittedly I'm kind of using IntelliJ Idea as a crutch to do the compiling work, and I'm not sure what's going on behind the scenes so I can do it myself on from a command line.  Can anyone advise either how to do that generally, or more specifically the like 'correct/programatic' way to do that?
## [7][Restul api using kotlin (Spring boot)](https://www.reddit.com/r/Kotlin/comments/g1prc9/restul_api_using_kotlin_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/g1prc9/restul_api_using_kotlin_spring_boot/
---
I wrote this post about how to use your kotlin knowledge from android to write down a restful api. This  is the first part and I am working on next parts now

[https://medium.com/@zahidrasheed/getting-started-with-restful-api-1d8ee5cac3cc](https://medium.com/@zahidrasheed/getting-started-with-restful-api-1d8ee5cac3cc)
## [8][Where can I find the 1.3.72 changelog?](https://www.reddit.com/r/Kotlin/comments/g1mo2w/where_can_i_find_the_1372_changelog/)
- url: https://www.reddit.com/r/Kotlin/comments/g1mo2w/where_can_i_find_the_1372_changelog/
---

## [9][Fixing timing issues in tests](https://www.reddit.com/r/Kotlin/comments/g1qwbz/fixing_timing_issues_in_tests/)
- url: https://blog.djy.io/fixing-timing-issues-in-tests/
---

## [10][Coroutines, exception handling and withTimeout. Can't wrap my head around this combination.](https://www.reddit.com/r/Kotlin/comments/g1kvnx/coroutines_exception_handling_and_withtimeout/)
- url: https://www.reddit.com/r/Kotlin/comments/g1kvnx/coroutines_exception_handling_and_withtimeout/
---
    import kotlinx.coroutines.*
    import kotlinx.coroutines.flow.*
    import kotlin.system.measureTimeMillis
    
    fun time() = System.currentTimeMillis()
    
    fun main(args: Array&lt;String&gt;) {
        runBlocking {
            repeat(5) {
                try {
                    println("${time()} withTimeout before")
                    withTimeout(1L) {
                        try {
                            delay(200L)
                        } catch (e: Exception) {
                            println("${time()} withTimeout exception " + e.localizedMessage)
                        }
                    }
                    println("${time()} withTimeout done")
                } catch (e: Exception) {
                    println("${time()} exception " + e.localizedMessage)
                }
                println("${time()} -----")
            }
        }
    }

This gives me the output:

    1586923684774 withTimeout before
    1586923684788 withTimeout exception Timed out waiting for 1 ms
    1586923684789 withTimeout done
    1586923684789 -----
    1586923684789 withTimeout before
    1586923684790 withTimeout exception Timed out waiting for 1 ms
    1586923684791 exception Timed out waiting for 1 ms
    1586923684791 -----
    1586923684791 withTimeout before
    1586923684792 withTimeout exception Timed out waiting for 1 ms
    1586923684792 exception Timed out waiting for 1 ms
    1586923684792 -----
    1586923684792 withTimeout before
    1586923684794 withTimeout exception Timed out waiting for 1 ms
    1586923684794 exception Timed out waiting for 1 ms
    1586923684794 -----
    1586923684794 withTimeout before
    1586923684795 withTimeout exception Timed out waiting for 1 ms
    1586923684795 exception Timed out waiting for 1 ms
    1586923684795 -----

On the first loop, `withTimeout done` gets printed. 

**Question 1: Why only on the first iteration?**

**Question 2: If I change the wait time in** `withTimeout` **from 1L to 10L, it does not get printed. Instead, the exception is propagated to the outer try/catch, like on the next iterations (See below). Why?**

    1586923727317 withTimeout before
    1586923727333 withTimeout exception Timed out waiting for 10 ms
    1586923727334 exception Timed out waiting for 10 ms
    1586923727334 -----
    1586923727334 withTimeout before
    1586923727347 withTimeout exception Timed out waiting for 10 ms
    1586923727347 exception Timed out waiting for 10 ms
    1586923727347 -----
    1586923727347 withTimeout before
    1586923727360 withTimeout exception Timed out waiting for 10 ms
    1586923727360 exception Timed out waiting for 10 ms
    1586923727360 -----
    1586923727360 withTimeout before
    1586923727370 withTimeout exception Timed out waiting for 10 ms
    1586923727371 exception Timed out waiting for 10 ms
    1586923727371 -----
    1586923727371 withTimeout before
    1586923727382 withTimeout exception Timed out waiting for 10 ms
    1586923727382 exception Timed out waiting for 10 ms
    1586923727382 -----

**Question 3: Why, even though the exception is being caught inside a try/catch block in** `withTimeout()`**, it is being propagated to the outer try/catch block?**

I appreciate and thank you in advance for your help.
