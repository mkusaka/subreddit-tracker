# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][TypeScript assertion signatures for Object.defineProperty](https://www.reddit.com/r/typescript/comments/ezpya1/typescript_assertion_signatures_for/)
- url: https://fettblog.eu/typescript-assertion-signatures/
---

## [3][How to execute TS code in Visual Studio (full IDE)?](https://www.reddit.com/r/typescript/comments/ezprmb/how_to_execute_ts_code_in_visual_studio_full_ide/)
- url: https://www.reddit.com/r/typescript/comments/ezprmb/how_to_execute_ts_code_in_visual_studio_full_ide/
---
I have a junior developer who sometimes needs examples how to use typescript properly. I'm a full .NET developer and I love my Visual Studio. I can create a new TS file with code highlighting and even proper intellisense to demonstrate code snippets. It would be perfect if I could interactively execute these snippets too without setting out the whole webpage or switching IDE. I mean, it runs the whole npm / ts engine when I open up the file, it must be possible to execute it somehow? 

What I'm looking for are C#-, F#- and Python-like interactive windows.

Has anyone got it working under full Visual Studio IDE?
## [4][Algebraic graphs implementation in TypeScript](https://www.reddit.com/r/typescript/comments/ezs5x0/algebraic_graphs_implementation_in_typescript/)
- url: https://github.com/algebraic-graphs/typescript
---

## [5][Verifying and inferring type](https://www.reddit.com/r/typescript/comments/ezpfk2/verifying_and_inferring_type/)
- url: https://www.reddit.com/r/typescript/comments/ezpfk2/verifying_and_inferring_type/
---
I'm working on a web app in a Node-ish environment (Next.js API routes) where the input for `POST` data comes in a `req.body` object. I've come up with the following (elegant, I think) way of both checking that the `POST` variables have the valid type at runtime while also inferring their type in the code:

    const {shouldBeString} = extract(req.body, {
        shouldBeString: isString,
        shouldBeNumber: isNumber
    });
    
    // hover here to see that the right type is inferred
    type ShouldAlsoBeString = typeof shouldBeString;

