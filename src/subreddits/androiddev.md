# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ijxa5y/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ijxa5y/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
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
## [3][I created this simple app that will notify you of any new releases to the current Android Jetpack libraries. You can also browse and bookmark specific libraries you're interested in for easier tracking. It's on GitHub and Play Store. [link in comments]](https://www.reddit.com/r/androiddev/comments/ikrbbs/i_created_this_simple_app_that_will_notify_you_of/)
- url: https://i.redd.it/a1r2hachalk51.png
---

## [4][Getting feedback on code](https://www.reddit.com/r/androiddev/comments/il59vb/getting_feedback_on_code/)
- url: https://www.reddit.com/r/androiddev/comments/il59vb/getting_feedback_on_code/
---
Hey, I was wondering if there is a platform where you can provide your code and somebody (probably an expert) reviews it and tells you how to improve and what you made wrong. Of course I would pay for such an service.

I am asking because throughout the last year I was reworking a lot of code again and again because I was unsatisfied with my solutions, and I was getting crazy because I wanted to know how an expert would have solved the problem.
## [5][The internals of Android Stack Architecture - Article](https://www.reddit.com/r/androiddev/comments/ikl56a/the_internals_of_android_stack_architecture/)
- url: https://www.reddit.com/r/androiddev/comments/ikl56a/the_internals_of_android_stack_architecture/
---
## Table of Contents

* Early Days of Android
* Android Stack
* Linux Kernel
* Secure Element
* Hardware abstraction layer
* Native Libraries
* Runtime
* Framework
* Apps
* References

**Understanding the internals of Android Stack Architecture and how it relates to Linux** blog post aims to be the starting point for developers to get familiar and have an overview of core components in the Android Stack Architecture.

## Early Days of Android

Founded in 2003, Android Inc, in the early days began as an operating system for *Digital Cameras.* Due to the low market for Digital Cameras, the Android Inc teams' intentions slowly diverted to Mobile Devices. **Becoming the rivals of then Symbian and Windows Mobile.**

Acquired by Google in 2005, the development team worked on an Operating System based on Linux Kernel in shadows until the unveil of **Open Handset Alliance in 2007.**

https://preview.redd.it/wjoa0sysojk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=fd78d0546420b115f5f32acad6c093bfc4688f81

## Android Stack

By Google, Android is described officially as an open-source,**Linux-based software stack** created for a wide array of devices and form factors.

The graphical user interface environment, middlewares, libraries, APIs... sitting on top of Linux Kernel and shell binaries are Software Stack Layers that make the bulk of Android and which makes it **much more than a variation of the Linux system.**

https://preview.redd.it/jcmf41nuojk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=3cf4d97d8a32277f6ec8faf1634a3ec6e9e87b98

## Linux Kernel

Android's Kernel is based on the Long Term Support (LTS) branches of the Linux Kernel.

The kernel provides drivers for filesystem access, process management, hardware, networking.

The Android Kernel differs from vanilla Linux Kernel due to the differences called as **Androidisms**

**Some of the notable Androidisms added to Kernel are** IPC Binder, Wavelocks, Low-Memory Killer, Dalvik, and Android Runtime, Anonymous Shared Memory (ashmem), Alarm, paranoid network, RAM console, Physical memory (pmem), Sync driver, Timed Output, and GPIO, memory and logging enhancements.

**Android utilizes many unused/less-popular features in desktop distributions of Linux** such as control groups, Low Memory Killer Daemon, Security-Enhanced Linux (SELinux), and open source projects like a racoon for VPN, mdns for network service discovery, and many more.

https://preview.redd.it/j2cq9uhxojk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=ec4bc88f683efe190610532db27e9b3e83f781e3

## Secure Element

To provide better security, some devices have an embedded Secure Element (SE), which is dedicated, separate tamper-resistant hardware to store cryptographic data.

## Hardware abstraction layer

Android runs on TV, Mobile, Refrigerator and almost everywhere that the underlying hardware may greatly differ in its' scope and support.

**To solve this the Android stack typically relies on shared libraries provided by manufacturers to interact with hardware.**

Android relies on what can be considered a Hardware Abstraction Layer (HAL), although the interface, behavior, and function of abstracted hardware components differ greatly from type to type.

The idea of GPS, sensor, TV, camera, audio, input media components and other components behavior is defined by HAL and how it should behave in Android.

**The vendors are still not allowed to make unnecessary modifications so as to not fail the Compatibility Test Suite, Vendor Test Suite.**

https://preview.redd.it/3xx6e0z0pjk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=a190299b4aba07053e67fd2df69e66222582b4ea

## Native Libraries

**The native libraries layer is responsible for providing support for the core features.**

The WebKit Web rendering engine, Audio Manager, LIBC, Secure Sockets Layer (SSL), FreeType for rendering fonts, Media, OpenGL ES graphics API, SQLite database, Surface Manager.

