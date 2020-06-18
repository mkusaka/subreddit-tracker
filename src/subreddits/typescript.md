# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Variadic Kinds arriving in TypeScript 4!](https://www.reddit.com/r/typescript/comments/haw15b/variadic_kinds_arriving_in_typescript_4/)
- url: https://github.com/microsoft/TypeScript/issues/5453#issuecomment-644984977
---

## [3][Using a conditional filter, error still occurs: Property 'header' does not exist on type 'never'.ts(2339)](https://www.reddit.com/r/typescript/comments/hbe5ct/using_a_conditional_filter_error_still_occurs/)
- url: https://www.reddit.com/r/typescript/comments/hbe5ct/using_a_conditional_filter_error_still_occurs/
---
    export default abstract class StepsBase {
        protected description: string = ``;
     
        /**** Duck Typed DEFAULT Methods ****/
    
        public explain() {
            // Property 'header' does not exist on type 'never'.ts(2339)
            if ("header" in this) { console.log(this.header); }
    
            if ("description" in this) { console.log(this.description); }
        }

The lint is on the header property of `this.header`, I guess since `this` points back to this class definition.

I thought my if statement would tell the interpreter to only type check inside that control flow block if the condition passed.

I also tried the following type assertions but both failed:

            if ("header" in this) { console.log(this.header as unknown); }
            if ("header" in this) { console.log(this.header as any); }

This is meant to be a base class where some inheriting subclasses define their own `header` and `description` properties. How should the base class be setup to pass type checking? I can leave an empty default value for `header` as I did for `description`, but I would like to know the best practice (maybe it is exactly that).
## [4][Type parameters: What is the name for the part that goes in front of &lt;&gt;?](https://www.reddit.com/r/typescript/comments/hbasvp/type_parameters_what_is_the_name_for_the_part/)
- url: https://www.reddit.com/r/typescript/comments/hbasvp/type_parameters_what_is_the_name_for_the_part/
---
I see the syntax for generics is:

    function doX &lt;T&gt; (args: T): T {...}

And sometimes there is a word in front:

    function doY &lt;T&gt; (args: Array&lt;T&gt;): Partial&lt;T&gt; {...}

I know what these represent. `Array` specifies an array of values of type `T`. `Partial` is a utility type that specific some k/v pairs of the generic are there, but not all are required.

My understanding is not complete. I understand  `Array` and `Partial` are modifiers of a generic type... but how exactly? Are they keywords? Can I make my own? Do they, as a whole, have a deeper syntactical structure (`Partial.If` for example)?

Any help clarifying what that syntax is exactly is appreciated, thanks.
## [5][Help needed with error 'Argument of type 'any' is not assignable to parameter of type 'never'.'](https://www.reddit.com/r/typescript/comments/hb90fk/help_needed_with_error_argument_of_type_any_is/)
- url: https://www.reddit.com/r/typescript/comments/hb90fk/help_needed_with_error_argument_of_type_any_is/
---
Hi! I've been stuck on this error all day: Argument of type 'any' is not assignable to parameter of type 'never'.

 I am using Object.entries(constraints).forEach(([k, v]) =&gt; strConstraint[k as keyof StringValidator](v)) to loop over a constraints object which can contain any combination of the following: 
maxChars?: number; minChars?: number; regexToValidate? RegExp; equals?: string; notEquals?: string;. 

Here is the code for my stringValidator (error on line: 232) https://codesandbox.io/s/festive-bartik-6hdhm?file=/index.ts. If maxChars method (97-103) of number type is commented out, it works fine. 

As a last resort, I could parse a string to a number which works since everything is a string then. Any thoughts?
## [6][Non-hacky ways of doing nested classes?](https://www.reddit.com/r/typescript/comments/hbbif6/nonhacky_ways_of_doing_nested_classes/)
- url: https://www.reddit.com/r/typescript/comments/hbbif6/nonhacky_ways_of_doing_nested_classes/
---
Hi.  I want to create a private class within a class.
```ts
class Foo {
  private class Bar extends Baz {
    …
  }
}
```
But as of TypeScript 3.9.5, it is not possible.  How do I implement that in a non-hacky sane manner?
## [7][Rulr 📐 TypeScript package to save you time writing validation.](https://www.reddit.com/r/typescript/comments/hb1nt6/rulr_typescript_package_to_save_you_time_writing/)
- url: https://github.com/ryansmith94/rulr
---

## [8][A collection of Algorithms and Data Structures with video lectures in Typescript [Update 2]](https://www.reddit.com/r/typescript/comments/haoirx/a_collection_of_algorithms_and_data_structures/)
- url: https://github.com/jeffzh4ng/algorithms-and-data-structures
---

## [9][Understanding Mixins in TypeScript](https://www.reddit.com/r/typescript/comments/hazjt0/understanding_mixins_in_typescript/)
- url: https://blog.bitsrc.io/understanding-mixins-in-typescript-3c2c9a545d87
---

## [10][How To Build A Shopify Headless eCommerce Storefront](https://www.reddit.com/r/typescript/comments/haq4kv/how_to_build_a_shopify_headless_ecommerce/)
- url: http://selleo-shopify.xyz
---

## [11][How to iterate over the properties of an interface/class in order to default initialize an object?](https://www.reddit.com/r/typescript/comments/hav4nv/how_to_iterate_over_the_properties_of_an/)
- url: https://www.reddit.com/r/typescript/comments/hav4nv/how_to_iterate_over_the_properties_of_an/
---
I have interfaces and classes which are rather cumbersome and, for various reasons, require creating "default" forms; for example, the strings need to be initialized to empty strings, the numbers need to be initialized to 0, the arrays need to be initialized to [], etc.

Can anyone point me in the direction of information on how best to accomplish this? I tried looking at the "ts-transformer-keys" package but ran into the same error about "ts_transformer_keys_1.keys is not a function" people have reported from time to time.

Thanks in advance.
