# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/f588f0/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/f588f0/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 17, 2020](https://www.reddit.com/r/androiddev/comments/f574wo/weekly_questions_thread_february_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f574wo/weekly_questions_thread_february_17_2020/
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
## [3][Coding Exercises Should Not Be Used to Rank Candidates](https://www.reddit.com/r/androiddev/comments/f5q8mv/coding_exercises_should_not_be_used_to_rank/)
- url: https://medium.com/better-programming/coding-exercises-should-not-be-used-to-rank-candidates-7b0fbde86fe6
---

## [4][GitHub - JakeWharton/shimo: Shimo is an adapter for Moshi which randomizes the order of keys when serializing and deserializing](https://www.reddit.com/r/androiddev/comments/f5efzj/github_jakewhartonshimo_shimo_is_an_adapter_for/)
- url: https://github.com/JakeWharton/shimo/
---

## [5][Is it really impossible to secure an API key in an Android app?](https://www.reddit.com/r/androiddev/comments/f5rrt4/is_it_really_impossible_to_secure_an_api_key_in/)
- url: https://www.reddit.com/r/androiddev/comments/f5rrt4/is_it_really_impossible_to_secure_an_api_key_in/
---
So I need to make sure that a request to my API really comes from my app. My solution is to use an API key. But putting an API key in the app isn't safe from reverse engineering. Here are some solutions I found online:

**Solution 1: Hide it using NDK**

I can hide the API key inside a native code or do the entire API call there. But an attacker can still extract my .so file and use it to access my API.

**Solution 2: Save it in the server**

I can save the key in the server and my app will download the key everytime it needs to access the API. Or download once, encrypt it using Android Keystore, then save the encrypted key in the shared preferences. But an attacker can look for the code to download the key and do it manually without my app. Or he/she can sniff the downloading process easily using Fiddler, Charles, or Wireshark.

Is it really impossible to secure an API from unauthorized access?
## [6][How to test Android apps in Virtual machine (from Android studio)](https://www.reddit.com/r/androiddev/comments/f5qmsg/how_to_test_android_apps_in_virtual_machine_from/)
- url: https://www.reddit.com/r/androiddev/comments/f5qmsg/how_to_test_android_apps_in_virtual_machine_from/
---
I am running Android studio inside VMware virtual machine. When I tried to launch Android emulator, Android studio says that it can't run inside vm. (Although x86 version of Android emulator can be run, it is very very slow) Genymotion is also failing to run on vm.
Is there any way to run android apps inside vm?
## [7][Marketing a game as a spare time developer with a small-ish budget](https://www.reddit.com/r/androiddev/comments/f5lamj/marketing_a_game_as_a_spare_time_developer_with_a/)
- url: https://www.reddit.com/r/androiddev/comments/f5lamj/marketing_a_game_as_a_spare_time_developer_with_a/
---
Hi all,

I've spent the last few years creating Android games with some varying success in my spare time, a game a released last year is called 'Tennis Superstars' I was really happy with how the game turned out but have had a very limited success in the app store. The feedback I've had has been great (avg 4.6 reviews [https://play.google.com/store/apps/details?id=com.moz.tennis](https://play.google.com/store/apps/details?id=com.moz.tennis&amp;hl=en_US)) but I think there is a visibility problem with the game, I'm way down in the Play Store when I search for 'Tennis games'.

I am more than willing to invest some money into promoting this game, but I don't know what would be effective with a budget in the range of $100's of dollars, is anything achievable? In the past I have spent $100+ overall in Facebook advertising app installs with little success, I don't feel like this is the right way. I've also done a whole bunch of app giveaways and free advertising on sites such as reddit (/r/Tennis for example) which seemed to work better than the Facebook adverts.

Is it possible to submit my app to receive a paid review from either a website or a social media influencer? Is there a reliable company that can manage this process on a small scale? I thought paid reviews could be submitted to [www.androidpolice.com](https://www.androidpolice.com) but I can't see where? I've also seen a lot of scammy looking websites offering to market games that obviously I don't want to get involved in. Maybe it would be a good idea to get a paid review from an Android gaming YouTuber?

I imagine a lot of spare time developers are in this situation, I'm not expecting something from nothing, I want to invest money in my marketing - I just want to do it in a sensible way.
## [8][Desktop Entry for Stable And Preview version of Android Studio](https://www.reddit.com/r/androiddev/comments/f5r3ek/desktop_entry_for_stable_and_preview_version_of/)
- url: https://www.reddit.com/r/androiddev/comments/f5r3ek/desktop_entry_for_stable_and_preview_version_of/
---
Android Studio allow to create desktop entry from "Tools-&gt;Create Desktop Entry", but if you try to do it with two versions of IDE (Stable, Preview) at the same time, desktop entry will be overridden

And i'd like to share solution and make two versions available from desktop:

For Stable version you can use  "Tools-&gt;Create Desktop Entry"

For Preview version you can use this files: [gist](https://gist.github.com/lndmflngs/7edfee54adf378f5db1886e08eadbe66)

Edit files with your properties, put them at one folder and run `sudo ./create_desktop_entry.sh`  

Result:  


https://preview.redd.it/pu3ctv69foh41.png?width=498&amp;format=png&amp;auto=webp&amp;s=c0e1eef62da3e1344081a8dccdbbf01cf0c3ee2f
## [9][Hypercube. How to provide developers with phones for testing without losing any](https://www.reddit.com/r/androiddev/comments/f5r1nz/hypercube_how_to_provide_developers_with_phones/)
- url: https://habr.com/en/company/yandex/blog/487722/
---

## [10][[Kotlin] Store ArrayList persistent](https://www.reddit.com/r/androiddev/comments/f5qzyu/kotlin_store_arraylist_persistent/)
- url: https://www.reddit.com/r/androiddev/comments/f5qzyu/kotlin_store_arraylist_persistent/
---


My app is displaying a list of various categories (herbs, side dishes, ..) in a RecyclerView. Depending on the category you clicked on, a new Activity with a new RecylcerView opens containing all the ingredients.

Right now I have an ArrayList which gets filled with the ingredients via ".add" depending on the choosen category.

The problem im facing right now is, that I want to implement an option for the user to add own Ingredients. I tried storing the ArrayList containing the ingredients in SharedPreferences by using Gson, but I couldn't manage to add elements, since it always overwrote the current list.

What would be the best way to store the ingredients? A room, sqlite, ..? Without further explanation, the ingredient list will only contain about 70 items max.

Thanks in advance.
## [11][Google Play Issue with Apps](https://www.reddit.com/r/androiddev/comments/f5oqqm/google_play_issue_with_apps/)
- url: https://www.reddit.com/r/androiddev/comments/f5oqqm/google_play_issue_with_apps/
---
 I'm pretty new to developing android apps. I have published my first 2 apps on Google Play and the issue is that I can't have them both installed at the same time. I had the first one on my phone, but when I wanted to install my second one I couldn't do it unless I uninstalled the first app.

Anyone had this issue before?

The only thing the apps have in common is probably the .jks key used to generate the .aab bundle, but that shouldn't be a problem.

And no, the apps don't have the same package name.
## [12][Draw Path along stars with fingers in Canvas .](https://www.reddit.com/r/androiddev/comments/f5qxml/draw_path_along_stars_with_fingers_in_canvas/)
- url: https://www.reddit.com/r/androiddev/comments/f5qxml/draw_path_along_stars_with_fingers_in_canvas/
---
Hi there .I am able to draw line in canvas using path when user finger moves changing paint line colors and all ,  But my need is i need to draw stars along the line .Any hint please ...
