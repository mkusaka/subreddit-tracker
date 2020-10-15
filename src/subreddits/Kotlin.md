# Kotlin
## [1][.toUUID(): A tiny kotlin/jvm library for generating deterministic UUIDs for automated testing](https://www.reddit.com/r/Kotlin/comments/jbic8f/touuid_a_tiny_kotlinjvm_library_for_generating/)
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
## [2][Announcing Tinlok - A standard library for Kotlin/Native on desktop platforms](https://www.reddit.com/r/Kotlin/comments/jb5uot/announcing_tinlok_a_standard_library_for/)
- url: https://tinlok.lotte.tf
---

## [3][Announcing Kofiko: Code-First Configuration library for Kotlin](https://www.reddit.com/r/Kotlin/comments/jbm9b9/announcing_kofiko_codefirst_configuration_library/)
- url: https://github.com/davidohana/kofiko-kotlin
---

## [4][First attempt at a StringFormat for KXS-properties](https://www.reddit.com/r/Kotlin/comments/jbkt4e/first_attempt_at_a_stringformat_for_kxsproperties/)
- url: https://www.reddit.com/r/Kotlin/comments/jbkt4e/first_attempt_at_a_stringformat_for_kxsproperties/
---
Recently, I have made [a post](https://www.reddit.com/r/Kotlin/comments/jb1dcm/is_there_any_documentation_for/) about kotlinx.serialization.properties and if there is any documentation. Turns out there apparently isn't, but the comments and my sourcecode-digging helped me out. The problem however is that KXS-properties is only capable of converting an object of an arbitrary model class into a **flat** `Map&lt;String, Any&gt;`, the keys already being in the correct format. Now there is still the last step needed, namely converting the map into a string that conforms the properties format (and vice versa). You might think that cannot be too hard, but it turns out that the [specification of the properties format](https://docs.oracle.com/javase/10/docs/api/java/util/Properties.html#load(java.io.Reader)) is a lot more intricate than one might think.

I found a way to avoid writing the entire encoding / decoding on my own, just by using `java.util.Properties` and the `load()` and `store()` methods. To make this quick solution available to everyone, I created a public gist. If you need this functionality in your code, just copy paste the following:

https://gist.github.com/RaphaelTarita/748e02c06574b20c25ab96c87235096d

However, there are two major problems with this solution:

1. The usage of `java.util.Properties` makes the solution JVM-dependent. And this kinda kills the idea of KXS, which is mainly to be multiplatform. I'm currently considering to actually implement the encoding / decoding on my own, but it's gonna be tedious.

2. Right now, all values are mapped to their `toString()` results. Normally, you'd have to inject all custom serializers at this point, but I have no idea how to do this so that it does not involve runtime reflection. Any help is greatly appreciated.

Anyways, for now this Gist exists, despite the problems it has. So if you need to serialize a model class and encode it to a stringified `.properties` file (or vice versa), feel free to make use of this snippet.
## [5][How can I give "generic" interface functions a concrete type on override?](https://www.reddit.com/r/Kotlin/comments/jb24nz/how_can_i_give_generic_interface_functions_a/)
- url: https://www.reddit.com/r/Kotlin/comments/jb24nz/how_can_i_give_generic_interface_functions_a/
---
I believe this shouldn't be a hard question, but I somehow can't come up with the clean solution I'm hoping for.

Basically, I'd like to have several interfaces with one function each, and several classes implementing any number of those interfaces. The thing is, that each implementing class is dealing with one specific type, but a different type each.

Here is a solution that would be working, but that I'd like to improve on:

    interface add &lt;E&gt; { fun add(entity: E) }
    
interface remove &lt;E&gt; { fun remove(entity: E) }
    
    class TypeA_Handler: add&lt;TypeA&gt; {
     //implementation of add()
}
    
    class TypeB_Handler: add&lt;TypeB&gt;, remove&lt;TypeB&gt; {
     //implementation of add() &amp; remove
}

What bothers me is that I will have quite a few such interfaces, and each handler will always use the same type. Being forced to explicitly state the same type on the implements statement seems cumbersume (and slightly error prone).

The interfaces' purpose here is to force the implementation of a certain set of functions (different combination for each handler) and those functions to have some uniformity across all implementing handlers (e.g. so that the function will always be "remove", and cannot be called "delete" instead).

&amp;#x200B;

In case you haven't guessed it already, this approach is inspired by the DAO Design Pattern. What I dislike about that pattern is that every DAO\_Impl has it's own interface with all the operations you want to allow for said DAO. Considering (at least in my case) DAOs have no need for non-private functions I want to mask via a interface and there's no reusability for a specific interface, that's seems to me like a whole bunch of nonsensical overhead.

So what I'm instead trying to to is define a interface for each operation, and have my DAOs implement all the interfaces I want to allow on it. Thus I need my interfaces to be generalized and my implementations to specify these generalizations. 

&amp;#x200B;

I'm really open to any kind of suggestion here. 

I've been looking into generics (see example), and into using `Any` instead of generics in my interfaces (which I couldn't override to a certain type in my handlers).   
While I'm also open to arguments in favor of the classic DAO Design pattern, this idea of overriding generalized types somehow peaked my interest independently of the thing I'm working on right now... \^\^
## [6][[100% OFF] Android and Kotlin From Zero to Full Professional +45 hours](https://www.reddit.com/r/Kotlin/comments/jbd20k/100_off_android_and_kotlin_from_zero_to_full/)
- url: https://unitedaca.com/100-off-android-y-kotlin-desde-cero-a-profesional-completo-45-horas/
---

