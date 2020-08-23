# androiddev
## [1][App Feedback Thread - August 22, 2020](https://www.reddit.com/r/androiddev/comments/iehl0b/app_feedback_thread_august_22_2020/)
- url: https://www.reddit.com/r/androiddev/comments/iehl0b/app_feedback_thread_august_22_2020/
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
## [2][We’re on the engineering team for Android Jetpack &amp; Jetpack Compose. Ask us Anything! (starts August 27)](https://www.reddit.com/r/androiddev/comments/idefss/were_on_the_engineering_team_for_android_jetpack/)
- url: https://www.reddit.com/r/androiddev/comments/idefss/were_on_the_engineering_team_for_android_jetpack/
---
We’re on the engineering team for Android Jetpack &amp; Jetpack Compose, and we are excited to participate in another AMA on r/androiddev on **Thursday, August 27**!

For our [launch](https://android-developers.googleblog.com/2020/06/unwrapping-android-11-beta-plus-more.html) of the [Android 11 Beta](http://d.android.com/android11), we introduced [\#11WeeksOfAndroid](https://developer.android.com/11weeksofandroid), focusing on a new topic every week. We’re excited to close out our [\#11WeeksOfAndroid](https://developer.android.com/11weeksofandroid) with a focus on UI, and on Thursday we’ll be hosting an AMA on the Android Jetpack and Jetpack Compose!

[Android Jetpack](https://developer.android.com/jetpack) is our suite of libraries to help developers follow best practices, reduce boilerplate code, and write code that works consistently across Android versions and devices. We launched several new libraries like Hilt for Dependency Injection, App Startup, and updates to Paging, Navigation, CameraX, and more.  Check out the Jetpack updates we made during #11WeeksOfAndroid [here](https://android-developers.googleblog.com/2020/07/11-weeks-of-android-jetpack.html).

[Jetpack Compose](https://developer.android.com/jetpack/compose) is Android’s modern toolkit for building native UI - declarative, composable, and more! We’re currently in developer preview and excited to show you more next week, when we will focus the week on UI! Check out our overview [video](https://www.youtube.com/watch?v=U5BwfqBpiWU) and try the [tutorial](https://developer.android.com/jetpack/compose/tutorial) to learn more.

For this AMA, we want to answer your *technical* questions about Android Jetpack and Compose.  No roadmaps :)

We'll start answering questions on **Thursday, August 27 at 12:00 PM PDT / 3:00 PM EDT (UTC 1900)** and will continue until 1:20 PM PDT / 4:20 PM EDT.

**Feel free to submit your questions ahead of time.** This thread will be used for both questions and answers. Please adhere to our [community guidelines](https://developer.android.com/community-guidelines) when participating in this conversation.

**Here are some topics we’re looking forward to talking about, but feel free to ask anything!**

* Jetpack Compose
* Hilt
* Navigation
* WorkManager
* Paging
* Permissions
* App Startup
* AppCompat
* CameraX
* ...[and more](https://developer.android.com/jetpack/androidx/explorer)!

&amp;#x200B;

**Participants from the Android team:**

* **Adam Powell** \- Tech Lead on Jetpack Compose
* **Alan Viverette** (/u/alanviverette) - Tech Lead on Android Jetpack
* **Alex Elias** (/u/alex_elias) - Tech Lead on Jetpack Compose
* **Amanda Alexander** \- Product Manager on Jetpack and Jetpack Compose
* **Anna-Chiara Bellini** \- Product Manager on Jetpack Compose
* **Chris Banes** (/u/chrisbanes) - Android Developer Relations
* **Chris Craik** \- Tech Lead on Paging, Benchmark
* **Clara Bayarri** (/u/clarabayarri) - Tech Lead on Jetpack Compose
* **Dany Santiago** (/u/danyaguacate) - Tech Lead on Hilt &amp; Room
* **Diana Wong** (/u/androiddiana) - Product Manager on Android Jetpack &amp; App Compatibility
* **Doris Liu** \- (/u/doris4lt) Engineer on Jetpack Compose Animation
* **George Mount** \- Tech Lead on Jetpack Compose core
* **Ian Lake** \- Tech Lead on Navigation, Fragments, Lifecycle
* **Jamal Eason** \- Senior Product Manager, Android Studio
* **Jim Sproch** \- Engineer on Jetpack Compose
* **Karen Ng** (/u/nkaren) - Director of Product, Jetpack and Compose
* **Leland Richardson** (/u/lrichardson) - Jetpack Compose Compiler &amp; Runtime
* **Nick Butcher** \- Android Developer Relations
* **Romain Guy** (/u/romainguy) - Manager of the Android Toolkit/Jetpack team
* **Scott Swarthout** \- Product Manager on Jetpack Compose Motion Tools
* **Sergey Vasilinetc** \- Tech Lead on Arch Components
* **Siyamed Sinir** (/u/siyamed) - Tech Lead on Android Toolkit &amp; Compose
* **Stephan Linzer** \- Test
* **Sumir Kataria** (/u/SumirKodes) - Tech Lead on Android Jetpack
* **Trevor McGuire** (/u/teamcguire) - Engineer on CameraX
* **Vinit Modi** \- Product Manager on CameraX &amp; Camera
* **Yigit Boyar** (/u/yboyar) - Tech Lead on Android Jetpack
## [3][Request Android permissions using annotations only](https://www.reddit.com/r/androiddev/comments/if0tka/request_android_permissions_using_annotations_only/)
- url: https://www.reddit.com/r/androiddev/comments/if0tka/request_android_permissions_using_annotations_only/
---
[https://github.com/LikeTheSalad/aaper](https://github.com/LikeTheSalad/aaper)
## [4][Business logic to Transform the incoming data in Repository or ViewModel](https://www.reddit.com/r/androiddev/comments/iezuix/business_logic_to_transform_the_incoming_data_in/)
- url: https://www.reddit.com/r/androiddev/comments/iezuix/business_logic_to_transform_the_incoming_data_in/
---
Hello everyone,
I was recently experimenting with MVVM and i have this query.. 

I have different data sources form where the data is being fetched

1. Api/Room database(cached api data) 
2. Firebase Firestore
3. Another table form Room database


Now the data that will be arriving from all the sources are in different format

Now my query is where should i put the business logic that converts them to a common format so that all the data can be displayed in a single recyclerview 

Should i put the logic in the repository or should i put the logic in ViewModel?
## [5][Just Released RxResult, Android Activity Result as RxJava Streams](https://www.reddit.com/r/androiddev/comments/iexkf1/just_released_rxresult_android_activity_result_as/)
- url: https://www.reddit.com/r/androiddev/comments/iexkf1/just_released_rxresult_android_activity_result_as/
---
For all RxJava users, I just released RxResult, which allows you to get Activity Result as RxJava streams.

Repo: [https://github.com/psuzn/RxResult](https://github.com/psuzn/RxResult)

This library allows the usage of RxJava with the new Activity Result APIs are made available on the [Activity 1.2.0-alpha02](https://developer.android.com/jetpack/androidx/releases/activity#1.2.0-alpha02), [Fragment 1.3.0-alpha02](https://developer.android.com/jetpack/androidx/releases/fragment#1.3.0-alpha02)

This is my second android library, there might be some mistakes. So any feedback is highly appreciated.
## [6][New ratings filters on google play store and fake review problem.](https://www.reddit.com/r/androiddev/comments/if0vht/new_ratings_filters_on_google_play_store_and_fake/)
- url: https://www.reddit.com/r/androiddev/comments/if0vht/new_ratings_filters_on_google_play_store_and_fake/
---
I just noticed something today. In google play store at the top of the search results near the search bar there were two tags "4.0+ and 4.5+". Which would make it easy for users to filter out apps that have less than 4 rating.

In theory this would provide users better choices and it would force developers to make better products, but as we all know the rating system on play store is far from perfect.

You can post several reviews from different profiles to improve your rating. And then there's the problem of fake reviews.

At a time where we are getting emails like these

"I will provide You many good reviews with 5 star rating and download app promotion your app in low price plz msg me my whatsapp no +\*\* \*\*\*\*\*\*\*\*\*\*  bcz I have above 5,000-10,000 users review makers I have. plz call me or 📷 share your whatsapp no. I will msg you.. .. !!"

And shady reviews Like these

" This app hacked all my data from my phone. Do not download this app."

"This program does not work. Do not Download this program "3\*Devil Emojis""

One of my app did not receive any review for 6 months. Then suddenly I started receiving 1 star reviews, and coincidently at the same time I started getting these emails promising me to provide fake positive reviews. This is just a theory but I think these fake review makers post negative reviews on apps in hope that the developer may get desperate for positive reviews and then they will slide in with a service that does just that.

It seems that these innocent rating tags on the top of search results can have a lot of negative effects. What are your thought on this.
## [7][Helpful glossary For new developers](https://www.reddit.com/r/androiddev/comments/iephhj/helpful_glossary_for_new_developers/)
- url: https://www.reddit.com/r/androiddev/comments/iephhj/helpful_glossary_for_new_developers/
---
A long time ago when i started learning developing apps for android, I started with well-known udacity course for beginners, and one of the best resources I discovered is this glossary:
 
[Android For all ](https://developers.google.com/android/for-all/vocab-words)

I just opened it today by coincidence, and remembered how it helped me in the beginning of journey as a programmer. And so i recommend it to developers who are new or people who wanna take programming as a hoppy.

What about you other developers? Is there anyone who find it helpful for beginners like me ?
## [8][Comparing Android Build Times by CPU](https://www.reddit.com/r/androiddev/comments/if2i04/comparing_android_build_times_by_cpu/)
- url: https://chrishorner.codes/post/cpu-build-comparison/
---

## [9][16 - 22 August Android Newsletter](https://www.reddit.com/r/androiddev/comments/if1mks/16_22_august_android_newsletter/)
- url: https://www.reddit.com/r/androiddev/comments/if1mks/16_22_august_android_newsletter/
---
Stay up to date with Android development, in this week's edition:  
🤓 Control of Android versions  
💡 Fragments improvements  
👮 Lint framework custom rules  
💅 Transitions with Material Motionand much more!

🔥Featuring [@ianhlake](https://twitter.com/ianhlake) [@rockandnull](https://twitter.com/rockandnull) [@the\_gabrielfv](https://twitter.com/the_gabrielfv) [@NikitBhandari](https://twitter.com/NikitBhandari) [@wkalic](https://twitter.com/wkalic)

Enjoy and have a nice week!  
[https://vladsonkin.com/android-newsletter-8/](https://vladsonkin.com/android-newsletter-8/)
## [10][Does anyone else notice a drop in ranking on the play store after updating the store listing (specifically the full description)?](https://www.reddit.com/r/androiddev/comments/if0cs7/does_anyone_else_notice_a_drop_in_ranking_on_the/)
- url: https://www.reddit.com/r/androiddev/comments/if0cs7/does_anyone_else_notice_a_drop_in_ranking_on_the/
---
First time I noticed a change even when typing the name of our app into the play store it was ranking ~15th (previously first). Could be coincidence but noticed it shortly after I fixed a typo in the description. Updated the description and a few other things yesterday and today it is down to 3rd. Normally I never see it leave 1st, though I am not checking often.

Has anyone seen this - is this something to be aware of when updating the store listing, I couldn't find any documentation around it?
## [11][I have not much programming background but due to quarantine, I thought of learning to create a running track app with android studio. why not make one instead of paying a monthly subscription for premium features. If that is the case, I am wondering why there are not so many customized apps](https://www.reddit.com/r/androiddev/comments/if072p/i_have_not_much_programming_background_but_due_to/)
- url: https://www.reddit.com/r/androiddev/comments/if072p/i_have_not_much_programming_background_but_due_to/
---
Anyone can make their app of their choice instead of buying a subscription and so many individuals made apps should've been floating around in websites if making an app is easier which includes all the premium feature running app provides. Why that's not the case?  Am I too amateur to know about a wall which is there which I can't go past because of various reasons other than the difficulty of making one or Is it possible to make one If I am really into? Can you give me a starting point If it's possible? I am committed and ready for a lot of hours into it. I know a little bit of C++ and python.
## [12][What is the reason why I can't get libvlc to draw images on an android.view.Surface object?](https://www.reddit.com/r/androiddev/comments/if1yxw/what_is_the_reason_why_i_cant_get_libvlc_to_draw/)
- url: https://www.reddit.com/r/androiddev/comments/if1yxw/what_is_the_reason_why_i_cant_get_libvlc_to_draw/
---
I have been trying to create my own TV Framework module in conjunction with libvlc, but I have encountered serious problems when using it at runtime, starting with the fact that, despite the fact that from the java code the surface that I  delivers the framework, libvlc is not able to use it.  I link you the question of Stackoverflow where I uploaded the code and the registry.

[https://stackoverflow.com/questions/63542220/problems-with-libvlc-android-and-tvplayer-from-android-tv-framework](https://stackoverflow.com/questions/63542220/problems-with-libvlc-android-and-tvplayer-from-android-tv-framework)
