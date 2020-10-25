# Kotlin
## [1][Choosing the right architecture for a [new] Kotlin Multiplatform, Jetpack Compose and SwiftUI app](https://www.reddit.com/r/Kotlin/comments/jh9xk2/choosing_the_right_architecture_for_a_new_kotlin/)
- url: https://www.marcogomiero.com/posts/2020/kmm-shared-app-architecture
---

## [2][Trying to get property reference from a type given by function receiver](https://www.reddit.com/r/Kotlin/comments/jhr1au/trying_to_get_property_reference_from_a_type/)
- url: https://www.reddit.com/r/Kotlin/comments/jhr1au/trying_to_get_property_reference_from_a_type/
---
I'm making a builder for any class. It's based on KProperty reflections and it works well for now, but I want to improve an interface of it. Here is an example of current usage:  


    data class Data(
        val nonDef: String,
        val defValue: String = "DEF_VALUE",
        val copyValue: String = nonDef,
        val nonDef2: String,
        val aMap: Map&lt;String, String&gt; = emptyMap(),
)
    val builtData = Builder&lt;Data&gt;()
        .assign(Data::aMap, mapOf("key" to "value"))
        .assign(Data::nonDef2, "10")
        .assign(Data::nonDef, "10")
        .build()

I don't like, that I have to specify Data:: receiver for every property reference in assign method. Data class is already given as generic for Builder, so I can pass only reference with receiving Data class.

So, I want to make new function prettyAssign() with callable argument, where class reference receiver  is passed as function receiver for callable, so I could address ::aMap, ::nonDef and ::nonDef2 to implicit 'this'. Here how it would look like:

    class Builder&lt;T : Any&gt;(val kClass: KClass&lt;T&gt;) {
        fun &lt;K&gt; assign(prop: KProperty1&lt;T, K&gt;, value: K)
            : Builder&lt;T&gt; = TODO("doing all reflection magic here")
            
        fun &lt;K&gt; prettyAssign(call: KClass&lt;T&gt;.() -&gt; Pair&lt;KProperty1&lt;T, K&gt;, K&gt;) {
            val (prop, value) = call(kClass)
            return assign(prop, value)
        }
    }
    
    val builtData = DataBuilder&lt;Data&gt;()
        .prettyAssign { ::aMap to mapOf("key" to "value") }
        .prettyAssign { ::nonDef2 to "10" }
        .prettyAssign { ::nonDef to "10" }
        .build()

In that approach it doesn't work, because KClass&lt;T&gt; doesn't have references to T properties.

So, I don't understand, is there a way to make prettyAssign() working
## [3][How to create a fat jar for Kotlin Multiplatform project without main class?](https://www.reddit.com/r/Kotlin/comments/jhh0hc/how_to_create_a_fat_jar_for_kotlin_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/jhh0hc/how_to_create_a_fat_jar_for_kotlin_multiplatform/
---
 Hello! 

I'm trying to create fat jar for my kotlin multiplatform project which will then be used as a dependency on another project. When I'm trying to use the class defined in the multiplatform project on the other project, it's not able to find/import it. However, when I opened the created jar, I can see it's there (along with all the dependencies).

There is a [stackoverflow post](https://stackoverflow.com/a/62770101) that I found on how to build a fat jar for multiplatform project and I pretty much just followed that. Correct me if I'm wrong, since my multiplatform project doesn't have a main class, I do not need to include the part where it adds the main class to the attributes.

I have been having hard time to get this to work. Most of the examples online are for multiplatform projects with main class. Is there any tutorial on how to build a fat jar for a multiplatform project that doesn't have main class? Any help is greatly appreciated!

Thanks in advance! 
## [4][How to handle keyboard input in windows floating over other apps on Android?](https://www.reddit.com/r/Kotlin/comments/jhbocw/how_to_handle_keyboard_input_in_windows_floating/)
- url: https://vaclavhodek.com/6-floating-windows-on-android-keyboard-input
---

## [5][Can't Run Kotlin Project IntelliJ](https://www.reddit.com/r/Kotlin/comments/jhkzaa/cant_run_kotlin_project_intellij/)
- url: https://www.reddit.com/r/Kotlin/comments/jhkzaa/cant_run_kotlin_project_intellij/
---
Just started using Kotlin and intelli j like last week. First time I used it I had no problems but today idk if there was a recent update or something but the UI changed when creating a project. I right clicked on the src folder and created a new kotlin file but I can't run it at all Idk what happened. Can someone please help
## [6][Kotlin Team AMA session is coming!](https://www.reddit.com/r/Kotlin/comments/jgmvgz/kotlin_team_ama_session_is_coming/)
- url: https://www.reddit.com/r/Kotlin/comments/jgmvgz/kotlin_team_ama_session_is_coming/
---
Hi! The Kotlin team are getting ready to answer your questions on October 26-27. 

During this AMA, we'll also cover some of the unanswered questions from the [Kotlin 1.4 Online Event](https://kotlinlang.org/lp/event-14/). 

Keep an eye on the notifications in this subreddit, we will start at 9:00 am CET on October 26.
## [7][Choosing the right architecture for a [new] Kotlin Multiplatform, Jetpack Compose and SwiftUI app](https://www.reddit.com/r/Kotlin/comments/jgllev/choosing_the_right_architecture_for_a_new_kotlin/)
- url: https://www.marcogomiero.com/posts/2020/kmm-shared-app-architecture/
---

## [8][A good game library for Kotlin?](https://www.reddit.com/r/Kotlin/comments/jgr4wz/a_good_game_library_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jgr4wz/a_good_game_library_for_kotlin/
---
Hey, does someone know a good game library for Kotlin (for JVM, native, script/js or wasm)?
## [9][JUnit 4.12 to 4.13.1: assertThrows deprecates @Test(expected) and @Rule ExpectedException](https://www.reddit.com/r/Kotlin/comments/jgpx4d/junit_412_to_4131_assertthrows_deprecates/)
- url: https://www.youtube.com/watch?v=qsPDdlL-iqA
---

## [10][Kotlin JVM Gradle plugin documentation](https://www.reddit.com/r/Kotlin/comments/jgbscp/kotlin_jvm_gradle_plugin_documentation/)
- url: https://www.reddit.com/r/Kotlin/comments/jgbscp/kotlin_jvm_gradle_plugin_documentation/
---
Suppose I wanted to know what properties are available on the Kotlin object introduced by the Kotlin JVM Gradle plugin. From the [Using Gradle reference](https://kotlinlang.org/docs/reference/using-gradle.html) I can see that kotlin.incremental is a valid property, but is there a real reference of all properties and methods of the kotlin JVM plugin (preferably w/o having to use reflection to find them)?
