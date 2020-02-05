# Kotlin
## [1][Opened the Call for Papers for KotlinConf 2020, until 31st of March](https://www.reddit.com/r/Kotlin/comments/ez8gpg/opened_the_call_for_papers_for_kotlinconf_2020/)
- url: https://sessionize.com/kotlinconf/
---

## [2][How to declare unitialized map with value type of Array&lt;String&gt;](https://www.reddit.com/r/Kotlin/comments/eyuz3n/how_to_declare_unitialized_map_with_value_type_of/)
- url: https://www.reddit.com/r/Kotlin/comments/eyuz3n/how_to_declare_unitialized_map_with_value_type_of/
---
Hi, general noob here. Having trouble declaring a mutable map with key type Char and value type Array&lt;String&gt; of size 10 (values unitialized). I think it should look something like the following:

`val example = mutableMapOf&lt;Char, Array&lt;String&gt;&gt;()`

Help greatly appreciated ⌨️
## [3][Kotlin Coroutines: What happens to the caller thread when context changes?](https://www.reddit.com/r/Kotlin/comments/eyjfz7/kotlin_coroutines_what_happens_to_the_caller/)
- url: https://www.reddit.com/r/Kotlin/comments/eyjfz7/kotlin_coroutines_what_happens_to_the_caller/
---
I have a doubt that even after a long time of studying coroutines, I never managed to understand it in a simple way. The doubt is about the difference between suspending and blocking code and how suspending functions relate to threads and coroutines.

