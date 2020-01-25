# androiddev
## [1][App Feedback Thread - January 25, 2020](https://www.reddit.com/r/androiddev/comments/etq7hg/app_feedback_thread_january_25_2020/)
- url: https://www.reddit.com/r/androiddev/comments/etq7hg/app_feedback_thread_january_25_2020/
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
## [2][Weekly Questions Thread - January 20, 2020](https://www.reddit.com/r/androiddev/comments/erbsa6/weekly_questions_thread_january_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/erbsa6/weekly_questions_thread_january_20_2020/
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
## [3][On the Storage Apocalypse: Isolating existing apps from negative reviews from Android 10 users](https://www.reddit.com/r/androiddev/comments/etm048/on_the_storage_apocalypse_isolating_existing_apps/)
- url: https://www.reddit.com/r/androiddev/comments/etm048/on_the_storage_apocalypse_isolating_existing_apps/
---
**NOTE:** jump to the "Splitting an app" section below to skip the discussion section

**EDIT:** as pointed out in Tolriq's comment thread - https://www.reddit.com/r/androiddev/comments/etm048/_/ffhoicp - distributing the two APKs as Multiple APKs or if you are willing to use App Bundles, then as an App Bundle - these are also options available to you.  However, with these you do not get an opportunity to phrase your Description differently to the two demographics - for example if you want to NOT publicize an old feature that you can no longer support with the new crippled storage options, then you wont be able to say that (unless you want to add if-then statements to your Description that will confuse earlier android version users as well).

&amp;nbsp;

----

With the upcoming extinction of storage as we know it in Android 10, and with the alternatives being still not clearly defined:

- Scoped Storage (using MediaStore)

- Storage Access Framework (SAF) - which may require approval by Permissions Declaration Form in the future (see the first link in the References section at the bottom)

&amp;nbsp;


Some developers may be unable to update their apps successfully to the new models - some may be short of time to update all their apps, and some apps' feature sets may not translate well to the new limitations (even using the MediaStore/SAF alternatives).

Once Android 10 users encounter their app, it will lead to bad ratings, and support issues.  And we are not aware if the storage changes for Android 10 are the final change - it could be reversed in the future if it turns out to be a disaster.

In such a situation, some developers may want to isolate their previous apps from the headache that will accompany Android 10 (as users ask where their files went and cannot be found).

Some apps which were designed from the ground up to have access to persistent storage access on local storage, may not be fully constrainable to the new model (where files saved are non-persistent or stored in some sandbox no longer visible to other apps).

Knowing the limitations of Android 10, developers may want to not promise some features (that they used to promise for previous Android versions).

All these changes militate for a split in how your apps are going to target users - if you want to save yourself from needless support distractions.

&amp;nbsp;


The solutions are:

- limit the apps to no longer be visible to Android 10 users

- split those apps so the older version is not visible to Android 10 users, but newer versions (which you can appropriately set Description and capabilities to reflect the newer limited features)



&amp;nbsp;



