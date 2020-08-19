# Kotlin
## [1][Better way to do this?](https://www.reddit.com/r/Kotlin/comments/iclh08/better_way_to_do_this/)
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
## [2][Ktor 1.4.0 now available! – Ktor Blog](https://www.reddit.com/r/Kotlin/comments/ic7v62/ktor_140_now_available_ktor_blog/)
- url: https://blog.jetbrains.com/ktor/2020/08/18/ktor-1-4-0-now-available/
---

## [3][Testing time based code with JodaTime](https://www.reddit.com/r/Kotlin/comments/icjjfz/testing_time_based_code_with_jodatime/)
- url: https://kotlintesting.com/testing-timestampts-with-joda-time/
---

## [4][How can I use sealed classes with Paint into Fill, Line, and Circle if my type names are outside the data (LayerDetail.type)?](https://www.reddit.com/r/Kotlin/comments/icj0mj/how_can_i_use_sealed_classes_with_paint_into_fill/)
- url: https://i.redd.it/rrhiz7ootwh51.png
---

## [5][Our Next Step in Multiplatform Application Development: Trikot](https://www.reddit.com/r/Kotlin/comments/ic26fi/our_next_step_in_multiplatform_application/)
- url: https://shift.mirego.com/en/trikot
---

## [6][Lightweight Object Mapping Over jasync-sql](https://www.reddit.com/r/Kotlin/comments/icawo8/lightweight_object_mapping_over_jasyncsql/)
- url: https://github.com/danpozmanter/jk-mapper
---

## [7][Random feed for social media app in kotlin](https://www.reddit.com/r/Kotlin/comments/iccg6o/random_feed_for_social_media_app_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/iccg6o/random_feed_for_social_media_app_in_kotlin/
---
Working on a social media app and I want to create a feed that displays posts at random. So far I have a recent feed working and a following feed. Not sure how to get the posts at random. Ideas?
## [8][Exploring coroutines by Venkat subramaniam](https://www.reddit.com/r/Kotlin/comments/ic4no5/exploring_coroutines_by_venkat_subramaniam/)
- url: https://youtu.be/jT2gHPQ4Z1Q
---

## [9][How to return value from from a listener?](https://www.reddit.com/r/Kotlin/comments/iceztd/how_to_return_value_from_from_a_listener/)
- url: https://www.reddit.com/r/Kotlin/comments/iceztd/how_to_return_value_from_from_a_listener/
---
    fun getDeviceToken() : String {
        FirebaseInstanceId.getInstance().instanceId.addOnCompleteListener{
            return if (it.isSuccessful) it.result?.token else "" //return is not allowed here
        }
    }

I'm getting the error "return is not allowed here", so how do I return a value from a listener?
## [10][Kotlin 1.4 Released with a Focus on Quality and Performance](https://www.reddit.com/r/Kotlin/comments/ibill6/kotlin_14_released_with_a_focus_on_quality_and/)
- url: https://blog.jetbrains.com/kotlin/2020/08/kotlin-1-4-released-with-a-focus-on-quality-and-performance/
---

