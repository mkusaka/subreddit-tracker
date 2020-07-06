# Kotlin
## [1][Getting Started with Kotlin on iOS, Part 3: The Build](https://www.reddit.com/r/Kotlin/comments/hm1jck/getting_started_with_kotlin_on_ios_part_3_the/)
- url: https://benasher.co/kotlin-ios-getting-started-build/
---

## [2][Wrapper for Wikia API with minmal abstractions in Kotlin (WIP)](https://www.reddit.com/r/Kotlin/comments/hlw66z/wrapper_for_wikia_api_with_minmal_abstractions_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hlw66z/wrapper_for_wikia_api_with_minmal_abstractions_in/
---
[Github](https://github.com/Gabgab2003/WikiaKT/)

I wrote an API Wrapper (wrote = about 60% of API methods implemented and no tests yet, pretty WIP) for the Wikia API. If someone wants to offer Feedback, implement some methods, or write Tests, you are free to do so.

Id love some feedback on the namespacing, as I wanted the API to be as close to the original API as possible, which is why every module (\`Activity\` for example) is a class, as that allows people to call \`api.activity.(somemethod)\` which looks better than \`api.activity(Somemethod)\`
## [3][Kotlin 1.4-M3-eap-* of the IDEA Plugin](https://www.reddit.com/r/Kotlin/comments/hljfbh/kotlin_14m3eap_of_the_idea_plugin/)
- url: https://www.reddit.com/r/Kotlin/comments/hljfbh/kotlin_14m3eap_of_the_idea_plugin/
---
I can see Kotlin 1.4-M3 artifacts are being pushed to the `kotlin-dev` channel.  

However I can't find a build of the IDEA Plugin. Any idea where I could look? I would like to avoid building it myself.
## [4][VisualVM with Kotlin?](https://www.reddit.com/r/Kotlin/comments/hlqex6/visualvm_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hlqex6/visualvm_with_kotlin/
---
Can I use VisualVM with Kotlin applications?
## [5][lndmflngs/line-to-telegram - Tool to prepare a set of Line stickers into Telegram stickers [Rewrote on the use of Executors, Need some code review]](https://www.reddit.com/r/Kotlin/comments/hlkmj5/lndmflngslinetotelegram_tool_to_prepare_a_set_of/)
- url: https://github.com/lndmflngs/line-to-telegram
---

## [6][Released KHipster 1.9.0](https://www.reddit.com/r/Kotlin/comments/hld2qu/released_khipster_190/)
- url: https://www.npmjs.com/package/generator-jhipster-kotlin
---

## [7][java.lang.ExceptionInInitializerError in self-referential objects?](https://www.reddit.com/r/Kotlin/comments/hlhuvu/javalangexceptionininitializererror_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hlhuvu/javalangexceptionininitializererror_in/
---
I'm having problems with the jvm compiler.

I'm trying to write a factory method for classes. The factory method has an init() block that helps to define behaviour for the new object. While this method compiles for JVM, I encounter a problem when running it:

java.lang.ExceptionInInitializerError

Caused by: java.lang.IllegalArgumentException: Parameter specified as non-null is null: method type.ProblematicKt.nullable, parameter $this$nullable

&amp;#x200B;

Apparently, the object isn't yet defined when I attempt to run the problematicInit() block. How do I fix this?

Seems to be a JVM problem. It seems to work for Javascript compilations. My understanding was that getProblematic would be hoisted, but what's inside the scope would be deferred until it's run designed to be run later - after the factory method is completed.

    interface ProblematicBuilderScope {
      fun problematicInit(getX: () -&gt; ProblematicInterface)
    }
    
    fun getProblematic() = X
    class Problematic(...): ProblematicInterface
    
    // Factory method with init() block
    val X = Problematic.factory(...) {
      problematicInit{ getProblematic() }
    } 
    
    fun factory(init: ProblematicBuilderScope.() -&gt; Unit): Problematic {
      val newObject = Problematic(...)
      val scope = ProblematicBuilderScope(newObject)
      scope.init()
      return newObject
    }
## [8][Can Kotlin.JS And Javascript be in the same project?](https://www.reddit.com/r/Kotlin/comments/hl1440/can_kotlinjs_and_javascript_be_in_the_same_project/)
- url: https://www.reddit.com/r/Kotlin/comments/hl1440/can_kotlinjs_and_javascript_be_in_the_same_project/
---
I was making some basic website after going through the React and Kotlin.JS tutorial of JetBrains and then my friend expressed his desire to contribute. But the problem is that he doesn't knows Kotlin and only Javascript and I have just started so I only know Javascript.
AFAIK, Java and Kotlin can be used in the same project but is there any way to use Kotlin.JS with simple Javascript. Or will one of us have to learn a new language?
## [9][Announcing LA4K: Logging API for Kotlin](https://www.reddit.com/r/Kotlin/comments/hkrz4x/announcing_la4k_logging_api_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hkrz4x/announcing_la4k_logging_api_for_kotlin/
---
Greetings!

I have been working on an original logging API for Kotlin Multiplatform for over a year now. I have done this because every other API I have found has really just been an extension of an existing JVM-only API, usually SLF4J. Right now, it can be bound to Android (Logcat), Log4j2, SLF4J and JUL. JS and native targets are planned for future releases.

I invite everyone to [check out its GitHub repo](https://github.com/wswartzendruber/la4k) (which includes a ReadMe). I have also [published the initial release, v0.1.0, to Bintray](https://bintray.com/wswartzendruber/la4k). I am especially interested in the community's assessment of the project, as any shortcommings need to be identified.

It is my intention for this to mature as Kotlin Multiplatform does the same. I would like for the v1.0.0 release to coincide with Multiplatform's initial, stable release.

I appreciate any feedback on this project.
## [10][Multiplatform Persistence with SQLDelight](https://www.reddit.com/r/Kotlin/comments/hkqf1e/multiplatform_persistence_with_sqldelight/)
- url: http://gh.jdoneill.com/2020/06/28/sqldelight/
---

