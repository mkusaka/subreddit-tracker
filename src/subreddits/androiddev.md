# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/emqf2y/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/emqf2y/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - January 06, 2020](https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/
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
## [3][Is there a way for developers to programmatically avoid triggering Huawei's HiTouch without foregoing the entire idea of having two fingers on the screen at the same time?](https://www.reddit.com/r/androiddev/comments/emn8ho/is_there_a_way_for_developers_to_programmatically/)
- url: https://www.reddit.com/r/androiddev/comments/emn8ho/is_there_a_way_for_developers_to_programmatically/
---

## [4][First commits of ConstraintLayout in AOSP Jetpack Compose repo](https://www.reddit.com/r/androiddev/comments/emo2do/first_commits_of_constraintlayout_in_aosp_jetpack/)
- url: https://www.reddit.com/r/androiddev/comments/emo2do/first_commits_of_constraintlayout_in_aosp_jetpack/
---
[Here's a merged commit](https://android-review.googlesource.com/c/platform/frameworks/support/+/1147725) for `ConstraintLayout` in Jetpack Compose. You can see an [example usage here](https://android-review.googlesource.com/c/platform/frameworks/support/+/1147725/22/ui/ui-layout/integration-tests/layout-demos/src/main/java/androidx/ui/layout/demos/ConstraintLayoutActivity.kt). It looks like they separated out declaring the children (which you do in the trailing lamda) from declaring the constraints (which you do as part of the `ConstraintSet` parameter to the `ConstraintLayout` composable function). They use the `Tag` modifier to link them together. My understanding is that the tag plays the same role as IDs did for referring to other views in the  ConstraintLayout XML.

Of course - this is all very early and a lot could change. And of course a lot of stuff is missing (most notably, animations).
## [5][Update not live yet, a week after submission - anyone else?](https://www.reddit.com/r/androiddev/comments/emm7em/update_not_live_yet_a_week_after_submission/)
- url: https://www.reddit.com/r/androiddev/comments/emm7em/update_not_live_yet_a_week_after_submission/
---
Hey all - is this happening to anyone else? The update I published after new year's has not gone live yet.

&amp;#x200B;

I usually receive the "update is live" message within 24 hours.

&amp;#x200B;

For reference, there were only a couple of crash fixes in the new update so not sure why this would happen. The app has been out there for 5 years or so.
## [6][Observing datastores with coroutines: Jetpack's Room now supports Flow in Kotlin](https://www.reddit.com/r/androiddev/comments/emcyda/observing_datastores_with_coroutines_jetpacks/)
- url: https://medium.com/androiddevelopers/room-flow-273acffe5b57
---

## [7][What actions have you taken to comply with CCPA? Will you take any actions?](https://www.reddit.com/r/androiddev/comments/emnne1/what_actions_have_you_taken_to_comply_with_ccpa/)
- url: https://www.reddit.com/r/androiddev/comments/emnne1/what_actions_have_you_taken_to_comply_with_ccpa/
---

## [8][Flutter and Android docker images with support for emulators, VNC, Fastlane and more.](https://www.reddit.com/r/androiddev/comments/emqjyi/flutter_and_android_docker_images_with_support/)
- url: https://www.reddit.com/r/androiddev/comments/emqjyi/flutter_and_android_docker_images_with_support/
---
Hi,  
I have published docker images for android and flutter.  


The docker images contains sdks and also tools that will help you:  
\* automate code review, code style review, execute tests and generate report and also html report.  
\* docker images also has support for android emulators even on device without KVM.  
\* support for VNC, FTP, QR code generation and more.

Android images: [https://hub.docker.com/repository/docker/bitsydarel/android](https://hub.docker.com/repository/docker/bitsydarel/android)  
Flutter images: [https://hub.docker.com/repository/docker/bitsydarel/flutter-ci](https://hub.docker.com/repository/docker/bitsydarel/flutter-ci)  


Docker allow you to build and run software without compromizing your host computer.
## [9][Converting PyTorch float tensor to Android RGBA Bitmap with Kotlin](https://www.reddit.com/r/androiddev/comments/empazq/converting_pytorch_float_tensor_to_android_rgba/)
- url: https://medium.com/@philipplies/converting-pytorch-float-tensor-to-android-rgba-bitmap-with-kotlin-ffd4602a16b6?source=friends_link&amp;sk=685a9b193132d2db33110e6dc4933d90
---

## [10][Scratch That Itch - An intro to scratch files.](https://www.reddit.com/r/androiddev/comments/em9ywc/scratch_that_itch_an_intro_to_scratch_files/)
- url: https://zarah.dev/2019/11/21/scratch-files.html
---

## [11][Google Play remove my app because app upload user data to another app what I never use it in my code? Why?](https://www.reddit.com/r/androiddev/comments/emkjep/google_play_remove_my_app_because_app_upload_user/)
- url: https://www.reddit.com/r/androiddev/comments/emkjep/google_play_remove_my_app_because_app_upload_user/
---
Google Play remove my app because app upload user data to https://apps.applozic.com without a prominent disclosure. but I never do that. And there is noy any code about https://apps.applozic.com. Why google play do this decision?  do anyone know this app(https://apps.applozic.com)?
## [12][Firebase specific question for importing a String](https://www.reddit.com/r/androiddev/comments/emnwd9/firebase_specific_question_for_importing_a_string/)
- url: /r/Firebase/comments/emhhym/android_coding_question_reading_from_firebase/
---

