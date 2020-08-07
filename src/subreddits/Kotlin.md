# Kotlin
## [1][Kotlin JSON serialization crash course](https://www.reddit.com/r/Kotlin/comments/i504u9/kotlin_json_serialization_crash_course/)
- url: https://www.rockandnull.com/kotlin-json/
---

## [2][How Do I Convert My build.gradle.kts To Use The Multiplatform Module Instead Of The JVM One?](https://www.reddit.com/r/Kotlin/comments/i55fsf/how_do_i_convert_my_buildgradlekts_to_use_the/)
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
## [3][Advanced Kotlin and Coroutines resources](https://www.reddit.com/r/Kotlin/comments/i4vfyr/advanced_kotlin_and_coroutines_resources/)
- url: https://www.reddit.com/r/Kotlin/comments/i4vfyr/advanced_kotlin_and_coroutines_resources/
---
I'm a developer writing in Kotlin for 2 years but I wasn't using Coroutines yet that much.  Are there any courses/training/coding puzzles you could recommend to write a better idiomatic code?  It doesn't have to be a free one. 

Same standing for coroutines.
## [4][Where can you learn about all of the list methods similar to map, zip etc (and when to use them)?](https://www.reddit.com/r/Kotlin/comments/i53niw/where_can_you_learn_about_all_of_the_list_methods/)
- url: https://www.reddit.com/r/Kotlin/comments/i53niw/where_can_you_learn_about_all_of_the_list_methods/
---
I want to replace normal loops but every time I need a solution I have no idea what to Google to find something that would solve my usecase
## [5][AndroidBites | Common Mistakes Everyone Makes In Accumulator Pattern | Kotlin Fold and Reduce](https://www.reddit.com/r/Kotlin/comments/i5693w/androidbites_common_mistakes_everyone_makes_in/)
- url: https://chetangupta.net/reduce-fold/
---

## [6][KMP with Android and JVM](https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/i4ojmr/kmp_with_android_and_jvm/
---
I'm trying to setup a project with three targets Android, JVM(Ktor) and JS.

Everything works fine when only JS and JVM targets are included , the moment I add android target, the JVM target can't recognize anything from the *common* module but the android target can, and vice versa if I remove the android target.

Has anyone had this problem?
## [7][I will code for your project (want experience)](https://www.reddit.com/r/Kotlin/comments/i46kea/i_will_code_for_your_project_want_experience/)
- url: https://www.reddit.com/r/Kotlin/comments/i46kea/i_will_code_for_your_project_want_experience/
---
Hi. If anyone has a project they are working on, I would like to contribute. I do not expect compensation, i simply want something meaningful to do and some real-world experience.

Edit: I was thinking of something like a personal/hobby project but if you are representing a commercial project, I would like to get paid appropriately. Thanks for the tip
## [8][ELI5: Nested for-loops](https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/)
- url: https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/
---
Hi guys, I'm currently learning Kotlin, it's my first programming language - so I'm totally new to this.

Everything so far has made sense, and I've managed to figure it out, but this for-loop is driving me crazy!

`fun main() {`

  `for (i in 1..3) {`

  `for (j in 1..i) {`

`print(j)`

`}`

`}`

`}`

From the IDE, I know the print output is:  112123, but I cannot figure out why!

1..i should be 1, and i is 1..3 which is 123, so shouldn't the output be 1231 or 1123?

Can someone please ELI5 this for me!
## [9][OAS3 Implementation](https://www.reddit.com/r/Kotlin/comments/i4asnh/oas3_implementation/)
- url: https://www.reddit.com/r/Kotlin/comments/i4asnh/oas3_implementation/
---
Hi, I'm making some tests with the Open API Specs.  


I made a request to the Marvel API and I need to decode the JSON response to an OAS3 file, I found [this](https://github.com/schehata/oas3) package that lets you make an OAS3 document within the IDE but I can't make it work. I already compiled the package in a .jar file and add it as a module in the main project, but it doesn't make the auto-import and doesn't recognize the package.  


Thanks in advance.
## [10][Best resource for switching from Java to Kotlin - Android dev](https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/
---
Forgive me if this has been asked and answered a million times, I ran a search and only found threads about learning Android/Kotlin from scratch.

I'm a professional Android Dev and have been for years, very familiar with Java, use Python and C a fair bit for work as well. I find I can read Kotlin fine, but figure it's about time I started writing it. Hoping to get a few pointers to some good resources for people who know Android/Java already, and just want to learn Kotlin, not Android or Programming from scratch.

&amp;#x200B;

Edit: Thanks everyone, this looks like exactly what I was after! Much appreciated
