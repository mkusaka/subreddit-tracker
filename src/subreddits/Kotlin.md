# Kotlin
## [1][Construct object through parent sealed class, and child private constructor](https://www.reddit.com/r/Kotlin/comments/iyw1mo/construct_object_through_parent_sealed_class_and/)
- url: https://www.reddit.com/r/Kotlin/comments/iyw1mo/construct_object_through_parent_sealed_class_and/
---
Hi,

I want to implement some Scala inspired `Try` monad

I wanted to know if it is possible to make child class (`Success` and `Failure`) not instantiate by themselves (no public constructor), but only from the `Try.Companion.invoke` method. Here's my current code with public constructor for both child, when I make them private I got an error in `invoke` method (which I understand). I found nothing in the Kotlin documentation about this kind of thing and I don't think it is possible, so I post here just in case

    sealed class Try&lt;T&gt; {
        companion object {
            operator fun &lt;T&gt;invoke(elem: () -&gt; T) =
                try {
                    Success(elem())
                } catch (e: Exception) {
                    Failure&lt;T&gt;(e)
                }
        }
    }
    
    data class Success&lt;T&gt;(val value: T) : Try&lt;T&gt;()
    data class Failure&lt;T&gt;(val e: Exception) : Try&lt;T&gt;()

Thanks in advance :D
## [2][[HELP] Struggling at pair programming interviews](https://www.reddit.com/r/Kotlin/comments/iymjtq/help_struggling_at_pair_programming_interviews/)
- url: https://www.reddit.com/r/Kotlin/comments/iymjtq/help_struggling_at_pair_programming_interviews/
---
Hello everyone,

&amp;#x200B;

*TL;DR: trying to get a job but keep failing at pair programming interviews over and over. Don't know what to do and am considering leaving the tech industry to seek something else out of frustration.*

&amp;#x200B;

I'll start with a little bit of background so you can understand where I'm coming from. I started out as a web/java developer(4\~ years exp)  then became an Android developer (5+ years exp) and have led small sized Android teams over the last 2.5years. 

&amp;#x200B;

Unfortunately, I lost my job in late March. Over the last 3 months I have been aggressively looking for a job but I seem to keep failing in the technical department when it comes to pair programming interviews. Though, I must say that I feel a lot more confident when I'm given one of those take-home tests where I feel like I can excel without being judged or something but that doesn't seem to be the case nowadays as 90% of the companies I've interviewed with thus far didn't do it. Out of all the companies I interviewed with, 4 of them I managed to get to very last step of their hiring process. But like you'd guess, they turned me down in the end. I tried to practice more, going to those code challenge websites along with working on my own small project to try and keep my mind sharp but that also doesn't seem to help at all. With that being said, like the title says, I keep failing at pair programming interviews and I don't know what to do about it anymore and decided to take it to the internet.

&amp;#x200B;

As a matter of fact, I do understand how software architecture works and what purpose they serve (MVVM for instance) and I can talk about it for hours because I really find it interesting. Not to mention the new Google Jetpack APIs etc. But for some weird reason my mind blacks out when in that scenario.  


Lastly, what really intrigues me too is the fact that I see so many developers getting new jobs within weeks and I'm really starting to feel that I might've just "got lucky" so far in my career and that I might not be as good as I thought I was to begin with. I'm actually considering completely deserting the industry altogether to pursue something else instead.  


  
Any thoughts?
## [3][Cute trick for making sealed classes have a constructor/factory](https://www.reddit.com/r/Kotlin/comments/iywyvf/cute_trick_for_making_sealed_classes_have_a/)
- url: https://www.reddit.com/r/Kotlin/comments/iywyvf/cute_trick_for_making_sealed_classes_have_a/
---
EDIT: The code I posted didn't actually work, but /u/SkiFire13 posted a correct version:

    sealed class Foo {
        companion object {
            operator fun invoke(p1: String): Foo {
                return if (p1 == "bar") Bar() else Baz()
            }
        }
    }
    val f1 = Foo("bar") // is Bar
    val f2 = Foo("something") // is Baz

