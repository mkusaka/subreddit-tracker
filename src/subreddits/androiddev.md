# androiddev
## [1][App Feedback Thread - June 27, 2020](https://www.reddit.com/r/androiddev/comments/hgsozi/app_feedback_thread_june_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hgsozi/app_feedback_thread_june_27_2020/
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
## [2][Announcement: Three upcoming AMAs with the Android team!](https://www.reddit.com/r/androiddev/comments/hf6l0h/announcement_three_upcoming_amas_with_the_android/)
- url: https://www.reddit.com/r/androiddev/comments/hf6l0h/announcement_three_upcoming_amas_with_the_android/
---
Hey folks! The Android team at Google will be hosting three upcoming AMAs. We've updated the sidebar with details (and will keep updating it with more info as it comes along).

The team there will submit posts for collecting questions ~1 week before answering them, first post will be sometime next week. We'll sticky these posts when they're up!

|   Date  |  Time  |  Who   |   What  | Link |
|:--------:|:-------:|:-------:|:-------:|:-------:|
| 9th July | TBD | Android Team | Android Platform | TBD |
| 30th July | TBD | Android Team | Android Studio | TBD |
| 27th August | TBD | Android Team | Android Jetpack &amp; Jetpack Compose | TBD |
## [3][Hi guys! I've made a simple android game, where you need to identify characters based on their visual clues (they are randomly generated). Github link: https://github.com/rosspetukhov/SuspectMatch. Let me know your thoughts :)](https://www.reddit.com/r/androiddev/comments/hga8rv/hi_guys_ive_made_a_simple_android_game_where_you/)
- url: https://i.redd.it/igkec4vm1a751.png
---

## [4][Why Android Gradle documentation sucks and feels like it is not updated frequently?](https://www.reddit.com/r/androiddev/comments/hghzs8/why_android_gradle_documentation_sucks_and_feels/)
- url: https://www.reddit.com/r/androiddev/comments/hghzs8/why_android_gradle_documentation_sucks_and_feels/
---
I don't know much about Gradle for Android

