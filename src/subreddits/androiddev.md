# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/jmmr24/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jmmr24/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - November 02, 2020](https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/
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
## [3][Android 11 dodges a bullet - apps creating a folder at top level maybe able to simply move that to Music/Photos "shared storage" folder (requiring single line change in java) - without needing to resort to complications of SAF](https://www.reddit.com/r/androiddev/comments/jn5azi/android_11_dodges_a_bullet_apps_creating_a_folder/)
- url: https://www.reddit.com/r/androiddev/comments/jn5azi/android_11_dodges_a_bullet_apps_creating_a_folder/
---
**EDIT:** what is described below applies not only for File API for java - but also for your C code i.e. apps using JNI/NDK native C libraries (if you are doing fopen(), and other standard file io).  I say this because our tests included native file io using C as well.


**Summary**

Google is moving to restrict android storage.  They had initially telegraphed a much stronger change that would have broken android.  For Android 11 someone at Google seems to have convinced the others that retaining file paths and fopen() is essential (this was something we have been harping about for ages on reddit - as absence of file paths and fopen() spelled the death of standard storage).

Here I provide a quick overview of the storage changes, and advice for migrating for app developers who **do not want to spend time on storage migration**.  Specifically developers who have no interest in spending time on Storage Access Framework (SAF) - the flawed and inefficient "alternative" that Google tried to push devs to adopt (much like they pushed SAF as the alternative when they killed seamless ext SD card access in KitKat).

Many apps just need ability to save files to a location that will be persistent (not go away once app is uninstalled).  This is the case for apps like audio recorders, camera apps and such.

That is now possible with something as little as a one line change to your code for Android 11.

The end result will be that you will not need to change your app's file handling (except one or two lines of java code).  The simplest of apps (like audio recorder apps) will only need to change one line, and keep behaving much as before.

&amp;nbsp;


**Backstory**

As discussed here before, Google has been on a march to kill traditional storage on Android.

Just as Google killed seamless external SD access with KitKat (and later providing an inadequate replacement - SAF - which expectedly never took off, leading to the demise of seamless ext SD card storage) - similarly Google had announced a flurry of changes for storage.  These changes are expected to make persistent storage as before harder to do.  Because the only way to continue using old storage code was to use the app-specific folders (which are removed when app is uninstalled).  This would have left cloud storage as an attractive alternative (to mirror the app-specific folders) - with few other easy options for storage persistence.

Use of SAF is non-trivial for devs, and it comes with it's own set of caveats and performance limitations.  In addition, there was earlier a shadow over use of SAF as well (whether one would need Google Permissions Declaration Form for this as well - since SAF does allow writing in many more places and currently is used to routinely grant top folder access).  Now for Android 11, Google medium.com post has clarified that SAF does not require special permission from Google - and Google themselves will limit SAF so it cannot access the top level folder, and some other folders (this means those devs using SAF will need to check user flows to ensure their SAF use works under new restrictions).

&amp;nbsp;


**Android 11 solution**

Android 11 arrives with changes:

- file paths can be used as before and File API - for a few specific folders (Music, Photos .. i.e. the so-called "shared storage" folders).

- fopen(), delete, instantaneous move of files - can be done  (again for a few specific folder locations)

- these capabilities were not available in Android 10


In practice this means an app could choose to no longer house it's app folder (where it stores persistent audio recordings etc.) at the top level folder on internal storage - but instead locate it in the Music folder (which is one of the "shared storage" folders).

If your app saves files in a folder "folder1" (that was previously located at top folder) - that "folder1" now can be saved in the Music folder.

Just change this line in your code - where you discover the parent directory where "folder1" should be stored:

&gt;**File sdcardRoot = Environment.getExternalStorageDirectory();**

to:

&gt;**File sdcardRoot = Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_MUSIC);**

