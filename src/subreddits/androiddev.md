# androiddev
## [1][App Feedback Thread - October 31, 2020](https://www.reddit.com/r/androiddev/comments/jlhjcn/app_feedback_thread_october_31_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jlhjcn/app_feedback_thread_october_31_2020/
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
## [2][Weekly Questions Thread - October 26, 2020](https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/
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
## [3][Looking for a code review for a small todo-list app using Flows, Channels, DataStore, Hilt, Room, etc.](https://www.reddit.com/r/androiddev/comments/jlgi1h/looking_for_a_code_review_for_a_small_todolist/)
- url: https://www.reddit.com/r/androiddev/comments/jlgi1h/looking_for_a_code_review_for_a_small_todolist/
---
My focus in this small project was to get MVVM right and separate the logic between the fragment and the ViewModel properly. In the past, I've been writing a lot of logic in my fragments and I'm trying to learn how to do that better.

I use Channels converted to Flow to emit events from my ViewModels for my fragments to listen to.

Can someone take a look at this project and give some feedback? Any big mistakes?

[https://github.com/codinginflow/ToDoPrep](https://github.com/codinginflow/ToDoPrep)
## [4][AndroidBites | Java ☕️ Maps 🗺 on the Kotlin.](https://www.reddit.com/r/androiddev/comments/jle5oc/androidbites_java_maps_on_the_kotlin/)
- url: https://chetan-garg36.medium.com/java-%EF%B8%8F-maps-on-the-kotlin-8930b9f55d8d
---

## [5][How do I show a toast just below the Toolbar](https://www.reddit.com/r/androiddev/comments/jl5wcl/how_do_i_show_a_toast_just_below_the_toolbar/)
- url: https://www.reddit.com/r/androiddev/comments/jl5wcl/how_do_i_show_a_toast_just_below_the_toolbar/
---
I need to show a toast just below the toolbar. It is a standard activity with no frills

&amp;#x200B;

&gt;override fun onCreateView(savedInstance: Bundle?) {  
&gt;  
&gt;super.onCreateView(savedBundle)  
&gt;  
&gt;setContentView(R.layout.my\_layout)  
&gt;  
&gt;val toast = Toast(this).apply {  
&gt;  
&gt;setGravity(Gravity.***TOP*** | Gravity.***CENTER\_HORIZONTAL***, 0, 0)setDuration(Toast.***LENGTH\_LONG***)  
&gt;  
&gt;show(R.string.success\_string)  
&gt;  
&gt;}  
&gt;  
&gt;. . . . . . . Other code in the function . . . .  
&gt;  
&gt;}

&amp;#x200B;

Activity definition in the manifest

&amp;#x200B;

&gt;&lt;**activity android:name=".activities.InputActivity"** /&gt;

&amp;#x200B;

I have  `setGravity(Gravity.TOP | Gravity.CENTER_HORIZONTAL, 0, 0)`. I know that I can set the y offset but I don't know how the y offset should be calculated. The most relevant article on SO is from 2013. Any suggestions would be gratefully appreciated.

[https://stackoverflow.com/questions/18420445/how-to-show-custom-toast-below-actionbar](https://stackoverflow.com/questions/18420445/how-to-show-custom-toast-below-actionbar)
## [6][How to fix app:DexBuilderDebug build failed error?](https://www.reddit.com/r/androiddev/comments/jlgryb/how_to_fix_appdexbuilderdebug_build_failed_error/)
- url: https://www.reddit.com/r/androiddev/comments/jlgryb/how_to_fix_appdexbuilderdebug_build_failed_error/
---
This is my second time building an app with android studio. I’m trying to build a currency conversion app that calls an API to do the conversion. However, when I run the app, I get the dexBuilderDebug error. I’m only getting this error after creating the interface and class for retrofit. 

I have my code and errors posted on [stackoverflow ](https://stackoverflow.com/questions/64620823/android-studio-build-failed-task-appdexbuilderdebug-failed). Has anyone ran into the same error before? How did you fix it? I can’t even understand what is causing the error
## [7][How can I do custom notch on CardView like this?](https://www.reddit.com/r/androiddev/comments/jksyrl/how_can_i_do_custom_notch_on_cardview_like_this/)
- url: https://i.redd.it/uncfvma607w51.png
---

## [8][New android dev, what do I need to know to do C/C++?](https://www.reddit.com/r/androiddev/comments/jlc1tm/new_android_dev_what_do_i_need_to_know_to_do_cc/)
- url: https://www.reddit.com/r/androiddev/comments/jlc1tm/new_android_dev_what_do_i_need_to_know_to_do_cc/
---
Hi, I'm totally new to Android dev. I'm hoping someone can give me a brief roadmap on what I'm going to need to know to work effectively with Android Studio and C++ so I can start reading the right articles and so on. For example, do I need to know "normal" Java android-development for the NDK to make sense? Are there decent resources for learning C++ oriented Android dev that aren't out of date? Are all NDK apps hybrids where at some point I'm dealing with the JNI?

I'm mucking about with JUCE but I'm a big believer in understanding what your framework is doing and not developing more married to the framework than need be, so I feel I should be educating myself outside of just JUCE oriented resources.

Thanks!
## [9][Avoiding I/O in gradle ?](https://www.reddit.com/r/androiddev/comments/jkyr0v/avoiding_io_in_gradle/)
- url: https://www.reddit.com/r/androiddev/comments/jkyr0v/avoiding_io_in_gradle/
---
I have seen some threads recommending that you avoid I/O during gradle configuration. My company currently does this to get the branch name, build time, and git SHA, into BuildConfig. I'm trying to understand the reasoning of why I should avoid it before I bring it up to my coworkers.

For these specific examples I replaced the git command with a static string, and did `./gradlew assemble* --dry-run --rerun-tasks --profile` and those configuration times dropped from around 3s to 1.5s. that's a nice percentage but not all that significant since that's only for configuration.

What else should I know? I have read some threads alluding to impacting incremental builds but can't find concrete info about why
## [10][Personal or business gmail address for developer account](https://www.reddit.com/r/androiddev/comments/jlboit/personal_or_business_gmail_address_for_developer/)
- url: https://www.reddit.com/r/androiddev/comments/jlboit/personal_or_business_gmail_address_for_developer/
---
I plan on releasing a paid app to the play store soon. I don’t really expect it to make me much more than beer money, so I’m not incorporating as an LLC or anything. 

I will be getting a separate gmail account though for the developer account. When I sign up for a new gmail address should it be a personal one since I’m just doing business as a sole proprietor (US based) or a business gmail account since I am selling something, which is a business transaction for the little it’ll be worth?
## [11][Is there a known problem with chrome, and samsung galaxy phones having broken grayscale/black levels?](https://www.reddit.com/r/androiddev/comments/jl702w/is_there_a_known_problem_with_chrome_and_samsung/)
- url: https://www.reddit.com/r/androiddev/comments/jl702w/is_there_a_known_problem_with_chrome_and_samsung/
---
Chrome seems to make almost every video that should have either black become light gray, or a value that should be 100% white becomes grey instead.
## [12][Mobile Developers Cafe Weekly Issue #12 🎉🎊 is out](https://www.reddit.com/r/androiddev/comments/jlcm71/mobile_developers_cafe_weekly_issue_12_is_out/)
- url: https://archive.mobiledeveloperscafe.com/p/issue12
---

