# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Return type of a function that takes any number of functions as argument, and return the array of returned values](https://www.reddit.com/r/typescript/comments/gi84kp/return_type_of_a_function_that_takes_any_number/)
- url: https://www.reddit.com/r/typescript/comments/gi84kp/return_type_of_a_function_that_takes_any_number/
---
Not sure if my title makes sense, here is my problem:

I have a function that runs multiple functions concurrently, and returns the array of results (ignore the absence of promises in what follows, I'm using coroutines).

    parrallel(() =&gt; 'a', () =&gt; 1, () =&gt; [0, 1, 2]);
    // Returns ['a', 1, [0, 1, 2]]

As of now, I'm using the following type declaration that handles calling this function with up to 7 arguments while keeping type-safety:

    type Fn&lt;T&gt; = () =&gt; T;
    export function parallel&lt;T&gt;(fcts: Fn&lt;T&gt;[]): T[]; 
    export function parallel&lt;T&gt;(fcts: [Fn&lt;T&gt;]): [T]; 
    export function parallel&lt;T, U&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;]): [T, U]; 
    export function parallel&lt;T, U, V&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;, Fn&lt;V&gt;]): [T, U, V]; 
    export function parallel&lt;T, U, V, W&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;, Fn&lt;V&gt;, Fn&lt;W&gt;]): [T, U, V, W]; 
    export function parallel&lt;T, U, V, W, X&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;, Fn&lt;V&gt;, Fn&lt;W&gt;, Fn&lt;X&gt;]): [T, U, V, W, X]; 
    export function parallel&lt;T, U, V, W, X, Y&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;, Fn&lt;V&gt;, Fn&lt;W&gt;, Fn&lt;X&gt;, Fn&lt;Y&gt;]): [T, U, V, W, X, Y]; 
    export function parallel&lt;T, U, V, W, X, Y, Z&gt;(fcts: [Fn&lt;T&gt;, Fn&lt;U&gt;, Fn&lt;V&gt;, Fn&lt;W&gt;, Fn&lt;X&gt;, Fn&lt;Y&gt;, Fn&lt;Z&gt;]): [T, U, V, W, X, Y, Z];

I did not find any way to find a type declaration that work for any number of arguments, is it even possible ?
## [3][Get a better grip on TS at JSNation, the largest remote JavaScript conf](https://www.reddit.com/r/typescript/comments/giau0e/get_a_better_grip_on_ts_at_jsnation_the_largest/)
- url: https://sfree.life/jsnation-javascript-live-conference-2020-free/
---

## [4][Polymorphic TypeScript - Function overloading with rest parameters](https://www.reddit.com/r/typescript/comments/gi7ljb/polymorphic_typescript_function_overloading_with/)
- url: https://tane.dev/2020/05/polymorphic-typescript-function-overloading-with-rest-parameters/
---

## [5][Prototope - TailwindCSS-inspired CSS-in-JS library](https://www.reddit.com/r/typescript/comments/gi5zen/prototope_tailwindcssinspired_cssinjs_library/)
- url: https://github.com/Isotope-js/isotope/tree/master/packages/prototope
---

## [6][Whats the benefit of declaring a functions argument types as ({arg1, arg2} : {arg1: string, arg2: number}) rather than just ({arg1: string, arg2: number})](https://www.reddit.com/r/typescript/comments/ghwggh/whats_the_benefit_of_declaring_a_functions/)
- url: https://www.reddit.com/r/typescript/comments/ghwggh/whats_the_benefit_of_declaring_a_functions/
---
Why did the typescript programmers add this complexity to the language?In my opinion, being able to declare a function like

`const func = ({arg1: string, arg2: number}) =&gt; {}`

is way cleaner and more intuitive, so why would they make it the default that to declare a functions argument types you would have to do something like

`const func = ({arg1, arg2} : {arg1: string, arg2: number}) =&gt; {}`
## [7][mockzilla: Yet another mocking library for TypeScript and jest.](https://www.reddit.com/r/typescript/comments/ghv6q5/mockzilla_yet_another_mocking_library_for/)
- url: https://www.reddit.com/r/typescript/comments/ghv6q5/mockzilla_yet_another_mocking_library_for/
---
I'm well aware that there are many mocking libraries out there, but I went ahead and created another one (for use with jest): [https://lusito.github.io/mockzilla/](https://lusito.github.io/mockzilla/)

