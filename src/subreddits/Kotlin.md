# Kotlin
## [1][Kotlin, Eclipse and Executable JARs](https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/)
- url: https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/
---
Sooo, I just wanted to compile my Kotlin application into an executable JAR file. No big deal, right? I didn't do that too often when developing with Java, but how much easier than "Export to Executable JAR" can it get?

So that's what I do. And end up with an completely empty error message. The file still got created, but executing it does nothing. 

I'm thinking my IDE (Eclipse) doesn't work all that well with Kotlin - it's not exactly a great IDE after all, and I only use it because I'm used to it and too lazy to get into another one. But maybe, since Kotlin runs in the JVM, I can just export it as a Java Application. A quick test at least shows, that I can successfully run it as a Java Application from within the IDE.  
But same result: empty error message and a file that does nothing.

So like the good little programmer I am, I hop onto google. 3 pages filled with nothing but "use Gradle", only with the occasional "Maven works too" sprinkled in for flavor. Nothing else I could find.

&amp;#x200B;

So - is there really no way of creating an Executable JAR without using Gradle?
## [2][Good Kotlin Video courses](https://www.reddit.com/r/Kotlin/comments/gscxdr/good_kotlin_video_courses/)
- url: https://www.reddit.com/r/Kotlin/comments/gscxdr/good_kotlin_video_courses/
---
Hi!

I'm searching for good Video Courses to learn Kotlin. I have experience in Dart/Flutter, Javascript and Typescript. Does anyone know some good Video courses on Udemy or similar platforms? I mean a little bit more than Youtube Walkthroughs. It should cover a little bit more than just build an Android App because I want also little bit knowledge for building something like REST Endpoints.

Thanks!
## [3][Host your Github Actions generated Dokka documentation on Netlify](https://www.reddit.com/r/Kotlin/comments/gsnt68/host_your_github_actions_generated_dokka/)
- url: https://pgreze.dev/posts/2020-05-28-static-doc-netlify/
---

