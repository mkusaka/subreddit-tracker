# Kotlin
## [1][Graal vs C2. Who runs Kotlin faster?](https://www.reddit.com/r/Kotlin/comments/gatx48/graal_vs_c2_who_runs_kotlin_faster/)
- url: https://medium.com/javarevisited/graal-vs-c2-who-runs-kotlin-faster-82f03f1b11dd
---

## [2][What are the idiomatic differences between Kotlin libraries and Java libraries? Is there a Kotlin way of creating library APIs?](https://www.reddit.com/r/Kotlin/comments/gackvy/what_are_the_idiomatic_differences_between_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gackvy/what_are_the_idiomatic_differences_between_kotlin/
---
Let's say I have two libraries that do the same task that are written in Java and in Kotlin. Assuming the Kotlin library is written following Kotlin best practices, what are the advantages of using the Kotlin library from a Kotlin code base? What are the disadvantages of using the one written in Java? Are there any killer features? And generally how does the Kotlin way of writing APIs differ from the one in Java?

Thank you.
## [3][Using JUnit 5 with Spring Boot 2, Kotlin and Mockito](https://www.reddit.com/r/Kotlin/comments/gajpzl/using_junit_5_with_spring_boot_2_kotlin_and/)
- url: https://medium.com//using-junit-5-with-spring-boot-2-kotlin-and-mockito-d5aea5b0c668?source=friends_link&amp;sk=d01893fe8bdc26687ff8c30be708b8b3
---

## [4][Spatial K - A WIP pure Kotlin Turfjs port and implementation of GeoJSON](https://www.reddit.com/r/Kotlin/comments/gak42w/spatial_k_a_wip_pure_kotlin_turfjs_port_and/)
- url: https://www.reddit.com/r/Kotlin/comments/gak42w/spatial_k_a_wip_pure_kotlin_turfjs_port_and/
---
Docs: https://dellisd.github.io/spatial-k/  
Repo: https://github.com/dellisd/spatial-k/

This is my first time going about making and publishing a library, so I'm mainly hoping for some feedback about the design of the library and its APIs before I commit to a proper first release.

The design goals of this library were to be compatible with Kotlin Multiplatform projects (i.e. written in pure Kotlin) while also having an API that could still be called from Java. 

The project is made up of three libraries (which are covered in more detail in the docs): 

**geojson**

Just contains a set of class definitions for GeoJSON objects with built in conversion to/from JSON strings.

**geojson-dsl**

This library is meant for Kotlin projects only and contains a DSL for building GeoJSON objects.

**turf**

Contains the port of Turfjs written in Kotlin that uses the geojson library for representing the GeoJSON input to the functions.

The library is still under development, mainly porting a more Turf functions, but like I mentioned at the top I'd love to hear any feedback about the library. Thanks!
## [5][Since we're in lockdown and I had nothing interesting to do. I spent all the free time I had last week on this and I would like your feedback, bug reports, feature requests, contributions, anything. It initially started as a learning exercise but I feel others can use it too.](https://www.reddit.com/r/Kotlin/comments/gaf32z/since_were_in_lockdown_and_i_had_nothing/)
- url: https://www.reddit.com/r/Kotlin/comments/gaf32z/since_were_in_lockdown_and_i_had_nothing/
---
I'm relatively new to Kotlin. I have some experience with Java so I picked it up easily.

I built this with Kotlin and Springboot this week. I'd like you guys to take a look.

https://github.com/BBloggsbott/ProfileServer
## [6][Printing a reflected hollow right triangle pattern (*) in Kotlin](https://www.reddit.com/r/Kotlin/comments/gamszr/printing_a_reflected_hollow_right_triangle/)
- url: https://www.reddit.com/r/Kotlin/comments/gamszr/printing_a_reflected_hollow_right_triangle/
---
 0

I am currently taking a programming principles class with a Kotlin language focus at my school due to my growing interest in programming.

I am to create a program that takes user input for an integer. After I collect this information, I need to print two triangles in a asterisk (\*) pattern (I had no problem completing these steps).

In addition, there is potential for extra credit which I am not working on more for experience than the extra points as I am looking good grade wise. I need to take the first triangle and reflect it while maintaining the properties. So far, I am unable to figure out how I would keep the triangle hollow and get the third side on to it.

My code:

    fun main() {
    
        // Prompt user for the triangle's size
        print( "How big do you want your triangles? " )
    
        // Get the triangle's size from the user and store in a variable
        var triangleSize = readLine()!!.toInt()
    
        // Decrementor for internal spaces
        var decrementor = triangleSize
    
        // Incrementor for internal spaces
        var incrementor = 2
    
        // Case given a 0 is entered
        if ( triangleSize == 0 ) {
    
            // Case given 1 is entered
        } else if ( triangleSize == 1 ) {
    
            println( "*" )
    
            println( "1" )
    
            for ( i in 1..( triangleSize - 1 )) {
    
                print( " " )
    
            }
    
            print( "*" )
    
            // Case given number entered that is not 0 or 1
        } else {
    
            // Print the top point of the triangle
            println( "*" )
    
            // Print the sides of the triangle
            for ( i in 0..( triangleSize - 3 )) {
    
                print( "*" )
    
                for ( i in 1..( triangleSize - decrementor )) {
    
                    print( " " )
    
                }
    
                println( "*" )
    
                decrementor--
    
            }
    
            // Print the bottom of the triangle
            for ( i in 1..triangleSize ) {
    
                print( "*" )
    
            }
    
            println()
    
            // Print the numbers triangle
            for ( i in 1..triangleSize ) {
    
                for ( j in 1..i ) {
    
                    print( "$j" )
    
                }
    
                println()
    
            }
    
            // Print the top of the inverse triangle
            for ( i in 1..( triangleSize - 1 )) {
    
                print( " " )
    
            }
    
            println( "*" )
    
            for ( i in 1..( triangleSize - 2 )) {
    
                for ( i in 1..( triangleSize - incrementor )) {
    
                    print( " " )
    
                }
    
                println( "*" )
    
                incrementor++
    
            }
    
            // Print the bottom of the inverse triangle
            for ( i in 1..triangleSize ) {
    
                print( "*" )
    
            }
    
        }
    
    }
## [7][How do you package your application](https://www.reddit.com/r/Kotlin/comments/ga6obi/how_do_you_package_your_application/)
- url: https://www.reddit.com/r/Kotlin/comments/ga6obi/how_do_you_package_your_application/
---
At work we have spring-boot and for my private projects I just use the gradle Shadow plugin. but I was just wondering if there is something better
What do you guys prefer to package and distribute applications?
## [8][A new DynamoDB Kotlin DSL generated based on AWS SDK v2](https://www.reddit.com/r/Kotlin/comments/ga56cg/a_new_dynamodb_kotlin_dsl_generated_based_on_aws/)
- url: https://www.reddit.com/r/Kotlin/comments/ga56cg/a_new_dynamodb_kotlin_dsl_generated_based_on_aws/
---
Hi there! I created a DSL generator for AWS SDK 2 and started with DynamoDB:

[https://github.com/cylab/awssdk-dynamodb-kotlin-dsl](https://github.com/cylab/awssdk-dynamodb-kotlin-dsl)

Feel free to try it out! Any feedback would be greatly appreciated!
## [9][Kotlin Coroutines](https://www.reddit.com/r/Kotlin/comments/g9ybag/kotlin_coroutines/)
- url: https://medium.com/@magdamiu/kotlin-coroutines-7938b5c7a11b
---

## [10][Trouble with generics and asList()](https://www.reddit.com/r/Kotlin/comments/ga1kyt/trouble_with_generics_and_aslist/)
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
