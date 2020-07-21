# Kotlin
## [1][An old unnoticed bug causes memory safety issue in Kotlin Native.](https://www.reddit.com/r/Kotlin/comments/hv5uqz/an_old_unnoticed_bug_causes_memory_safety_issue/)
- url: https://www.reddit.com/r/Kotlin/comments/hv5uqz/an_old_unnoticed_bug_causes_memory_safety_issue/
---
[KT-7972](https://youtrack.jetbrains.com/issue/KT-7972) is a bug that causes unsafe cast when smartcasting with generics. The following code

```kotlin
open class Base&lt;out T&gt;
class Derive&lt;T&gt;(var value: T) : Base&lt;T&gt;()

fun unsound() {
    val d = Derive(0)
    val b: Base&lt;Any&gt; = d
    if (b is Derive) {
        b.value = "0123456789"
    }
    println(d.value + 2)
}
```

will crashes on Kotlin JVM.

This bug is quite old and not be considered as a huge problem. Maybe it is not on JVM because the program would crash eventually if any incorrect cast involved. But when I run it with Kotlin Native targeting Linux x64, the program does not crash and output `12`. 

I guess the reason is that it unsafely cast `"0123456789"` to `Int` , and `d.value` reads the first 32 bits of the string object, which happens to be the length of the string.

Unsafe memory access may involve because of this bug, which is a big problem. I think we should pay more attention to this bug, either fix it at language level, or at least do more runtime check to prevent unsafe memory access on Kotlin Native.
## [2][Kotlin/Native Concurrency Changes…](https://www.reddit.com/r/Kotlin/comments/hutcjb/kotlinnative_concurrency_changes/)
- url: https://medium.com/@kpgalligan/kotlin-native-concurrency-changes-bbb1a5147e6
---

## [3][Kotlin/Native Memory Management Roadmap](https://www.reddit.com/r/Kotlin/comments/hup358/kotlinnative_memory_management_roadmap/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-native-memory-management-roadmap/
---

## [4][[X-POST] JetBrains is porting Compose to desktop and other platforms : androiddev](https://www.reddit.com/r/Kotlin/comments/huidyh/xpost_jetbrains_is_porting_compose_to_desktop_and/)
- url: https://www.reddit.com/r/androiddev/comments/hufnjq/_/
---

## [5][Can I use Kotlin serialization to serialize to very compact ByteArrays?](https://www.reddit.com/r/Kotlin/comments/huxjsz/can_i_use_kotlin_serialization_to_serialize_to/)
- url: https://www.reddit.com/r/Kotlin/comments/huxjsz/can_i_use_kotlin_serialization_to_serialize_to/
---
I have a use-case where I need to serialize data to very compact `ByteArray`s, in particular serializing `ByteArray`s properties of an object to compact representations - which need to fit in a UDP packet - max 1040 bytes.

I've been playing with CBOR, but it only seems marginally more efficient than JSON, `ByteArray`s are more than doubled in size, and - like JSON - the format includes property names in the serialized data.

[Avro](https://github.com/sksamuel/avro4k)'s `binary` serialization format looks promising, but the documentation isn't great, it's not clear how to configure it to do `binary` serialization to `ByteArray`s.

Has anyone dug into this?  Would appreciate any suggestions.
## [6][How do I check if a date minus a number of days is still in the future in Kotlin JVM?](https://www.reddit.com/r/Kotlin/comments/huyhvz/how_do_i_check_if_a_date_minus_a_number_of_days/)
- url: https://www.reddit.com/r/Kotlin/comments/huyhvz/how_do_i_check_if_a_date_minus_a_number_of_days/
---
I need to take a date (where the raw input is a string in the format `YYYY-MM-DD`) and an integer, where the integer represents the number of days *earlier* from that date, and determine if that new date is in the past, today, or the future.

Basically, how do I subtract a certain number of days from a date, and how do I determine if a date is today, or before or after today?
## [7][Is using named arguments in your function calls encouraged in kotlin?](https://www.reddit.com/r/Kotlin/comments/huy757/is_using_named_arguments_in_your_function_calls/)
- url: https://www.reddit.com/r/Kotlin/comments/huy757/is_using_named_arguments_in_your_function_calls/
---
I'm learning kotlin now and just learned about named arguments, where you can provide the argument name when calling a function. This seems like a lot, considering most modern IDEs should allow me to preview the arguments, and even if not this is something that shouldn't be too difficult to figure out using a ctrl+f if I have to.

    fun namedArguments() {
        printUserInfo(firstName = "Donn", lastName = "Felker", age = 32, isSunburned = false, likesMovies = true, lovesPopcorn = true)
    }
    
    fun printUserInfo(firstName: String, lastName: String, age: Int, isSunburned: Boolean, likesMovies: Boolean, lovesPopcorn: Boolean) {
        println("$firstName $lastName is of age $age. Sunburned: $isSunburned, Likes Movies: $likesMovies, Loves Popcorn: $lovesPopcorn")
    }

This is a concise example (from [this](https://caster.io/courses/kotlin-programming-language) great course)

Do you always explicitly name your argument when calling a function? I can see how it improves readability, but it is also pretty verbose.
## [8][Learn Kotlin by building Android AR app series by Radoslaw Fabisiak](https://www.reddit.com/r/Kotlin/comments/hundrn/learn_kotlin_by_building_android_ar_app_series_by/)
- url: https://www.reddit.com/r/Kotlin/comments/hundrn/learn_kotlin_by_building_android_ar_app_series_by/
---
In the first lesson, we will build the Android augmented reality mobile app&amp; we will learn how to setup Android Studio with ARCore for augmented reality projects.

In the second lesson of the Kotlin course, we will learn how to build a login form.

[Lesson1: Android Studio setup](https://itnext.io/learn-kotlin-by-building-android-ar-app-lesson1-android-studio-setup-d51a3568453a?source=friends_link&amp;sk=7ebd18d7243dde9d96a0cf6854643bcd) 

[Lesson2: How to build login form](https://itnext.io/learn-kotlin-by-building-android-ar-app-lesson2-how-to-build-login-form-725dcafbd00d?source=friends_link&amp;sk=634a0208d0ca97bdbd40d1bafdf51353)
## [9][Tips for newbies](https://www.reddit.com/r/Kotlin/comments/hukd3y/tips_for_newbies/)
- url: https://www.reddit.com/r/Kotlin/comments/hukd3y/tips_for_newbies/
---
                Hello people!
 
                I have not any experience at programming. I started to learn Java. But java seems to me difficult to learn. Then I tried Kotlin. It is looking not easy. But not hard as Java.

                I am watching Udemy videos. When I print "Hello World!" I was so excited. I said "That's magic!". I know it is a veeeeerryy basic thing but i liked that.

                My goal is coding budget app. There is so much and so good apps in Google Play but i want to build my own app.

                What do you advise for starters who know nothing?
## [10][Android app with the Room database](https://www.reddit.com/r/Kotlin/comments/hug5wx/android_app_with_the_room_database/)
- url: https://youtu.be/kF0TPehhIx0
---