My main focus was to be able to easily write and verify mocks of a deeply nested API like the webextensions browser object.

In webextensions you would have code like this one:

    browser.webRequest.onBeforeRedirect.addListener(callback, filter);

But you don't have access to an actual implementation of this API, which is why some libraries already didn't work for my use-case. Other libraries would require a lot of typing and I loathe that.

With [mockzilla](https://lusito.github.io/mockzilla), I can now just write a simple line like this:

    mockBrowser.webRequest.onBeforeRedirect.addListener.expect(listener, expect.anything());

This contains types for auto-completion and type-safety. It is auto-verified, so no need to call verify() at the end of a test.

Another possibility is to assimilate an actual instance to become partially ~~borg~~ mock:

    const myInstance = new MyClass();
    const mock = mockAssimilate(myInstance, "myInstance", {
        mock: ["runA"],
        whitelist: ["run", "someProp"],
    });
    mock.runA.expect(expect.anything(), true).andReturn(true);

mockAssimilate works with types as well, even on private methods.

There are a few more things included. Take a look at the documentation.

I am still working on the API design, looking for some feedback and ideas.
## [8][Castellated: An Adaptable, Robust Password Storage System for Node.js](https://www.reddit.com/r/typescript/comments/ghx52r/castellated_an_adaptable_robust_password_storage/)
- url: http://www.wumpus-cave.net/2020/05/11/castellated-an-adaptable-robust-password-storage-system-for-node-js/
---

## [9][Question about the Pick utility type](https://www.reddit.com/r/typescript/comments/ghwpyr/question_about_the_pick_utility_type/)
- url: https://www.reddit.com/r/typescript/comments/ghwpyr/question_about_the_pick_utility_type/
---
I've been using utility types like [Pick](https://www.typescriptlang.org/docs/handbook/utility-types.html#picktk) and [Omit](https://www.typescriptlang.org/docs/handbook/utility-types.html#omittk) for some time now to reduce type repetition. I just stumbled across a scenario that struck me as odd though and would appreciate some explanation.

Say I have an interface like so:

    interface ComponentState {
       user?: User;
       loading: boolean;
       error: Error | null;
    }

And then I want to define a discriminated union that references properties from that interface:

    type ComponentActionTypes = 
    | {
        type: ComponentActions.SET_USER,
        payload: Pick&lt;ComponentState, 'user'&gt; //I would expect User | undefined
      }
    | {
       type: ComponentActions.SET_LOADING,
       payload: Pick&lt;ComponentState, 'loading'&gt; //I would expect boolean
      }
    | {
       type: ComponentActions.SET_ERROR,
       payload: Pick&lt;ComponentState, 'error'&gt; //I would expect Error | null
      }

I would expect the `payload` in each of the above `ComponentActionTypes` scenarios to conform to whatever the property type is for the `ComponentState` interface but the compiler interprets each of the payloads as objects with a key being that of the respective interface property and the value being the respective type. So the `payload` for `ComponentActionTypes` type `ComponentActions.SET_USER` is actually

    { user: User | undefined } //I want action.payload to be User | undefined

Clearly I misunderstood how `Pick` works. Any suggestions on how I can achieve what I'm after without having to just replicate each of the types defined in the `ComponentState` interface?

&amp;#x200B;

Much appreciated
## [10][Make a defined interface property optional](https://www.reddit.com/r/typescript/comments/ghwcjb/make_a_defined_interface_property_optional/)
- url: https://www.reddit.com/r/typescript/comments/ghwcjb/make_a_defined_interface_property_optional/
---
I have type:

\`\`\`ts  
interface INameType {  
   name: string;  
   typeId: number  
   ...

}

\`\`\`

And I want to have a second type which takes all of the first interface properties but makes the ones I specify optional, like: 

\`\`\`ts  
interface INameType {  
   name: string;  
   typeId?: number  
   ...

}

\`\`\`  
How can I accomplish this? Thank you
## [11][TypeScript alias for asserting value type for keyof](https://www.reddit.com/r/typescript/comments/ghx4tg/typescript_alias_for_asserting_value_type_for/)
- url: https://stackoverflow.com/questions/61739893/typescript-alias-for-asserting-value-type-for-keyof
---

