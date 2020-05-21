# Kotlin
## [1][Introducing Gryphon, the Swift to Kotlin translator](https://www.reddit.com/r/Kotlin/comments/gnm5ed/introducing_gryphon_the_swift_to_kotlin_translator/)
- url: https://www.reddit.com/r/Kotlin/comments/gnm5ed/introducing_gryphon_the_swift_to_kotlin_translator/
---
Hi all, I’ve just published the first full version Gryphon, a program that translates Swift code into Kotlin code. It is meant to let developers share their iOS app's code with Android.

I encourage anyone interested to [check out the video](https://twitter.com/gryphonblog/status/1263233844519620610?s=20) and [the website](https://vinivendra.github.io/Gryphon/index.html).

This is the first version where all features are working:

* The generated Kotlin code works the same as the Swift code it came from, no edits needed. There’s an automated test that makes Gryphon translate its own source code, with around 12k lines, and ensures the translation passes all the same tests as the original executable.
* There’s a templates system that’s used to automatically translate many standard library types and methods.
* Xcode integration, with Xcode showing any warnings or errors raised by the Kotlin compiler next to the Swift lines that originated them, so users can fix them at the source.
* The output code is readable, which is useful for minimizing the risks for new users. If Gryphon isn’t the right fit for them, they’re left with a Kotlin codebase they can keep maintaining.

This list of features isn’t to say [there aren’t any bugs](https://github.com/vinivendra/Gryphon/issues), of course - I’m doing my best to keep track of them and fix them on a daily basis.

Any questions, comments, or feedback in general is welcome!
## [2][Scripting in Python vs Kotlin](https://www.reddit.com/r/Kotlin/comments/gnvjya/scripting_in_python_vs_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gnvjya/scripting_in_python_vs_kotlin/
---
How easy and supported is scripting and automating daily stuff in kotlin
## [3][I am loving Kotlin, but...](https://www.reddit.com/r/Kotlin/comments/gnhmbb/i_am_loving_kotlin_but/)
- url: https://www.reddit.com/r/Kotlin/comments/gnhmbb/i_am_loving_kotlin_but/
---
I don't know Java

  


I started learning Android development some months ago, and I don't know Java, so I went straight to Kotlin. The problem is, everytime a problem comes up, I look for a solution on Stack Overflow and there it is, a solution in Java. So I try to apply it, and sometimes it works, but when it don't I have to make a new question on Stack Overflow, and mostly they mark my post as duplicate and send me to other solution... In Java.
## [4][[Help] Can anyone explain why this code sample about coroutines from Kotlinlang.org works the way it does?](https://www.reddit.com/r/Kotlin/comments/gnozxp/help_can_anyone_explain_why_this_code_sample/)
- url: https://www.reddit.com/r/Kotlin/comments/gnozxp/help_can_anyone_explain_why_this_code_sample/
---
Hello everyone. 

I'm trying to understand coroutines in Kotlin. While following up the official docs, I came across a code sample I am unable to make sense of. 

Here's the [code sample](https://kotlinlang.org/docs/reference/coroutines/basics.html#coroutines-are-light-weight)

    import kotlinx.coroutines.*

    fun main() = runBlocking {
        repeat(100_000) { // launch a lot of coroutines
            launch {
                delay(1000L)
                print(".")
            }
        }
    }


What's happening is: the program runs, shows nothing for one second, and then prints `.` 1,00,000 times one after another with no delay at all!

I think this should not be happening, and I know I'm wrong but don't know where.

To me, for an iteration let's say first iteration (it = 0):

1. `launch` should be building a coroutine
2. Then that particular coroutine should be suspended for 1 second
3. Then `.` should be printed as output. 
4. Meanwhile, when coroutine built during first iteration was suspended, repeat loop would go to next iteration and the aforesaid procedure in steps 1-3 should be repeated.

I am expecting synchronous code execution here. So for every dot to be printed, there has to have a delay of 1 second. But the overall delay during execution is 1 second only and they all get printed than!

Can you please show me the flaw in my logic?
## [5][Tell me why Kotlin is more than the current Fad](https://www.reddit.com/r/Kotlin/comments/gnucq9/tell_me_why_kotlin_is_more_than_the_current_fad/)
- url: https://www.reddit.com/r/Kotlin/comments/gnucq9/tell_me_why_kotlin_is_more_than_the_current_fad/
---
Hi! Im a Java Developer with JS Knowledge looking for a third language to learn.
Im choosing between Kotlin and Python.
I need it for smaller stuff, scripting and more modern coding, but im afraid kotlin will just go away once google embraces flutter
## [6][Is Ktor production ready?](https://www.reddit.com/r/Kotlin/comments/gnmkn2/is_ktor_production_ready/)
- url: https://www.reddit.com/r/Kotlin/comments/gnmkn2/is_ktor_production_ready/
---
If I'm making a new website for a business, should I use Ktor? How stable is it and how likely is it to see sustained development?

Part of the source code is already in Java, so I thought that by rewriting the site in Kotlin (which I prefer to Java) I can reuse some of the old Java code by converting it to Kotlin using Intellij's converter tool. Is this actually a good idea?

Finally, does Ktor support rendering HTML on the server side with a template language?
## [7][Get top tracks from last.fm](https://www.reddit.com/r/Kotlin/comments/gnhxsf/get_top_tracks_from_lastfm/)
- url: https://www.reddit.com/r/Kotlin/comments/gnhxsf/get_top_tracks_from_lastfm/
---
Hey, guys.

I wanted to start learning Kotlin and I created app which lets You generate Your top yearly playlists from [last.fm](https://last.fm) in spotify. Any feedback is appreciated. 

&amp;#x200B;

 [https://github.com/lisu188/spotify-web-api-demo](https://github.com/lisu188/spotify-web-api-demo) 

 [https://spotify-web-api-demo.herokuapp.com/](https://spotify-web-api-demo.herokuapp.com/)
## [8][Diffing with annotations and Kotlin](https://www.reddit.com/r/Kotlin/comments/gng62u/diffing_with_annotations_and_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gng62u/diffing_with_annotations_and_kotlin/
---
I wrote a small project to help with diffing between states in a MVI type setup.

https://github.com/iFanie/Stateful

The processor generates a wrapper and a listener interface for the annotated model. The wrapper invokes the listener only when there is a change in the individual properties of the model.

The listener is very granular, with 4 different functions for each property (new value, old &amp; new value, new model, old &amp; new model), all of which have default implementations to avoid clutter.

This is just a proof of concept at the moment so any and all feedback would be appreciated. Cheers :)
## [9][Creating an object in Kotlin](https://www.reddit.com/r/Kotlin/comments/gni8ls/creating_an_object_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gni8ls/creating_an_object_in_kotlin/
---
So I come from java and I'm still quite new to Kotlin and have been using it on and off for a few months now. 

Only recently have I been able to write objects in Kotlin without looking it up on Google. I don't why but it didn't stuck in my head that Kotlin doesn't have the "new" keyword. 

I kept forgetting how to write new objects and it was even to the point that I was getting frustrated because let's be honest, it's not hard at all. 

I'm good now but is there anyone else that had this problem?
## [10][GraalVM 20.1 comes with improved Kotlin coroutine support](https://www.reddit.com/r/Kotlin/comments/gn63sm/graalvm_201_comes_with_improved_kotlin_coroutine/)
- url: https://medium.com/graalvm/graalvm-20-1-7ce7e89f066b
---

