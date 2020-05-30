# Kotlin
## [1][TechEmpower released round 19 of their server middleware benchmark (including Kotlin ones)](https://www.reddit.com/r/Kotlin/comments/gsuvvo/techempower_released_round_19_of_their_server/)
- url: https://www.reddit.com/r/Kotlin/comments/gsuvvo/techempower_released_round_19_of_their_server/
---
TechEmpower just released their latest round (19) of their server middleware benchmark. You can check the results here: https://www.techempower.com/benchmarks
## [2][I Need Help Adding To Hashmap From a Text Input](https://www.reddit.com/r/Kotlin/comments/gt99bu/i_need_help_adding_to_hashmap_from_a_text_input/)
- url: https://www.reddit.com/r/Kotlin/comments/gt99bu/i_need_help_adding_to_hashmap_from_a_text_input/
---
Can anyone help me add functionality to my program to add items to my hashmap from a text input?

&amp;#x200B;

var products = *hashMapOf*&lt;Int, Pair&lt;String, Double&gt;&gt;(  
 111 *to* Pair("shoes", 59.99),  
 222 *to* Pair("shirt", 19.99),  
 333 *to* Pair("socks", 3.99)

*print*("Please Enter The Description Of The New Item: ")  


var addItemDescription = *readLine*()!!  


*print*("Please Enter The Price For $addItemDescription: ")  


var addItemPrice = *readLine*()!!.*toFloat*()
## [3][Codota's free ai code completion plugin now supports kotlin for android studio and eclipse.](https://www.reddit.com/r/Kotlin/comments/gt0u32/codotas_free_ai_code_completion_plugin_now/)
- url: https://www.reddit.com/r/Kotlin/comments/gt0u32/codotas_free_ai_code_completion_plugin_now/
---
Has anyone tried this plugin with Kotlin? I'm curious about your thoughts on it
## [4][Java interop: Kotlin gets tripped up by capitalization.](https://www.reddit.com/r/Kotlin/comments/gsw920/java_interop_kotlin_gets_tripped_up_by/)
- url: https://www.reddit.com/r/Kotlin/comments/gsw920/java_interop_kotlin_gets_tripped_up_by/
---
The java.net.URLClassLoader class provides a getURLs() method:

&gt; public URL[] getURLs()

&gt; Returns the search path of URLs for loading classes and resources. This includes the original list of URLs specified to the constructor, along with any URLs subsequently appended by the addURL() method.

And the Kotlin documentation claims:

&gt; Methods that follow the Java conventions for getters and setters (no-argument methods with names starting with get and single-argument methods with names starting with set) are represented as properties in Kotlin. Boolean accessor methods (where the name of the getter starts with is and the name of the setter starts with set) are represented as properties which have the same name as the getter method.

Well, all righty then, I should have a .URLs property I can use. But no:

    &gt;&gt;&gt; import java.net.URLClassLoader
    &gt;&gt;&gt; val cl = ClassLoader.getSystemClassLoader()
    &gt;&gt;&gt; cl is URLClassLoader
    res4: kotlin.Boolean = true
    &gt;&gt;&gt; (cl as URLClassLoader).URLs
    error: unresolved reference: URLs
    (cl as URLClassLoader).URLs
                           ^

However, the old reliable fallback works:

    &gt;&gt;&gt; (cl as URLClassLoader).getURLs()
    res6: kotlin.Array&lt;(out) java.net.URL!&gt;! = [Ljava.net.URL;@5b6ac24b
    
Of course, the rule is that .getFoo() maps to a .foo property. Note the downcasing; the property on the Kotlin side is called .foo and not .Foo. So maybe it’s called .uRLs? Close, but no cigar!

    &gt;&gt;&gt; (cl as URLClassLoader).uRLs
    error: unresolved reference: uRLs
    (cl as URLClassLoader).uRLs
                           ^
    &gt;&gt;&gt; (cl as URLClassLoader).urLs
    res8: kotlin.Array&lt;(out) java.net.URL!&gt;! = [Ljava.net.URL;@5dddd61f

Ugh. Maybe I’ll just stick to calling the Java-style getters and setters.
## [5][How To Prevent ERP Implementation Failure](https://www.reddit.com/r/Kotlin/comments/gtbbta/how_to_prevent_erp_implementation_failure/)
- url: https://www.odooimplementation.com/blog/erp_implementation_failure
---

## [6][Kotlin, Eclipse and Executable JARs](https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/)
- url: https://www.reddit.com/r/Kotlin/comments/gsoqrd/kotlin_eclipse_and_executable_jars/
---
Sooo, I just wanted to compile my Kotlin application into an executable JAR file. No big deal, right? I didn't do that too often when developing with Java, but how much easier than "Export to Executable JAR" can it get?

So that's what I do. And end up with an completely empty error message. The file still got created, but executing it does nothing. 

I'm thinking my IDE (Eclipse) doesn't work all that well with Kotlin - it's not exactly a great IDE after all, and I only use it because I'm used to it and too lazy to get into another one. But maybe, since Kotlin runs in the JVM, I can just export it as a Java Application. A quick test at least shows, that I can successfully run it as a Java Application from within the IDE.  
But same result: empty error message and a file that does nothing.

So like the good little programmer I am, I hop onto google. 3 pages filled with nothing but "use Gradle", only with the occasional "Maven works too" sprinkled in for flavor. Nothing else I could find.

&amp;#x200B;

So - is there really no way of creating an Executable JAR without using Gradle?
## [7][Do I need an Android phone to learn Kotlin?](https://www.reddit.com/r/Kotlin/comments/gsuq6h/do_i_need_an_android_phone_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gsuq6h/do_i_need_an_android_phone_to_learn_kotlin/
---
So I'm interested in learning Kotlin but... I don't have an Android phone. I know I can learn the language just like any other language, but I'm curious if it's pointless to learn it if I can't test it out on an actual phone. Not sure how elaborate Android Studio is. 

Do I need to get an Android phone to learn Kotlin?
## [8][Host your Github Actions generated Dokka documentation on Netlify](https://www.reddit.com/r/Kotlin/comments/gsnt68/host_your_github_actions_generated_dokka/)
- url: https://pgreze.dev/posts/2020-05-28-static-doc-netlify/
---

## [9][Good Kotlin Video courses](https://www.reddit.com/r/Kotlin/comments/gscxdr/good_kotlin_video_courses/)
- url: https://www.reddit.com/r/Kotlin/comments/gscxdr/good_kotlin_video_courses/
---
Hi!

I'm searching for good Video Courses to learn Kotlin. I have experience in Dart/Flutter, Javascript and Typescript. Does anyone know some good Video courses on Udemy or similar platforms? I mean a little bit more than Youtube Walkthroughs. It should cover a little bit more than just build an Android App because I want also little bit knowledge for building something like REST Endpoints.

Thanks!
## [10][Bug in Kotlin 1.3.72 (java.lang.AssertionError: Built-in class kotlin.Any is not found)?](https://www.reddit.com/r/Kotlin/comments/gs9ywj/bug_in_kotlin_1372_javalangassertionerror_builtin/)
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
