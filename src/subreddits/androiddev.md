# androiddev
## [1][App Feedback Thread - June 06, 2020](https://www.reddit.com/r/androiddev/comments/gxq3m1/app_feedback_thread_june_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/gxq3m1/app_feedback_thread_june_06_2020/
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
## [2][Community Megathread](https://www.reddit.com/r/androiddev/comments/gwpxlc/community_megathread/)
- url: https://www.reddit.com/r/androiddev/comments/gwpxlc/community_megathread/
---
Good morning/afternoon/evening everyone!

Let's get right into it. Recent events have lead to a lot of debate and deliberation internally and externally. I'd like to reach out to everyone and open a dialogue between us and the community.

We will not be allowing several posts discussing the subreddit and past events, this is not the proper method to reach us, and I don't want to stifle or drown out the great discussion that happens here with too many posts. Instead, I'd like to open this thread as a place to discuss. In response to past events I would like to state the following will be happening in short order.

 * We will be restructuring our leadership internally as some mods have differing activity levels and some wish to retire. We recognize that we are also severely understaffed which is hurting our ability to serve the community, so we will soon be recruiting additional volunteers from the community to help out. More on this will be announced soon.

 * Any action we take is as a team. At the end of the day we are volunteers doing this in our free time with the best interests of our community in mind. With everything that is going on in the world right now, now is not time for bickering, from anyone. Now is the time for coming together and solving problems. Remember that everyone is a human being. Harassment is zero tolerance.

 * In response to the above point, I would like to ask for everyone's feedback on our current rule set in the comment below. Please keep the discussion calm and collected, or it will be unproductive and removed. I am however encouraging everyone to provide their feedback and suggestions on how we can improve our community.

Expect to see more from me personally as I take a bigger role in trying to help restructure our team and improve our community. 

Have a great day everyone!
## [3][JetBrains is going to monetize JetBrains Academy register quickly to get free access until Jan 1st](https://www.reddit.com/r/androiddev/comments/gy70u5/jetbrains_is_going_to_monetize_jetbrains_academy/)
- url: https://blog.jetbrains.com/blog/2020/06/02/jetbrains-academy-is-getting-ready-to-hit-the-market/
---

## [4][Android Architecture: MVP to MVVM](https://www.reddit.com/r/androiddev/comments/gy65bu/android_architecture_mvp_to_mvvm/)
- url: https://medium.com/@kgaurav23/android-architecture-journey-from-mvp-to-mvvm-29bf469a8b59
---

## [5][Google is OK with violating the GDPR?](https://www.reddit.com/r/androiddev/comments/gy9uro/google_is_ok_with_violating_the_gdpr/)
- url: https://i.redd.it/q3vf7oqssg351.png
---

## [6][[DEV] I made a wrapper around official Spotify android SDK](https://www.reddit.com/r/androiddev/comments/gy8xnx/dev_i_made_a_wrapper_around_official_spotify/)
- url: https://www.reddit.com/r/androiddev/comments/gy8xnx/dev_i_made_a_wrapper_around_official_spotify/
---
Recently for some reason, I've to implement Spotify API in my personal android app project. So I went to the official android [docs](https://developer.spotify.com/documentation/android/) and saw that they actually provide an authentication library.

It's good until I came across a problem regarding refreshing an access token like there is no built-in way to generate an access token from refresh token, also for capturing result from OAuth we've to override onActivityResult() which is OK, but I think there is a better solution to use [ActivityResultContract](https://developer.android.com/reference/androidx/activity/result/contract/ActivityResultContracts.StartActivityForResult) from the latest androidx [Activity](https://developer.android.com/jetpack/androidx/releases/activity) (alpha release) which lets you register &amp; listens for Activity Result Callback.

There are other limitations regarding the official SDK which I've mentioned in my GitHub repo's readme.

So I thought to create a wrapper overcoming these limitations. The library is open-source and licensed under GPL v3. It is also well documented, so if you've any spare time take a look at it and let me know how I can improve it further.

Library - [https://github.com/KaustubhPatange/Unofficial-Spotify-SDK](https://github.com/KaustubhPatange/Unofficial-Spotify-SDK)
## [7][1.54k requests in AdMob and my app is not live? It was sent to 5 people to test quickly. Where are they coming from? I do not want to violate any rules.](https://www.reddit.com/r/androiddev/comments/gxrptc/154k_requests_in_admob_and_my_app_is_not_live_it/)
- url: https://i.redd.it/gyicf1ugsa351.png
---

## [8][TheMovieDB: An app for TMDB API featuring MVVM, Koin and AAC](https://www.reddit.com/r/androiddev/comments/gy9qxy/themoviedb_an_app_for_tmdb_api_featuring_mvvm/)
- url: https://www.reddit.com/r/androiddev/comments/gy9qxy/themoviedb_an_app_for_tmdb_api_featuring_mvvm/
---
Hey guys, I'm working on this app for a long time and it has always been a playground for learning new things. Some parts of the app are very well written and other parts are just horrible IMO. Still, here it is hope you guys like it, and please point out any improvements that you guys think can be made (or just something that I did horribly wrong).

EDIT: [https://github.com/skrilltrax/TheMovieDB](https://github.com/skrilltrax/TheMovieDB)
## [9][Android distribution chart new location](https://www.reddit.com/r/androiddev/comments/gy7plh/android_distribution_chart_new_location/)
- url: https://androiddistribution.io
---

## [10][I have a jar which depends on old Android support, and there is no way I can update the Jar. What are the possible solutions?](https://www.reddit.com/r/androiddev/comments/gybvlo/i_have_a_jar_which_depends_on_old_android_support/)
- url: https://www.reddit.com/r/androiddev/comments/gybvlo/i_have_a_jar_which_depends_on_old_android_support/
---
I'm building an app and it depends on a jar/rar file, which uses old \`ActionBarActivity\`  from \`android.support.v7.app\` package.  Now this class has been deprecated and removed (I can't see it in AndroidX library). 

&amp;#x200B;

What are my options to continue using the jar?.

**Note:** I'm using AndroidX libraries

&amp;#x200B;

**What I have tried**

1. Copy pasting the \`ActionBarActivity\` to my project under \`androidx.appcompat.app\` package name (Androidx because **Jetifier** is enabled.) However lot of the APIs that \`ActionBarActivity\` is using is blocked by \`@RestrictTo\` annotation. So it gives me a error. Not sure if I can get pass this error!!
## [11][Issue in split screen](https://www.reddit.com/r/androiddev/comments/gyaki2/issue_in_split_screen/)
- url: https://www.reddit.com/r/androiddev/comments/gyaki2/issue_in_split_screen/
---
When I use android:windowFullScreen true in activity style it works perfectly in normal screen mode i.e., it hides the status bar but in split screen mode status bar becomes visible and content is being drawn beneath the status bar. Is there any recommended way to handle this? I have tried official documentation around this and couldn't find anything
## [12][Adding support functionality in the app](https://www.reddit.com/r/androiddev/comments/gyaej8/adding_support_functionality_in_the_app/)
- url: https://www.reddit.com/r/androiddev/comments/gyaej8/adding_support_functionality_in_the_app/
---
I've seen an app named **FASTHUB** which has a support us, a screen where you can donate to the dev, is that allowed? 

Payment is through **Google Play Billing.**

Thanks!
