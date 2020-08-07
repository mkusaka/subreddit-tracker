# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/i5d2yl/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/i5d2yl/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
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
## [3][How do I improve my Android skills and learn "Good practices"?](https://www.reddit.com/r/androiddev/comments/i577pd/how_do_i_improve_my_android_skills_and_learn_good/)
- url: https://www.reddit.com/r/androiddev/comments/i577pd/how_do_i_improve_my_android_skills_and_learn_good/
---
I've a good idea of app development and have made 3 apps yet. But whenever I get stuck I always use YouTube videos. I don't know if they are good enough.

I need to learn and make my code upto the accepted standards. Like how to write good code, what all things to use , how to improve security and all this. How are codes written by professionals at Google and Amazon. I don't have any idea how close random Indian YouTube videos are to their standards.

Can someone guide me on how I can improve my coding skills. Thank you
## [4][Using `ContentProvider` to receive context in your library.](https://www.reddit.com/r/androiddev/comments/i5a2cf/using_contentprovider_to_receive_context_in_your/)
- url: https://www.reddit.com/r/androiddev/comments/i5a2cf/using_contentprovider_to_receive_context_in_your/
---
Was curious about how `Picasso.get()` works so I took a look at the main classes it has. (in Github). I noticed that the `Context` is received using a `ContentProvider`. For a library, this is very cool so that you don't have to ask for the context on the clients app side. However, I want to know, is there any risk, problem, something you should care about with this approach? 
If it's safe, why are still a lot of libraries asking for a context?
## [5][Are there APIs that allow devs to replicate Google Translate's "brush over the words to translate" feature?](https://www.reddit.com/r/androiddev/comments/i5akyc/are_there_apis_that_allow_devs_to_replicate/)
- url: https://www.reddit.com/r/androiddev/comments/i5akyc/are_there_apis_that_allow_devs_to_replicate/
---
To those who aren't familiar with Google Translate: there's a feature where you take a picture of text (e.g. a road sign or a book), and you use your finger to "paint over" the words you want to translate.

Are there ways to replicate this in an app? Are there APIs for this?
## [6][[UPDATE] I am a self-taught developer who just finished working on an app for controlling Philips Hue via Samsung Galaxy Edge panel.](https://www.reddit.com/r/androiddev/comments/i5d93x/update_i_am_a_selftaught_developer_who_just/)
- url: https://www.reddit.com/r/androiddev/comments/i5d93x/update_i_am_a_selftaught_developer_who_just/
---
# OP: [https://www.reddit.com/r/androiddev/comments/hzjuu9/i\_am\_a\_selftaught\_developer\_who\_just\_finished/](https://www.reddit.com/r/androiddev/comments/hzjuu9/i_am_a_selftaught_developer_who_just_finished/)

# Galaxy Store: [http://apps.samsung.com/appquery/appDetail.as?appId=com.nilstrubkin.hueedge](http://apps.samsung.com/appquery/appDetail.as?appId=com.nilstrubkin.hueedge)

&amp;#x200B;

