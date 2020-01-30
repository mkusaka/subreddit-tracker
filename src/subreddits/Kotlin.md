# Kotlin
## [1][GitHub - VMadalin/kotlin-sample-app: 📚 Android Sample App using modular, clean, scalable, testable Architecture written in Kotlin following the best practices with Jetpack.](https://www.reddit.com/r/Kotlin/comments/ew30x4/github_vmadalinkotlinsampleapp_android_sample_app/)
- url: https://github.com/VMadalin/kotlin-sample-app
---

## [2][Ktor 1.3 Release](https://www.reddit.com/r/Kotlin/comments/evp3qa/ktor_13_release/)
- url: https://blog.jetbrains.com/kotlin/2020/01/ktor-1-3-release/
---

## [3][Is there any way define protected constant without using object in kotlin?](https://www.reddit.com/r/Kotlin/comments/ew2jyk/is_there_any_way_define_protected_constant/)
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
## [4][Deep copy data classes](https://www.reddit.com/r/Kotlin/comments/evnmdx/deep_copy_data_classes/)
- url: https://www.reddit.com/r/Kotlin/comments/evnmdx/deep_copy_data_classes/
---
Is there a way where we can override `copy()` method to copy members which are not declared in constructor.

for eg. 

`data class Dummy(val a: Int , val b: Int) {`

`var list = mutableListOf&lt;String&gt;()`

`}`
## [5][Android Programming vs. Backend Programming with Kotlin](https://www.reddit.com/r/Kotlin/comments/evo7hw/android_programming_vs_backend_programming_with/)
- url: https://www.reddit.com/r/Kotlin/comments/evo7hw/android_programming_vs_backend_programming_with/
---
I only have experiences doing backend programming with Kotlin (API's, Kafka Streams, etc.).

For those of you on both sides, how do they compare and contrast? For example, how is it dealing with UI related code in Kotlin? Does it have some CSS like thing similar to CSS in JS?

&amp;#x200B;

If you've made a transition from one to the other, how did that go? What were the challenges?
## [6][Fragments Tutorial With Example In Android Studio](https://www.reddit.com/r/Kotlin/comments/evnuk4/fragments_tutorial_with_example_in_android_studio/)
- url: https://www.reddit.com/r/Kotlin/comments/evnuk4/fragments_tutorial_with_example_in_android_studio/
---
Learn about the Fragments and how to reuse your fragments in a different layout using Kotlin.

.

[Medium Tutorial Link](https://medium.com/@martinbaraya/fragments-tutorial-with-example-in-android-studio-6f92f53ad8cd?source=friends_link&amp;sk=9f18662fa8b00e746bf2a7beb899dc07)
## [7][KotlinConf 2020 will be September 9-11 in Montreal, Canada](https://www.reddit.com/r/Kotlin/comments/ev5gac/kotlinconf_2020_will_be_september_911_in_montreal/)
- url: https://twitter.com/kotlinconf/status/1222142913280385024?s=20
---

## [8][Arbitrarily ordered parameters](https://www.reddit.com/r/Kotlin/comments/evisnc/arbitrarily_ordered_parameters/)
- url: https://www.reddit.com/r/Kotlin/comments/evisnc/arbitrarily_ordered_parameters/
---
In a recent code review I came across a snippet that bothered me, to the effect of

    first.observe { postValue(calculate(it, second.value)) }
    second.observe { postValue(calculate(first.value, it)) }

`calculate` is a function typed `(A, B) -&gt; C`. In this case, the function requires one each of `A` and `B` in order to generate a `C`, but the order shouldn't matter. I wanted the code to more clearly convey that we are mixing the two data sources, and the forced ordering of the parameters was obscuring that.

My solution, after playing with a few different implementations:

    operator fun &lt;A, B, C&gt; ((A?, B?) -&gt; C).invoke(b: B?, a: A?) = this(a, b)

And that allows me to write the previous block as

    first.observe { postValue(calculate(it, second.value)) }
    second.observe { postValue(calculate(it, first.value)) }

The small change was very satisfying for me, and I think it makes the code easier to understand at a glance.

The solution can also be scaled up to functions with `n` parameters, but will require `nPn - 1` function declarations, ie `(A, B, C) -&gt; D` would require 5 functions, while `(A, B, C, D) -&gt; E` would require 23 functions.

Anyway, I thought this was a neat snippet, so I thought I'd share.
## [9][Help Removing Recursive Generics for Kotlin Compatibility](https://www.reddit.com/r/Kotlin/comments/evi8h2/help_removing_recursive_generics_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/evi8h2/help_removing_recursive_generics_for_kotlin/
---
In a previous discussion the topic of how Kotlin doesn't support recursive generics came up. I'm working on adding better Kotlin support for my Java computer vision library. A significant pain point is that It uses recursive Generics in the image class as a way to keep strong typing. Below is minimalist example showing the problem. Any ideas on how to do this without recursive generics?

Example of the real code on [Github](https://github.com/lessthanoptimal/BoofCV/blob/SNAPSHOT/main/boofcv-types/src/main/java/boofcv/struct/image/GrayU8.java)

    //-----------------------------------------------------------------------------------------------
    // With recursive generics
    // - No warnings and no need for explicit type-cast in newInstance()
    // - Maintains strong typing
    
    public interface BaseA&lt;T extends BaseA&lt;T&gt;&gt; {
        T newInstance();
    }
    
    public static &lt;T extends BaseA&lt;T&gt;&gt; T functionA( T a ) {
        return a.newInstance();
    }
    
    //-----------------------------------------------------------------------------------------------
    // Without recursion there's the loss of type information
    //
    // 1) Raw use of parameterized class warnings
    // 2) Have to type-cast in functionB
    
    interface BaseB&lt;T extends BaseB&gt; {
        T newInstance();
    }
    
    public static &lt;T extends BaseB&gt; T functionB( T a ) {
        return (T)a.newInstance();
    }
## [10][Smart casting doesn't work when non-nullability is given indirectly](https://www.reddit.com/r/Kotlin/comments/ev3cl2/smart_casting_doesnt_work_when_nonnullability_is/)
- url: https://i.redd.it/s5s73a28mhd41.png
---

