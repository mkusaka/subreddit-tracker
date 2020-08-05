# Kotlin
## [1][ELI5: Nested for-loops](https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/)
- url: https://www.reddit.com/r/Kotlin/comments/i4460x/eli5_nested_forloops/
---
Hi guys, I'm currently learning Kotlin, it's my first programming language - so I'm totally new to this.

Everything so far has made sense, and I've managed to figure it out, but this for-loop is driving me crazy!

`fun main() {`

  `for (i in 1..3) {`

  `for (j in 1..i) {`

`print(j)`

`}`

`}`

`}`

From the IDE, I know the print output is:  112123, but I cannot figure out why!

1..i should be 1, and i is 1..3 which is 123, so shouldn't the output be 1231 or 1123?

Can someone please ELI5 this for me!
## [2][Dokka 1.4.0-rc released](https://www.reddit.com/r/Kotlin/comments/i3i4if/dokka_140rc_released/)
- url: https://blog.jetbrains.com/kotlin/2020/08/dokka-preview-based-on-kotlin-1-4-0-rc/
---

## [3][Best resource for switching from Java to Kotlin - Android dev](https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/i41o2k/best_resource_for_switching_from_java_to_kotlin/
---
Forgive me if this has been asked and answered a million times, I ran a search and only found threads about learning Android/Kotlin from scratch.

