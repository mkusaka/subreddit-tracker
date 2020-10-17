# Kotlin
## [1][Kweb 0.7.33 released, Kotlin web framework for backend developers adds support for FileReader / file uploads](https://www.reddit.com/r/Kotlin/comments/jcscyo/kweb_0733_released_kotlin_web_framework_for/)
- url: https://github.com/kwebio/kweb-core
---

## [2][http4k blog: Nanoservices: The Power of Composition](https://www.reddit.com/r/Kotlin/comments/jc8tez/http4k_blog_nanoservices_the_power_of_composition/)
- url: https://www.http4k.org/blog/nanoservices/
---

## [3][Kotlin 1.4 Online Event Recap](https://www.reddit.com/r/Kotlin/comments/jcbdi9/kotlin_14_online_event_recap/)
- url: https://link.medium.com/OzC2ERGhDab
---

## [4][Accessing class member name via interpolation of number?](https://www.reddit.com/r/Kotlin/comments/jcimht/accessing_class_member_name_via_interpolation_of/)
- url: https://www.reddit.com/r/Kotlin/comments/jcimht/accessing_class_member_name_via_interpolation_of/
---
I just want to state that I know I'm doing this in a dumb way. I'm a newb.

I have a class paragraphNumber that holds 20 different paragraphText. paragraphText1 all the way to paragraphText20. 

In my MainActivity.kt, I have a button. This button figures out which paragraph it should be loading, and that's all it does. I have some logic to figure out the next paragraph (which isn't the next one in line!) 

How can I interpolate a number into the name? I mean, now I have

rollButton.setOnClickListener **{** textView1.*text* = paragraphNumber().**paragraphText1}**

but, I would rather be able to call on paragraphTextX where X would be a number 1 - 20. Is this possible? Please let me know if I'm not making any sense.
## [5][OutOfMemoryError related JVM Arguments](https://www.reddit.com/r/Kotlin/comments/jc73ux/outofmemoryerror_related_jvm_arguments/)
- url: https://blog.heaphero.io/2019/06/21/outofmemoryerror-related-jvm-arguments/
---

## [6][How do you add a KLib to a Native Gradle Project?](https://www.reddit.com/r/Kotlin/comments/jcdgc6/how_do_you_add_a_klib_to_a_native_gradle_project/)
- url: https://www.reddit.com/r/Kotlin/comments/jcdgc6/how_do_you_add_a_klib_to_a_native_gradle_project/
---
I have one project that looks like this:

`build.gradle.kts`
```kotlin
plugins {
    kotlin("multiplatform") version "1.4.10"
}

repositories {
    mavenCentral()
}

kotlin {
    linuxX64("native") {
        binaries {
            executable {
                entryPoint = "foo.main"
            }
            staticLib()
        }
    }
}
```

`src/nativeMain/kotlin/Main.kt`
```kotlin
package foo

fun main() {
    println("hi!")
    println(foo(bar(listOf(3, 5, 10))))
}

fun foo(x: Int): List&lt;Int&gt; = (0..x).map { y -&gt; x * y }

fun bar(x: List&lt;Int&gt;): Int = x.last()
```

Gradle compiles this to a KLib. How can I use this KLib in another project?
## [7][Hi what is the problem? Dont see button id and textView](https://www.reddit.com/r/Kotlin/comments/jci3z3/hi_what_is_the_problem_dont_see_button_id_and/)
- url: https://i.redd.it/il6afq33tit51.jpg
---

## [8][Announcing Kofiko: Code-First Configuration library for Kotlin](https://www.reddit.com/r/Kotlin/comments/jbm9b9/announcing_kofiko_codefirst_configuration_library/)
- url: https://github.com/davidohana/kofiko-kotlin
---

## [9][.toUUID(): A tiny kotlin/jvm library for generating deterministic UUIDs for automated testing](https://www.reddit.com/r/Kotlin/comments/jbic8f/touuid_a_tiny_kotlinjvm_library_for_generating/)
- url: https://www.reddit.com/r/Kotlin/comments/jbic8f/touuid_a_tiny_kotlinjvm_library_for_generating/
---
Let's rip off the bandage: UUIDs is a pain to work with within the context of automated testing.

* Creating UUIDs look ugly
* Creating UUIDs takes up a lot of space
* Creating UUIDs decrease readability.

This might all be minor gripes I have, but I like my tests to be short and clean.

The most common way to generate a human-readable UUID is doing the following:

    UUID.fromString("00000000-0000-0000-0000-000000000001")

If you use .toUUID() you can get the same result like this:

    1.toUUID()

There's also a few more bells and whistles like generating UUIDs based on a collection, varargs, range of integers or even in a sequence.

If  you find yourself having to deal with UUIDs in automated tests, or you  want to generate a set of simple UUIDs quickly, then .toUUID() might be  up your alley:

[https://github.com/atomfinger/toUUID](https://github.com/atomfinger/toUUID)

*This post was also posted on* r/Java *a* [*few days ago*](https://www.reddit.com/r/java/comments/j9mv3x/touuid_a_tiny_library_for_generating/)*. If you find yourself sharing some of the same thoughts and opinions expressed in that thread, do head over to the* [*writeup I did in response to it*](https://github.com/atomfinger/toUUID/blob/main/why_touuid.md)*, which hopefully does a better job explaining the value of .toUUID().*
## [10][Need some help](https://www.reddit.com/r/Kotlin/comments/jbv4rv/need_some_help/)
- url: https://www.reddit.com/r/Kotlin/comments/jbv4rv/need_some_help/
---
I´m a beginner in programming and have a problem implementing the inmobi ads in my app, banner and instertitial are the ads that i´m trying to put in,  if someone can help i will be pleased

 some extra info: i´m venezuelan so my english is a bit orthopedic
