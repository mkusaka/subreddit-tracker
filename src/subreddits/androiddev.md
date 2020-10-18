# androiddev
## [1][App Feedback Thread - October 17, 2020](https://www.reddit.com/r/androiddev/comments/jcua3k/app_feedback_thread_october_17_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jcua3k/app_feedback_thread_october_17_2020/
---
This thread is for getting feedback on your own apps.

####Developers:

- must **provide feedback** for others
- must include **Play Store**, **GitHub**, or **BitBucket** link
- must make top level comment
- must make effort to respond to questions and feedback from commenters
- may be open or closed source

####Commenters:

- must give **constructive feedback** in replies to top level comments
- must not include links to other apps

To cut down on spam, accounts who are too young or do not have enough karma to post will be removed. Please make an effort to contribute to the community before asking for feedback.

As always, the mod team is only a small group of people, and we rely on the readers to help us maintain this subreddit. Please report any rule breakers. Thank you.

\- Da Mods
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
## [3][Here is my custom prototype version of DOOM running on an android headunit...in a car!](https://www.reddit.com/r/androiddev/comments/jdeix0/here_is_my_custom_prototype_version_of_doom/)
- url: https://youtu.be/5mmiPT2avrY
---

## [4][Create a google play account for the company instead of individual](https://www.reddit.com/r/androiddev/comments/jdctrb/create_a_google_play_account_for_the_company/)
- url: https://www.reddit.com/r/androiddev/comments/jdctrb/create_a_google_play_account_for_the_company/
---
 

Hello,

I'm having some difficulties trying to register a google play account for my company. I've started by creating a new google account by selecting "To manage my business" and after that I've selected payment method and added a new card. This has automatically created a payment profile with the account type as individual with no option to change it to business. 

I looked over different help sections and from what I saw, I should have seen an option to chose from individual/business when creating the payment profile. I deleted the payment profile and create it again several times and couldn't see any option.

Did someone else had the same problem and could provide some exact steps on how can I create a payment profile with the account type as business?
## [5][If you use Glide, and if you use WebView, then you should use them together. Checkout the link and share your reviews.](https://www.reddit.com/r/androiddev/comments/jddp29/if_you_use_glide_and_if_you_use_webview_then_you/)
- url: https://muditsen.medium.com/enhance-android-webview-performance-using-glide-aba4bbc41bc7
---

## [6][Android CodeView: an Easy way to create Code Editor app](https://www.reddit.com/r/androiddev/comments/jdc694/android_codeview_an_easy_way_to_create_code/)
- url: https://medium.com/@amrdeveloper/android-codeview-an-easy-way-to-create-code-editor-app-5d67c3534f84
---

## [7][Loading Arraylist of data from firebase using ListIterator returning null](https://www.reddit.com/r/androiddev/comments/jdem7s/loading_arraylist_of_data_from_firebase_using/)
- url: https://www.reddit.com/r/androiddev/comments/jdem7s/loading_arraylist_of_data_from_firebase_using/
---
Hi Im asking my question here cause , its been 3 days and I dont find any answer on my own  , hope one of you could help me 

I want to load post from my Users inside an Recyclerview ; I'm loading post corretly . Inside of a Post I got the Firebase uid of the user who posted the post , I added all post users uid inside a arraylist and im trying to load all Users info data inside a model User . Before loading all data from user uid , Im trying to see if the user sill exist inside the database with a ondatange methode inside ListIterator , but some id who are still inside my database return null where they should return true , someone can explain me why ? Sorry for my bad english

 

