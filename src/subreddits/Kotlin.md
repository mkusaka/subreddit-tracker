# Kotlin
## [1][Good API design patterns to help imports be more seamless?](https://www.reddit.com/r/Kotlin/comments/fk1r38/good_api_design_patterns_to_help_imports_be_more/)
- url: https://www.reddit.com/r/Kotlin/comments/fk1r38/good_api_design_patterns_to_help_imports_be_more/
---
My web framework [Kweb](https://docs.kweb.io/) incorporates quite a sophisticated DSL involving a large number of quite short functions (often corresponding to HTML tags).

My instinct was to break these small functions up into different packages according to purpose, but this had the effect of relying much more heavily on IntelliJ to auto-import symbols when necessary.

Is there a better pattern?  Should I just say "screw it" and put them all in `kweb.*` for ease of importing?
## [2][I'm beginner I'm trying to make Tic Tac Toe game but when I press any button the app crashes and the device don't choose any button](https://www.reddit.com/r/Kotlin/comments/fk22et/im_beginner_im_trying_to_make_tic_tac_toe_game/)
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
## [3][MVIKotlin - new Kotlin Multiplatform framework for MVI](https://www.reddit.com/r/Kotlin/comments/fjjupo/mvikotlin_new_kotlin_multiplatform_framework_for/)
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
## [4][How brevity in Kotlin increases developers' product as compare to Java with example](https://www.reddit.com/r/Kotlin/comments/fk1k4s/how_brevity_in_kotlin_increases_developers/)
- url: https://www.vtnetzwelt.com/mobile/java-vs-kotlin/
---

## [5][Which programming languages and libraries should I learn to become a complete kotlin android developer?](https://www.reddit.com/r/Kotlin/comments/fj35j3/which_programming_languages_and_libraries_should/)
- url: https://www.reddit.com/r/Kotlin/comments/fj35j3/which_programming_languages_and_libraries_should/
---

## [6][Any of you guys switched to Kotlin and then back to Java?](https://www.reddit.com/r/Kotlin/comments/fj6a2c/any_of_you_guys_switched_to_kotlin_and_then_back/)
- url: /r/java/comments/fj24n8/any_of_you_guys_switched_to_kotlin_and_then_back/
---

## [7][Starting a project, quick question on language](https://www.reddit.com/r/Kotlin/comments/fjciq0/starting_a_project_quick_question_on_language/)
- url: https://www.reddit.com/r/Kotlin/comments/fjciq0/starting_a_project_quick_question_on_language/
---
Hi, I'm about to start a project that involves managing files on devices such as my Android phone and my computer. I know I need to use something on the JVM and I'm trying to decide between just Java, Scala and kotlin, but if a different language would definitely be better than those feel free to suggest it. I know this is a Kotlin page but if possible please try to be unbiased. 

Thanks in advance.
## [8][From Apache Cordova to Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/fj4fuo/from_apache_cordova_to_kotlin_multiplatform/)
- url: http://whereisdarran.com/2020/03/from-apache-cordova-to-kotlin-multiplatform/
---

## [9][How the JVM converts bytecode to machine code.](https://www.reddit.com/r/Kotlin/comments/fipcai/how_the_jvm_converts_bytecode_to_machine_code/)
- url: https://www.reddit.com/r/programming/comments/fipbgk/how_the_jvm_converts_bytecode_to_machine_code/?utm_medium=android_app&amp;utm_source=share
---

## [10][Kotlin data class builders](https://www.reddit.com/r/Kotlin/comments/fja3v4/kotlin_data_class_builders/)
- url: https://www.rockandnull.com/kotlin-data-class-builders/
---

