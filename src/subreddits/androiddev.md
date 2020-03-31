# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/frq4i4/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/frq4i4/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - March 30, 2020](https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/
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
## [3][Introducing dual-screen layouts for Android | Surface Duo Blog](https://www.reddit.com/r/androiddev/comments/fs39j7/introducing_dualscreen_layouts_for_android/)
- url: https://devblogs.microsoft.com/surface-duo/introducing-dual-screen-layouts-android/
---

## [4][BL Taxi is an open-source app built to showcase the latest technologies. It uses Kotlin, Coroutines, Koin, Architecture Components, MVVM, Room, Retrofit, Material Components, and other Jetpack libraries.](https://www.reddit.com/r/androiddev/comments/frvv8s/bl_taxi_is_an_opensource_app_built_to_showcase/)
- url: https://github.com/VladimirWrites/BLTaxi
---

## [5][Is it worth it to follow the Android Kotlin course by Google on Udacity?](https://www.reddit.com/r/androiddev/comments/fsba8g/is_it_worth_it_to_follow_the_android_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/fsba8g/is_it_worth_it_to_follow_the_android_kotlin/
---
I'm a beginner here trying to learn different Jetpack packages such as Architecture and Navigation.

 
I have taken the Android Kotlin Udacity by Google course and I have noticed in their examples that they have used android arch libraries and on the documentation page it has been mentioned androidx. 

I know what AndroidX is but is it okay to replace those examples but using the versions mentioned in the docs? 

Thanks in advance!
## [6][How to test api calls from android app to localhost service?](https://www.reddit.com/r/androiddev/comments/fsa8as/how_to_test_api_calls_from_android_app_to/)
- url: https://www.reddit.com/r/androiddev/comments/fsa8as/how_to_test_api_calls_from_android_app_to/
---
I'm trying to understand how to test http calls from my app (which is running in an emulator) to my backend service, which is running in debug on my pc on https://localhost:8080.

I get that talking to https://localhost:8080 directly from android won't work because it means "emulator's localhost", so I'm talking to my pc local ip, but I got CERTIFICATE_VERIFY_FAILED because of course https://192.168.1.123:8080 doesn't have a signed cert.

What is the correct way to handle this situation?
## [7][Is there anyway I could test an Android application (via Android Studio) without ADB?](https://www.reddit.com/r/androiddev/comments/fs91sl/is_there_anyway_i_could_test_an_android/)
- url: https://www.reddit.com/r/androiddev/comments/fs91sl/is_there_anyway_i_could_test_an_android/
---
Hi all,

I'm currently studying Android developing(personally) using Android Studio. Due to my company's security policy, we must gain permission to execute adb.exe. I think my boss would not allow me to gain this permission, so I'm looking for ways to test out my app. Is this possible? Would there be any way for me to do this?
## [8][Where do you guys find good svgs and image assets for your apps?](https://www.reddit.com/r/androiddev/comments/fs6olp/where_do_you_guys_find_good_svgs_and_image_assets/)
- url: https://www.reddit.com/r/androiddev/comments/fs6olp/where_do_you_guys_find_good_svgs_and_image_assets/
---
I want to make a simple app as my first, however I'm looking for some crisp images assets I can use. Is there a better alternative than using non copyrighted google images?
## [9][Unit testing help](https://www.reddit.com/r/androiddev/comments/fs777k/unit_testing_help/)
- url: https://www.reddit.com/r/androiddev/comments/fs777k/unit_testing_help/
---
I've recently been through u/VasiliyZukanov 's Android unit testing course, but I'm still at a bit of a loss here with how to attack this problem.

I've been trying to get Unit tests into my project at work for a while now which has zero test coverage and severely lacks architecture. Fortunately I'm able to refactor where necessary, although I'm really not sure what to do about the these kind of scenarios of callbacks containing business logic, or even in some cases, nested callbacks which are extremely common in our codebase.

For example how would I possibly approach testing the logic within onAvailablePlansLoaded() or onFail()?

    billingCenter.connect(object :
            billingCenter.BillingConnectionCallback {
        override fun onConnected() {
            billingCenter.loadAvailablePlans(object : BillingCenter.LoadAvailablePlansCallback {
                override fun onAvailablePlansLoaded() {
                    when {
                        billingCenter.availablePlans.isEmpty() -&gt; showError()
                        else -&gt; checkDiscounts()                
                    }
                }
    
                override fun onFail() {
                    showError()
                }
            }
        })

I feel like I'm missing something fundamental here, please enlighten me!
## [10][Is there a library for this style dialog? Looks really nice.](https://www.reddit.com/r/androiddev/comments/fryfih/is_there_a_library_for_this_style_dialog_looks/)
- url: https://i.imgur.com/tWxtZ1c.jpg
---

## [11][How to write in a txt file in my internal storage](https://www.reddit.com/r/androiddev/comments/fs7sw0/how_to_write_in_a_txt_file_in_my_internal_storage/)
- url: https://www.reddit.com/r/androiddev/comments/fs7sw0/how_to_write_in_a_txt_file_in_my_internal_storage/
---
I want to write in a file, for example in 
/storage/emulated/0/My Files/abc.txt

How can I do that while mentioning the path in the code. I want to mention the path in tbe code.
## [12][Why is the emoji on Gboard became inconsistent recently. Before, it was uniform Samsung Emoji across apps. Now it differs per app.](https://www.reddit.com/r/androiddev/comments/fsckc6/why_is_the_emoji_on_gboard_became_inconsistent/)
- url: https://i.redd.it/fzvw5sjkb0q41.jpg
---

