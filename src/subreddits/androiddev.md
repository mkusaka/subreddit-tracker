# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/huklpf/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/huklpf/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - July 20, 2020](https://www.reddit.com/r/androiddev/comments/huj9uy/weekly_questions_thread_july_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/huj9uy/weekly_questions_thread_july_20_2020/
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
## [3][JetBrains is porting Compose to desktop and other platforms](https://www.reddit.com/r/androiddev/comments/hufnjq/jetbrains_is_porting_compose_to_desktop_and_other/)
- url: https://www.reddit.com/r/androiddev/comments/hufnjq/jetbrains_is_porting_compose_to_desktop_and_other/
---
It's no news that JetBrains employees are working on Compose. But this job offer: https://www.jetbrains.com/careers/jobs/ui-framework-developer-401/ makes it clear what the goal is - "Our team is working to port it to desktop and other platforms.".
## [4][I have made this 🗽 NY Times App with JSoup to Crawl data from NYtimes with MVVM Architecture 💡](https://www.reddit.com/r/androiddev/comments/huea2b/i_have_made_this_ny_times_app_with_jsoup_to_crawl/)
- url: https://github.com/TheCodeMonks/NYTimes-App
---

## [5][Beagle debug menu library v2 + showcase app](https://www.reddit.com/r/androiddev/comments/huh589/beagle_debug_menu_library_v2_showcase_app/)
- url: https://www.reddit.com/r/androiddev/comments/huh589/beagle_debug_menu_library_v2_showcase_app/
---
This debug menu library has been rewritten from scratch to support more ways to display its UI (side drawer, bottom sheet, dialog, etc), custom modules, easier integration and a simpler API. I'm still actively testing it (that's the reason for the beta tag) but I'm not aware of any blockers at the moment so I'm very much looking forward to see what you think about it!

I've also created a [showcase app](https://play.google.com/store/apps/details?id=com.pandulapeter.beagle) to convince you that you need it in your project, and it also acts as a tutorial for implementing the different features. The source code for the app, as well as the library itself can be found in this [GitHub repo](https://github.com/pandulapeter/beagle).

Without going into details about what can be done with the various modules: one of my main goals was to make this as simple to use as possible, and based on the feedback I got from colleagues as well as QA people working on projects where Beagle is used, it's well worth the 5-10 minutes it takes to add it to the apps. Let me know if you have any questions / issues / ideas to make it better :)
## [6][Hilt in Android 11 (Android Dependency Injection)](https://www.reddit.com/r/androiddev/comments/hul3on/hilt_in_android_11_android_dependency_injection/)
- url: https://www.youtube.com/watch?v=B56oV3IHMxg
---

## [7][Android Studio + headless Emulator knocking down Macbook](https://www.reddit.com/r/androiddev/comments/hul1fp/android_studio_headless_emulator_knocking_down/)
- url: https://www.reddit.com/r/androiddev/comments/hul1fp/android_studio_headless_emulator_knocking_down/
---
Hi,  


I'm running 4.2 (canary 4) to use the embedded emulator inside Studio.  
Unfortunately it's actually not usable as the emulator is taking so much CPU of my new Macbook pro (2019, 2,6 Ghz i7, 32g ram).  
Is anyone else having this issue?
## [8][Exoplayer not work in huawei devices containing Huawei Mobile Services](https://www.reddit.com/r/androiddev/comments/huk17t/exoplayer_not_work_in_huawei_devices_containing/)
- url: https://www.reddit.com/r/androiddev/comments/huk17t/exoplayer_not_work_in_huawei_devices_containing/
---
In my huawei test device which contains huawei mobile services, exoplayer does not work both audio and video files. Do you know why? Does exoplayer not support devices including hms?
## [9][What are the ways to read SMS and sub-second latencies?](https://www.reddit.com/r/androiddev/comments/hujtcv/what_are_the_ways_to_read_sms_and_subsecond/)
- url: https://www.reddit.com/r/androiddev/comments/hujtcv/what_are_the_ways_to_read_sms_and_subsecond/
---
I am making an app that reads users SMS and sends them to a secure server and then notifies about some actions on a web-portal.

This complete flow is striving to do this in 4-5 seconds since the SMS is received. I want to discuss what are the pain points I can anticipate and if possible, how I can mitigate them.

There are a couple on the top of my head:

* There are some OS and mobile models that put the app in the sleep mode and then we can miss getting to know about the new SMS.
   * This might be handled by making the app exempt from the \`battery saver\` mode. 
* Playstore might not allow an app to read the users' SMS (for good reasons, but the use case, in particular, can be solved by reading SMS).
   * I can have this app available at other app stores like f-droid.  


Can there be more pain points that I am not considering at this moment? Any tips on how to take care of those.
## [10][Distributing android app outside of play store possible with app signing enabled for Google Play?](https://www.reddit.com/r/androiddev/comments/huaoe7/distributing_android_app_outside_of_play_store/)
- url: https://www.reddit.com/r/androiddev/comments/huaoe7/distributing_android_app_outside_of_play_store/
---
 0

I have an android app that is currently being distributed through google play store. Google upcoming deadline requires that [apps must target API level 29](https://developer.android.com/distribute/play-policies#apps-must-target-api-level-29) or I won't be able to update app after 2nd November 2020.

I can not change targetSdkVersion  
 29 or above because of the app core functionality relies on an Android API that has been removed in API level 29 and above.

My plan was to acquire new installs via Google Play Store and then update the app later from my own server or somewhere else. But I have opted for [App signing for Google Play](https://support.google.com/googleplay/android-developer/answer/7384423?hl=en) which resigns the APKs before redistributing them to users.

My Second Plan was to upload the APK to play store and not publish there and download the signed APK and distribute that from outside. Assuming google play will let me upload the APK after 2 November 2020. To test it out I changed targetSdkVersion  
 from 28 to 27. Current google play requirement is app must target API level 28. So google play rejected my APK.

I am wondering if there is a way to distribute app updates to the existing user outside of Google Play store.
## [11][Meet Google Play Billing Library Version 3](https://www.reddit.com/r/androiddev/comments/hugs84/meet_google_play_billing_library_version_3/)
- url: https://android-developers.googleblog.com/2020/06/meet-google-play-billing-library.html
---

## [12][Begginers guide to build modules on Android](https://www.reddit.com/r/androiddev/comments/hubk7z/begginers_guide_to_build_modules_on_android/)
- url: https://medium.com/@shalommathews05/modularization-in-android-make-your-code-reusable-and-maintainable-beginner-db5d6185a247
---

