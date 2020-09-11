# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/iqpolu/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/iqpolu/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - September 07, 2020](https://www.reddit.com/r/androiddev/comments/io5e9j/weekly_questions_thread_september_07_2020/)
- url: https://www.reddit.com/r/androiddev/comments/io5e9j/weekly_questions_thread_september_07_2020/
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
## [3][Reasons for 1 star review on Google Play.](https://www.reddit.com/r/androiddev/comments/iqmirx/reasons_for_1_star_review_on_google_play/)
- url: https://www.reddit.com/r/androiddev/comments/iqmirx/reasons_for_1_star_review_on_google_play/
---
1️⃣

User: Expects z feature.

Your app: Has the best x, y features. Doesn't have z feature.

User: "Garbage app. I want z feature."

You: Never said I have z feature at all. 👁👄👁

2️⃣

User: "I hate ads. Too many ads."

Your app: "One small banner ads at the bottom peacefully or pay $0.99 to remove ads."

3️⃣

Your app: "Pay $0.99 to unlock some premium features."

User: "One star. Because I think it should be free. Greedy scammer developer!"

You: 👁👄👁

4️⃣

You: "Don't ask user to rate so frequently. Don't annoy them. Only ask at the right time."

Your app: Showing pop up dialog for the first time after 10 launched count. "If you love it, please write us a review on Google Play. Thanks :)"

User: "One star. Because this app forced me to write review."

Cancel button: "Am I a joke to you?"

5️⃣

User: Rated One star. "Great. I love it."●

You: 👁👄👁
## [4][Progress Indicators – Material Components For Android](https://www.reddit.com/r/androiddev/comments/iqjgye/progress_indicators_material_components_for/)
- url: https://howtodoandroid.com/progress-indicators/
---

## [5][Firebase Cloud Messaging (FCM) Android RANT](https://www.reddit.com/r/androiddev/comments/iqn47n/firebase_cloud_messaging_fcm_android_rant/)
- url: https://www.reddit.com/r/androiddev/comments/iqn47n/firebase_cloud_messaging_fcm_android_rant/
---
This is a **rant**. 

First of all I want to say that Firebase is great and Google did a great job there. But this doesn't mean it's perfect.

The following is something relatively small but it has been nagging me for a while now.

Google Firebase Cloud Messaging has 3 ways of sending push notification

1. Notification: this is received by your app if it is in foreground, and directly goes to the notification tray if your app is in background
2. Data: always received by your app
3. Notification+Data: mix of the two, it works like 1) but you get data when you open the notification.

The Firebase Console has a cool tool to build notification of type 1) which give you a preview, let you chose when to send a notification, who to send to... yeah great.

Small, tiny issue.... when you get type 1) in Android and your app is in background your app doesn't receive it, when you click the notification your app is opened but the Intent associated with it **CONTAINS NO DATA** of the notification (title, text, image, whatever).

This is what you get in the extras: 

```
'google.delivered_priority' : 'high'
'google.sent_time' : '1599808968376'
'google.ttl' : '2419200'
'google.original_priority' : 'high'
'from' : '23094812394'
'google.message_id' : '0:198347123094%9d7575989d2304341'
'collapse_key' : 'your.app.package.id'
```

Completely useless, specially cause there's no service to obtain the notification from `google.message_id`. You can't even show the notification inside the app, your app "just opens". 

