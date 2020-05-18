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
## [3][Google bans Podcast Addict app after 9 years for letting users play podcasts that reference COVID-19](https://www.reddit.com/r/androiddev/comments/glx7q2/google_bans_podcast_addict_app_after_9_years_for/)
- url: https://reclaimthenet.org/google-play-suspends-podcast-addict/
---

## [4][AVD Query](https://www.reddit.com/r/androiddev/comments/gm0826/avd_query/)
- url: https://www.reddit.com/r/androiddev/comments/gm0826/avd_query/
---
What to choose for hardware acceleration for your android emulator when you don't have a graphics card installed in your computer? GLES hardware or Software?
## [5][Kevin Lynch (now VP of Technology at Apple) in 2010 on cross-platform apps](https://www.reddit.com/r/androiddev/comments/glxh5e/kevin_lynch_now_vp_of_technology_at_apple_in_2010/)
- url: https://www.youtube.com/watch?v=Z512TwwyRWM
---

## [6][Migrating paid app users to free app](https://www.reddit.com/r/androiddev/comments/glt218/migrating_paid_app_users_to_free_app/)
- url: https://www.reddit.com/r/androiddev/comments/glt218/migrating_paid_app_users_to_free_app/
---
Hi, i hope to get advise here.  i have decided to just have one app from a paid app and free app version before and new users will use iap to make it ad free.

how did you guys manager to detect paid app users and unlock them in the free version?

there does not seem to be any api in the billing api that does it since it involves only iap.

i had read about forcing the user to install the pro version so the free version can detect if the paid version is installed. this is not an ideal approach though since it forces the user to have 2 installed apps (free and paid) even if it is only a 1 time step to verify that the paid version is installed.

thoughts?
## [7][Dear Developers, there is a big difference % wise in earnings on AdMob in this month, compared to the last. Can you help me figure what's the problem?](https://www.reddit.com/r/androiddev/comments/gly9kf/dear_developers_there_is_a_big_difference_wise_in/)
- url: https://i.redd.it/8vy7qx99shz41.png
---

## [8][Hey, this is my first android app please have a look and give me suggestions. https://play.google.com/store/apps/details?id=com.aniket.randomize](https://www.reddit.com/r/androiddev/comments/gm12l9/hey_this_is_my_first_android_app_please_have_a/)
- url: https://i.redd.it/n6rczolzuiz41.png
---

## [9][Type com.github.ybq.android.spinkit.R is defined multiple times, because of aar library](https://www.reddit.com/r/androiddev/comments/gm11pa/type_comgithubybqandroidspinkitr_is_defined/)
- url: https://www.reddit.com/r/androiddev/comments/gm11pa/type_comgithubybqandroidspinkitr_is_defined/
---
Hello, I have integrated a third party SDK (aar file), which uses com.github.ybq.android.spinkit inside of it. This aar library, does not declare com.github.ybq.android.spinkit as gradle depency and this causes an issue, when we **try to create a release version of our app:**

**Type com.github.ybq.android.spinkit.R is defined multiple times**

Debug builds work without any issues. 

What I have tried:

1. Exclude a com.github.ybq.android.spinkit from this aar file. This did not worked, because it seems that it does not support **aar files**. Here is my code:

&amp;#8203;

     implementation ('example.library@aar') {
            exclude group: 'com.github.ybq', module: 'Android-SpinKit'
        }

I know that it is a fault of library creator to build **a fat aar** with compiled library additionally. However, I can't just force them to change it... 

Perhaps there is a way to somehow **exclude Android-SpinKit from that AAR library**?
## [10][Client/server CRUD application with client side cache](https://www.reddit.com/r/androiddev/comments/gm118e/clientserver_crud_application_with_client_side/)
- url: https://www.reddit.com/r/androiddev/comments/gm118e/clientserver_crud_application_with_client_side/
---
Recently I've been hacking away on a few different applications. The common denominator is that the Android applications are clients for RESTful backends. The way I've implemented it so far is by synchronizing the backend to local storage on the client. Specifically I'm using Retrofit and Room for that task.

The applications mostly consist of CRUD functionality and the business logic mostly consist of the way the data is structured and its relations.

The problem is that I feel I'm coding the application twice. I'm pretty much duplicating the database layout both client- and server-side. With the limitations of Room versus JPA and Postgresql.

To some extend I feel that Room/Sqlite is overkill and I could do with a simple key value store.

Any comments?
## [11][How to protect my personal play store account as a freelancer?](https://www.reddit.com/r/androiddev/comments/gm10b2/how_to_protect_my_personal_play_store_account_as/)
- url: https://www.reddit.com/r/androiddev/comments/gm10b2/how_to_protect_my_personal_play_store_account_as/
---
I will be mostly working with companies who never had an app. I will be building apps for them and guiding them through the process of publishing it on play store.

What are things I should keep in mind so that if in future if they get their account ban, my personal account is not affected by it?

If you have any other tips for someone starting freelancing please share them as well.

Thanks.

P.s. I am asking this because I have heard stories where a company got their account ban because a developer who was banned had access to it and google associated the company's account with his. I just want to be sure that I am not affected by it the other way round.
## [12][Someone modified my app and uploaded it to some app store](https://www.reddit.com/r/androiddev/comments/gm0zse/someone_modified_my_app_and_uploaded_it_to_some/)
- url: https://www.reddit.com/r/androiddev/comments/gm0zse/someone_modified_my_app_and_uploaded_it_to_some/
---
Hi all. Few days ago I was looking in Google Analytics at the most frequently viewed activities in my app. I noticed that there are many views of "TencentActivity". I'm sure that there's no class with that name in my app. I looked further and it seems that it's used by users in Thailand and Malaysia on Oppo devices. It also uses an app version that I didn't set. I'm curious about this modification. Did anyone had similar case?

I tried downloading an apk of the Oppo app store, but it's all in Chinese on my phone, and I couldn't find my app there (probably the name is changed).
