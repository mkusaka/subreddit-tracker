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
## [2][Question about code reuse and inheritance](https://www.reddit.com/r/typescript/comments/ez7sfn/question_about_code_reuse_and_inheritance/)
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
## [3][Learn how to make a TypeScript Playground Plugin](https://www.reddit.com/r/typescript/comments/eywcn8/learn_how_to_make_a_typescript_playground_plugin/)
- url: https://www.typescriptlang.org/v2/dev/playground-plugins/
---

## [4][Demystifying a TypeScript quirk](https://www.reddit.com/r/typescript/comments/eyothb/demystifying_a_typescript_quirk/)
- url: https://matthewmiller.dev/blog/demystifying-typescript-quirk/
---

## [5][How to deal with leftover compiled outputs after pulling changes or switching branches?](https://www.reddit.com/r/typescript/comments/eyt2vb/how_to_deal_with_leftover_compiled_outputs_after/)
- url: https://www.reddit.com/r/typescript/comments/eyt2vb/how_to_deal_with_leftover_compiled_outputs_after/
---
I've been running into an issue that's kind of a pain and I was hoping someone has an idea on how to deal with it.

We have a typescript (v3.7) project in Visual Studio and we're using Git. Compiled outputs (.js, .d.ts, etc.) are in a .gitignore file so they aren't being tracked.

My problem is that if I delete/rename a few .ts files and commit the changes, and another developer pulls that change, the old compiled outputs are still in their file structure and the compiler (Visual Studio, in this case) will start throwing errors as it reads the .d.ts files and tries to resolve references.

**What doesn't work:**

The "clean" or "rebuild" project actions in Visual Studio. VS will only clean outputs for *existing* .ts files, but since I've deleted the .ts file it will no longer clean the now-orphaned outputs.

**Options I've considered:**

1. **"git clean -xf &lt;project_directory&gt;"** -- this works to clean up those orphaned outputs, but I'd rather not have to make people run this every time something changes.

1. **Using the .tsconfig "include" or "files" properties so I'm only trying to compile files that I know exist.** This seems super unwieldy and hard to maintain -- directory-level inclusions wouldn't be good enough (what if a directory contains multiple files but only one gets deleted?) and there are too many files to list each one individually.

1. **Using "outDir" to move all outputs to a completely separate directory, which will also automatically exclude them from the compiler.** I'm not sure of all the pros/cons of this option, honestly. It might require updating our packaging stuff to use the new location.

Is what I'm saying making sense? And is there some easier way to manage this that I'm not thinking of?

**EDIT:** I'm going with my option #3 based on the responses. It does seem pretty simple now that I look at it more closely. Thank you!
## [6][Applying SOLID principles to your TypeScript code](https://www.reddit.com/r/typescript/comments/ey8ilv/applying_solid_principles_to_your_typescript_code/)
- url: https://wanago.io/2020/02/03/applying-solid-principles-to-your-typescript-code/
---

## [7][Typing objects in TypeScript – 2ality](https://www.reddit.com/r/typescript/comments/eyen40/typing_objects_in_typescript_2ality/)
- url: https://2ality.com/2020/01/typing-objects-typescript.html
---

## [8][Playground that runs on Android](https://www.reddit.com/r/typescript/comments/eygewa/playground_that_runs_on_android/)
- url: https://www.reddit.com/r/typescript/comments/eygewa/playground_that_runs_on_android/
---
Hi,

Very often I would like to try out some ideas in TypeScript but I don't have access to a computer for various reasons, but I do have my phone.

Unfortunately the official TypeScript playground doesn't support mobile browsers.

Do you know about anything else?
## [9][A hasOwnProperty helper function/types to check for existing properties in objects](https://www.reddit.com/r/typescript/comments/ey99th/a_hasownproperty_helper_functiontypes_to_check/)
- url: https://fettblog.eu/typescript-hasownproperty/
---

## [10][Marshal.ts serializer got a JIT engine + ORM abstraction. It's now up to 300x faster than class-transformer. Time to switch guys](https://www.reddit.com/r/typescript/comments/exzdsd/marshalts_serializer_got_a_jit_engine_orm/)
- url: https://github.com/marcj/marshal.ts#marshalts
---

## [11][Question about Typescript objects assign](https://www.reddit.com/r/typescript/comments/ey4vbb/question_about_typescript_objects_assign/)
- url: https://www.reddit.com/r/typescript/comments/ey4vbb/question_about_typescript_objects_assign/
---
Hi. I'm having trouble with TypeScript objects assign. For example:

    class A {
        x: number;
        y: number;
    };
    class B {
        x: number;
        y: number;
        z: number;
    };
    const b: B = {
        x: 0,
        y: 1,
        z: 2,
    };
    const a: A = b;
    console.log(a);

I expect the result to be:

    { x: 0, y: 1 }

But I get:

    { x: 0, y: 1, z: 2 }

Is there any way to get the first result?
