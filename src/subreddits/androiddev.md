# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - April 06, 2020](https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/
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
## [3][Just found out that i can drag expressions into watches list in Android Studio .. very handy!](https://www.reddit.com/r/androiddev/comments/fwzhra/just_found_out_that_i_can_drag_expressions_into/)
- url: https://i.redd.it/9ld40zbzqir41.gif
---

## [4][Android R, Package Visibility, and Some Holes](https://www.reddit.com/r/androiddev/comments/fwxdbw/android_r_package_visibility_and_some_holes/)
- url: https://commonsware.com/blog/2020/04/05/android-r-package-visibility-holes.html
---

## [5][Android Studio 4.0 Beta 4 available](https://www.reddit.com/r/androiddev/comments/fx0t01/android_studio_40_beta_4_available/)
- url: https://androidstudio.googleblog.com/2020/04/android-studio-40-beta-4-available.html
---

## [6][What is current status of Kotlin Native ?](https://www.reddit.com/r/androiddev/comments/fx530w/what_is_current_status_of_kotlin_native/)
- url: https://www.reddit.com/r/androiddev/comments/fx530w/what_is_current_status_of_kotlin_native/
---
As far as I got in this topic I have found these answers 

**Kotlin Native** - sharing code between platforms - already used in prod by companies in following article : 

https://medium.com/@touchlab/top-8-mobile-apps-in-2020-built-with-kotlin-multiplatform-3e9fea10e2af

**SQLDelight** -&gt; Used by CashApp in production -&gt; Can someone confirm this ?

**Multithreading** -&gt; NOT READY for production, many people that have tried KN on iOS complained about unstable memory model and multithreading on iOS - can someone explain / elaborate on current status of this ?

**Coroutines** -&gt; I have not found much about this topic

**KTOR** -&gt; Not used by anyone , e.g. VMWare engineer in this article states, that at VMWare they are not using multiplatfrom http client, they only share logic that generated requests and requests are being executed using HTTP client separatedly on each platform.

Is Kotlin Native production ready in 04/2020 ?

Any success / failure stories from using Kotlin Native in prod ? 

Thanks.
## [7][Scrollable page?](https://www.reddit.com/r/androiddev/comments/fx3ypc/scrollable_page/)
- url: https://www.reddit.com/r/androiddev/comments/fx3ypc/scrollable_page/
---
Hey, guys! I am in college for Software App Development. I just finished my first class (Basically an intro to Android Studio) and now I am taking Graphics and Interface Design. Long story short, we have a client that needs a website, but during our class interview with her, the client also stated she would like an app. No one in my class has any experience with apps besides me, so my professor asked if I would like to build an app for the client instead of a webpage. Since I want to eventually pursue a job with Android Studio, I agreed.

So, my question; how do I  makes my page be able to scroll? I know there is scrollview, but I can't seem to get it to work. I drag it over to my .xml design, but I cant add anything over it? I don't know. Sorry if this sounds dumb, I only have a month of experience.

I will post my code when I get home later today.

Thanks
## [8][What do you call apps that are like a google site where you can put info on it and people can look through ur info?](https://www.reddit.com/r/androiddev/comments/fx6czg/what_do_you_call_apps_that_are_like_a_google_site/)
- url: https://www.reddit.com/r/androiddev/comments/fx6czg/what_do_you_call_apps_that_are_like_a_google_site/
---
Kinda like those tourist apps where you can search for attractions or Airbnb where you can search places for countries?
## [9][Clearable edit text inheritance hell, is there a better way? (Kotlin)](https://www.reddit.com/r/androiddev/comments/fx1p96/clearable_edit_text_inheritance_hell_is_there_a/)
- url: https://www.reddit.com/r/androiddev/comments/fx1p96/clearable_edit_text_inheritance_hell_is_there_a/
---
My company's app wants to transition to using edit texts with a clear icon, similar to what is detailed in [this blog](https://medium.com/@dimabatyuk/adding-clear-button-to-edittext-9655e9dbb721)

The problem is that while that blog favors a composition strategy (adding an extension function to enable clearable functionality), it requires adding code to all screens and manually calling it on every text field.

Additionally, it automatically sets an OnTouchListener and OnFocusChangeListener, which might clash with other listeners that are used.

So far the only other solution I've found is to create new classes like ClearableEditText overriding onTouchEvent and onFocusChanged (this way I can use search-and-replace in XMLs), but this introduces yet another problem, because the app uses three kinds of EditTexts (AppCompatEditText, TextInputEditText, AppCompatAutoCompleteTextView) that all want this functionality, requiring me to triplicate this code. Is there a better way I can do this?
## [10][Long review times (Google Play)](https://www.reddit.com/r/androiddev/comments/fx332t/long_review_times_google_play/)
- url: https://www.reddit.com/r/androiddev/comments/fx332t/long_review_times_google_play/
---
Hi,  
Our update has been in review for 14 days. We understand that they have mentioned a longer review time of 7 days or more, but there's not much information more than that. I've not received any response from Google so it's hard to know if we're doing something wrong or if we should just wait.  


Is there anyone else here who is experiencing anything similar?  


Thanks!
## [11][I just finished my first full-size MVVM app - Upcoming Games. Can I get a code review?](https://www.reddit.com/r/androiddev/comments/fwissh/i_just_finished_my_first_fullsize_mvvm_app/)
- url: https://www.reddit.com/r/androiddev/comments/fwissh/i_just_finished_my_first_fullsize_mvvm_app/
---
[https://github.com/gavingt/upcoming-games](https://github.com/gavingt/upcoming-games)

[https://youtu.be/RS\_684kyOEk](https://youtu.be/RS_684kyOEk)

The app lets you view release dates and other info for every video game ever made.  It allows you to easily track your favorite and most anticipated games. Contains full sorting and search capabilities. The database is updated periodically using WorkManager.

This is my first time using several modern Android libraries, so I'm sure I've made some mistakes along the way. Any feedback would be appreciated.

Also, what more do I need to learn to become hireable?

Thanks!
## [12][App update stuck from more than 20 days. Are you in a similiar situation?](https://www.reddit.com/r/androiddev/comments/fx5766/app_update_stuck_from_more_than_20_days_are_you/)
- url: https://www.reddit.com/r/androiddev/comments/fx5766/app_update_stuck_from_more_than_20_days_are_you/
---
Hi,

I sent an update to the store listing (only description) 20 days ago, and the update is still not available for download. Only the store listing was updated (I edited the store description only).

History: 23 days ago I sent an update o the store listing of my app. After 3 days, seeing that the update was not released yet, I just sent it again doing a little edit to the description. Now, 20 days later, the update is still not available.

I am pretty sure that the process is stuck for some reasons, as 20 days for only descriptions is really a lot.

I wrote to the support about 10 days ago, and they told me to just wait. About 4 days ago I wrote to the support again and they didn't respond.

Context: the app is a wallpaper editor app that has never received suspensions/removal (published without issues from years).

Are you in a similiar situation? I just don't know what to do. Thanks in advance
