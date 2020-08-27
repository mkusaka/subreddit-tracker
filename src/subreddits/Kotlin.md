# Kotlin
## [1][Introducing Kotlin for Apache Spark Preview – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/ihkw3n/introducing_kotlin_for_apache_spark_preview/)
- url: https://blog.jetbrains.com/kotlin/2020/08/introducing-kotlin-for-apache-spark-preview/
---

## [2][Infixation for non-extension functions](https://www.reddit.com/r/Kotlin/comments/ihgl6t/infixation_for_nonextension_functions/)
- url: https://www.reddit.com/r/Kotlin/comments/ihgl6t/infixation_for_nonextension_functions/
---
Quick question: Is there a specific reason why infixation is only available for member/extension functions or is it just an arbitrary design choice?
## [3][Passing Java classes as arguments to annotation](https://www.reddit.com/r/Kotlin/comments/ihc6eq/passing_java_classes_as_arguments_to_annotation/)
- url: https://www.reddit.com/r/Kotlin/comments/ihc6eq/passing_java_classes_as_arguments_to_annotation/
---
I tried to create a test suite with Kotlin that runs Java test classes. As far as I see, I followed the Android guides to a T: https://developer.android.com/training/testing/unit-testing/instrumented-unit-tests#test-suites

But for my @ Suite annotation, I am getting the error "An annotation argument must be a compile-time constant" and "Unresolved reference" errors for the test classes even though they are being imported.

`@Suite.SuiteClasses( TestA::class, TestB::class, TestC::class, TestD::class )`

Is this because they are Java files and they can't be converted to KClass like this? How can I pass in Java files?
## [4][Shrinking a Kotlin binary by 99.2%](https://www.reddit.com/r/Kotlin/comments/igrt6a/shrinking_a_kotlin_binary_by_992/)
- url: https://jakewharton.com/shrinking-a-kotlin-binary/
---

## [5][How hard would it be to switch from Java to Kotlin?](https://www.reddit.com/r/Kotlin/comments/ihbo3i/how_hard_would_it_be_to_switch_from_java_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ihbo3i/how_hard_would_it_be_to_switch_from_java_to_kotlin/
---
I'm currently learning Java because my college uses it, and I'm wondering if it would be easy to switch to Kotlin later? 

Any answer would be appreciated!
## [6][How to convert inputStream into string](https://www.reddit.com/r/Kotlin/comments/ih7b50/how_to_convert_inputstream_into_string/)
- url: https://www.reddit.com/r/Kotlin/comments/ih7b50/how_to_convert_inputstream_into_string/
---
Currently: `val str : String =` [`client.inputStream.read`](https://client.inputStream.read)`().toString();`

I get ASCII values one by one, I tried with readBytes but no luck. What's the simplest solution to this?
## [7][Can someone tell me how to NOT get this runtime error over and over again? I tried using "%" all over the place:](https://www.reddit.com/r/Kotlin/comments/ih35eb/can_someone_tell_me_how_to_not_get_this_runtime/)
- url: https://www.reddit.com/r/Kotlin/comments/ih35eb/can_someone_tell_me_how_to_not_get_this_runtime/
---
 

A right rotation is an operation which shifts each element of the array to the right.

**For example**, if a right rotation is 1 and array is {1,2,3,4,5}, the new array will be {5,1,2,3,4}.  
**Another example**, if a right rotation is 2 and array {1,2,3,4,5}, the new array will be {4,5,1,2,3}, because  
{1,2,3,4,5} -&gt; {5,1,2,3,4} -&gt; {4,5,1,2,3}.

The first line of the input contains the number of elements in the array. The second line is the elements of the array. The third one is the number of right shifts.

The output contains all the shifted elements of the array. Separate the elements by the space character.

Report a typo

**Sample Input 1:**

5 1 2 3 4 5 1

**Sample Output 1:**

5 1 2 3 4

**Sample Input 2:**

5 1 2 3 4 5 2

**Sample Output 2:**

4 5 1 2 3 

============  
import java.util.\*   
fun main(args: Array&lt;String&gt;) {  
val scanner = Scanner(System.\`in\`)    
val size = scanner.nextInt()   
val a = IntArray(size) { scanner.nextInt() }   
val rotate = scanner.nextInt()    
val b = IntArray(size)    
var counter = 0          
repeat(size) { var key = ((rotate + counter) % size)              
counter += 1              
b.set(a\[key\], key)          
}    println(b.joinToString(" "))  
}   
=================

**Wrong, but keep on trying and never give up!** [**Install an IDE** ](https://hyperskill.org/learn/step/4698#)to get access to powerful debugging tools which let you examine your solution step by step. **1 runtime errorFEEDBACK**Failed test #1 of 9. Runtime error  This is a sample test from the problem statement!  Test input: 5 1 2 3 4 5 1 Correct output: 5 1 2 3 4  Your code output:   Error: Exception in thread "main" java.lang.ArrayIndexOutOfBoundsException: Index 5 out of bounds for length 5 	at MainKt.main(main.kt:15) 	at java.base/jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method) 	at java.base/jdk.internal.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62) 	at java.base/jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43) 	at java.base/java.lang.reflect.Method.invoke(Method.java:566) 	at org.jetbrains.kotlin.runner.AbstractRunner.run(runners.kt:61) 	at org.jetbrains.kotlin.runner.Main.run(Main.kt:109) 	at org.jetbrains.kotlin.runner.Main.main(Main.kt:119)
## [8][[Article] Snippets | List to String with Examples | Join operations and advance use cases](https://www.reddit.com/r/Kotlin/comments/igxzfd/article_snippets_list_to_string_with_examples/)
- url: https://chetangupta.net/list-to-string
---

## [9]["Wrap" an abstract method](https://www.reddit.com/r/Kotlin/comments/igway6/wrap_an_abstract_method/)
- url: https://www.reddit.com/r/Kotlin/comments/igway6/wrap_an_abstract_method/
---
I'm sure there is probably a design pattern or standard way of doing this but I don't know what it is, so rather than invent my own, I thought I'd ask!

I have a 3rd party library which has an abstract class like this

`abstract class Job{ abstract fun run() }`

Now, I want to automatically log the start and the end of all jobs into a file / table / whatever, so what I want to do is have the person writing the job not need to worry about it - they just implement the run method and the logging gets added automatically.

I could do this by having a class "LoggedJob", which implements run and has a differently named abstract method

    abstract class LoggedJob{
      override fun run(){
        log.start()
        myRun()
        log.end()
      }
      abstract fun myRun()
    }

but is there no way for me to keep the method name "run" and add logging automatically - basically wrap the abstract method with some extra functionality (open a transaction, write to a log, whatever).

I realise AOP could probably do this but it seems overkill for such a simple requirement - is there an idiomatic way of doing this?
## [10][Writing J2ME app in Kotlin](https://www.reddit.com/r/Kotlin/comments/igyt7k/writing_j2me_app_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/igyt7k/writing_j2me_app_in_kotlin/
---
I'm interested in developing a Java ME app using LWUIT. I've developed Java server-side and client-side for Android, but never used ME. Recently, I started working with Kotlin and found it saves a great deal of effort, is less verbose, and the resulting code is easier to read. 

Are there guides or examples of a J2ME app developed using Kotlin or Kotlin and Java? Alternatively, are there reasons why this might not be a great idea? I'm aware that Kotlin JAR files will be larger, and not all features (i.e. Reflection) are available.
