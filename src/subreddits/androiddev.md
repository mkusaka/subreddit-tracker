# androiddev
## [1][App Feedback Thread - October 24, 2020](https://www.reddit.com/r/androiddev/comments/jh83jn/app_feedback_thread_october_24_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jh83jn/app_feedback_thread_october_24_2020/
---
This thread is for getting feedback on your own apps.

####Developers:

- must **provide feedback** for others
- must include **Play Store**, **GitHub**, or **BitBucket** link
- must make top level comment
- must make effort to respond to questions and feedback from commenters
- may be open or closed source

####Commenters:

- must give **constructive feedback** in replies to top level comments
- must not include links to other apps

To cut down on spam, accounts who are too young or do not have enough karma to post will be removed. Please make an effort to contribute to the community before asking for feedback.

As always, the mod team is only a small group of people, and we rely on the readers to help us maintain this subreddit. Please report any rule breakers. Thank you.

\- Da Mods
## [2][Weekly Questions Thread - October 19, 2020](https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/
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
## [3][I want to revise and relearn the android development after working 4 years as a Android developer.](https://www.reddit.com/r/androiddev/comments/jh4z3r/i_want_to_revise_and_relearn_the_android/)
- url: https://www.reddit.com/r/androiddev/comments/jh4z3r/i_want_to_revise_and_relearn_the_android/
---
Hello,

I have been working as a android developer for last 4 years and worked on many projects. But still i find myself lost when i try to go in depth of any complex issue. I was a lazy developer and didnt give much time for learning new things and go in depth.  I didnt use any design pattern in my projects. And when i give  any interview i cant show confidence. This is making me frustrated. 

so i want to relearn everything and go in depth of everything. I wanna know everything about android development and architecture. 

Now i need your suggestion. Please suggest me which things i should learn to be a good android developer and save me from my frustration? Please guide me. 

Thank you.
## [4][Improving TextView &amp; Spannable Performance](https://www.reddit.com/r/androiddev/comments/jh7k2f/improving_textview_spannable_performance/)
- url: https://androidexplained.github.io/android/ui/2020/10/12/improving-spannable-performance.html
---

## [5][Legitimacy of a post that claims: Aarogya Setu app, the covid 19 tracking app of India is recording video while in background.](https://www.reddit.com/r/androiddev/comments/jgwrwm/legitimacy_of_a_post_that_claims_aarogya_setu_app/)
- url: https://www.reddit.com/r/androiddev/comments/jgwrwm/legitimacy_of_a_post_that_claims_aarogya_setu_app/
---
Hi all, while my usual reddit scrolling, I found a post on r/India claiming that the Aarogya Setu app, the covid 19 tracking app of India is recording video while in background.

Here's the link to post : https://www.reddit.com/r/india/comments/jgrbt1/the_aarogya_setu_app_on_my_moms_phone_was/?utm_medium=android_app&amp;utm_source=share

Here's my comment on the same post :
https://www.reddit.com/r/india/comments/jgrbt1/the_aarogya_setu_app_on_my_moms_phone_was/g9sz2zg?utm_medium=android_app&amp;utm_source=share&amp;context=3

What are your thoughts? This feels somewhat important as it can totally be a fake news and people are overreacting to it. Or at the same time, something really messed up is happening.
## [6][Observations of "requestLegacyExternalStorage" behavior on Android 10 and Android 11](https://www.reddit.com/r/androiddev/comments/jh5uif/observations_of_requestlegacyexternalstorage/)
- url: https://www.reddit.com/r/androiddev/comments/jh5uif/observations_of_requestlegacyexternalstorage/
---
I have tested an app on emulator running Android 10 (API 29) and Android 11 (API 30), with the addition of the line:

AndroidManifest.xml:

&gt;android:requestLegacyExternalStorage="true"

This is the method suggested by Google to allow apps to continue to work as before (avoiding the new "scoped storage" restrictions in Android 10 and Android 11).

----

**Results on Android 10 (API 29)**

When the app is targeting Android 10 (API 29) - i.e. using build.gradle:

&gt;compileSdkVersion 30
&gt;
&gt;targetSdkVersion 29

then the app works as before (legacy behavior) on BOTH Android 10 and Android 11:

- emulator running Android 10 (API 29)

- emulator running Android 11 (API 30)

That is, the files your app creates in a top level folder on internal storage ARE created correctly.

The app can also delete/rename files as well.

AND these files remain visible and accessible when you try to view them using the built-in Files app (file manager).

This means if you choose to transition to one of the "shared storage" folders like Music or Downloads - you will be able to copy over files there, and you will be able to delete the files from the old location.

