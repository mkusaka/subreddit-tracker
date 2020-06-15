# Kotlin
## [1][Building a Reactive Oauth Client App with SpringBoot and Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/h96ypw/building_a_reactive_oauth_client_app_with/)
- url: https://www.shiveenp.com/posts/spring-boot-reactive-oauth-client-with-coroutines/
---

## [2][Learning Kotlin - Android Studio](https://www.reddit.com/r/Kotlin/comments/h9bu9j/learning_kotlin_android_studio/)
- url: https://www.reddit.com/r/Kotlin/comments/h9bu9j/learning_kotlin_android_studio/
---
Guys, I'm working on app which have different aspects, feeds, navigation, sharing.. but I trying to find a good source to learn Kotlin and Android development. I need a course that teach me in exact. how to build for ex "Main activity" what is each part of it means? class, appcompat()..etc, who to interupt button? navigation drawer? link? I need to learn more about the methods used in real life situation, on how to transfer data between activities and how to handle motions and charts.

I already looked watched two courses from "Linkedin Learning" I watched few youtube lessons, I watched kotlin lessons, I watch how to build an app lesson. I downloaded open source apps, but they either showing me basic tutorials, or just how to use android studio tutorials, or a general tutorial for kotlin, which very confusing when I'm trying to apply it in Android studio. and regarding the open source apps, I've downloaded only one apps, if someone can guide me for bigger library for open source apps I will be thankful. it kindof hard to search in github.

&amp;#x200B;

btw, I have around 3 years experience in web development, I've studies java on windows, [visualbasic.net](https://visualbasic.net), php, swift.
## [3][Android Memory Dump Analyzer](https://www.reddit.com/r/Kotlin/comments/h9bbph/android_memory_dump_analyzer/)
- url: https://heaphero.io/
---

## [4][Trying to get access to errors in Ktor JWTAuth custom challenge](https://www.reddit.com/r/Kotlin/comments/h9b83y/trying_to_get_access_to_errors_in_ktor_jwtauth/)
- url: https://www.reddit.com/r/Kotlin/comments/h9b83y/trying_to_get_access_to_errors_in_ktor_jwtauth/
---
Hey everyone!

I'd like to be able to respond differently to different JWTAuth errors, so have customised the challenge (below is some code I've been using to test). Unfortunately the context is returning an empty list for the errors. I've tested with an expired token and with no token at all.

```
fun JWTAuthenticationProvider.Configuration.customConfigure() {
    verifier(verifier)
    realm = ISSUER
    challenge { _, _ -&gt;
      call.respond(HttpStatusCode.Unauthorized, JSONObject(mapOf("err" to context.authentication.allErrors)))
    }
    validate { 
      if (it.payload.audience.contains(AUDIENCE)) {
        it.payload.getClaim("id").asString().let { id -&gt; userDao.getUserById(id) }
      } else null
    }
  }
```

I was wondering if I was perhaps missing something. Hoping someone can help me!

Link to Stackoverflow post: [https://stackoverflow.com/q/62377411/6374965](https://stackoverflow.com/q/62377411/6374965)
## [5][Destructuring to existing vars](https://www.reddit.com/r/Kotlin/comments/h90p8r/destructuring_to_existing_vars/)
- url: https://www.reddit.com/r/Kotlin/comments/h90p8r/destructuring_to_existing_vars/
---
I was just fooling around with destructuring of data classes, and I noticed that you _can_ do

    var (a, b) = returnsDataClassWith2Members()
 
But if the `var`s `a` and `b` are already declared, you _cannot_ do

    (a, b) = returnsDataClassWith2Members()

Although it would be pretty useful. Am I just doing it wrong or is it generally not possible? And if it's not possible, why? And is it probable that this feature might be added in a future version?

I'm just curious, any help is appreciated. Thanks!
## [6][What does '?' after className means?](https://www.reddit.com/r/Kotlin/comments/h8xa7k/what_does_after_classname_means/)
- url: https://www.reddit.com/r/Kotlin/comments/h8xa7k/what_does_after_classname_means/
---
what does `?` means here `fun onCreate(container: ViewGroup?): View? { ... }`? There are two `?` here.
## [7][Why nonlocal return isn't allowed in noinline function? Please, explain me](https://www.reddit.com/r/Kotlin/comments/h8puxc/why_nonlocal_return_isnt_allowed_in_noinline/)
- url: https://www.reddit.com/r/Kotlin/comments/h8puxc/why_nonlocal_return_isnt_allowed_in_noinline/
---
Now I am reading "Kotlin in Action" and this book describes about nonlocal return. But I don't understand why nonlocal return isn't allowed in noinline func. I reread some sentences about it a lot of times. Give me a simple example why it isn't allowed
## [8][Is Kotlin commonly used with Spring?](https://www.reddit.com/r/Kotlin/comments/h8egj6/is_kotlin_commonly_used_with_spring/)
- url: https://www.reddit.com/r/Kotlin/comments/h8egj6/is_kotlin_commonly_used_with_spring/
---
Hi guys, I've known about Kotlin existence for some time but just recently I decided to give it a proper go and I really like some aspects of it. Mostly how it deals with collections and functional aspects such as HOFs and I'm sure there is more I just don't know about yet. I also know that it can be used with Spring, but when I was trying to find content about it I couldn't find much stuff. So my question is: is it because it works fairly similar to Java and there isn't much stuff specific to Kotlin or because it is rarely used with Spring for some reasons I am not aware of?
## [9][How to pass userinput data from another activity to recyclerview?](https://www.reddit.com/r/Kotlin/comments/h8j8gm/how_to_pass_userinput_data_from_another_activity/)
- url: https://www.reddit.com/r/Kotlin/comments/h8j8gm/how_to_pass_userinput_data_from_another_activity/
---
I am having a hard time getting my user input to show in the recyclerview. I have two activities. The Note activity has 2 edit text and a button and the MainActivity has my recyclerview. I am trying to pass the first activity input to the other and place that in the recyclerview. I also have a data class which just stores the two input data. Here are my codes.

Recyclerview adapter class :  [https://pastebin.com/4jBgS2Q5](https://pastebin.com/4jBgS2Q5)

MainActivity class:  [https://pastebin.com/yhCBwcD3](https://pastebin.com/yhCBwcD3)// main activity has my recyclerview

NoteActivity class:  [https://pastebin.com/kMqFzX29](https://pastebin.com/kMqFzX29)   // this class takes user input

NoteItem:  [https://pastebin.com/8WKEimNJ](https://pastebin.com/8WKEimNJ)   // data I want to pass into recyclerview.

&amp;#x200B;

Edit: I fix this issue. In my OnactivityResult after I added  my note to the list,  I needed to let my adapter know something happen by calling noteAdapter.notifyDataSetChanged(). noteAdapter is the name of my adapter btw. So you need to use the name you assign to your adapter.

&amp;#x200B;

&amp;#x200B;
## [10][The KorGE (Kotlin game engine) Gamejam is now live](https://www.reddit.com/r/Kotlin/comments/h7ixxb/the_korge_kotlin_game_engine_gamejam_is_now_live/)
- url: https://blog.korge.org/2020/06/korge-gamejam-1-theme.html
---

