# Kotlin
## [1][I’m excited to announce that my open-source sample project about “Android Modular Architecture” just reach 1K ⭐️ on Github.](https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/)
- url: https://www.reddit.com/r/Kotlin/comments/hnydey/im_excited_to_announce_that_my_opensource_sample/
---
Maybe sounds insignificant if you compare with video reproductions or a post likes on twitter. But isn’t, the project is among top 10 github projects in Kotlin related with Android Architecture Sample with several mentions/references like Android Weekly, AndroidSweets, Droidcon and KotlinBy.

I want to say thank you ❤️ in this way for all the support received by the android/kotlin developer community and I hope that some of the architecture decisions taken help you to apply/inspire to your current or next project.

Of course, nothing is perfect and this project isn’t an exception, there are a lot of things to improve, but the iteration is the key to my point of view.

Project Link: [https://github.com/VMadalin/android-modular-architecture](https://github.com/VMadalin/android-modular-architecture)
## [2][Kotlin/Native desktop UI](https://www.reddit.com/r/Kotlin/comments/hnppf2/kotlinnative_desktop_ui/)
- url: https://www.reddit.com/r/Kotlin/comments/hnppf2/kotlinnative_desktop_ui/
---
Does anyone know how to get simple UI for the desktop (mainly Linux and Windows) with Kotlin/Native? 

I tried the [libui](https://github.com/msink/kotlin-libui) which has nice DSL, sadly it's more of a prototype. Secondly, the cinterop of gtk3 library, which works but as of today it takes long time to compile (with Kotlin 1.4-M3) and the code is all but not nice / simple.
## [3][Netty (Jooby/Ktor) will leak memory significantly on a simple hello world. What could the culprit be?](https://www.reddit.com/r/Kotlin/comments/hnvaga/netty_joobyktor_will_leak_memory_significantly_on/)
- url: https://www.reddit.com/r/Kotlin/comments/hnvaga/netty_joobyktor_will_leak_memory_significantly_on/
---
The following app - with an empty conf, will (if hit intensely via wrk) jumps up to over half a gig of memory. The same app in c#/f# will stay solid around 78mb. Python through gunicorn will stay solid around 40mb (per worker).

It seems like something is going nuts in netty (running this through ktor runs into the same problem), but I can't seem to find a good discussion of why this would be a problem. Any ideas? Code below:


    import io.jooby.runApp

    fun main(args: Array&lt;String&gt;) {
        runApp(args) {
            get ("/hello") {
                "hi"
            }
        }
    }


Dependencies in gradle:

    dependencies {
        compile group: 'org.jetbrains.kotlin', name: 'kotlin-stdlib', version:'1.3.72'
        compile group: 'io.jooby', name: 'jooby-netty', version:'2.8.9'
        compile group: 'ch.qos.logback', name: 'logback-classic', version:'1.2.3'
    }

wrk command:

    wrk -d60s -t6 -c100 http://localhost:8000/hello
## [4][String to operation](https://www.reddit.com/r/Kotlin/comments/hnmxhk/string_to_operation/)
- url: https://www.reddit.com/r/Kotlin/comments/hnmxhk/string_to_operation/
---
How can I turn a string such as “5*8+3-1” to an operation 5*8+3-1 which would equal 42.
## [5][7 JVM arguments of Highly Effective Applications](https://www.reddit.com/r/Kotlin/comments/hnae5p/7_jvm_arguments_of_highly_effective_applications/)
- url: https://blog.gceasy.io/2020/03/18/7-jvm-arguments-of-highly-effective-applications/
---

## [6][How to run application forever?](https://www.reddit.com/r/Kotlin/comments/hnkwp4/how_to_run_application_forever/)
- url: https://www.reddit.com/r/Kotlin/comments/hnkwp4/how_to_run_application_forever/
---
Maybe a dumb question but I come from web development world and I have written nodejs application which I run using forever package:

https://github.com/foreversd/forever 

So whenever the application crashes it restart automatically. Is there something the same for Kotlin?
## [7][Tee is a Kotlin micro-library (just 2kb jar) to help debug computational flows [disclaimer: core contributor]](https://www.reddit.com/r/Kotlin/comments/hmytjr/tee_is_a_kotlin_microlibrary_just_2kb_jar_to_help/)
- url: https://github.com/medly/ktee
---

## [8][[Question] What code coverage tool to use for the javascript code in a Kotlin Multiplatform project?](https://www.reddit.com/r/Kotlin/comments/hn6dq6/question_what_code_coverage_tool_to_use_for_the/)
- url: https://www.reddit.com/r/Kotlin/comments/hn6dq6/question_what_code_coverage_tool_to_use_for_the/
---

## [9][Ktor re-authentication after inactivity](https://www.reddit.com/r/Kotlin/comments/hn97c8/ktor_reauthentication_after_inactivity/)
- url: https://www.reddit.com/r/Kotlin/comments/hn97c8/ktor_reauthentication_after_inactivity/
---
I have just started my first large project using Kotlin and plan to run the backend of my application on an embedded Netty server in Ktor.  I am trying to setup user authentication so that it will expire after a certain period of inactivity, as the application will run on a few shared computers that many users will access.  If possible in Ktor, is there some property I can change on my server to enable this or do I have to find a way to manually implement it?
## [10][How to know when to use class and keyword object](https://www.reddit.com/r/Kotlin/comments/hmz76m/how_to_know_when_to_use_class_and_keyword_object/)
- url: https://www.reddit.com/r/Kotlin/comments/hmz76m/how_to_know_when_to_use_class_and_keyword_object/
---
Ever since I first learned about the keyword object I've been making everything an object when I don't see a need to make more than one of something ( I seem to have trouble thinking in terms of objects, so I rarely ever use class ). But I'm now wondering if I should even be declaring a group of functions as an object. I've tried searching how to know when to make something an object vs just making a file with top level functions, and the views seem to be split on the issue ( the views seem to be split between: object should be used, object should almost never be used, and it doesn't really matter which is used ). I have my latest project up on GitHub, I've made two versions of it, one with using the keyword [object](https://github.com/PhenixFine/smart_calculator/tree/master/src) and one [without](https://github.com/PhenixFine/smart_calculator/tree/master/src/alt), and I was wondering which version is better ( I should note that other students were able to solve the project with 2/3 less code using regex, but I haven't learned how to use that yet ). Here's a brief example of my PostFix as an object for those that don't want to look at my GitHub files:

    object Postfix {
        private val STACK = Stack&lt;String&gt;()
        private val POSTFIX = mutableListOf&lt;String&gt;()
        ::more values::
        
        fun convert(infix: String): Array&lt;String&gt; { code }
    
        private fun reset() { code }
        ::more private functions::
    }

And this is it adjusted as just a file with top level functions:

    private val STACK = Stack&lt;String&gt;()
    private val POSTFIX = mutableListOf&lt;String&gt;()
    ::more values::
    
    fun infixToPostfix(infix: String): Array&lt;String&gt; { code } // convert was renamed to be more concise, since it is no longer being called with the object name.
    
    private fun reset() { code }
    ::more private functions::

Also, for the purpose of the class I'm in, the main file has to be named Main and I was wondering if that is common practice to name the main file that? ( as with my alt version it's not descriptive of what it contains, like with my version with object use, the main file just makes a call to the SmartCalculator object, where the alt version has the guts of it within the main file ).
