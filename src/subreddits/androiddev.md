# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - January 27, 2020](https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/
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
## [3][Why is our rating from "Unknown" app version in Google Play Console so low](https://www.reddit.com/r/androiddev/comments/ev3u3j/why_is_our_rating_from_unknown_app_version_in/)
- url: https://www.reddit.com/r/androiddev/comments/ev3u3j/why_is_our_rating_from_unknown_app_version_in/
---
In Google Play Console/User Feedback/Ratings this statistics is presented:

As you can see the "Unknown" app version has significantly lower rating than the rest, and it's quite a large amount of ratings. Does anyone have similar statistics?

It might be that Unknown are people who didn't like my game and uninstalled it, but it could also be that some hack-bots are involved?

https://preview.redd.it/eoc4bf3qwhd41.png?width=919&amp;format=png&amp;auto=webp&amp;s=b22943110dae3e4d265a5093545284e1afb6543b
## [4][How to build secure Android Apps: the tough job everybody ignores](https://www.reddit.com/r/androiddev/comments/euq23j/how_to_build_secure_android_apps_the_tough_job/)
- url: https://quickbirdstudios.com/blog/android-app-security-best-practices/?utm_source=reddit.com
---

## [5][How would you plan a big new project with the intention of moving to Jetpack Compose later?](https://www.reddit.com/r/androiddev/comments/ev5kj4/how_would_you_plan_a_big_new_project_with_the/)
- url: https://www.reddit.com/r/androiddev/comments/ev5kj4/how_would_you_plan_a_big_new_project_with_the/
---
* What Jetpack components would you use?
* What kind of architecture would you use?  ex:  MVI or MVVM

Would you do things same way you do now and refactor later little by little since it is incrementally adoptable?
## [6][Has anyone tested Admob banner refresh rates? Is "Google optimized refresh rate" really the best one?](https://www.reddit.com/r/androiddev/comments/ev5hf2/has_anyone_tested_admob_banner_refresh_rates_is/)
- url: https://www.reddit.com/r/androiddev/comments/ev5hf2/has_anyone_tested_admob_banner_refresh_rates_is/
---

## [7][How to effectively manage your multi-service on-demand business using a GoJek clone app?](https://www.reddit.com/r/androiddev/comments/ev3zu2/how_to_effectively_manage_your_multiservice/)
- url: https://www.codemade.io/how-to-effectively-manage-your-multi-service-on-demand-business-using-a-gojek-clone-app/
---

## [8][Testing AndroidX Room + Kotlin Coroutines](https://www.reddit.com/r/androiddev/comments/euy4e7/testing_androidx_room_kotlin_coroutines/)
- url: https://medium.com/@eyalg/testing-androidx-room-kotlin-coroutines-2d1faa3e674f
---

## [9][Up to what "layer" should DI be used?](https://www.reddit.com/r/androiddev/comments/ev3bip/up_to_what_layer_should_di_be_used/)
- url: https://www.reddit.com/r/androiddev/comments/ev3bip/up_to_what_layer_should_di_be_used/
---
I'm new to using DI (Dagger 2) so my question might be a bit confusing at first, let me explain.

My app architecture looks something like this:

Repository(Retrofit, Api, Gson, Storage1, Storage 2..) -&gt; ViewModel(Repository1, Repository2...) -&gt; View(ViewModel)

So View's only dependency is a ViewModel, which can have multiple Repositories as dependencies. Each repository then has 3 fixed dependencies and optional Storage dependencies.

Let me explain the reason for 3 fixed dependencies of the Repository:

* Retrofit is needed so that the instance of \`responseBodyConverter\` can be retreived in case we need to parse error body of a response.
* Api is an instance of Retrofit interface for making API calls.
* Gson is needed to prpare JSON body when making certain requests.

Storage is only added when Repository also uses local storage to store API responses and vice versa. This is only needed in very rare cases as most of the app only shows live data and doesn't care about working offline. 

My main question is at what "layer" here can I stop using DI and decide to just instantiate objects inside the class?

I'm currently dealing with this when writing my Storage classes. My (base) Storage class currently has 3 dependencies:

* SharedPreferences for storing data
* Gson for turning objects into JSON before saving them as string in SharedPreferences
* LoginManager for seeing the currently logged in user when dealing with "user specific" storage

Maybe I'm just overthinking here because I'm new to DI so having constructors with 3+ arguments like this is simply weird to me. Is it normal to have structure like this or am I over-engineering or simply misunderstanding some concepts?
## [10][What is the info button called?](https://www.reddit.com/r/androiddev/comments/ev2yzy/what_is_the_info_button_called/)
- url: https://www.reddit.com/r/androiddev/comments/ev2yzy/what_is_the_info_button_called/
---
I've been searching for the "info button" for quite a while but I can't find its specific name. I'm talking about the button with "i" which shows a "bubble" with text on the same screen (not dialog). I'd really appreciate if some can tell me what it's called or how it's implemented. Thanks.
## [11][Ad Network Certified by Google play description confusion](https://www.reddit.com/r/androiddev/comments/ev2lve/ad_network_certified_by_google_play_description/)
- url: https://www.reddit.com/r/androiddev/comments/ev2lve/ad_network_certified_by_google_play_description/
---
So in their description for trusted ad networks/ad networks certified by Google play it says:  


&gt; The use of certified ad networks applies for all apps in the Designed for Families program and apps outside the Designed for Families program that serve ads to children. ***Apps that are not for children (e.g. apps with content ratings of “T” or “MA”) or apps that do not serve ads to children*** are not required to use a certified ad network. Apps that include children and adults in their target audience are allowed to use non-certified ad network SDKs for serving ads only to users above the age of 13. For example, apps that have both audiences may implement a [neutral age screen](https://support.google.com/googleplay/android-developer/answer/9285070#neutral-agescreen) to be compliant.  

The bolded text, it seems a bit confusing, what about apps rated E or Pegi 3 Pegi 7.  
Does that mean that apps rated E must use only listed ad networks?
## [12][What's wrong with the BiometricPrompt dialog with phones that have a cutout?](https://www.reddit.com/r/androiddev/comments/euwyxr/whats_wrong_with_the_biometricprompt_dialog_with/)
- url: https://www.reddit.com/r/androiddev/comments/euwyxr/whats_wrong_with_the_biometricprompt_dialog_with/
---
Why dim does not appear below the cutout ([1 screenshot](https://i.imgur.com/CksmaG8.png))?  However, everything works fine after hiding the cutout in the developer options ([2 screenshot](https://i.imgur.com/XkxO6mO.png)).Can I fix it?

[https://imgur.com/a/fV2cmGq](https://imgur.com/a/fV2cmGq)

Device: **Pixel 3 XL**

Lib version: **1.0.1** ([androidx.biometric:biometric](https://developer.android.com/jetpack/androidx/releases/biometric))
