# androiddev
## [1][App Feedback Thread - February 29, 2020](https://www.reddit.com/r/androiddev/comments/fbbuz9/app_feedback_thread_february_29_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fbbuz9/app_feedback_thread_february_29_2020/
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
## [2][Weekly Questions Thread - February 24, 2020](https://www.reddit.com/r/androiddev/comments/f8pltv/weekly_questions_thread_february_24_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f8pltv/weekly_questions_thread_february_24_2020/
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
## [3][Scabbard (Dagger Visualizer) 0.2.0 - SVG support for generated graphs, uses simple names, improved gradle plugin and minor fixes.](https://www.reddit.com/r/androiddev/comments/fbqsbw/scabbard_dagger_visualizer_020_svg_support_for/)
- url: https://github.com/arunkumar9t2/scabbard/releases/tag/0.2.0
---

## [4][Rx is not the problem; Asynchronous programming is.](https://www.reddit.com/r/androiddev/comments/fbpgza/rx_is_not_the_problem_asynchronous_programming_is/)
- url: https://medium.com/@steven.popovich/rx-is-not-the-problem-asynchronous-programming-is-6055b5839b1a
---

## [5][Showing the Android Keyboard Reliably](https://www.reddit.com/r/androiddev/comments/fbiakx/showing_the_android_keyboard_reliably/)
- url: https://developer.squareup.com/blog/showing-the-android-keyboard-reliably/
---

## [6][Is there any FREE, SIMPLE, COMMUNITY based translation tool as Google Translator toolkit?](https://www.reddit.com/r/androiddev/comments/fbro2j/is_there_any_free_simple_community_based/)
- url: https://www.reddit.com/r/androiddev/comments/fbro2j/is_there_any_free_simple_community_based/
---
Everything I tried, and its been a lot, is either paid, or directed towards professionals, or with landing page which takes an hour to read and offering tons of functionality I dont really need.

I just want something simple. Where I:

1. upload text I want to translate
2. share link with people who help me with translation

Thats all. Exactly like Google Translator Toolkit. Can you recommend any tools? Thanks a lot!
## [7][How would I go about creating an app to intercept and edit all mobile hotspot traffic?](https://www.reddit.com/r/androiddev/comments/fbto0z/how_would_i_go_about_creating_an_app_to_intercept/)
- url: https://www.reddit.com/r/androiddev/comments/fbto0z/how_would_i_go_about_creating_an_app_to_intercept/
---
I would like to build an app with xamarin using c# which should be able to:
1: setup a mobile hotspot with a custom name and password.
2: intercept and edit all traffic passing through the hotspot.

Would the WifiManager class give me enough control to do this? If not, what tools could give me this level of control? If there are no tools to do this then how would I get low level access to networking on the android device?

Thanks in advance for any help! :)
## [8][How to speed up image processing on android?](https://www.reddit.com/r/androiddev/comments/fbsnpt/how_to_speed_up_image_processing_on_android/)
- url: https://www.reddit.com/r/androiddev/comments/fbsnpt/how_to_speed_up_image_processing_on_android/
---
I am working on an image enhancement android app. In one part i loop over a single channel 2d array with dimensions equal to that of image. Doing this on a full resolution image takes too much time and the app crashes. So right now i downscale the image and work on it but I have to do it on full resolution image.

What i am doing is that i have already created 10 different segment of image. I have labelled the above mentioned 2d array where each index is the segment number from 1 to 10. I then sum pixel values at each belonging to a specific segment. For this I loop over the 2d array and check which segment the pixel belongs to and add the pixel value to corresponding sum.

To speed up this process I can think of two techniques.
1. Create 10 different images making pixels that don't belong to that segment equal to 0 and than using opencv's built in sum function.
2. I am doing this in C++ using jni and I have read online about Halide and how it is used to speed up such tasks by parallelising the tasks. I currently don't know how to do it but I can learn. 

Which of the above mentioned technique is better or if there exists an even better technique which I have missed?
## [9][I opened my minesweeper game source code](https://www.reddit.com/r/androiddev/comments/fbfomo/i_opened_my_minesweeper_game_source_code/)
- url: https://github.com/lucasnlm/antimine-android
---

## [10][Any way to set default mode to "Split" view in Layout Editor (Android Studio 3.6)?](https://www.reddit.com/r/androiddev/comments/fbc5f3/any_way_to_set_default_mode_to_split_view_in/)
- url: https://www.reddit.com/r/androiddev/comments/fbc5f3/any_way_to_set_default_mode_to_split_view_in/
---
I went to `Settings-&gt;Editor-&gt;Layout Editor` but it only has "Prefer XML editor" checkbox, which means it'll either open the Design view, or Code view, but not split view. Is there any way I can make it open Split view by default?
## [11][Android generic receipt printing](https://www.reddit.com/r/androiddev/comments/fbppn8/android_generic_receipt_printing/)
- url: https://www.reddit.com/r/androiddev/comments/fbppn8/android_generic_receipt_printing/
---
I am currently using Epson sdk to connect to epson thermal printers, but I would like to be able to connect to other manufacter's printers too, is there any generic way to print? I tried using sockets but it doesn't work.
## [12][New Android App Ideas](https://www.reddit.com/r/androiddev/comments/fbrq8w/new_android_app_ideas/)
- url: https://www.reddit.com/r/androiddev/comments/fbrq8w/new_android_app_ideas/
---
A recent app I was working on I had to take down unfortunately because the API is no longer operational  
[https://pitted.app/](https://pitted.app/)  
So now I am looking to build a new Android app as a side project

I have created a couple other Android apps [https://play.google.com/store/apps/developer?id=Etienne+Lawlor&amp;hl=en\_US](https://play.google.com/store/apps/developer?id=Etienne+Lawlor&amp;hl=en_US)

and have worked on several open source projects [https://github.com/lawloretienne?tab=repositories](https://github.com/lawloretienne?tab=repositories)

But I want to work on something new.

Does anyone have any ideas, or has any interest in possibly partnering up?
