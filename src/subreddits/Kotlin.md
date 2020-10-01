# Kotlin
## [1][Kotlinx Serialization: How to circumvent reified typeargs for deserialization?](https://www.reddit.com/r/Kotlin/comments/j37tyn/kotlinx_serialization_how_to_circumvent_reified/)
- url: https://www.reddit.com/r/Kotlin/comments/j37tyn/kotlinx_serialization_how_to_circumvent_reified/
---
Actually, the main problem is still that there are no reified typeargs for classes in Kotlin (see [my other reddit post](https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/?utm_source=share&amp;utm_medium=web2x&amp;context=3)). But here is why this bothers me in this specific case:

Suppose you have a wrapper class `Wrapper` that takes in a string `content` and a class\* `type` and can output an object of class `type` retrieved by parsing `content` as JSON by demand by calling the function `getObj()`:

    class Wrapper&lt;T&gt;(private val content: String, /*private val type: KClass&lt;*&gt;*/) {
        fun getObj(): T {
            // ?
        }
    }

And I want to use kotlinx.serialization. Now, you might have noticed how I put an asterisk after "class" before. Here's the reason: Yes, `Wrapper` has to take the target class in _some_ way, but how? Should it be just the typearg (won't work because type erausre) or a `KClass` reference (won't work because I need a reified typearg)? 

The thing is that as far as I know, the only way to decode a generic JSON to a serializable target class is to use `Json.decodeFromString&lt;T&gt;(content)`, where `T` is the target type and `content` is the JSON string. Now, `T` is defined to be reified (so that the type can be processed at runtime) and can only be filled with another reified typearg or an actual class reference. I can't use another reified typearg because I am in the context of a class and a class cannot have reified typeargs. I can also not use an actual class reference because the user of the class should be able to construct it with different targets, e.g. they decide what the target is, not me.

So, how do I do this with kotlinx.serialization? Is it even possible?
## [2][Autocompletion not working with JPA repository in intellij when using kotlin](https://www.reddit.com/r/Kotlin/comments/j33pf0/autocompletion_not_working_with_jpa_repository_in/)
- url: https://www.reddit.com/r/Kotlin/comments/j33pf0/autocompletion_not_working_with_jpa_repository_in/
---
Please help
## [3][Chucker (OkHTTP Debugger for Android) v3.3.0 is out 🎉 - This version comes with minor features, improvement and bugfixes.](https://www.reddit.com/r/Kotlin/comments/j37ep4/chucker_okhttp_debugger_for_android_v330_is_out/)
- url: https://github.com/ChuckerTeam/chucker/releases/tag/3.3.0
---

## [4][Coding Style](https://www.reddit.com/r/Kotlin/comments/j2zvwk/coding_style/)
- url: https://www.reddit.com/r/Kotlin/comments/j2zvwk/coding_style/
---
I have been working on a few Android apps in Kotlin to place in my resume and I have deviated from the colon spacing rule listed [here](https://kotlinlang.org/docs/reference/coding-conventions.html), as well as spacing in general. 

An example:

    val currencySymbolKey  : String  = sp[Constants.KEY_CURRENCY_SYMBOL , "dollar"]!!
    val dateFormatKey      : String  = sp[Constants.KEY_DATE_FORMAT     , "0"     ]!!
    val decimalSymbolKey   : String  = sp[Constants.KEY_DECIMAL_SYMBOL  , "period"]!!
    val thousandsSymbolKey : String  = sp[Constants.KEY_THOUSANDS_SYMBOL, "comma" ]!!
    val decimalPlaces      : Boolean = sp[Constants.KEY_DECIMAL_PLACES  , true    ]!!
    val symbolSide         : Boolean = sp[Constants.KEY_SYMBOL_SIDE     , true    ]!!
Another:

    val currencySymbol  : String = getCurrencySymbol (currencySymbolKey )
    val dateFormat      : Int    = getDateFormat     (dateFormatKey     )
    val decimalSymbol   : Char   = getSeparatorSymbol(decimalSymbolKey  )
    val thousandsSymbol : Char   = getSeparatorSymbol(thousandsSymbolKey)
The rule states that colon should have no spaces when declaring a variable and its type, but I like the way it looks when they are lined up as shown above. I find it easier to read compared to:

    val currencySymbolKey: String = sp[Constants.KEY_CURRENCY_SYMBOL , "dollar"]!!
    val dateFormatKey: String = sp[Constants.KEY_DATE_FORMAT, "0"]!!
    val decimalSymbolKey: String = sp[Constants.KEY_DECIMAL_SYMBOL, "period"]!!
    val thousandsSymbolKey: String = sp[Constants.KEY_THOUSANDS_SYMBOL, "comma"]!!
    val decimalPlaces: Boolean = sp[Constants.KEY_DECIMAL_PLACES, true]!!
    val symbolSide: Boolean = sp[Constants.KEY_SYMBOL_SIDE, true]!!

My question is whether or not this is good coding practice. Should I go back and remove all unnecessary white space?  The only time I do this is when defining multiple variables at once as shown above.
## [5][Kotlin exercises](https://www.reddit.com/r/Kotlin/comments/j2k1oc/kotlin_exercises/)
- url: https://www.reddit.com/r/Kotlin/comments/j2k1oc/kotlin_exercises/
---
Hey, does anyone know something similar to [scalatron](https://scalatron.github.io/) but for kotlin or something similar [Codin game clash of code private server](https://www.codingame.com/multiplayer) just with bot writing? Of course any exercises that have some kind of competitiveness in kotlin would be extremely useful.
## [6][Kotlin Loops/Regular Expressions](https://www.reddit.com/r/Kotlin/comments/j2zqrm/kotlin_loopsregular_expressions/)
- url: https://www.reddit.com/r/Kotlin/comments/j2zqrm/kotlin_loopsregular_expressions/
---
 

Hello,

I am working on writing a program in Kotlin that utilizes regular expressions to read data from a file to give class info.

The only issues I seem to be having are with two of the loops that are within my program.

For example, I would like the program to run again and again until the user inputs exit/EXIT when prompted to enter the department. I thought I would do this with a do...while loop, however, whenever I type exit/EXIT the program continues to run.

As for my other problem loop, I would like it to run until the patternTwo it is searching for stops being recognized or pattern is. For example, there might be more than 1 instructor who would be listed on the next line. I know that the loop must be nested within the first for loop but I just cannot seem to get it to work without it being endless and so I have it commented out.

My relevant code is below and I appreciate all of the help ahead of time.

`do {`               

`val file = "src/Enrollment.txt"`       

`var lines = File(file).readLines()`               

`println()`                

`print("Enter Department: ")`              

`var department = readLine()!!.toUpperCase()`                

`print("Enter Class Number: ")`              

`var classNumber = readLine()!!`                

`val pattern = """\A[\d\s]*?$department+?\s*?$classNumber+?""".toRegex()`               

`val patternTwo = """\A\s+.*""".toRegex()`                

`for (i in 0..(lines.size - 1)) {`                       

`if (pattern.containsMatchIn(lines[i])) {`                               

`println(lines[i])`                               

`if (patternTwo.containsMatchIn(lines[i + 1])) {`                                         
`println(lines[i + 1])`                               

`}`                                  

`/*for (i in 0..(lines.size - 1)) {                      }*/`                       

`}`               

`}`       

`} while ( department != "EXIT" )`   

&amp;#x200B;

`sample output:`  

`Enter Department: CIS` 

`Enter Class Number: 282`      

`2007 CIS   282  AN        HD  PROG PRN I    5.0  5.0 Hybrid Class 001 1116  F          ARR    ARR &lt;- line 1 matches pattern`                                                            

`Jones D      001 1116  MTWTh      10:30A 11:30A    24     15     9  5.0 &lt;- line 2 matches patternTwo`
## [7][Testing MVI View Models with Kotlin on Android](https://www.reddit.com/r/Kotlin/comments/j2opwl/testing_mvi_view_models_with_kotlin_on_android/)
- url: https://quickbirdstudios.com/blog/android-mvi-testing-view-model-kotlin/
---

## [8][Weekly curated newsletter for Mobile developer with loads of Kotlin and Android content.](https://www.reddit.com/r/Kotlin/comments/j2wt6g/weekly_curated_newsletter_for_mobile_developer/)
- url: https://mobiledeveloperscafe.com
---

## [9][How to change text color of a tab in TornadoFX](https://www.reddit.com/r/Kotlin/comments/j2kinu/how_to_change_text_color_of_a_tab_in_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/j2kinu/how_to_change_text_color_of_a_tab_in_tornadofx/
---
Hi guys,

Yesterday I made a post about whether or not I should use TornadoFX. The result was that I decided to give it a shot. While fooling around with it yesterday, I noticed that I am able to change the background color or the background radius of a tab in a tab pane, but I wasn't able to figure out how to change the text color. Normally, it would be `textFill`. Does anyone know how to do that?
## [10][Effective Class Delegation - zsmb.co](https://www.reddit.com/r/Kotlin/comments/j277pm/effective_class_delegation_zsmbco/)
- url: https://zsmb.co/effective-class-delegation/
---

