# Kotlin
## [1][GraalVM 20.1 comes with improved Kotlin coroutine support](https://www.reddit.com/r/Kotlin/comments/gn63sm/graalvm_201_comes_with_improved_kotlin_coroutine/)
- url: https://medium.com/graalvm/graalvm-20-1-7ce7e89f066b
---

## [2][The Debate: Which is Better For Android App Development, Java Vs. Kotlin?](https://www.reddit.com/r/Kotlin/comments/gnb3mg/the_debate_which_is_better_for_android_app/)
- url: https://kodytechnolab.com/java-vs-kotlin-comparison
---

## [3][Idiomatic way to handle traditional usage of null in function parameters](https://www.reddit.com/r/Kotlin/comments/gn7pfp/idiomatic_way_to_handle_traditional_usage_of_null/)
- url: https://www.reddit.com/r/Kotlin/comments/gn7pfp/idiomatic_way_to_handle_traditional_usage_of_null/
---
It's fairly normal in Java and many other languages to use `null` as a parameter default in functions in places where null would indicate "perform no functionality based on that parameter"

Here's an example of the pattern in question:

    class foo {
        void bar(int number, Baz exampleBaz) {
            //do some kind of functionality
            //if exampleBaz is not null, do something related to exampleBaz as well
        }
        
        void bar(int number) {
            bar(number, null);
        }
    }

Is there a more idiomatic way of achieving this then just checking the type of the parameter and relying on smart casts? A lot of weird design patterns I did to work around null safety ended up being antipatterns so I am hesitant to resort to nullable types immediately.
## [4][What's the best way to build a cross-platform app today, for Android and iOS?](https://www.reddit.com/r/Kotlin/comments/gmomkw/whats_the_best_way_to_build_a_crossplatform_app/)
- url: https://www.reddit.com/r/Kotlin/comments/gmomkw/whats_the_best_way_to_build_a_crossplatform_app/
---
For business logic, I am definitely looking at Kotlin Multiplatform. But I'm not sure of the UI part.   
React Native vs Flutter - I'm not sure which of them works best with Multiplatform. And I'm not conversant with either of them.

It is very important that I have to write as less code as possible to release the apps on both the platforms and maintain them for a long time.

Can somebody help in making me decide?
## [5][Conflict between Kotlin and project dependencies](https://www.reddit.com/r/Kotlin/comments/gmu950/conflict_between_kotlin_and_project_dependencies/)
- url: https://www.reddit.com/r/Kotlin/comments/gmu950/conflict_between_kotlin_and_project_dependencies/
---
I've been using Kotlin for quite a while now and I've never encountered this issue before, but now I really don't know how to go about it.

I'm using gradle to build my Kotlin application. I recently added a feature that uses the google custom search API and I therefore include a dependency on the official google library for this service.

Now, this broke the application and I get a `NoSuchMethodError` when initialising the new library. It turns out that the library uses a method of the Google guava `Preconditions` class that doesn't exist at runtime, for some reason. So I ran `gradle dependencies` and looked for incompatible guava versions, but only found versions that have the method in question. I then decided to manually declare the guava dependency, hoping that this would override the clash.

When this didn't work either, I figured out where the guava at runtime is coming from. It turns out that it is kotlin, more specifically the `kotlin-compiler-1.3.70` jar. Now, the problem is that guava is *inside* that jar, so I don't know how to exclude it with gradle.

Has anyone ever experienced something similar and has an idea how to resolve this conflict? It is surprising to me that kotlin even uses such an old and incompatible version. I haven't checked what exactly it is, but it must be at least guava 25.0 or lower.
## [6][JVM Performance Engineering and Troubleshooting](https://www.reddit.com/r/Kotlin/comments/gn4r90/jvm_performance_engineering_and_troubleshooting/)
- url: https://blog.gceasy.io/2015/07/17/jvm-performance-engineering-troubleshooting/
---

## [7][Discord](https://www.reddit.com/r/Kotlin/comments/gmw2xh/discord/)
- url: https://www.reddit.com/r/Kotlin/comments/gmw2xh/discord/
---
Does the Discord channel not work?
## [8][How to avoid "import hell" when your DSL has a large number of extension functions](https://www.reddit.com/r/Kotlin/comments/gmr50a/how_to_avoid_import_hell_when_your_dsl_has_a/)
- url: https://www.reddit.com/r/Kotlin/comments/gmr50a/how_to_avoid_import_hell_when_your_dsl_has_a/
---
I've been working on a Kotlin web framework called [Kweb](https://kweb.io/), which is designed to allow backend developers to build modern websites with minimal fuss.

Kweb incorporates a DSL which corresponds roughly to various aspects of HTML and JavaScript - and to make this work it relies on a lot of extension functions, you can find an example [here](https://github.com/kwebio/kweb-core/blob/master/src/main/kotlin/kweb/prelude.kt).

Previously these extension functions were categorized into different packages, but I found that this often required a large number of `import` statements at the top of every file.  Even when these are created automatically by the IDE, I still found it rather unweildy.

So I made the decision to move almost all extension functions to the `kweb` package so that they could be imported with a simple `import kweb.*`.

Is this a common pattern, or is there a better way to avoid needing a large number of import statements for extension functions?
## [9][[Spanish] Online{ it.kt } - Kotlin Remote Latam &lt;3](https://www.reddit.com/r/Kotlin/comments/gmcaka/spanish_online_itkt_kotlin_remote_latam_3/)
- url: https://www.reddit.com/r/Kotlin/comments/gmcaka/spanish_online_itkt_kotlin_remote_latam_3/
---
[Kotlin Remote Latam \&lt;3](https://preview.redd.it/kxaxz3acslz41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=89dc1c24bfb3512673d31f96a14f73abfdf36ff3)

We are going to have our first Online{ it.kt } talks, an event with 8 Latam speakers in two days of content.

We are gonna talk about DI, Android, Firebase, Coroutines, Channels and more.

Join us 

[https://www.youtube.com/channel/UC0KbJMASJI0QUohkYYiDtjg](https://www.youtube.com/channel/UC0KbJMASJI0QUohkYYiDtjg)

[https://www.meetup.com/Kotlin-CDMX/events/270719229/](https://www.meetup.com/Kotlin-CDMX/events/270719229/)

[https://www.meetup.com/Kotlin-CDMX/events/270720465/](https://www.meetup.com/Kotlin-CDMX/events/270720465/)
## [10][Download Files in Kotlin for Android Using Ktor and Intents](https://www.reddit.com/r/Kotlin/comments/gm1a6j/download_files_in_kotlin_for_android_using_ktor/)
- url: https://spin.atomicobject.com/2020/05/18/android-download-files-kotlin/
---