## [4][Bug in Kotlin 1.3.72 (java.lang.AssertionError: Built-in class kotlin.Any is not found)?](https://www.reddit.com/r/Kotlin/comments/gs9ywj/bug_in_kotlin_1372_javalangassertionerror_builtin/)
- url: https://www.reddit.com/r/Kotlin/comments/gs9ywj/bug_in_kotlin_1372_javalangassertionerror_builtin/
---
Or maybe it’s not a bug, and I’m not using reflection property. Digression: I generally like Kotlin, but its reflection support has to be the least-favorite part of the language I have run across so far. Poorly-documented and outright clunky in places.

    Exception in thread "main" java.lang.AssertionError: Built-in class kotlin.Any is not found
            at kotlin.reflect.jvm.internal.impl.builtins.KotlinBuiltIns$3.invoke(KotlinBuiltIns.java:113)
            at kotlin.reflect.jvm.internal.impl.builtins.KotlinBuiltIns$3.invoke(KotlinBuiltIns.java:108)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$MapBasedMemoizedFunction.invoke(LockBasedStorageManager.java:512)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$MapBasedMemoizedFunctionToNotNull.invoke(LockBasedStorageManager.java:587)
            at kotlin.reflect.jvm.internal.impl.builtins.KotlinBuiltIns.getBuiltInClassByName(KotlinBuiltIns.java:362)
            at kotlin.reflect.jvm.internal.impl.builtins.KotlinBuiltIns.getAny(KotlinBuiltIns.java:367)
            at kotlin.reflect.jvm.internal.impl.builtins.KotlinBuiltIns.getAnyType(KotlinBuiltIns.java:642)
            at kotlin.reflect.jvm.internal.impl.descriptors.NotFoundClasses$MockClassDescriptor.&lt;init&gt;(NotFoundClasses.kt:60)
            at kotlin.reflect.jvm.internal.impl.descriptors.NotFoundClasses$classes$1.invoke(NotFoundClasses.kt:44)
            at kotlin.reflect.jvm.internal.impl.descriptors.NotFoundClasses$classes$1.invoke(NotFoundClasses.kt:22)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$MapBasedMemoizedFunction.invoke(LockBasedStorageManager.java:512)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$MapBasedMemoizedFunctionToNotNull.invoke(LockBasedStorageManager.java:587)
            at kotlin.reflect.jvm.internal.impl.descriptors.NotFoundClasses.getClass(NotFoundClasses.kt:92)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.TypeDeserializer$typeConstructor$1.invoke(TypeDeserializer.kt:109)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.TypeDeserializer.typeConstructor(TypeDeserializer.kt:113)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.TypeDeserializer.simpleType(TypeDeserializer.kt:75)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.TypeDeserializer.type(TypeDeserializer.kt:63)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedClassDescriptor$DeserializedClassTypeConstructor.computeSupertypes(DeserializedClassDescriptor.kt:180)
            at kotlin.reflect.jvm.internal.impl.types.AbstractTypeConstructor$supertypes$1.invoke(AbstractTypeConstructor.kt:80)
            at kotlin.reflect.jvm.internal.impl.types.AbstractTypeConstructor$supertypes$1.invoke(AbstractTypeConstructor.kt:26)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedLazyValue.invoke(LockBasedStorageManager.java:355)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedLazyValueWithPostCompute.invoke(LockBasedStorageManager.java:428)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedNotNullLazyValueWithPostCompute.invoke(LockBasedStorageManager.java:459)
            at kotlin.reflect.jvm.internal.impl.types.AbstractTypeConstructor.getSupertypes(AbstractTypeConstructor.kt:27)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedClassDescriptor$DeserializedClassMemberScope.getNonDeclaredVariableNames(DeserializedClassDescriptor.kt:296)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope$variableNamesLazy$2.invoke(DeserializedMemberScope.kt:77)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope$variableNamesLazy$2.invoke(DeserializedMemberScope.kt:40)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedLazyValue.invoke(LockBasedStorageManager.java:355)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedNotNullLazyValue.invoke(LockBasedStorageManager.java:474)
            at kotlin.reflect.jvm.internal.impl.storage.StorageKt.getValue(storage.kt:42)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope.getVariableNamesLazy(DeserializedMemberScope.kt)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope.getVariableNames(DeserializedMemberScope.kt:85)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope.addFunctionsAndProperties(DeserializedMemberScope.kt:210)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedMemberScope.computeDescriptors(DeserializedMemberScope.kt:181)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedClassDescriptor$DeserializedClassMemberScope$allDescriptors$1.invoke(DeserializedClassDescriptor.kt:217)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedClassDescriptor$DeserializedClassMemberScope$allDescriptors$1.invoke(DeserializedClassDescriptor.kt:210)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedLazyValue.invoke(LockBasedStorageManager.java:355)
            at kotlin.reflect.jvm.internal.impl.storage.LockBasedStorageManager$LockBasedNotNullLazyValue.invoke(LockBasedStorageManager.java:474)
            at kotlin.reflect.jvm.internal.impl.serialization.deserialization.descriptors.DeserializedClassDescriptor$DeserializedClassMemberScope.getContributedDescriptors(DeserializedClassDescriptor.kt:227)
            at kotlin.reflect.jvm.internal.impl.resolve.scopes.ResolutionScope$DefaultImpls.getContributedDescriptors$default(ResolutionScope.kt:52)
            at kotlin.reflect.jvm.internal.KDeclarationContainerImpl.getMembers(KDeclarationContainerImpl.kt:65)
            at kotlin.reflect.jvm.internal.KClassImpl$Data$declaredNonStaticMembers$2.invoke(KClassImpl.kt:159)
            at kotlin.reflect.jvm.internal.KClassImpl$Data$declaredNonStaticMembers$2.invoke(KClassImpl.kt:44)
            at kotlin.reflect.jvm.internal.ReflectProperties$LazySoftVal.invoke(ReflectProperties.java:92)
            at kotlin.reflect.jvm.internal.ReflectProperties$Val.getValue(ReflectProperties.java:31)
            at kotlin.reflect.jvm.internal.KClassImpl$Data.getDeclaredNonStaticMembers(KClassImpl.kt)
            at kotlin.reflect.jvm.internal.KClassImpl$Data$allNonStaticMembers$2.invoke(KClassImpl.kt:168)
            at kotlin.reflect.jvm.internal.KClassImpl$Data$allNonStaticMembers$2.invoke(KClassImpl.kt:44)
            at kotlin.reflect.jvm.internal.ReflectProperties$LazySoftVal.invoke(ReflectProperties.java:92)
            at kotlin.reflect.jvm.internal.ReflectProperties$Val.getValue(ReflectProperties.java:31)
            at kotlin.reflect.jvm.internal.KClassImpl$Data.getAllNonStaticMembers(KClassImpl.kt)
            at kotlin.reflect.full.KClasses.getMemberProperties(KClasses.kt:149)
            at kotlin.reflect.jvm.ReflectJvmMapping.getKotlinProperty(ReflectJvmMapping.kt:100)
            at name.blackcap.xtest.XtestKt.getVars(Xtest.kt:99)
            at name.blackcap.xtest.XtestKt.main(Xtest.kt:54)



