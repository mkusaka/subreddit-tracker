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

On October 12–15 we hosted the Kotlin 1.4 Online Event, and over the course of those four days we received many interesting questions. Since we did not manage to cover all of them due to time constraints, we will answer them in a specific thread under this post.

Please post your questions as top-level comments to this post.

We look forward to hearing from you!
## [2][Lint warning "assignment should be lifted out of if"](https://www.reddit.com/r/Kotlin/comments/jlhm1g/lint_warning_assignment_should_be_lifted_out_of_if/)
- url: https://www.reddit.com/r/Kotlin/comments/jlhm1g/lint_warning_assignment_should_be_lifted_out_of_if/
---
I am sorry if this question is kinda lame.. but i am killing myself last hour or so to figure out what's the issue... In Android Studio Kotlin I want to clean my code based on lint warnings. Cleaned everything except this:

&amp;#x200B;

var myCoolImage = BitmapFactory.decodeFile(myCoolFile.toString())

val hhh = myCoolImage.height

val www = myCoolImage.width

if (hhh &gt; www) myCoolImage = resizeBitmap(myCoolImage,600,400)

else myCoolImage = resizeBitmap(myCoolImage,400,600)

&amp;#x200B;

Simple code to check if photo is horizontal or vertical and to resize it accordingly to that.

This code works but gives warning "assignment should be lifted out of if"

Sooo... Any idea what to do? :)
## [3][AndroidBites | Java ☕️ Maps 🗺 on the Kotlin.](https://www.reddit.com/r/Kotlin/comments/jle9xu/androidbites_java_maps_on_the_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jle9xu/androidbites_java_maps_on_the_kotlin/
---
Most of the Kotlin developers are migrated from the Java environment, coming to Kotlin the collection framework has some tweaks which makes It so awesome to work with, In Today’s article, let's understand how Kotlin catered to Java Maps in his its colors, and few other concepts that will help you understand stdlib of collection

https://chetan-garg36.medium.com/java-%EF%B8%8F-maps-on-the-kotlin-8930b9f55d8d
## [4][kotlinx.coroutines 1.4.0: Introducing StateFlow and SharedFlow](https://www.reddit.com/r/Kotlin/comments/jkuyy2/kotlinxcoroutines_140_introducing_stateflow_and/)
- url: https://blog.jetbrains.com/kotlin/2020/10/kotlinx-coroutines-1-4-0-introducing-stateflow-and-sharedflow/
---

## [5][We are participating in the development of the international Kotlin Multiplatform community](https://www.reddit.com/r/Kotlin/comments/jkzqc5/we_are_participating_in_the_development_of_the/)
- url: https://www.reddit.com/r/Kotlin/comments/jkzqc5/we_are_participating_in_the_development_of_the/
---

