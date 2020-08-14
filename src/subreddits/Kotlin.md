# Kotlin
## [1][I used the ExperimentalStdlibApi to annotate the ExperimentalStdlibApi](https://www.reddit.com/r/Kotlin/comments/i9h8i6/i_used_the_experimentalstdlibapi_to_annotate_the/)
- url: https://i.redd.it/kjili1lo1xg51.jpg
---

## [2][KotlinX multiplatform date/time library](https://www.reddit.com/r/Kotlin/comments/i9l6fe/kotlinx_multiplatform_datetime_library/)
- url: https://github.com/Kotlin/kotlinx-datetime
---

## [3][kotlinx-datetime 0.1 has been published](https://www.reddit.com/r/Kotlin/comments/i94r6i/kotlinxdatetime_01_has_been_published/)
- url: https://discuss.kotlinlang.org/t/kotlinx-datetime-0-1-has-been-published/18766
---

## [4][I am here to Praise Android Studio and KOTLIN. Story inside](https://www.reddit.com/r/Kotlin/comments/i9iyuf/i_am_here_to_praise_android_studio_and_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/i9iyuf/i_am_here_to_praise_android_studio_and_kotlin/
---
So I decided I needed more control over some home automation things and that using an Android would be my best bet.  I am not a professional by any means, I only say this to give warning tht im google taught.

I got into coding with Cheat Engine by learning ASM then learning CE's included LUA.  Which got me into VBA a dead language IK, but it was very LUA like so it was easy to catch on.  Then I moved into C# which I HATE . Hate Arduino IDE. Hate the syntax, the overuse of defines, terrible arrays and SEMICOLONS. A uint8\_t cost me 3 nights of troubleshooting before. I dont like C# in Visual Studio either.

Back to this Android App im writing. So I try to learn Xamarin. NOPE to that. idk who had the bright idea of breaking Visual Studio's interface with the Xamarin extension but that team should be fired. The Official Xamarin youtube channel actually steers you towards a PAID extension for GUI design.

So I reluctantly downloaded Android Studio and decided to learn Kotlin over Java. Couple reasons why Kotlin over Java 1. the tutorial I was watching was in Kotlin 2. MOST important Java uses this "  ;  " , ugh semicolons. Them things suck they a relic of old language compilers that couldnt understand lines or defines properly.

&amp;#x200B;