Here’s the code that prompts the error (it’s barfing in the field.kotlinProperty?.let expression):

    fun getVars(cls: Class&lt;*&gt;): String {
        val inst = cls.getField("INSTANCE").get(cls)
        var needsDelim = false
        return StringBuilder().apply {
            for (field in cls.declaredFields) {
                field.kotlinProperty?.let { member -&gt;
                    member.getter.call(inst)?.let { value -&gt;
                        if (value is String || value is Number) {
                            if (needsDelim)
                                append(';')
                            needsDelim = true
                            append("set ")
                            append(member.name)
                            append(' ')
                            if (value is String)
                                tclQuote(this, value)
                            else
                                append(value.toString())
                        } else {
                            System.err.println("${MYNAME}: warning - ${member.name} not string or number, ignored")
                        }
                    }
                }
            }
        }.toString()
    }
## [5][Talking Kotlin - Jetpack Compose with Leland Richardson](https://www.reddit.com/r/Kotlin/comments/gs5p1t/talking_kotlin_jetpack_compose_with_leland/)
- url: https://talkingkotlin.com/jetpack-compose-with-leland-richardson/
---

## [6][Object - Kotlin Vocabulary](https://www.reddit.com/r/Kotlin/comments/gs75fd/object_kotlin_vocabulary/)
- url: https://www.youtube.com/watch?v=KUk6k865Vgg
---

## [7][Kotlin/Native to avoid JNI](https://www.reddit.com/r/Kotlin/comments/gs42vh/kotlinnative_to_avoid_jni/)
- url: https://www.reddit.com/r/Kotlin/comments/gs42vh/kotlinnative_to_avoid_jni/
---
I recently discovered the great multi-platform capabilities of Kotlin/Native with Kotlin Multiplatform projects.
However, I want it to use it the other way around: Use Kotlin/Native's `cinterop` tool to generate Kotlin bindings for native C (or even Objective-C/Swift) libraries and use them in Kotlin directly.

I got this to work on my Mac. I built a Swift package into a `.a` file, generated a `klib` using `cinterop` and then linked that when building a simple `main.kt` program. Calling the Swift classes worked as expected.

I am currently trying to get this running on Android, where I face multiple issues. Aside from a Swift compiler for Android, I am trying to make Gradle build a `Klib` which I can call from my Android activity.
Building the klib for macOS works (so switching to ARM with correct compiler and so forth should work, too), but how to integrate it with Android?

Is this even possible? Is the klib compatible with calls from Kotlin, running on Android's ART runtime?
## [8][No resources out there to learn from - Android + iOS Development with Kotlin Multi-platform](https://www.reddit.com/r/Kotlin/comments/gs2nc1/no_resources_out_there_to_learn_from_android_ios/)
- url: https://www.reddit.com/r/Kotlin/comments/gs2nc1/no_resources_out_there_to_learn_from_android_ios/
---
I have a passion to learn about building mobile applications using Kotlin. I would like them to run on both Android and iOS. As I understand Kotlin Multi-Platform is the way to go, but there are no resources for me to learn from out there. I tried searching in:

* PluralSight
* Udemy
* Udacity
* Google Search

Could you please suggest me a path that I should take in order to learn the skills required? To my knowledge I should use Kotlin Multi-Platform to create a Shared Library, then consume that library from XCode Swift (for iOS)  and I could then also use Android Studio for the Android application. Building Android applications with Kotlin seems to be very common so I have no trouble finding tutorials there, but I need your help for the iOS part. Thanks!
## [9][How do I compare a variable bitmask to several constant bitmasks elegantly in Kotlin?](https://www.reddit.com/r/Kotlin/comments/gs38tj/how_do_i_compare_a_variable_bitmask_to_several/)
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
## [10][Are ranges sequences?](https://www.reddit.com/r/Kotlin/comments/gs8aou/are_ranges_sequences/)
- url: https://www.reddit.com/r/Kotlin/comments/gs8aou/are_ranges_sequences/
---
So when writing collection code, I prefer to focus on lazy iteration. This allows for leveraging the intermediate vs terminal operation behavior, similar to Java's Stream. My understanding is that Sequences are the only collection type in Kotlin that has lazy iteration and intermediate operations, so I prefer either calling toSequence() on collections or building sequences to begin with.

So for my question, are ranges sequence? Is the collection built by rangeTo or rangeUtil a sequence?
