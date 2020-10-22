# Kotlin
## [1][What are some libraries everyone should be using that would make life easier?](https://www.reddit.com/r/Kotlin/comments/jfk44r/what_are_some_libraries_everyone_should_be_using/)
- url: https://www.reddit.com/r/Kotlin/comments/jfk44r/what_are_some_libraries_everyone_should_be_using/
---

## [2][Unable to create a simple project following the Spring Guide tutorial](https://www.reddit.com/r/Kotlin/comments/jfyohl/unable_to_create_a_simple_project_following_the/)
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
## [3][trying Kotlin on a new project, but nothing is working, details inside](https://www.reddit.com/r/Kotlin/comments/jfy4o4/trying_kotlin_on_a_new_project_but_nothing_is/)
- url: https://www.reddit.com/r/Kotlin/comments/jfy4o4/trying_kotlin_on_a_new_project_but_nothing_is/
---
So I'm trying to make  anew Kotlin project but there's a problem, I can't even produce a Hello World.  Using intellij 2020.2.3 there appears to be two ways to create a Kotlin project (Either selecting Kotlin directly, or choosing Java and selecting Kotlin/JVM as an additional framework), and neither of them work out of the box, and all of the documentation I can find is seemingly out of date ([https://kotlinlang.org/docs/tutorials/jvm-get-started.html](https://kotlinlang.org/docs/tutorials/jvm-get-started.html) for example, the UI options it walks you through don't exist),

Once you successfully create a project, it fails to setup because the configuration doesn't work.  I had to update some gradle options to get the correct version of gradle for my version of the JVM. Thankfully I was able to find that solution on Stack Overflow because I would have never figured that out on my own or with the available documentation.  So it downloads the new gradle version, and the example main function created by the IDE has no green Run button in the gutter, and attempting to build the project yielded some 2000 lines of LLVM optimization errors.

Is there a better "hello world" walkthrough somewhere on the internet?  I'm not really finding a lot of help with all of this, and frankly this is kind of a nightmare compared to pretty much every other environment I've ever worked with.
## [4][Is the IntelliJ experience with Kotlin as good or better than Java in terms of IDE features?](https://www.reddit.com/r/Kotlin/comments/jfj46v/is_the_intellij_experience_with_kotlin_as_good_or/)
- url: https://www.reddit.com/r/Kotlin/comments/jfj46v/is_the_intellij_experience_with_kotlin_as_good_or/
---

## [5][Ktor create long running API JOB endpoint with status](https://www.reddit.com/r/Kotlin/comments/jfun3f/ktor_create_long_running_api_job_endpoint_with/)
- url: https://www.reddit.com/r/Kotlin/comments/jfun3f/ktor_create_long_running_api_job_endpoint_with/
---
I need to create an API endpoint that will create a long running JOB managed through a coroutine and return the use a JOB ID through which the user can poll for status. JOB normally lasts between 5 to 15 minutes and its OK if this JOB is not shared between instances.

Is there anything available already or i should write my own JOB queue class?
## [6][Should I learn Kotlin?](https://www.reddit.com/r/Kotlin/comments/jfmdjs/should_i_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/jfmdjs/should_i_learn_kotlin/
---
I’ve just finished my 1st Semester of Java on my way to a Computer Science degree.

Very impressed with Java and will continue to become proficient in it through further study.

Have also heard a lot about Kotlin and wondering if all the enthusiasm for it was before the latest changes in Java from Java 8 to Java 15?

I have a lot of other other languages to learn, plus the front end, so would like an honest opinion if I should just continue to do a deep dive with Java or also take in Kotlin.

Would appreciate if you could provide any reasons why, when comparing Kotlin to Java 15. Of relevance is that I am a fan of JetBrains and IntelliJ Idea and believe Kotlin was originally developed by JetBrains?

Thanks, very much appreciated!
## [7][Announcing mapneat - my second Kotlin project attempt](https://www.reddit.com/r/Kotlin/comments/jfk9w0/announcing_mapneat_my_second_kotlin_project/)
- url: https://www.reddit.com/r/Kotlin/comments/jfk9w0/announcing_mapneat_my_second_kotlin_project/
---
After writing a small project called **serverneat** a few months back, my adventure in kotlin-land continues with [https://github.com/nomemory/mapneat](https://github.com/nomemory/mapneat).

As the rest of us here, I simply love the language, and I cannot stop writing code in it. So my intention is to practice it in non-trivial hello world examples.

\-----

**MapNeat** is a Kotlin DSL / library that simplifies JSON to JSON, XML to JSON and POJO to JSON transformations.

It can be particularly useful when you plan to use it inside an integration layer that converts data from one format to another, creating DTOs, etc.

For the moment the library is not Maven Central, because it's not yet covered by tests or in a final state.

Any feedback is very much appreciated (in terms of features, making the API more friendly, etc.);

Code review would be also good - my Kotlin experience is limited, so I tend to still write Kotlin with a Java mindset.

Contributors always wanted.

Cheers!
## [8][Kotlin and Spring boot rest api with Google oauth](https://www.reddit.com/r/Kotlin/comments/jffwfp/kotlin_and_spring_boot_rest_api_with_google_oauth/)
- url: https://www.reddit.com/r/Kotlin/comments/jffwfp/kotlin_and_spring_boot_rest_api_with_google_oauth/
---
Is there any good article on implementing spring boot rest api with Google oauth in kotlin/java? I would like to use Google access token via Android app and verify details from spring boot server via rest api.
## [9][How to run Kotlin Multiplatform Web Application in continous mode](https://www.reddit.com/r/Kotlin/comments/jex6ej/how_to_run_kotlin_multiplatform_web_application/)
- url: https://www.reddit.com/r/Kotlin/comments/jex6ej/how_to_run_kotlin_multiplatform_web_application/
---
I started a project using the default Kotlin Multiplatform Full-Stack Web Application. (Ktor + Kotlin React with styled components)

Can I somehow run the entire application in continous mode, so that the website gets refreshed automatically whenever I change either the front end or the back end?

Edit: I use the gradle wrapper in version 6.7.
## [10][jOOQ 3.14 includes support for kotlin code generation, among other features](https://www.reddit.com/r/Kotlin/comments/jeortf/jooq_314_includes_support_for_kotlin_code/)
- url: https://blog.jooq.org/2020/10/20/jooq-3-14-released-with-sql-xml-and-sql-json-support/
---

