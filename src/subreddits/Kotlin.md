# Kotlin
## [1][Published my first app to the play store today 😁](https://www.reddit.com/r/Kotlin/comments/haoxom/published_my_first_app_to_the_play_store_today/)
- url: https://www.reddit.com/r/Kotlin/comments/haoxom/published_my_first_app_to_the_play_store_today/
---

I started learning kotlin for Android app development a few weeks ago and I’m loving it


Published my first android app to the play store using Kotlin. Yes I know it’s a ***VERY basic app*** However I’m learning new stuff everyday.

The ***primary focus*** of making this simple app was just to understand the fundamentals of kotlin,  how to use android studio and how to implement ads using admob.

Check it out if you want to see what a terrible app this is lol. The main thing is that I’m learning anyway.

https://play.google.com/store/apps/details?id=com.herd.whattodo
## [2][Kotlin Cheat Sheet - a basic reference for beginner and advanced](https://www.reddit.com/r/Kotlin/comments/hal4zf/kotlin_cheat_sheet_a_basic_reference_for_beginner/)
- url: https://simplecheatsheet.com/tag/kotlin-cheat-sheet/
---

## [3][Elide - JSON:API or GraphQL web service starting from a JPA annotated data model (by Yahoo!)](https://www.reddit.com/r/Kotlin/comments/haq2ht/elide_jsonapi_or_graphql_web_service_starting/)
- url: https://blog.graphqleditor.com/elide-opinionated-apis/
---

## [4][[Doubt] What is the function of init code block in following program ?](https://www.reddit.com/r/Kotlin/comments/hannfg/doubt_what_is_the_function_of_init_code_block_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hannfg/doubt_what_is_the_function_of_init_code_block_in/
---
    class Food(_name: String) {
        var name = _name
            get() = "I'm eating $field"
            set(value) {
                field = value.toLowerCase().reversed().capitalize()
            }
    /*    init {
            name=_name
        }*/
    }
    
    fun main(args: Array&lt;String&gt;) {
        val food = Food("Banana")
        println(food.name)
    }

In above program why commenting out init code block causing program to fail  to reverse  print banana ?  why I need to  initialize name=\_name in 'init code block' even though I did it at start in Class Food?

    ("I'm eating Banana"-&gt;"I'm eating Ananab")

&amp;#x200B;
## [5][Routing in Ktor](https://www.reddit.com/r/Kotlin/comments/ha8eds/routing_in_ktor/)
- url: https://hadihariri.com/2020/04/02/Routing-in-Ktor/
---

## [6][Any examples/engineering blog posts of companies using Kotlin for server side? Better, if paired with spring boot.](https://www.reddit.com/r/Kotlin/comments/haeghi/any_examplesengineering_blog_posts_of_companies/)
- url: https://www.reddit.com/r/Kotlin/comments/haeghi/any_examplesengineering_blog_posts_of_companies/
---
thanks!
## [7][Building a Reactive Oauth Client App with SpringBoot and Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/ha3u2m/building_a_reactive_oauth_client_app_with/)
- url: https://www.shiveenp.com/posts/spring-boot-reactive-oauth-client-with-coroutines/
---

## [8][Where are the answers to the online Kotlin Bootcamp homework questions?](https://www.reddit.com/r/Kotlin/comments/haj1k5/where_are_the_answers_to_the_online_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/haj1k5/where_are_the_answers_to_the_online_kotlin/
---
Does anyone know where I could find the answers to these:
https://codelabs.developers.google.com/codelabs/kotlin-bootcamp-introduction/#7

(The homework questions to the Kotlin Bootcamp online)

Most can be found within the tutorials but later on some questions can't necessarily be found in the exact text.
## [9][Inheritance, Overriding using Open Keyword in Kotlin](https://www.reddit.com/r/Kotlin/comments/hainpg/inheritance_overriding_using_open_keyword_in/)
- url: https://www.youtube.com/watch?v=s2UwnlAxstA&amp;t=6s
---

## [10][Connect 4](https://www.reddit.com/r/Kotlin/comments/ha8sbz/connect_4/)
- url: https://www.reddit.com/r/Kotlin/comments/ha8sbz/connect_4/
---
Hello,

I am working on a program that allows the user to play a game of Connect 4 within the IDE command console.

Currently, I have my playerVsPlayer version working reasonably. However, I am having trouble checking for a win in any given direction. I am also having an issue implementing the  computer for a playerVsComputer or computerVsComputer version. I tried copy pasting my pvp code but I am unsure how to put the computer into the loop.

Thanks in advance.

