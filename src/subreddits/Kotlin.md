# Kotlin
## [1][Small problem with rest Api. I need help](https://www.reddit.com/r/Kotlin/comments/gm0bs8/small_problem_with_rest_api_i_need_help/)
- url: https://www.reddit.com/r/Kotlin/comments/gm0bs8/small_problem_with_rest_api_i_need_help/
---
Hello. I am making an application that connects to Punk API and I have a small problem. I don't know how to download the ingredients list into my Recycler View. I put a link to the application where my attempts are. I was able to successfully download everything except ingredients and MashTemp. Thank you for any tips.

[https://github.com/ciechanek1/PunkApi](https://github.com/ciechanek1/PunkApi)  


[https://punkapi.com/documentation/v2](https://punkapi.com/documentation/v2)
## [2][Registration for Kotlin Heroes Coding Contest (Episode 4) is open!](https://www.reddit.com/r/Kotlin/comments/gm0ehw/registration_for_kotlin_heroes_coding_contest/)
- url: https://www.reddit.com/r/Kotlin/comments/gm0ehw/registration_for_kotlin_heroes_coding_contest/
---
Join a programming challenge developed by JetBrains and Codeforces, practice your skills, and compete for the prizes! Learn more [https://jb.gg/kv1azl](https://jb.gg/kv1azl)  and save the date: May 29, 14:35 - 17:05 UTC.
## [3][KVision 3.9.0 is released (with support for React components)](https://www.reddit.com/r/Kotlin/comments/glmzy8/kvision_390_is_released_with_support_for_react/)
- url: https://www.reddit.com/r/Kotlin/comments/glmzy8/kvision_390_is_released_with_support_for_react/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.9.0, with extremely useful support for embedding React components in KVision applications. Now you can use not only the rich set of built-in KVision components, but also literally any free React component available in the NPM repository. Full integration with KVision DSL and built-in state management makes using external React components even easier then with [kotlin-react](https://github.com/JetBrains/kotlin-wrappers)!

You can look at the [new chapter of the guide](https://kvision.gitbook.io/kvision-guide/part-2-advanced-features/using-react-components) to get started right now.

For more details about this release see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.9.0).

As always any feedback is welcomed :-)
## [4][Download Files in Kotlin for Android Using Ktor and Intents](https://www.reddit.com/r/Kotlin/comments/gm1a6j/download_files_in_kotlin_for_android_using_ktor/)
- url: https://spin.atomicobject.com/2020/05/18/android-download-files-kotlin/
---

## [5][String.split(Pattern) is not symmetrical with Pattern.split(String) for even-length strings](https://www.reddit.com/r/Kotlin/comments/gls1ko/stringsplitpattern_is_not_symmetrical_with/)
- url: https://www.reddit.com/r/Kotlin/comments/gls1ko/stringsplitpattern_is_not_symmetrical_with/
---
**Note: If you're using old Reddit, see Wolfsdale's reply below. I'm new to Reddit and using new Reddit in Markdown mode.**

I was working on a simple exercise to split a string up into pairs of characters when I noticed something surprising with the behavior of `String.split(Pattern)`: it gives a different result from `Pattern.split(String)` when the string has an even number of characters. Here's the JUnit 5 test code that shows the problem:
```
private val expr = "(?&lt;=\\G.{2})" // group into pairs of characters
private val regex: Regex = expr.toRegex()
private val pattern: Pattern = expr.toPattern()

@ParameterizedTest
@ValueSource(strings = ["a", "ab", "abc", "abcd", "abcde", "abcdef"])
fun `split should be symmetrical`(s: String) {
    assertAll(
        { assertEquals(regex.split(s).size, s.split(regex).size,
                "regex.split(\"$s\").size != \"$s\".split(regex).size") },

        { assertEquals(pattern.split(s).size, s.split(pattern).size,
                "pattern.split(\"$s\").size != \"$s\".split(pattern).size") }
    )
}
```
The `Pattern.split()` assertion fails with even-length strings. The test can be made to pass if you have this instead:
```
assertEquals(pattern.split(s, -1).size, s.split(pattern).size)
```
This doesn't seem right though since the documentation for String.split shows `limit: Int = 0` which is contradictory to the above assertion passing. The documentation also makes it seem like `split(Regex)` and `split(Pattern)` should behave the same way for all cases but that is contradicted by the following assertion failing for even-length strings:
```
assertEquals(pattern.split(s).size, regex.split(s).size)
```
This can also be made to pass by specifying a limit of -1 for `pattern.split()`, which doesn't make sense. Per the documentation of [`java.util.regex.Pattern.split()`](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/util/regex/Pattern.html#split(java.lang.CharSequence,int)) regarding the limit parameter, `regex.split(s)` behaves as though it has a negative limit instead of the declared default of `0`.

Is there any reason this shouldn't be reported as a bug in `String.split()`?
## [6][Blazing fast Fibonacci with Kotlin and Arrow library](https://www.reddit.com/r/Kotlin/comments/glm2cy/blazing_fast_fibonacci_with_kotlin_and_arrow/)
- url: https://medium.com/@cesar.tronlozai/blazing-fast-fibonacci-with-kotlin-and-arrow-library-33c1d7eca0bb?source=friends_link&amp;sk=83adb400a6d637fc80d08b52073b30d4
---

## [7][I've created a web-app for self-testing / knowledge retrieval](https://www.reddit.com/r/Kotlin/comments/glea87/ive_created_a_webapp_for_selftesting_knowledge/)
- url: https://www.reddit.com/r/Kotlin/comments/glea87/ive_created_a_webapp_for_selftesting_knowledge/
---
It's written completely in Kotlin (Kotlin/JS/React + Kotlin/JVM/Spring Boot).

This pet project was inspired by a book called 'Make it Stick'. Its authors claim that learning can be more efficient if combined with information retrieval at points of forgetting.

You can:

* Create questions with answers (markdown, code highlighting, and multiple-choice questions are supported)
* Add tags and topics to questions
* Combine questions to quizzes
* Test yourself with your questions
* Get some simple tag/topic-based analysis
* Share quiz links with the world

Short video:

https://reddit.com/link/glea87/video/qcfjtux8fbz41/player

Repo links:

* [https://github.com/night-crawler/quiz](https://github.com/night-crawler/quiz)
* [https://github.com/night-crawler/quiz-frontend](https://github.com/night-crawler/quiz-frontend)

Project: [https://quiz.force.fm/](https://quiz.force.fm/)

P.S. I'm a Python developer but I decided to give up on Python. I have finally figured out that Kotlin\*Java &gt; Python.
## [8][Kotlin Collection’s appended To, By, and With](https://www.reddit.com/r/Kotlin/comments/glc1a7/kotlin_collections_appended_to_by_and_with/)
- url: https://medium.com/@elye.project/kotlin-collections-appended-to-by-and-with-205d9540208?source=friends_link&amp;sk=e08ab97f670d9f6b488a8d67cd67a58c
---

## [9][new android developer from ios](https://www.reddit.com/r/Kotlin/comments/glh4w9/new_android_developer_from_ios/)
- url: https://www.reddit.com/r/Kotlin/comments/glh4w9/new_android_developer_from_ios/
---
hey everyone, i'm a swift developer with ~ 3.5 years professional experience, starting to get into android dev, and am hoping for some advice primarily when it comes to ui.

after looking around just a bit, it seems like two popular options that have the features i want (primarily programmatic, declarative ui) are jetpack compose and anko. compose looks great because it's backed by google, but it seems that, like swiftui, it's still in beta and won't be production ready for a while. anko looks great as well (read about it in Kotlin in Action) but was surprised to see that it's been deprecated. i wanted to start android dev for fun and profit, as i'll likely have the chance to work on android projects at work, so i'm a little hesitant to learn a technology that's not viable for a production app.

so, my question is, what do you think i should do? learn compose, because it seems like that will become the standard? learn anko, or another third party library (even though anko is deprecated, i'm wondering if it's still viable)? or is my best bet to just learn creating ui through xml, which correct me if i'm wrong, seems like the current standard way of doing things.

thanks in advance for any help, and if you have any advice related to coming from ios to android, that would be greatly appreciated as well!
## [10][Interesting problems that I met using Kotlin/IntelliJ, looking for explanations.](https://www.reddit.com/r/Kotlin/comments/gljj33/interesting_problems_that_i_met_using/)
- url: https://www.reddit.com/r/Kotlin/comments/gljj33/interesting_problems_that_i_met_using/
---
I have a bunch of kt files in one module, each of them has a main function

    fun main(){}

Recently I opened my IDE(IntelliJ CE), and all of those kt files gave me an error saying

    Conflicting overloads:public fun main(): Unit defined in xxx package

I found two ways to fix it.

First I could change

    fun main()
into

    fun main(args: Array&lt;String&gt;)

Or I could set the kotlin version back to 1.1 from 1.3, which I already have set for the project, and then change it back to 1.3 in IntellJ settings(not module settings). And then they would compile and run.


Any idea why this is happening. By changing kotlin version to 1.1 then back to 1.3, I effectively changed nothing at all, but it fixed my problem.
