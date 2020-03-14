# Kotlin
## [1][Kotlin newbie](https://www.reddit.com/r/Kotlin/comments/fiicgz/kotlin_newbie/)
- url: https://www.reddit.com/r/Kotlin/comments/fiicgz/kotlin_newbie/
---
Hi! I've started to learn the Kotlin language two weeks ago because I want to develop my first app. At first, I jumped right into coding a login start screen and include Firebase, but I don't understand some things and it became frustrating. After that, I've considered that it's best to start with the basics and try to structure the info in lessons. Usually, the best learning method for me is to take a quiz/test after finishing a "lesson". Can you recomend websites or apps where I can do this?
## [2][Smart Casts via Assertions + Kotlin Contracts](https://www.reddit.com/r/Kotlin/comments/fi1dac/smart_casts_via_assertions_kotlin_contracts/)
- url: https://medium.com/@akperkins1/smart-casts-via-assertions-kotlin-contracts-fd737a92013b
---

## [3][Do we have to master java before we can dominate Kotlin in 2020?](https://www.reddit.com/r/Kotlin/comments/fhsigx/do_we_have_to_master_java_before_we_can_dominate/)
- url: https://www.reddit.com/r/Kotlin/comments/fhsigx/do_we_have_to_master_java_before_we_can_dominate/
---
I want to be an android developer and  I'm interested in the language of Kotlin. Do I need to specialize in java before kotlin to make my career easier in 2020 and future? And what are the features an android developer should have in our time? What are the factors that make it easier to find a job?
## [4][Kotlin compilation fails in IntelliJ with KotlinSDK selected, but not with JavaSDK?](https://www.reddit.com/r/Kotlin/comments/fhnosw/kotlin_compilation_fails_in_intellij_with/)
- url: https://www.reddit.com/r/Kotlin/comments/fhnosw/kotlin_compilation_fails_in_intellij_with/
---
I set up a simple Kotlin/Gradle project in IntelliJ, and when I select java8-openjdk oder java11-openjdk as the project SDK, then everything compiles and works fine. However, if I select Kotlin SDK (1.3.61), it fails with the following output: 

`22:09:22: Executing tasks ':classes :testClasses'...`

`22:09:22: Tasks execution finished ':classes :testClasses'.`

Why is this or should this even work? How can I fix this? Thanks!
## [5][How to use the onRestoreInstanceState Method with TextView string values](https://www.reddit.com/r/Kotlin/comments/fhyjq2/how_to_use_the_onrestoreinstancestate_method_with/)
- url: https://www.reddit.com/r/Kotlin/comments/fhyjq2/how_to_use_the_onrestoreinstancestate_method_with/
---
Hi dudes  


Im not a professionnal of Kotlin, but I'm working in a big project for a quite simple app using Android Studio.  


My app is working as I wished but I have an issue to save my datas when I switch of activity.  
I think its link to my Textview values that I always need to toString/toInt for treat them, but Idk how to unblock me.  


Please take a look to my code and tell me how and where create that onRestoreInstanceState method. Maybe my code is not as professional than yours but please don't mind.  


Tks a lot for the help I'm sure ur gonna do for me, heres the code :  


  


[First part](https://preview.redd.it/5omgqnakjfm41.png?width=1001&amp;format=png&amp;auto=webp&amp;s=0cf909ddae37fbcf4a50879355d41159543edb4f)

&amp;#x200B;

[Second part](https://preview.redd.it/i5h2s9hnjfm41.png?width=1368&amp;format=png&amp;auto=webp&amp;s=5d36d3dee6e7fd8bca797bb797cf38116e62fa1e)

  
PS :  tell me too if my onSaveInstanceState method is correct or not, tks :)
## [6][why does it show error when trying to print enum variables without using toString() method?](https://www.reddit.com/r/Kotlin/comments/fhuolf/why_does_it_show_error_when_trying_to_print_enum/)
- url: https://www.reddit.com/r/Kotlin/comments/fhuolf/why_does_it_show_error_when_trying_to_print_enum/
---
    fun main(args: Array&lt;String&gt;) {
        val mobile1: Mobile = Mobile("IPhone",MobileColor.GOLD)
        val mobile2: Mobile = Mobile("SONY",MobileColor.BLACK)
        println("The color of my "+mobile1.name+" is "+mobile1.color)
        println("The color of my "+mobile2.name+" is "+mobile2.color)
        
    //error code block start
        println(mobile1.color.toString() + " value is "+mobile1.color.value)
        //no error   
        
        println(mobile1.color + " value is "+mobile1.color.value) 
        // the mobile1.color shows error if used without toString() method
    
        println("${mobile1.color} value is "+mobile1.color.value)
        // no error when string template is used
    
    //error code block end
    }
    data class Mobile(val name:String, val color: MobileColor)
    enum class MobileColor(val value: Int) {
        GOLD(0xffd323),
        SILVER(0xeaeaea),
        WHITE(0xffffff),
        BLACK(0x000000),
        RED(0xFF0000)
    }
## [7][Newcomer Tutorial: Conditionals in Kotlin - Deal with if and when](https://www.reddit.com/r/Kotlin/comments/fhqkhg/newcomer_tutorial_conditionals_in_kotlin_deal/)
- url: https://marcuseisele.com/pages/learningKotlin/kotlin-conditionals
---

## [8][Kotlin in Top 20 - Redmonk Programming languages](https://www.reddit.com/r/Kotlin/comments/fhb8lp/kotlin_in_top_20_redmonk_programming_languages/)
- url: https://redmonk.com/sogrady/2020/02/28/language-rankings-1-20/
---

## [9][Help with Goocle Cloud Platform tutorial](https://www.reddit.com/r/Kotlin/comments/fhevqs/help_with_goocle_cloud_platform_tutorial/)
- url: https://www.reddit.com/r/Kotlin/comments/fhevqs/help_with_goocle_cloud_platform_tutorial/
---
Hi everyone,

 I finished the [Build a Kotling Spring Application with Google Cloud Platform](https://codelabs.developers.google.com/codelabs/cloud-spring-cloud-gcp-kotlin/index.html?index=..%2F..index#0) tutorial, where we create an application that accepts registrants information, publishes it to a Cloud Pub/Sub topic,  and persists this to a Cloud MySQL database. 

It is all fine but now I would like to remove registrants from the database and I don't know how to proceed. Either from the all registrants page or creating a new page that takes an input like last name to remove the registrant.

Can anyone help me?
## [10][The state of kotlin-ee](https://www.reddit.com/r/Kotlin/comments/fhfmx8/the_state_of_kotlinee/)
- url: https://www.reddit.com/r/Kotlin/comments/fhfmx8/the_state_of_kotlinee/
---
Hello everyone!  


i've spent a few days prospecting on how kotlin would behave on enterprise java development, with clear separation between app server resources and application resources, jndi context and so on.  


Heavily inspired by [this article](https://vorozco.com/blog/2019/2019-04-17-Evaluating-Kotlin-caveats-JavaEE.html), i've built [this sample project](https://github.com/sombriks/sample-jee-kotlin) and i'm looking for feedback and also to hear other experiences like this one.

Has anyone had better luck with other IDEs hather intellij ultimate?

Any suggestions on what the showcase should cover?

Please share your toughts.
