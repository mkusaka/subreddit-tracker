# Kotlin
## [1][Any suggestions for Kotlin certification?](https://www.reddit.com/r/Kotlin/comments/ey5pla/any_suggestions_for_kotlin_certification/)
- url: https://www.reddit.com/r/Kotlin/comments/ey5pla/any_suggestions_for_kotlin_certification/
---
Hello,

I’m a junior dev and I’m working on the server side with kotlin. I don’t have a car degree and I wanted a certification to establish myself as a dev. I’m based in the UK and everything that I found (47degrees and jetBrains) is located in the states. 

Is there any certification that I could read for/take classes and then take an exam? 

Thank you
## [2][Why everything has to return something in Kotlin?](https://www.reddit.com/r/Kotlin/comments/ey5i2s/why_everything_has_to_return_something_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ey5i2s/why_everything_has_to_return_something_in_kotlin/
---
When doing the codelabs from the Kotlin Bootcamp for Programmers, I stumbled across the topic in lesson 3 where they talk about how everything in Kotlin returns something. 

I tested it and in fact you can nest the println() statements in Kotlin whereas JAVA just doesn't let you do that. However the output is a bit confusing to interpret for me. 

I ran this code : 

`println({print(print("This!"))})`

and got this as output : 

`() -&gt; kotlin.Unit`

and for statement :

`println(print(print("This!")))`

I got this as output : 

`This!kotlin.Unitkotlin.Unit`

I am trying to dig deeper to interpret the output (just started learning Kotlin).

I am curious as to if there are hazards, if any, to this approach, or is it just that Kotlin realized what JAVA missed out on.
## [3][Loading and caching images into ImageViews using one annotation](https://www.reddit.com/r/Kotlin/comments/ey6idr/loading_and_caching_images_into_imageviews_using/)
- url: https://www.reddit.com/r/Kotlin/comments/ey6idr/loading_and_caching_images_into_imageviews_using/
---
I created an annotation which will load image from a URL directly into an ImageView (caching it for future use) . It can also be used to show a placeholder image and animation until the image is downloaded into the device.

Usage :-

&gt;"@LoadImage(  
R.id.*image\_1*,  
 "https://images.unsplash.com/photo-1549740425-5e9ed4d8cd34?ixlib=rb-1.2.1&amp;w=1000&amp;q=80",  
 placeHolder = R.drawable.*loading*, enableLoadingAnimation = true  
)  
lateinit var imageWithAnimation: ImageView"

Check out the post for more details :-  [https://medium.com/@crypticmindscom\_5258/caching-made-easy-on-android-with-kotlin-part-4-18e7b066e9c2](https://medium.com/@crypticmindscom_5258/caching-made-easy-on-android-with-kotlin-part-4-18e7b066e9c2) 

Library :-  [https://github.com/crypticminds/ColdStorage](https://github.com/crypticminds/ColdStorage)
## [4][Getting started with web development (Tips &amp; Recommendations?)](https://www.reddit.com/r/Kotlin/comments/ey6du2/getting_started_with_web_development_tips/)
- url: https://www.reddit.com/r/Kotlin/comments/ey6du2/getting_started_with_web_development_tips/
---
Hey there,

I'm looking into building a system in Kotlin in which I have a server containing some user-bound data (sent from a mobile client that can be omitted from this discussion), and a web client to retrieve and visualize said data. In upcoming iterations I want users of the web client to be able to add information to the server as well.

So far I've build a [server according to this example](https://github.com/gimlet2/kottpd), and added GSON for parsing JSON-Objects. So far, things work as expected: the server receives JSON-Objects and parses them to Koktlin objects (and vice versa). The server can display a very basic website containt a simple form to input a User ID, and upon submitting said form displays the chosen user data in text form.

I'm now looking for ways to get started with proper development of my web frontent. Unfortunately, while having some theoretical understand of the subject, I have little to no practical experience in this (other than working with JSPs, Servelts and JavaBeans once 6 years ago). So I'm looking for some guidance on what I need to get started.

I'm specifically trying to avoid build tools such as Maven, Gradle, etc., so I would ask you to kindly respect that.  
I'm also trying to keep the use of libraries to a minimum. I understand that I can hardly go without them, but I'm generally trying to keep as independant of third party software as possible.
## [5][Best resource for learning Kotlin for android dev.](https://www.reddit.com/r/Kotlin/comments/exvxwl/best_resource_for_learning_kotlin_for_android_dev/)
- url: https://www.reddit.com/r/Kotlin/comments/exvxwl/best_resource_for_learning_kotlin_for_android_dev/
---
Hello, I want to build an android app and this is why I choose to learn kotlin.

It has been some time with kotlin now and I think in order to develop my skills I should just straight into trying to build some apps in android studio.

Ideally I would like to follow some online tutorial as reference to get me started.

There doesn't seem to be much content online especially tutorials that actually walk you through a whole build, it seems to just be one off tutorials for very specific parts.

Can anyone suggest something that will help guide me on this journey?

Thank you
## [6][Choosing the right scope function](https://www.reddit.com/r/Kotlin/comments/ey5u7i/choosing_the_right_scope_function/)
- url: https://www.atomiccommits.io/scope-functions/
---

## [7][React in Kotlin/JS: What I learned (long but useful read)](https://www.reddit.com/r/Kotlin/comments/exslpk/react_in_kotlinjs_what_i_learned_long_but_useful/)
- url: https://discuss.kotlinlang.org/t/react-in-kotlin-js-what-i-learned-long-but-useful-read/16168
---

## [8][What have been the main objections against Kotlin transition from Java amongst your colleagues?](https://www.reddit.com/r/Kotlin/comments/exuo7x/what_have_been_the_main_objections_against_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/exuo7x/what_have_been_the_main_objections_against_kotlin/
---

## [9][JSON to Kotlin data class](https://www.reddit.com/r/Kotlin/comments/exmp2s/json_to_kotlin_data_class/)
- url: https://www.rockandnull.com/json-to-kotlin-data-class/
---

## [10][How to read (Kotlin) documentation properly?](https://www.reddit.com/r/Kotlin/comments/ex8uz4/how_to_read_kotlin_documentation_properly/)
- url: https://www.reddit.com/r/Kotlin/comments/ex8uz4/how_to_read_kotlin_documentation_properly/
---
Hi I need some help on how to understand the kotlin documentation, although I guess this is a problem for me in other languages also. 

Say I want to learn about lambdas, I go here: 

 [https://kotlinlang.org/docs/reference/lambdas.html](https://kotlinlang.org/docs/reference/lambdas.html) 

and begin reading.

I don't really understand but read up until I reach the code:

`fun &lt;T, R&gt; Collection&lt;T&gt;.fold(`

`initial: R,` 

`combine: (acc: R, nextElement: T) -&gt; R`

`): R {`

`var accumulator: R = initial`

`for (element: T in this) {`

`accumulator = combine(accumulator, element)`

`}`

`return accumulator`

`}`

&amp;#x200B;

Now I am really lost.

I don't know what these mean &lt;&gt;, they normally come with something inside, I have seen ArrayList&lt;String&gt; for example in java so maybe it means type?

I don't know what &lt;T,R&gt; means, I don't know what Collection&lt;T&gt;.fold is and so on.

I feel this is a common problem for me with documentation, I only can understand the documentation when I understand the subject it is talking about, hence rendering the documentation a lot less useful.

Could somebody, explain in detail how to understand something like this and how I should approach the documentation in order to get the best comprehension?

Thank you,
