# Kotlin
## [1][I want to develop a custom app for my DJI Drone. DJI SDK is in Java. However, i'm attracted to learning Kotlin rather than Java as it seems the future (for Android). So will i be able to develop an app for my drone in Kotlin?](https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/)
- url: https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/
---
Hi guys, I'm an absolute beginner Android dev, with an end goal to develop a custom app for my DJI drone. I know very little Java &amp; none of Kotlin. The [DJI SDK](https://github.com/dji-sdk/Mobile-SDK-Android) is written in Java.

With my end goal in mind, will it be possible to make a custom app for my drone in Kotlin? Or do I need to stick to Java since the DJI SDK itself is in Java?

Basis the answer to the above question I intend to start learning Android Studio tutorials either specifically for Java or Kotlin. Any guidance would be appreciated, thank you!
## [2][Is there a way to do nested destructuring?](https://www.reddit.com/r/Kotlin/comments/jt37lg/is_there_a_way_to_do_nested_destructuring/)
- url: https://www.reddit.com/r/Kotlin/comments/jt37lg/is_there_a_way_to_do_nested_destructuring/
---
There are two places where it comes up quasi-frequently for me: lambda arguments, and unwrapping data classes.

Example 1:

    data class Foo (val s: String, val i: Int)
    data class Bar (val f: Foo)
    
    fun doStuff () {
        val (f) = makeBar()
        val (_, i) = f
        useI(i)
    }
    
    fun useI(i: Int) {
        TODO()
    }
    
    fun makeBar(): Bar {
        TODO()
    }

I'd like it if I could've written `val ((_, i)) = makeBar()` instead of having `f` as an intermediate.

Example 2:

    data class Foo (val s: String, val i: Int)
    
    fun doStuff () {
        val sum = makeListOfFoo().fold(0) { acc, foo -&gt;
            acc + foo.i
        }
        
        println(sum)
    }
    
    fun makeListOfFoo(): List&lt;Foo&gt; {
        TODO()
    }

I'd like it if I could write my lambda args like `fold(acc) { acc, (_, i) -&gt; `

None of that is possible, right?
## [3][Kotlin Compose + web](https://www.reddit.com/r/Kotlin/comments/jsiggw/kotlin_compose_web/)
- url: https://www.reddit.com/r/Kotlin/comments/jsiggw/kotlin_compose_web/
---
Since Kotlin supports JS, is it possible for Kotlin Compose to be ported to web? (I was thinking whether is possible to match Flutter's platform set)
## [4][The best books to learn Kotlin](https://www.reddit.com/r/Kotlin/comments/jskqd6/the_best_books_to_learn_kotlin/)
- url: https://medium.com/p/the-best-books-to-learn-kotlin-894f9a83e64a
---

## [5][Announcing the http4k Toolbox: Guns for show, knives for a pro](https://www.reddit.com/r/Kotlin/comments/js8iiw/announcing_the_http4k_toolbox_guns_for_show/)
- url: https://http4k.org/blog/toolbox/
---

## [6][Free Udavity Android Kotlin Developer Nanodegree Course material](https://www.reddit.com/r/Kotlin/comments/js3y3n/free_udavity_android_kotlin_developer_nanodegree/)
- url: https://www.reddit.com/r/Kotlin/comments/js3y3n/free_udavity_android_kotlin_developer_nanodegree/
---
Udacity Just lanched their Android Kotlin Developer Nanodegree program, So I thought of compiling all the free resources for interested devs who want to follow that same syllabus for free.

I hope this help :) 

## Free course material for Udacity's Android Kotlin Developer Nanodegree

Syllabus ▶

## Course 1

* Build your First App
* Layouts
* App Navigation
* Activity and Fragment Lifecycle
* App Architecture (UI Layer)

## Course 2

* Recycler View
* Connect to the Internet
* Behind the Scenes
* Designing for
* Everyone

## Link To Course 1 and Course 2

