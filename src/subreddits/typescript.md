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
## [2][How can I solve this Object is possibly 'null' error when using array destructuring?](https://www.reddit.com/r/typescript/comments/hwd5d6/how_can_i_solve_this_object_is_possibly_null/)
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
## [3][Gamedev Patterns and Algorithms in Action with TypeScript. Entity Component System](https://www.reddit.com/r/typescript/comments/hwesaz/gamedev_patterns_and_algorithms_in_action_with/)
- url: https://medium.com//entity-component-system-in-action-with-typescript-f498ca82a08e?source=friends_link&amp;sk=e02d1158e11b10330edcb8965b05cd6d
---

## [4][TypeScript complaining about class improperly implementing an interface, but it seems to be fine - what am I missing?](https://www.reddit.com/r/typescript/comments/hw4kih/typescript_complaining_about_class_improperly/)
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
## [5][Any way to type this array in a generic function?](https://www.reddit.com/r/typescript/comments/hw3d9f/any_way_to_type_this_array_in_a_generic_function/)
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
## [6][Given the type of a class , how can I get the type of the instances it creates ?](https://www.reddit.com/r/typescript/comments/hw1vm1/given_the_type_of_a_class_how_can_i_get_the_type/)
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
## [7][Converting a generic React Component to TypeScript throws error](https://www.reddit.com/r/typescript/comments/hvutjs/converting_a_generic_react_component_to/)
- url: https://www.reddit.com/r/typescript/comments/hvutjs/converting_a_generic_react_component_to/
---
I am trying to port https://github.com/catalinmiron/react-typical to TypeScript. However, I am facing some issues.

Here's the screenshot with errors in VSCode:

https://i.stack.imgur.com/IKZK2.png

Here's the same code for brevity:

```ts
import React from 'react'
import { type, type as loopedType } from '@camwiegert/typical'
import styles from './styles.module.css'

type Props = {
	steps: Array&lt;any&gt;
	loop: number
	className: string
	wrapper: React.Component
}

const Typical: React.FC&lt;Props&gt; = ({ steps, loop, className, wrapper = 'p' }) =&gt; {
	const typicalRef = React.useRef&lt;HTMLElement&gt;(null)
	const Component: string = wrapper
	const classNames: string[] = [styles.typicalWrapper]

	if (className) {
		classNames.unshift(className)
	}

	React.useEffect(() =&gt; {
		if (loop === Infinity) {
			type(typicalRef.current, ...steps, loopedType)
		} else if (typeof loop === 'number') {
			type(typicalRef.current, ...Array(loop).fill(steps).flat())
		} else {
			type(typicalRef.current, ...steps)
		}
	}, [typicalRef])

	return &lt;Component ref={typicalRef} className={classNames.join(' ')} /&gt;
}

export default React.memo(Typical)
```

I am unable to write type for `Component`.

I tried doing the following too:

```ts
const Component = React.Component | string
```

But it says `'Component' refers to a value, but is being used as a type here. Did you mean 'typeof Component'?` near `return &lt;Component .../&gt;` with underline over `Component`.

I am also unable to convert the `typicalRef` as `typicalRef.current` always throws error by showing red squiggly lines under it. Same thing with `flat()` as well as `classNames.join(' ')`.

I am losing my brain over it. Can't seem to figure it out. Would love any pointers?
## [8][Need help with tsconfig](https://www.reddit.com/r/typescript/comments/hvwcsj/need_help_with_tsconfig/)
- url: https://www.reddit.com/r/typescript/comments/hvwcsj/need_help_with_tsconfig/
---
Hi everyone! I cannot find solution anywhere so I decided to post here.

I've made simple express API with typescript.

I've got this eslint error on \*.test.ts files and \_\_mocks\_\_ dir:

&gt;Parsing error: "parserOptions.project" has been set for u/typescript-eslint/parser.  
&gt;  
&gt;The file does not match your project config: &lt;filename\_here&gt;  
&gt;  
&gt;The file must be included in at least one of the projects provided.

Where &lt;filename\_here&gt; is the name of file giving me this error.

To reproduce just clone the repo, install deps and open any of these files.

I think the problem is in tsconfig, jest config or eslint config. No idea how to fix it and i'm stuck.

Repo:

[https://github.com/nemmtor/love-matcher-api](https://github.com/nemmtor/love-matcher-api)
## [9][Typing the iterator in a for/of loop](https://www.reddit.com/r/typescript/comments/hvz27q/typing_the_iterator_in_a_forof_loop/)
- url: https://www.reddit.com/r/typescript/comments/hvz27q/typing_the_iterator_in_a_forof_loop/
---
TSC is not letting me type annotate the iterator `sentence` in the below loop. I switched to using `as` syntax to assert what `sentence` should be. Is this the best that can be done in this loop?

            for (const sentence of buildOrchestrator.qualifiedSentences) {
              const folderExists: boolean = checkForExistingFolder(sentence as Sentence);
    
              /* eslint-disable no-await-in-loop */
              if (folderExists === false) {
                await buildOrchestrator.makeFolderAndAudioFiles(sentence as Sentence);
              }
            }
## [10][Type 'string' not assignable to a custom type](https://www.reddit.com/r/typescript/comments/hvre75/type_string_not_assignable_to_a_custom_type/)
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
## [11][How to map API reponse property name to internal property name with different casing?](https://www.reddit.com/r/typescript/comments/hvxhg3/how_to_map_api_reponse_property_name_to_internal/)
- url: https://www.reddit.com/r/typescript/comments/hvxhg3/how_to_map_api_reponse_property_name_to_internal/
---
I googled but couldn't seem to find the solution. 

Let's say I have this response in Pascal case returning from an external API. 

    {
        "Name": "hey",
    }

In my class, I am using camelCase for the properties.

    export interface Person {
      name: string;
    }
    

How do I map them? 

    const response = await this.httpService
            .get&lt;Person&gt;(getPersonAPIUrl, { headers: requestHeaders })
            .toPromise();

Right now, I will get **undefined** when I try to get the value of **response.data.name**.
