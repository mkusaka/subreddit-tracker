# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Runtime JSON type checks with Typescript interfaces](https://www.reddit.com/r/typescript/comments/ifm356/runtime_json_type_checks_with_typescript/)
- url: https://medium.com/bytecodeagency/runtime-json-type-checks-with-typescript-interfaces-379e8ea81258?sk=489009cb893e8a700956d9e51ff855d7
---

## [3][tsc vs ts-loader in the context of project references](https://www.reddit.com/r/typescript/comments/iflwtl/tsc_vs_tsloader_in_the_context_of_project/)
- url: https://www.reddit.com/r/typescript/comments/iflwtl/tsc_vs_tsloader_in_the_context_of_project/
---
Wondering what is better in terms of performance and project reference support:

1. Compilation done by `tsc -b -w` only; `webpack` and `webpack-dev-server` only bundle already compiled `js` files and only frontend code

2. `webpack` + `ts-loader` for frontend code

You find more background in this [Github issue](https://github.com/TypeStrong/ts-loader/issues/1157#issuecomment-679012708)
## [4][TypeScript ignores type annotation for string literals?](https://www.reddit.com/r/typescript/comments/iffdlr/typescript_ignores_type_annotation_for_string/)
- url: https://www.reddit.com/r/typescript/comments/iffdlr/typescript_ignores_type_annotation_for_string/
---
    const foo: string | undefined = "";
    // TS ignores the type annotation and decides the type is `string`.
    foo.toUpperCase();
    
    const bar: string | undefined = Math.random() &gt; 0 ? "" : undefined;
    // tsc error as expected
    bar.toUpperCase();

Is this new or has this always been the case? Is this documented anywhere? Is it a bug?
## [5][Looking for best practices, style guidelines and naming conventions](https://www.reddit.com/r/typescript/comments/ifjxam/looking_for_best_practices_style_guidelines_and/)
- url: https://www.reddit.com/r/typescript/comments/ifjxam/looking_for_best_practices_style_guidelines_and/
---
Are there any style guidelines, projects to look at, or best practices available on how to use typescript?  
Currently I am bit of a loss on how I should structure the project I am working on.

Do all the type declarations belong at the beginning of the file? Do I put the into a separate declaration file? Do I spread them out across the source file? How about naming convention?

Tutorials usually cover how to use the language but I haven't found a style which feels organized. e.g. the types packages all look vastly different and some of them are just incomprehensible at first glance.
## [6][rxjs filter pipe doesn't narrow type?](https://www.reddit.com/r/typescript/comments/ifklp4/rxjs_filter_pipe_doesnt_narrow_type/)
- url: https://www.reddit.com/r/typescript/comments/ifklp4/rxjs_filter_pipe_doesnt_narrow_type/
---
Consider this example:

`const source$ = of(5, 'abc');`  
`source$.pipe(filter(val =&gt; typeof val === 'string')).subscribe(val =&gt; {`  
 `// val: string | number;`  
`});`

Why doesn't the filter pipe here narrow down the type to only string? Is there a rxjs operator that could do the filtering?

Thanks
## [7][Narrowing strings to keys of own type?](https://www.reddit.com/r/typescript/comments/ifmrob/narrowing_strings_to_keys_of_own_type/)
- url: https://www.reddit.com/r/typescript/comments/ifmrob/narrowing_strings_to_keys_of_own_type/
---
I am designing a library with an API like this:

    let myThing = thing({
      identifier: ["name", "size", "foobarbaz"],
      properties: {
        name: { type: String },
        size: { type: Number },
        color: { type: Color },
        text: { type: String }
      }
    });

It basically means that you create a kind of thing with these properties, and properties "name" and "size" identify it unambiguously.

I want it to be impossible to mention a field inside `identifier` which is not a part of `properties`. So the example above should show an error at `"foobarbaz"`, which is not a valid property name. My attempt looks like this:

    function thing&lt;TProp extends string, T extends {
        identifier: TProp[],
        properties: { [key in TProp]: any }
    }&gt;(data: T): T {
        return data;
    }

However, this doesn't work. I assume that TypeScript reads the `identifier` first and assumes that whatever is written there is a part of `TProp` and therefore valid. What can I do to make it show an error in my example?

To clarify, I want the type to be inferred completely, without requiring the user to provide the type themselves.

Thanks!
## [8][[NEED HELP] I need to build APIs which in turn call 3rd party APIs using NodeJS and TypeScript. Any preferred path or template?](https://www.reddit.com/r/typescript/comments/ifm18a/need_help_i_need_to_build_apis_which_in_turn_call/)
- url: https://www.reddit.com/r/typescript/comments/ifm18a/need_help_i_need_to_build_apis_which_in_turn_call/
---
Title.
## [9][How can I extend an interface with increased specificity for a type?](https://www.reddit.com/r/typescript/comments/if7ytw/how_can_i_extend_an_interface_with_increased/)
- url: https://www.reddit.com/r/typescript/comments/if7ytw/how_can_i_extend_an_interface_with_increased/
---
Say I have the following "generic" interface that defines props for a component:

```
export type MyType = string | number;

export interface BaseProps {
    ...
    onChange: (newValue: MyType) =&gt; void;
}
```

Now what I'd like to do is write an interface that extends this, but specifies which of the two values is actually going to be used. The type used depends on the component. Both are never valid at the same time.

```
interface Props extends BaseProps {
    ...
    onChange: (newValue: string) =&gt; void;
}
```

This does not work, with the error that "string is not assignable to number", which makes sense. I understand that this isn't a correct usage of TS types. 

However, is there a way to do what I want here? I could just slap an `any` in `BaseProps` but I like that `MyType` clearly defines what `onChange` will actually do. Otherwise you'd have to go find where it's extended to understand what kind of values it uses.

It would also be nice so that I could ensure I wasn't using a type that the base did not allow. For example: 

```
interface Props extends BaseProps {
    ...
    onChange: (newValue: boolean) =&gt; void;
}
```

Having this be rejected would be quite convenient. 

I'd like to declare at a high level all possible options and gain specificity as I go deeper down, if that makes sense.
## [10][How do I contribute to the TS community](https://www.reddit.com/r/typescript/comments/if7iu0/how_do_i_contribute_to_the_ts_community/)
- url: https://www.reddit.com/r/typescript/comments/if7iu0/how_do_i_contribute_to_the_ts_community/
---
I have been coding in JavaScript since last 6 years. Recently I completed a project in TypeScript. I am in love with this language.

I really want to contribute to TS community. I'm hoping that this will also help me get a good grip on the language.

I'll appreciate any suggestions. Feel free to ask any questions.
## [11][Why does TypeScript still think this number could be undefined?](https://www.reddit.com/r/typescript/comments/if2zc6/why_does_typescript_still_think_this_number_could/)
- url: https://www.reddit.com/r/typescript/comments/if2zc6/why_does_typescript_still_think_this_number_could/
---
I've just spent 20+ minutes playing around with a function like this and I'm lost. Is TypeScript's type inference not able to figure this out, or am I missing something obvious? 

(I'm running pretty low on sleep so it's very possibly the latter)

edit: The main point is by the time we get to `const num = a ?? b;`, we already know that at least one of those values is guaranteed to be a number, but the compiler still treats num as potentially undefined.

```typescript
function myFunc(a?: number, b?: number): number {
  if (!a &amp;&amp; !b) throw new Error('No value provided');

  const num = a ?? b;

  return num; // Type 'number | undefined' is not assignable to type 'number'
}
```
