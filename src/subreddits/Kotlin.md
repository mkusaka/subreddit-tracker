# Kotlin
## [1][For people like me who are into listening music while coding... [off]](https://www.reddit.com/r/Kotlin/comments/gdpgmf/for_people_like_me_who_are_into_listening_music/)
- url: https://www.reddit.com/r/Kotlin/comments/gdpgmf/for_people_like_me_who_are_into_listening_music/
---
... check out the [list](https://spoti.fi/3cPiDAs) I use when I’m coding: 8+ hours of retro synth music inspired by Stranger Things.

Also curious to know what kind of stuff do you listen to, if any.
## [2][(Advice) What after Head First Kotlin?](https://www.reddit.com/r/Kotlin/comments/gdvfzo/advice_what_after_head_first_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gdvfzo/advice_what_after_head_first_kotlin/
---
I just finished the book ' Head First Kotlin ' and have no prior java experience. However I do have some python and C++ experience.

*Question*  :  I am looking to get started in Android dev. Should I learn Java or just brush up on the abstraction Kotlin provides? Also please include some resources, if you have any, for the same.

Ps: books preferably

Thank you for your time.
## [3][Free Kotlin course for beginners by Donn Felker](https://www.reddit.com/r/Kotlin/comments/gdffkk/free_kotlin_course_for_beginners_by_donn_felker/)
- url: https://www.reddit.com/r/Kotlin/comments/gdffkk/free_kotlin_course_for_beginners_by_donn_felker/
---
Links:  
[https://twitter.com/donnfelker/status/1257341939109842944](https://twitter.com/donnfelker/status/1257341939109842944)

[https://caster.io/courses/kotlin-programming-language](https://caster.io/courses/kotlin-programming-language)
## [4][Kotlin’s Elvis better than Swift’s Guard](https://www.reddit.com/r/Kotlin/comments/gd9c4k/kotlins_elvis_better_than_swifts_guard/)
- url: https://medium.com/@elye.project/kotlins-elvis-better-than-swift-s-guard-53030d403c3f?source=friends_link&amp;sk=bf7fced446cc161095a9d281054fb927
---

## [5][Kotlin/Native can't use a class reference in a lambda passed to a class constructor or in a method of an object extending that class](https://www.reddit.com/r/Kotlin/comments/gdg1dd/kotlinnative_cant_use_a_class_reference_in_a/)
- url: https://www.reddit.com/r/Kotlin/comments/gdg1dd/kotlinnative_cant_use_a_class_reference_in_a/
---
I have an instance of a class [here](https://github.com/DeflatedPickle/ducknroll/blob/map-command/src/commonMain/kotlin/example/main.kt#L27), which if I throw `println(player)` under it, it'll rightly print the instance. However, either of these [calls](https://github.com/DeflatedPickle/ducknroll/commit/374f72de259d4606e20dc0369f9c2fcf72ca9cbb#diff-9f158b81a446ad38bfadaa041350dda7L37), both throw a non-sense error.

What's the error? Why `finished with non-zero exit value -1073741571`, of course!

Now, I can comment out that `print`, and it'll run fine. Until I run the [map command](https://github.com/DeflatedPickle/ducknroll/blob/map-command/src/commonMain/kotlin/com/deflatedpickle/ducknroll/common/command/MapCommand.kt#L20), which takes in the player, where (for testing), it `prints` out the executor. If, in the command, I take out the `print`, again, it works perfectly fine. I can pass around `player`, but not use it.

&amp;#x200B;

Now, to test if it was a scoping issue, I moved it into its own `ThreadLocal` tagged `object`, but the same error was thrown. Note: You could still `print` it fine outside of the lambda or extended class update method.

&amp;#x200B;

I've never encountered an error that's like this, where I can't use an object in a nested scope, even if it's static. Is it a problem with Kotlin/Native or am I being an idiot?

&amp;#x200B;

**Edit**: I don't have the ultimate version of IDEA, so I can't debug it (if anyone can, that'd be great!).

&amp;#x200B;

**Edit 2**: One of my friends has ultimate, they cloned the repo, downloaded all the LLVM stuff, ran it in debug with breakpoints and... got the same result. Debugging is no help here. Either in doing something wrong (which we spent awhile looking for) or it's a bug.
## [6][Kotlin: Sealed Classes (enum 2.0)](https://www.reddit.com/r/Kotlin/comments/gdqm5q/kotlin_sealed_classes_enum_20/)
- url: https://flutteryapps.com/blogs/kotlin/sealed-classes.html
---

## [7][Kotlin tutorial 01 for beginners - Introduction](https://www.reddit.com/r/Kotlin/comments/gdiw9z/kotlin_tutorial_01_for_beginners_introduction/)
- url: https://youtu.be/cQSRB_exCGM
---

## [8][The one and only object](https://www.reddit.com/r/Kotlin/comments/gdcw2h/the_one_and_only_object/)
- url: https://medium.com/androiddevelopers/the-one-and-only-object-5dfd2cf7ab9b
---

## [9][Create Retrofit CallAdapter for Coroutines to handle response as states](https://www.reddit.com/r/Kotlin/comments/gd8p84/create_retrofit_calladapter_for_coroutines_to/)
- url: https://medium.com/@melegy/create-retrofit-calladapter-for-coroutines-to-handle-response-as-states-c102440de37a
---

## [10][GraphQL API integration tests in a Spring Boot 2.x Kotlin application](https://www.reddit.com/r/Kotlin/comments/gd7ff5/graphql_api_integration_tests_in_a_spring_boot_2x/)
- url: https://www.reddit.com/r/Kotlin/comments/gd7ff5/graphql_api_integration_tests_in_a_spring_boot_2x/
---
In order to test end-to-end scenarios, the author wanted an easy way to…

* start the back-end exposing the GraphQL API endpoint
* define test GraphQL queries
* load the test GraphQL queries
* prepare &amp; send GraphQL HTTP queries

In this post, the author assumes that you’re already using graphql-java and exposing it using graphql-java-servlet or equivalent means: https://medium.com/dSebastien/graphql-api-integration-tests-in-a-spring-boot-2-x-kotlin-application-5840d3c5d66f?source=friends_link&amp;sk=9f22a7426e83f8ae02ace369174c2b62
