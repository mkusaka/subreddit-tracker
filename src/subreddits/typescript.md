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
## [2][Single value as a return type?](https://www.reddit.com/r/typescript/comments/ha1pog/single_value_as_a_return_type/)
- url: https://www.reddit.com/r/typescript/comments/ha1pog/single_value_as_a_return_type/
---
VSCode doesn't lint if I set the return type as a static value. I have an orchestrator method that should only ever err or return true:

       public run(): true {

Another option is to just set it to `void`. That actually sounds more sensible.

Does it ever make sense to set the return type to a static non-void value?
## [3][thetutlage/japa the missing Typescript-Typescript unit test library](https://www.reddit.com/r/typescript/comments/h9p03q/thetutlagejapa_the_missing_typescripttypescript/)
- url: https://github.com/thetutlage/japa
---

## [4][Show Reddit: Computing with types in TypeScript](https://www.reddit.com/r/typescript/comments/h9vaxq/show_reddit_computing_with_types_in_typescript/)
- url: https://2ality.com/2020/06/computing-with-types.html
---

## [5][[Newbie] Type annotations for functions](https://www.reddit.com/r/typescript/comments/h9xl4p/newbie_type_annotations_for_functions/)
- url: https://www.reddit.com/r/typescript/comments/h9xl4p/newbie_type_annotations_for_functions/
---
I came accros this code! I'm a bit confused on the first code and what the differences between the 2 are! 

    // first
    const logNumber: (i: number) =&gt; void = (i: number) =&gt; {
      console.log(i);
    };
    //Annotating the  variable with (i:number)? And the function must return void? 
    //Why are we anotating the variable with (i:number)?
     
    // second
    const logNumber = (i:number): void =&gt; {
      console.log(i);
    };
    //annotaitng the function, where i must be a number, return type must be void
## [6][Type guard is acting strange. Is this a bug, or am I confused?](https://www.reddit.com/r/typescript/comments/h9isnt/type_guard_is_acting_strange_is_this_a_bug_or_am/)
- url: https://www.reddit.com/r/typescript/comments/h9isnt/type_guard_is_acting_strange_is_this_a_bug_or_am/
---
Idk if I found a bug or something is going over my head, but I've written out a super simplified example of code I wrote where the type error is displayed, and one where it is not.

[example code 1 (error)](https://www.typescriptlang.org/play/#code/MYewdgzgLgBAygFQKIAUD6BJAIjAvDAbwCgZSYAzASwCdoAueATwFsAjEAGwDpyRqAKAOQBGQQEoANCTIQApqDAATBnBbtuvAYIBM4qQF8iRKIwAOs+FFmm8haaUrKYENZwDc9mGFkAPKAH4VK1MPQyIFaGdgiCDrAG0AXVs4z2IydJhHFWR0bB4aaCkMsm8-BgJMp0RUTCwuOQVFGH0isha7YocqnNr6+XBFVuLSqAYAVyVZKm9Bz0MEjyJyCeAoSnAKSiU4YIAhRgxFfmhrQ4YAa1lGEHIYE3Mb+B7sMQ6yDllYE9NYmwAfGATRRTLayJr4b4QOIABgWRnSlFu-AAhN9XmlOtRPmNqGAPOkwukAO4AC0oHxgx2CXEcMGRuHw1VyWDi30OCXRnnSEVgI1s3y4I0WnUySORI1eWKgOLx8M631sQrmctIUplUWsoSIQA)

[example code 2 (working)](https://www.typescriptlang.org/play/#code/MYewdgzgLgBAygFQKIAUD6BJAIjAvDAbwCgZSYAzASwCdoAueATwFsAjEAGwDpyRqAKAOQBGQQEoANCTIQApqDAATBnBbtuvAYIBM4qQF8iRKIwAOs+FFmm8haaUrKYENZwDc9mGFkAPKAH4VK1MPQyIFaGdgiCDrAG0AXVs4z2IydJhHFWR0bB4aaCkMsm8-BgJMp0RUTCwuOQVFGH0isha7YocqnNr6+XBFVuLSqAYAVyVZKm9Bz0MEjyJyCeAoSnAKSiU4YIAhRgxFfmhrQ4YAa1lGEHIYE3Mb+B7sMQ6yDllYE9NYmwAfGATRRTLayJr4b4QOIABgWRnSlFu-AAhN9XmlOtRPmNqGAPOkwukAO4AC0oHxgx2CXEcMGRuHw1VyWDi30OCXRngRSNR1JGryxUBxePhnW+tm+XBG+LaotIguFUWsoSIQA)

Does anyone know why the guard at line 24-26 in example code 1 does not work, but does work in example 2? I'm fairly confident in my TypeScript abilities but this one is not making sense in my brain.
## [7][How are branded types valid types?](https://www.reddit.com/r/typescript/comments/h9udyv/how_are_branded_types_valid_types/)
- url: https://www.reddit.com/r/typescript/comments/h9udyv/how_are_branded_types_valid_types/
---
An example of a branded type is:

`type USD = number &amp; {__brand: "USD"}`

You can then cast numbers to this type with \`as USD\`. However, I don't understand how this works. How can you intersect a primitive with an object? I thought objects could only be intersected with other objects, `object`, or `any`.
## [8][How to create a JSON Object's element based on if it is null or not?](https://www.reddit.com/r/typescript/comments/h9xfpa/how_to_create_a_json_objects_element_based_on_if/)
- url: https://www.reddit.com/r/typescript/comments/h9xfpa/how_to_create_a_json_objects_element_based_on_if/
---
    async homeApi(_source: any, _args: any) { const body = {           door: _args.door,           window: _args.window         }; }

I have a typescript code like above where I create a JSON Object called bodyusing doorand windowarguments. Typically the body should be:

    {     door: 3,     window: 4 }

What I want is if \_args.dooris empty/blank, the bodyshould be

    {window: 4}

It should not be: {         door: undefined,         window: 4        }
## [9][Beginner needs help writing his first .d.ts file](https://www.reddit.com/r/typescript/comments/h9o9g9/beginner_needs_help_writing_his_first_dts_file/)
- url: https://www.reddit.com/r/typescript/comments/h9o9g9/beginner_needs_help_writing_his_first_dts_file/
---
Hello everyone!

I’m currently working on a react component library for my own personal use and so far I’ve built my components, compiled them to JS and it all works as expected when I import my lib from a new project.

However when I import the lib, TS demands a .d.ts file... to override this while working on my lib I just had TS ignore the missing .d.ts file.

But now I’m ready to actually supply this file and I’ve never written one.

I read some of the TS docs this morning and although I’ll appreciate any input, my main question is:

Where does the .d.ts file go?

Do I put it in the root of the lib where node_modules is?

Do I put it inside of the src folder?

Or am I creating a .d.ts file for every module of my lib?

I’m just throwing this out there, any input is appreciated! Thank you! 🙏🏽
## [10][String Literal Propagation Issue](https://www.reddit.com/r/typescript/comments/h9laux/string_literal_propagation_issue/)
- url: https://www.reddit.com/r/typescript/comments/h9laux/string_literal_propagation_issue/
---
I'm trying to propagate a string literal through a function argument while also accepting another generic type. What I'm trying to do is get the result of the action and put it in a list of actions for a switch case Here is where I'm at.
```
export function generateAction&lt;Payload, T extends string = string&gt;(
	type: T
): [T, (payload: Payload) =&gt; { type: T; payload: Payload }] {
	return [type, payload =&gt; ({ type, payload })];
}

export const [SWITCH_SECTION, switchSection] = generateAction&lt;{ section: number }&gt;('SWITCH_SECTION');
export const [SWITCH_SUBSECTION, switchSubsection] = generateAction&lt;{ subsection: number }, 'SWITCH_SUBSECTION'&gt;(

type actionTypes = ReturnType&lt;typeof switchSection&gt; | ReturnType&lt;typeof switchSubsection&gt;;
```

By specifying the payload type it makes the string generic dumb and doesn't use the string literal. If I use the below the string literal is preserved.

```
export const [SWITCH_SECTION, switchSection] = generateAction('SWITCH_SECTION');
export const [SWITCH_SUBSECTION, switchSubsection] = generateAction('SWITCH_SUBSECTION');

type actionTypes = ReturnType&lt;typeof switchSection&gt; | ReturnType&lt;typeof switchSubsection&gt;;
```
[Playground Link](https://www.staging-typescript.org/play?#code/KYDwDg9gTgLgBAMwK4DsDGMCWEVwObArBQCGMwAghtigDwAKJAngDYQkAmANHACpyhyKDgGc4ImFEwo8cALzjJ0vAD4AFACg42uDCZhgALj4aAlMYDavHmrDM2nY41bsOp+SrgBvXfqN8Abjg7F0c4ZwcOOABfAF1vLR0oYBgkKFwLPQMeEMiPODUfLOAc+1cY01iAjWiNDVBIWDg0HAk4CwBlAHUASV4AYQAJAH0OgFF+3h6AeQA5HhEAd0wYNAALDuBqHHiFAiJSciosHFofES2TlGMUJABbACNiGPUAcm6+odGJqbnX02qDWg8BaKDanV6AxGHQAqgAhcaTGbzcTLVYbJAPC7bFC7fCEYhkSg4s7iTHYq43e5PKAxHjvSFfWEIn7I15vD5Q0bwxG-Wb-QHgYHNVrwCGfaE81lzABMCzR602OJleP2hKOOLUDIl3JZSNlArqxTgJBxvD8YgUACUUmkUOaDLRihAEKiVorLjRPAAfOA21LpB3AJ1+F1u9EdcmenA+v22wN+EMGMNLd0baMoGUqAJAA)
Any idea how to preserve the string literal while also specifying another generic?
## [11][Must I always extend protected members to be protected?](https://www.reddit.com/r/typescript/comments/h9g59c/must_i_always_extend_protected_members_to_be/)
- url: https://www.reddit.com/r/typescript/comments/h9g59c/must_i_always_extend_protected_members_to_be/
---
I have something like this:

    abstract class Base {
        protected description: string = ``;
    }
    
    class Subclass extends Base {
        private description: string = ``;
    
    }
    
    /* class Subclass
    Class 'Subclass' incorrectly extends base class 'Base'.
      Property 'description' is private in type 'Subclass' but not in type 'Base'.ts(2415) */

I don't plan to extend `Subclass` so I thought marking `description` as `protected` would not be necessary or very clear to others. I therefore thought `description` in `Subclass` should be marked `private`, but only the `protected` keyword avoids errors.

Can anyone help me understand why `protected` is needed here?

Is the case simply that any time you extend a class you must identify overwritten members with the `protected` keyword?