Here is my code:

    var gameBoard = Array( 8 ) { Array( 8 ) { " . " } }
    var player = "X"
    var turnCounter = 1
    var winner = false
    var computer1 = ( 1..8 ).random()
    var computer2 = ( 1..8 ).random()
    
    fun main() {
     println()
     println( "Welcome to the game of Connect Four!" )
     var menuItems = arrayOf(
     "Player vs. Player",
     "Player vs. Computer",
     "Computer vs. Computer",
     "Quit"
     )
     var quitOption = menuItems.size
     var userChoice = 0
     while( userChoice != quitOption ) {
     println()
            userChoice = menuOptions( menuItems, "\nPlease enter your selection: " )
     when( userChoice ) {
     1 -&gt; playerVsPlayer()
     2 -&gt; playerVsComputer()
     3 -&gt; computerVsComputer()
     else -&gt; {
     if( userChoice != quitOption ) {
     println()
     println( "ERROR: Please select a valid menu option." )
                    }
                }
            }
        }
     println()
     println( "Thank you for playing!" )
    }
    
    fun menuOptions(items: Array&lt;String&gt;, prompt: String): Int {
     for( ( index, item ) in items.withIndex() ) {
     println( "${index + 1}. $item" )
        }
     print( prompt )
     return readLine()!!.toInt()
    }
    
    fun playerVsPlayer() {
     while(!winner) {
     var validMove: Boolean
     var play: Int
     do {
     printGameBoard(gameBoard)
     print("Player $player, please choose a column: ")
                play = readLine()!!.toInt() - 1
     printGameBoard(gameBoard)
                validMove = validateMove(play, gameBoard)
            } while( validMove == false )
     for (row in gameBoard.indices.reversed()) {
     if (gameBoard[row][play] == " . ") {
     gameBoard[row][play] = " ${player} "
     break
     }
            }
     player = if (player == "X") {
     "O"
     } else {
     "X"
     }
     turnCounter++
        }
     printGameBoard(gameBoard)
     if (winner) {
     if (player == "X") {
     winner = true
     println("X won")
            } else {
     winner = true
     println("O won")
            }
        } else {
     println("Tie game")
        }
    }
    
    fun playerVsComputer() {
     while(!winner) {
     var validMove: Boolean
     var play: Int
     do {
     printGameBoard(gameBoard)
     print("Player $player, please choose a column: ")
                play = readLine()!!.toInt() - 1
     printGameBoard(gameBoard)
                validMove = validateMove(play, gameBoard)
            } while( validMove == false )
     for (row in gameBoard.indices.reversed()) {
     if (gameBoard[row][play] == " . ") {
     gameBoard[row][play] = " ${player} "
     break
     }
            }
     player = if (player == "X") {
     "O"
     } else {
     "X"
     }
     turnCounter++
        }
     printGameBoard(gameBoard)
     if (winner) {
     if (player == "X") {
     winner = true
     println("X won")
            } else {
     winner = true
     println("O won")
            }
        } else {
     println("Tie game")
        }
    }
    
    fun computerVsComputer() {
     while(!winner) {
     var validMove: Boolean
     var play: Int
     do {
     printGameBoard(gameBoard)
     print("Player $player, please choose a column: ")
                play = readLine()!!.toInt() - 1
     printGameBoard(gameBoard)
                validMove = validateMove(play, gameBoard)
            } while( validMove == false )
     for (row in gameBoard.indices.reversed()) {
     if (gameBoard[row][play] == " . ") {
     gameBoard[row][play] = " ${player} "
     break
     }
            }
     player = if (player == "X") {
     "O"
     } else {
     "X"
     }
     turnCounter++
        }
     printGameBoard(gameBoard)
     if (winner) {
     if (player == "X") {
     winner = true
     println("X won")
            } else {
     winner = true
     println("O won")
            }
        } else {
     println("Tie game")
        }
    }
    
    fun printGameBoard(gameBoard: Array&lt;Array&lt;String&gt;&gt;) {
     println()
     println( "          Connect Four!          " )
     println( "╔===╦===╦===╦===╦===╦===╦===╦===╗" )
     println( "║ 1 ║ 2 ║ 3 ║ 4 ║ 5 ║ 6 ║ 7 ║ 8 ║" )
     println( "╠===╬===╬===╬===╬===╬===╬===╬===╣" )
     for( row in 0 until gameBoard.size ) {
     print( "║" )
     for( col in 0 until gameBoard[row].size ) {
     print( gameBoard[row][col] )
     print( "║" )
            }
     println()
     println( "╠===╬===╬===╬===╬===╬===╬===╬===╣" )
        }
     println( "║ 1 ║ 2 ║ 3 ║ 4 ║ 5 ║ 6 ║ 7 ║ 8 ║" )
     println( "╚===╩===╩===╩===╩===╩===╩===╩===╝" )
     println()
    }
    
    fun validateMove(col: Int, gameBoard: Array&lt;Array&lt;String&gt;&gt;): Boolean {
     if( col &lt; 0 || col &gt; gameBoard[0].size ) {
     return false
     }
     return if( gameBoard[0][col] != " . " ) {
     false
     } else true
    }
    
    fun fourInARow(player: String, gameBoard: Array&lt;Array&lt;String&gt;&gt;): Boolean {
     //check for 4 across
     for (row in gameBoard.indices) {
     for (col in 0 until gameBoard[0].size - 4) {
     if (gameBoard[row][col] == player &amp;&amp;
                        gameBoard[row][col + 1] == player &amp;&amp;
                        gameBoard[row][col + 2] == player &amp;&amp;
                        gameBoard[row][col + 3] == player) {
     return true
     }
            }
        }
     //check for 4 up and down
     for (row in 0 until gameBoard.size - 4) {
     for (col in 0 until gameBoard[0].size) {
     if (gameBoard[row][col] == player &amp;&amp;
                        gameBoard[row + 1][col] == player &amp;&amp;
                        gameBoard[row + 2][col] == player &amp;&amp;
                        gameBoard[row + 3][col] == player) {
     return true
     }
            }
        }
     //check upward diagonal
     for (row in 2 until gameBoard.size) {
     for (col in 0 until gameBoard[0].size - 4) {
     if (gameBoard[row][col] == player &amp;&amp;
                        gameBoard[row - 1][col + 1] == player &amp;&amp;
                        gameBoard[row - 2][col + 2] == player &amp;&amp;
                        gameBoard[row - 3][col + 3] == player) {
     return true
     }
            }
        }
     //check downward diagonal
     for (row in 0 until gameBoard.size - 4) {
     for (col in 0 until gameBoard[0].size - 4) {
     if (gameBoard[row][col] == player &amp;&amp;
                        gameBoard[row + 1][col + 1] == player &amp;&amp;
                        gameBoard[row + 2][col + 2] == player &amp;&amp;
                        gameBoard[row + 3][col + 3] == player) {
     return true
     }
            }
        }
     return false
    }
    

&amp;#x200B;
