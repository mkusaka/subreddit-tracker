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
## [2][Demystifying a TypeScript quirk](https://www.reddit.com/r/typescript/comments/eyothb/demystifying_a_typescript_quirk/)
- url: https://matthewmiller.dev/blog/demystifying-typescript-quirk/
---

## [3][Applying SOLID principles to your TypeScript code](https://www.reddit.com/r/typescript/comments/ey8ilv/applying_solid_principles_to_your_typescript_code/)
- url: https://wanago.io/2020/02/03/applying-solid-principles-to-your-typescript-code/
---

## [4][Typing objects in TypeScript – 2ality](https://www.reddit.com/r/typescript/comments/eyen40/typing_objects_in_typescript_2ality/)
- url: https://2ality.com/2020/01/typing-objects-typescript.html
---

## [5][Playground that runs on Android](https://www.reddit.com/r/typescript/comments/eygewa/playground_that_runs_on_android/)
- url: https://www.reddit.com/r/typescript/comments/eygewa/playground_that_runs_on_android/
---
Hi,

Very often I would like to try out some ideas in TypeScript but I don't have access to a computer for various reasons, but I do have my phone.

Unfortunately the official TypeScript playground doesn't support mobile browsers.

Do you know about anything else?
## [6][A hasOwnProperty helper function/types to check for existing properties in objects](https://www.reddit.com/r/typescript/comments/ey99th/a_hasownproperty_helper_functiontypes_to_check/)
- url: https://fettblog.eu/typescript-hasownproperty/
---

## [7][Marshal.ts serializer got a JIT engine + ORM abstraction. It's now up to 300x faster than class-transformer. Time to switch guys](https://www.reddit.com/r/typescript/comments/exzdsd/marshalts_serializer_got_a_jit_engine_orm/)
- url: https://github.com/marcj/marshal.ts#marshalts
---

## [8][Question about Typescript objects assign](https://www.reddit.com/r/typescript/comments/ey4vbb/question_about_typescript_objects_assign/)
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
## [9][What's been your experience using TypeScript with React Hooks?](https://www.reddit.com/r/typescript/comments/exwqq2/whats_been_your_experience_using_typescript_with/)
- url: https://www.reddit.com/r/typescript/comments/exwqq2/whats_been_your_experience_using_typescript_with/
---
I've been working on a project involving both React and TypeScript. I've decided to start moving the project in the direction of hooks. What I didn't realize when starting this process was that React hooks provide a way to work in React without having to use classes. This kind of bummed me out since I've been learning a lot about TypeScript's added class features like modifying class properties to be private / public / protected, etc.

This got me thinking about using React and TypeScript together, and about how Facebook has also made Flow, which is kind of Facebook's analog to TypeScript (at least in terms of type declarations). 

Are React and TypeScript at odds with each other? Do you think these two technologies are moving away from each other? Would you recommend using them together? Why/why not?
## [10][Could I get a quick code review on these three simple functions?](https://www.reddit.com/r/typescript/comments/exzdjn/could_i_get_a_quick_code_review_on_these_three/)
- url: https://www.reddit.com/r/typescript/comments/exzdjn/could_i_get_a_quick_code_review_on_these_three/
---
I'm just wondering if you think these three functions could be written any better. Their purpose is to attach/inject a hook into the invocation lifecycle of a target function.

Thanks.

```
export enum HookType {
    BEFORE,
    AFTER
};

export type TargetContext = { [key: string]: any };
export type AnyFunction = (...args: any[]) =&gt; any;

export const isTypeFunction = (candidate: any): candidate is AnyFunction =&gt; {
    return (typeof candidate === 'function');
}

/**
 * Attaches a hook function of a specific type to fire for a target function invocation.
 * @param hookType   The type of hook to apply to the target function.
 * @param context    The context that the target function exists within.
 * @param targetName The name of the target function.
 * @param hook       The hook function execute.
 */
export type ReplaceTarget = &lt;T extends TargetContext, K extends keyof T&gt;(
    hookType: HookType, 
    context: T,
    targetName: K,
    hook: AnyFunction
) =&gt; void;

export const replaceTarget: ReplaceTarget = &lt;T extends TargetContext, K extends keyof T&gt;(
    hookType: HookType, 
    context: T,
    targetName: K,
    hook: AnyFunction
): void =&gt; {
    const originalFunc = context[targetName];
    
    if (!isTypeFunction(originalFunc))
        throw new Error(`TypeError: context[${targetName}] is not a function.`);

    const callOriginal = (...args: any[]) =&gt; originalFunc.apply(context, args);
    const callHook = (...args: any[]) =&gt; hook.apply(context, args);

    let originalReturnValue;

    context[targetName] = ((...args: any[]): typeof originalFunc =&gt; {
        switch (hookType) {
            case HookType.BEFORE:
                callHook(...args);
                originalReturnValue = callOriginal(...args);
                break;
            case HookType.AFTER:
                originalReturnValue = callOriginal(...args);
                callHook(...args);
                break;
            default:
                throw new Error(`HookError: The provided HookType is invalid.`);
        }

        return originalReturnValue;
    }) as T[K];
};

export const backupTargets = &lt;T extends TargetContext, K extends keyof T&gt;(
    context: T, 
    targetNames: K[]
): T =&gt; {
    const targetContextBackup: T = {} as T;

    Object.getOwnPropertyNames(context).forEach(propName =&gt; {
        if (targetNames.indexOf(propName as K) &gt; -1) {
            targetContextBackup[propName as K] = context[propName];
        }
    });

    return targetContextBackup;
};

export const restoreTargets = &lt;
    T extends TargetContext, 
    K extends keyof T,
&gt;(
    contextToRestore: TargetContext,
    backupContext: T,
    targetNames: K[]
): void =&gt; {
    Object.getOwnPropertyNames(contextToRestore).forEach(propName =&gt; {
        if (targetNames.indexOf(propName as K) &gt; -1) {
            contextToRestore[propName] = backupContext[propName];
        }
    });
}
```
I'm not sure if that type casting I'm doing is a best practice, for example.

**EDIT:** It was reported that the code is not formatted probably here. Here is a picture of the snippet: https://imgur.com/a/dVP5Tho

Thank you.
## [11][What is the VScode plugin Anders Hejlsberg is using to get the types underlined?](https://www.reddit.com/r/typescript/comments/exkh20/what_is_the_vscode_plugin_anders_hejlsberg_is/)
- url: https://i.redd.it/o780xxzgige41.png
---

