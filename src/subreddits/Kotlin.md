# Kotlin
## [1][Doodle Tutorials released](https://www.reddit.com/r/Kotlin/comments/juu69x/doodle_tutorials_released/)
- url: https://www.reddit.com/r/Kotlin/comments/juu69x/doodle_tutorials_released/
---
You can now see example apps and get more in-depth explanation of Doodle's concepts at the  [tutorials](https://nacular.github.io/doodle-tutorials) site. The first example walks through building a simple calculator app.

[Doodle](https://github.com/nacular/doodle) is a pure Kotlin UI framework for the Web, that lets you create rich applications without worrying about Javascript, HTML or CSS. Check out the [documentation](https://nacular.github.io/doodle) to learn more.

Feedback is always appreciated.
## [2][Shared flows, broadcast channels](https://www.reddit.com/r/Kotlin/comments/jv6o4e/shared_flows_broadcast_channels/)
- url: https://medium.com/@elizarov/shared-flows-broadcast-channels-899b675e805c
---

## [3][From Spring Boot to Quarkus](https://www.reddit.com/r/Kotlin/comments/juoaz6/from_spring_boot_to_quarkus/)
- url: https://blog.frankel.ch/spring-to-quarkus/
---

## [4][Is there a stdlib functionality to "add" to an immutable list, so that a new list is returned?](https://www.reddit.com/r/Kotlin/comments/juqktq/is_there_a_stdlib_functionality_to_add_to_an/)
- url: https://www.reddit.com/r/Kotlin/comments/juqktq/is_there_a_stdlib_functionality_to_add_to_an/
---
Really quick question, is it possible to do something like

    val list1 = listOf(1, 2)
    val list2 = list1.someKindOfAdd(5)

Inside the stdlib, so that list1 remains unchanged and list2 has the elements 1, 2, 5? I mean it would be quite simple to implement it on my own, but I wanna know if there's some standard way for this.
## [5][Started learning Kotlin, but build time is really stressing me out.](https://www.reddit.com/r/Kotlin/comments/jusy5a/started_learning_kotlin_but_build_time_is_really/)
- url: https://www.reddit.com/r/Kotlin/comments/jusy5a/started_learning_kotlin_but_build_time_is_really/
---
So I decided to start learning Kotlin. This is my first experience with coding. I'm using some Google resources taking things slow step by step learning the basics and I've downloaded Android Studio.

My problem is that when I want to test what I've learned to follow the examples and run the code on my Galaxy S9 it takes like 11 minutes just to build a single page app that has two strings of text and a single image.

I'm not sure if it's because of my laptop which is pretty old because laptops in the country I'm in are expensive. I'm using an HP Elitebook 8570 with i5 processor and 4 gb of RAM so I know it's far from the best and I can't run the emulator which is why I plug my S9 in to run the program.

Any advice? Are there any options I can disable to make the grade build time faster? Once it finishes it installs onto my phone in seconds its everything in between that takes 10 to 15 minutes. Should I buy another stick of 4gb RAM? Or is this just how it is?

I'm really enthusiastic about learning Kotlin but these past two days the build time has really but a damper on it.

Edit: 

I seemed to have solved the problem, I did two things so I'm not sure which did the trick perhaps you'll know.

First I went to File&gt;Power Save Mode and enabled that, which from my understanding disables a few background operations.

Second, I went to Tools&gt;Kotlin&gt;Enable migrations detection (experimental) and Disabled that. (Not sure why it was automatically enabled if it's experimental). 

Anyways I tried running the app to my S9 and in 30 seconds it was done. I'll repeat with and without said options disabled/enabled to see if its a fluke or if it actually is the solution I require.
## [6][How to avoid passing objects between activities](https://www.reddit.com/r/Kotlin/comments/juiyk4/how_to_avoid_passing_objects_between_activities/)
- url: https://www.reddit.com/r/Kotlin/comments/juiyk4/how_to_avoid_passing_objects_between_activities/
---
I suspect that I am using the wrong code design for my app. I am making a note taking app where each note is an object. In the main activity I create an empty `Notebook` object. In the `Notebook` is an empty mutableListOf of `Notes`s. Each `Note` has the context of the note. (text, date created, date updated ...). When the user clicks on the "Add a note" button it jumps to the CreateNote activity. 

I can think of two ways of adding that note. Either, 
 1. Serialize the `Notebook` object
 2. Send the `Notebook` by using 
```
// In main activity
intent.putExtra(NOTEBOOK, notebook)
// In CreateNote activity
val notebook = intent.getSerializableExtra(NOTEBOOK) as? Notebook
```
 3. Add a newly created note to the notebook
 4. Send back the updated notebook
 5. Update the notebook in main activity.

Alternatively, I can create a new `Note`, serialize it, and send it back. 


Is that the best coding style? If there was a way of storing the `Notebook` object as a shared object between activities I think that would be the best. I could then simply access that object (and potentially send an index if I want to edit a specific note). Is there a way to do that?
## [7][Any doc/article/reference for the new Kotlin's IR (intermediate representation)?](https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/)
- url: https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/
---
My mind was blown after seeing the [kotlin-power-assert](https://github.com/bnorm/kotlin-power-assert), and I'm eager to know about this (IR).

The [sources](https://github.com/JetBrains/kotlin/tree/master/compiler/ir) are so wide (even though well documented almost one third of those), its hard to follow a path i.e. where to see and understand the things.

I was trying to find a way to write a compiler plugin which can modify the intermediate code, so that it will easily target multiplatform, but not able to find any quick-start or something for where to start or start looking at the sources.

If anybody know already about this, he/she can share his experience and help for how to get started (basically I want to know how to traverse all the function/classes and then some way to find linked/called functions, etc).
## [8][Kotlin Coroutines in Spring Webflux](https://www.reddit.com/r/Kotlin/comments/ju7td5/kotlin_coroutines_in_spring_webflux/)
- url: https://medium.com/@tien.donam/asynchronous-spring-using-kotlin-coroutines-and-r2dbc-93b3a079ac22
---

## [9][Help, Why isnt my smile moving? I made the mouse event and the key event but nothing happens](https://www.reddit.com/r/Kotlin/comments/juet5m/help_why_isnt_my_smile_moving_i_made_the_mouse/)
- url: https://www.reddit.com/r/Kotlin/comments/juet5m/help_why_isnt_my_smile_moving_i_made_the_mouse/
---
import pt.isel.canvas.\*  
import java.nio.file.WatchEvent  
import kotlin.math.round  
import kotlin.random.Random  


data class Point(val x: Int, val y: Int)  


data class Smile(val center: Point, val radius: Int, val startAng: Int , val endAng: Int ,val color: Int, val thickness: Int)  


fun createSmile(smile: Smile, x: Int = smile.center.x, y: Int = smile.center.y, radius: Int = smile.radius, startAng: Int = smile.startAng, endAng: Int = smile.endAng, color: Int = smile.color, thickness: Int = smile.thickness): Smile = *createSmile*(x, y, radius, startAng, endAng, color, thickness)  


fun createSmile(x: Int, y: Int, radius: Int, startAng: Int, endAng: Int, color: Int, thickness: Int): Smile {  
 val center = Point(x, y)  
 return Smile(center, radius, startAng, endAng, color, thickness)  


}  


fun drawSmile(smile: Smile, canvas: Canvas): Unit {  
canvas.drawCircle(smile.center.x, smile.center.y, smile.radius, smile.color)  
canvas.drawCircle(smile.center.x, smile.center.y, smile.radius, *BLACK*, smile.thickness)  
canvas.drawArc(smile.center.x, smile.center.y, 30, 200, 340, *BLACK*, smile.thickness)  
canvas.drawCircle(smile.center.x - 17, smile.center.y - 17, 6, *BLACK* )  
canvas.drawCircle(smile.center.x + 17, smile.center.y - 17, 6, *BLACK* )  


}  


fun main() {  
 val WIDTH = 600  
 val HEIGHT = 400  
 val RADIUS = 50  
 *onStart* **{**  
 val cv = Canvas(WIDTH, HEIGHT, *CYAN*)  


var smile = *createSmile*(WIDTH/2, HEIGHT / 2, 50, 0, 360, *YELLOW*, 3)  
 *drawSmile*(smile, cv)  


cv.onMouseDown **{** mouse **-&gt;**  
 smile = *createSmile*(smile, x = mouse.x, y = mouse.y)  


**}**  
 cv.onKeyPressed **{** ke **-&gt;**  
 smile = when (ke.code) {  
 *LEFT\_CODE* \-&gt; *createSmile*(smile, (smile.center.x - 4), smile.center.y)  
 *RIGHT\_CODE* \-&gt; *createSmile*(smile, (smile.center.x + 4), smile.center.y)  
 *UP\_CODE* \-&gt; *createSmile*(smile, smile.center.x, (smile.center.y - 4))  
 *DOWN\_CODE* \-&gt; *createSmile*(smile, smile.center.x, (smile.center.y + 4))  


else -&gt; when (ke.char) {  
 '+' -&gt; *createSmile*(smile, (smile.radius + 1))  
 '-' -&gt; *createSmile*(smile, (smile.radius - 1))  
 else -&gt; smile  
}  


}  


**}**  
**}**  
 *onFinish* **{**  
**}**  
 }
## [10][Plan from a theoretical point of view and implement in Kotlin the solution to the following problem:](https://www.reddit.com/r/Kotlin/comments/juboc8/plan_from_a_theoretical_point_of_view_and/)
- url: https://www.reddit.com/r/Kotlin/comments/juboc8/plan_from_a_theoretical_point_of_view_and/
---
Plan from a theoretical point of view and implement in Java the solution to the following problem:

We want to know how many different ways exist to obtain certain integer number as a sum of smaller positive numbers, taking into account a certain list of numbers that cannot be used.

Considerations

• If A + B = B + A, only one of them should be indicated

• In addition, A &gt; 0 and B &gt; 0

Input

Each case is described in one line with an integer number X, which is the value to obtain, and a list of M positive numbers N0, N1 ... NM-1 separated by a space. These values cannot be used as operands.

Output

For each case T, the output must be the string "#t:" followed by the total number of possible ways. For example

Sample input

3 1

6 1 3

Sample output

\#1: 0

\#2: 2

In the first case (3), there is no way to get 3 from 1 and 2 without using operand 1. The second case (6) can be calculated with operations 2 + 2 + 2 and 4 + 2, so there are 2 possible ways without using operands 1 and 3.

The objective is to create the most efficient algorithm able to perform this calculation

Any tips )
