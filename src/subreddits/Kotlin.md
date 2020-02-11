# Kotlin
## [1][The state of Kotlin for microservices](https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/)
- url: https://www.reddit.com/r/Kotlin/comments/f26a7w/the_state_of_kotlin_for_microservices/
---
Hi all 

I need your advices and would like to know, if it is the right choice to do microservices in Kotlin.
Of course, Golang would be the better choice, but I do not like it because of boilerplate codes.
Rust seems to be interesting, but libraries are not stable yet.
One point, that I concern about it, is the big image size of JVM. But I think with GraalVM is good to go.

The question is, should I use Kotlin to build my microservices or Kotlin is only suitable for frontend programming? What is your opinion? 

Thanks
## [2][Getting Started with Kotlin on iOS, Part 1](https://www.reddit.com/r/Kotlin/comments/f1uhwh/getting_started_with_kotlin_on_ios_part_1/)
- url: https://benasher.co/kotlin-ios-getting-started/
---

## [3][does anyone know if kotlin/native already support curses/ncurses on linux?](https://www.reddit.com/r/Kotlin/comments/f1pmwt/does_anyone_know_if_kotlinnative_already_support/)
- url: https://www.reddit.com/r/Kotlin/comments/f1pmwt/does_anyone_know_if_kotlinnative_already_support/
---

## [4][Accelerate Your Kotlin Multiplatform Evaluation with KaMP Kit](https://www.reddit.com/r/Kotlin/comments/f1p6nq/accelerate_your_kotlin_multiplatform_evaluation/)
- url: https://blog.jetbrains.com/kotlin/2020/02/accelerate-your-kotlin-multiplatform-evaluation-with-kamp-kit/
---

