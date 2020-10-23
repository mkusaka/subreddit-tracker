# Kotlin
## [1][Choosing the right architecture for a [new] Kotlin Multiplatform, Jetpack Compose and SwiftUI app](https://www.reddit.com/r/Kotlin/comments/jgllev/choosing_the_right_architecture_for_a_new_kotlin/)
- url: https://www.marcogomiero.com/posts/2020/kmm-shared-app-architecture/
---

## [2][Kotlin Team AMA session is coming!](https://www.reddit.com/r/Kotlin/comments/jgmvgz/kotlin_team_ama_session_is_coming/)
- url: https://www.reddit.com/r/Kotlin/comments/jgmvgz/kotlin_team_ama_session_is_coming/
---
Hi! The Kotlin team are getting ready to answer your questions on October 26-27. 

During this AMA, we'll also cover some of the unanswered questions from the [Kotlin 1.4 Online Event](https://kotlinlang.org/lp/event-14/). 

Keep an eye on the notifications in this subreddit, we will start at 9:00 am CET on October 26.
## [3][Kotlin JVM Gradle plugin documentation](https://www.reddit.com/r/Kotlin/comments/jgbscp/kotlin_jvm_gradle_plugin_documentation/)
- url: https://www.reddit.com/r/Kotlin/comments/jgbscp/kotlin_jvm_gradle_plugin_documentation/
---
Suppose I wanted to know what properties are available on the Kotlin object introduced by the Kotlin JVM Gradle plugin. From the [Using Gradle reference](https://kotlinlang.org/docs/reference/using-gradle.html) I can see that kotlin.incremental is a valid property, but is there a real reference of all properties and methods of the kotlin JVM plugin (preferably w/o having to use reflection to find them)?
## [4][Object and Companion object In Kotlin (For beginners)](https://www.reddit.com/r/Kotlin/comments/jg5pq9/object_and_companion_object_in_kotlin_for/)
- url: https://rylexr2678.hashnode.dev/object-and-companion-object-in-kotlin-1
---

## [5][Can we fix the documentation layout? Specifically the standard lib section?](https://www.reddit.com/r/Kotlin/comments/jg3p65/can_we_fix_the_documentation_layout_specifically/)
- url: https://www.reddit.com/r/Kotlin/comments/jg3p65/can_we_fix_the_documentation_layout_specifically/
---
The layout of the STD LIB docs is so shitty it hurts.  There are no collapsable sections once you enter into a package.  Why can't it be laid out like the JavaDocs? Trying to find something on the page requires you to scroll half a mile to find the sub-heading you're even looking for.  Entire sections should be collapsable at least on the subheading
## [6][trying Kotlin on a new project, but nothing is working, details inside](https://www.reddit.com/r/Kotlin/comments/jfy4o4/trying_kotlin_on_a_new_project_but_nothing_is/)
- url: https://www.reddit.com/r/Kotlin/comments/jfy4o4/trying_kotlin_on_a_new_project_but_nothing_is/
---
I don't know how I bungled it up so badly this morning, but I've managed to get it working.  I suspect I didn't actually select Kotlin and I selected Java with Kotlin as an additional framework... and man I don't know I must have just needed coffee or something.  Thanks for chiming in, I'm finally in business now!

~~So I'm trying to make  anew Kotlin project but there's a problem, I can't even produce a Hello World.  Using intellij 2020.2.3 there appears to be two ways to create a Kotlin project (Either selecting Kotlin directly, or choosing Java and selecting Kotlin/JVM as an additional framework), and neither of them work out of the box, and all of the documentation I can find is seemingly out of date (~~[~~https://kotlinlang.org/docs/tutorials/jvm-get-started.html~~](https://kotlinlang.org/docs/tutorials/jvm-get-started.html) ~~for example, the UI options it walks you through don't exist),~~

~~Once you successfully create a project, it fails to setup because the configuration doesn't work.  I had to update some gradle options to get the correct version of gradle for my version of the JVM. Thankfully I was able to find that solution on Stack Overflow because I would have never figured that out on my own or with the available documentation.  So it downloads the new gradle version, and the example main function created by the IDE has no green Run button in the gutter, and attempting to build the project yielded some 2000 lines of LLVM optimization errors.~~

