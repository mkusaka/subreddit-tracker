# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/i745ei/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/i745ei/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - August 10, 2020](https://www.reddit.com/r/androiddev/comments/i72odw/weekly_questions_thread_august_10_2020/)
- url: https://www.reddit.com/r/androiddev/comments/i72odw/weekly_questions_thread_august_10_2020/
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
## [3][(Self-Taught) Android Devs of Reddit: Show your portfolio](https://www.reddit.com/r/androiddev/comments/i83su4/selftaught_android_devs_of_reddit_show_your/)
- url: https://www.reddit.com/r/androiddev/comments/i83su4/selftaught_android_devs_of_reddit_show_your/
---
Hello there! (*General Kenobi!)*

As a self-taught, future super awesome Android developer :D I am really curious about your portfolios and especially those of self-taught and more especially that got you the job :)

&amp;#x200B;

Show off what you got and feel free to add some comments :)
## [4][Gallerit, a sample Reddit gallery](https://www.reddit.com/r/androiddev/comments/i8amd7/gallerit_a_sample_reddit_gallery/)
- url: https://www.reddit.com/r/androiddev/comments/i8amd7/gallerit_a_sample_reddit_gallery/
---
&amp;#x200B;

https://preview.redd.it/zmtij01whjg51.png?width=2560&amp;format=png&amp;auto=webp&amp;s=8083011caaf4c3822ccbab2e3be568caea18b240

Hi guys , I created a small Android application to search images on Reddit. Just search for a subreddit and the top images will be shown (they can also be saved). The goal of the project is to demonstrate best practices using modern Android development tools.

Project characteristics:

* 100% Kotlin
* MVVM
* Repository Pattern
* A single activity architecture (Navigation component)
* Android Jetpack
* Coroutines &amp; Flow
* Reactive UI
* Testing
* Dependency Injection (Koin)
* Gradle Kotlin DSL
* Material Design

Hope it can be of help or inspiration for someone!

Gallerit: [https://github.com/auron567/Gallerit](https://github.com/auron567/Gallerit)
## [5][FCM and other Firebase services on Huawei phones: What's the latest situation?](https://www.reddit.com/r/androiddev/comments/i8as4l/fcm_and_other_firebase_services_on_huawei_phones/)
- url: https://www.reddit.com/r/androiddev/comments/i8as4l/fcm_and_other_firebase_services_on_huawei_phones/
---
In the rapidly evolving situation with Huawai's divorce from Google Play Services, has anyone had any luck converting an app that relies on Firebase Cloud Messaging to receive messages on newer Huawei phones that don't have Play Services on them? 

If so, what were your experiences? Did you end up changing your backend to emit Huawei Push Kit messages or is there some hack that lets you continue to register for Firebase? Looking around on the Web there are a number of discussions from this summer suggesting that the trouble starts when you notice one day that your app is no longer obtaining Firebase IDs...
## [6][Records and Pattern Matching for Instanceof Finalized in JDK 16](https://www.reddit.com/r/androiddev/comments/i8b3mu/records_and_pattern_matching_for_instanceof/)
- url: https://www.infoq.com/news/2020/08/java16-records-instanceof/
---

## [7][How do you deal with Multiple Themes?](https://www.reddit.com/r/androiddev/comments/i8c159/how_do_you_deal_with_multiple_themes/)
- url: https://www.reddit.com/r/androiddev/comments/i8c159/how_do_you_deal_with_multiple_themes/
---
Hi, 

I was wondering how other developers deal with having multiple themes in their apps, now that   
Material Components and DayNight are widely used?

It's easy when you have a light and a dark version, since you can handle everything with minimal code through AppCompatDelegate. But let's say you want light, dark and AMOLED. Or a couple of variations for each. What is the best approach?

This is how I did:

\- I have my main theme: MyStyle, inheriting from DayNight, representing light and dark.

\- I have another one: MyStyle\_amoled, inheriting from MyStyle, with attributes override only in the dark version to make it AMOLED.

This way it follows the system toggle when changing from light to dark, and in the settings, the user can choose the type of dark mode to get, Default or AMOLED.

This is ugly IMO, since you still have to make a base activity class overriding OnCreate() and OnResume() and use setTheme() accordingly. And your user clicking the back button will see the theme changing for a split second as it recreate().

Is there a better way?

It would be nice to manage themes through qualifiers having: -light, -night, -amoled, -version3... folders in your res/ and then just bind the qualifier you choose to the system setting.
## [8][Hello guys! I have a Question related to AlarmManager and Notifications.](https://www.reddit.com/r/androiddev/comments/i8aifq/hello_guys_i_have_a_question_related_to/)
- url: https://www.reddit.com/r/androiddev/comments/i8aifq/hello_guys_i_have_a_question_related_to/
---
I am working on an app that will send notifications to the user at a specific point in time defined by the user. So it will remind a user with 15 minutes before an event is happening. I've tried using AlarmManager and a BroadcastReceiver to send the notification, but it doesn't work when the app is closed. Do you have any idea or recommendation on what I should check?
## [9][What CPC and eCPM does Admob have for USA?](https://www.reddit.com/r/androiddev/comments/i89eyi/what_cpc_and_ecpm_does_admob_have_for_usa/)
- url: https://www.reddit.com/r/androiddev/comments/i89eyi/what_cpc_and_ecpm_does_admob_have_for_usa/
---
I removed Admob an year ago because they suddenly started giving me very bad CPC ($0.05-0.1) for USA.

I use FAN and here's what I get for USA users (Android only):

* Interstitials

CPC = $0.54

eCPM = $9.12


* Banner

CPC = $0.27

eCPM = $0.49

I was wondering what Admob provides to see if FAN is paying good?
## [10][Assets or Resource Raw folder of Android?](https://www.reddit.com/r/androiddev/comments/i8brko/assets_or_resource_raw_folder_of_android/)
- url: https://medium.com/mobile-app-development-publication/assets-or-resource-raw-folder-of-android-5bdc042570e0?source=friends_link&amp;sk=708031a2e38cc7b9d681553ba7a6dd66
---

## [11][Why you should embrace Dependency Injection](https://www.reddit.com/r/androiddev/comments/i7w1bi/why_you_should_embrace_dependency_injection/)
- url: https://codingtroops.com/android/why-do-we-need-dependency-injection-in-our-lives/
---

## [12][I challenge to recreate this.](https://www.reddit.com/r/androiddev/comments/i8dlmi/i_challenge_to_recreate_this/)
- url: https://www.reddit.com/r/androiddev/comments/i8dlmi/i_challenge_to_recreate_this/
---
😌

https://reddit.com/link/i8dlmi/video/f5tcv319okg51/player
