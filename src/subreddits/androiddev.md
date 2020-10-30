# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/jkw00q/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jkw00q/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
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
## [3][How can I do custom notch on CardView like this?](https://www.reddit.com/r/androiddev/comments/jksyrl/how_can_i_do_custom_notch_on_cardview_like_this/)
- url: https://i.redd.it/uncfvma607w51.png
---

## [4][Leak detection: Android Studio vs LeakCanary](https://www.reddit.com/r/androiddev/comments/jkmcu5/leak_detection_android_studio_vs_leakcanary/)
- url: https://dev.to/pyricau/leak-detection-android-studio-vs-leakcanary-35j5
---

## [5][I made a library alternative to CustomTabs](https://www.reddit.com/r/androiddev/comments/jkr2so/i_made_a_library_alternative_to_customtabs/)
- url: https://www.reddit.com/r/androiddev/comments/jkr2so/i_made_a_library_alternative_to_customtabs/
---
The main reason to develop this was the fact that CustomTabs needs a service provider installed in the device to launch the tabs eg: Chrome, Brave, etc.

So I made my own version of it which uses `WebView` to load pages. Plus, it supports private browsing &amp; dark mode (for API &gt;= 29).

Any suggestions &amp; improvements are most welcome!

[Github](https://github.com/KaustubhPatange/CustomWebviewTabs)
## [6][ConstraintLayout 2.0.4](https://www.reddit.com/r/androiddev/comments/jkkziq/constraintlayout_204/)
- url: https://androidstudio.googleblog.com/2020/10/constraintlayout-204.html
---

## [7][Multiple Ad Banners for every Fragment vs Single Shared Ad Banner. Revenue difference? (AdMob or others)](https://www.reddit.com/r/androiddev/comments/jkw7ri/multiple_ad_banners_for_every_fragment_vs_single/)
- url: https://www.reddit.com/r/androiddev/comments/jkw7ri/multiple_ad_banners_for_every_fragment_vs_single/
---
So I was curious whether it makes a difference, in terms of revenue, to show a different Ad Banner on every Fragment or if it would get the same results as just showing the same one from the Activity.

I'm not sure if I explained myself properly:

OPTION 1:

Activity1-&gt;\[no ad banner\]

* Fragment1-&gt;AdBanner1
* Fragment2-&gt;AdBanner2
* Fragment3-&gt;AdBanner3

OPTION 2:

Activity1-&gt;AdBanner1

* Fragment1-&gt;(same AdBanner1)
* Fragment2-&gt;(same AdBanner1)
* Fragment3-&gt;(same AdBanner1)
## [8][Java 9+ support for Android projects is being worked on](https://www.reddit.com/r/androiddev/comments/jkbj97/java_9_support_for_android_projects_is_being/)
- url: https://issuetracker.google.com/issues/139013660#comment6
---

## [9][Scrollytelling on native Android?](https://www.reddit.com/r/androiddev/comments/jkve8a/scrollytelling_on_native_android/)
- url: https://www.reddit.com/r/androiddev/comments/jkve8a/scrollytelling_on_native_android/
---
There are quite a lot of websites out there that use scrolling as an element of storytelling (scrubbing through a video, pinch into images, etc.). A good example for that is this website: AirPods Pro (sorry for the Apple reference).

For a project we're trying to accomplish something similar, but native. I haven't come across many Android related references but came up with a few options:

**Build the website and wrap it in a WebView.**  
Pro: We could use it for several platforms  
Contra: I think the UX will suffer quite a bit from that. Also, we need a click interaction to trigger sound.

**Use a list (like RecyclerView) with video players and some custom views**  
Pro: Native experience  
Contra: Not sure if this is feasible. f.ex. linking the video progress to scroll is either very hard or not possible after all.

Does anyone of you already accomplished something similar or has some tips how to create such an experience?
## [10][Google Elevation APIs](https://www.reddit.com/r/androiddev/comments/jkt7iv/google_elevation_apis/)
- url: https://www.reddit.com/r/androiddev/comments/jkt7iv/google_elevation_apis/
---
Hello! We will have our project at Microwave Design, which one of the crucial data we need to get is the ELEVATION.

Fortunately, Google offers such API: [https://developers.google.com/maps/documentation/elevation/overview.](https://developers.google.com/maps/documentation/elevation/overview)

But I don't really get the...

&amp;#x200B;

"**Pay only for what you use.**

Our pricing scales to fit your needs with no commitments, termination fees, or usage limits — and you can use mobile Maps at no charge."

&amp;#x200B;

Does it mean that I will get charged ONLY for every request? or does google will charge me EVERY month?  So it means that my card would not be credited unless the bulk number of requests on API happens?

Since I will be only using this API for around 100 to 500 requests. Thank you!

PS. Im using a python script to extract data.
## [11][I need ideas for features for a mobile app that teaches programming in Minecraft.](https://www.reddit.com/r/androiddev/comments/jkvgnk/i_need_ideas_for_features_for_a_mobile_app_that/)
- url: https://www.reddit.com/r/androiddev/comments/jkvgnk/i_need_ideas_for_features_for_a_mobile_app_that/
---
What kind of features and functionalities should I integrate into an app that is teaching programming in Minecraft? So far I have quizzes and explanatory videos and I need a lot more.
## [12][AR Core Android Tutorial without Sceneform?](https://www.reddit.com/r/androiddev/comments/jkp2cs/ar_core_android_tutorial_without_sceneform/)
- url: https://www.reddit.com/r/androiddev/comments/jkp2cs/ar_core_android_tutorial_without_sceneform/
---
Hi everyone! 

I am new to ARCore and it seems that most tutorials I find are related to Sceneform. However, sceneform doesn't seem to be supported in the newer versions of Android Studio ( which can only mean that it is being dropped eventually ). So it would really help me if there were some tutorials on ARCore android development without Sceneform.
