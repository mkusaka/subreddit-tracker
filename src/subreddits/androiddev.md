# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ektosq/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ektosq/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - January 06, 2020](https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/
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
## [3][🤯 Just realized Dagger’s naming double entendre...](https://www.reddit.com/r/androiddev/comments/em495b/just_realized_daggers_naming_double_entendre/)
- url: https://www.reddit.com/r/androiddev/comments/em495b/just_realized_daggers_naming_double_entendre/
---
Obviously the first one: daggers are able to pierce into things, thereby injecting dependencies. 

The second: DAG’s (directed acyclic graphs) are how dependency creation order is determined. The “Dagger” library is a DAG creator or “DAGger”...

Show of hands: how stupid am I? Did everyone already realize this?
## [4][How did Telegram implement this theme change effect ? I know I can look at the code but needed the gist of the implementation. It animates the status bar color change as well! I'm fairly impressed.](https://www.reddit.com/r/androiddev/comments/eltrvl/how_did_telegram_implement_this_theme_change/)
- url: https://i.redd.it/b68e2zi7kk941.jpg
---

## [5][Write your JSON in Kotlin DSL | Create your JSON in a more easy and readable way | (github.com/hitanshu-dhawan/JsonDSL)](https://www.reddit.com/r/androiddev/comments/em6x1a/write_your_json_in_kotlin_dsl_create_your_json_in/)
- url: https://i.redd.it/va9k5s1brp941.png
---

## [6][What is the requirement to be a "Senior" Android Developer?](https://www.reddit.com/r/androiddev/comments/em5vib/what_is_the_requirement_to_be_a_senior_android/)
- url: https://www.reddit.com/r/androiddev/comments/em5vib/what_is_the_requirement_to_be_a_senior_android/
---
I know "Senior" require some of these soft skills:

* Interpersonal Skill
* Analytical Skill
* Problem Solving Skill
* Able to lead juniors
* Able to be role model
* Decent Communication Skill
* Detail Oriented

But how about tech stack? Which tech skills to master to be more considered as "Senior"?

Some that I know: Unit Testing, understand UX, familiar with Android JetPack.
## [7][Anyone notice any increase in RPM after using app-ads.txt?](https://www.reddit.com/r/androiddev/comments/em75c1/anyone_notice_any_increase_in_rpm_after_using/)
- url: https://www.reddit.com/r/androiddev/comments/em75c1/anyone_notice_any_increase_in_rpm_after_using/
---

## [8][What the hell do I have to do to my notifications for my wearOS app to be accepted as a standalone app?!?!](https://www.reddit.com/r/androiddev/comments/em750k/what_the_hell_do_i_have_to_do_to_my_notifications/)
- url: https://www.reddit.com/r/androiddev/comments/em750k/what_the_hell_do_i_have_to_do_to_my_notifications/
---
some context, I have a simple app that fetches a list of activities for the current day from a server and notifies you when they're about to start.

on said notifications I have actions to both open the app to view activity details and to either start or stop activities as appropriate. it's set as quick actions and are displaying properly on every watch that I can get my hands on.

Yet I get a "your notification does not provide enhanced functions to Wear OS users." message when submiting my app for review

I EVEN REMOVED ALL NOTIFICATIONS FROM MY APP AND GOT THE EXACT SAME FUCKING MESSAGE (making it just a portal to view todays activities)!!!! WHAT THE FLYING FUCK DO YOU MEAN MY NOTIFICATIONS DON'T PROVIDE ENHANCED FUNCTIONS, I REMOVED ALL NOTIFICATIONS FROM THE FUCKING APP!!!!! AAAAAAAAARGH
## [9][LeakCanary leak deobfuscation feature explained](https://www.reddit.com/r/androiddev/comments/em912f/leakcanary_leak_deobfuscation_feature_explained/)
- url: https://polidea.com/blog/leakcanary-deobfuscation-feature-explained/
---

## [10][Review my MVI study repo please](https://www.reddit.com/r/androiddev/comments/em903j/review_my_mvi_study_repo_please/)
- url: https://www.reddit.com/r/androiddev/comments/em903j/review_my_mvi_study_repo_please/
---
Hey fellow android devs. I started looking at MVI recently with talks such as [this](https://www.youtube.com/watch?v=0IKHxjkgop4) or [this](https://www.youtube.com/watch?v=Ls0uKLqNFz4&amp;list=PLsAfcuwrBov7NJOcGxBN7FU0SQdid6L0j). After I digested the information I tried creating a small sample app with the MVI architecture and the result is this [repo](https://github.com/acristescu/mvitest). All the app does is show a paginated list of github repos and loads the next page when the user reaches the bottom of the list.

Could I please ask members of the community that have more experience with MVI for some constructive criticism on this implementation?
## [11][Accessibility text to speach.](https://www.reddit.com/r/androiddev/comments/em8m8v/accessibility_text_to_speach/)
- url: https://www.reddit.com/r/androiddev/comments/em8m8v/accessibility_text_to_speach/
---
I'm improving accessibility(TalkBack) on one of the screens and noticed that acronyms like API or HTML are not read like they should only on beloved Samsung device, others seems not to have issue with them. 

Any idea how this could be handled without additional string as a content description?
This would require to have additional string that have to be translated in multiple supported languages.
## [12][Lightweight PDF Library ?](https://www.reddit.com/r/androiddev/comments/em85ca/lightweight_pdf_library/)
- url: https://www.reddit.com/r/androiddev/comments/em85ca/lightweight_pdf_library/
---
Does anyone have any recommendations for a lightweight/low-size PDF viewer library to display PDF documents in-app ?

I've recently discovered my project's PDF library has a native dependency that's minimum 5mb per ABI (I use app bundles so it's not compounded for the total size, but still)
