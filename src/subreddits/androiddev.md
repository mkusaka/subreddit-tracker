# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/j9q7kr/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/j9q7kr/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 12, 2020](https://www.reddit.com/r/androiddev/comments/j9oso5/weekly_questions_thread_october_12_2020/)
- url: https://www.reddit.com/r/androiddev/comments/j9oso5/weekly_questions_thread_october_12_2020/
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
## [3][[OPEN-SOURCE] Safe Dot - iOS 14 privacy in Android](https://www.reddit.com/r/androiddev/comments/jabro9/opensource_safe_dot_ios_14_privacy_in_android/)
- url: https://www.reddit.com/r/androiddev/comments/jabro9/opensource_safe_dot_ios_14_privacy_in_android/
---
I've recently built an app that mimics the iOS 14's privacy feature which alerts the user when third-party app uses your device camera or microphone.

Link to Google Play : https://play.google.com/store/apps/details?id=com.aravi.dot

So, the thing is I'm not really good at programming yet. I can build things and make them work but I can't make them efficient and 100% bug free. So I need help from you guys. 

I've made the source code open-source feel free to contribute to the app even it can be a translation to your language it really is helpful.

Here is the link to the repo : https://github.com/kamaravichow/safe-dot-android

If you have any suggestions or found bugs in the app feel free to use issues section in the repo or the comments of this post. I'll be happy to respond 😃.
## [4][Is it really that difficult to get app downloads as a new developer on the Play Store?](https://www.reddit.com/r/androiddev/comments/ja9378/is_it_really_that_difficult_to_get_app_downloads/)
- url: https://www.reddit.com/r/androiddev/comments/ja9378/is_it_really_that_difficult_to_get_app_downloads/
---
I have heard many developers say that it is very difficult to get downloads on the Google Play store. Especially when people are downloading less apps and keeping fewer apps on their devices, how has the struggle been for new developers?

I am curious to know this because I am currently developing an app and want to be aware of all of this. I would love to hear thoughts from other Android developers! Thank you for all those who respond!
## [5][Android Studio 4.1 now available with new features including Database Inspector, Native Memory Profiler, Hilt/Dagger Navigation Support and TensorFlow Lite Support](https://www.reddit.com/r/androiddev/comments/j9uv1w/android_studio_41_now_available_with_new_features/)
- url: https://android-developers.googleblog.com/2020/10/android-studio-41.html
---

## [6][Animating photos using OpenGL ES](https://www.reddit.com/r/androiddev/comments/jacvar/animating_photos_using_opengl_es/)
- url: https://www.reddit.com/r/androiddev/comments/jacvar/animating_photos_using_opengl_es/
---
I made an app which uses OpenGL ES fragment shaders and PyTorch foreground object segmentation to create live photo filters.

This is my first project working with the GPU.

Some interesting tech details:

\-The color matching when using live background is achieved by exporting color lookup tables (LUTs) from photoshop and applying them by a custom fragment shader.

\-The foreground object segmentation model is U2Net

\-Some shaders are from Shadertoy (pretty large collection of fragment shaders) 

https://reddit.com/link/jacvar/video/l2vuqj3t0vs51/player
## [7][Has anyone figured out the secure implementation of HostNameVerifier ?](https://www.reddit.com/r/androiddev/comments/jackpy/has_anyone_figured_out_the_secure_implementation/)
- url: https://www.reddit.com/r/androiddev/comments/jackpy/has_anyone_figured_out_the_secure_implementation/
---
This is my current implementation:

`private HostnameVerifier getHostnameVerifier(final URL url) {`

`return new HostnameVerifier() {`

`public boolean verify(String hostname, SSLSession session) {`

`String host_name = "examplesomething.com"`

`if(url.getHost().equals(host_name) &amp;&amp; session.getPeerHost().equals(host_name)){`

`return true;`

`}else{`

`return false;`

`}`

`}`

`};`

&amp;#x200B;

Even the official google support page [here](https://support.google.com/faqs/answer/7188426) gives the same advice:

If you are using the [HostnameVerifier](https://developer.android.com/reference/javax/net/ssl/HostnameVerifier.html) interface, change the implementation of the [verify](https://developer.android.com/reference/javax/net/ssl/HostnameVerifier.html#verify(java.lang.String,%20javax.net.ssl.SSLSession)) method to return false whenever the hostname of the server does not meet your expectations.

&amp;#x200B;

Even then i still get insecure implementation of HostnameVerifier Error. Any ideas why? Thank you
## [8][A Look Into the Future by Roman Elizarov](https://www.reddit.com/r/androiddev/comments/j9xujt/a_look_into_the_future_by_roman_elizarov/)
- url: https://youtu.be/0FF19HJDqMo
---

## [9][Released my first game](https://www.reddit.com/r/androiddev/comments/jad1hf/released_my_first_game/)
- url: https://www.reddit.com/r/androiddev/comments/jad1hf/released_my_first_game/
---
https://play.google.com/store/apps/details?id=com.vishwah.FlatLandInChaos

Released my first game on Play store please play my game and leave review , and I needed thoughts and suggestions
## [10][Storybook UI component explorer but for Android?](https://www.reddit.com/r/androiddev/comments/jacrii/storybook_ui_component_explorer_but_for_android/)
- url: https://www.reddit.com/r/androiddev/comments/jacrii/storybook_ui_component_explorer_but_for_android/
---
I'm looking to visualize my Android component library, and have used a solution called [Storybook JS](https://storybook.js.org/) in the past for Javascript component libraries. Is there a similar solution for Android? Specifically Kotlin.
## [11][How to get back into shared Play Console account?](https://www.reddit.com/r/androiddev/comments/jac4gm/how_to_get_back_into_shared_play_console_account/)
- url: https://www.reddit.com/r/androiddev/comments/jac4gm/how_to_get_back_into_shared_play_console_account/
---
A few months ago an organization invited me into their account to deploy an App, I logged in and created some of the info for the App, images etc. 

Today I logged back into my console account and cant find how to get back into the other shared account, would I receive any notification if I get removed or access expires?

On the top right profile/dropdown thing there is "manage dev accounts" but they only lead to a screen to select my own gmail accounts.
## [12][Replacing Mocks by Ryan Harter](https://www.reddit.com/r/androiddev/comments/j9wq0j/replacing_mocks_by_ryan_harter/)
- url: https://ryanharter.com/blog/2020/06/replacing-mocks/
---

