# Kotlin
## [1][Where can I find the 1.3.72 changelog?](https://www.reddit.com/r/Kotlin/comments/g1mo2w/where_can_i_find_the_1372_changelog/)
- url: https://www.reddit.com/r/Kotlin/comments/g1mo2w/where_can_i_find_the_1372_changelog/
---

## [2][Fixing timing issues in tests](https://www.reddit.com/r/Kotlin/comments/g1qwbz/fixing_timing_issues_in_tests/)
- url: https://blog.djy.io/fixing-timing-issues-in-tests/
---

## [3][[HELP] Using Spring and Kotlin to make API - Getting a java.lang.VerifyError: Bad local variable type](https://www.reddit.com/r/Kotlin/comments/g1flsu/help_using_spring_and_kotlin_to_make_api_getting/)
- url: https://www.reddit.com/r/Kotlin/comments/g1flsu/help_using_spring_and_kotlin_to_make_api_getting/
---
Hello, 

This might not be the best place to be asking it, if so could someone point me in the right direction.

I have built out a JPA Repository and Spring Service and Entity using Kotlin ( I have done this 15 times over and none before this had issues.)

Whenever I run the save method on the JpaRepository, I end up with one of two errors. One being 

    Exception Details: 
    Location: com/example/api/persistant/model/timetracking/TimeEntry_Accessor_8i288y.setProperty(Lorg/springframework/data/mapping/PersistentProperty;Ljava/lang/Object;)V @145: aload_3    
    Reason: Type top (current frame, locals[3]) is not assignable to reference type    
    Current Frame:      
        bci: @145      
        flags: { } 
        locals: { 'com/example/api/persistant/model/timetracking/TimeEntry_Accessor_8i288y', 'org/springframework/data/mapping/PersistentProperty', 'java/lang/Object' }      
        stack: { 'com/example/api/persistant/model/timetracking/TimeEntry_Accessor_8i288y' }

The other which I understand is a linking error because it thinks its a duplicate class but I can't find where it thinks its a duplicate class.

     java.lang.LinkageError: loader (instance of  sun/misc/Launcher$AppClassLoader): attempted  duplicate class definition for name: "com/example/api/persistant/model/timetracking/TimeEntry_Accessor_oqv37" 

I can edit this with any more info if need be. Thank you in advance for the help. :) 

&amp;#x200B;