[Screenshots of the app](https://preview.redd.it/0582acvvukf51.png?width=1703&amp;format=png&amp;auto=webp&amp;s=3767a1bfeabe9534fab928fad9f404b20b8f8f44)

I would like to thank every person that has participated in the beta testing of my newly created app. The beta app has reached over 150 downloads just over a couple of days. I have received an impressive amount of feedback from the users. After some issues with the approval, I am now releasing it out of beta, as the Samsung Galaxy store has just approved my submission.

**Feedback**

I have received many positive and insightful comments and suggestions all of which I have read and noted down. I have tried to refine the problematic areas and eliminate the bugs. I’ve rewritten large chunks of code in order to implement the latest and greatest libraries (Moshi + OkHttp) and make it easier for anyone to understand my code.

**Bugs**

The most common problem was the text size being too big or too small on the buttons, some users also reported the text or symbols being uncentered and unaligned. Both issues were addressed.

Some users reported issues with the scroll list of the edit activity. Both the scroll and the associated Drag&amp;drop action were refined. It now has haptic feedback (with the ability to turn it off).

Some users reported issues with buttons not registering when pressed. I have improved the communication with the bridge which should fix that. I have also added haptic feedback (with the ability to turn it off).

**Features**

I have improved the sorting of the buttons in edit activity. They now are not just alphabetically sorted, but first and foremost by type of resource. All your lights, groups, and scenes are now easier to find.

Ability to switch symbols indicating on and off states to words “On” and “Off”

As you might have noticed when long-holding on a button, you enter the adjustment menu. The color and brightness options reflect the current color of the light selected. The rooms and zones now do the same (with some reservations for API limitations).

The adjustment menu as well as the refresh function are now described in the setup's final screen.

**Open-Source**

Even though the app is published to the Galaxy Store for the price of 2$ I decide to keep it Open-Source. If you want to create a similar app to control some other type of equipment, just do it. I will also gladly take any feedback on my code.

Source: [https://github.com/nils-trubkin/Hue-Edge](https://github.com/nils-trubkin/Hue-Edge)

**Conclusion**

This app turned out great and I feel confident enough to release it. I will still take any feedback and suggestions to further improve the app.
## [7][Android 10 EACCES (Permission denied) issue](https://www.reddit.com/r/androiddev/comments/i5d3zz/android_10_eacces_permission_denied_issue/)
- url: https://www.reddit.com/r/androiddev/comments/i5d3zz/android_10_eacces_permission_denied_issue/
---
I have updated my app to API level 29 so it broke my image picking mechanism. Eariler everything was working fine. Now when I searched over forums/stackoverflow I came to know that I have to migrate to scoped storage framework in order to make it work. I found 2 things so far 
1. opt out of scoped storage 
2. Read the stream and write into cache directory so that my app can access it.

My question is that is it mandatory to copy the file for using it. If yes I'm afriad to use this approach as it will consume lot of time because my app uses large videos which we upload to server.

Is it possible to use this framework to just read the URI of image/video without copying to cache directory ? Did anyone worked on it ?

for reference I followed these articles
https://developer.android.com/guide/topics/providers/document-provider
https://medium.com/@sriramaripirala/android-10-open-failed-eacces-permission-denied-da8b630a89df
https://issuetracker.google.com/issues/130494105
## [8][AndroidBites | Understand Accumulator Pattern Before You Regret.](https://www.reddit.com/r/androiddev/comments/i568al/androidbites_understand_accumulator_pattern/)
- url: https://chetangupta.net/reduce-fold
---

## [9][How would one demo an android app on a portfolio? Are there better options than screenshots?](https://www.reddit.com/r/androiddev/comments/i5bwrn/how_would_one_demo_an_android_app_on_a_portfolio/)
- url: https://www.reddit.com/r/androiddev/comments/i5bwrn/how_would_one_demo_an_android_app_on_a_portfolio/
---
Hi all, I'm curious how most people share their work.

Having static screenshots seems like it isn't very interactable (compared to deploying a front-end site and having a link in your portfolio to xyz.com), and asking people to download an app just to see what it does seems like too much to ask.

 Is there an online emulator that can be used to demo mobile apps?
## [10][Is there any free API that can fetch me all the series available on online steam platform such as Netflix, Amazon Prime?](https://www.reddit.com/r/androiddev/comments/i5dfwj/is_there_any_free_api_that_can_fetch_me_all_the/)
- url: https://www.reddit.com/r/androiddev/comments/i5dfwj/is_there_any_free_api_that_can_fetch_me_all_the/
---

## [11][How can we use existing styles in styles.xml in Jetpack Compose? Or even is that possible?](https://www.reddit.com/r/androiddev/comments/i57xmo/how_can_we_use_existing_styles_in_stylesxml_in/)
- url: https://www.reddit.com/r/androiddev/comments/i57xmo/how_can_we_use_existing_styles_in_stylesxml_in/
---

## [12][Flutter vs React Native: Best for Cross-Platform Development](https://www.reddit.com/r/androiddev/comments/i5d5zy/flutter_vs_react_native_best_for_crossplatform/)
- url: https://tekkiwebsolutions.com/blog/flutter-vs-react-native/
---

