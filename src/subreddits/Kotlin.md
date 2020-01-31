# Kotlin
## [1][Regex matcher using TCR with Jordan Stewart](https://www.reddit.com/r/Kotlin/comments/ewo85o/regex_matcher_using_tcr_with_jordan_stewart/)
- url: https://www.youtube.com/watch?v=QEd2anW86YQ
---

## [2][What is the future of Kotlin?](https://www.reddit.com/r/Kotlin/comments/ewacns/what_is_the_future_of_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ewacns/what_is_the_future_of_kotlin/
---
We know it's the first language in Android development. What do you think is the future in mobile and backend development? Is it possible to be a developer sought in business areas with Kotlin?
## [3][GitHub - VMadalin/kotlin-sample-app: 📚 Android Sample App using modular, clean, scalable, testable Architecture written in Kotlin following the best practices with Jetpack.](https://www.reddit.com/r/Kotlin/comments/ew30x4/github_vmadalinkotlinsampleapp_android_sample_app/)
- url: https://github.com/VMadalin/kotlin-sample-app
---

## [4][As a JavaScript developer should I learn Java before Kotlin?](https://www.reddit.com/r/Kotlin/comments/ewc6k5/as_a_javascript_developer_should_i_learn_java/)
- url: https://www.reddit.com/r/Kotlin/comments/ewc6k5/as_a_javascript_developer_should_i_learn_java/
---
Im already familiar with JavaScript but I'm looking for a way to properly learn OOP. I'm not doing this to get a new job or for any particular project, I just want to improve my understanding of programming in general. 

Some people tell me I should learn Java but others say it's too verbose and Kotlin is better. I have limited time outside of my job so this could be a selling point. Is Kotlin easier to learn? Would I miss out on any fundamentals of OOP if I skipped Java?
## [5][Convert your regular class to a cache layer using just one annotation (android)](https://www.reddit.com/r/Kotlin/comments/ewbg78/convert_your_regular_class_to_a_cache_layer_using/)
- url: https://www.reddit.com/r/Kotlin/comments/ewbg78/convert_your_regular_class_to_a_cache_layer_using/
---
 [https://medium.com/@crypticmindscom\_5258/caching-made-easy-on-android-with-kotlin-part-3-3d4cfcb57df0](https://medium.com/@crypticmindscom_5258/caching-made-easy-on-android-with-kotlin-part-3-3d4cfcb57df0) 

&amp;#x200B;

Check out the repository here:-  [https://github.com/crypticminds/ColdStorage](https://github.com/crypticminds/ColdStorage)
## [6][Ktor 1.3 Release](https://www.reddit.com/r/Kotlin/comments/evp3qa/ktor_13_release/)
- url: https://blog.jetbrains.com/kotlin/2020/01/ktor-1-3-release/
---

## [7][Is there any way define protected constant without using object in kotlin?](https://www.reddit.com/r/Kotlin/comments/ew2jyk/is_there_any_way_define_protected_constant/)
- url: https://www.reddit.com/r/Kotlin/comments/ew2jyk/is_there_any_way_define_protected_constant/
---
Hi,

i want to define protected constant with Kotlin but i am stuck. Using of companion object and object seems like over kill to me. Unnecessary object creation , getter and setter. I know I can use const or @ JvmField for get rid of getter and setter but Object creation still is there.

When I put constant value to top level, I can"t use protected visibility and Kotlin generate another public final class for just one constant. So how can i define protected constant? Is there any optimization for companion object when i use with  const or @ JvmField because when i inspect bytecode just see unused companion object. Will the jit optimize companion object?

    public final class Something {
       private static final String constantValue = "BLABLA";
       public static final Something.Companion Companion = new Something.Companion((DefaultConstructorMarker)null);
    
       @Metadata(
          mv = {1, 1, 16},
          bv = {1, 0, 3},
          k = 1,
          d1 = {"\u0000\u0012\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0010\u000e\n\u0000\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002¢\u0006\u0002\u0010\u0002R\u000e\u0010\u0003\u001a\u00020\u0004X\u0082T¢\u0006\u0002\n\u0000¨\u0006\u0005"},
          d2 = {"Lcom/imobilecode/ekotaksi/Something$Companion;", "", "()V", "constantValue", "", "app"}
       )
       public static final class Companion {
          private Companion() {
          }
    
          // $FF: synthetic method
          public Companion(DefaultConstructorMarker $constructor_marker) {
             this();
          }
       }
    }
## [8][Deep copy data classes](https://www.reddit.com/r/Kotlin/comments/evnmdx/deep_copy_data_classes/)
- url: https://www.reddit.com/r/Kotlin/comments/evnmdx/deep_copy_data_classes/
---
Is there a way where we can override `copy()` method to copy members which are not declared in constructor.

for eg. 

`data class Dummy(val a: Int , val b: Int) {`

`var list = mutableListOf&lt;String&gt;()`

`}`
## [9][Android Programming vs. Backend Programming with Kotlin](https://www.reddit.com/r/Kotlin/comments/evo7hw/android_programming_vs_backend_programming_with/)
- url: https://www.reddit.com/r/Kotlin/comments/evo7hw/android_programming_vs_backend_programming_with/
---
I only have experiences doing backend programming with Kotlin (API's, Kafka Streams, etc.).

For those of you on both sides, how do they compare and contrast? For example, how is it dealing with UI related code in Kotlin? Does it have some CSS like thing similar to CSS in JS?

&amp;#x200B;

If you've made a transition from one to the other, how did that go? What were the challenges?
## [10][Fragments Tutorial With Example In Android Studio](https://www.reddit.com/r/Kotlin/comments/evnuk4/fragments_tutorial_with_example_in_android_studio/)
- url: https://www.reddit.com/r/Kotlin/comments/evnuk4/fragments_tutorial_with_example_in_android_studio/
---
Learn about the Fragments and how to reuse your fragments in a different layout using Kotlin.

.

[Medium Tutorial Link](https://medium.com/@martinbaraya/fragments-tutorial-with-example-in-android-studio-6f92f53ad8cd?source=friends_link&amp;sk=9f18662fa8b00e746bf2a7beb899dc07)
