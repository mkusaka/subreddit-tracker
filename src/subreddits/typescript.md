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
## [2][Marshal.ts serializer got a JIT engine + ORM abstraction. It's now up to 300x faster than class-transformer. Time to switch guys](https://www.reddit.com/r/typescript/comments/exzdsd/marshalts_serializer_got_a_jit_engine_orm/)
- url: https://github.com/marcj/marshal.ts#marshalts
---

## [3][Question about Typescript objects assign](https://www.reddit.com/r/typescript/comments/ey4vbb/question_about_typescript_objects_assign/)
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
## [4][What's been your experience using TypeScript with React Hooks?](https://www.reddit.com/r/typescript/comments/exwqq2/whats_been_your_experience_using_typescript_with/)
- url: https://www.reddit.com/r/typescript/comments/exwqq2/whats_been_your_experience_using_typescript_with/
---
I've been working on a project involving both React and TypeScript. I've decided to start moving the project in the direction of hooks. What I didn't realize when starting this process was that React hooks provide a way to work in React without having to use classes. This kind of bummed me out since I've been learning a lot about TypeScript's added class features like modifying class properties to be private / public / protected, etc.

This got me thinking about using React and TypeScript together, and about how Facebook has also made Flow, which is kind of Facebook's analog to TypeScript (at least in terms of type declarations). 

Are React and TypeScript at odds with each other? Do you think these two technologies are moving away from each other? Would you recommend using them together? Why/why not?
## [5][Could I get a quick code review on these three simple functions?](https://www.reddit.com/r/typescript/comments/exzdjn/could_i_get_a_quick_code_review_on_these_three/)
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

Thank you.
## [6][What is the VScode plugin Anders Hejlsberg is using to get the types underlined?](https://www.reddit.com/r/typescript/comments/exkh20/what_is_the_vscode_plugin_anders_hejlsberg_is/)
- url: https://i.redd.it/o780xxzgige41.png
---

## [7][Inversify based dependency injection in React - create-react-app-inversify](https://www.reddit.com/r/typescript/comments/exqmz8/inversify_based_dependency_injection_in_react/)
- url: https://www.reddit.com/r/typescript/comments/exqmz8/inversify_based_dependency_injection_in_react/
---
 Hi All

I've created a few packages to help with constructor (or function) based dependency injection using inversify in React. Feedback or comments would be welcome.

I understand that this isn't really the way that react apps are usually written so I don't expect many react developers to use or even to like this!! I am an angular developer but want to build some react apps. I don't want to give up the ease of testing that you get with constructor based dependency injection so I created this.

As I said comments and constructive feedback welcome.

npm package:

