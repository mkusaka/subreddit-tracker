# androiddev
## [1][App Feedback Thread - July 25, 2020](https://www.reddit.com/r/androiddev/comments/hxlxir/app_feedback_thread_july_25_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hxlxir/app_feedback_thread_july_25_2020/
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
## [2][We’re on the engineering team for Android Studio. Ask us Anything! (starts July 30)](https://www.reddit.com/r/androiddev/comments/hwiffd/were_on_the_engineering_team_for_android_studio/)
- url: https://www.reddit.com/r/androiddev/comments/hwiffd/were_on_the_engineering_team_for_android_studio/
---
We’re on the engineering team for Android Studio, and we are excited to participate in another AMA on r/androiddev on **Thursday, July 30**!

For our [launch](https://android-developers.googleblog.com/2020/06/unwrapping-android-11-beta-plus-more.html) of the [Android 11 Beta](http://d.android.com/android11), we introduced [#11WeeksOfAndroid](https://developer.android.com/11weeksofandroid), and next week is Week 7, in which we’ll focus on Android Developer Tools. 

Also as part of the Android 11 Beta launch, we announced Android Studio 4.1 Beta and 4.2 Canary, and released an [overview video and demo](https://youtu.be/NMFGuy6TRqk) of the latest features in action. The preview versions of Android Studio have a host of [new features](https://developer.android.com/studio/preview/features) and include behind the scenes polish and stability work. 

For this AMA, we want to answer your *technical* questions about Android Studio.  Please note that we want to keep the conversation focused strictly on the engineering and usage of Android developer tools.

We'll start answering questions on **Thursday, July 30 at 12:00 PM PDT / 3:00 PM EDT (UTC 1900)** and will continue until 1:20 PM PDT / 4:20 PM EDT. 

**Feel free to submit your questions ahead of time.** This thread will be used for both questions and answers. Please adhere to our [community guidelines](https://developer.android.com/community-guidelines) when participating in this conversation.

**Here are some topics we’re looking forward to talking about, but feel free to ask anything!**

* Android Studio 
* Android Gradle Plugin &amp; Toolchain
* R8/D8 
* Apply Changes
* Android Emulator
* App Performance Tools
* UI Design Tools
* Constraint Layout &amp; Motion Layout
* Kotlin Tools
* C++ Tools
* Jetpack Compose Tools

**Participants from the Android team:**

* **Adarsh Fernando** (u/adarshf) - Android Studio Product Manager
* **Arif Sukoco**‎  - Android Studio C++ Tech Lead
* **David Herman** (u/adt\_dherman) - Data Binding Tech Lead
* **David Winer** \- Kotlin Product Manager
* **Diego Perez** (u/nosuid) - Layout Editor Tech Lead
* **Esteban de la Canal** (u/estebandlc) - Apply Changes Tech Lead
* **Frank Yang** (u/lfy\_google) - Android Emulator Tech Lead
* **Jamal Eason** (u/easonj) - Android Studio Product Manager
* **Jeffrey Van Gogh** \- Android Studio Tech Lead, &amp; Kotlin
* **Jerome Dochez** (u/jdochez) - Gradle Plugin Tech Lead
* **Jon Tsao**‎ (u/twilligy) - Android Studio Product Manager
* **Karen Ng** (u/nkaren) - Android Studio Product Director
* **Mads Ager** (u/madsager) - Android Studio Compiler Tech Lead
* **Nicolas Roard** (u/nicolasroard) - Design Tools &amp; ConstraintLayout Tech Lead
* **Raluca Sauciuc**‎ (u/raluca\_sauciuc) - Android Studio Tech Lead
* **Renaud Paquay‎** \- Android Studio Tech Lead
* **Scott Swarthout** (u/scottswarthout) - Android Studio Product Manager - Motion Tools
* **Shukang Zhou** (u/shukang) - Android Studio Profiler Tech Lead
* **Siva Velusamy** \- Android Studio Tech Lead
* **Tor Norbye** (u/tnorbye) - Engineering Director for Android Studio
* **Wojtek Kaliciński** (u/wkalicinski) - Android Developer Relations
* **Xavier Ducrohet** (u/droidxav) - Tech lead for Android SDK
## [3][How does your CI/CD workflow look like?](https://www.reddit.com/r/androiddev/comments/hy44kb/how_does_your_cicd_workflow_look_like/)
- url: https://www.reddit.com/r/androiddev/comments/hy44kb/how_does_your_cicd_workflow_look_like/
---
While this is the workflow we follow I am interested to know how does your workflow looks and any feedback on the above workflow.  


  
Whether you are working for a services company, a start-up, or an indie developer, we can share our workflow here.  


I'll start with mine:  


**develop branch:** Whenever a feature branch PR raised to be merged in currently active development branch, we run tests to check if all test cases are passing and code coverage to check via jacoco to ensure our test code coverage does not go below 70%.  
After the PR is merged build is generated with Bitrise and the latest build is sent to developers, testers, and PMs for testing and other stuff. (staging flavour)  


**pre-master branch:** We run a 2-week sprint and after every 2 weeks, we send a build to the client to see the work we have done till yet and give feedback.   
The code on develop branch is merged in the pre-master branch and the latest build is sent to the client, developers, testers, and PMs. (staging flavor)  


**master/main branch:**  The code in pre-master branch is merged into the master to push build live to Play store. (Production flavor)  


**note**: All the above-mentioned branches are protected i.e no one can push directly to them, it has to be done via PR.  


While this is the workflow we follow I am interested to know how does your workflow looks like and any feedback on the above workflow.
## [4][Is this even possible without bots? - this app is a 1:1 stickman Hook clone not published through by a published](https://www.reddit.com/r/androiddev/comments/hxsq49/is_this_even_possible_without_bots_this_app_is_a/)
- url: https://i.redd.it/mj6a1gcs22d51.jpg
---

## [5][Android Studio with AMD chipset](https://www.reddit.com/r/androiddev/comments/hy6si6/android_studio_with_amd_chipset/)
- url: https://www.reddit.com/r/androiddev/comments/hy6si6/android_studio_with_amd_chipset/
---
Has anyone run into any issues with utilising the inbuilt emulator on AS with any AMD A-series processors? I thought it may be a Windows-specific problem with Hyper-V but even after compiling and trying to test emulation on my Linux machine with an A9 CPU I'm still seeing major problems. 
Has anyone come across this before? Switch to Intel entirely?
## [6][What’s your favorite language for Android app development.](https://www.reddit.com/r/androiddev/comments/hy5xso/whats_your_favorite_language_for_android_app/)
- url: https://www.reddit.com/r/androiddev/comments/hy5xso/whats_your_favorite_language_for_android_app/
---
Am kind of stuck in choosing my next programming language. For those of you who’ve had the chance to explore the different language options to develop android apps. From your experience, which have you found to be the best for you and why?
## [7][Urgent advice needed. When is it better to use SQLite and when to use FireBase?](https://www.reddit.com/r/androiddev/comments/hy55oz/urgent_advice_needed_when_is_it_better_to_use/)
- url: https://www.reddit.com/r/androiddev/comments/hy55oz/urgent_advice_needed_when_is_it_better_to_use/
---
My only experience with databases so far is postgres but I need to make an app in 3 weeks that requires a database. The database would mostly handle text and I would need it to have String questions stored in it and I also would need the app to write to the database on user performance (to keep it simple, it will be a quiz app of sorts).

In case it makes any difference, I am using Java, not Kotlin.

\*This is general and should not be indicated as a "help me" post as indicated by the rules. I urgently need advice for this matter. If for whatever reason a mod sees it as unacceptable, please direct me to where I can ask this.
## [8][Making An MMORPG Platformer (Like Mario)](https://www.reddit.com/r/androiddev/comments/hy4d2v/making_an_mmorpg_platformer_like_mario/)
- url: https://www.reddit.com/r/androiddev/comments/hy4d2v/making_an_mmorpg_platformer_like_mario/
---
Hello!

I've made plenty of games before and have been burned out from my web-dev job. I started to mess around in Unity again, going back to my roots, and started falling in love with coding again. I'm hoping to start a big project that will keep me pre-occupied when I don't wanna hang out with people (my girlfriend because I don't have many friends).

TMI SORRY

I've seen a large amount of MMO released for android and iOS. UNITY is great BUT they are planning on depreciating their current MMO solution and I'm wondering what all these games in the android store are built on. Any ideas?

I'm wondering if it is GameMaker Studio , Unity, or UnRealEngine.

I want to be pre-occupied but I want it to go smoothly, so I actually have a finished product and not just another one of the many games I've started and abandoned (It makes me feel like that dad that went to go get cigarettes).

ADVICE APPRECIATED!! Also my game idea isn't anything special but if it's finished I think it will at least be fun and funny - til it isn't. :) It will be physics based (platformer 2d sidescroller) with projectiles.

&amp;#x200B;

Edit: In case no one answers, I'll post a bit of what I'm finding here for future search queries.

\- [https://improbable.io/blog/spatialos-mobile-pre-alpha](https://improbable.io/blog/spatialos-mobile-pre-alpha)

\-uMMORPG is not good for anyone making a platform game because it doesn't use gravity. It uses navmesh. It's great if you are doing zelda style games, based on what I'm seeing.

\- Because Unity is the best at 2D games, and I'm very familiar with it, going to make it using that. I'm going to use SpatialOS and have them host it for me. I'm not 100% sure how it is going to work but it looks very promising at first glance and it appears that many indie studios have used it. Here is a link to their documentation for anyone who is interested.

 [https://documentation.improbable.io/gdk-for-unity/docs/setup-and-dependencies](https://documentation.improbable.io/gdk-for-unity/docs/setup-and-dependencies) 

comments and advice appreciated still, even though I'm pretty sure this is the path I'm going to take. Now I need sleep. I still get crazy on weekends and don't sleep til, apparently about 7:33am
## [9][Video Walkie Talkie](https://www.reddit.com/r/androiddev/comments/hy18oq/video_walkie_talkie/)
- url: https://www.reddit.com/r/androiddev/comments/hy18oq/video_walkie_talkie/
---
How hard would it be to develop a video walkie talkie app, like Marco Polo or Glide?
## [10][COVID-19 tracing apps may fail to notify exposed users due to aggressive OEM battery saving measures](https://www.reddit.com/r/androiddev/comments/hxduj8/covid19_tracing_apps_may_fail_to_notify_exposed/)
- url: https://www.androidpolice.com/2020/07/24/covid-19-tracing-apps-may-fail-to-notify-exposed-users-due-to-aggressive-oem-battery-saving-measures/
---

## [11][Glad to share that my library 'TextWriter' crossed 100 stars ⭐ on Github. Thanks a lot to everyone 🙏](https://www.reddit.com/r/androiddev/comments/hy4fiu/glad_to_share_that_my_library_textwriter_crossed/)
- url: https://github.com/sarnavakonar/TextWriter
---

## [12][What is your experience in encountering large legacy android codebase?](https://www.reddit.com/r/androiddev/comments/hxiotr/what_is_your_experience_in_encountering_large/)
- url: https://www.reddit.com/r/androiddev/comments/hxiotr/what_is_your_experience_in_encountering_large/
---
Hi,  
The reason I wanted your opinion is that I just moved to a new team and their project is being developed for 4 years now. and this is the first time in my career that I'm experiencing this.

There are tons of classes, tons of base classes..., some parts don't have any architectures at all(imagine Activities with 600 lines of code, no DI, API calls in Activity and etc.) and some minor parts do have MVVM/CleanArch implemented.

\-How do you start reading these large codebases that have mixed architectures/no architecture in order to be able to understand and contribute to?