and here a link for my [question closed](https://stackoverflow.com/questions/64404330/loading-arraylist-of-data-from-firebase-using-listiterator-returning-null)

&amp;#x200B;

    public void CheckifUserstillexist(FragmentActivity activity, List&lt;String&gt; listuidusers) {     if (list.size() != 0){         List&lt;ModelUserProfil &gt; Listofusers = new Arraylist&lt;&gt;;          int i = 0;         for (ListIterator&lt;String&gt; it = list.listIterator(); it.hasNext(); i++) {              String UserUid= it.next();              isAccountExistSingleValue(UserUid, new OnObjectExistListener&lt;String&gt;() {                 @Override                 public void onDataChanged(boolean exist) {                      ModelUserProfil  Userprofil  =  new ModelUserProfil  ();                     Userprofil.setUserid(UserUid);                     Userprofil.setisAccountExisting(exist);                     Listofusers.add(Userprofil);                      if (!it.hasNext()) {                         ListIterator&lt;ModelUserProfil&gt; TheBigList  = Listofusers.listIterator();                         Collection&lt;ModelUserProfil  &gt; itemtodeletefromlist = new ArrayList&lt;&gt;();                          while (TheBigList.hasNext()) {                             ModelUserProfil  User = TheBigList.next();                             if (User .getisAccountExisting() != null){                                  if (!User.getisAccountExisting()){                                      itemtodeletefromlist .add(OldList);                                 }                             }else {                                 Useritemtodeletefromlist.add(OldList);                                 // Items that should not be add here  are add here                              }                         }                          if (Useritemtodeletefromlist.size() != 0){                             Listofusers.removeAll(itemtodeletefromlist );                         }                         LoadUsersInfo(Listofusers);                     }                 }             });         }     }else {         // List empty     } }
## [8][Google removed app by mistake](https://www.reddit.com/r/androiddev/comments/jcsasd/google_removed_app_by_mistake/)
- url: https://www.reddit.com/r/androiddev/comments/jcsasd/google_removed_app_by_mistake/
---
# TL;DR

Google wrongfully removed an app from the Google Play Store. The reviewer somehow made a mistake and even documented it with screenshots attached to the removal notice. With no way to contact Google, the app is off the app store for hours, money being lost every hour. Has Google completely lost it?

# Long story

The app metrics looked slightly odd all Friday, but I thought, well, it’s a Friday and with ever changing Corona lockdowns all over the globe, we live in odd times anyway.

Until I realized at night that the user count had dropped significantly. I searched on Google Search for the app store link, but clicking on the result led to an error page:

&gt; We're sorry, the requested URL was not found on this server.

I then searched on the Google Play Store for the app, and it really could not be found. Rushing to the Google Play Developer Console I see a notice that the app has been removed and I should have received an email with details.

&gt; Your app has been removed from Google Play for not adhering to the Google Play Developer Program policies. New users can't find and install your app, and existing users won't receive updates. We've sent more information to the account owner's email address.

I looked for the email but I could find any, nor in the Trash, Junk, Archive. I later found out that Google is apparently falsely claiming to send out these emails, because so many developers [complain](https://support.google.com/googleplay/thread/32034066?hl=en) that they didn’t receive one. I noticed that I hadn’t opted into the new Developer Console website, and on the new site I actually found a link to display the email that was allegedly sent out.

The reasons for removal were related to the User Generated Content policy, specifically:
- the app doesn’t require a user to accept the Terms of Service 
- the Terms of Service do not state the restrictions on user generated content 

Both claims are false. The reviewer attached some screenshots of the walkthrough that is displayed after app launch as proof. The screen where the user has to accept the TOS and even explicitly agree to a summary regarding restrictions on user generated content was simply missing. It is not even technically possible to skip it.

How could it happen that an app with millions of downloads is removed without prior notice due to a single reviewer’s mistake? Looking into the server logs it seems that Google off-shores app reviewers to the Philippines. Going offshore into low labor cost markets is a common practice, often with implications on the quality of work. Procedurally, the removal should not have happened, but Google’s internal review processes are apparently a mess. They only recently introduced a more stringent app review process (rightly so) but this seems to be plagued by [misinformation and lack of transparency](https://www.androidpolice.com/2019/09/03/google-play-expands-store-app-review/) ever since.

Luckily, the removal notice contains an appeal link to quickly resolve this mistake - or so I thought:

&gt; If you’ve reviewed the policy and feel our decision may have been in error, please reach out to our policy support team. We'll get back to you within 2 business days.

I submitted an appeal with screenshots showing that there is in fact a TOS screen. But already at the beginning of the appeal form, Google actually discourages from filing an appeal:

&gt; For a quicker resolution, we highly encourage you to refer to the original notification email to see how you can resubmit your app without an appeal. In most cases, you can make changes to bring your app into compliance and submit your app again.

Google even requires to check a box that one has read the above and agrees to it. Whoever of their legal team though of that and for what reason, but it makes an appeal look fishy - against which claims are they trying to protect themselves?

I try to get more clarification on that whole mess by heading back to the Developer Console and opening the Developer Support Chat. The link now leads to a new form that I fill out to get connected to an agent - or so I thought, because after filling out the whole form, it fails to submit:

&gt; There were problems sending your form. Please try again.

I remember the same error from some months ago and now that I need something urgent, of course it still doesn’t work. Random? Doesn’t look like it. Google seems to be trying everything they can to insulate themselves from contact with their developers.

At this point I should mention that I have a call-through phone number to our reviewer on the Apple App Review Team for urgent cases regarding the Apple App Store, a stark contrast in company culture and how they value developer relations.

I am left to submit the appeal, but it is Friday night and with the weekend ahead, when will the app be available again? Mid next week? End of next week? Due to COVID19 even later? Impossible to wait. So I quickly compile the unchanged app code with a new app version code and submit the update for review, in the hope that this works faster as Google itself suggests.

It is the next day, the app is now removed from the store for approximately 24 hours and still in review. It is a Saturday, with weekends usually generating the highest download numbers, customer complaints are flooding in why the app is unavailable on the Google Play Store. Every hour waiting means more upset customers and more money lost.

Is this the way Google intents to treat their developers? How is it possible that a company with a clear monopolistic market dominance has free reign without proper regulation to impose arbitrary rules and low quality processes that are lacking due diligence on hundreds of thousands of developers who are economically dependent on their platform?

There is something severely wrong here and it needs to be changed, through government policy, because self-regulation is a myth.

# Update
Google has reinstated the app after approximately 24h.

It is unclear at this point whether ~24h was the usual response time or whether either the appeal or the new build upload did speed things up. It is unlikely that this post has contributed because I did not reveal the app name to users who have reached out to me privately and asked for it.

The app has been unavailable for ~24h, which has a monetary impact that I would preliminarily estimate to be a 4 figure USD loss from what I can see in the analytics so far. The number seems large enough to bother but too small to follow up on it considering the effort required to argue with Google, which probably has a low chance of success anyway. Another contra argument is possible retaliation from Google, which could be so sublime that it would be difficult to proof.

This predicament is just another example that proper legislation is needed to protect the weaker ones, in this case the developers. And it should start at the basics: one of the largest tech companies in the world with an app store market size of billions of dollars and their developer contact form does not work for months, just as they are rolling out a more stringent app review process that seems to be poorly executed? Random? Possibly, but unlikely. Why? Because they can.
## [9][Why is hilt the officially recommended DI library even though its in alpha?](https://www.reddit.com/r/androiddev/comments/jdc5s4/why_is_hilt_the_officially_recommended_di_library/)
- url: https://www.reddit.com/r/androiddev/comments/jdc5s4/why_is_hilt_the_officially_recommended_di_library/
---
Hilt is now being recommended throughout Android docs (under Best practices -&gt; Dependency Injection &amp; Guide to app architecture) to be used as the DI library. However, the library is still in alpha and as I understand it, its API can still drastically change, correct?  


For new apps should I start using Hilt or is it safer bet to use existing established DI libraries?
## [10][Review Times Google Play?](https://www.reddit.com/r/androiddev/comments/jd5v94/review_times_google_play/)
- url: https://www.reddit.com/r/androiddev/comments/jd5v94/review_times_google_play/
---
Since COVID, it seems like review times have gone way up. Is this still the case, how long have you guys had to wait?
## [11][11 — 17 October Android Newsletter](https://www.reddit.com/r/androiddev/comments/jdd0w8/11_17_october_android_newsletter/)
- url: https://www.reddit.com/r/androiddev/comments/jdd0w8/11_17_october_android_newsletter/
---
Stay up to date with Android development, in this week's edition:  
💨 Sharing Shortcuts  
🗺️ Navigation Component  
🥰 StateFlow &amp; UI  
💡 Kotlin Scope Functions  
and much more!

Read it here 👉[https://vladsonkin.com/android-newsletter-16/](https://vladsonkin.com/android-newsletter-16/)  
What's your favorite article in this edition? I would love to see your comments!

🔥Featuring [@RaulHernandezL](https://twitter.com/RaulHernandezL) [@chethaase](https://twitter.com/chethaase) [@safaorhanEN](https://twitter.com/safaorhanEN) [@tfcporciuncula](https://twitter.com/tfcporciuncula) [@informramiz](https://twitter.com/informramiz) [@SG5202](https://twitter.com/SG5202) [@thecodeside\_AL](https://twitter.com/thecodeside_AL) and many other great authors!

💚 Subscribe and receive new editions directly to your email. Weekly, no spam, unsub anytime.  
Here is an example: [https://mailchi.mp/0426e7d6455d/android-newsletter-16](https://mailchi.mp/0426e7d6455d/android-newsletter-16)
## [12][Why I Almost Gave Up On My Indie Game - Devlog](https://www.reddit.com/r/androiddev/comments/jcymy4/why_i_almost_gave_up_on_my_indie_game_devlog/)
- url: https://www.youtube.com/watch?v=3LG4zlCfnmw
---

