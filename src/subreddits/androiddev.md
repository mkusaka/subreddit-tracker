# androiddev
## [1][Weekly Questions Thread - June 22, 2020](https://www.reddit.com/r/androiddev/comments/hdq8ko/weekly_questions_thread_june_22_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hdq8ko/weekly_questions_thread_june_22_2020/
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
## [3][How do I test my app with premium features? Will it impact on my future sales? Should I do my semi-private beta with or without adverts?](https://www.reddit.com/r/androiddev/comments/hfktuw/how_do_i_test_my_app_with_premium_features_will/)
- url: https://www.reddit.com/r/androiddev/comments/hfktuw/how_do_i_test_my_app_with_premium_features_will/
---
So I've spent the last 6 months (excluding the long hiatuses in-between) working on my study timer app: Pomato. 

It's basically a study app which follows the [pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique) . I've tested it heavily myself, and I'm more or less happy with how it functions. Now I want to test it with other peoples phones. I've got a beta testing track on the Play Console (though I can't seem to make it work without adding peoples emails. I just want the invites to work with links). 

Only thing is, a lot of features on my app I want to reserve for the premium version which will be out alongside the ads supported free version. How do I go about testing in this situation? Do I test with all the premium stuff unlocked and then lock it all down before release and then add the adverts?

Or do I do all this now? 

How have you guys managed situations like this and how did it work out?
## [4][How to reliably calculate the screen size / screen zoom level when you can set different screen resolutions?](https://www.reddit.com/r/androiddev/comments/hfibpg/how_to_reliably_calculate_the_screen_size_screen/)
- url: https://www.reddit.com/r/androiddev/comments/hfibpg/how_to_reliably_calculate_the_screen_size_screen/
---
Is there a way to reliably calculate the screen size / screen zoom level, which can be changed in Settings App -&gt; Display?

I found [here](https://www.reddit.com/r/androiddev/comments/80y795/how_to_programmatically_check_if_user_has/) that you can use [DisplayMetrics.DENSITY_DEVICE_STABLE](https://developer.android.com/reference/android/util/DisplayMetrics#DENSITY_DEVICE_STABLE) and [densityDpi](https://developer.android.com/reference/android/util/DisplayMetrics#densityDpi) to calculate the "zoom level".
This generally works for most devices, but not on certain Samsung devices where you can change the resolution.

On a Samsung Galaxy S10 and the resolution set to FHD+:

* DENSITY_DEVICE_STABLE reports 420
* standard display zoom: densityDpi reports 420
* next zoom level: densityDpi reports 450

420 / 420 = 1

450 / 420 = 1.0714285

Works as expeced, yay!

Now we change the resolution:

Samsung Galaxy S10 and the resolution set to WQHD+:

* DENSITY_DEVICE_STABLE reports 420
* standard display zoom: densityDpi reports 560

560 / 420 = 1.3333334

We can't reliably calculate the zoom level since densityDpi reports different values for the same zoom level on different resolutions, but DENSITY_DEVICE_STABLE stays the same, regardless of what resolution is set.

Is there any other way you can do this?
## [5][Merge Adapter : Merge your lists seamlessly in Android](https://www.reddit.com/r/androiddev/comments/hflav6/merge_adapter_merge_your_lists_seamlessly_in/)
- url: https://blog.yudiz.com/merge-adapter-merge-your-lists-seamlessly-in-android/
---

## [6][The Bifurcation of Android](https://www.reddit.com/r/androiddev/comments/hf47vz/the_bifurcation_of_android/)
- url: https://medium.com/nala-money/the-bifurcation-of-android-6fa1cced074d
---

## [7][Making The Most of a Rewrite](https://www.reddit.com/r/androiddev/comments/hfbjuf/making_the_most_of_a_rewrite/)
- url: https://eng.snap.com/making-the-most-of-a-rewrite
---

## [8][Looking for any library to create good looking blur layout/view like ios blur. I have tried BlurKit but can't make it looks good](https://www.reddit.com/r/androiddev/comments/hflr88/looking_for_any_library_to_create_good_looking/)
- url: https://www.reddit.com/r/androiddev/comments/hflr88/looking_for_any_library_to_create_good_looking/
---

## [9][Advice on feature separation of concerns.](https://www.reddit.com/r/androiddev/comments/hfllh5/advice_on_feature_separation_of_concerns/)
- url: https://www.reddit.com/r/androiddev/comments/hfllh5/advice_on_feature_separation_of_concerns/
---
Hello,

I have a project coming up where I need to build admin related features on top of standard user based features.

I thought about using dynamic delivery here.  
[https://developer.android.com/guide/app-bundle/dynamic-delivery#customize\_delivery](https://developer.android.com/guide/app-bundle/dynamic-delivery#customize_delivery)   


It would make for a nice clean separation of concerns. If you login as admin then we add the admin module using the above service.

&amp;#x200B;

But I am interested to know what happens if we want to just make a small change to a screen thats already available. Im not sure how I can utilize dynamic delivery for this?

&amp;#x200B;

For example a user sees the cost of a product, but if they are  admin level a button appears next to it for the admin to change the cost. 

&amp;#x200B;

If I ignore the idea of using dynamic delivery I would just go with some boolean type thing   
"if admin enable/disable visibility of button component" 

&amp;#x200B;

Any thoughts on this appreciated.
## [10][Veracode: Open-source libraries cause security flaws in 70% of apps](https://www.reddit.com/r/androiddev/comments/hfgy0q/veracode_opensource_libraries_cause_security/)
- url: https://developer-tech.com/news/2020/may/27/veracode-open-source-libraries-security-flaws-apps/
---

## [11][What is the standard way for getting an access token directly from a Google API using Google Sign-In for Android client app?](https://www.reddit.com/r/androiddev/comments/hfkyan/what_is_the_standard_way_for_getting_an_access/)
- url: https://www.reddit.com/r/androiddev/comments/hfkyan/what_is_the_standard_way_for_getting_an_access/
---
 

I am working on an android app that connects to Google Books API directly. Part of its working is to fetch user's private book data as well. (such as bookshelves)

I used Google Sign-in for Android for authentication. However, i needed an access token for the request's authorization as well.

**THE PROBLEM:-**

Google Sign-in does a great job handling the authentication part but i was stuck in implementing the authorization part for getting an access token since it provided no methods for it.

It only provides us a server auth code through its `requestServerAuthCode(String web_client_id)`  
 method but it expects us to exchange the retrieved code with OUR OWN backend server to get an access token (which is not fit for my use case since i don't have any backend). 

**POSSIBLE SOLUTIONS:-**

**#1:** One way for achieving this is by using the GoogleAuthUtils.getToken(Context, Account, String scope)  
 method. Although [this method's doc](https://developers.google.com/android/reference/com/google/android/gms/auth/GoogleAuthUtil#public-static-string-gettoken-context-context,-account-account,-string-scope,-bundle-extras)  states to use this method for getting an access token directly, but it is mentioned in [this other doc](https://android-developers.googleblog.com/2016/05/improving-security-and-user-experience.html)  

about some security issues with this approach. I can't figure out if this method is fit for my use case or not?

**#2:** Another way for getting an access token is by exchanging the auth code manually:-

    OkHttpClient client = new OkHttpClient();
            RequestBody requestBody = new FormEncodingBuilder()
                    .add("grant_type", "authorization_code")
                    .add("client_id", "{client_id}")
                    .add("client_secret", "{clientSecret}")
                    .add("redirect_uri","")
                    .add("code", "4/4-GMMhmHCXhWEzkobqIHGG_EnNYYsAkukHspeYUk9E8")    //this code is received from requestServerAuthCode(...) method
                    .build();
            final Request request = new Request.Builder()
                    .url("https://www.googleapis.com/oauth2/v4/token")
                    .post(requestBody)
                    .build();
            client.newCall(request).enqueue(new Callback() {
                u/Override
                public void onFailure(final Request request, final IOException e) {
                    Log.e(LOG_TAG, e.toString());                
                }
        
                u/Override
                public void onResponse(Response response) throws IOException {
                    try {
                        JSONObject jsonObject = new JSONObject(response.body().string());
                        final String message = jsonObject.toString(5);
                        Log.i(LOG_TAG, message);                    
                    } catch (JSONException e) {
                        e.printStackTrace();
                    }
                }
            });

 With a sucessful response, you will have the following info in logcat: 

    I/onResponse: {
                      "expires_in": 3600,
                      "token_type": "Bearer",
                      "refresh_token": "1\/xz1eb0XU3....nxoALEVQ",
                      "id_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQxMWY1Ym......yWVsUA",
                      "access_token": "ya29.bQKKYah-........_tkt980_qAGIo9yeWEG4"
                 }

 

I am currently using this approach. However the problem with this approach is the client credentials (client\_id and client\_secret) are being exposed.

**QUESTION:-**

Which of the above solutions should i apply for my problem?
## [12][Your Recommendations Are Important To Me!](https://www.reddit.com/r/androiddev/comments/hfka99/your_recommendations_are_important_to_me/)
- url: https://www.reddit.com/r/androiddev/comments/hfka99/your_recommendations_are_important_to_me/
---
Hi people,

I'm a college student and trying to be a good app developer. I developed an app but I couldn't be able to get more than one thousand downloads. My app provides the people who hides their instagram stories from you and lets you to save and watch other people's stories anonymously. If anybody can help me to increase the downloads, I would appreciate it. Thanks for your help.

Web: [https://socialhunters.app/landing](https://socialhunters.app/landing)

Google Play: [https://play.google.com/store/apps/details?id=socialhunters.app](https://play.google.com/store/apps/details?id=socialhunters.app)

&amp;#x200B;