Anyway Ive been learning Kotlin for a week or so and OMG this languange and the way Android Studio handles it is so intuitive they deserve reward. Need an Import for [file.IO](https://file.IO) its write/paste the code, highlight the RED, click Import. Thts incredible!  I love how forced returns work eg.(return@setOnClickListener). The loop setups are done right with *downTo* and *step*. The greyed out indicator of var's sayin Int,String,etc. The Colorcoding with Android Studio is the best ive seen. But most importantly Kotlin has NO semicolons.
## [5][Use Kotlin USE to shorten file reading](https://www.reddit.com/r/Kotlin/comments/i9kabg/use_kotlin_use_to_shorten_file_reading/)
- url: https://medium.com/mobile-app-development-publication/use-kotlin-use-to-shorten-file-reading-cc57565d2e22?source=friends_link&amp;sk=a03daae414adb0e4f62c8d3474dee0c2
---

## [6][How to implement Dagger2 in an Android application](https://www.reddit.com/r/Kotlin/comments/i9j4h7/how_to_implement_dagger2_in_an_android_application/)
- url: https://youtu.be/bV3jf8VpbSY
---

## [7][Reflection &amp; Casting from generic to a class - dynamically?](https://www.reddit.com/r/Kotlin/comments/i9fqlq/reflection_casting_from_generic_to_a_class/)
- url: https://www.reddit.com/r/Kotlin/comments/i9fqlq/reflection_casting_from_generic_to_a_class/
---
Let's say I have a class Example, that grabs a list of rows from a database, for a table elephant. I can pass the columns for each row into a call to the constructor:
    class Example(kclass: KClass&lt;*&gt;) {
        // KotlinCasing -&gt; sql_casing:
        private val constructor: KFunction&lt;*&gt;
        private val klass:KClass&lt;*&gt; = kclass
        init {
            constructor = klass.constructors.first()
        }
        fun map(rows: ResultSet): List&lt;Any&gt; {
            // Map a result set into instances of the data class
            return rows.map {
                val args = parameters.zip(it).toMap()
                constructor.callBy(args)!!
            }
        }
    }

This works great for a data class Elephant.

Problem is - I'd like to be able to pass in Elephant::class (perhaps as elephantClass) so I could do:
        
    class Example(kclass: KClass&lt;*&gt;) {
        // KotlinCasing -&gt; sql_casing:
        private val constructor: KFunction&lt;*&gt;
        private val klass:KClass&lt;*&gt; = kclass
        init {
            constructor = klass.constructors.first()
        }
        fun map(rows: ResultSet): List&lt;Any&gt; {
            // Map a result set into instances of the data class
            return rows.map {
                val args = parameters.zip(it).toMap()
                constructor.callBy(args)!! as kclass
            }
        }
    }

That won't work. I can call Example.map() in app code, and cast those results  directly to Elephant:

    val x = Example(Elephant::class)
    val y = x.map(results).first() as Elephant

y will be an Elephant instance.

I can also redefine "map" as follows:

    fun &lt;T&gt; map(rows: ResultSet): List&lt;T&gt; {
        // Map a result set into instances of the data class
        return rows.map {
            val args = parameters.zip(it).toMap()
            constructor.callBy(args)!! as T
        }
    }

It just means I need to pass the class in twice, in the call to map(), and while initializing Example. Is there a way to just pass in Example in some fashion while initializing my class? I played with reified - but it seems to really make that work well you need an inline function.

Thanks for the help!
## [8][Best &amp; Most Up-To-Date Kotlin Book](https://www.reddit.com/r/Kotlin/comments/i97b7b/best_most_uptodate_kotlin_book/)
- url: https://www.reddit.com/r/Kotlin/comments/i97b7b/best_most_uptodate_kotlin_book/
---
I have several years of experience with Kotlin, mostly on Android, but also some JVM and Native. I'm totally comfortable with the language and I feel like I already know many ins-and-outs, along with the right ways to do things (scoped functions, sequences vs. lists, coroutines, extension functions, sealed classes, etc).

I love learning, so I want to find a good book to describe what's idiomatic and also best practices, in order to augment my skills. Just doing a cursory search on Amazon, there's a lot to look through. I only want to buy one book, especially since they are kind of expensive. Which is the the best, most-comprehensive, and most up-to-date one, that's not at an introductory level, and can also go deep into implementation (JVM details, vs. Java, Native/Multiplatform, etc)?

I've read Jake's blog, which is sometimes fascinating to me how in-depth it gets, albeit somewhat overwhelming at times. I guess what I want is that in book-form, but structured.
## [9][Professional Work: Libraries Licensed as MPL-2.0](https://www.reddit.com/r/Kotlin/comments/i9bgqw/professional_work_libraries_licensed_as_mpl20/)
- url: https://www.reddit.com/r/Kotlin/comments/i9bgqw/professional_work_libraries_licensed_as_mpl20/
---
I'm trying to gauge the community's reaction for a scenario. Let's say you're working on a professional project (not personal) and you see a library you want to check out. But you notice that it's licensed as MPL-2.0, and not Apache 2 or MIT. You're only looking to include the library and not modify it.

What would your reaction be?

[View Poll](https://www.reddit.com/poll/i9bgqw)
## [10][How to preserve text formatting of string in kotlin](https://www.reddit.com/r/Kotlin/comments/i9bcg2/how_to_preserve_text_formatting_of_string_in/)
- url: /r/learnprogramming/comments/i9bc6a/how_to_preserve_text_formatting_of_string_in/
---

