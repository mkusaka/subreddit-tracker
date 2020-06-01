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
## [2][Do you guys use eslint on your typescript projects?](https://www.reddit.com/r/typescript/comments/guiovl/do_you_guys_use_eslint_on_your_typescript_projects/)
- url: https://www.reddit.com/r/typescript/comments/guiovl/do_you_guys_use_eslint_on_your_typescript_projects/
---
I used to use eslint with the plugins it suggested to me to install after executing :

    npx eslint --init;

and chooseing the option : I use typescript  , but I have found it to produce some akward linting errors recently like not allowing me to put type `any` on a function parameter. 

I have stopped using eslint since then. 

If you use eslint then what you use it for?

I was using it just because it was linting errors while I was writing js and so I believed it will also be useful in typescript.
## [3][Are there any online courses that covers TS with JS concepts?](https://www.reddit.com/r/typescript/comments/guehwr/are_there_any_online_courses_that_covers_ts_with/)
- url: https://www.reddit.com/r/typescript/comments/guehwr/are_there_any_online_courses_that_covers_ts_with/
---
I am interested in learning TS, but since it builds to JS, I would also need to learn JS. 

Are there any courses that covers TS entirely such that i won't be needing to learn JS before it?
## [4][Can enums be treated as numbers for incrementation purposes?](https://www.reddit.com/r/typescript/comments/gugq66/can_enums_be_treated_as_numbers_for/)
- url: https://www.reddit.com/r/typescript/comments/gugq66/can_enums_be_treated_as_numbers_for/
---
By default I believe enums are assigned numbers starting from 0. I was thinking of using an enum to define a series of steps for a CLI script. that way I could name each step, for example, "intro", "gatherName", "configureGoogle" etc. There are no branching paths so in this setup, reaching the final enum would signal that the next step should perform an exit procedure.

Does it ever make sense to increment enums? Are there side effects to consider?

If this should not be done, I suppose additional logic will be needed to define step pathing. I'm open to any ideas you guys have. The first idea I have is to iterate over the enumb definition, push each enum into an array to create a sequence, and then allow the indexes to become the sequence values.
## [5][Solving Riddles with 0 lines of program | Using typescript types](https://www.reddit.com/r/typescript/comments/gu390u/solving_riddles_with_0_lines_of_program_using/)
- url: https://medium.com/@damodharanjay/solving-riddles-with-0-lines-of-program-b23994072dd0
---

## [6][Typescript in React, tough example to understand](https://www.reddit.com/r/typescript/comments/gubcvg/typescript_in_react_tough_example_to_understand/)
- url: https://www.reddit.com/r/typescript/comments/gubcvg/typescript_in_react_tough_example_to_understand/
---
&amp;#x200B;

    const Form = &lt;T extends {}&gt;({values, children}: FormProps&lt;T&gt;) =&gt; {
        return children(values);
    };
    
    const App: React.FC = () =&gt; {
        return {
            &lt;div&gt;
    	    &lt;Form &lt;{lastName:string|null}&gt; values={{lastName: ""}}&gt;
    		{values =&gt; &lt;div&gt;{values.lastName}&lt;/div&gt;}
    	    &lt;/Form&gt;
    	&lt;/div&gt;
        }
    }

My interpretation (as best as I can):

1. The generic T is extending an empty object, meaning it can take any properties
2. He must have an interface `FormProps` somewhere else. I have a hard time reading what is going on exactly inside the parameter definitions. Destructuring of the props object? Or something else? `FormProps` can't be a return type because `:` should go outside the `()`
3. Inside the JSX portion, he has a `&lt;Form&gt;` component. The component is typing the `lastName` prop to a string or null.
4. The values prop appears to be escaping JSX and passing an object with `lastName` set to empty string by default, not sure why
5. Finally on the last line JSX is escaped and on every render an arrow function prints the value of `values.lastName` to the div if not null.

I know my assessment above is wrong, I appreciate any feedback to help me read this better.
## [7][What's the return type of a static method that returns new instances of its class?](https://www.reddit.com/r/typescript/comments/gui766/whats_the_return_type_of_a_static_method_that/)
- url: https://www.reddit.com/r/typescript/comments/gui766/whats_the_return_type_of_a_static_method_that/
---
Check this first:
```
class A {
    static someStaticMethod(): any[] {
        return [new this(), new this()]
    }
}

class B extends A {
    someClassProp = 0
}

console.log(B.someStaticMethod())
```
1. What is the actual return type of `someStaticMethod` beyond `any[]`? How do I express the class which was used to call `someStaticMethod`? In this example it's `B` and the static's return type something like `new B()[]` but the `B` should be a generic. Do I use generics here? If yes, how?

2. Why does the TS compiler does not complain about the `any` type despite having all options set to strict:

Code sandbox: https://www.typescriptlang.org/play?#code/MYGwhgzhAECC0G8BQ1XQgFzBglsdA9gLYCmAylrsALIkYAWBAJgBQCUAXNGAHYCeAbQC6iFGnEAnOgFcJPaAJ4kA7tAY4I7ADTQlq9ZrZCxqAL5JzSUJBgAhaCQAeGEjyYx4ycRGIkAwuBQAAoSBAAO0AC80AAMFkhWBDw+ICQAdCAEAOYstmk+pBTYeLQMzOxsQA
## [8][Programmer without JS experience: where to start with TS?](https://www.reddit.com/r/typescript/comments/gui6ai/programmer_without_js_experience_where_to_start/)
- url: https://www.reddit.com/r/typescript/comments/gui6ai/programmer_without_js_experience_where_to_start/
---
Hey lads,

