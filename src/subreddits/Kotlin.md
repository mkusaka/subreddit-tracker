# Kotlin
## [1][Kotlin GDNative](https://www.reddit.com/r/Kotlin/comments/f717nb/kotlin_gdnative/)
- url: https://www.reddit.com/r/godot/comments/f07564/kotlin_gdnative/?utm_source=share&amp;utm_medium=web2x
---

## [2][Testing my apps](https://www.reddit.com/r/Kotlin/comments/f7ak7c/testing_my_apps/)
- url: https://www.reddit.com/r/Kotlin/comments/f7ak7c/testing_my_apps/
---
Hi.Do i need a high end android phone to test my mobile apps?Or any minimum android version?
## [3][Encrypted shared preferences](https://www.reddit.com/r/Kotlin/comments/f6xzub/encrypted_shared_preferences/)
- url: https://www.rockandnull.com/encrypted-shared-preferences/
---

## [4][Kotlin class cant reach Java Class fields](https://www.reddit.com/r/Kotlin/comments/f6t1h2/kotlin_class_cant_reach_java_class_fields/)
- url: https://www.reddit.com/r/Kotlin/comments/f6t1h2/kotlin_class_cant_reach_java_class_fields/
---
Im making a Spring Project basically in Java but for some classes I want to use Kotlin practice. I have a Pojo Java Class named Item with 2 private fields(email and URL) and with the @Data annotation so it should have getters and setters.
In my Kotlin class I want to use those fields and when i write item.email as I should in Kotlin the IntelliJ lets me but on compile it says that those are private and I cant reach them.

I had searched for this problem but didnt found anything usefull.
## [5][Java lazy](https://www.reddit.com/r/Kotlin/comments/f6ynz9/java_lazy/)
- url: https://www.reddit.com/r/Kotlin/comments/f6ynz9/java_lazy/
---
Would you use the class below, such as

    private Lazy&lt;TextView&gt; text = Lazy.of(() -&gt; findViewById(R.id.text));

?


    public class Lazy&lt;Data&gt; {
        public static &lt;Data&gt; Lazy&lt;Data&gt; of(Provider&lt;Data&gt; provider) {
            return new Lazy&lt;&gt;(provider);
        }

        private final Provider&lt;Data&gt; provider;
        private Data field;
        private boolean instantiated;

        private Lazy(Provider&lt;Data&gt; provider) {
            this.provider = provider;
        }

        public Data get() {
            if (!instantiated) {
                field = provider.provide();
                instantiated = true;
            }
            return field;
        }

        public void reset(Data newData) {
            field = newData;
        }

        public interface Provider&lt;Data&gt; {
            Data provide();
        }
    }
## [6][Kotlin: Fun with “in”](https://www.reddit.com/r/Kotlin/comments/f6aq8r/kotlin_fun_with_in/)
- url: https://proandroiddev.com/kotlin-fun-with-in-8a425704b635
---

## [7][swagger-gradle-codegen v1.4.0 just got released with support for moshi-codegen and a lot of bugfixes/improvements.](https://www.reddit.com/r/Kotlin/comments/f6h990/swaggergradlecodegen_v140_just_got_released_with/)
- url: https://github.com/Yelp/swagger-gradle-codegen/releases/tag/1.4.0
---

## [8][LiveStream-kt - Library to emit &amp; observe data from any module of the application.](https://www.reddit.com/r/Kotlin/comments/f6lx2c/livestreamkt_library_to_emit_observe_data_from/)
- url: https://github.com/PatilShreyas/LiveStream-kt
---

## [9][Can't figure out how to put this into a function to make the code D.R.Y.](https://www.reddit.com/r/Kotlin/comments/f6h24j/cant_figure_out_how_to_put_this_into_a_function/)
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
## [10][Can I download Kotlin on a Chromebook?](https://www.reddit.com/r/Kotlin/comments/f6fl69/can_i_download_kotlin_on_a_chromebook/)
- url: https://www.reddit.com/r/Kotlin/comments/f6fl69/can_i_download_kotlin_on_a_chromebook/
---

