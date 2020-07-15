# Kotlin
## [1][A problem with reading files. Might be encoding-related.](https://www.reddit.com/r/Kotlin/comments/hrkxg7/a_problem_with_reading_files_might_be/)
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
## [2][Multiplatform Encryption with SQLDelight and SQLCipher](https://www.reddit.com/r/Kotlin/comments/hra59x/multiplatform_encryption_with_sqldelight_and/)
- url: https://dev.to/samhill303/multiplatform-encryption-with-sqldelight-and-sqlcipher-5do4
---

## [3][Help. Started using Kotlin few hours ago. Got compiler. Started with atomic kotlin course. First Exercise got me error. Installed other version of gradle. But now there is simple problem i can't resolve](https://www.reddit.com/r/Kotlin/comments/hr842l/help_started_using_kotlin_few_hours_ago_got/)
- url: https://www.reddit.com/r/Kotlin/comments/hr842l/help_started_using_kotlin_few_hours_ago_got/
---
Please tell me where i am wrong

https://preview.redd.it/anzc5okijva51.png?width=908&amp;format=png&amp;auto=webp&amp;s=6c17a06b1da554d4c901b3f37302970e1121d72e
## [4][Talking Kotlin - Gradient Descent](https://www.reddit.com/r/Kotlin/comments/hr7z89/talking_kotlin_gradient_descent/)
- url: https://talkingkotlin.com/gradient-descent/
---

## [5][Kotlin - A Better, More Cloud-Friendly Java](https://www.reddit.com/r/Kotlin/comments/hqx6fs/kotlin_a_better_more_cloudfriendly_java/)
- url: https://www.youtube.com/watch?v=e0y1GJ7CfAk
---

## [6][Using C in Kotlin](https://www.reddit.com/r/Kotlin/comments/hr7zr7/using_c_in_kotlin/)
- url: https://www.youtube.com/watch?v=WlNzx6GEpkM
---

## [7][How do I make a custom getter for a hashmap?](https://www.reddit.com/r/Kotlin/comments/hrcj55/how_do_i_make_a_custom_getter_for_a_hashmap/)
- url: https://www.reddit.com/r/Kotlin/comments/hrcj55/how_do_i_make_a_custom_getter_for_a_hashmap/
---
    fun getArenaSpawn(id: Int): Location = Location(
            plugin.config.getString("arenas.${id}.world")?.let { Bukkit.getWorld(it) },
            plugin.config.getDouble("arenas.${id}.x", 0.0),
            plugin.config.getDouble("arenas.${id}.y", 0.0),
            plugin.config.getDouble("arenas.${id}.z", 0.0),
            plugin.config.getDouble("arenas.${id}.yaw", 0.0).toFloat(),
            plugin.config.getDouble("arenas.${id}.pitch", 0.0).toFloat())

So I have the following getter code but I don't want to use getArenaSpawn(id) I want to be able to do arenaSpawn\[id\] but when I try making a hashmap with a custom getter I'm completely lost because I don't know what the keyword to reference the index being used is. So my question is, for a HashMap*&lt;*Int, Location*&gt;* hashmap how would I reference the Int in the getter?
## [8][Kotlin Microservices On Kubernetes - Part 4: Best Practices](https://www.reddit.com/r/Kotlin/comments/hr1vo4/kotlin_microservices_on_kubernetes_part_4_best/)
- url: https://www.youtube.com/watch?v=ayoT9U4Pd8Y&amp;feature=share
---

## [9][Should we always use the `this` keyword?](https://www.reddit.com/r/Kotlin/comments/hr3rak/should_we_always_use_the_this_keyword/)
- url: https://www.reddit.com/r/Kotlin/comments/hr3rak/should_we_always_use_the_this_keyword/
---
IMO the `this` keyword is an outdated java construct that should only be used when absolutely necessary in Kotlin. However, at my current job, we are forced to use the `this` keyword everywhere we can because it makes your code more concise and easier to understand in a text editor (or GitHub). I do not believe we should write code for GitHub, I believe we write code for our environment. Modern IDE's have plenty of syntaxes that makes the `this` keyword obsolete. What are your thoughts on using `this` in your Kotlin code?
## [10][Getting started with Ktor](https://www.reddit.com/r/Kotlin/comments/hqwhqz/getting_started_with_ktor/)
- url: https://www.mscharhag.com/kotlin/ktor-getting-started
---

