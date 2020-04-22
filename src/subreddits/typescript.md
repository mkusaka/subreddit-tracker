# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Ofnotes: Offline note taking application with live markdown rendering rewritten in typescript, hooks, context, and material-ui with improved user experience and 100% test coverage](https://www.reddit.com/r/typescript/comments/g5pz1f/ofnotes_offline_note_taking_application_with_live/)
- url: https://v.redd.it/cfgv4ht379u41
---

## [3][Parallelizing Work via a Semaphore](https://www.reddit.com/r/typescript/comments/g5xtkf/parallelizing_work_via_a_semaphore/)
- url: https://alexn.org/blog/2020/04/21/javascript-semaphore.html
---

## [4][How to make prop mandatory based on another prop being passed](https://www.reddit.com/r/typescript/comments/g60wei/how_to_make_prop_mandatory_based_on_another_prop/)
- url: https://www.reddit.com/r/typescript/comments/g60wei/how_to_make_prop_mandatory_based_on_another_prop/
---
I have an interface:

```
interface FormProps {
  regexPreset?: RegexPresets;
  customRegex?: RegExp;
  description?: string;
  inputTitle: string;
  errorMessage?: string;
}
```
Implemented like this:
```
const Form: React.FC&lt;FormProps&gt; = props =&gt; {
  return &lt;div&gt; formstuff &lt;/div&gt;
} 
```
If I pass in a `customRegex`, I want the compiler to throw an error if `errorMessage` is not passed (make errorMessage a mandatory property). 

If `customRegex` is **not** passed, then **don't** make `errorMessage` mandatory.

The closest thing that I've come to is this [StackOverflow post](https://stackoverflow.com/questions/52771626/typescript-react-conditionally-optional-props-based-on-the-type-of-another-prop), but I'm unsure whether I can apply this to my use case.