I'm a professional Android Dev and have been for years, very familiar with Java, use Python and C a fair bit for work as well. I find I can read Kotlin fine, but figure it's about time I started writing it. Hoping to get a few pointers to some good resources for people who know Android/Java already, and just want to learn Kotlin, not Android or Programming from scratch.
## [4][TIL Exploring Kotin E-Book is free to download from LeanPub](https://www.reddit.com/r/Kotlin/comments/i3eqa9/til_exploring_kotin_ebook_is_free_to_download/)
- url: https://www.reddit.com/r/Kotlin/comments/i3eqa9/til_exploring_kotin_ebook_is_free_to_download/
---
[Exploring Kotlin](https://leanpub.com/exploring-kotlin) is one of the [recommended books](https://kotlinlang.org/docs/books.html) by JetBrains. Today I found out that it is free to download from LeanPub. It is available in PDF, EPUB &amp; MOBI formats to download. 
## [5][Use cases for init blocks](https://www.reddit.com/r/Kotlin/comments/i3x9nn/use_cases_for_init_blocks/)
- url: https://www.reddit.com/r/Kotlin/comments/i3x9nn/use_cases_for_init_blocks/
---
I'm trying to understand the use cases for init blocks. Per the docs

&gt;The primary constructor cannot contain any code. Initialization code can be placed in **initializer blocks**, which are prefixed with the *init* keyword.

but as far as I can tell, you can achieve the same result by doing it in the class body.

I found [this post](https://www.reddit.com/r/Kotlin/comments/hannfg/doubt_what_is_the_function_of_init_code_block_in/) from a month ago and it looks like one of the use cases is to call a custom setter after declaring it. Are there any other use cases?
## [6][Coroutines: How do you create a coroutine that stops when it is started again (only runs when it's the most recent one).](https://www.reddit.com/r/Kotlin/comments/i3m748/coroutines_how_do_you_create_a_coroutine_that/)
- url: https://www.reddit.com/r/Kotlin/comments/i3m748/coroutines_how_do_you_create_a_coroutine_that/
---
I'm calling a method with a coroutine multiple times but I want only the newest one to actually execute and the old one stopping once a new one starts.  


Basically what I have in a non-suspend function:  
scope.launch(Dispatchers.Default){  
list = someLongCall()

}
## [7][TornadoFX - Bind 2D boolean array to grid of buttons](https://www.reddit.com/r/Kotlin/comments/i3py4t/tornadofx_bind_2d_boolean_array_to_grid_of_buttons/)
- url: https://www.reddit.com/r/Kotlin/comments/i3py4t/tornadofx_bind_2d_boolean_array_to_grid_of_buttons/
---
Hi, I have a model class containing a plain 2D array of booleans and I want to be able to bind this array to a grid of buttons on my UI, but I'm not sure how.
## [8][Why aren't my Shared Preferences saving correctly?](https://www.reddit.com/r/Kotlin/comments/i3x2eg/why_arent_my_shared_preferences_saving_correctly/)
- url: https://www.reddit.com/r/Kotlin/comments/i3x2eg/why_arent_my_shared_preferences_saving_correctly/
---
This is my code for saving and changing the values of preferences in a preference activity.

    companion object {
    
            private fun bindPreferenceSummaryToValue(preference: Preference) {
                preference.onPreferenceChangeListener = sBindPreferenceSummaryToValueListener
                val sharedPreferences = PreferenceManager.getDefaultSharedPreferences(preference.context)
    
                if (preference is CheckBoxPreference) {
    
                    sBindPreferenceSummaryToValueListener.onPreferenceChange(
                        preference,
                        sharedPreferences
                            .getBoolean(preference.key, false)
                    )
                    println(sharedPreferences.getBoolean(preference.key, false))
    
                } else if (preference is EditTextPreference){
                    sBindPreferenceSummaryToValueListener.onPreferenceChange(
                        preference,
                        sharedPreferences
                            .getString(preference.key, "")
                    )
                    println(sharedPreferences.getString(preference.key, ""))
    //                editor.putString(preference.key, preference.text)
                }
    
    //            editor.apply()
            }
    
            private val sBindPreferenceSummaryToValueListener = Preference.OnPreferenceChangeListener { preference, newValue -&gt;
                val stringValue = newValue.toString()
    
                val sharedPreferences = PreferenceManager.getDefaultSharedPreferences(preference.context)
                val editor = sharedPreferences.edit()
    
    
                if (preference is ListPreference) {
                    // For list preferences, look up the correct display value in
                    // the preference's 'entries' list.
                    val index = preference.findIndexOfValue(stringValue)
    
                    // Set the summary to reflect the new value.
                    preference.value =
                        if (index &gt;= 0)
                            preference.entries[index] as String?
                        else
                            null
    
                } else if (preference is EditTextPreference) {
                    if (preference.getKey() == "signature") {
                        // update the changed gallery name to summary filed
                        preference.summaryProvider = null
                        preference.text = stringValue
                    }
                } else if (preference is CheckBoxPreference) {
                    preference.isChecked = stringValue == "true"
                    println("STRINGVALUE $stringValue")
                    println("PREFERENCE: ${preference.isChecked}")
                    editor.putBoolean(preference.key, preference.isChecked)
                    editor.apply()
                    println(sharedPreferences.getBoolean(preference.key, false))
                } else if (preference is SwitchPreferenceCompat) {
                    preference.isChecked = stringValue == "true"
                    editor.putBoolean(preference.key, preference.isChecked)
                    editor.apply()
                }
                else {
                    preference.summary = stringValue
                }
                true
            }
        }
    

This is the output from the CheckBoxPreference. 

    I/System.out: STRINGVALUE false
        PREFERENCE: false
        false
    I/System.out: false
    I/System.out: STRINGVALUE true
        PREFERENCE: true
        true

So, at first it is always false, and when i tick it, it is true. I put a boolean in the SharedPreferences, as you can see, and applied it with editor.apply(). But, every time I close and open the app, it is back to false. Why is this happening?
## [9][Validata: Type-safe and highly customizable validation library](https://www.reddit.com/r/Kotlin/comments/i3gqhd/validata_typesafe_and_highly_customizable/)
- url: https://www.reddit.com/r/Kotlin/comments/i3gqhd/validata_typesafe_and_highly_customizable/
---
[https://github.com/njanma/validata](https://github.com/njanma/validata)
## [10][Turbine: A small testing library for kotlinx.coroutines Flow](https://www.reddit.com/r/Kotlin/comments/i34bvg/turbine_a_small_testing_library_for/)
- url: https://github.com/cashapp/turbine
---

