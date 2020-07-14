# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Cool helper type that will allow you to recursively replace all types in an object with another type.](https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/)
- url: https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/
---
Figured this out recently, when I had to transform large plain objects of different shapes, and replace some leaf value with a different one. `Replaced` type takes 4 generic arguments:

* `T` the object to be transformed.
* `TReplace` the type to be replaced with
* `TWith` the type to use instead of `TReplace`
* `TKeep` the type to not transform. Defaults to primitive values.  


Obviously to be used as the return type of some function that does the actual transformation.

    type Primitive = string | number | bigint | boolean | null | undefined;
    
    type Replaced&lt;T, TReplace, TWith, TKeep = Primitive&gt; = T extends TReplace | TKeep
        ? (T extends TReplace
            ? TWith | Exclude&lt;T, TReplace&gt;
            : T)
        : {
            [P in keyof T]: Replaced&lt;T[P], TReplace, TWith, TKeep&gt;
        }
    
    type Foo = symbol;
    type Bar = "3";
    
    type ToReplace = {
        x?: {
            y: Foo | number;
            z: string;
        }[]
    } | undefined;
    
    type WasReplaced = Replaced&lt;ToReplace, Foo, Bar&gt;;
    /* Output:
    {
        x?: {
            y: number | "3";
            z: string;
        }[];
    }
    */
## [3][I made a Snapchat clone in the browser!](https://www.reddit.com/r/typescript/comments/hqfzmn/i_made_a_snapchat_clone_in_the_browser/)
- url: https://v.redd.it/gnbkqf8duma51
---

## [4][Discord Community: Code Review, Paired Programming, Collab, and Live Instruction for TS Devs!](https://www.reddit.com/r/typescript/comments/hqw7ts/discord_community_code_review_paired_programming/)
- url: https://www.reddit.com/r/typescript/comments/hqw7ts/discord_community_code_review_paired_programming/
---
 A community of student and professional developers. Join us for free instruction, paired programming, code/project review, collaboration, and more. Looking for exposure to new concepts? Engage in peer review on your projects? Looking for help getting through your latest assignment or project? **Looking to spread the "Good News" and show the rest of us Typescript isn't just good for Angular?** If you can ask a good question you're like to find a good answer here!

Engage in the Feinberg method of learning, where you reinforce what you do know and spot weaknesses in your own studies, by teaching a few code-along courses or just share your latest project.

**In short:** We'd love to have you. Come on in!

