# Kotlin
## [1][Kotlin Team AMA #2: Ask Us Anything](https://www.reddit.com/r/Kotlin/comments/ji9z19/kotlin_team_ama_2_ask_us_anything/)
- url: https://www.reddit.com/r/Kotlin/comments/ji9z19/kotlin_team_ama_2_ask_us_anything/
---
**UPDATE:** All done, thanks for all the questions! Note that you can always contact us via the Kotlin Slack or Twitter, so if you missed the AMA, feel free to post your questions there. If you’re yet not on Kotlin Slack, please request your invite here: [https://surveys.jetbrains.com/s3/kotlin-slack-sign-up](https://surveys.jetbrains.com/s3/kotlin-slack-sign-up)

Hi everyone! The JetBrains team is here and we will be answering your questions over the next \~36 hours:

* [/u/abreslav](https://www.reddit.com/u/abreslav/), Project Lead for Kotlin
* [/u/elizarov](https://www.reddit.com/u/elizarov/), Team Lead for Kotlin Language Research
* [/u/etolstoy](https://www.reddit.com/user/etolstoy/), Team Lead for Kotlin Product Management
* [/u/erokhins](https://www.reddit.com/user/erokhins), Development Lead for Kotlin
* [/u/svtk\_](https://www.reddit.com/u/svtk_/), Team Lead for Kotlin Advocacy
* [/u/udalov](https://www.reddit.com/user/udalov/), Team Lead for Kotlin/JVM
* [/u/qwwdfsad](https://www.reddit.com/user/qwwdfsad/), Team Lead for Kotlin Libraries
* [/u/dsavvinov](https://www.reddit.com/user/dsavvinov/), Team Lead for Kotlin Multiplatform
* [/u/terrakok](https://www.reddit.com/user/terrakok/), Team Lead for KMM (Kotlin Multiplatform Mobile)
* [/u/ekaterina-petrova](https://www.reddit.com/user/ekaterina-petrova/), Developer Advocate for KMM (Kotlin Multiplatform Mobile)
* [/u/akapanina](https://www.reddit.com/user/akapanina/), Product Manager for KMM (Kotlin Multiplatform Mobile)
* [/u/demiurg\_906](https://www.reddit.com/user/demiurg_906/), Software Developer for Kotlin Compiler
* [/u/sebi\_io](https://www.reddit.com/user/sebi_io/), Developer Advocate for Kotlin
* [/u/bashor\_](https://www.reddit.com/user/bashor_), Team Lead for Kotlin/JS
* [/u/EugeneAndroid](https://www.reddit.com/user/EugeneAndroid/), Product Manager for Kotlin for Android, Kotlin Build Tools
* [/u/fundamentalparticle](https://www.reddit.com/user/fundamentalparticle/), Developer Advocate for Kotlin Server-side
* [/u/hhariri](https://www.reddit.com/u/hhariri/), Team Lead for Ktor, Developer Advocate for Kotlin
* [/u/anton-yalyshev](https://www.reddit.com/user/anton-yalyshev/), Product Manager for Kotlin IDE
* [/u/mzarechenskiy](https://www.reddit.com/user/mzarechenskiy), Team Lead for Kotlin IDE, Compiler Core (Front-end), Language Specification
* [/u/yanex](https://www.reddit.com/user/yanex/), Software Developer for Kotlin IDE
* [/u/Artyom\_Degtyarev](https://www.reddit.com/user/Artyom_Degtyarev/), Support Engineer for Kotlin/Native
* [/u/SvyatoslavScherbina](https://www.reddit.com/user/SvyatoslavScherbina), Team Lead for Kotlin/Native
* [/u/\_Wo\_0dy\_](https://www.reddit.com/user/_Wo_0dy_/), Product Manager for Kotlin Website, Kotlin/JS Teams, and Kotlin Analytics
* [u/wild\_\_lynx](https://www.reddit.com/user/wild__lynx/), QA for Kotlin Multiplatform
* [/u/alinagrebenkina](https://www.reddit.com/user/alinagrebenkina), Team Lead for Kotlin Marketing
* [/u/Belosnegova](https://www.reddit.com/user/Belosnegova), Kotlin Marketing Manager for Education
* [/u/meilalina](https://www.reddit.com/user/meilalina), Kotlin Marketing Manager for Community and Kotlin Server-side
* [/u/Anisim\_1](https://www.reddit.com/user/Anisim_1), Kotlin Marketing Manager for KMM (Kotlin Multiplatform Mobile)
* [/u/yole](https://www.reddit.com/user/yole/), Product Manager for IntelliJ IDEA
* /u/ilya-g, Developer for Kotlin Standard Library and kotlinx-datetime

&amp;#x200B;

We’ll be joined by other developers and team leads from Kotlin, as well as two special guests: Wojtek Kaliciński /u/wkalicinski from Google and Sebastien Deleuze /u/sdeleuze from Spring.

On October 12–15 we hosted the [Kotlin 1.4 Online Event](https://kotlinlang.org/lp/event-14/), and over the course of those four days we received many interesting questions. Since we did not manage to cover all of them due to time constraints, we will answer them in a specific thread under this post.

Please post your questions as top-level comments to this post.

We look forward to hearing from you!
## [2][What do you think of this little extension function?](https://www.reddit.com/r/Kotlin/comments/jjmgb0/what_do_you_think_of_this_little_extension/)
- url: https://www.reddit.com/r/Kotlin/comments/jjmgb0/what_do_you_think_of_this_little_extension/
---
If you are a fan of chained calls, for example for string manipulation, you are probably sometimes as frustrated as I am because you have to "break the flow". I hit such a case again, but this time I decided to write an extension function so I can comfortably continue the chained call.

I actually don't know if it already exists in some form, but I haven't found it on the first look and it was so little code that it actually doesn't matter. If you actually know that this exists already, let me know.

So, for example let's say you have a filename (taken out of a list, for example) and you want to strip certain file endings (let's just say `.txt`, `.exe` and `.pdf`). But that's not the only thing you want to do. In fact, you already have a chained call with different string manipulations, and as the next step you want to strip the extension if it is one of the mentioned ones. Now the `removeSuffix()` method comes in handy, but the problem is that you don't have a _single_ suffix, but a _collection_ of suffixes. And _for each_ of these suffixes `x`, you need to call `removeSuffix(x)`. Normally this would mean that you have to break your chained call flow from

    string.callA()
        .callB()
        .callC()
        .callD()

To

    var res = string.callA()
        .callB()
        .callC()
        .callD()
    
    extensions.forEach {
        res = res.removeSuffix(it)
    }

Ugly, right? That's why I just patched together the extension function `forEachOf()`. Basically it is called with an arbitrary object as the receiver, but operates a lambda over every element of a given `Iterable`:

    inline fun &lt;T, U&gt; T.forEachOf(other: Iterable&lt;U&gt;, block: T.(U) -&gt; T): T {
        var result = this
        other.forEach {
            result = result.block(it)
        }
        return result
    }

Now I can perform the desired action without breaking the chained call, I can even continue it afterwards:

    string.callA()
        .callB()
        .callC()
        .callD()
        .forEachOf(extensions) { removeSuffix(it) }
        .callE()

Neat, right? Note that the lambda and the `forEachOf()` function (have to) return `T` to 1. be able to continue the chain and 2. deal with immutable classes like `String`.

What do you think? Is it useful? Would you use it? Can you imagine more practical use cases? Or would you improve the extension in some way?

For me it just shows the beauty of the expressiveness and extensibility of Kotlin. A similar problem in Java would be impossible to solve, but with Kotlin you have multiple ways of doing it!
## [3][Using withContext: is it possible to resume execution on the same thread in the same dispatcher?](https://www.reddit.com/r/Kotlin/comments/jjnuji/using_withcontext_is_it_possible_to_resume/)
- url: https://www.reddit.com/r/Kotlin/comments/jjnuji/using_withcontext_is_it_possible_to_resume/
---
Imagine a suspend function running on `Dispatchers.Default` similar to this one:

    suspend fun doSomething() {
        pre()
        withContext(myContext) {
            // do some heavy stuff here
            work()
        }
        post()
    }

According to the docs `withContex` will switch to a new thread if a new dispatcher is provided (which is the case here) and resume using the original dispatcher, `Dispatchers.Default`. However, I observe that sometimes it resumes and executes `post()` on a different thread than the one used to execute `pre()`.

This poses a challenge in a problem I need to solve, so I'm wondering whether it is possible to enforce the execution of `post()` and `pre()` on the same thread.
## [4][Is Coroutines turning into Rx?](https://www.reddit.com/r/Kotlin/comments/jj5lss/is_coroutines_turning_into_rx/)
- url: https://www.reddit.com/r/Kotlin/comments/jj5lss/is_coroutines_turning_into_rx/
---
One of the biggest reasons why Coroutines was so hyped compared with Rx is that it allowed to write async code in a synchronous way. But with recent trend of Flow and Channels, I feel like people are building their apps now heavily around these concepts, adding on top couple of aggregation operators on it. This reminds me all of the early days of Rx, where people used Subjects heavily.

What is the reason for the sudden popularity of flows? And how do they differ from Rx Subjects
## [5][Kotlin data type](https://www.reddit.com/r/Kotlin/comments/jj6qm8/kotlin_data_type/)
- url: https://www.reddit.com/r/Kotlin/comments/jj6qm8/kotlin_data_type/
---
Hi im new in kotlin and im really confused about float and double
What is the exact min and max value of float and double can have
There is so many diffrent value in internet and im confused
Thanks
## [6][Released kotlinx.coroutines 1.4.0](https://www.reddit.com/r/Kotlin/comments/jikkd1/released_kotlinxcoroutines_140/)
- url: https://github.com/Kotlin/kotlinx.coroutines/releases/tag/1.4.0
---

## [7][Usage of SharedFlow](https://www.reddit.com/r/Kotlin/comments/jj5kr0/usage_of_sharedflow/)
- url: https://coroutinedispatcher.com/posts/shared-flow/
---

## [8][Extend your code readability with Kotlin extensions](https://www.reddit.com/r/Kotlin/comments/jj3k1h/extend_your_code_readability_with_kotlin/)
- url: https://medium.com/androiddevelopers/extend-your-code-readability-with-kotlin-extensions-542bf702aa36
---

## [9][I published a small Android library written in Kotlin inspired by the uncaught exception widget from flutter. A screen that is shown when your app crashes and includes the crash details instead of the normal crash dialog. Should be used only in debug builds.](https://www.reddit.com/r/Kotlin/comments/jioz0c/i_published_a_small_android_library_written_in/)
- url: https://github.com/mlegy/red-screen-of-death
---

## [10][Multiplatform Java/JS examples?](https://www.reddit.com/r/Kotlin/comments/jimrqv/multiplatform_javajs_examples/)
- url: https://www.reddit.com/r/Kotlin/comments/jimrqv/multiplatform_javajs_examples/
---
I am trying to find an ideal example of using the MPP (yes, I know it's alpha) plugin to generate both JVM Maven package output and a NodeJS package of some kind from the same code base. In Gradle KTS format.

I've gone over dozens of projects I've found, and for the life of me can't quite figure out how they generate JS output into the right place, for NPM publishing, and also JVM output and a POM file for Maven publishing.

I do intend to use commonMain, jvmMain and jsMain stuff. Some types need to differ between platforms, etc.
## [11][Kotlin tutorials free download](https://www.reddit.com/r/Kotlin/comments/jicji8/kotlin_tutorials_free_download/)
- url: https://www.reddit.com/r/Kotlin/comments/jicji8/kotlin_tutorials_free_download/
---
If anyone can help, I need free download tutorials for Kotlin for beginners , (no torrent please).

I am learning and I want to create inspection checklist which saves my notes and can add remarks on each item inside list.

Would be great if you could help.
