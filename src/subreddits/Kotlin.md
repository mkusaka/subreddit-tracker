# Kotlin
## [1][KotlinConf 2020 has been postponed until 2021](https://www.reddit.com/r/Kotlin/comments/fluber/kotlinconf_2020_has_been_postponed_until_2021/)
- url: https://kotlinconf.com/?2020Postponed
---

## [2][Unit Testing Delays, Errors &amp; Retries with Kotlin Flows](https://www.reddit.com/r/Kotlin/comments/fll2a7/unit_testing_delays_errors_retries_with_kotlin/)
- url: https://medium.com/@heyitsmohit/unit-testing-delays-errors-retries-with-kotlin-flows-77ce00d0c2f3
---

## [3][How to build Interactive App with Kotlin](https://www.reddit.com/r/Kotlin/comments/flj8ij/how_to_build_interactive_app_with_kotlin/)
- url: https://medium.com/@martinbaraya/how-to-build-interactive-app-with-kotlin-447b81f0d5cb?source=friends_link&amp;sk=45bcd21b90f8f2c12a6f2c6d94a109b9
---

## [4][Kotlin pattern matching with instanceof (the Java 14 way)](https://www.reddit.com/r/Kotlin/comments/fl9ry0/kotlin_pattern_matching_with_instanceof_the_java/)
- url: https://www.reddit.com/r/Kotlin/comments/fl9ry0/kotlin_pattern_matching_with_instanceof_the_java/
---
In Java 14 we are now enabled to use pattern matching with `instanceof` expressions.
e.g. on my code

```
if (psiVariableReferenceContext instanceof PsiExpressionList psiExpressionList &amp;&amp;
    psiExpressionList.getContext() instanceof PsiMethodCallExpression psiMethodCallExpr &amp;&amp;
    !isMethodParameterCharSequence(psiVariableReference, psiMethodCallExpr)) {
  return false;
}
```

Is there a similar Kotlin alternative?  
Currently I need to re-cast the instance and call the getter a second time.

```
if (psiVariableReferenceContext is PsiExpressionList &amp;&amp;
    psiVariableReferenceContext.context is PsiMethodCallExpression &amp;&amp;
    !isMethodParameterCharSequence(
      psiVariableReference,
      psiVariable.context as PsiMethodCallExpression
    )) {
  return false
}
```
## [5][Basic(?) coroutines question for Ktor project](https://www.reddit.com/r/Kotlin/comments/flfxui/basic_coroutines_question_for_ktor_project/)
- url: https://www.reddit.com/r/Kotlin/comments/flfxui/basic_coroutines_question_for_ktor_project/
---
Kind of a basic question but I can't seem to Google it. I'm building an API with Ktor, the main objective being to aggregate and clean up data from multiple sources for an Android app.

When I receive a request to my API I'm launching a request to THIRD_PARTY_API_1 to obtain the results, if THIRD_PARTY_API_1 is down or the results are invalid I try with THIRD_PARTY_API_2 as fallback and if that also fails I try THIRD_PARTY_API_3.

That's easy enough but I thought: I don't really care which THIRD_PARTY_API gives me the data, I should just request to all of them from the start and respond once one of them answers successfully. This way it works *in parallel*. That sounds like something that coroutines can do.

