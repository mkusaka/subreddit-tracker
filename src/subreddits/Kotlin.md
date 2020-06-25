# Kotlin
## [1][Announcing Lingua 1.0.0: The most accurate natural language detection library for the JVM](https://www.reddit.com/r/Kotlin/comments/hf59dc/announcing_lingua_100_the_most_accurate_natural/)
- url: https://github.com/pemistahl/lingua
---

## [2][How do I deep copy a Kotlin object?](https://www.reddit.com/r/Kotlin/comments/hfcrzu/how_do_i_deep_copy_a_kotlin_object/)
- url: https://www.reddit.com/r/Kotlin/comments/hfcrzu/how_do_i_deep_copy_a_kotlin_object/
---
I have a method that takes a MutableList of objects and needs return a MutableList of the same objects but with significant modifications. However, the original versions of the objects are still needed. Running the method on objects that have already been modified would generate invalid data. WHat's the best way to deep copy the list object so I can make modifications without affecting the originals?
## [3][Kotlin/Native extra imports](https://www.reddit.com/r/Kotlin/comments/hf59n1/kotlinnative_extra_imports/)
- url: https://www.reddit.com/r/Kotlin/comments/hf59n1/kotlinnative_extra_imports/
---
I'm trying to import an existing swift library of ours, and while it exports to obj-c just fine, when kotlin goes to generate headers it adds in some classes related to UIPointer and an import for UIAxis, which then cant be resolved. These aren't used anywhere in our swift library, and I'm not sure how to get rid of them. Using Kotlin 1.3.72, Xcode 11.5, targeting iOS 13.4 (i've tried targeting older to see if that would fix it. it did not). Current best guess is that kotlin/native is importing the fields from the UIKit header in the swift library, but not mapping these particular ones across to the provided platform library correctly. Is there a way I can force that to regenerate?
## [4][Discussion: Do you think this would be a cool feature?](https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/)
- url: https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/
---
TL;DR let the Kotlin compiler generate the boilerplate code for reifying typeargs in order to circumvent type erasure. Or, in other words, a `reified` keyword for class / non-inline function typeargs.

Working with Java (and reflection) a lot, I frequently feel the urge of banging my head against the table because of type erasure. Kotlin having a `reified` keyword for inline function typeargs is already a helpful thing, but I feel like there could be done more. In order to manually reify typeargs, I found myself using this boilerplate pattern:

1. Store a `Class&lt;?&gt;` instance in the class (this usually requires passing that instance to the constructor)
2. Perform instance checks using `clazz.isInstance(obj)`
3. Perform an unchecked cast to the requested typearg.

Kotlin is already known as a language that mostly uses syntactic sugar to improve Java, while still ensuring interoperability. Therefore it would be quite fitting if Kotlin provided an universally usable `reified` keyword that hides the boilerplate pattern from you and lets you use your typearg as if it wasn't subject to type erasure. So you could do

```
var inner: T
// ...
fun set(newval: Any) {
    if (newval is T) {
        inner = newval // smartcast
    }
}
```

Which would actually store the KClass representation of T and performs an isInstance check in the background.

What are your thoughts? Is this a good idea? Is it useful? How would you improve it? And ultimately, do you think it should be a feature request on YouTrack?
## [5][Kotlin Coroutines vs Java Threads](https://www.reddit.com/r/Kotlin/comments/hecrel/kotlin_coroutines_vs_java_threads/)
- url: https://piotrminkowski.com/2020/06/23/kotlin-coroutines-vs-java-threads/
---

## [6][SQLDelight 1.4.0 released](https://www.reddit.com/r/Kotlin/comments/hdvn5g/sqldelight_140_released/)
- url: https://github.com/cashapp/sqldelight/releases/tag/1.4.0
---

## [7][variable name vs type position](https://www.reddit.com/r/Kotlin/comments/hedgeg/variable_name_vs_type_position/)
- url: https://www.reddit.com/r/Kotlin/comments/hedgeg/variable_name_vs_type_position/
---
Hey everyone, trying to learn Kotlin, coming from Java.

Why did Kotlin decide to switch the variable name and the type positions around? Is there a good reason for this?  
I find it so much easier to write code like this:

Dog dog; (Something like D&lt;ctrl-space&gt;&lt;space&gt;&lt;ctrl-space&gt;&lt;enter&gt;)

vs.

dog: Dog (Where the autocomplete is only available towards the end of the line)

What advantages are there to do it the Kotlin way?
## [8][What's the best way to check if something is null in Kotlin?](https://www.reddit.com/r/Kotlin/comments/he5792/whats_the_best_way_to_check_if_something_is_null/)
- url: https://www.reddit.com/r/Kotlin/comments/he5792/whats_the_best_way_to_check_if_something_is_null/
---
Should I use `variable == null`, `variable === null` or something else?
## [9][Where can I see a list of the built-in exceptions in Kotlin?](https://www.reddit.com/r/Kotlin/comments/he6nv6/where_can_i_see_a_list_of_the_builtin_exceptions/)
- url: https://www.reddit.com/r/Kotlin/comments/he6nv6/where_can_i_see_a_list_of_the_builtin_exceptions/
---
I like to throw exceptions built into the language with a custom message explaining the error when possible, instead of making my own exceptions every time. Howver, I couldn't find a definitive list of all the exceptions Kotlin defines in the standard library and their usages. Does such a list exist? Is this way of throwing exceptions even a good idea?
## [10][New Dokka - Designed for Fearless Creativity](https://www.reddit.com/r/Kotlin/comments/hdxgj4/new_dokka_designed_for_fearless_creativity/)
- url: https://www.youtube.com/watch?v=OvFoTRhqaKg
---

