# Kotlin
## [1][What I Learnt from Benchmarking Http4k, Ktor (Kotlin) and Actix v2, v3 (Rust) Microservices](https://www.reddit.com/r/Kotlin/comments/isgvm6/what_i_learnt_from_benchmarking_http4k_ktor/)
- url: https://matej.laitl.cz/bench-rust-kotlin-microservices/
---

## [2][Talking Kotlin - Fritz2](https://www.reddit.com/r/Kotlin/comments/iskdpk/talking_kotlin_fritz2/)
- url: https://talkingkotlin.com/fritz2/
---

## [3][Kotlin Multiplatform: The Secret to Kotlin's Rising Popularity](https://www.reddit.com/r/Kotlin/comments/isico9/kotlin_multiplatform_the_secret_to_kotlins_rising/)
- url: https://blog.codota.com/kotlin-multiplatform/
---

## [4][Looking for Kotlin micro-service with excellent testing and DevOps practices](https://www.reddit.com/r/Kotlin/comments/ise78l/looking_for_kotlin_microservice_with_excellent/)
- url: https://www.reddit.com/r/Kotlin/comments/ise78l/looking_for_kotlin_microservice_with_excellent/
---
I'm looking for an open source micro-service, written in Kotlin, that has excellent testing practices to get inspiration from: unit, integration, and e2e tests with CI/CD. On AWS would be the cherry on top.

If you know of one, please share!
## [5][Customizing your GitHub profile - Scripting](https://www.reddit.com/r/Kotlin/comments/is155t/customizing_your_github_profile_scripting/)
- url: https://blog.frankel.ch/customizing-github-profile/1/
---

## [6][What is the fastest way Combination in Kotlin](https://www.reddit.com/r/Kotlin/comments/isg16h/what_is_the_fastest_way_combination_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/isg16h/what_is_the_fastest_way_combination_in_kotlin/
---
As far as I know, Kotlin doesn't support Combination library.  


So what is the way to make the fastest element Combination in List?

I think recursive solution is not fast,,
## [7][Primitive obsession](https://www.reddit.com/r/Kotlin/comments/isf8tj/primitive_obsession/)
- url: https://www.reddit.com/r/Kotlin/comments/isf8tj/primitive_obsession/
---
Is it good to stick to primitives in kotlin ? 
Or will it cost me too much to wrap some of my code in a class and play with types ? 
I watched this video but again the discussion on this says that it’s always not the best thing to do. 

https://m.youtube.com/watch?v=t3DBzaeid74

Confused. 
Any clear suggestion
## [8][Validating function return type](https://www.reddit.com/r/Kotlin/comments/isah5s/validating_function_return_type/)
- url: https://www.reddit.com/r/Kotlin/comments/isah5s/validating_function_return_type/
---
I am converting some Kotlin code to work in KMP and apparently there are some limitations with reflection there, plus I always wondered if there was a more straightforward way to do this since reflection is not always the preferred solution.  Using the following simplified code is there a way to check the eval() instance of 2 classes will return the same type?  

&amp;#x200B;

abstract class Function {

abstract fun eval(): Number

}

&amp;#x200B;

class Function1 : Function() {

override fun eval(): Int = 55

}

&amp;#x200B;

class Function2 : Function() {

override fun eval(): Double = 5.5

}

&amp;#x200B;

fun foo(a: Function, b: Function) {

// TODO without calling eval(), check if a.eval() returns the same type as b.eval()

}
## [9][How to run process in background?](https://www.reddit.com/r/Kotlin/comments/is19pk/how_to_run_process_in_background/)
- url: https://www.reddit.com/r/Kotlin/comments/is19pk/how_to_run_process_in_background/
---
So i need to validate username/password and then redirect to another screen. Validation needs to be done in the background process while, and redirection on main thread. How should i do that? Using coroutines?
## [10][http4k graalvm native-image example.](https://www.reddit.com/r/Kotlin/comments/irm41u/http4k_graalvm_nativeimage_example/)
- url: https://github.com/http4k/examples/tree/master/graalvm
---

