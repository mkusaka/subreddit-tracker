# androiddev
## [1][App Feedback Thread - February 08, 2020](https://www.reddit.com/r/androiddev/comments/f0qz4w/app_feedback_thread_february_08_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f0qz4w/app_feedback_thread_february_08_2020/
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
## [3][HTML Admin CMS for Firebase?](https://www.reddit.com/r/androiddev/comments/f0q9kr/html_admin_cms_for_firebase/)
- url: https://www.reddit.com/r/androiddev/comments/f0q9kr/html_admin_cms_for_firebase/
---
Firebase's own admin layout is pretty much useless and wastes a lot of time so is there any other HTML based admin sites/projects that I can connect my database with and do changes in Firestore? Most data on my Firestore is input from me and it's very hard to create individual nodes again and again since there is a lot of data that I enter.
## [4][This is my fault. Somehow.](https://www.reddit.com/r/androiddev/comments/f0bvgu/this_is_my_fault_somehow/)
- url: https://www.reddit.com/r/androiddev/comments/f0bvgu/this_is_my_fault_somehow/
---
https://imgur.com/a/AFQcVvK

Edit: so yeah, I requested a removal and the bots just told me to eat the review, it's staying.
Thanks Google Dev support!
## [5][Android TV, and interstitial ads.](https://www.reddit.com/r/androiddev/comments/f0rlcv/android_tv_and_interstitial_ads/)
- url: https://www.reddit.com/r/androiddev/comments/f0rlcv/android_tv_and_interstitial_ads/
---
I developed a small app for a client, the app contains AdMob interstitial ads.  


the client asked me to develop the same kind of app but for Android TV. from research online I found that AdMob is not officially supported on Android TV yet.   


is there an alternative ad network that will work with Android TV? does anyone here have any experience with such an issue?
## [6][I just published an article about ViewBinding on medium. Please check it out.](https://www.reddit.com/r/androiddev/comments/f0rk0e/i_just_published_an_article_about_viewbinding_on/)
- url: https://www.reddit.com/r/androiddev/comments/f0rk0e/i_just_published_an_article_about_viewbinding_on/
---
“Exploring ViewBinding in depth” by Somesh Kumar https://link.medium.com/Jprz4AjpU3
## [7][Help me with my Kotlin Flow hangups](https://www.reddit.com/r/androiddev/comments/f0iddv/help_me_with_my_kotlin_flow_hangups/)
- url: https://www.reddit.com/r/androiddev/comments/f0iddv/help_me_with_my_kotlin_flow_hangups/
---
Or more precisely (but less catchy as a title) my Cold vs Hot sequence hangups.

First of all, it seems like 3 rather distinct concepts are associated with Hot vs. Cold, maybe because they're not as distinct as I think they are:

1. Whether the Observable is generating items when no subscriber is listening (Hot), or whether it waits until there is a subscriber before generating items (Cold)
2. Whether the Observable can be back-pressured (Cold) or not (Hot)
3. Whether the Observable broadcasts its events to all subscribers (Hot) or whether an Observable generates new items for each subscriber (Cold) -- i.e. whether the Observable is shared.

I can definitely see how they are related. Perhaps these concepts go together because they just make sense together. Anyway, that's not my question.

It seems that A LOT of events in mobile app dev (user input, data updates pushed from the back end, etc.) are inherently Hot. You don't have any control over when they happen. You generally want to broadcast/share them. You could try to adopt a back-pressure strategy (e.g. buffer) but you can't back-pressure the event itself.

Notable exception of course is I/O, e.g. unary network calls, or file I/O. These can be "Single-like" in nature, or maybe you have a paginated network API or you're reading individual lines from a file, which would involve multiple events and is "Flowable-like" in nature. These seem Cold to me.

I could also imagine having a Presenter in which you've subscribed to Flowable&lt;ViewState&gt; and you request 1 item at a time because there's no sense in processing multiple ViewStates each frame.

So I see a use for both, but the Kotlin language designers don't seem to be feeling the love for Hot channels. Channels are still experimental, and many of their (seemingly useful!) operators are being deprecated (with instructions to use Flow instead). I definitely understand [their limitations](https://medium.com/@elizarov/cold-flows-hot-channels-d74769805f9), but they seem necessary.

So my question is, should we be minimizing the use of Hot sequences? Should we be converting them to Cold sequences with some sort of Back-Pressure Strategy as immediately as possible? Should the frame timing of our UI generally be the tick rate at which we calculate new app state, with "Hot" events buffered somehow until the next tick?
## [8][I need help, google play leaderboard won't record scores](https://www.reddit.com/r/androiddev/comments/f0p1ma/i_need_help_google_play_leaderboard_wont_record/)
- url: https://www.reddit.com/r/androiddev/comments/f0p1ma/i_need_help_google_play_leaderboard_wont_record/
---
In my game I am able to log into google play and even have the leaderboard pop open, but i can't record scores into the leaderboard. I used this code I found in a tutorial, and it seems to be the only method I can find anywhere:

 

public static void AddScoreToLeaderboard(string leaderboardId, long score)

{

Social.ReportScore(score, leaderboardId, success =&gt; { });

}

&amp;#x200B;

But it does nothing. It used to work for me a year ago, but since google now requires games to be in 64bit, i had to update everything, and now this code doesn't work for me. I'm guessing maybe google changed it in an update, but I can't find any other code to record to leaderboards. Plz help. Thanks.
## [9][Critical Bluetooth Vulnerability in Android (CVE-2020-0022)](https://www.reddit.com/r/androiddev/comments/f0dflc/critical_bluetooth_vulnerability_in_android/)
- url: https://insinuator.net/2020/02/critical-bluetooth-vulnerability-in-android-cve-2020-0022/
---

## [10][Android 11 will add a new App Compatibility feature to ease testing](https://www.reddit.com/r/androiddev/comments/f081sw/android_11_will_add_a_new_app_compatibility/)
- url: https://www.developer-tech.com/news/2020/jan/22/android-11-new-app-compatibility-feature-testing/
---

## [11][View Binding: Merge](https://www.reddit.com/r/androiddev/comments/f0corf/view_binding_merge/)
- url: https://blog.stylingandroid.com/view-binding-merge/
---

## [12][Changing minSdkVersion from 21 to 23?](https://www.reddit.com/r/androiddev/comments/f09jni/changing_minsdkversion_from_21_to_23/)
- url: https://www.reddit.com/r/androiddev/comments/f09jni/changing_minsdkversion_from_21_to_23/
---
What is your minSdkVersion? Right now in our apps it's mostly API 21, but we are considering to move to API 23, no specific reason except some code optimization and possibly APK size reduction.   


According to analytics 13% of our users are on API level 21, but here is the catch, 69% of them are from India and Mexico and those countries tend to have very low eCPM and are not very profitable.   
What are your thoughts?  
What is your minSdkVersion?
