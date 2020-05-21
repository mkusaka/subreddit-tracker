# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/gm0rsw/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/gm0rsw/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - May 18, 2020](https://www.reddit.com/r/androiddev/comments/glzbpr/weekly_questions_thread_may_18_2020/)
- url: https://www.reddit.com/r/androiddev/comments/glzbpr/weekly_questions_thread_may_18_2020/
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
## [3][“Postmortem: How I got banned from /r/androiddev for being ‘too’ active as a member of the Community” by Gabor Varadi](https://www.reddit.com/r/androiddev/comments/gnv79i/postmortem_how_i_got_banned_from_randroiddev_for/)
- url: https://link.medium.com/IO2y8v1qF6
---

## [4][Introducing Gryphon, the Swift to Kotlin translator](https://www.reddit.com/r/androiddev/comments/gnm5da/introducing_gryphon_the_swift_to_kotlin_translator/)
- url: https://www.reddit.com/r/androiddev/comments/gnm5da/introducing_gryphon_the_swift_to_kotlin_translator/
---
Hi all, I’ve just published the first full version Gryphon, a program that translates Swift code into Kotlin code. It is meant to let developers share their iOS app's code with Android.

I encourage anyone interested to [check out the video](https://twitter.com/gryphonblog/status/1263233844519620610?s=20) and [the website](https://vinivendra.github.io/Gryphon/index.html).

This is the first version where all features are working:

* The generated Kotlin code works the same as the Swift code it came from, no edits needed. There’s an automated test that makes Gryphon translate its own source code, with around 12k lines, and ensures the translation passes all the same tests as the original executable.
* There’s a templates system that’s used to automatically translate many standard library types and methods.
* Xcode integration, with Xcode showing any warnings or errors raised by the Kotlin compiler next to the Swift lines that originated them, so users can fix them at the source.
* The output code is readable, which is useful for minimizing the risks for new users. If Gryphon isn’t the right fit for them, they’re left with a Kotlin codebase they can keep maintaining.

This list of features isn’t to say [there aren’t any bugs](https://github.com/vinivendra/Gryphon/issues), of course - I’m doing my best to keep track of them and fix them on a daily basis.

Any questions, comments, or feedback in general is welcome!
## [5][Multiple navigation graphs (with bottom navigation bar)](https://www.reddit.com/r/androiddev/comments/gnipdq/multiple_navigation_graphs_with_bottom_navigation/)
- url: https://www.rockandnull.com/jetpack-navigation-bottom-navigation-view/
---

## [6][Customize text for app share.](https://www.reddit.com/r/androiddev/comments/gnuu56/customize_text_for_app_share/)
- url: https://www.reddit.com/r/androiddev/comments/gnuu56/customize_text_for_app_share/
---
I added my app to the group of the apps that appears on the share menu, so I can get data from orhers app.  
Now how can I customize the text below. For example for Reddit it says: "Post on Reddit".
## [7][Kotlin Extension function to chain ConstraintLayout animations with TransitionManager](https://www.reddit.com/r/androiddev/comments/gnu92q/kotlin_extension_function_to_chain/)
- url: https://www.reddit.com/r/androiddev/comments/gnu92q/kotlin_extension_function_to_chain/
---
So this basically just runs one Transition after the previous one ends, useful where you want to chain several ConstraintLayout Transitions one after another. As there was no "true" built-in method to do this, I decided to create a small extension function for it. 

&amp;#x200B;

PS: TransitionSet with the SEQUENTIAL option didn't work for me

&amp;#x200B;

[https://gist.github.com/DDihanov/6624925ced3b4db6f4ce6cbe1704a891](https://gist.github.com/DDihanov/6624925ced3b4db6f4ce6cbe1704a891)
## [8][Coroutines &amp; Patterns for work that shouldn’t be cancelled](https://www.reddit.com/r/androiddev/comments/gntqtr/coroutines_patterns_for_work_that_shouldnt_be/)
- url: https://medium.com/androiddevelopers/coroutines-patterns-for-work-that-shouldnt-be-cancelled-e26c40f142ad
---

## [9][How important is familiarity with Gradle as a skill for an Android developer?](https://www.reddit.com/r/androiddev/comments/gnm5k4/how_important_is_familiarity_with_gradle_as_a/)
- url: https://www.reddit.com/r/androiddev/comments/gnm5k4/how_important_is_familiarity_with_gradle_as_a/
---
I understand that Gradle helps automate dependency management, task automation, and other things, but I can't find anything online that discusses its importance in the daily use by Android developers. Do you use the features with Gradle a lot in your work? As a beginner learning it in Udacity, I'm just being introduced to it. How important is this skill? When you're hiring someone, is this something you'd value in them?
## [10][Interest in Google Photos app for Android TV? Looking to get some public sentiment/opinion.](https://www.reddit.com/r/androiddev/comments/gnr0xc/interest_in_google_photos_app_for_android_tv/)
- url: https://www.reddit.com/r/androiddev/comments/gnr0xc/interest_in_google_photos_app_for_android_tv/
---
Google Photos is an amazing product from Google. As a photo enthusiast, I've been using it since it's inception. For the longest time, there was no dedicated API for Google Photos until [mid-2018](https://developers.googleblog.com/2018/05/introducing-google-photos-partner.html). Before that, I believe the photos were accessible via [Picasa Web Albums API](https://support.google.com/picasa/answer/6383491?hl=en).

I'm an Android developer and own nVidia Shield [Android TV](https://www.android.com/intl/en_us/tv/). Even though the Google Photos app has Google [cast](https://developers.google.com/cast) support to show your photos or slideshows on Android TV/Chromecast, I miss the native Android TV photo browsing experience.

I was wondering if there are people like me who has Android TV and have a similar craving to have a native app for Google Photos. ***If there were an Android TV native app for Google Photos what would you wish it had?***

\----

ps. As far as I know, the following features are already supported by Google Photos

* Show pictures on TV via Google Cast
* Show slideshow of photo album via Google Cast
* Select photos from albums to use as screen-saver via Google Home app
## [11][DO Android Jobs pay well?](https://www.reddit.com/r/androiddev/comments/gnjjzm/do_android_jobs_pay_well/)
- url: https://www.reddit.com/r/androiddev/comments/gnjjzm/do_android_jobs_pay_well/
---
I feel like everyone wants to get into Android and so there is a lot of competition and it keeps the pay down. Kinda like the gaming industry?
## [12]["Clean Dagger - A pragmatic guide to dependency injection on Android" - what do you think? Do you also use this approach in your applications?](https://www.reddit.com/r/androiddev/comments/gn8spj/clean_dagger_a_pragmatic_guide_to_dependency/)
- url: https://proandroiddev.com/clean-dagger-f248eda5790b
---

