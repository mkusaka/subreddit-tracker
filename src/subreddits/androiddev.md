# androiddev
## [1][App Feedback Thread - October 17, 2020](https://www.reddit.com/r/androiddev/comments/jcua3k/app_feedback_thread_october_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jcua3k/app_feedback_thread_october_17_2020/
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
## [2][Weekly Questions Thread - October 12, 2020](https://www.reddit.com/r/androiddev/comments/j9oso5/weekly_questions_thread_october_12_2020/)
- url: https://www.reddit.com/r/androiddev/comments/j9oso5/weekly_questions_thread_october_12_2020/
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
## [3][Google removed app by mistake](https://www.reddit.com/r/androiddev/comments/jcsasd/google_removed_app_by_mistake/)
- url: https://www.reddit.com/r/androiddev/comments/jcsasd/google_removed_app_by_mistake/
---
# TL;DR

Google wrongfully removed an app from the Google Play Store. The reviewer somehow made a mistake and even documented it with screenshots attached to the removal notice. With no way to contact Google, the app is off the app store for hours, money being lost every hour. Has Google completely lost it?

# Long story

The app metrics looked slightly odd all Friday, but I thought, well, it’s a Friday and with ever changing Corona lockdowns all over the globe, we live in odd times anyway.

Until I realized at night that the user count had dropped significantly. I searched on Google Search for the app store link, but clicking on the result led to an error page:

&gt; We're sorry, the requested URL was not found on this server.

I then searched on the Google Play Store for the app, and it really could not be found. Rushing to the Google Play Developer Console I see a notice that the app has been removed and I should have received an email with details.

&gt; Your app has been removed from Google Play for not adhering to the Google Play Developer Program policies. New users can't find and install your app, and existing users won't receive updates. We've sent more information to the account owner's email address.

