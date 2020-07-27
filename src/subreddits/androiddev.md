# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/hyrl0d/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/hyrl0d/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - July 27, 2020](https://www.reddit.com/r/androiddev/comments/hyq58w/weekly_questions_thread_july_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hyq58w/weekly_questions_thread_july_27_2020/
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
## [3][What is the best (and hopefully simplest) resource for creating a privacy policy for a mobile app?](https://www.reddit.com/r/androiddev/comments/hym3ym/what_is_the_best_and_hopefully_simplest_resource/)
- url: https://www.reddit.com/r/androiddev/comments/hym3ym/what_is_the_best_and_hopefully_simplest_resource/
---
What is the best (and hopefully simplest) resource for creating a privacy policy for a mobile app?
## [4][Anyone received reviews like these that says the game/app is nice but vote 1 star? I'm just curious, are they bots or are they just being sarcastic? None of them reply when I tried replying to them. Flagging them or contacting support will not work as they didn't violate the comment policy](https://www.reddit.com/r/androiddev/comments/hyqcwl/anyone_received_reviews_like_these_that_says_the/)
- url: https://i.redd.it/cqvhf216udd51.jpg
---

## [5][For all who use Admob, don't target Android 11 (API 30) just yet](https://www.reddit.com/r/androiddev/comments/hyqap5/for_all_who_use_admob_dont_target_android_11_api/)
- url: https://www.reddit.com/r/androiddev/comments/hyqap5/for_all_who_use_admob_dont_target_android_11_api/
---
Why? Because it doesn't support it yet.

I've tried to target Android API 30, and tested on my Pixel 4 R beta 2.5, and I can't see native ads, can't see interstitial ads, and probably also banner ads (this I didn't test yet).

This is true at least for this version:

        implementation 'com.google.android.gms:play-services-ads:19.3.0'

You can try it for yourself if you wish. It will happen even with Admob's sample projects:

[https://github.com/googleads/googleads-mobile-android-examples/tree/master/kotlin/admob](https://github.com/googleads/googleads-mobile-android-examples/tree/master/kotlin/admob)

The error you will get when trying it out :

    0 - Internal error

I remember I saw similar issues on older versions of Android when they were on beta/preview, too.

 ¯\_(ツ)_/¯ 
## [6][Is there a reason that there's no native version of a collapsing/expandable layout?](https://www.reddit.com/r/androiddev/comments/hyp2ef/is_there_a_reason_that_theres_no_native_version/)
- url: https://www.reddit.com/r/androiddev/comments/hyp2ef/is_there_a_reason_that_theres_no_native_version/
---
And no, DrawerLayout and CollapsingToolbarLayout don't count, these are hard-wired to be window-related and toolbar-related (which is also window-related) respectively.

While I use [this version](https://github.com/cachapa/ExpandableLayout) and it works fine, it doesn't really makes any sense to not have a native control/view for this since it's such a trivial functionality...
## [7][How do I beat the '4+ years experience' requirement in job posting with building my own apps?](https://www.reddit.com/r/androiddev/comments/hyp7k0/how_do_i_beat_the_4_years_experience_requirement/)
- url: https://www.reddit.com/r/androiddev/comments/hyp7k0/how_do_i_beat_the_4_years_experience_requirement/
---
What sort of demo or published apps must I develop to beat this requirement? I am facing this situation where nobody wants to hire me if I don't have work experience, and the only way to have work experience is to work. 

I have a feeling that those tutorial apps I see online just won't cut it.
## [8][Optimizing ‘new byte[]’](https://www.reddit.com/r/androiddev/comments/hycbw4/optimizing_new_byte/)
- url: https://publicobject.com/2020/07/26/optimizing-new-byte/
---

## [9][AdMob vs Facebook Ads, which is better?](https://www.reddit.com/r/androiddev/comments/hys10z/admob_vs_facebook_ads_which_is_better/)
- url: https://www.reddit.com/r/androiddev/comments/hys10z/admob_vs_facebook_ads_which_is_better/
---
I know this has been asked a lot, but I've heard FB ads can only show ads to people that have facebook installed? Is this true? 

Also which do you think has the potential to earn more profit? And which network is more strict and tight? I don't want to go to a network that would get me banned for some stupid reasons.
## [10][Need help for Collage maker](https://www.reddit.com/r/androiddev/comments/hyrct8/need_help_for_collage_maker/)
- url: https://www.reddit.com/r/androiddev/comments/hyrct8/need_help_for_collage_maker/
---
I am developing an app which also includes a feature of collage maker, can someone help me and show how to do that or any reference for open source project.

As of now i am using [this code](https://github.com/Azoft/CollageMaker-Android)
## [11][Beta users not able to review production release?](https://www.reddit.com/r/androiddev/comments/hyqa16/beta_users_not_able_to_review_production_release/)
- url: https://www.reddit.com/r/androiddev/comments/hyqa16/beta_users_not_able_to_review_production_release/
---
Anybody know why beta users can't review the game on play store after it's fully released on production? I removed all testers from beta, but none of my users seem to be able to review? Is this a policy on google's side?
## [12][A way to call out PS4 controller buttons command using on-Screen buttons on Android](https://www.reddit.com/r/androiddev/comments/hyq6ze/a_way_to_call_out_ps4_controller_buttons_command/)
- url: https://www.reddit.com/r/androiddev/comments/hyq6ze/a_way_to_call_out_ps4_controller_buttons_command/
---
Sorry if I post it in wrong sub-reddit, but to be honest Idon't where else to post this. Soas title says I'm trying to find out is it posiible and if how to call out DUALSHOCK4 button command with on-screen button. For example I connect PS4 conntroller to my Android and press a one button on a controller (may be any △ or O )and then controller send specific command to device and device act it out how it can,so would it be possible to have a on screen button that when pressed device would think that same specific controller button command was send and would react the same way? The reason for asking is that there is a game that has specific command it can do but there's is no on screen button for that and it can only be called out with PS4 controller buttons I don't wanna pay tons of cash for a genuine PS4 controller just for two button on it.I you know where this kind of question would belong in another sub-reddit please link it in reply.
