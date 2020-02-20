# Kotlin
## [1][Kotlin class cant reach Java Class fields](https://www.reddit.com/r/Kotlin/comments/f6t1h2/kotlin_class_cant_reach_java_class_fields/)
- url: https://www.reddit.com/r/Kotlin/comments/f6t1h2/kotlin_class_cant_reach_java_class_fields/
---
Im making a Spring Project basically in Java but for some classes I want to use Kotlin practice. I have a Pojo Java Class named Item with 2 private fields(email and URL) and with the @Data annotation so it should have getters and setters.
In my Kotlin class I want to use those fields and when i write item.email as I should in Kotlin the IntelliJ lets me but on compile it says that those are private and I cant reach them.

I had searched for this problem but didnt found anything usefull.
## [2][Kotlin: Fun with “in”](https://www.reddit.com/r/Kotlin/comments/f6aq8r/kotlin_fun_with_in/)
- url: https://proandroiddev.com/kotlin-fun-with-in-8a425704b635
---

## [3][swagger-gradle-codegen v1.4.0 just got released with support for moshi-codegen and a lot of bugfixes/improvements.](https://www.reddit.com/r/Kotlin/comments/f6h990/swaggergradlecodegen_v140_just_got_released_with/)
- url: https://github.com/Yelp/swagger-gradle-codegen/releases/tag/1.4.0
---

## [4][LiveStream-kt - Library to emit &amp; observe data from any module of the application.](https://www.reddit.com/r/Kotlin/comments/f6lx2c/livestreamkt_library_to_emit_observe_data_from/)
- url: https://github.com/PatilShreyas/LiveStream-kt
---

## [5][Can't figure out how to put this into a function to make the code D.R.Y.](https://www.reddit.com/r/Kotlin/comments/f6h24j/cant_figure_out_how_to_put_this_into_a_function/)
- url: https://www.reddit.com/r/Kotlin/comments/f6h24j/cant_figure_out_how_to_put_this_into_a_function/
---
&amp;#x200B;

Hi, so I was trying to make this code more dry (code here: [https://codeshare.io/214R8z](https://codeshare.io/214R8z), and this is editable). The part I was trying to change is in the placePiece() function. It's just too repetitive.

(It's a tictactoe cli game)

So I made this function:

    fun checkPlacePiece(pos: Int,input: Int, gameBoard: Array&lt;Array&lt;Char&gt;&gt;, player: Char, row: Int, col: Int, playerHasToBeSwitched: Boolean) {
        var playerHasToBeSwitched = playerHasToBeSwitched
        if (pos == input) {
            if (gameBoard[row][col] != 'X' &amp;&amp; gameBoard[row][col] != 'O') {
                gameBoard[row][col] = player
            } else if (gameBoard[row][col] == 'X' || gameBoard[row][col] == 'O') {
                println("This place is already taken!")
    
              
                playerHasToBeSwitched = true
            }
        }
    }
    
        But that doesn't work as expected 
    Can anyone help me, I am struggling so much
## [6][Can I download Kotlin on a Chromebook?](https://www.reddit.com/r/Kotlin/comments/f6fl69/can_i_download_kotlin_on_a_chromebook/)
- url: https://www.reddit.com/r/Kotlin/comments/f6fl69/can_i_download_kotlin_on_a_chromebook/
---

## [7][FYI Kotlin programs on benchmarks game](https://www.reddit.com/r/Kotlin/comments/f6kkyl/fyi_kotlin_programs_on_benchmarks_game/)
- url: https://discuss.kotlinlang.org/t/fyi-kotlin-programs-on-benchmarks-game/16360
---

## [8][How to check if T is sometype ?](https://www.reddit.com/r/Kotlin/comments/f67z36/how_to_check_if_t_is_sometype/)
- url: https://www.reddit.com/r/Kotlin/comments/f67z36/how_to_check_if_t_is_sometype/
---
I have code like this 


    object BindUtil {
        fun &lt;T&gt; bind(editText: EditText, model: KMutableProperty0&lt;out T&gt;) {
            editText.afterTextChanged {
                if (**T is double**) {
                    model.setter.call(it.toDouble())
                } else {
                    model.setter.call(it)
                }
            }
        }
    }

it doesnt like "T is Double", how do I fix this ?
## [9][Can't figure out how to put this into a function to make the code D.R.Y.](https://www.reddit.com/r/Kotlin/comments/f69sv0/cant_figure_out_how_to_put_this_into_a_function/)
- url: https://www.reddit.com/r/Kotlin/comments/f69sv0/cant_figure_out_how_to_put_this_into_a_function/
---
Hi, so I was trying to make this code more dry (code here: [https://codeshare.io/214R8z](https://codeshare.io/214R8z), and this is editable). The part I was trying to change is in the placePiece() function. It's just too repetitive.

(It's a tictactoe cli game)

&amp;#x200B;

So I made this function:

&amp;#x200B;

    fun checkPlacePiece(pos: Int,input: Int, gameBoard: Array&lt;Array&lt;Char&gt;&gt;, player: Char, row: Int, col: Int, playerHasToBeSwitched: Boolean) {
        var playerHasToBeSwitched = playerHasToBeSwitched
        if (pos == input) {
            if (gameBoard[row][col] != 'X' &amp;&amp; gameBoard[row][col] != 'O') {
                gameBoard[row][col] = player
            } else if (gameBoard[row][col] == 'X' || gameBoard[row][col] == 'O') {
                println("This place is already taken!")
    
                // todo update comment here
                playerHasToBeSwitched = true
            }
        }
    }
    
    
    
    But that doesn't work as expected
    
    Can anyone help me, I am struggling so much
## [10][Hello Kotlin Coroutine - basic intro to coroutines](https://www.reddit.com/r/Kotlin/comments/f6cq4y/hello_kotlin_coroutine_basic_intro_to_coroutines/)
- url: https://medium.com/@MohamedISoliman/hello-kotlin-coroutines-4e40cb9a106c
---

