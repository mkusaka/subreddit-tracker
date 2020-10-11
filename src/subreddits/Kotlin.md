# Kotlin
## [1][AndroidBites: Awesome-Kotlin-Extensions [Resource]](https://www.reddit.com/r/Kotlin/comments/j91eaw/androidbites_awesomekotlinextensions_resource/)
- url: https://www.reddit.com/r/Kotlin/comments/j91eaw/androidbites_awesomekotlinextensions_resource/
---
Hi Guys, I have started an awesome series, where I plan to host a list of Kotlin extensions, Please do check it out and if you like to contribute anything form your arsenal then your always welcome!

Resource link 👉 [https://github.com/ch8n/awesome-kotlin-extensions.git](https://github.com/ch8n/awesome-kotlin-extensions.git)
## [2][Kotlin Bootcamp for Programmers 3: Functions](https://www.reddit.com/r/Kotlin/comments/j8n5p1/kotlin_bootcamp_for_programmers_3_functions/)
- url: https://codelabs.developers.google.com/codelabs/kotlin-bootcamp-functions/#9
---

## [3][Having a hard time deciding between java and kotlin](https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/)
- url: https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/
---
Well i still consider myself as a beginner even if I started 3 years (kept switching language look at my reddit profile it's a mess...) and I still do the same mistake, anyway some guy told me to learn kotlin and not java (well he got mad after he called me a thickheaded because I told him that kotlin is only for mobile dev lmao), because java had some 'flaws'
So any suggestion?.
Thanks.
## [4][A TornadoFx question](https://www.reddit.com/r/Kotlin/comments/j8vphl/a_tornadofx_question/)
- url: https://www.reddit.com/r/Kotlin/comments/j8vphl/a_tornadofx_question/
---
Hey guys,

I have a question about the usage of tableview in TornadoFx. I have this simple code:

main.kt

    class Person (val name:String, val age:Int) {}
    
    fun main () {
        val listOfPeople = listOf&lt;Person&gt; (
            Person("Jhon",34),
            Person("Mary",30),
            Person("Jack",32)
    }

MainApp class:

    import tornadofx.*
    
    class MainApp :App(MainView::class)

MainView class:

    import tornadofx.*
    
    class MainView :View() {
        override val root = vbox {  
            tableview {
            } 
        }
    }

Now my quetion is: how do I populate the tableview in my MainView with the list in the main function? How, in general, can I refer to elements of the main module in the Views in TornadoFx? And also do I need to update the tableview or other components fields when properties change at runtime?

I'm a little bit of a newbie with Tornado, so I apologize if this is a naive question, but still I can't figure out how to make a referral to these elements.

Thanks in advance!
## [5][Learning Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/j8kyy4/learning_kotlin_coroutines/)
- url: https://www.reddit.com/r/Kotlin/comments/j8kyy4/learning_kotlin_coroutines/
---
I am trying to learn Kotlin corountines through [Kotlin Corountines Official Documentation](https://kotlinlang.org/docs/reference/coroutines/basics.html) but I'm having hard time understanding what's going on.  
Can you suggest  other tutorials (with JVM as target and not android)?
## [6][How to serialize an array of different types of value with `kotlinx.serialization`.](https://www.reddit.com/r/Kotlin/comments/j8gjt8/how_to_serialize_an_array_of_different_types_of/)
- url: https://www.reddit.com/r/Kotlin/comments/j8gjt8/how_to_serialize_an_array_of_different_types_of/
---
I try the following code.

\`\`\` kotlin

import kotlinx.serialization.json.Json

Json.encodeToString(arrayOf("gid", arrayOf("status", "c", "d"), 3))

\`\`\`

&amp;#x200B;

It gives the error message.

&amp;#x200B;

\`\`\`

Exception in thread "main" kotlinx.serialization.SerializationException: Class 'String' is not registered for polymorphic serialization in the scope of 'Serializable'.

Mark the base class as 'sealed' or register the serializer explicitly.

\`\`\`
## [7][SpringBoot alternative for Identity Server 4](https://www.reddit.com/r/Kotlin/comments/j8hwla/springboot_alternative_for_identity_server_4/)
- url: https://www.reddit.com/r/Kotlin/comments/j8hwla/springboot_alternative_for_identity_server_4/
---
Hi, is there something like Asp.Net Identity Server 4 for SpringBoot?
## [8][What is most advanced web framework/libraries in kotlin?](https://www.reddit.com/r/Kotlin/comments/j8gy4u/what_is_most_advanced_web_frameworklibraries_in/)
- url: https://www.reddit.com/r/Kotlin/comments/j8gy4u/what_is_most_advanced_web_frameworklibraries_in/
---
I am looking for something that has a lot of things working out-of-the-box. Something like symphony in PHP, where you have many ready to work solutions and you just adjust them. Is there something like this in kotlin? What is important for me is quickly building front-end and back-end so it will be more like building blocks. Spring is too complex, i want something easy. Ktor is very good but doesnt have solutions out of the box, you must create everything by yourself.
## [9][MDC Weekly Blend — Edition #9 - with loads of Android and Kotlin curated articles.](https://www.reddit.com/r/Kotlin/comments/j8icak/mdc_weekly_blend_edition_9_with_loads_of_android/)
- url: https://mobiledeveloperscafe.substack.com/p/mdc-weekly-blend-edition-9
---

## [10][New Release Cadence for Kotlin and the IntelliJ Kotlin Plugin – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/j81868/new_release_cadence_for_kotlin_and_the_intellij/)
- url: https://blog.jetbrains.com/kotlin/2020/10/new-release-cadence-for-kotlin-and-the-intellij-kotlin-plugin/
---

