# Kotlin
## [1][Should I use 'and' and 'or' instead of &amp;&amp; and ||?](https://www.reddit.com/r/Kotlin/comments/f55bup/should_i_use_and_and_or_instead_of_and/)
- url: https://www.reddit.com/r/Kotlin/comments/f55bup/should_i_use_and_and_or_instead_of_and/
---

## [2][understanding coroutines / channels](https://www.reddit.com/r/Kotlin/comments/f4sakk/understanding_coroutines_channels/)
- url: https://www.reddit.com/r/Kotlin/comments/f4sakk/understanding_coroutines_channels/
---
Hi,

I have a hard time understanding coroutines and channels.

I have the following szenario:

1. I want to read some URLs from a database  
2. I want to download the content from these URLs  
3. I want to store the content in a database  

So I thought kotlin channels would be the right way to implement this pipeline. Read the URLs from the database, download them in parallel and store them as soon as they are downloaded.

However, I have the problem that only two downloads are started at a time.

So I wrote this short example to understand what's happening:


    fun main(args: Array&lt;String&gt;) {
        val context = newFixedThreadPoolContext(5, "TEST")
        runBlocking(context) {
            produce {  // 1. read from database
                for (i in 1..5) {
                    send(i)
                }
            }
            .map {  // 2. download content
                async {
                    println("$it downloading")
                    delay(5000)
                    it
                }
            }
            .consumeEach{  // 3. store to database
                val content = it.await()
                println("$content consumed")
            }
        }
    }


with the Output

    1 downloading
    2 downloading
    1 consumed
    2 consumed
    4 downloading
    3 downloading
    3 consumed
    4 consumed
    5 downloading
    5 consumed

So why are there two downloads running at a time. I thought with the context of 5 Threads it would start all 5 downloads.
## [3][Stream processing for computing approximations](https://www.reddit.com/r/Kotlin/comments/f4udv4/stream_processing_for_computing_approximations/)
- url: https://blog.frankel.ch/stream-processing/2/
---

## [4][Kotlin scoped functions](https://www.reddit.com/r/Kotlin/comments/f4ec6n/kotlin_scoped_functions/)
- url: https://www.rockandnull.com/kotlin-let/
---

## [5][Hyderabad Kotlin User Group (Hyderabad, India)](https://www.reddit.com/r/Kotlin/comments/f4jncs/hyderabad_kotlin_user_group_hyderabad_india/)
- url: https://www.reddit.com/r/Kotlin/comments/f4jncs/hyderabad_kotlin_user_group_hyderabad_india/
---
Hi,

If you are living in are around  Hyderabad or Secendrabad, Telanganam India. Feel free to join us in our Kotlin User Group. We are just started so be connected to monthly meetings and other events. 

URL https://www.meetup.com/Hyderabad-Kotlin-User-Group/

Feel free to share with your colleagues and friends
## [6][Hyderabad Kotlin User Group (Hyderabad, India)](https://www.reddit.com/r/Kotlin/comments/f4jnvq/hyderabad_kotlin_user_group_hyderabad_india/)
- url: https://www.meetup.com/Hyderabad-Kotlin-User-Group/
---

## [7][Setter annotation inheritance](https://www.reddit.com/r/Kotlin/comments/f4i72w/setter_annotation_inheritance/)
- url: https://www.reddit.com/r/Kotlin/comments/f4i72w/setter_annotation_inheritance/
---
Hi

I have these classes.

    @Target(AnnotationTarget.FUNCTION,AnnotationTarget.VALUE_PARAMETER,AnnotationTarget.CLASS,AnnotationTarget.FIELD,AnnotationTarget.PROPERTY_SETTER) u/Inherited annotation class Name(val name:String,val valueName:String=""){}
    
    @Target(AnnotationTarget.CLASS)
    @Inherited
    annotation class Context(val name:String){}
    
    @Target(AnnotationTarget.FUNCTION,AnnotationTarget.PROPERTY_SETTER)
    @Inherited
    annotation class ArgumentSetter {
    }
    
    @Name("test")
    @Context("test")
    internal interface ParentWithAnnotations {
        var test: String
            @Name("field", valueName = "value") @ArgumentSetter set
    }
    
    internal class AnnotatedParent : ParentWithAnnotations {
        override var test: String = ""
    }
    
    assertEquals(1,data::test::setter.annotations.stream().peek({println(it)}).filter { it == ArgumentSetter::class.java }.count())

I want to make it so I can read the setter annotations of the interface property from the implementation class. Is this possible?
## [8][Is there a way to make a data class open?](https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/)
- url: https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/
---
I'm trying to use the ModelMapper library with data classes. To build an explicit type map it needs to proxy my class, and data classes cannot be extended and therefore cannot be proxied. Is there a way around this? Some compiler setting? I've tried all-open but I can't seem to get it to work.

Thanks.
## [9][Why are parameters passed into a function always val?](https://www.reddit.com/r/Kotlin/comments/f4fdvw/why_are_parameters_passed_into_a_function_always/)
- url: https://www.reddit.com/r/Kotlin/comments/f4fdvw/why_are_parameters_passed_into_a_function_always/
---
Title. And how can I make a parameter mutable?

e.g. 

&amp;#x200B;

    fun testFunction(a: String)
    {
        a = a + "yelp"
    }
    
    fun main() {
        var a: String = "Help, "
        testFunction(a)
    }
## [10][how do you even print a 2d array in kotlin??](https://www.reddit.com/r/Kotlin/comments/f4avak/how_do_you_even_print_a_2d_array_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/f4avak/how_do_you_even_print_a_2d_array_in_kotlin/
---
Hi, I am new to kotlin, I have been using Java for a while, and I still have to get used to Kotlin.

The main thing is that Kotlin's for loop is quite different (I'm pretty sure that's similar to python's ones).

&amp;#x200B;

If for example I have to print this 2d array:

&amp;#x200B;

    var 2dArray = arrayOf(
        arrayOf(1, 2, 3, 4, 5),
        arrayOf(6, 7, 8, 9, 10)
    )

how do you do that in kotlin?

&amp;#x200B;

&amp;#x200B;

I tried with:

    for (i in 2dArray) {
            for (j in 2dArray) {
                print("$j")
            }
        }

&amp;#x200B;
