# androiddev
## [1][App Feedback Thread - April 04, 2020](https://www.reddit.com/r/androiddev/comments/fusw1l/app_feedback_thread_april_04_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fusw1l/app_feedback_thread_april_04_2020/
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
## [2][Weekly Questions Thread - March 30, 2020](https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/
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
## [3][Google is trying to fix Android camera apps but it's not going to work - AndroidCentral](https://www.reddit.com/r/androiddev/comments/fva9rr/google_is_trying_to_fix_android_camera_apps_but/)
- url: https://www.androidcentral.com/google-trying-fix-android-camera-apps-its-not-going-work?amp
---

## [4][Chucker v3.2.0 is out 🚀 - with a lot of UI/perf. improvements, RTL support, URL decoding and much more.](https://www.reddit.com/r/androiddev/comments/fvbmw8/chucker_v320_is_out_with_a_lot_of_uiperf/)
- url: https://github.com/ChuckerTeam/chucker/releases/tag/3.2.0
---

## [5][I need to do some large math operations in my app, and i used AsyncTask till now, what is a good alternative?](https://www.reddit.com/r/androiddev/comments/fvb71f/i_need_to_do_some_large_math_operations_in_my_app/)
- url: https://www.reddit.com/r/androiddev/comments/fvb71f/i_need_to_do_some_large_math_operations_in_my_app/
---
The app is in Java, so can't use coroutines. Also i am a multithreading noob.
## [6][Greg Kroah Hartman AMA coming next week, write your questions now!](https://www.reddit.com/r/androiddev/comments/fvahmf/greg_kroah_hartman_ama_coming_next_week_write/)
- url: https://www.reddit.com/r/linux/comments/fu9sv5/greg_kh_ama_coming_next_week_write_your_questions/
---

## [7][Android Studio Layout Inspector](https://www.reddit.com/r/androiddev/comments/fuxi57/android_studio_layout_inspector/)
- url: https://medium.com/androiddevelopers/layout-inspector-1f8d446d048
---

## [8][Image from camera or gallery comes rotated sometimes](https://www.reddit.com/r/androiddev/comments/fvdic3/image_from_camera_or_gallery_comes_rotated/)
- url: https://www.reddit.com/r/androiddev/comments/fvdic3/image_from_camera_or_gallery_comes_rotated/
---
Sometimes, when I load an image from the gallery or the camera the image gets rotated not in a 90Degrees angle but just a little bit something like 5 Degrees and I can't explain it since it's not even happening consistently but random, and it has nothing to do with the actual image sometimes it rotates a certain image and sometimes it doesn't rotate the same image. I already have tried setting the rotation to the actual rotation of the image after selecting but somehow it doesn't care and I can't find people with the same problem, so thanks in advance for everyone helping me with this.

Here is my code for the selection of the image:

