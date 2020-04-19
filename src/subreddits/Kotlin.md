# Kotlin
## [1][Does literally anyone besides me use Kotlin for development outside of Android? (JVM Kotlin).](https://www.reddit.com/r/Kotlin/comments/g3r6yr/does_literally_anyone_besides_me_use_kotlin_for/)
- url: https://www.reddit.com/r/Kotlin/comments/g3r6yr/does_literally_anyone_besides_me_use_kotlin_for/
---
Anytime I look for new libraries in Kotlin or fixes, they always have fixes for Android Studio or the android library system. I love Kotlin, as I know both Java and Kotlin, Kotlin in every way is better than Java. But, I feel like Kotlin is just used for Android.

&amp;#x200B;

Edit: I obviously know Kotlin is used heavily, it was more I was curious of where in the real world Kotlin is being used outside of Android. Most everyone understood this and replied with their experiences &amp; thoughts on Kotlin.   
## [2][VKUG - Fight Complexity With Functional Programming](https://www.reddit.com/r/Kotlin/comments/g457o1/vkug_fight_complexity_with_functional_programming/)
- url: https://youtu.be/uGxx01yYAgk
---

## [3][Some improvements in Eclipse Kotlin Plugin](https://www.reddit.com/r/Kotlin/comments/g46ngy/some_improvements_in_eclipse_kotlin_plugin/)
- url: https://www.reddit.com/r/Kotlin/comments/g46ngy/some_improvements_in_eclipse_kotlin_plugin/
---
Usually if I open .kt file in eclipse, mostly it will hang

But the new version  [https://marketplace.eclipse.org/content/kotlin-plugin-eclipse](https://marketplace.eclipse.org/content/kotlin-plugin-eclipse) 

Seems to improved a bit and the files are opening without much hanging
## [4][Coroutines and exceptions: things to know](https://www.reddit.com/r/Kotlin/comments/g3mbju/coroutines_and_exceptions_things_to_know/)
- url: https://www.rockandnull.com/coroutines-and-exceptions/
---

## [5][Kotlin Microservices With Spring Boot And Spring Cloud: Part 1 - Intro to Spring Boot](https://www.reddit.com/r/Kotlin/comments/g3wpdz/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://www.youtube.com/watch?v=KnbH4hOswLA
---

## [6][JSON library](https://www.reddit.com/r/Kotlin/comments/g3p86q/json_library/)
- url: https://www.reddit.com/r/Kotlin/comments/g3p86q/json_library/
---
Hello you all, I have done a library for json manipulation [https://github.com/Barddo/kson](https://github.com/Barddo/kson) it's a Kotlin friendly fork of the org.json. Feel free to open issues and pull requests.
## [7][DDD Domain entity as data class?](https://www.reddit.com/r/Kotlin/comments/g3tghh/ddd_domain_entity_as_data_class/)
- url: https://www.reddit.com/r/Kotlin/comments/g3tghh/ddd_domain_entity_as_data_class/
---
As I understand the primary use case for data classes is DTOs, classes which represent API response or a response of an external service, etc. These classes are simply data holders and they don't need any behavior attached to them. This is an obvious use case.

What about domain entities? They are supposed to have both data and behavior. Is it okay to define them as data class or that's semantically incorrect?

(same question could stand for Java 14 records, although they are more limited, so maybe not a proper comparison...)

EDIT1:

Example:

    data class Person(
        private var name: String,
        val age: Int
    ) {
        fun isAdult() = age &gt;= 18
    
        fun changeName(newName: String) {
            this.name = newName
        }
    }

EDIT2:

Actually, there is a detekt rule which should fail on data classes which have custom functions (if I understand correctly):  [https://arturbosch.github.io/detekt/style.html#dataclasscontainsfunctions](https://arturbosch.github.io/detekt/style.html#dataclasscontainsfunctions)
## [8][Multiple Kotlin Frameworks in an Application](https://www.reddit.com/r/Kotlin/comments/g3ssyv/multiple_kotlin_frameworks_in_an_application/)
- url: https://touchlab.co/multiple-kotlin-frameworks-in-application/
---

## [9][Micronaut JWT auth mock](https://www.reddit.com/r/Kotlin/comments/g3opms/micronaut_jwt_auth_mock/)
- url: https://www.reddit.com/r/Kotlin/comments/g3opms/micronaut_jwt_auth_mock/
---
It took me several hours to figure out how to mock the JWT authorization in my Micronaut application tests, but finally I found a way to do it. So that you don't waste too much time on the same problem I have published a sample project on Github. I hope there are some people I can help with this :)   


[https://github.com/lstoeferle/micronaut-jwt-auth-mock](https://github.com/lstoeferle/micronaut-jwt-auth-mock)
## [10][MVI in Kotlin Multiplatform — part 1](https://www.reddit.com/r/Kotlin/comments/g35sle/mvi_in_kotlin_multiplatform_part_1/)
- url: https://badootech.badoo.com/mvi-in-kotlin-multiplatform-part-1-1-of-3-205c6feb4ac7
---

