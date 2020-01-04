# androiddev
## [1][App Feedback Thread - January 04, 2020](https://www.reddit.com/r/androiddev/comments/ejw3fu/app_feedback_thread_january_04_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ejw3fu/app_feedback_thread_january_04_2020/
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
## [2][Weekly Questions Thread - December 30, 2019](https://www.reddit.com/r/androiddev/comments/ehkp99/weekly_questions_thread_december_30_2019/)
- url: https://www.reddit.com/r/androiddev/comments/ehkp99/weekly_questions_thread_december_30_2019/
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
## [3][GitHub - lndmflngs/hat-view: Library that allow to put "hat" on TextView. Inspired by Telegram appbar title with Santa Claus hat 🎅🏻](https://www.reddit.com/r/androiddev/comments/ejvovk/github_lndmflngshatview_library_that_allow_to_put/)
- url: https://github.com/lndmflngs/hat-view
---

## [4][How do some launcher apps import shortcuts and other settings from Nova Launcher?](https://www.reddit.com/r/androiddev/comments/ejv4lz/how_do_some_launcher_apps_import_shortcuts_and/)
- url: https://www.reddit.com/r/androiddev/comments/ejv4lz/how_do_some_launcher_apps_import_shortcuts_and/
---
 

I've noticed that if you used Nova launcher, and you try out some launchers (such as Apex Launcher, Evie Launcher, ADW Launcher 2, and maybe others), they let you import various settings from Nova launcher, including the shortcuts to apps and even the widgets.

How do they do it? How come it doesn't exist for other launchers? Are there other launchers that you can switch from, knowing that another launcher could import the data from?

Is this a known API ?
## [5][Kotlinx Json vs Gson - a short article about Kotlin oriented cross-platform reflectionless serialization library from JetBrains.](https://www.reddit.com/r/androiddev/comments/ejjroe/kotlinx_json_vs_gson_a_short_article_about_kotlin/)
- url: https://medium.com/@jurajkunier/kotlinx-json-vs-gson-4ba24a21bd73
---

## [6][what’s the current best practice for implementing Android Room?](https://www.reddit.com/r/androiddev/comments/ejwcpk/whats_the_current_best_practice_for_implementing/)
- url: https://www.reddit.com/r/androiddev/comments/ejwcpk/whats_the_current_best_practice_for_implementing/
---
I’ve never done any MVVM thing before, and I don’t know what’s the significance of using such a headache!

