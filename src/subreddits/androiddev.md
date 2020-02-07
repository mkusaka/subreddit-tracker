# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/f09xxs/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/f09xxs/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - February 03, 2020](https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/
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
## [3][Android 11 will add a new App Compatibility feature to ease testing](https://www.reddit.com/r/androiddev/comments/f081sw/android_11_will_add_a_new_app_compatibility/)
- url: https://www.developer-tech.com/news/2020/jan/22/android-11-new-app-compatibility-feature-testing/
---

## [4][All-in-One — Android TabLayout and TabItem](https://www.reddit.com/r/androiddev/comments/f09nss/allinone_android_tablayout_and_tabitem/)
- url: https://medium.com/@myrickchow32/android-tablayout-and-tabitem-268ac06ba966?source=friends_link&amp;sk=65ae680fe463e95d0c7301fc64cc057d
---

## [5][Huawei, Xiaomi, Oppo and Vivo creating a Google Play Store alternative](https://www.reddit.com/r/androiddev/comments/ezufny/huawei_xiaomi_oppo_and_vivo_creating_a_google/)
- url: https://www.theverge.com/2020/2/6/21126118/huawei-xiaomi-vivo-oppo-app-store-platform-google-play-dominance-worldwide
---

## [6][Changing minSdkVersion from 21 to 23?](https://www.reddit.com/r/androiddev/comments/f09jni/changing_minsdkversion_from_21_to_23/)
- url: https://www.reddit.com/r/androiddev/comments/f09jni/changing_minsdkversion_from_21_to_23/
---
What is your minSdkVersion? Right now in our apps it's mostly API 21, but we are considering to move to API 23, no specific reason except some code optimization and possibly APK size reduction.   


According to analytics 13% of our users are on API level 21, but here is the catch, 69% of them are from India and Mexico and those countries tend to have very low eCPM and are not very profitable.   
What are your thoughts?  
What is your minSdkVersion?
## [7][Implementing Scoped Storage in Android 10](https://www.reddit.com/r/androiddev/comments/f05rzn/implementing_scoped_storage_in_android_10/)
- url: https://heartbeat.fritz.ai/implementing-scoped-storage-in-android-10-4b657280c066
---

## [8][Android UI Testing with Azure DevOps](https://www.reddit.com/r/androiddev/comments/f0a47i/android_ui_testing_with_azure_devops/)
- url: https://www.reddit.com/r/androiddev/comments/f0a47i/android_ui_testing_with_azure_devops/
---
I recently published an article, for those of you who need to configure a pipeline for Android Instrumented Tests within Azure DevOps! :)

[https://medium.com/genetec-tech/android-ui-testing-in-azure-devops-81bbe7cea9fd](https://medium.com/genetec-tech/android-ui-testing-in-azure-devops-81bbe7cea9fd)
## [9][Tips to land first job as Jr](https://www.reddit.com/r/androiddev/comments/f09owu/tips_to_land_first_job_as_jr/)
- url: https://www.reddit.com/r/androiddev/comments/f09owu/tips_to_land_first_job_as_jr/
---
What you guys recommend doing to land a first job as a Jr Android Dev without college and professional experience?

I live in Spain (if that matters), and I thought it would be easier, I know a lot of tools the industry uses and have more than 2 years of experience with Android.

My only plan is to 1) Get the Google certificate, as it would validate my knowledge without going to college. 2) Build some open source projects to have a portfolio to show. 3) Maybe lie about professional experience, as I have seen people saying that they only put that requirement to weed people out. But I really don't like lying..

Do you guys have any tips? Any thing I am missing out? How did you get your first job?
## [10][Is it okay to play video at app first launch?](https://www.reddit.com/r/androiddev/comments/f08tud/is_it_okay_to_play_video_at_app_first_launch/)
- url: https://www.reddit.com/r/androiddev/comments/f08tud/is_it_okay_to_play_video_at_app_first_launch/
---
My company planning to show a video on app first launch. I just wondering if it's violate google play policy or something?
## [11][[Help] Query regarding Android Keystore system.](https://www.reddit.com/r/androiddev/comments/f05fg6/help_query_regarding_android_keystore_system/)
- url: https://www.reddit.com/r/androiddev/comments/f05fg6/help_query_regarding_android_keystore_system/
---
Hello everyone,I want to send a file, say "secret.txt", from my server to my client app. But I want this to be completely secured to the degree in which even on a rooted phone user can't see what's inside secret.txt. It will be a simple file in fs but encrypted. I don't know how Android keystore system works to achieve this.

I want my app to access this data pragmatically in such a way that the app won't know what the key is before the server sends the file.

I've figured out some approaches but all of them are flawed:

1. Storing decryption key in code in the apk. (Can be easily extracted?)
2. Sending random key from client to server, so I can send the data encrypted with that key. (Random key can be seen using apps such as Fiddler/Wireshark/[Canary](https://play.google.com/store/apps/details?id=com.guoshi.httpcanary).)

Can anybody please help me solving this problem?
## [12][Game over for Android?](https://www.reddit.com/r/androiddev/comments/f0aags/game_over_for_android/)
- url: https://bgr.com/2020/02/04/android-vs-iphone-google-earnings-show-android-needs-to-die/
---