[code](https://hastebin.com/bahunifanu.cs)
## [9][Android Navigation component button navigation to another fragment and back shows blank screen](https://www.reddit.com/r/androiddev/comments/fvdhxq/android_navigation_component_button_navigation_to/)
- url: https://www.reddit.com/r/androiddev/comments/fvdhxq/android_navigation_component_button_navigation_to/
---
Am using navigation component to handle navigation in my app, i have a user profile with two buttons , one is to navigate to users posts when clicked and the other is to log out the user 

`class MeFragment : BaseFragment() {`  


`override fun onCreateView(`  
`inflater: LayoutInflater, container: ViewGroup?,`  
 `savedInstanceState: Bundle?`  
`): View? {`  
 `// Inflate the layout for this fragment`  
 `return inflater.inflate(R.layout.fragment_me, container, false)`  
`}`  


`override fun onViewCreated(view: View, savedInstanceState: Bundle?) {`  
 `super.onViewCreated(view, savedInstanceState)`  


`myProperty.setOnClickListener {`  
 `val action = MeFragmentDirections.actionMeFragmentToMyUploadHousesFragment()`  
 `findNavController().navigate(R.id.myUploadHousesFragment)`  
 `}`  
 `logout.setOnClickListener {`  
 `FirebaseAuth.getInstance().signOut()`  
 `}`  
 `}`  


`}`  
but when user clicks on user profile they can navigate just fine, issue is when they click from the posts back to the meFragment, the whole app breaks and shows a blank, here is my navigation component graph 

&amp;#x200B;

`&lt;?xml version="1.0" encoding="utf-8"?&gt;`  
`&lt;navigation xmlns:android="http://schemas.android.com/apk/res/android"`  
 `xmlns:app="http://schemas.android.com/apk/res-auto"`  
 `xmlns:tools="http://schemas.android.com/tools"`  
 `android:id="@+id/main_graph.xml"`  
 `app:startDestination="@id/onBoarding"&gt;`  
`&lt;action`  
 `android:id="@+id/action_logout"`  
 `app:destination="@id/authFragment"`  
 `app:enterAnim="@anim/nav_default_enter_anim"`  
 `app:exitAnim="@anim/nav_default_exit_anim"`  
 `app:popEnterAnim="@anim/nav_default_pop_enter_anim"`  
 `app:popExitAnim="@anim/nav_default_pop_exit_anim"`  
 `app:popUpTo="@id/main_graph.xml"`  
 `app:popUpToInclusive="true" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/onBoarding"`  
 `android:name="org.hero76.zedhousely.ui.fragments.onBoarding.OnBoarding"`  
 `android:label="OnBoarding"`  
 `tools:layout="@layout/fragment_on_boarding"&gt;`  
`&lt;action`  
 `android:id="@+id/action_onBoarding_to_authFragment"`  
 `app:destination="@id/authFragment"`  
 `app:enterAnim="@anim/nav_default_enter_anim"`  
 `app:exitAnim="@anim/nav_default_exit_anim"`  
 `app:popEnterAnim="@anim/nav_default_pop_enter_anim"`  
 `app:popExitAnim="@anim/nav_default_pop_exit_anim"`  
 `app:popUpTo="@id/main_graph.xml"`  
 `app:popUpToInclusive="true" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/authFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.authHost.AuthFragment"`  
 `android:label="AuthFragment"`  
 `tools:layout="@layout/fragment_auth"&gt;`  
`&lt;action`  
 `android:id="@+id/action_authFragment_to_phoneVerificationFragment"`  
 `app:destination="@id/phoneVerificationFragment"`  
 `app:enterAnim="@anim/nav_default_enter_anim"`  
 `app:exitAnim="@anim/nav_default_exit_anim"`  
 `app:popEnterAnim="@anim/nav_default_pop_enter_anim"`  
 `app:popExitAnim="@anim/nav_default_pop_exit_anim"/&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/loginFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.login.LoginFragment"`  
 `android:label="LoginFragment"`  
 `tools:layout="@layout/fragment_login" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/signUpFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.signUp.SignUpFragment"`  
 `android:label="SignUpFragment"`  
 `tools:layout="@layout/fragment_sign_up" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/homeFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.home.HomeFragment"`  
 `android:label="Home"`  
 `tools:layout="@layout/fragment_home" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/homepageFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.homepage.HomepageFragment"`  
 `android:label="fragment_homepage"`  
 `tools:layout="@layout/fragment_homepage" &gt;`  
`&lt;action`  
 `android:id="@+id/action_homepageFragment_to_categoryFragment"`  
 `app:destination="@id/categoryFragment"`  
 `app:popUpToInclusive="false" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/meFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.me.MeFragment"`  
 `android:label="fragment_me"`  
 `tools:layout="@layout/fragment_me" &gt;`  
`&lt;action`  
 `android:id="@+id/action_meFragment_to_myUploadHousesFragment"`  
 `app:destination="@id/myUploadHousesFragment" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/postFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.post.PostFragment"`  
 `android:label="fragment_post"`  
 `tools:layout="@layout/fragment_post" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/searchFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.search.SearchFragment"`  
 `android:label="fragment_search"`  
 `tools:layout="@layout/fragment_search" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/servicesFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.services.ServicesFragment"`  
 `android:label="fragment_services"`  
 `tools:layout="@layout/fragment_services" /&gt;`  
`&lt;fragment`  
 `android:id="@+id/otpVerificationFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.phone.OtpVerificationFragment"`  
 `android:label="OtpVerificationFragment" &gt;`  
`&lt;action`  
 `android:id="@+id/action_otpVerificationFragment_to_signUpFragment"`  
 `app:destination="@id/signUpFragment" /&gt;`  
`&lt;action`  
 `android:id="@+id/action_otpVerificationFragment_to_homeFragment"`  
 `app:destination="@id/homeFragment" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/phoneVerificationFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.phone.PhoneVerificationFragment"`  
 `android:label="PhoneVerificationFragment" &gt;`  
`&lt;action`  
 `android:id="@+id/action_phoneVerificationFragment_to_otpVerificationFragment"`  
 `app:destination="@id/otpVerificationFragment" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/myUploadHousesFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.me.MyUploadHousesFragment"`  
 `android:label="fragment_my_upload_houses"`  
 `tools:layout="@layout/fragment_my_upload_houses" &gt;`  
`&lt;action`  
 `android:id="@+id/action_myUploadHousesFragment_pop"`  
 `app:popUpTo="@id/meFragment"`  
 `app:popUpToInclusive="true" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;fragment`  
 `android:id="@+id/categoryFragment"`  
 `android:name="org.hero76.zedhousely.ui.fragments.category.CategoryFragment"`  
 `android:label="CategoryFragment" &gt;`  
`&lt;action`  
 `android:id="@+id/action_categoryFragment_to_homepageFragment"`  
 `app:destination="@id/homepageFragment" /&gt;`  
`&lt;/fragment&gt;`  
`&lt;/navigation&gt;`

&amp;#x200B;

`Asking for help on how to fix the issue.`
## [10][What subjects do I need to learn for a stick balance game?](https://www.reddit.com/r/androiddev/comments/fvdafj/what_subjects_do_i_need_to_learn_for_a_stick/)
- url: https://www.reddit.com/r/androiddev/comments/fvdafj/what_subjects_do_i_need_to_learn_for_a_stick/
---
Our professor wanted us to build something simple with Android sensors. A simple app that displays sensor values would probably do just fine but I want to build a fun app, an app that would teach me more than a 5-10 minutes tutorial does.

So I decided to build a stick balance game where a long stick slowly tips one side and the player uses accelerometer to try to keep it standing up.

I am a novice Android dev but I thought it is still doable for someone on my level, with a bit of a challenge of course. 

Now, my question is, what topics should I be looking at to build something like this? I hope I don't have to use a physics engine or something
## [11][How to add margin to the MaterialButton inside MaterialButtonToggleGroup?](https://www.reddit.com/r/androiddev/comments/fvd614/how_to_add_margin_to_the_materialbutton_inside/)
- url: https://www.reddit.com/r/androiddev/comments/fvd614/how_to_add_margin_to_the_materialbutton_inside/
---
It's not working by setting `android:layout_margin="16dp"` to the buttons
## [12][[Lyla Fujiwara / Shailen Tull] Android Conference Talks - Easy Android accessibility](https://www.reddit.com/r/androiddev/comments/fv57zq/lyla_fujiwara_shailen_tull_android_conference/)
- url: https://www.youtube.com/watch?v=yxNROzu9nQQ
---

