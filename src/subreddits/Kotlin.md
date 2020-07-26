# Kotlin
## [1][Why does Kotlin have so many keywords or modifiers?](https://www.reddit.com/r/Kotlin/comments/hy43ej/why_does_kotlin_have_so_many_keywords_or_modifiers/)
- url: https://www.reddit.com/r/Kotlin/comments/hy43ej/why_does_kotlin_have_so_many_keywords_or_modifiers/
---
I generally (as a Java developer) find all of the keywords very confusing. There's data, enum, lateinit, open, typealias, and a lot more that I probably don't know. Why are there so many? Is there a good resource that lists them all?
## [2][Is There A Way To Define A Generic Number Parameter?](https://www.reddit.com/r/Kotlin/comments/hy44vy/is_there_a_way_to_define_a_generic_number/)
- url: https://www.reddit.com/r/Kotlin/comments/hy44vy/is_there_a_way_to_define_a_generic_number/
---
Let's say I want to make a generic function that works with all number types, how would I do that because the Number class can't really be used without manually checking and casting it to one of the types.
## [3][Why do data classes require at least one parameter in its constructor?](https://www.reddit.com/r/Kotlin/comments/hxm13d/why_do_data_classes_require_at_least_one/)
- url: https://www.reddit.com/r/Kotlin/comments/hxm13d/why_do_data_classes_require_at_least_one/
---
I'm trying to make an object to simply hold data. However, since the data is coming from HTTP requests, the object needs to be instantiated before any data can be added to it. Kotlin doesn't allow data classes to be created without an argument to its primary constructor, making this approach impossible without having some arbitrary value to pass to it.

Why are data classes like this? Why can't you have an empty data class?
## [4][Using reflection at all still considered bad practise?](https://www.reddit.com/r/Kotlin/comments/hxn6u3/using_reflection_at_all_still_considered_bad/)
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
## [5][Kotlin programming language: How Google is using it to squash the code bugs that cause most crashes](https://www.reddit.com/r/Kotlin/comments/hx6esj/kotlin_programming_language_how_google_is_using/)
- url: https://www.zdnet.com/article/google-were-using-kotlin-programming-language-to-squash-the-bugs-that-cause-most-crashes/
---

## [6][val me : Java-&gt;Kotlin](https://www.reddit.com/r/Kotlin/comments/hx4evi/val_me_javakotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hx4evi/val_me_javakotlin/
---
I'm not 100 percent sure this is the place for this but it might be useful for other java devs looking at this sub to see whether it's worth making the switch.

I recently enrolled in Svetlana Isakova @ JetBrains' excellent coursera course "Kotlin for Java developers" I found it supremely informative and though I have not finished yet, I was motivated enough to complete 4.5 of the five weeks of material in just a few days. While I am used to Java and personally do not like using some other languages like python, I could immediately see the great advantages that Kotlin offers. I have decided to list some of my top 5 reasons to use Kotlin so that others may be swayed.

1. Java interoperability. Kotlin can be used with java code in a project directly, can be seamlessly converted to from java and can use any Java library.

2.:Runs on JVM, Android, JavaScript(!) and native executables on certain platforms. You can use it everywhere you can use java and beyond.

3.:It's clean and concise. Seriously, it is just so much more compact. On occasions I've been able to condense 20 lines of java into 4 lines of Kotlin. I think that speaks for itself.

4.: Functional And Oop. I think Kotlin strikes a great balance between a pure functional style as seen e.g. in Haskell and a classic Object oriented approach like I was familiar with from Java and c#. Depending on your preferences, you can use both styles. I think this aids a great deal in transitioning to the language.

5. Easy to learn from a Java background. As I have already said, I went from knowing literally nothing in Kotlin to a rather respectable level in just a few days. Transposing and transferring prior knowledge from Java to Kotlin works seamlessly and if you forget how to do something in Kotlin, you can just write it in java and either use the java code directly or convert it to Kotlin code.

Tl:Dr just learn kotlin
I hope this helps someone

Edit: HOW did I forget about nullability? Of course! I suppose it's just something that wasn't on my mind when I wrote this.

6. Nullability. Do I need to say more?
## [7][Proposal: Community Content Creation](https://www.reddit.com/r/Kotlin/comments/hx7p4y/proposal_community_content_creation/)
- url: https://www.reddit.com/r/Kotlin/comments/hx7p4y/proposal_community_content_creation/
---
Personally, I come here to learn and ask questions when needed. But it would be nice if we can have some active content creation on a regular schedule by the members of the community from different levels of experience.

This can get a little bit messy if people are just posting their courses for sale or linking to their business, basically not following rule #2: no spamming.

But as long as the content is accessible by all (a reddit post, your blog post or a youtube video, etc.) and high quality, I think it can make the community more lively and therefore provide a richer experience for current and future members.

What do you guys think?
## [8][It would be great if we could get some active moderation here.](https://www.reddit.com/r/Kotlin/comments/hww5d8/it_would_be_great_if_we_could_get_some_active/)
- url: https://www.reddit.com/r/Kotlin/comments/hww5d8/it_would_be_great_if_we_could_get_some_active/
---
This place kinda feels like a desert. No weekly threads, no apparent moderation. Half of the questions posted are about the Android SDK and have 0 relevancy to Kotlin.

Not trying to diss the moderators but I haven't seen a sign of life from any of them.
## [9][Options for managing derived attributes](https://www.reddit.com/r/Kotlin/comments/hx5kgb/options_for_managing_derived_attributes/)
- url: https://blog.frankel.ch/options-manage-derived-attributes-kotlin/
---

## [10][I need (technical) help with the first assignment of the "Kotlin for Java developers" course on Coursera](https://www.reddit.com/r/Kotlin/comments/hx49k2/i_need_technical_help_with_the_first_assignment/)
- url: https://www.reddit.com/r/Kotlin/comments/hx49k2/i_need_technical_help_with_the_first_assignment/
---
I'm new to Kotlin but familiar with Java, so I decided to start this course by jetbrains, but when installing the eduTools plugin and selecting the assignment of week 2, nothing showed up on the course view of the project, and on the regular view I saw only the gradle packages. Did anyone else encounter this problem and how do I fix it?
