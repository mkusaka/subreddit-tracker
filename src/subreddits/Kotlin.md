# Kotlin
## [1][KVision v3.16.0 is released (with printing support)](https://www.reddit.com/r/Kotlin/comments/j6145b/kvision_v3160_is_released_with_printing_support/)
- url: https://www.reddit.com/r/Kotlin/comments/j6145b/kvision_v3160_is_released_with_printing_support/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.16.0, with the new module for printing. Based on an excellent [Print.js](https://printjs.crabbly.com/) library, it allows to easily print the content of any KVision component, directly from the browser window. It also supports printing PDF documents, images and any JSON-formatted tabular data.

Going further with the news, Kotlin sealed classes are now fully supported as parameters and return values in KVision fullstack interfaces. New multiplatform annotation based on Jackson `@JsonTypeInfo` has been introduced. See the [documentation](https://kvision.gitbook.io/kvision-guide/part-3-server-side-interface/common-code) for details.

Last but not least, we continue to enhance state management experience by introducing simple sub-store implementation for any instance of `ObservableState` interface. It's a way to partition your data and optimize rendering at the same time.

For more details about this release see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.16.0).

As always, any feedback is appreciated.
## [2][Kotlin Public Roadmap Through Spring 2021 – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/j5oich/kotlin_public_roadmap_through_spring_2021_kotlin/)
- url: https://blog.jetbrains.com/kotlin/2020/10/kotlin-public-roadmap-through-spring-2021/
---

## [3][On Files and Okio](https://www.reddit.com/r/Kotlin/comments/j648jd/on_files_and_okio/)
- url: https://publicobject.com/2020/10/06/files/
---

## [4][Kotlin Telegram Bot looking for maintainers](https://www.reddit.com/r/Kotlin/comments/j5ptrk/kotlin_telegram_bot_looking_for_maintainers/)
- url: https://www.reddit.com/r/Kotlin/comments/j5ptrk/kotlin_telegram_bot_looking_for_maintainers/
---
Hi!

I'm Iván, the creator of [Kotlin Telegram Bot](https://github.com/kotlin-telegram-bot/kotlin-telegram-bot) a library to help you build telegram bots. 

Since a year a go or so I have not been able to dedicate the time needed to maintain the project, so I looked for more people to help me with that. At the moment 2 other persons are helping with the project, but sometimes a big [PR which needs review](https://github.com/kotlin-telegram-bot/kotlin-telegram-bot/pull/103) or a bug comes out and neither of us is able to check it as fast as we would like (like months).

So, we are looking for more people! If you like telegram and/or Kotlin and are interested in contributing to the project DM me or send me an email at [ivanmartinga@gmail.com](mailto:ivanmartinga@gmail.com) 

Thanks!
## [5][Compile to native](https://www.reddit.com/r/Kotlin/comments/j5zixi/compile_to_native/)
- url: https://www.reddit.com/r/Kotlin/comments/j5zixi/compile_to_native/
---
What is the rational behind the Kotlin/Native compiler on a desktop if I can't use the Java API? What other API can I use, e.g., to access the file system or GUI libraries?
## [6][Android Model-View-Intent with Unit Tests - Using Kotlin Flow](https://www.reddit.com/r/Kotlin/comments/j5kd1x/android_modelviewintent_with_unit_tests_using/)
- url: https://medium.com/@AdamHurwitz/android-model-view-intent-with-unit-tests-260e9a0cdd64
---

## [7][New to kotlin, gradle. Does it get easier? Otherwise, kotlin is a beauty.](https://www.reddit.com/r/Kotlin/comments/j53sv3/new_to_kotlin_gradle_does_it_get_easier_otherwise/)
- url: https://www.reddit.com/r/Kotlin/comments/j53sv3/new_to_kotlin_gradle_does_it_get_easier_otherwise/
---
hi. I am new to Kotlin, "Gradle and the Kotlin DSL" build system. I am not new to programming though. I am using Intellij. I am trying to learn kotlin and the gradle build system using kotlin dsl all at the same time. And it is super confusing. Nothing seems straight forward. I am not new to programming at all by any stretch. But I am new to kotlin and the entire java ecosystem. And I want to learn this as soon as possible because, I will start work on a very large kotlin/gradle project at work very soon.

At the moment, I am trying to do the following just to test out the kotlin language, and learn more about the gradle/kotlin DSL build system.

1. create a kotlin project with gradle/kotlinDSL as build system
2. add the fuel (http client) as a dependency.
3. use fuel to make a http request.

At the moment I am stuck at "adding fuel as a dependency". In intellij, I added fuel as a dependency in project settings but still i keep getting "unresolved reference"


Here are some of the problems I faced and does not seem straight-forward for someone that has been coding professionally for a long time, but is very new to the java ecosystem.

* couldn't find the one package manager. python has pypi. rust has crates. c# has nuget. So I naively assumed I would find something like that in the java/jvm/kotlin world. Oh boy, how wrong I was.
* theres like 10 different JDKs. I am not sure which one is the most used and which one I should use.
* confusion about what the JDK is. Is it the compiler , runtime and standard libraries. ?
* there is jscentral. maven central. and many others. Are they repositories of java libraries? kotlin libraries? or a combination of both.
* can i use any java libary with kotlin? Or are there some java libraries that won't work with kotlin? If so how do I tell if a java library will not work with kotlin
* are libraries restricted to working on only certain SDKs and not others?
* what is a module? can modules exist within other modules? In the c# world a solution can have many projects. Is it something like that?

I have so many other questions and If I seem frustrated, my apologies. I have been trying to get fuel added as a dependency and It just wont work. I've been trying to work this out for the past one day.

Otherwise, I think kotlin is a beauty. I just wish it was easier to understand the ecosytem.

Any help or document/tutorial/book that will help me to answer some or all these questions will be appreciated.
## [8][Non-intrusive, AI aided, instant root cause analysis tool](https://www.reddit.com/r/Kotlin/comments/j5dr6x/nonintrusive_ai_aided_instant_root_cause_analysis/)
- url: https://ycrash.io/
---

## [9][A minimal real Kotlin Multiplatform Showcase](https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/)
- url: https://www.reddit.com/r/Kotlin/comments/j4v42x/a_minimal_real_kotlin_multiplatform_showcase/
---
Recently I created a minimal app to illustrate the power of kotlin multiplatform. It currently runs on:

* Android (available on Google Play)
* iOS
* Web

You can checkout the [repository](https://github.com/moallemi/kotlin-multiplatform-showcase) for more info

&amp;#x200B;

https://preview.redd.it/ba04e0zr81r51.png?width=1000&amp;format=png&amp;auto=webp&amp;s=ed7a02304c2d3a362883b5845fb07ac17b0cd84b
## [10][Kotlin for competitive programming](https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/)
- url: https://www.reddit.com/r/Kotlin/comments/j4yakf/kotlin_for_competitive_programming/
---
I want to know if I can use Kotlin for Competitive programming without taking support of Java(Java collections).
Is there anyone here who uses kotlin for cp? 
If so would you like to share the details :)
