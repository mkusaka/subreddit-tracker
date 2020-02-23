# androiddev
## [1][App Feedback Thread - February 22, 2020](https://www.reddit.com/r/androiddev/comments/f7rxif/app_feedback_thread_february_22_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f7rxif/app_feedback_thread_february_22_2020/
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
## [2][Weekly Questions Thread - February 17, 2020](https://www.reddit.com/r/androiddev/comments/f574wo/weekly_questions_thread_february_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f574wo/weekly_questions_thread_february_17_2020/
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
## [3][Random Musings on the R Developer Preview 1 [CommonsWare]](https://www.reddit.com/r/androiddev/comments/f7pz5n/random_musings_on_the_r_developer_preview_1/)
- url: https://www.reddit.com/r/androiddev/comments/f7pz5n/random_musings_on_the_r_developer_preview_1/
---
&gt;**What Will Piss Off Users**  
&gt;  
&gt;The phrasing around the [file and directory access restrictions](https://developer.android.com/preview/privacy/storage#file-directory-restrictions) puts the blame on developers. How *dare* we ask for the Downloads/document tree, or for files in Android/data/directories!  
In  reality, developers don’t do that. Developers simply fire off  ACTION\_OPEN\_DOCUMENT, ACTION\_OPEN\_DOCUMENT\_TREE, or  ACTION\_CREATE\_DOCUMENTIntents. Where the user winds up choosing is up to  the user.  
So, in reality, **the restrictions listed in this section will be placed on** ***users***\*\*, not developers\*\*. And I have absolutely no idea **what sort of messaging Google could use** in the Storage Access Framework UI **that would justify to the user blocking them from navigating to certain portions of external and removable storage**. “Sorry, you can’t choose Downloads/, because ¯\\\_(ツ)\_/¯” doesn’t strike me as something that will make users very happy.  
I  am all for improved user control over privacy and security concerns,  but the boundaries between what is and is not allowed need to make sense  to users and, secondarily, to developers.  
&gt;  
&gt;  
&gt;  
&gt;**What Will Piss Off Developers**  
&gt;  
&gt;On the one hand, the [“All Files Access”](https://developer.android.com/preview/privacy/storage#all-files-access) support should be a good thing. It gives file managers and similar sorts of apps an option for accessing more stuff on external and removable storage. However, it is not really “all files”, as Android/data/  
 is still blocked off. IMHO, the next step is to use the one-time permission system to allow these apps to have access to these protected areas (e.g., so a backup app can actually do a backup).  
&gt;  
&gt;The [background location restrictions](https://developer.android.com/preview/privacy/location#background-location) are confusing as documented. I’ll be playing around with these and hopefully will work out some samples soon.

&amp;#x200B;

[https://commonsware.com/blog/2020/02/21/random-musings-r-dp1.html](https://commonsware.com/blog/2020/02/21/random-musings-r-dp1.html)
## [4][How long did it take you to get your first dev job?](https://www.reddit.com/r/androiddev/comments/f7ihtr/how_long_did_it_take_you_to_get_your_first_dev_job/)
- url: https://www.reddit.com/r/androiddev/comments/f7ihtr/how_long_did_it_take_you_to_get_your_first_dev_job/
---
I'm sure this has been posted here before but my search didn't return much. I have been learning Android and Kotlin for the past two years and started applying for positions last July. In those nine months I have had three phone interviews (two of which were with recruiters) and no face to face/actual coding interviews. My job search has mainly been through job boards like LinkedIn and indeed. Most recruiters say that my resume is on point, I have two apps publish to the Google play store and several other projects on my GitHub. Why am I not getting anywhere? How long did it take you guys to get your first engineer position and how did you get it? Feeling like maybe I need to switch up my approach! Any info helps, thanks people!
## [5][Compatibility](https://www.reddit.com/r/androiddev/comments/f7qoo4/compatibility/)
- url: https://www.reddit.com/r/androiddev/comments/f7qoo4/compatibility/
---
Hi all,

My app seems to work OK across all Android versions - except Oreo 8.1 - the audio does not work on Oreo 8.1. Is there anything special about this version of Android? Has anyone else encountered this?

Thanks.
## [6][So Google has just banned over 600 Apps for ad policy violation, including 40 apps from Cheetah Mobile, wondering how long until Cheetah Mobile returns?](https://www.reddit.com/r/androiddev/comments/f7bbip/so_google_has_just_banned_over_600_apps_for_ad/)
- url: https://www.reddit.com/r/androiddev/comments/f7bbip/so_google_has_just_banned_over_600_apps_for_ad/
---
Here is more info:  [https://www.androidauthority.com/google-play-store-app-ban-cheetah-1085601/](https://www.androidauthority.com/google-play-store-app-ban-cheetah-1085601/)   


Who wants to bet that by the end of the month Cheetah Mobile will be back on GooglePlay store?
## [7][AdSense disabled for no apparent reason](https://www.reddit.com/r/androiddev/comments/f7s0zo/adsense_disabled_for_no_apparent_reason/)
- url: https://www.reddit.com/r/androiddev/comments/f7s0zo/adsense_disabled_for_no_apparent_reason/
---
I have an app that I released 4-5 months ago just for fun and since I didn't really market it, it didn't have many downloads. A few days ago I started sharing it here and there, without any paid advertisement so I started getting some downloads. (Almost 300 now). 
A few days ago, I got an email from Google saying that they have put a limit on my Ads. I didn't really care since I wasn't making any money from AdMob. Yesterday I got an email from Google again saying that my account is banned because there was "Invalid traffic" on my app... it's has barely any users lol. Anyways, I replied by saying that I had no idea about this and they said my account will never be reopened again.

Anyone else experienced the same thing? And what are some good alternatives?
## [8][I was scammed by Appodeal](https://www.reddit.com/r/androiddev/comments/f7fhm6/i_was_scammed_by_appodeal/)
- url: https://www.reddit.com/r/androiddev/comments/f7fhm6/i_was_scammed_by_appodeal/
---
Appodeal is a mobile ad provider similar to Google Ads which I used to for a smaller app for a bit of time.

I had a balance of over 130$ revenues on my account and requested a payout. Guess what they did: They terminated my account and told me I was violating their Terms of Service while completely unspecifically quoting their ToS. After multiple requests they sent me a made up chart which listed device ids an ip addresses with large numbers of impressions and stating that this would be "invalid activity" leading to termination of account.

To be clear: I never engaged in fake clicking and just from seeing the ips being american where i don't know anyone I can safely assume that these fake impressions are just made up so that they don't have to pay out anything.

I requested to pay out the "valid" revenue but they just stopped responding to my emails.

So I guess the money is gone.

This story just as a warning: Don't use Appodeal!
## [9][Keep it or leave it](https://www.reddit.com/r/androiddev/comments/f7r1n4/keep_it_or_leave_it/)
- url: https://www.reddit.com/r/androiddev/comments/f7r1n4/keep_it_or_leave_it/
---
Hey guys. I'm planning to make a flutter app for my graduation project I'm very good at Java but have 0 knowledge in Dart.
I found a fantastic flutter UI to start with but I'm having trouble understanding Dart. Should i just go back to Java (suggestion on nice UI templates please?) or just spend the time to learn Dart and everything else related to flutter?
In other words; can i have the same flexability and access great UIs in Java as in Flutter?
## [10][The Most Interesting Takeaways From Android 11 So Far](https://www.reddit.com/r/androiddev/comments/f7qwhw/the_most_interesting_takeaways_from_android_11_so/)
- url: https://i.redd.it/z4cxp9kvcgi41.png
---

## [11][Where would I go to learn about working with more advanced graphics (Such as those in the Timely app)?](https://www.reddit.com/r/androiddev/comments/f7jirt/where_would_i_go_to_learn_about_working_with_more/)
- url: https://www.reddit.com/r/androiddev/comments/f7jirt/where_would_i_go_to_learn_about_working_with_more/
---
I've always liked the look of the Timely app. It has some of the most smooth and enjoyable graphics and UI interactions I've seen in an android app. Here's a video:

https://youtu.be/8hYQobl5xL4?t=9

Unfortunately, I have no idea how they achieved that. It seems a long way from the stock xml-based layouts, but I'm a beginner and I could be wrong and it could be easier than it looks. 

Anyone more experienced here care to chip in with some advice on how they could have made such graphics and interactions? For example at 13 seconds in the video, they swipe the screen to set the alarm, and the time to the left grows and shrinks beautifully. Think there could be some framework they used or would one have to code in raw OpenGL to make something like that?

Just any advice on where to start looking would help, really. Thanks!
## [12][Would a Firebase database be a sufficient way of incorporating simple leaderboards and global user statistics to a game?](https://www.reddit.com/r/androiddev/comments/f7qikx/would_a_firebase_database_be_a_sufficient_way_of/)
- url: https://www.reddit.com/r/androiddev/comments/f7qikx/would_a_firebase_database_be_a_sufficient_way_of/
---
Also, could the unique device ID be used to recognizing individual users in the database?
