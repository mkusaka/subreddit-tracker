# Kotlin
## [1][Which Kotlin web development framework to learn?](https://www.reddit.com/r/Kotlin/comments/hs80qw/which_kotlin_web_development_framework_to_learn/)
- url: https://www.reddit.com/r/Kotlin/comments/hs80qw/which_kotlin_web_development_framework_to_learn/
---
I have been using Java/Spring for web development for years.

I recently have started learning Kotlin. I want to explore more and use/try different web development framework with Kotlin.

I know Spring  has a good support for Kotlin but I want to learn new framework that I can use in Production if chance permits.

Any suggestion between Ktor, Vertx or Micronaut and why?
## [2][Kotlin 1.4-M3: Generating Default Methods in Interfaces](https://www.reddit.com/r/Kotlin/comments/hs92c0/kotlin_14m3_generating_default_methods_in/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-m3-generating-default-methods-in-interfaces/
---

## [3][Should I first learn Kotlin/Native (multiplatform) or Swift?](https://www.reddit.com/r/Kotlin/comments/hs8e2d/should_i_first_learn_kotlinnative_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/hs8e2d/should_i_first_learn_kotlinnative_multiplatform/
---
I've done a lot of research, and Kotlin/Native and JVM finally seems like the answer to the endless quest for the best "write once, run anywhere" toolset.

I have very minimal programming experience, but I refuse to start with Python first for a few reasons. I'd eventually like to be able to have high-performance iOS and Android apps, macOS, Linux, and Windows apps using TornadoFX (or Kotlin/Native?), and run on the web via WebAssembly.

My problems are:

* I'll end up having to learn multiple languages (Swift, Kotlin, TypeScript, maybe Go) if I don't choose Kotlin Multiplatform, though I feel like this will happen anyway since I'm working with multiple SDKs and frameworks
* I don't have access to an Android device (Apple-based household), so physically running Android apps isn't viable and I'm not familiar with Google's ecosystem (don't want to be, but I have to be)
* Unlike Swift, there aren't a ton of resources to learn Kotlin, though I do own the continuously updated Atomic Kotlin book/Stepik course
* Multiplatform projects are still experimental, but I doubt I'll have a professional app ready before it's considered stable. I'm afraid a lot will change and break, making a more stable solution better for learning and getting things up and running
* I've heard Xcode is garbage compared to IntelliJ IDEA/Android Studio
## [4][jasync-sql and object mapping](https://www.reddit.com/r/Kotlin/comments/hs1isx/jasyncsql_and_object_mapping/)
- url: https://www.reddit.com/r/Kotlin/comments/hs1isx/jasyncsql_and_object_mapping/
---
Coming from the dotnet world and using dapper to map to a class, wondering if there's a way to use the speed of jasync-sql along with either a supporting library or by itself to do something like:

     val future = con.sendPreparedStatement(query, values)
     val result = future.get&lt;MyDataClass&gt;()

Is there an easy way to map to a data class - or should I be looking into jdbi, jooq, exposed, or something else entirely instead?
## [5][Kotlin extensions (KTX) for the Places SDK for Android](https://www.reddit.com/r/Kotlin/comments/hrtjfg/kotlin_extensions_ktx_for_the_places_sdk_for/)
- url: https://github.com/googlemaps/android-places-ktx
---

## [6][[Open Source] I'm working on an selfhosted web app for remote planning poker using Ktor and KotlinJS/React](https://www.reddit.com/r/Kotlin/comments/hrt1ul/open_source_im_working_on_an_selfhosted_web_app/)
- url: https://github.com/Foso/Showdown
---

## [7][Flym is a android app build kotlin it is open source and anyone can contribute to it.](https://www.reddit.com/r/Kotlin/comments/hs6d1v/flym_is_a_android_app_build_kotlin_it_is_open/)
- url: https://www.reddit.com/r/Kotlin/comments/hs6d1v/flym_is_a_android_app_build_kotlin_it_is_open/
---
Try the app it is I think best RSS reader app for Android and developer try to contribute to the I am not promoting the app I is really too good and serve best job
## [8][passing safe args to one fragment](https://www.reddit.com/r/Kotlin/comments/hrwd5o/passing_safe_args_to_one_fragment/)
- url: https://www.reddit.com/r/Kotlin/comments/hrwd5o/passing_safe_args_to_one_fragment/
---
I am currently working on todolist app and got stuck working with safe args. I have recyclerview adapter items and when I click on one item I want to navigate to that item and pass my data from the adapter to the detail fragment. If I tried creating a new item, I get the toast message "updated" instead of "created" and my cardview is blank. So it somehow is updating a non existence item. I have one layout and want to use this layout to either create a new item or pass existing item data into it and updated that said item.

Detail Fragment:  [https://pastebin.com/bYTxTWj7](https://pastebin.com/bYTxTWj7)

nav\_graph xml:  [https://pastebin.com/9cWbWWrY](https://pastebin.com/9cWbWWrY)

List adapter:  [https://pastebin.com/khsDFh4L](https://pastebin.com/khsDFh4L)

Cardview : [https://pastebin.com/XxAJgnnY](https://pastebin.com/XxAJgnnY)

github: [https://github.com/ikassim9/TodoList](https://github.com/ikassim9/TodoList) 
## [9][A problem with reading files. Might be encoding-related.](https://www.reddit.com/r/Kotlin/comments/hrkxg7/a_problem_with_reading_files_might_be/)
- url: https://www.reddit.com/r/Kotlin/comments/hrkxg7/a_problem_with_reading_files_might_be/
---
Hello! So I'm going through JetBrains Academy's Kotlin course, and I'm stuck with one of their projects, ASCII New Roman. Basically I am to render the text in ASCII form.

Provided is a file containing the "font". I first started with reading the file then simply spitting out all the lines, using this code:

    package signature
    
    import java.io.File
    import java.nio.charset.Charset
    import java.util.*
    
    fun main() {
        val roman = File("roman.txt").bufferedReader().readLines()

        for (s in roman) {
            println(s)
        }
    }

A portion of the file looks like this:

    10 52
    a 10
              
              
     .oooo.   
    `P  )88b  
     .oP"888  
    d8(  888  
    `Y888""8o 
              
              
              

Unfortunately the output spits out the following:

    10 52
    a 10
              
              
     .oooo.   
    `P  )88b  
     .oP&amp;quot;888  
    d8(  888  
    `Y888&amp;quot;&amp;quot;8o 
              
              
      

I suspected this might be some encoding error, but I already verified that the file is in UTF-8. I already tinkered with a few settings in the IDE to no avail. Other users seem to be experiencing this problem as well. I really don't know why special characters are getting converted to HTML codes.
## [10][Multiplatform Encryption with SQLDelight and SQLCipher](https://www.reddit.com/r/Kotlin/comments/hra59x/multiplatform_encryption_with_sqldelight_and/)
- url: https://dev.to/samhill303/multiplatform-encryption-with-sqldelight-and-sqlcipher-5do4
---

