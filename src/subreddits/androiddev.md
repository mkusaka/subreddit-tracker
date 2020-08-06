# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/i2wdrd/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/i2wdrd/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - August 03, 2020](https://www.reddit.com/r/androiddev/comments/i2uyzu/weekly_questions_thread_august_03_2020/)
- url: https://www.reddit.com/r/androiddev/comments/i2uyzu/weekly_questions_thread_august_03_2020/
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
## [3][I started my channel a week ago. I convert dribbble shots to android apps. I just posted another video. I love to explore and come up with creative ways to make learning and coding fun and simple. Do watch it and provide suggestions. Github repo is in description feel free to use it. Thanks ✌🏻](https://www.reddit.com/r/androiddev/comments/i4nurf/i_started_my_channel_a_week_ago_i_convert/)
- url: https://youtu.be/Apqeg-OUheo
---

## [4][Android Vitals - Is this a cold start?](https://www.reddit.com/r/androiddev/comments/i4kluc/android_vitals_is_this_a_cold_start/)
- url: https://dev.to/pyricau/android-vitals-is-this-a-cold-start-3m44
---

## [5][Why are people trying to build Android apps with Python?](https://www.reddit.com/r/androiddev/comments/i4myou/why_are_people_trying_to_build_android_apps_with/)
- url: https://www.reddit.com/r/androiddev/comments/i4myou/why_are_people_trying_to_build_android_apps_with/
---
On my blog, I previously made a post about the Chaquopy plugin which lets you run Python scripts in an Android app.

This has been one of my most popular posts since it was posted, but I'm just befuddled, why?

I don't see the benefits of using Python over Java / Kotlin (or C++ if you're into that), especially considering the fact that you'll run with compatibility problems when you try to import other Python scripts and libraries.

I'm aware of other Python for Android libraries like PyPi and Kivy, and I don't know how well these can do to make solid Android apps.

Is this a case of Python Developers trying to make use of their existing skills to make Android apps? Or is there something I'm missing?
## [6][AltBeacon - need help with ranging beacons when app in foreground, and also monitoring them when it goes to background.](https://www.reddit.com/r/androiddev/comments/i4qlil/altbeacon_need_help_with_ranging_beacons_when_app/)
- url: https://www.reddit.com/r/androiddev/comments/i4qlil/altbeacon_need_help_with_ranging_beacons_when_app/
---
Hi, I'm asking here hoping that someone has had a similar use case as me and implemented it using [AltBeacon](https://altbeacon.github.io/android-beacon-library/).

Firstly, let me explain what exactly I want to achieve.

**When app is opened** on main view, I want to run beacon detection in the foreground, so user can actively see which beacons are detected. I want to **range beacons** here, so I know the proximity as well.

**When app is closed** or in background, I want to run beacon detection again, but I only want to monitor beacons here, so I don't care how far away the user is, just if user enters or exits the beacon range.

**Now here's the problem:**

AltBeacon clearly states that in order to detect beacons in the background, you need to create an instance of \`RegionBootstrap\` that lives through the entire App lifecycle. Sample code is found [here](https://altbeacon.github.io/android-beacon-library/samples.html) in "Background" section. 

If \`RegionBootstrap\` is implemented correctly, the app should relaunch the detection service after it's been killed, so it works in the background as well.

My main issue currently is that \`RegionBootstrap\` seems to "override" the beacon detection that's implemented inside the view, using \`BeaconManager.startRangingBeaconsInRegion(region)\`. So the beacon detection works either in the foreground or background, but not both at the same time.

Has anyone implemented something similar yet and how did you solve it? How can I implement beacon ranging when app is opened using \`BeaconManager\`, but also implement background monitoring with \`RegionBootstrap\`?
## [7][Master switch preference](https://www.reddit.com/r/androiddev/comments/i4f3d4/master_switch_preference/)
- url: https://www.reddit.com/r/androiddev/comments/i4f3d4/master_switch_preference/
---
There didn't seem to be any implementations of the master on/off and individual on/off preference switches as recommended by the Material Design guidelines and making them was pretty time consuming, so I'm releasing it as an open source library.

[https://github.com/svenoaks/MasterSwitchPreference](https://github.com/svenoaks/MasterSwitchPreference)
## [8][Leverage the In-App Review API for your Google Play reviews](https://www.reddit.com/r/androiddev/comments/i450fh/leverage_the_inapp_review_api_for_your_google/)
- url: https://android-developers.googleblog.com/2020/08/in-app-review-api.html
---

## [9][Organizing resource layout files in Android Studio .](https://www.reddit.com/r/androiddev/comments/i4rg9t/organizing_resource_layout_files_in_android_studio/)
- url: https://www.reddit.com/r/androiddev/comments/i4rg9t/organizing_resource_layout_files_in_android_studio/
---
we can organize the Java or Kaitlin files creating separate folder packages but how to manage layout resource files , us there any way to create a folder to separate some files for better readability and organizing .
Thank you for reading any for t highly appreciable .
## [10][Material Components Release 1.2.0](https://www.reddit.com/r/androiddev/comments/i3zgys/material_components_release_120/)
- url: https://github.com/material-components/material-components-android/releases/tag/1.2.0
---

## [11][How is Android 11 going to impact your mobile application?](https://www.reddit.com/r/androiddev/comments/i4muz9/how_is_android_11_going_to_impact_your_mobile/)
- url: https://www.metizsoft.com/blog/android-11-features
---

## [12][ExpandableFab: a customizable, orientation aware Android widget with animations](https://www.reddit.com/r/androiddev/comments/i4a0gk/expandablefab_a_customizable_orientation_aware/)
- url: https://www.reddit.com/r/androiddev/comments/i4a0gk/expandablefab_a_customizable_orientation_aware/
---
Just an easy to use library that allows you to beautifully show users the available actions they can take on the current screen.

Orientation Aware meaning you can show different options between portrait and landscape, easily.

Nearly everything is customizable (colors, text, size, placement, margins, animations and more)  and optional (don't need an Overlay, or FabOptions, or Labels, or icons, etc). Every property can be accessed or set through XML layouts or programmatically - whatever you prefer.

Written 100% in Kotlin but comes with full JavaDoc and KDoc (published API is well documented). Also comes with an example app so you can see different use cases with 0 coding.

Let me know what you think.

Github:  [https://github.com/nambicompany/expandable-fab](https://github.com/nambicompany/expandable-fab)

Library website (w/ links to full documentation):  [https://nambicompany.github.io/expandable-fab/](https://nambicompany.github.io/expandable-fab/)

(the smoothness and color of the gifs are a little off. See repo's README for more gifs of better quality, or better yet, just run the example app)

https://i.redd.it/fdj3agif7cf51.gif

https://i.redd.it/5lrz6lkf8cf51.gif
