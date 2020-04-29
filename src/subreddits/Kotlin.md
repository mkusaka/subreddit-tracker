# Kotlin
## [1][Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/g9ybag/kotlin_coroutines/)
- url: https://medium.com/@magdamiu/kotlin-coroutines-7938b5c7a11b
---

## [2][A new DynamoDB Kotlin DSL generated based on AWS SDK v2](https://www.reddit.com/r/Kotlin/comments/ga56cg/a_new_dynamodb_kotlin_dsl_generated_based_on_aws/)
- url: https://www.reddit.com/r/Kotlin/comments/ga56cg/a_new_dynamodb_kotlin_dsl_generated_based_on_aws/
---
Hi there! I created a DSL generator for AWS SDK 2 and started with DynamoDB:

[https://github.com/cylab/awssdk-dynamodb-kotlin-dsl](https://github.com/cylab/awssdk-dynamodb-kotlin-dsl)

Feel free to try it out! Any feedback would be greatly appreciated!
## [3][Trouble with generics and asList()](https://www.reddit.com/r/Kotlin/comments/ga1kyt/trouble_with_generics_and_aslist/)
- url: https://www.reddit.com/r/Kotlin/comments/ga1kyt/trouble_with_generics_and_aslist/
---
I have the following code:

        sealed class Tile {
            class ResourceTile: Tile() { ... }
            class AnotherTile: Tile() { ... }
            ...
        }
     
        class Hex&lt;out T : Tile&gt;(val tile: T) { ... }
    
        inline fun &lt;reified T:Tile&gt; hexes(): List&lt;Hex&lt;T&gt;&gt; {
            return allHexes().filter { it.tile is T } as List&lt;Hex&lt;T&gt;&gt;
        }
    
    

The usage is `foo = hexes&lt;ResourceTile&gt;()`, where `foo` will have the type `List&lt;Hex&lt;ResourceTile&gt;&gt;`.

This works. It compiles, it runs, life is grand. But IntelliJ flags \`as List&lt;Hex&lt;T&gt;&gt;\` as being an unchecked cast. IntelliJ normally knows its business; so, when it gives me a warning I try to make it happy. 

Is there a syntax/rewrite of this that will get rid of IntelliJ's warning message?
## [4][How do you package your application](https://www.reddit.com/r/Kotlin/comments/ga6obi/how_do_you_package_your_application/)
- url: https://www.reddit.com/r/Kotlin/comments/ga6obi/how_do_you_package_your_application/
---
At work we have spring-boot and for my private projects I just use the gradle Shadow plugin. but I was just wondering if there is something better
What do you guys prefer to package and distribute applications?
## [5][Kotlin Microservices With Spring Boot And Spring Cloud: Part 4 - API gateway](https://www.reddit.com/r/Kotlin/comments/g9nnua/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://youtu.be/XIkSWHX38Tg
---

## [6][Kotlin, Spring boot and MySql](https://www.reddit.com/r/Kotlin/comments/g9ubkm/kotlin_spring_boot_and_mysql/)
- url: https://www.reddit.com/r/Kotlin/comments/g9ubkm/kotlin_spring_boot_and_mysql/
---
[https://medium.com/@zahidrasheed/part-2-connect-restful-api-with-mysql-bc21d212de48](https://medium.com/@zahidrasheed/part-2-connect-restful-api-with-mysql-bc21d212de48)
## [7][RESTful pagination in Spring using Link header](https://www.reddit.com/r/Kotlin/comments/g9xkhm/restful_pagination_in_spring_using_link_header/)
- url: https://www.matchilling.com/restful-pagination-in-spring-using-link-header/
---

## [8][Android VideoView in Kotlin with Examples](https://www.reddit.com/r/Kotlin/comments/g9xbaz/android_videoview_in_kotlin_with_examples/)
- url: https://kotlin-android.com/android-videoview-kotlin/
---

## [9][Best android knowledges in 2020](https://www.reddit.com/r/Kotlin/comments/g9zgsk/best_android_knowledges_in_2020/)
- url: https://www.reddit.com/r/Kotlin/comments/g9zgsk/best_android_knowledges_in_2020/
---
##Hey everyone

###I'm a new in Android development world and I want to know, what I should use to build a amazing mobile app? What technologies, patterns, best pratices, style guide do you recommend?

I'm a React JS, React Native and Node JS developer and the develop environment is so different, I have been studying Kotlin for a month to build a application backend and now I want to go to mobile
## [10][Is kotlin the primary language to develop Android applications?](https://www.reddit.com/r/Kotlin/comments/g9u3x3/is_kotlin_the_primary_language_to_develop_android/)
- url: https://www.reddit.com/r/Kotlin/comments/g9u3x3/is_kotlin_the_primary_language_to_develop_android/
---
Hi,

I haven't done native Android app development but come from hybrid mobile app development. I wanted to get started with native android application development. I know java used to be the official language but now since kotlin is also officially supported. How are the new Android apps written? Is kotlin the primary choice? Why?