## [5][Run code with "delay" in a browser](https://www.reddit.com/r/Kotlin/comments/f1c31k/run_code_with_delay_in_a_browser/)
- url: https://www.reddit.com/r/Kotlin/comments/f1c31k/run_code_with_delay_in_a_browser/
---
Hi everyone! I was looking at code examples here [https://kotlinlang.org/docs/tutorials/coroutines/coroutines-basic-jvm.html](https://kotlinlang.org/docs/tutorials/coroutines/coroutines-basic-jvm.html)

and you can notice that if we'll run this code we won't see delays between printing lines, because the code sends to the server and executes there, that's why we just receive the result.

https://preview.redd.it/581tl4k9pxf41.png?width=752&amp;format=png&amp;auto=webp&amp;s=5a517c804ff7bfe810ad57b905ca16651148c4f2

Is it possible to execute such codes somehow so we'll be able to see delays? (For example, after pushing Run button there appears a window just like on the screen, but "Start", "Hello" and "Stop" appear with the 1 and 2 seconds intervals) Maybe we can convert kotlin to JavaScript and run it in the browser? I'm only studying and not a professional so I'm not even sure if it's possible. I tried to google this problem but didn't succeed. It would be great if you can give me advice or suggestion about how this problem can possibly be solved or where to find the information.
## [6][abstract fun](https://www.reddit.com/r/Kotlin/comments/f1cait/abstract_fun/)
- url: https://www.reddit.com/r/Kotlin/comments/f1cait/abstract_fun/
---
What is the difference between using this:

`abstract val sleepDatabaseDao: SleepDatabaseDao`

vs using this:

`abstract fun chapterDao(): ChapterDao`

Cant I in both cases as the methods inside the DAO interface?
## [7][[Discussion] Thoughts on making keywords like "data" in "data class" definable?](https://www.reddit.com/r/Kotlin/comments/f1el4x/discussion_thoughts_on_making_keywords_like_data/)
- url: https://www.reddit.com/r/Kotlin/comments/f1el4x/discussion_thoughts_on_making_keywords_like_data/
---
I recently thought about this, the benefit would be that classes that for example are used by libraries can be defined in a way that (often) doesn't require annotations.  


Please give feedback and don't just downvote when you disagree, I want to hear what speaks against this and what you think.  


For example when looking at the Room Android library (a ORM library) a class like this:  


`@Entity`  
`data class User(`  
 `@PrimaryKey val uid: Int,`  
 `val firstName: String?,`  
 `val lastName: String?`  
`)`   


Could look like this:  


  
`entity class User(`  
 `val uid: PrimaryKey&lt;Int&gt;,`  
 `val firstName: String?,`  
 `val lastName: String?`  
`)`   


Keywords like data could also be able to inherit so **entity** could inherit **data** but add some extra functionality, in this case for example uid doesn't really have to even have a PrimaryKey specification at all when the "**entity**" keyword takes the first constructor parameter as PrimaryKey by default. This would look like this:  


`entity class User(val uid: Int,val firstName: String?,val lastName: String?)`   


 Kotlin got rid of @ Override as well, so why should we have to annotate stuff so frequently when it could also be done like this?  


Another example would be TestClasses in Android, You could just use a "test class XX" keyword and it could get rid of some basic stuff.  


This would basically be like inheritance but instead of fields and methods being inherited you actually get the way that you declare stuff inherited.   


**Let me address some concerns:**  
*- "This would be weird and looking up how to define stuff for every library that uses it would take time"*  
\-&gt; No, it would probably take less time than learning how stuff has to be annotated and it would be way easier to understand what e.g. Room entity classes want from you by simply looking at their source  


*- "This would clutter the namespace"*  
\-&gt; Fair point, but this shouldn't be too much of a problem as it's only before class declarations. Also the keyword could just be imported like everything else so no need to specifiy "room.entity class" when importing "room.entity" as a keyword (e.g. when you have multiple "entity" defining libraries)  


*-"Infrequent use case"*  
\-&gt; I don't think so, Kotlin could even have some extra keywords included in it's standard library  


The Room library has 2 other important class annotations which are: @ Dao and @ Database  
Both of those need to be put before abstract classes / interfaces so you could put them into keywords that extend the "**abstract class**" keyword.  


So yeah, the thing I'm proposing is not having a special case for the data keyword but rather giving us the option to make up our own keywords like it. These classes shouldn't be created frequently but if this feature would exist I think it would be used in a lot of cases.  


Tell me what you think
## [8][Trying to proxy data class using ModelMapper](https://www.reddit.com/r/Kotlin/comments/f183zk/trying_to_proxy_data_class_using_modelmapper/)
- url: https://www.reddit.com/r/Kotlin/comments/f183zk/trying_to_proxy_data_class_using_modelmapper/
---
So I'm using ModelMapper to convert between DTO and DB objects. I've got a scenario where my DB object has a few fields that I do not want to expose via the API. So I'm running an update operation, and I'm trying to use ModelMapper to overwrite the DB entity with data from the DTO. See below:

&amp;#x200B;

`return videoFileRepo.findById(fileId)`  
`.map { videoFile -&gt;`  
`modelMapper.map(payload, videoFile)`  
`videoFile.fileId = fileId`  
`val savedVideoFile =` [`videoFileRepo.save`](https://videoFileRepo.save)`(videoFile)`  
`modelMapper.map(savedVideoFile, VideoFilePayload::class.java)`  
`}`  
`.orElse(null)`

My problem is that I have nested entities in the VideoFile object. The DTO has DTO versions of those entities, the DB entity has DB entity versions. So a set of CategoryPayload vs a set of Category, basically. Using ModelMapper this way, with a destination instance instead of a class, doesn't map the nested objects properly.

The solution, I would think, is to create a ModelMapper TypeMap for my classes. That way I could provide more explicit instructions. However, my VideoFilePayload and VideoFile entity objects are both data classes, and ModelMapper throws an error about how it needs to be able to proxy the classes for a TypeMap, and it can't do that with data classes.

So, I'm hoping there is some workaround here where I can proxy my data classes and create the necessary TypeMap. Thanks.
## [9][From Java to Kotlin: life without static](https://www.reddit.com/r/Kotlin/comments/f0t306/from_java_to_kotlin_life_without_static/)
- url: https://jelmini.dev/post/from-java-to-kotlin-life-without-static/
---

## [10][Migrating Java to kotlin (4 minute video)](https://www.reddit.com/r/Kotlin/comments/f0trz5/migrating_java_to_kotlin_4_minute_video/)
- url: https://www.youtube.com/watch?v=w9WE0c6fjRQ
---

