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
## [3][Android 11 scoped storage - MediaStore can create sub-directories, delete - but are rename-file/move-file instantaneous as before ?](https://www.reddit.com/r/androiddev/comments/jib3l7/android_11_scoped_storage_mediastore_can_create/)
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
## [4][Can't start rollout of an app due to app permissions error](https://www.reddit.com/r/androiddev/comments/ji5y4q/cant_start_rollout_of_an_app_due_to_app/)
- url: https://www.reddit.com/r/androiddev/comments/ji5y4q/cant_start_rollout_of_an_app_due_to_app/
---
I'm rolling out a new version of an app but for the first time needed to submit a declaration regarding the location permissions I'm requesting.

I've submitted my declaration (as "no, I'm not accessing location in the background") and the page shows a green tick with a "Submitted" label, but my "Review and release" page is still showing it as an error, and the "Start rollout to production" button is disabled.

I thought the review process was supposed to start by clicking "Start rollout to production"? Or is Google still reviewing the declaration, even though this release is still only a draft? It's been like this since last Friday.

[https://imgur.com/a/hyuRkbb](https://imgur.com/a/hyuRkbb)
## [5][Android Always On Display api](https://www.reddit.com/r/androiddev/comments/jiavb2/android_always_on_display_api/)
- url: https://www.reddit.com/r/androiddev/comments/jiavb2/android_always_on_display_api/
---
Hi, how can I create an AOD app?(using android studio) 
Anyone know about an android AOD api guide or something?
## [6][What is READ_EXTERNAL_STORAGE permission needed for?](https://www.reddit.com/r/androiddev/comments/jibvaj/what_is_read_external_storage_permission_needed/)
- url: https://www.reddit.com/r/androiddev/comments/jibvaj/what_is_read_external_storage_permission_needed/
---
I noticed that an app is able to access any audio file on a device without having been granted the READ_EXTERNAL_STORAGE permission.

I checked the official Android documentation and it's written in poor English:

https://developer.android.com/reference/android/Manifest.permission#READ_EXTERNAL_STORAGE

I was under the assumption that "External Storage" in Android is any storage that is not specific to your app. Am I wrong in that? Perhaps this permission is only needed to access sd cards?
## [7][Firebase Crashlytics Integration With Github](https://www.reddit.com/r/androiddev/comments/jie4lc/firebase_crashlytics_integration_with_github/)
- url: https://www.reddit.com/r/androiddev/comments/jie4lc/firebase_crashlytics_integration_with_github/
---
Good day guys,

Hope you are safe and in good health.

Can anyone here guild this fellow android developer on how to integrate firebase Crashlytics with Github.

The idea is if there's any major crash on app, it should automatically get issued on our github repository.

Thank you.
## [8][Android Studio 4.1 now has a SQL database viewer.](https://www.reddit.com/r/androiddev/comments/jid9we/android_studio_41_now_has_a_sql_database_viewer/)
- url: https://www.reddit.com/r/androiddev/comments/jid9we/android_studio_41_now_has_a_sql_database_viewer/
---
How do you browse your Android app databases? do you really have to copy the files back and forth?
## [9][My Play Console says I have a message but nothing is in the inbox. What if this is an app suspension? This is a very bad bug to have.](https://www.reddit.com/r/androiddev/comments/jiafeb/my_play_console_says_i_have_a_message_but_nothing/)
- url: https://i.redd.it/8fq7psl88ev51.png
---

## [10][Just published a new Android library. A simple screen that is shown when your app crashes that includes the crash details instead of the normal crash dialog. It's very similar to the one in Flutter.](https://www.reddit.com/r/androiddev/comments/jhjfzu/just_published_a_new_android_library_a_simple/)
- url: https://github.com/mlegy/red-screen-of-death
---

## [11][Phone going on Reboot cycle.](https://www.reddit.com/r/androiddev/comments/ji6ybg/phone_going_on_reboot_cycle/)
- url: https://www.reddit.com/r/androiddev/comments/ji6ybg/phone_going_on_reboot_cycle/
---
Phone in reboot cycle and it's 4 year old phone.

Any help form this short log between reboots - on what is casuing it to reboot often.

1/4 CPU gone bad?

[https://filebin.net/j9astmfteeo1umew/last\_kernel-1026-065243.log?t=1mo4ab3u](https://filebin.net/j9astmfteeo1umew/last_kernel-1026-065243.log?t=1mo4ab3u)
## [12][apps supporting the full screen size/shape of the pixel 4 and 5](https://www.reddit.com/r/androiddev/comments/jhxbyy/apps_supporting_the_full_screen_sizeshape_of_the/)
- url: https://www.reddit.com/r/androiddev/comments/jhxbyy/apps_supporting_the_full_screen_sizeshape_of_the/
---
hi

what do app developers need to do to fully support the screen size/shape of the pixel 5 (and partially the pixel 4)?

i have a pixel 5 (and previously the pixel 4), and some, but not all, apps support using all of the screen, however some apps don't use the very bottom part, and others don't support both the bottom and the  top?

this is very jarring when an app ends up with a black section at the top and bottom of the screen.  
it's even worse when apps that play video don't support it. prime example being the plex app. i play a video, in landscape mode, and you end up with a black section on the left.

see examples of two apps that have either one or both the "issues".  


  


https://preview.redd.it/k1g8magw0av51.jpg?width=1080&amp;format=pjpg&amp;auto=webp&amp;s=595cedc86540298c7ea2f779a7a46b5fe4318d12

https://preview.redd.it/1ahrepov0av51.jpg?width=1080&amp;format=pjpg&amp;auto=webp&amp;s=0cffccb4a26e6ddc3af9eaf97289b68556c9ef4d
