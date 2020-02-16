# Kotlin
## [1][Kotlin scoped functions](https://www.reddit.com/r/Kotlin/comments/f4ec6n/kotlin_scoped_functions/)
- url: https://www.rockandnull.com/kotlin-let/
---

## [2][Hyderabad Kotlin User Group (Hyderabad, India)](https://www.reddit.com/r/Kotlin/comments/f4jncs/hyderabad_kotlin_user_group_hyderabad_india/)
- url: https://www.reddit.com/r/Kotlin/comments/f4jncs/hyderabad_kotlin_user_group_hyderabad_india/
---
Hi,

If you are living in are around  Hyderabad or Secendrabad, Telanganam India. Feel free to join us in our Kotlin User Group. We are just started so be connected to monthly meetings and other events. 

URL https://www.meetup.com/Hyderabad-Kotlin-User-Group/

Feel free to share with your colleagues and friends
## [3][Beginner needs a little help](https://www.reddit.com/r/Kotlin/comments/f4itxi/beginner_needs_a_little_help/)
- url: https://www.reddit.com/r/Kotlin/comments/f4itxi/beginner_needs_a_little_help/
---
I am a beginner trying to learn Kotlin. I have an assignment that I am working on. See below.  I can get the odd numbers but how do I add them up at the end ? 

 1. The program prompts a user to enter a number. 2. The program adds all the odd numbers between 1 and the number the user has entered and displays the result. For instance, if the user enters 5, the program adds all odd numbers between 1 and 5 (i.e., 1, 3, and 5) and displays the result (i.e., 9). 3. The program asks the user whether he/she wants to continue. If the user enters c or C, the program continues and repeats the process again. If the user enters any other key, the program stops 

&amp;#x200B;

import java.util.Scanner  
fun main(args: Array&lt;String&gt;) {  


val userInput = Scanner(System.*\`in\`*)  
 *print*("Please enter a number: ")  
 var num:Int = userInput.nextInt()  
 *println*("Your number is: $num")  


*print*("The sum of your odd numbers is: ")  
 for (i in 1..num-1)  
 if(i%2!=0)*print*("$i ")  
}
## [4][Hyderabad Kotlin User Group (Hyderabad, India)](https://www.reddit.com/r/Kotlin/comments/f4jnvq/hyderabad_kotlin_user_group_hyderabad_india/)
- url: https://www.meetup.com/Hyderabad-Kotlin-User-Group/
---

## [5][Is there a way to make a data class open?](https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/)
- url: https://www.reddit.com/r/Kotlin/comments/f48mnz/is_there_a_way_to_make_a_data_class_open/
---
I'm trying to use the ModelMapper library with data classes. To build an explicit type map it needs to proxy my class, and data classes cannot be extended and therefore cannot be proxied. Is there a way around this? Some compiler setting? I've tried all-open but I can't seem to get it to work.

Thanks.
## [6][Setter annotation inheritance](https://www.reddit.com/r/Kotlin/comments/f4i72w/setter_annotation_inheritance/)
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
## [7][Why are parameters passed into a function always val?](https://www.reddit.com/r/Kotlin/comments/f4fdvw/why_are_parameters_passed_into_a_function_always/)
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
## [8][how do you even print a 2d array in kotlin??](https://www.reddit.com/r/Kotlin/comments/f4avak/how_do_you_even_print_a_2d_array_in_kotlin/)
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
## [9][java.lang.UnsupportedOperationException: Reflective setAccessible(true) disabled](https://www.reddit.com/r/Kotlin/comments/f4dfrx/javalangunsupportedoperationexception_reflective/)
- url: https://stackoverflow.com/questions/60241857/java-lang-unsupportedoperationexception-reflective-setaccessibletrue-disabled
---

## [10][How to run ktor application?](https://www.reddit.com/r/Kotlin/comments/f47a4r/how_to_run_ktor_application/)
- url: https://stackoverflow.com/questions/60233921/how-to-run-ktor-application/60237120#60237120
---

