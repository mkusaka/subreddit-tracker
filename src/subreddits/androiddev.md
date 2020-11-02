# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/jmmr24/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jmmr24/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - November 02, 2020](https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jmley8/weekly_questions_thread_november_02_2020/
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
## [3][Netflix Android and iOS Studio Apps — now powered by Kotlin Multiplatform](https://www.reddit.com/r/androiddev/comments/jmdoux/netflix_android_and_ios_studio_apps_now_powered/)
- url: https://netflixtechblog.com/netflix-android-and-ios-studio-apps-kotlin-multiplatform-d6d4d8d25d23
---

## [4][Play protect Warning Pop-Up While installing Android app](https://www.reddit.com/r/androiddev/comments/jmkc2j/play_protect_warning_popup_while_installing/)
- url: https://www.reddit.com/r/androiddev/comments/jmkc2j/play_protect_warning_popup_while_installing/
---
 So we have an app that's deployed via the Microsoft AppCenter and one of the users received the following pop-up when trying to install it. 

&amp;#x200B;

&amp;#x200B;

[Installation Pop-Up](https://preview.redd.it/d1qzmdvxmsw51.png?width=380&amp;format=png&amp;auto=webp&amp;s=ca2890bc82ab24c582683c1716bd1f876d3bbc87)

 

Couldn't find any particular resource when I googled it except for the Minimum Target APK Version policy. The app in question has a minSdkVersion="21" and compile version(Target Framework) = Android 8.1 (27). I tried installing the same application on my device that has Android 11 but it works fine for me. Are we looking at a region specific issue? 

Would be super grateful if anybody can help me out with this.
## [5][Can an app identify other apps which are present on a phone](https://www.reddit.com/r/androiddev/comments/jmk7ie/can_an_app_identify_other_apps_which_are_present/)
- url: https://www.reddit.com/r/androiddev/comments/jmk7ie/can_an_app_identify_other_apps_which_are_present/
---
Sorry in advanced if this is the wrong sub but i have a question for any app developer. for context i have online exams coming up and i am required to download and invigilation app,at the same time for privacy reasons i want to download a microphone disabling app. The invigilation app requires storage permission ,so my question is will the invigilation app be able to identify what other apps are present on my phone
## [6][does System.exit(int) kill the process and clear up all the resources?](https://www.reddit.com/r/androiddev/comments/jmj19y/does_systemexitint_kill_the_process_and_clear_up/)
- url: https://www.reddit.com/r/androiddev/comments/jmj19y/does_systemexitint_kill_the_process_and_clear_up/
---
According to [my previous post](https://www.reddit.com/r/androiddev/comments/jmfw8v/weird_behaviour_of_service_and_media_player/), and some research online, I found out that the sytem might be playing the music even after the service's `onDestroy` was called because it was holding on to some resources like the `MediaPlayer`, so I decided to actually kill the process instead of calling the  `stopForeground` or `stopSelf`, because they simply are not working. However, I synced all the application states and called `System.exit(0)` and it kills the foreground process as well as the activity as intended.

The only bit of confusion I had is if `System.exit(0)` killed the process and also freed the app resources? That is, does it release every bit of resource held by the app? I am worried about it because I don't need to call `MediaPlayer.release()` and `stopForeground()` and other similar cleanup methods, just one `System.exit()` does the trick.

It's okay to call this for my app, because it triggers the behaviour I intend, but, I wonder if its okay for the android system, because lingering memories would be a bad thing. Thanks in advance, any insight would be helpful.
## [7][Im doing English-Spanish translations on small apps for free.](https://www.reddit.com/r/androiddev/comments/jm31ho/im_doing_englishspanish_translations_on_small/)
- url: https://www.reddit.com/r/androiddev/comments/jm31ho/im_doing_englishspanish_translations_on_small/
---
 That's it. I have a lot of free time since i ~~work from home~~ lost my job.  


 I posted the same thing a few months ago when i was working from home and translated about 50 apps and a few random things  


So im doing it again because honestly im bored and i really need to do something besides trying to get a new job lol.  


Im doing up to 500-1000 words, not doing ToS or anything like that.

Anyone can PM me (:
## [8][Fewer crashes and more stability with Kotlin](https://www.reddit.com/r/androiddev/comments/jmlyyf/fewer_crashes_and_more_stability_with_kotlin/)
- url: https://medium.com/androiddevelopers/fewer-crashes-and-more-stability-with-kotlin-b606c6a6ac04
---

## [9][Something like LFS.](https://www.reddit.com/r/androiddev/comments/jmkdl3/something_like_lfs/)
- url: https://www.reddit.com/r/androiddev/comments/jmkdl3/something_like_lfs/
---
I'm always fascinated by Android and various ports of Android.
We know that there are tons and tons of "build" guides on the internet but non of them actually explain how things work.
If only we had something like LFS for android, to teach people how android actually works, what goes where, etc, we'd be able to create more ports for various devices.
Like right now, I wanna know how the port of android x86 actually works and how can I port other custom android like Lineage OS to work on x86 platform. (I know android x86 has Lineage OS port already, but that's version 14.1, Lineage is on 17.1 now)
## [10][Auto Revoke Permissions settings panel missing on some apps (Android 11 API/System)](https://www.reddit.com/r/androiddev/comments/jmlsfe/auto_revoke_permissions_settings_panel_missing_on/)
- url: https://www.reddit.com/r/androiddev/comments/jmlsfe/auto_revoke_permissions_settings_panel_missing_on/
---
Hi,

I'm trying to whitelist some of my enterprise apps, but it seems ([https://issuetracker.google.com/issues/170968725](https://issuetracker.google.com/issues/170968725)) that a new manifest attribute " autoRevokePermissions " isn't working (anymore or for now, I don't know because its new on API lvl 30 BUT its also not working AND bypassed by system...logic).

So I'm trying to ask user permisison for this action, like explained here : [https://developer.android.com/guide/topics/permissions/overview#auto-reset-permissions-unused-apps](https://developer.android.com/guide/topics/permissions/overview#auto-reset-permissions-unused-apps)

My problem is, on some of the apps, this section of settings is totally gone. I've updated targetSdkVersion on all, the configurations seems similar...

Does anyone know why I encounter this issue ? (Sorry for bad english I'm french ;)
## [11][WebSockets and Android apps - client-side considerations](https://www.reddit.com/r/androiddev/comments/jmlnh7/websockets_and_android_apps_clientside/)
- url: https://www.ably.io/topic/websockets-android/?utm_source=reddit&amp;utm_medium=sc&amp;utm_campaign=cm_wand
---

## [12][What type of APP is used to communicate with micro-controllers ?](https://www.reddit.com/r/androiddev/comments/jmllzh/what_type_of_app_is_used_to_communicate_with/)
- url: https://www.reddit.com/r/androiddev/comments/jmllzh/what_type_of_app_is_used_to_communicate_with/
---
Hi guys!

**Before I begin, if you find this post inappropriate for this subreddit, please let me know and direct me elsewhere if you don't mind. Thank you!**

I'm planning a project which will require an app to control it, but as I'm very new to this, I genuinely have no idea what am I looking for. I'm not asking for help with the development, but I could seriously use a few pointers just so I know where to start. There's an overwhelming amount of programming languages / frameworks / other relevant things and it's incredibly easy to become confused by it

**In a single sentence**: the app should be able to communicate with an **external micro-controller** (ESP32) via Bluetooth and/or WIFI in order to control an individually addressable LED strip (such as WS2812B &amp; SK6812B). I'd like it to be both Android &amp; IOS compatible if possible.

I'll continue to describe the plan in detail so you can get a better idea of what I'm trying to accomplish give feedback (if you want to help) accordingly. If something on this list is not possible or is extremely complicated, please take a moment to let me know if you don't mind.

\----------------------------------------------------------------------------------------------------------------------------------------

**App Setup plans**

1.Download &amp; install APP (Android + IOS compatibility)2.Select a control method - Bluetooth or Wifi- this option can be changed later but active option would apply to all devices (if multiple are added)- if both is not possible without user having to flash an ESP32 manually, then only one choice would be available.

3.Select the LED strip chipset (WS2812B / SK6812, etc.)4.Select color order (RGB / GRB, etc.)5.Input the total number of LEDs6.Continue to APP / add more devices

Since I'm likely to make custom models with separate dedicated white LEDs, I'd like to be able to add X models through APP updates.

\----------------------------------------------------------------------------------------------------------------------------------------

**App Feature plans**

1.Brightness control (global)2.Speed control (global)3.Mic sensitivity adjustment4.On/off switch (turns the LEDs off via SSR to eliminate idle power consumption &amp; puts ESP in idle mode)5.Auto-mode (each effect is displayed for X amount of time, then a new effect begins)6.Manual solid color presets + RGB wheel / color picker7."Add to favorites" button &amp; list (quick access to saved pre-made effects)8.Save changes button &amp; list (pre-made effects with custom brightness and/or speed settings)9.Option to "sync" multiple lamps so they can be controlled simultaneously10.Update APP if controls were changed using built-in push buttons, encoders/pots and switches11.Language change (2 languages, without "auto-translate")

\----------------------------------------------------------------------------------------------------------------------------------------

\*\*Advanced Features (\*\*This option would require a small one-time payment)

**1**.Wake up feature // using built-in clock or RTC module if needed

a) Gradual - lights turn on at set time and brightness gradually increases. Lights stay this way until user input occurs.b) Flashing - lights continuously flash until cancelled but stop after 30 minutes or something like that.c) Lights &amp; Sound - combines option a or b (user's choice) with a user selected premade sounds (uses a small built-in speaker)

**2**.Timer &amp; schedule function

a) Schedule the lights to turn off after a set amount of timeb) Lights turn on &amp; off automatically at set times (+ choice for which days of the week)

**3**.Proximity switch (Using a micro-wave radar or a passive infra-red sensor to detect nearby people)

a) When a person is detected, the lights will automatically turn on (dedicated white light as a base, but changeable)b) When a person is not detected for X amount of time, lights turn off or go into idle mode

// I don't know if this is possible to achieve without the lights turning off if a person isn't moving a lot. Why - PIR sensors usually use timers and only work with motion, if a person is sitting, they go out.

**4**.Alexa &amp; Google Home compatibility

\-- Not as important, but would be nice. Not sure if it's possible with a custom app or worth the integration price I'd probably have to pay since I know nothing about this but I reckon it would be a cool feature to have.

**Commands** (just a few basic stuff):

Lights on/offChange effect to \_preset keyword or numberDim the X light to/by X percentIncrease brightness to X percent / set to full

**5**.Clock / temperature / humidity (if the model has a built-in LED segment display or an OLED module)

Turn on/off  
Auto cycle through stats / modes by default

\----------------------------------------------------------------------------------------------------------------------------------------So Based on all that info, could you recommend a place to start ? Selecting an appropriate programming language is a priority, but any additional tips or suggestions are welcome.

I was also wondering, how much would an APP like this cost ? Just in case someone has an approximation. In case this is an inappropriate question, ignore it.

Anyways, thank you for taking the time &amp; have a great day!
