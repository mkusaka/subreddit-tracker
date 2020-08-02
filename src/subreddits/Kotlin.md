# Kotlin
## [1][Newbie question about Kotlin Native and sharing code between Android and iOS](https://www.reddit.com/r/Kotlin/comments/i2bf32/newbie_question_about_kotlin_native_and_sharing/)
- url: https://www.reddit.com/r/Kotlin/comments/i2bf32/newbie_question_about_kotlin_native_and_sharing/
---
Dear Friends,

&amp;#x200B;

I developed an Android app that makes intensive use of the JSoup library to parse HTML. As an example, I will say that this happens inside a class called CustomHtmlAnalyzer.kt. After a couple of years developing this app, my CustomHtmlAnalyzer is grown, is complex, and works very well. I have covered it with a lot of unit tests, and I have all under control.

Now I'm porting the Android app to a native iOS app, and I reached the point where a I need the same CustomHtmlAnalyzer class but for iOS. Fortunately, there is a pod library for iOS called SwiftSoup, which works like Jsoup. I could write a copy of my CustomHtmlAnalyzer but in Swift. Besides that, I could translate each function and line of code. They would be pretty much the same, only written in Swift.

Would it be possible to reuse the class that I wrote in Kotlin, using kotlin native? I can abstract this class from jsoup, and then somehow configure the shared code that in case of the Android app, the jsoup library is used, and in case of the iOS version, the SwiftSoup library is used.

&amp;#x200B;

I saw a similar sample regarding the Android SharedPreferences, and the iOS UserDefaults. In this case, using the expected/actual feature, it is possible to call platform-dependent functions. Of course, the android SharedPreferences and iOS UserDefaults are supported natively by Kotlin Native.  


And because of that my question. Would it be achievable somehow, using libraries that are not natively supported, like in this case Jsoup and SwiftSoup?

If you know some examples or post in the internet where this is explained, I would be very grateful.

&amp;#x200B;

Tanks in advance.
## [2][AndroidBites | Snippet | three most useful but least known kotlin list functions](https://www.reddit.com/r/Kotlin/comments/i263xd/androidbites_snippet_three_most_useful_but_least/)
- url: https://chetangupta.net/union-intersect-subs/
---

## [3][How Do I Change The Naming Format Gradle Produces?](https://www.reddit.com/r/Kotlin/comments/i241sf/how_do_i_change_the_naming_format_gradle_produces/)
- url: https://www.reddit.com/r/Kotlin/comments/i241sf/how_do_i_change_the_naming_format_gradle_produces/
---
Whenever I build my binary the name that's outputted is "ColoredEverything-\[1.0.0\]" but I want the name to be "\[1.0.0\] ColoredEverything" how would I do that?

My build.gradle.kts is:

    import com.github.jengelman.gradle.plugins.shadow.tasks.ConfigureShadowRelocation
    import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar
    plugins {
        java
        val kotlinVersion: String by System.getProperties()
        kotlin("jvm").version(kotlinVersion)
        id("com.github.johnrengelman.shadow").version("6.0.0")
    }
    group = "com.smushytaco"
    version = "[1.0.0]"
    repositories {
        mavenCentral()
        maven("https://papermc.io/repo/repository/maven-public/")
    }
    dependencies {
        val kotlinVersion: String by System.getProperties()
        implementation(kotlin("stdlib", kotlinVersion))
        //https://papermc.io/javadocs/paper/1.16/overview-summary.html
        compileOnly("com.destroystokyo.paper", "paper-api", "1.16.1-R0.1-SNAPSHOT")
    }
    val autoRelocate by tasks.register&lt;ConfigureShadowRelocation&gt;("configureShadowRelocation", ConfigureShadowRelocation::class) {
        target = tasks.getByName("shadowJar") as ShadowJar?
        val packageName = "${project.group}.${project.name.toLowerCase()}"
        prefix = "$packageName.shaded"
    }
    tasks {
        shadowJar {
            archiveClassifier.set("")
            project.configurations.implementation.get().isCanBeResolved = true
            configurations = listOf(project.configurations.implementation.get())
            dependsOn(autoRelocate)
        }
        build {
            dependsOn(shadowJar)
        }
    }

My settings.gradle.kts is:

    rootProject.name = "ColoredEverything"

My gradle.properties is:

    kotlin.code.style=official
    systemProp.kotlinVersion=1.3.72
