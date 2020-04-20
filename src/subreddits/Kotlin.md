# Kotlin
## [1][New release of MockK](https://www.reddit.com/r/Kotlin/comments/g4h4x7/new_release_of_mockk/)
- url: https://i.redd.it/ikg2bwkmsut41.png
---

## [2][Kotlin Microservices With Spring Boot And Spring Cloud: Part 2 - Distributed Configuration &amp; Service Discovery](https://www.reddit.com/r/Kotlin/comments/g4m8og/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://youtu.be/laI2yxthk3c
---

## [3][KVision 3.7.0 is released (Object oriented web framework for Kotlin/JS)](https://www.reddit.com/r/Kotlin/comments/g4sla6/kvision_370_is_released_object_oriented_web/)
- url: https://www.reddit.com/r/Kotlin/comments/g4sla6/kvision_370_is_released_object_oriented_web/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.7.0. The highlights of this release include:

* new server-side module with full support for the [Vert.x](https://vertx.io) backend (with examples and documentation)
* support for awesome toast messages based on [Toastr library](https://codeseven.github.io/toastr/)
* support for method annotations for compiler plugin code generation
* many dependencies upgraded (including Kotlin 1.3.72, Jooby 2.8.1 and Tabulator 4.6.2)
* bugfixes and improvements for different components

This version contains also some workarounds for [jQuery breaking issue with 3.5.0](https://github.com/jquery/jquery/issues/4665).

For more details about this release see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.7.0). There are also some new fullstack example apps for Vert.x server in the [examples repository](https://github.com/rjaros/kvision-examples).

Don't forget to ⭐️ my [project](https://github.com/rjaros/kvision) if you like it!

Any feedback is appreciated!
## [4][Some trick with functions](https://www.reddit.com/r/Kotlin/comments/g4p61j/some_trick_with_functions/)
- url: https://www.reddit.com/r/Kotlin/comments/g4p61j/some_trick_with_functions/
---
During writing some project which heavily uses dsl builders I found out about equality of functions signatures or at least their ability to substitute each other:

    fun Type.() -&gt; Unit 

and

    fun(type: Type): Unit

This allows some handy tricks, the example below: 

    class Example(val someField:Int = 1,
                  val someOther:String = "hello")
        fun writeTo(sb: StringBuilder) {
            sb.append(someField).append(" ").append(someOther)
        }
    }
    val e = Example()
    val string = buildString(e::writeTo)
    
    println(string)

This code successfully writes object to the StringBuilder, and stdlib function buildString properly returns written string.
## [5][Kotlin Multiplatform With JVM for IOS and Android](https://www.reddit.com/r/Kotlin/comments/g4azee/kotlin_multiplatform_with_jvm_for_ios_and_android/)
- url: https://www.reddit.com/r/Kotlin/comments/g4azee/kotlin_multiplatform_with_jvm_for_ios_and_android/
---
Okay just looking for any clarification, and I'm sure there are places I could google but nothing has been straight forward. Is it possible to use Java/JVM libraries in a Multiplatform Project, and have the application run on Android &amp; IOS?

&amp;#x200B;

Feel free to post links and give personal experience on this topic.
## [6][Does literally anyone besides me use Kotlin for development outside of Android? (JVM Kotlin).](https://www.reddit.com/r/Kotlin/comments/g3r6yr/does_literally_anyone_besides_me_use_kotlin_for/)
- url: https://www.reddit.com/r/Kotlin/comments/g3r6yr/does_literally_anyone_besides_me_use_kotlin_for/
---
Anytime I look for new libraries in Kotlin or fixes, they always have fixes for Android Studio or the android library system. I love Kotlin, as I know both Java and Kotlin, Kotlin in every way is better than Java. But, I feel like Kotlin is just used for Android.

&amp;#x200B;

Edit: I obviously know Kotlin is used heavily, it was more I was curious of where in the real world Kotlin is being used outside of Android. Most everyone understood this and replied with their experiences &amp; thoughts on Kotlin.   
## [7][VKUG - Fight Complexity With Functional Programming](https://www.reddit.com/r/Kotlin/comments/g457o1/vkug_fight_complexity_with_functional_programming/)
- url: https://youtu.be/uGxx01yYAgk
---

## [8][Some improvements in Eclipse Kotlin Plugin](https://www.reddit.com/r/Kotlin/comments/g46ngy/some_improvements_in_eclipse_kotlin_plugin/)
- url: https://www.reddit.com/r/Kotlin/comments/g46ngy/some_improvements_in_eclipse_kotlin_plugin/
---
Usually if I open .kt file in eclipse, mostly it will hang

But the new version  [https://marketplace.eclipse.org/content/kotlin-plugin-eclipse](https://marketplace.eclipse.org/content/kotlin-plugin-eclipse) 

Seems to improved a bit and the files are opening without much hanging
## [9][Kotlin Microservices With Spring Boot And Spring Cloud: Part 1 - Intro to Spring Boot](https://www.reddit.com/r/Kotlin/comments/g3wpdz/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://www.youtube.com/watch?v=KnbH4hOswLA
---

## [10][Coroutines and exceptions: things to know](https://www.reddit.com/r/Kotlin/comments/g3mbju/coroutines_and_exceptions_things_to_know/)
- url: https://www.rockandnull.com/coroutines-and-exceptions/
---

