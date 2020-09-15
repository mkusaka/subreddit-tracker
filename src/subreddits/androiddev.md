# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/isk58n/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/isk58n/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - September 14, 2020](https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/)
- url: https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/
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
## [3][Kaspresso 1.2.0: AdbServer](https://www.reddit.com/r/androiddev/comments/it5bg1/kaspresso_120_adbserver/)
- url: https://www.reddit.com/r/androiddev/comments/it5bg1/kaspresso_120_adbserver/
---
Hello everyone!

We are happy to announce a new version of Kaspresso!

Shortly, Kaspresso is a great framework for UI testing. Based on [Espresso](https://developer.android.com/training/testing/espresso) and [UI Automator](https://developer.android.com/training/testing/ui-automator), Kaspresso provides a wide range of additional amazing features, such as:

* 100% stability, no flakiness.
* Significantly faster execution of UI Automator commands. With Kaspresso, some UI Automator commands run **10 times faster**!
* Excellent readability due to human DSL.
* Incredible mechanism of interceptors that allows you to catch all actions and assertions in one place.
* Full logging.
* Ability to call ADB commands.
* UI tests writing philosophy, implemented with DSL.
* Features screenshotting.

All details about the new version of Kaspresso are [here](https://github.com/KasperskyLab/Kaspresso/releases/tag/v1.2.0).
## [4][What are the current caveats in KMM?](https://www.reddit.com/r/androiddev/comments/it2c9r/what_are_the_current_caveats_in_kmm/)
- url: https://www.reddit.com/r/androiddev/comments/it2c9r/what_are_the_current_caveats_in_kmm/
---
Recently JB released [KMM 1.0 alpha](https://www.reddit.com/r/androiddev/comments/ik1zu8/kotlin_multiplatform_mobile_goes_alpha_kotlin_blog/) and they are [addressing memory model](https://www.google.com/search?q=memory+model+kotlin+native) issue but what about other issues?

* Tooling : IDE, Gradle
* Ecosystem
* iOS performance issues

Does Android / iOS app 

* performance get affected
* .apk / .ipa size increases 

with KMM?

Can anyone please share your experience?  Can we confidently share code between iOS and Android platforms without much hassle?
## [5][Java 15 is now available.](https://www.reddit.com/r/androiddev/comments/it85li/java_15_is_now_available/)
- url: http://openjdk.java.net/projects/jdk/15/
---

## [6][Turning it up to 11: Android 11 for developers](https://www.reddit.com/r/androiddev/comments/it5o5a/turning_it_up_to_11_android_11_for_developers/)
- url: https://android-developers.googleblog.com/2020/09/android11-final-release.html
---

## [7][Kotlin 1.4 Online Event – Kotlin Blog](https://www.reddit.com/r/androiddev/comments/islbfb/kotlin_14_online_event_kotlin_blog/)
- url: https://blog.jetbrains.com/kotlin/2020/09/kotlin-1-4-online-event/
---

## [8][What is really going in these lines (except the intent declarartion part)](https://www.reddit.com/r/androiddev/comments/it8gx9/what_is_really_going_in_these_lines_except_the/)
- url: https://www.reddit.com/r/androiddev/comments/it8gx9/what_is_really_going_in_these_lines_except_the/
---
 public static Drawable getActivityIcon(Context context, String packageName, String activityName) { PackageManager pm = context.getPackageManager(); 

Intent intent = new Intent(); 

intent.setComponent(new ComponentName(packageName, activityName));

ResolveInfo resolveInfo = pm.resolveActivity(intent, 0);

return resolveInfo.loadIcon(pm); 

}
## [9][Anyone got affected by new google algorithm change?](https://www.reddit.com/r/androiddev/comments/it84t1/anyone_got_affected_by_new_google_algorithm_change/)
- url: https://www.reddit.com/r/androiddev/comments/it84t1/anyone_got_affected_by_new_google_algorithm_change/
---
I checked App Tweak, and it shows that on 15th of August, Google changed their algorithm again and that's exactly when I started to lose a lot of downloads, has anyone got affected by it? I was getting around 1500 downloads daily, but now I am getting around 500, and it's still dropping more and more, I have no idea what's happening.
## [10][Printing document](https://www.reddit.com/r/androiddev/comments/it1xnn/printing_document/)
- url: https://www.reddit.com/r/androiddev/comments/it1xnn/printing_document/
---
Hello Redditors, I’m working on a app in Java for billing purposes. I need to get printed invoices on a Bluetooth printer. Can someone tell me how to get it done?
Thank you
## [11][Custom component with own logic](https://www.reddit.com/r/androiddev/comments/it78fj/custom_component_with_own_logic/)
- url: https://www.reddit.com/r/androiddev/comments/it78fj/custom_component_with_own_logic/
---
Hey guys! I'm trying to create a custom component that has its own logic. I'm using MVI architecture and for each feature I have a fragment, view, presenter, viewstate. I want to include my component inside one of my fragments. Is it bad for my component to also be a fragment? I know that having a fragment inside a fragment is an antipattern, but I don't know how to create something else that will have its own logic, presenter, states and everything. 

&amp;#x200B;

Thank you!
## [12][Mixin Messenger, an open-source End-to-End Encryption Messenger.](https://www.reddit.com/r/androiddev/comments/it76c0/mixin_messenger_an_opensource_endtoend_encryption/)
- url: https://github.com/mixinnetwork/android-app
---

