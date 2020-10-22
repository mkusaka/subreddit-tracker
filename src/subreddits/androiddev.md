# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 19, 2020](https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/
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
## [3][Got Mail From Google “Your app(s) are vulnerable to Intent Redirection” but cannot find vulnerability](https://www.reddit.com/r/androiddev/comments/jfx18a/got_mail_from_google_your_apps_are_vulnerable_to/)
- url: https://www.reddit.com/r/androiddev/comments/jfx18a/got_mail_from_google_your_apps_are_vulnerable_to/
---
When I update my app then I got a mail from google showing

Hello Google Play Developer,

We reviewed your app and found that your app uses software that contains security vulnerabilities for users. Apps with these vulnerabilities can expose user information or damage a user’s device, and may be considered to be in violation of our Malicious Behavior policy.

Vulnerability - **Your app(s) are vulnerable to Intent Redirection.**

But I am not doing any intent extraction or redirection. After searching a little bit I apply some solution like `android:exported="false"`  
 but none of them works.

Now, I think some third party plugin is responsible for the vulnerability. How to find that vulnerability in my code?

Below is my build.gradle file

    apply plugin: 'com.android.application' apply plugin: 'com.google.gms.google-services' apply plugin: 'com.google.firebase.crashlytics'  android {     compileSdkVersion 29     defaultConfig {         minSdkVersion 21         targetSdkVersion 29         renderscriptTargetApi 18         renderscriptSupportModeEnabled true         multiDexEnabled true         testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"         resConfigs "en"      }     buildTypes {         debug {             firebaseCrashlytics {                 mappingFileUploadEnabled false             }         }         release {             minifyEnabled true             shrinkResources true             proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'         }     }     compileOptions {         sourceCompatibility = '1.8'         targetCompatibility = '1.8'     }     externalNativeBuild {         ndkBuild {             path 'src/main/jni/Android.mk'         }     }     dependencies {         androidTestImplementation('androidx.test.espresso:espresso-core:3.1.1', {             exclude group: 'com.android.support', module: 'support-annotations'         })         testImplementation 'junit:junit:4.12'         implementation 'androidx.multidex:multidex:2.0.1'         implementation fileTree(dir: 'libs', include: ['*.jar'])          implementation 'androidx.appcompat:appcompat:1.2.0'         implementation 'androidx.constraintlayout:constraintlayout:2.0.2'         implementation 'androidx.legacy:legacy-support-v4:1.0.0'         implementation 'androidx.recyclerview:recyclerview:1.1.0'         implementation 'androidx.cardview:cardview:1.0.0'         implementation "androidx.viewpager2:viewpager2:1.0.0"         implementation 'com.github.bumptech.glide:glide:4.11.0'         annotationProcessor 'com.github.bumptech.glide:compiler:4.11.0'         implementation 'com.google.android.material:material:1.3.0-alpha03'         //fix         implementation 'com.google.guava:listenablefuture:9999.0-empty-to-avoid-conflict-with-guava'         implementation 'com.google.android.gms:play-services-basement:17.5.0'          //Firebase         implementation platform('com.google.firebase:firebase-bom:25.12.0')         implementation 'com.google.firebase:firebase-auth'         implementation 'com.google.firebase:firebase-database'         implementation 'com.google.firebase:firebase-storage'         implementation 'com.google.firebase:firebase-firestore'         implementation 'com.google.firebase:firebase-messaging'         implementation 'com.google.firebase:firebase-config'         implementation 'com.google.firebase:firebase-analytics'         implementation 'com.google.firebase:firebase-crashlytics'         implementation 'com.google.firebase:firebase-inappmessaging-display'         implementation 'com.firebaseui:firebase-ui-auth:6.3.0'          // Google Sign In SDK (only required for Google Sign In)         implementation 'com.google.android.gms:play-services-auth:18.1.0'         implementation 'com.google.android.gms:play-services-identity:17.0.0'          // Facebook Android SDK (only required for Facebook Login)         implementation 'androidx.browser:browser:1.2.0'         implementation 'com.facebook.android:facebook-android-sdk:7.1.0'          //retrofit         implementation 'com.squareup.retrofit2:retrofit:2.2.0'         implementation 'com.squareup.retrofit2:converter-gson:2.2.0'         implementation 'com.squareup.okhttp3:logging-interceptor:3.9.0'          //json utilities         implementation 'com.fasterxml.jackson.core:jackson-core:2.10.1'         implementation 'com.fasterxml.jackson.core:jackson-annotations:2.10.1'         implementation 'com.fasterxml.jackson.core:jackson-databind:2.10.1'          //Paytm All-in-one SDK Payment Gateway         implementation'com.paytm.appinvokesdk:appinvokesdk:1.5'         //Circular ImageView         implementation 'de.hdodenhof:circleimageview:2.2.0'         //Number Picker         implementation 'com.shawnlin:number-picker:2.4.7'         //Version Compare         implementation 'com.g00fy2:versioncompare:1.3.2'         //Seekbar         implementation 'com.github.warkiz.widget:indicatorseekbar:2.1.2'         //country picker         implementation 'com.hbb20:ccp:2.3.1'         //TrueTime         implementation 'com.github.instacart.truetime-android:library:3.4'         //facebook ads         implementation 'com.facebook.android:audience-network-sdk:6.1.0'         //shimmer         implementation 'com.facebook.shimmer:shimmer:0.5.0'         //Loading Animation         implementation 'com.wang.avi:library:2.1.3'         //likeButton         implementation 'com.github.varunest:sparkbutton:1.0.6'     }      configurations.all {         resolutionStrategy.force 'com.android.support:support-annotations:28.0.0'     }  }
## [4][Help: Mass spam and doxxing in play store reviews, no action taken by Google](https://www.reddit.com/r/androiddev/comments/jfycgb/help_mass_spam_and_doxxing_in_play_store_reviews/)
- url: https://www.reddit.com/r/androiddev/comments/jfycgb/help_mass_spam_and_doxxing_in_play_store_reviews/
---
A month back we started a Discord group for my app to build a community and get user feedback. It's a gamification app where users get vouchers for some tasks. But soon, Discord became a channel for users to blackmail us for vouchers.  


Few days back one user spammed 1 star reviews to "bargain" for a voucher he didn't earn. After banning him, he doxxed our identity in a review. The review is not removed by Google for supposedly not violating the "[Google Play Comment posting policy](https://play.google.com/about/comment-posting-policy/)". What can be done?  


[Doxxed- this does not violate Google Play comment posting policy](https://preview.redd.it/qx2e0xoiymu51.png?width=922&amp;format=png&amp;auto=webp&amp;s=0e8bdd3eaf80cc90672a939fabc74b4b02f8f27d)

[1 star spam- no action taken](https://preview.redd.it/gytzia0rxmu51.png?width=367&amp;format=png&amp;auto=webp&amp;s=68755a188fd315028484b4d7deaf4c1afdc81a57)
## [5][Will there be much performance impact from me developing from a USB stick as opposed to having the files on my computer?](https://www.reddit.com/r/androiddev/comments/jfxwl3/will_there_be_much_performance_impact_from_me/)
- url: https://www.reddit.com/r/androiddev/comments/jfxwl3/will_there_be_much_performance_impact_from_me/
---
Apologies if I've explained that terribly. 

Basically, if I open my Android project from my USB stick as opposed to opening it from the SSD inside my computer, will it affect load times all that much?
## [6][Streaming API returning multiple type objects](https://www.reddit.com/r/androiddev/comments/jfqqxe/streaming_api_returning_multiple_type_objects/)
- url: https://www.reddit.com/r/androiddev/comments/jfqqxe/streaming_api_returning_multiple_type_objects/
---
I'm working with the lichess.org API. It has some methods that stream events, each event having a different fields in their JSON representation ( the only common field "type" representing the type of event ). How should I deal with this ? I want to convert each event to a java object. 

 I'm using java and retrofit for network calls, any library or framework is welcome.
## [7][Help me with some research: we want to create a better monetisation option for app developers.](https://www.reddit.com/r/androiddev/comments/jfzaiq/help_me_with_some_research_we_want_to_create_a/)
- url: https://www.reddit.com/r/androiddev/comments/jfzaiq/help_me_with_some_research_we_want_to_create_a/
---
Hey! If you have ever used ads to monetize your app, we would love to hear from you. We want to create a better option for app developers to monetize their hard work, but first we need to learn some more. If you can spare &lt;2 minutes to answer 5 short questions, it would be appreciated! 🙏🏼
https://docs.google.com/forms/d/e/1FAIpQLSfWblwaigbavrAN5_k5vCmWtJE1AGt09u-Tqve55KNZqrHIXw/viewform

Thanks so much for your help.
## [8][Feedback on blog post - Network request and battery drain](https://www.reddit.com/r/androiddev/comments/jfz26q/feedback_on_blog_post_network_request_and_battery/)
- url: https://www.reddit.com/r/androiddev/comments/jfz26q/feedback_on_blog_post_network_request_and_battery/
---
Hi guys,

Ive just started blogging on medium. This is my second blog post that post about why network calls affect device battery.

[https://anirudhmenon.medium.com/android-network-requests-and-battery-drain-2af7283f48ec](https://anirudhmenon.medium.com/android-network-requests-and-battery-drain-2af7283f48ec)

I'd appreciate feedback. Thank you so much. :)
## [9][Developer Option: Stay awake while USB debugging is connected](https://www.reddit.com/r/androiddev/comments/jfz1g3/developer_option_stay_awake_while_usb_debugging/)
- url: https://www.reddit.com/r/androiddev/comments/jfz1g3/developer_option_stay_awake_while_usb_debugging/
---
I use my own phone for development and I find it convenient to enable the stay awake while I'm working. The issue is that I constantly forget to disable when I'm done working. Later, when I go to bed and charge my phone, I'll get some notification and my screen stays on all night (probably not the best for my hardware). Wouldn't it be better if the stay awake option only applied while the USB debugger was connected? It could even be an alternative option. I think I could sleep more soundly if I didn't have to worry about my screen being on all night.

Are there any other options that you use or could see being part of the developer options?
## [10][Check out my new pal :)](https://www.reddit.com/r/androiddev/comments/jf8sqy/check_out_my_new_pal/)
- url: https://i.redd.it/tscwa1jfoeu51.jpg
---

## [11][How can I underline bottom navigation bar items?](https://www.reddit.com/r/androiddev/comments/jfxhuq/how_can_i_underline_bottom_navigation_bar_items/)
- url: https://www.reddit.com/r/androiddev/comments/jfxhuq/how_can_i_underline_bottom_navigation_bar_items/
---
Exactly like in the image, but I have also text below icons.  
I want to change the color &amp; width &amp; height of the underline so spannable didn't help.  iOS equivalent is  

    tabBar.selectionIndicatorImage

&amp;#x200B;

https://preview.redd.it/2p88i8uynmu51.png?width=696&amp;format=png&amp;auto=webp&amp;s=1c0bc4cf8416ea1bc5d50a051f1235bf8610cbfe
## [12][The US government has filed antitrust charges against Google](https://www.reddit.com/r/androiddev/comments/jf6b8n/the_us_government_has_filed_antitrust_charges/)
- url: https://www.theverge.com/2020/10/20/21454192/google-monopoly-antitrust-case-lawsuit-filed-us-doj-department-of-justice
---

