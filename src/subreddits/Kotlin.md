# Kotlin
## [1][The Power of Types](https://www.reddit.com/r/Kotlin/comments/etpck5/the_power_of_types/)
- url: https://www.youtube.com/watch?v=t3DBzaeid74
---

## [2][What I learned from Kotlin Flow API](https://www.reddit.com/r/Kotlin/comments/eta1en/what_i_learned_from_kotlin_flow_api/)
- url: https://www.coroutinedispatcher.com/2020/01/what-i-learned-from-kotlin-flow-api.html
---

## [3][Integrating Github Actions for Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/etczp8/integrating_github_actions_for_kotlin/)
- url: https://www.alecstrong.com/2020/01/github-actions-mpp/
---

## [4][On overloading](https://www.reddit.com/r/Kotlin/comments/ete8cq/on_overloading/)
- url: https://www.reddit.com/r/Kotlin/comments/ete8cq/on_overloading/
---
    operator fun &lt;T&gt; T.plus(elements: Iterable&lt;T&gt;): List&lt;T&gt; {

What is your opinion on overloading like above?
## [5][Kotlin Vs Java](https://www.reddit.com/r/Kotlin/comments/eto2xu/kotlin_vs_java/)
- url: https://www.reddit.com/r/Kotlin/comments/eto2xu/kotlin_vs_java/
---

## [6][Using of ConstraintLayout to support different Android screen sizes with Material Design](https://www.reddit.com/r/Kotlin/comments/etbetv/using_of_constraintlayout_to_support_different/)
- url: https://medium.com/@martinbaraya/using-of-to-support-different-android-screen-sizes-with-material-design-30e8669d7b39?source=friends_link&amp;sk=a23294b64f412db61b53529de11c3133
---

## [7][Kotlin background workers/job queue -- retry, exponential back-off, persistence?](https://www.reddit.com/r/Kotlin/comments/esxoag/kotlin_background_workersjob_queue_retry/)
- url: https://www.reddit.com/r/Kotlin/comments/esxoag/kotlin_background_workersjob_queue_retry/
---
I am starting a large project in Kotlin and using Ktor for my API. I am now wondering what the de-facto robust job queue system is in the Java world, something akin to Sidekiq in Ruby?

I can't find much Kotlin-specific, and am even having trouble finding Java stuff. Everywhere I look (including many posts on this subreddit and Java) mention threads or coroutines and such.. which doesn't even begin to approach the requirements of a proper job queue system. I can only assume I am missing something fundamental in my shift from separate background worker processes and what folks do on the JVM.

What do people use to track their background jobs, add retries, back-offs, persist for pause/resume between process restarts, etc, etc?

Refs:

Ruby - [https://github.com/mperham/sidekiq](https://github.com/mperham/sidekiq)

Node - [https://github.com/OptimalBits/bull](https://github.com/OptimalBits/bull)

Go - [https://gobuffalo.io/en/docs/workers](https://gobuffalo.io/en/docs/workers)

Kotlin - ?

Java - ?
## [8][Do I need to import some library to use this nextLong function?](https://www.reddit.com/r/Kotlin/comments/et5q4b/do_i_need_to_import_some_library_to_use_this/)
- url: https://i.redd.it/cpmwk83qtnc41.jpg
---

## [9][Looking for web project tutorials / tips](https://www.reddit.com/r/Kotlin/comments/esvix2/looking_for_web_project_tutorials_tips/)
- url: https://www.reddit.com/r/Kotlin/comments/esvix2/looking_for_web_project_tutorials_tips/
---
Hey there,

I'm new to Kotlin, just finished working my way through the basic concepts and syntax.

For a project I'm working on, I need to build a web server (REST in the long run, but we'll build up to that with time), and I'm supposed to keep things as independant as possible.

So basically I'm looking for a tutorial / tips / documentation to implement a web server without Maven, Gradle or any kind of external library.

&amp;#x200B;

What I already have is a small Kotlin/JS project from [this tutorial](https://www.raywenderlich.com/201669-web-app-with-kotlin-js-getting-started), that's only missing a server to run on. So any input on how to set up a server running this web application would be incredibly helpful!
## [10][Snakecase Long Type varibales value?](https://www.reddit.com/r/Kotlin/comments/esqod6/snakecase_long_type_varibales_value/)
- url: https://www.reddit.com/r/Kotlin/comments/esqod6/snakecase_long_type_varibales_value/
---
I wanted to know if there's any difference of if it's better to use one type of casing, as I've seen Long variable values written normally and snakecased, such as

val regularLong: Long = 1000000L -&gt; Regular declaration

var snakeCasedLong: Long = 1\_000\_000L -&gt; Snakecased declaration

Is there any real difference? Also, if anybody knows where to learn the best practices for Kotlin, I would appreciate if you shared it. Thanks!
