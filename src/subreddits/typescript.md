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
## [2][Typing promisify with variadic tuple types](https://www.reddit.com/r/typescript/comments/hbpbwq/typing_promisify_with_variadic_tuple_types/)
- url: https://oida.dev/variadic-tuple-types-preview/
---

## [3][Why You Should Use TypeScript in 2020](https://www.reddit.com/r/typescript/comments/hbhfsh/why_you_should_use_typescript_in_2020/)
- url: https://serokell.io/blog/why-typescript
---

## [4][[Help Wanted] Mutation Testing with Typescript](https://www.reddit.com/r/typescript/comments/hbqgfg/help_wanted_mutation_testing_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hbqgfg/help_wanted_mutation_testing_with_typescript/
---
Supposedly Stryker as it working for nearly 3 years: https://stryker-mutator.io/blog/2017-10-06/typescript-support

There's even a [short guide](https://github.com/stryker-mutator/stryker-handbook/blob/master/stryker/guides/vuejs.md#jest-configuration) suitable for [my small project](https://gitlab.com/bss03/vue-webpack-ts-lambda).

Unfortunately I'm experiencing a failure before even the first test: https://pastebin.com/WWdPEXf4

I commented on a closed [issue](https://github.com/stryker-mutator/stryker/issues/1229#issuecomment-645721873), but it hasn't been reopened (nor has anyone commented on it since I did).

Anyone got mutation testing (Skryker or otherwise) working with typescript that could point me the right direction?  It looks like I might need to tell stryker to do some transpiling, but I'm not sure.  I think I tried `"transpiler": [ "typescript" ]` in the Stryker configuration, but it didn't help.

Thanks for reading, and thanks in advance for any help you can provide.
## [5][Those of you who came from dynamic languages, how did you survive without compile-time checks and autocomplete?](https://www.reddit.com/r/typescript/comments/hbhajh/those_of_you_who_came_from_dynamic_languages_how/)
- url: https://www.reddit.com/r/typescript/comments/hbhajh/those_of_you_who_came_from_dynamic_languages_how/
---
So, I come from a Java / Android / Kotlin background, where the IDE is top notch and warns you over every little mistake. Maybe it warns you a bit too much, but at least I know what I'm doing wrong.

For the last 2 years I've been working with Javascript and Ruby folk. Fortunately, on my own projects I'm allowed to write Typescript, but once in awhile I have to cross over to their projects and it's so frustrating, both for me and for them.

- Sometimes they import files that don't exist. 
- ~~Sometimes~~ Most of the time they spell variable / function names wrong and the code doesn't work.
- They can't remember the whether it's an array or an object or a hash or if it's "user_id" or "userID" or "userId" or "user.id" or "user['id']".
- null pointer bugs. So many.
- Forgetting to pass variables into functions / passing in the wrong kind
- If the documentation isn't good they're screwed, because they have no autocomplete and they can't navigate to the source file to read it.
- Renaming files. Painful Ctrl-F everywhere. Half the time they'll still miss out some references anyway.
- Moving file is more painful because I have to watch them do mental sums on how many '../../' to type.
- Nobody can remember what data their objects contain after going on vacation and coming back.

I get really annoyed whenever I have to interact with their code bases. Either I make so many stupid bugs that would never happen in Typescript, or I have to sit through and painfully watch them make stupid bugs that would never happen in Typescript and waste half a day over stupid things like wrong variable names.

Been trying to convince them to use Typescript because it feels like you're coding blindly without it, but all of them just don't see any point to it? I don't know. Are they just too blind to realize the advantages? Or maybe I'm the one that's too dumb and everyone else is really smart and that's why I need a crutch like Intellisense.

Those of you who came from non-typed backgrounds like Javascript / Ruby, how on earth did you survive?
## [6][Newbie Help with union types](https://www.reddit.com/r/typescript/comments/hbh47x/newbie_help_with_union_types/)
- url: https://www.reddit.com/r/typescript/comments/hbh47x/newbie_help_with_union_types/
---
I am attempting to do this

&amp;#x200B;

```typescript

let username: String | HTMLInputElement = document.querySelector('#username')



if (username.reportValidity()) {

  username = (username as HTMLInputElement).value

  verify({ username })

}

```

&amp;#x200B;

But typescript wont let me make a union type with String and HTMLInputElement

&amp;#x200B;

```typescript

Type 'Element' is not assignable to type 'String | HTMLInputElement'.

  Type 'Element' is missing the following properties from type 'HTMLInputElement': accept, align, alt, autocomplete, and 159 more.ts(2322)

```

what do I need to do to make this happen?

Thank you so much
## [7][Variadic Kinds arriving in TypeScript 4!](https://www.reddit.com/r/typescript/comments/haw15b/variadic_kinds_arriving_in_typescript_4/)
- url: https://github.com/microsoft/TypeScript/issues/5453#issuecomment-644984977
---

## [8][Type parameters: What is the name for the part that goes in front of &lt;&gt;?](https://www.reddit.com/r/typescript/comments/hbasvp/type_parameters_what_is_the_name_for_the_part/)
- url: https://www.reddit.com/r/typescript/comments/hbasvp/type_parameters_what_is_the_name_for_the_part/
---
I see the syntax for generics is:

    function doX &lt;T&gt; (args: T): T {...}

And sometimes there is a word in front:

    function doY &lt;T&gt; (args: Array&lt;T&gt;): Partial&lt;T&gt; {...}

I know what these represent. `Array` specifies an array of values of type `T`. `Partial` is a utility type that specific some k/v pairs of the generic are there, but not all are required.

My understanding is not complete. I understand  `Array` and `Partial` are modifiers of a generic type... but how exactly? Are they keywords? Can I make my own? Do they, as a whole, have a deeper syntactical structure (`Partial.If` for example)?

Any help clarifying what that syntax is exactly is appreciated, thanks.
## [9][Using a conditional filter, error still occurs: Property 'header' does not exist on type 'never'.ts(2339)](https://www.reddit.com/r/typescript/comments/hbe5ct/using_a_conditional_filter_error_still_occurs/)
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
## [10][Non-hacky ways of doing nested classes?](https://www.reddit.com/r/typescript/comments/hbbif6/nonhacky_ways_of_doing_nested_classes/)
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
## [11][Help needed with error 'Argument of type 'any' is not assignable to parameter of type 'never'.'](https://www.reddit.com/r/typescript/comments/hb90fk/help_needed_with_error_argument_of_type_any_is/)
- url: https://www.reddit.com/r/typescript/comments/hb90fk/help_needed_with_error_argument_of_type_any_is/
---
Hi! I've been stuck on this error all day: Argument of type 'any' is not assignable to parameter of type 'never'.

 I am using Object.entries(constraints).forEach(([k, v]) =&gt; strConstraint[k as keyof StringValidator](v)) to loop over a constraints object which can contain any combination of the following: 
maxChars?: number; minChars?: number; regexToValidate? RegExp; equals?: string; notEquals?: string;. 

Here is the code for my stringValidator (error on line: 232) https://codesandbox.io/s/festive-bartik-6hdhm?file=/index.ts. If maxChars method (97-103) of number type is commented out, it works fine. 

As a last resort, I could parse a string to a number which works since everything is a string then. Any thoughts?
