# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ibdg88/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ibdg88/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - August 17, 2020](https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/
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
## [3][New language features and more in Kotlin 1.4](https://www.reddit.com/r/androiddev/comments/ibxm5o/new_language_features_and_more_in_kotlin_14/)
- url: https://android-developers.googleblog.com/2020/08/new-language-features-and-more-in.html
---

## [4][Kotlin 1.4 released](https://www.reddit.com/r/androiddev/comments/ibixya/kotlin_14_released/)
- url: https://blog.jetbrains.com/kotlin/2020/08/kotlin-1-4-released-with-a-focus-on-quality-and-performance/
---

## [5][About MVI](https://www.reddit.com/r/androiddev/comments/ibxcq1/about_mvi/)
- url: https://www.reddit.com/r/androiddev/comments/ibxcq1/about_mvi/
---
I've been following MVI courses. I'm already very familiar with MVVM, but there is a one thing I can't wrap my head around.   
Most of MVI implementations follows a princible that you put every UI related data into ViewState and observe it in the fragment. When you want to trigger an event, you changed the StateEvent value on viewmodel. When StateEvent's value is changed, you trigger the actual logic.   
But why is there a single LiveData? What if I have 3 seperate things to show in Fragment? In MVVM I would have 3 seperate LiveData. But with MVI, I have only one. If one of the three data is changed, three of them observed again. Why is that?

  
For example: I have UI that has:  
1. User Info  
2. Message List  
3. Any other reactive value

  
If there is a slight change in the 3rd value, User Info will be binded again and Message list will be submited again. I know it's not changing so it's not a big of a deal. But if there is no change why do we even bind them again? Why do we encapsulate when we can seperate things?
## [6][Android apps now available for Windows 10](https://www.reddit.com/r/androiddev/comments/ibdrjf/android_apps_now_available_for_windows_10/)
- url: https://i.redd.it/6vk9fjyv9kh51.gif
---

## [7][How to connect android app to a remote linux server?](https://www.reddit.com/r/androiddev/comments/ibxv6i/how_to_connect_android_app_to_a_remote_linux/)
- url: https://www.reddit.com/r/androiddev/comments/ibxv6i/how_to_connect_android_app_to_a_remote_linux/
---
Hello,
I want to know, how could I make an app with Java which connects to a linux server(like aws ec2 instance or a server installed in my PC) which can give commands to the server(like SSH) upload a file, process it again download it, etc.
## [8][Is Material Design dying?](https://www.reddit.com/r/androiddev/comments/ibrn38/is_material_design_dying/)
- url: https://www.reddit.com/r/androiddev/comments/ibrn38/is_material_design_dying/
---
Looking around at UI design websites and those design videos you see on YouTube it seems this "flat-rounded" design is all the rage right now.
**** Edit
Examples of what I mean: https://imgur.com/a/vr1FyEX
## [9][Developers communities on Facebook or LinkedIn](https://www.reddit.com/r/androiddev/comments/ic0k5a/developers_communities_on_facebook_or_linkedin/)
- url: https://www.reddit.com/r/androiddev/comments/ic0k5a/developers_communities_on_facebook_or_linkedin/
---
Hi! Do you know any groups of developers on Facebook or LinkedIn, where participants are actively sharing opportunities and thoughts with each other?
## [10][How long will I wait for a response to my appeal to suspend my application ?](https://www.reddit.com/r/androiddev/comments/ibzy31/how_long_will_i_wait_for_a_response_to_my_appeal/)
- url: https://www.reddit.com/r/androiddev/comments/ibzy31/how_long_will_i_wait_for_a_response_to_my_appeal/
---
I've been waiting 20 days for reply from Google Play support. Does anyone know how much it will take ?
## [11][Let an app power on the device like Alarms does](https://www.reddit.com/r/androiddev/comments/ibz9g9/let_an_app_power_on_the_device_like_alarms_does/)
- url: https://www.reddit.com/r/androiddev/comments/ibz9g9/let_an_app_power_on_the_device_like_alarms_does/
---
Hi Android Devs, I was wondering if there is a way to make an app that powers the device on/off.

It should work like alarms, where you set one. And if you want to power off the phone there is a checkbox if you want, that the alarm is active while the phone is off. When the time comes it powers on and the alarm starts ringing shortly after.

&amp;#x200B;

Sorry when this is complicated written, english is not my first language
## [12][Using Hilt in production.](https://www.reddit.com/r/androiddev/comments/ibz68d/using_hilt_in_production/)
- url: https://www.reddit.com/r/androiddev/comments/ibz68d/using_hilt_in_production/
---
I'm trying to introduce dependency injection in my app.  I am planning to use dagger 2, but with Hilt coming soon, how should I go with it?

Also, [docs](https://developer.android.com/training/dependency-injection/dagger-android) suggest using hilt instead, I believe it's still in alpha. Would it be a good idea to use hilt in production?