Instead of the typical approach of having a companion object with a factory method.
## [4][Dependency injection with function types has some shortcomings compared to interfaces](https://www.reddit.com/r/Kotlin/comments/iyfv3q/dependency_injection_with_function_types_has_some/)
- url: https://www.reddit.com/r/Kotlin/comments/iyfv3q/dependency_injection_with_function_types_has_some/
---
One thing you sometimes find is that a code base may have a Foo interface that has several methods. When you're writing a new function, you may only need one of those methods to operate:


    interface Foo {
        fun doSomethingWithInt(x: Int): Int?
        fun doSomethingWithString(x: String): String?
        fun doSomethingWithBar(x: Bar): Bar?
    }
    
    fun coolNewStringFeature(foo: Foo): String? {
        // stuff
        return foo.doSomethingWithString("string")
    }

Now, if you want to test your `coolNewStringFeature` you have to mock/stub the *whole* interface. Either by hand, or with some mock library like mockk. On the other hand, if you just wrote your function as:


    fun coolNewStringFeature(doSomethingWithString: (String) -&gt; String?): String? {
        // stuff
        return doSomethingWithString("string")
    }

Now you just have to write a trivial closure to test the function! And it's easy to quickly test a bunch of different inputs, instead of fiddling with your mocks or whatever.

But there are problems. 

Problem #1: Function types can't have arg labels.


    fun withRange(operation: (min: Int, max: Int) -&gt; Boolean) {
        operation(min = 0, max = 2) // compile error
        operation(0, 2) // which one means what?
    }

    interface MinMaxinator {
        fun operation(min: Int, max: Int) = //...
    }

    fun withRange(operator: MinMaxinator) {
        operator.operation(min = 0, max = 2) // okay!
    }

Problem #2: Functions with default args can't be used as function types without those args.


    fun importantWork(x: Int, y: Int, z: Int = 10) = //...

    fun doWith(work: (Int, Int) -&gt; Unit) = //...

    doWith(::importantWork) // compile error


Just putting this out there. No real question. Just generally complaining about stuff :)

I like the idea of passing functions around instead of objects, but I keep running into papercuts.
## [5][Artificial Intelligence in Kotlin: coding a planning optimization application from scratch](https://www.reddit.com/r/Kotlin/comments/iycc4e/artificial_intelligence_in_kotlin_coding_a/)
- url: https://www.youtube.com/watch?v=n6fl60gR8Gc
---

## [6][Kotlin - Room Architecture Components - Room Database with CRUD Operations](https://www.reddit.com/r/Kotlin/comments/iypvuw/kotlin_room_architecture_components_room_database/)
- url: https://www.rrtutors.com/SQLite-Database-with-CRUD-Operations
---

## [7][Create Hello World App with KMM 📱- Android &amp; IOS](https://www.reddit.com/r/Kotlin/comments/iy7hd5/create_hello_world_app_with_kmm_android_ios/)
- url: https://medium.com/@iamanbansal/create-hello-world-app-in-kmm-android-ios-79cc8c9bb84b
---

## [8][Can paging 3.0 be used to paginate data from an online db such as cloud firestore?](https://www.reddit.com/r/Kotlin/comments/iyjrvg/can_paging_30_be_used_to_paginate_data_from_an/)
- url: https://www.reddit.com/r/Kotlin/comments/iyjrvg/can_paging_30_be_used_to_paginate_data_from_an/
---
While building a chat application,i came across the need to paginate the amount of text messages between 2 users. Adding stack overflow link that contains the whole shebang below. How should i proceed? 

https://stackoverflow.com/questions/64027103/can-paging-3-0-be-used-to-paginate-data-from-an-online-db-such-as-cloud-firestor
## [9][SQLCipher and KMM](https://www.reddit.com/r/Kotlin/comments/iycdzh/sqlcipher_and_kmm/)
- url: https://medium.com/@kpgalligan/sqlcipher-and-kmm-58d96ea8095d
---

## [10][I am trying to create a terminal-like app in Kotlin using TornadoFX](https://www.reddit.com/r/Kotlin/comments/iy4q8x/i_am_trying_to_create_a_terminallike_app_in/)
- url: https://www.reddit.com/r/Kotlin/comments/iy4q8x/i_am_trying_to_create_a_terminallike_app_in/
---
I am currently not concerned with commands, or any sort of restriction concerning user input. I was hoping if somebody could point me out in as to what controls should I use to just have a prompt and the user input anything. Should I use a textfield and make it the same color in my styles files?
