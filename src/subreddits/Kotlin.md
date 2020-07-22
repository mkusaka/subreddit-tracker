# Kotlin
## [1][Kotlin Microservices On Kubernetes: Part 6 - Service mesh](https://www.reddit.com/r/Kotlin/comments/hvql7d/kotlin_microservices_on_kubernetes_part_6_service/)
- url: https://www.youtube.com/watch?v=TYpRzK03N80&amp;feature=share
---

## [2][Immutability we can afford](https://www.reddit.com/r/Kotlin/comments/hvtlzd/immutability_we_can_afford/)
- url: https://medium.com/@elizarov/immutability-we-can-afford-10c0dcb8351d
---

## [3][Squasher](https://www.reddit.com/r/Kotlin/comments/hvtl4k/squasher/)
- url: https://dkandalov.github.io/squasher
---

## [4][Kotlin’s Noinline &amp; Crossline, once for all](https://www.reddit.com/r/Kotlin/comments/hvt4vh/kotlins_noinline_crossline_once_for_all/)
- url: https://medium.com/@cortinico/kotlins-noinline-crossline-once-for-all-c942fd07b7a3
---

## [5][What's up with Kotlin/native?](https://www.reddit.com/r/Kotlin/comments/hvhpvd/whats_up_with_kotlinnative/)
- url: https://www.reddit.com/r/Kotlin/comments/hvhpvd/whats_up_with_kotlinnative/
---
As a .NET developer, I love Kotlin. It allows me to take advantage of the JVM ecosystem without THAT language, and also introduce it to people at work(I love Scala as well, but selling it to the average corporate dev is a different story).     
That being said, Kotlin is obviously a JVM language. All the primitive wrapper magic, type erasure, collection wrappers, lack of value types and so on scream "I'm made to fix Java's crap".      
My question is, when do you actually use Kotlin/Native? For low-level stuff/SP, Rust is amazing and very expressive and fast. For mobile development, I'd rather use the native languages. For tools, well, I can use F# or Python to drop some quick scripts. For me, it's a bit weird to offer a native implementation of a language spec that's realistically stunted by the JVM limitations.
## [6][Do you consider using labels in kotlin bad practice?](https://www.reddit.com/r/Kotlin/comments/hvq24v/do_you_consider_using_labels_in_kotlin_bad/)
- url: https://www.reddit.com/r/Kotlin/comments/hvq24v/do_you_consider_using_labels_in_kotlin_bad/
---
[SOLVED]

Hey!Im going through the kotlin reference right now and i came across the section break and continue labels a while ago.I tried to find some usage examples in code of others but i didnt found any,so it came to my mind that using labels in kotlin could be considered bad practice just like in other languages.I wonder what you guys think about it.Is it bad practice?Is it just not used so much but actually even  encouraged to use it??If you got production code experience with kotlin i would like to know about your experience and view onto the use of labels.
## [7][An old unnoticed bug causes memory safety issue in Kotlin Native.](https://www.reddit.com/r/Kotlin/comments/hv5uqz/an_old_unnoticed_bug_causes_memory_safety_issue/)
- url: https://www.reddit.com/r/Kotlin/comments/hv5uqz/an_old_unnoticed_bug_causes_memory_safety_issue/
---
[KT-7972](https://youtrack.jetbrains.com/issue/KT-7972) is a bug that causes unsafe cast when smartcasting with generics. The following code

```kotlin
open class Base&lt;out T&gt;
class Derive&lt;T&gt;(var value: T) : Base&lt;T&gt;()

fun main() {
    val d = Derive(0)
    val b: Base&lt;Any&gt; = d
    if (b is Derive) {
        b.value = "0123456789"
    }
    println(d.value + 2)
}
```

will crash on Kotlin JVM.

This bug is quite old and not be considered as a huge problem. Maybe it is not on JVM because the program would crash eventually if any incorrect cast involved. But when I run it with Kotlin Native targeting Linux x64, the program does not crash and output `12`. 

I guess the reason is that it unsafely cast `"0123456789"` to `Int` , and `d.value` reads the first 32 bits of the string object, which happens to be the length of the string.

Unsafe memory access may involve because of this bug, which is a big problem. I think we should pay more attention to this bug, either fix it at language level, or at least do more runtime check to prevent unsafe memory access on Kotlin Native.

=========================

Edit: the reason it compiles is that, `d` has type `Derive&lt;Int&gt;`, which is a subtype of `Base&lt;Int&gt;` because `Derive` is a subclass of `Base`, which is a subtype of `Base&lt;Any&gt;` because `T` in `Base&lt;out T&gt;` is covariance. Expressing by code,

```kotlin
val d: Derive&lt;Int&gt; = Derive(0)
val tmp: Base&lt;Int&gt; = d
val b: Base&lt;Any&gt; = tmp
```
## [8][An example of stateful sequence operation](https://www.reddit.com/r/Kotlin/comments/hvbixu/an_example_of_stateful_sequence_operation/)
- url: https://www.reddit.com/r/Kotlin/comments/hvbixu/an_example_of_stateful_sequence_operation/
---
I am currently studying Kotlin [Sequences - Sequence Operations](https://kotlinlang.org/docs/reference/sequences.html#sequence-operations)  
Is there an example to demonstrate a stateful sequence operation?
## [9][Android Model-View-Intent with the new Kotlin StateFlow! - Replacing LiveData](https://www.reddit.com/r/Kotlin/comments/hvdkda/android_modelviewintent_with_the_new_kotlin/)
- url: https://proandroiddev.com/android-model-view-intent-with-kotlin-flow-ca5945316ec
---

## [10][Best non-mutating way to add an element to a Map&lt;T, List&lt;R&gt;&gt;?](https://www.reddit.com/r/Kotlin/comments/hv7tii/best_nonmutating_way_to_add_an_element_to_a_mapt/)
- url: https://www.reddit.com/r/Kotlin/comments/hv7tii/best_nonmutating_way_to_add_an_element_to_a_mapt/
---
So, I have a `Map&lt;T, List&lt;R&gt;&gt;` and I have elements `val t: T = foo()` and `val r: R = bar()`.

I want to return a new map with the list at key `t` replaced with a list that is the original + `r`. 

What's the most elegant way to do so?
