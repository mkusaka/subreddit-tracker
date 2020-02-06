# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 03, 2020](https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/
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
## [3][Whats the point of "Developer Support" when you can't tell me what i need to fix to deploy an update? Directing me to the Developer Program Policies does NOT help - especially when the app has been live with no major changes.](https://www.reddit.com/r/androiddev/comments/ezk7lu/whats_the_point_of_developer_support_when_you/)
- url: https://i.redd.it/91106ggxe7f41.png
---

## [4][AndroidX security](https://www.reddit.com/r/androiddev/comments/ezs8ro/androidx_security/)
- url: https://www.netguru.com/codestories/androidx-security
---

## [5][Android Dark Theme](https://www.reddit.com/r/androiddev/comments/ezs5bk/android_dark_theme/)
- url: https://www.netguru.com/codestories/android-dark-theme
---

## [6][Motion Pattern (Gesture) Recognition AI: MiVRy](https://www.reddit.com/r/androiddev/comments/eznvek/motion_pattern_gesture_recognition_ai_mivry/)
- url: https://v.redd.it/ar8wr1tst8f41
---

## [7][How can I fix this type of crashes which are apparently not exactly related to my code?](https://www.reddit.com/r/androiddev/comments/ezqd15/how_can_i_fix_this_type_of_crashes_which_are/)
- url: https://i.redd.it/68nr826h3af41.png
---

## [8][Live Data + Sealed Classes are an easy way to lose a reference to data. How can I go around this?](https://www.reddit.com/r/androiddev/comments/ezqbgb/live_data_sealed_classes_are_an_easy_way_to_lose/)
- url: https://www.reddit.com/r/androiddev/comments/ezqbgb/live_data_sealed_classes_are_an_easy_way_to_lose/
---
It's becoming very popular to use sealed classes as a way to handle errors when observing LiveData. One common example would be to have a sealed class named Result with two Data Classes: Success and Failure.

Success holds a reference to our data and Failure holds a reference to Exception.

If we have a very simple data flow where we request some data *once* and wait for it to be read, this works very well. Our observable will either give us the Success state, after which we retrieve our value and draw UI - or it'll give us a Failure state, after which we handle the error gracefully and deem that we were unable to get the data for whatever reason.

Now here are two relatively common cases from the top of my head when things can go wrong and I'm not sure how to handle them:

1. We read some data and pass it to UI. User decides to refresh the view, but this time the network request fails for whatever reason. Now our LiveData&lt;Result&gt; turns from Success (which holds a reference to our previous data) to Failure (which only holds a reference to our Exception). We can decide to ignore the error in this case, since our UI is already drawn and we can tell the user that the refresh failed some other way (e.g. Toast) but we don't have a reference to our data any more.
2. Our Repository read the data from both a local cache and a network request. We don't have a single point of truth implemented (we don't read all data from a database because we only cache data rarely) so the way it works is we read data from SharedPreferences and pass it to LiveData as a Result.Success. Observable is fired and UI is drawn. Now because the network request failed, just like last time, this Result.Success turns into Result.Failure again. The UI is already drawn, we don't really care that the request failed because we'd rather display outdated data than an error message. But again we lost the reference to our data because Result.Success was overwritten by Result.Failure.

How should this be handled? I don't see how Sealed Classes are so popular and they seemed to have completely replaced callbacks when it's so easy to lose a reference to your data. Error states will simply overwrite any other states.
## [9][Which course to pick in 2020 for learning android development?](https://www.reddit.com/r/androiddev/comments/ezqhe0/which_course_to_pick_in_2020_for_learning_android/)
- url: https://www.reddit.com/r/androiddev/comments/ezqhe0/which_course_to_pick_in_2020_for_learning_android/
---
I am totally a beginner and I know some java basics. While searching yesterday, I find out that people suggest Udemy course: The Complete Android Oreo Developer Course - Build 23 Apps!   
OR   
Android basics on Udacity. On the other hand, people say these courses are outdated and contain some topics or functions which are legacy.  


I am confused, I don't know what to do. I have to work on the university project and I still have 6 months. I will be spending 6 hrs each day for learning, but I really need to know that even if these courses are 4 years old, are they worth to start now in 2020? I hope I am not wasting my time on learning outdated courses since time is really short for the project.
## [10][How to become a better android programmer?](https://www.reddit.com/r/androiddev/comments/ezb4y6/how_to_become_a_better_android_programmer/)
- url: https://www.reddit.com/r/androiddev/comments/ezb4y6/how_to_become_a_better_android_programmer/
---
Hi all,

I'm a junior android developer and I want to improve. I would like to know, which in your opinion are the best libraries,frameworks,design patterns, etc... to focus on.

For example I've read about Dagger and Retrofit (I'm using Volley) and about MVVM, even RxAndroid seems cool. I want to start to implement unit tests and I'm also learning Kotlin. 

There are a lot of things, but which are the things that are worth to learn for real?
## [11][Easily create beautiful app screenshots for the Play Store](https://www.reddit.com/r/androiddev/comments/ezf8hz/easily_create_beautiful_app_screenshots_for_the/)
- url: https://www.reddit.com/r/androiddev/comments/ezf8hz/easily_create_beautiful_app_screenshots_for_the/
---
Hey folks! Have you ever dreaded uploading screenshots when listing your app on the Play Store because they were annoying to make? I've had the same experience, so I decided to build a super easy-to-use tool to streamline making beautiful looking screenshots.

It has support for the latest phones, like the Pixel 4 and Pixel 4 XL, Samsung Galaxy S10 and S10+, as well as smaller phone sizes like the Nexus 6P and 5X. You can customize a lot of aspects of the screenshot, such as font family, font size, background color, background images, etc. You can also save your screenshots directly in the tool, access them on any device, and edit them easily at any point in the future!

It has a generous free tier and I'm continually adding new features based on users' feedback. Check it out at [https://screenshotcreator.com](https://screenshotcreator.com/), and let me know what you think!

Here's a preview of some of the screenshots that you could make:

https://preview.redd.it/b4x7r5lcs5f41.png?width=3726&amp;format=png&amp;auto=webp&amp;s=073d68d6750cc3b57fee8cd412665e8822f5d4f6
## [12][Why is my Android virtual device so slow on Ubuntu 18?](https://www.reddit.com/r/androiddev/comments/ezrkpl/why_is_my_android_virtual_device_so_slow_on/)
- url: https://www.reddit.com/r/androiddev/comments/ezrkpl/why_is_my_android_virtual_device_so_slow_on/
---
I've been developing an app for Android using the Pixel 3a Oreo virtual device. Every time I start up the device, my computer completely tanks, to the point of not being able to move my mouse for 30 seconds at a time. It only runs slightly better once set up, and practically freezes my computer for 5 minutes when shutting down.

I followed [these](https://developer.android.com/studio/run/emulator-acceleration) [tutorials](https://help.ubuntu.com/community/KVM/Installation) on VM acceleration fully and they don't appear to have helped.

This isn't the case for my friends on Windows, and I'm sure it isn't the case for most people on Ubuntu, so do you know how I can make the device run better?

I will gladly provide any information needed in solving this problem.
