# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - June 01, 2020](https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/)
- url: https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/
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
## [3][A small Pokedex project using Dagger Hilt, Motion animations, Jetpack based on MVVM architecture.](https://www.reddit.com/r/androiddev/comments/gukp2e/a_small_pokedex_project_using_dagger_hilt_motion/)
- url: https://i.redd.it/8r8v0adkwa251.gif
---

## [4][Retrofit CallAdapter for Coroutines to handle response as states](https://www.reddit.com/r/androiddev/comments/gv518p/retrofit_calladapter_for_coroutines_to_handle/)
- url: https://www.reddit.com/r/androiddev/comments/gv518p/retrofit_calladapter_for_coroutines_to_handle/
---
Hi guys, taking inspiration from this interesting [article](https://proandroiddev.com/create-retrofit-calladapter-for-coroutines-to-handle-response-as-states-c102440de37a), I created a sample Android application to handle response from **Retrofit** \+ **Coroutines** as states. I must say that everything is clean and easy to manage.

Source code: [https://github.com/auron567/Retrofit2CoroutinesCallAdapter](https://github.com/auron567/Retrofit2CoroutinesCallAdapter)
## [5][icons not full](https://www.reddit.com/r/androiddev/comments/gv71jr/icons_not_full/)
- url: https://www.reddit.com/r/androiddev/comments/gv71jr/icons_not_full/
---
I've tried putting legacy icons but I don't have it on my phone for some reason (umidigi power 3) and the icons look smaller so is there other ways to fix this?
## [6][Offering 50$ for solving this Android Bundle bug](https://www.reddit.com/r/androiddev/comments/gv6glx/offering_50_for_solving_this_android_bundle_bug/)
- url: https://www.reddit.com/r/androiddev/comments/gv6glx/offering_50_for_solving_this_android_bundle_bug/
---
See here:

&amp;#x200B;

[https://stackoverflow.com/questions/62148182/android-app-bundle-kills-google-play-games-integration-working-with-plain-old](https://stackoverflow.com/questions/62148182/android-app-bundle-kills-google-play-games-integration-working-with-plain-old)
## [7][Does anyone know what does this mean of console pricing templates : "List includes countries where users make payments using local currency. In the 59 other countries where you distribute your app, your price in JPY will be used instead."](https://www.reddit.com/r/androiddev/comments/gv662l/does_anyone_know_what_does_this_mean_of_console/)
- url: https://www.reddit.com/r/androiddev/comments/gv662l/does_anyone_know_what_does_this_mean_of_console/
---
With Google Play we can distribute apps in 151 countries, when we set pricing template we only set prices for some 92 countries, for the rest of the 59 countries our main template price is set instead.

*"List includes countries where users make payments using local currency. In the 59 other countries where you distribute your app, your price in JPY will be used instead."*

When they say,"your price in JPY" do they mean they will show the Japanese price **and Japanese currency.**

I ask this question because if someone is not familiar with JPY they might be reluctant to pay. If this is the case how may I change this to a more familiar currency such as a USD.
## [8][If you're someone who adds data manually in Firestore backend, I found a very useful thing accidentally!](https://www.reddit.com/r/androiddev/comments/guo5pd/if_youre_someone_who_adds_data_manually_in/)
- url: https://www.reddit.com/r/androiddev/comments/guo5pd/if_youre_someone_who_adds_data_manually_in/
---
I have a wallpaper module in my app, to which I add Wallpapers manually in Firestore backend and which has similar kind of data but only with different keys.

It gets very frustrating to create the whole model class again with the same key pairs, and there's a way to automate this on Google Cloud Console. **Yes, you read that right, this thing is available on Google Cloud Platform and not Firebase.**

Here's how you avoid the headache of manually creating the same data model again and again.

1. Go to [https://console.cloud.google.com](https://console.cloud.google.com)
2. From the top left corner, select the project that you have your Firestore database in. Here's where you select

[Click on the blacked out portion and select the project](https://preview.redd.it/njisn0ycub251.png?width=560&amp;format=png&amp;auto=webp&amp;s=84f84d5f02677436d3fbd4eee169e55486152b0d)

3. After selecting the project, from the left navigation panel, select Firestore -&gt; Data

https://preview.redd.it/1453o57vub251.png?width=361&amp;format=png&amp;auto=webp&amp;s=719a7fdd5e8183d664fde1bf9b40794704c2ab03

4. It will load your Firestore database (it'll look just like Firebase console).

5. But here's where the magic happens:

[See that copy button?](https://preview.redd.it/t1x2awpovb251.png?width=663&amp;format=png&amp;auto=webp&amp;s=c18a8bf5c762fafd6193c47486b2e2b01ae1a27c)

Clicking on that button brings up a new window, where it copies this data with a new random document ID.

Look at the window below and you'll get the idea.

https://preview.redd.it/iqo8ype5wb251.png?width=1146&amp;format=png&amp;auto=webp&amp;s=bf7809c48c5ac08c6fc37385379d46005f028267

You can now save the data with new values, and choose whether to add another document with "Save &amp; Add another" button.

And people at Firebase, when are you adding this in the console itself?
## [9][Developed an app using Flutter and Firebase, for exploring nearby restaurants and checking out their reviews. Also includes a 'Likes section' where users liked restaurants are displayed!](https://www.reddit.com/r/androiddev/comments/gv5iml/developed_an_app_using_flutter_and_firebase_for/)
- url: https://github.com/ahmedgulabkhan/Foodspace
---

## [10][installing android studio on windows 10 problem](https://www.reddit.com/r/androiddev/comments/gv4zxx/installing_android_studio_on_windows_10_problem/)
- url: https://www.reddit.com/r/androiddev/comments/gv4zxx/installing_android_studio_on_windows_10_problem/
---
i tried to install android studio on windows 10 but i keep getting the same error as shown in the picture . i searched for the sdk folder inside appdata and it is not there.

&amp;#x200B;

https://preview.redd.it/24ml8lc47h251.png?width=1366&amp;format=png&amp;auto=webp&amp;s=0e21badad7b41a19b769adb37874a3c30bd87b64
## [11][A library which draws dividers in a RecyclerView (it supports LinearLayoutManager, GridLayoutManager and StaggeredGridLayoutManager) with one line of code](https://www.reddit.com/r/androiddev/comments/guwsiz/a_library_which_draws_dividers_in_a_recyclerview/)
- url: https://www.reddit.com/r/androiddev/comments/guwsiz/a_library_which_draws_dividers_in_a_recyclerview/
---
[https://github.com/fondesa/recycler-view-divider](https://github.com/fondesa/recycler-view-divider)

Hi guys, here is my library which can be used to draw automatically dividers between the items of a `RecyclerView`. It supports `LinearLayoutManager`, `GridLayoutManager` and `StaggeredGridLayoutManager`.

The dividers by default pick the configuration from the theme, but they can be easily configured runtime for further customizations.

The default configuration can be attached with `recyclerView.addDivider()` and the library will handle the rest.

The library is pretty mature, but if you have any questions or feedbacks, I'm here :)
## [12][[HELP] What should I learn Ruby on rails or React js ?](https://www.reddit.com/r/androiddev/comments/gv4bm6/help_what_should_i_learn_ruby_on_rails_or_react_js/)
- url: https://www.reddit.com/r/androiddev/comments/gv4bm6/help_what_should_i_learn_ruby_on_rails_or_react_js/
---
I am and android dev and my manager is asking me to learn any of this. ( ROR or React js)  


I do know what this technologies are there for, but I just want to know if someone was in this situation before and how did it turn out. 

Any advice would be great.

Thanks a lot in advance.