## [7][Kotlin MPP: Unable to compile C bridges](https://www.reddit.com/r/Kotlin/comments/jb01c9/kotlin_mpp_unable_to_compile_c_bridges/)
- url: https://localazy.com/blog/kotlin-mpp-unable-to-compile-c-bridges
---

## [8][Start with just Android or KMM?](https://www.reddit.com/r/Kotlin/comments/jb4xcr/start_with_just_android_or_kmm/)
- url: https://www.reddit.com/r/Kotlin/comments/jb4xcr/start_with_just_android_or_kmm/
---
Hi, I want to develop an android app. I looked into some docs about Kotlin, and I found this KMM - Kotlin Multiplatform Mobile. Now I wonder, should I start with just android or start with KMM? Is it hard to use KMM and take an existing business logic from the project and implement it to iOS? Of course, it would be better to have both Apple and Android simultaneously, but I am not sure about the costs. I tried to find some discussions about it but I didn't find any. Thank a lot :)
## [9][Update yaml properties using yaml path?](https://www.reddit.com/r/Kotlin/comments/jb3q8l/update_yaml_properties_using_yaml_path/)
- url: https://www.reddit.com/r/Kotlin/comments/jb3q8l/update_yaml_properties_using_yaml_path/
---
I'm working on a project where I have the need to overwrite some yaml properties in a document.

Lets say I have something like the following

    image:
      name: repo/whoami:0.1
      pullPolicy: IfNotPresent

Then I'd like to update pullPolicy by issuing a string like the following

    image.pullPolicy=Always

For simple structures like the above this is fairly simple but I need to support all valid yaml structures. With lists and dictionaries it quickly becomes less simple.

Does anyone here know of a library which does the above?

Helm support this functionality through it's [--set flag](https://helm.sh/docs/chart_template_guide/values_files/). But that's programmed using go.

Update: To clarify, I'd like to avoid parsing "image.pullPolicy=Always" but rather leave that to a library since it easily get rather complex
## [10][Is Kotlin Multiplatform Mobile closed source?](https://www.reddit.com/r/Kotlin/comments/jb3h26/is_kotlin_multiplatform_mobile_closed_source/)
- url: https://www.reddit.com/r/Kotlin/comments/jb3h26/is_kotlin_multiplatform_mobile_closed_source/
---
I see KotlinMultiplatformPlayground and kmm-sample in the Jetbrains/Kotlin github pages but I don't see the code base for KMM. Is it not open-source?