~~Is there a better "hello world" walkthrough somewhere on the internet?  I'm not really finding a lot of help with all of this, and frankly this is kind of a nightmare compared to pretty much every other environment I've ever worked with.~~
## [7][Does anyone have any experience with Kotlin and Leaflet JS?](https://www.reddit.com/r/Kotlin/comments/jg3ebc/does_anyone_have_any_experience_with_kotlin_and/)
- url: https://www.reddit.com/r/Kotlin/comments/jg3ebc/does_anyone_have_any_experience_with_kotlin_and/
---
I'd like to create a web and mobile app that uses an interactive [choropleth map](http://en.wikipedia.org/wiki/Choropleth_map).  Since the release of Kotlin Multiplatform, I was thinking about trying to use [https://leafletjs.com/](https://leafletjs.com/), but I'm not sure how well it will work.  Any advice is much appreciated!
## [8][VCS with Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/jg3lf2/vcs_with_kotlin_multiplatform/)
- url: https://www.reddit.com/r/Kotlin/comments/jg3lf2/vcs_with_kotlin_multiplatform/
---
I just wanted to ask how someone should structure ist Multiplatform Project for Git? With the new Plugin for Android Studio a big Project is created, containing iOS and Android App and Multiplatform code. I think this isn't ideal but isn't this somehow related to the build process and where the framework file is put and how we link the Multiplatform code to the android project? Is it necessary?
## [9][Unable to create a simple project following the Spring Guide tutorial](https://www.reddit.com/r/Kotlin/comments/jfyohl/unable_to_create_a_simple_project_following_the/)
- url: https://www.reddit.com/r/Kotlin/comments/jfyohl/unable_to_create_a_simple_project_following_the/
---
This is the tutorial in Java: [https://spring.io/guides/gs/validating-form-input/](https://spring.io/guides/gs/validating-form-input/)

I cloned the above repo and ran in local and it is working fine. 

I'm trying to create the same with Kotlin. Project Link: [https://github.com/imranpasha-tech/Kotlin-springweb-thymeleaf-example](https://github.com/imranpasha-tech/Kotlin-springweb-thymeleaf-example)

This project is also working except it is unable to catch errors and forward them to view. 

So, when I try to insert Null values against the allowed values it is supposed to show the validation errors on the form. But that is being redirected to 404

&amp;#x200B;

[Entering empty\/null values here](https://preview.redd.it/5rjxojcc3nu51.png?width=466&amp;format=png&amp;auto=webp&amp;s=99d82212f5a0e0593cd7e683766975118f8daf6c)

This empty form should be validated as shown in the tutorial. However, it is not working as expected. Instead it is being redirected to:

&amp;#x200B;

[Not supposed to happen](https://preview.redd.it/gistdkak3nu51.png?width=671&amp;format=png&amp;auto=webp&amp;s=82c2ab0282ca84ba7db12fd6e96f19e71eb6e871)

&amp;#x200B;

Origin tutorial screenshots: 

&amp;#x200B;

[from original tutorial in java working fine. ](https://preview.redd.it/575k45iq3nu51.png?width=507&amp;format=png&amp;auto=webp&amp;s=6fd5869033034606ec7c84d55c0cea304aa6f4f9)
## [10][Just started BNR’s Book on Kotlin](https://www.reddit.com/r/Kotlin/comments/jg4k0l/just_started_bnrs_book_on_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jg4k0l/just_started_bnrs_book_on_kotlin/
---
I’ve been having a hard time staying focused and disciplined trying to learn Kotlin. I’ve read the first two chapters I want to read at least a chapter a day and spend some time with the challenges and the read more section. Has anyone completed the book? The book is just ok Kotlin I want to try to get into Android Development but since Kotlin has sooooo many different uses I will be experimenting. If anyone has read the book did you like it and did you feel comfortable with Kotlin and your ability to problem solve after?
