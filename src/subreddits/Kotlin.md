# Kotlin
## [1][No resources out there to learn from - Android + iOS Development with Kotlin Multi-platform](https://www.reddit.com/r/Kotlin/comments/gs2nc1/no_resources_out_there_to_learn_from_android_ios/)
- url: https://www.reddit.com/r/Kotlin/comments/gs2nc1/no_resources_out_there_to_learn_from_android_ios/
---
I have a passion to learn about building mobile applications using Kotlin. I would like them to run on both Android and iOS. As I understand Kotlin Multi-Platform is the way to go, but there are no resources for me to learn from out there. I tried searching in:

* PluralSight
* Udemy
* Udacity
* Google Search

Could you please suggest me a path that I should take in order to learn the skills required? To my knowledge I should use Kotlin Multi-Platform to create a Shared Library, then consume that library from XCode Swift (for iOS)  and I could then also use Android Studio for the Android application. Building Android applications with Kotlin seems to be very common so I have no trouble finding tutorials there, but I need your help for the iOS part. Thanks!
## [2][Kotlin/Native to avoid JNI](https://www.reddit.com/r/Kotlin/comments/gs42vh/kotlinnative_to_avoid_jni/)
- url: https://www.reddit.com/r/Kotlin/comments/gs42vh/kotlinnative_to_avoid_jni/
---
I recently discovered the great multi-platform capabilities of Kotlin/Native with Kotlin Multiplatform projects.
However, I want it to use it the other way around: Use Kotlin/Native's `cinterop` tool to generate Kotlin bindings for native C (or even Objective-C/Swift) libraries and use them in Kotlin directly.

I got this to work on my Mac. I built a Swift package into a `.a` file, generated a `klib` using `cinterop` and then linked that when building a simple `main.kt` program. Calling the Swift classes worked as expected.

I am currently trying to get this running on Android, where I face multiple issues. Aside from a Swift compiler for Android, I am trying to make Gradle build a `Klib` which I can call from my Android activity.
Building the klib for macOS works (so switching to ARM with correct compiler and so forth should work, too), but how to integrate it with Android?

Is this even possible? Is the klib compatible with calls from Kotlin, running on Android's ART runtime?
## [3][Talking Kotlin - Jetpack Compose with Leland Richardson](https://www.reddit.com/r/Kotlin/comments/gs5p1t/talking_kotlin_jetpack_compose_with_leland/)
- url: https://talkingkotlin.com/jetpack-compose-with-leland-richardson/
---

## [4][How do I compare a variable bitmask to several constant bitmasks elegantly in Kotlin?](https://www.reddit.com/r/Kotlin/comments/gs38tj/how_do_i_compare_a_variable_bitmask_to_several/)
- url: https://www.reddit.com/r/Kotlin/comments/gs38tj/how_do_i_compare_a_variable_bitmask_to_several/
---
I'm using the `android.bluetooth.BluetoothGattCharacteristic` API.   
`BluetoothGattCharacteristic.properties` is an `Int` bitmask, and I'm trying to see which of the API's properties it matches. This just feels wrong:

    val foo = bluetoothGattCharacteristic.properties
    if (foo and BluetoothGattCharacteristic.PROPERTY_WRITE &gt; 0) doBar()
    if (foo and BluetoothGattCharacteristic.PROPERTY_READ &gt; 0) doBaz()
    if (foo and BluetoothGattCharacteristic.PROPERTY_NOTIFY &gt; 0) doQux()
    //etc

I feel like there should be a `when` block here, but I can't figure out how I'm supposed to express it condition in Kotlin. Help?
## [5][FXGL game engine 11.9 release](https://www.reddit.com/r/Kotlin/comments/grrt73/fxgl_game_engine_119_release/)
- url: https://www.reddit.com/r/Kotlin/comments/grrt73/fxgl_game_engine_119_release/
---
FXGL is a Java/Kotlin game engine that targets the JVM, mobile (via gluon-client) and web (via jpro).

