# Kotlin
## [1][Kotlin pattern matching with instanceof (the Java 14 way)](https://www.reddit.com/r/Kotlin/comments/fl9ry0/kotlin_pattern_matching_with_instanceof_the_java/)
- url: https://www.reddit.com/r/Kotlin/comments/fl9ry0/kotlin_pattern_matching_with_instanceof_the_java/
---
In Java 14 we are now enabled to use pattern matching with `instanceof` expressions.
e.g. on my code

```
if (psiVariableReferenceContext instanceof PsiExpressionList psiExpressionList &amp;&amp;
    psiExpressionList.getContext() instanceof PsiMethodCallExpression psiMethodCallExpr &amp;&amp;
    !isMethodParameterCharSequence(psiVariableReference, psiMethodCallExpr)) {
  return false;
}
```

Is there a similar Kotlin alternative?  
Currently I need to re-cast the instance and call the getter a second time.

```
if (psiVariableReferenceContext is PsiExpressionList &amp;&amp;
    psiVariableReferenceContext.context is PsiMethodCallExpression &amp;&amp;
    !isMethodParameterCharSequence(
      psiVariableReference,
      psiVariable.context as PsiMethodCallExpression
    )) {
  return false
}
```
## [2][How do you work with Network System Discovery? [Question]](https://www.reddit.com/r/Kotlin/comments/fl5yb6/how_do_you_work_with_network_system_discovery/)
- url: https://www.reddit.com/r/Kotlin/comments/fl5yb6/how_do_you_work_with_network_system_discovery/
---
1. I am having a hard time checking whether or not [NSD](https://developer.android.com/training/connect-devices-wirelessly/nsd) is working or not. The resources online are scarce, slack isn't answering, documentation is well made but nothing past the bare minimum. All in all catastrophe... I am basically registering a service, discovering, and then I keep discovering myself on my emulators (I run 2 emulators at the same time). By myself I mean I discover the emulator I registered on and not the other....
2.  Could someone explain a little bit what "service type" and "service name" are. As service type when registering equals to \_http.\_tcp NO DOT, and when listening to a protocol it's with A DOT at the end.  And an example for this with real machine (printer, PC, another phone). 
3. Is there an easier way to work with 2 emulators without having to go into android studio and change the device everytime? Is there a way to split screen logcat and just have both emulators logcats open? (if not where can i ask for this feature?) Or is there another way to test an application that talks to another device. I mean you have to  have 2 emulators open?  


Thank you for your time. I am stuck on this for like 3 or so days. It's been frustrating. The examples online are horrible and just a copy paste of the documentation.
## [3][State-full sequences](https://www.reddit.com/r/Kotlin/comments/fl3nq7/statefull_sequences/)
- url: https://www.reddit.com/r/Kotlin/comments/fl3nq7/statefull_sequences/
---
So here's the behaviour I want:

```Kotlin
val nums = sequence {
    var n = 0
    while (true) { yield(n++) }
}

nums.find { isPrime(it) } // 2
nums.first()              // Want : 3 but I get 0
nums.take(5).toList()     // Want : [4, 5, 6, 7, 8] but I get [0, 1, 2, 3, 4]
```

Any ideas on how I could do something like this?

---

Here's my actual code which essentially require a state-full sequence.

Here's my actual code:
```Kotlin
// See 0 for setup.
while (true) {
    // See 1 for the logic that's implemented below (if sequences had state)
    val data = (blockingQueuesLoopingSequence.take(numQueues) 
        .find { it.peek() != null } ?: blockingQueuesLoopingSequence.first())
        .get()
    // Do stuff with data
}
```
### 0
Here i have a sequence that loops infinitely over a list of `blockingQueues`. 

### 1
Think of a Round-Robin scheduler, current position in the list 
could be any particular queue which would be our state.  
Say we are on the `m^th` queue and we have a total on `N` unique queues.

From all Queues starting from position `m` (our state) looping back to `m-1th` queue.  
Find the first queue that has something available. (Now state is m')

If no such queue choose the first queue from current position 
(which will be `m' = (m + N) mod N = m`)

Call a blocking get on the chosen queue.

Unfortunately, the state(`m`) of `blockingQueuesLoopingSequence` doesn't update and `m is always 0`.
## [4][Equalizer type Progress Indicator](https://www.reddit.com/r/Kotlin/comments/fl6cz3/equalizer_type_progress_indicator/)
- url: /r/androiddev/comments/fko6rs/equalizer_type_progress_indicator/
---

## [5][How can I pass different version of object without altering it as argument?](https://www.reddit.com/r/Kotlin/comments/fl6dom/how_can_i_pass_different_version_of_object/)
- url: https://www.reddit.com/r/Kotlin/comments/fl6dom/how_can_i_pass_different_version_of_object/
---
I have a code like this

val x = MyObject("first", "second")

function1(x)

function2(x)

I want to send different version of x to function 1 but without actually altering it so function2 can have the first version. I have tried let, run, apply and also but couldn't get a solution.
## [6][can we use Spring with Kotlin/Native?](https://www.reddit.com/r/Kotlin/comments/fkn5ko/can_we_use_spring_with_kotlinnative/)
- url: https://www.reddit.com/r/Kotlin/comments/fkn5ko/can_we_use_spring_with_kotlinnative/
---
Hi, we have some java business logic that we want to migrate to Kotlin and use as library on multiplatform. Spring is used in the java code, can we reuse it directly in the new library that will be also integrated in iOS/Andoird solutions?

If we have to use Kotlin DI solution in the library, can it be mixed back with Spring-based java apps that would consume this new library?

thanks!
## [7][Good API design patterns to help imports be more seamless?](https://www.reddit.com/r/Kotlin/comments/fk1r38/good_api_design_patterns_to_help_imports_be_more/)
- url: https://www.reddit.com/r/Kotlin/comments/fk1r38/good_api_design_patterns_to_help_imports_be_more/
---
My web framework [Kweb](https://docs.kweb.io/) incorporates quite a sophisticated DSL involving a large number of quite short functions (often corresponding to HTML tags).

My instinct was to break these small functions up into different packages according to purpose, but this had the effect of relying much more heavily on IntelliJ to auto-import symbols when necessary.

Is there a better pattern?  Should I just say "screw it" and put them all in `kweb.*` for ease of importing?
## [8][I'm beginner I'm trying to make Tic Tac Toe game but when I press any button the app crashes and the device don't choose any button](https://www.reddit.com/r/Kotlin/comments/fk22et/im_beginner_im_trying_to_make_tic_tac_toe_game/)
- url: https://www.reddit.com/r/Kotlin/comments/fk22et/im_beginner_im_trying_to_make_tic_tac_toe_game/
---
    
    import android.graphics.Color
    import androidx.appcompat.app.AppCompatActivity
    import android.os.Bundle
    import android.view.View
    import android.widget.Button
    import android.widget.Toast
    import kotlinx.android.synthetic.main.activity_main.*
    import kotlin.random.Random
    
    
    class MainActivity : AppCompatActivity() {
    
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
        }
         fun buClick(view:View){
       val buSelected:Button = view as Button
            var cellID=0
            when(buSelected.id){
                R.id.bu1-&gt; cellID=1
                R.id.bu2-&gt; cellID=2
                R.id.bu3-&gt; cellID=3
                R.id.bu4-&gt; cellID=4
                R.id.bu5-&gt; cellID=5
                R.id.bu6-&gt; cellID=6
                R.id.bu7-&gt; cellID=7
                R.id.bu8-&gt; cellID=8
                R.id.bu9-&gt; cellID=9
            }
    
             //Log.d("buClick:",buSelected.id.toString())
             //Log.d("buClick:cellID:",cellID.toString())
           PlayGame(cellID,buSelected)
        }
        var Player1=ArrayList&lt;Int&gt;()
        var Player2=ArrayList&lt;Int&gt;()
        var ActivePlayer=1
        fun PlayGame(cellID:Int,buSelected:Button){
            if (ActivePlayer==1){
                buSelected.text="X"
                buSelected.setBackgroundColor(Color.BLUE)
                Player1.add(cellID)
                ActivePlayer=2
                AutoPlay()
            }else{
                buSelected.text="O"
                buSelected.setBackgroundColor(Color.RED)
                Player2.add(cellID)
                ActivePlayer=1
            }
            buSelected.isEnabled=false
            CheckWinner()
        }
        fun CheckWinner() {
            var Winner = -1
            //Row1
            if (Player1.contains(1) &amp;&amp; Player1.contains(2) &amp;&amp; Player1.contains(3)) {
                Winner = 1
            }
            if (Player2.contains(1) &amp;&amp; Player1.contains(2) &amp;&amp; Player2.contains(3)) {
                Winner = 2
            }
    
            //Row2
            if (Player1.contains(4) &amp;&amp; Player1.contains(5) &amp;&amp; Player1.contains(6)) {
                Winner = 1
            }
            if (Player2.contains(4) &amp;&amp; Player2.contains(5) &amp;&amp; Player2.contains(6)) {
                Winner = 2
            }
    
            //Row3
            if (Player1.contains(7) &amp;&amp; Player1.contains(8) &amp;&amp; Player1.contains(9)) {
                Winner = 1
            }
            if (Player2.contains(7) &amp;&amp; Player2.contains(8) &amp;&amp; Player2.contains(9)) {
                Winner = 2
            }
    
    
            //Column1
            if (Player1.contains(1) &amp;&amp; Player1.contains(4) &amp;&amp; Player1.contains(7)) {
                Winner = 1
            }
            if (Player2.contains(1) &amp;&amp; Player2.contains(4) &amp;&amp; Player2.contains(7)) {
                Winner = 1
            }
    
            //Column 2
            if (Player1.contains(2) &amp;&amp; Player1.contains(5) &amp;&amp; Player1.contains(8)) {
                Winner = 1
            }
            if (Player2.contains(2) &amp;&amp; Player2.contains(5) &amp;&amp; Player2.contains(8)) {
                Winner = 2
            }
    
            //Column 3
            if (Player1.contains(3) &amp;&amp; Player1.contains(6) &amp;&amp; Player1.contains(9)) {
                Winner = 1
            }
            if (Player2.contains(3) &amp;&amp; Player2.contains(6) &amp;&amp; Player2.contains(9)) {
                Winner = 2
            }
    
    
            //Cross from Left
            if (Player1.contains(1) &amp;&amp; Player1.contains(5) &amp;&amp; Player1.contains(9)) {
                Winner = 1
            }
    
            if (Player2.contains(1) &amp;&amp; Player2.contains(5) &amp;&amp; Player2.contains(9)) {
                Winner = 2
            }
    
            //Cross from Left
            if (Player1.contains(3) &amp;&amp; Player1.contains(5) &amp;&amp; Player1.contains(7)) {
                Winner = 1
            }
            if (Player2.contains(3) &amp;&amp; Player2.contains(5) &amp;&amp; Player2.contains(7)) {
                Winner = 2
            }
    
    
            if (Winner != -1) {
                if (Winner == 1) {
                    Toast.makeText(this, "X is the WINNER", Toast.LENGTH_LONG).show()
                } else {
                    Toast.makeText(this, "O is the WINNER", Toast.LENGTH_LONG).show()
                }
    
            }
    
    
        }
    private fun AutoPlay(){
        //SCAN EMPTY CELLS
        var EmptyCell=ArrayList&lt;Int&gt;()
    for (CellID in 1..9){
        if (!Player1.contains(CellID)||!Player2.contains(CellID)){
            EmptyCell.add(CellID)
        }
        //SELECT RANDOM INDEX
        val r = Random
        val randIndex=r.nextInt(EmptyCell.size-0)+0
        val CellID=EmptyCell[randIndex]
        //CONVERT INDEX INTO BUTTON
        var buSelect:Button?
        when(CellID){
            1-&gt;buSelect=bu1
            2-&gt;buSelect=bu2
            3-&gt;buSelect=bu3
            4-&gt;buSelect=bu4
            5-&gt;buSelect=bu5
            6-&gt;buSelect=bu6
            7-&gt;buSelect=bu7
            8-&gt;buSelect=bu8
            9-&gt;buSelect=bu9
            else-&gt;{
                buSelect=bu1
            }
        }
        PlayGame(CellID,buSelect)
    }
    }
    
    
    }
## [9][MVIKotlin - new Kotlin Multiplatform framework for MVI](https://www.reddit.com/r/Kotlin/comments/fjjupo/mvikotlin_new_kotlin_multiplatform_framework_for/)
- url: https://www.reddit.com/r/Kotlin/comments/fjjupo/mvikotlin_new_kotlin_multiplatform_framework_for/
---
Checkout it here: [https://github.com/arkivanov/MVIKotlin](https://github.com/arkivanov/MVIKotlin)

Features:

* Multiplatform: Android, JVM, JavaScript, iosX64, iosArm64, linuxX64
* Does not depend on any reactive library or coroutines
* Extensions for [Reaktive](https://github.com/badoo/Reaktive) library
* Extensions for [Coroutines](https://github.com/Kotlin/kotlinx.coroutines)
* Multithreading friendly (freezable in Kotlin Native if needed)
* Logging functionality with adjustable verboseness and ability to provide custom logger
* Time travel feature:
   * Multiplatform for all supported targets
   * Plug-and-play UI for Android and iOS
   * Export/import events for JVM and Android
## [10][How brevity in Kotlin increases developers' product as compare to Java with example](https://www.reddit.com/r/Kotlin/comments/fk1k4s/how_brevity_in_kotlin_increases_developers/)
- url: https://www.vtnetzwelt.com/mobile/java-vs-kotlin/
---

