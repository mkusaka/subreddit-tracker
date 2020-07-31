# Kotlin
## [1][Why can’t I use val inside Plugins {}?](https://www.reddit.com/r/Kotlin/comments/i15dfs/why_cant_i_use_val_inside_plugins/)
- url: https://www.reddit.com/r/Kotlin/comments/i15dfs/why_cant_i_use_val_inside_plugins/
---
    val kotlinVersion = "1.3.72"
    plugins {
        // Error: 'val kotlinVersion: String' can't be called in this context by implicit receiver. Use the explicit one if necessary
        kotlin("jvm").version(kotlinVersion)]
    }

 I use Kotlin’s standard library as a dependency too and I only want to have to specify the version in one place but when I try something like I did above in my build.gradle.kts I get that error you see in the comment. How do I resolve this?
## [2][How Do I Resolve This Type Mismatch?](https://www.reddit.com/r/Kotlin/comments/i10xri/how_do_i_resolve_this_type_mismatch/)
- url: https://www.reddit.com/r/Kotlin/comments/i10xri/how_do_i_resolve_this_type_mismatch/
---
So I was working on converting my build.gradle (Which uses Groovy) to a build.gradle.kts (Which uses Kotlin) and this is what I got so far:

Groovy Original:

    plugins {
        id 'java'
        id 'org.jetbrains.kotlin.jvm' version '1.3.72'
        id 'com.github.johnrengelman.shadow' version '5.2.0'
    }
    group 'com.smushytaco'
    version '1.0-SNAPSHOT'
    repositories {
        mavenCentral()
        maven {
            url 'https://papermc.io/repo/repository/maven-public/'
        }
    }
    dependencies {
        implementation group: 'org.jetbrains.kotlin', name: 'kotlin-stdlib', version: '1.3.72'
        //https://papermc.io/javadocs/paper/1.16/overview-summary.html
        compileOnly group: 'com.destroystokyo.paper', name: 'paper-api', version: '1.16.1-R0.1-SNAPSHOT'
    }
    shadowJar {
        getArchiveClassifier().set('')
        project.configurations.implementation.canBeResolved = true
        configurations = [project.configurations.implementation]
        relocate 'kotlin', 'com.smushytaco.src.main.kotlin.com.smushytaco.plugin'
    }
    build.dependsOn shadowJar

Kotlin Conversion:

    plugins {
        java
        kotlin("jvm") version("1.3.72")
        id("com.github.johnrengelman.shadow") version("5.2.0")
    }
    group = "com.smushytaco"
    version = "1.0-SNAPSHOT"
    repositories {
        mavenCentral()
        maven("https://papermc.io/repo/repository/maven-public/")
    }
    dependencies {
        implementation("org.jetbrains.kotlin", "kotlin-stdlib", "1.3.72")
        //https://papermc.io/javadocs/paper/1.16/overview-summary.html
        compileOnly("com.destroystokyo.paper", "paper-api", "1.16.1-R0.1-SNAPSHOT")
    }
    tasks {
        shadowJar {
            archiveClassifier.set("")
            project.configurations.implementation.get().isCanBeResolved = true
            // Type Mismatch:
            // Required: (Mutable)List&lt;Configuration!&gt;!
            //Found: Array&lt;NamedDomainObjectProvider&lt;Configuration&gt;&gt;
            configurations = [project.configurations.implementation]
            relocate("kotlin", "com.smushytaco.src.main.kotlin.com.smushytaco.plugin")
        }
        build {
            dependsOn(shadowJar)
        }
    }

 I commented the error I'm getting with the Kotlin rewrite, why does it work with the original Groovy one and not with this one? How do I fix this error?
## [3][Androibites | Destructuring Params with safety | Overcoming limitation of Positional Destructuring](https://www.reddit.com/r/Kotlin/comments/i11b5x/androibites_destructuring_params_with_safety/)
- url: https://chetangupta.net/destructuring-safely/
---