Room is currently working fine for me without too much complications, but I’m using GlobalScope launch to get access the database, which is highly discouraged as I read [here](https://kotlin.github.io/kotlinx.coroutines/kotlinx-coroutines-core/kotlinx.coroutines/-global-scope/). 

I wanna do it in a better way and learn more about the subject (MVVM, lifecycles, LiveData, etc), can you provide some links that help with that?

I’m hoping that implementing these will result in a better code.

Thanks, appreciate the help!
## [7][SSL certificate unknown](https://www.reddit.com/r/androiddev/comments/ejwhst/ssl_certificate_unknown/)
- url: https://www.reddit.com/r/androiddev/comments/ejwhst/ssl_certificate_unknown/
---
Hello,

Trying to create secure MQTT connection to Mosquitto server.  I'm using CA.crt certificate to get secure connection with help of [sample application](https://github.com/eurbon/Paho-MQTT-Android-TCP-TLS-WSS-Example/blob/master/app/src/main/java/paho/android/mqtt_example/MainActivity.java). 

&amp;#x200B;

Mosquitto server reports with error:

`1578142589: OpenSSL Error: error:14094416:SSL routines:SSL3_READ_BYTES:sslv3 alert certificate unknown`

`1578142589: OpenSSL Error: error:140940E5:SSL routines:SSL3_READ_BYTES:ssl handshake failure`

&amp;#x200B;

I'm using same CA.crt certificate to connect from linux machine at it works fine:

`mosquitto_sub -h` [`192.168.1.8`](https://192.168.1.8) `-t s1/ss --cafile ca.crt -p 8883 -d -u c -P c -i dd`

&amp;#x200B;

What might be wrong with my connection and how to solve this problem.
## [8][[Need advice] AdMob restricted ad serving due to curse words](https://www.reddit.com/r/androiddev/comments/ejvto1/need_advice_admob_restricted_ad_serving_due_to/)
- url: https://www.reddit.com/r/androiddev/comments/ejvto1/need_advice_admob_restricted_ad_serving_due_to/
---
Hello, fellow Android devs,

I need advice on what to do in such a situation:

My app plays Russian phrases that are popular in a game called Counter Strike: Global Offensive. I have released 2 versions of an app, one that contains no bad words and another that is basically full of curse words both in English and in Russian (I call this 'rage' version). I had integrated AdMob interstitial ads into both of the apps and it was working fine. However, yesterday I have received a notification that ads were restricted in 'rage' version. The issue is "Adult: Sexual content":

&amp;#x200B;

https://preview.redd.it/04ik6beg4r841.png?width=941&amp;format=png&amp;auto=webp&amp;s=1868a1cd06b06123fa02ab21a8b6088e0cd08c13

As I said, the app contains no visual sexual content, only written and spoken phrases, here are some examples (upon pressing the play button the phrase is said in Russian):

*Oh fuck - Блять*

*Rush B whore bitch - Раш Б cука блять*

*You make my dick hard with your performance - У меня встал от твоей игры*

&amp;#x200B;

This is how the app looks when you open it:

&amp;#x200B;

[Russian rage phrases for CS:GO](https://preview.redd.it/8gsv35nh3r841.jpg?width=606&amp;format=pjpg&amp;auto=webp&amp;s=65816dafcd08c52007ef2697eb88024c0b32d561)

**Basically, these Ads are my only source of income from these apps. If I want to restore the AdMob serving to my app, what choices do I have?**

**Will it be enough to blur curse words (whole or just one letter)? I don't want to lose user experience.**

**Are there any alternatives to AdMob that does not have such restrictions?**

&amp;#x200B;

For those, who are interested, here are the links to these apps:

mild version:  [https://play.google.com/store/apps/details?id=ether.paul.lt.russiansphrasesforcsgo](https://play.google.com/store/apps/details?id=ether.paul.lt.russiansphrasesforcsgo)

rage version:  [https://play.google.com/store/apps/details?id=ether.paul.lt.russiansragephrasesforcsgo](https://play.google.com/store/apps/details?id=ether.paul.lt.russiansragephrasesforcsgo)
## [9][I have created a tutorial for Recycler View Animation using Custom Layout Manager in Kotlin.](https://www.reddit.com/r/androiddev/comments/ejd9m6/i_have_created_a_tutorial_for_recycler_view/)
- url: https://youtu.be/7gnBZ_fitdE
---

## [10][ML Kit face identification](https://www.reddit.com/r/androiddev/comments/ejuw5w/ml_kit_face_identification/)
- url: https://www.reddit.com/r/androiddev/comments/ejuw5w/ml_kit_face_identification/
---
Hey everyone! 

I need to identify people with face detection in my app. I haven't been able to find such thing in ml docs. 

Is there any way to identify and compare different faces with ml kit?

Like for example does ml provide an ID for each individual face?
## [11][Compile c source code with iconv](https://www.reddit.com/r/androiddev/comments/ejujw2/compile_c_source_code_with_iconv/)
- url: https://www.reddit.com/r/androiddev/comments/ejujw2/compile_c_source_code_with_iconv/
---
I am trying to compile c source code with iconv for Android. But getting error to reinstall iconv. 

How to solve it ?
## [12][How to create a custom music database like in Poweramp (my goal is get rid of MediaStore)?](https://www.reddit.com/r/androiddev/comments/ejudns/how_to_create_a_custom_music_database_like_in/)
- url: https://www.reddit.com/r/androiddev/comments/ejudns/how_to_create_a_custom_music_database_like_in/
---
Poweramp has its own databse with albums/tracks/artists. It detects file changes, links tracks with albums, with artists etc.

I want stop using MediaStore because it is buggy and many users complains about their music not updated, not displaying and so on.