I looked for the email but I could find any, nor in the Trash, Junk, Archive. I later found out that Google is apparently falsely claiming to send out these emails, because so many developers [complain](https://support.google.com/googleplay/thread/32034066?hl=en) that they didn’t receive one. I noticed that I hadn’t opted into the new Developer Console website, and on the new site I actually found a link to display the email that was allegedly sent out.

The reasons for removal were related to the User Generated Content policy, specifically:
- the app doesn’t require a user to accept the Terms of Service 
- the Terms of Service do not state the restrictions on user generated content 

Both claims are false. The reviewer attached some screenshots of the walkthrough that is displayed after app launch as proof. The screen where the user has to accept the TOS and even explicitly agree to a summary regarding restrictions on user generated content was simply missing. It is not even technically possible to skip it.

How could that happen? Looking into the server logs later on it seems that Google outsources its app reviews to the Philippines for cheap labor costs, possibly with the result of low quality app reviews, leading to severe mistakes like this one where an app with millions of downloads is removed without prior notice due to a single reviewer’s mistake. 

Procedurally, this should have not been possible, but Google’s internal review processes are apparently a mess. They only recently introduced a more stringent app review process (rightly so) but this is plagued by [misinformation and lack of transparency](https://www.androidpolice.com/2019/09/03/google-play-expands-store-app-review/) ever since.

Luckily, the removal notice contains an appeal link to quickly resolve this mistake - or so I thought:

&gt; If you’ve reviewed the policy and feel our decision may have been in error, please reach out to our policy support team. We'll get back to you within 2 business days.

I submitted an appeal with screenshots showing that there is in fact a TOS screen. But already at the beginning of the appeal form, Google actually discourages from filing an appeal:

&gt; For a quicker resolution, we highly encourage you to refer to the original notification email to see how you can resubmit your app without an appeal. In most cases, you can make changes to bring your app into compliance and submit your app again.

Google even requires to check a box that one has read the above and agrees to it. Whoever of their legal team though of that and for what reason, but it makes an appeal look fishy - against which claims are they trying to protect themselves?

I try to get more clarification on that whole mess by heading back to the Developer Console and opening the Developer Support Chat. The link now leads to a new form that I fill out to get connected to an agent - or so I thought, because after filling out the whole form, it fails to submit:

&gt; There were problems sending your form. Please try again.

I remember the same error from some months ago and now that I need something urgent, of course it still doesn’t work. Random? Doesn’t look like it. Google seems to be trying everything they can to insulate themselves from contact with their developers.

At this point I should mention that I have a call-through phone number to our reviewer on the Apple App Review Team for urgent cases regarding the Apple App Store, a stark contrast in company culture and how they value developer relations.

I am left to submit the appeal, but it is Friday night and with the weekend ahead, when will the app be available again? Mid next week? End of next week? Due to COVID19 even later? Impossible to wait. So I quickly compile the unchanged app code with a new app version code and submit the update for review, in the hope that this works faster as Google itself suggests.

It is the next day, the app is now removed from the store for approximately 24 hours and still in review. It is a Saturday, with weekends usually generating the highest download numbers, customer complaints are flooding in why the app is unavailable on the Google Play Store. Every hour waiting means more upset customers and more money lost.

Is this the way Google intents to treat their developers? How is it possible that a company with a clear monopolistic market dominance has free reign without proper regulation to impose arbitrary rules and low quality processes that are lacking due diligence on hundreds of thousands of developers who are economically dependent on their platform?

There is something severely wrong here and it needs to be changed, through government policy, because self-regulation is a myth.
## [4][Ugly Android Studio 4.1 icon](https://www.reddit.com/r/androiddev/comments/jcceq3/ugly_android_studio_41_icon/)
- url: https://www.reddit.com/r/androiddev/comments/jcceq3/ugly_android_studio_41_icon/
---
Am I the only one who noticed ugly 4.1 icon? Before it looked stylish and on par with Visual Studio, Xcode, other development tools and now it looks like a cheap knockoff made by a student for a bowl of rice. It looks especially bad on Mac OS.
## [5][Supercharing mobile app development with GPU accelerated machine learning using Android NDK and Vulkan Kompute](https://www.reddit.com/r/androiddev/comments/jcuwwe/supercharing_mobile_app_development_with_gpu/)
- url: https://towardsdatascience.com/gpu-accelerated-machine-learning-in-your-mobile-applications-using-the-android-ndk-vulkan-kompute-1e9da37b7617
---

## [6][Where to store the data for a social media app?](https://www.reddit.com/r/androiddev/comments/jcrdw2/where_to_store_the_data_for_a_social_media_app/)
- url: https://www.reddit.com/r/androiddev/comments/jcrdw2/where_to_store_the_data_for_a_social_media_app/
---
Hi, if I want to create a social media app for learning purposes. I am going to create not the clone social media app. I have an own idea for social media app.  Which options for store the database are better? AWS? Firebase? or other options? Is it free to use if I use it for learning purpose?
## [7][How to query for autocomplete suggestions?](https://www.reddit.com/r/androiddev/comments/jcuwcl/how_to_query_for_autocomplete_suggestions/)
- url: https://www.reddit.com/r/androiddev/comments/jcuwcl/how_to_query_for_autocomplete_suggestions/
---
There are a list of autocomplete word suggestions shown as we type on the android keyboard. Is there any API through which I can pass some text and get back similar suggestions?  
Note that my app is not a keyboard app.
## [8][[DEV][GIVEAWAY] Big update for AutoTagger - music tag editor. 100 promo codes inside!](https://www.reddit.com/r/androiddev/comments/jcrwhh/devgiveaway_big_update_for_autotagger_music_tag/)
- url: https://www.reddit.com/r/androiddev/comments/jcrwhh/devgiveaway_big_update_for_autotagger_music_tag/
---
Hi folks! I'm a dev of *AutoTagger* app.  If you haven't heard about *AutoTagger* \- it's music tag editor with tag search feature (also it can search tags using track recognition) and many batch tag editing operations.

&amp;#x200B;

Since my last post about *AutoTagger*, i've added a big portion of features to the app.

Now it's only tag editor across Google Play which has batch editing features like on desktop apps.

So, here are some of new features:

* 6 full-featured batch operations: tag search, tag editing (with case transforming and tags deletion), files renaming, tracks numerating, extracting tags from filenames, batch file info viewing (bitrate, etc).
* Many new tag fields for editing
* Smooth material design animations and haptic feedback
* Search by files, tracks, albums and artists
* In-built audio player
* Fullscreen cover art viewing
* Displaying cover art size in tag search

&amp;#x200B;

[Download from Google Play](https://play.google.com/store/apps/details?id=com.serg.chuprin.tageditor)

Of course, *AutoTagger* has free and paid version which allows you to feel the full power of the app. However, you can try most of features for free (all batch operations are free for 5 tracks at once).

I attach **100 promo codes** for you guys. Please mark code as taken after use.

[Extended version (50 items)](https://docs.google.com/spreadsheets/d/1HfIMCD1Qd0-kcKsRHTq03yYQooJWiEZu5v6StxliX4k/edit?usp=sharing)

[Tag search using track recognition (50 items)](https://docs.google.com/spreadsheets/d/1NtMpLuJ9ChBotBx4N4Ts0_54r6HnHAasPMtRApSN_aA/edit?usp=sharing)

Promo codes are valid untill oct. 21.

&amp;#x200B;

Hope you'll like it.
## [9][Android-R scoped storage and SD card](https://www.reddit.com/r/androiddev/comments/jcrs18/androidr_scoped_storage_and_sd_card/)
- url: https://www.reddit.com/r/androiddev/comments/jcrs18/androidr_scoped_storage_and_sd_card/
---
Hi guys,
I am working with scoped storage to make my app compatible with it.

According to scoped storage we can't create any folder on the main storage of the device we have to use DCIM or Pictures folder for that.

But i am confused about SD card. Does same rule apply to it? Because i am testing on pixel device and it doesn't have sd card slot so i am unable to test this usecase.

If anyone has worked with scoped then kindly answer.
Thanks
## [10][Best NPM Packages to build Animations using ReactJS | Ashutosh Hathidara](https://www.reddit.com/r/androiddev/comments/jctxu5/best_npm_packages_to_build_animations_using/)
- url: https://www.youtube.com/watch?v=OGGuw2dLBjk&amp;feature=share
---

## [11][Android control photo metadata](https://www.reddit.com/r/androiddev/comments/jcraga/android_control_photo_metadata/)
- url: https://www.reddit.com/r/androiddev/comments/jcraga/android_control_photo_metadata/
---
Say I want to create an android app that takes photos. Simple stuff. Now if I want the app to also control the metadata generated for the photo, how do I do that? Your answer is much appreciated.
## [12][Saving files on Android: total madness](https://www.reddit.com/r/androiddev/comments/jcbpfi/saving_files_on_android_total_madness/)
- url: https://www.reddit.com/r/androiddev/comments/jcbpfi/saving_files_on_android_total_madness/
---
Recently I was trying a achieve what I thought a relatively simple goal: copying an audio file from the app's internal storage to the external storage.

So I started to study the new MediaStore API: that's total madness. You need a 4 year degree to understand how it works. That API is pure garbage.

In the end I managed to copy the file on API 29+, but on API 28- no way: I had to use a deprecated field (DATA) and the audio file is not shown.

What do you think of data management in Android. Do you agree with me that it is pure garbage?
