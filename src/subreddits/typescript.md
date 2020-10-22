# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][How can I get the keys from a Map.keys() iterator as a type?](https://www.reddit.com/r/typescript/comments/jfzkor/how_can_i_get_the_keys_from_a_mapkeys_iterator_as/)
- url: https://www.reddit.com/r/typescript/comments/jfzkor/how_can_i_get_the_keys_from_a_mapkeys_iterator_as/
---
```
const entities = new Map&lt;string, []&gt;([
  ["hero", []],
  ["zombies", []],
  ["bullets", []],
  ["text", []],
]);

const mapIter = entities.keys();

for (const keys of mapIter) {
  console.log(keys); // hero, zombies, bullets text
}
```
I want to generate `type EntityKeys = "hero" | "zombies" | "bullets" | "text";` programatically. 

I want to use `entities.get("hero")` and have `"hero"` type-checked, which won't work with `Map&lt;string, []&gt;`. I can use `new Map&lt;EntityKeys, Entity[]&gt;` using `type EntityKeys =...` but then it is not dynamic and I have to create `type EntityKeys` manually.

Cheers!
## [3][TypeScript JSX transform vs Babel JSX transform - pros and cons](https://www.reddit.com/r/typescript/comments/jffysy/typescript_jsx_transform_vs_babel_jsx_transform/)
- url: https://www.reddit.com/r/typescript/comments/jffysy/typescript_jsx_transform_vs_babel_jsx_transform/
---
I'm building a project with both TypeScript and Babel (and React) and I see that both TypeScript and Babel support the same JSX transform step. Either TypeScript can emit transformed JSX in JS code directly, or it can preserve the JSX in JSX files so Babel can handle it.

It looks like each compiler has their own spec compliant implementation of this transformation. I cannot find any details as to which one is better to use and why. There must be some differences (speed, reliability, spec compliance, support, bug fixes, type checking, etc) but I cannot seem to identify any.

Given the little information I have, it seems like Babel is the "canonical" JSX transformer as React specifically works with Babel to implement the spec. From the TS JSX PR I see that they are just trying to copy what Babel implemented. This leads me to believe that I should use the TS JSX preserve setting and let Babel handle it as it'll likely be more spec complainant and more stable.

Does anyone have any additional information to add here or know of anything which I should consider in making this decision? Thanks!

----