[https://discord.gg/5vphZ5s](https://discord.gg/5vphZ5s)
## [5][Create a library with only type definitions](https://www.reddit.com/r/typescript/comments/hqnc8g/create_a_library_with_only_type_definitions/)
- url: https://www.reddit.com/r/typescript/comments/hqnc8g/create_a_library_with_only_type_definitions/
---
Hello, I'm a novice typescript developer that is loving the language so far but is having some issues understanding how types definitions works.

What I'm trying to achieve is creating a library that I would like to publish in a github repo and that should contains type definitions shared between multiple projects.   


What's the best approach to achieve something similar?
## [6][ESLint TS-specific errors (interfaces/enums "not used")](https://www.reddit.com/r/typescript/comments/hqlv93/eslint_tsspecific_errors_interfacesenums_not_used/)
- url: https://www.reddit.com/r/typescript/comments/hqlv93/eslint_tsspecific_errors_interfacesenums_not_used/
---
    import ConfigData from '../setupWizard/ConfigDataInterface';
    
    // (alias) interface ConfigData
    // import ConfigData
    // 'ConfigData' is defined but never used.eslintno-unused-vars

When removed the interface references become undefined as one would expect.

&amp;#x200B;

    export enum translationDirection { toForeign, toEnglish }
    // (alias) interface ConfigData
    // import ConfigData
    // 'ConfigData' is defined but never used.eslintno-unused-vars
    // same for toEnglish
    
    export enum voiceGender { male = 'MALE', female = 'FEMALE' }
    
    // (enum member) voiceGender.male = "MALE"
    'male' is assigned a value but never used.eslintno-unused-vars
    // (enum member) voiceGender.female = "FEMALE"
    'female' is assigned a value but never used.eslintno-unused-vars

There is another enum defined inside of a method in the consuming file, even that one throws the lint/error for "defined but not used".

&amp;#x200B;

Anyone know how to fix these? The error fixing page seems JS-specific: [https://eslint.org/docs/rules/no-unused-vars](https://eslint.org/docs/rules/no-unused-vars)

    // .eslintrc.js
    module.exports = {
      env: {
        browser: false,
        es2020: true,
      },
      extends: [
        'airbnb-base',
      ],
      parser: '@typescript-eslint/parser',
      parserOptions: {
        ecmaVersion: 11,
        sourceType: 'module',
      },
      plugins: [
        '@typescript-eslint',
      ],
      rules: {
        'import/extensions': 0, // off
      },
      settings: {
        'import/resolver': {
          node: {
            extensions: ['.ts'],
            moduleDirectory: ['node_modules', 'src'],
          },
        },
      },
    };
## [7][Am I being pedantic when it comes to helper methods in classes?](https://www.reddit.com/r/typescript/comments/hqar0k/am_i_being_pedantic_when_it_comes_to_helper/)
- url: https://www.reddit.com/r/typescript/comments/hqar0k/am_i_being_pedantic_when_it_comes_to_helper/
---
For the most part, I try to program in a functional style - simple, reusable functions, accessing state, where needed, via closure, etc.  And certainly \*anything\* you can write as a class in JS/TS can be written as functions.  But sometimes classes are a simpler, easier solution and I find that when you have complex state to manage, they're nice to use. (Basically, when I start thinking that something will require a monad, I think it might be okay to switch to classes. :) ) 

The question occurs when I have a class definition that Is more or less useless outside the class, but is purely functional - that is, no access to state, no mutations. 

Now, I would insist these functions be defined \*outside the class\* so that they can be independently unit tested.  (Though to keep concerns together, I'd put them in the same \*file\* where the class is the default export).  But is that being pedantic?    


In other words, I'm refactoring this: 

```typescript
export default class Match {
  private checkCondition = (condition: Condition, input: any): boolean =&gt; {
    if (typeof condition === "function") {
      return condition(input);
    }
    return condition === input;
  };
}
```

to this:
```typescript
export const checkCondition = (condition: Condition, input: any): boolean =&gt; {
  if (typeof condition === "function") {
    return condition(input);
  }
  return condition === input;
};

export default class Match {
  /* OPTIONAL!
  private checkCondition = checkCondition;
  */
}
```
## [8][New to TypeScript. How do you know what types to use for things other than strings and booleans?](https://www.reddit.com/r/typescript/comments/hqty7o/new_to_typescript_how_do_you_know_what_types_to/)
- url: https://www.reddit.com/r/typescript/comments/hqty7o/new_to_typescript_how_do_you_know_what_types_to/
---
I've never been able to figure this out. For complex things like array, objects and other components what type do you give them?
## [9][Express.js course with TypeScript Lesson 2 — Apollo &amp; WebSockets](https://www.reddit.com/r/typescript/comments/hqcfqa/expressjs_course_with_typescript_lesson_2_apollo/)
- url: https://medium.com//express-js-course-with-typescript-lesson-2-apollo-websockets-7eb2063186bf?source=friends_link&amp;sk=123f4b00527f5080717f1cbfa6f34ac2
---

## [10][Access mapped type by index](https://www.reddit.com/r/typescript/comments/hqc4h2/access_mapped_type_by_index/)
- url: https://www.reddit.com/r/typescript/comments/hqc4h2/access_mapped_type_by_index/
---
I'm trying to wrap my head around the following behavior, say I have the following type definitions

    type RMap = {
        a: { foo: string }
        b: { foo: unknown }
    }
    
    type RType = keyof RMap
    
    type RState = {
        [k: string]: {
            s: RMap[RType]
        }
    }

Given these, what's the difference between f1 and f2 below? Why does the first one not compile?

    declare const map: RState
    
    function f1&lt;R extends RType&gt;(k: string, f: (r: RMap[R]) =&gt; void): void {
        f(map[k].s)
    }
    
    function f2(k: string, f: (r: RMap[RType]) =&gt; void): void {
        f(map[k].s)
    }
## [11][Having a hard time Typing a ref that is a function that gets assigned to ref.curent](https://www.reddit.com/r/typescript/comments/hq8tnc/having_a_hard_time_typing_a_ref_that_is_a/)
- url: https://www.reddit.com/r/typescript/comments/hq8tnc/having_a_hard_time_typing_a_ref_that_is_a/
---
Code:

    const memoizedHandler = useCallback((e) =&gt; handler(e), dependencies);useEventListener({ eventName: 'click', handler: memoizedHandler }); // &lt;-- error here
    ERROR:Error:(17, 42) TS2741: Property 'current' is missing in type '(e: any) =&gt; void' but required in type 'RefObject&lt;HTMLDivElement&gt;'.

I am assigning to "current", WHY do I need it on the function I am trying to assign to it?

which says it is originating from:

    useEffect(() =&gt; {
      savedHandler.current = handler;
    }, [handler]);
    
    useEffect(() =&gt; {
      const isSupported = element &amp;&amp; element.addEventListener;
      if (!isSupported) return;
    
      const eventListener = (event: React.ChangeEvent&lt;HTMLDivElement&gt;) =&gt; 
      savedHandler?.current?.(event);
    
    // error at "current" says TS2349: This expression is not callable.
    // Type 'RefObject&lt;HTMLDivElement&gt;' has no call signatures

getting so frustrated. lol
