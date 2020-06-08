# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/gyyn9x/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/gyyn9x/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - June 08, 2020](https://www.reddit.com/r/androiddev/comments/gyx6z3/weekly_questions_thread_june_08_2020/)
- url: https://www.reddit.com/r/androiddev/comments/gyx6z3/weekly_questions_thread_june_08_2020/
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
## [3][TIL: Be really careful with proguard wildcards](https://www.reddit.com/r/androiddev/comments/gyu1ar/til_be_really_careful_with_proguard_wildcards/)
- url: https://medium.com/@woitaschek/how-to-break-your-android-app-with-proguard-r8-6566bc387b63
---

## [4][I created this simple app that will notify you of any new releases to the current Android Jetpack libraries. You can also browse and bookmark specific libraries you're interested in for easier tracking. It's on GitHub and Play Store. Not sure if I'll add more features, but I'm open to suggestions.](https://www.reddit.com/r/androiddev/comments/gyyhrm/i_created_this_simple_app_that_will_notify_you_of/)
- url: https://i.redd.it/jzc3t3ofjo351.png
---

## [5][Mobile app CI with AppCenter — good, but could be better](https://www.reddit.com/r/androiddev/comments/gyxtol/mobile_app_ci_with_appcenter_good_but_could_be/)
- url: https://medium.com/variant-as/mobile-app-ci-with-appcenter-good-but-could-be-better-7c7eb515dcea
---

## [6][Finding Android dev a bit tougher as a Java developer as time goes on.](https://www.reddit.com/r/androiddev/comments/gyeoko/finding_android_dev_a_bit_tougher_as_a_java/)
- url: https://www.reddit.com/r/androiddev/comments/gyeoko/finding_android_dev_a_bit_tougher_as_a_java/
---
I've been working on my pet project for Android for a while now on and off, but lately I've been finding it increasingly difficult to find java code references and examples that are up to date. Everything new is almost always in Kotlin. Now I have nothing against Kotlin, I want to focus on increasing my proficiency with Java and direct energy into actually developing a full fledged app. 

I'd be interested to hear about if any other developers have experienced this and how they approached it. I am fine to learn Kotlin in the future, and hopefully when it's better than Java in all respects (hear a lot about performance issues in compile and ide).

Edit: Wow this kind of blew up. Thank you everyone for the feedback and advice. Means a lot. :)
## [7][What's a good 2020 source to learn app marketing, pricing, IAP strategies, ad placement etc.?](https://www.reddit.com/r/androiddev/comments/gyx5ms/whats_a_good_2020_source_to_learn_app_marketing/)
- url: https://www.reddit.com/r/androiddev/comments/gyx5ms/whats_a_good_2020_source_to_learn_app_marketing/
---
Preferable simple and compact written articles, do's and don'ts, starter strategies, geo-specific pricing etc.
## [8][What are the best places to stay on top of Android app development news/tools/ideas/knowledge?](https://www.reddit.com/r/androiddev/comments/gyx2n9/what_are_the_best_places_to_stay_on_top_of/)
- url: https://www.reddit.com/r/androiddev/comments/gyx2n9/what_are_the_best_places_to_stay_on_top_of/
---
Hey, I just want to get new sources into my Android info sources. What do you like recently?

Thanks for any tips in advance!
## [9][What are good examples for shared element transition implementation?](https://www.reddit.com/r/androiddev/comments/gyz54v/what_are_good_examples_for_shared_element/)
- url: https://www.reddit.com/r/androiddev/comments/gyz54v/what_are_good_examples_for_shared_element/
---
I'm trying to find apps that implement shared element transition in a meaningful way. I came across [this](https://medium.com/google-design/implementing-motion-9f2839002016) blog post which includes some nice animated gifs but those are only designs and not the real thing.
## [10][App update rejected (Default phone handler)](https://www.reddit.com/r/androiddev/comments/gyyyp7/app_update_rejected_default_phone_handler/)
- url: https://www.reddit.com/r/androiddev/comments/gyyyp7/app_update_rejected_default_phone_handler/
---
 We have an android app that we developed as a default dialer that can identify incoming calls if they are registered to a company.

At first Play store told us to implement the asking for default before asking for any permissions. We fixed that and our app was accepted. After that all updates have failed due to this error:

&gt;You declared Default Phone handler (and any other core functionality usage while default handler), Caller ID, spam detection, and /or spam blocking, Write and Show Call History in Dialer as the core functionality of your app. However, after review, we found that your app does not match the declared use case(s). Learn more about permitted uses and exceptions.

