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
## [2][Type 'string' not assignable to a custom type](https://www.reddit.com/r/typescript/comments/hvre75/type_string_not_assignable_to_a_custom_type/)
- url: https://www.reddit.com/r/typescript/comments/hvre75/type_string_not_assignable_to_a_custom_type/
---
(Typescript beginner using React here)

I am trying to use a custom type definition (`type Ingredients = 'bread-bottom' | 'bread-top' | 'meat' | 'cheese' | 'salad' | 'bacon';`) to set the type of the input value in a map method.

Here is the code I am running

      const ingredients: ReactElement[] = Object.entries(props.ingredients)
        .reduce((acc: Array&lt;string&gt;, [ing, num]: [string, number]) =&gt; {
          let arr: Array&lt;string&gt; = [];
          for (let i = 0; i &lt; num; i++) {
            arr.push(ing)
          };
          return [...acc, ...arr];
        }, [])
        .map((ingredient: Ingredients, i: number) =&gt; {
          return (
            &lt;Ingredient
              key={i}
              type={ingredient}
            /&gt;
          );
        });

I know the exact string values of the elements that will be returned in the array, as I have set a `Props` interface above to check. Therefore, in the `.map` method, I know that the value, which I name `ingredient` will have the type `Ingredients`, however, I believe that the compiler does not know this. The type property in the `&lt;Ingredient&gt;` element already has type checking, the same as the `Ingredients` type. The error thrown is

    Argument of type '(ingredient: Ingredients, i: number) =&gt; JSX.Element' is not assignable to parameter of type '(value: string, index: number, array: string[]) =&gt; Element'.
      Types of parameters 'ingredient' and 'value' are incompatible.
        Type 'string' is not assignable to type 'Ingredient'.  TS2345
    
        14 |       return [...acc, ...arr];
        15 |     }, [])
      &gt; 16 |     .map((ingredient: Ingredients, i: number) =&gt; {
           |          ^
        17 |       return (
        18 |         &lt;Ingredient
        19 |           key={i}

Changing the type of `ingredient` to `string` will remove the error, but cause another one on the `type` attribute as I implement the same type checking on it through the props in `&lt;Ingredient&gt;`. Removing both specific type string checks will remove the errors. Should I leave it as that, or is there a way to fix this?
## [3][Question around Using an enum as an interface type](https://www.reddit.com/r/typescript/comments/hvkm3s/question_around_using_an_enum_as_an_interface_type/)
- url: https://www.reddit.com/r/typescript/comments/hvkm3s/question_around_using_an_enum_as_an_interface_type/
---
If I'm defining an enum to use as a type on an interface, which of these options would it be better to write as?

    export enum someEnum {
      A = "A",
      B = "B"
    }
    
    export interface Foo {
        / * Option 1 */ 
       prop: someEnum.A | someEnum.B
        / * Option 2 */ 
       props: someEnum;
    }

I'm not sure if I should assign the whole enum as the type or the individual pieces, especially if there are only two. Could anyone tell me which one is better here?
## [4][Enforcing type on the result of Papaparse](https://www.reddit.com/r/typescript/comments/hvtkta/enforcing_type_on_the_result_of_papaparse/)
- url: https://www.reddit.com/r/typescript/comments/hvtkta/enforcing_type_on_the_result_of_papaparse/
---
Here's the thing. I have a RAW CSV file, and I am trying to change into an array of, say, `TestType`. I am using PapaParser for the job. Unfortunately, the result has the type `Papa.ParseResult&lt;unknown&gt;.data` . I can transform it to an array of objects of type `TestType` like so:

    const objs: TestType[] = (parsedResult.data as unknown) as TestType[];

However it feels like a bad pattern. Is there any way to use duck typing to coerce the result into an array of different type?
## [5][Intro to Typescript Generics](https://www.reddit.com/r/typescript/comments/hvccoj/intro_to_typescript_generics/)
- url: https://www.youtube.com/watch?v=h6L7egCOai4
---

## [6][Conversion of String to Enum](https://www.reddit.com/r/typescript/comments/hvigps/conversion_of_string_to_enum/)
- url: https://www.reddit.com/r/typescript/comments/hvigps/conversion_of_string_to_enum/
---
Here's the thing. I have an Enum in my Prisma schema like so:

    enum Exchange {
      NSE
      BSE
      MCX
    }

For all intents and purposes, it can be considered as typescript enum. Now let's say, I have a string `const x = "NSE"`. Is there any way to convert x into `Exchange.NSE`, without using pesky ternary operations or a pattern matching switch-case or cascading if statements?
## [7][It looks like there’s yarn-based lerna alternative to consider now](https://www.reddit.com/r/typescript/comments/hvbvgo/it_looks_like_theres_yarnbased_lerna_alternative/)
- url: https://github.com/guigrpa/oao
---

## [8][String not being detected in else.](https://www.reddit.com/r/typescript/comments/hvg22e/string_not_being_detected_in_else/)
- url: https://www.reddit.com/r/typescript/comments/hvg22e/string_not_being_detected_in_else/
---
The following code doesn't compile. It complains

`"Argument of type 'T' is not assignable to parameter of type 'string'."`

    type StringReturn = {};
    type MinusOneReturn = {};
    
    
    const myFunction = &lt;T extends string | -1&gt; ( 
        param: T
    ): T extends string ? StringReturn : MinusOneReturn  =&gt; { 
        if (param === -1) { 
            minusOneFunction(param); // error
        } else { 
            stringFunction(param); // error
        } 
    }
    
    const stringFunction = (stringParam: string): StringReturn;
    const minusOneFunction = (minusOneParam: -1): MinusOneReturn;

Any idea how to solve it?
## [9][What's your preferred tool for testing type definitions?](https://www.reddit.com/r/typescript/comments/hvhrhm/whats_your_preferred_tool_for_testing_type/)
- url: https://www.reddit.com/r/typescript/comments/hvhrhm/whats_your_preferred_tool_for_testing_type/
---
What do you use to write tests for asserting that things have correct type, and to assert that there are type error when you expect? I work on libraries and it's something would help a lot. I tried dts-lint before but it was too opinionated, I want something which just tests types and nothing else.

The new `$ExpectError` in TS is nice, but since you don't specify what error it is, it's not very useful for testing.
## [10][TypeTypeScript with Promises, Async/Await, and Generator Functions](https://www.reddit.com/r/typescript/comments/hvh6op/typetypescript_with_promises_asyncawait_and/)
- url: https://medium.com/@binyamin/typescript-with-promises-and-async-await-63623b8e5e2a?source=friends_link&amp;sk=3b00110853a3ef98a7797d7ce83a0f9f
---

## [11][Using an AWS sdk call in class contructor](https://www.reddit.com/r/typescript/comments/hv5qbr/using_an_aws_sdk_call_in_class_contructor/)
- url: https://www.reddit.com/r/typescript/comments/hv5qbr/using_an_aws_sdk_call_in_class_contructor/
---
Hey guys, I am trying to use a call to AWS Secret Manager in the constructor of an API class. The instance will be used to send calls to that API, and the call to AWS Secret Manager is needed to fill the auth.

However, Typescript warns me to not use await in the constructor. Since the AWS Secret Manager calls seem to be async, how would I use this service in my API class? Any advice would be welcome.


```js
export class SOMEServiceAPI {
    private readonly options: any;

    constructor() {
        this.options = {
            timeout: 60000,
            method: 'PUT'
        };
        await this.makeOptions();
    }


    makeRequest(body: string) {
        return Object.assign({}, this.options, {
            body
        });
    }


    async put(event: IncomingEvent): Promise&lt;Response&gt; {
        const request = this.makeRequest(JSON.stringify(event.event_payload));
        const date: Date = new Date();
   &lt;....&gt;
    }

    private makeOptions = async() =&gt; {
        const Auth = await awssdkcall; 
&lt;..fills in the authHeader and Api key..&gt;



        return this.options.headers = {
            Authorization: authHeader,
            'x-api-key': apiHeader
        };
    }
}

```
