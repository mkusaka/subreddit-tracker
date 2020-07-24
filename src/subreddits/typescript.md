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
## [2][How to increase code coverage for mocha testing on basic Typescript example](https://www.reddit.com/r/typescript/comments/hwtsd3/how_to_increase_code_coverage_for_mocha_testing/)
- url: https://www.reddit.com/r/typescript/comments/hwtsd3/how_to_increase_code_coverage_for_mocha_testing/
---
I want to get a better understanding of how to test these kinds of functions in Typescript using mocha. In the StackOverflow post below I've outlined a basic small Typescript file, which I want to test and increase code coverage. Right now I'm not sure why the breakdown of the testing is the way that it is (50% Stmts | 100% Branch | 0% Funcs | 50% Lines) and how I would be able to get everything to 100%.

[https://stackoverflow.com/questions/63065938/how-to-increase-code-coverage-for-mocha-testing-on-basic-typescript-example](https://stackoverflow.com/questions/63065938/how-to-increase-code-coverage-for-mocha-testing-on-basic-typescript-example)

Any advice would greatly be appreciated!
## [3][Rest API app built with Deno and TypeScript](https://www.reddit.com/r/typescript/comments/hwzjdu/rest_api_app_built_with_deno_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hwzjdu/rest_api_app_built_with_deno_and_typescript/
---
Part of the reason many people are now paying attention to Deno is not only its ambitious goal to compete with Node.js, but also its built-in TypeScript support. Do you want to know what it's like to develop a REST API application with Deno. My friend [recounted](https://tsh.io/blog/deno-rest-api/) his experience with that. Thought you might find it interesting.
## [4][TS2488: Type 'string | null' must have a '[Symbol.iterator]()' method that returns an iterator](https://www.reddit.com/r/typescript/comments/hwzeo6/ts2488_type_string_null_must_have_a/)
- url: https://www.reddit.com/r/typescript/comments/hwzeo6/ts2488_type_string_null_must_have_a/
---
I am trying to convert https://github.com/camwiegert/typical/blob/master/typical.js into TypeScript.

In Vanilla JS, it looks like:

# typical.js

```js
export async function type(node, ...args) {
    for (const arg of args) {
        switch (typeof arg) {
            case 'string':
                await edit(node, arg);
                break;
            case 'number':
                await wait(arg);
                break;
            case 'function':
                await arg(node, ...args);
                break;
            default:
                await arg;
        }
    }
}

async function edit(node, text) {
    const overlap = getOverlap(node.textContent, text);
    await perform(node, [...deleter(node.textContent, overlap), ...writer(text, overlap)]);
}

async function wait(ms) {
    await new Promise(resolve =&gt; setTimeout(resolve, ms));
}

async function perform(node, edits, speed = 60) {
    for (const op of editor(edits)) {
        op(node);
        await wait(speed + speed * (Math.random() - 0.5));
    }
}

export function* editor(edits) {
    for (const edit of edits) {
        yield (node) =&gt; requestAnimationFrame(() =&gt; node.textContent = edit);
    }
}

export function* writer([...text], startIndex = 0, endIndex = text.length) {
    while (startIndex &lt; endIndex) {
        yield text.slice(0, ++startIndex).join('');
    }
}

export function* deleter([...text], startIndex = 0, endIndex = text.length) {
    while (endIndex &gt; startIndex) {
        yield text.slice(0, --endIndex).join('');
    }
}

export function getOverlap(start, [...end]) {
    return [...start, NaN].findIndex((char, i) =&gt; end[i] !== char);
}
```

And my TS conversion looks like:

# typical.ts

```ts
export async function type(node: HTMLElement, ...args: any[]): Promise&lt;void&gt; {
  for (const arg of args) {
    switch (typeof arg) {
      case 'string':
        await edit(node, arg);
        break;
      case 'number':
        await wait(arg);
        break;
      case 'function':
        await arg(node, ...args);
        break;
      default:
        await arg;
    }
  }
}

async function edit(node: HTMLElement, text: string): Promise&lt;void&gt; {
  const overlap = getOverlap(node.textContent, text);
  await perform(node, [
    ...deleter(node.textContent, overlap),
    ...writer(text, overlap),
  ]);
}

async function wait(ms: number): Promise&lt;void&gt; {
  await new Promise(resolve =&gt; setTimeout(resolve, ms));
}

async function perform(
  node: HTMLElement,
  edits: Iterable&lt;string | null&gt;,
  speed: number = 60
): Promise&lt;void&gt; {
  for (const op of editor(edits)) {
    op(node);
    await wait(speed + speed * (Math.random() - 0.5));
  }
}

export function* editor(
  edits: Iterable&lt;string | null&gt;
): Generator&lt;(node: any) =&gt; number, void, unknown&gt; {
  for (const edit of edits) {
    yield node =&gt; requestAnimationFrame(() =&gt; (node.textContent = edit));
  }
}

export function* writer(
  [...text]: string,
  startIndex: number = 0,
  endIndex: number = text.length
): Generator&lt;string, void, unknown&gt; {
  while (startIndex &lt; endIndex) {
    yield text.slice(0, ++startIndex).join('');
  }
}

export function* deleter(
  [...text]: string | null,
  startIndex: number = 0,
  endIndex: number = text.length
): Generator&lt;string, void, unknown&gt; {
  while (endIndex &gt; startIndex) {
    yield text.slice(0, --endIndex).join('');
  }
}

export function getOverlap(start: any, [...end]: Iterable&lt;any&gt;): number {
  return [...start, NaN].findIndex((char, i) =&gt; end[i] !== char);
}
```

Mostly, I followed VSCode's advice on hover to type it &amp; some logic.

However, it gives me an error saying:

&gt; typical.ts(61,3): semantic error TS2488: Type 'string | null' must have a '[Symbol.iterator]()' method that returns an iterator.

Line 61 is `deleter()` function's `[...text]`, i.e,:

```ts
[...text]: string | null,
```

How do I solve this?
## [5][Gamedev Patterns and Algorithms in Action with TypeScript. Entity Component System](https://www.reddit.com/r/typescript/comments/hwesaz/gamedev_patterns_and_algorithms_in_action_with/)
- url: https://medium.com//entity-component-system-in-action-with-typescript-f498ca82a08e?source=friends_link&amp;sk=e02d1158e11b10330edcb8965b05cd6d
---

## [6][Typescript Mocha Testing, How to solve TypeError: is not a function?](https://www.reddit.com/r/typescript/comments/hwrhwn/typescript_mocha_testing_how_to_solve_typeerror/)
- url: https://www.reddit.com/r/typescript/comments/hwrhwn/typescript_mocha_testing_how_to_solve_typeerror/
---
I'm getting the error `TypeError: is not a function` for several Typescript methods I'm testing using mocha. All the importing/exporting for the mocha test seems to be working fine. I'm receiving the error for the specific methods referenced in the code in my StackOverflow post. 

[https://stackoverflow.com/questions/63064952/typescript-mocha-testing-how-to-solve-typeerror-is-not-a-function](https://stackoverflow.com/questions/63064952/typescript-mocha-testing-how-to-solve-typeerror-is-not-a-function)

Any advice would be greatly appreciated!
## [7][How to test process.on() using Mocha](https://www.reddit.com/r/typescript/comments/hwnsbt/how_to_test_processon_using_mocha/)
- url: https://www.reddit.com/r/typescript/comments/hwnsbt/how_to_test_processon_using_mocha/
---
Apologies if this is a naive question. I'm relatively new to Typescript and Mocha testing. I have the following question posted on StackOverflow about how to test \`process.on()\` using mocha testing, so that I may increase code coverage:

[https://stackoverflow.com/questions/63062843/how-to-test-process-on-using-mocha](https://stackoverflow.com/questions/63062843/how-to-test-process-on-using-mocha)

Any advice would be greatly appreciated!
## [8][How can I solve this Object is possibly 'null' error when using array destructuring?](https://www.reddit.com/r/typescript/comments/hwd5d6/how_can_i_solve_this_object_is_possibly_null/)
- url: https://www.reddit.com/r/typescript/comments/hwd5d6/how_can_i_solve_this_object_is_possibly_null/
---
I have an API request helper function with the following return type:

     Promise&lt;[APIError, null] | [null, ResponseData]&gt;

APIError and ResponseData are custom types that I have defined elsewhere.

My understanding of this return type is that only two scenarios are possible:

* The first item in the array is not-null, and the second item is null.
* The first item in the array is null, and the second item is not-null.

The idea behind this type signature is that we would use destructuring to get the API result, and always have to acknowledge the scenario that the request failed. If we ignored the first argument when destructuring we would get an unused variable linter error.

    const [errors, responseData] = await apiRequest();
    
    if (errors != null) {
        ...
        return;
    }
    
    console.log(responseData.message);

In the code example above I have returned early if errors is not null, this means that beyond this point responseData is guaranteed to not be null. However I get the following error when trying to access the message property of responseData.

&gt;Object is possibly 'null'.ts(2531)

If I remove the destructuring it works as expected with no error.

    const apiResponse = await apiRequest();
    
    if (apiResponse[0] != null) {
        ...
        return;
    }
    
    console.log(apiResponse[1].message);

Is there any way around this, I understand I could use a 'non-null assertion operator' but I would rather not have the added overhead of doing this.
## [9][TypeScript complaining about class improperly implementing an interface, but it seems to be fine - what am I missing?](https://www.reddit.com/r/typescript/comments/hw4kih/typescript_complaining_about_class_improperly/)
- url: https://www.reddit.com/r/typescript/comments/hw4kih/typescript_complaining_about_class_improperly/
---
Hi everyone,

I have an interface that I'm using to identify classes that interact with REST APIs. At the lowest level, the interface looks something like this:

    interface Foo {
        config: Config;

        constructor(config: Config);
    }

To reuse some general functionality and reduce boilerplate between clients, I have an abstract class similar to this:

    abstract class Bar {
        config: Config;

        constructor(public config: Config) {}

        getGreeting(): string {
            return `Hello ${this.config.subject}!`;
        }
    }

So the API client classes extend this abstract class and implement a child type of `Foo` (depending on the kind of API/entity it represents - I excluded this layer for brevity since it's irrelevant):

    class Bam extends Bar implements Foo {
        sayHello(): string {
            return this.getGreeting();
        }
    }

When I instantiate `Bam` and log the value returned by `sayHello()`, it is what I expect it to be. However, TypeScript is complaining that:

    Class 'Bam' incorrectly implements interface 'Foo'.
        Type 'Bam' provides no match for the signature 'constructor(config: Config): any'

What am I missing about this? `Bam` inherits the injected dependency, and I'm not overriding the constructor in any of the children - the signature in the abstract class is the same as in the interface.
## [10][Any way to type this array in a generic function?](https://www.reddit.com/r/typescript/comments/hw3d9f/any_way_to_type_this_array_in_a_generic_function/)
- url: https://www.reddit.com/r/typescript/comments/hw3d9f/any_way_to_type_this_array_in_a_generic_function/
---
I wrote this function with a generic type:

`export const listToArray: &lt;T&gt;(list: {[key: string]: T[]}) =&gt; T[] = (list) =&gt; {`  
 `let res: any[] = [];`  
 `for (let key in list) {`  
 `res = res.concat(list[key])`  
`}`  
 `return res;`  
`}`

I'm not too experienced in TS but I can't really see a way to type the res array correctly since it's value depends on the generic. If there any way to type the res array that I'm simply overlooking?

Thanks!
## [11][Given the type of a class , how can I get the type of the instances it creates ?](https://www.reddit.com/r/typescript/comments/hw1vm1/given_the_type_of_a_class_how_can_i_get_the_type/)
- url: https://www.reddit.com/r/typescript/comments/hw1vm1/given_the_type_of_a_class_how_can_i_get_the_type/
---
I have a module that exports a factory but also has as a side effect the registration of that factory to a diContainer . I can take the factory return type from the import via `ReturnType&lt;typeof factory&gt;` but I am supposed to get its return value from the diContainer , which is responsible for resolving its dependencies . The factory returns a class . How can I create the ts type that corresponds to the instance of the class?

^(Every get that I do on the diContainer has as type any . Can I somehow make the diContainer give me the correct type when I get on it ?)

Edit : For some strange reason the following works :

    import {diContainer} from "./diContainer";
    import {factory} from "./factory";
    
    type myClassInstance = ReturnType&lt;typeof factory&gt;["prototype"];
    
    let myClassInstance : myClassInstance;

But I can use `myClassInstance` for type on the instances only and not the class .
