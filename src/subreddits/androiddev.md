# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/fhz0zc/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fhz0zc/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - March 09, 2020](https://www.reddit.com/r/androiddev/comments/ffsx46/weekly_questions_thread_march_09_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ffsx46/weekly_questions_thread_march_09_2020/
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
## [3][Securing Data Storage in Preparation for Pentesting](https://www.reddit.com/r/androiddev/comments/fhyjla/securing_data_storage_in_preparation_for/)
- url: https://infinum.com/the-capsized-eight/securing-data-storage-in-preparation-for-pentesting
---

## [4][Someone stole my app and modified the code a little bit and uploaded it to the PlayStore. What action should I take? In order to take that app down?](https://www.reddit.com/r/androiddev/comments/fhjw79/someone_stole_my_app_and_modified_the_code_a/)
- url: https://i.redd.it/hp8jauipx9m41.png
---

## [5][How do you manage app updates that rely on new web service versions with breaking database changes?](https://www.reddit.com/r/androiddev/comments/fhz2ts/how_do_you_manage_app_updates_that_rely_on_new/)
- url: https://www.reddit.com/r/androiddev/comments/fhz2ts/how_do_you_manage_app_updates_that_rely_on_new/
---
So we have app version A which relies on web service version 1 and the new app version B which should use the new web service version 2. No problem so far, both can run parallel.

But what do you do if version 2 of the web service has breaking database changes which won't allow web service 1 to work anymore.

Why did we choose two different parallel web service versions anyway? 
Because we want to make sure, that users who haven't updated from app A to app B can still use the app, even if web service endpoints have changed in the meantime.

So, is this a common problem? How do you manage it?
Or what alternative approaches do you use in order to not break old versions of your app?
## [6][The De-Googled Android Fork is Making Good Progress - It's FOSS](https://www.reddit.com/r/androiddev/comments/fhgknb/the_degoogled_android_fork_is_making_good/)
- url: https://itsfoss.com/gael-duval-interview/
---

## [7][How to use the onRestoreInstanceState Method with TextView string values](https://www.reddit.com/r/androiddev/comments/fhyyq3/how_to_use_the_onrestoreinstancestate_method_with/)
- url: https://www.reddit.com/r/androiddev/comments/fhyyq3/how_to_use_the_onrestoreinstancestate_method_with/
---
Hi dudes

Im not a professionnal of Kotlin, but I'm working in a big project for a quite simple app using Android Studio.

My app is working as I wished but I have an issue to save my datas when I switch of activity.I think its link to my Textview values that I always need to toString/toInt for treat them, but Idk how to unblock me.

Please  take a look to my code and tell me how and where create that onRestoreInstanceState method. Maybe my code is not as professional than  yours but please don't mind.

Tks a lot for the help I'm sure ur gonna do for me, heres the code :

&amp;#x200B;

[First Part](https://preview.redd.it/hfwbkbikpfm41.png?width=1001&amp;format=png&amp;auto=webp&amp;s=598c5f0c07f50b71bb9ddd3fe02bceddea4ee720)

&amp;#x200B;

[Second Part](https://preview.redd.it/cw7au9erpfm41.png?width=1368&amp;format=png&amp;auto=webp&amp;s=89d5c919016d6f857f6faa28e42422fd3f1b5c4d)

&amp;#x200B;

PS :  tell me too if my onSaveInstanceState method is correct or not, tks :)
## [8][Should I abandon ObjectAnimator for Motionlayout](https://www.reddit.com/r/androiddev/comments/fhyr8g/should_i_abandon_objectanimator_for_motionlayout/)
- url: https://www.reddit.com/r/androiddev/comments/fhyr8g/should_i_abandon_objectanimator_for_motionlayout/
---
I keep seeing that MotionLayout is the new thing, and was wondering will it replace ObjectAnimator and should I stop using ObjectAnimator?
## [9][Is it possible to create a Custom View Base class? Especially when a custom view inherits either View, Button, ConstraintLayout etc.](https://www.reddit.com/r/androiddev/comments/fhy2wx/is_it_possible_to_create_a_custom_view_base_class/)
- url: https://www.reddit.com/r/androiddev/comments/fhy2wx/is_it_possible_to_create_a_custom_view_base_class/
---
Seems hard to do as the custom view class would have to be a abstract class to get rid of the constructors.  
So when we previously had `class CustomView: ConstraintLayout`  
I now want to have something like:  
`class CustomView: BaseView&lt;ConstraintLayout&gt;`

unfortunately it seems like Kotlin doesn't allow extending generic types, which would have been the solution here so  
`abstract class BaseView&lt;V: ViewGroup&gt; : V` 

Is not possible, any other way that you can achieve something like this?   
As I said the constructors and inflating are always the same, but the subclasses still need to inherit from other classes as well.   
I know multiple inheritance isn't supported but there should be a solution to this :/
## [10][How do you handle factor specifig issues](https://www.reddit.com/r/androiddev/comments/fhvtdq/how_do_you_handle_factor_specifig_issues/)
- url: https://www.reddit.com/r/androiddev/comments/fhvtdq/how_do_you_handle_factor_specifig_issues/
---
I'm developing an Android-App and test it mainly on different Pixel devices. Sometime a colleague uses a Huawei P30 Pro and that's where the fun part starts. I see behaviour issues, I never had on any of the Pixels.Yeah, I could buy multiple devices from different brands, but well how should I fix specific issues, with some device name/brand checks in my app?I would appreciate any help on this topic.
## [11][“How to Create Notification Badges With Google Bottom Navigation Bar”](https://www.reddit.com/r/androiddev/comments/fhv2xq/how_to_create_notification_badges_with_google/)
- url: https://medium.com/better-programming/how-to-create-notification-badges-with-google-bottom-navigation-bar-330925474f6d?source=email-d70ea3470e6a-1584027034387-layerCake.autoLayerCakeWriterNotification----------
---

## [12][Where to start in an android application](https://www.reddit.com/r/androiddev/comments/fhx497/where_to_start_in_an_android_application/)
- url: https://www.reddit.com/r/androiddev/comments/fhx497/where_to_start_in_an_android_application/
---
Hi I'm trying to modify an existing Android application which wasn't bult by me.it is an application that uses native c++ library. The code is huge . Where to start reading to properly understand the whole pipeline of the application. It is a bit overwhelming to not know where or how to start.
