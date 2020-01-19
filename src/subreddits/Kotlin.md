# Kotlin
## [1][Please give this sub an appropriate title image of the kotlin logo!](https://www.reddit.com/r/Kotlin/comments/eqtoru/please_give_this_sub_an_appropriate_title_image/)
- url: https://www.reddit.com/r/Kotlin/comments/eqtoru/please_give_this_sub_an_appropriate_title_image/
---

## [2][A small hack to add state to interfaces (or, in other words, have Mixins) in Kotlin](https://www.reddit.com/r/Kotlin/comments/equfsz/a_small_hack_to_add_state_to_interfaces_or_in/)
- url: https://discuss.kotlinlang.org/t/you-can-apparently-add-state-to-classes-with-extension-properties-with-almost-no-overhead-a-k-a-mixins-in-kotlin/15312
---

## [3][Kotlin Multiplatform Starter Project for iOS and Android](https://www.reddit.com/r/Kotlin/comments/eqsbkq/kotlin_multiplatform_starter_project_for_ios_and/)
- url: https://www.reddit.com/r/Kotlin/comments/eqsbkq/kotlin_multiplatform_starter_project_for_ios_and/
---
I found myself making many projects again and again for kotlin multiplatform over the last 6 months. I have written plenty of multiplatform project in enterprise settings in past few months, and expect to write many more in the coming days. So, to help folks out there looking to get their hands on the kotlin multiplatform, I made this. Feel free to use as you see fit. Made this mostly for my own necessity initially, but thought it would also help people out there.

[https://github.com/sujitpoudel/kmp-android-ios-starter](https://github.com/sujitpoudel/kmp-android-ios-starter)

EDIT: Just wanted to point out that the project has setup to produce output on maven local, or publish to aws(what I generally do), and for iOS it has a custom gradle task to produce a "Fat Framework", that also does setup for compilation of that framework right from xcode itself.
## [4][What's New in Java 19: The end of Kotlin?](https://www.reddit.com/r/Kotlin/comments/eqgfc3/whats_new_in_java_19_the_end_of_kotlin/)
- url: https://www.youtube.com/watch?v=te3OU9fxC8U
---

## [5][Is it possible in 2020 to learn Kotlin from books/courses if you have zero experience with other languages?](https://www.reddit.com/r/Kotlin/comments/eq7yxn/is_it_possible_in_2020_to_learn_kotlin_from/)
- url: https://www.reddit.com/r/Kotlin/comments/eq7yxn/is_it_possible_in_2020_to_learn_kotlin_from/
---
The title.
The information I found is at least one yeard old.
I have 0 knowledge in programming but usually, I am a fast learner. I can learn foreign languages (not programming) very fast. So I want to try and see how it will go.

It may sound weird but I want to build a couple of apps that are not too complicated like a Tv Guide for my country because there is no good one right now.

Any recommendations are welcome - books, online courses. I can search for myself but there are so many out there and would like to pick something that someone else on my level (0) learned something and was good enough to make them start programming.

Thanks.
## [6][Converting callback functions to ordinary functions](https://www.reddit.com/r/Kotlin/comments/eqevsd/converting_callback_functions_to_ordinary/)
- url: https://www.kotlintips.com/convert-a-callback-function-to-ordinary-function/
---

## [7][Coroutine Dispatcher Selection For JDBC](https://www.reddit.com/r/Kotlin/comments/eq4e62/coroutine_dispatcher_selection_for_jdbc/)
- url: https://www.reddit.com/r/Kotlin/comments/eq4e62/coroutine_dispatcher_selection_for_jdbc/
---
Hey, I need to know on which coroutine dispatcher should I use for querying the data or insert data from the database if my service and MySQL database are on the same machine.

    withContext(whatDispatcher) {   
        database.getData()
    }

Thanks for your time.
## [8][Mutual TLS authentication in REST services with Kotlin](https://www.reddit.com/r/Kotlin/comments/eq32hp/mutual_tls_authentication_in_rest_services_with/)
- url: https://www.reddit.com/r/Kotlin/comments/eq32hp/mutual_tls_authentication_in_rest_services_with/
---
Do you want to develop mutual TLS Web services in Kotlin? Hexagon 1.2 got released today with that feature, check the following section to learn how: https://hexagonkt.com/port_http_server/#https 

You can check all project's examples and source code at GitHub : https://github.com/hexagonkt

If you give it a go, please share your feedback!

Disclaimer: I'm the project's maintainer
## [9][Caching with annotations in android (Kotlin)](https://www.reddit.com/r/Kotlin/comments/epwyeb/caching_with_annotations_in_android_kotlin/)
- url: https://medium.com/@crypticmindscom_5258/caching-made-easy-in-android-with-kotlin-part-2-61bb476063b4
---

## [10][kapt3 stub does not include package name of a generated class annotation parameter](https://www.reddit.com/r/Kotlin/comments/epxmc7/kapt3_stub_does_not_include_package_name_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/epxmc7/kapt3_stub_does_not_include_package_name_of_a/
---
This works:

    @Subcomponent(modules = [shiv.SharedViewModelProviders::class])
    interface ViewModels {
        val fragmentsComponentFactory: Fragments.Factory
    
        @Subcomponent.Factory
        interface Factory {
            fun create(@BindsInstance owner: ViewModelStoreOwner): ViewModels
        }
    }

but when I import the `shiv.SharedViewModelProviders` class and change the first line to

    @Subcomponent(modules = [SharedViewModelProviders::class])

I suddenly get this compile error

    /home/mz/projects/android/FavoritePrimes/app/build/tmp/kapt3/stubs/debug/ph/codeia/favoriteprimes/ViewModels.java:6: error: cannot find symbol
    @dagger.Subcomponent(modules = {SharedViewModelProviders.class})
                                    ^
      symbol: class SharedViewModelProviders

and the stub file looks like

    package ph.codeia.favoriteprimes;
    
    import java.lang.System;
    
    @kotlin.Metadata(/* snip */)
    @dagger.Subcomponent(modules = {SharedViewModelProviders.class})
    public abstract interface ViewModels {
        // [snip]

There's no import for `shiv.SharedViewModelProviders` and the parameter is not fully qualified. This class is generated by an annotation processor I made. What's going on? The import is clearly in the source file. Is there some flag I should enable in my annotation processor to fix this?


**[UPDATE]**

`kapt { correctErrorTypes = true }` in the build script fixes this. Unfortunately, it has to be done on the client side and not in my library. I can only add a note in the docs.

BTW, [this is the library in question](https://github.com/monzee/shiv). It allows you to use constructor injection in fragments and view models using dagger with very little boiler plate (constant wrt the number of fragment and view model classes). Somebody might find it handy.
