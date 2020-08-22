# androiddev
## [1][App Feedback Thread - August 22, 2020](https://www.reddit.com/r/androiddev/comments/iehl0b/app_feedback_thread_august_22_2020/)
- url: https://www.reddit.com/r/androiddev/comments/iehl0b/app_feedback_thread_august_22_2020/
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
## [2][Weekly Questions Thread - August 17, 2020](https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ibbyvc/weekly_questions_thread_august_17_2020/
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
## [3][ConstraintLayout 2.0.0](https://www.reddit.com/r/androiddev/comments/ie9auq/constraintlayout_200/)
- url: https://androidstudio.googleblog.com/2020/08/constraintlayout-200.html
---

## [4]["Rate this app" functionality using In-App Review API](https://www.reddit.com/r/androiddev/comments/iefq78/rate_this_app_functionality_using_inapp_review_api/)
- url: https://www.rockandnull.com/rate-this-app-android/
---

## [5][Performance loss using compound components?](https://www.reddit.com/r/androiddev/comments/ieg2pg/performance_loss_using_compound_components/)
- url: https://www.reddit.com/r/androiddev/comments/ieg2pg/performance_loss_using_compound_components/
---
Hey, so some time ago I discovered that we could create compound components by extending some kind of a ViewGroup and inflating an layout into it. After some time I noticed that I should use the merge tag as the top most tag in that layouts to prevent unnecessary nesting.

Recently I saw a post that stated that nesting ViewGroups is very bad for performance and we should keep our view hierarchy flat. I 100% agreed with that.
 
Finally my problem: I got in love with compound views as they make some components reusable (not only the layout but also some logic that is coupled to them - that's why I don't consider that include tag as satisfying as compound components) and reduce complexity as layout files get smaller and so on. But with every compound component you add to your layout, you add an ViewGroup that is nested. Sometimes this can introduce multiple layers of nesting (as an compound component could be used in another one).

Now I would like to know how big the performance penalty is if I would like to continue compound views to split up my layout and if there is a way to eliminate the issue of nesting compound components?
## [6][Why can't you set an ActionMode bar's style programmatically?](https://www.reddit.com/r/androiddev/comments/iehj83/why_cant_you_set_an_actionmode_bars_style/)
- url: https://www.reddit.com/r/androiddev/comments/iehj83/why_cant_you_set_an_actionmode_bars_style/
---
ActionMode bars... is there ANY way to set the height or style via Java (not xml)????

If not, why hasn't this feature been added? I feel like this is really bad.

I can't change every single ActionMode bar style via the theme just for ONE actionmode, specific situation.

I need to add more room when there's a subtitle.
## [7][Google releases Admob for Android](https://www.reddit.com/r/androiddev/comments/iefw8o/google_releases_admob_for_android/)
- url: https://www.reddit.com/r/androiddev/comments/iefw8o/google_releases_admob_for_android/
---
Google has released Admob app for android and looks like they are using Flutter for this just like they did with Goohle Ads app. What are your thoughts on this?

https://9to5google.com/2020/08/21/google-admob-android/
## [8][Turn off deprecation warning from Hilt compiled code](https://www.reddit.com/r/androiddev/comments/iefnff/turn_off_deprecation_warning_from_hilt_compiled/)
- url: https://www.reddit.com/r/androiddev/comments/iefnff/turn_off_deprecation_warning_from_hilt_compiled/
---
My gradle build system is set to stop when there are warnings on source files.

Hilt generated code uses some deprecated call

    Task :home:compileReleaseJavaWithJavac /.../home/fragment/Hilt_HomeFragment.java:42: warning:
    [deprecation] onAttach(Activity) in Fragment has been deprecated
    public void onAttach(Activity activity) {

so my build miserably fails

For now I allow to compile with warnings but the question is (drum roll):

Is it possible to turn off deprecated warnings coming **ONLY** from generated code (specifically Hilt)?
## [9][How do apps handle Periodic Jobs in 2020 that work outside app, with latest android version limiting background functionality and OEM's prioritizing battery life over app functionality ?](https://www.reddit.com/r/androiddev/comments/ie86ko/how_do_apps_handle_periodic_jobs_in_2020_that/)
- url: https://www.reddit.com/r/androiddev/comments/ie86ko/how_do_apps_handle_periodic_jobs_in_2020_that/
---

## [10][Projects using ButterKnife will have to migrate away? "Resource IDs will be non-final in Android Gradle Plugin 5.0"](https://www.reddit.com/r/androiddev/comments/idve1u/projects_using_butterknife_will_have_to_migrate/)
- url: https://i.redd.it/qlmylj022bi51.png
---

## [11][Learn Java REST API in 10 minutes (Web Service Coding Tutorial 2020)](https://www.reddit.com/r/androiddev/comments/ief1e6/learn_java_rest_api_in_10_minutes_web_service/)
- url: https://youtu.be/m2aqTS3EYnE
---

## [12][SQLDelight release with Kotlin Native 1.4.0](https://www.reddit.com/r/androiddev/comments/ieaob7/sqldelight_release_with_kotlin_native_140/)
- url: https://github.com/cashapp/sqldelight/releases/tag/1.4.1
---