Oh, but wait, you [could](https://firebase.google.com/docs/cloud-messaging/http-server-ref#notification-payload-support) at least add `click_action` to the notification to specific what to do when clicked **but** you cannot set it from the firebase console.

Even more outraging: iOS receive that notification and all its data when the user clicks it (title, text, image and everything).

I think we can all agree that a notification with text/image in it that when clicked disappear forever is never gonna be a good UX experience.

So I really don't understand the logic behind this, why would Google give you this kind of behavior and a default backoffice that provides this crippled UX on Android only.

It basically force everyone to always implement a backoffice even for the most basic of the use case even when you wouldn't need it for iOS.

I really can't believe this is what they actually wanted.

BONUS: when the app is in background it builds a notification using all the variable set in the Notification body (image, sound, ... etc) but if you receive it when the app is in foreground and you want to build that same exact notification you need to do it yourself, why not just expose an utility method to do exactly that, Google you already have the code, right?


I just hope someone in Google agree with this and do something to put iOS and Android on par. Just make the notification data available in the extras effectively making your app receive the same information regardless.

(Would be actually AWESOME if it could give a notification history service with read / unread notification out of the box but this is most likely asking too much).

Am I missing something here? Do you agree?
## [6][PSA: Android 11 leaks view hierarchy if you pass a Spannable to a TextView](https://www.reddit.com/r/androiddev/comments/iq56zn/psa_android_11_leaks_view_hierarchy_if_you_pass_a/)
- url: https://www.reddit.com/r/androiddev/comments/iq56zn/psa_android_11_leaks_view_hierarchy_if_you_pass_a/
---
edit: Someone pointed out these issues are for Android 9+, which I initially missed. However, something must've changed with Android 11 (or Android 11 emulator), for me this issue is now triggered while on previous versions it's not. Don't believe me though, just install LeakCanary in your apps and check if your apps are still OK

This just popped up from LeakCanary when I was testing app on a new Android 11 emulator. If the reports are correct (I haven't thoroughly tested myself), the view hierarchy will leak through `android.app.assist.AssistStructure`as soon as you set a `Spannable` as the `TextView`'s content. Some links for reference:

* A twitter thread: https://twitter.com/piwai/status/1107731823625265152
* Tracker issue (not even acknowledged): https://issuetracker.google.com/issues/146100180
* Another issue, closed _because_: https://issuetracker.google.com/issues/62204023
* LeakCanary thread: https://github.com/square/leakcanary/issues/1573

edit: typo
## [7][Looking for a experienced Livedata mentor.](https://www.reddit.com/r/androiddev/comments/iqo4zi/looking_for_a_experienced_livedata_mentor/)
- url: https://www.reddit.com/r/androiddev/comments/iqo4zi/looking_for_a_experienced_livedata_mentor/
---
I'm looking for someone that can help me understand  livedata concept and all the complex thing that goes with it ( map, witchmap, custom mediator...), in a real project context.  
I dont know why i'm struggling so much to understand these, i guess i miss a big point  but i cant get things done the way i want to.
## [8][Why didn't Nintendo go with a heavily modified version of Android, as opposed to developing their own OS?](https://www.reddit.com/r/androiddev/comments/iqn66r/why_didnt_nintendo_go_with_a_heavily_modified/)
- url: https://www.reddit.com/r/androiddev/comments/iqn66r/why_didnt_nintendo_go_with_a_heavily_modified/
---
I understand that Nintendo wanted a closed ecosystem, and that they wanted to reduce bloat. But couldn't they have just modified Android Go or something, and taken steps to remove all the bloat, in order to ensure that they would have a system that was still closed, but able to benefit from Android development libraries?

Essentially... Couldn't Nintendo have gone with Android whilst ensuring a closed ecosystem?

I think Amazon have succeeded in this, as they have the Fire OS which is used for their Kindle Fire tablet line.
## [9][Android 11 tightens restrictions on CA certificates](https://www.reddit.com/r/androiddev/comments/iq77mf/android_11_tightens_restrictions_on_ca/)
- url: https://httptoolkit.tech/blog/android-11-trust-ca-certificates
---

## [10][Play 2 songs from 2 apps at once](https://www.reddit.com/r/androiddev/comments/iqppxa/play_2_songs_from_2_apps_at_once/)
- url: https://www.reddit.com/r/androiddev/comments/iqppxa/play_2_songs_from_2_apps_at_once/
---
Hello,
I need help, I'm looking for a way to read several audio sources at the same time on my phone, without the first one being cut off. Example: Youtube and Spotify at the same time.
Can you tell me how I could do this? Do I need a specific programme? I'm willing to pay for it. Thank you.
## [11][How can I access all my monthly reports?](https://www.reddit.com/r/androiddev/comments/iqpbeq/how_can_i_access_all_my_monthly_reports/)
- url: https://www.reddit.com/r/androiddev/comments/iqpbeq/how_can_i_access_all_my_monthly_reports/
---
I can only access a few of my monthly reports...

&amp;#x200B;

https://preview.redd.it/23dhdu79gim51.png?width=1126&amp;format=png&amp;auto=webp&amp;s=732f6a7c92e7d2e6e92a6f05263a937a28ecb1c2
## [12][Need help to make my app compatible for android 10 &amp; 11](https://www.reddit.com/r/androiddev/comments/iqpbe9/need_help_to_make_my_app_compatible_for_android/)
- url: https://www.reddit.com/r/androiddev/comments/iqpbe9/need_help_to_make_my_app_compatible_for_android/
---
I have source code for an android app that currently supports up to android 9. How can I make it compatible to run on android 10 &amp; 11

BTW I am new to this sub and don't know which flair to choose.
