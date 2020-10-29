# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/jidusx/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jidusx/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 26, 2020](https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jicfz7/weekly_questions_thread_october_26_2020/
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
## [3][Netflix Android and iOS Studio Apps — now powered by Kotlin Multiplatform](https://www.reddit.com/r/androiddev/comments/jk0rur/netflix_android_and_ios_studio_apps_now_powered/)
- url: https://medium.com/@NetflixTechBlog/netflix-android-and-ios-studio-apps-kotlin-multiplatform-d6d4d8d25d23
---

## [4][What modern day app looks like](https://www.reddit.com/r/androiddev/comments/jjk9ks/what_modern_day_app_looks_like/)
- url: https://i.redd.it/wmptdhg5vsv51.jpg
---

## [5][How do I fill the whole popup menu?, don't want the white in the square.](https://www.reddit.com/r/androiddev/comments/jk6yxg/how_do_i_fill_the_whole_popup_menu_dont_want_the/)
- url: https://i.redd.it/alq0uqkq60w51.png
---

## [6][How can I unit test my view utils functions?](https://www.reddit.com/r/androiddev/comments/jk6xaq/how_can_i_unit_test_my_view_utils_functions/)
- url: https://www.reddit.com/r/androiddev/comments/jk6xaq/how_can_i_unit_test_my_view_utils_functions/
---
I  gathered quite a few utils functions in a library, using kotlin   extension functions to simplify common operations when Android framework   API was too cumbersome. Like in this very simple example:

    val View.contentWidth: Int get() = width - paddingLeft - paddingRight  

I would like to add unit  tests to this library but have no idea how  to test this. I'm new to  unit tests and already read a lot of  documentation on how to getting started but feel a bit lost about it at  the moment and trying to search directly for a solution yielded no result.
## [7][How much and why is accessibility in Android Apps important for you? What problems do you find in implementing the accessibility guidelines?](https://www.reddit.com/r/androiddev/comments/jk9ist/how_much_and_why_is_accessibility_in_android_apps/)
- url: https://www.reddit.com/r/androiddev/comments/jk9ist/how_much_and_why_is_accessibility_in_android_apps/
---

## [8][Released integration between the Navigation component and Jetpack Compose as androidx.navigation:navigation-compose:1.0.0-alpha01](https://www.reddit.com/r/androiddev/comments/jjt4b1/released_integration_between_the_navigation/)
- url: https://developer.android.com/jetpack/compose/navigation
---

## [9][Twitter like feed](https://www.reddit.com/r/androiddev/comments/jka41h/twitter_like_feed/)
- url: https://www.reddit.com/r/androiddev/comments/jka41h/twitter_like_feed/
---
How to accomplish feed item like twitter where user gets to see not only the like counts but also if he has liked that post or not. The heart is red if he has liked the post.
## [10][What kind of data type column you use in SQLite, to store LocalDate?](https://www.reddit.com/r/androiddev/comments/jka2ek/what_kind_of_data_type_column_you_use_in_sqlite/)
- url: https://www.reddit.com/r/androiddev/comments/jka2ek/what_kind_of_data_type_column_you_use_in_sqlite/
---
Currently, we need to store holidays information in SQLite.

We prefer to use LocalDate in our application code.

As, Christmas always happen on 2020-12-25 doesn't matter the user is from United Kingdom, or from Thailand.

In SQLite side, we tend not to store it in INTEGER timestamp, as that requires correct timezone information.

So far, we choose to store it in TEXT.

    public class LocalDateConverter {
        // 2020-12-25
        // 2019-11-08
    
        private static final ThreadLocal&lt;DateTimeFormatter&gt; threadLocal = ThreadLocal.withInitial(() -&gt; {
            DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd", Locale.US);
            return formatter;
        });
    
        @TypeConverter
        public static String toString(LocalDate localDate) {
            return threadLocal.get().format(localDate);
        }
    
        @TypeConverter
        public static LocalDate toLocalDate(String string) {
            return LocalDate.parse(string, threadLocal.get());
        }
    }

I was wondering, is that a good approach?

Is there a better data type in SQLite to store LocalDate using Android Room?
## [11][Which is the security process behind Android Apk signature key and the playstore ?](https://www.reddit.com/r/androiddev/comments/jk9tmv/which_is_the_security_process_behind_android_apk/)
- url: https://www.reddit.com/r/androiddev/comments/jk9tmv/which_is_the_security_process_behind_android_apk/
---
Hi :)

When we decide to create our app we create a keystore, we decide the keystore password, the keyAlias and the keyPassword and after that we call assembleRelease.

We go to the play console and we upload our fantastic apk.

But in which way google checks the apk is signed with the correct signature key ? Which is the checks ?

Which is the role of the keystore ?
## [12][NotyKT is a complete Kotlin-stack application built with Ktor (REST API Backend) + Android](https://www.reddit.com/r/androiddev/comments/jk9i15/notykt_is_a_complete_kotlinstack_application/)
- url: https://github.com/PatilShreyas/NotyKT
---