----

**Results on Android 11 (API 30)**

However, when you target Android 11 - i.e. using build.gradle:

&gt;compileSdkVersion 30
&gt;
&gt;targetSdkVersion 30


then the app works as before (legacy behavior) on Android 10 (API 29):

- emulator running Android 10 (API 29)


but does NOT work as before on Android 11 (API 30):

- emulator running Android 11 (API 30)


i.e. on Android 11, it no longer supports the legacy behavior fully - details:

- the app can read the files it had created before

- but the app cannot write files in the old location

- the app cannot delete/rename files in the old location


Essentially the app retains access to the old top level folder in read-only mode only.

This means that once you target Android 11 (API 30) - you will be able to copy over the old files (to the Music or Downloads "shared storage" folder using MediaStore methods), but you will NOT be able to clean up the old files (since  will not have write access).

----

**Conclusion**

This means there IS a way to make your app work as before for storage (legacy behavior) - on Android 10, and on Android 11.

Just target Android 10 (API 29) - with build.gradle:

&gt;compileSdkVersion 30
&gt;
&gt;targetSdkVersion 29

This will cover devs up to late 2021 or thereabouts - at which point they will be forced by Google to start targeting Android 11 (just as they are being forced to target Android 10 by Nov 2020).

And as described above, if you choose to target Android 11 (API 30), your app will not fully support legacy behavior for storage.

- the app will still be able to read old files

- but will not be able to create new files, or delete the files it created before


----

**For those thinking of transitioning files to a MediaStore location**

This means if an app wants to transition from old location to new "shared storage" locations (like Music or Downloads using MediaStore methods) - it will be able to copy the files even if it starts targeting Android 11 (API 30).

But if it wants to also cleanup (delete) the old files - it will only be able to do so while it is still targeting Android 10 (API 29) - not once it starts targeting Android 11 (API 30).

----

**Questions**

Android 11 emulator did demonstrate bad behaviors - sluggish notifications (not disappearing when they should) - this issue has also been reported by others on stackoverflow.

Same app targeting Android 11 misbehaves when run on Android 11 emulator - widgets misbehave (perhaps this requires some change to code for this to work on Android 11).
## [7][Free messaging for Android app](https://www.reddit.com/r/androiddev/comments/jh71u8/free_messaging_for_android_app/)
- url: https://www.reddit.com/r/androiddev/comments/jh71u8/free_messaging_for_android_app/
---
Is there any service which gives you ACTUALLY free messaging for your app?  What I mean by this is not having the *indirect* costs of Firebase.
## [8][Curated UDACITY free course A-Z Android Development](https://www.reddit.com/r/androiddev/comments/jh5knl/curated_udacity_free_course_az_android_development/)
- url: https://docs.google.com/document/d/1sF2z3XEkKM9h-mL-v96_MwdxFJFWXEk7FlehUUJgBvI/edit?usp=drivesdk
---

## [9][Material icons seem to be down?](https://www.reddit.com/r/androiddev/comments/jh0b66/material_icons_seem_to_be_down/)
- url: https://material.io/resources/icons
---

## [10][“Custom column types” in SQLDelight: Android](https://www.reddit.com/r/androiddev/comments/jh54s0/custom_column_types_in_sqldelight_android/)
- url: https://developerkp16.medium.com/custom-column-types-in-sqldelight-android-a6f166635464
---

## [11][How to keep foreground service running even if i kill the app manually?](https://www.reddit.com/r/androiddev/comments/jh4d1h/how_to_keep_foreground_service_running_even_if_i/)
- url: https://www.reddit.com/r/androiddev/comments/jh4d1h/how_to_keep_foreground_service_running_even_if_i/
---
I have created a foreground service which works on main thread i.e Service and not Intend service and it is working fine if i open all of my phone app. The service keeps running in background but as soon as i kill the app from task bar it get stopped. So how do i keep the serving running even if i kill the app manually.
(now some of you may say after android O, services can't run in background but im using foreground service and user is being notified through notification that service of X app is running in background)
## [12][3 Practical Steps To Master Android DataStore](https://www.reddit.com/r/androiddev/comments/jh70xg/3_practical_steps_to_master_android_datastore/)
- url: https://www.reddit.com/r/androiddev/comments/jh70xg/3_practical_steps_to_master_android_datastore/
---
3 practical steps with migration: a brief story of how we implemented Jetpack DataStore in Dott.  
Do you use it in your project? 😊  
[https://vladsonkin.com/3-practical-steps-to-master-android-datastore/](https://vladsonkin.com/3-practical-steps-to-master-android-datastore/)
