# Kotlin
## [1][Bootiful GraphQL with Kotlin](https://www.reddit.com/r/Kotlin/comments/hzzleu/bootiful_graphql_with_kotlin/)
- url: https://youtu.be/1siPT1pTXFU?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

## [2][Super NOOB question: Why does everything before the variable disappear if I put something after the variable?](https://www.reddit.com/r/Kotlin/comments/hzs795/super_noob_question_why_does_everything_before/)
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
## [3][Kotlin Cheat Sheet](https://www.reddit.com/r/Kotlin/comments/hzdk56/kotlin_cheat_sheet/)
- url: https://i.redd.it/w4warklofid51.jpg
---

## [4][AndroidBites | How to catch ?: Errors !!](https://www.reddit.com/r/Kotlin/comments/hzu76x/androidbites_how_to_catch_errors/)
- url: https://chetangupta.net/kotlin-try-null/
---

## [5][What's the best way to observe an ArrayList in Koltin?](https://www.reddit.com/r/Kotlin/comments/i007m1/whats_the_best_way_to_observe_an_arraylist_in/)
- url: https://www.reddit.com/r/Kotlin/comments/i007m1/whats_the_best_way_to_observe_an_arraylist_in/
---

## [6][Why using SAM lambda in Flink causes `InvalidTypesException`](https://www.reddit.com/r/Kotlin/comments/hzwea5/why_using_sam_lambda_in_flink_causes/)
- url: https://discuss.kotlinlang.org/t/using-sam-lambda-in-flink-causes-invalidtypesexception/18615
---

## [7][Zephyria - Kotlin RPG](https://www.reddit.com/r/Kotlin/comments/hzbupy/zephyria_kotlin_rpg/)
- url: https://www.reddit.com/r/Kotlin/comments/hzbupy/zephyria_kotlin_rpg/
---
Hi,

This is a post to make those interested in Kotlin and games aware of the Zephyria RPG. Here's a 4-minute gameplay demo: [https://www.youtube.com/watch?v=LIV8UJAokA0](https://www.youtube.com/watch?v=LIV8UJAokA0)

The game is originally a testbed (sandboxed development area so to speak) for the FXGL engine. As it currently stands, the game can be compiled and run. Whilst some actions, such as movement, are fully implemented, other actions, such as combat, might be glitchy, whereas talking to NPCs will outright crash the game.

The quality of code also varies. The newly added code uses latest FXGL features and is concise. The leftover legacy code is messy in places with a few dozen "TODO"s.

Contributions *are* welcome, however further refactoring and code cleanup are needed from my side to make the repository contributor-friendly. Once it is ready, there will be a further post to call for contributions.

I hope this is a useful addition to the Kotlin game development community.

Thank you!

Game sources: [https://github.com/AlmasB/Zephyria](https://github.com/AlmasB/Zephyria) 

Engine sources: [https://github.com/AlmasB/FXGL](https://github.com/AlmasB/FXGL)

[Screenshot](https://preview.redd.it/vlcl5ziqlkd51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=08737f1fd9b3f92bd7577f126c64437fe0703786)
## [8][Is there a workaround for this error](https://www.reddit.com/r/Kotlin/comments/hzmc03/is_there_a_workaround_for_this_error/)
- url: https://www.reddit.com/r/Kotlin/comments/hzmc03/is_there_a_workaround_for_this_error/
---
[https://discuss.kotlinlang.org/t/accidental-override-when-using-third-party-java-interface-and-extending-runtimeexception-in-kotlin/7658](https://discuss.kotlinlang.org/t/accidental-override-when-using-third-party-java-interface-and-extending-runtimeexception-in-kotlin/7658)

That is an existing post about an error that I just ran into myself. There is no information about a workaround. Basically I'm extending RuntimeException and implementing GraphQLError. They both have a method (in Java) called `String getMessage()`. Those signatures seem to clash. Is there any known way around this? Thanks.
## [9][What's the best way to launch a small coroutine?](https://www.reddit.com/r/Kotlin/comments/hzd4u3/whats_the_best_way_to_launch_a_small_coroutine/)
- url: https://www.reddit.com/r/Kotlin/comments/hzd4u3/whats_the_best_way_to_launch_a_small_coroutine/
---
I think this is more related to Kotlin than Android, that's why I post here.

What's the best/fastest way to launch a small thread from a RecyclerView's Adapter that has no direct access to the ViewModel (to get the viewModelScope)? The coroutine is mostly used for choreographing animations.

`MainScopre().lauch { }` or should I pass the ViewModel somehow indirectly? 

Sorry in advance if it's a stupid question.
## [10][Best book/resource to learn Android in Kotlin language?](https://www.reddit.com/r/Kotlin/comments/hzf6vs/best_bookresource_to_learn_android_in_kotlin/)
- url: /r/androiddev/comments/hzf59x/best_bookresource_to_learn_android_in_kotlin/
---

