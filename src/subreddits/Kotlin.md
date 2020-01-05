# Kotlin
## [1][KotlinConf 2019 Videos](https://www.reddit.com/r/Kotlin/comments/ebd0np/kotlinconf_2019_videos/)
- url: https://www.youtube.com/playlist?list=PLQ176FUIyIUY6SKGl3Cj9yeYibBuRr3Hl
---

## [2][FXGL (Java / Kotlin Game Framework) 11.7 with A* support and improved Kotlin API](https://www.reddit.com/r/Kotlin/comments/ekbp3k/fxgl_java_kotlin_game_framework_117_with_a/)
- url: https://www.reddit.com/r/Kotlin/comments/ekbp3k/fxgl_java_kotlin_game_framework_117_with_a/
---
Check out the new version at:  [https://github.com/AlmasB/FXGL](https://github.com/AlmasB/FXGL) 

Some example screenshots from existing games / demos:

[FXGL Games](https://preview.redd.it/ca4igbjjzx841.jpg?width=2561&amp;format=pjpg&amp;auto=webp&amp;s=04863a2c9c9081aab328b820a98142eea2233117)

There is a wiki for new users:  [https://github.com/AlmasB/FXGL/wiki](https://github.com/AlmasB/FXGL/wiki) 

New contributors are always welcome!
## [3][Can the Kotlin type system be leveraged to programmatically generate the map of HTTP and message bus inputs to all possible message bus outputs for a Kotlin service? Can the Arrow library help accomplish this?](https://www.reddit.com/r/Kotlin/comments/ek5030/can_the_kotlin_type_system_be_leveraged_to/)
- url: https://www.reddit.com/r/Kotlin/comments/ek5030/can_the_kotlin_type_system_be_leveraged_to/
---
My team is currently developing an application using a microservice architecture with pub-sub interservice communication over a message bus. Certain types of HTTP requests and messages received as inputs to a service might cause that service to subsequently publish certain types of messages over the bus as a result, and these published messages might trigger further publication events from other services as well.

I'd like to be able to programmatically generate a graph of all of the possible message flows without having to run any of these services - I know it's possible to do this through runtime analysis, but I'd like to solve this problem statically if possible. In our Kotlin services, messages are statically typed, and I believe it should be possible to programmatically determine message flow inputs and outputs for each Kotlin service at compile time or through static code analysis.

How should I start to tackle this problem? As one possible approach, I'm trying to determine if the Arrow functional programming library can encode potential message output side effects (of specific message types) at the type level, but I'm not sure if that's possible. For instance, is there a way to extend the Arrow IO type so that functions that directly or indirectly produce a certain type of message as a side effect can only be run in a monad of such an extended type which explicitly allows that (and which therefore implicitly disallows producing messages of types that are not whitelisted)?
## [4][Anyone would like to add support for Kotlin/Native in FOSS Godot game engine?](https://www.reddit.com/r/Kotlin/comments/ejuxss/anyone_would_like_to_add_support_for_kotlinnative/)
- url: https://twitter.com/reduzio/status/1213082947965853696
---

## [5][Book for Mastering Kotlin](https://www.reddit.com/r/Kotlin/comments/ejwqk4/book_for_mastering_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ejwqk4/book_for_mastering_kotlin/
---
Hey guys, could you recommend me some Kotlin advanced books?  


Thanks in advance
## [6][Is it better to use Kotlin for web dev?](https://www.reddit.com/r/Kotlin/comments/ejw236/is_it_better_to_use_kotlin_for_web_dev/)
- url: https://www.reddit.com/r/Kotlin/comments/ejw236/is_it_better_to_use_kotlin_for_web_dev/
---
I'm learning Kotlin and plan to create an app for android and also a web version. Is it better to stick in Kotlin (then transpile to JS) for the web version or just program my web version in JS?
## [7][Is it possible to code this style game in Kotlin alone?](https://www.reddit.com/r/Kotlin/comments/ejst8w/is_it_possible_to_code_this_style_game_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ejst8w/is_it_possible_to_code_this_style_game_in_kotlin/
---
I want to make something much more simple but the game boy styled game graphics and simple movements. Can anyone recommend reputable teachers or online sources on kotlin game development? 
https://www.youtube.com/watch?v=7dr0nABlGNM
## [8][Kotlin learning resources for a web developer trying to move to mobile?](https://www.reddit.com/r/Kotlin/comments/ejpgz9/kotlin_learning_resources_for_a_web_developer/)
- url: https://www.reddit.com/r/Kotlin/comments/ejpgz9/kotlin_learning_resources_for_a_web_developer/
---
First off, I apologize if this has been asked 100 times before.  I searched and found lots of questions, but none that seemed exactly in my sweet spot.

I am a experienced web developer (mostly React on the front end &amp; node/rails on the backend) and am an aws solutions architect who has had almost zero exposure to mobile other than PWAs and playing with React Native.  I have a couple ideas that I'd like to try to build and figured it would be a good chance to dip my toes in the mobile waters.

I'd love any advice on resources/books/courses/videos to get me started as someone who doesn't need an intro to programming and OO &amp; functional techniqures, but does need an intro to mobile and Kotlin.  Also, as an aside, I have ZERO java experience, and I am wondering if Kotlin is the way to go, or am I going to need to learn java at some point anyway?

Thanks for any and all the help.  I'm in a state of analysis by paralysis in getting started.
## [9][How to start with Tornadofx ?](https://www.reddit.com/r/Kotlin/comments/ejkls7/how_to_start_with_tornadofx/)
- url: https://www.reddit.com/r/Kotlin/comments/ejkls7/how_to_start_with_tornadofx/
---
I want to try Tornadofx with Kotlin. But I haven't found any step by step tutorial. Seems like Kotlin is still a new language :/

If anyone can help me with Tornadofx then it will be very helpful !!
## [10][Learning Materials for Kotlin](https://www.reddit.com/r/Kotlin/comments/eja5n5/learning_materials_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/eja5n5/learning_materials_for_kotlin/
---
Hi guys, I was recently asked to learn kotlin2.0, I have a background with Python and some Ruby.   
What resources would you suggest I get to learn Kotlin?
## [11][How to persuade the management to try Ktor?](https://www.reddit.com/r/Kotlin/comments/eja2ei/how_to_persuade_the_management_to_try_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/eja2ei/how_to_persuade_the_management_to_try_ktor/
---
Hi everyone, I tried Ktor with my pet project and it was a pretty interesting experience. It was something new after 3 years of using spring, my questions are:
1) Is it anyone who using ktor in production, and what is your opinion?
2) How to persuade the management(CTO, team leads) to start using Ktor?