If we define a function like this

    suspend fun wait() { // some suspending code here }

From what I understand, this function definition does not create new threads, nor coroutines. It's only meant to signal to the compiler that this function might suspend, so it should only be called from a suspendable context.

But if we declare it like this instead:

    suspend fun waitOnContext() = withContext(Dispatchers.IO) {
        // some suspending code here
    }

I believe (and correct if I'm wrong) that three things happens when a function defined like this is called:

1. Since the function is suspending (not blocking), the caller thread is free to do whaterver else it needs to do.
2. The sequential execution flow stops, meaning that the code after the function will only run after it completes, like a callback.
3. A new thread is created so that the code runs in the background.

How do the environment (I don't know which name to use here) decides what to do next with the caller thread? Does it uses an event loop, OS scheduling facilities or maybe the compiler organizes the code in such a way so that other things can be done while the new thread is running on the background.

One simple example is the Android main thread: How does the environment knows that it needs to run view rendering and gesture handling code after the new thread is created? And how does it knows to come back to the rest of the coroutine once the background thread finishes?

Another doubt that I have is if `CoroutineScope` is meant to group related coroutines so that they doesn't leak or if it's just meant to represent parent-child relationships so that when `Job` is cancelled, all of the children coroutines are cancelled as well.

One thing that made me confused was [this example](https://kotlinlang.org/docs/reference/coroutines/basics.html#scope-builder) from the Kotlin documentation. In it one `coroutineScope` is defined that suspends while waiting for a child coroutines to finish. Does this happen because the `CoroutineScope` needs to wait for all children coroutines to finish so it can be finished as well?

Am I correct in assuming that a `CoroutineScope` works similarly to a function scope in the sense that while a function needs to clear its local variables once it finishes, a `CoroutineScope` needs to clear it's children coroutines by waiting for them?

Another similarity is that when an exception happens inside a function, it has to clear all its local variables (stack unwinding), while a `CoroutineScope` needs to clear all it's coroutines by canceling them. Am I correct in using this analogy?
## [4][Kotlin Multiplatform header files (iOS)](https://www.reddit.com/r/Kotlin/comments/eyp9ue/kotlin_multiplatform_header_files_ios/)
- url: https://www.reddit.com/r/Kotlin/comments/eyp9ue/kotlin_multiplatform_header_files_ios/
---
So, I am working on a kotlin multiplatform project and I need to use some objective C classes for iOS part. I would like to know if there is a way to import and use Header (.h) files in kotlin MPP?
## [5][Kotlin for crossplatform iOS/Android/Web](https://www.reddit.com/r/Kotlin/comments/eynl35/kotlin_for_crossplatform_iosandroidweb/)
- url: https://www.reddit.com/r/Kotlin/comments/eynl35/kotlin_for_crossplatform_iosandroidweb/
---
I am currently developing alone a project available on iOS Android and Web, each project is coded natively in Java, swift and HTML and its very time consuming to maintain obviously
I heard about Kotlin Native but I can’t find big existing projects which run on these 3 platform

Is Kotlin reliable enough to recode all my project from scratch on all platforms ? I don’t use a lot of physical features, mostly graphics with REST API.

Do you know good tutorials for starting crossplatform development ?

Thanks
## [6][2020 Java Technology Report](https://www.reddit.com/r/Kotlin/comments/eynj9b/2020_java_technology_report/)
- url: https://www.jrebel.com/blog/2020-java-technology-report
---

## [7][Why is asSequence() not reporting the expected result?](https://www.reddit.com/r/Kotlin/comments/eyflvs/why_is_assequence_not_reporting_the_expected/)
- url: https://www.reddit.com/r/Kotlin/comments/eyflvs/why_is_assequence_not_reporting_the_expected/
---

    fun main(args: Array&lt;String&gt;) {
        val animals = listOf( "Cow", "Dog", "Cat", "Chicken", "Frog")
        val animal = animals.asSequence()
            .filter { it.startsWith("C" ) }
            .map { "Animal: $it" }
            .also { println("Intermediary result: ${it.count()}") }
            .take(1)

        println("Result: ${animal.count()}")
    }

This code results in outputting the line:

    Intermediary result: 3

I would expect (from my reading of asSequence()) that only one item from the list of *animals* should be processed at a time.
## [8][Why everything has to return something in Kotlin?](https://www.reddit.com/r/Kotlin/comments/ey5i2s/why_everything_has_to_return_something_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ey5i2s/why_everything_has_to_return_something_in_kotlin/
---
When doing the codelabs from the Kotlin Bootcamp for Programmers, I stumbled across the topic in lesson 3 where they talk about how everything in Kotlin returns something. 

I tested it and in fact you can nest the println() statements in Kotlin whereas JAVA just doesn't let you do that. However the output is a bit confusing to interpret for me. 

I ran this code : 

`println({print(print("This!"))})`

and got this as output : 

`() -&gt; kotlin.Unit`

and for statement :

`println(print(print("This!")))`

I got this as output : 

`This!kotlin.Unitkotlin.Unit`

I am trying to dig deeper to interpret the output (just started learning Kotlin).

I am curious as to if there are hazards, if any, to this approach, or is it just that Kotlin realized what JAVA missed out on.
## [9][Any suggestions for Kotlin certification?](https://www.reddit.com/r/Kotlin/comments/ey5pla/any_suggestions_for_kotlin_certification/)
- url: https://www.reddit.com/r/Kotlin/comments/ey5pla/any_suggestions_for_kotlin_certification/
---
Hello,

I’m a junior dev and I’m working on the server side with kotlin. I don’t have a CS degree and I wanted a certification to establish myself as a dev. I’m based in the UK and everything that I found (47degrees and jetBrains) is located in the USA. 

Is there any certification that I could read for/take classes and then take an exam? 

Thank you
## [10][Loading and caching images into ImageViews using one annotation](https://www.reddit.com/r/Kotlin/comments/ey6idr/loading_and_caching_images_into_imageviews_using/)
- url: https://www.reddit.com/r/Kotlin/comments/ey6idr/loading_and_caching_images_into_imageviews_using/
---
I created an annotation which will load image from a URL directly into an ImageView (caching it for future use) . It can also be used to show a placeholder image and animation until the image is downloaded into the device.

Usage :-

&gt;"@LoadImage(  
R.id.*image\_1*,  
 "https://images.unsplash.com/photo-1549740425-5e9ed4d8cd34?ixlib=rb-1.2.1&amp;w=1000&amp;q=80",  
 placeHolder = R.drawable.*loading*, enableLoadingAnimation = true  
)  
lateinit var imageWithAnimation: ImageView"

Check out the post for more details :-  [https://medium.com/@crypticmindscom\_5258/caching-made-easy-on-android-with-kotlin-part-4-18e7b066e9c2](https://medium.com/@crypticmindscom_5258/caching-made-easy-on-android-with-kotlin-part-4-18e7b066e9c2) 

Library :-  [https://github.com/crypticminds/ColdStorage](https://github.com/crypticminds/ColdStorage)
