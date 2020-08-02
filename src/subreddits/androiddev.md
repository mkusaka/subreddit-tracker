# androiddev
## [1][App Feedback Thread - August 01, 2020](https://www.reddit.com/r/androiddev/comments/i1rfl3/app_feedback_thread_august_01_2020/)
- url: https://www.reddit.com/r/androiddev/comments/i1rfl3/app_feedback_thread_august_01_2020/
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
## [2][How android OS Boots up - An in-depth look](https://www.reddit.com/r/androiddev/comments/i27q6d/how_android_os_boots_up_an_indepth_look/)
- url: https://ayusch.com/android-internals-the-android-os-boot-process/
---

## [3][AndroidBites | Snippets | Three Most useful but least known List functions | Union, Intersection, Subtraction in Kotlin](https://www.reddit.com/r/androiddev/comments/i265tr/androidbites_snippets_three_most_useful_but_least/)
- url: https://chetangupta.net/union-intersect-subs
---

## [4][Android 10: “tile memory limits exceeded, some content may not draw” Multiple WebViews inside my app](https://www.reddit.com/r/androiddev/comments/i29nsp/android_10_tile_memory_limits_exceeded_some/)
- url: https://www.reddit.com/r/androiddev/comments/i29nsp/android_10_tile_memory_limits_exceeded_some/
---
Since one or two weeks some ads inside my app wont show up and would just flicker! Its random when this happens, and it flickers continuously

I have banner ads and video ads in my app and thats the way i call them:

    holder.wv.setInitialScale(1); holder.wv.setLayerType(View.LAYER_TYPE_SOFTWARE, null); //i call that on banner ads, stackoverflow fix 
    holder.wv.getSettings().setJavaScriptEnabled(true); 
    holder.wv.setBackgroundColor(Color.argb(1, 0, 0, 0)); //stackoverflow fix 
    holder.wv.getSettings().setUseWideViewPort(true); 
    holder.wv.getSettings().setLoadWithOverviewMode(true); 
    holder.wv.setScrollBarStyle(WebView.SCROLLBARS_OUTSIDE_OVERLAY); 
    holder.wv.setScrollbarFadingEnabled(false); //loadurl etc
    

I already googled a lot and found out that the hardware acceleration might be problem so I am turning it off for the banner ads but for my video ads i cant because otherwise they wouldnt work

My ads are integrated in a Recyclerview that loads a CardView with the Webview inside. I have 4-5 on each site but there were no problems with the amount of ads, never. There must be kind of issue with the Webview or the Chrome Driver itself? Thats what came up when I searched for my problem

Does anyone encountered this problem?

E/chromium: \[ERROR:tile\_manager.cc(793)\] WARNING: tile memory limits exceeded, some content may not draw
## [5][SpotiFlyer Update v1.2](https://www.reddit.com/r/androiddev/comments/i2b6zd/spotiflyer_update_v12/)
- url: https://www.reddit.com/r/androiddev/comments/i2b6zd/spotiflyer_update_v12/
---
Provide Suggestions on how to improve my app code base and my Developer Skills.

I am just a Beginner trying to improve by making.

SpotiFlyer is an Android App(Written in Kotlin), which aims to work as:

Downloader For Spotify:Albums, Tracks and Playlists and Even Direct Youtube Videos.

Extension To mute whenever a Spotify Ad Plays.(TODO)

Save your Data ,by not listening your Fav Songs Online again &amp; again(Just Download Them!)

No ADS!

Works straight out of the box and does not require you to generate or mess with your API keys (already included).


Release Apk: https://github.com/Shabinder/SpotiFlyer/releases/download/v1.2/SpotiFlyer-v1.2.apk

Github Repo Link: https://github.com/Shabinder/SpotiFlyer/

Contributions are Welcome!

