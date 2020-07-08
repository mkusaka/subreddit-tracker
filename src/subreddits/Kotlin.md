# Kotlin
## [1][7 JVM arguments of Highly Effective Applications](https://www.reddit.com/r/Kotlin/comments/hnae5p/7_jvm_arguments_of_highly_effective_applications/)
- url: https://blog.gceasy.io/2020/03/18/7-jvm-arguments-of-highly-effective-applications/
---

## [2][Tee is a Kotlin micro-library (just 2kb jar) to help debug computational flows [disclaimer: core contributor]](https://www.reddit.com/r/Kotlin/comments/hmytjr/tee_is_a_kotlin_microlibrary_just_2kb_jar_to_help/)
- url: https://github.com/medly/ktee
---

## [3][[Question] What code coverage tool to use for the javascript code in a Kotlin Multiplatform project?](https://www.reddit.com/r/Kotlin/comments/hn6dq6/question_what_code_coverage_tool_to_use_for_the/)
- url: https://www.reddit.com/r/Kotlin/comments/hn6dq6/question_what_code_coverage_tool_to_use_for_the/
---

## [4][Ktor re-authentication after inactivity](https://www.reddit.com/r/Kotlin/comments/hn97c8/ktor_reauthentication_after_inactivity/)
- url: https://www.reddit.com/r/Kotlin/comments/hn97c8/ktor_reauthentication_after_inactivity/
---
I have just started my first large project using Kotlin and plan to run the backend of my application on an embedded Netty server in Ktor.  I am trying to setup user authentication so that it will expire after a certain period of inactivity, as the application will run on a few shared computers that many users will access.  If possible in Ktor, is there some property I can change on my server to enable this or do I have to find a way to manually implement it?
## [5][How to know when to use class and keyword object](https://www.reddit.com/r/Kotlin/comments/hmz76m/how_to_know_when_to_use_class_and_keyword_object/)
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
## [6][vscode autoimport](https://www.reddit.com/r/Kotlin/comments/hmre0s/vscode_autoimport/)
- url: https://www.reddit.com/r/Kotlin/comments/hmre0s/vscode_autoimport/
---
Привет! Do you know some ways, to use autoimport in vscode? Or any external tools? Thanks in advance!

in autoimport - I mean, that vscode automatilally import needed packages.

P.S: I use neovim plugin, and I can additionaly use vim's plugins
## [7][Anyone here using Ktor and / or Exposed in production ?](https://www.reddit.com/r/Kotlin/comments/hmo9d4/anyone_here_using_ktor_and_or_exposed_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hmo9d4/anyone_here_using_ktor_and_or_exposed_in/
---
I know Exposed is 0.x software, but I was wondering if it or Ktor was in production use by any of you guys.

I'm currently using both for a professional project for a client.

Would like to hear your experiences.
## [8][Gradle Plugins and Composite Builds - Using Kotlin DSL](https://www.reddit.com/r/Kotlin/comments/hmvkiy/gradle_plugins_and_composite_builds_using_kotlin/)
- url: https://medium.com/@cortinico/gradle-plugins-and-composite-builds-962350501c38
---

## [9][Anyone mutation-testing Kotlin projects with Pitest?](https://www.reddit.com/r/Kotlin/comments/hmz4s3/anyone_mutationtesting_kotlin_projects_with_pitest/)
- url: https://www.reddit.com/r/Kotlin/comments/hmz4s3/anyone_mutationtesting_kotlin_projects_with_pitest/
---
How valuable you find integrating [pitest](https://pitest.org/) on a mid  to large project (microservice or library)?  

Does it give lot of false positives/negatives? 

Does it  take too long to finish?

Do you run it on  CI for every commit or nightly?

stack to add **PiTest** to **⇒** **Kotlin (JVM) + Micronaut + Kotest + Gradle**
## [10][Kotlin 1.4-M3 is Out: Standard Library Changes – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/hmaqpa/kotlin_14m3_is_out_standard_library_changes/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-m3-is-out-standard-library-changes/?utm_source=feedburner&amp;utm_medium=feed&amp;utm_campaign=Feed%3A+kotlin+%28Project+Kotlin%29
---

