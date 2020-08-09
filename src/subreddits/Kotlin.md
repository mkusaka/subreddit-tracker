# Kotlin
## [1][Why JetBrains created Kotlin? - YouTube - Credits to all the articles from around 2011](https://www.reddit.com/r/Kotlin/comments/i6fkjq/why_jetbrains_created_kotlin_youtube_credits_to/)
- url: https://youtu.be/zTUZ7PQgWsI
---

## [2][Guide to Quarkus with Kotlin](https://www.reddit.com/r/Kotlin/comments/i6g3pg/guide_to_quarkus_with_kotlin/)
- url: https://piotrminkowski.com/2020/08/09/guide-to-quarkus-with-kotlin/
---

## [3][Good places for help with Kotlin?](https://www.reddit.com/r/Kotlin/comments/i6hatf/good_places_for_help_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/i6hatf/good_places_for_help_with_kotlin/
---
G'day Kotlin folk! 

I'm fairly new to the Kotlin community, but I am a web and iOS developer. I have recently done a (just OK, not great) kotlin udemy course, and have been doing the rounds on youtube. 

Are there any recommended discord servers, subreddits (besides this one) or other forums for fledgling Kotlin developers such as myself? Stack Overflow is not the place for newbie questions, but sometimes I would like to ask something that may seem obvious to experienced Android devs.

You'll see me around... 🤙Thanks champions.
## [4][Code Discussion](https://www.reddit.com/r/Kotlin/comments/i6a07a/code_discussion/)
- url: https://www.reddit.com/r/Kotlin/comments/i6a07a/code_discussion/
---
    I was attending an android development course, in Member classes the instructor said that to change value of variable in class we should type 

`"fun main() {`

`var umar = Person("Umar", "Akram", 16)`

`umar.hobby = "Playing chess"`

`umar.getHobby()`

`}`

`class Person(firstName:String = "John", secondName:String = "Doe"){`

`var age : Int? = null`

`var hobby : String = "Watching Netflix"`

`var firstName :String? = null ".`

    

 Then I was wondering about that we can use ( `var hobby = "Anything"`)  instead  
what do you think?
## [5][KVision 3.12.0 is released (with support for Micronaut framework)](https://www.reddit.com/r/Kotlin/comments/i5y5an/kvision_3120_is_released_with_support_for/)
- url: https://www.reddit.com/r/Kotlin/comments/i5y5an/kvision_3120_is_released_with_support_for/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.12.0, with a new module supporting the [Micronaut](https://micronaut.io) web framework. Micronaut provides a simple compile-time aspect-oriented programming API that does not use reflection. It's perfect for building modular, easily testable microservice and serverless applications.

There are some [new example applications](https://github.com/rjaros/kvision-examples), including the new [fullstack template](https://github.com/rjaros/kvision-examples/tree/master/template-fullstack-micronaut), which present how to build KVision fullstack apps with Micronaut.

For more details about this release see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.12.0).

As always, any feedback is appreciated!
## [6][Guide/order of concepts to understand Recyclerview in Android?](https://www.reddit.com/r/Kotlin/comments/i6929g/guideorder_of_concepts_to_understand_recyclerview/)
- url: https://www.reddit.com/r/Kotlin/comments/i6929g/guideorder_of_concepts_to_understand_recyclerview/
---
I started to learn Recyclerview in Kotlin and I see there are so many things to take care about, I'm new programming so I'm trying to figure out concepts like inflate, adapters, inheritance, etc.

Can you suggest an order of things to understand and use Recyclerviews?
## [7][Automatically generating Dokka API docs using Github Actions and deploying to Github Pages](https://www.reddit.com/r/Kotlin/comments/i5m4km/automatically_generating_dokka_api_docs_using/)
- url: https://www.reddit.com/r/Kotlin/comments/i5m4km/automatically_generating_dokka_api_docs_using/
---
I just got this working and though it would be useful to share.

