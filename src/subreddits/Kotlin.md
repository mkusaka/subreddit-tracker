# Kotlin
## [1][ktfmt, a Kotlin code formatter based on google-java-format, now with a Plugin for IntelliJ](https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/)
- url: https://www.reddit.com/r/Kotlin/comments/id7uze/ktfmt_a_kotlin_code_formatter_based_on/
---
[https://github.com/facebookincubator/ktfmt/](https://github.com/facebookincubator/ktfmt/)

This came up in a couple of archived posts here, so I'm excited to say that `ktfmt` now has an IntelliJ plugin (as well as a Spotless plugin).

Also, a lot has changed in `ktfmt` since it was announced ~8 months ago - give it a try! we're listening for feedback over at https://github.com/facebookincubator/ktfmt/issues :)
## [2][To initialize a new variable that is only used once, or just use the code instead of the variable?](https://www.reddit.com/r/Kotlin/comments/idcvaz/to_initialize_a_new_variable_that_is_only_used/)
- url: https://www.reddit.com/r/Kotlin/comments/idcvaz/to_initialize_a_new_variable_that_is_only_used/
---
I was wondering what the best practice is.

Case example:

    fun foo() {
        for (x in y) {
            val bar = [some super long series of calls involving x and y]
            if (bar fulfills some condition) 
                doThing()
        }
    }

VS

    fun foo() {
        for (x in y) {
            if ([some super long series of calls involving x and y] fulfills some condition) 
                doThing()
        }
    }

I like the first one because to me it's more readable. I can assign a meaningful name to those series of calls and later I won't have to think much to get what it means.

But... is it wrong? What if I expect to run that loop many many times. Is creating variables all those times a bad thing to do? Especially when it's not needed because I only use that variable in one place.
## [3][Need Help with Ktor](https://www.reddit.com/r/Kotlin/comments/ided9u/need_help_with_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/ided9u/need_help_with_ktor/
---
Hi guys I'm a new kotlin developer and was looking into ktor. I couldn't find any videos or tutorials(other than its doc) on how to use it. I'm used to react but I've never used anything like ktor and wanted to know how to work with HTML and css on ktor and how hard it is. Can I use HTML css directly with it??
## [4][User Activity Recognition || Activity Transition API || Kotlin](https://www.reddit.com/r/Kotlin/comments/idcxpp/user_activity_recognition_activity_transition_api/)
- url: https://youtu.be/2Em5ep_KYIY
---

## [5][Video game engine](https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/)
- url: https://www.reddit.com/r/Kotlin/comments/icupw4/video_game_engine/
---
Hello I am learning to program with kotlin, my idea is to create video games for android. I wanted to ask you if you know of any game engine for programming that supports kotlin.

Note: Sorry for my English I speak Spanish.
## [6][Critical bug found in latest Kotlin](https://www.reddit.com/r/Kotlin/comments/idd0vj/critical_bug_found_in_latest_kotlin/)
- url: https://youtrack.jetbrains.com/issue/KT-41273
---

## [7][Better way to do this?](https://www.reddit.com/r/Kotlin/comments/iclh08/better_way_to_do_this/)
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
## [8][Jetpack compose release date](https://www.reddit.com/r/Kotlin/comments/icvh54/jetpack_compose_release_date/)
- url: https://www.reddit.com/r/Kotlin/comments/icvh54/jetpack_compose_release_date/
---
Is Jetpack compose production ready !?
If not when to expect??
## [9][How can I use sealed classes with Paint into Fill, Line, and Circle if my type names are outside the data (LayerDetail.type)?](https://www.reddit.com/r/Kotlin/comments/icj0mj/how_can_i_use_sealed_classes_with_paint_into_fill/)
- url: https://i.redd.it/rrhiz7ootwh51.png
---

## [10][Problem with Google kotlin course](https://www.reddit.com/r/Kotlin/comments/icps1w/problem_with_google_kotlin_course/)
- url: https://www.reddit.com/r/Kotlin/comments/icps1w/problem_with_google_kotlin_course/
---
Hi guys!
I am new in Kotlin... I am learning from Google kotlin course. Now, I have problem. Now, I am on lesson 3 and they have start codes. When I download start code and import it as a project I got error that some files didnt support in my version of android studio... I have latest android studio, but Google start code is published 18 months ago. 
So, anyone have solution for this? :)
