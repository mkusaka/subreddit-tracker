# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/hqeluj/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/hqeluj/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - July 13, 2020](https://www.reddit.com/r/androiddev/comments/hqd7s0/weekly_questions_thread_july_13_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hqd7s0/weekly_questions_thread_july_13_2020/
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
## [3][A Rick and Morty app tutorial using Hilt, MVVM, LiveData, Kotlin Coroutines and Repository pattern :)](https://www.reddit.com/r/androiddev/comments/hqsy98/a_rick_and_morty_app_tutorial_using_hilt_mvvm/)
- url: https://medium.com/@sberoch/android-architecture-hilt-mvvm-kotlin-coroutines-live-data-room-and-retrofit-ft-8b746cab4a06
---

## [4][How to plan a refactor for better architecture](https://www.reddit.com/r/androiddev/comments/hqzwrf/how_to_plan_a_refactor_for_better_architecture/)
- url: https://www.reddit.com/r/androiddev/comments/hqzwrf/how_to_plan_a_refactor_for_better_architecture/
---
The current structure of my app is as follows:

* One activity and multiple fragments (about 25 fragments)
* One huge monolithic repository
* One huge monolithic viewmodel
* Databinding is being used but with one huge monolithic handler class
* Navigation Component with one nav graph for the entire application

I've been learning as I go and I now realize my architecture is very poor. I want to achieve proper MVVM and also to modularize the project. My end goal is to have an architecture that is more maintainable and testable. I'm aiming for the following:

* Separate viewmodels for each fragment
* Separate repositories for each entity (or roughly each data access object)
* Separate data binding handlers for each fragment
* The entire project split into different library modules, with each module being specific to a feature or core function

In order to achieve the above outcome, is the following a reasonable plan?

1. Create separate viewmodels for each fragment
2. Create separate repositories for each entity
3. Create separate handler classes for each fragment
4. Divide the entire project into separate modules

Some of the sources I've come across while researching are:

