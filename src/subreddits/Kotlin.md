# Kotlin
## [1][Coroutines and exceptions: things to know](https://www.reddit.com/r/Kotlin/comments/g3mbju/coroutines_and_exceptions_things_to_know/)
- url: https://www.rockandnull.com/coroutines-and-exceptions/
---

## [2][MVI in Kotlin Multiplatform — part 1](https://www.reddit.com/r/Kotlin/comments/g35sle/mvi_in_kotlin_multiplatform_part_1/)
- url: https://badootech.badoo.com/mvi-in-kotlin-multiplatform-part-1-1-of-3-205c6feb4ac7
---

## [3][[Newbie-question] [Kotlin/Android] ListView + setOnItemClickListener -&gt; how to get value?](https://www.reddit.com/r/Kotlin/comments/g3bths/newbiequestion_kotlinandroid_listview/)
- url: https://www.reddit.com/r/Kotlin/comments/g3bths/newbiequestion_kotlinandroid_listview/
---
I'm new to Android/Kotlin and try to understand ListView:

let say my App shows:

yellow

blue

orange

....

in a ListView.

\-&gt; when I use:

mylist.setOnItemClickListener **{** parent, **view**, position, id **-&gt;**  
 Toast.makeText(this, "Clicked item/fruit is **:"**\+**" "**\+position+**view**,Toast.*LENGTH\_SHORT*).show()  
}

&amp;#x200B;

I get the position (0/1/2/3/...) + the view as an array from the clicked item.

How can I get the text-value that is shown  (e.g. "yellow", "blue" or "orange")  of the item?

thx
## [4][In march, the german government started a hackathon to develop solutions for the corona crisis. We are one of the top 20 solutions and are in need of volunteers for Kotlin development!](https://www.reddit.com/r/Kotlin/comments/g2jwq3/in_march_the_german_government_started_a/)
- url: https://www.reddit.com/r/Kotlin/comments/g2jwq3/in_march_the_german_government_started_a/
---
Hi, I'm on the volunteer team for one of the top 20 projects in the German government's Wir Vs Virus Hackathon. We (nexd) are developing a platform to coordinate grocery shopping helpers for all those who cannot leave the house at the moment. The twist here is that even people without a smartphone and technical knowledge can submit their shopping list via a simple phone call. This way, we reach the people most affected by the crisis: seniors.

We are about to launch the beta test but are in dire need of Kotlin developers. If you are interested in helping (our slack is in English) please contact mail@nexd.app More information about the project (German only, sadly): https://www.youtube.com/watch?v=0PmuLrq4Hno

Thank you and stay healthy!
## [5][Gradle &amp; Kotlin — create tests artifact](https://www.reddit.com/r/Kotlin/comments/g2ygpq/gradle_kotlin_create_tests_artifact/)
- url: https://medium.com//gradle-kotlin-create-tests-artifact-1e22bb00b1e1?source=friends_link&amp;sk=5bc717c40eb14b24c190e686d2834f21
---

## [6][gRPC, meet Kotlin | Google Cloud Blog](https://www.reddit.com/r/Kotlin/comments/g2mbe9/grpc_meet_kotlin_google_cloud_blog/)
- url: https://cloud.google.com/blog/products/application-development/use-grpc-with-kotlin
---

## [7][My first Kotlin Project: ServerNeat - a standalone Mock/Stub server that is able to generate dynamic data](https://www.reddit.com/r/Kotlin/comments/g2hhxd/my_first_kotlin_project_serverneat_a_standalone/)
- url: https://github.com/nomemory/serverneat
---

## [8][A comparison between Kotlin Native approach and native solutions by Marek Multarzyński](https://www.reddit.com/r/Kotlin/comments/g2memo/a_comparison_between_kotlin_native_approach_and/)
- url: https://www.youtube.com/watch?v=u6u_MaQBxAs
---

## [9][Implementing Retrofit in Android | Request methods | Part 6](https://www.reddit.com/r/Kotlin/comments/g2l4z8/implementing_retrofit_in_android_request_methods/)
- url: https://www.youtube.com/watch?v=Bke3kmujg_A
---

## [10][Question about nullable generics](https://www.reddit.com/r/Kotlin/comments/g2e4mh/question_about_nullable_generics/)
- url: https://www.reddit.com/r/Kotlin/comments/g2e4mh/question_about_nullable_generics/
---
Let's say I have a class with a single type parameter, T.

In this class, I can use T? when I want a value of type T to be nullable.

But T could already be a nullable type. 

My question is, how do I do the opposite, i.e. safely represent that a member function that returns type T is NOT nullable even if T happens to be nullable?

Or if that is confusing let me phrase it in another way: given intended construction like MyClass&lt;MyType?&gt;(), how could I write a function inside MyClass that returns a (MyType) instead of a (MyType?)

Thanks for your help