Our specialists at IceRock have contributed to the documentation for the Kotlin Multiplatform Mobile portal (https://kotlinlang.org/lp/mobile/). This has been prepared by the JetBrains team, the creators of the Kotlin programming language. The portal helps programmers who study this language independently and use the multi-platform technology.  

Android developers are familiar with Kotlin and understand its benefits – it's capacious, secure, and user-friendly. And iOS app developers can easily learn the syntax because it is similar to Swift.

Kotlin Multiplatform Mobile (multi-platform) is a set of tools for developing cross-platform apps. It enables business logic to be written just once for two operating systems, Android and iOS. The multi-platform helps reduce the time spent on writing, testing and debugging apps.

We have been using this technology at IceRock for more than 2 years and have released 14 projects based on it. We will talk about this soon in our business cases.
Thanks to their experience in Kotlin development, our specialists have participated in writing documentation for the portal. It contains the following information:

●      multi-platform projects and how to get started with them;
●      the mobile Kotlin multi-platform allowing such projects to be created;
●      what libraries to use and where to find them;
●      the business cases that allow the experience of other development teams to be studied. 

It also provides technical documentation and guides for common problem-solving. 
Several portal sections have been written by our specialists: Vladislav Areshkin @tetraquark_v and Andrey Chernov.
 
Links are provided below. The content is in English.

Organizing the code writing process using the Kotlin mobile multi-platform: https://kotlinlang.org/docs/mobile/organize-process-around-kmm.html

Using the SQLDelight database for the multi-platform: https://kotlinlang.org/docs/mobile/configure-sqldelight-for-data-storage.html

Ktor framework for building an asynchronous client/server architecture that provides high program performance: https://kotlinlang.org/docs/mobile/use-ktor-for-networking.html

The Kotlin Multiplatform language libraries: https://libs.kmp.icerock.dev
## [6][Struggling with Leet and Kotlin](https://www.reddit.com/r/Kotlin/comments/jl0b8q/struggling_with_leet_and_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jl0b8q/struggling_with_leet_and_kotlin/
---
My first post here, so please forgive me if this is inappropriate.

I have been a professional programmer for a very long time (more than 35 years), and have moved over time from (including others) assembler, to C, to C++, to Java, to Scala, to Kotlin.

To pass the time during "covid lockdown" I decided a couple of weeks ago to flex my Kotlin skills and tackle lots of "programming challenges" on sites like HackerRank and Leet.

Now that I have tackled about 100 such problems, I have noticed that my clean and functional Kotlin code almost always rates "slower than 95% of submissions" or some such, and the only way to get it ranking higher is to go really low-level in the code, doing pointer-based stuff in the style of C (or, specifically, index into arrays using Java-style code).

I do understand that the purpose of functional programming is to write elegant and expressive code, but time and again I end up rewriting my submissions to improve my ranking on these sites.

Of course, in "real life" performance isn't always critical, but at the same time it has me pondering hard on whether my love for "functional" code means I am doomed to writing slow code.
## [7][Junior iOS developer looking to expand my skill set !](https://www.reddit.com/r/Kotlin/comments/jkuhoy/junior_ios_developer_looking_to_expand_my_skill/)
- url: https://www.reddit.com/r/Kotlin/comments/jkuhoy/junior_ios_developer_looking_to_expand_my_skill/
---
Hey yall! 

So, i'm a self taught junior iOS developer. I've been learning in the academy of YouTube for for 6 months before i got my first job 3 months ago ( a grand total of 9 months, MASSIVE!)

So, i know there's still a lot for me to learn about iOS, but - at the company i work for right now there aren't any android developers, so the head of development is forced to work as an Android dev. 

That's where i try to come in ! I want to learn Kotlin in order for me to be able to work on our app from both platforms, and have a wider set of skill. 

From what i've seen so far the differences between Swift and Kotlin aren't massive, so i was wondering, is there anyone else here who is an iOS developer as well? If so, how did you go about adapting to Kotlin? What are some good sources to learn for someone with a bit of history? for iOS there's hacking with swift and Angela Yu, are there any equivalencies in Kotlin? 

&amp;#x200B;

Thanks!
## [8][Best Way To Learn Kotlin](https://www.reddit.com/r/Kotlin/comments/jks1yh/best_way_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jks1yh/best_way_to_learn_kotlin/
---
I am a CS graduate and I have experience with C++, C#.NET and Python. But I have never worked with Java.

I want to learn Kotlin from the ground-up, and not just learn it quickly and without depth.

I REALLY want to know how everything works and how each piece fits in the language. My goal is to become an Android Developer, but after learning Kotlin first.

Are there any resources you would recommend to me? I really prefer resources with plenty of exercies, videos are okay only if they contain exercises as well, and no, *think of something, and build it* is not much of a help, because I have no practical problems that I can currently solve with plain Kotlin on the commandline.
## [9][Koin automatic injection](https://www.reddit.com/r/Kotlin/comments/jkuhap/koin_automatic_injection/)
- url: https://www.reddit.com/r/Kotlin/comments/jkuhap/koin_automatic_injection/
---
Libraries like Guice or Dagger allow you to inject a dependency, if all of this dependency dependencies are satisfied (through your module definition).

This comes in handy to avoid having huge module definitions and also you can inject classes without dependencies without adding them to the module definition.

Any idea on how to do this on Koin or Kodein?
## [10][Migrating the deprecated Kotlin Android Extensions compiler plugin to ViewBinding](https://www.reddit.com/r/Kotlin/comments/jkt7up/migrating_the_deprecated_kotlin_android/)
- url: https://melegy.medium.com/migrating-the-deprecated-kotlin-android-extensions-compiler-plugin-to-viewbinding-d234c691dec7
---

## [11][Suggestion: an option to add automatic reversals for boolean functions](https://www.reddit.com/r/Kotlin/comments/jkuowv/suggestion_an_option_to_add_automatic_reversals/)
- url: https://www.reddit.com/r/Kotlin/comments/jkuowv/suggestion_an_option_to_add_automatic_reversals/
---
So a way to set an option that will create opposite functions for boolean functions, just like it creates getters and setters.  
For example:

`class Foo {`  
`...`  
`fun isEmpty() = this.amountOfBar == 0`  
`fun isBig() = this.amountOfBar &gt; 0`  
`}`  
It'll then automatically generate these functions:  
`@Kotlin.internal.InlineOnly inline fun Foo.isNotEmpty() = !this.isEmpty()`  
`@Kotlin.internal.InlineOnly inline fun Foo.isNotBig() = !this.isBig()`  
Which will inline on compilation. This will make coding more fluent, so no need to go to the back of the function call to add a `!`.
