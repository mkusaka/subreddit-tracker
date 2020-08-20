# Kotlin
## [1][ktfmt, a Kotlin code formatter based on google-java-format, now with a Plugin for IntelliJ](https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/)
- url: https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/
---
[https://github.com/facebookincubator/ktfmt/](https://github.com/facebookincubator/ktfmt/)

This came up in a couple of archived posts here, so I'm excited to say that `ktfmt` now has an IntelliJ plugin (as well as a Spotless plugin).

Also, a lot has changed in `ktfmt` since it was announced ~8 months ago - give it a try! we're listening for feedback over at https://github.com/facebookincubator/ktfmt/issues :)
## [2][Video game engine](https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/)
- url: https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/
---
Hello I am learning to program with kotlin, my idea is to create video games for android. I wanted to ask you if you know of any game engine for programming that supports kotlin.

Note: Sorry for my English I speak Spanish.
## [3][Better way to do this?](https://www.reddit.com/r/Kotlin/comments/iclh08/better_way_to_do_this/)
- url: https://www.reddit.com/r/Kotlin/comments/iclh08/better_way_to_do_this/
---
So I'm trying to build a map from 3 different lists and I am using this code here:

    map = categories!!.groupBy { it.id }.mapValues {
        map-&gt; list1!!.filter {
            a -&gt;list2!!.any { b -&gt; a.id == b.id } &amp;&amp; a.category_id == map.key
    }}.toMutableMap()

So it goes like this:

* list1 contains object1 which has the category id and an id
* list2 contains object2 which has only the id and not the category id
* The keys needs to be the category ids
* The values are the elements of list1 that have the same ids of the elements of list2
* Multiple elements can be in one category
* The end result is like this: Map&lt;Long, List&lt;object1&gt;&gt;

There is quite a lot of data to go through and right now it takes somewhere between 30sec to 1min on my phone to execute (in a coroutine).
## [4][Why there's no goodl package manager for Kotlin?](https://www.reddit.com/r/Kotlin/comments/id7ixs/why_theres_no_goodl_package_manager_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/id7ixs/why_theres_no_goodl_package_manager_for_kotlin/
---
Pretty much every language has it, would be nice to have one for Kotlin too.

P.S.

I know Maven [and Gradle](https://www.reddit.com/r/java/comments/i5oab3/how_do_you_people_use_such_a_piece_of_shit_as/), both are total garbage.
## [5][Jetpack compose release date](https://www.reddit.com/r/Kotlin/comments/icvh54/jetpack_compose_release_date/)
- url: https://www.reddit.com/r/Kotlin/comments/icvh54/jetpack_compose_release_date/
---
Is Jetpack compose production ready !?
If not when to expect??
## [6][How can I use sealed classes with Paint into Fill, Line, and Circle if my type names are outside the data (LayerDetail.type)?](https://www.reddit.com/r/Kotlin/comments/icj0mj/how_can_i_use_sealed_classes_with_paint_into_fill/)
- url: https://i.redd.it/rrhiz7ootwh51.png
---

## [7][Problem with Google kotlin course](https://www.reddit.com/r/Kotlin/comments/icps1w/problem_with_google_kotlin_course/)
- url: https://www.reddit.com/r/Kotlin/comments/icps1w/problem_with_google_kotlin_course/
---
Hi guys!
I am new in Kotlin... I am learning from Google kotlin course. Now, I have problem. Now, I am on lesson 3 and they have start codes. When I download start code and import it as a project I got error that some files didnt support in my version of android studio... I have latest android studio, but Google start code is published 18 months ago. 
So, anyone have solution for this? :)
## [8][Ktor 1.4.0 now available! – Ktor Blog](https://www.reddit.com/r/Kotlin/comments/ic7v62/ktor_140_now_available_ktor_blog/)
- url: https://blog.jetbrains.com/ktor/2020/08/18/ktor-1-4-0-now-available/
---

## [9][Testing time based code with JodaTime](https://www.reddit.com/r/Kotlin/comments/icjjfz/testing_time_based_code_with_jodatime/)
- url: https://kotlintesting.com/testing-timestampts-with-joda-time/
---

## [10][Our Next Step in Multiplatform Application Development: Trikot](https://www.reddit.com/r/Kotlin/comments/ic26fi/our_next_step_in_multiplatform_application/)
- url: https://shift.mirego.com/en/trikot
---