I am really trying to learn something from this documentation [https://developer.android.com/studio/build](https://developer.android.com/studio/build) but I feel it is just an unordered piece of \*\*\*\* that just mention some "tips" put in some legacy type of documentation

I thought that it can be some just ordered type of documentation when you learn stuff incrementally. But in this type of documentation it is really hard to understand some topic, when it includes several types of information (considered even advanced) from other topic.

What is more I feel that this documentation is not updated frequently

For example:

\- here [https://developer.android.com/studio/build/dependencies](https://developer.android.com/studio/build/dependencies) not mention anything about kapt which is required for Kotlin code

\- here [https://developer.android.com/studio/build/optimize-your-build#optimize](https://developer.android.com/studio/build/optimize-your-build#optimize) mentioning something about Android Studio 2.3 (really?)

&amp;#x200B;

It is really totally pushing me away when try to learn something from it and I give up

Do you guys know some good source from which I can learn Gradle for Android?
## [5][I want to fire off a notification when entering a room.](https://www.reddit.com/r/androiddev/comments/hgs7d6/i_want_to_fire_off_a_notification_when_entering_a/)
- url: https://www.reddit.com/r/androiddev/comments/hgs7d6/i_want_to_fire_off_a_notification_when_entering_a/
---
I basically want to make a function where If you go into a room (doesn't have to be super accurate) a notification can fire.  

I am just not sure which way to go.  I was thinking anything from something like an RFID in the room or could an alexa, echo dot or other device solve this? I would prefer cheaper hardware if possible.

Or could I use just the phone alone to say "when you get 10 feet near this point in space, fire a notification.  I don't need insane accuracy like realtime tracking but I would like to denote if someone goes into one room vs another.  

Is there a platform for this already? Just having a hard time figuring out what would make the most sense to start with.
## [6][Has anyone noticed that Firebase Authentication doesn't even follow Google's one activity guideline. Does it matter?](https://www.reddit.com/r/androiddev/comments/hgimem/has_anyone_noticed_that_firebase_authentication/)
- url: https://www.reddit.com/r/androiddev/comments/hgimem/has_anyone_noticed_that_firebase_authentication/
---
[If you go to the Firebase Authentication guide it directly contradicts having one activity.](https://firebase.google.com/docs/auth/android/phone-auth?fbclid=IwAR0Do4bdXfPA8_YXAigjxguUe2VQ7aWpEPp3xsaFr0cfJYAPrM1YkJx3mT8#kotlin+ktx) It uses PhoneAuthentication as it's own activity? Is this recommended or just outdated, and on top of that they write all their business logic in the Activity. So what's the deal? Auth Management is tricky as hell to manage the "right way" did they also just say screw it throw it all in an activity because it's easier.

https://github.com/firebase/quickstart-android/blob/09b3db061c66499ff8a48d5a1f3d84acf6f56bed/auth/app/src/main/java/com/google/firebase/quickstart/auth/java/PhoneAuthActivity.java#L191-L191
## [7][How to use/control RxJava Observable.cache](https://www.reddit.com/r/androiddev/comments/hgptcg/how_to_usecontrol_rxjava_observablecache/)
- url: https://www.reddit.com/r/androiddev/comments/hgptcg/how_to_usecontrol_rxjava_observablecache/
---
I am trying to use the RxJava caching mechanism but i can't seem to  catch how it works or how can i control the cached contents since there  is the `cache` operator.

I want to verify the cached data with some conditions before emitting the new data.

for example

    someObservable.
    repeat().
    filter { it.age &lt; maxAge }.
    map(it.name).
    cache() 

There is also `ObservableCache&lt;T&gt;` class but i can't find any resources of using it.

how can i get the cached content and verify it (or do any operation on it) before emitting a new value

Any help would be much appreciated. Thanks.
## [8][Was Google measure ar app made in Android studio or unity?](https://www.reddit.com/r/androiddev/comments/hgpscf/was_google_measure_ar_app_made_in_android_studio/)
- url: https://www.reddit.com/r/androiddev/comments/hgpscf/was_google_measure_ar_app_made_in_android_studio/
---

## [9][Android Context Full Guide](https://www.reddit.com/r/androiddev/comments/hga4fe/android_context_full_guide/)
- url: https://www.reddit.com/r/androiddev/comments/hga4fe/android_context_full_guide/
---
We have different types of Android Context, and sometimes developers get confused about which Context to use 🤷‍♂️  
More than that, wrong context usage can lead to problems, for example, memory leaks 😕

Master the Android Context and use it correctly.  
[https://vladsonkin.com/android-context-full-guide/](https://vladsonkin.com/android-context-full-guide/)
## [10][Do you use any mobile attribution partners (Adjust, Appsflyer, Kochava, Branch, Tenjin...)?](https://www.reddit.com/r/androiddev/comments/hgk3jc/do_you_use_any_mobile_attribution_partners_adjust/)
- url: https://www.reddit.com/r/androiddev/comments/hgk3jc/do_you_use_any_mobile_attribution_partners_adjust/
---
We are looking to buy installs from the top non-self attributing video ad networks i.e. UnityAds, Vungle, etc. and are looking for any recommendations or inside knowledge about attribution partners from those that may already be doing so. If you are open to discussing which genre your app/game is in, and any experiences you may have using attribution partners and non-SAN networks for your UA, please feel free to connect with us via DM.
## [11][Android Studio 4.0 removed the option to disable "Include fragment factory methods" and "Include interface callbacks"](https://www.reddit.com/r/androiddev/comments/hfzgue/android_studio_40_removed_the_option_to_disable/)
- url: https://i.redd.it/xvieraqv26751.png
---

## [12][Should I continue to learn native development with kotlin](https://www.reddit.com/r/androiddev/comments/hgcfou/should_i_continue_to_learn_native_development/)
- url: https://www.reddit.com/r/androiddev/comments/hgcfou/should_i_continue_to_learn_native_development/
---
I am quite new to android development and have beem developing apps with kotlin for a 6 moths now.
But I am concerned that when Fuchsia arrives, all the things I have learnt will go useless.
Moreover there's is always a chance that with tools like react native, native app development may not be required.
What are your views regarding the demand of native app development in the future?
