# Kotlin
## [1][Kotlin and Exceptions](https://www.reddit.com/r/Kotlin/comments/h0alcw/kotlin_and_exceptions/)
- url: https://medium.com/@elizarov/kotlin-and-exceptions-8062f589d07
---

## [2][KSP - Kotlin Symbol Processing API](https://www.reddit.com/r/Kotlin/comments/h0gdth/ksp_kotlin_symbol_processing_api/)
- url: https://github.com/android/kotlin/tree/ksp/libraries/tools/kotlin-symbol-processing-api
---

## [3][Developing serverless applications with Kotlin, Micronaut and GraalVM](https://www.reddit.com/r/Kotlin/comments/h0bj5d/developing_serverless_applications_with_kotlin/)
- url: https://medium.com/tech-hunters/developing-production-ready-serverless-applications-with-kotlin-micronaut-and-graalvm-fff72d5c804b
---

## [4][Database suggestions for Kotlin applications](https://www.reddit.com/r/Kotlin/comments/h09y50/database_suggestions_for_kotlin_applications/)
- url: https://www.reddit.com/r/Kotlin/comments/h09y50/database_suggestions_for_kotlin_applications/
---
Hey guys,

I'm currently looking into setting up a relational database for my Kotlin project, and was wondering if any of you would like to share some experiences with DBMS that work well with Kotlin.

I'm specifically interested in DBMS that include object relational mapping to more easily parse between my object structure and the entity relationship model of the database.

I'm also open to any experiences or suggestions regarding the ERM design based on my class-model. I do have some theoretical knowledge, but little to know practical experience.

Also, while I don't mind additional knowledge, I'm not particularly interested in NoSQL databases, as they don't fit well with my requirements.

&amp;#x200B;

Edit: for anyone else interested in the issue, I'll compile the replies here:

1. anything that works well with Java works well with Kotlin. More specifically: anything that provides a JDBC driver
2. Jetbrains has their own ORM called Exposed. It's supposedly lightweight and easy to use, but apparently lacks documentation and some features
3. Just using the Java Ecosystem may be a viable option
4. Popular solutions include:
   1. Postgres
   2. Hibernate
   3. MySQL / MariaDB
   4. jOOQ + PostgreSQL

&amp;#x200B;
## [5][Is Android development in a functional programming style possible with kotlin?](https://www.reddit.com/r/Kotlin/comments/h0c4rd/is_android_development_in_a_functional/)
- url: https://www.reddit.com/r/Kotlin/comments/h0c4rd/is_android_development_in_a_functional/
---
I have been researching lately the current state of Android development since when I wanted to learn it the only choice was OOP Java (early Android years) and I stumbled over kotlin which seem a nice language with an easy to use/understand syntax and it does not force you to use classes at all.

&amp;#x200B;

I have been researching the kotlin environment and I found a lot of resources about kotlin functional programming but none of them were centered on android programming so I wanted to know if it is possible to develop Android apps using a fully functional approach (i.e.: writing only data classes and top level functions).
## [6][Connect Four Game](https://www.reddit.com/r/Kotlin/comments/h03gqz/connect_four_game/)
- url: https://www.reddit.com/r/Kotlin/comments/h03gqz/connect_four_game/
---
I'm wanting to create a connect four game in Kotlin for my final in college. I have to use a double array and play against the "computer". I just need help with where I need to start.

&amp;#x200B;

I wrote some pseudocode:

Prompt the user for their turn

Get the user's selection  (Column 1 through 8)

Place the piece in the double array,

Reprint the screen

Internally check to see if the user won and quit if they did with a "win" message

Show the computer's turn

Get the computer's selection  (Random number for Column 1 through 8)

Place the piece in the double array

Reprint the screen

Internally check to see if the computer won and quit if they did with a "win" message.

&amp;#x200B;

All of this would be done via Intellij console.

**Any help would be greatly appreciated.**
## [7][What are the possible ways/methods for connecting Kotlin/JVM (backend) with Dart (frontend, flutter) with minimum resource loss, i.e. in serialization/IO/deserialization?](https://www.reddit.com/r/Kotlin/comments/gzpq29/what_are_the_possible_waysmethods_for_connecting/)
- url: https://www.reddit.com/r/Kotlin/comments/gzpq29/what_are_the_possible_waysmethods_for_connecting/
---

## [8][Article: MVI in Kotlin Multiplatform — part 2](https://www.reddit.com/r/Kotlin/comments/gzrn2b/article_mvi_in_kotlin_multiplatform_part_2/)
- url: https://www.reddit.com/r/Kotlin/comments/gzrn2b/article_mvi_in_kotlin_multiplatform_part_2/
---
I have published the second part of the article "MVI in Kotlin Multiplatform — part 2". Here is the [Friend Link](https://badootech.badoo.com/mvi-in-kotlin-multiplatform-part-2-2-of-3-3faab535de02?source=friends_link&amp;sk=a7a347e49202e139d5cd7533d2a97141).
## [9][How is Android App UI done in Kotlin?](https://www.reddit.com/r/Kotlin/comments/gzhpvg/how_is_android_app_ui_done_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gzhpvg/how_is_android_app_ui_done_in_kotlin/
---
Hi, I am currently learning Kotlin basics on JetBrains Academy, because I think it's an awesome language. I learnt that it is often seen as the first choice for Android developers. I already have a little experience with Android apps via Dart/Flutter, but while the framework is generally good, I had a few problems with wrapping my head around some concepts. I hope that I will be able to make a fresh start with Kotlin.

While I still have to learn a lot about the basics before I can start developing apps, I was wondering about how an app UI is designed in Kotlin. I know that Kotlin is runnable on Android because of the JVM, but how do I design the frontend? Are there frameworks similar to Flutter or do I have to resort to Java frameworks like SWING or JavaFX (hopefully not)? And how is it done usually?

Currently I am not trying to develop an app, so I am just curious at this point, but later I might actually need this information.

Thanks!
## [10][Is there a Ktor equivalent to Django's built-in admin panel features?](https://www.reddit.com/r/Kotlin/comments/gzevqn/is_there_a_ktor_equivalent_to_djangos_builtin/)
- url: https://www.reddit.com/r/Kotlin/comments/gzevqn/is_there_a_ktor_equivalent_to_djangos_builtin/
---
I really like Django's built-in admin app for managing database entries for the models. Is there an equivalent for the Ktor framework?
