# Kotlin
## [1][Delegated properties](https://www.reddit.com/r/Kotlin/comments/fqnfgm/delegated_properties/)
- url: https://www.rockandnull.com/delegated-properties-in-kotlin/
---

## [2][Any way to initialize an ArrayList with random data without using a for loop?](https://www.reddit.com/r/Kotlin/comments/fqhgce/any_way_to_initialize_an_arraylist_with_random/)
- url: https://www.reddit.com/r/Kotlin/comments/fqhgce/any_way_to_initialize_an_arraylist_with_random/
---
I need an ArrayList of random data to pass to another function. In python I would just do `[random() for i in range(0, numberOfPoints)]` but the example code I was reading uses a loop with `.add()` and I thought "surely there's a better way?" However in practice I ended up writing the below code (which doesn't even work since ArrayList can't be instantiated with this argument and is harder to read than the loop). Advice?

    var data = ArrayList&lt;PointValue&gt;((0..numberOfPoints).map{
        it -&gt; Math.random().toFloat() * 100f // fill with random data 
    })
## [3][Null Safety](https://www.reddit.com/r/Kotlin/comments/fqp54e/null_safety/)
- url: https://www.reddit.com/r/Kotlin/comments/fqp54e/null_safety/
---
This might be a silly question, but sometimes when working with android and calling methods that are already written using Java in my `.kt` class, intellij will give an error that I should use either the safe calls(?.) or the npe lover operator (!!). Now I understand both, already read the docs:

[https://kotlinlang.org/docs/reference/null-safety.html](https://kotlinlang.org/docs/reference/null-safety.html)

But what I dont understand, how should we know when to use `?.` when calling java code? Now intellij will give me an error which is great! According to what I read if a method contains the annotation `@Nullable` it means it can be null then we have to use `"?."`

&amp;#x200B;

Also my second question, is that sometimes we are forced to use !! instead of ?., why is that?

For example, if I am inside a fragment (inside `onViewCreated()`), and i call `findNavController:`

`val navController : NavController = activity!!.findNavController(R.id.navhostfragment)`

I have to use `!!` instead of `?.` I can do this:

`val navController : NavController? = activity?.findNavController(R.id.navhostfragment)` but then later on I get an error when calling `setupWithNavController()` since it doesnt have an argument of type `NavController?`
## [4][Difference between lateinit and Type declaration ( val variableName : type)](https://www.reddit.com/r/Kotlin/comments/fqfvrw/difference_between_lateinit_and_type_declaration/)
- url: https://www.reddit.com/r/Kotlin/comments/fqfvrw/difference_between_lateinit_and_type_declaration/
---
Hi everyone, I am learning Kotlin in order to be able to program Android apps and I noticed that the lateinit modifier has the same functionality as declaring a variable and its type like this:

var myName : String 

myName = "Tony Stark"


With lateinit it would be:

lateinit var myName : String

myName = "Isaac Asimov"

I tried this on the Kotlin playground web page and both syntaxes returned the value as expected.

So is there a difference between one and the other? 

Thanks in advance to everyone that takes the time to help me figure this out, Please be safe from covid!

Regards
Keyzo
## [5][A good Spring/Reactive project to learn from?](https://www.reddit.com/r/Kotlin/comments/fpxh1u/a_good_springreactive_project_to_learn_from/)
- url: https://www.reddit.com/r/Kotlin/comments/fpxh1u/a_good_springreactive_project_to_learn_from/
---
Last year I got into Kotlin + Spring Reactive project and while I am able to perform, my code starts to seem as overbloated and poorly composed. 

Could somebody please recommend a well written Spring Reactive project to be used as a source of learning? I am intereted in improving the way I handle and compose all the Mono/Flux transformations, errors, etc.
## [6][A smooth refactor using sealed classes and a factory function](https://www.reddit.com/r/Kotlin/comments/fq75u4/a_smooth_refactor_using_sealed_classes_and_a/)
- url: http://le0nidas.gr/2020/03/08/a-smooth-refactor-using-sealed-classes-and-a-factory-function/
---

## [7][Please criticize my code](https://www.reddit.com/r/Kotlin/comments/fpyxgt/please_criticize_my_code/)
- url: https://www.reddit.com/r/Kotlin/comments/fpyxgt/please_criticize_my_code/
---
To enhance my Kotlin skills I've implemented Solitaire - Not the card game.

I couldn't make the launcher work with pure Kotlin so I had to use Java. I used version 14 while developing but others probably work.

The entrance point to the application can be found her.

https://github.com/tonsV2/solitaire/blob/master/src/main/java/dk/fitfit/solitaire/LaunchFX.java

And the readme file can be seen here.

https://github.com/tonsV2/solitaire

I'm well aware that tests are lacking and several parts are pure spaghetti. But please let me know which parts can be written better.
## [8][Introductory Post about the Kotlin/Native Compiler](https://www.reddit.com/r/Kotlin/comments/fpdq7p/introductory_post_about_the_kotlinnative_compiler/)
- url: https://www.bignerdranch.com/blog/exploring-kotlin-native-part-1/
---

## [9][Comments vs. descriptive method names?](https://www.reddit.com/r/Kotlin/comments/fpa4v1/comments_vs_descriptive_method_names/)
- url: https://www.reddit.com/r/Kotlin/comments/fpa4v1/comments_vs_descriptive_method_names/
---
I've recently been playing around with writing a gradle plugin. While doing so I found myself adding some comments.

For the last couple of years I've been arguing that only complicated (or just not obvious) business logic should be commented. If the code needed comments it should probably be refactored. Possible into smaller methods with descriptive names.

What are your thoughts on the matter?

https://github.com/tonsV2/source-release/blob/master/src/main/kotlin/dk/fitfit/gradle/SourceRelease.kt

Which version do you guys prefer?

Edit: Thanks for the feedback everyone! I stripped the comment and kept my code as refactored methods.
## [10][Mocking is not practical — Use fakes](https://www.reddit.com/r/Kotlin/comments/fovn9z/mocking_is_not_practical_use_fakes/)
- url: https://medium.com/@june.pravin/mocking-is-not-practical-use-fakes-e30cc6eaaf4e
---

