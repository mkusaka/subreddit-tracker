# Kotlin
## [1][gRPC with Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/ig51fl/grpc_with_kotlin_coroutines/)
- url: https://codingwithmohit.com/grpc/grpc-kotlin-coroutines/
---

## [2][I made a thing, I'm sure all Kotlin lovers will appreciate!](https://www.reddit.com/r/Kotlin/comments/iflmgn/i_made_a_thing_im_sure_all_kotlin_lovers_will/)
- url: https://i.redd.it/l5np0cey2xi51.jpg
---

## [3][Does anyone know an extensive and detailed guide to use scope functions? I haven't been able to find anything good online. Thanks a lot :]](https://www.reddit.com/r/Kotlin/comments/igaoqc/does_anyone_know_an_extensive_and_detailed_guide/)
- url: https://www.reddit.com/r/Kotlin/comments/igaoqc/does_anyone_know_an_extensive_and_detailed_guide/
---

## [4][A Kotlin Library for the Twitch API](https://www.reddit.com/r/Kotlin/comments/ifp51g/a_kotlin_library_for_the_twitch_api/)
- url: https://www.reddit.com/r/Kotlin/comments/ifp51g/a_kotlin_library_for_the_twitch_api/
---
Hello guys,

I recently implemented a [Kotlin Twitch Helix API wrapper](https://github.com/frozencure/twitch-client) for some of the other projects I’m working on and I decided to open-source it. I’m hoping some of you will use it and would also love to have some feedback on it.

It might be an interesting example to people that want to use the \`1.4.0-rc\` version of the Dokka documentation. You can find the documentation \[here\]([https://frozencure.github.io/twitch-client/twitch-client/](https://frozencure.github.io/twitch-client/twitch-client/))

Here are is some information about the library:

## About

[Twitch-Client](https://github.com/frozencure/twitch-client) is a modern Kotlin library that permits an easy and quick interaction with the Twitch services. Currently, the library can be used in *Java* or *Kotlin/JVM* projects and will be extended to other platforms in the future with the help of *Kotlin-Multiplatform*. The library is built on top of the [Ktor Client](https://ktor.io/clients/index.html) which is one of the best performing HTTP clients for Kotlin.

## Features

* All [Twitch Helix](https://dev.twitch.tv/docs/api) endpoints are implemented and tested
* Works with Kotlin/JVM and Java projects
* Works with Android Studio projects
* Authentication service included for an easy OAuth flow implementation
* Nullable types are used for all Twitch API response objects
* 100% unit test coverage
* All public methods and members are documented
* Documentation includes code samples for all endpoints
* Easy handling of multi-page Twitch collection responses
## [5][I'm trying to enable Kotlin Result monad in a multi-platform project (more specifically in the :shared module) so i tried doing it like in the image, but none of them has worked, anyone know how?](https://www.reddit.com/r/Kotlin/comments/ifw542/im_trying_to_enable_kotlin_result_monad_in_a/)
- url: https://i.redd.it/4b3h2tgw60j51.png
---

## [6][I decided to move from nodejs to kotlin, I need some help in choosing backend libs](https://www.reddit.com/r/Kotlin/comments/iffb3f/i_decided_to_move_from_nodejs_to_kotlin_i_need/)
- url: https://www.reddit.com/r/Kotlin/comments/iffb3f/i_decided_to_move_from_nodejs_to_kotlin_i_need/
---
I have very small exp with java/kotlin, all I tried in the past was a classic spring boot with its own jpa and so on  

Given that, I'm currently trying to choose the libraries to use for an asynchronous graphql server with a code first approach,  I'm looking for some tips

Is expedia/graphql-kotlin good or are there alternatives?  
Which web framework should I use? Spring or spark/other?  
I don't really understand how asynchronous database access works (I'm used to node where everything is async u know). It is not enabled by default while using hibernate/any jdbc right, what am I supposed to use?

Secondary: logging and validation libs suggestions

Thanks for the help

EDIT NON RELATED QUESTION: Is it fine using Apache Spark in kotlin or does it have any downside compared to java?
## [7][Problems with Hyperskill courses](https://www.reddit.com/r/Kotlin/comments/iflwru/problems_with_hyperskill_courses/)
- url: https://www.reddit.com/r/Kotlin/comments/iflwru/problems_with_hyperskill_courses/
---
Hi everyone!

I'm trying to do one of the hyperskill's kotlin courses, but it seems like there's some issues with their config and/or my jdk version.

I've selected one of the projects from the edutools and idea set it up and installed everything. When I press "check", though, I always get build errors.

I've tried different jdk settings with no luck (changing the gradle jvm seems to make no difference at all):

\- JDK 1.8 =&gt; Could not target platform: 'Java SE 11' using tool chain: 'JDK 8 (1.8)'

\- 14 =&gt; java.lang.noClassDefFoundError: Could not initialize class org.codehaus.groovy.vmplugin.VMPluginFactory

\- openjdk-14 =&gt; Could not initialize class org.codehaus.groovy.runtime.InvokerHelper

\- kotlin-sdk 1.4 =&gt; JAVA\_HOME is set to an invalid directory

&amp;#x200B;

If I create a kotlin project from scratch everything works fine (and it selects the openjdk-14 automatically).
## [8][Error deploying Ktor app to Heroku](https://www.reddit.com/r/Kotlin/comments/ifkvee/error_deploying_ktor_app_to_heroku/)
- url: https://www.reddit.com/r/Kotlin/comments/ifkvee/error_deploying_ktor_app_to_heroku/
---
Hello, i am struggling deploying ktor app to Heroku, i followed the documentation but no success. The problem is when i run 'git push heroku master' and i keep getting the same error(the error below). I have tried every solution i found on net but no success. Any suggestions? Btw. I have added a global variable ANDROID_SDK_ROOT and also the sdk.dir shows the right path of my sdk.

SDK location not found. Define location with an ANDROID_SDK_ROOT enviroment variabkr or by setting the sdk.dir in your project's local.properties file at 'tmp/build_b06419f6/local.properties'
## [9][Need help with timer](https://www.reddit.com/r/Kotlin/comments/if6fis/need_help_with_timer/)
- url: https://www.reddit.com/r/Kotlin/comments/if6fis/need_help_with_timer/
---
Hi, I’m completely new to kotlin. I want to make a for loop that wait's a second between every iteration, but I can’t figure out how to delay it. Any suggestions?
## [10][My solution for mandatory parameters in Kotlin DSL](https://www.reddit.com/r/Kotlin/comments/if2hgj/my_solution_for_mandatory_parameters_in_kotlin_dsl/)
- url: https://www.reddit.com/r/Kotlin/comments/if2hgj/my_solution_for_mandatory_parameters_in_kotlin_dsl/
---
Hey all,

I've been looking for a clean solution to mandatory parameters in Kotlin DSL.

Ended up writing a linter for it, tell me what you guys think

[https://blog.kotlin-academy.com/kotlin-dsl-know-your-limits-2deaef1bab66](https://blog.kotlin-academy.com/kotlin-dsl-know-your-limits-2deaef1bab66)
