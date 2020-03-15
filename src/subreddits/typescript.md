# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Put together an illustration to clear up my own confusion as to why the Union and Intersection operators are named as such](https://www.reddit.com/r/typescript/comments/fisz4q/put_together_an_illustration_to_clear_up_my_own/)
- url: https://www.reddit.com/r/typescript/comments/fisz4q/put_together_an_illustration_to_clear_up_my_own/
---
The naming of the Union and Intersection operators had always felt curious to me, as it seemed in direct contrast to their respective functions. Why do we call `&amp;` the Intersection operator when the declared object possesses not only the mutual members, but the members of all concerning types? Why do we call `|` Union when the declared object can only reference the mutual members? Recently stumbled upon an SO question asking the same things and decided to finally sit down and reason through it. While the existing answers are great and helped a good deal, I figured it might be of value to put my own understanding down (in good old graphic form) for others who may fall for the same misinterpretive line of thinking.

SO answer:
https://stackoverflow.com/a/60686242/1107110

Illustration:
https://i.stack.imgur.com/fY4gL.png

Feel free to share your thoughts. Does it make sense? Does it clear things up? Does it confuse you more? All feedback welcome. 🙂
## [3][Correct way of storing interfaces in TypeScript](https://www.reddit.com/r/typescript/comments/fi8v85/correct_way_of_storing_interfaces_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fi8v85/correct_way_of_storing_interfaces_in_typescript/
---
New to TS here. When I have to declare objects, I create their corresponding interfaces in the same file.

```ts
// Interfaces
interface User {
  name: string;
  age: number;
}

interface City {
  state: string;
  zipcode: string;
}

// Objects
const bob: User = {
  name: 'Bob',
  age: 20
}

const boston: City = {
  state: 'MA',
  zipcode: '02022'
}
```
1. What is the correct way of storing interfaces? Is it storing all of them in one file and importing as needed or is it storing one interface per file?

2. What is the role of `.d.ts` files here? Can we store all interfaces under one namespace in a `.d.ts` file?
## [4][Designing the perfect TypeScript schema validation library](https://www.reddit.com/r/typescript/comments/fhwj5m/designing_the_perfect_typescript_schema/)
- url: https://vriad.com/blog/zod/
---

## [5][Something like "Nested" Pick](https://www.reddit.com/r/typescript/comments/fi2za5/something_like_nested_pick/)
- url: https://www.reddit.com/r/typescript/comments/fi2za5/something_like_nested_pick/
---
Hi there,

I've just recently discovered *Pick* and I'm really enjoying the idea of using it on our code base and get ride of most of our *Partial's*.

However, I have a doubt, imagine that I have a type called *Company*.

    export interface Company {
      id: string;
      industry: Industry;
      name: string;
      offices?: Office[];
      (...)
    }

I can use *Pick* to create a type that only includes the name and the array of offices.

    Pick&lt;Company, 'name' | 'offices'&gt;

However, my *Office* type is quite big and, for this specific use case, I'll only need the city of each office

    export interface Office {
      id: string;
      addressLine: string;
      addressLineOther: string;
      city: string;
      country: string;
      phoneNumber: string;
      postalCode: string;
      state: string;
    }

Can I use *pick* (or something else) to only pick the cities and end up with a type like this?

    {
      name: string;
      offices?: {
        city
      }[];
    }
## [6][Up and running in a breeze with Firebase](https://www.reddit.com/r/typescript/comments/fhxjcr/up_and_running_in_a_breeze_with_firebase/)
- url: https://blog.pragmatists.com/up-and-running-in-a-breeze-with-firebase-72ca889d22cb
---

## [7][[Question] extending a type for a generic function](https://www.reddit.com/r/typescript/comments/fhkgcg/question_extending_a_type_for_a_generic_function/)
- url: https://www.reddit.com/r/typescript/comments/fhkgcg/question_extending_a_type_for_a_generic_function/
---
I have a generic function that should only accept a generic type that extends another type

