# Kotlin
## [1][Kotlin 1.4.0-RC: Debugging coroutines – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/i01l6p/kotlin_140rc_debugging_coroutines_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-rc-debugging-coroutines/
---

## [2][AndroidBites | 6 things that you need to know before destructuring Parameters](https://www.reddit.com/r/Kotlin/comments/i0fcjz/androidbites_6_things_that_you_need_to_know/)
- url: https://chetangupta.net/destructuring-limitations/
---

## [3][I made a puzzle game and an automatic solver](https://www.reddit.com/r/Kotlin/comments/i0cgta/i_made_a_puzzle_game_and_an_automatic_solver/)
- url: https://www.reddit.com/r/Kotlin/comments/i0cgta/i_made_a_puzzle_game_and_an_automatic_solver/
---
Hi there! I spent the last couple days working on this project, and I wanted to share it.

Here it is: https://github.com/amtejani/Number-Game

I hope some of you will have the patience to take a look at the game and maybe give me feedback on my code. I (obviously) didn't put much work into the UI since my goal here was to write an algorithm to solve the puzzle. I also didn't handle exceptions/unexpected states very gracefully.

I do have a couple questions for those who can answer. 

I don't have much experience with Gradle other than just adding dependencies in Android Gradle scripts. How do I add 3rd party libraries with  this project? The build.gradle file was generated with the Korge intellij plugin. 

Also, does anyone have suggestions for another algorithms project I could pick up next?

Thanks for reading!
## [4][Bootiful GraphQL with Kotlin](https://www.reddit.com/r/Kotlin/comments/hzzleu/bootiful_graphql_with_kotlin/)
- url: https://youtu.be/1siPT1pTXFU?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

## [5][Why am I getting an error code of 400 with Fuel?](https://www.reddit.com/r/Kotlin/comments/i0e0g7/why_am_i_getting_an_error_code_of_400_with_fuel/)
- url: https://www.reddit.com/r/Kotlin/comments/i0e0g7/why_am_i_getting_an_error_code_of_400_with_fuel/
---
I have this Kotlin class:

    package eu.fliney.spica
    
    import com.github.kittinunf.fuel.Fuel
    import com.github.kittinunf.fuel.core.extensions.jsonBody
    
    class API {
    
        private val urlBase = "https://alles.cx/api"
    
        fun login(username: String, password: String) {
            Fuel.post("$urlBase/login")
                .jsonBody("{\"username\": \"$username\", \"username\": \"$password\"}")
                .also { println(it) }
                .response { result -&gt;  println(result.toString()) }
        }
    }

Everything in that should work, because I found it as an example and tweaked it. I have tried this exact request with postman. The exact same username and password. It returns the desired outcome. 400 Bad Request is used by the API for this error:

    { 
    "err": "badRequest" 
    }

I am wondering if this is the error occurring, and that fuel doesn't return the actual response if I get a 400 error. Does anyone know the problem here?

Here's the output:

    I/System.out: [Failure: HTTP Exception 400 Bad Request
    I/System.out:     com.github.kittinunf.fuel.core.FuelError$Companion.wrap(FuelError.kt:84)
            com.github.kittinunf.fuel.core.FuelError$Companion.wrap$default(FuelError.kt:83)
            com.github.kittinunf.fuel.core.requests.RequestTaskCallbacks.call(RequestTaskCallbacks.kt:30)
            com.github.kittinunf.fuel.core.requests.RequestTaskCallbacks.call(RequestTaskCallbacks.kt:20)
            java.util.concurrent.FutureTask.run(FutureTask.java:266)
            java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1167)
            java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:641)
            java.lang.Thread.run(Thread.java:919)
        Caused by: HTTP Exception 400 Bad Request
            com.github.kittinunf.fuel.core.FuelError$Companion.wrap(FuelError.kt:86)
        Caused by: com.github.kittinunf.fuel.core.HttpException: HTTP Exception 400 Bad Request
            com.github.kittinunf.fuel.core.requests.RequestTask.prepareResponse(RequestTask.kt:35)
            com.github.kittinunf.fuel.core.requests.RequestTask.call(RequestTask.kt:47)
            com.github.kittinunf.fuel.core.requests.RequestTask.call(RequestTask.kt:14)
            com.github.kittinunf.fuel.core.requests.RequestTaskCallbacks.call(RequestTaskCallbacks.kt:28)
            com.github.kittinunf.fuel.core.requests.RequestTaskCallbacks.call(RequestTaskCallbacks.kt:20)
            java.util.concurrent.FutureTask.run(FutureTask.java:266)
            java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1167)
            java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:641)
            java.lang.Thread.run(Thread.java:919)
        ]
    I/eu.fliney

