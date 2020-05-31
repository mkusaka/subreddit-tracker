# Kotlin
## [1][When one assertions is not enough - Kotlin Testing](https://www.reddit.com/r/Kotlin/comments/gtxlci/when_one_assertions_is_not_enough_kotlin_testing/)
- url: https://kotlintesting.com/assert-softly-when-one-assertion-is-not-enough/
---

## [2][What is your preferred stack for building a RESTful web app in Kotlin?](https://www.reddit.com/r/Kotlin/comments/gtpg4s/what_is_your_preferred_stack_for_building_a/)
- url: https://www.reddit.com/r/Kotlin/comments/gtpg4s/what_is_your_preferred_stack_for_building_a/
---
I'm new to Kotlin and was wondering what people preferred to use. Especially interested in anything that supports fat jar for easy deployment.
## [3][Realm 7, the frozen throne](https://www.reddit.com/r/Kotlin/comments/gtigx3/realm_7_the_frozen_throne/)
- url: https://www.coroutinedispatcher.com/2020/05/realm-7-frozen-throne.html
---

## [4][TechEmpower released round 19 of their server middleware benchmark (including Kotlin ones)](https://www.reddit.com/r/Kotlin/comments/gsuvvo/techempower_released_round_19_of_their_server/)
- url: https://www.reddit.com/r/Kotlin/comments/gsuvvo/techempower_released_round_19_of_their_server/
---
TechEmpower just released their latest round (19) of their server middleware benchmark. You can check the results here: https://www.techempower.com/benchmarks
## [5][I Need Help Adding To Hashmap From a Text Input](https://www.reddit.com/r/Kotlin/comments/gt99bu/i_need_help_adding_to_hashmap_from_a_text_input/)
- url: https://www.reddit.com/r/Kotlin/comments/gt99bu/i_need_help_adding_to_hashmap_from_a_text_input/
---
Can anyone help me add functionality to my program to add items to my hashmap from a text input?

&amp;#x200B;

var products = *hashMapOf*&lt;Int, Pair&lt;String, Double&gt;&gt;(  
 111 *to* Pair("shoes", 59.99),  
 222 *to* Pair("shirt", 19.99),  
 333 *to* Pair("socks", 3.99)

*print*("Please Enter The Description Of The New Item: ")  


var addItemDescription = *readLine*()!!  


*print*("Please Enter The Price For $addItemDescription: ")  


var addItemPrice = *readLine*()!!.*toFloat*()
## [6][Codota's free ai code completion plugin now supports kotlin for android studio and eclipse.](https://www.reddit.com/r/Kotlin/comments/gt0u32/codotas_free_ai_code_completion_plugin_now/)
- url: https://www.reddit.com/r/Kotlin/comments/gt0u32/codotas_free_ai_code_completion_plugin_now/
---
Has anyone tried this plugin with Kotlin? I'm curious about your thoughts on it
## [7][Java interop: Kotlin gets tripped up by capitalization.](https://www.reddit.com/r/Kotlin/comments/gsw920/java_interop_kotlin_gets_tripped_up_by/)
- url: https://www.reddit.com/r/Kotlin/comments/gsw920/java_interop_kotlin_gets_tripped_up_by/
---
The java.net.URLClassLoader class provides a getURLs() method:

&gt; public URL[] getURLs()

&gt; Returns the search path of URLs for loading classes and resources. This includes the original list of URLs specified to the constructor, along with any URLs subsequently appended by the addURL() method.

And the Kotlin documentation claims:

&gt; Methods that follow the Java conventions for getters and setters (no-argument methods with names starting with get and single-argument methods with names starting with set) are represented as properties in Kotlin. Boolean accessor methods (where the name of the getter starts with is and the name of the setter starts with set) are represented as properties which have the same name as the getter method.

Well, all righty then, I should have a .URLs property I can use. But no:

    &gt;&gt;&gt; import java.net.URLClassLoader
    &gt;&gt;&gt; val cl = ClassLoader.getSystemClassLoader()
    &gt;&gt;&gt; cl is URLClassLoader
    res4: kotlin.Boolean = true
    &gt;&gt;&gt; (cl as URLClassLoader).URLs
    error: unresolved reference: URLs
    (cl as URLClassLoader).URLs
                           ^

However, the old reliable fallback works:

    &gt;&gt;&gt; (cl as URLClassLoader).getURLs()
    res6: kotlin.Array&lt;(out) java.net.URL!&gt;! = [Ljava.net.URL;@5b6ac24b
    
Of course, the rule is that .getFoo() maps to a .foo property. Note the downcasing; the property on the Kotlin side is called .foo and not .Foo. So maybe it’s called .uRLs? Close, but no cigar!

    &gt;&gt;&gt; (cl as URLClassLoader).uRLs
    error: unresolved reference: uRLs
    (cl as URLClassLoader).uRLs
                           ^
    &gt;&gt;&gt; (cl as URLClassLoader).urLs
    res8: kotlin.Array&lt;(out) java.net.URL!&gt;! = [Ljava.net.URL;@5dddd61f

Ugh. Maybe I’ll just stick to calling the Java-style getters and setters.
## [8][Kotlin, Eclipse and Executable JARs](https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/)
- url: https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/
---
Sooo, I just wanted to compile my Kotlin application into an executable JAR file. No big deal, right? I didn't do that too often when developing with Java, but how much easier than "Export to Executable JAR" can it get?

So that's what I do. And end up with an completely empty error message. The file still got created, but executing it does nothing. 

I'm thinking my IDE (Eclipse) doesn't work all that well with Kotlin - it's not exactly a great IDE after all, and I only use it because I'm used to it and too lazy to get into another one. But maybe, since Kotlin runs in the JVM, I can just export it as a Java Application. A quick test at least shows, that I can successfully run it as a Java Application from within the IDE.  
But same result: empty error message and a file that does nothing.

So like the good little programmer I am, I hop onto google. 3 pages filled with nothing but "use Gradle", only with the occasional "Maven works too" sprinkled in for flavor. Nothing else I could find.

&amp;#x200B;

So - is there really no way of creating an Executable JAR without using Gradle?
## [9][How To Prevent ERP Implementation Failure](https://www.reddit.com/r/Kotlin/comments/gtbbta/how_to_prevent_erp_implementation_failure/)
- url: https://www.odooimplementation.com/blog/erp_implementation_failure
---

## [10][Do I need an Android phone to learn Kotlin?](https://www.reddit.com/r/Kotlin/comments/gsuq6h/do_i_need_an_android_phone_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gsuq6h/do_i_need_an_android_phone_to_learn_kotlin/
---
So I'm interested in learning Kotlin but... I don't have an Android phone. I know I can learn the language just like any other language, but I'm curious if it's pointless to learn it if I can't test it out on an actual phone. Not sure how elaborate Android Studio is. 

Do I need to get an Android phone to learn Kotlin?