🤩No Limited Quota from this Update 
🎧
## [6][Any issues with Android Studio/native NDK development workflows if one updates to MacOS Catalina (which only allows 64-bit code to run)](https://www.reddit.com/r/androiddev/comments/i28t4b/any_issues_with_android_studionative_ndk/)
- url: https://www.reddit.com/r/androiddev/comments/i28t4b/any_issues_with_android_studionative_ndk/
---
MacOS Catalina is the first mac os version that is going to stop supporting 32-bit code - only 64-bit programs/libraries
will run.

My question is whether any of the android developer (Android Studio/NDK compilation) workflows are going to be affected by this ?

Are there any components there that are still 32-bit and may fail on MacOS Catalina?
## [7][Jetpack compose for desktop](https://www.reddit.com/r/androiddev/comments/i1t0i8/jetpack_compose_for_desktop/)
- url: https://zeromod.in/posts/jetpack-compose-desktop/
---

## [8][Don't use email magic link for auth...](https://www.reddit.com/r/androiddev/comments/i283p9/dont_use_email_magic_link_for_auth/)
- url: https://www.reddit.com/r/androiddev/comments/i283p9/dont_use_email_magic_link_for_auth/
---
When we considered user auth methods, we first looked at using only using "Sign in with Google/Facebook". However, App Store guidelines now require that any "app that use a social login service ... must also offer Sign in With Apple as an equivalent option" -- which means you won't be given the true email of most users (given [Hide My Email](https://support.apple.com/en-us/HT210425)) ... **so if you're developing X-Platform, simply using Google &amp; Facebook is no longer viable.**

Given we wanted to allow users to find other users on the app, we needed a unique user identifier per user, giving us 4 options:

1. **Email + Password** ... requires forgot password and users to prove email ownership
2. **Support all third-party (Apple, Google, Facebook)** ... and add a new unique identifier to allow friends to find each other, given email will not be sufficient
3. **Phone number magic link** ... we were relying on Expo (previously did not support phone-number auth), and we also felt emails would be a good/cheap tool for communication
4. **Email magic link**

Email magic link felt... perfect, but it was pretty much a disaster, and you should be pretty wary of doing it without implementing safeguards. I imagine email + login code is a better option.

Here are the various safeguards we tried with metrics and graphs: [https://snaphabit.app/blog/password-less-login/](https://snaphabit.app/blog/password-less-login/). Happy to answer any additional questions!
## [9][Rise widgets?](https://www.reddit.com/r/androiddev/comments/i2avio/rise_widgets/)
- url: https://www.reddit.com/r/androiddev/comments/i2avio/rise_widgets/
---
I love my widgets but man they can take up a lot of screen space.. I try to resize the widget by long pressing and scale down but it just goes back to normal size.. is there any way of doing this?
## [10][ViewSwapper: a new library for single activity architecture fans](https://www.reddit.com/r/androiddev/comments/i26z7v/viewswapper_a_new_library_for_single_activity/)
- url: https://www.reddit.com/r/androiddev/comments/i26z7v/viewswapper_a_new_library_for_single_activity/
---
Self plugging here for a recent library I made:

ViewSwapper is a simple view container (think like FrameLayout) that supports transitions (API level 19+) and can be used to swap out screens without having to bother with fragments: [https://github.com/grant-park/ViewSwapper](https://github.com/grant-park/ViewSwapper)

it's meant to be a convenient basis for UI in any MVU/MVI apps, but also great for anyone who needs some flexible and lightweight view replacement

would love to start some discussion and answer any questions regarding this or single activity
## [11][As senior Android dev, how hard/time consuming is it to learn Spring/Hibernate for BE development?](https://www.reddit.com/r/androiddev/comments/i1w45i/as_senior_android_dev_how_hardtime_consuming_is/)
- url: https://www.reddit.com/r/androiddev/comments/i1w45i/as_senior_android_dev_how_hardtime_consuming_is/
---
Just something that's been on my mind for some time. I'm sure I'm not the first one.

I've been doing Android development for nearly 5 years now, currently working as platform lead. I actually still like the platform (although it's not perfect) and write code regularly but was thinking about broadening my skill set.

Anybody here who's gone through this? And if you switched to a BE job, do you prefer it now?