Our manifest:

    &lt;manifest
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:versionCode="5"
    android:versionName="1.0.4"
    android:compileSdkVersion="28"
    android:compileSdkVersionCodename="9"
    package="com.storybookdialerv4"
    platformBuildVersionCode="5"
    platformBuildVersionName="1.0.4"&gt;
    &lt;application
        android:theme="(reference) @0x7f0f0006"
        android:label="(reference) @0x7f0e002e"
        android:icon="(reference) @0x7f0d0005"
        android:name="com.storybookdialerv4.App"
        android:allowBackup="true"
        android:largeHeap="true"
        android:supportsRtl="true"
        android:roundIcon="(reference) @0x7f0d0006"
        android:appComponentFactory="android.support.v4.app.CoreComponentFactory"&gt;
        &lt;activity
            android:theme="(reference) @0x7f0f00c6"
            android:name="com.nabinbhandari.android.permissions.PermissionsActivity"/&gt;
        &lt;service
            android:name="com.storybookdialerv4.NotificationService"/&gt;
        &lt;service
            android:name="com.storybookdialerv4.CallService"
            android:permission="android.permission.BIND_INCALL_SERVICE"&gt;
            &lt;intent-filter&gt;
                &lt;action
                    android:name="android.telecom.InCallService"/&gt;
            &lt;/intent-filter&gt;
            &lt;meta-data
                android:name="android.telecom.IN_CALL_SERVICE_UI"
                android:value="true"/&gt;
        &lt;/service&gt;
        &lt;activity
            android:name="com.storybookdialerv4.CallActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="2"/&gt;
        &lt;activity
            android:name="com.storybookdialerv4.DialerActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="2"&gt;
            &lt;intent-filter&gt;
                &lt;category
                    android:name="android.intent.category.DEFAULT"/&gt;
                &lt;action
                    android:name="android.intent.action.DIAL"/&gt;
            &lt;/intent-filter&gt;
            &lt;intent-filter&gt;
                &lt;data
                    android:scheme="tel"/&gt;
                &lt;category
                    android:name="android.intent.category.BROWSABLE"/&gt;
                &lt;category
                    android:name="android.intent.category.DEFAULT"/&gt;
                &lt;action
                    android:name="android.intent.action.DIAL"/&gt;
                &lt;action
                    android:name="android.intent.action.VIEW"/&gt;
            &lt;/intent-filter&gt;
            &lt;intent-filter&gt;
                &lt;category
                    android:name="android.intent.category.LAUNCHER"/&gt;
                &lt;action
                    android:name="android.intent.action.MAIN"/&gt;
            &lt;/intent-filter&gt;
        &lt;/activity&gt;
        &lt;activity
            android:name="com.storybookdialerv4.ContactsActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="20"/&gt;
        &lt;activity
            android:name="com.storybookdialerv4.AddContactActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="20"/&gt;
        &lt;activity
            android:name="com.storybookdialerv4.RecentsActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="2"/&gt;
        &lt;activity
            android:name="com.storybookdialerv4.CallKeypadActivity"
            android:screenOrientation="1"
            android:configChanges="80"
            android:windowSoftInputMode="2"/&gt;
    &lt;/application&gt;
    &lt;uses-permission
        android:name="android.permission.WAKE_LOCK"/&gt;
    &lt;uses-permission
        android:name="android.permission.WRITE_CALL_LOG"/&gt;
    &lt;uses-permission
        android:name="android.permission.READ_CALL_LOG"/&gt;
    &lt;uses-permission
        android:name="android.permission.WRITE_CONTACTS"/&gt;
    &lt;uses-permission
        android:name="android.permission.READ_CONTACTS"/&gt;
    &lt;uses-permission
        android:name="android.permission.MODIFY_AUDIO_SETTINGS"/&gt;
    &lt;uses-permission
        android:name="android.permission.RECORD_AUDIO"/&gt;
    &lt;uses-permission
        android:name="android.permission.CALL_PHONE"/&gt;
    &lt;uses-permission
        android:name="android.permission.PROCESS_OUTGOING_CALLS"/&gt;
    &lt;uses-permission
        android:name="android.permission.READ_PHONE_STATE"/&gt;
    &lt;uses-permission
        android:name="android.permission.ACCESS_NETWORK_STATE"/&gt;
    &lt;uses-permission
        android:name="android.permission.INTERNET"/&gt;
    &lt;uses-sdk
        android:minSdkVersion="23"
        android:targetSdkVersion="28"/&gt;

 We've choosen three core functionalities in our permissions declaration: 

\* Default Phone handler   
\* Caller ID, spam detection, and blocking   
\* Write and Show Call History in Dialer

We have tried contacting play store developer support but Google seems busy and we're waiting for a response for about a week already. The chat option says they can not do anything about permissions errors and that we have to wait for an email from Google. I'm wondering if anyone can see anything fishy about our manifest or permissions declaration.

The app is not listed in play store because it's in open beta and is limited to Estonia at our testing phase.

I'm kind of lost how to proceed from here. Google wont answer (or offer normal paid support like almost every other software/service company). Really hoping some redditors have ideas.
## [11][Android example project of LiveData, Retrofit and infinity scroll RecyclerView](https://www.reddit.com/r/androiddev/comments/gyveub/android_example_project_of_livedata_retrofit_and/)
- url: https://www.reddit.com/r/androiddev/comments/gyveub/android_example_project_of_livedata_retrofit_and/
---
Hi,

I've created an example project for anyone who don't know how to implement MVVM with LiveData and RecyclerView.

Link:

 [https://github.com/tRamboli/LiveData/](https://github.com/tRamboli/LiveData/) 

Feel free to fork, share, edit and improve the code.

Thanks!
## [12]["Introducing AndroidDistribution․io, a new home for the Android distribution chart"](https://www.reddit.com/r/androiddev/comments/gykslm/introducing_androiddistributionio_a_new_home_for/)
- url: https://www.reddit.com/r/androiddev/comments/gykslm/introducing_androiddistributionio_a_new_home_for/
---
In case you wanted to view it without Android Studio, someone made it into a website:

https://9to5google.com/2020/06/05/introducing-androiddistribution-io/

https://androiddistribution.io/

If you want, I've made today an app that shows it in textual format, open sourced too, based on the JSON we can have from Android Studio code (it also has an APK) :

https://github.com/AndroidDeveloperLB/AnddroidVersionsStats
