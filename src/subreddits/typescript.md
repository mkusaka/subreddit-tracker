# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Drash - A microframework for deno](https://www.reddit.com/r/typescript/comments/esnv1t/drash_a_microframework_for_deno/)
- url: https://www.reddit.com/r/typescript/comments/esnv1t/drash_a_microframework_for_deno/
---
Would love to get feedback on this:

[https://www.reddit.com/r/Deno/comments/esazl4/drash\_a\_microframework\_for\_deno/](https://www.reddit.com/r/Deno/comments/esazl4/drash_a_microframework_for_deno/)

Thanks!
## [3][Unexpected Lessons from 100% Test Coverage](https://www.reddit.com/r/typescript/comments/eso0pv/unexpected_lessons_from_100_test_coverage/)
- url: https://medium.com/@EyasSH/unexpected-lessons-from-100-test-coverage-eebeee211b7a
---

## [4][TypeScript generic higher order functions](https://www.reddit.com/r/typescript/comments/esf793/typescript_generic_higher_order_functions/)
- url: https://www.reddit.com/r/typescript/comments/esf793/typescript_generic_higher_order_functions/
---
I was wondering if anyone happens to know if what I am trying to achieve is even possible.

What I have in regular Javascript is the following:

    let functions = {
        add: (state, x, y) =&gt; state + x + y,
        subtract: (state, x) =&gt; state - x,
        decrement: (state) =&gt; state - 1,
        increment: (state) =&gt; state + 1,
    };
    
    function wrapper(func) {
        return (...args) =&gt; func(5, ...args);
    }
    
    function wrapObject(obj) {
        const newFunctions = {};
        Object.keys(obj).map(key =&gt; {
            newFunctions[key] = wrapper(obj[key]);
        });
        return newFunctions;
    }
    
    const newFunctions = wrapObject(functions);
    
    newFunctions.add()

The idea is simple. You have an object with functions that all have state as the first argument. I want to create a new object where I return the same functions except where the state is injected. In the above example it's just the number 5 that's always injected but that's beside the point. I achieve this by simply looping over the keys and returning an object with those keys and the values wrapped in a wrapper function.

The problem with this code is that the editor doesn't know what parameters the function add (or any of them) takes. I would like to keep this intact so I figured I would use typescript.

I came up with the following:

    type Dictionary = { [key: string]: (state: number, ...args: any[]) =&gt; any }
    type StateFunction = (state: any, ...args: any[]) =&gt; any
    type OmitFirstParam&lt;T&gt; = T extends (state: number, ...args: infer P) =&gt; any ? P : never;
    type NewFunction&lt;T&gt; = T extends (state: number, ...args: infer P) =&gt; any ? (...args: P) =&gt; void :  never;
    type NewDictionary&lt;T&gt; = T extends { [key: string]: infer F } ? { [key: string]:  NewFunction&lt;F&gt;  } : never;
    
    
    let functions = {
        add: (state: number, x: number, y: number) =&gt; state + x + y,
        subtract: (state: number, x: number) =&gt; state - x,
        decrement: (state: number) =&gt; state - 1,
        increment: (state: number) =&gt; state + 1,
    };
    
    
    function wrapper&lt;T extends StateFunction&gt;(func: T): (...funcArgs: OmitFirstParam&lt;T&gt;) =&gt; void {
        return (...args) =&gt; func(5, ...args);
    }
    
    
    function wrapObject&lt;T extends { [key: string]: (state: number, ...args: any[]) =&gt; any }&gt;(obj: T): NewDictionary&lt;T&gt; {
        const newFunctions: any = {};
        Object.keys(obj).map(key =&gt; {
            newFunctions[key] = wrapper(obj[key])
        });
        return newFunctions;
    }
    
    
    const newFunctions = wrapObject(functions);
    
    newFunctions.increment();

the code in [typescriptlang playground](https://www.typescriptlang.org/play/?ssl=32&amp;ssc=26&amp;pln=1&amp;pc=1#code/C4TwDgpgBAIglgY2HA9gOwIYCcRQLxQDeUA2gNYQgBcUAzsFnGgOYC6NAFPRsBDWgFcAtgCMIWADRQAdLOzNaNDGhAlWASnwA+KMtwBfALAAoUJCgBlYDwgAxAWiSo0+KF2u8lKqbOnzFuipqmng6eiZm0ADyQnDAtnBY9AAK2BhCADwAKjoEWVAQAB68aAAmtG7cnlCCouI+clgKNEwAZuJQySFhKlAA-J1Q-BAAbuIA3BHg0AByEADu9o7I6Nm5UPlFJeWVHnw1wmKSMo3NUG0dXdqBuAMcvv40V6FQIyhwpUNQNaMTU+ZzebwJzobAgNauTbFCBlCrEciUGj0RgsdjnNDtLBQWxQfT9IikCjUOgMJhsGhQQFLEFoDK2HS4r5oX5YSbGEwmAA2EGAUFaDhpFQIhBM32+GFKpU4VX2tSOUkK-EO9SgxLl4m6JJsUAA1FBCrrVRJRWLaAIRAwMEhpXslXVjoqDvbNTKoABafXG4xiqClCAILAQIQw4A2mx2o4uvbuqAARi9PqYAaDIbD1XVWCj2r18ZM+jZHOM-OWzig8ywGDAkCw2QK0Nhlj21JWaC0HGLCBoWXUnF8HYAgk0AjE4gkksBUhXMjlNW8PkQTd9A8ABFgXPdTrRNR2OABWBp+IfqNlGdlnjstssVsBREQAK39wFrWxhO3hRKRpNRadlyuODyHLxVA0a49FxNsUHvLse0pBZgRbMEIRFb0xQQdB6B+RYBRbAIwOFfNFygW8HyQaQiVoDhILvdRpCESsOCJa5kJ9H1mSwkt0IREBWFcctK2rSj7y4jRCP0Y9COXVcXDY5tnFoE9CxMNC0AwmTsLk3jr2Ix923U9DxLPNSOJU6Qk0DYM0GADhjyAA)

It almost works. It correctly realizes that the state parameter is not needed anymore but it thinks that every function takes the same arguments namely x and y. What I think is going wrong is that typescript has to pick one function for the type `key: string]: (state: number, ...args: any[]) =&gt; any` and uses that to build the Dictionary type that I created.

I have a feeling it's not possible.

EDIT: 

Thanks to /u/Erelde I got something that works [link](https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=38&amp;pc=19#code/C4TwDgpgBAysCGwIDECuA7AxsAlge3SgF4oAKAZwSQC4p50QAaKAOjfgCcBzc2+kANoBdAJTEAfHQYBuALAAoUJCgB5ALY5gyHB0oBBbgB5kkksigQAHknQATcmUqIIfBszYtOPWjnQAzCA4oAAUxIkl+KAB+Mg8vXhCwyQA3PBxbKFp0CGTAuUVwaABZeDBIW0MAFQtrCDsHACUITDwOCsoOXy5mOGc0LFwCcVMoAG8oAQBrKF8oSYgQPD8oSqFadU1tXWADLiqpoUkAX3yFABsIYCg-DGx8dAcSUYUoV7pbW1oKKhcodFQ1AAjQLMSxZAHAjjMEDgoGBJJQJxIKAAaigllRUCYLze5FQgOAHHg2C+SN+-zhUPRsMhCLJUAAtOjGDjXrZmhwIGo6sBST8afCJIifoyoABGFnyN4zLCc7noXmOfl-CGC8LC5yYiUKE4KPXyG4De5QADuRLKgSqNRs9ig-TuQ1IhswtEqYmeUrenOAqA4hFIcW4CX4wgRztIAFZ3OwgyJ8kd9QpnYNCGbSipAQArZrAK1WG0OcZTBa0DpdNawH72lNQI7iUh4LOukS0EoWiqVSQe6UtB5XbIm6v3YMMYhjXWe14Z7PYFjzEDkBtZkQsNSlUjzoXd6VvAdDgjkYsgIRjtMWjhLzNH0Ss2tx2-e32EPe3FPkeP63uUP4QQev4enua045k6-4Hve8jnJcUDJPAZyoNAJAvkaB4sL4mByjypAiAoX54BcLBnHgXCkLB8EQDhkHyGRCFjshDoPCw7IYVyWGUXhBFESRNEUfqPF0b++6MXiBJEtgpAAMzsQe+EQIRxGkXBCGUQo-FIYJYGMfAHySYwABM0kPLJ8ncUpFFAA)
## [5][How to add fixed properties to an object](https://www.reddit.com/r/typescript/comments/esdo91/how_to_add_fixed_properties_to_an_object/)
- url: https://www.reddit.com/r/typescript/comments/esdo91/how_to_add_fixed_properties_to_an_object/
---
For example

    interface iObj1 { x:string }
    interface iObj2 extends iObj1 { y:number }
    const someObj : iObj1 = {x:'hello'};
    someObj.y = 1
    someFunc (someObj as iObj2);

I could just be doing something stupid but this doesn't seem to work
## [6][TypeScript Tip of the Week — Nullish Coalescing vs Logical OR Operator](https://www.reddit.com/r/typescript/comments/eskcbq/typescript_tip_of_the_week_nullish_coalescing_vs/)
- url: https://medium.com/@sredmond/typescript-tip-of-the-week-nullish-coalescing-vs-logical-or-operator-72779807051
---

## [7][Does someone have a simple bash script to include non .ts files in /dist after running tsc?](https://www.reddit.com/r/typescript/comments/esie55/does_someone_have_a_simple_bash_script_to_include/)
- url: https://www.reddit.com/r/typescript/comments/esie55/does_someone_have_a_simple_bash_script_to_include/
---

## [8][Is there a easy way for us to extend the language of typescript](https://www.reddit.com/r/typescript/comments/es8hdh/is_there_a_easy_way_for_us_to_extend_the_language/)
- url: https://www.reddit.com/r/typescript/comments/es8hdh/is_there_a_easy_way_for_us_to_extend_the_language/
---
Hi All

Is there a easy way for us to extend the language of typescript?  adding my own syntax?

thanks

Peter
## [9][Type vs Interface (React/React-native)](https://www.reddit.com/r/typescript/comments/es2e4c/type_vs_interface_reactreactnative/)
- url: https://www.reddit.com/r/typescript/comments/es2e4c/type_vs_interface_reactreactnative/
---
Hey beautiful people!

I just ported my project (react native) over to typescript but dont get the hang when to use **type and when to use interface**

I looked at some examples, but just added to the confusion, because apparently both have their place but i cannot figure it out

Could you tell me on a high level the difference between these two and if you have expereicne with React/React-Native where to use which?

&amp;#x200B;

Thanks a lot!

&amp;#x200B;

Edit: So as far as I understood it, it doesn't matter too much. To get started I will stick to types and then have a look at where it goes. If this is completely wrong let me now.

Thanks for all of your answers!
## [10][Dev Diary: 5 Tips for Building Beautiful CLIs with Node.js, Yargs, &amp; Ink](https://www.reddit.com/r/typescript/comments/erzcbn/dev_diary_5_tips_for_building_beautiful_clis_with/)
- url: https://medium.com/eximchain/dev-diary-5-tips-for-building-beautiful-clis-with-node-js-yargs-ink-16d184ea0d14
---

## [11][Thoughts on Map and Set in TypeScript](https://www.reddit.com/r/typescript/comments/es2fav/thoughts_on_map_and_set_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/thoughts-on-map-and-set-in-typescript-efd43c3bf2ad?source=friends_link&amp;sk=5863d3228b0d37f3d493bb1d001a2c13
---

