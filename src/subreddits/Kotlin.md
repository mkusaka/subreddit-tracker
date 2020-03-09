# Kotlin
## [1][Koltin with more than one class in the same file, is this break the single responsibility concept?](https://www.reddit.com/r/Kotlin/comments/fftmpv/koltin_with_more_than_one_class_in_the_same_file/)
- url: https://www.reddit.com/r/Kotlin/comments/fftmpv/koltin_with_more_than_one_class_in_the_same_file/
---
Hey folks, I have a question that in my team we have different opinions about this question. 

Is it better to use a KT file that holds more than one class or not?  for example, in android, I have a file that holds an Adapter class and its ViewHolder class. A friend of mine says that breaks the single responsibility in SOLID. Can you tell me about what do you think?
## [2][Kotlin Note App (Github Repository)](https://www.reddit.com/r/Kotlin/comments/fffjht/kotlin_note_app_github_repository/)
- url: https://github.com/maxwellnewage/kotlin-notes
---

## [3][Kotlin + JPA - val vs var](https://www.reddit.com/r/Kotlin/comments/fffgev/kotlin_jpa_val_vs_var/)
- url: https://www.reddit.com/r/Kotlin/comments/fffgev/kotlin_jpa_val_vs_var/
---
I'm new to Kotlin and I was wondering what's the best way to design your JPA entities? I've seen some examples for very basic Kotlin+JPA usage and they all used `val` for all fields within the entity which is a nice design but introduces the problem that you can't manipulate the values of an existing entity.

Those examples never manipulate an existing entity attached to the Hibernate session. 

What do you think is the best way?
## [4][why do i need the (Int)-&gt;Unit for the lambda function definition ?](https://www.reddit.com/r/Kotlin/comments/ffe7uz/why_do_i_need_the_intunit_for_the_lambda_function/)
- url: https://www.reddit.com/r/Kotlin/comments/ffe7uz/why_do_i_need_the_intunit_for_the_lambda_function/
---
        fun main(args: Array&lt;String&gt;){  
            val myLambda: (Int) -&gt; Unit= {s: Int -&gt; println(s) } //lambda function  
            addNumber(5,10,myLambda)  
        }  
        fun addNumber(a: Int, b: Int, mylambda: (Int) -&gt; Unit ){   //high level function lambda as parameter  
            val add = a + b  
            mylambda(add) // println(add)  
        }  

what is the need to use it in both the definition and lambda function call in the addnumber function?
## [5][Kotlin + Jersey + Jetty + MongoDB – creating a scalable RESTful API](https://www.reddit.com/r/Kotlin/comments/ffjj0v/kotlin_jersey_jetty_mongodb_creating_a_scalable/)
- url: https://blog.gikken.co/kotlin-jersey-jetty-mongodb-creating-a-restful-api/
---

## [6][Return result to activity](https://www.reddit.com/r/Kotlin/comments/ffh26r/return_result_to_activity/)
- url: https://www.reddit.com/r/Kotlin/comments/ffh26r/return_result_to_activity/
---
Im trying to use coroutines in android and I want to return result to the UI(fragment), so I did the following:

Inside the repository class:

`suspend fun saveNotes(text : String) : Deferred&lt;String?&gt; {return CoroutineScope(Dispatchers.IO).async {saveNotesToDb(text)}}`

`suspend fun saveNotesToDb(text: String) : String? {var result : String? = ""try{val notes = Notes(notes = text)notesDb.notesDAO.insertNote(notes)result = "saved data to the database"return result}catch (e: Exception){print(e.stackTrace)result = e.messagereturn result}}`

So basically I create suspend function saveNotesToDB, and then in another function saveNotes I return a Deferred object so I can use the result later. Then inside the viewmodel i do the following:

`suspend fun setNotes(text : String) : String? {return notesRepository.saveNotes(text).await()}`

So I thought I can use await() here to get the result, but then again in the fragment I need to create a suspend function, how can I leave the suspend function. All i want is for the result to go from the repository --&gt; viewmodel --&gt; display in fragment inside a snackbar

&amp;#x200B;

Can someone please tell me how to fix this
## [7][JetBrains Space](https://www.reddit.com/r/Kotlin/comments/fev7jg/jetbrains_space/)
- url: https://www.reddit.com/r/Kotlin/comments/fev7jg/jetbrains_space/
---
Lately, Jetbrains just announced their brand new product Space. They provide us programmers, an early access preview. Jetbrains had improved our experiences to create new applications and collaborations with the team members.

Learn more about them at [https://www.jetbrains.com/](https://www.jetbrains.com/?fbclid=IwAR232BrBEzCOX6Sg-uEhpzpmA3pf0jz61sWNQe852DPe3GHQWOqYAOD7Ijc) and the blog about Space [https://blog.jetbrains.com/space/2019/12/05/welcome-space/](https://blog.jetbrains.com/space/2019/12/05/welcome-space/?fbclid=IwAR3PYv7-xGVO_vz6QAidKwtm8_ipSg9lfiiK3XCgUlI3uZKZMz-coX4M2cw).
## [8][Tutorial: Variables, Data Types, Immutability and User Input with Kotlin](https://www.reddit.com/r/Kotlin/comments/fethd6/tutorial_variables_data_types_immutability_and/)
- url: https://marcuseisele.com/pages/learningKotlin/variables
---

## [9][Kotlin implementation of CommonMark spec](https://www.reddit.com/r/Kotlin/comments/feq3fw/kotlin_implementation_of_commonmark_spec/)
- url: https://www.reddit.com/r/Kotlin/comments/feq3fw/kotlin_implementation_of_commonmark_spec/
---
[https://github.com/matt-belisle/CommonMarK/tree/master/src/main/kotlin/com/matt/belisle/commonmark](https://github.com/matt-belisle/CommonMarK/tree/master/src/main/kotlin/com/matt/belisle/commonmark)

&amp;#x200B;

Finished this recently, a kotlin only implementaiton of the CommonMark MarkDown spec, was a pretty interesting deeper dive into kotlin than I have ever done before, did profiling on a project for the first time, had to design the architecture of a medium sized project, and face some of the consequences of some of my decisions and how they impacted the programming experience
## [10][Making my own apps or freelance to make money under 18 with kotlin?](https://www.reddit.com/r/Kotlin/comments/fexgu8/making_my_own_apps_or_freelance_to_make_money/)
- url: https://www.reddit.com/r/Kotlin/comments/fexgu8/making_my_own_apps_or_freelance_to_make_money/
---
Hi.I am 15 and I want to make money from programming because i don't want to take a summer job:)).I am still learning android development and I have 2 ways of making money:

1.Building my own app:I have some ideas but i think is very hard to publish and have success with it if i am under 18 and have little knowleadge about app marketing but i whould like to try this thought.

2.Freelancing:I heard that i can make some  money out of freelancing even if i am under 18 but I am not sure if it is true.I whould like to freelance too.

Soo out of this ways of making income which one do you guys thing i should use?And why?Please help me if you can,any advice is good.

Btw:I DON'T WANT TO MAKE MUCH MONEY ONLY A FEW JUST TO BE ABLE TO STAY HOME AND CODE ALL SUMMER AND NOT TAKING A SUMMER JOB
