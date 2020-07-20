# Kotlin
## [1][[X-POST] JetBrains is porting Compose to desktop and other platforms : androiddev](https://www.reddit.com/r/Kotlin/comments/huidyh/xpost_jetbrains_is_porting_compose_to_desktop_and/)
- url: https://www.reddit.com/r/androiddev/comments/hufnjq/_/
---

## [2][Android app with the Room database](https://www.reddit.com/r/Kotlin/comments/hug5wx/android_app_with_the_room_database/)
- url: https://youtu.be/kF0TPehhIx0
---

## [3][Tips for newbies](https://www.reddit.com/r/Kotlin/comments/hukd3y/tips_for_newbies/)
- url: https://www.reddit.com/r/Kotlin/comments/hukd3y/tips_for_newbies/
---
                Hello people!
 
                I have not any experience at programming. I started to learn Java. But java seems to me difficult to learn. Then I tried Kotlin. It is looking not easy. But not hard as Java.

                I am watching Udemy videos. When I print "Hello World!" I was so excited. I said "That's magic!". I know it is a veeeeerryy basic thing but i liked that.

                My goal is coding budget app. There is so much and so good apps in Google Play but i want to build my own app.

                What do you advise for starters who know nothing?
## [4][Artifacts failing to build for my Kotlin program, getting a frustrating "main class not found" error at compile time.](https://www.reddit.com/r/Kotlin/comments/hu1gq7/artifacts_failing_to_build_for_my_kotlin_program/)
- url: https://www.reddit.com/r/Kotlin/comments/hu1gq7/artifacts_failing_to_build_for_my_kotlin_program/
---
I've moved a Kotlin project of mine into a new repository, which has necessitated re-configuring IntelliJ to use a new project. I had an artifact output which was a fat JAR containing all my dependencies, however this move has caused an extremely frustrating issue in that the JAR will build, but it won't run. Looking around in my manifest file, the IDE is moaning that my Main-Class attribute is invalid, but I can't figure out why. 

&gt;Manifest-Version: 1.0

&gt;Main-Class: uniFessSvr.MainKt

&gt;Class-Path: h2-1.4.199.jar websocket-api-9.3.6.v20151106.jar [...]

Any ideas what might be causing this?
## [5][Kotlin Actors – No Drama Concurrency](https://www.reddit.com/r/Kotlin/comments/htlgjt/kotlin_actors_no_drama_concurrency/)
- url: https://handstandsam.com/2020/07/18/kotlin-actors-no-drama-concurrency/
---

## [6][How to update mutable state in a gRPC client app?](https://www.reddit.com/r/Kotlin/comments/htu9af/how_to_update_mutable_state_in_a_grpc_client_app/)
- url: https://www.reddit.com/r/Kotlin/comments/htu9af/how_to_update_mutable_state_in_a_grpc_client_app/
---
This is a design question about the optimal/recommended way of maintaining shared mutable state using Kotlin concurrency primitives. For that, I need to establish the context first, please bear with me.  
I’ve a Spring Boot web application that has a gRPC client. Following are the requirements/constraints on the system:

1. Messages from the gRPC server could be informational or *commands*.
2. The client has internal state, that may change depending on the commands received from the server. Informational messages don’t affect client state.
3. Based on the commands received from the server, the client may generate more messages for the server, or update its internal state, or do both.
4. Every client message generated due to a server command *must happen after* any state changes.

My current design is as follows;

1. The client is a Spring bean. It has a reference to its internal state, which is a POJO (or POKO, if you want)
2. Upon receiving a server message, it constructs a command instance and passes the message content and the client state to the command. This happens in a gRPC streaming receiver, which runs on a gRPC thread pool (default).
3. The command instance may update the client state, and returns one or messages to the client.
4. The client, running within an Unconfined dispatcher, puts the messages in a buffered Channel, and waits for the next server message.
5. A Channel consumer, running on Kotlin default dispatcher pool, picks up the messages and sends them to the server.

Notice that requirement #4 is satisfied by design steps #3 and #5.

Questions:

1. I keep hearing that Flow is recommended over Channel. In this case though, the Channel  
 needs to stay open throughout the lifecycle of the client, and is very easy to use as an async queue. What benefits, if any, are there for using a Flow in this case?
2. I don’t feel particularly good about passing the client state to the commands. Is there a different way of achieving requirement #4 without doing this? (actors, may be)

*Disclaimer: Also posted on kotlin forum, but haven't received any replies. People tend to avoid long posts.*
## [7][📣 Calling all Kotlin/JS developers! 📣](https://www.reddit.com/r/Kotlin/comments/hth3e9/calling_all_kotlinjs_developers/)
- url: https://twitter.com/kotlin/status/1284114694702346240?s=19
---

## [8][Firebase competitors.](https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/)
- url: https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/
---
I'm interested in building my first Android app using Kotlin.  Base on my research thus far, Firebase seem to be a good solution to help minimize some of the development start up work.

--

I was wondering if there are alternatives / competitors to Firebase that I should do some research on before committing to Firebase.

--

Thanks.
## [9][Android Kotlin RecyclerView Tutorial - In-Depth Guide incl. Different View Types](https://www.reddit.com/r/Kotlin/comments/hth7c2/android_kotlin_recyclerview_tutorial_indepth/)
- url: https://www.reddit.com/r/Kotlin/comments/hth7c2/android_kotlin_recyclerview_tutorial_indepth/
---
I watched this video. It is amazing it is really very depth knowledge as I am familiar with android development but if any beginner wants to learn android it is really amazing. I loved it.

Video Link: [https://youtu.be/oDfl-xLXiac](https://youtu.be/oDfl-xLXiac)
## [10][Java Slack SDK 1.1.0 includes a Kotlin DSL for building messages](https://www.reddit.com/r/Kotlin/comments/hswsr4/java_slack_sdk_110_includes_a_kotlin_dsl_for/)
- url: https://slack.dev/java-slack-sdk/guides/composing-messages#block-kit-kotlin-dsl
---

