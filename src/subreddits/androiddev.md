# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/fecwy1/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fecwy1/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - March 02, 2020](https://www.reddit.com/r/androiddev/comments/fcaf42/weekly_questions_thread_march_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fcaf42/weekly_questions_thread_march_02_2020/
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
## [3][PSA: Google survey about Scoped Storage for Android 11](https://www.reddit.com/r/androiddev/comments/febs6h/psa_google_survey_about_scoped_storage_for/)
- url: https://google.qualtrics.com/jfe/form/SV_9HOzzyeCIEw0ij3?Source=scoped-storage
---

## [4][Aspiring Android developer discouraged from reading about bans - would you still choose Android as a career if you did it over?](https://www.reddit.com/r/androiddev/comments/fe9laj/aspiring_android_developer_discouraged_from/)
- url: https://www.reddit.com/r/androiddev/comments/fe9laj/aspiring_android_developer_discouraged_from/
---
TL;DR at the bottom.

A bit of background about me: I'm 27 and finishing up a degree in Computer Science. Lately I've been interested in Android development and working through their official code labs and concept chapters. I love Android Studio and enjoy the Android platform with its fun phone sensors and other features to work with. I've been planning on focusing on this and pursuing a career in it.

However, I'm discouraged after reading about Google's treatment of developers (for example, removing apps without clear instructions and support about policy violations, or even worse, an outright developer ban).

My biggest fear is if I somehow got permanently banned from the Playstore, that would severely limit my employability as an Android developer - especially as a new one. I'm not planning on making spammy apps, but from what I've read it seems that Google can sometimes be vague and capricious in removing apps and banning developers.

How likely is a permanent ban if you try to follow their policies? And if worst case scenario you do get banned, would that actually affect your employability as an Android developer? Or am I worrying about this too much?

I actually have a course at school for Android and need what I've studied so far to pass it, so even if I don't pursue Android development professionally, my studying hasn't been for nought. And I've got front end experience as well as Java, so I may go one of those routes instead. Or possibly learn React Native and try to find a job doing that.

Am I overthinking this? Would you, a professional Android developer, still choose this as a career if you were in my shoes?

TL;DR Considering a career in Android development, but worried about being banned from the Playstore even if I try to follow policy. Biggest worry is a ban affecting my employability. Is that a legit worry?
## [5][How do reach out to someone at Google Play? I tried updating the logo for the application I am developing, got rejected saying we are impersonating a brand. I didn't know there was chat support in Google Play, I tried last night and this morning and I see this. What is 12am - 12am (GMT)?](https://www.reddit.com/r/androiddev/comments/fe1rqy/how_do_reach_out_to_someone_at_google_play_i/)
- url: https://i.redd.it/stw6w4gz1xk41.png
---

## [6][Facebook punishes my Android game: Your app cannot be installed by new users. 🤡](https://www.reddit.com/r/androiddev/comments/febcjn/facebook_punishes_my_android_game_your_app_cannot/)
- url: https://www.reddit.com/r/androiddev/comments/febcjn/facebook_punishes_my_android_game_your_app_cannot/
---
Hello,

since over 2 years I run a small (150 mobile users, a bit more on desktops) word game, without any IAP or ads (I don't like ads and I haven't implemented IAP yet).

The mobile version is published at Google Play, Huawei App Gallery, Amazon.

Users can login through Google, HMS, Amazon Login and - Facebook login.

Yesterday I have gotten the warning mail by Facebook: "Your app is in violation of our Platform Policies"

Being an indy dev I live in constant Angst and thus I have immediately updated my app at Google Play (changed the Facebook icon from greyscale to blue-white) and re-entered the base64 hashes in Facebook settings. I haven't fixed anything more because actually the Facebook Login just works in my little app.

And of course this didn't help and now my app is:

**Your app XXXXX XXXXXX has been restricted for creating a negative user experience in violation of our Platform Policies.**  Your app XXXXX XXXXXX has been restricted for creating a negative user experience that is in violation of our Platform Policies.

* Your app cannot be installed by new users.

I understand many things, been watching devs in this Reddit being banned last year by Google robots, understand that Facebook uses similar robot scripts...

But how on Earth does Facebook want to forbid new users to install my Android app? What do they mean by the above restriction (and yes the Facebook login still works in my game).

