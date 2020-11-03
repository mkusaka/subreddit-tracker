# Kotlin
## [1][Kotlin Team AMA #2: Ask Us Anything](https://www.reddit.com/r/Kotlin/comments/ji9z19/kotlin_team_ama_2_ask_us_anything/)
- url: https://www.reddit.com/r/Kotlin/comments/ji9z19/kotlin_team_ama_2_ask_us_anything/
---
**UPDATE:** All done, thanks for all the questions! Note that you can always contact us via the Kotlin Slack or Twitter, so if you missed the AMA, feel free to post your questions there. If you’re yet not on Kotlin Slack, please request your invite here: [https://surveys.jetbrains.com/s3/kotlin-slack-sign-up](https://surveys.jetbrains.com/s3/kotlin-slack-sign-up)

Hi everyone! The JetBrains team is here and we will be answering your questions over the next \~36 hours:

* [/u/abreslav](https://www.reddit.com/u/abreslav/), Project Lead for Kotlin
* [/u/elizarov](https://www.reddit.com/u/elizarov/), Team Lead for Kotlin Language Research
* [/u/etolstoy](https://www.reddit.com/user/etolstoy/), Team Lead for Kotlin Product Management
* [/u/erokhins](https://www.reddit.com/user/erokhins), Development Lead for Kotlin
* [/u/svtk\_](https://www.reddit.com/u/svtk_/), Team Lead for Kotlin Advocacy
* [/u/udalov](https://www.reddit.com/user/udalov/), Team Lead for Kotlin/JVM
* [/u/qwwdfsad](https://www.reddit.com/user/qwwdfsad/), Team Lead for Kotlin Libraries
* [/u/dsavvinov](https://www.reddit.com/user/dsavvinov/), Team Lead for Kotlin Multiplatform
* [/u/terrakok](https://www.reddit.com/user/terrakok/), Team Lead for KMM (Kotlin Multiplatform Mobile)
* [/u/ekaterina-petrova](https://www.reddit.com/user/ekaterina-petrova/), Developer Advocate for KMM (Kotlin Multiplatform Mobile)
* [/u/akapanina](https://www.reddit.com/user/akapanina/), Product Manager for KMM (Kotlin Multiplatform Mobile)
* [/u/demiurg\_906](https://www.reddit.com/user/demiurg_906/), Software Developer for Kotlin Compiler
* [/u/sebi\_io](https://www.reddit.com/user/sebi_io/), Developer Advocate for Kotlin
* [/u/bashor\_](https://www.reddit.com/user/bashor_), Team Lead for Kotlin/JS
* [/u/EugeneAndroid](https://www.reddit.com/user/EugeneAndroid/), Product Manager for Kotlin for Android, Kotlin Build Tools
* [/u/fundamentalparticle](https://www.reddit.com/user/fundamentalparticle/), Developer Advocate for Kotlin Server-side
* [/u/hhariri](https://www.reddit.com/u/hhariri/), Team Lead for Ktor, Developer Advocate for Kotlin
* [/u/anton-yalyshev](https://www.reddit.com/user/anton-yalyshev/), Product Manager for Kotlin IDE
* [/u/mzarechenskiy](https://www.reddit.com/user/mzarechenskiy), Team Lead for Kotlin IDE, Compiler Core (Front-end), Language Specification
* [/u/yanex](https://www.reddit.com/user/yanex/), Software Developer for Kotlin IDE
* [/u/Artyom\_Degtyarev](https://www.reddit.com/user/Artyom_Degtyarev/), Support Engineer for Kotlin/Native
* [/u/SvyatoslavScherbina](https://www.reddit.com/user/SvyatoslavScherbina), Team Lead for Kotlin/Native
* [/u/\_Wo\_0dy\_](https://www.reddit.com/user/_Wo_0dy_/), Product Manager for Kotlin Website, Kotlin/JS Teams, and Kotlin Analytics
* [u/wild\_\_lynx](https://www.reddit.com/user/wild__lynx/), QA for Kotlin Multiplatform
* [/u/alinagrebenkina](https://www.reddit.com/user/alinagrebenkina), Team Lead for Kotlin Marketing
* [/u/Belosnegova](https://www.reddit.com/user/Belosnegova), Kotlin Marketing Manager for Education
* [/u/meilalina](https://www.reddit.com/user/meilalina), Kotlin Marketing Manager for Community and Kotlin Server-side
* [/u/Anisim\_1](https://www.reddit.com/user/Anisim_1), Kotlin Marketing Manager for KMM (Kotlin Multiplatform Mobile)
* [/u/yole](https://www.reddit.com/user/yole/), Product Manager for IntelliJ IDEA
* /u/ilya-g, Developer for Kotlin Standard Library and kotlinx-datetime

&amp;#x200B;

We’ll be joined by other developers and team leads from Kotlin, as well as two special guests: Wojtek Kaliciński /u/wkalicinski from Google and Sebastien Deleuze /u/sdeleuze from Spring.

On October 12–15 we hosted the Kotlin 1.4 Online Event, and over the course of those four days we received many interesting questions. Since we did not manage to cover all of them due to time constraints, we will answer them in a specific thread under this post.

Please post your questions as top-level comments to this post.

We look forward to hearing from you!
## [2][Wow...Compose desktop (mac/linux/win) is now available with the latest plugin.](https://www.reddit.com/r/Kotlin/comments/jn6wio/wowcompose_desktop_maclinuxwin_is_now_available/)
- url: https://i.redd.it/h3ul68wqlzw51.png
---

## [3][My first Kotlin project](https://www.reddit.com/r/Kotlin/comments/jnaank/my_first_kotlin_project/)
- url: https://v.redd.it/ogqte3sr11x51
---

## [4][!null -&gt; !null contract](https://www.reddit.com/r/Kotlin/comments/jn5hxj/null_null_contract/)
- url: https://www.reddit.com/r/Kotlin/comments/jn5hxj/null_null_contract/
---
Is there a Kotlin construct (e.g. a contract) where I can define that `fun(blah: X?): X`? will return X when given an X; something similar to what `@org.jetbrains.annotations.Contract("!null -&gt; !null")` implies for Java?
## [5][Can't solve a problem because of ConcurrentModificationException](https://www.reddit.com/r/Kotlin/comments/jn91zm/cant_solve_a_problem_because_of/)
- url: https://www.reddit.com/r/Kotlin/comments/jn91zm/cant_solve_a_problem_because_of/
---
I think I found the solution for this problem, no idea how efficient it is, but it looks good. The only problem is that I have to change the list while iterating over its elements and it throws ConcurrentModificationException. I tried what I found on the internet, some iterator functions, but wherever I try to put it in relation to the for-loops it doesn't work. How can I avoid this error?

Here's the code. I need to find the longest non-descending sequence in an N elements list(array originally)

&amp;#x200B;

    import java.util.*
    
    //sample input
    //10                                           number of numbers
    //1 2 4 1 2 2 5 7 4 3
    //Sample Output 1:
    //5                              the longest non-descending sequence (1 2 2 5 7)
    
    fun main() {
        val scanner = Scanner(System.`in`)
        val arrayLength = scanner.nextInt()
        val array = MutableList(arrayLength) { scanner.nextInt() }
    
        val iterator = array.iterator()
        var longestNonDescendingSequence = 0
    
        while (iterator.hasNext()) {
            for (number in array) {
                val item = iterator.next()
                var lastNumOfSeq = number
                val nonDesSeq = mutableListOf&lt;Int&gt;()
    
                for (i in array) {
                    if (i &gt;= lastNumOfSeq) {
                        nonDesSeq.add(1)
                        lastNumOfSeq = i
                        if (item == lastNumOfSeq) {
                            iterator.remove()
                        }
                    } else {
                        break
                    }
                }
                if (nonDesSeq.size &gt; longestNonDescendingSequence) longestNonDescendingSequence = nonDesSeq.size
            }
        }
    
        println(longestNonDescendingSequence)
    }
    
    12
    1 2 4 1 2 3 5 7 8 9 10 11
    Exception in thread "main" java.util.ConcurrentModificationException
    	at java.base/java.util.ArrayList$Itr.checkForComodification(ArrayList.java:1013)
    	at java.base/java.util.ArrayList$Itr.next(ArrayList.java:967)
    	at MainKt.main(Main.kt:19)
    	at MainKt.main(Main.kt)
    
    Process finished with exit code 1

I probably should add the original code without me trying to solve this with an iterator

    fun main() {
        val scanner = Scanner(System.`in`)
        val arrayLength = scanner.nextInt()
        val array = MutableList(arrayLength) { scanner.nextInt() }
        var longestNonDescendingSequence = 0
    
        for (number in array) {
            var lastNumOfSeq = number
            val nonDesSeq = mutableListOf&lt;Int&gt;()
            for (i in array) {
                if (i &gt;= lastNumOfSeq) {
                    nonDesSeq.add(1)
                    lastNumOfSeq = i
                    array.removeAt(0)
                } else {
                    break
                }
            }
            if (nonDesSeq.size &gt; longestNonDescendingSequence) longestNonDescendingSequence = nonDesSeq.size
        }
    
    
        println(longestNonDescendingSequence)
    }

&amp;#x200B;
## [6][How to get a job as a Kotlin dev?](https://www.reddit.com/r/Kotlin/comments/jn6w45/how_to_get_a_job_as_a_kotlin_dev/)
- url: https://www.reddit.com/r/Kotlin/comments/jn6w45/how_to_get_a_job_as_a_kotlin_dev/
---
Hey Guys, let me explain this a bit why i ask this (tldr at the end). I'm a Dev with a finished industrial job training as a developer and working for around 6 years now for a company using xBase++ as programming language. The language itself comes from clipper, is quite old and is an OOP language. I have some minor basic Java skills, but never worked with toolkits or frameworks for testing, automating workflows or anything normal for java/kotlin devs, but i'm confident in learning a new programminglanguage and i'm interested in doing kotlin with android programming. For my current job i did frontend and backend development both with xBase++. I also never did any normal workflows used in the developerbranch like crunch, because we don't use this stuff at my current job.

TLDR: I have no Knowledge about Kotlin/Java and it's frameworks with minor basic knowledge about the language itself but i know how to develop some stuff.

So how can i get into a developer job with kotlin with my current knowledge? Or how would you guys do it?
## [7][Best course for beginners ?](https://www.reddit.com/r/Kotlin/comments/jn6suq/best_course_for_beginners/)
- url: https://www.reddit.com/r/Kotlin/comments/jn6suq/best_course_for_beginners/
---
What course / learning material for absolute beginners in Kotlin do you recommend ?

Preference: free
+ what IDE for kotlin is the best ? Intellij ?
## [8][Kotlin Cheat Sheet](https://www.reddit.com/r/Kotlin/comments/jmfjs6/kotlin_cheat_sheet/)
- url: https://i.redd.it/hm1fj7innqw51.jpg
---

## [9][Setting an audio file as phone's ringtone](https://www.reddit.com/r/Kotlin/comments/jn6h1u/setting_an_audio_file_as_phones_ringtone/)
- url: https://www.reddit.com/r/Kotlin/comments/jn6h1u/setting_an_audio_file_as_phones_ringtone/
---
I've tried so many code so far now and none of em seem to work. Maybe I'm doing something ,I don't know, I'm absolutely new to Kotlin. Can somebody please help me on this?
## [10][OpenCV](https://www.reddit.com/r/Kotlin/comments/jmu8b8/opencv/)
- url: https://www.reddit.com/r/Kotlin/comments/jmu8b8/opencv/
---
I am porting some code that I wrote in python into Kotlin so it will work on android. I am wondering if I can use OpenCV in Kotlin or do I need to use Java?
## [11][Facts You May Not Know About Kotlin](https://www.reddit.com/r/Kotlin/comments/jmn85w/facts_you_may_not_know_about_kotlin/)
- url: https://youtu.be/XwzStZaDpH0?list=PLEx5khR4g7PIiAEHCt6LGMFnzq7JjO8we
---