And similarly for Photos etc.  For Downloads there is some additional restriction (apps cannot see files created by other apps).  While for Music/Photos etc. apps CAN see files (read-only) created by other apps (as long as you keep using the READ_EXTERNAL_STORAGE permission in AndroidManifest.xml.

Now your "folder1" will be located in Music/folder1, but you can continue to use the rest of your code as before.  Manipulating file path strings etc. ..

&amp;nbsp;


**Android 11 caveats**

The only caveat or restriction is:

- if you use the Music folder, you can only create "audio" files there (.wav, .ogg, .mp3 and perhaps others).  If you need to create a dummy file "dummy", you can create it, but you will have to name it "dummy.mp3" etc. i.e. with an audio-like extension.

- you can create folders within the Music folder - example: Music/folder1

- two apps can use the same folder i.e. app1 creates folder1 and app2 also creates folder1.  One app can delete the folder created by another app (if folder is empty).  Files created by app1 can be read by app2 (if it uses the READ_EXTERNAL_STORAGE permission), but cannot be written or deleted by app2.  This means if you delete folder1 from app1, it will delete all the app1-created files in folder1, but will leave the files created by app2 there untouched (and so folder1 will not be deleted).  But if Music/folder1 was created by app1, it can be deleted by app2 (if the folder1 is empty or only contains files created by app2).

&amp;nbsp;


**Android 10 and earlier**

Since Android 10 was missing these file path and fopen() capabilities, that means it will cause problems if you don't use "requestLegacyExternalStorage=true" in your AndroidManifest.xml.

This is why Google also recommends that apps use this flag in your AndroidManifest.xml:

requestLegacyExternalStorage="true"

This will allow their app to perform the same as before all through to Android 10.  And somewhat so on Android 11 as well (as long as app is targeting below Android 11).

Once your app starts targeting Android 11, this "requestLegacyExternalStorage" will be ignored.

This means once you start targeting Android 11 (targetSdkVersion=30) your app should be using "Music/folder1" etc. instead of "folder1".


Thus, the app developer HAS to ship his app for Android 10 using the "requestLegacyExternalStorage" flag set to TRUE (to opt out of the new storage changes) - if they want to not change their app code.

If you don't use this for Android 10, then your app will be subject to Android 10 rules, and because Android 10 did not have file path and fopen() support, you will not be able to introduce the "Music/folder1" way of doing things.

So keep using "requestLegacyExternalStorage" while you targetSdkVersion=29 (Android 10).

Once you targetSdkVersion=30 (Android 11), the "requestLegacyExternalStorage" is ignored, and your app should be ready to use "Music/folder1" etc.  So you should have a behavior in place so files are stored in the Music folder or Photos folder (one of the "shared storage" folders) instead of at top level folder of internal storage.

&amp;nbsp;


**How to adapt to new restrictions**

Google has announced that Android 11 will now again support File API and fopen() type methods (Android 10 did not - i.e. if you were targeting Android 10).

The only restriction in Android 11 is that these capabilities can only be used for files and folders that are stored within Music, Photos etc. - the so-called "shared storage" folders.

This means all you have to do is ensure the folder where you saved audio recorder files (usually a folder at top level of internal storage), can now be saved within the Music folder on internal storage:

change:

&gt;**File sdcardRoot = Environment.getExternalStorageDirectory();**

to:

&gt;**File sdcardRoot = Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_MUSIC);**


And make sure you are using this in your AndroidManifest.xml (as Google recommends, this is to cover for the aberration that was Android 10 which does not support file paths and fopen() - Android 11 will ignore this flag):

&gt;**requestLegacyExternalStorage="true"**

&amp;nbsp;


Eveything else can be kept as before - you can:

- create a folder within this Music folder (just as you created a folder at top level on internal storage)

- you can manpipulate the path, create a path for a sub-directory by appending to the file path

- you can create a folder, and create files there

- basically nearly all your old java code and NDK/JNI native C code will work as before - use fopen() using file path strings, manipulate path strings etc. (just make sure the paths you want to reference are within the Music folder)

&amp;nbsp;


What you cannot do:

- you can only create audio files (more precisely files that have extension that indicate it is a file like .wav, .ogg, .mp3 etc.) within the Music folder (similar restrictions may apply to Photos).

- evidently the file extension is the only thing used to screen - so you can create a file holding arbitrary data - just ensure it is named file.mp3 etc. (standard music file extensions)

- if you try to create a file that does not have an audio extension, or another type of extension, it will fail

&amp;nbsp;


Some other different behaviors:

- two apps can write to the same folder

- so you can have two of your apps write to the same folder (within Music for example)

- a folder created by app1 can be deleted by another app2 (if it is empty)

- a file created by app1 cannot be deleted by another app2

- this means app2 cannot delete a folder that contains a file created by app1

