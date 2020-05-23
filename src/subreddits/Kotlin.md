# Kotlin
## [1][Trying to understand something about Kotlin's Collections interface and Java](https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/)
- url: https://www.reddit.com/r/Kotlin/comments/gp1hil/trying_to_understand_something_about_kotlins/
---
I'm using Data Binding on Android and I'm trying to understand why I have to import `java.util.List` into my data binding layouts even tho the rest of the code is using Kotlin:

 [https://imgur.com/a/jFiIfy9](https://imgur.com/a/jFiIfy9) 

I tried replacing this for `kotlin.collections.List` but the code doesn't compile. Online I read something about that the Kotlin collection interfaces "don't actually exist". Can someone shed some light on this?
## [2][Code review/help for a Calculator project I'm working on](https://www.reddit.com/r/Kotlin/comments/gp37ai/code_reviewhelp_for_a_calculator_project_im/)
- url: https://www.reddit.com/r/Kotlin/comments/gp37ai/code_reviewhelp_for_a_calculator_project_im/
---
I'm working on a smart calculator project for a site that teaches you Kotlin.

## ---- Description ----

Write a program that reads two numbers in a loop and prints the sum in the standard output. If a user enters only a single number, the program should print the same number. If a user enters an empty line, the program should ignore it.  


When the command 

    /exit

is entered, the program must print 

    "Bye!"

(without quotes), and then stop.  
\----- End Description ----

I wrote some code which I thought meets the requirements but when the project is checked, I get an error stating this input is incorrect, which I'm not sure why it's incorrect  


    java.lang.AssertionError: Wrong answer in test #2  Please find below the output of your program during this failed test. Note that the '&gt;' character indicates the beginning of the input line.  
    ---  
    &gt; 17 9 
    26

Here's the bulk of my logic. I read the inputs as String initially so that I can perform my check. Then convert to Int when I need to calculate.   


    when {
            // if user enters "/exit", print "Bye!"
            a == "/exit" -&gt; {
                print("Bye!")
            }
            // if user enters empty line, ignore it
            a.isEmpty() -&gt; {
            }
            // if user enters single #, print same #
            b.isEmpty()-&gt; {
                print(a)
            }
            // if user enter 2 #'s, print the sum
            else -&gt; {
                print(addTwoNumbers(a.toInt(), b.toInt()))
            }
        }

Are there any issues w/ my logic? What would you do differently?
## [3][30 days of Kotlin - Seminar 2 - Decoding Kotlin: The Modern Way To Build on Android](https://www.reddit.com/r/Kotlin/comments/gp044m/30_days_of_kotlin_seminar_2_decoding_kotlin_the/)
- url: /r/androiddev/comments/gp038k/30_days_of_kotlin_seminar_2_decoding_kotlin_the/
---

## [4][The Result Monad](https://www.reddit.com/r/Kotlin/comments/gonrds/the_result_monad/)
- url: https://adambennett.dev/2020/05/the-result-monad/
---

## [5][Can I make a function similar to the repeat() function?](https://www.reddit.com/r/Kotlin/comments/gp2rbh/can_i_make_a_function_similar_to_the_repeat/)
- url: https://www.reddit.com/r/Kotlin/comments/gp2rbh/can_i_make_a_function_similar_to_the_repeat/
---
If I understand it correctly, repeat() is a special function. You can pass in parameters and you also need to specify the body of the function, like this:

repeat(3) {
  println("This will print 3 times.")
}

Can I make my own function that will have this functionality? That the user can pass his own code like in the repeat function.
## [6][🌊StateFlow, End of LiveData?](https://www.reddit.com/r/Kotlin/comments/gowz14/stateflow_end_of_livedata/)
- url: https://link.medium.com/sm7NfjtoH6
---

## [7][Absolem - a knowledge sharing platform](https://www.reddit.com/r/Kotlin/comments/gp1bpj/absolem_a_knowledge_sharing_platform/)
- url: https://www.reddit.com/r/Kotlin/comments/gp1bpj/absolem_a_knowledge_sharing_platform/
---
 Hi, I just published my real-world knowledge platform project using Vert.x Kotlin, DDD, Event Driven Design, GraphQL and lots of stuff. Come and take a look!

 [https://github.com/UkonnRa/wonderland](https://github.com/UkonnRa/wonderland)
## [8][Secondary Constructors](https://www.reddit.com/r/Kotlin/comments/gomb0f/secondary_constructors/)
- url: https://www.reddit.com/r/Kotlin/comments/gomb0f/secondary_constructors/
---
I've just started learning Kotlin as I have free time on my hands right now; I have prior experience in Java and python (may be relevant). When learning about secondary constructors I understand how to use them, just not why they would be useful, why not just include the extra values in the first constructor and assign them as null unless used? Or use an overload to handle objects with different constructors? I'm really confused as to how they are beneficial, any help is appreciated :)))
## [9][The schedule &amp; talks are out for Kotliners 2020](https://www.reddit.com/r/Kotlin/comments/golgpv/the_schedule_talks_are_out_for_kotliners_2020/)
- url: https://kotliners.com/conference
---

## [10][Question about singletons](https://www.reddit.com/r/Kotlin/comments/gojnfo/question_about_singletons/)
- url: https://www.reddit.com/r/Kotlin/comments/gojnfo/question_about_singletons/
---
I've been studying Kotlin and I just got to the section about creating a singleton with the keyword object, and I was wondering if my understanding of objects and companion objects is correct? With the following Player class, from what I understand if there was 20 Player objects, they each would have a reference to the one companion object ( and any other object that may get added to the class ) that then can be used. But if the companion object was replaced with individual functions they'd all have their own copy of it that then could lead to different results for each Player if there was something within it that could be changed. 
```
class Player(val id: Int) {
 
    /* ... */
    companion object Properties {
        /* Default player speed in playing field - 7 cells per turn */
        val defaultSpeed = 7
 
        fun calcMovePenalty(cell: Int): Int {
            /* calc move speed penalty */
        }
    }
}
```
