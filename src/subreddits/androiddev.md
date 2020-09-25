# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/izipbu/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/izipbu/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - September 21, 2020](https://www.reddit.com/r/androiddev/comments/iwy6cm/weekly_questions_thread_september_21_2020/)
- url: https://www.reddit.com/r/androiddev/comments/iwy6cm/weekly_questions_thread_september_21_2020/
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
## [3][Why use Room for offline caching when Retrofit/ OkHttp provide a better alternative?](https://www.reddit.com/r/androiddev/comments/izeone/why_use_room_for_offline_caching_when_retrofit/)
- url: https://www.reddit.com/r/androiddev/comments/izeone/why_use_room_for_offline_caching_when_retrofit/
---
I have been going through multiple tutorials and it seems the common/ accepted practice is to use Room to save data for offline caching. 

The first problem I faced with Room as the single source of truth was that I had to manually delete data from the app that has been removed from the backend. This made the process like this -

1. Make API call to get data
2. If API call is a success, delete old entries and add new ones from the API response and then display the data from the DB (Room with LiveData makes this easy)
3. If API call fails, fetch data from the app's database directly

Now, this seems unnecessary complex - inserting/ deleting entries from db after every API request.

I came across [this article](https://krtkush.com/2016/06/01/caching-using-okhttp-part-1.html) which explains that we can use an okHttp interceptor to store response for a defined time period and fetch it only if the API fails. 
    
    /**
     * Interceptor to cache data and maintain it for four weeks.
     *
     * If the device is offline, stale (at most four weeks old)
     * response is fetched from the cache.
     */
    private static class OfflineResponseCacheInterceptor implements Interceptor {
        @Override
        public okhttp3.Response intercept(Chain chain) throws IOException {
            Request request = chain.request();
            if (!UtilityMethods.isNetworkAvailable()) {
                request = request.newBuilder()
                        .header("Cache-Control",
                          "public, only-if-cached, max-stale=" + 2419200)
                        .build();
            }
            return chain.proceed(request);
        }
    }

I found this a better approach and less complex. 

Thoghts?
## [4][Jetpack Compose image loading library for requesting and displaying images using Fresco.](https://www.reddit.com/r/androiddev/comments/izgukk/jetpack_compose_image_loading_library_for/)
- url: https://github.com/skydoves/Frescomposable
---

## [5][Announcing a painless Kotlin/Multiplatform NoSQL embedded database](https://www.reddit.com/r/androiddev/comments/iziuf1/announcing_a_painless_kotlinmultiplatform_nosql/)
- url: https://medium.com/kodein-koders/announcing-a-painless-kotlin-multiplatform-nosql-embedded-database-30fed677549c
---

## [6][All developers will get the new Google Play Console on November 2, 2020](https://www.reddit.com/r/androiddev/comments/iz2avj/all_developers_will_get_the_new_google_play/)
- url: https://android-developers.googleblog.com/2020/09/all-developers-will-get-new-google-play.html
---

## [7][Rant: Surprised how uncomfortable/bad Room DB is for developers](https://www.reddit.com/r/androiddev/comments/izjbhq/rant_surprised_how_uncomfortablebad_room_db_is/)
- url: https://www.reddit.com/r/androiddev/comments/izjbhq/rant_surprised_how_uncomfortablebad_room_db_is/
---
After developing backend applications with Spring Boot for a while, I recently joined a mobile dev team again. They use Room as DB.    
I can see the pros: like being developed by Google, optimised for mobile and so on....   


But i was shocked how complicated such a trivial task as one-to-many relationships are with Room.  


You can't just have a list of items inside an entity. No, you need to manually persist all items in their own table and manually set the foreign keys. But if you let Room generate the ID of the "One" for you, it's pretty complicated to get this generated ID back, because room will only tell you the row number as result of the insert, but not the ID of the inserted item.   


Fortunately Room has at least some sort of support for reading one-to-many relationships. So you can define a class with the relation and room will at least create a join query.   


Maybe i miss something but based on the [official documentation](https://developer.android.com/training/data-storage/room/relationships) and some research i did that's how it needs to be done :(  


Honestly, in 2020 i would have expected way better tooling for RDBMS even on mobile platforms.
## [8][I want to create PixelArt icon packs](https://www.reddit.com/r/androiddev/comments/izjb7z/i_want_to_create_pixelart_icon_packs/)
- url: https://www.reddit.com/r/androiddev/comments/izjb7z/i_want_to_create_pixelart_icon_packs/
---
Hello.

&amp;#x200B;

First of all, let me be clear: i'm a noob and have 0 programming experience. Still, i have an Android publisher licence because i used app builders for things. I'm a PixelArtist and i want to create icon packs for Android (and possibly iOS if that's possible). Is there any understandable tutorial that might help me? If not, any developer that might want to collaborate in RevShare?
## [9][COVID related apps](https://www.reddit.com/r/androiddev/comments/izikfk/covid_related_apps/)
- url: https://www.reddit.com/r/androiddev/comments/izikfk/covid_related_apps/
---
Was reading an android games website, and it informed that google has a policy regarding COVID related apps (it was about a virus fighting game, but they couldn't add a COVID name in the game, even though the virus is round):

[https://support.google.com/googleplay/android-developer/answer/9889712?hl=en](https://support.google.com/googleplay/android-developer/answer/9889712?hl=en)

During quarantine, I practiced my app developing skills making an app that got COVID-19 related data from an API, and shows in a human comprehensible form, no judgement on data or info on what should be done with it, in local language (brazilian portuguese). So I guess that due to that policy I wouldn't be able to publish that app, is that right?
## [10][Migrating SharedPreferences to Jetpack DataStore](https://www.reddit.com/r/androiddev/comments/izijtq/migrating_sharedpreferences_to_jetpack_datastore/)
- url: https://medium.com/@jurajkunier/migrating-sharedpreferences-to-jetpack-datastore-9deb8259063
---

## [11][PSA: Android Studio's 'Local history' has all your changes if you need to revert something](https://www.reddit.com/r/androiddev/comments/iyvtdu/psa_android_studios_local_history_has_all_your/)
- url: https://www.reddit.com/r/androiddev/comments/iyvtdu/psa_android_studios_local_history_has_all_your/
---
I didn't even know what was Local history until now. 

I thought i was done with the current feature, so i made my "final commit". But i forgot something, so i reseted the commit, did my change...  Then all the sudden i had a brain fart, instead of making a commit again, i reseted my files. So seeming i lost the past few hours work. 

Turns out, AS / IJ keeps track of your changes. So i was able to reset the last 'External change' (git), reverting my changes (restoring the commit).

I have no affiliation with Coding in Flow, but he has a 5min tutorial on [Local History](https://www.youtube.com/watch?v=2A5F34TofMY)
## [12][An Android transformation library providing a variety of image transformations for Coil, Glide, Picasso, and Fresco.](https://www.reddit.com/r/androiddev/comments/izgpqz/an_android_transformation_library_providing_a/)
- url: https://github.com/wasabeef/transformers
---

