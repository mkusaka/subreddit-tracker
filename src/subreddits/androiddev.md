# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/fnjcmy/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fnjcmy/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - March 23, 2020](https://www.reddit.com/r/androiddev/comments/fni20d/weekly_questions_thread_march_23_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fni20d/weekly_questions_thread_march_23_2020/
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
## [3][App with 1 Million downloads suspended because of the word "Corona"](https://www.reddit.com/r/androiddev/comments/fn0wzl/app_with_1_million_downloads_suspended_because_of/)
- url: https://www.reddit.com/r/androiddev/comments/fn0wzl/app_with_1_million_downloads_suspended_because_of/
---
I'm the creator of the Meme and Vine Soundboard which had nearly 1 Million downloads. Sadly it got suspended 2 days ago because I wasn't aware of the fact that anything which mentions "Corona" is getting suspended instantly.

The app contained more than 600 different meme sounds and one of them was unfortunately labeled "Corona time"  
I would love to get the chance to remove that one string and upload an update. Sadly I can't because the app got suspended instantly.  
My two employees and I rely on the App because it accounts for 90% of our income.  
Without it we will probably have to close our small studio.  


We are very sorry for unknowingly violating the guidelines.  
Does someone has an idea on how to get our app back online (besides of writing a ticket)?   


Aaron Waller,  
FujiGames