[https://www.npmjs.com/package/inversify-constructor-injection](https://www.npmjs.com/package/inversify-constructor-injection)

to create a new app:

`npx create-react-app-inversify my-app`
## [8][How can I reference a string literal in an object created using 'as const' in a callback?](https://www.reddit.com/r/typescript/comments/exs8kq/how_can_i_reference_a_string_literal_in_an_object/)
- url: https://www.reddit.com/r/typescript/comments/exs8kq/how_can_i_reference_a_string_literal_in_an_object/
---
I have an object, here is a sample of it. I use `as const`, and now the keys and values are no longer type `string`, but the string literal:
```
readonly additional: readonly ["no-console"];
```
```
export const rules = {
  add: {},
  remove: {
    additional: ['no-console', 'another-rule']
  }
} as const;
```
In my function I can no longer type the array contents as `string` as it is now a tuple, so how do i  type them as the correct reference to the object rules... `rules` is not an interface, how do i get at the types in this context?
```
import { rules } from './rules.ts';

interface EslintRules {
  [key: string]: any[];
}

export const conditions = (key: string, val: any[]): boolean =&gt;
  !!(
    !rules.remove.basicPrettierConflicts.includes(key) &amp;&amp; // &lt;&lt; key is not a string, it needs to refer to the actual string literal
  );

export function ruleFilter(
  esLintRules: EslintRules,
  conditionToAccept: { (key: string, val: any[]): boolean }
): [EslintRules, string[]] {
  const removedRules = [];
  return [
    Object.fromEntries(
      Object.entries(esLintRules).filter(([key, val]) =&gt; {
        if (conditionToAccept(key, val)) {
          return true;
        }
        removedRules.push(key);
        return false;
      })
    ),
    removedRules
  ];
}
```
## [9][Railroad Programming in TypeScript](https://www.reddit.com/r/typescript/comments/exoctl/railroad_programming_in_typescript/)
- url: https://medium.com//railroad-programming-in-typescript-21d69f486f6e?source=friends_link&amp;sk=07669f7de4932860744c4eb267bad179
---

## [10][How to map variable length tuple argument to a second variable length tuple argument with different item types?](https://www.reddit.com/r/typescript/comments/exru7o/how_to_map_variable_length_tuple_argument_to_a/)
- url: https://www.reddit.com/r/typescript/comments/exru7o/how_to_map_variable_length_tuple_argument_to_a/
---
I want to write a function 'foo' that meets these requirements:

    interface idToDataType {
      'id1': string,
      'id2': number,
      'id3': boolean,
    }
    
    // Compiles
    foo(['id2', 'id1', 'id3'], [6, 'happy', true]);
    foo(['id1', 'id2'], ['happy', 6]);
    foo(['id2', 'id2'], [6, 7]);
    
    // Should not compile!
    foo(['id1', 'id2'], ['sad']); // Data too short!
    foo(['id1', 'id2'], ['sad', 'sad']); // wrong mapped tuple type!
    foo(['id1', 'id2'], ['sad', 6, 'sad']); // Data too long

So far the closest I've gotten is here:  
[http://www.typescriptlang.org/play/#code/FDAuE8AcFMAIAUBOB7GiIBUpwLywNbTjIBmCKaEAagIYA2ArtFjALI2QDcY25q06cLUbNs7SLDwBvYLFgByAJYATAIzyAXLADOoRIoB2AcwA0shSoBMm2AYYBbAEYCzcpcoDMNx8mR1oNAZmAL4gEDCwAJLKGMgAIjSgNAA8AIKw0AAeoNAGytqwqYiINODJSPyCLNAAfDWSsDJyANqRsIYERKSFALpaqa09Gdm5+XyU4NWwAPzjAtT0TNXizQORPUMa5nI7u3v7B4dHBtAAbgLcocDK0ADGdDSIcCQMBregisgGsDm6yRjDHJ5AoVCbVZo9GoACm2sGUiRoWmisQSSX+NRMsAAdDiVNotBgAJRaU7IFTcYAAekpsAAwsh7JBFP5tGBoLooc0AGyY+QACw4kHA8kxeiYPV5VhFFjU0vcXkJ3F+oE5-MFwsxXIlMvUkuU1kVbI53MxAHZte5rHqDRTqbAAMp85AMOjKWzIUCwW4Mpn+ACERpVzXk2hoynkFpUuplNtgdtRNB+vh0TvQAeVqtD4d5WYjeujlvkirjNIA7ihjLB7ILoG7QAxIP4fth0+ygyGw9KeQpc5HZdai5wS7AE0nkLA6F8jCAqTSKo5-PYCqXFKA+T8+YoCoKUDRbnytnbVFiRwjYIE3dFt092t813By1PYMhEDdEH7T0kU87XV6GXBtDuL5lBPDA+TgGgfHOdoCjoRRCFgYh4R+eh8C0IwnlyDc4AwmhtAKLdMRQBgjHXe9n1uAJvkQBwF3ZWdYEsUDwPaMYBWg0BxwYQCFCeXR5HPfBExxLFMRfDI6B48jwnZc8b1eT4TjdC9sNgF43g+L5YGcb17Fkzjx38AxFFyUAsSAA](http://www.typescriptlang.org/play/#code/FDAuE8AcFMAIAUBOB7GiIBUpwLywNbTjIBmCKaEAagIYA2ArtFjALI2QDcY25q06cLUbNs7SLDwBvYLFgByAJYATAIzyAXLADOoRIoB2AcwA0shSoBMm2AYYBbAEYCzcpcoDMNx8mR1oNAZmAL4gEDCwAJLKGMgAIjSgNAA8AIKw0AAeoNAGytqwqYiINODJSPyCLNAAfDWSsDJyANqRsIYERKSFALpaqa09Gdm5+XyU4NWwAPzjAtT0TNXizQORPUMa5nI7u3v7B4dHBtAAbgLcocDK0ADGdDSIcCQMBregisgGsDm6yRjDHJ5AoVCbVZo9GoACm2sGUiRoWmisQSSX+NRMsAAdDiVNotBgAJRaU7IFTcYAAekpsAAwsh7JBFP5tGBoLooc0AGyY+QACw4kHA8kxeiYPV5VhFFjU0vcXkJ3F+oE5-MFwsxXIlMvUkuU1kVbI53MxAHZte5rHqDRTqbAAMp85AMOjKWzIUCwW4Mpn+ACERpVzXk2hoynkFpUuplNtgdtRNB+vh0TvQAeVqtD4d5WYjeujlvkirjNIA7ihjLB7ILoG7QAxIP4fth0+ygyGw9KeQpc5HZdai5wS7AE0nkLA6F8jCAqTSKo5-PYCqXFKA+T8+YoCoKUDRbnytnbVFiRwjYIE3dFt092t813By1PYMhEDdEH7T0kU87XV6GXBtDuL5lBPDA+TgGgfHOdoCjoRRCFgYh4R+eh8C0IwnlyDc4AwmhtAKLdMRQBgjHXe9n1uAJvkQBwF3ZWdYEsUDwPaMYBWg0BxwYQCFCeXR5HPfBExxLFMRfDI6B48jwnZc8b1eT4TjdC9sNgF43g+L5YGcb17Fkzjx38AxFFyUAsSAA)

I've also listed the problems with the approach in the playground link.

&amp;#x200B;

Does anyone have any better ideas? I'd really appreciate the help, as I have already spent hours learning about index types, mapped types, tuples, etc and now I'm stuck with this working-but-subpar approach.

And yes, I do want this function. I do indeed need a pair of tuple args, and I do not want an array of tuple pairs, or a map. This function is interfacing with a very high performance piece of software that operates on a Structure of Arrays rather than an Array of Structures, and passing in a pair of arrays is much faster than passing in an array of pairs. In my testing, it was between 2x (Small arrays) to 30x faster (large arrays).

Thanks in advance!
## [11][Using paths property of tsconfig to specify a specific file?](https://www.reddit.com/r/typescript/comments/exjk6d/using_paths_property_of_tsconfig_to_specify_a/)
- url: https://www.reddit.com/r/typescript/comments/exjk6d/using_paths_property_of_tsconfig_to_specify_a/
---
Basically i want to be able to do something like this:

    "paths": { "@constants": ["src/constants.ts"] }

So i don't have to deal with this mess:

    import { FOO } from '../../../../constants'

Instead i would rather have:

    import { FOO } from '@constants'

I am on TS 3.7, and unfortunately doing the above results in a compilation error indicating it cannot find the file. Which doesn't make much sense to me. If i remove the '@constants\` from the path it works fine.

I also know i can do something like this:

    "paths": { "@app": ["src/*"] }
    
    import { FOO } from '@app/constants'

However this makes Webstorm go completely dumb, as it will ignore all other paths, and instead focus solely upon that single path entry and ignore the rest. Guessing it is because it considers all files now under that path instead of caring about any of the others.