The problem is that [every example](https://ktor.io/clients/index.html#parallel-requests) I find seems to assume that I want to await for every request or I want to cancel the other requests when one of them fails. I want to:

- Launch the three requests.
- Don't cancel the others if one fails or succeeds with invalid data.
- Cancel the others once one is successful with valid data.
- Get that data or an error if noone succeeded to respond from my endpoint.

I usually can get around issues by googling but it seems like I just don't understand asynchronous programming jargon well enough to even look up examples of what I'm trying to do. Any kind of pointers are appreciated &lt;3
## [6][Help With Next Steps](https://www.reddit.com/r/Kotlin/comments/flesct/help_with_next_steps/)
- url: https://www.reddit.com/r/Kotlin/comments/flesct/help_with_next_steps/
---
TLDR: I'm trying to move from beginner programming concepts (in Kotlin) to more advanced things. Kotlin Koans are too much for me. What can I do to bridge the gap?

I started with the codecademy free Java course. That was a great and very simple introduction to programming. I then did the Udemy beginner Kotlin course. I was able to complete the programming challenges and I felt really good about the course. I feel like I have a very simple and basic grasp of object oriented concepts.

I went to do the Kotlin Koans like everyone suggests, but I'm just getting blasted by them. The Kotlin documentation is concept heavy and I struggle to understand what I'm reading. I can't understand what most of the Koans are doing so I can't even try to get the output they're looking for the long way.

I feel like I'm climbing up this mountain and I just arrived at a rock face. I know I need to climb up the rocks to keep advancing but I don't know how to strap on my harness. I can't even begin. I'm just looking for something that can help me bridge the gap from beginner to intermediate.

I appreciate any suggestions. Blog posts, podcasts, books, courses, you name it. I'm just looking for anything that can help me. Thanks in advance!
## [7][How do you work with Network System Discovery? [Question]](https://www.reddit.com/r/Kotlin/comments/fl5yb6/how_do_you_work_with_network_system_discovery/)
- url: https://www.reddit.com/r/Kotlin/comments/fl5yb6/how_do_you_work_with_network_system_discovery/
---
1. I am having a hard time checking whether or not [NSD](https://developer.android.com/training/connect-devices-wirelessly/nsd) is working or not. The resources online are scarce, slack isn't answering, documentation is well made but nothing past the bare minimum. All in all catastrophe... I am basically registering a service, discovering, and then I keep discovering myself on my emulators (I run 2 emulators at the same time). By myself I mean I discover the emulator I registered on and not the other....
2.  Could someone explain a little bit what "service type" and "service name" are. As service type when registering equals to \_http.\_tcp NO DOT, and when listening to a protocol it's with A DOT at the end.  And an example for this with real machine (printer, PC, another phone). 
3. Is there an easier way to work with 2 emulators without having to go into android studio and change the device everytime? Is there a way to split screen logcat and just have both emulators logcats open? (if not where can i ask for this feature?) Or is there another way to test an application that talks to another device. I mean you have to  have 2 emulators open?  


Thank you for your time. I am stuck on this for like 3 or so days. It's been frustrating. The examples online are horrible and just a copy paste of the documentation.
## [8][State-full sequences](https://www.reddit.com/r/Kotlin/comments/fl3nq7/statefull_sequences/)
- url: https://www.reddit.com/r/Kotlin/comments/fl3nq7/statefull_sequences/
---
So here's the behaviour I want:

```Kotlin
val nums = sequence {
    var n = 0
    while (true) { yield(n++) }
}

nums.find { isPrime(it) } // 2
nums.first()              // Want : 3 but I get 0
nums.take(5).toList()     // Want : [4, 5, 6, 7, 8] but I get [0, 1, 2, 3, 4]
```

Any ideas on how I could do something like this?

---

Here's my actual code which essentially require a state-full sequence.

Here's my actual code:
```Kotlin
// See 0 for setup.
while (true) {
    // See 1 for the logic that's implemented below (if sequences had state)
    val data = (blockingQueuesLoopingSequence.take(numQueues) 
        .find { it.peek() != null } ?: blockingQueuesLoopingSequence.first())
        .get()
    // Do stuff with data
}
```
### 0
Here i have a sequence that loops infinitely over a list of `blockingQueues`. 

### 1
Think of a Round-Robin scheduler, current position in the list 
could be any particular queue which would be our state.  
Say we are on the `m^th` queue and we have a total on `N` unique queues.

From all Queues starting from position `m` (our state) looping back to `m-1th` queue.  
Find the first queue that has something available. (Now state is m')

If no such queue choose the first queue from current position 
(which will be `m' = (m + N) mod N = m`)

Call a blocking get on the chosen queue.

Unfortunately, the state(`m`) of `blockingQueuesLoopingSequence` doesn't update and `m is always 0`.
## [9][Equalizer type Progress Indicator](https://www.reddit.com/r/Kotlin/comments/fl6cz3/equalizer_type_progress_indicator/)
- url: /r/androiddev/comments/fko6rs/equalizer_type_progress_indicator/
---

## [10][How can I pass different version of object without altering it as argument?](https://www.reddit.com/r/Kotlin/comments/fl6dom/how_can_i_pass_different_version_of_object/)
- url: https://www.reddit.com/r/Kotlin/comments/fl6dom/how_can_i_pass_different_version_of_object/
---
I have a code like this

val x = MyObject("first", "second")

function1(x)

function2(x)

I want to send different version of x to function 1 but without actually altering it so function2 can have the first version. I have tried let, run, apply and also but couldn't get a solution.
