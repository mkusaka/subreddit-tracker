# Kotlin
## [1][From RxJava to Kotlin Flow: Error Handling](https://www.reddit.com/r/Kotlin/comments/f8lo0l/from_rxjava_to_kotlin_flow_error_handling/)
- url: https://proandroiddev.com/from-rxjava-to-kotlin-flow-error-handling-da1f6a4f2708
---

## [2][Does anyone here have an example of a Kotlin (JVM) project that enforces test coverage?](https://www.reddit.com/r/Kotlin/comments/f8qmgu/does_anyone_here_have_an_example_of_a_kotlin_jvm/)
- url: https://www.reddit.com/r/Kotlin/comments/f8qmgu/does_anyone_here_have_an_example_of_a_kotlin_jvm/
---
I work best off of real examples
## [3][Interaction with user and coroutines](https://www.reddit.com/r/Kotlin/comments/f8bkcw/interaction_with_user_and_coroutines/)
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
## [4][Kotlin Flow Assert Delight](https://www.reddit.com/r/Kotlin/comments/f8d7yw/kotlin_flow_assert_delight/)
- url: https://medium.com/@heyitsmohit/kotlin-flow-assert-ff45465c01c0
---

## [5][Ktor no netty_transport_native_epoll_x86_64 in java.library.path](https://www.reddit.com/r/Kotlin/comments/f8gqtq/ktor_no_netty_transport_native_epoll_x86_64_in/)
- url: https://stackoverflow.com/questions/60367302/ktor-no-netty-transport-native-epoll-x86-64-in-java-library-path
---

## [6][What else can I use kotlin for](https://www.reddit.com/r/Kotlin/comments/f81tag/what_else_can_i_use_kotlin_for/)
- url: https://www.reddit.com/r/Kotlin/comments/f81tag/what_else_can_i_use_kotlin_for/
---
Apart from android development which is the reason I started to learn kotlin what else can I use the language for?
## [7][Which course to take to learn kotlin](https://www.reddit.com/r/Kotlin/comments/f8d258/which_course_to_take_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/f8d258/which_course_to_take_to_learn_kotlin/
---
Hi.I want to learn kotlin to start mobile developing but i don't know which course to take.I saw this one https://www.udacity.com/course/kotlin-bootcamp-for-programmers--ud9011 from google and it seems pretty good but i don't know if there is other better on udemy or something.I am able to pay for it if it is good and i can learn to make mobile apps.Which are your recomandations?
## [8][Any suggestion about how to make my code a bit pretty?](https://www.reddit.com/r/Kotlin/comments/f84989/any_suggestion_about_how_to_make_my_code_a_bit/)
- url: https://www.reddit.com/r/Kotlin/comments/f84989/any_suggestion_about_how_to_make_my_code_a_bit/
---
On the server side I use Micronaut to return a LocalDateTime object and sometimes it's returned without the nano second. Therefor my deserializer contains a, not so pretty, if statement.

        val localDateTimeDeserializer = JsonDeserializer { json, _, _ -&gt;
            if (json.asJsonArray.size() == 7) {
                LocalDateTime.of(
                    json.asJsonArray[0].asInt, // Year
                    json.asJsonArray[1].asInt, // Month
                    json.asJsonArray[2].asInt, // Day of month
                    json.asJsonArray[3].asInt, // Hour
                    json.asJsonArray[4].asInt, // Minute
                    json.asJsonArray[5].asInt, // Second
                    json.asJsonArray[6].asInt // Nano of second
                )
            } else {
                LocalDateTime.of(
                    json.asJsonArray[0].asInt, // Year
                    json.asJsonArray[1].asInt, // Month
                    json.asJsonArray[2].asInt, // Day of month
                    json.asJsonArray[3].asInt, // Hour
                    json.asJsonArray[4].asInt, // Minute
                    json.asJsonArray[5].asInt // Second
                )
            }
        }
## [9][Trying to unit test my app. Rx, Coroutines, Mockk, Junit5. Cannot solve it on my own and no luck on slack nor discord...](https://www.reddit.com/r/Kotlin/comments/f7rqv1/trying_to_unit_test_my_app_rx_coroutines_mockk/)
- url: https://stackoverflow.com/questions/60288406/android-kotlin-unit-test-failing-with-io-mockk-mockkexception-no-answer-found-f
---

## [10][How to make a telnet app with kotlin](https://www.reddit.com/r/Kotlin/comments/f82n6j/how_to_make_a_telnet_app_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/f82n6j/how_to_make_a_telnet_app_with_kotlin/
---
What's the best way to make a telnet app with kotlin?

&amp;#x200B;

I'm a newb in kotlin...

&amp;#x200B;

Thanks :)
