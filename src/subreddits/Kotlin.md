# Kotlin
## [1][A few questions about Kotlin Native / MPP](https://www.reddit.com/r/Kotlin/comments/fo4phk/a_few_questions_about_kotlin_native_mpp/)
- url: https://www.reddit.com/r/Kotlin/comments/fo4phk/a_few_questions_about_kotlin_native_mpp/
---
Hey I'm mostly an iOS developer and new to Kotlin. I'm building a Kotlin Multiplatform app for fun and to learn a bit about the ecosystem. I had a pretty rough time setting everything up and getting the right versions of dependencies to play well together as well as a hard time wrapping my head around async tasks between my iOS code and my shared Kotlin framework, I'm also seeing some really awful compile times for a fairly simple app (+1.5 minutes). How much of this will likely change in the future and how much of this is just down to my inexperience with the platform? on paper it looks great but I'm finding it pretty difficult to get behind for a serious project.
## [2][Kotlin Named Tuple equivalent?](https://www.reddit.com/r/Kotlin/comments/fnzfry/kotlin_named_tuple_equivalent/)
- url: https://www.reddit.com/r/Kotlin/comments/fnzfry/kotlin_named_tuple_equivalent/
---
Hey all, I am in the process of converting my Blackjack game from Python to Kotlin. One of the tools I used was collections, where I could use collections.namedtuple('Card',['rank','suit','faceval']) function. 

This would set up a list of cards, where each card would have a rank ("Ace" through "King"), suit, and faceval (0-9,10,10,10,10). I could then create two loops and use .append() to add card values. 

Does anyone have any idea on how I should do this in Kotlin? I've tried making a  mutableList of each item in the loop, but I'm not sure how to add that list to another list. 

Thx in advance.

Here is my code: https://i.imgur.com/F2mdaTI.png
## [3][Kotlin 1.4-M1 Released](https://www.reddit.com/r/Kotlin/comments/fnjpot/kotlin_14m1_released/)
- url: https://blog.jetbrains.com/kotlin/2020/03/kotlin-1-4-m1-released/
---

## [4][Kotlin 1.3.71](https://www.reddit.com/r/Kotlin/comments/fnmug8/kotlin_1371/)
- url: https://github.com/JetBrains/kotlin/releases/tag/v1.3.71
---

## [5][Authorization for a Kotlin Spring backend, using JSON Web Tokens](https://www.reddit.com/r/Kotlin/comments/fngfcj/authorization_for_a_kotlin_spring_backend_using/)
- url: https://medium.com//authorization-for-a-kotlin-spring-backend-using-json-web-tokens-9f3c7b0d1ee7?source=friends_link&amp;sk=d1da9d3f7594d2b66cc882f260fd000f
---

## [6][Has anyone here used/seen kotlin/native used with flutter via dart ffi?](https://www.reddit.com/r/Kotlin/comments/fnwii0/has_anyone_here_usedseen_kotlinnative_used_with/)
- url: https://www.reddit.com/r/Kotlin/comments/fnwii0/has_anyone_here_usedseen_kotlinnative_used_with/
---
asking for a friend ;)
## [7][Getting Started with Kotlin on iOS, Part 2: Interop](https://www.reddit.com/r/Kotlin/comments/fnbx8x/getting_started_with_kotlin_on_ios_part_2_interop/)
- url: https://benasher.co/kotlin-ios-getting-started-interop/
---

## [8][Ktor, nested features](https://www.reddit.com/r/Kotlin/comments/fnjzrz/ktor_nested_features/)
- url: https://www.reddit.com/r/Kotlin/comments/fnjzrz/ktor_nested_features/
---
Maybe I have luck here, stackoverflow doesn't have a big ktor community.

Has anyone achieved that only the last interceptor is executed when they are nested? Like

&amp;#x200B;

    a(1) {
      a(2) {
       call.respond("")
      }
    }

I do not want a(1) to be executed for the call pipeline at all as soon as one child also has a registered.
## [9][Is there a better way to get a non-nullable object from a map?](https://www.reddit.com/r/Kotlin/comments/fn9njv/is_there_a_better_way_to_get_a_nonnullable_object/)
- url: https://www.reddit.com/r/Kotlin/comments/fn9njv/is_there_a_better_way_to_get_a_nonnullable_object/
---
Hi!

I'm pretty new to Kotlin. I've rewritten a smaller Java component of mine into Kotlin throughout the weekend and I'm liking what I'm seeing thus far. I have but one question for now:

&gt;if (name !in catsByName) {  
 return  
}  
val cat: Cat = catsByName\[name\]!!

There are some occasions when I ran into code like this. I don't want to do a thing when there is no value under the key specified, but if there is I want my code to run past through the first if guard. The problem is that the get method of the map (here it is \[\]) gives back a Cat? type when in realty I need a non-nullable Cat. Is there a better practice to get a non-nullable Cat type or is using the !! operator the best thing to do here? I know it is safe, there is no chance it will throw an exception here but still, I guess it looks a bit ugly.

Thanks!
## [10][Keycloak Authentication Infrastructure as Code, inspired by LiquiBase](https://www.reddit.com/r/Kotlin/comments/fn1a88/keycloak_authentication_infrastructure_as_code/)
- url: https://github.com/klg71/keycloakmigration
---