NOTE: I did also post this on [StackOverflow](https://stackoverflow.com/questions/61167193/bad-local-variable-type-when-calling-save-on-jpa-repository-in-kotlin) but have not gotten a response. So looking elsewhere for help.
## [4][Restul api using kotlin (Spring boot)](https://www.reddit.com/r/Kotlin/comments/g1prc9/restul_api_using_kotlin_spring_boot/)
- url: https://www.reddit.com/r/Kotlin/comments/g1prc9/restul_api_using_kotlin_spring_boot/
---
I wrote this post about how to use your kotlin knowledge from android to write down a restful api. This  is the first part and I am working on next parts now

[https://medium.com/@zahidrasheed/getting-started-with-restful-api-1d8ee5cac3cc](https://medium.com/@zahidrasheed/getting-started-with-restful-api-1d8ee5cac3cc)
## [5][Coroutines, exception handling and withTimeout. Can't wrap my head around this combination.](https://www.reddit.com/r/Kotlin/comments/g1kvnx/coroutines_exception_handling_and_withtimeout/)
- url: https://www.reddit.com/r/Kotlin/comments/g1kvnx/coroutines_exception_handling_and_withtimeout/
---
    import kotlinx.coroutines.*
    import kotlinx.coroutines.flow.*
    import kotlin.system.measureTimeMillis
    
    fun time() = System.currentTimeMillis()
    
    fun main(args: Array&lt;String&gt;) {
        runBlocking {
            repeat(5) {
                try {
                    println("${time()} withTimeout before")
                    withTimeout(1L) {
                        try {
                            delay(200L)
                        } catch (e: Exception) {
                            println("${time()} withTimeout exception " + e.localizedMessage)
                        }
                    }
                    println("${time()} withTimeout done")
                } catch (e: Exception) {
                    println("${time()} exception " + e.localizedMessage)
                }
                println("${time()} -----")
            }
        }
    }

This gives me the output:

    1586923684774 withTimeout before
    1586923684788 withTimeout exception Timed out waiting for 1 ms
    1586923684789 withTimeout done
    1586923684789 -----
    1586923684789 withTimeout before
    1586923684790 withTimeout exception Timed out waiting for 1 ms
    1586923684791 exception Timed out waiting for 1 ms
    1586923684791 -----
    1586923684791 withTimeout before
    1586923684792 withTimeout exception Timed out waiting for 1 ms
    1586923684792 exception Timed out waiting for 1 ms
    1586923684792 -----
    1586923684792 withTimeout before
    1586923684794 withTimeout exception Timed out waiting for 1 ms
    1586923684794 exception Timed out waiting for 1 ms
    1586923684794 -----
    1586923684794 withTimeout before
    1586923684795 withTimeout exception Timed out waiting for 1 ms
    1586923684795 exception Timed out waiting for 1 ms
    1586923684795 -----

On the first loop, `withTimeout done` gets printed. 

**Question 1: Why only on the first iteration?**

**Question 2: If I change the wait time in** `withTimeout` **from 1L to 10L, it does not get printed. Instead, the exception is propagated to the outer try/catch, like on the next iterations (See below). Why?**

    1586923727317 withTimeout before
    1586923727333 withTimeout exception Timed out waiting for 10 ms
    1586923727334 exception Timed out waiting for 10 ms
    1586923727334 -----
    1586923727334 withTimeout before
    1586923727347 withTimeout exception Timed out waiting for 10 ms
    1586923727347 exception Timed out waiting for 10 ms
    1586923727347 -----
    1586923727347 withTimeout before
    1586923727360 withTimeout exception Timed out waiting for 10 ms
    1586923727360 exception Timed out waiting for 10 ms
    1586923727360 -----
    1586923727360 withTimeout before
    1586923727370 withTimeout exception Timed out waiting for 10 ms
    1586923727371 exception Timed out waiting for 10 ms
    1586923727371 -----
    1586923727371 withTimeout before
    1586923727382 withTimeout exception Timed out waiting for 10 ms
    1586923727382 exception Timed out waiting for 10 ms
    1586923727382 -----

**Question 3: Why, even though the exception is being caught inside a try/catch block in** `withTimeout()`**, it is being propagated to the outer try/catch block?**

I appreciate and thank you in advance for your help.
## [6][[side-project] Daft Punk - Da Funk Quarantine Remix in Kotlin](https://www.reddit.com/r/Kotlin/comments/g12l51/sideproject_daft_punk_da_funk_quarantine_remix_in/)
- url: https://www.youtube.com/watch?v=OdQQJPpL6Lo&amp;t=138s
---

## [7][Going Full Kotlin Multiplatform - Talking Kotlin](https://www.reddit.com/r/Kotlin/comments/g14tij/going_full_kotlin_multiplatform_talking_kotlin/)
- url: https://talkingkotlin.com/going-full-kotlin-multiplatform/
---

## [8][Animations in Jetpack Compose using Transition](https://www.reddit.com/r/Kotlin/comments/g1bj13/animations_in_jetpack_compose_using_transition/)
- url: https://proandroiddev.com/animations-in-jetpack-compose-using-transition-25d5d2143401
---

## [9][[Help] Layout in service](https://www.reddit.com/r/Kotlin/comments/g184ub/help_layout_in_service/)
- url: https://www.reddit.com/r/Kotlin/comments/g184ub/help_layout_in_service/
---
Hello, in this project the layout is opened when it draws over all apps... but i can't make them be where i want them to be.  
Example: up img is what i want, bottom is what i get almost never mind what i do  


https://preview.redd.it/ncj51a6n2ts41.jpg?width=524&amp;format=pjpg&amp;auto=webp&amp;s=0dd146b3cfe51e74df8eb8f67434bd2beb2adfb8

Thank you guys in advance :)

XML file:

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        xmlns:tools="http://schemas.android.com/tools"
        android:id="@+id/cont"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:theme="@style/AppTheme"
        tools:context=".MainActivity"&gt;
    
        &lt;com.google.android.material.floatingactionbutton.FloatingActionButton
            android:id="@+id/btnService"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:backgroundTint="@android:color/holo_blue_dark"
            android:clickable="true"
            android:focusable="true"
            app:layout_constraintBottom_toBottomOf="parent"
            app:srcCompat="@android:drawable/btn_star_big_on"
            tools:ignore="MissingConstraints"
            tools:layout_editor_absoluteX="489dp" /&gt;
    
        &lt;ScrollView
            android:id="@+id/scrlView"
            android:layout_width="80dp"
            android:layout_height="match_parent"
            android:background="#2196F3"
            android:visibility="invisible"
            app:layout_constraintEnd_toEndOf="parent"
            tools:ignore="MissingConstraints"&gt;
    
            &lt;LinearLayout
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:orientation="vertical"&gt;
    
                &lt;Button
                    android:id="@+id/button1"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="1 M" /&gt;
    
                &lt;Button
                    android:id="@+id/button2"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="5 M" /&gt;
    
                &lt;Button
                    android:id="@+id/button3"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="10 M" /&gt;
    
                &lt;Button
                    android:id="@+id/button4"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="15 M" /&gt;
    
                &lt;Button
                    android:id="@+id/button5"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="30 M" /&gt;
    
                &lt;Button
                    android:id="@+id/button6"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="60 HM" /&gt;
    
                &lt;Button
                    android:id="@+id/button7"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="3 H" /&gt;
    
                &lt;Button
                    android:id="@+id/button8"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="8 H" /&gt;
    
                &lt;Button
                    android:id="@+id/button9"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="15 H" /&gt;
    
                &lt;Button
                    android:id="@+id/button10"
                    android:layout_width="match_parent"
                    android:layout_height="80dp"
                    android:background="@drawable/speedup"
                    android:text="24 H" /&gt;
    
            &lt;/LinearLayout&gt;
        &lt;/ScrollView&gt;
    
    
    &lt;/androidx.constraintlayout.widget.ConstraintLayout&gt;
## [10][Beginners Kotlin Android Development](https://www.reddit.com/r/Kotlin/comments/g16u6n/beginners_kotlin_android_development/)
- url: https://www.reddit.com/r/Kotlin/comments/g16u6n/beginners_kotlin_android_development/
---
Hey, so I have an interest in app development but am having A LOT of difficulty learning Kotlin on my own. I  have Android studio and I'm finding it very difficult to use as I am on a Windows and most tutorials are set for a mac and it's very hard to follow along. Is Android Studio the best for beginners with Kotlin? The app I'm trying to design is fairly simple, it's not a game and I believe it would have a similar layout to say Instagram or Pinterest; a main activity page that opens a new page that would store data like videos/switches and stores data. 

Does anyone know which development program would best suit this/have any tips or resources for beginner developers teaching themselves?