P.S. My game is super clean: no ads, no IAP, all-native Java, here few screenshots: [https://imgur.com/a/2m2X3Pa](https://imgur.com/a/2m2X3Pa)
## [7][Glide image library being reported as malware by AVG, Avast](https://www.reddit.com/r/androiddev/comments/fe35hv/glide_image_library_being_reported_as_malware_by/)
- url: https://www.reddit.com/r/androiddev/comments/fe35hv/glide_image_library_being_reported_as_malware_by/
---
I spent 6 hours paring down this issue, removing one library at a time. Hopefully this information is useful for others who are hitting this issue.

Starting today, AVG and Avast (they are the same reskinned app and have combined over hundreds of millions of downloads) started reporting my apps that use Glide (4.8.0 and 4.11.0) as malware. I reported it to AVG as a false positive already, but don't expect a quick fix from them.

When I replaced Glide with Picasso, the apps are no longer marked as malware. I've already received about 10 bad reviews and dozens of e-mails about this issue and have not updated my app recently, which leads me to think they updated their algo very recently. I have tried disabling proguard but it makes no difference.


Quick fix is to replace Glide with Picasso and the issue goes away. Replace:

-            Glide.with(context).load(uri).into(imageView)

With:

+            Picasso.get().load(uri).into(imageView)


In gradle, replace:

-    api "com.github.bumptech.glide:glide:4.8.0"

With:

+    api 'com.squareup.picasso:picasso:2.71828'
## [8][[OC] I’ve made a LIVE INTERACTIVE dashboard to track COVID19](https://www.reddit.com/r/androiddev/comments/fedaw3/oc_ive_made_a_live_interactive_dashboard_to_track/)
- url: https://v.redd.it/0n0d35ls8uk41
---

## [9][Animated Play/Pause Button: Animated Vector Drawables tutorial](https://www.reddit.com/r/androiddev/comments/fecdf8/animated_playpause_button_animated_vector/)
- url: https://blog.stylingandroid.com/animatedicons-play-pause-circle/
---

## [10][anoop44/3dview](https://www.reddit.com/r/androiddev/comments/fe7xp5/anoop443dview/)
- url: https://github.com/anoop44/3dview
---

## [11][Developing your own in-app update system. Best way to avoid random app/account ban?](https://www.reddit.com/r/androiddev/comments/fe9yve/developing_your_own_inapp_update_system_best_way/)
- url: https://www.reddit.com/r/androiddev/comments/fe9yve/developing_your_own_inapp_update_system_best_way/
---
I am seeing a lot of horrifying stories of developers whose apps are getting banned for frivolous reasons. An account ban automatically tarnishes your image as an android developer, and people may no longer be ready to work with you. All because you did a silly mistake or there was no mistake at all and Google bots just nuked your account due to false alarm.

I was thinking this. Once your app does get published on Google Play, you are pretty much safe till you try to upload a next update. At that point of time, your app is scanned again, and you keep your fingers crossed that it won't be rejected/lead to a ban.

So why not add an in-app updater?
It's pretty easy actually. You put a file on GitHub/Gitlab with update definition (current version of the app and url for direct download of the update apk) compare the current version of the app with the update definition and if your current version is outdated, download the new update apk and prompt the user to install it.

Atleast once your app gets published for the first time, you won't have to deal with unnecessary fears of getting banned in a future update. (Unless Google re-scans the app and finds it violating policies or anything similar)
## [12][App removed from playstore.](https://www.reddit.com/r/androiddev/comments/fdswex/app_removed_from_playstore/)
- url: https://www.reddit.com/r/androiddev/comments/fdswex/app_removed_from_playstore/
---
Every morning when I wake up I check my phone. I am always a little bit anxious of seeing this little icon from the google console in my notification bar. Usually not good news.This morning I saw that icon and almost got a heart attack when I read the notification:all my apps removed because of violation of the developer content policy. They did not send me an email at all. What the hell, this cannot be true. Not again please!Immediately contacted my account manager at Google admob. Asked him if he could forward my issue directly to the playstore team.And of course i quickly went over all my apps. Checking for missing things like app ratings etc. Found an email from last year from the Google Playstore saying: \[Action Required\] Switch to the Play Referrer API by March 1, 2020. I remember that the issue was actually in one of the firebase libs which were not fixed even a month ago. Could it be that? Lets update the libs first. I also noticed that inside the user aqcuisition section of my apps there was a warning about some kind of expired payment method. Weird, because I am not using that functionality. Checked [google.ads](https://google.ads) and found out that my creditcard was expired. Updated it immediately.

I filed quickly a new update of all the apps without testing them. Published it and 2 hours laters i got a notification again that my apps are back live! Amazing. My entire morning's mood changed from a hopeless , feared and angered person to the happiest man on earth.

Nevertheless i find it still unacceptable that these kind of things are possible. I still don't know what the real issue was. My apps are having about 5m daily users and are for a large part generating income for my company. App business remains a still risky one!

//edit

now 10 hours i receive an email from Admob that action is required.   
Looking at the policy violations details inside the admob screen they saw that I need to fix the app because of removal. But the app is back live again ?!?!.  to be continued.....
