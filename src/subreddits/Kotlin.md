# Kotlin
## [1][Stop scan and scanReduce before it's too late](https://www.reddit.com/r/Kotlin/comments/ftnvzd/stop_scan_and_scanreduce_before_its_too_late/)
- url: https://www.reddit.com/r/Kotlin/comments/ftnvzd/stop_scan_and_scanreduce_before_its_too_late/
---
**TLDR**

The names `scan` and `scanReduce` are poorly chosen and should be renamed as long as they are still in experimental. That is if Kotlin wants to remain a well readable language. Other languages such as Python avoided the name 'scan' as well.

&amp;#x200B;

**1.  Introduction**

With version 1.7.0 `scan` and `scanReduce` made it finally into the Kotlin programming language, but the name choice was poorly made. These functions are still in experimental state, so this is the last chance to get it right.

In Advance I want to appologize for my strong opinion on this topic, but if I didn't love Kotlin I woudn't care so much.

&amp;#x200B;

**2. Conventional names in functional programming**

Functions to process collections such as `map`, `filter` and `reduce` are well known across many programming languages.

`reduce` (also called `fold`) is a function which accumulates/reduces many values to a single result value.

`scan` is like reduce, but instead of returning a single result it yields all intermediate results.

In the example of adding a list of numbers `reduce` would simply be a sum and `scan` would be a running sum/cumulative sum.

&amp;#x200B;

**3. Names in Kotlin**

For `reduce` Kotlin uses two different names: `reduce` and `fold`. The difference is in the initial value of the operation. Kotlin's `reduce` starts operating on the first element of the collection and `fold` expects an initial value as parameter.

In similar fashion Kotlin recently introduced two versions of `scan`. Kotlin's `scan` is a running `fold` and Kotlin's `scanReduce` is a running `reduce`.

|returns \\ initial value|by paramter|first element|
|:-|:-|:-|
|end result|**fold**|**reduce**|
|all intermediate result|**scan** 😢|**scanReduce** 🤮|

&amp;#x200B;

**4. Naming Problem**

There are two naming problems. The first one being that `scanReduce` as name is totally wrong by definition. The second one being my personal opinion why the name `scan` is bad in general.

**4.1 scanReduce must not exist**

Earlier we saw that `reduce` operates on a collection and returns a single value and that `scan` does the same returning all intermediate results. If we use this definition to understand the difference between the two, what in God's name does `scanReduce` mean? This term is an **abomination** and **must not exist**. It is clear to me that `fold` relates to `scan` and there has to be a name for `reduce` to relate to. But it cannot be `scanReduce`!

**4.2 The never ending legacy of scan**

Not only the name `scanReduce` is wrong. Also the name `scan` is flawed. Many programming languages such as Haskell, Scala or Typescript have a `scan` function. `map`, `filter`, `reduce` all make sense, but the name `scan` is more than awkward. Outside of programming scan has the following meanings:

1. to look quickly but not very carefully at a document, etc.
2. to look at every part of something carefully, especially because you are looking for a particular thing or person