We did this with a hearing aid app we made some time ago which was using a new android audio engine with low latency.  This audio engine was promised to be available for Oreo 8.0.  However, on arrival even though their docs and support forum still said it worked great for Oreo 8.0, the reality was that it did not work for half the Oreo 8.0 devices out in the wild (Google engineers were focused primarily on testing on Pixel devices and didn't know their much publicized engine did not work on half the Android Oreo 8.0 devices out there).

As a result the app we created was a disaster from day one, as all manner of devices were showing up as being unable to run the app.

Some time later, and after prodding by developers, Google started acknowledging that the engine won't work for Oreo 8.0 after all (they took even longer to change their promotional material which continued to lure developers into using it for Oreo 8.0).  In the end, the engine was stated to work on Oreo 8.1.

Because our app was already out there, and suffering from hordes of 1-stars, we split the app - one targeting up to Oreo 8.0, and one (new app) targeting Oreo 8.1.

This may not have been ideal, but it allowed all the 1-star reviews to go to the Oreo 8.0 and lower version, while at least the Oreo 8.1 version was saved from such criticism.

The older app (which now became an Oreo 8.0 and lower app) was not visible to Oreo 8.1 users and above, and the new Oreo 8.1+ app was not visible to Oreo 8.0 and lower users.  Thus we split the demographic - directing all the problematic Oreo 8.0 users to go to the old app (which the by now bad ratings) - while directing new  Oreo 8.1+ users to the new app (which thus had a chance to get good ratings).

For Storage we need to so something similar.

We need your old app to be limited to below Android 10, so it is saved from the ire of Android 10 users.  And if you want to have a version for Android 10 and above - that you do with a new app.


&amp;nbsp;


**Splitting an app**

Here are the API numbers for Pie and Android 10:

https://developer.android.com/guide/topics/manifest/uses-sdk-element

&gt;Android 10.0 - API 29
&gt;
&gt;Android Pie 9.0 - API 28

OldApp's AndroidManifest.xml probably did not have a maxSdkVersion - because you were not previously limiting it's visibility to higher android versions (since Android previously used to assure that all old apps would run on newer Android versions - no more).

&amp;nbsp;


**OldApp's AndroidManifest.xml** now should limit it's visibility up to below Android 10 - i.e. up to Android Pie 9.0 (API 28):

Add:

&gt;android:maxSdkVersion=28

You can keep your old minSdkVersion etc. as before, since you are not limiting behavior for older versions of Android.

&amp;nbsp;



**Reference:** As you can see the Android documentation is already obsolete - here Google argues how old versions of the app will always work on new Android versions.  The old mantra - now discarded with these Storage changes in Android 10 - is no longer being honored by Google (but it's vestiges still remain in the docs, misguiding users):

https://developer.android.com/guide/topics/manifest/uses-sdk-element

&gt;By design, **new versions of the platform are fully backward-compatible**. Your application should work properly on new versions, provided it uses only standard APIs and follows development best practices. 


&amp;nbsp;


With the changes to maxSdkVersion above, Android 10 users will not see OldApp on Google Play Store.

If you decide to create a new app - or even just use the same app (**Google will not penalize you for Repetitive Content because these two apps will address separate demographics** - each seeing only one version) - with this new app you now get the advantage that you can couch it's Description to the new Android 10 users, caution them about limitations to storage in the Description, and possibly in the app's help section or what's new section as well.  And you get the ability to actually LIMIT or remove those features which you could not translate well over to the new Android 10 storage model (where files saved are saved instead to a sandbox and are not persistent in the old sense).  Even if you use MediaStore or SAF, use of those may still limit some features from being translated well, and having a separate version of the app will allow you the freedom to tinker with the presentation of those features (independent from how your old app used to do it).

&amp;nbsp;


**NewApp's AndroidManifest.xml:**

You probably already had a minSdkVersion setting (probably was set to API 22 or some such) - well now you set it to Android 10 (API 29) - as this new app will target only Android 10 and above users:

change:

&gt;android:minSdkVersion=22 (or whatever you have)

to:

&gt;android:minSdkVersion=29

You don't need to specify the maxSdkVersion for the NewApp because this will work for all versions Android 10 and above.

&amp;nbsp;


----

**References:**

- [Android apps will start to lose ability to access local persistent storage and ext SD card - in a move which will boost Google's cloud storage for app data backup](https://www.reddit.com/r/androiddev/comments/e5fb9m/android_apps_will_start_to_lose_ability_to_access/)

- [The Death of External Storage: The End of the Saga(?)](https://commonsware.com/blog/2019/06/07/death-external-storage-end-saga.html)

- [Hey, Google. Where is your roadmap ? Why commercial viability for indie devs is going down, and Google Play is dead for indie developers](https://www.reddit.com/r/androiddev/comments/b24i3d/hey_google_where_is_your_roadmap_why_commercial/)
## [4][Completely custom animations using DiffUtil](https://www.reddit.com/r/androiddev/comments/etptz0/completely_custom_animations_using_diffutil/)
- url: https://kiranrao.in/blog/2020/01/12/diff-util-part3/
---

## [5][How long does the approval of an app lasts?](https://www.reddit.com/r/androiddev/comments/etq5gs/how_long_does_the_approval_of_an_app_lasts/)
- url: https://www.reddit.com/r/androiddev/comments/etq5gs/how_long_does_the_approval_of_an_app_lasts/
---
Its been 20 hours now, would like to know whats the normal time for that. In german forums they say around 3-5 hours :/
## [6][Many apps still can pass Google Play policy...](https://www.reddit.com/r/androiddev/comments/etaccg/many_apps_still_can_pass_google_play_policy/)
- url: https://imgur.com/x8qz8Zg
---

## [7][State of Magisk: 2020 - topjohnwu](https://www.reddit.com/r/androiddev/comments/etkvje/state_of_magisk_2020_topjohnwu/)
- url: https://medium.com/@topjohnwu/state-of-magisk-2020-21de32721d65
---

## [8][Ad network for non Google Play app](https://www.reddit.com/r/androiddev/comments/etokvv/ad_network_for_non_google_play_app/)
- url: https://www.reddit.com/r/androiddev/comments/etokvv/ad_network_for_non_google_play_app/
---
Hey, my App was removed from Google Play, does anyone know an Ad network for which it’s not necessary to be published in Google Play?
## [9][Benefits of Writing as a Software Engineer](https://www.reddit.com/r/androiddev/comments/etmble/benefits_of_writing_as_a_software_engineer/)
- url: https://ayusch.com/every-software-engineer-should-write/
---

## [10][The Story of Our Big Android App Rewrite at Gojek](https://www.reddit.com/r/androiddev/comments/etcxbl/the_story_of_our_big_android_app_rewrite_at_gojek/)
- url: https://blog.gojekengineering.com/the-story-of-our-big-android-app-rewrite-6ede1cc3ad9a
---

## [11][Now in Android: January 23, 2020](https://www.reddit.com/r/androiddev/comments/ethdqr/now_in_android_january_23_2020/)
- url: https://medium.com/androiddevelopers/now-in-android-january-23-2020-587378171ae3
---

## [12][Anyone here using a collapsing toolbar with an image that draws under the status bar?](https://www.reddit.com/r/androiddev/comments/etpino/anyone_here_using_a_collapsing_toolbar_with_an/)
- url: https://www.reddit.com/r/androiddev/comments/etpino/anyone_here_using_a_collapsing_toolbar_with_an/
---
Does your status bar jump around like this?

[https://imgur.com/DMzi68O](https://imgur.com/DMzi68O)

&amp;#x200B;

The frame where it happens:  
[https://imgur.com/sSfYtW9](https://imgur.com/sSfYtW9)

&amp;#x200B;

Looks super ugly.

It starts happening when I set fitsSystemWindows on the AppBarLayout. But without that attribute, the AppBarLayout doesn't draw under the status bar.
