# Kotlin
## [1][Kobot - configurable conversational bot framework, open source](https://www.reddit.com/r/Kotlin/comments/g77jye/kobot_configurable_conversational_bot_framework/)
- url: https://www.reddit.com/r/Kotlin/comments/g77jye/kobot_configurable_conversational_bot_framework/
---
Hi folks,

In these days of quarantine, I started an open source project. I'd like to share it with you and gently ask, to everyone interested, feedback or critiques.

The project is hosted here: [https://github.com/besil/kobot](https://github.com/besil/kobot).

It is a framework to create conversational bots, using a simple configuration file.

The goal is to help non technical users to create conversational channels in an easy way. I know some proprietary software already do this, but I didn't find any open source project about it.

Main feature include:

* multiple conversational channels: telegram, http, cli, ...
* Database interaction through sql language, in order to send dynamic messages or to store information
* External HTTP API integration, in order to interact with web services
* Published as docker image to simplify devops

Framework is implemented 100% in Kotlin, with Spring Boot integration.

Any critique or feedback is welcome!

Thank you!
## [2][Kotlin Primary &amp; Secondary Constructors with Examples](https://www.reddit.com/r/Kotlin/comments/g77bfr/kotlin_primary_secondary_constructors_with/)
- url: https://kotlin-android.com/kotlin-primary-and-secondary-constructors/
---

## [3][How to diagnose OutOfMemoryError in Android?](https://www.reddit.com/r/Kotlin/comments/g75wh6/how_to_diagnose_outofmemoryerror_in_android/)
- url: https://blog.heaphero.io/2018/06/18/how-to-diagnose-outofmemoryerror-in-android/
---

## [4][Help with Ktor mainClassName](https://www.reddit.com/r/Kotlin/comments/g72lf2/help_with_ktor_mainclassname/)
- url: https://www.reddit.com/r/Kotlin/comments/g72lf2/help_with_ktor_mainclassname/
---
I have a Ktor application with fun main() inside Application.kt with the uk.astro at the top of the file. What should my mainClassName be inside Gradle? I've tried 'uk.astro.MainKt' but the resulting jar produces a classpath error.
## [5][Kotlin Microservices With Spring Boot And Spring Cloud: Part 3 - Inter-service communication](https://www.reddit.com/r/Kotlin/comments/g6n7sd/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://youtu.be/EJUtU_VGpuk
---

## [6][Kotlin/Native usage of an object created by invoking the constructor](https://www.reddit.com/r/Kotlin/comments/g6pfl5/kotlinnative_usage_of_an_object_created_by/)
- url: https://www.reddit.com/r/Kotlin/comments/g6pfl5/kotlinnative_usage_of_an_object_created_by/
---
For a new project of mine, I've been working on a command system that uses a registry.

[Registry](https://github.com/DeflatedPickle/ducknroll/blob/master/src/mingwMain/kotlin/com/deflatedpickle/ducknroll/common/common/registry/Registry.kt), [Command Example](https://github.com/DeflatedPickle/ducknroll/blob/master/src/mingwMain/kotlin/com/deflatedpickle/ducknroll/common/common/command/HelpCommand.kt), Commands are registered like [this](https://github.com/DeflatedPickle/ducknroll/blob/master/src/mingwMain/kotlin/example/main.kt#L19). When you want to parse user input for a command, you use [this](https://github.com/DeflatedPickle/ducknroll/blob/master/src/mingwMain/kotlin/example/main.kt#L34), which is then called like [this](https://github.com/DeflatedPickle/ducknroll/blob/master/src/mingwMain/kotlin/com/deflatedpickle/ducknroll/common/common/parser/CommandParser.kt).

&amp;#x200B;

You'll probably notice the fatal flaw, I didn't think about it before committing. It's hardcoded to use `HelpCommand`.

So, I went back to the Scrabble board the next morning. I made [these changes](https://github.com/DeflatedPickle/ducknroll/commit/7b04fa5e8741c64856fa193a1846998f62af8fac). The most notable is that commands are now registered with the constructor reference (`Registries.command.register("help", ::HelpCommand)`) as I decided it would be easiest to construct a zero-argument object, then change the arguments of the command after construction.

It runs, great. It constructs the command, it registers it fine, it even executes it fine. But here's the issue: I have no access to that constructed instance to change the arguments.

In the context of `command.get(split[0])?.run { }`, `this` is a `KFunction&lt;IHasRegistry&gt;`. I have access to the type the function returns, the argument types it takes, but I have no access to the actual instance.

&amp;#x200B;

Am I doing something blatantly wrong ~~as usual~~? Is this yet to be implemented? Are there any hacky ways I can get that *darn* instance?
## [7][Correct usage of also scope function](https://www.reddit.com/r/Kotlin/comments/g6oq6c/correct_usage_of_also_scope_function/)
- url: https://www.reddit.com/r/Kotlin/comments/g6oq6c/correct_usage_of_also_scope_function/
---
I'm asking myself what the real intention behind the `also` scope function is. In the [documentation](https://kotlinlang.org/docs/reference/scope-functions.html#also) they are saying 

&gt;also is good for performing some actions that take the context object as an argument. Use also for additional actions that don't alter the object, such as logging or  printing debug information. Usually, you can remove the calls of also from the call chain without breaking the program logic.

So far, I haven't seen any usage of `also` where you can remove the scope function without breaking the program logic. What is your opinion about it? How are you using it?

I've also created a post on stack overflow about it: [https://stackoverflow.com/questions/61367281/whats-the-real-intention-behind-kotlins-also-scope-function](https://stackoverflow.com/questions/61367281/whats-the-real-intention-behind-kotlins-also-scope-function)
## [8][Kotlin Koin Scope Illustrated](https://www.reddit.com/r/Kotlin/comments/g6n0w0/kotlin_koin_scope_illustrated/)
- url: https://medium.com/@elye.project/kotlin-koin-scope-illustrated-3bfa6c7ae98?source=friends_link&amp;sk=f771bc0469b042d155b6a2af5ace065c
---

## [9][Do you know which one is the best for developing cross-platform application: Flutter Vs Kotlin?](https://www.reddit.com/r/Kotlin/comments/g7297y/do_you_know_which_one_is_the_best_for_developing/)
- url: https://www.bacancytechnology.com/blog/flutter-vs-kotlin
---

## [10][Opinions on if without parentheses](https://www.reddit.com/r/Kotlin/comments/g6cw4c/opinions_on_if_without_parentheses/)
- url: https://www.reddit.com/r/Kotlin/comments/g6cw4c/opinions_on_if_without_parentheses/
---
Languages like Rust, Go or Swift decided to have mandatory curly braces and do away with parentheses for the condition:

Kotlin:
```kotlin
val result = if (x &lt; 5) true else false
```

Rust:
```rust
let result = if x &lt; 5 { true } else { false };
```

This is, in my opinion, better. There are many mistakes that are made by forgetting braces for `if`, for example [Apple's goto fail](https://nakedsecurity.sophos.com/2014/02/24/anatomy-of-a-goto-fail-apples-ssl-bug-explained-plus-an-unofficial-patch/).
```Kotlin:
if (x &lt; 5)
    do_something_1()
    do_something_2()
```
In the example above, `do_something_2()` will be always executed, even though someone might think it's part of the `if` because of the indentation.

Now, I don't want to make a breaking change the core of the language, of course. However, I think that an addition of a braces-mandatory `if` (which would be used by not adding parentheses around the `if` condition) could be a nice addition to the language.

What do you think?
