# Kotlin
## [1][KotlinConf 2019 Videos](https://www.reddit.com/r/Kotlin/comments/ebd0np/kotlinconf_2019_videos/)
- url: https://www.youtube.com/playlist?list=PLQ176FUIyIUY6SKGl3Cj9yeYibBuRr3Hl
---

## [2][Kotlin First Impressions](https://www.reddit.com/r/Kotlin/comments/elfqqs/kotlin_first_impressions/)
- url: https://www.reddit.com/r/Kotlin/comments/elfqqs/kotlin_first_impressions/
---
In short, I’m impressed.

I’ve been doing software development for over three decades. Long ago, I had the chance to do some work in Lisp, and thoroughly enjoyed the power and productivity that functional programming enabled. Unfortunately, Lisp tends to turn off many programmers, simply because of its unconventional syntax (which isn’t hard to get used to, but is unconventional, and as the saying goes “you never get a second chance to make a first impression”).

Kotlin is sneaky. It’s a very Scheme-ish functional programming language disguised as a conventional language. That makes it much less likely to inspire the sort of negative gut reactions that the Lisp family of languages does. (Ruby tried to do the same thing, but falling for the anti-patterns of Do What I Mean and There Is More Than One Way to Do It really hurt that language, IMHO.)

The biggest handicap that any new programming language has is libraries. It’s the existing libraries you can leverage that do so much to drive the decision of what programming language is best for a project, so starting out from scratch with a brand-new and mostly empty ecosystem of libraries sets up a simply huge hurdle for any new programming language to overcome. Targeting the JVM was an enormous win. (Getting endorsed by Google as the preferred language for Android app development was another coup, but it would have never happened without the one of targeting the JVM.)

The question is, what happens now. With great power comes great responsibility. Most programmers are mediocre programmers; will they tend to shy away from the more powerful features of Kotlin and mostly write Java code without semicolons, or will they tend to carelessly play with creating new syntax-level features and create unmaintainable code? The answer to that question will, I think, be the determining factor as to how popular and widespread Kotlin becomes.
## [3][Let's Review: Pokedex - zsmb.co](https://www.reddit.com/r/Kotlin/comments/eld6sb/lets_review_pokedex_zsmbco/)
- url: https://zsmb.co/lets-review-pokedex/
---

## [4][What tool do the presenters use in KotlinConf](https://www.reddit.com/r/Kotlin/comments/el742o/what_tool_do_the_presenters_use_in_kotlinconf/)
- url: https://www.reddit.com/r/Kotlin/comments/el742o/what_tool_do_the_presenters_use_in_kotlinconf/
---
I want to give a talk at my company to embrace Kotlin. We are currently a Java shop and I believe that since we are going to be addressing some technical debt, we might as well clean stuff up.

I am supposed to give a presentation on Kotlin soon and I wanted to know how do the presenters at KotlinConf make the presentations. In particular, how do they get the code snippets to fold, expand, and change as they give their talks. 

Thanks in advance
## [5][EvoMaster: Automated Test Generation for Java/Kotlin RESTful APIs](https://www.reddit.com/r/Kotlin/comments/el935m/evomaster_automated_test_generation_for/)
- url: https://github.com/EMResearch/EvoMaster
---

## [6][Android App architecture. Everything in kotlin vs client - server interaction](https://www.reddit.com/r/Kotlin/comments/elc5p4/android_app_architecture_everything_in_kotlin_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/elc5p4/android_app_architecture_everything_in_kotlin_vs/
---
I am trying to build an android app which will include google maps API, some user data storing, and payments.

I have zero experience with kotlin and android development so I would like to know what is the best practice.

Is it safe to put the business logic including payments in the kotlin code? 

As an alternative, I was thinking to create a server in another language, i.e Java using akka-http for example and make REST API calls from the android client. This would be a nice option if for example I want to make an iOS App out of it in the future.

&amp;#x200B;

Any suggestions?
## [7][OpenAPI/Swagger specs (and clients for specs) using Javalin and Maven](https://www.reddit.com/r/Kotlin/comments/ekv8ht/openapiswagger_specs_and_clients_for_specs_using/)
- url: https://javalin.io/tutorials/openapi-example?language=kotlin
---

## [8][Custom Android XML Lint Rules with Kotlin](https://www.reddit.com/r/Kotlin/comments/ekwlfr/custom_android_xml_lint_rules_with_kotlin/)
- url: https://coderamblings.dev/posts/custom-android-xml-lint-rules-with-kotlin/
---

## [9][Request for Kotlin Android learning](https://www.reddit.com/r/Kotlin/comments/el0k7c/request_for_kotlin_android_learning/)
- url: https://www.reddit.com/r/Kotlin/comments/el0k7c/request_for_kotlin_android_learning/
---
Does anyone have any good resources for learning android for Kotlin? Like starting from downloading the program to beginning your first app
## [10][Sealed class -&gt; Subclass: Sealed() -&gt; Subclass.Inner: Sealed() = error](https://www.reddit.com/r/Kotlin/comments/ekzv0v/sealed_class_subclass_sealed_subclassinner_sealed/)
- url: https://www.reddit.com/r/Kotlin/comments/ekzv0v/sealed_class_subclass_sealed_subclassinner_sealed/
---
Worst title ever, right?

Here's the issue. I have a sealed class and I would like a class and its inner class to both be members of the sealed class. It works if the class is nested in the sealed class but not if it's not nested. Is this a bug?

This one works:  

    sealed class Foo {
        class Bar : Foo() {
            inner class Baz : Foo()
        }
    }

This gets an error: "Cannot access '&lt;init&gt;': it is private in 'Foo'"

    sealed class Foo
    
    class Bar: Foo() {
        inner class Baz: Foo()
    }
## [11][Kotlin Tutorials](https://www.reddit.com/r/Kotlin/comments/ekro1y/kotlin_tutorials/)
- url: https://kotlintutorialblog.com
---

