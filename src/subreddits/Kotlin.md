# Kotlin
## [1][Can I use a lambda for null checking?](https://www.reddit.com/r/Kotlin/comments/i2qozw/can_i_use_a_lambda_for_null_checking/)
- url: https://www.reddit.com/r/Kotlin/comments/i2qozw/can_i_use_a_lambda_for_null_checking/
---
    fun exampleFunction() : Int? = null
    val example = exampleFunction() ?: {
        
    }

I want to use a lambda like so to run more code in case where the value is null like so but I get an error saying:

    Type mismatch.
    Required: Int
    Found: () → Unit

How do I do what I'm trying to do?
## [2][Small, but interesting RxJava2 Puzzlers](https://www.reddit.com/r/Kotlin/comments/i2vxuy/small_but_interesting_rxjava2_puzzlers/)
- url: https://www.reddit.com/r/Kotlin/comments/i2vxuy/small_but_interesting_rxjava2_puzzlers/
---
Really cool [RxJava2 Puzzlers](https://mailchi.mp/631bd69d4410/smallengineer-rxjava2puzzlers-free)! Check it out
## [3][Need Help Learning Kotlin?, Check out our community, https://devguildweb.xyz](https://www.reddit.com/r/Kotlin/comments/i2omyz/need_help_learning_kotlin_check_out_our_community/)
- url: https://www.reddit.com/r/Kotlin/comments/i2omyz/need_help_learning_kotlin_check_out_our_community/
---
&amp;#x200B;

[https://devguildweb.xyz](https://devguildweb.xyz/)

Hello!,  We are The Developers Guild, a friendly discord community offering free  help and chat about java and many other programming languages.

We  want more Kotlin programmers to join our ranks, so we together can teach and learn from each other

[https://discord.gg/xqUN8KY](https://discord.gg/xqUN8KY)

&gt;Discord Community: Script Review, Paired Programming, Collab, and Live Instruction for API Engineers!  
A  community of student and professional developers. Join us for free  instruction, paired programming, code/project review, collaboration, and  more. Looking for exposure to new concepts? Engage in peer review on  your projects? Want to shout to the world that your API is online and  ready for use? Looking for help getting through your latest assignment  or project? If you can ask a good question you're like to find a good  answer here!  If you're more on the expert side and you're so inclined  you can engage in the Feinberg method of learning, where you reinforce  what you do know and spot weaknesses in your own studies, by teaching a  few code-along courses or just share your latest project.  Regardless of  your skill level though we are eager to have individuals from the  software development sphere to collaborate and converse with. To teach  and be taught. In short: We'd love to have you. Come on in!  
[https://discord.gg/xqUN8KY](https://discord.gg/xqUN8KY)
## [4][Any other solution to this issue with type checking?](https://www.reddit.com/r/Kotlin/comments/i2u69v/any_other_solution_to_this_issue_with_type/)
- url: https://www.reddit.com/r/Kotlin/comments/i2u69v/any_other_solution_to_this_issue_with_type/
---
Hey! I would like to know if there's a non-"workaround" solution to a type checking issue I'm facing:

Imagine the following example function:
```kt
/** Calls [fn] with `num + 1`. */
fun &lt;T&gt; withPlus1(num: Int, fn: (num: Int) -&gt; T): T = fn(num + 1)
```
What I would like to do is to provide a default `fn` argument to `withPlus1`.

Something like:
```kt
/** Calls [fn] with `num + 1`. [fn] defaults to stringifying its input. */
fun &lt;T&gt; withPlus1(
    num: Int,
    fn: (num: Int) -&gt; T = { num: Int -&gt; num.toString() }
): T = fn(num + 1)
```
The above doesn't compile because `num.toString()` is of type `String` rather than `T`.

Casting `num.toString()` to `T` makes it so that I can't call `withPlus1(5)` because there isn't enough type information to infer that `T` is `String` (i.e. I would have to call `withPlus1&lt;String&gt;(5)`), which is a no-go.

The obvious "workaround" solution here is to create a second function that calls the first function with the my default `fn`, i.e.:

```kt
fun withPlus1(num: Int): String = withPlus1(num) { num: Int -&gt; num.toString() }
```
However, is there any way to make this work without relying on a second function? I.e., by providing a default argument to the first version of the function? I played around with the `@BuilderInference` experimental directive but couldn't get it to work. Am I missing something?
## [5][Composing in the wild](https://www.reddit.com/r/Kotlin/comments/i2el3z/composing_in_the_wild/)
- url: https://medium.com/@shikasd/composing-in-the-wild-145761ad62c3
---

## [6][Newbie question about Kotlin Native and sharing code between Android and iOS](https://www.reddit.com/r/Kotlin/comments/i2bf32/newbie_question_about_kotlin_native_and_sharing/)
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
## [7][AndroidBites | Snippet | three most useful but least known kotlin list functions](https://www.reddit.com/r/Kotlin/comments/i263xd/androidbites_snippet_three_most_useful_but_least/)
- url: https://chetangupta.net/union-intersect-subs/
---

## [8][How Do I Change The Naming Format Gradle Produces?](https://www.reddit.com/r/Kotlin/comments/i241sf/how_do_i_change_the_naming_format_gradle_produces/)
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
## [9][What would you say is the best way for a beginner to learn about GUIs and such in Kotlin?](https://www.reddit.com/r/Kotlin/comments/i1bmvk/what_would_you_say_is_the_best_way_for_a_beginner/)
- url: https://www.reddit.com/r/Kotlin/comments/i1bmvk/what_would_you_say_is_the_best_way_for_a_beginner/
---
I'm currently learning Kotlin, and after learning Java, one of the things I always wanted to do is to learn how to create GUIs. Are there any good courses or material on the subject? What about the best frameworks for the job?
## [10][Can an API throw a compile-time error instead of a runtime one?](https://www.reddit.com/r/Kotlin/comments/i1htcd/can_an_api_throw_a_compiletime_error_instead_of_a/)
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
