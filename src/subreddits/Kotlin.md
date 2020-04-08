# Kotlin
## [1][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fwnsyu/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [2][The Kotlin discord](https://www.reddit.com/r/Kotlin/comments/fwu8vh/the_kotlin_discord/)
- url: https://www.reddit.com/r/Kotlin/comments/fwu8vh/the_kotlin_discord/
---
Heya, just wondering if anyone has an invite for the Kotlin discord server? It'd be more convenient for me if i was able to use it
## [3][Python style sort with key?](https://www.reddit.com/r/Kotlin/comments/fwtyva/python_style_sort_with_key/)
- url: https://www.reddit.com/r/Kotlin/comments/fwtyva/python_style_sort_with_key/
---
In python you can take a list and sort it by applying a "key" function that determines the sorted order. Is there something similar in Kotlin other than something like this?
```
list.map { elem -&gt; 
    MyDataClass(elem, key(elem)) 
}
    .sortedBy{it.key}
    .map {it.elem}
```
Thanks!
## [4][[Question] Overload resolution ambiguity](https://www.reddit.com/r/Kotlin/comments/fwohn2/question_overload_resolution_ambiguity/)
- url: https://www.reddit.com/r/Kotlin/comments/fwohn2/question_overload_resolution_ambiguity/
---
https://preview.redd.it/bj8e691rifr41.png?width=794&amp;format=png&amp;auto=webp&amp;s=0062da1a51bbcf2ad26efb28d66a1c4bfc4f21cf

From [this answer](https://stackoverflow.com/questions/47884934/im-getting-overload-resolution-ambiguity-error-on-kotlin-safe-call) I think this might be due to smart casting. But if that is the case, why is it only happening inside the if block (line: 42) and not outside (line: 39)?

  
**Code:**

    fun main(){
        var greeting: String? = "Hello"
        greeting = null
    
        println(greeting)
    
        if(greeting!=null){
            println(greeting)
        }
    }
## [5][What UI toolkit to use?](https://www.reddit.com/r/Kotlin/comments/fwmusp/what_ui_toolkit_to_use/)
- url: https://www.reddit.com/r/Kotlin/comments/fwmusp/what_ui_toolkit_to_use/
---
First time coding in Kotlin, and I'm creating an app for both iOS and Android. I've decided to go with Jetpack for the Android UI, but what do I use for iOS? Is Flutter an option? Because the common consensus is that I'll be coding in Dart if I use it. 

Do I go with the default SwiftUI, or is there an alternative?  
Also, I understand if this question sounds completely asinine - I've never coded mobile apps before, so all help is appreciated, thanks!
## [6][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fw3wkt/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [7][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fwxmmc/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [8][How to Kotlin/Native KDocs documentation in built .frameworks for iOS/XCode?](https://www.reddit.com/r/Kotlin/comments/fwnvjj/how_to_kotlinnative_kdocs_documentation_in_built/)
- url: https://www.reddit.com/r/Kotlin/comments/fwnvjj/how_to_kotlinnative_kdocs_documentation_in_built/
---
I'm developing a Kotlin/Native library for both iOS and Android.

When running ./gradlew assemble I get the release/debug .frameworks with Obj-C Headers.

I can use these .frameworks fine. But I would also like to carry on with documentation of classes and functions when using the code on XCode.

If I use KDoc and document my code:

    /**
     * This is a comment about the class I developed.
     * [BaseService] is a class.
     */
    class BaseService {
    
        /**
         * Comments about this val
         */
        val greeting = "Hello"
    
        /**
         * This function greets the user
         */
        fun greetUser(){
            println(greeting)
        }
    
    }

All of this is lost when bringing it to my iOS project. Because the .framework don't carry the KDocs with it.

Is there any way to carry the documentation over to the Obj-C Headers or other auxiliary files when assembling the .framework files? Maybe some Gradle configuration I am missing?

I also asked on StackOverflow. Maybe someone in can grab it:

[https://stackoverflow.com/questions/61082784/include-kotlin-native-kdocs-documentation-in-built-frameworks-for-ios-xcode](https://stackoverflow.com/questions/61082784/include-kotlin-native-kdocs-documentation-in-built-frameworks-for-ios-xcode)  


Edit: That was fast.  
Question was answered: [https://github.com/JetBrains/kotlin-native/issues/4000](https://github.com/JetBrains/kotlin-native/issues/4000)  

## [9][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fwt7dr/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [10][Will Kotlin ever be able to compile to a stand-alone exe?](https://www.reddit.com/r/Kotlin/comments/fwch4x/will_kotlin_ever_be_able_to_compile_to_a/)
- url: https://www.reddit.com/r/Kotlin/comments/fwch4x/will_kotlin_ever_be_able_to_compile_to_a/
---
I really like the syntax of Kotlin and I'm wondering if it'll ever be possible to easily compile to a stand-alone executable (like you'd get with "gcc myprog.c").  I'm aware I can do things like append the .jar to a shell script to create an "executable", but that feels pretty clunky.
