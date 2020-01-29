# Kotlin
## [1][KotlinConf 2020 will be September 9-11 in Montreal, Canada](https://www.reddit.com/r/Kotlin/comments/ev5gac/kotlinconf_2020_will_be_september_911_in_montreal/)
- url: https://twitter.com/kotlinconf/status/1222142913280385024?s=20
---

## [2][Arbitrarily ordered parameters](https://www.reddit.com/r/Kotlin/comments/evisnc/arbitrarily_ordered_parameters/)
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
## [3][Help Removing Recursive Generics for Kotlin Compatibility](https://www.reddit.com/r/Kotlin/comments/evi8h2/help_removing_recursive_generics_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/evi8h2/help_removing_recursive_generics_for_kotlin/
---
In a previous discussion the topic of how Kotlin doesn't support recursive generics came up. My computer vision library uses them in the image class as a way to keep strong typing when working with generics. Below is minimalist example showing the problem. Any ideas on how to do this without recursive generics?

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
## [4][Smart casting doesn't work when non-nullability is given indirectly](https://www.reddit.com/r/Kotlin/comments/ev3cl2/smart_casting_doesnt_work_when_nonnullability_is/)
- url: https://i.redd.it/s5s73a28mhd41.png
---

## [5][Java Streams vs. Kotlin Sequences](https://www.reddit.com/r/Kotlin/comments/ev5aqx/java_streams_vs_kotlin_sequences/)
- url: https://www.reddit.com/r/Kotlin/comments/ev5aqx/java_streams_vs_kotlin_sequences/
---
This article explores the pros and cons of Streams vs. Sequences since this has been a common question among Kotlin developers:  

[https://proandroiddev.com/java-streams-vs-kotlin-sequences-c9ae080abfdc](https://proandroiddev.com/java-streams-vs-kotlin-sequences-c9ae080abfdc) 

I also included a few surprises so I hope you find it interesting.
## [6][What do i need to learn if i want to start developing android apps in kotlin rather than java ?](https://www.reddit.com/r/Kotlin/comments/evdpa2/what_do_i_need_to_learn_if_i_want_to_start/)
- url: https://www.reddit.com/r/Kotlin/comments/evdpa2/what_do_i_need_to_learn_if_i_want_to_start/
---
I have basic knowledge of the syntax (i know swift).
I have very good knowledge of java.
## [7][Underrated blog post explaining -Xjvm-default](https://www.reddit.com/r/Kotlin/comments/ev5wzr/underrated_blog_post_explaining_xjvmdefault/)
- url: https://www.realjenius.com/2018/06/29/jvm-default/
---

## [8][I've created a Kotlin JWT (Json Web Token) library specifically designed for APN (Apple Push Notifications) and Sign in with Apple - feedback appreciated! 👍](https://www.reddit.com/r/Kotlin/comments/ev5rgy/ive_created_a_kotlin_jwt_json_web_token_library/)
- url: https://github.com/PhilJay/JWT
---

## [9][Extending default spring data repository methods](https://www.reddit.com/r/Kotlin/comments/ev4ogw/extending_default_spring_data_repository_methods/)
- url: https://www.reddit.com/r/Kotlin/comments/ev4ogw/extending_default_spring_data_repository_methods/
---
Hey guys,

I recently published a medium related to how can you extend your reactive default spring data repository methods. I hope it helps you not to face the same problem that I did. I was struggling a lot to figure out how can I achieve this and finally; I found it. Here is the link:

[https://medium.com/@ghahremani/extending-default-spring-data-repository-methods-patch-example-a23c07c35bf9](https://medium.com/@ghahremani/extending-default-spring-data-repository-methods-patch-example-a23c07c35bf9)

Let me know your thoughts.
## [10][When to use inline functions](https://www.reddit.com/r/Kotlin/comments/euoxfb/when_to_use_inline_functions/)
- url: https://www.atomiccommits.io/inline-functions/
---

