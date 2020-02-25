# Kotlin
## [1][What are the best selling points for Kotlin over Java?](https://www.reddit.com/r/Kotlin/comments/f9960k/what_are_the_best_selling_points_for_kotlin_over/)
- url: https://www.reddit.com/r/Kotlin/comments/f9960k/what_are_the_best_selling_points_for_kotlin_over/
---
Working with a group of researchers and they mainly know Java. What are the biggest advantages that Kotlin has over Java ?
## [2][Questions about fkorotkov's Kubernetes DSL](https://www.reddit.com/r/Kotlin/comments/f99nrd/questions_about_fkorotkovs_kubernetes_dsl/)
- url: https://www.reddit.com/r/Kotlin/comments/f99nrd/questions_about_fkorotkovs_kubernetes_dsl/
---
Hey guys,

At my job we are about to use Helm charts to configure more easily our k8s cluster, but somewhere I've seen there's this Kotlin DSL for Kubernetes and I was looking at it and seems pretty good to be honest, specially by the fact I want to propose the usage of Kotlin for other aspects in the project.My questions are:

1. Is there any project example that uses this dsl besides of the ones that are documented in the repository?
2. Under the hood, how does this DSL connect with the kubernetes cluster? Is there any configuration that has to be done beforehand 
3. Or does this generate the yml configurations ready to use in kubernetes?

  
Thanks in advance people.
## [3][How to build a GraphQL with Authentication](https://www.reddit.com/r/Kotlin/comments/f92u2u/how_to_build_a_graphql_with_authentication/)
- url: https://youtu.be/0Vs_AfJltLI
---

## [4][Does anyone here have an example of a Kotlin (JVM) project that enforces test coverage?](https://www.reddit.com/r/Kotlin/comments/f8qmgu/does_anyone_here_have_an_example_of_a_kotlin_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/f8qmgu/does_anyone_here_have_an_example_of_a_kotlin_jvm/
---
I work best off of real examples
## [5][Enable Android ViewModel to pass in custom parameters with the new Saved State with Android Lifecycle 2.2.0](https://www.reddit.com/r/Kotlin/comments/f8tgkl/enable_android_viewmodel_to_pass_in_custom/)
- url: https://proandroiddev.com/optimizing-viewmodel-with-lifecycle-2-2-0-a2895b5c01fd
---

## [6][What about kotlin documentation](https://www.reddit.com/r/Kotlin/comments/f8sa1y/what_about_kotlin_documentation/)
- url: https://www.reddit.com/r/Kotlin/comments/f8sa1y/what_about_kotlin_documentation/
---
Hi.I want to learn kotlin and i have background in python and c++(i know only the basics in both).Do you guys think that the official kotlin documentation is good to learn kotlin?or should i take other online course?I want to be an android developer btw.
## [7][From RxJava to Kotlin Flow: Error Handling](https://www.reddit.com/r/Kotlin/comments/f8lo0l/from_rxjava_to_kotlin_flow_error_handling/)
- url: https://proandroiddev.com/from-rxjava-to-kotlin-flow-error-handling-da1f6a4f2708
---

## [8][Interaction with user and coroutines](https://www.reddit.com/r/Kotlin/comments/f8bkcw/interaction_with_user_and_coroutines/)
- url: https://www.reddit.com/r/Kotlin/comments/f8bkcw/interaction_with_user_and_coroutines/
---
Hi!

I'm starting to use Kotlin coroutines within my project's codebase, and they're so easy to use and make everything so much simpler, that I'm starting to have a tendency to use them for everything asynchronous, even for interactions with user.

For example, at some point the app needs to ask the user if they would really like to perform a dangerous action, and displays a dialog to them. The classic approach would be to show a dialog and subscribe to its buttons, but with coroutines there's another (very seductive) approach:

    suspend fun askUserToMoveAccount():
            TwoOptionsDialog.UserChoice = suspendCoroutine { continuation -&gt;
        // NOTE: it's ok to lose the dialog state on Activity recreation
        // and configuration changes
        val dialog = TwoOptionsDialog.showDialog(
                activity.supportFragmentManager, TAG,
                R.string.gp_account_move_request,
                R.string.gp_account_move_request_confirmation,
                R.string.gp_account_move_request_cancellation)
    
        dialog.setOnUserChoiceListener {
            dialog.dismiss()
            continuation.resume(it)
        }
    }
    
    ...
    
    val userChoice = askUserToMoveAccount()
    if (userChoice == TwoOptionsDialog.UserChoice.NEGATIVE) {
        return ObtainResult.CanceledByUser
    }

I.e. I move the dialog creation into a suspend function, which returns the user choice, and code, which calls the function, looks like it's synchronous.

Is it ok to do such things?

Coroutines are advertised and explained in the docs as a nice way to write concurrent code (e.g. working with DB or network) - I think I never saw coroutines usages similar to the code above.
## [9][Kotlin Flow Assert Delight](https://www.reddit.com/r/Kotlin/comments/f8d7yw/kotlin_flow_assert_delight/)
- url: https://medium.com/@heyitsmohit/kotlin-flow-assert-ff45465c01c0
---

## [10][Ktor no netty_transport_native_epoll_x86_64 in java.library.path](https://www.reddit.com/r/Kotlin/comments/f8gqtq/ktor_no_netty_transport_native_epoll_x86_64_in/)
- url: https://stackoverflow.com/questions/60367302/ktor-no-netty-transport-native-epoll-x86-64-in-java-library-path
---

