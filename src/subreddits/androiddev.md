# androiddev
## [1][App Feedback Thread - August 08, 2020](https://www.reddit.com/r/androiddev/comments/i5yb1v/app_feedback_thread_august_08_2020/)
- url: https://www.reddit.com/r/androiddev/comments/i5yb1v/app_feedback_thread_august_08_2020/
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
## [3][Is it possible to make a reminder app that acts like an alarm app?](https://www.reddit.com/r/androiddev/comments/i5udj4/is_it_possible_to_make_a_reminder_app_that_acts/)
- url: https://www.reddit.com/r/androiddev/comments/i5udj4/is_it_possible_to_make_a_reminder_app_that_acts/
---
I was just watching this podcast where the hosts were complaining about not getting reminder notifications. Apparently they just blend in with all their other unimportant notifications and sometimes their phone does this battery saver thing where it just won't receive notifications until they unlock. They mentioned they'd resorted to setting alarms but this was impractical because you can't set an alarm more than 24 hours in advance.

I have third party alarm apps and they have a completely custom ui when they go off but they are 100% reliable in my experience. So I'm guessing android somehow knows that they're alarms and somehow wakes the phone when it's time. So my question is what api are they using and will it allow me to set an "alarm" for 30 days from now at 2:15pm?
## [4][[UPDATE] I am a self-taught developer who just finished working on an app for controlling Philips Hue via Samsung Galaxy Edge panel.](https://www.reddit.com/r/androiddev/comments/i5d93x/update_i_am_a_selftaught_developer_who_just/)
- url: https://www.reddit.com/r/androiddev/comments/i5d93x/update_i_am_a_selftaught_developer_who_just/
---
# OP: [https://www.reddit.com/r/androiddev/comments/hzjuu9/i\_am\_a\_selftaught\_developer\_who\_just\_finished/](https://www.reddit.com/r/androiddev/comments/hzjuu9/i_am_a_selftaught_developer_who_just_finished/)

# Galaxy Store (if you get an error, search for "Hue" on Galaxy Store app): [http://apps.samsung.com/appquery/appDetail.as?appId=com.nilstrubkin.hueedge](http://apps.samsung.com/appquery/appDetail.as?appId=com.nilstrubkin.hueedge)

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
## [5][Can I use a prepaid Visa/MasterCard card to register for the Dev Console?](https://www.reddit.com/r/androiddev/comments/i5xh23/can_i_use_a_prepaid_visamastercard_card_to/)
- url: https://www.reddit.com/r/androiddev/comments/i5xh23/can_i_use_a_prepaid_visamastercard_card_to/
---
I have a client who needs me to make an app and list it on the Google Play Store. I have a fresh Google account to register with, but I'll have to give this account to the customer after the work is done, so can I use a prepaid debit card to ensure my privacy?
## [6][Playing Netflix content in your own app](https://www.reddit.com/r/androiddev/comments/i5x8c0/playing_netflix_content_in_your_own_app/)
- url: https://www.reddit.com/r/androiddev/comments/i5x8c0/playing_netflix_content_in_your_own_app/
---
I've looked into the Rave Android app [https://play.google.com/store/apps/details?id=com.wemesh.android&amp;hl=en](https://play.google.com/store/apps/details?id=com.wemesh.android&amp;hl=en) which plays Netflix content in its app and does sync across multiple users. I was curious as to how they by-pass the DRM protection of the content.

However, this is already done for Kodi  [https://github.com/CastagnaIT/plugin.video.netflix](https://github.com/CastagnaIT/plugin.video.netflix)  
Rave does it for Android.

I've had some luck to decompile the Rave APK and constructed a sample app that fetches the Manifest file responsible for playback. [https://github.com/Veeresh8/InputStream-Netflix](https://github.com/Veeresh8/InputStream-Netflix) (sample app)

Has anyone done something like this before? I'm having problems playing the content getting an EOF exception but can't seem to figure out why.

[https://youtu.be/3Y3R\_snaDDc?t=1115](https://youtu.be/3Y3R_snaDDc?t=1115) this will give a complete picture of what I'm talking about.

Any help is appreciated, thanks!
## [7][square/anvil - A Kotlin compiler plugin to make dependency injection with Dagger 2 easier](https://www.reddit.com/r/androiddev/comments/i5i552/squareanvil_a_kotlin_compiler_plugin_to_make/)
- url: https://github.com/square/anvil
---

## [8][Custom Rating With kotlin](https://www.reddit.com/r/androiddev/comments/i5v1h8/custom_rating_with_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/i5v1h8/custom_rating_with_kotlin/
---
Source code link inside article

https://medium.com/@rithvik2994/custom-rating-system-android-3185b4f94596
## [9][How do I improve my Android skills and learn "Good practices"?](https://www.reddit.com/r/androiddev/comments/i577pd/how_do_i_improve_my_android_skills_and_learn_good/)
- url: https://www.reddit.com/r/androiddev/comments/i577pd/how_do_i_improve_my_android_skills_and_learn_good/
---
I've a good idea of app development and have made 3 apps yet. But whenever I get stuck I always use YouTube videos. I don't know if they are good enough.

I need to learn and make my code upto the accepted standards. Like how to write good code, what all things to use , how to improve security and all this. How are codes written by professionals at Google and Amazon. I don't have any idea how close random Indian YouTube videos are to their standards.

Can someone guide me on how I can improve my coding skills. Thank you
## [10][How to become a proficient Android Mobile Developer](https://www.reddit.com/r/androiddev/comments/i5jx3x/how_to_become_a_proficient_android_mobile/)
- url: https://www.reddit.com/r/androiddev/comments/i5jx3x/how_to_become_a_proficient_android_mobile/
---
I have recently started learning JAVA and Android studio, I do have so much more to learn, however I started getting curious on whether "learning by watching tutorials" was the best way to learn.  I personally feel like I do have a relative grasp of basic concepts, however, I do not feel like I remember any of the syntax. Would it be better to stop "learning" and just start making personal projects and later google questions I have relating some issues ? or should I follow what I am currently doing ? 

&amp;#x200B;

The reason why I am sticking to "tutorial hell" is that I feel like I could miss on some major concept, (like interfaces or threading) and not know that i missed out on it, making me not know that the possibility even exists. I don't know if that makes sense... If it doesnt please let me know i will edit and try to rephrase my last sentence. 

&amp;#x200B;

Anyways, thanks in advance for the help !
## [11][Using `ContentProvider` to receive context in your library.](https://www.reddit.com/r/androiddev/comments/i5a2cf/using_contentprovider_to_receive_context_in_your/)
- url: https://www.reddit.com/r/androiddev/comments/i5a2cf/using_contentprovider_to_receive_context_in_your/
---
Was curious about how `Picasso.get()` works so I took a look at the main classes it has. (in Github). I noticed that the `Context` is received using a `ContentProvider`. For a library, this is very cool so that you don't have to ask for the context on the clients app side. However, I want to know, is there any risk, problem, something you should care about with this approach? 
If it's safe, why are still a lot of libraries asking for a context?
## [12][Return Early Pattern](https://www.reddit.com/r/androiddev/comments/i5f1a1/return_early_pattern/)
- url: https://medium.com/@leonelmenaia/return-early-pattern-3d18a41bba8
---