https://preview.redd.it/hmiokov4pjk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=e40d6a952df4d26ed3e9e2951dd3ed98d4c9c652

## Runtime

Before Android 4.4, each Android app would run on its own virtual machine, which is called "Dalvik" which got superseded by the Android RunTime (ART).**In Android 4.4, along with Dalvik, Google experimentally introduced a new Android Runtime called "ART" which still today remains the standard.**

ART introduces ahead-of-time (AOT) compilation, which can improve app performance.

**App runs and launches faster on ART than Dalvik** because DEX bytecode gets translated into machine code during installation which means no compilation during runtime and thus seemingly faster!

Because Dalvik requires extra memory for Just-in-time code cache, an app occupies a smaller memory footprint when it runs on ART."

https://preview.redd.it/x8b29li6pjk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=028cf01c7fa5b66d7b926a1434a697530941457b

## Framework

**The Android application creation process is provided by Application frameworks/libraries** which allows developers to use the higher-level Kotlin or Java language, rather than low-level C/C++.

**The framework includes the basic blocks for building Android applications** such as Content Provider, Activity Manager, Location Manager, View System, Package Manager, Notification Manager, Resource Manager, Telephony Manager, Window Manager.

**Android frameworks are divided into separate namespaces using the Java package naming and according to their functionality.**

Packages in the android.\* namespace is available for use by developers.

Packages in com.android.\* are internal.

Android also supports most of the standard Java runtime packages in the java.\* namespace.

https://preview.redd.it/9vhtlhu8pjk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=feab2c94bfda06af4c0ae0ac051cd1d1772feccc

## Apps

The topmost layer in the Android Stack is the Applications layer which can be categorized into System apps and user-installed apps.

System apps cannot be uninstalled or changed by users and are read-only in production devices. **System apps are included in the OS image, mounted as /system**

User-installed apps can be uninstalled at will. Each application lives in a dedicated security sandbox and cannot affect other applications or access their data.

**User-installed apps are installed on a dedicated read-write partition, mounted as /data that host user data.**

https://preview.redd.it/09tg4sv9pjk51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=8a75bf2267c51eba9d18628a7dc7a4b1727f7db3

## References