The Kotlin version of a simple game tutorial is available on [FXGL wiki](https://github.com/AlmasB/FXGL/wiki/Kotlin---Basic-Game-Example-(FXGL-11))

Published game demos are on [itch.io](https://fxgl.itch.io/)

Full source code of the engine and games can be found on [GitHub](https://github.com/AlmasB/FXGL)

Give it a go and join the [dev chat](https://gitter.im/AlmasB/FXGL) if there are any issues

[FXGL game demos](https://preview.redd.it/xu36bv8ebd151.jpg?width=1246&amp;format=pjpg&amp;auto=webp&amp;s=1fdf05778bbb011057ce676783283b9897014c3e)
## [6][How does Kotlin detect inappropriate blocking method calls?](https://www.reddit.com/r/Kotlin/comments/grxer4/how_does_kotlin_detect_inappropriate_blocking/)
- url: https://www.reddit.com/r/Kotlin/comments/grxer4/how_does_kotlin_detect_inappropriate_blocking/
---
I was recently working with a JedisPool in a suspending function and only found out from reading that Jedis is synchronous. If I understand correctly, that'd make using it blocking right? Do I misunderstand, or should Kotlin have warned it was inappropriate to use it within a suspending function? How does Kotlin detect something like that? Is it based on the exceptions it can throw (some testing showed IOException causes the warning)?
## [7][Kotlin Microservices With Spring Boot And Spring Cloud: Part 5 - Event-driven microservices](https://www.reddit.com/r/Kotlin/comments/groqvt/kotlin_microservices_with_spring_boot_and_spring/)
- url: https://youtu.be/fwhP9k0e1BY
---

## [8][Functional Programming Question About Kotlin](https://www.reddit.com/r/Kotlin/comments/grw9r2/functional_programming_question_about_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/grw9r2/functional_programming_question_about_kotlin/
---
So I used to love Kotlin, it was my favorite language. Recently, though, I've discovered Scala, and it changed how I think about writing code. The functional Scala way I have become quite fond of, and now I find myself trying to write code that way pretty much in everything I do.

Anyway, I've got Kotlin, and I'm trying to write some code. I've got a list, and I'm trying to fold() over it and re-build it into a map of lists:

    records.fold(mapOf&lt;String,List&lt;Record&gt;&gt;(), { acc, entry -&gt;
    	// What to do?
    })

Now, in Scala it would be something like this (I didn't compile this code, I wrote it on the fly):

    records.fold(Map[String,List[Record]](), (acc, entry) =&gt; {
    	acc ++ acc.get(entry.name) match {
    		case Some(list) =&gt; Map(entry.name -&gt; list + entry)
    		case None =&gt; Map(entry.name -&gt; List(entry))
    	}
    })

I'm trying to figure out the equivalent in Kotlin. My goals:

1) Use immutable Map/List.

2) Efficiently combine the collections. I've read that Kotlin has horrible performance in this area, unlike Scala.

3) Create a new entry in the map with a new list if it doesn't exist, append to the existing one if it does.

Basically just do that all in a simple, elegant, functional way. Essentially I still have a love of Kotlin, but this is my first time coming back to it after working with Scala, and it's just a very different experience. To be honest, I feel like even the Vavr library and Java would be able to get closer to my goal here.

But again, I feel like I haven't dived far enough into functional Kotlin, so I'm looking for being pointed in the right direction. If there are some good libraries (like Vavr for Java) that help, that's fine too.

PS. I know I could easily write this using mutable map/list. that's not the point.
## [9][Create progressive web apps with 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/gr4wz6/create_progressive_web_apps_with_100_kotlin/)
- url: https://github.com/grantas33/Kotlin-PWA-starter-kit
---

## [10][Getting the reference of delegated object](https://www.reddit.com/r/Kotlin/comments/gr74xp/getting_the_reference_of_delegated_object/)
- url: https://www.reddit.com/r/Kotlin/comments/gr74xp/getting_the_reference_of_delegated_object/
---
When using delegation, is it possible to gain access to the references of delegated objects?

&amp;#x200B;

    val b = B() // get access to this reference inside A
    val a = A(B)
    
    class A(delegation: B): B by delegation {
        fun something() {
            val x = this // reference to instance of A
            val y = (this as B) // same as x, how to get the original reference of b?    
        }
    }
