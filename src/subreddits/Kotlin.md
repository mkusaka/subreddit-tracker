# Kotlin
## [1][Design Patterns and Principles in Kotlin | Notes from Head First Design Patterns](https://www.reddit.com/r/Kotlin/comments/j9m8oc/design_patterns_and_principles_in_kotlin_notes/)
- url: https://www.reddit.com/r/Kotlin/comments/j9m8oc/design_patterns_and_principles_in_kotlin_notes/
---
Here are my notes I made while reading the book **Head First Design Patterns**. It has a full and quality implementation of every design pattern in the book with notes on different patterns and principles. The examples are written in Kotlin.

The Head First Design Pattern is a great read if you are just starting to learn design patterns.

Check out the Github repo below and do star it if you find it useful: [https://github.com/Devansh-Maurya/Design-Patterns-And-Principles](https://github.com/Devansh-Maurya/Design-Patterns-And-Principles)

https://preview.redd.it/dirikolr5ms51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=109648c86dd47e25cbd863c5d9c919569ead1639
## [2][How to build Slackbot with Kotlin, Ktor &amp; deploy to Heroku](https://www.reddit.com/r/Kotlin/comments/j9ddk5/how_to_build_slackbot_with_kotlin_ktor_deploy_to/)
- url: https://www.reddit.com/r/Kotlin/comments/j9ddk5/how_to_build_slackbot_with_kotlin_ktor_deploy_to/
---
I wrote a Slackbot at work and I thought it might be helpful to make an article with a simple example of how I went about creating and deploying one. Hope this helps someone and always welcome feedback!

https://plusmobileapps.com/2020/10/09/ktor-slackbot-heroku.html
## [3][How to fix “e: Unable to compile C bridges” with Kotlin Multiplatform.](https://www.reddit.com/r/Kotlin/comments/j9n6lw/how_to_fix_e_unable_to_compile_c_bridges_with/)
- url: https://www.reddit.com/r/Kotlin/comments/j9n6lw/how_to_fix_e_unable_to_compile_c_bridges_with/
---
We’ve come across this problem several times, and for the first time, it was a bit problem to find what’s wrong.

https://preview.redd.it/pxhtz5hgkms51.png?width=1200&amp;format=png&amp;auto=webp&amp;s=939b8851b66d696b00825ce66fdba1ae95c3509d

Basically, it’s missing libncurses5which you can simply install on Ubuntu/Debian:

    apt-get install libncurses5

This problem usually hits us on CI/CD, where you are in a clean environment. The last time, it was with Github Actions after moving our runner to a new set of servers.

This one is just a quick tip that can save you nerves eventually 📷.

[https://localazy.com/blog/kotlin-mpp-unable-to-compile-c-bridges](https://localazy.com/blog/kotlin-mpp-unable-to-compile-c-bridges)
## [4][Why Kotlin is so popular for Android development and not general Java Development?](https://www.reddit.com/r/Kotlin/comments/j9qhcn/why_kotlin_is_so_popular_for_android_development/)
- url: https://www.reddit.com/r/Kotlin/comments/j9qhcn/why_kotlin_is_so_popular_for_android_development/
---
So I already know Java and working on it for 5 year. 

Now I am transining to Kotlin and all I hear is "Kotlin for Android". I know this is easy and kind of language which reduces the code base of app and main Null Pointer.

But why so much for Android and not anything else. I worked on Spring boot and want to convert whole Rest API in Kotlin. 

Let me know what is reason if there is any to become very good language for Android development and not for other kind of S/W and web development on Java.
## [5][Trying to upload an image. Correct way to get image file from uri?](https://www.reddit.com/r/Kotlin/comments/j9ljv8/trying_to_upload_an_image_correct_way_to_get/)
- url: https://www.reddit.com/r/Kotlin/comments/j9ljv8/trying_to_upload_an_image_correct_way_to_get/
---
Hi there, 

I'm a reasonably new Kotlin developer, having come from an iOS/web and sysadmin background. I'm looking for a little guidance as for the 'correct' way to create and upload an image a user has picked (all within my app, nothing being shared) . I've spent literally all day on this without getting a 100% working solution. I've tried a few hacky copypastas from SO, but none have been a good overall solution (I'm also using cached data from Glide, which I can get a uri for, but the hacky "get the real filepath" methods fall apart here)...

It seems that trying to get the 'real' filepath given a uri is not meant to be the best way to go about it, and I should be using a fileprovider / contentresolver (I have these set up in my manifest/xml dir). However, using a uri to get a file reference seems to be able to detect the type of file, but creating a file from a uri always results in a 0 size / and a java.io.FileNotFoundException: /external/images/media/26: open failed: ENOENT (No such file or directory) error when I execute the upload.

I'm using okhttp3 for networking, for a multipartform upload with multiple images... the code snippet in question goes like..

val image = File(uri.*path*)  
val imageType = *contentResolver*.getType(uri) //seems to work

val fileRequestBody = image.*asRequestBody*(imageType!!.*toMediaType*())  
requestBody.addFormDataPart("image\_$index", "image\_$index.$fileExtension", fileRequestBody)

This compiles and runs, up until I execute the request, where I get the file not found error. 

If anyone has had similar issues, or can point me in the right direction I would be very appreciative, as I'm getting seriously over it (after 8 hours+ of failure)
## [6][Kotlin Bootcamp for Programmers 3: Functions](https://www.reddit.com/r/Kotlin/comments/j8n5p1/kotlin_bootcamp_for_programmers_3_functions/)
- url: https://codelabs.developers.google.com/codelabs/kotlin-bootcamp-functions/#9
---

## [7][Having a hard time deciding between java and kotlin](https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/)
- url: https://www.reddit.com/r/Kotlin/comments/j90qnv/having_a_hard_time_deciding_between_java_and/
---
Well i still consider myself as a beginner even if I started 3 years (kept switching language look at my reddit profile it's a mess...) and I still do the same mistake, anyway some guy told me to learn kotlin and not java (well he got mad after he called me a thickheaded because I told him that kotlin is only for mobile dev lmao), because java had some 'flaws'
So any suggestion?.
Thanks.
## [8][A TornadoFx question](https://www.reddit.com/r/Kotlin/comments/j8vphl/a_tornadofx_question/)
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
## [9][Learning Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/j8kyy4/learning_kotlin_coroutines/)
- url: https://www.reddit.com/r/Kotlin/comments/j8kyy4/learning_kotlin_coroutines/
---
I am trying to learn Kotlin corountines through [Kotlin Corountines Official Documentation](https://kotlinlang.org/docs/reference/coroutines/basics.html) but I'm having hard time understanding what's going on.  
Can you suggest  other tutorials (with JVM as target and not android)?
## [10][AndroidBites: Awesome-Kotlin-Extensions [Resource]](https://www.reddit.com/r/Kotlin/comments/j91eaw/androidbites_awesomekotlinextensions_resource/)
- url: https://www.reddit.com/r/Kotlin/comments/j91eaw/androidbites_awesomekotlinextensions_resource/
---
Hi Guys, I have started an awesome series, where I plan to host a list of Kotlin extensions, Please do check it out and if you like to contribute anything form your arsenal then your always welcome!

Resource link 👉 [https://chetangupta.net/awesome-kotlin-extensions/](https://chetangupta.net/awesome-kotlin-extensions/)