`extract` is the function that does the extraction by taking `req.body` as a parameter and a definition object that maps field names to type narrowing functions (such as the commonplace `isString` -- also found in the [TypeScript documentation](https://www.typescriptlang.org/docs/handbook/advanced-types.html#typeof-type-guards), but can be any custom type narrowing function);

The code that makes this happen is the following:

    type Narrower&lt;T&gt; = (v: unknown) =&gt; v is T;
    type Definition = { [k: string]: Narrower&lt;unknown&gt; };

    type Narrowed&lt;T&gt; = T extends Narrower&lt;infer X&gt; ? X : never;
    type Extracted&lt;T&gt; = { [P in keyof T]: Narrowed&lt;T[P]&gt; }

    type PostData = { [k: string]: unknown };
    
    const extract = &lt;D extends Definition&gt;(post: PostData, def: D): Extracted&lt;D&gt; =&gt;
        Object.entries(def).reduce((result, [field, isRightType]) =&gt; {
            const value = post[field];
            if (!isRightType(value)) {
                throw Error(`Invalid value for field ${field}`);
            }
            return ({...result, [field]: value});
        }, {} as Extracted&lt;D&gt;);
    
    // hover here to see that the right type is inferred
    type ShouldBeString = Narrowed&lt;typeof isString&gt;;

What do you think? Any suggestions as to naming? Code style? Anything else?

(This is minimalist by design, it doesn't support optional inputs or fields having more than one value per field. Might work on that if the use case arises.)

Thanks!
## [6][How do I setup a Typescript project?](https://www.reddit.com/r/typescript/comments/ezr6hg/how_do_i_setup_a_typescript_project/)
- url: https://www.reddit.com/r/typescript/comments/ezr6hg/how_do_i_setup_a_typescript_project/
---
I want to try using typescript within a browser context but I haven't really found what is the best way to set-up such a project. I'm mainly having trouble with what I should do with npm downloaded packages, I'm using PHPStorm so the typescript compiling itself is handled.

I'm looking to use systemjs as a loader as well.
## [7][TS server detects error in IDE, ts-jest passes the test without raising the type error.](https://www.reddit.com/r/typescript/comments/ezqp3i/ts_server_detects_error_in_ide_tsjest_passes_the/)
- url: https://i.redd.it/jmywutwy9af41.gif
---

## [8][Create single ES6 module?](https://www.reddit.com/r/typescript/comments/ezqnp4/create_single_es6_module/)
- url: https://www.reddit.com/r/typescript/comments/ezqnp4/create_single_es6_module/
---
I'm just getting started with typescript and what I'd like to achieve is to export/build a single ES6 module from all files with just typescript. It seems like typescript can do that for AMD or system but I didn't find a way to do it with ES6. Is this not possible? Tbh, that would be a dealbreaker for me. I want to avoid webpack and rollup because they add unnecessary overhead. As the AMD option shows, tsc should already be capable of doing single file builds.
## [9][Can you create a subset of a const enum?](https://www.reddit.com/r/typescript/comments/ezdawh/can_you_create_a_subset_of_a_const_enum/)
- url: https://www.reddit.com/r/typescript/comments/ezdawh/can_you_create_a_subset_of_a_const_enum/
---
I'm trying to create a  subset of a const enum, but it's [not working](https://www.typescriptlang.org/play/#code/MYewdgzgLgBApmArgWxgQQDYYHIoEZwBOMA3gFAyUwDyYcMAvDAEThzMA0FVAKgO4hGLKAM7dKPABaE49JsyjTZzMgF8yZUJFgIUNACb7cyAsXJUadIZhz4iAOlpwuFqTLnosx0-beyu6mQAZohgwFAAluAwERDUht5EABQAHgBcMNCEEWAA5gCUGSkxEAZGdmbimXwRUMCSMKn5pFUWwACGEPTx5SYOTmmtVB1dZYmEvkpwgxazVDJQiIRgMFCEiHAA3EOU+nBB7YgYUDNzswtLKwcYXdsW6qpAA) as I'd expect. Is something like this possible, or would you recommend another approach?
## [10][Question about code reuse and inheritance](https://www.reddit.com/r/typescript/comments/ez7sfn/question_about_code_reuse_and_inheritance/)
- url: https://www.reddit.com/r/typescript/comments/ez7sfn/question_about_code_reuse_and_inheritance/
---
I have a question regarding the best way to structure code for reuse with inheritance in cases, where some methods in the base class does not make sense in the inheriting class.

I think this is a general OOP issue and not only peculair to TypeScript...

So basically, in some form of pseudo code, the issue is similar to this:

`class BaseClassmakesSenseForAmakesSsenseForA....doesnotMakeSenseForAdoesnotMakeSenseForA....`

How then to have a `Class A` inherit BaseClass. Just direct inheritance means that `Class A`  would have member methods that does not really make sense.

Not using inheritance means the things that code would be duplicated in both `BaseClass` and `Class A`

Moving  the `makesSenseForA` methods to an external class and have both `BaseClass` and it's children class depend on this extracted class (ie using composition) does not work in this particular scenario, because the extracted class should actually be seen as A `BaseClass.`

How really is the best way to handle this kind of OOP modeling situation?

EDIT: example

For lack of better example, it is like trying to model \`Human\` and \`Cyborg\` - this two would share a ton loads of similar implementation and hence first thought might be to have \`Cyborg\` extends \`Human\`...

&amp;#x200B;

But also there would be a ton of implementation that \`Human\` has that \`Cyborg\` should not have.

&amp;#x200B;

Doing away with inheritance means those similar functionality would have to be duplicated in \`Human\` and Cyborg\`.

&amp;#x200B;

And composition also does not work, because if you extract those implementation that are similar into a separate object, that object, and it's method would constitute human like characteristics, and hence should be seen as a \`Human\` in our modelling.
## [11][Tip: type npx typescript-react , to bootstrap a starter application](https://www.reddit.com/r/typescript/comments/ezgfpe/tip_type_npx_typescriptreact_to_bootstrap_a/)
- url: https://www.reddit.com/r/typescript/comments/ezgfpe/tip_type_npx_typescriptreact_to_bootstrap_a/
---

