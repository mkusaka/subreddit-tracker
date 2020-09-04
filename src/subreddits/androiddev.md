# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/imf7bk/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/imf7bk/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - August 31, 2020](https://www.reddit.com/r/androiddev/comments/ijvutk/weekly_questions_thread_august_31_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ijvutk/weekly_questions_thread_august_31_2020/
---
This thread is for simple questions that don't warrant their own thread (although we suggest checking the sidebar, [the wiki](http://www.reddit.com/r/androiddev/wiki/), [our Discord](https://discord.gg/D2cNrqX), or [Stack Overflow](http://stackoverflow.com) before posting). Examples of questions:

* How do I pass data between my Activities?
* Does anyone have a link to the source for the AOSP messaging app?
* Is it possible to programmatically change the color of the status bar without targeting API 21?

**Important: Downvotes are strongly discouraged in this thread. Sorting by new is strongly encouraged.**

Large code snippets don't read well on reddit and take up a lot of space, so please don't paste them in your comments. Consider linking [Gists](https://gist.github.com) instead.

Have a question about the subreddit or otherwise for /r/androiddev mods? [We welcome your mod mail!](http://www.reddit.com/message/compose?to=%2Fr%2Fandroiddev)

Also, please don't link to Play Store pages or ask for feedback on this thread. Save those for the App Feedback threads we host on Saturdays.

Looking for all the Questions threads? Want an easy way to locate this week's thread? Click [this link](https://www.reddit.com/r/androiddev/search?q=title%3A%22questions+thread%22+author%3A%22AutoModerator%22&amp;restrict_sr=on&amp;sort=new&amp;t=all)!
## [3][Android hobbyist dev and working as a dev - advice](https://www.reddit.com/r/androiddev/comments/ime1kh/android_hobbyist_dev_and_working_as_a_dev_advice/)
- url: https://www.reddit.com/r/androiddev/comments/ime1kh/android_hobbyist_dev_and_working_as_a_dev_advice/
---
Hello,  
Hope this post is not against the rules. Just seeking some advice on a situation I have.  
So in my last couple of years I have been doing some Android development as a hobby and have released several apps to Google Play Store. One of them has done quite well, and has been a source of small income since, and I am trying to actively work on it.  
However, now at my actual job, I am being put on an Android project, which means I'll (most likely) be using Google Play using my work account.

Having read many stories about hobbyist devs having their accounts suspended by Google, and then in turns any work they do by association can be suspended as well, I am faced with a decision.  
Should I remove all my apps that I have in Play Store, and just forget about my hobby of being a solo android dev in Play Store? Even if I remove them - does it give a decent safety that no apps I remove risk of getting any strikes in the future? Will I be OK to use my work account for Google Play console, despite having used it with my personal account?  


Thanks in advance for any advice given, really appreciate it.  
Andrew
## [4][Calculator library (pure Kotlin module)](https://www.reddit.com/r/androiddev/comments/imeek5/calculator_library_pure_kotlin_module/)
- url: https://www.reddit.com/r/androiddev/comments/imeek5/calculator_library_pure_kotlin_module/
---
Many apps have number inputs, and in many cases, having a calculator function to provide the number inputs is desirable. This is my first time to publish a library, and hope you will find this helpful. Any feedback is much appreciated. To import:

    repositories {
        maven { url 'https://jitpack.io' }
    }
    
    dependencies {
        implementation 'com.github.jairrab:KotlinCalculator:1.0.0'
    }

To get an instance of the `Calculator` library, call `Calculator.getInstance()` and pass a listener to receive calculator updates.

    class MainActivity : AppCompatActivity(), Calculator.Listener {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
    
            val calculator = Calculator.getInstance(this)
        }
    
        override fun onCalculatorUpdate(key: String?, entries: List&lt;String&gt;, result: Double) {
            Log.v("CALC", "key: $key | entries: $entries | result: $result")
        }
    }

`Calculator.getInstance()` returns a `Calculator` interface that provides the following functions:

    calculator.clear()
    calculator.pressOne()
    calculator.pressTwo()
    calculator.pressThree()
    calculator.pressFour()
    calculator.pressFive()
    calculator.pressSix()
    calculator.pressSeven()
    calculator.pressEight()
    calculator.pressNine()
    calculator.pressZero()
    calculator.pressDecimal()
    calculator.pressPlus()
    calculator.pressMinus()
    calculator.pressMultiply()
    calculator.pressDivide()
    calculator.backSpace()
    calculator.pressEquals()

The `Calculator.Listener` receives updates when a calculator function is called. Here's an example of a sequence of calculator operations and it's resulting logs:

    I/CalculatorLog: Key: initializing | Entries: [] | Result: 0.0
    I/CalculatorLog: Key: 1 | Entries: [1] | Result: 1.0
    I/CalculatorLog: Key: 2 | Entries: [12] | Result: 12.0
    I/CalculatorLog: Key: 3 | Entries: [123] | Result: 123.0
    I/CalculatorLog: Key: backspace | Entries: [12] | Result: 12.0
    I/CalculatorLog: Key: + | Entries: [12, +] | Result: 12.0
    I/CalculatorLog: Key: 1 | Entries: [12, +, 1] | Result: 13.0
    I/CalculatorLog: Key: 3 | Entries: [12, +, 13] | Result: 25.0
    I/CalculatorLog: Key: * | Entries: [12, +, 13, *] | Result: 25.0
    I/CalculatorLog: Key: 2 | Entries: [12, +, 13, *, 2] | Result: 38.0

The calculator is a pure Kotlin module library and does not have any Android dependency, thus you will need to provide a UI to call the and display the calculator functions.  This makes the library lightweight and very flexible to suit your needs. See the link below for a sample UI.

Note: The calculator library currently only supports MDAS operation. That is, it prioritizes multiplication and division over addition and subtraction as shown on the image above.

Link: [https://github.com/jairrab/KotlinCalculator/](https://github.com/jairrab/KotlinCalculator/)
## [5][The internals of Android APK build process - Article](https://www.reddit.com/r/androiddev/comments/ilu7ye/the_internals_of_android_apk_build_process_article/)
- url: https://www.reddit.com/r/androiddev/comments/ilu7ye/the_internals_of_android_apk_build_process_article/
---
## Table of Contents

* CPU Architecture and the need for Virtual Machine
* Understanding the Java Virtual Machine
* Compiling the Source Code
* Android Virtual Machine
* Compilation Process to .dex
* ART over Dalvik
* Understanding each part of the build process.
* Source Code
* Resource Files
* AIDL Files
* Library Modules
* AAR Libraries
* JAR Libraries
* Android Asset Packaging Tool
* resources.arsc
* D8 and R8
* Dex and Multidex
* Signing the APK
* References

**Understanding the flow of the Android APK build process, the execution environment, and code compilation** blog post aims to be the starting point for developers to get familiar with the build process of Android APK.

## CPU Architecture and the need for Virtual Machine

Unveiled in 2007, Android has undergone lots of changes related to its build process, the execution environment, and performance improvements.

There are many fascinating characteristics in Android and one of them is different CPU architectures like **ARM64 and x86**

It is not realistic to compile code that supports each and every architecture. **This is where Java Virtual Machine is used.**

https://preview.redd.it/91nrrk3twxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=a95b8cf916f000e94c6493a5780d9244e8d27517

## Understanding the Java Virtual Machine

**JVM is a virtual machine that enables a computer to run applications that are compiled to Java bytecode.** It basically helps us in converting the compiled java code to machine code.

By using the JVM, the issue of dealing with different types of CPU architecture is resolved.

**JVM provides portability** and it also allows Java code to be executed in a virtual environment rather than directly on the underlying hardware.

But JVM is designed for systems with huge storages and power, whereas **Android has comparatively low memory and battery capacity.**

**For this reason, Google has adopted an Android JVM called Dalvik.**

&amp;#x200B;

https://preview.redd.it/up2os7juwxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=2a290bdc9be86fb08d67228c730329130da3bc63

## Compiling the Source Code

Our Java source code for the Android app is **compiled into a .class file bytecode by the javac compiler and executed on the JVM.**

For Kotlin source code, when targeting JVM, Kotlin produces Java-compatible bytecode, thanks to **kotlinc compiler.**

To understand bytecode, it is a form of instruction set designed for efficient execution by a software interpreter.

Whereas Java bytecode is the instruction set of the Java virtual machine.

&amp;#x200B;

https://preview.redd.it/w2uzoicwwxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=b122e0781bf9e9ba236d34a87a636c9218f7ea35

## Android Virtual Machine

Each Android app runs on its own virtual machine. From version 1.0 to 4.4, it was 'Dalvik'. **In Android 4.4, along with Dalvik, Google experimentally introduced a new Android Runtime called 'ART'.**

Android users had the option to choose either Dalvik or ART runtime in Android 4.4.

**The .class files generated contains the JVM Java bytecodes.**

But Android has its own optimized bytecode format called Dalvik from version 1.0 to 4.4. **Dalvik bytecodes, like JVM bytecodes, are machine-code instructions for a processor.**

&amp;#x200B;

https://preview.redd.it/sqychk81xxk51.png?width=217&amp;format=png&amp;auto=webp&amp;s=49445fa42e4aa6f4008114a822f364580649fcdf

## Compilation Process to .dex

The compilation process **converts the .class files and .jar libraries into a single classes.dex file** containing Dalvik byte-codes. This is possible with the dx command.

**The dx command turns all of the .class and .jar files together into a single classes.dex file is written in Dalvik bytecode format.**

To note, dex means Dalvik Executable.

https://preview.redd.it/g4z3tb95xxk51.jpg?width=831&amp;format=pjpg&amp;auto=webp&amp;s=1cdbaacaf10cc529cccca2ba016583596781ee88

## ART over Dalvik

Since Android 4.4, Android migrated to ART, the Android runtime from Dalvik. **This execution environment executes .dex as well.**

The benefit of ART over Dalvik is that the app runs and launches faster on ART, this is because DEX bytecode has been translated into machine code during installation, **no extra time is needed to compile it during the runtime.**

ART and Dalvik are **compatible runtimes** running Dex bytecode, so apps developed for Dalvik should work when running with ART.

The JIT based compilation in the previously used Dalvik has disadvantages of poor battery life, application lag, and performance.

**This is the reason Google created Android Runtime(ART).**

ART is based on Ahead - Of - Time (AOT) based compilation process where **compilation happens before application starts.**

In ART, the compilation process happens during the app installation process itself. Even though this leads to higher app installation time, it reduces app lag, increases battery usage efficiency, etc.

**Even though dalvik was replaced as the default runtime, dalvik bytecode format is still in use (.dex)**

**In Android version 7.0, JIT came back.** The hybrid environment combining features from both a JIT compiler and ART was introduced.

*The bytecode execution environment of Android is important as it is involved in the application startup and installation process.*

https://preview.redd.it/qh9bxsplzxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=bc40ba6c69cec2110b7d695fe23df094bf5aea6c

# Understanding each part of the process.

&amp;#x200B;

https://preview.redd.it/obelgd7axxk51.png?width=950&amp;format=png&amp;auto=webp&amp;s=299abcf4798ad4d2de93f4eb18b9d9e70dd54297

## Source Code

**Source code** is the Java and Kotlin files in the src folder.

## Resource Files

**The resource files** are the ones in the res folder.

## AIDL Files

**Android Interface Definition Language (AIDL)** allows you to define the programming interface for client and service to communicate using IPC.

IPC is interprocess communication.

AIDL can be used between any process in Android.

## Library Modules

**Library module** contains Java or Kotlin classes, Android components, and resources though assets are not supported.

**The code and resources of the library project** are compiled and packaged together with the application.

*Therefore a library module can be considered to be a compile-time artifact.*

## AAR Libraries

**Android library compiles into an Android Archive (AAR) file that you can use as a dependency for an Android app module.**

**AAR files can contain Android resources and a manifest file**, which allows you to bundle in shared resources like layouts and drawables in addition to Java or Kotlin classes and methods.

## JAR Libraries

JAR is a Java library and unlike AAR it cannot contain Android resources and manifests.

## Android Asset Packaging Tool

**Android Asset Packaging Tool (aapt2) compiles the AndroidManifest and resource files into a single APK.**

At this point, it is divided into two steps, **compiling and linking.** It improves performance, since if only one file changes, you only need to recompile that one file and link all the intermediate files with the 'link' command.

**AAPT2 supports the compilation of all Android resource types, such as drawables and XML files.**

When you invoke AAPT2 for compilation, you should pass a single resource file as an input per invocation.

AAPT2 then parses the file and generates an intermediate binary file with a .flat extension.

**The link phase merges all the intermediate files generated in the compile phase and outputs one .apk file.** You can also generate R.java and proguard-rules at this time.

## resources.arsc

The output .apk file does not include the DEX file, so the DEX file is not included, and since it is not signed, it is an APK that cannot be executed.

**This APK contains the AndroidManifest, binary XML files, and resources.arsc.**

*This resource.arsc contains all meta-information about a resource, such as an index of all resources in the package.*

It is a binary file, and the APK that can be actually executed, and the APK that you often build and execute are uncompressed and can be used simply by expanding it in memory.

**The R.java that is output with the APK is assigned a unique ID, which allows the Java code to use the resource during compilation.**

arsc is the index of the resource used when executing the application.

&amp;#x200B;

https://preview.redd.it/hmmlfwhdxxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=b2fe2b6ad998594a5364bb6af6b5cbd880a2452c

## D8 and R8

**Starting from android studio 3.1 onwards, D8 was made the default compiler.**

D8 produces smaller dex files with better performance when compared with the old dx.

**R8 is used to compile the code. R8 is an optimized version of D8.**

*D8 plays the role of dexer that converts class files into DEX files and the role of desugar that converts Java 8 functions into bytecode that can be executed by Android.*

**R8 further optimizes the dex bytecode. R8 provides features like optimization, obfuscation, remove unused classes.**

**Obfuscation reduces the size of your app by shortening the names of classes, methods, and fields.**

Obfuscation has other benefits to prevent easy reverse engineering, but the goal is to reduce size.

**Optimization reduces the DEX file size by rewriting unnecessary parts and inlining.**

By doing Desugaring we can use the convenient language features of Java 8 in older devices.

https://preview.redd.it/so424bxwxxk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=0ad2df5bd194ec770d453f620aae9556e14ed017

## Dex and Multidex

**R8 outputs one DEX file called classes.dex.**

If you are using Multidex, that is not the case, but multiple DEX files will appear, but for the time being, classes.dex will be created.

**If the number of application methods exceeds 65,536 including the referenced library, a build error will occur.**

The method ID range is 0 to 0xFFFF.

*In other words, you can only refer to 65,536, or 0 to 65,535 in terms of serial numbers.*

This was the cause of the build error that occurred above 64K.

**In order to avoid this, it is useful to review the dependency of the application and use R8 to remove unused code or use Multidex.**

&amp;#x200B;

https://preview.redd.it/kjyychmzxxk51.png?width=1261&amp;format=png&amp;auto=webp&amp;s=18bea3bf9f7920a4701c2db9714dc53ae6cc5f82

## Signing the APK

**All APKs require a digital signature before they can be installed or updated on your device.**

**For Debug builds**, Android Studio automatically signs the app using the debug certificate generated by the Android SDK tools when we run.

A debug Keystore and a debug certificate is automatically created.

**For release builds**, you need a Keystore and upload the key to build a signed app. You can either make an APK file with apkbuilder and finally optimize with zipalign on cmd or have Android Studio handle it for you with the 'Generated Signed Apk option'.

&amp;#x200B;

https://preview.redd.it/10m8rjl0yxk51.png?width=1468&amp;format=png&amp;auto=webp&amp;s=078c4ab3f41c7d08e7c2280555ef2038cc04c5b0

## References

[https://developer.android.com/studio/build](https://developer.android.com/studio/build)

[https://github.com/dogriffiths/HeadFirstAndroid/wiki/How-Android-Apps-are-Built-and-Run](https://github.com/dogriffiths/HeadFirstAndroid/wiki/How-Android-Apps-are-Built-and-Run)

[https://logmi.jp/tech/articles/322851](https://logmi.jp/tech/articles/322851)

[https://android-developers.googleblog.com/2017/08/next-generation-dex-compiler-now-in.html](https://android-developers.googleblog.com/2017/08/next-generation-dex-compiler-now-in.html)

[https://speakerdeck.com/devpicon/uncovering-the-magic-behind-android-builds-droidfestival-2018](https://speakerdeck.com/devpicon/uncovering-the-magic-behind-android-builds-droidfestival-2018)

by [androiddevnotes on GitHub](https://github.com/androiddevnotes)

🐣
## [6][Users can't redeem promotional codes anymore](https://www.reddit.com/r/androiddev/comments/imctsz/users_cant_redeem_promotional_codes_anymore/)
- url: https://www.reddit.com/r/androiddev/comments/imctsz/users_cant_redeem_promotional_codes_anymore/
---
Since August 31, users are reporting that they cannot redeem our promotional codes for our in-app purchase anymore.

The error they see is '**Could not redeem code at this time. Please try again later**'. In our tests we see the error too.

The promo codes from that promotion worked fine before, but just in case we created a new promotion to see if the error was happening only on the old promotion. Unfortunately, the error still happens with the new promotion.

Is anyone else experiencing this problem?
## [7][what are some free face filter library for android?](https://www.reddit.com/r/androiddev/comments/imftqb/what_are_some_free_face_filter_library_for_android/)
- url: https://www.reddit.com/r/androiddev/comments/imftqb/what_are_some_free_face_filter_library_for_android/
---
I have a project to build something similar to TikTok, hence I am looking for some free opensource library, by which I will be able to enable face filtering features like TikTok?
## [8][Android Studio extremely slow with autocompletion, code analysis, etc.](https://www.reddit.com/r/androiddev/comments/im0d54/android_studio_extremely_slow_with_autocompletion/)
- url: https://www.reddit.com/r/androiddev/comments/im0d54/android_studio_extremely_slow_with_autocompletion/
---
I'm wondering if anyone else has observed something similar:

in the past month or two, I observed progressively more problems with AS speed in things like:

\- opening a new file, it takes like 10-20+ seconds for the IDE to "color up" the code (from the initial gray)

\- even when a file is opened (and "colored"), oftentimes I start typing and autocompletion just doesn't work, it takes another several dozen seconds for it to finally "sync up"

\- general random actions just "freeze" for a few seconds

&amp;#x200B;

**however, build times are as fast as always.**

\- 2018 MBP, 6 core i7, 32 GB RAM. I allocated 4GBs of RAM to AS (never been a problem before). I disabled all the external plugins.
## [9][RPG Database App](https://www.reddit.com/r/androiddev/comments/imdf33/rpg_database_app/)
- url: https://www.reddit.com/r/androiddev/comments/imdf33/rpg_database_app/
---
A friend of mine is launching his own RPG series and wants an app for the content database as well as local character sheet storage. I know I need to break this down further into making a searchable database and... I don't even know where to start for the character sheets. I've taken a few classes in programming for game development but could use help figuring out where to start with this.
## [10][Display on Top permission granted automatically on install from Play Store](https://www.reddit.com/r/androiddev/comments/imccdx/display_on_top_permission_granted_automatically/)
- url: https://www.reddit.com/r/androiddev/comments/imccdx/display_on_top_permission_granted_automatically/
---
Hi All.  


I need some help finding a form for Google Play Support and my Google Fu has failed me.  
The Form is to grant display over apps permission on app install from Google Play Store, very much how Truecaller is granted this permission. I'm aware that this form has existed atleast a couple months ago but I can't seem to find it anymore  


Any help will be appreciated. Thanks :)
## [11][I'm new to the whole android thing and I'm trying to develop a new skill in it.](https://www.reddit.com/r/androiddev/comments/imb4h6/im_new_to_the_whole_android_thing_and_im_trying/)
- url: https://www.reddit.com/r/androiddev/comments/imb4h6/im_new_to_the_whole_android_thing_and_im_trying/
---
I'm trying to connect to a specific mac address from Bluetooth socket   
after clicking a button and perform some task after establishing the connection 

`connectBtn.setOnClickListener(new View.OnClickListener() {`  
`@Override`  
 `public void onClick(View view) {`  
 `if (!bluetoothAdapter.isEnabled()) {`  
`showToast("turning on bluetooth");`  
 `Intent enableIntent = new Intent(BluetoothAdapter.ACTION_REQUEST_ENABLE);`  
 `startActivityForResult(enableIntent, REQUEST_ENABLE_BT);`  
 `}else {`  
`showToast("Bluetooth ON Retrying connection");`  
 `}`  
`establishConnection();`  
 `System.out.println(bluetoothSocket);`  
 `}`  
`});`  
`if(bluetoothSocket!=null&amp;&amp;bluetoothSocket.isConnected()) {`  
 `try {`  


`OutputStream outputStream = bluetoothSocket.getOutputStream();`  
 `outputStream.write(48);`  
 `System.out.println("out\n");`  
 `} catch (IOException e) {`  
`e.printStackTrace();`  
 `}`  
 `try {`  
`System.out.println("in\n");`  
 `InputStream inputStream = bluetoothSocket.getInputStream();`  
 `inputStream.skip(inputStream.available());`  
`byte b = (byte) inputStream.read();`  
`if (b == 65) {`  
`mediaPlayer.start();`  
 `}`  
`} catch (IOException e) {`  
`e.printStackTrace();`  
 `}`  
`}`  


`}`  
`private void establishConnection(){`  
 `try {`  
 `bluetoothSocket = hc05.createRfcommSocketToServiceRecord(mUUID);`  
 `System.out.println(bluetoothSocket);`  
 `bluetoothSocket.connect();`  
 `showToast("!!Cane is Connected!!");`  
 `System.out.println("!!yes connected!!"+bluetoothSocket.isConnected());`  
 `} catch (IOException e) {`  
`e.printStackTrace();`  
 `}`  


`}`  
 the rest of the code after OnClickListeners is not working  The 

    if(...){...}

block after OnClickListeners
## [12][Help me design a DB for my app](https://www.reddit.com/r/androiddev/comments/im9b3q/help_me_design_a_db_for_my_app/)
- url: https://www.reddit.com/r/androiddev/comments/im9b3q/help_me_design_a_db_for_my_app/
---
I am building my first android app. It's gonna be a workout log app, where the user can pick a date from the Calendar View to enter a workout, each workout will have a table that consists of all the exercises, number of sets for each exercise, and average\_reps, average\_weight, and average\_rpe which should be calculated from the Sets table. Each Set of the exercise, should have its own row that stores reps, weight, and rpe.

Please advise me if this is a good practice of building such database, and any advice on how to do this in Android &amp; Kotlin would be appreciated. 

https://preview.redd.it/9ir7tphk42l51.png?width=1834&amp;format=png&amp;auto=webp&amp;s=1440acee1bc89a7c8e77ee9522a36c249d2049c8
