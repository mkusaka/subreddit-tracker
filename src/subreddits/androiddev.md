# androiddev
## [1][App Feedback Thread - November 14, 2020](https://www.reddit.com/r/androiddev/comments/ju180o/app_feedback_thread_november_14_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ju180o/app_feedback_thread_november_14_2020/
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
## [2][Weekly Questions Thread - November 09, 2020](https://www.reddit.com/r/androiddev/comments/jqvlp4/weekly_questions_thread_november_09_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jqvlp4/weekly_questions_thread_november_09_2020/
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
## [3][Vigilante - iOS 14 alike privacy features in Android](https://www.reddit.com/r/androiddev/comments/jude8s/vigilante_ios_14_alike_privacy_features_in_android/)
- url: https://www.reddit.com/r/androiddev/comments/jude8s/vigilante_ios_14_alike_privacy_features_in_android/
---
Hello folks,

I've recently built an app that mimics the iOS 14's privacy feature that sends an alert as a dot on your screen when your microphone/camera is being used.

You can download it directly from [GitHub](https://github.com/CraZyLegenD/Vigilante/releases/tag/1.0.0)

As you can see, it's an open source project that showcases some of the latest Android development's goodies as you can see on the [GitHub](https://github.com/CraZyLegenD/Vigilante) page.

It's open source since it's security based app, your data is yours and there are no trackers or libraries that talk to Google, feel free to contribute to the app especially if you can do a translation to your language.

If you have any suggestions or you've found an bugs in the app feel free to open an issue in the repo or the comments of this post.

Once again a link to the repo [https://github.com/CraZyLegenD/Vigilante](https://github.com/CraZyLegenD/Vigilante)
## [4][Jintin/FancyLocationProvider Wrapper of FusedLocationProviderClient for Android to support modern usage like LiveData or Flow.](https://www.reddit.com/r/androiddev/comments/jujqay/jintinfancylocationprovider_wrapper_of/)
- url: https://github.com/Jintin/FancyLocationProvider
---

## [5][Escape your Full Text Search Queries](https://www.reddit.com/r/androiddev/comments/jujq52/escape_your_full_text_search_queries/)
- url: https://blog.haroldadmin.com/escape-fts-queries/
---

## [6][How to avoid passing objects between activities](https://www.reddit.com/r/androiddev/comments/juj6su/how_to_avoid_passing_objects_between_activities/)
- url: https://www.reddit.com/r/androiddev/comments/juj6su/how_to_avoid_passing_objects_between_activities/
---
I don't know if this counts as "help me" or as "general discussion about [...] design"

I suspect that I am using the wrong code design for my app. I am making a note taking app where each note is an object. In the main activity I create an empty `Notebook` object. In the `Notebook` is an empty mutableListOf of `Notes`s. Each `Note` has the context of the note. (text, date created, date updated ...). When the user clicks on the "Add a note" button it jumps to the CreateNote activity. 

I can think of two ways of adding that note. Either, 
 1. Serialize the `Notebook` object
 2. Send the `Notebook` by using 
```
// In main activity
intent.putExtra(NOTEBOOK, notebook)
// In CreateNote activity
val notebook = intent.getSerializableExtra(NOTEBOOK) as? Notebook
```
 3. Add a newly created note to the notebook
 4. Send back the updated notebook
 5. Update the notebook in main activity.

Alternatively, I can create a new `Note`, serialize it, and send it back. 


Is that the best coding design? If there was a way of storing the `Notebook` object as a shared object between activities I think that would be the best. I could then simply access that object (and potentially send an index if I want to edit a specific note). Is there a way to do that?

I am writing in Kotlin, if that matters.
## [7][8 — 14 November Android Newsletter](https://www.reddit.com/r/androiddev/comments/jujnmb/8_14_november_android_newsletter/)
- url: https://www.reddit.com/r/androiddev/comments/jujnmb/8_14_november_android_newsletter/
---
Stay up to date with Android development, in this week's edition:  
🤓 Android Coroutine Scopes   
🤩 Run Android Studio on any device   
🧪 Experiment with a new framework Simple-Stack   
💭 Master the Bubbles API  
and much more!

Read it here 👉 [https://vladsonkin.com/android-newsletter-20/](https://vladsonkin.com/android-newsletter-20/)  
What's your favorite one?

🔥Featuring [@joenrv](https://twitter.com/joenrv) [@FilipStanis](https://twitter.com/FilipStanis) [@Zhuinden](https://twitter.com/Zhuinden) [@filipebatista](https://twitter.com/filipebatista) [@balarka](https://twitter.com/balarka) [@SG5202](https://twitter.com/SG5202) [@rockandnull](https://twitter.com/rockandnull) [@iabhishek1041](https://twitter.com/iabhishek1041) [@gokhanalicciii](https://twitter.com/gokhanalicciii) and many other great authors!

💚 Subscribe and receive new editions directly to your email. Weekly, no spam, unsub anytime.  
Here is an example: [https://mailchi.mp/46c5d65d2315/android-newsletter-20](https://mailchi.mp/46c5d65d2315/android-newsletter-20)
## [8][Jetpack Compose or Flutter... which is the safest bet?](https://www.reddit.com/r/androiddev/comments/judj8o/jetpack_compose_or_flutter_which_is_the_safest_bet/)
- url: https://www.reddit.com/r/androiddev/comments/judj8o/jetpack_compose_or_flutter_which_is_the_safest_bet/
---
I know this sub will mostly pick native, but just assume you can only pick a cross-platform solution. Here's what I've gathered so far. Please correct me if I'm wrong.


Flutter pros:

* Easy to learn

* Used by Google themselves in Google Pay, Google Assistant, Google Ads, etc.
* Used by Ebay, Alibaba, Baidu, Square, Tencent
* Used by popular apps like Insight Timer and Reflectly
* Will work on Fuschia OS

Flutter cons: 

* Devs seem to prefer Kotlin over Dart

* If Google drops Flutter, your investment in learning Dart will be for nothing
* iOS animation bugs 

* Not many first party plug-ins, and some of the few that exist have issues

Jetpack Compose pros:

* Similar to SwiftUI, so easier to get into for new iOS devs and  current Android devs 

* Built by Google *and* Jetbrains, who have a track record of excellent products, and they created Kotlin

* devs who prefer native android won't have much of a problem with this 

Compose cons: 

* Younger project, currently in Alpha, so who knows if it'll catch up to Flutter by the time it's mature
* If Flutter continues to grow, it'll be harder to convince Flutter devs to switch 

* With so many companies (including Google) already invested in Flutter, what reason do they have to even consider Compose? 

* If Fuchsia turns out to be a success and not a toy project, Flutter will be ahead in this new OS. Doesn't mean Android will just disappear though 

What do you think? I'm still in university so I don't have to worry as much about the current cross-platform job market. I wanna use the cross-platform framework that will be more likely to stay relevant or at least alive by 2022. Not sure if I should dive into Flutter now or just work on other stuff while I wait for Compose to become production ready
## [9][Jetpack Compose: intro &amp; basic layouts](https://www.reddit.com/r/androiddev/comments/juc0n0/jetpack_compose_intro_basic_layouts/)
- url: https://www.rockandnull.com/jetpack-compose/
---

## [10][Enabling databinding issue](https://www.reddit.com/r/androiddev/comments/julgc8/enabling_databinding_issue/)
- url: https://www.reddit.com/r/androiddev/comments/julgc8/enabling_databinding_issue/
---
Good day everyone,

I have this issue on a kotlin dsl Gradle multi module project I have been practicing with.

It prompts me to upgrade Gradle plugin from 4.0.2 to 4.1.0.

After updating I noticed the buildFeatures has no databinding option.

So I have no idea how to enable databinding on it .
## [11][How to change the Floating label(Hint) background in material design?](https://www.reddit.com/r/androiddev/comments/ju2b2r/how_to_change_the_floating_labelhint_background/)
- url: https://i.redd.it/yit485ypo7z51.png
---

## [12][Graph (network) library for Android app](https://www.reddit.com/r/androiddev/comments/jufh33/graph_network_library_for_android_app/)
- url: https://www.reddit.com/r/androiddev/comments/jufh33/graph_network_library_for_android_app/
---
Graph as in nodes/ edges graph (i.e. graph theory)

I'm a newbie Android Java developer, and I'm looking for Graph library for Android (one that has some good contributers or still active in maintaining it). I've googled and used some but they are either not maintained (e.g. have low committs) or no longer used by the committee.

Any suggestion is welcome!