* [https://developer.android.com/jetpack/guide](https://developer.android.com/jetpack/guide)
* [https://proandroiddev.com/intro-to-app-modularization-42411e4c421e](https://proandroiddev.com/intro-to-app-modularization-42411e4c421e)
* [https://stackoverflow.com/questions/44375276/when-using-mvvm-on-android-should-each-activity-have-one-and-only-one-viewmod?rq=1](https://stackoverflow.com/questions/44375276/when-using-mvvm-on-android-should-each-activity-have-one-and-only-one-viewmod?rq=1)
* [https://stackoverflow.com/questions/51263301/multi-module-navigation-with-architecture-components](https://stackoverflow.com/questions/51263301/multi-module-navigation-with-architecture-components)
* [https://medium.com/swlh/using-the-navigation-component-in-a-modular-world-e7578825962](https://medium.com/swlh/using-the-navigation-component-in-a-modular-world-e7578825962)
* [https://stackoverflow.com/questions/50466743/should-we-create-a-separate-repository-for-each-activity-or-single-repository-fo](https://stackoverflow.com/questions/50466743/should-we-create-a-separate-repository-for-each-activity-or-single-repository-fo)

Is the above rough plan a good one? Are there any pitfalls I should look out for?
## [5][Lightweight android library for image compression](https://www.reddit.com/r/androiddev/comments/hqod1r/lightweight_android_library_for_image_compression/)
- url: https://github.com/jeziellago/image-minifier
---

## [6][How to Monetize Hyper-Casual Games with Success](https://www.reddit.com/r/androiddev/comments/hqzfiw/how_to_monetize_hypercasual_games_with_success/)
- url: https://www.appodeal.com/home/blog/monetize-hyper-casual-games-success/
---

## [7][How to create expandable ViewGroup like this?](https://www.reddit.com/r/androiddev/comments/hqylxt/how_to_create_expandable_viewgroup_like_this/)
- url: https://v.redd.it/wkl3nrblmsa51
---

## [8][Android assistant - MediaBrowserService - Voice command - “Play x on appName” not working](https://www.reddit.com/r/androiddev/comments/hqxoqz/android_assistant_mediabrowserservice_voice/)
- url: https://www.reddit.com/r/androiddev/comments/hqxoqz/android_assistant_mediabrowserservice_voice/
---
I'm having trouble with the Play \[song\] on \[appName\] command; specifically the "app" is not being recognised by google assistant.

&amp;#x200B;

I followed the instruction form [https://developer.android.com/guide/topics/media-apps/interacting-with-assistant](https://developer.android.com/guide/topics/media-apps/interacting-with-assistant) and the demo app form [https://github.com/android/uamp](https://github.com/android/uamp)

My AndroidManifest:

&amp;#x200B;

&lt;?xml version="1.0" encoding="utf-8"?&gt;

&lt;manifest xmlns:android="[http://schemas.android.com/apk/res/android](http://schemas.android.com/apk/res/android)"

xmlns:tools="[http://schemas.android.com/tools](http://schemas.android.com/tools)"

package="com.xx"&gt;

&amp;#x200B;

&lt;uses-permission android:name="android.permission.FOREGROUND\_SERVICE" /&gt;

&lt;uses-permission android:name="android.permission.INTERNET" /&gt;

&lt;uses-permission android:name="android.permission.ACCESS\_NETWORK\_STATE" /&gt;

&lt;uses-permission android:name="android.permission.ACCESS\_FINE\_LOCATION" /&gt;

&lt;uses-permission android:name="android.permission.ACCESS\_COARSE\_LOCATION" /&gt;

&amp;#x200B;

&lt;application

android:name=".XApplication"

android:allowBackup="false"

android:icon="@mipmap/ic\_launcher"

android:label="@string/app\_name" // app\_name = appX 101.1

android:roundIcon="@mipmap/ic\_launcher\_round"

android:supportsRtl="false"

android:theme="@style/AppTheme"&gt;

&amp;#x200B;

&lt;activity

android:name=".ui.splash.SplashActivity"

android:exported="true"

android:launchMode="singleTop"

android:screenOrientation="portrait"

android:theme="@style/AppTheme.AppFullScreenTheme"

android:windowSoftInputMode="adjustPan|stateHidden"&gt;

&lt;intent-filter&gt;

&lt;action android:name="android.intent.action.MAIN" /&gt;

&lt;category android:name="android.intent.category.LAUNCHER" /&gt;

&lt;/intent-filter&gt;

&lt;intent-filter&gt;

&lt;action android:name="android.media.action.MEDIA\_PLAY\_FROM\_SEARCH" /&gt;

&lt;category android:name="android.intent.category.DEFAULT" /&gt;

&lt;/intent-filter&gt;

&lt;/activity&gt;

...

...

&lt;service

android:name="com.music.android.xx.media.MusicService"

android:enabled="true"

android:exported="true"

tools:ignore="ExportedService"&gt;

&lt;intent-filter&gt;

&lt;action android:name="android.media.browse.MediaBrowserService" /&gt;

&lt;/intent-filter&gt;

&lt;/service&gt;

&lt;service

android:name=".service.LocationUpdatesService"

android:foregroundServiceType="location" /&gt;

&amp;#x200B;

&lt;receiver android:name=".service.StopServiceReceiver" /&gt;

&lt;receiver android:name=".service.PlayRadioReceiver" /&gt;

&amp;#x200B;

&lt;service

android:name=".service.MyFirebaseService"

android:exported="false"&gt;

&lt;intent-filter&gt;

&lt;action android:name="com.google.firebase.MESSAGING\_EVENT" /&gt;

&lt;/intent-filter&gt;

&lt;/service&gt;

&lt;/application&gt;

&lt;/manifest&gt;

&amp;#x200B;

I have a problem figuring out what I did wrong or what I missed . Any help will be greatly appreciated..

PS: My app had been in the store for 3days
## [9][How to update an app on a KIOSK device using a launcher.](https://www.reddit.com/r/androiddev/comments/hqzozd/how_to_update_an_app_on_a_kiosk_device_using_a/)
- url: https://www.reddit.com/r/androiddev/comments/hqzozd/how_to_update_an_app_on_a_kiosk_device_using_a/
---
I have three different tablets (no-name brand) that will only run my app in a KIOSK mode. I want to update the app on these devices.

So far my progress:

\- I have a "Launcher" application that starts and handles the updating of the "Main" application.

\- The Launcher app receives the new APK over HTTP and saves it in its getFilesDir() (/data/data/packagename/files/)

My goal at this point to install this APK without prompting the user to allow this installation. AFAIK this is not possible without root access

I tried this:

    Runtime.getRuntime().exec("chmod 777" + apkLocation);
    String command; command = "pm install -r " + apkLocation;
    Process proc = Runtime.getRuntime().exec(new String[]{"su", "-c", command});
    proc.waitFor();

But this fails and returns 1. I assume this is because the application does not have root access.

Also, the application is the active-admin and the device-owner of the device, but still, these permissions don't seem like they gave me the right to do what I mentioned above.

EDIT: My target API version is 24.

EDIT2: Another question. Is it possible to silently install an APK by calling the hidden API using reflection?
## [10][How to see number of app downloads from google play Developer console? Not Graphs, But Total in number](https://www.reddit.com/r/androiddev/comments/hqvb36/how_to_see_number_of_app_downloads_from_google/)
- url: https://www.reddit.com/r/androiddev/comments/hqvb36/how_to_see_number_of_app_downloads_from_google/
---
I have uploaded an app to google play. I want to know how many people already downloads the app. Not active users or any other.  Not graphs. I need to know the total. How to do that?
## [11][Updating our game to Google Play is a nightmare...](https://www.reddit.com/r/androiddev/comments/hqywbq/updating_our_game_to_google_play_is_a_nightmare/)
- url: https://www.reddit.com/r/androiddev/comments/hqywbq/updating_our_game_to_google_play_is_a_nightmare/
---
I am the developer of world-famous board game Tic Tac Toe. 

Recently, I have a problem updating the application. Google bot rejects our updates because of ... well ... using in long description word "COCK"

"Jogo de Galo" is the common name for Tic Tac Toe in Brazil. I hope one day machine will learn or it will be a living hell of political correctness... 

[Tic Tac Toe Club](https://play.google.com/store/apps/details?id=com.mobilearts.game.tictactoe)

https://preview.redd.it/k5lctdw7qsa51.png?width=1310&amp;format=png&amp;auto=webp&amp;s=9804643e3064a140510f4e58a81c3b36c55da67d
## [12][Multiple APK support with Android TV](https://www.reddit.com/r/androiddev/comments/hqvtfp/multiple_apk_support_with_android_tv/)
- url: https://www.reddit.com/r/androiddev/comments/hqvtfp/multiple_apk_support_with_android_tv/
---
I have been using android app bundle to deploy my app, which currently supports phone and tablets, to play store. Now I am working on TV app but as a separate project, but I wish to keep the package id same as my phone app (single play store entry for both apps).

My TV app will generate a separate APK and I was wondering how do I upload the apk to my existing play store application. I read through \[multiple apk support\] ([https://developer.android.com/google/play/publishing/multiple-apks](https://developer.android.com/google/play/publishing/multiple-apks)), but I still have questions.

1. Can I upload both AAB and APK for the same app in play store. I know I can uploaded multiple APKs.

2. How do I clearly mark my APK to be for TV only. Current way of differentiating is using screen size &lt;supports-screens&gt; or &lt;compatible-screens&gt;, which is not really a very clear way of differentiation.
