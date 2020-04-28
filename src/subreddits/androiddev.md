# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/g900i2/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/g900i2/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - April 27, 2020](https://www.reddit.com/r/androiddev/comments/g8yltx/weekly_questions_thread_april_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g8yltx/weekly_questions_thread_april_27_2020/
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
## [3][Was my game hacked? What can I do?](https://www.reddit.com/r/androiddev/comments/g99pla/was_my_game_hacked_what_can_i_do/)
- url: https://www.reddit.com/r/androiddev/comments/g99pla/was_my_game_hacked_what_can_i_do/
---
I have a game developed in Unity that is currently in the Google Play Store.

Today, when I looked on Firebase Console in Latest Release, unexpectedly, I see a newly released version which I don't recognize! According to Firebase, the Latest Release is version 17.1.1701, but the latest version which I published myself is version 2.3! Looking further, it appears that these rogue installs have started to appear in the past week, and evidently there have been 4 releases like this in total.

Currently, this latest version which I do not recognize is almost 10% of my active users. I checked my Google Play Publish Console to make sure I wasn't hacked there, and everything looks good. It is showing that version 2.3 is indeed my latest official release.

I did some checking in Google Analytics, and it appears that these installs are almost exclusively in China.

So my questions are, where are these installs coming from (I mean specifically is there a website dedicated to this type of thing)? Was my game somehow hacked / hijacked? How is this possible and is there anything I can do about it? Does anyone else have any experience with this?
## [4][High refresh rate rendering on Android](https://www.reddit.com/r/androiddev/comments/g96kmb/high_refresh_rate_rendering_on_android/)
- url: https://android-developers.googleblog.com/2020/04/high-refresh-rate-rendering-on-android.html
---

## [5][Is it true that the Android Keystore can at best protect the keys but not the data?](https://www.reddit.com/r/androiddev/comments/g9ilc5/is_it_true_that_the_android_keystore_can_at_best/)
- url: https://www.reddit.com/r/androiddev/comments/g9ilc5/is_it_true_that_the_android_keystore_can_at_best/
---
I've been reading up on the benefits of protecting sensitive data using the Android Keystore. So far, what I've understood is that the Keystore, when hardware backed, will safe guard my keys, preventing extraction. However, it cannot guard against an attacker using them to decrypt the data. A root user would be able to impersonate the app that owns the keys, and thereby use them to decrypt data without the need for extraction. The user won't be able to benefit outside of the device however, as extraction is not possible.

So then, what is the point of encrypting data that is anyway not accessible without root, given that with root it is possible to decrypt without the need to extract the key. I'm looking at this from the angle of safeguarding things like API keys, etc. that are not user specific, and when leaked, could be used outside of the user's scope.

For reference, [here](https://security.stackexchange.com/questions/201068/android-what-are-the-practical-security-benefits-of-using-a-hardware-backed-ke)'s a post I read on stack exchange. There's also [this research](https://cypherpunk.nl/papers/spsm14.pdf) that seems to back up the claim.
## [6][PSA: Google Play appears to be down-ranking my app while it's under review](https://www.reddit.com/r/androiddev/comments/g9mbpe/psa_google_play_appears_to_be_downranking_my_app/)
- url: https://www.reddit.com/r/androiddev/comments/g9mbpe/psa_google_play_appears_to_be_downranking_my_app/
---
I have good reason to believe that whenever my app is listed as "We are currently reviewing your app." (in the popup that appears when you click the info button next to your app name at the top of the Console) my app is severely penalized in the search results on Play Store. This may be some feature Google is experimenting with or already implemented.

This "under review" status happens even when I increase the staged rollout percentage, so limiting this activity may be prudent to keep installs and revenue up. Also, I'd be concerned if my app got flagged for manual review on updates, right now my updates go live in the store in a couple hours, and the "under review" flag disappears after a couple more. 

This is following years of my app being stably ranked #1 for a few search terms, a couple of them such as the name of the app it will even give half the screen space with the app screenshots instead of a list of apps. Within some minutes of submitting an update, I lose #1 in every search term dropping to #4 or 6 for some of them. As soon as the flag is cleared, within minutes it recovers. I've checked the listing from different devices from different accounts and from devices and accounts not associated with me in any way and located in different geographic region even.
## [7][Mobile Game Playtime Increases By 62% Due To Coronavirus](https://www.reddit.com/r/androiddev/comments/g9m6bi/mobile_game_playtime_increases_by_62_due_to/)
- url: https://mobilemarketingreads.com/mobile-game-playtime-increases-by-62-due-to-coronavirus/
---

## [8][Sealed classes + RecyclerView with headers = ❤️](https://www.reddit.com/r/androiddev/comments/g9a52n/sealed_classes_recyclerview_with_headers/)
- url: https://medium.com/@patxi/sealed-classes-recyclerview-with-headers-%EF%B8%8F-14b87d41deb6?source=friends_link&amp;sk=2cb3fac8e2bce1ecbb8488d36fe1cf24
---

## [9][Google Bots back from lockdown &amp; now just showing notifications from last week](https://www.reddit.com/r/androiddev/comments/g9n5ma/google_bots_back_from_lockdown_now_just_showing/)
- url: https://imgur.com/a/z7SXBB1
---

## [10][Droidcon Online 2020: Become A Composer](https://www.reddit.com/r/androiddev/comments/g9mxp6/droidcon_online_2020_become_a_composer/)
- url: https://www.droidcon.com/media-detail?video=412304809
---

## [11][Make art with Android Development](https://www.reddit.com/r/androiddev/comments/g9m38d/make_art_with_android_development/)
- url: https://medium.com/@elye.project/make-art-with-android-development-598fe21afaae?source=friends_link&amp;sk=3bcfee75c8f2451a667a94acafbc60d6
---

## [12][How Much Will It Cost to Design a Mobile App in 2020? Let us look at these questions by estimating the mobile application design cost of Uber and Instagram](https://www.reddit.com/r/androiddev/comments/g9lqaz/how_much_will_it_cost_to_design_a_mobile_app_in/)
- url: https://kodytechnolab.com//cost-to-design-mobile-application
---

