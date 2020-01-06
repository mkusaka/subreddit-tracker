# Kotlin
## [1][KotlinConf 2019 Videos](https://www.reddit.com/r/Kotlin/comments/ebd0np/kotlinconf_2019_videos/)
- url: https://www.youtube.com/playlist?list=PLQ176FUIyIUY6SKGl3Cj9yeYibBuRr3Hl
---

## [2][Kotlin Tutorials](https://www.reddit.com/r/Kotlin/comments/ekro1y/kotlin_tutorials/)
- url: https://kotlintutorialblog.com
---

## [3][Kotlin closed classes: how to compose with them](https://www.reddit.com/r/Kotlin/comments/ekjb5c/kotlin_closed_classes_how_to_compose_with_them/)
- url: https://jelmini.dev/post/2020-01-02-kotlin-classes-final-by-default/
---

## [4][Handling Clicks in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/eksurq/handling_clicks_in_jetpack_compose/)
- url: https://proandroiddev.com/handling-clicks-in-jetpack-compose-3800379845c4
---

## [5][Kotlin Algorithm Challenge No. 2](https://www.reddit.com/r/Kotlin/comments/ekjscp/kotlin_algorithm_challenge_no_2/)
- url: https://www.alexhwoods.com/blog/longest-substring
---

## [6][FXGL (Java / Kotlin Game Framework) 11.7 with A* support and improved Kotlin API](https://www.reddit.com/r/Kotlin/comments/ekbp3k/fxgl_java_kotlin_game_framework_117_with_a/)
- url: https://www.reddit.com/r/Kotlin/comments/ekbp3k/fxgl_java_kotlin_game_framework_117_with_a/
---
Check out the new version at:  [https://github.com/AlmasB/FXGL](https://github.com/AlmasB/FXGL) 

Some example screenshots from existing games / demos:

[FXGL Games](https://preview.redd.it/ca4igbjjzx841.jpg?width=2561&amp;format=pjpg&amp;auto=webp&amp;s=04863a2c9c9081aab328b820a98142eea2233117)

There is a wiki for new users:  [https://github.com/AlmasB/FXGL/wiki](https://github.com/AlmasB/FXGL/wiki) 

New contributors are always welcome!
## [7][Executing static block outside of class?](https://www.reddit.com/r/Kotlin/comments/ekgzu1/executing_static_block_outside_of_class/)
- url: https://www.reddit.com/r/Kotlin/comments/ekgzu1/executing_static_block_outside_of_class/
---
Hello. A little new to Kotlin, with some familarity in java.

As a learning exercise, I am trying to convert a Java library to Kotlin. (If it matters, I am talking about algs4 from Sedgewick and Wayne companion site to algorithms 4th edition).

I am trying to convert a static class to a bunch of standalone functions.
Now, what is the proper way to do it?
I understand, that I could just create kotlin class and call all the static methods with companion object, however I am curious what is the "right" way to do it.
1. Create an object, with the same name as the static class and put everything inside of it? 
2. Just write standalone functions together with global objects?
3. Do it as it was done in java - a static class (companion object + JvmStatic).

Also, there is a static block in the static class I am trying to convert. If I understand correctly, the code inside of static block should be put inside of init function in (companion) object.
However, what is the alternative to static block (I mean static{ somecode }) in java if my code is not inside of an object or a class?

Thanks.

Thanks.
## [8][How do I make a database connection a member variable?](https://www.reddit.com/r/Kotlin/comments/eki8a3/how_do_i_make_a_database_connection_a_member/)
- url: https://www.reddit.com/r/Kotlin/comments/eki8a3/how_do_i_make_a_database_connection_a_member/
---
I'm new to Kotlin, coming from the C++ world, so making an uninitialized member variable is stumping me:

`class Dao {`  
 `private var connectResult = String()`  
 `private val dbcon: Connection = ???`  


`public fun connect() {`  
 `val properties = Properties()`  
 `with(properties) {`  
 `put("user", "root")`  
 `put("password", "foo")`  
 `}`  
 `connectResult = "connection failed"`  
 `println("Attempting to connect to database")`  
 `try {`  
 `dbcon = DriverManager.getConnection("jdbc:mariadb://database/data", properties)`  
 `println("Connection to mariadb succeded!")`  
 `connectResult = "connection to mariadb succeded"`  
 `}`  
`//        catch(e: SQLException)`  
 `catch (e: SQLException) {`  
 `connectResult = e.toString()`  
`//            connectResult = "SQLException thrown"`  
 `}`  
   
`}`  
`}`

So up at the top, my private val dbcon: Connection, I'm not sure what to set it to. It doesn't get initialized until inside the try statement. I want it as a member variable - or even at least a variable that isn't scoped inside the try statement so that I can pass it around and continue to perform operations on the database.

How do I accomplish this correctly with kotlin? I could also be using the JDBC library incorrectly too, as I'm coming from the C++ world.
## [9][Can the Kotlin type system be leveraged to programmatically generate the map of HTTP and message bus inputs to all possible message bus outputs for a Kotlin service? Can the Arrow library help accomplish this?](https://www.reddit.com/r/Kotlin/comments/ek5030/can_the_kotlin_type_system_be_leveraged_to/)
- url: https://www.reddit.com/r/Kotlin/comments/ek5030/can_the_kotlin_type_system_be_leveraged_to/
---
My team is currently developing an application using a microservice architecture with pub-sub interservice communication over a message bus. Certain types of HTTP requests and messages received as inputs to a service might cause that service to subsequently publish certain types of messages over the bus as a result, and these published messages might trigger further publication events from other services as well.

I'd like to be able to programmatically generate a graph of all of the possible message flows without having to run any of these services - I know it's possible to do this through runtime analysis, but I'd like to solve this problem statically if possible. In our Kotlin services, messages are statically typed, and I believe it should be possible to programmatically determine message flow inputs and outputs for each Kotlin service at compile time or through static code analysis.

How should I start to tackle this problem? As one possible approach, I'm trying to determine if the Arrow functional programming library can encode potential message output side effects (of specific message types) at the type level, but I'm not sure if that's possible. For instance, is there a way to extend the Arrow IO type so that functions that directly or indirectly produce a certain type of message as a side effect can only be run in a monad of such an extended type which explicitly allows that (and which therefore implicitly disallows producing messages of types that are not whitelisted)?
## [10][Anyone would like to add support for Kotlin/Native in FOSS Godot game engine?](https://www.reddit.com/r/Kotlin/comments/ejuxss/anyone_would_like_to_add_support_for_kotlinnative/)
- url: https://twitter.com/reduzio/status/1213082947965853696
---

## [11][Book for Mastering Kotlin](https://www.reddit.com/r/Kotlin/comments/ejwqk4/book_for_mastering_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ejwqk4/book_for_mastering_kotlin/
---
Hey guys, could you recommend me some Kotlin advanced books?  


Thanks in advance
