# Kotlin
## [1][Export Typescript definitions in 1.4-M1](https://www.reddit.com/r/Kotlin/comments/gqoacw/export_typescript_definitions_in_14m1/)
- url: https://www.reddit.com/r/Kotlin/comments/gqoacw/export_typescript_definitions_in_14m1/
---
I'm trying to build my shared library with the new backend (IR) in order to export Typescript definition files.

Before adding any annotation (such as JsExport), I just tried to build my project "as-is" with 1.4-M1 (without any modification) but I get this very strange error :

&amp;#x200B;

&amp;#x200B;

    java.lang.IllegalArgumentException: Collection contains more than one matching element. at org.jetbrains.kotlin.ir.backend.js.lower.EnumClassConstructorLowering.transformEnumConstructor(EnumClassLowering.kt:561) at org.jetbrains.kotlin.ir.backend.js.lower.EnumClassConstructorLowering.transformFlat(EnumClassLowering.kt:109) at org.jetbrains.kotlin.backend.common.LowerKt$transformFlatRestricted$1.invoke(Lower.kt:169) at org.jetbrains.kotlin.backend.common.LowerKt$transformFlatRestricted$1.invoke(Lower.kt) at org.jetbrains.kotlin.ir.declarations.StageController$DefaultImpls.restrictTo(PersistentApi.kt:29) at org.jetbrains.kotlin.ir.backend.js.CompilerKt$compile$1.restrictTo(compiler.kt:54) at org.jetbrains.kotlin.backend.common.LowerKt.transformFlatRestricted(Lower.kt:168) at org.jetbrains.kotlin.backend.common.LowerKt$runPostfix$1$transformFlat$1.visitClass(Lower.kt:240) at org.jetbrains.kotlin.ir.visitors.IrElementVisitorVoid$DefaultImpls.visitClass(IrElementVisitorVoid.kt:44) at org.jetbrains.kotlin.backend.common.LowerKt$runPostfix$1$transformFlat$1.visitClass(Lower.kt:184) at org.jetbrains.kotlin.backend.common.LowerKt$runPostfix$1$transformFlat$1.visitClass(Lower.kt:184) at org.jetbrains.kotlin.ir.declarations.impl.IrClassImpl.accept(IrClassImpl.kt:172) at org.jetbrains.kotlin.ir.visitors.IrElementVisitorVoidKt.acceptVoid(IrElementVisitorVoid.kt:267) at org.jetbrains.kotlin.backend.common.LowerKt$runPostfix$1.transformFlat(Lower.kt:184) at org.jetbrains.kotlin.backend.common.LowerKt$toFileLoweringPass$1.lower(Lower.kt:176) at org.jetbrains.kotlin.backend.common.DeclarationTransformer$DefaultImpls.lower(Lower.kt:163) at org.jetbrains.kotlin.ir.backend.js.lower.EnumClassConstructorLowering.lower(EnumClassLowering.kt:97) at org.jetbrains.kotlin.backend.common.LowerKt.lower(Lower.kt:68) at org.jetbrains.kotlin.backend.common.phaser.PhaseBuildersKt$makeIrModulePhase$1.invoke(PhaseBuilders.kt:240) at org.jetbrains.kotlin.backend.common.phaser.PhaseBuildersKt$makeIrModulePhase$1.invoke(PhaseBuilders.kt:233) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper$runBody$1.invoke(CompilerPhase.kt:128) at org.jetbrains.kotlin.backend.common.phaser.CompilerPhaseKt.downlevel(CompilerPhase.kt:24) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper.runBody(CompilerPhase.kt:127) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper.invoke(CompilerPhase.kt:105) at org.jetbrains.kotlin.backend.common.phaser.CompositePhase.invoke(PhaseBuilders.kt:30) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper$runBody$1.invoke(CompilerPhase.kt:128) at org.jetbrains.kotlin.backend.common.phaser.CompilerPhaseKt.downlevel(CompilerPhase.kt:24) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper.runBody(CompilerPhase.kt:127) at org.jetbrains.kotlin.backend.common.phaser.AbstractNamedPhaseWrapper.invoke(CompilerPhase.kt:105) at org.jetbrains.kotlin.backend.common.phaser.CompilerPhaseKt.invokeToplevel(CompilerPhase.kt:42) at org.jetbrains.kotlin.ir.backend.js.CompilerKt.compile(compiler.kt:118) at org.jetbrains.kotlin.ir.backend.js.CompilerKt.compile$default(compiler.kt:52) at org.jetbrains.kotlin.cli.js.K2JsIrCompiler.doExecute(K2JsIrCompiler.kt:218) at org.jetbrains.kotlin.cli.js.K2JSCompiler.doExecute(K2JSCompiler.java:192) at org.jetbrains.kotlin.cli.js.K2JSCompiler.doExecute(K2JSCompiler.java:85) at org.jetbrains.kotlin.cli.common.CLICompiler.execImpl(CLICompiler.kt:86) at org.jetbrains.kotlin.cli.common.CLICompiler.execImpl(CLICompiler.kt:44) at org.jetbrains.kotlin.cli.common.CLITool.exec(CLITool.kt:105) at org.jetbrains.kotlin.daemon.CompileServiceImpl.compile(CompileServiceImpl.kt:1558) at jdk.internal.reflect.GeneratedMethodAccessor26.invoke(Unknown Source) at java.base/jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43) at java.base/java.lang.reflect.Method.invoke(Method.java:566) at java.rmi/sun.rmi.server.UnicastServerRef.dispatch(UnicastServerRef.java:359) at java.rmi/sun.rmi.transport.Transport$1.run(Transport.java:200) at java.rmi/sun.rmi.transport.Transport$1.run(Transport.java:197) at java.base/java.security.AccessController.doPrivileged(Native Method) at java.rmi/sun.rmi.transport.Transport.serviceCall(Transport.java:196) at java.rmi/sun.rmi.transport.tcp.TCPTransport.handleMessages(TCPTransport.java:562) at java.rmi/sun.rmi.transport.tcp.TCPTransport$ConnectionHandler.run0(TCPTransport.java:796) at java.rmi/sun.rmi.transport.tcp.TCPTransport$ConnectionHandler.lambda$run$0(TCPTransport.java:677) at java.base/java.security.AccessController.doPrivileged(Native Method) at java.rmi/sun.rmi.transport.tcp.TCPTransport$ConnectionHandler.run(TCPTransport.java:676) at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128) at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628) at java.base/java.lang.Thread.run(Thread.java:834)

