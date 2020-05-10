# androiddev
## [1][App Feedback Thread - May 09, 2020](https://www.reddit.com/r/androiddev/comments/ggemd1/app_feedback_thread_may_09_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ggemd1/app_feedback_thread_may_09_2020/
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
## [2][Weekly Questions Thread - May 04, 2020](https://www.reddit.com/r/androiddev/comments/gd9hik/weekly_questions_thread_may_04_2020/)
- url: https://www.reddit.com/r/androiddev/comments/gd9hik/weekly_questions_thread_may_04_2020/
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
## [3][An open source minimalistic notes app (Link in comments)](https://www.reddit.com/r/androiddev/comments/ggz0ul/an_open_source_minimalistic_notes_app_link_in/)
- url: https://i.redd.it/7tk00j8sywx41.png
---

## [4][I published a sample Android App repository with MVVM, Clean Architecture, Modularization, Repository Pattern](https://www.reddit.com/r/androiddev/comments/gglb2q/i_published_a_sample_android_app_repository_with/)
- url: https://www.reddit.com/r/androiddev/comments/gglb2q/i_published_a_sample_android_app_repository_with/
---
I published a repository might be a starting point for building Android interview tasks.

[https://github.com/ibrahimsn98/web-dev-tools-android](https://github.com/ibrahimsn98/web-dev-tools-android)

**Tech Stack**

    - Kotlin
    - MVVM
    - Modularization
    - Repository Pattern
    - Coroutines
    - Koin
    - Architecture Components
    - Navigation Component
    - LiveData
    - Paging
    - Kotlin DSL 
    - OkHttp3

&amp;#x200B;

https://preview.redd.it/in6yqnowesx41.jpg?width=1076&amp;format=pjpg&amp;auto=webp&amp;s=b3eae77a0463df97f5a3148fa246b6250677ca34
## [5][Australian government publishes source code of COVIDSafe app on Github](https://www.reddit.com/r/androiddev/comments/ggznh5/australian_government_publishes_source_code_of/)
- url: https://github.com/AU-COVIDSafe/mobile-android
---

## [6][Need advice - Correct Architecture for Bottom Nav?](https://www.reddit.com/r/androiddev/comments/ggzw60/need_advice_correct_architecture_for_bottom_nav/)
- url: https://www.reddit.com/r/androiddev/comments/ggzw60/need_advice_correct_architecture_for_bottom_nav/
---
I want to make a music app, very simple, using ExoPlayer. It will have a bottom nav with "songs" as the landing page, a recycler view of all songs in the library, with a player at the bottom of the view. Other areas of the app will be "albums", "artists" and "playlists". These again will be recycler views, but you click on something in these recycler views, and instead of playing, like in "songs", you will be taken to a recycler view of songs contained within that item i.e. I select Artist "Daft Punk", I get a list of Daft Punk songs with the player at the bottom, I click a song in that list and it plays.

So my decision is how to architect this and as I am new to android I honestly don't know the right way. I have done some reading and thing I am right in saying that I would have 1 activity (the nav bar activity). From this, I would have 4 fragments, one for each of the items in the bottom nav. I would then pass a list of all songs to the songs fragment on opening the app, or tapping the "Songs" nav item. But I would pass a different list of songs to the "Songs" fragment if I were navigating through one of the other areas.

Is this the correct approach? I haven't worked much with fragments so am a little wary that this might be a bit dumb... The alternative I though of was to have one child fragment each off of artists, playlists and albums which are essentially a copy of the songs fragment, but this seems like too much duplication. A little guidance here would be appreciated.

I have drawn something simple here to explain what I mean.

https://preview.redd.it/crezrdedbxx41.png?width=1430&amp;format=png&amp;auto=webp&amp;s=095129dc7dafd244a367e8a5e7beeeb03d4bc863
## [7][Admob Banner placement policy on scrolling list?](https://www.reddit.com/r/androiddev/comments/ggyj4a/admob_banner_placement_policy_on_scrolling_list/)
- url: https://www.reddit.com/r/androiddev/comments/ggyj4a/admob_banner_placement_policy_on_scrolling_list/
---
Hi,

Came across this confusing admob banner placement policy over scrolling content saying it should have a clear border between the banner and scrolling content.

&gt; Banner ads that remain on the screen while a user scrolls through other content may be placed at the top or the bottom of the screen, as long as there are appropriate borders and/or sufficient separation from the app content, especially from navigational elements such as menu bars   
&gt;  
&gt;src:  [https://support.google.com/admob/answer/6275335?hl=en&amp;ref\_topic=2936214](https://support.google.com/admob/answer/6275335?hl=en&amp;ref_topic=2936214) 

 This seems like an awful UI experience for the user. Does anyone know any popular apps with banner at the bottom with scrolling content? Are they really forcing this policy?

Seems like its bit unnecessary since, while scrolling, user usually performs action in the middle of the screen and there is no risk of accidental clicks at the bottom of the page? Understand that it might be important to clearly separate ad content from app content in that case but are there any apps really adhering to that policy? Dont think I have ever seen an app providing extra border between recycler view and bottom banner...

&amp;#x200B;

https://preview.redd.it/jntjboxypwx41.png?width=250&amp;format=png&amp;auto=webp&amp;s=911f03c8fba62cfbce940c5b0fd21dfbe2e0c395
## [8][[GooglePlay Console] Would my approach fall into "Repetitive Content" policy by google.](https://www.reddit.com/r/androiddev/comments/ggz9st/googleplay_console_would_my_approach_fall_into/)
- url: https://www.reddit.com/r/androiddev/comments/ggz9st/googleplay_console_would_my_approach_fall_into/
---
I build an app which provide two main features, let call them FeatureA and FeatureB. I plan to publish three app from the same set of source code:  
\- App A -&gt; Only provide FeatureA and it's free

\- App B -&gt; Only provide FeatureB and it's free

\- App C -&gt; Provide FeatureA &amp; FeatureB, and it's a paid version.   
My question is, would this approach fall into "Repetitive Content" policy from Google.

[https://play.google.com/about/spam-min-functionality/#!?zippy\_activeEl=repetitive-content#repetitive-content](https://play.google.com/about/spam-min-functionality/#!?zippy_activeEl=repetitive-content#repetitive-content)

Anyone have experience, please help to advice. 

Thanks.
## [9][RenderScript](https://www.reddit.com/r/androiddev/comments/gh0u95/renderscript/)
- url: https://www.reddit.com/r/androiddev/comments/gh0u95/renderscript/
---
I understand that RenderScript is a way to speed up SIMD operations by parallelizing computation across multiple CPU and GPU cores and there is also JNI which allows you to accelerate code by running it in C instead. 

My question is do you use it in production applications and if so what's been your experience with it and if you have ever worked on JNI and how these two approaches are different in your experience
## [10][How was your experience using apps like appEnhancer, and is Google really okay with them?](https://www.reddit.com/r/androiddev/comments/gh0cfu/how_was_your_experience_using_apps_like/)
- url: https://www.reddit.com/r/androiddev/comments/gh0cfu/how_was_your_experience_using_apps_like/
---
I'm talking about apps, which for every review you give to other random apps, your apps also receive a review from a random developer. I've been using it for about half an hour and gave three opinions and already received two. However it's apparent that people don't even bother playing your games and just give them five stars, but does it help Google to reccomend my games?
## [11][How to have fragment add behaviour using Navigation Components? The fragment is always recreated when visiting back to it, which leads to re-query of data and re-rendering of data.](https://www.reddit.com/r/androiddev/comments/ggym39/how_to_have_fragment_add_behaviour_using/)
- url: https://www.reddit.com/r/androiddev/comments/ggym39/how_to_have_fragment_add_behaviour_using/
---
Does navigation components currently only support fragment replacement? Due to this whole screen is rebuild every time fragment is revisited again rather than just the 'onResume' being called in activity scenario.
## [12][I have made an open source PDF editor that works cooperatively like Google Docs (Github in comments)](https://www.reddit.com/r/androiddev/comments/ggdu3m/i_have_made_an_open_source_pdf_editor_that_works/)
- url: https://v.redd.it/dwt0aydy4qx41
---

