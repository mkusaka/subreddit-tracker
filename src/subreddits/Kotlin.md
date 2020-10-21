# Kotlin
## [1][How to run Kotlin Multiplatform Web Application in continous mode](https://www.reddit.com/r/Kotlin/comments/jex6ej/how_to_run_kotlin_multiplatform_web_application/)
- url: https://www.reddit.com/r/Kotlin/comments/jex6ej/how_to_run_kotlin_multiplatform_web_application/
---
I started a project using the default Kotlin Multiplatform Full-Stack Web Application. (Ktor + Kotlin React with styled components)

Can I somehow run the entire application in continous mode, so that the website gets refreshed automatically whenever I change either the front end or the back end?

Edit: I use the gradle wrapper in version 6.7.
## [2][jOOQ 3.14 includes support for kotlin code generation, among other features](https://www.reddit.com/r/Kotlin/comments/jeortf/jooq_314_includes_support_for_kotlin_code/)
- url: https://blog.jooq.org/2020/10/20/jooq-3-14-released-with-sql-xml-and-sql-json-support/
---

## [3][Kotlin Illustrated Guide - Enum Classes](https://www.reddit.com/r/Kotlin/comments/jeo9up/kotlin_illustrated_guide_enum_classes/)
- url: https://typealias.com/start/kotlin-enum-classes/
---

## [4][Suggestion: Kotlin for the .NET platform](https://www.reddit.com/r/Kotlin/comments/jea7s5/suggestion_kotlin_for_the_net_platform/)
- url: https://www.reddit.com/r/Kotlin/comments/jea7s5/suggestion_kotlin_for_the_net_platform/
---
The .NET platform has some various advantages over the JVM.  
Having a Kotlin port to the .NET platform could be very nice.  
JVM bytecode is pretty close to CLI bytecode, so a conversion would not be too complicated. Java bytecode to CLI bytecode tools already exist, [like IKVM.](https://en.wikipedia.org/wiki/IKVM.NET)  
However, since Kotlin has many features .NET has but Java doesn't (like delegates), a more efficient approach would be direct Kotlin to CLI conversion. 

Could it be considered?
## [5][Kotlin Design Tutorial For Beginners](https://www.reddit.com/r/Kotlin/comments/jee92u/kotlin_design_tutorial_for_beginners/)
- url: https://youtu.be/3rO01JsQ_DU
---

## [6][[Article] Using Mockito in Kotlin projects](https://www.reddit.com/r/Kotlin/comments/jek8zj/article_using_mockito_in_kotlin_projects/)
- url: https://kotlintesting.com/using-mockito-in-kotlin-projects/?utm_source=reddit&amp;utm_medium=kotlin
---

## [7][What is Kotlin's strategy for multiplatform UI?](https://www.reddit.com/r/Kotlin/comments/jdxvoe/what_is_kotlins_strategy_for_multiplatform_ui/)
- url: https://www.reddit.com/r/Kotlin/comments/jdxvoe/what_is_kotlins_strategy_for_multiplatform_ui/
---
I know Jetbrains folks are working with Google on porting the Jetpack compose to Desktop (https://android.googlesource.com/platform/frameworks/support/+/refs/heads/androidx-master-dev/compose/ui/ui/src/desktopMain/kotlin/androidx/compose?autodive=0%2F) and I am curious if Jetbrains/Google have any plans to port the same to Web or making it a multiplatform library?  I am asking here as I asked this question multiple times in the recent Kotlin event and got no answer :)
## [8][Express.js with Kotlin?](https://www.reddit.com/r/Kotlin/comments/je4j3w/expressjs_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/je4j3w/expressjs_with_kotlin/
---
I wrote an API with Express and I have been writing an Android app to work with the API in Kotlin.

As I have tested the API I came across weird quirks with JavaScript that caused some bugs.

I don't like using JavaScript on the server side. I think Kotlin is a well designed language and I think it would be worth my time to rewrite the API with Kotlin.

&amp;#x200B;

But I need to know will Kotlin have the same quirks as JavaScript? For example if I send a double - 1.0 it truncates the decimal and the 0, but this causes some bugs and I need to work around JavaScript by sending decimal numbers as a String. edit: JavaScript does not give a fuck if I give it a string or a Boolean or anything. So I changed the Android app to send strings and everything seems to be working.

If the Express.js API was written in Kotlin, could I send a Double for example, and have it still be a double, even if there is a 0 after the decimal? Or will Kotlin act like JavaScript does?

&amp;#x200B;

This is one of about 3 quirks JavaScript has that make me uncomfortable. I liked JavaScript at first because it was easy to get up and running but some of these shortcuts that I guess are designed to make it easier to code have caused some bugs. And JavaScript sucks at type.

&amp;#x200B;

Kotlin for Android seems to give me the best of both worlds and I was simply wondering if the language functions the same when it compiles to JavaScript.
## [9][New Dokka plugins](https://www.reddit.com/r/Kotlin/comments/je0np8/new_dokka_plugins/)
- url: https://www.reddit.com/r/Kotlin/comments/je0np8/new_dokka_plugins/
---
Dokka team has created a list of community plugins!  
You can share your plugins with the world [here](https://github.com/Kotlin/dokka/blob/master/docs/src/doc/docs/community/plugins-list.md)
## [10][Production build of Ktor](https://www.reddit.com/r/Kotlin/comments/je3p0f/production_build_of_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/je3p0f/production_build_of_ktor/
---
Hi, I'm pretty new to Kotlin / JVM. My primary platform is .NET, where I'm used to deploy to IIS. While learning SpringBoot, I deployed war on Tomcat. It was very similar to IIS. Now I'm learning Ktor, where it's recommended to use Netty. I was looking for an installer for Windows, but I didn't find any. How does deploying Ktor on Windows work when I use Netty? Or should I use something else?