I am sure it is from the .response { result -&gt; println(result.toString()) } because the message goes away when i remove the print statement.
## [6][Super NOOB question: Why does everything before the variable disappear if I put something after the variable?](https://www.reddit.com/r/Kotlin/comments/hzs795/super_noob_question_why_does_everything_before/)
- url: https://www.reddit.com/r/Kotlin/comments/hzs795/super_noob_question_why_does_everything_before/
---
I'm extremely new to Kotlin and have zero background in coding so I might not even ask this question correctly. I'm just playing around trying to learn something new. Anyway... I am trying to read data from a text file and use the data as a variable. It's a super simple list with one item per line:

Bacon

Steak

Ice Cream

Can someone explain to me why writing anything after the variable makes the variable and everything before it not print... except for the last item on the list? So, for the following code:

    import java.io.File
    
    fun main() {
        val foodList = File("data/food-list.txt")
                .readText()
                .split("\n")
    
    //  NOTHING AFTER VARIABLE: OUTPUT IS EXACTLY WHAT I EXPECTED
        foodList.forEach { food -&gt;
            println("I want $food")
        }
    //  PERIOD AFTER VARIABLE: ONLY THE PERIOD FOR THE FIRST 2 ITEMS
        foodList.forEach { food -&gt;
            println("I want $food.")
        }
    //  VARIABLE AT THE BEGINNING: ONLY "sounds delicious!" FOR FIRST 2 ITEMS
        foodList.forEach {food -&gt;
            println("$food sounds delicious!")
        }
    //    VARIABLE IN THE MIDDLE: ONLY "is delicious?" FOR FIRST 2 ITEMS
        foodList.forEach {food -&gt;
            println("Don't you think $food is delicious?")
        }
    }

&amp;#x200B;

The output is:

I want Bacon

I want Steak

I want Ice Cream

.

.

I want Ice Cream.

 sounds delicious!

 sounds delicious!

Ice Cream sounds delicious!

 is delicious?

 is delicious?

Don't you think Ice Cream is delicious?

&amp;#x200B;

Can someone explain to me (ELI5) what is going on here?
## [7][Kotlin Cheat Sheet](https://www.reddit.com/r/Kotlin/comments/hzdk56/kotlin_cheat_sheet/)
- url: https://i.redd.it/w4warklofid51.jpg
---

## [8][AndroidBites | How to catch ?: Errors !!](https://www.reddit.com/r/Kotlin/comments/hzu76x/androidbites_how_to_catch_errors/)
- url: https://chetangupta.net/kotlin-try-null/
---

## [9][What's the best way to observe an ArrayList in Koltin?](https://www.reddit.com/r/Kotlin/comments/i007m1/whats_the_best_way_to_observe_an_arraylist_in/)
- url: https://www.reddit.com/r/Kotlin/comments/i007m1/whats_the_best_way_to_observe_an_arraylist_in/
---

## [10][Why using SAM lambda in Flink causes `InvalidTypesException`](https://www.reddit.com/r/Kotlin/comments/hzwea5/why_using_sam_lambda_in_flink_causes/)
- url: https://discuss.kotlinlang.org/t/using-sam-lambda-in-flink-causes-invalidtypesexception/18615
---

