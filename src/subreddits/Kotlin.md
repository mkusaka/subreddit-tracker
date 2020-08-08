# Kotlin
## [1][KVision 3.12.0 is released (with support for Micronaut framework)](https://www.reddit.com/r/Kotlin/comments/i5y5an/kvision_3120_is_released_with_support_for/)
- url: https://www.reddit.com/r/Kotlin/comments/i5y5an/kvision_3120_is_released_with_support_for/
---
[KVision](https://github.com/rjaros/kvision) is an open source web framework created for Kotlin/JS. It allows developers to build modern web applications with the Kotlin language.

I have released KVision 3.12.0, with a new module supporting the [Micronaut](https://micronaut.io) web framework. Micronaut provides a simple compile-time aspect-oriented programming API that does not use reflection. It's perfect for building modular, easily testable microservice and serverless applications.

There are some [new example applications](https://github.com/rjaros/kvision-examples), including the new [fullstack template](https://github.com/rjaros/kvision-examples/tree/master/template-fullstack-micronaut), which present how to build KVision fullstack apps with Micronaut.

For more details about this release see the [changelog](https://github.com/rjaros/kvision/releases/tag/3.12.0).

As always, any feedback is appreciated!
## [2][Automatically generating Dokka API docs using Github Actions and deploying to Github Pages](https://www.reddit.com/r/Kotlin/comments/i5m4km/automatically_generating_dokka_api_docs_using/)
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
## [3][Learning Kotlin From Scratch](https://www.reddit.com/r/Kotlin/comments/i5x16o/learning_kotlin_from_scratch/)
- url: https://www.reddit.com/r/Kotlin/comments/i5x16o/learning_kotlin_from_scratch/
---
I am a CS graduate, I know basic programming concepts such as loops, conditionals, variables, etc.

However, I have never in my life worked with Java. I recently decided I want to learn Android Development, and before diving straight into Android development I want to learn the Kotlin language properly.

Which resource would you recommend for someone with no prior Java experience.
## [4][Zoe - The Kafka CLI for humans (written in Kotlin): Version 0.26.0 is out with the ability to filter consumed records based on Kafka headers, variable substitutions and more!](https://www.reddit.com/r/Kotlin/comments/i5enrg/zoe_the_kafka_cli_for_humans_written_in_kotlin/)
- url: https://github.com/adevinta/zoe/releases/tag/v0.26.0
---

## [5][Kotlin JSON serialization crash course](https://www.reddit.com/r/Kotlin/comments/i504u9/kotlin_json_serialization_crash_course/)
- url: https://www.rockandnull.com/kotlin-json/
---

## [6][Advanced Kotlin and Coroutines resources](https://www.reddit.com/r/Kotlin/comments/i4vfyr/advanced_kotlin_and_coroutines_resources/)
- url: https://www.reddit.com/r/Kotlin/comments/i4vfyr/advanced_kotlin_and_coroutines_resources/
---
I'm a developer writing in Kotlin for 2 years but I wasn't using Coroutines yet that much.  Are there any courses/training/coding puzzles you could recommend to write a better idiomatic code?  It doesn't have to be a free one. 

Same standing for coroutines.
## [7][How Do I Convert My build.gradle.kts To Use The Multiplatform Module Instead Of The JVM One?](https://www.reddit.com/r/Kotlin/comments/i55fsf/how_do_i_convert_my_buildgradlekts_to_use_the/)
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
## [8][Where can you learn about all of the list methods similar to map, zip etc (and when to use them)?](https://www.reddit.com/r/Kotlin/comments/i53niw/where_can_you_learn_about_all_of_the_list_methods/)
- url: https://www.reddit.com/r/Kotlin/comments/i53niw/where_can_you_learn_about_all_of_the_list_methods/
---
I want to replace normal loops but every time I need a solution I have no idea what to Google to find something that would solve my usecase
## [9][AndroidBites | Common Mistakes Everyone Makes In Accumulator Pattern | Kotlin Fold and Reduce](https://www.reddit.com/r/Kotlin/comments/i5693w/androidbites_common_mistakes_everyone_makes_in/)
- url: https://chetangupta.net/reduce-fold/
---

## [10][KMP with Android and JVM](https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/
---
I'm trying to setup a project with three targets Android, JVM(Ktor) and JS.

Everything works fine when only JS and JVM targets are included , the moment I add android target, the JVM target can't recognize anything from the *common* module but the android target can, and vice versa if I remove the android target.

Has anyone had this problem?
