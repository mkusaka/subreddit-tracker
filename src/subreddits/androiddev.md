# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/jidusx/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jidusx/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 26, 2020](https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/
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
## [3][Released an android sdk to reduce video streaming costs by up-to 90%](https://www.reddit.com/r/androiddev/comments/jixujb/released_an_android_sdk_to_reduce_video_streaming/)
- url: https://www.reddit.com/r/androiddev/comments/jixujb/released_an_android_sdk_to_reduce_video_streaming/
---
Hi everyone

I am working on an android sdk which can reduce video streaming costs of CDN by up-to 90% using a hybrid decentralized technology. I have opened it up for beta-access for developers to try it out. Looking for feedback w.r.t the technology and any features you would want to have.

A web demo is available here [https://api.peervadoo.com/test](https://api.peervadoo.com/test) . Click on **Add new peer** to see the tech in action,

Android sdk link :- [https://github.com/vadootvpeer/sdk-android](https://github.com/vadootvpeer/sdk-android)

We are giving free video streaming cdn credits for those interested from here [https://api.peervadoo.com/vadoo\_cdn](https://api.peervadoo.com/vadoo_cdn)
## [4][Released kotlinx.coroutines 1.4.0](https://www.reddit.com/r/androiddev/comments/jikkvp/released_kotlinxcoroutines_140/)
- url: https://github.com/Kotlin/kotlinx.coroutines/releases/tag/1.4.0
---

## [5][Navigating with SafeArgs](https://www.reddit.com/r/androiddev/comments/jixmlu/navigating_with_safeargs/)
- url: https://medium.com/androiddevelopers/navigating-with-safeargs-bf26c17b1269
---

## [6][Android development material](https://www.reddit.com/r/androiddev/comments/jj0yd2/android_development_material/)
- url: https://www.reddit.com/r/androiddev/comments/jj0yd2/android_development_material/
---
Hello I was an Android developer and stopped developing about 2 years ago I'm still a novice and was when I quit. If I want to get back on track what's a course or material I should study to up my skills and get a better understanding? Thanks!
## [7][Have you ever heard of undocumented QUICKBOOT_POWERON for boot receiver?](https://www.reddit.com/r/androiddev/comments/jj0cxj/have_you_ever_heard_of_undocumented_quickboot/)
- url: https://dev.to/localazy/7-floating-windows-on-android-boot-receiver-4flc
---

## [8][Saving Stack(java.util.Stack) Data on Application](https://www.reddit.com/r/androiddev/comments/jiz4ld/saving_stackjavautilstack_data_on_application/)
- url: https://www.reddit.com/r/androiddev/comments/jiz4ld/saving_stackjavautilstack_data_on_application/
---
Hi! How do you save stack data when you use java.util.Stack? Like say save the stack contents when your app closes and loads it when app reopens?  I'd like to know what concepts should I be looking into. Thanks!
## [9][Android 11 scoped storage - MediaStore can create sub-directories, delete - but are rename-file/move-file instantaneous as before ?](https://www.reddit.com/r/androiddev/comments/jib3l7/android_11_scoped_storage_mediastore_can_create/)
- url: https://www.reddit.com/r/androiddev/comments/jib3l7/android_11_scoped_storage_mediastore_can_create/
---
The first video in the references below has a good (non-Google) overview of the issues with MediaStore for Android 11 - and how some apps are having to use a mixture of SAF and MediaStore to cobble things together.


In this video they suggest keeping your files in your own folder - this way you avoid the 128 persistent permissions issue, as all the files within that folder get permission (if you have permission for that one folder).

Still some issues with System Picker - as varies by manufacturer.

They also resolve one question I had - whether one can create a folder within Music and other "shared storage" areas.

They mention you can delete as well.

We have had experience with MediaStore from some years ago - when it was really kludgy - sometimes a file would appear in MediaStore, sometimes not.

Perhaps it has improved by now.

**Questions:**

- **how easy is it to move a file from your own folder (within Music etc.) - and move it to a further sub-folder there.  Is that instantaneous (as it used to be earlier when both source and destination were on internal storage) ?**

- **when the Android 11 FAQ by Google on medium - https://medium.com/androiddevelopers/android-11-storage-faq-78cefea52b7c - Android 11 storage FAQ - suggests that fopen() can be used now - does that mean you can programmatically deal with file paths strings, adding on a suffix to point to a sub-folder (as you used to do before) ?  That is, can one extract a file path from a Uri that MediaStore returns ?  And then add on a sub-folder to the file path string and use that with fopen() ?**


----

**References:**

https://www.youtube.com/watch?v=32Jox0itYKI
Android 11 #1: Deep Dive into Scoped Storage &amp; Privacy
Android Academy Global
July 7, 2020

rough transcript:

&gt;2:19:30 minute mark
&gt;
&gt;128 persistent URIs ..
&gt;
&gt;when ask for 129 will have issue ..
&gt;
&gt;
&gt;2:21:00 minute mark
&gt;
&gt;suggest keep documents inside own folder
&gt;
&gt;(so means can do sub-folder etc. !?)
&gt;
&gt;
&gt;can hold permission to directory .. which easier ..
&gt;
&gt;
&gt;can create directory
&gt;
&gt;can delete files
&gt;
&gt;
&gt;2:25:00 minute mark
&gt;
&gt;migration ..
&gt;
&gt;
&gt;2:28:00 minute mark
&gt;
&gt;major problems with SAF
&gt;
&gt;
&gt;system picker provided by manufacturer ..
&gt;
&gt;problems with cloud based providers
&gt;
&gt;
&gt;FileProvider for sharing with other apps ..
&gt;
&gt;
&gt;they had issues with files ..
&gt;
&gt;migrated to MediaStore .. still have issues ..
&gt;
&gt;
&gt;so still requesting legacy storage ..
&gt;
&gt;barrier is getting less ..
&gt;
&gt;also using SAF .. still using legacy storage sometimes (!?)
&gt;
&gt;
&gt;1:09:00 - users not found user-friendly
&gt;
&gt;so had to use legacy ..
&gt;

----

https://medium.com/androiddevelopers/android-11-storage-faq-78cefea52b7c
Android 11 storage FAQ
## [10][Firebase Crashlytics Integration With Github](https://www.reddit.com/r/androiddev/comments/jie4lc/firebase_crashlytics_integration_with_github/)
- url: https://www.reddit.com/r/androiddev/comments/jie4lc/firebase_crashlytics_integration_with_github/
---
Good day guys,

Hope you are safe and in good health.

Can anyone here guild this fellow android developer on how to integrate firebase Crashlytics with Github.

The idea is if there's any major crash on app, it should automatically get issued on our github repository.

Thank you.
## [11][Android Studio 4.1 now has a SQL database viewer.](https://www.reddit.com/r/androiddev/comments/jid9we/android_studio_41_now_has_a_sql_database_viewer/)
- url: https://www.reddit.com/r/androiddev/comments/jid9we/android_studio_41_now_has_a_sql_database_viewer/
---
How do you browse your Android app databases? do you really have to copy the files back and forth?
## [12][Can't start rollout of an app due to app permissions error](https://www.reddit.com/r/androiddev/comments/ji5y4q/cant_start_rollout_of_an_app_due_to_app/)
- url: https://www.reddit.com/r/androiddev/comments/ji5y4q/cant_start_rollout_of_an_app_due_to_app/
---
I'm rolling out a new version of an app but for the first time needed to submit a declaration regarding the location permissions I'm requesting.

I've submitted my declaration (as "no, I'm not accessing location in the background") and the page shows a green tick with a "Submitted" label, but my "Review and release" page is still showing it as an error, and the "Start rollout to production" button is disabled.

I thought the review process was supposed to start by clicking "Start rollout to production"? Or is Google still reviewing the declaration, even though this release is still only a draft? It's been like this since last Friday.

[https://imgur.com/a/hyuRkbb](https://imgur.com/a/hyuRkbb)

EDIT: at the moment, it looks like updating the target version from 28 (Android 9) to 29 or 30 (Android 10/11) makes the Play Console accept it.
