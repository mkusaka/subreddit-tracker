# Kotlin
## [1][I've created a project that shows how you can write Grease-/Tamper-/Violent)-monkey user scripts with KotlinJS](https://www.reddit.com/r/Kotlin/comments/jd0guk/ive_created_a_project_that_shows_how_you_can/)
- url: https://github.com/Foso/gtvmonkey-scripts#greasetamperviolent-monkey-scripts-with-kotlinjs
---

## [2][Making my Kotlin code more "Kotlin-y"](https://www.reddit.com/r/Kotlin/comments/jcyd2s/making_my_kotlin_code_more_kotliny/)
- url: https://www.reddit.com/r/Kotlin/comments/jcyd2s/making_my_kotlin_code_more_kotliny/
---
Just a quick question about syntax and the difference from Java.

I've recently written my first actual project in Kotlin and plan to release it to the Play Store, but I'm concerned that, since Kotlin also supports Java syntax, that my code isn't as 'Kotlin-y' as it could be. Is there some service that can scan my code and reformat it to more Kotlin-accurate ssyntax, like how the IntelliJ IDEA converts Java files to Kotlin files? I'd like to see how it should look for the sake of learning. 

Thanks in advance.
## [3][Kweb 0.7.33 released, Kotlin web framework for backend developers adds support for FileReader / file uploads](https://www.reddit.com/r/Kotlin/comments/jcscyo/kweb_0733_released_kotlin_web_framework_for/)
- url: https://github.com/kwebio/kweb-core
---

## [4][Resources to learn Kotlin?](https://www.reddit.com/r/Kotlin/comments/jczxx6/resources_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jczxx6/resources_to_learn_kotlin/
---
I’m mainly a C++ / Haskell developer, plus I know Python and am planning to learn Rust (probably before Kotlin), but I have zero experience with JVM languages.

For the purposes of Android development, I want to learn Kotlin. Can you please recommend any good resources to do that (books, free online courses etc.), **preferably without touching Java** (if at all possible)?
## [5][[Academic] Software Development Survey (Software Developers)](https://www.reddit.com/r/Kotlin/comments/jcz9md/academic_software_development_survey_software/)
- url: /r/SampleSize/comments/jbsju0/academic_software_development_survey_software/
---

## [6][http4k blog: Nanoservices: The Power of Composition](https://www.reddit.com/r/Kotlin/comments/jc8tez/http4k_blog_nanoservices_the_power_of_composition/)
- url: https://www.http4k.org/blog/nanoservices/
---

## [7][Kotlin 1.4 Online Event Recap](https://www.reddit.com/r/Kotlin/comments/jcbdi9/kotlin_14_online_event_recap/)
- url: https://link.medium.com/OzC2ERGhDab
---

## [8][Accessing class member name via interpolation of number?](https://www.reddit.com/r/Kotlin/comments/jcimht/accessing_class_member_name_via_interpolation_of/)
- url: https://www.reddit.com/r/Kotlin/comments/jcimht/accessing_class_member_name_via_interpolation_of/
---
I just want to state that I know I'm doing this in a dumb way. I'm a newb.

I have a class paragraphNumber that holds 20 different paragraphText. paragraphText1 all the way to paragraphText20. 

In my MainActivity.kt, I have a button. This button figures out which paragraph it should be loading, and that's all it does. I have some logic to figure out the next paragraph (which isn't the next one in line!) 

How can I interpolate a number into the name? I mean, now I have

rollButton.setOnClickListener **{** textView1.*text* = paragraphNumber().**paragraphText1}**

but, I would rather be able to call on paragraphTextX where X would be a number 1 - 20. Is this possible? Please let me know if I'm not making any sense.
## [9][OutOfMemoryError related JVM Arguments](https://www.reddit.com/r/Kotlin/comments/jc73ux/outofmemoryerror_related_jvm_arguments/)
- url: https://blog.heaphero.io/2019/06/21/outofmemoryerror-related-jvm-arguments/
---

## [10][How do you add a KLib to a Native Gradle Project?](https://www.reddit.com/r/Kotlin/comments/jcdgc6/how_do_you_add_a_klib_to_a_native_gradle_project/)
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
