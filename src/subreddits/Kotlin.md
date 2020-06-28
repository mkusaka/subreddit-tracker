# Kotlin
## [1][Physics Algorithm (in Kotlin)](https://www.reddit.com/r/Kotlin/comments/hha9gn/physics_algorithm_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hha9gn/physics_algorithm_in_kotlin/
---
Hi all, this is my first post on this. I've been writing a book about Physics Algorithms (still fairly early)- how to do simulations. I'm aiming it at undergraduate physics students (who aren't necessarily good at programming), and computer scientists, engineers, mathematicians, and people with coding experience already.

I chose Kotlin for most of the algorithms in the book, and source code is presented too --- I did this because Kotlin is an amazing language for doing physics simulations in: it's fast, concise and easy to work with --- I would say that C and Python, although popular for this, are inadequate for teaching physics algorithms because C is too difficult to manage for largish projects, and python is too slow. I'm not expecting much disagreement on that here.

Anyway, so I'm looking for feedback or even just interest. Part of my reason to post this is to get the motivation to finish the book.

So if you're interested, the link is here:

[http://www.articlesbyaphysicist.com/book\_physics\_algorithm.html](http://www.articlesbyaphysicist.com/book_physics_algorithm.html)

And please do just leave a message here saying whether the idea is interesting, or whether you have definite feedback.

Bear in mind that it's early days, and I'll be adding a lot to this in all chapters.
## [2][Creating a GraphQL Api automatically with Kotlin, Ktor and Postgres](https://www.reddit.com/r/Kotlin/comments/hgzxb6/creating_a_graphql_api_automatically_with_kotlin/)
- url: https://medium.com/@chris.hinshaw/creating-a-graphql-database-service-with-kotlin-ktor-and-postgres-18b3a5ea9998
---

## [3][How could you clear the terminal in Kotlin?](https://www.reddit.com/r/Kotlin/comments/hhbat8/how_could_you_clear_the_terminal_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hhbat8/how_could_you_clear_the_terminal_in_kotlin/
---
Hi! I'm new to programming in Kotlin and I want know how you could clear the terminal. (I'm not using Android Studio)

code:

    fun clear() {
        // code to clear screen
    }
    
    fun main() {
        print("enter text: ")
        e = readLine()
        if (e == "clear") {
            clear()
        }
    }

Any help appreciated!
## [4][Can JS Kotlin and Java Kotlin interact?](https://www.reddit.com/r/Kotlin/comments/hgwjbf/can_js_kotlin_and_java_kotlin_interact/)
- url: https://www.reddit.com/r/Kotlin/comments/hgwjbf/can_js_kotlin_and_java_kotlin_interact/
---
Title.
## [5][I have worked with Python extensively. Can any one suggest me resources to get started with Kotlin?](https://www.reddit.com/r/Kotlin/comments/hgpao5/i_have_worked_with_python_extensively_can_any_one/)
- url: https://www.reddit.com/r/Kotlin/comments/hgpao5/i_have_worked_with_python_extensively_can_any_one/
---
I would prefer it would direct me somehow toward mobile application development for Android.
## [6][Can I import classes from other modules when using spring boot?](https://www.reddit.com/r/Kotlin/comments/hgvcac/can_i_import_classes_from_other_modules_when/)
- url: https://www.reddit.com/r/Kotlin/comments/hgvcac/can_i_import_classes_from_other_modules_when/
---
For some weird reasons, I kept on getting Unresolved reference: &lt;ClassName&gt; when I try to import a class from different module.

When I try to remove the spring dependencies in application module and just make a basic kotlin class and import that,, I don't get those unresolved reference issues. I don't understand what's happening now.

First time trying out spring so maybe they're doing some kind of magic here.

Can anyone check my config? Really weird.

&amp;#x200B;

    Application module build.gradle
    
    plugins {
        kotlin("jvm")
        id("org.springframework.boot")
        id("io.spring.dependency-management")
        kotlin("plugin.spring")
        kotlin("plugin.jpa")
    }
    
    dependencies {
        implementation(project(":domain"))
        implementation("org.springframework.boot:spring-boot-starter-web")
        implementation("org.springframework.boot:spring-boot-starter-data-jpa")
        implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
        implementation("org.jetbrains.kotlin:kotlin-reflect")
        implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
        developmentOnly("org.springframework.boot:spring-boot-devtools")
        runtimeOnly("com.h2database:h2")
    }
    
    Delivery module
    plugins {
    	kotlin("jvm")
    	id("org.springframework.boot")
    	id("io.spring.dependency-management")
    	kotlin("plugin.spring")
    	kotlin("plugin.jpa")
    }
    
    dependencies {
    	implementation(project(":domain"))
    	implementation(project(":application"))
    
    	implementation("org.springframework.boot:spring-boot-starter-web")
    	implementation("org.springframework.boot:spring-boot-starter-data-jpa")
    	implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
    	implementation("org.jetbrains.kotlin:kotlin-reflect")
    	implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    	developmentOnly("org.springframework.boot:spring-boot-devtools")
    	runtimeOnly("com.h2database:h2")
    	testImplementation("org.springframework.boot:spring-boot-starter-test") {
    		exclude(group = "org.junit.vintage", module = "junit-vintage-engine")
    	}
    }
    
    Root gradle.build
    import org.jetbrains.kotlin.gradle.tasks.KotlinCompile
    
    plugins {
        kotlin("jvm") version "1.3.72"
        id("org.springframework.boot") version "2.3.1.RELEASE"
        id("io.spring.dependency-management") version "1.0.9.RELEASE"
        kotlin("plugin.spring") version "1.3.72"
        kotlin("plugin.jpa") version "1.3.72"
    }
    
    buildscript {
        repositories {
            mavenCentral()
        }
    }
    
    allprojects {
        group = "com.example"
        version = "0.0.1-SNAPSHOT"
    
        tasks.withType&lt;JavaCompile&gt; {
            sourceCompatibility = "1.8"
            targetCompatibility = "1.8"
        }
    
        tasks.withType&lt;KotlinCompile&gt; {
            kotlinOptions {
                freeCompilerArgs = listOf("-Xjsr305=strict")
                jvmTarget = "1.8"
            }
        }
    
        tasks.withType&lt;Test&gt; {
            useJUnitPlatform()
        }
    
        repositories {
            mavenCentral()
        }
    }
## [7][Creating a BitTorrent client in Kotlin](https://www.reddit.com/r/Kotlin/comments/hg7ym2/creating_a_bittorrent_client_in_kotlin/)
- url: https://kcianfarini.dev/corbit-bencoding/
---

## [8][A question about nulls](https://www.reddit.com/r/Kotlin/comments/hgjdy1/a_question_about_nulls/)
- url: https://www.reddit.com/r/Kotlin/comments/hgjdy1/a_question_about_nulls/
---
Hey guys!

I wanted to ask you a (quick?) question about null types in Kotlin. To give you an idea about my background, I'm a little bit of a newbie programmer that started with Python about a year ago. I started because I wanted to automate some of the activities we do at my office and I really liked it, so I dived into it. Now I wanted to try to get into some app programming and I understood that Kotlin is great for Android development.  And It's also pretty fun actually.

I'm taking some courses and trying things step by step but, maybe mostly because I come from Python as first language, I'm having a really hard time with the static type way of creating variables and stating everything. The hardest thing to understand to me is the "null" type. I can't understand what it stands for and why I should use a nullable type in general.

In Python everything is dynamically stated, so I don't have to think about what type of var I'm creating. Here, on the other hand, it seems to be the base of the workflow. And the fact that I don't seem to properly understand this worries me a bit. Even worse with the null type.

Can you please give me an advice about this? Maybe then I can stop to wrap my head around this with your help.

Thanks guys!
## [9][Kotlin Flows and Channels for Android](https://www.reddit.com/r/Kotlin/comments/hg6940/kotlin_flows_and_channels_for_android/)
- url: https://youtu.be/xch4aw7hNcY?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

## [10][Possible bug with generic where clause and star projection.](https://www.reddit.com/r/Kotlin/comments/hg6npk/possible_bug_with_generic_where_clause_and_star/)
- url: https://www.reddit.com/r/Kotlin/comments/hg6npk/possible_bug_with_generic_where_clause_and_star/
---
It appears that star projections only understand the first where clause. I'm not sure how to best articulate it, so have some code.

Here's a playground link: https://pl.kotl.in/a6vtI8tLG

Here's the code:

    interface Foo {
        fun foo()
    }
    
    interface Bar {
        fun bar()
    }
    
    interface FooAndBar&lt;T&gt; where T: Foo, T: Bar {
        val value: T
    }
    
    interface BarAndFoo&lt;T&gt; where T: Bar, T: Foo {
        val value: T
    }
    
    fun doSomething(thing: FooAndBar&lt;*&gt;) {
        thing.value.foo() // fine
        thing.value.bar() // compile error
    }
    
    fun doSomething(thing: BarAndFoo&lt;*&gt;) {
        thing.value.foo() // compile error
        thing.value.bar() // fine
    }