[Course 1 and Course 2](https://www.udacity.com/course/developing-android-apps-with-kotlin--ud9012)

## Course 3

* Using Notifications
* Creating Custom Views
* Drawing on Canvas Objects
* Clipping Canvas Objects
* Android Property Animations
* Using Motion Layout to Animate Android Apps

## Course 4

* Wandering in Google Maps with Kotlin
* Virtual Treasure Hunt with Geofences
* Testing: Basics
* Introduction to Test Doubles and Dependency Injection
* Survey of Advanced Testing Topics
* Implementing Login on Android with FirebaseUI

## Link to Course 3 and Course 4 ▶

[Course 3 and Course 4](https://www.udacity.com/course/advanced-android-with-kotlin--ud940)
## [7][mltype - Typing practice for Kotlin and other languages](https://www.reddit.com/r/Kotlin/comments/js76mq/mltype_typing_practice_for_kotlin_and_other/)
- url: https://www.reddit.com/r/Kotlin/comments/js76mq/mltype_typing_practice_for_kotlin_and_other/
---
**What is it?**

Command line tool that uses a character-level LSTM model to generate text that resembles a real language (including programming languages). One can both train a network from scratch or download a pretrained one (Kotlin, Java, Go, JavaScript, C, C++, Python,..).

**Motivation**

I recently switched to touch typing and I realized that there is basically no way to practise typing of programming languages (other than actually programming). Additionally, I revisited the famous blog post [http://karpathy.github.io/2015/05/21/rnn-effectiveness/](http://karpathy.github.io/2015/05/21/rnn-effectiveness/) and thought it would be cool to use a model like this to generate infinite amount of custom text to type.

&amp;#x200B;

[Demo](https://i.redd.it/7bm2qysxsly51.gif)

**Links**

* docs: [https://mltype.readthedocs.io/en/latest/](https://mltype.readthedocs.io/en/latest/)
* github: [https://github.com/jankrepl/mltype](https://github.com/jankrepl/mltype)

**PS**

I trained the Kotlin model on LeakCanary.
## [8][How can I create another smile in order to move the ball to the right?? PS:the smile's variables have to stay as vals, not vars](https://www.reddit.com/r/Kotlin/comments/jsahix/how_can_i_create_another_smile_in_order_to_move/)
- url: https://www.reddit.com/r/Kotlin/comments/jsahix/how_can_i_create_another_smile_in_order_to_move/
---
 Teacher asked me to create a Smile in canvas and when I press right the smile moves right. He also told me to leave the smile's variables as vals not as vals, for me to move the ball I had to create a new smile with the old smile's values. How?? 

&amp;#x200B;

&amp;#x200B;

    data class Smile(val x: Int, val y: Int, val  radius: Int)
    
    fun drawSmile( smile: Smile, canvas: Canvas):Unit{
        canvas.drawCircle(smile.x,smile.y, smile.radius, pt.isel.canvas.YELLOW)
        canvas.drawCircle(smile.x, smile.y, smile.radius, 0, (0.07*smile.radius).toInt())
        canvas.drawArc(smile.x,smile.y, (0.6*smile.radius).toInt(),200,340, 0, (0.07*smile.radius).toInt())
        canvas.drawCircle(smile.x-(0.35*smile.radius).toInt(),smile.y-(0.35*smile.radius).toInt(),(0.125*smile.radius).toInt(), color= BLACK)
        canvas.drawCircle(smile.x+(0.35*smile.radius).toInt(),smile.y-(0.35*smile.radius).toInt(),(0.125*smile.radius).toInt(), color= BLACK)
    
    fun move(smile: Smile, canvas: Canvas, ke: Int, increment: Int = 4): Unit {
      canvas.erase()
      val x: Int = smile.x
      val y: Int = smile.y
      val rad = smile.radius
      val smile: Smile
      when (ke) {
        39 -&gt; smile = Smile(x + 4, y, rad)
        else -&gt; smile = Smile(x, y, rad)
      }
      drawSmile(smile, canvas)
    }
    
    fun main() {
      println("Hi, Smile! :)")
    
      onStart {
        val canvas = Canvas(600, 400, CYAN)
        //FIRST SMILE
        val xCenter = 600 / 2
        val yCenter = 400 / 2
        val radius = 50
        var smile = Smile(xCenter, yCenter, radius)
        println("Move the smile with the arrows or mouse")
        println("Grow and shrink with + and -")
        drawSmile(smile, canvas)
        canvas.onKeyPressed { ke -&gt;
          move(smile, canvas, ke.code)
        }
    
        canvas.onMouseDown { MousePlacer -&gt;
          placeSmile(smile, canvas, MousePlacer.x, MousePlacer.y)
        }
    
        onFinish {
          println("BYE")
        }
      }
    }
## [9][Ktor 1.4.2 has been released](https://www.reddit.com/r/Kotlin/comments/jrn38o/ktor_142_has_been_released/)
- url: https://ktor.io/changelog/#version-1-4-2
---

## [10][Kotlin Trivia Quiz App](https://www.reddit.com/r/Kotlin/comments/js33ly/kotlin_trivia_quiz_app/)
- url: https://youtu.be/uB4ePepyVBo
---