So where does the name in programming come from? The term scan was coined in the programming language APL by Kenneth E. Iverson at IBM in the early 1960s. I could not find out why the `\` operator was named scan operator. In this book "A Programming Language" (1962) he uses the term scan quite frequently to describe traversing a vector/matrix without changing its rank.

Since a lot of operations (`forEach`, `map`, `find`/`first`, and even `reduce`!) traverse a collection they are all technically 'scanning'. What `scan` should have been named is `accumulate`. "Accumulative" means:

1. having a result that increases in strength or importance each time more of something is added
2. including all the amounts that have been added previously

Isn't that a much better name for what the function is doing? By the way this is how the function is named in Python.

**5 Conclusion**

`accumulate` is a much better name for `scan`. Since Kotlin has two names for `reduce` (`reduce` \+ `fold`) it also needs two counterparts for `scan`. I suggest one of the follwing:

1. `accumulate` with overloads
2. `runningFold` \+ `runningReduce`
3. `accumulativeFold` \+ `accumulativeReduce`.

Python has made the right choice and it is well known for its readability and use in data science. The scan function is actually quite heavily used in data science. Kotlin is about readability too and has recently started to position itself in data science too. Will it also make the right choice?
## [2][Working with KTOR and multi-platform projects in IntelliJ](https://www.reddit.com/r/Kotlin/comments/fu6bcq/working_with_ktor_and_multiplatform_projects_in/)
- url: https://www.reddit.com/r/Kotlin/comments/fu6bcq/working_with_ktor_and_multiplatform_projects_in/
---
Hi all,

I've been struggling to create a good 'kotlin-for-all' project, with a KTOR server and and a kotlin/JS client, that runs nicely in IntelliJ

&amp;#x200B;

I've used the Kotlin Project template 'JS Client and JVM Server' which gets me a working project, but I can't seem dynamically recompile the Javascript. I'm not even sure where the JS is coming from - presumably the assembled JS jar? If I make a change to the Kotlin JS, do I have to stop the server, re-assemble the JS JAR, and restart?

&amp;#x200B;

I've managed to get dynamic JS recompilation working in a project at work but only with a small raft of hacks held together with string and sticky tape - what I want is a standardized, JetBrains-approved, non-hacky, clear way of accomplishing this goal.

&amp;#x200B;

Can anyone give me a hint on how to proceed?
## [3][Android Custom Notification in 6 Mins](https://www.reddit.com/r/Kotlin/comments/fu60b8/android_custom_notification_in_6_mins/)
- url: https://medium.com//android-custom-notification-in-6-mins-c2e7e2ddadab?source=friends_link&amp;sk=1d157d143c877d03370e311f2206e230
---

## [4][Coroutines flow with Kotlin](https://www.reddit.com/r/Kotlin/comments/fu7dl3/coroutines_flow_with_kotlin/)
- url: https://youtu.be/CIvjwIfOG5A
---

## [5][Help with when expression](https://www.reddit.com/r/Kotlin/comments/fty435/help_with_when_expression/)
- url: https://www.reddit.com/r/Kotlin/comments/fty435/help_with_when_expression/
---
Hi, recently started learning Kotlin, I encounter a question that really confuses me

 What is true about  **when**?

1. We can not use multiple statements in branches of expression style ***when***

2. Any ***if*** expression can be written with ***when***.

3. If we use ***when*** without an argument, all the branch conditions are to be boolean expressions.

4. ***else*** branch is to be use in expression ***when*** only.

5. We can not execute more than one branch in ***when***.

Can someone not just give the answer but explain why?
## [6][New to Kotlin](https://www.reddit.com/r/Kotlin/comments/ftjw23/new_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ftjw23/new_to_kotlin/
---
Hey,

I am currently studying Kotlin but having a hard time since I feel every tutorial I find somehow expects you to be a bit experienced in Java. In which I have zero experience. Are there any books which really offer the reader an extensive overview of the language? I prefer reading over videos.  
Thanks!
## [7][Math solver android app question](https://www.reddit.com/r/Kotlin/comments/ftkzsu/math_solver_android_app_question/)
- url: https://www.reddit.com/r/Kotlin/comments/ftkzsu/math_solver_android_app_question/
---
I am thinking of a making a math solver app, mostly linear algebra stuff, matrixes, vectors etc.

If i were to make something like symbolab/wolfram (much less complex, only some features), would it be resonable to do the computation on the mobile app, or should i do it with api calls to a backend?

Thank you :)
## [8][How much can a kotlin android developer make as a freelancer pe hour?](https://www.reddit.com/r/Kotlin/comments/ftt6lm/how_much_can_a_kotlin_android_developer_make_as_a/)
- url: https://www.reddit.com/r/Kotlin/comments/ftt6lm/how_much_can_a_kotlin_android_developer_make_as_a/
---
Hi.I have a friend who started freelancing and he doesn't know how much should he charge per hour.He has some experience because he already published about 7 apps on the market and got 5k or more downloads on each of them.He knows how to build good apps and I whould say he is preatty good.He is also studing CS.So guys,how much do you thing he should ask per hour?And I am curios which is the realistic maximum amount that a android dev can make per hour(one with a lot of experience).
## [9][Compared to Java I like Kotlin's typesafety and flexibility, but...](https://www.reddit.com/r/Kotlin/comments/ftrnrq/compared_to_java_i_like_kotlins_typesafety_and/)
- url: https://www.reddit.com/r/Kotlin/comments/ftrnrq/compared_to_java_i_like_kotlins_typesafety_and/
---
sometimes it bites me:

&amp;#x200B;

`fun main(args: Array&lt;String&gt;) {`

  `data class Transaction(`

  @`SerializedName("amount") val amountInCents: Int,`

  @`SerializedName("description") val description: String)`

  `val tx = Transaction(amountInCents = 42, description = "For Bob")`

  `val payloadStringOrNull: String? = Gson().toJson(tx, Transaction::class.java)`

  `payloadStringOrNull?.let { payloadStr: String -&gt;`

`val payloadIsNull: Boolean = payloadStr == null`

  `println("I feel cheated ${payloadIsNull} -- $payloadStr")`

  `}`

`}`

&amp;#x200B;

Moving the \`data class\` definition out fixes it. Hard to find this one.

But should this compile? I try ot pass something to Java, if that is not supposed to work for non-top-level-defined things then why accept it at compile time?

Also Gson using Unsafe to gimme nulls where I did not expect them really messed me up.

Is there some things to keep in mind so I will not get bitten like that again?
## [10][Suppress redundant semicolon warning](https://www.reddit.com/r/Kotlin/comments/ftr9jv/suppress_redundant_semicolon_warning/)
- url: https://www.reddit.com/r/Kotlin/comments/ftr9jv/suppress_redundant_semicolon_warning/
---
Hey guys,

Like most people, I am using kotlin in android studio, and it keeps telling me that the semicolon is redundant, by greying it out and even underlining it.

Now I don't mind that it's redundant, but I do work in languages where semicolon IS enforced (fx Flutter/Dart), so I'm not about to change my habits.

Instead, I'd love to just suppress the warning. Is there anyway to do that?
