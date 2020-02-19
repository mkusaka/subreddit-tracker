# Kotlin
## [1][How to check if T is sometype ?](https://www.reddit.com/r/Kotlin/comments/f67z36/how_to_check_if_t_is_sometype/)
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
## [2][Can't figure out how to put this into a function to make the code D.R.Y.](https://www.reddit.com/r/Kotlin/comments/f69sv0/cant_figure_out_how_to_put_this_into_a_function/)
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
## [3][JTRANSC. Convert your Java, Kotlin and Scala code into JavaScript, C++, D, C#, PHP,](https://www.reddit.com/r/Kotlin/comments/f60et4/jtransc_convert_your_java_kotlin_and_scala_code/)
- url: https://www.reddit.com/r/Kotlin/comments/f60et4/jtransc_convert_your_java_kotlin_and_scala_code/
---
[https://www.kotlinresources.com/library/jtransc/](https://www.kotlinresources.com/library/jtransc/)
## [4][Name Mangling in Kotlin](https://www.reddit.com/r/Kotlin/comments/f5sjjh/name_mangling_in_kotlin/)
- url: https://medium.com/@cortinico/name-mangling-in-kotlin-7d0e2a7a173a
---

## [5][How do you call this extension?](https://www.reddit.com/r/Kotlin/comments/f63tft/how_do_you_call_this_extension/)
- url: https://www.reddit.com/r/Kotlin/comments/f63tft/how_do_you_call_this_extension/
---
Is this even possible? if not, why is it allowed?

https://preview.redd.it/niwyp62emsh41.png?width=292&amp;format=png&amp;auto=webp&amp;s=db40b2e31665f0e81dde47cde1b1b6db87e81781
## [6][Escaping Strings to Valid Kotlin String Literals](https://www.reddit.com/r/Kotlin/comments/f64s3z/escaping_strings_to_valid_kotlin_string_literals/)
- url: https://www.reddit.com/r/Kotlin/comments/f64s3z/escaping_strings_to_valid_kotlin_string_literals/
---
I'm wondering if I've overlooked a function in the base Kotlin libraries, that accepts a string, and produces a valid Kotlin literal string.

Example Input:
```
Hello!
How are you doing today?
I'm doing "great"!
```

Example Output:
```
"Hello!\nHow are you doing today?\nI'm doing \"great\"!"
```

---

I can achieve similarly (functionally close enough) with FasterXML's `ObjectMapper::writeValueAsString`, passing a `kotlin.String`, but if there's something in the base library/libraries that would remove the need for a static `ObjectMapper`, that'd be awesome.
## [7][Looking for feedback on first Kotlin project](https://www.reddit.com/r/Kotlin/comments/f62tvd/looking_for_feedback_on_first_kotlin_project/)
- url: https://www.reddit.com/r/Kotlin/comments/f62tvd/looking_for_feedback_on_first_kotlin_project/
---
I've just started work on my first Kotlin project after taking the Coursera Kotlin course some time ago. I'm mostly coming from a Scala &amp; Rust background right now.

I'd love to get some feedback on my code and if there are things I'm doing wrong. Better to catch some bad behaviors before I get too far into this project!

[https://github.com/andygrove/kotlin-query](https://github.com/andygrove/kotlin-query)

Any feedback is appreciated!
## [8][Is it a best practice to directly use the property to assign value now?](https://www.reddit.com/r/Kotlin/comments/f5p3fg/is_it_a_best_practice_to_directly_use_the/)
- url: https://www.reddit.com/r/Kotlin/comments/f5p3fg/is_it_a_best_practice_to_directly_use_the/
---
For example, View has setVisibility.

In Java, I always use view.setVisibility(...), but Android Studio gives me warning of "Use of setter method instead of property access syntax" in Kotlin, while if I use view.visibility directly, the warning is gone.
## [9][Can't figure out what's wrong with the code (tictactoe cli game)](https://www.reddit.com/r/Kotlin/comments/f5sm3p/cant_figure_out_whats_wrong_with_the_code/)
- url: https://www.reddit.com/r/Kotlin/comments/f5sm3p/cant_figure_out_whats_wrong_with_the_code/
---
Hi, I was trying to make a command line tictactoe that should be played in 2, using  kotlin.

here the source code: [https://codeshare.io/am1w0o](https://codeshare.io/am1w0o) (the source code is editable, and I know that this is kind of messy, but I will fix it later on tomorrow).

The problem occurs in the placePiece function.

Basically I am trying to make it not possible for the users to put an X or an O when the place is already occupied, if that makes sense (I explained a little bit better in the code)

But if I try and run the code, the user can still put pieces wherever they want. What am I doing wrong?
## [10][Kotlin User Group Mumbai](https://www.reddit.com/r/Kotlin/comments/f5r5f1/kotlin_user_group_mumbai/)
- url: https://www.reddit.com/r/Kotlin/comments/f5r5f1/kotlin_user_group_mumbai/
---
Hello Everyone if you are an Android Developer living in Mumbai,India we have a community called Kotlin User Group Mumbai so if you are interested and want Explore the world of Kotlin join us here

 [https://www.meetup.com/Kotlin-User-Group-Mumbai/](https://www.meetup.com/Kotlin-User-Group-Mumbai/)