## [4][What would you say is the best way for a beginner to learn about GUIs and such in Kotlin?](https://www.reddit.com/r/Kotlin/comments/i1bmvk/what_would_you_say_is_the_best_way_for_a_beginner/)
- url: https://www.reddit.com/r/Kotlin/comments/i1bmvk/what_would_you_say_is_the_best_way_for_a_beginner/
---
I'm currently learning Kotlin, and after learning Java, one of the things I always wanted to do is to learn how to create GUIs. Are there any good courses or material on the subject? What about the best frameworks for the job?
## [5][Can an API throw a compile-time error instead of a runtime one?](https://www.reddit.com/r/Kotlin/comments/i1htcd/can_an_api_throw_a_compiletime_error_instead_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/i1htcd/can_an_api_throw_a_compiletime_error_instead_of_a/
---
Take this code for example:

    fun example(string: String) {
        if (string != "123") {
            println(string)
        } else {
            throw Exception("string cannot be equal to \"123\"")
        }
    }
    fun main() {
        example("123")
    }

It's obviously going to crash at runtime but Kotlin gives no compile-time warning or error regarding it. Is there a way to throw an error at compile-time and if not is it a planned feature?
## [6][Is There An Annotation To Do What I'm Trying To Achieve?](https://www.reddit.com/r/Kotlin/comments/i1ny9c/is_there_an_annotation_to_do_what_im_trying_to/)
- url: https://www.reddit.com/r/Kotlin/comments/i1ny9c/is_there_an_annotation_to_do_what_im_trying_to/
---
Is there an annotation to turn this:

    class example : Cancellable {
        var isCancelled = false
    }

into this:

    class example : Cancellable {
        private var isCancelled = false
        override fun setCancelled(cancel: Boolean) {
            isCancelled = cancel
        }
        override fun isCancelled(): Boolean = isCancelled
    }

or do I just have to live with doing the second one?
## [7][Introduction/Beginner Question](https://www.reddit.com/r/Kotlin/comments/i19mek/introductionbeginner_question/)
- url: https://www.reddit.com/r/Kotlin/comments/i19mek/introductionbeginner_question/
---
Hey what's up everyone! I'm new to this Reddit community, Kotlin, and programming in general. After looking through alot of the posts and seeing how everyone helps each other out I'm happy to have found this group and be apart of it. Just started with Kotlin and learning the basics from a uDemy course which has been great so far. I am definitely motivated and know I have a long way to go to reach my goals. 

My goal is to create Android apps and just learn the Kotlin language to the point where I can think of an idea and just have a good idea about how I'd make it happen. My question for the community is what other knowledge should I be seeking besides just Kotlin and Android Studio knowledge. The course I'm taking right now with Udemy has me working in IntelliJ which has been actually really nice to work in as an IDE. I have read that Data Structures and Algorithms and Linear Algebra are also go knowledge to have. What do you guys think about the other areas? Should I be learning these things in conjunction with learning the basics of Kotlin.

I've been really motivated and have been trying to work 2-4 hours a day learning. I started an excel spreadsheet (lol I love spreadsheets) to track my time studying. Ive  been equally motivated to learn Kotlin as I am just to add time to my spreadsheet. I know I have to put in the work, so far it's been a week and I've got 20 hours in so far, which I don't think is too bad working a full time job. But any other advice you guys have is appreciated. Thanks for reading.
## [8][Why can’t I use val inside Plugins {}?](https://www.reddit.com/r/Kotlin/comments/i15dfs/why_cant_i_use_val_inside_plugins/)
- url: https://www.reddit.com/r/Kotlin/comments/i15dfs/why_cant_i_use_val_inside_plugins/
---
    val kotlinVersion = "1.3.72"
    plugins {
        // Error: 'val kotlinVersion: String' can't be called in this context by implicit receiver. Use the explicit one if necessary
        kotlin("jvm").version(kotlinVersion)]
    }

 I use Kotlin’s standard library as a dependency too and I only want to have to specify the version in one place but when I try something like I did above in my build.gradle.kts I get that error you see in the comment. How do I resolve this?
## [9][How Do I Resolve This Type Mismatch?](https://www.reddit.com/r/Kotlin/comments/i10xri/how_do_i_resolve_this_type_mismatch/)
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
## [10][Androibites | Destructuring Params with safety | Overcoming limitation of Positional Destructuring](https://www.reddit.com/r/Kotlin/comments/i11b5x/androibites_destructuring_params_with_safety/)
- url: https://chetangupta.net/destructuring-safely/
---

