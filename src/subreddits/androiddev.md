# androiddev
## [1][App Feedback Thread - June 13, 2020](https://www.reddit.com/r/androiddev/comments/h86msj/app_feedback_thread_june_13_2020/)
- url: https://www.reddit.com/r/androiddev/comments/h86msj/app_feedback_thread_june_13_2020/
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
## [3][Google Play Store will make app bundles a requirement in 2021](https://www.reddit.com/r/androiddev/comments/h8cv3v/google_play_store_will_make_app_bundles_a/)
- url: https://www.androidpolice.com/2020/06/12/google-play-store-will-make-app-bundles-a-requirement-in-2021/
---

## [4][Google Play Console Beta is beautiful &amp; also packs great new features](https://www.reddit.com/r/androiddev/comments/h8opqc/google_play_console_beta_is_beautiful_also_packs/)
- url: https://www.androidpolice.com/2020/06/13/hands-on-with-the-new-google-play-console-easier-publishing-and-a-new-material-theme/
---

## [5][Some clarification about the terminlogy of TESTING](https://www.reddit.com/r/androiddev/comments/h8subw/some_clarification_about_the_terminlogy_of_testing/)
- url: https://www.reddit.com/r/androiddev/comments/h8subw/some_clarification_about_the_terminlogy_of_testing/
---
Hello,

I am getting into testing and I find the terminology a bit confusing and at times mixed up.

Can someone tell me if I get this right:

The `test` source set in Android Studio contains **local** tests. All test in this folder run without an emulator or real device (but not necessarily without Android dependencies, since they can be mocked).

The `androidTest` directory contains all tests that need an **emulator or real device** to run. They are called **instrumented** tests since they use the instrumentation API.

We also divide tests into **3 categories** by their **size: Unit** tests ( smallest), **Integration** tests (medium-sized), and **end-to-end** tests (largest).

Usually, **unit** tests are **local** tests (`test` package) but there are also instrumented unit tests that need an Android device to run (an so go into the `androidTest` package). 

**Integration** tests usually go into the `androidTest` package but there are also integration tests that don't need the Android framework (or where the Android parts can be mocked?) to run and can, therefore, go into the `test` (local) package.

End-to-end tests go into the `androidTest` package and might need external tools (? unsure). Do end-to-end tests need **manual** interaction or can they be fully automated?

Where do UI tests fit here? Are all UI tests end-to-end tests? Sorry for my ignorance.
## [6][Second app for Firebase Admin.](https://www.reddit.com/r/androiddev/comments/h8s7at/second_app_for_firebase_admin/)
- url: https://www.reddit.com/r/androiddev/comments/h8s7at/second_app_for_firebase_admin/
---
Hello there. Hope all are good.

I have an app that list promotions for retailers in a certain field of the economy. So products constantly have to be added and removed from promotions from 9 different retailers. I do that, not the retailer (at this stage).

I'm thinking of building a total separate new app to maintain the database of that app. Have any of you people done it? My thinking is if I put a link to an activity (with a login obviously) where I do that, it will almost double the size.

At this stage I do it with a website I created, but it is very clumsy, because my JavaScript and HTML is not up to par. I can build a much better way to do that with an app, but I'm not sure if I should add it to app or build  a separate app for that. Doing it directly on Firebase is totally out of the question. 

Just a little input please, what would you do?
## [7][Possible to implement a scrollable container with cumulative percentage height of child views that exceeds 100%?](https://www.reddit.com/r/androiddev/comments/h8r3fk/possible_to_implement_a_scrollable_container_with/)
- url: https://www.reddit.com/r/androiddev/comments/h8r3fk/possible_to_implement_a_scrollable_container_with/
---
I tried to do this with ConstraintLayout's percent height. The cumulative is 120%. But it can't be scrolled. match_parent obviously makes the layout unscrollable in the ScrollView. wrap_content seems to not work with scrolling, even though the child percentages still add up to 120%.
## [8][Google Play Store app details: "varies with device"](https://www.reddit.com/r/androiddev/comments/h8r2a9/google_play_store_app_details_varies_with_device/)
- url: https://www.reddit.com/r/androiddev/comments/h8r2a9/google_play_store_app_details_varies_with_device/
---
Something that's puzzled me for a while. When looking at apps in the Play Store, I often see: Current version: varies with device / Size: varies with device / Requires Android: varies with device. If that's accurate, then that's fine I guess.

But for example I see the same info (or lack thereof) for one of my own apps, uploaded as an apk. AFAIK, the version number is unique, the size is the size (all assets are loaded irrespective of the device), the minimum Android version is specified and unique.

Does anyone know why this information isn't being shown?
## [9][As promised, my insights on the effect of a 1 week ad on reddit for my app. Is it worth it? In short: Nope, not for me.](https://www.reddit.com/r/androiddev/comments/h8369y/as_promised_my_insights_on_the_effect_of_a_1_week/)
- url: https://i.redd.it/7ya29writm451.jpg
---

## [10][Screen pinning programmatically](https://www.reddit.com/r/androiddev/comments/h8pqf0/screen_pinning_programmatically/)
- url: https://www.reddit.com/r/androiddev/comments/h8pqf0/screen_pinning_programmatically/
---
How Touch Lock Lite achieve screen pinning of other apps? The official Android documentation says only the current activity, from which the pinning is triggered, can be pinned. Any hint?
## [11][Android emulator behaving weird after AS updated to 4.0](https://www.reddit.com/r/androiddev/comments/h8os2p/android_emulator_behaving_weird_after_as_updated/)
- url: https://www.reddit.com/r/androiddev/comments/h8os2p/android_emulator_behaving_weird_after_as_updated/
---
I have been having problems with the Android emulator lately after upgrading to the 4.0 version. The emulator doesn't start, I have tried every method i.e. cold boot, wiping data, even deleting the AS with SDK, and installing it afresh but everything went in vain. A few days ago it just started to work normally without me doing anything and today it again stopped working. Every time I start it this warning appears "Resetting for cold boot: Emulation engine failed". Any suggestions on fixing it?

P.S I am on a Windows machine
## [12][GitHub Actions used to build an Android project (with working UI tests - on emulators)](https://www.reddit.com/r/androiddev/comments/h86yju/github_actions_used_to_build_an_android_project/)
- url: https://www.reddit.com/r/androiddev/comments/h86yju/github_actions_used_to_build_an_android_project/
---
I have converted my android-showcase project to use. Now I am running UI tests on Android emulators and I have an infinite number of parallel executors.

Here is the full workflow config for GitHub Actions
https://github.com/igorwojda/android-showcase/tree/master/.github/workflows