&amp;#x200B;

~~Everything works well if I go back to Kotlin 1.3 &amp; Kotlin multiplatform 1.3.7X~~

EDIT : Just by removing the flag

    kotlin.js.compiler=ir 

in

    gradle.settings

&amp;#x200B;

Fixes the build.

So it is not related to the plugins (even when using 1.4-M1 version)

    id 'org.jetbrains.kotlin.multiplatform' version '1.4-M1'
    id "org.jetbrains.kotlin.plugin.noarg" version "1.4-M1"
    id 'org.jetbrains.kotlin.plugin.allopen' version '1.4-M1'

Of course, I need to keep the IR back-end to be able to do  d.ts export

&amp;#x200B;

Any ideas ?
## [2][Issue with Jetbrain's kotlin jvm maven archetype.](https://www.reddit.com/r/Kotlin/comments/gqt927/issue_with_jetbrains_kotlin_jvm_maven_archetype/)
- url: https://www.reddit.com/r/Kotlin/comments/gqt927/issue_with_jetbrains_kotlin_jvm_maven_archetype/
---
Started learning Kotlin last week. I'm using VS code and generated a maven project using Jetbrain's Kotlin single module jvm [archetype](https://mvnrepository.com/artifact/org.jetbrains.kotlin/kotlin-archetype-jvm).

Code compiles and jars are built. But when I run: `java -jar app.jar` I get: `no main manifest attribute, in .\app-1.0-SNAPSHOT.jar`

I've also tried the build plugin referred by [them](https://kotlinlang.org/docs/reference/using-maven.html#self-contained-jar-file)  to generate a self contained jar with "-with-dependencies.jar" name suffix. It says to point main class name in the plugin. But the main function being a top level function, I don't know what values to provide there. Nonetheless, even the built jar doesn't run and says "no main manifest..."

I've looked online and everyone suggests using IDEA. It is very heavy for my machine. Is there anyway I can rectify this issue?
## [3][Designing and Working with Single View States on Android - zsmb.co](https://www.reddit.com/r/Kotlin/comments/gqfe6n/designing_and_working_with_single_view_states_on/)
- url: https://zsmb.co/designing-and-working-with-single-view-states-on-android/
---

## [4][Do you know any framework to build Apps for Linux Arm coding with Kotlin?](https://www.reddit.com/r/Kotlin/comments/gqf4co/do_you_know_any_framework_to_build_apps_for_linux/)
- url: https://www.reddit.com/r/Kotlin/comments/gqf4co/do_you_know_any_framework_to_build_apps_for_linux/
---

## [5][kotlin programming: The Big Nerd Ranch Guide Book](https://www.reddit.com/r/Kotlin/comments/gqh3go/kotlin_programming_the_big_nerd_ranch_guide_book/)
- url: https://www.reddit.com/r/Kotlin/comments/gqh3go/kotlin_programming_the_big_nerd_ranch_guide_book/
---
\#kotlin u/kotlin   
Is thi book good for someone who never programming in java ? or java is necessary ?

https://preview.redd.it/gvgt9wu5my051.png?width=360&amp;format=png&amp;auto=webp&amp;s=76dab4a99ebb8c4932504a883f3efc9742224782
## [6][viewModels&lt;&gt;() vs activityViewModels&lt;&gt;() vs createViewModelLazy()](https://www.reddit.com/r/Kotlin/comments/gq5wjk/viewmodels_vs_activityviewmodels_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/gq5wjk/viewmodels_vs_activityviewmodels_vs/
---
Recently I've been going thorough the **KTX** libraries and I found several ways to instantiate VM using ktx. 

I tried following up on the differences between by **viewModels&lt;&gt;()**, **activityViewModels&lt;&gt;()** and **createViewModelLazy()** but couldn't figure out much.  

While all this time, I've been using **KOIN viewmodel artifact** for initializing my VMs, I am just keen to know the differences between the above mentioned.
## [7][Need help with this confusion in data binding](https://www.reddit.com/r/Kotlin/comments/gq6r2e/need_help_with_this_confusion_in_data_binding/)
- url: https://www.reddit.com/r/Kotlin/comments/gq6r2e/need_help_with_this_confusion_in_data_binding/
---
So i am still learning about Android app development with Kotlin using a Udacity course.

I was learning about data binding in the course and they used the following code;

&amp;#x200B;

My name.kt

    package com.example.aboutme
    data class MyName(var name: String = "",var nickname: String = "")

&amp;#x200B;

MainActivity.kt

    import androidx.databinding.DataBindingUtil
    import com.rackaapps.aboutme.databinding.ActivityMainBinding
    
    class MainActivity : AppCompatActivity() {
    
        private lateinit var binding: ActivityMainBinding
        private val myName = MyName("Racka Legit")

&amp;#x200B;

activity\_main.xml (for a text view that's supposed to show the name &amp; one for nickname)

    &lt;TextView
                android:id="@+id/name_text"
                style="@style/NameStyle"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:text="@={myName.name}"
                android:textAlignment="center" /&gt;
    
    &lt;TextView
                android:id="@+id/nickname_text"
                style="@style/NameStyle"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:textAlignment="center"
                android:visibility="gone"
                android:text="@={myName.nickname}"/&gt;

&amp;#x200B;

Now with this code i get an error saying `"cannot find symbol class ActivityMainBindingImpl"`

And its telling me to `"import com.example.aboutme.databinding.ActivityMainBindingImpl"` which gives an error too.

**Note that,**

    dataBinding {enabled = true}

was enabled too.

&amp;#x200B;

**After some Googling i found out that i should use;**

    android:text="@{myName.name}"

&amp;

    android:text="@{myName.nickname}"

**Instead of;**

    android:text="@={myName.name}"

&amp;

    android:text="@={myName.nickname}"

and it worked like charm!

&amp;#x200B;

Now i want to know what difference does the equal (=) make in data binding in the context of the code above?

And why did the code used in the course work for them but didn't work for me?

I didn't find any info about the use of this yet.
## [8][Annotation processor Extensions](https://www.reddit.com/r/Kotlin/comments/gppuj9/annotation_processor_extensions/)
- url: https://www.reddit.com/r/Kotlin/comments/gppuj9/annotation_processor_extensions/
---
I created a repo with the intention of gathering useful extensions/utils for anyone working with Kotlin annotation processing and code generation. Have a look.

https://github.com/iFanie/KTAP

Also, what extension/utils do you find useful for Kotlin annotation processor projects?
## [9][Understand Kotlin Collection Function Past Tense](https://www.reddit.com/r/Kotlin/comments/gpk2w7/understand_kotlin_collection_function_past_tense/)
- url: https://medium.com/@elye.project/understand-kotlin-collection-function-past-tense-59f592af9436?source=friends_link&amp;sk=e22ccd272ebd28ef6f417e7b455b5b4f
---

## [10][How to secure Ktor web app with Keycloak Jetty 9.x Adapters?](https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/)
- url: https://www.reddit.com/r/Kotlin/comments/gpnxq8/how_to_secure_ktor_web_app_with_keycloak_jetty_9x/
---
Hi all 

I would like to secure the Ktor webapp, that is run on top of Jetty server with [Keycloak Jetty 9.x Adapters][1]. 

Ktor provides a [hook][2] for Jetty server initialization and maybe it is the right place in integrate the Keycloak Jetty 9.x Adapters.  

How to integrate Keycloak Jetty 9.x Adapters into Ktor app?  

Thanks


  [1]: https://www.keycloak.org/docs/latest/securing_apps/index.html#_jetty9_adapter
  [2]: https://ktor.io/servers/configuration.html#jetty
