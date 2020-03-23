# Kotlin
## [1][Authorization for a Kotlin Spring backend, using JSON Web Tokens](https://www.reddit.com/r/Kotlin/comments/fngfcj/authorization_for_a_kotlin_spring_backend_using/)
- url: https://medium.com//authorization-for-a-kotlin-spring-backend-using-json-web-tokens-9f3c7b0d1ee7?source=friends_link&amp;sk=d1da9d3f7594d2b66cc882f260fd000f
---

## [2][Getting Started with Kotlin on iOS, Part 2: Interop](https://www.reddit.com/r/Kotlin/comments/fnbx8x/getting_started_with_kotlin_on_ios_part_2_interop/)
- url: https://benasher.co/kotlin-ios-getting-started-interop/
---

## [3][Ktor, nested features](https://www.reddit.com/r/Kotlin/comments/fnjzrz/ktor_nested_features/)
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
## [4][Kotlin 1.4-M1 Released](https://www.reddit.com/r/Kotlin/comments/fnjpot/kotlin_14m1_released/)
- url: https://blog.jetbrains.com/kotlin/2020/03/kotlin-1-4-m1-released/
---

## [5][Is there a better way to get a non-nullable object from a map?](https://www.reddit.com/r/Kotlin/comments/fn9njv/is_there_a_better_way_to_get_a_nonnullable_object/)
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
## [6][Keycloak Authentication Infrastructure as Code, inspired by LiquiBase](https://www.reddit.com/r/Kotlin/comments/fn1a88/keycloak_authentication_infrastructure_as_code/)
- url: https://github.com/klg71/keycloakmigration
---

## [7][Coroutines: a practical vocabulary](https://www.reddit.com/r/Kotlin/comments/fmnllf/coroutines_a_practical_vocabulary/)
- url: https://www.rockandnull.com/coroutines-practical/
---

## [8][A template for building Alfred workflows in Kotlin/JS](https://www.reddit.com/r/Kotlin/comments/fmu4hu/a_template_for_building_alfred_workflows_in/)
- url: https://twitter.com/VladyslavSitalo/status/1241584247812448256?s=20
---

## [9][A sample app built with Kotlin, Coroutines, Flow, Dagger 2, Architecture Components, MVVM, Room, Retrofit, Moshi, Material Components and other Jetpack libraries.](https://www.reddit.com/r/Kotlin/comments/fm91pb/a_sample_app_built_with_kotlin_coroutines_flow/)
- url: https://github.com/PatilShreyas/Foodium
---

## [10][Try make code friends](https://www.reddit.com/r/Kotlin/comments/fmnen3/try_make_code_friends/)
- url: https://www.reddit.com/r/Kotlin/comments/fmnen3/try_make_code_friends/
---
Which is the best way to make friends for talk and code?
