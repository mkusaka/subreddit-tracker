# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/huklpf/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/huklpf/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - July 20, 2020](https://www.reddit.com/r/androiddev/comments/huj9uy/weekly_questions_thread_july_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/huj9uy/weekly_questions_thread_july_20_2020/
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
## [3][The greatest versionCode that is allowed by Google Play is 2100000000. Does that mean once uploading with that versionCode will lead to no possibility to update the app afterwards?](https://www.reddit.com/r/androiddev/comments/hur9m3/the_greatest_versioncode_that_is_allowed_by/)
- url: https://www.reddit.com/r/androiddev/comments/hur9m3/the_greatest_versioncode_that_is_allowed_by/
---
Not that it's likely or a problem I'm encountering, just wondering what the options would be.
## [4][Multiple Backstack Management using Navigation Component](https://www.reddit.com/r/androiddev/comments/hv1i76/multiple_backstack_management_using_navigation/)
- url: https://github.com/vipulasri/Multiple-Backstack-Navigation
---

## [5][Colors with opacity performance](https://www.reddit.com/r/androiddev/comments/hv4oys/colors_with_opacity_performance/)
- url: https://www.reddit.com/r/androiddev/comments/hv4oys/colors_with_opacity_performance/
---
Hello, I was wondering whether it is more efficient to have a color with opacity built in for example #80FFFFFF which is white with 50% transparency, or have it be #FFFFFF and then add alpha value on the view 0.5f, either on XML or through code. Are there any differences that you know of? Or are they basically the same?
Thank you!
## [6][Getting on the same page with Paging 3](https://www.reddit.com/r/androiddev/comments/hv7cyg/getting_on_the_same_page_with_paging_3/)
- url: https://android-developers.googleblog.com/2020/07/getting-on-same-page-with-paging-3.html
---

## [7][Hilt and Dagger annotations cheat sheet](https://www.reddit.com/r/androiddev/comments/huo9jp/hilt_and_dagger_annotations_cheat_sheet/)
- url: https://medium.com/androiddevelopers/hilt-and-dagger-annotations-cheat-sheet-9adea070e495
---

## [8][How do I merge/aggregate unit test reports from an multi-module android project?](https://www.reddit.com/r/androiddev/comments/hv5qo1/how_do_i_mergeaggregate_unit_test_reports_from_an/)
- url: https://www.reddit.com/r/androiddev/comments/hv5qo1/how_do_i_mergeaggregate_unit_test_reports_from_an/
---

## [9][What do I need to do to implement a system similar to a smart home, but for storages?](https://www.reddit.com/r/androiddev/comments/hv33q7/what_do_i_need_to_do_to_implement_a_system/)
- url: https://www.reddit.com/r/androiddev/comments/hv33q7/what_do_i_need_to_do_to_implement_a_system/
---
I have a main storage on which an android tablet is installed as a control center.  and there are the same storage modules, only without screens.  you need to implement their connection to the main control center.  what technologies can i use to achieve the goal?
## [10][I'm confused about 'Maps API for Android' API Key!](https://www.reddit.com/r/androiddev/comments/hv4new/im_confused_about_maps_api_for_android_api_key/)
- url: https://www.reddit.com/r/androiddev/comments/hv4new/im_confused_about_maps_api_for_android_api_key/
---
Everywhere it says that you need a billing account to get an API Key. Even though you never cross the free quota billing account is a must to start using the API. That's what I've read in the official documentations and StackOverflow.

Now the thing is, yesterday I tried to get an API Key without a billing account and I was able get one. I enabled the Maps API for Android and added the API Key to an app. Tried to load the map. It worked.

My question is – how is this possible? What's happening? 

I'm really confused. Can anyone clear me up?
## [11][Does DI code belong to modules?](https://www.reddit.com/r/androiddev/comments/hv0c93/does_di_code_belong_to_modules/)
- url: https://www.reddit.com/r/androiddev/comments/hv0c93/does_di_code_belong_to_modules/
---
I'm working on a multi app modular codebase. Its natural to put dagger/koin/whatever DI code modules inside given gradle modules. Also, Hephaestus and Hilt sort of enable you to do that easier.

But, should you?

Granted, module owner knows best about scopes of partucilar dependencies

However, with dagger, you need to slap explicit scope annotation, effectively hardcoding it to that scope. You cannot just say "Single instance per whatever component I'm in". This limits reuse -- what if you want to use the module in a different scope per app it is in

Also, retrofit doesnt provide no dagger modules

Also, multiplatform?

So, this leads me to believe DI setup should maybe only live in the top level :app module

However if it only lives in app, then it needs to be duplicated per app, as it cannot be shared. Which is a bummer

Also, not sure how would inter module communication would work if component interfaces were not available. This is necessary, until we get ui constructor back (compose)

I think what I really want is for dagger to implement the mentioned @SinglePerWhateverComponentThisIsIn 🤔

Thoughts?
## [12][JetBrains is porting Compose to desktop and other platforms](https://www.reddit.com/r/androiddev/comments/hufnjq/jetbrains_is_porting_compose_to_desktop_and_other/)
- url: https://www.reddit.com/r/androiddev/comments/hufnjq/jetbrains_is_porting_compose_to_desktop_and_other/
---
It's no news that JetBrains employees are working on Compose. But this job offer: https://www.jetbrains.com/careers/jobs/ui-framework-developer-401/ makes it clear what the goal is - "Our team is working to port it to desktop and other platforms.".