https://preview.redd.it/u0vbc8kuk8o41.jpg?width=750&amp;format=pjpg&amp;auto=webp&amp;s=7ce6594ce006faf7a75647ef03ab6551fb4f0847
## [4][Anyone not get their admob payment this month?](https://www.reddit.com/r/androiddev/comments/fnixqf/anyone_not_get_their_admob_payment_this_month/)
- url: https://www.reddit.com/r/androiddev/comments/fnixqf/anyone_not_get_their_admob_payment_this_month/
---
I received my payment on the 15th for IAP but have not received ad income. I'm assuming it's because of this covid business..
## [5][Follow up: Transferring app to new Play account, what happens with purchase validation API calls?](https://www.reddit.com/r/androiddev/comments/fne1kx/follow_up_transferring_app_to_new_play_account/)
- url: https://www.reddit.com/r/androiddev/comments/fne1kx/follow_up_transferring_app_to_new_play_account/
---
As outlined in [this post](https://www.reddit.com/r/androiddev/comments/fmtab1/transferring_app_to_new_play_account_what_happens/) I was wondering when transferring apps between Google Play Developer accounts, **what happens with the existing Google API Console project?** Will in-app purchases just start failing until the new account creates and deploys its own API Console project? Does the old project get linked to the new account?

Well, I grit my teeth and jumped in. The transfer went through within about 8 hours (on a Sunday)... at which point we started seeing errors in the server logs and the Google API console. **Nobody could purchase our app anymore. We also couldn't process renewals.** Fantastic.

Even though I added the new account owner as the owner of the Google Play Android Developer API project, Google did not transfer the link to the new account. This resulted in 403 errors saying:

&gt;The project id used to call the Google Play Developer API has not been linked in the Google Play Developer Console.

We unlinked the old account and linked the new one, and gave the correct service accounts access. Even after doing so, we still had errors.

Finally, we added a dummy managed in-app purchase, inactive, to the apps under the new account. The errors went away, and we were able to then delete the dummy in-app purchase.

Google really needs to do a better job with this process, but we seem to be back up and running!
## [6][Is FAN banner placement practically impossible to use in mediation?](https://www.reddit.com/r/androiddev/comments/fngroe/is_fan_banner_placement_practically_impossible_to/)
- url: https://www.reddit.com/r/androiddev/comments/fngroe/is_fan_banner_placement_practically_impossible_to/
---
Hey, did anyone have luck using Facebook Audience Network banner ads with AdMob? I had &gt;200k impressions, 76% fill rate, 1,48 CTR and that didn’t pass FAN quality check. Support is useless, as they point to guidelines, which are useless.
## [7][Getting downloads - but don't know what the users want!](https://www.reddit.com/r/androiddev/comments/fniw77/getting_downloads_but_dont_know_what_the_users/)
- url: https://www.reddit.com/r/androiddev/comments/fniw77/getting_downloads_but_dont_know_what_the_users/
---
Hi guys,

I have a simple emoji translator app, in which you can select preset questions and get the emoji translations for it.
I am getting around 10 downloads a day mostly from Russia &amp; Ukraine but from US &amp; UK as well. 
Now all my users uninstall the app the same day. A couple of them have also left negative reviews and I am having a hard time understanding what is it that they want exactly.

Please help me out!

The apps name is Travemo

Thank you
## [8][How do you get context into ViewModel?](https://www.reddit.com/r/androiddev/comments/fniftf/how_do_you_get_context_into_viewmodel/)
- url: https://www.reddit.com/r/androiddev/comments/fniftf/how_do_you_get_context_into_viewmodel/
---
Is extending AndroidViewModel and using application context the most efficient solution?
## [9][Order list items in room db?](https://www.reddit.com/r/androiddev/comments/fnida1/order_list_items_in_room_db/)
- url: https://www.reddit.com/r/androiddev/comments/fnida1/order_list_items_in_room_db/
---
I've been trying to order list items in room database using `ItemTouchHelper`, where the item has a `position` property, I posted this before but I got no valid response [here](https://www.reddit.com/r/androiddev/comments/fcveue/is_there_a_way_to_update_a_column_of_entity_in/?utm_source=share&amp;utm_medium=web2x). Can someone please help me with this? 

I spent a lot of time trying to implement it but it just doesn't work, I searched a lot but everyone is just implementing swiping and ignoring the dragging.
## [10][Architecture for connecting touch events and openGLES (for pinch/scale, drag/pan)?](https://www.reddit.com/r/androiddev/comments/fnib8k/architecture_for_connecting_touch_events_and/)
- url: https://www.reddit.com/r/androiddev/comments/fnib8k/architecture_for_connecting_touch_events_and/
---
The effect is as in maps or a browser or an image viewer, where you can pinch-zoom and drag-pan.
I do it by handling pinch/drag events, then scaling and translating a matrix accordingly, which is applied in a vertex shader.

It's all working -- but the communicating is very awkward: from `onScale`/`onDrag` (in a `GLSurfaceView`) to `draw` (in a `Draw` object). I've just hacked a chain of objects to do it (`mRenderer.sim.coords.draw`), but there's got to be a better way!

Is there a standard way to do it? Thanks for any pointers!
## [11][Android App- WiFi Password Show For rooted phones](https://www.reddit.com/r/androiddev/comments/fniawv/android_app_wifi_password_show_for_rooted_phones/)
- url: https://www.reddit.com/r/androiddev/comments/fniawv/android_app_wifi_password_show_for_rooted_phones/
---
  You neer ROOT your device to use 'Wifi Password show'.Because on Android the wifi passwords are stored in "/data/misc/wifi/wpa\_supplicant.conf" or "/data/misc/wifi//WifiConfigStore.xml"  
 **Wifi password show app features:**  
Wifi password recovery  
share wifi password  
copy wifi password  
Show wifi info  
show wifi help section  
Generate qr code/barcode with wifi ssd and wifi password details 

 [https://play.google.com/store/apps/details?id=com.sbacham.srinu.showwifipassword2017](https://play.google.com/store/apps/details?id=com.sbacham.srinu.showwifipassword2017)
## [12][AndroidX camera in production?](https://www.reddit.com/r/androiddev/comments/fnhnpy/androidx_camera_in_production/)
- url: https://www.reddit.com/r/androiddev/comments/fnhnpy/androidx_camera_in_production/
---
Hello, I hope you all doing well :) We are planning to switch [Fotoappart](https://github.com/RedApparat/Fotoapparat) library for camerax, due to dropped maintenance of that library and some errors occurring in phones like Nokia. [Here](https://github.com/android/camera-samples/issues/17#issuecomment-556661413) is an error example.    


However, I doubt that it is still **production-ready** (*even though that API is beta now***)** and will not cause issues... Working with camera API, we have encountered various issues, like black previews, crashes, strange lifecycle issues ant etc.. Perhaps someone has been using androidx camera for a while and can share their experience?

Also, maybe I should rather use [CameraView](https://github.com/natario1/CameraView)? Seems the most stable and maintanable currently, while CameraKit and Fotoapparat seems dead :(
