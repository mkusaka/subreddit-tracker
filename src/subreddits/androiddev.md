# androiddev
## [1][App Feedback Thread - April 18, 2020](https://www.reddit.com/r/androiddev/comments/g3mqb2/app_feedback_thread_april_18_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g3mqb2/app_feedback_thread_april_18_2020/
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
## [2][Weekly Questions Thread - April 13, 2020](https://www.reddit.com/r/androiddev/comments/g0grat/weekly_questions_thread_april_13_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g0grat/weekly_questions_thread_april_13_2020/
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
## [3][Why is colorPrimary, colorPrimaryDark, colorAccent in res/colors in camelcase when other default values is in underscore case? Should colors then be named in camel case or should one stick to name everything in /res folders in underscore?](https://www.reddit.com/r/androiddev/comments/g3kydw/why_is_colorprimary_colorprimarydark_coloraccent/)
- url: https://www.reddit.com/r/androiddev/comments/g3kydw/why_is_colorprimary_colorprimarydark_coloraccent/
---

## [4][Google shenanigans](https://www.reddit.com/r/androiddev/comments/g33cx7/google_shenanigans/)
- url: https://i.redd.it/4zorclc05et41.jpg
---

## [5][Database Inspector in Android Studio 4.1 Canary version](https://www.reddit.com/r/androiddev/comments/g3aqx6/database_inspector_in_android_studio_41_canary/)
- url: https://goo.gle/34xvGDP
---

## [6][Flutter Countdown Timer](https://www.reddit.com/r/androiddev/comments/g3mijm/flutter_countdown_timer/)
- url: https://youtu.be/uPhhEgpyVY8
---

## [7][Coroutines and exceptions: things to know](https://www.reddit.com/r/androiddev/comments/g3mbpo/coroutines_and_exceptions_things_to_know/)
- url: https://www.rockandnull.com/coroutines-and-exceptions/
---

## [8][Android Studio 3.6.3 available](https://www.reddit.com/r/androiddev/comments/g36s3v/android_studio_363_available/)
- url: https://androidstudio.googleblog.com/2020/04/android-studio-363-available.html
---

## [9][Code review](https://www.reddit.com/r/androiddev/comments/g3n9dx/code_review/)
- url: https://www.reddit.com/r/androiddev/comments/g3n9dx/code_review/
---
For the last week I've been trying to brush up on my Android skills. It's been a few years since I last did any mobile development.

I've created the following app for the sake of learning. The idea is to present the user with a list of exercises which can be edited and updated with new exercises. The exercises are retrieved from a restful api which can be found [here](https://github.com/tonsV2/fitlog-api). And cached locally using SQLite.

https://github.com/tonsV2/my-koin-application

I'm interested in all the feedback you have. I'm fairly new to Kotlin as well. But mainly feedback of the architectural design is of interest. Please let me know what you think.
## [10][Securing Room DB](https://www.reddit.com/r/androiddev/comments/g3n616/securing_room_db/)
- url: https://www.reddit.com/r/androiddev/comments/g3n616/securing_room_db/
---
Hey, any hints on securing/encrypting Room DB? Link to a tutorial/example would be great.
## [11][I'm new to server side programming: How can I avoid storing user credential such as email and password and still have that functionality ?](https://www.reddit.com/r/androiddev/comments/g3ez7w/im_new_to_server_side_programming_how_can_i_avoid/)
- url: https://www.reddit.com/r/androiddev/comments/g3ez7w/im_new_to_server_side_programming_how_can_i_avoid/
---
I've created a small confession style app where users can confess whatever they want. This is based in Firebase realtime database with authentication. I store both username and email in server.

But now I am porting the backend to php/mysql and I don't' want to store user's email and password, and use something else instead so that even if the data is breached (remember I'm only learning right now), there would not be big repercussion.

How can I do that?
## [12][Implementing clean architecture in android application](https://www.reddit.com/r/androiddev/comments/g3j3vm/implementing_clean_architecture_in_android/)
- url: https://www.reddit.com/r/androiddev/comments/g3j3vm/implementing_clean_architecture_in_android/
---
Hey guys! I would love to get your feedback on two approaches i have seen while implementing clean architecture in android. While there are multiple ways one can implement clean architecture in android when in comes to making different modules or keeping everything in the app itself and rather slicing the project both vertically, as in feature, and horizontally, as in layers, the approach we will take is rather converting these features into android libraries themselves.

So here are the two approaches:
1. [Approach 1](https://github.com/igorwojda/android-showcase)
:APP(com.android.application), :BASE(com.android.library), :FEATURE1(com.android.library), :FEATURE2(com.android.library)

In this case, :APP is dependent on :BASE and :FEATURE1 and :FEATURE2 are dependent on :APP module.

2. [Approach 2](https://youtu.be/TWLkswxjSr0?t=1437)
:APP(com.android.application), :BASE(com.android.library), :FEATURE1(com.android.library), :FEATURE2(com.android.library)

In this case, :APP is **NOT** dependent on :BASE. Rather :FEATURE1 and :FEATURE2 are dependent on :BASE and :APP is dependent on the feature modules.

I personally like approach 2. Do let me know which one you prefer and why. I am building an android application and want to do it with clean architecture, however can't see beyond the lines of which approach is better than which. Thank you.
