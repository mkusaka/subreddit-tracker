# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/g4s3y9/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/g4s3y9/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - April 20, 2020](https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/
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
## [3][What are we supposed to use instead of AsyncTask ?](https://www.reddit.com/r/androiddev/comments/g59hoh/what_are_we_supposed_to_use_instead_of_asynctask/)
- url: https://www.reddit.com/r/androiddev/comments/g59hoh/what_are_we_supposed_to_use_instead_of_asynctask/
---
The documentation for [AsynchTask](https://developer.android.com/reference/android/os/AsyncTask) is crystal clear: 

This class was deprecated in API level R.  
Use the standard java.util.concurrent or Kotlin concurrency utilities instead.  


However in other articles in the documentation on processes and threading such as [here](https://developer.android.com/guide/components/processes-and-threads#AsyncTask) and [here](https://developer.android.com/topic/performance/threads#asynctask) we are still referred to AsynchTask. What is the recommended way to do some processing in the background telling the UI thread to update every now and then?
## [4][Will anyone be interested to use a Open API that emulates some of the functionality of some social media app to create app for learning?](https://www.reddit.com/r/androiddev/comments/g58i54/will_anyone_be_interested_to_use_a_open_api_that/)
- url: https://www.reddit.com/r/androiddev/comments/g58i54/will_anyone_be_interested_to_use_a_open_api_that/
---
I am thinking to create an API that emulates FB or Instagrams basic feature so mobile or frontend devs can use that API to learn about JWT auth flows and create practice apps to add into their portofolio. Once I am done with the basic functionality I would not mind if someone wanna pitch in their time help me maintain it.
## [5][Kotlin Discord Server, because why not?](https://www.reddit.com/r/androiddev/comments/g5f6jh/kotlin_discord_server_because_why_not/)
- url: https://discord.gg/yrabuRp
---

## [6][Best native Android Development learning resource](https://www.reddit.com/r/androiddev/comments/g5c6q1/best_native_android_development_learning_resource/)
- url: https://www.reddit.com/r/androiddev/comments/g5c6q1/best_native_android_development_learning_resource/
---
Here we go again.
I messed with android studio, java and kotlin before. But soon I got to the point where it just wasn't fun anymore. I have a bit of knowledge and experience in many fields all from web dev, to game dev... Recently I've been messing around with Flutter, but honestly, I don't see future with it. There zero job opportunities in my country with flutter.. So I would like to once more give chance to native Android dev. I've seen jetpack compose, and as I always struggled with making UI in Android Studio and all the constraints and making everything look good on all screen sizes, I think Compose would make difference in that field? To cut the story short, I'm interested in some Udemy course which covers everything and is up to date with the new stuff. Of course I searched myself, but they all look kinda outdated to me. So I need recommendations. Also is Kotlin the right way to go these days? I think most of the companies still use Java? Anyway recommend something you know for sure is good, and it possibly helped you.
## [7][JetNews, a small open-source Jetpack(AndroidX) news app](https://www.reddit.com/r/androiddev/comments/g5bz4p/jetnews_a_small_opensource_jetpackandroidx_news/)
- url: https://www.reddit.com/r/androiddev/comments/g5bz4p/jetnews_a_small_opensource_jetpackandroidx_news/
---
When quarantine takes away your sleep, might as well make something.

It's basically the AndroidX news site in the bottom navigation view.

Why did I make it?CBA to open the web, if you're like me you've come to the right place, plus you might learn a thing or two.

[Github](https://github.com/CraZyLegenD/JetNews)

&amp;#x200B;

Edit: this isn't a compose app nor anyhow related to [https://github.com/android/compose-samples/tree/master/JetNews](https://github.com/android/compose-samples/tree/master/JetNews)
## [8][Button Image Tint change possible?](https://www.reddit.com/r/androiddev/comments/g5enua/button_image_tint_change_possible/)
- url: https://www.reddit.com/r/androiddev/comments/g5enua/button_image_tint_change_possible/
---
Greetings,

Im working on an Android app.  I have a Button that is just using a black vector SVG image on line `android:drawableTop="@drawable/main_report"`. It is colored by a tint in the XML layout.


        &lt;Button
        android:id="@+id/buttonstage6"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:background="@android:color/transparent"
        android:drawableTop="@drawable/main_report"
        android:drawableTint="#3F51B5" &gt;


Im trying to use the OnClick listener event to change the *drawableTint* value to something like red (#FF0000) from within the MainActivity java file.

 I cant figure out how to do this and I've tried and not been able to get it to work. I think I read that once certain values are defined in the XML that it becomes non-changeable. Is that the case?  If not, how should I implement?
## [9][It's just a makeup kit. Had to learn and use flutter for company project (from Dec 2018). This is what I feel now.](https://www.reddit.com/r/androiddev/comments/g4m6w2/its_just_a_makeup_kit_had_to_learn_and_use/)
- url: https://i.redd.it/ffavl8dphwt41.jpg
---

## [10][Problem with loading text from assets file (Problem with reading text files from assets folder (Unmarshaling error, Parcelable class)](https://www.reddit.com/r/androiddev/comments/g5dzz8/problem_with_loading_text_from_assets_file/)
- url: https://www.reddit.com/r/androiddev/comments/g5dzz8/problem_with_loading_text_from_assets_file/
---
 

I developed my first android app, so i will describe it and the problem briefly...

It's a song book app. The app reads lines from text file that contains numbers and tile of the song then puts that data to an ArrayList  
 that is used to set up the RecyclerView  
.  When the certain item is clicked it checks for the song number that is  in array list item and then based on that opens and reads text from a  file that is named for example 23.txt  
.  I have three languages (menu items) that can be checked in navigation  menu, they have click listeners that change the current list for the RecyclerView.  
  The problem is that when the third language (menu item) is checked  whatever the first item is in the RecyclerView it throws an exception  that the file doesn't exist (while it exists in the assets folder, and  works for all other items in the list) and occasionally it says (on  Samsung Galaxy A8):

i have shared my code here [https://answiz.com/questions/90319/problem-with-loading-text-from-assets-file-problem-with-reading-text-files-from-assets-folder-unmarshaling-error-parcelable-class](https://answiz.com/questions/90319/problem-with-loading-text-from-assets-file-problem-with-reading-text-files-from-assets-folder-unmarshaling-error-parcelable-class)
## [11][Question: How to remove / customise top row of 'share with' contacts (blurred on the picture). Android 10, Pixel 2 XL. Don't want the contacts who are always appearing there.](https://www.reddit.com/r/androiddev/comments/g5dwfp/question_how_to_remove_customise_top_row_of_share/)
- url: https://i.redd.it/1fn3h8pru5u41.jpg
---

## [12][Google Play Developer Subscription Age Policy Question](https://www.reddit.com/r/androiddev/comments/g5dkv6/google_play_developer_subscription_age_policy/)
- url: https://www.reddit.com/r/androiddev/comments/g5dkv6/google_play_developer_subscription_age_policy/
---
I know, and I'm sure that many know, that Google Play has a policy where you must be at least 18 years of age to have a developer account. I'm 13, and i would want to post on Google Play in the future, so would it be ok if i had an account but my dad controls everything when it comes to payment and receiving money, as in, I have an account, but we buy the Developer options for the account through his card, and if i were to make some kind of income it would be received on his card, would that be ok with Google Play, or would we have to post everything on his account, as in I make a game, but he buys developer options for his account, instead of mine, and then if I were to make a game, we would have to post it on his account and in general do everything on his account? My dad told me that he would be happy to help me, and that we could simply just do everything on his account, but i would prefer if we could do it on mine cause i would have a little bit more accessibility and could modify my account, as in for example, if i would want to change my picture for my account, more easily.

TL;DR I'm 13 and would want to upload games on Google Play, so could my dad and i manage to do everything on my account, or would we have to do it on his, and not interfere with Google Play's policy?
