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
## [2][Kotlin Cheat Sheet](https://www.reddit.com/r/Kotlin/comments/jmfjs6/kotlin_cheat_sheet/)
- url: https://i.redd.it/hm1fj7innqw51.jpg
---

## [3][Ktor embedded server on production?](https://www.reddit.com/r/Kotlin/comments/jmmtr9/ktor_embedded_server_on_production/)
- url: https://www.reddit.com/r/Kotlin/comments/jmmtr9/ktor_embedded_server_on_production/
---
I was wondering if you guys know some resources to learn about ktor prod deployment.

I have been deploying with the embedded Netty and running into some issues.

What would be the best way to go about this?

Thanks a lot!
## [4][Facts You May Not Know About Kotlin](https://www.reddit.com/r/Kotlin/comments/jmn85w/facts_you_may_not_know_about_kotlin/)
- url: https://youtu.be/XwzStZaDpH0?list=PLEx5khR4g7PIiAEHCt6LGMFnzq7JjO8we
---

## [5][Why no "indexOfFirstOrNull" ?](https://www.reddit.com/r/Kotlin/comments/jm8vi1/why_no_indexoffirstornull/)
- url: https://www.reddit.com/r/Kotlin/comments/jm8vi1/why_no_indexoffirstornull/
---
I am still relatively new to Kotlin, albeit a professional programmer for a very long time.

One common pattern in Kotlin seems to be that various search-type functions have a "...OrNull" variant, so that you can add an elvis operator down the line.

One glaring exception to this is "indexOfFirst" which returns the index of the first element matching some predicate, or -1 if no such element found.

Is there some reason for this? And is there some specific convention for handling the -1 case? In my own code, I find it far more clumsy to deal with than the "...OrNull" pattern, but maybe I am just too dumb to see how to "recover" from the "-1" elegantly.
## [6][Does anyone else have an issue with the Android Studio file type icons when using log tags?](https://www.reddit.com/r/Kotlin/comments/jmlucv/does_anyone_else_have_an_issue_with_the_android/)
- url: https://www.reddit.com/r/Kotlin/comments/jmlucv/does_anyone_else_have_an_issue_with_the_android/
---
I've had this issue for a long time with the JetBrains IDE where the file type icons show as generic Kotlin files. I'm working on a project where I have many files in the same package with similar names so this issue has really reared its ugly head. It would be great if I could quickly identify the files with type icons. I refuse to follow the "IConcept" and "ConceptImpl" paradigm - nor am I going to separate them into different packages.

The alternative to top-level \`private const val\` declarations is a companion object, creating an entire object which will ultimately be empty at runtime (~~As far as I can find,~~ the unused, empty object is not optimized away). Yes, I could declare them as \`val\` properties, but that's not a compile-time constant. I know, I know - it's just a single read operation, but it shouldn't need to be.

Essentially any class I want to perform logging in won't display as the appropriate type.

[https://youtrack.jetbrains.com/issue/KT-19583](https://youtrack.jetbrains.com/issue/KT-19583)

I may have worded my issue a bit strongly. Someone please tell me if there's a workaround I'm unaware of.

Edit: Using the Kotlin bytecode inspector, I've confirmed that a completely empty companion object is still generated even if unused anywhere.
## [7][Started learning Kotlin and I love it!](https://www.reddit.com/r/Kotlin/comments/jlz97r/started_learning_kotlin_and_i_love_it/)
- url: https://www.reddit.com/r/Kotlin/comments/jlz97r/started_learning_kotlin_and_i_love_it/
---
Started learning Kotlin for Android app development and I love it. I have experience in other programming languages and Kotlin has become my favourite by far. It's just that good it has increased my productivity by a lot
## [8][From Reactor to Coroutines](https://www.reddit.com/r/Kotlin/comments/jm5h8s/from_reactor_to_coroutines/)
- url: https://blog.frankel.ch/reactor-to-coroutines/
---

## [9][Python in Kotlin](https://www.reddit.com/r/Kotlin/comments/jm6mei/python_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jm6mei/python_in_kotlin/
---
I'm starting to learn Kotlin, but I have a question before continuing, it is possible to run python programs inside a Kotlin Android app?
## [10][A starter project for creating admin websites using Kotlin](https://www.reddit.com/r/Kotlin/comments/jlluda/a_starter_project_for_creating_admin_websites/)
- url: https://github.com/tipsy/kotlin-admin-template
---

## [11][Something like automapper for kotlin?](https://www.reddit.com/r/Kotlin/comments/jm0c5w/something_like_automapper_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jm0c5w/something_like_automapper_for_kotlin/
---
Hi, I have Ktor + Exposed example app. Can i do this better? Is there any mapping libraby like automapper? Thx

PersonTable.selectAll().map {
                Person(it[PersonTable.Id],
                        it[PersonTable.Name],
                        it[PersonTable.Age])