## [4][Why am I getting org.json.JSONException: Not a primitive array?](https://www.reddit.com/r/Kotlin/comments/i103nz/why_am_i_getting_orgjsonjsonexception_not_a/)
- url: https://www.reddit.com/r/Kotlin/comments/i103nz/why_am_i_getting_orgjsonjsonexception_not_a/
---
I have this function:

    fun getFeed(token: String): Result&lt;Array&lt;Post?&gt;, Exception&gt; {
            if (token != "") {
                val (request, response, result) = "https://alles.cx/api/feed"
                    .httpGet()
                    .header("Authorization", token.toString())
                    .response()
    
                when (result) {
                    is Result.Success -&gt; {
                        println("Result success")
                        try {
                            val responseJSON = JSONObject(String(response.data))
                            if (!responseJSON.has("err")) {
                                println("Has no error")
                                var feedArray = arrayOfNulls&lt;Post&gt;(0)
                                println("HERE?")
                                var feed = JSONArray(responseJSON["feed"])   // &lt;-----ERROR HERE 
                                println("not here")
                                println("Made to the for")
                                for (i in 0 until feed.length()) {
                                    println("For $i")
                                    val jsonObject = feed.getJSONObject(i)
    
                                    var imageURL: URL? = null
                                    if (jsonObject.has("image")) {
                                        imageURL = URL(jsonObject["image"].toString())
                                    }
                                    feedArray[i] = Post(
                                        jsonObject.getString("id"),
                                        User(jsonObject.getJSONObject("author").getString("id"), jsonObject.getJSONObject("author").getString("username"), jsonObject.getJSONObject("author").getString("name")),
                                        jsonObject.getString("createdAt"),
                                        jsonObject.getInt("replyCount"),
                                        jsonObject.getInt("score"),
                                        jsonObject.getString("content"),
                                        null,
                                        imageURL,
                                        jsonObject.getInt("vote")
    
                                    )
                                    println("Added post $i")
                                }
                                println(feedArray)
                                return Result.Success(feedArray)
                            }
                            else {
                                println("yes error")
                                return Result.Failure(Exception(responseJSON["err"].toString()))
                            }
                        }
                        catch(e: Exception) {
                            return Result.Failure(e)
                        }
                    }
                    is Result.Failure -&gt; {
                        return Result.Failure(Exception(result.component2()?.localizedMessage))
                    }
                }
            }
            else {
                return Result.Failure(Exception("badAuth"))
            }
        }

An error is occurring on this line. I know this because of the print statements that are in the code:

    var feed = JSONArray(responseJSON["feed"])   // &lt;-----ERROR HERE 

This is what the error that is occurring is.

    [Failure: org.json.JSONException: Not a primitive array: class org.json.JSONArray]

This is what the feed object from the get request will look like, just so you have an idea of what it is:

    {
        "feed": [
            {
                "type": "post",
                "slug": "equ2ZfaX2UThtkS8DzkXdc",
                "author": {
                    "id": "00000000-0000-0000-0000-000000000000",
                    "name": "Archie",
                    "username": "archie",
                    "plus": true
                },
                "content": "I know it's quite inactive here atm. But I'm doing tons of stuff behind the scenes - basically rewriting the entire platform from scratch. Micro will be done by the end of the summer.",
                "image": null,
                "createdAt": "2020-07-30T20:39:23.000Z",
                "score": 5,
                "vote": 0,
                "replyCount": 0,
                "hasParent": false
            }
        ]
    }

What could possibly be going wrong here. I am using Fuel for the request, but I don't think that is the problem
## [5][VKUG - Apollo-Android: a Journey to Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/i0psls/vkug_apolloandroid_a_journey_to_kotlin/)
- url: https://www.youtube.com/watch?v=u_CtCWFH8jA
---

## [6][I made a puzzle game and an automatic solver](https://www.reddit.com/r/Kotlin/comments/i0cgta/i_made_a_puzzle_game_and_an_automatic_solver/)
- url: https://www.reddit.com/r/Kotlin/comments/i0cgta/i_made_a_puzzle_game_and_an_automatic_solver/
---
Hi there! I spent the last couple days working on this project, and I wanted to share it.

Here it is: https://github.com/amtejani/Number-Game

I hope some of you will have the patience to take a look at the game and maybe give me feedback on my code. I (obviously) didn't put much work into the UI since my goal here was to write an algorithm to solve the puzzle. I also didn't handle exceptions/unexpected states very gracefully.

I do have a couple questions for those who can answer. 

I don't have much experience with Gradle other than just adding dependencies in Android Gradle scripts. How do I add 3rd party libraries with  this project? The build.gradle file was generated with the Korge intellij plugin. 

Also, does anyone have suggestions for another algorithms project I could pick up next?

Thanks for reading!
## [7][Kotlin 1.4.0-RC: Debugging coroutines – Kotlin Blog](https://www.reddit.com/r/Kotlin/comments/i01l6p/kotlin_140rc_debugging_coroutines_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-rc-debugging-coroutines/
---

## [8][AndroidBites | 6 things that you need to know before destructuring Parameters](https://www.reddit.com/r/Kotlin/comments/i0fcjz/androidbites_6_things_that_you_need_to_know/)
- url: https://chetangupta.net/destructuring-limitations/
---

## [9][Computer Science Needs the Equivalent of Zero](https://www.reddit.com/r/Kotlin/comments/i0rzam/computer_science_needs_the_equivalent_of_zero/)
- url: https://www.reddit.com/r/Kotlin/comments/i0rzam/computer_science_needs_the_equivalent_of_zero/
---
[We Haven't Invented Zero Yet](https://www.bruceeckel.com/2020/07/30/we-havent-invented-zero-yet/)
## [10][Bootiful GraphQL with Kotlin](https://www.reddit.com/r/Kotlin/comments/hzzleu/bootiful_graphql_with_kotlin/)
- url: https://youtu.be/1siPT1pTXFU?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