First step is to configure Dokka in your [build.gradle](https://github.com/kwebio/kweb-core/blob/master/build.gradle#L15):

    plugins {
        ...
        id 'org.jetbrains.dokka' version '1.4.0-rc'
        ...
    }

Next configure the Github Action, you can see how I've done it [here](https://github.com/kwebio/kweb-core/blob/master/.github/workflows/build.yml#L25).

This will deploy the documentation to the `gh-pages` branch of your respository, see [here](https://pages.github.com/)  for more information.

You can see the result here: http://dokka.kweb.io/kweb-core/
## [8][Learning Kotlin From Scratch](https://www.reddit.com/r/Kotlin/comments/i5x16o/learning_kotlin_from_scratch/)
- url: https://www.reddit.com/r/Kotlin/comments/i5x16o/learning_kotlin_from_scratch/
---
I am a CS graduate, I know basic programming concepts such as loops, conditionals, variables, etc.

However, I have never in my life worked with Java. I recently decided I want to learn Android Development, and before diving straight into Android development I want to learn the Kotlin language properly.

Which resource would you recommend for someone with no prior Java experience.
## [9][Zoe - The Kafka CLI for humans (written in Kotlin): Version 0.26.0 is out with the ability to filter consumed records based on Kafka headers, variable substitutions and more!](https://www.reddit.com/r/Kotlin/comments/i5enrg/zoe_the_kafka_cli_for_humans_written_in_kotlin/)
- url: https://github.com/adevinta/zoe/releases/tag/v0.26.0
---

## [10][How Do I Convert My build.gradle.kts To Use The Multiplatform Module Instead Of The JVM One?](https://www.reddit.com/r/Kotlin/comments/i55fsf/how_do_i_convert_my_buildgradlekts_to_use_the/)
- url: https://www.reddit.com/r/Kotlin/comments/i55fsf/how_do_i_convert_my_buildgradlekts_to_use_the/
---
    plugins {
        java
        val kotlinVersion: String by System.getProperties()
        kotlin("jvm").version(kotlinVersion)
        id("com.github.johnrengelman.shadow").version("6.0.0")
    }
    group = "com.smushytaco"
    version = "1.0.0"
    repositories {
        mavenCentral()
        maven("https://papermc.io/repo/repository/maven-public/")
    }
    dependencies {
        val kotlinVersion: String by System.getProperties()
        implementation(kotlin("stdlib", kotlinVersion))
        implementation("org.rocksdb", "rocksdbjni", "6.11.4")
        //https://papermc.io/javadocs/paper/1.16/overview-summary.html
        compileOnly("com.destroystokyo.paper", "paper-api", "1.16.1-R0.1-SNAPSHOT")
    }
    tasks {
        shadowJar {
            archiveClassifier.set("")
            project.configurations.implementation.get().isCanBeResolved = true
            configurations = listOf(project.configurations.implementation.get())
            relocate("kotlin", "com.smushytaco.src.main.kotlin")
            minimize()
        }
        build {
            dependsOn(shadowJar)
        }
    }

That's the build.gradle.kts that uses the plain old jvm plugin module and I'm trying to convert it to use the multiplatform one in preparation of Kotlin 1.4 and this was my attempt at it:

    plugins {
        java
        val kotlinVersion: String by System.getProperties()
        kotlin("multiplatform").version(kotlinVersion)
        id("com.github.johnrengelman.shadow").version("6.0.0")
    }
    kotlin {
        jvm()
        sourceSets {
            val jvmMain by getting {
                dependencies {
                    val kotlinVersion: String by System.getProperties()
                    implementation(kotlin("stdlib", kotlinVersion))
                    implementation("org.rocksdb:rocksdbjni:6.11.4")
                    //https://papermc.io/javadocs/paper/1.16/overview-summary.html
                    compileOnly("com.destroystokyo.paper:paper-api:1.16.1-R0.1-SNAPSHOT")
                }
            }
        }
    }
    group = "com.smushytaco"
    version = "1.0.0"
    repositories {
        mavenCentral()
        maven("https://papermc.io/repo/repository/maven-public/")
    }
    tasks {
        shadowJar {
            archiveClassifier.set("")
            project.configurations.implementation.get().isCanBeResolved = true
            configurations = listOf(project.configurations.implementation.get())
            relocate("kotlin", "com.smushytaco.src.main.kotlin")
            minimize()
        }
        build {
            dependsOn(shadowJar)
        }
    } 

This is very broken and doesn't work. For starters it's like the dependencies get ignored. With my MPP attempted build.gradle.kts whenever I go to my code there are errors everywhere because the dependencies aren't properly build handled. Another issue I'm facing is the fact nothing is being shaded in the jars (Because dependencies are being ignored). The last issue I'm dealing with is that this MPP build.gradle.kts creates 3 jars instead of 1, before it just created  Ride-1.0.0.jar but now it creates  Ride-1.0.0.jar, Ride-jvm-1.0.0.jar, and Ride-metadata-1.0.0.jar. How do I fix all these issues?