I've been bid to learn TS for my job, which revolves around Amazon Web Services. The new AWS "CDK" (cloud dev kit) has by far the best support for TS, so there's no getting around it. 

I've got multiple years of exp working with primarily Java and a bit of python, but none so far with web languages like JavaScript. 

Where do I start with learning TS, assuming my main goal is to write architectures for AWS (so backend stuff)? Do I just start hacking away, do I dive into a Typescript tutorial, or should I start out reading up on JavaScript fundamentals before starting with TS? 

Appreciate the help!
## [9][How to get TypeScript to stop throwing an error when accessing a field on an object that has been determined to be an instance of a particular object](https://www.reddit.com/r/typescript/comments/gublkg/how_to_get_typescript_to_stop_throwing_an_error/)
- url: https://www.reddit.com/r/typescript/comments/gublkg/how_to_get_typescript_to_stop_throwing_an_error/
---
I've been writing some code, and let's say I have an array of several objects; e.g.

    const ar:(A | B | C)[] = [new A(), new B(), new C()];

where the A object is the only one of them that has a method called `match`.

and then I have a for loop that iterates through it

    for (let e of ar) {
      if (e instanceof A) {
        if (a.match('hello')) {
          print('good');
        }
      }
    }

But I keep getting an error that says something along the lines of `Property 'match' does not exist on type 'A | B | C'. Property 'match' does not exist on types 'B | C'.`

How should I resolve this, seeing as the actual error would not take place as I've already confirmed that it is the proper object as thus the object would have said field.

Any help is appreciated.

Thanks!
## [10][Typing for higher-order components](https://www.reddit.com/r/typescript/comments/gu657l/typing_for_higherorder_components/)
- url: https://www.reddit.com/r/typescript/comments/gu657l/typing_for_higherorder_components/
---
I'm trying to add typing for a higher-order component that takes in `mapStateToProps` and `mapDispatchToProps` functions to inject into a React component, but I can't seem to get generics to work. I've simplified my problem down to this example ([playground link](https://www.typescriptlang.org/play?#code/JYOwLgpgTgZghgYwgAgCoCE4GcIGkICeWyA3gLABQy1yARthAFzIgCuAtrdANyUC+lSggD2ILGGTtoAczyFiAXmQAeVAEEAJhogb8RZBAAekEBuLDaAKwgIwAPgAUAByjCAbsG1RmDgJTIFOzRNbV15f0CVShpggBtYvWIjEzNgrR1E5AAyNEwcRMpHaJoYEB8Aa3lmdXjEiKC3YU9KetJi6hExCTh0sP0lF3dPaD9eKhpO8ToGROq8uX7SaZxmABYAJmQ+MZioCDBWKBBkUocSZAA6K-p8+QAaS6ue0My+XzHtwQpJiVLUAAs4GBUHBKlgALIyDLyALIBzFc43CAPQYeUJbZiIhjMNicHjIVHDDTMcRQUDSLatSbCWIQC6xYTSBxI5AAamQACJOWyCa40Tp3l8pFBZIkHH4AkEzryhqFmByYMJhFy3r4HH9AcDQRAIVC+lh3kA)):

    interface TBaseKeys {
        base: number;
    }
    
    const mergeKeys = &lt;TAddedKeys extends object&gt;(provider: () =&gt; TAddedKeys) =&gt; &lt;
        TAllKeys extends TAddedKeys &amp; TBaseKeys
    &gt;(
        fn: (keys: TAllKeys) =&gt; void
    ) =&gt; {
        const addedKeys = provider();
        const baseKeys: TBaseKeys = { base: 42 };

        // Fails with:
        // Argument of type '{ base: number; } &amp; TAddedKeys' is not assignable to parameter of type 'TAllKeys'.
        //  '{ base: number; } &amp; TAddedKeys' is assignable to the constraint of type 'TAllKeys',
        //    but 'TAllKeys' could be instantiated with a different subtype of constraint 'object &amp; TBaseKeys'.(2345)
        return fn({ ...baseKeys, ...addedKeys });
    };

In the code above, `fn` takes in an object, where some of the keys are hardcoded (`TBaseKeys`), while others are generic (`TAddedKeys`, which you get from `provider()`). This is how `mergeKeys` could be used:

    const fnThatTakesMergedKeys = (
      { base, provided }: { base: number; provided: string }) =&gt; console.log(base + " " + provided);
    
    mergeKeys(() =&gt; ({ provided: "foo" }))(fnThatTakesMergedKeys);

Console output: `42 foo`

Does anyone have tips on resolving the TypeScript error above?
## [11][[NOOB QUESTION] How is typescript better if it is just compiled to javascript?](https://www.reddit.com/r/typescript/comments/gttwke/noob_question_how_is_typescript_better_if_it_is/)
- url: https://www.reddit.com/r/typescript/comments/gttwke/noob_question_how_is_typescript_better_if_it_is/
---
Since it is compiled to JS, the performance cannot be better than a natively written JS. 

Besides static type checking, what benefits do TS offer? I don't think it offers a performance boost. I'm confused if I should port all my JS code to TS.
