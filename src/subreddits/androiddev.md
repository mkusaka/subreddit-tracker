# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
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
## [3][The US government has filed antitrust charges against Google](https://www.reddit.com/r/androiddev/comments/jf6b8n/the_us_government_has_filed_antitrust_charges/)
- url: https://www.theverge.com/2020/10/20/21454192/google-monopoly-antitrust-case-lawsuit-filed-us-doj-department-of-justice
---

## [4][Check out my new pal :)](https://www.reddit.com/r/androiddev/comments/jf8sqy/check_out_my_new_pal/)
- url: https://i.redd.it/tscwa1jfoeu51.jpg
---

## [5][🔥 Quickly distribute your app with Firebase App Distribution using GitHub Actions + Fastlane 🚀 | by Shreyas Patil | Firebase Developers | Oct, 2020](https://www.reddit.com/r/androiddev/comments/jf9r6w/quickly_distribute_your_app_with_firebase_app/)
- url: https://medium.com/firebase-developers/quickly-distribute-app-with-firebase-app-distribution-using-github-actions-fastlane-c7d8eca18ee0
---

## [6][kotlin-android-extensions plugin will be deprecated by the next Kotlin release](https://www.reddit.com/r/androiddev/comments/jetvoa/kotlinandroidextensions_plugin_will_be_deprecated/)
- url: https://www.reddit.com/r/androiddev/comments/jetvoa/kotlinandroidextensions_plugin_will_be_deprecated/
---
As you can see from this ticket: [https://youtrack.jetbrains.com/issue/KT-42121](https://youtrack.jetbrains.com/issue/KT-42121)  


So you should stop using kotlinx.android.synthetic if you haven't already and switch to View Binding. The Parcelize feature will be moved into a separate plugin.
## [7][Android 11 and storage strategies](https://www.reddit.com/r/androiddev/comments/jfbdnc/android_11_and_storage_strategies/)
- url: https://www.reddit.com/r/androiddev/comments/jfbdnc/android_11_and_storage_strategies/
---
With Android 11 and it's compulsory restrictions on storage, I was trying to read the documents Google has released as guidance.  However the roadmap seems weak at best - even though some flexibility seems to have been shown - on use of fopen() in native code etc.

But these words are just an indication - there is no overtly obvious guarantee things will behave as the Google webpages seem to suggest.  That is, it depends on interpretation and essentially every developer will have to do their own testing of the hazy API changes to establish what the reality of the changes is.  And even that could change - roadmap has been fluid - which for something as basic as storage injects uncertainty.

For now, one path we are exploring is to use "maxSdkVersion" set to 29 in build.gradle, so that apps are not downloaded by Android 11 users.

And then use the requestLegacyExternalStorage flag set to true in AndroidManifest.xml to allow continuation of old behavior for Android 10 (API 29) at least, and set targetSdkVersion to 29 in build.gradle.

The reason for limiting use to below Android 11 is because Android 11 changes behavior dramatically.

Some of the transition suggests Google is giving is for developers to move the directory that the app saves audio recording etc to one of the "shared storage" places - which I assume means saving to Music.  This will supposedly allow those files to be readable by File Manager etc and such apps.

There is also an indication that requestLegacyExternalStorage set to true has some effect in Android 11 as well - allows for transitioning from old to new Music folder etc.

But I havent seen actual working code provided by Google which accomplishes all these tasks - if the same research/testing/behavior is going to be done by thousands of developers, it makes sense for Google to shoulder some of that effort and provide pre-tested methods to do this stuff reliably.  Otherwise there will be a lot repetition of effort.


Another thing I am not too comfortable in gauging is just how popular Android 11 will become.

I am half anticipating some resistance as some of the reduction in features and breaking down of compatibility between apps starts becoming known to users.  Previously they have taken Google's word for it that everything is better on Android 11 - but when they use it the reality will become apparent.

Anecdotally you do hear of users saying "I am never going to update to Android 11", but how prevalent is that - will it mean just addressing below Android 11 market could be a viable gameplan for developers.

Because we have little interest in spending the effort to update many of our more complex apps to comply with an android roadmap that is unclear whether it will change.


So are there any guesses how popular Android 11 is expected to get ?


In one of the Google medium posts they say use of SAF will not be restricted.  If the new iteration of SAF only limits by preventing writing to top folder - how is this preventing malware.


Wouldnt the whole exercise have been simpler if Google had just stopped the old APIs from working on top level folder.

Why all this complicated mess ?
## [8][Developer tips and guides: Common policy violations and how you can avoid them](https://www.reddit.com/r/androiddev/comments/jf8gya/developer_tips_and_guides_common_policy/)
- url: https://android-developers.googleblog.com/2020/10/developer-tips-and-guides-common-policy.html
---

## [9][At what point to I create a business?](https://www.reddit.com/r/androiddev/comments/jes1i3/at_what_point_to_i_create_a_business/)
- url: https://www.reddit.com/r/androiddev/comments/jes1i3/at_what_point_to_i_create_a_business/
---
Hello,

I'm (one person making an app as a hobby) currently working on an app that is going to use a tiered subscription. I was looking at some of the options for monetization in the Google play console and I was looking at setting up a payment profile. It was asking for a business name. Does this have to be a *real* business? Should I go start the process of getting an LLC? (Or some other business type? (Recommendations and reasons are appreciated) if it doesn't have to be a real business, is there any pros or cons to having an actual business vs not having an actual business? I'm not that new to app development, but all I've made before were simple free games for fun, I never really tried monetizing any of it, so any help/suggestions/advice is very appreciated!

Thank you
## [10][ServerSocket on android?](https://www.reddit.com/r/androiddev/comments/jfawq4/serversocket_on_android/)
- url: https://www.reddit.com/r/androiddev/comments/jfawq4/serversocket_on_android/
---
Can i host Server on android OS(I am doing it on Emui but it same API so nevermind ) ?

If so do i have to(and how to) port forwarded port on phone?

Or does it follow Ruter ?
## [11][AppCompat 1.2 Lint Checks - AppCompatResources or ContextCompat or ResourcesCompat](https://www.reddit.com/r/androiddev/comments/jfaodh/appcompat_12_lint_checks_appcompatresources_or/)
- url: https://saurabharora.dev/2020-10-20-appcompatresources-vs-contextcompat-resourcescompat/
---

## [12][Android Studio 4.2 Canary 14 available](https://www.reddit.com/r/androiddev/comments/jetqp9/android_studio_42_canary_14_available/)
- url: https://androidstudio.googleblog.com/2020/10/android-studio-42-canary-14-available.html
---

