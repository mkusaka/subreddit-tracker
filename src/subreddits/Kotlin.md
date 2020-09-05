# Kotlin
## [1][How we built Localazy CLI with Kotlin MPP and Github Actions](https://www.reddit.com/r/Kotlin/comments/in0c1c/how_we_built_localazy_cli_with_kotlin_mpp_and/)
- url: https://www.reddit.com/r/Kotlin/comments/in0c1c/how_we_built_localazy_cli_with_kotlin_mpp_and/
---
Hi, I shared our experience with building a CLI app using Kotlin MPP and automating the entire release with Github Actions.

[https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github](https://localazy.com/blog/localazy-cli-sofware-localization-tool-kotlin-github)

HTTP communication with Kotlin Native on Windows can be a bit tricky, don't forget to check also how to use WinINet API with Kotlin:

[https://localazy.com/blog/kotlin-mpp-and-http-with-wininet-api-on-windows](https://localazy.com/blog/kotlin-mpp-and-http-with-wininet-api-on-windows)
## [2][Are there any good GitHub projects that have implemented the Spring Framework/ Spring Boot using Kotlin? I'm trying to build some small test projects with Spring and Kotlin](https://www.reddit.com/r/Kotlin/comments/imcj2y/are_there_any_good_github_projects_that_have/)
- url: https://www.reddit.com/r/Kotlin/comments/imcj2y/are_there_any_good_github_projects_that_have/
---

## [3][Random numbers](https://www.reddit.com/r/Kotlin/comments/imr1e4/random_numbers/)
- url: https://www.reddit.com/r/Kotlin/comments/imr1e4/random_numbers/
---
I cannot get any sort of random number function to work in Kotlin, for the life of me. It's quite frustrating. Can someone show me an easy way?
## [4][Looking for advice to build a RESTful server written in kotlin on my raspberry pi](https://www.reddit.com/r/Kotlin/comments/imerkm/looking_for_advice_to_build_a_restful_server/)
- url: https://www.reddit.com/r/Kotlin/comments/imerkm/looking_for_advice_to_build_a_restful_server/
---
I have a lot of kotlin experience through Android dev and I absolutely love the language. I have a raspberry pi sitting around and I thought it would be neat as a fun side project to build a local server and app to communicate with each other.

I've been trying to decide on a framework to use but I haven't found any solid documentation and every project seems a bit different. I don't have much knowledge when it comes to web frameworks like spring or ktor so I'm not sure what I need.

Any advice is appreciated, thanks for reading!
## [5][Calculator library (pure Kotlin module)](https://www.reddit.com/r/Kotlin/comments/imegl1/calculator_library_pure_kotlin_module/)
- url: https://www.reddit.com/r/Kotlin/comments/imegl1/calculator_library_pure_kotlin_module/
---
Many apps have number inputs, and in many cases, having a calculator function to provide the number inputs is desirable. This is my first time to publish a library, and hope you will find this helpful. Any feedback is much appreciated. To import:

    repositories {
        maven { url 'https://jitpack.io' }
    }
    
    dependencies {
        implementation 'com.github.jairrab:KotlinCalculator:1.0.0'
    }

To get an instance of the `Calculator` library, call `Calculator.getInstance()` and pass a listener to receive calculator updates.

    class MainActivity : AppCompatActivity(), Calculator.Listener {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
    
            val calculator = Calculator.getInstance(this)
        }
    
        override fun onCalculatorUpdate(key: String?, entries: List&lt;String&gt;, result: Double) {
            Log.v("CALC", "key: $key | entries: $entries | result: $result")
        }
    }

`Calculator.getInstance()` returns a `Calculator` interface that provides the following functions:

    calculator.clear()
    calculator.pressOne()
    calculator.pressTwo()
    calculator.pressThree()
    calculator.pressFour()
    calculator.pressFive()
    calculator.pressSix()
    calculator.pressSeven()
    calculator.pressEight()
    calculator.pressNine()
    calculator.pressZero()
    calculator.pressDecimal()
    calculator.pressPlus()
    calculator.pressMinus()
    calculator.pressMultiply()
    calculator.pressDivide()
    calculator.backSpace()
    calculator.pressEquals()

The `Calculator.Listener` receives updates when a calculator function is called. Here's an example of a sequence of calculator operations and it's resulting logs:

    I/CalculatorLog: Key: initializing | Entries: [] | Result: 0.0
    I/CalculatorLog: Key: 1 | Entries: [1] | Result: 1.0
    I/CalculatorLog: Key: 2 | Entries: [12] | Result: 12.0
    I/CalculatorLog: Key: 3 | Entries: [123] | Result: 123.0
    I/CalculatorLog: Key: backspace | Entries: [12] | Result: 12.0
    I/CalculatorLog: Key: + | Entries: [12, +] | Result: 12.0
    I/CalculatorLog: Key: 1 | Entries: [12, +, 1] | Result: 13.0
    I/CalculatorLog: Key: 3 | Entries: [12, +, 13] | Result: 25.0
    I/CalculatorLog: Key: * | Entries: [12, +, 13, *] | Result: 25.0
    I/CalculatorLog: Key: 2 | Entries: [12, +, 13, *, 2] | Result: 38.0

The calculator is a pure Kotlin module library and does not have any Android dependency, thus you will need to provide a UI to call the and display the calculator functions.  This makes the library lightweight and very flexible to suit your needs. See the link below for a sample UI.

Note: The calculator library currently only supports MDAS operation. That is, it prioritizes multiplication and division over addition and subtraction as shown on the image above.

Link: [https://github.com/jairrab/KotlinCalculator/](https://github.com/jairrab/KotlinCalculator/)
## [6][Gradle modules: Running unit tests only in affected modules](https://www.reddit.com/r/Kotlin/comments/imdgos/gradle_modules_running_unit_tests_only_in/)
- url: https://medium.com/@lehen01/gradle-modules-running-unit-tests-only-in-affected-modules-fff89562339e?source=friends_link&amp;sk=1fe7e4375da04d3626c4f4d6b2b1b34d
---

## [7][Hello DataStore, bye SharedPreferences](https://www.reddit.com/r/Kotlin/comments/imi5zn/hello_datastore_bye_sharedpreferences/)
- url: https://medium.com/scalereal/hello-datastore-bye-sharedpreferences-android-f46c610b81d5
---

## [8][How to use frontend js library with kotlin and ktor](https://www.reddit.com/r/Kotlin/comments/imbeet/how_to_use_frontend_js_library_with_kotlin_and/)
- url: https://www.reddit.com/r/Kotlin/comments/imbeet/how_to_use_frontend_js_library_with_kotlin_and/
---
I can't find it anywhere. I'm trying to build a server side rendered ktor kotlin project and wanted to use a frontend js lib but I can't figure out how to do it.
## [9][Could anyone explain where the 'object' keyword is intended to be used? I found some articles online, but they aren't very intuitive.](https://www.reddit.com/r/Kotlin/comments/ilrm6g/could_anyone_explain_where_the_object_keyword_is/)
- url: https://www.reddit.com/r/Kotlin/comments/ilrm6g/could_anyone_explain_where_the_object_keyword_is/
---

## [10][Can you implement the data class copy() function in kotlin itself, using reflection?](https://www.reddit.com/r/Kotlin/comments/im20w6/can_you_implement_the_data_class_copy_function_in/)
- url: https://www.reddit.com/r/Kotlin/comments/im20w6/can_you_implement_the_data_class_copy_function_in/
---
Is it possible to use kotlin reflection, perhaps in combination with invocation handlers, to implement a generic version of copy(), as it works on data classes? 

Let's say we have a factory, something like:

    object Model {

        inline fun &lt;reified T&gt; init(): T {
            ...
        }
    }

...which given an interface as the type parameter, could give you back a decorated/proxied object that implemented the interface, and provided the copy() function.

We could say that allowable interfaces have to implement some other interface that has a generic definition of the copy function on it, perhaps:

    interface DecoratedInterface {
         fun copy(vararg params: Any) 
    }

However, we don't want to rely on positional parameters, and we don't want to use strings to identify the parameters (i.e. a Map&lt;String, Any&gt; is no bueno).

So, putting it together, what we'd want is to be able to do something like:

    interface Foo: DecoratedInterface {
         val marg: String
         val peanut: Int
         val marmite: Boolean
    } 

    val foo = Model.init&lt;Foo&gt;()  
    // somehow set values on it
    val bah = foo.copy(peanut=8)

Are there language features that let us get at this without rewriting bytecode? Or heck, with rewriting bytecode?
