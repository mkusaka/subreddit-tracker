# Kotlin
## [1][After 3 years of Kotlin, I finally found an understandable explanation of "in" and "out"](https://www.reddit.com/r/Kotlin/comments/hyo565/after_3_years_of_kotlin_i_finally_found_an/)
- url: https://www.reddit.com/r/Kotlin/comments/hyo565/after_3_years_of_kotlin_i_finally_found_an/
---
Stumbled upon this article after yet again trying to understand "in" and "out" and how it correlates to Java. Might come in handy for someone else!  [https://kotlin.christmas/2019/22](https://kotlin.christmas/2019/22)
## [2][Can Retrofit be used in Kotlin?](https://www.reddit.com/r/Kotlin/comments/hyjxni/can_retrofit_be_used_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hyjxni/can_retrofit_be_used_in_kotlin/
---
There documentation shows everything in Java so I don't know if Retrofit will work if I adjust my code to Kotlin.
## [3][What is there to do with Kotlin other than Android development?](https://www.reddit.com/r/Kotlin/comments/hykkul/what_is_there_to_do_with_kotlin_other_than/)
- url: https://www.reddit.com/r/Kotlin/comments/hykkul/what_is_there_to_do_with_kotlin_other_than/
---
I'm interested in learning Kotlin (and most likely will regardless of if there is much to do with Kotlin, since I'm into Android Development), but would just like to know what else I could do with the language. From what I've heard Kotlin is easier than Java and since I'm also interested in learning Golang I thought instead of learning Java first, learn Kotlin, then move onto Golang or Java or any other Java-Based languages since there are so many at this point.
Thanks in advance!
## [4][Is There A Way To Create Operators?](https://www.reddit.com/r/Kotlin/comments/hyl89c/is_there_a_way_to_create_operators/)
- url: https://www.reddit.com/r/Kotlin/comments/hyl89c/is_there_a_way_to_create_operators/
---
We can currently overload an operator for a class but is there a way to create one? For example, maybe instead of using .xor() for exclusive or I can create an operator for it to use ^ for .xor() for that specific class. If not there should be a way to do this because having the ability to create syntactic sugar like this could result in more readable and concise code.
## [5][How to implement interface with onCheckChangeListener in kotlin with recyclerview?](https://www.reddit.com/r/Kotlin/comments/hyktr5/how_to_implement_interface_with/)
- url: https://www.reddit.com/r/Kotlin/comments/hyktr5/how_to_implement_interface_with/
---
This probably a noob question but I am wondering how one can create a interface method for onCheckChangeListener. If you want to implement onCheckChangeListener in recyclerview but call that implementation in your MainActivity, how would you forward onCheckChangeListener from recyclerview class to MainActivity?
## [6][Why does Kotlin have so many keywords or modifiers?](https://www.reddit.com/r/Kotlin/comments/hy43ej/why_does_kotlin_have_so_many_keywords_or_modifiers/)
- url: https://www.reddit.com/r/Kotlin/comments/hy43ej/why_does_kotlin_have_so_many_keywords_or_modifiers/
---
I generally (as a Java developer) find all of the keywords very confusing. There's data, enum, lateinit, open, typealias, and a lot more that I probably don't know. Why are there so many? Is there a good resource that lists them all?
## [7][Is There A Way To Define A Generic Number Parameter?](https://www.reddit.com/r/Kotlin/comments/hy44vy/is_there_a_way_to_define_a_generic_number/)
- url: https://www.reddit.com/r/Kotlin/comments/hy44vy/is_there_a_way_to_define_a_generic_number/
---
Let's say I want to make a generic function that works with all number types, how would I do that because the Number class can't really be used without manually checking and casting it to one of the types.
## [8][Why do data classes require at least one parameter in its constructor?](https://www.reddit.com/r/Kotlin/comments/hxm13d/why_do_data_classes_require_at_least_one/)
- url: https://www.reddit.com/r/Kotlin/comments/hxm13d/why_do_data_classes_require_at_least_one/
---
I'm trying to make an object to simply hold data. However, since the data is coming from HTTP requests, the object needs to be instantiated before any data can be added to it. Kotlin doesn't allow data classes to be created without an argument to its primary constructor, making this approach impossible without having some arbitrary value to pass to it.

Why are data classes like this? Why can't you have an empty data class?
## [9][Using reflection at all still considered bad practise?](https://www.reddit.com/r/Kotlin/comments/hxn6u3/using_reflection_at_all_still_considered_bad/)
- url: https://www.reddit.com/r/Kotlin/comments/hxn6u3/using_reflection_at_all_still_considered_bad/
---
I have the following method which takes a pointer to an object's property and sets its value:

`private fun fetchResource(resUri: String, resPointer: KMutableProperty0&lt;String&gt;): Promise&lt;String&gt; {`  
   `return window.fetch(resUri, params)`  
`.then {response -&gt; response.text()}`  
`.then {text -&gt; resPointer.set(text);text} // This line sets the property pointer's value`  
`.catch {e -&gt; console.error("Fetch for $resUri failed with $e"); e.message!! }`  
`}`

This is passed as a KMutableProperty0&lt;String&gt; variable. With my experience in Java, I know that reflection in general should be avoided as it breaks encapsulation, but this sort of logic is akin to something in C++. Is this sort of logic still considered bad practise?
## [10][Kotlin programming language: How Google is using it to squash the code bugs that cause most crashes](https://www.reddit.com/r/Kotlin/comments/hx6esj/kotlin_programming_language_how_google_is_using/)
- url: https://www.zdnet.com/article/google-were-using-kotlin-programming-language-to-squash-the-bugs-that-cause-most-crashes/
---