- a file created by app1 CAN be read by app2 (if app2 uses the READ_EXTERNAL_STORAGE permission in it's AndroidManifest.xml)


&amp;nbsp;


**Thanks**

Thanks to /u/oneday111 for [outlining the possibilities](https://www.reddit.com/r/androiddev/comments/jib3l7/android_11_scoped_storage_mediastore_can_create/ga6i0lw/) - which led to testing app behavior when the app folder is simply relocated to Music folder etc.


&amp;nbsp;


**NOTE TO MANUFACTURERS**

Please ensure your devices running Android 11 use the source tree with the latest changes for Android 11.

Because (as has happened before) - manufacturers sometimes choose an earlier Beta as their starting point (which can sometimes miss the final behaviors promised).

So manufacturers, please don't mess up by failing to comply with this file path and fopen() behavior in Android 11 - since this is an essential feature of Android.  If you fail to ensure this is supported in your Android 11 version, a huge number of apps will break.

I say this because the storage nuances seem to have been changing a lot in the last few months - so it is possible that a manufacturer picks up a Beta version as their starting point - but which fails to have the final behaviors now promised for storage in Android 11.
## [4][Jetpack Compose on Desktop is in preview!](https://www.reddit.com/r/androiddev/comments/jn9swn/jetpack_compose_on_desktop_is_in_preview/)
- url: https://twitter.com/iateyourmic/status/1323551064512172033
---

## [5][Play Console: Way to see source of third-party referrals?](https://www.reddit.com/r/androiddev/comments/jn8wj3/play_console_way_to_see_source_of_thirdparty/)
- url: https://www.reddit.com/r/androiddev/comments/jn8wj3/play_console_way_to_see_source_of_thirdparty/
---
I have a small app, published on the Play Store, that is jugging along at a steady pace of about 200 downloads a day. However, a couple of weeks ago, I got two big spikes on specific days with thousands of downloads.

Looking at the statistics, I can see that these spikes came from one country and almost exclusively through third-party referrals. So someone must have posted something about this app online.

Is there a way to find the source of the referrals in the Play Console? I've tried googling for mentions from only that country and only in the past month, but have had no luck.

Thanks for any ideas :)
## [6][Android Intent &amp; Bundle extensions that insert and retrieve values elegantly.](https://www.reddit.com/r/androiddev/comments/jmqvh0/android_intent_bundle_extensions_that_insert_and/)
- url: https://github.com/skydoves/bundler
---

## [7][How to go back from AOSP to OHA](https://www.reddit.com/r/androiddev/comments/jna7rk/how_to_go_back_from_aosp_to_oha/)
- url: https://www.reddit.com/r/androiddev/comments/jna7rk/how_to_go_back_from_aosp_to_oha/
---
Hey ! 

Sorry if it's not the good place, i'm new to reddit

I tried rooting my new Google Pixel 4a with magisk and got some kind of error at the end (the phone said "corrupt" and wouldn't boot).  
So i use the automatic android flashing tool for pixels and nexus ([https://flash.android.com/welcome](https://flash.android.com/welcome))

It worked well, but it installed the AOSP version of android (open source and without google services). I just wanted to go back to a fresh install of Android 11 lol.  
The flashing tool only give me AOSP android version sadly.  
How to get a fresh OHA android 11 install ?

Thanks !
## [8][Versioning Android apps](https://www.reddit.com/r/androiddev/comments/jn9l5k/versioning_android_apps/)
- url: https://maxirosson.medium.com/versioning-android-apps-d6ec171cfd82
---

## [9][What Could a Debug Menu Contain?](https://www.reddit.com/r/androiddev/comments/jmqg7f/what_could_a_debug_menu_contain/)
- url: https://halcyonmobile.com/blog/mobile-app-development/android-app-development/what-could-a-debug-menu-contain/
---

## [10][How can I implement a voice chat bot with User Interface in Android Studio ?](https://www.reddit.com/r/androiddev/comments/jn8qzp/how_can_i_implement_a_voice_chat_bot_with_user/)
- url: https://i.redd.it/hm321hk8h0x51.png
---

## [11][Automate Dependencies Upgrades With Releases Hub](https://www.reddit.com/r/androiddev/comments/jn8835/automate_dependencies_upgrades_with_releases_hub/)
- url: https://www.reddit.com/r/androiddev/comments/jn8835/automate_dependencies_upgrades_with_releases_hub/
---
Automatically keep your Gradle project dependencies up to date.

The **Releases Hub Gradle Plugin** helps developers to keep their dependencies up to date, reducing some tedious manual tasks like remembering to look for dependencies upgrades, upgrading the dependencies on the Gradle configuration and creating a PR with the changes.

[https://medium.com/swlh/automate-dependencies-upgrades-with-releases-hub-8eac6d7f43d6](https://medium.com/swlh/automate-dependencies-upgrades-with-releases-hub-8eac6d7f43d6)
## [12][New subscription server notifications available in production - News](https://www.reddit.com/r/androiddev/comments/jn86m6/new_subscription_server_notifications_available/)
- url: https://developer.apple.com/news/?id=an960mux
---