I asked this [on StackOverflow](https://stackoverflow.com/questions/60915996/how-to-make-prop-mandatory-based-on-another-prop-being-passed) and came up blank. I'm hoping the minds of /r/TypeScript can help me solve this!
## [5][Is there any way for a function to have a narrowing effect on array/tuple of arguments? Type guards?](https://www.reddit.com/r/typescript/comments/g5y138/is_there_any_way_for_a_function_to_have_a/)
- url: https://www.reddit.com/r/typescript/comments/g5y138/is_there_any_way_for_a_function_to_have_a/
---
Let's say I have a bunch of optional variables

    let client: Client|undefined =  // ...
    let cache: Cache|undefined =  // ...
    let sessionId: string|undefined = // ...
    // etc.

Is there some way to have a generic function (any number of arguments), such that

    if (validateAll(client, cache, sessionId)) {
        // ...
    }

would narrow the types of all variables to be non-undefined inside that `if` block (when the function returns true)?

I was looking into [type guards](https://www.typescriptlang.org/docs/handbook/advanced-types.html#type-guards-and-differentiating-types), but it doesn't seem like those can achieve what I'm after.

I can make a type that transforms a tuple of possibly undefined (or null) values to non-undefined ones by adapting `Required&lt;T&gt;`

    // E.g. turns `[string|undefined, number|undefined]`
    // into `[string, number]`
    type RequiredStrict&lt;T&gt; = {
        [P in keyof T]-?: Exclude&lt;T[P], undefined|null&gt;;
    }

But I don't see a way to use that in a type guard function. With a rest argument

    function validateAll&lt;T extends CanValidate[]&gt;(
        ...items: T
    ): items is RequiredStrict&lt;T&gt;

I get an explicit error *A type predicate cannot reference a rest parameter*. With an array (`items: T[]`) it doesn't do any narrowing

    function validateAll&lt;T extends CanValidate&gt;(
        items: T[]
    ): items is RequiredStrict&lt;T[]&gt;

    if (validateAll([client, cache, sessionId])) {
        // `client` still has a `Client|undefined` type here
    }

mabye because it treats it like a regular array instead of a tuple? Maybe having some way to force it to treat it like a tuple would fix it? (e.g. https://github.com/microsoft/TypeScript/issues/16656).

Also, type guards may not be suitable for this in principle, because the compiler expects them to *only* do type checking (at least the docs suggest that), instead of other validation. Something like a more adavanced version of Kotlin's [contracts](https://kotlinlang.org/docs/reference/whatsnew13.html#contracts) (specifically the `returns (value) implies (type assertion)` contract) would help, but TypeScript doesn't have anything like that planned as far as I know.

Can this be achieved in any way? Or at least the multi-undefined check without additional validation logic?
## [6][I'm almost there to release the alpha but I need a little help](https://www.reddit.com/r/typescript/comments/g5xtw1/im_almost_there_to_release_the_alpha_but_i_need_a/)
- url: https://www.reddit.com/r/typescript/comments/g5xtw1/im_almost_there_to_release_the_alpha_but_i_need_a/
---
I am about to make an official announcement about the Alpha release of [Typetron](http://typetron.org/), the Node.js framework I am working on. 

Before that, I would need some help from 1 or 2 of you to take the first tutorial [here](https://typetron.org/tutorials) and give feedback about any issues found, so I can fix them before the release. 

I would be cool if we can have a call and screen-share so I can take notes about what steps you do and what is missing from the tutorial. This won't take more than 1 hour (unless we start discussing the cool stuff under the hood :D )

What do you say?
## [7][Term for the ability to create variants of object types?](https://www.reddit.com/r/typescript/comments/g5o41c/term_for_the_ability_to_create_variants_of_object/)
- url: https://www.reddit.com/r/typescript/comments/g5o41c/term_for_the_ability_to_create_variants_of_object/
---
Is there a general term for the ability to create modified versions of an object/struct type in a type system?

The obvious examples in TS would be \`Pick&lt;T&gt;\` , \`Omit&lt;T&gt;\`, \`Partial&lt;T&gt;\`, and \`Required&lt;T&gt;\`.

Context: I recently started learning Dart (having used TS almost exclusively for the past 2 years) and was stunned that a) the only way to create a object/struct type is by defining a new class and b) there's no equivalent to Pick, Omit, Partial, etc whatsoever. So I'm trying to find the words to describe my frustration 😅
## [8][Myzod v1.0.0-alphar release - Schema Validation and Type Inference](https://www.reddit.com/r/typescript/comments/g5eybi/myzod_v100alphar_release_schema_validation_and/)
- url: https://www.reddit.com/r/typescript/comments/g5eybi/myzod_v100alphar_release_schema_validation_and/
---
[Myzod](https://www.npmjs.com/package/myzod) is a runtime validation library who's goal is to use only typescript concepts to build type infer-able schemas.  The purpose of [myzod](https://www.npmjs.com/package/myzod) is to no longer have to match declared typescript types to the result of separately maintained validation logic, and in so doing minimise discrepancies between runtime  and the compile time types. Myzod is also inspired by [@hapi/joi](https://www.npmjs.com/package/@hapi/joi) and offers a similar validation api.  


At this time I am about ready to release version 1.0.0 but I am hoping to get more eyes on it before I do. Any feedback, issues, feature requests, or PRs would be extremely valuable.

Thanks in advance.
## [9][Vue + TS + Chart JS - Property 'renderChart' does not exist on type 'LineChartComponent'](https://www.reddit.com/r/typescript/comments/g5qm21/vue_ts_chart_js_property_renderchart_does_not/)
- url: https://stackoverflow.com/questions/61354901/property-renderchart-does-not-exist-on-type-linechartcomponent#61354901
---

## [10][Purify - Functional programming library for TypeScript](https://www.reddit.com/r/typescript/comments/g5s0ma/purify_functional_programming_library_for/)
- url: https://gigobyte.github.io/purify/
---

## [11][Generics headache](https://www.reddit.com/r/typescript/comments/g5icot/generics_headache/)
- url: https://www.reddit.com/r/typescript/comments/g5icot/generics_headache/
---
Hey guys, sorry if this question is dumb/has been answered elsewhere, but I'm trying to create a util function (not saying that it is a good a idea to do this or not), to only execute a code once, have it just kept under wraps by reference.

    const once: &lt;O, K extends keyof O, V extends O[K]&gt;(object: O, key: K, task: () =&gt; V) =&gt; V = (object, key, task) =&gt; {
        return (object[key] = object[key] || task());
    };

tsc is spitting an error at me tho:

&gt;'O\[string\]' is assignable to the constraint of type 'V', but 'V' could be instantiated with a different subtype of constraint '{}'

And I can't seem to understand it, where is the constraint coming from.

Is it because K can be symbol/number/string at once? But if I restrict it, it still doesn't work.

&amp;#x200B;

On the function interface I am asking for the object, for a key of that object, and the value generator.

Basically, I'm not seeing what im doing wrong. Does anybody have a pointer?
