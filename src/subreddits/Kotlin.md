# Kotlin
## [1][How to Make the Compiler Smarter with Kotlin Contracts](https://www.reddit.com/r/Kotlin/comments/jpcrpa/how_to_make_the_compiler_smarter_with_kotlin/)
- url: https://deniskrr.medium.com/how-to-make-the-compiler-smarter-b37f414875ac?sk=ab72c70068a9199f205a95567cf98207
---

## [2][JetBrains Knowledge Day](https://www.reddit.com/r/Kotlin/comments/jpd0k5/jetbrains_knowledge_day/)
- url: https://www.youtube.com/watch?v=kV0vZQe5zk0
---

## [3][Spring Data Search 1.0.1 released](https://www.reddit.com/r/Kotlin/comments/jp92s8/spring_data_search_101_released/)
- url: https://www.reddit.com/r/Kotlin/comments/jp92s8/spring_data_search_101_released/
---
I'm happy to share with you the release [1.0.1](https://github.com/Kobee1203/spring-data-search/releases/tag/1.0.1) of [Spring Data Search](https://github.com/Kobee1203/spring-data-search).

This library allows to automatically expose endpoints to search for data using a powerful query language.

What's New :

* Changing the base path
* Handle Map type fields on Entities
* Support for Reactive application
* 2 new sample applications added in Java and in Spring WebFlux, available in the distribution zip

[Example of complex query](https://i.redd.it/seenvlreknx51.gif)
## [4][LeetCode 11 - Container With Most Water (Medium)](https://www.reddit.com/r/Kotlin/comments/jpltr2/leetcode_11_container_with_most_water_medium/)
- url: https://www.reddit.com/r/Kotlin/comments/jpltr2/leetcode_11_container_with_most_water_medium/
---
 A new blog post discussing the approach to solve this problem - [https://redquark.org/leetcode/0011-container-with-most-water/](https://redquark.org/leetcode/0011-container-with-most-water/)
## [5][Jetpack Compose for Desktop: Milestone 1 Released](https://www.reddit.com/r/Kotlin/comments/johcrc/jetpack_compose_for_desktop_milestone_1_released/)
- url: https://blog.jetbrains.com/cross-post/jetpack-compose-for-desktop-milestone-1-released/
---

## [6][Scanner question](https://www.reddit.com/r/Kotlin/comments/jp6g76/scanner_question/)
- url: https://www.reddit.com/r/Kotlin/comments/jp6g76/scanner_question/
---
So they want us to do this:

input:

    Hello world
    this is monika

expected output:

    Hello
    world
    this
    is
    monika

2nd input:

    one two three four five

expected 2nd output:

    one
    two
    three
    four
    five

My code:

        val scanner = Scanner(System.`in`)
        val s1 = scanner.next()
        val s2 = scanner.next()
        val s3 = scanner.next()
        val s4 = scanner.next()
        val s5 = scanner.next()
        println(s1)
        println(s2)
        println(s3)
        println(s4)
        println(s5)
    
    
    
    
        val s6 = scanner.next()
        val s7 = scanner.next()
        val s8 = scanner.next()
        val s9 = scanner.next()
        val s10 = scanner.next()
        println(s6)
        println(s7)
        println(s8)
        println(s9)
        println(s10)
    }
    

Its wrong, i have no idea why, i have been trying to solve it for an hour
## [7][Are Kotlin moving away from the jvm?](https://www.reddit.com/r/Kotlin/comments/jos8bb/are_kotlin_moving_away_from_the_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/jos8bb/are_kotlin_moving_away_from_the_jvm/
---
Since the new Desktop compose framework is jetbrains moving intellij to that platform in the long run. Cutting of ties to the jvm?
## [8][Any possibility of releasing Jetpack Compose for Kotlin multiplatform mobile support (Android and iOS) ?](https://www.reddit.com/r/Kotlin/comments/jox14m/any_possibility_of_releasing_jetpack_compose_for/)
- url: https://www.reddit.com/r/Kotlin/comments/jox14m/any_possibility_of_releasing_jetpack_compose_for/
---
Jetpack Compose released for Desktop also 

Now I am thinking about what if it also released with Kotlin multiplatform mobile support?

Any possibility for that ?
## [9][JB Compose Web poc - https://zal.im/h8/](https://www.reddit.com/r/Kotlin/comments/jop7ba/jb_compose_web_poc_httpszalimh8/)
- url: https://www.reddit.com/r/Kotlin/comments/jop7ba/jb_compose_web_poc_httpszalimh8/
---
https://zal.im/h8/ - Glad they are not using Canvas for rendering the app like in the flutter web. Also, JB Compose team is looking for a Developer to work on Web version of the framework - https://www.jetbrains.com/careers/jobs/ui-framework-developer-400/
## [10][Help - Find View By ID doesn't seem to work](https://www.reddit.com/r/Kotlin/comments/joxc58/help_find_view_by_id_doesnt_seem_to_work/)
- url: https://www.reddit.com/r/Kotlin/comments/joxc58/help_find_view_by_id_doesnt_seem_to_work/
---
First of all, I'm a newbie and I'm not even a programmer. I'm following a Udemy course in order to learn how to get apps done.

So,  simply put I've created a button in a LinearLayout view and set an ID for it but I can't call it in my main activity - it just doesn't show up as I type its id inside the main class. I know Kotlin was supposed to call a view by its id, but why isnt't it working?

&amp;#x200B;

\[edit\] this is my main activity kt

    package com.example.idadeemminutos
    
    import androidx.appcompat.app.AppCompatActivity
    import android.os.Bundle
    
    class MainActivity : AppCompatActivity() {
        override fun onCreate(savedInstanceState: Bundle?) {
            setContentView(R.layout.activity_main)
            super.onCreate(savedInstanceState)
    
            // can't find btnCatchDate in here!
        }
    }

and this is the button in my main activity xml

        &lt;Button
            android:id="@+id/btnCatchDate"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:text="sua data de nascimento"
            android:layout_marginTop="15dp"
            android:textStyle="bold"
            android:textSize="20sp"
            /&gt;
    