In another file, I am passing that function as argument, and would like to create a `type` that details this.


    // form-values.ts

    export type FormValues = {
      [key: string]: string
    }

    export type SomeSpecificFormValuesType = FormValues &amp; {
      email: string
      password: string
    }

    export type AnotherFormValuesType = FormValues &amp; {
      firstName: string
    }

    export type FormatFormValues&lt;T extends FormValues&gt; = (formValues: T) =&gt; T

    // factory.ts
    import { FormValues, FormatFormValues, SomeSpecificFormValuesType, AnotherFormValuesType } from './form-values'

    // simplified factory
    export type Factory = {
      formatFormValues?: FormatFormValues&lt;SomeSpecificFormValuesType | AnotherFormValuesType&gt; 
      ...
    }

    export const factory: Factory (formatFormValues, ..args) =&gt; {
      return new FactoryItem(formatFormValues, ...args)
    }

    // a file that uses the factory
    import { FormatFormValues } from './form-values'
    import { factory } from './factory'

    // simplified example of this function
    const formatFormValues: FormatFormValues&lt;SomeSpecificFormValuesType&gt; = formValues =&gt; {
      const { email, password } = formValues
      return { email, password }
    }

    const instance = factory(formatFormValues, ...) // TypeScript Error...



Where the instance is created gives the following error:


    Argument of type 'FormatFormValues&lt;SomeSpecificFormValuesType&gt;' is not assignable to parameter of type 'FormatFormValues&lt;SomeSpecificFormValuesType | AnotherFormValuesType&gt;'
      Types of parameters 'formValues' and 'formValues' are incompatible.
          Type 'SomeSpecificFormValuesType | AnotherFormValuesType' is not assignable to type 'SomeSpecificFormValuesType'.
          Type 'SomeSpecificFormValuesType' is not assignable to type '{ email: string; }'


I don't understand why I am getting this error — the type passed to the `FormatFormValues` function extends `FormValues`. Can someone explain?
## [8][TicTacToe Typescript Backend With Sockets](https://www.reddit.com/r/typescript/comments/fhlcr5/tictactoe_typescript_backend_with_sockets/)
- url: https://www.reddit.com/r/typescript/comments/fhlcr5/tictactoe_typescript_backend_with_sockets/
---
&amp;#x200B;

https://reddit.com/link/fhlcr5/video/wuniv516eam41/player

I built a typescript backend that uses express endpoints and ws npm package for the sockets.  I used jest library with typescript for the testing.  Let me know what you think and if you have any suggestions, especially about the tsconfig.json or package.json.

TicTacToe Game:

[https://phptuts.github.io/tictactoe-svelte-app/](https://phptuts.github.io/tictactoe-svelte-app/)

Typescript Node Backend Code:

[https://github.com/phptuts/tictactoe-server](https://github.com/phptuts/tictactoe-server)

Frontend Svelte Code:

[https://github.com/phptuts/tictactoe-svelte-app](https://github.com/phptuts/tictactoe-svelte-app)
## [9][Native apps using typescript powered by rust](https://www.reddit.com/r/typescript/comments/fh4muu/native_apps_using_typescript_powered_by_rust/)
- url: https://v.redd.it/lpq08i5q14m41
---

## [10][Need help with TS + Node config with custom typed config interfaces | StackOverflow question](https://www.reddit.com/r/typescript/comments/fhe4sl/need_help_with_ts_node_config_with_custom_typed/)
- url: https://stackoverflow.com/q/60565055
---

## [11][How to add custom properties on the Request object in Express](https://www.reddit.com/r/typescript/comments/fgxob0/how_to_add_custom_properties_on_the_request/)
- url: https://www.reddit.com/r/typescript/comments/fgxob0/how_to_add_custom_properties_on_the_request/
---
Hi there!

How do I add custom properties on the Request object in Express, with Typescript?

I have on strict mode and need to somehow add a property on the request object so I have access to it in my controllers.
