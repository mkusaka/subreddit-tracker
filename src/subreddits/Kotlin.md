# Kotlin
## [1][Coroutines: a practical vocabulary](https://www.reddit.com/r/Kotlin/comments/fmnllf/coroutines_a_practical_vocabulary/)
- url: https://www.rockandnull.com/coroutines-practical/
---

## [2][A template for building Alfred workflows in Kotlin/JS](https://www.reddit.com/r/Kotlin/comments/fmu4hu/a_template_for_building_alfred_workflows_in/)
- url: https://twitter.com/VladyslavSitalo/status/1241584247812448256?s=20
---

## [3][A sample app built with Kotlin, Coroutines, Flow, Dagger 2, Architecture Components, MVVM, Room, Retrofit, Moshi, Material Components and other Jetpack libraries.](https://www.reddit.com/r/Kotlin/comments/fm91pb/a_sample_app_built_with_kotlin_coroutines_flow/)
- url: https://github.com/PatilShreyas/Foodium
---

## [4][My personal library of Kotlin extensions for Android development](https://www.reddit.com/r/Kotlin/comments/fmkwcl/my_personal_library_of_kotlin_extensions_for/)
- url: https://www.reddit.com/r/Kotlin/comments/fmkwcl/my_personal_library_of_kotlin_extensions_for/
---
Hey folks, I'm sharing my personal library that I use in every project of mine, it has things found working from StackOverflow to things I've written personally to take shortcuts, copied some from Github and modified them to work as intended just to make the Android development easier.

[Github Link](https://github.com/CraZyLegenD/Set-Of-Useful-Kotlin-Extensions-and-Helpers)

It lacks documentation at some places, but the methods are quite self explanatory, feel free to contribute or to use it in your projects, pull requests are welcomed as well as if something doesn't work feel free to open an issue or if you find your code add headers/credits that's yours and make a pull request.
## [5][Try make code friends](https://www.reddit.com/r/Kotlin/comments/fmnen3/try_make_code_friends/)
- url: https://www.reddit.com/r/Kotlin/comments/fmnen3/try_make_code_friends/
---
Which is the best way to make friends for talk and code?
## [6][Kotlin multiplatform annotation](https://www.reddit.com/r/Kotlin/comments/fmhlm0/kotlin_multiplatform_annotation/)
- url: https://www.reddit.com/r/Kotlin/comments/fmhlm0/kotlin_multiplatform_annotation/
---
Hello,

&amp;#x200B;

I'm trying to create a shared library of entities (target for Android, JVM and JS).

&amp;#x200B;

Everything goes well until implementing annotations. My goal is to have validator annotations directly written in the entities classes. I want to use JSR 303 (Bean validation) compliant annotation to validate the entities for JVM and Android and write a fallback for Javascript (using the npm class-validator package).

    import validation.Min
    
    class Marketable : Salable() {
    
        @Min(0)
        var stockAmount: Int = 0
    }

My approach is to define a set of "generic" (expected) annotations like this :

    package validation.Min 
    expect annotation class Min

And in JVM implementation

    package validation.Min 
    
    actual typealias Min = javax.validation.constraints.Min
    // ^----- error : Actual class 'Min' 
    // has no corresponding members for expected class members 

What I would like to achieve is to "proxy" the calls to the "actual" annotations because the parameters list is not the same for javax.validation.\* and the annotations from the class-validator npm package.

&amp;#x200B;

In Javascript/Typescript (and even Python), we can wrap decorator into decorator like this (since decorators are simply functions) :

    function MinValidator(value, message, ...) {
        // ... the base Min validator
        return function(target) {}
    }
    
    function MinProxy(value) {
        return MinValidator(value, "Some message", ...)
    }

Thanks
## [7][KotlinConf 2020 has been postponed until 2021](https://www.reddit.com/r/Kotlin/comments/fluber/kotlinconf_2020_has_been_postponed_until_2021/)
- url: https://kotlinconf.com/?2020Postponed
---

## [8][Are there any examples of websites built with kotlin/kotlinJS/react?](https://www.reddit.com/r/Kotlin/comments/fm3qa3/are_there_any_examples_of_websites_built_with/)
- url: https://www.reddit.com/r/Kotlin/comments/fm3qa3/are_there_any_examples_of_websites_built_with/
---
I know that react supports kotlin now.  There are a number of videos showing web applications being built with kotlin-react.  But are there any full websites made with kotlin that you can point me to?

Thanks in advance!
## [9][Help needed with YAML parsing and accessing specific keys(?)](https://www.reddit.com/r/Kotlin/comments/fm2jat/help_needed_with_yaml_parsing_and_accessing/)
- url: https://www.reddit.com/r/Kotlin/comments/fm2jat/help_needed_with_yaml_parsing_and_accessing/
---
Now I'm sure that there's something out there for this, but I haven't found it yet. I need a YAML parsing library that lets me access the keys(?) (I dunno the terminology, here) in a file, using the dot notation (like how we do in Spring Boot).

For example, say I have a YAML file like this . . .

```
some-property: C:\some\folder\path

root:
  foo: sample text
  bar: 1234
  baz: some@random.text
```

. . . and let's say that I wanna access ```"root.foo"``` . . . **is there a way to do so _without_ having to create nested maps or creating DTOs or (basically) data classes?**
## [10][TornadoFX Help](https://www.reddit.com/r/Kotlin/comments/fm0acd/tornadofx_help/)
- url: https://www.reddit.com/r/Kotlin/comments/fm0acd/tornadofx_help/
---
Hey all, I am trying to set up a tornadoFX project. I have openjdk 8, and I need to get javafx in there somewhere, preferably through gradle. Whats the best way to do this? Thanks!
