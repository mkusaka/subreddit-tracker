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
## [2][Help me understand .d.ts files](https://www.reddit.com/r/typescript/comments/hly19w/help_me_understand_dts_files/)
- url: https://www.reddit.com/r/typescript/comments/hly19w/help_me_understand_dts_files/
---
Every time I think I know enough about `.d.ts` files to use them in my own code, I'm soon reminded that I don't have a clue about them. I know the basics of JS and TS, but I'm not familiar with what ES* is or how JS import statements work.

As I understand it, `.d.ts` files:

- are "declaration" files that can only contain compile-time types. Is that right? Only `interface`s and `enum`s and `type`s, no `class`es or `function`s?
- are auto-imported by other files in the same directory, somehow? What's going on here?

Right now I'm using a `.d.ts` file to hold my API's types and enums for my project. My API client is in the same folder:

    src
    ├── api
    │   ├── api-client.ts
    │   └── interfaces.d.ts
    ├── login
    ├── navbar
    └── sidebar

I want to be able to use these types anywhere I might need them. I can import `api-client.ts` from a file in one of those other folders under `src/` but those files then can't see `interfaces`. If there's a way to import it, I don't know how.

If anyone knows any articles I can read to understand how all this importing stuff works, that would be awesome. Or you can try to ELI5 it for me.

A few other related questions that I fear may otherwise go unanswered:

- _Can_ you manually import a `.d.ts` file? If not, why not?
- Is there a compiler setting I can tweak to make it so that `interfaces` is imported into every `.ts` file under `src/`?
- ELI5 `export` / `declare` / `default export` assuming a linked article does not cover it
## [3][Function return type as an object where the keys are values from the input.](https://www.reddit.com/r/typescript/comments/hlwk9g/function_return_type_as_an_object_where_the_keys/)
- url: https://www.reddit.com/r/typescript/comments/hlwk9g/function_return_type_as_an_object_where_the_keys/
---
So I have an interface that looks like this:
```
interface Task {
   name: string;
   command: string;
   ...
}
```
And I want to make a function that would take an array of that interface, something like this: 
```
function execute([{ name: "TEST", ... }, { name: "MSG", ... }]) 
```
do some things and return an object looking like this: `{ TEST: any, MSG: any }`

But the trick is that the typescript compiler would be aware of the keys and not return the generic `[key: string]: any` or anything like that.

I've seen it being implemented in the `@types/inquirer` but I couldn't implement it myself even after looking at that code for several hours. 

Can someone guide me into solving this or give me a solution on how to do this? Please.
## [4][Strict null checks and default values](https://www.reddit.com/r/typescript/comments/hll24h/strict_null_checks_and_default_values/)
- url: https://www.reddit.com/r/typescript/comments/hll24h/strict_null_checks_and_default_values/
---
In my code I want to have optional output parameter (as ref or ptr in C++), and actual return value of the same parameter.

The following code generates TS2532 with Strict Null Checks enabled.

```TypeScript
const foo = (out?: Buffer): Buffer {
  if(!out) {
     out = Buffer.alloc(100);
  }

  // manipulate buffer
  out.length // Error: TS2532 Object is possibly 'undefined'


  return out;
}
```

Why `out.length` after `if(!out)` check still generates that error?
## [5][Union and Intersection notation](https://www.reddit.com/r/typescript/comments/hlnswd/union_and_intersection_notation/)
- url: https://www.reddit.com/r/typescript/comments/hlnswd/union_and_intersection_notation/
---
How is a Union different than a logical OR? How is an Intersection different than a logical AND. If they do conceptually correspond to the aforementioned logical operators, why do they use different notation (| instead of ||, and &amp; instead of &amp;&amp;)?
## [6][Levenshtein Distance in TypeScript](https://www.reddit.com/r/typescript/comments/hla0ei/levenshtein_distance_in_typescript/)
- url: https://medium.com//levenshtein-distance-in-typescript-6de81ea2fb63?source=friends_link&amp;sk=81db1fe3f82a039f967dd477029aa451
---