[https://developer.android.com/guide/platform](https://developer.android.com/guide/platform)

[https://source.android.com](https://source.android.com)

[https://en.wikipedia.org/wiki/Android\_(operating\_system)](https://en.wikipedia.org/wiki/Android_(operating_system))

[https://software.intel.com/content/www/us/en/develop/blogs/art-vs-dalvik-introducing-the-new-android-x86-runtime.html](https://software.intel.com/content/www/us/en/develop/blogs/art-vs-dalvik-introducing-the-new-android-x86-runtime.html)

Android Internals::A Confectioner's Cookbook

Android Security Internals: An In-Depth Guide to Android's Security Architecture

Embedded Android: Porting, Extending, and Customizing

by [androiddevnotes on GitHub](https://github.com/androiddevnotes)

🐣
## [6][Top 20 Android App Development Tools](https://www.reddit.com/r/androiddev/comments/il6bew/top_20_android_app_development_tools/)
- url: https://blog.svapinfotech.com/android-app-development-tools/
---

## [7][ADB over WiFi Batch Script (development on Windows)](https://www.reddit.com/r/androiddev/comments/il5jrv/adb_over_wifi_batch_script_development_on_windows/)
- url: https://www.reddit.com/r/androiddev/comments/il5jrv/adb_over_wifi_batch_script_development_on_windows/
---
I've created a simple batch script to manage some of the workflow needed to debug your Android app over WiFi.

My knowledge on batch scripting is very limited; I'm open for feedback.

I've created it, because my device's USB connection is faulty. Maybe it can help you as well.

[File and description (GitHub Gist)](https://gist.github.com/jpmcosta/aced6a1ac03e61a7b0fa29082d0b51ce)
## [8][How to make a backup low level only within shell /data/data/com.app](https://www.reddit.com/r/androiddev/comments/il3a0m/how_to_make_a_backup_low_level_only_within_shell/)
- url: https://www.reddit.com/r/androiddev/comments/il3a0m/how_to_make_a_backup_low_level_only_within_shell/
---
Is it possible to make a backup of an app inside /data/data/com.app ? Need to know of low-level method if possible, I know of higher level methods.

&amp;#x200B;

Was testing this today and wasn't successfull as root:

1. \# cp -r /data/data/com.app /sdcard/
2. \# pm uninstall com.app
3. \# cp -r /sdcard/com.app /data/data/
4. \# pm install /sdcard/com.app

&amp;#x200B;

I ran app but got a bunch of following errors:

`E SharedPreferencesImpl: Couldn't rename file /data/user/0/com.app/shared_prefs/com.google.android.gms.measurement.prefs.xml to backup file /data/user/0/com.app/shared_prefs/com.google.android.gms.measurement.prefs.xml.bak`

&amp;#x200B;

So my question is. Is there a low level approach where I can manually backup apps preferences, logins etc?

&amp;#x200B;

**EDIT: Found solution for myself**: [https://android.stackexchange.com/questions/136895/backup-an-apps-complete-data-when-phone-is-stuck-in-bootloop/137043#137043](https://android.stackexchange.com/questions/136895/backup-an-apps-complete-data-when-phone-is-stuck-in-bootloop/137043#137043)

Backup:

`busybox tar -C / -cvzhf /sdcard/PACKAGE_NAME.tgz data/app/PACKAGE_NAME* data/data/PACKAGE_NAME # replace PACKAGE_NAME with package name of the app`

Restore: (I restored before first running an app)

`busybox tar -C / -xvzhf BACKUP_PATH # replace BACKUP_PATH with the file path of your backup, such as /sdcard/PACKAGE_NAME.tgz`
## [9][DocumentFileX - java.io.File interoperable SAF implementation](https://www.reddit.com/r/androiddev/comments/ikykfr/documentfilex_javaiofile_interoperable_saf/)
- url: https://www.reddit.com/r/androiddev/comments/ikykfr/documentfilex_javaiofile_interoperable_saf/
---
SAF is a framework that brings everyone who wants to access filesystem pain and agony. Slow performance and a messy official library make using SAF literally a pain in the butt.

I can't do much about performance, but I can write a better library so that's what I did.

keep in mind that this library is in its early alpha stage, things might break down or change drastically in the future.

The feature of the library is as follows:

* Interoperable with java.io.File with some exceptions
* File.getName() without overhead
* Faster Directory IO
* Get child FileX without overhead
* *Slightly* improved listFiles() //√evil
* Automatic URI type detection
* File metadata caching
* Useful URI extensions
* (WIP) kotlin-stdlib File extension compatible extension methods

and here's some sample code

## File I/O

    val file = FileX(context, uri)
    
    file.writeText("Hello, SAF! You're dumb.")
    val text = file.readText() // "Hello, SAF! You're dumb."
    
    val data = listOf(0x82, 0x72, 0x82, 0x60, 0x82, 0x65, 0x82, 0xA4, 0x82, 0xF1, 0x82, 0xBF).map { 
        it.toByte()
    }.toByteArray()
    
    file.writeBytes(data)
    val text = file.readText(data, Charset.forName(&lt;small quiz for you&gt;))

## Directory I/O

Directory I/O is only supported by `tree://com.android.externalstorage.documents/...` URI

    val folder = FileX(context, uri, "akita") // No overhead
    val child = folder.getChild("daisen") // No overhead
    val neighbor = FileX(context, folder.parent, "yamagata") // No overhead
    val neice = FileX(context, folder.parent, "iwate/morioka/nakano.txt") // No overhead
    
    if (neice.parent.mkdirs()) {
        neice.createNewFile()
        neice.renameTo(FileX(context, neice.parent, "kurokawa.json"))
    }
    
    folder.listFiles().forEach { sichouson -&gt; // Returns FileX
        sichouson.list().forEach { // Returns Uri string
            ....
        }
    }

head over to [https://github.com/tom5079/DocumentFileX](https://github.com/tom5079/DocumentFileX) to learn more!
## [10][Understanding the Android APK build process, the execution environment, and code compilation.](https://www.reddit.com/r/androiddev/comments/ikr28m/understanding_the_android_apk_build_process_the/)
- url: https://www.youtube.com/watch?v=Qp-5stxpTz4
---

## [11][How to generate PR directly from Android Studio?](https://www.reddit.com/r/androiddev/comments/il2slq/how_to_generate_pr_directly_from_android_studio/)
- url: https://www.reddit.com/r/androiddev/comments/il2slq/how_to_generate_pr_directly_from_android_studio/
---
Every time I complete my work, my boss asks me to raise the **pull request** for the work done. To raise the **PR** I navigate to **Bitbucket** dashboard and then generate the **PR** against the desired branch manually. Although it's a minimal task, I would like to accomplish it without leaving the Android Studio.  


I am aware that the VCS option in Android Studio provides many options for commit, pull and push but I'm still missing  **Create Pull Reques**t option.  


https://preview.redd.it/x5nxjfz30pk51.png?width=1444&amp;format=png&amp;auto=webp&amp;s=bf51bbb325b50d87305747839f423725ade9b83a

Are there any plugins or options through which I can raise the PR  directly from the Android Studio itself?
## [12][Smaller APKs with resource optimization - Jake Wharton](https://www.reddit.com/r/androiddev/comments/ikkaec/smaller_apks_with_resource_optimization_jake/)
- url: https://jakewharton.com/smaller-apks-with-resource-optimization/
---