Edit: Follow up. After tons of very helpful comments I figured I'd share what I decided on. I am now using `babel` for all TypeScript compiling needs. `@babel/preset-env`, `@babel/preset-typescript` and `@babel/preset-react` specifically. To get builds to fail due to type errors and to see errors in the console during development, I'm using [fork-ts-checker-webpack-plugin](https://github.com/TypeStrong/fork-ts-checker-webpack-plugin). I now no longer need `tsc` for production or development at all and am only using it for debugging if I purely want to run a command to quickly see all type errors in the project.
## [4][Matching a destructured type based on transformation of argument type](https://www.reddit.com/r/typescript/comments/jfpy8t/matching_a_destructured_type_based_on/)
- url: https://www.reddit.com/r/typescript/comments/jfpy8t/matching_a_destructured_type_based_on/
---
I've been using typescript for a little while but haven't had the real need to dive into anything past simple typing until now. Basically, I have a function that takes in an obj map of promises and want the typing to understand it returns a new map of the resolved types.

The closest I've come is getting it to return the relevant keys where each is a union of the resolved promise types.

    type Unpromised&lt;T&gt; = T extends Promise&lt;infer R&gt; ? R : T;
    type ValueOf&lt;T&gt; = T[keyof T];
    type PromiseObjRet&lt;T&gt; = { [field: string]: Unpromised&lt;ValueOf&lt;T&gt;&gt;; // how can we make this respective?

    async function handlePromiseObj&lt;T&gt;(obj: T): Promise&lt;PromiseObjRet&lt;T&gt;&gt; {
        ... // resolves promises as a batch and maps them back to an obj by the original obj key
    }

I'm looking for:

    const { a, b } = handlePromiseObj({ a: Promise.resolve('hi'), b: Promise.resolve(6) });

to have a: string, b: number. 

Right now I'm getting a: string | number, b: | number.

In short, I'm looking to have types setup to be able to transform via typing 
    {
        a: Promise&lt;string&gt;,
        b: Promise&lt;number&gt;
    }

to

    {
        a: string,
        b: number
    }
## [5][Possible to get typeof generic type parameter?](https://www.reddit.com/r/typescript/comments/jfq79i/possible_to_get_typeof_generic_type_parameter/)
- url: https://www.reddit.com/r/typescript/comments/jfq79i/possible_to_get_typeof_generic_type_parameter/
---
I am writing a simple de/serializer that works for any type of \`State\`. Is there any way to get the concrete types of a generic type parameter T? Code sample is shown below.

[https://pastebin.com/84YabPtS](https://pastebin.com/84YabPtS)

ps: The code does not render in reddit, so have to post it somewhere else.
## [6][How can I resolve my async function so &lt;Promise&gt;Zombie becomes Zombie?](https://www.reddit.com/r/typescript/comments/jfhglw/how_can_i_resolve_my_async_function_so/)
- url: https://www.reddit.com/r/typescript/comments/jfhglw/how_can_i_resolve_my_async_function_so/
---
I am stuck, my `init` static method insists I return `Promise&lt;Zombie&gt;`, but it is meant to be `Zombie`, it should be resolved by the `awaitPromise.all`? The consumers of the class all throw errors now as it is no longer a `Zombie`

I think there is some 'gotcha' I am not seeing here? Cheers!

https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABAGzgQwCYEkC2aDmApgBTAzKEAKaUAFgFyIDOUATjGPgJSOWtw4YTQgB4AEgBUAsgBlcBQgFEKOQmCgA+RAG8AUIkQQELRDDxFlhVesQBeRGEIB3RPKLEuAbn2nzSlWpQAHRMrBB2iGQU1HTeBqyEUCCsSI4ufAJCopKybv5WgRrExAlMcMgAboQANIgJAFaE0Fx2WnoGBmYKltbBCKiYER6tdYRllSRdFgHqXj6dfj2BQQiErPysEQ1NUHGIAL5z+7q6hAAeAA5wrFCGyGhMTIgAWgIARjCEOj5ThHtgaFU3h8RjALFYIGg12Iv0YOTkixmUFqANUjHBHG43w6iDoQiCvwivz2BjxTCCqK+9kpe2OPhYNBg4QeAE9IKYwDAoB5ePxBMIRK8cB9CG15oZjLcANoAL3enzyAF0ImgnGguYgMvzCEE0MhkMQpeKDANsH5iAAiB7CKBMAD0cuFnwAbAAWAC0ZAByCCF04Fq41XFirm4oSSRSDmcL3lJEdIrytQttDUrBZAdpujpoEgsAQhgSNEIQpFPJjTq+7QlYNu8c+ERLnwJnO5c3iiWSSDrfyzJ1B4x1qHwxAghagxdjHjmQA

```typescript
function loadImage(filePath: string): Promise&lt;HTMLImageElement&gt; {
  const imageElement = new Image();
  imageElement.src = filePath;
  return new Promise&lt;HTMLImageElement&gt;((resolve, reject) =&gt; {
    imageElement.onload = () =&gt; resolve(imageElement);
    imageElement.onerror = reject;
  });
}

export class Zombie {
  image;
  name;

  constructor(image: HTMLImageElement, name: string) {
    this.image = image;
    this.name = name;
  }

  static async init(): Promise&lt;Zombie&gt; {
    const [zombieImage] = await Promise.all([
      loadImage("assets/zombie64-final.png"),
    ]);

    return new Zombie(zombieImage, "henry");
  }
}

function createZombie(): Zombie {
  const zombie = Zombie.init();
  return zombie;
}

console.log(createZombie());
```
## [7][View the API exposed by a given module](https://www.reddit.com/r/typescript/comments/jf4c4t/view_the_api_exposed_by_a_given_module/)
- url: https://www.reddit.com/r/typescript/comments/jf4c4t/view_the_api_exposed_by_a_given_module/
---
Is there a tool or convention to view/document the API exposed by a given module?

```typescript
/// example/index.ts

export * from "./core";
export * from "./util";
```

```typescript
/// example/core/index.ts

export class Application {
  private constructor (...args: unknown[]) { … }
  
  public get id(): string { … }
  
  public async start(): Promise&lt;void&gt; { … }
  
  public static create(process: NodeJS.Process): Application {
    // …
    
    return new Application(…);
  }
}
```

```typescript
/// example/util/index.ts

export const VERSION = "0.0.0" as const;
```

For example, the above would expose a very small public API, consisting of:
- `Application`
- `Application#id`
- `Application#start()`
- `Application.create()`
- `VERSION`

Are TypeScript definition files what I'm looking for?
## [8][InversifyJS - How can I use container.get in an object literal?](https://www.reddit.com/r/typescript/comments/jf4axz/inversifyjs_how_can_i_use_containerget_in_an/)
- url: https://www.reddit.com/r/typescript/comments/jf4axz/inversifyjs_how_can_i_use_containerget_in_an/
---
I'm not sure if this is the place to ask this, but it's the best place I could find. I'm new to dependency injection with TS but did use it in C# once. When I try doing this:

    export let COMMAND_MAP = {
        "join": container.get&lt;ChannelJoiner&gt;(TYPES.ChannelJoiner)
    }

I get a runtime error that "get" is not defined. What am I doing wrong? Am I allowed to do this? I can send other related files if necessary.  


Edit: I believe I have found the problem, but I am still unsure on how to fix it. The object above gets initialized before the container. Is there any way I can force the container to be initialized first?
## [9][Good book or other resource for learning how to define your models?](https://www.reddit.com/r/typescript/comments/jex71u/good_book_or_other_resource_for_learning_how_to/)
- url: https://www.reddit.com/r/typescript/comments/jex71u/good_book_or_other_resource_for_learning_how_to/
---
I’m comming from a c# background and currently have a huge model that is hard to understand and does nothing else but set the types. My model is a huge collection of classes inside classes that don’t have anything else but some properties(no constructor, methods getters or setters). I’m trying to learn how to improve this but can’t seem to find a good book or article
## [10][How to map an interface to another one](https://www.reddit.com/r/typescript/comments/jexpj2/how_to_map_an_interface_to_another_one/)
- url: https://www.reddit.com/r/typescript/comments/jexpj2/how_to_map_an_interface_to_another_one/
---
I'm sorry if this is improperly formatted, I copied this from another subreddit I posted it.
Okay folks, Idk where to ask (I got downvoted in [stack](https://stackoverflow.com/questions/64443221/how-to-map-different-interfaces-to-each-other-in-typescript)) but the question is clear there, it's more of a typescript question, I'm working on a store which has \`ICategory\`, \`IProduct\`, ...etc and I generated the GraphQL schema with graphql codegen, the schema has \`Category\`, \`Product\`, ...etc.

The ICategory is on the following format

```

export interface IBaseCategory {

id: string;

type: string;

name: string;

slug: string;

image: string|null;

items?: number;

parent?: this | null;

children?: this\[\];

customFields: ICustomFields; }
```

    and the GraphQl one is on the format of the server, names of attributes are slightly different, but it has children with has edges which have nodes, etc (GraphQL basic schema).
    Now, I've created a class for each type that queries that server and returns a `Category`, `Product`, ..etc (GraphQL format). and I want them to be mapped to the other format (The theme format, `ICategory, ...) so that I don't remove or alter the theme interfaces except as little as possible, details are in stackoverflow link.
    
    Thanks in advance.
## [11][Generate Google Presentation from Wikipedia article](https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/)
- url: https://www.reddit.com/r/typescript/comments/je6qg8/generate_google_presentation_from_wikipedia/
---
[GSlides Maker](https://github.com/vilmacio/gslides-maker) is a open-source project for create google presentations from wikipedia content like a robot. Try it [CLI version](https://github.com/vilmacio/gslides-maker) and help me with the [web application](https://gslidesmaker.com/). I hope you like it.

https://preview.redd.it/98yuu3ceb3u51.png?width=681&amp;format=png&amp;auto=webp&amp;s=5259f69c148faeab52c9a10af6490d366bfec5c5