## [7][A type-safe MobX router with parallel routing and full lifecycle hooks support](https://www.reddit.com/r/typescript/comments/hl87k3/a_typesafe_mobx_router_with_parallel_routing_and/)
- url: https://www.reddit.com/r/typescript/comments/hl87k3/a_typesafe_mobx_router_with_parallel_routing_and/
---
We have been actively developing [Boring Router](https://makeflow.github.io/boring-router/) along with our own main project for two years. It's no longer as light-weight as it was, but it also gets some unique and useful features.

There are several keywords you might be interested in with this router:

1. **State First**  
It manages route states (like `route.workbench.$matched`) and provides thin react components that reacts to those states.
2. **Type Safe**  
It has object-based route notation (like `route.workbench.task.taskId.$push({taskId: '124'})`) with type information, and you can also share it with Node.js server.
3. **Parallel Routing Support**  
Probably an exclusive feature. It allows developer to create a primary route with multiple parallel routes to describe multiple route views. E.g.: `/app/workbench?_sidebar=/notification`.
4. **Full Lifecycle Hooks Support**  
Dedicated `BrowserHistory` implementation that tracks the history stack and is capable of restoring the stack with a given snapshot. This makes using lifecycle hooks safe to browser history stack with browser navigation.

&amp;#x200B;

* GitHub [https://github.com/makeflow/boring-router](https://github.com/makeflow/boring-router)
* Examples [https://makeflow.github.io/boring-router/examples](https://makeflow.github.io/boring-router/examples)
## [8][Why? (Codesanbox included)](https://www.reddit.com/r/typescript/comments/hlgai2/why_codesanbox_included/)
- url: https://www.reddit.com/r/typescript/comments/hlgai2/why_codesanbox_included/
---
\`SwitchType\` has a name field, why do I get this error?

https://preview.redd.it/9hjpy4nvny851.png?width=777&amp;format=png&amp;auto=webp&amp;s=417193bfebdb05e0cafe208e697d023a00fcd3f7

Full code:  [https://codesandbox.io/s/unruffled-gauss-f08r8?file=/src/index.ts](https://codesandbox.io/s/unruffled-gauss-f08r8?file=/src/index.ts)
## [9][`unions can't be used in index signatures, use a mapped object type instead`](https://www.reddit.com/r/typescript/comments/hl3svn/unions_cant_be_used_in_index_signatures_use_a/)
- url: https://www.reddit.com/r/typescript/comments/hl3svn/unions_cant_be_used_in_index_signatures_use_a/
---
I am trying to clear the errors in the below code block. After getting the error \`unions can't be used in index signatures, use a mapped object type instead\` I am now attempting to convert a string literal union (key names of an interface) to a mapped object: [https://github.com/microsoft/TypeScript/issues/24220#issuecomment-390063153](https://github.com/microsoft/TypeScript/issues/24220#issuecomment-390063153)

But it still isn't working:

    const returnObject: { 
       [index : mappedObjectOfConfigDataKeys] : Partial&lt;ConfigData&gt;
     } = {...accumulator};
    
    // error TS1023: An index signature parameter type must be either 'string' or 'number'.
    // 51 [index : // mappedObjectOfConfigDataKeys] : // Partial&lt;ConfigData&gt;
    //     ~~~~~
    
     if (currentStepInstance.hasSaveableData &amp;&amp; currentStepInstance.configDataKey) {
        returnObject[currentStepInstance.configDataKey] = validatedUserInput;
     } 
    // error TS7053: Element implicitly has an 'any' type because expression of type 'string' can't be used to index type '{}'.
    //  No index signature with a parameter of type 'string' was found on type '{}'.
    // 55  returnObject[currentStepInstance.configDataKey] = validatedUserInput;
           ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

When trying to assign a string literal union:

    const returnObject: { 
       [index : stringLiteralUnion] : Partial&lt;ConfigData&gt;
     } = {...accumulator};
    // (parameter) index: stringLiteralUnion
    // An index signature parameter type cannot be a union type. Consider using a mapped object type instead.ts(1337)

Anyone know how to fix this? These last two attempts have me in circles. The compiler seems to  want the mapped object to pass a constrained set of strings; I don't know how to give it that in an acceptable way.

    export default interface ConfigData {
       languageCode : string
       , numberOfRepeats : number
       , projectId : string
    }
    
    /**** Index Signature handling for the above interface ****/
    export type stringLiteralUnion = "languageCode" | "numberOfRepeats" | "projectId";
    
    export type mappedObjectOfConfigDataKeys = {[key in stringLiteralUnion]: string}; 
    // also tried typing it as `any` as shown in the linked post
    // same error occured

[Link to TS Playground](https://www.typescriptlang.org/play/?ssl=29&amp;ssc=16&amp;pln=29&amp;pc=1#code/KYDwDg9gTgLgBAE2AMwIYFcA28CWA7GYKNAY2DgGEI9kcBzAEVRlTgG8AoObzVPO9KjrAqSOAC44AZxhR8dLtwA0cPOgC2AIyIB5ZACVgYYMykTVG7VEVwVYKBABWwEjACSCczLn8OAXw4OAHoAKjCQuDc8JBA4AGV6PGZ0KHIACz4ETHk4ZGg4GDTyVE0IADdyfEJiVDI4cJCgjlBIWAKAT2NpWXkAGRxq1EwAVTwcajgAXjgAIl5%20QWFRYBm4AB9ZtS1dAyMTGClVjZn7Jxd3BBmAbkCW6HgYTvJ1VDBjBB1NZ1c9Khp6JgsADSwHaZmmbDgAG0ANaguD4bo%20Oj9QYjMbUAC6km8OT8N2CQTgAHUcAAvVBQBBxQhgMxRaqkYCBKpEJkk8mU6m0sycbhwVKoBDUTDtOAZKRxVAVEqYYCA1CSUoQOV8GwqQXCvCi1TAYAIKQAMRwcoAakMcAhmOM8EqICqTHh1QKTFqdfgyhaEFEwOgYABZYBSKRCYAAfhxPX46zg6GiKHw%20udmpFYpI1FojGYqBB7QjSJyGzjSFoeCTzpavHwAAoAJSSMoQS3O07qMAwOuR5HOz3ZK2EH1%206voKREQcwfO46NF%20Ol-X1uDK1VO-kqXuW5jAY1ysPV2hygAKzDSk6jdBjxYTZYQC6XjpjjebAUCQSJFF4wbgAHcBmk4ERkTgdRgEKCAEECEgPzMOIQPQMBSQpKl2BsV84AAWgwzCsOwuADwcYxYBwIMbHsHBPUIbojCkSQAEEoCgVB2gAHgQrkaSogA%20G5%20RsdM8G8dBXGgashkwdi6Vo%20jGJYzkqXEqQONrZD%20X5QocCkAA6GQqKmOBRPk7juGfHj%20VQ7DzKw3D0E0bISDgQNQINLgUIiAAJTI5TgTB7TAHIIGQAoikosAvIgOgcBIGx%20QAUVqP9tJCn9MEwf9wCrPANPiCBgO-E0UvXftKjwX0YBsRoSOs2yETGGAcCGGCYDgztKAzAFs2UlS4FQqA4z05Lgs0lq2xNIMBtywoXUaqAkhs8gYSCMo4DAVAcCgMwqggPS4D4zM4AgL5zkCTrtuoGQTv%20LMWF0tTNISzTUgQQTgGrKLOpEkgSA0LBmGgSQj0IoYmL%20TMFQ4lRPvo4ACHEqIZD4MhJFYuSeQXf7asB4G2pYDiphxvlju4PipAdDTvLoasZlqT71G%20mBoBmFQqa%203g6agWsCQJ7hUPwATXFelSIdSaHaVhlg8DIDTKxWvA6w5znUO8iAQrjWqUtYArqpKhEzFSKbr350yiVAAYTrEdS%20uJ8UPKTTnzrOgrNwQYZRygccuxyaYbo04QYHNPtx2rQWoZgGH%20LFsh2YNwnTpJsnqwAcgdwgnZd8d47XL1Hedsdir9SOo66okbIgEgYTMFWTTgDX90qMw8mLFQJhrqvM9qiY7oLnAAsDlIhZDkWw-h4ANLLfUjRGv2NzbvA4AAMln7be%20D0O4fF4ek63EalPx23F8h4WjFFoeNI37dnsj23jPloldZSGfWHmxbltW3J8ju8a-xwNt7j4XBc4OAuRN4C32mp8b4MBJCGHTFSJicw%20ACFDMsI4mxLA7EMMYUwyCTgOHAR4BmuFKTo0wEDVql1UAcRxhCDS1CmY0xZtAfEgRO7dyDgfMAR814aQlFKGUs0FRzwXqw-uh9B6cJ2ljHMoJt4F35CAvAYDzhQiESvcOw9xFkNzJiXSG8U45xKoZAmfg4BHU5nIyad8FGuAMfyPwKg2B%20CUqgMwmMyFXELnpYMRBgGwWmh0LoEwNIPSenWExAsY5ylJmFBOyAsApX2uAyQ6iFTp3OiDbMtZ%20ZmKSdmAxV9jEoSJBZIpcBXLAEwARICIE0hgSkKE0I9lpa5DjK4G02sCibUVsrAglc%20Caz9G0pAwBgKeAKvzSe-Z8jmz4iwRMnhESFHIPE84CIDhlOQIzD6QZ9RVzqogQSMI-HyDKk0fkpxCCuG2T7cZAd1HiURrJbkVEFxTnPDvfkX40gjTgNWWQ6BgDSNtkAgUqAvxuwLNGaYtzaQaVbO2WWoTjqoRIEUUur8oB6RnsbeAvY-kFy9liigKKYTVgYqC-%20%20dAWnT-iVNwUhxl2gdL0yFpDxIn1bsAAOpLxwUs5l3b5%20AaV0q9AC3e5jfFcv-tYzqRiymjg6rvImscolQqMBpD0XpxyBmDKGDJl9%20Z5NUmkBwX5dQmuivRYSMw4jVKwJ4bQsY8CCmRbKFYF8jJMNOXIci5ACVEuHKnf%207t%20AirCWHVQ0AXjZDJPqbOrt-66RHHov0GkejqDrCmiAvQIBfiIBQJx58pX-i2KlAYyxeQlvgNMGYWL8EAEd0AmyrfWgYqwDXcD5dWPAEaLTRt0XGrWkxB0VrLZLEAJs1hrBkV2qAkbyQxoDQOodBKwJBg0s20qu8Q0E1OGQYMo6BjwuOgEd1-gOBAA)

&amp;#x200B;

**A couple of attempts at fixing this**

    const returnObject: Record&lt;keyof ConfigData, Partial&lt;ConfigData&gt;&gt; = {...accumulator};
    
    // error
    const returnObject: Record&lt;"languageCode" | "numberOfRepeats" | "projectId", Partial&lt;ConfigData&gt;&gt;
    
    Type '{ languageCode?: string | undefined; numberOfRepeats?: number | undefined; projectId?: string | undefined; }' is not assignable to type 'Record&lt;"languageCode" | "numberOfRepeats" | "projectId", Partial&lt;ConfigData&gt;&gt;'.
    
      Types of property 'languageCode' are incompatible.
        Type 'string | undefined' is not assignable to type 'Partial&lt;ConfigData&gt;'.
          Type 'undefined' is not assignable to type 'Partial&lt;ConfigData&gt;'.ts(2322)

&amp;#x200B;

    const returnObject: Record&lt;"languageCode" | "numberOfRepeats" | "projectId", Partial&lt;ConfigData&gt;&gt; = {...accumulator};
    
    // error
    const returnObject: Record&lt;"languageCode" | "numberOfRepeats" | "projectId", Partial&lt;ConfigData&gt;&gt;
    
    Type '{ languageCode?: string | undefined; numberOfRepeats?: number | undefined; projectId?: string | undefined; }' is not assignable to type 'Record&lt;"languageCode" | "numberOfRepeats" | "projectId", Partial&lt;ConfigData&gt;&gt;'.
    
      Types of property 'languageCode' are incompatible.
        Type 'string | undefined' is not assignable to type 'Partial&lt;ConfigData&gt;'.
          Type 'undefined' is not assignable to type 'Partial&lt;ConfigData&gt;'.ts(2322)
## [10][Typescript and mutually exclusive variables ?](https://www.reddit.com/r/typescript/comments/hl1af0/typescript_and_mutually_exclusive_variables/)
- url: https://www.reddit.com/r/typescript/comments/hl1af0/typescript_and_mutually_exclusive_variables/
---
I know that there are Boolean types for true or false. However I would like to have a value pairs, or some might call them interdependent polarities. The world around us is full of that. Like temperature = is warm or is cold. Or age = young vs old. Light = dark | bright. I was thinking about two methods in an object that will set to true or false the corresponding properties. But again, these are propreties and not actual variables. Or maybe there is some other way? It must be simpler than I think it is
## [11][WHERE to put my types/interfaces?](https://www.reddit.com/r/typescript/comments/hknqf4/where_to_put_my_typesinterfaces/)
- url: https://www.reddit.com/r/typescript/comments/hknqf4/where_to_put_my_typesinterfaces/
---
Hey,  
So, I am working on my first TS project and I want to make sure I start off correctly.   
What is the general concensus in terms of architecture? This is a website, not a library.  


1. Do you have your types/interfaces in a seperate file?
   1. If so, do you name it with a '\*.d.ts" or is that ONLY for libraries/modules to be used outside the codebase?
2. Do you put your types IN the same file that is being used? If you have to re-use that type to you just export it? (won't that get a bit disorganized after a while?)
3. Do you have a "types" folder that you define by feature?

&amp;#8203;

    /src
      /src/products
         --- product.tsx
         --- product.types.ts // these can be used in this feature or wherever?

Basically, what I am asking is HOW do you structure your typings that is robust &amp; expandable?
