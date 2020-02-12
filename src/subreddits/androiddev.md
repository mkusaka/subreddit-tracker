# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/f1psjk/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/f1psjk/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 10, 2020](https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/
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
## [3][Using Gesture Recognition AI (MiVRy) to detect drawing characters in the space.](https://www.reddit.com/r/androiddev/comments/f2mi56/using_gesture_recognition_ai_mivry_to_detect/)
- url: https://v.redd.it/bwpfmwaolfg41
---

## [4][Currently developing a community Android ad network and looking to partner up with other app developers](https://www.reddit.com/r/androiddev/comments/f2engl/currently_developing_a_community_android_ad/)
- url: https://www.reddit.com/r/androiddev/comments/f2engl/currently_developing_a_community_android_ad/
---
Evening all,

I've been working on a community ad network for Android. The basic gist is that you show banner ads for other apps in the network in exchange for being shown in other apps.

I'm trying to solve the exposure on a zero budget issue and also show something when other ads fail to fill.

Currently I have a [very simple SDK](https://github.com/laurencedawson/AdColliderSdk) live that's showing ads in Sync for reddit. 

Would really love to work with other developers on this.

Cheers,

Laurence
## [5][Why Rewarded Ads CTR is so Low? Is this Normal?](https://www.reddit.com/r/androiddev/comments/f2nx15/why_rewarded_ads_ctr_is_so_low_is_this_normal/)
- url: https://i.redd.it/jib1jqeg9gg41.png
---

## [6][Guide to host app-ads.txt on Github Pages (fast and free)](https://www.reddit.com/r/androiddev/comments/f2nji4/guide_to_host_appadstxt_on_github_pages_fast_and/)
- url: https://www.reddit.com/r/androiddev/comments/f2nji4/guide_to_host_appadstxt_on_github_pages_fast_and/
---
// NOTE: Because the [old guide](https://www.reddit.com/r/androiddev/comments/cp5n2m/the_developerfriendly_guide_to_appsadstxt_admobs/) says nothing about the Github Pages I decided to make my guide. Based on my [previous post](https://www.reddit.com/r/androiddev/comments/f15zty/how_to_host_appadstxt_on_github_pages/).

1. Register your personal page on Github Pages: [https://pages.github.com/](https://pages.github.com/) (TL;DR - just create a repository with name "your-github-username.github.io" and push everything only into the master branch)
2. Create your app-ads.txt.  
Long official version: [https://support.google.com/admob/answer/9363762](https://support.google.com/admob/answer/9363762)  
Short: [https://apps.admob.com/v2/apps/appadstxt](https://apps.admob.com/v2/apps/appadstxt) (press the "How to set up app-ads.txt" button at the top right)
3. Put app-ads.txt in the root of your Github website repository into master branch.Check availability by link [https://your-github-username.github.io/app-ads.txt](https://your-github-username.github.io/app-ads.txt) (update may take a couple of minutes).  
NOTE: Your repository should be public. If you switched it to private and revert back - try to push any commit. There is such a bug on Github Pages right now: [Github pages isn't working after switching from public to private and then back to public](https://github.community/t5/GitHub-Pages/Github-pages-isn-t-working-after-switching-from-public-to/td-p/16979) (please write to Github team about it, maybe then they will fix it faster)
4. Set your website url ([https://your-github-username.github.io/](https://your-github-username.github.io/)) on the page of each app page (App control panel -&gt; App page -&gt; App description -&gt; Website).  
NOTE: The website specified in the main settings in Google Dev Console (Main page with apps list -&gt; Settings -&gt; Developer account -&gt; Account details -&gt; Website) is not used for app-ads.txt detection.  
NOTE: Make sure you set your website url with https.
5. \[Optional\] Recheck the specified website. After you receive a notification about updating the app page go to the page of your app in Google Play ([https://play.google.com/store/apps/details?id=your.app.id](https://play.google.com/store/apps/details?id=your.app.id)) and view the html code. You should see the meta tag with name "appstore:developer\_url" and "content" with url of your site. These specifications are described in [https://iabtechlab.com/wp-content/uploads/2019/03/app-ads.txt-v1.0-final-.pdf](https://iabtechlab.com/wp-content/uploads/2019/03/app-ads.txt-v1.0-final-.pdf) the link to which is indicated in the official help [https://support.google.com/admob/answer/9363762](https://support.google.com/admob/answer/9363762)
6. After a few days (five in my case) you should see a successful status in [https://apps.admob.com/v2/apps/appadstxt](https://apps.admob.com/v2/apps/appadstxt)

You are breathtaking!
## [7][2 suspended apps, what to do next?](https://www.reddit.com/r/androiddev/comments/f2plwj/2_suspended_apps_what_to_do_next/)
- url: https://www.reddit.com/r/androiddev/comments/f2plwj/2_suspended_apps_what_to_do_next/
---
I have a google developer account with 2 suspended apps (got both suspensions in my first month (and same day) \~2015..), so i assume that a third suspension means my account termination.  
I plan to upload some real games, but I'm afraid making a mistake leading to my third suspension and hence banned for life.. is it possible to make a second account (with same admob and same credit card?) and upload some new apps? or the strikes count will be counted as a "3rd strike" and ban my accounts?
## [8][How we fought bad apps and malicious developers in 2019](https://www.reddit.com/r/androiddev/comments/f2m9xh/how_we_fought_bad_apps_and_malicious_developers/)
- url: https://android-developers.googleblog.com/2020/02/how-we-fought-bad-apps-and-malicious.html
---

## [9][Is there a way to test the Google Play Install Referrer?](https://www.reddit.com/r/androiddev/comments/f2p48s/is_there_a_way_to_test_the_google_play_install/)
- url: https://www.reddit.com/r/androiddev/comments/f2p48s/is_there_a_way_to_test_the_google_play_install/
---
I recently was trying out the "Play Install Referrer Library" (the latest one), but i couldn't find a way to test my code if it works or not, any known way to test this library?

For reference https://developer.android.com/google/play/installreferrer/library.html
## [10][EventBus 3.2 release: AndroidX and incremental annotation processing](https://www.reddit.com/r/androiddev/comments/f2r3x6/eventbus_32_release_androidx_and_incremental/)
- url: https://greenrobot.org/release/eventbus-3-2/
---

## [11][Mutli-Selection Recycler View](https://www.reddit.com/r/androiddev/comments/f2oa45/mutliselection_recycler_view/)
- url: https://www.reddit.com/r/androiddev/comments/f2oa45/mutliselection_recycler_view/
---
What are the best common ways to select multiple items in a Recycler View, I tried the androidx selection library. But I'd like to hear from you guys, cause I found there are many ways to do it and I don't know which way to go or should I just implement my own version and what works best for my app?

&amp;#x200B;

Edit - Typo
## [12][getting into android](https://www.reddit.com/r/androiddev/comments/f2q7zt/getting_into_android/)
- url: https://www.reddit.com/r/androiddev/comments/f2q7zt/getting_into_android/
---
hi, im a comp sci grad with some experience in programming, just wondering how long it would take to self teach myself android development, to an industry standard, also if someone can reference some things to help will be great!
