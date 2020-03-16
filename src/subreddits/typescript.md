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
## [2][Trouble in finding modules with '@' paths, getting "cannot find module '@/common/utils'" despite the file definitely existing and being recognised by my IDE.](https://www.reddit.com/r/typescript/comments/fjiqzr/trouble_in_finding_modules_with_paths_getting/)
- url: https://www.reddit.com/r/typescript/comments/fjiqzr/trouble_in_finding_modules_with_paths_getting/
---
I need help, `tsc` and `ts-node` do not recognise the custom path aliases specified in my `tsconfig.json` file. My file is as follows:

    {
      "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "strict": true,
        "jsx": "preserve",
        "importHelpers": true,
        "moduleResolution": "node",
        "experimentalDecorators": true,
        "esModuleInterop": true,
        "allowSyntheticDefaultImports": true,
        "sourceMap": true,
        "baseUrl": ".",
        "types": ["webpack-env", "jest", "node"],
        "paths": {
          "@/*": ["src/*"]
        },
        "lib": ["esnext", "dom", "dom.iterable", "scripthost"]
      },
      "include": [
        "src/**/*.ts",
        "src/**/*.tsx",
        "src/**/*.vue",
        "tests/**/*.ts",
        "tests/**/*.tsx"
      ],
      "exclude": ["node_modules"]
    }

My IDE recognises the paths and I'm able to use the imported values as expected but when I compile any TS file or use `ts-node` on them, I get the error that the module is not recognised. Any help would be much appreciated, thanks fellas.

EDIT: **SOLVED** thanks to /u/moremattymattmatt, I had to install the `tsconfig-paths` npm package to use path aliasing.
## [3][Example of simple Typescript util npm package on Azure DevOps](https://www.reddit.com/r/typescript/comments/fjjp7r/example_of_simple_typescript_util_npm_package_on/)
- url: https://www.reddit.com/r/typescript/comments/fjjp7r/example_of_simple_typescript_util_npm_package_on/
---
I was wondering if anyone had a simple and good example/tutorial/course on how to set up an NPM package written in Typescript on Azure DevOps?

Pretty new at Azure DevOps, and the repository itself is no problem, but I don't know how to:

  1. set up pipelines
  2. how to deal with npm package versioning (since package.json needs to be updated)
  3. how to publish it to the Artifacts npm feed.
## [4][best way to use images from an API?](https://www.reddit.com/r/typescript/comments/fjfisf/best_way_to_use_images_from_an_api/)
- url: https://www.reddit.com/r/typescript/comments/fjfisf/best_way_to_use_images_from_an_api/
---
using angular 7 iirc and TS.  I have an API written in go and part of this API call is to return data as well as a base64 image.  In my TS side, I have to use a santizer and then a bypasssecurity method which just seems really odd.

is there a better way to do something like this?
## [5]["Lazy" IoC on TypeScript](https://www.reddit.com/r/typescript/comments/fj4lgt/lazy_ioc_on_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fj4lgt/lazy_ioc_on_typescript/
---
Recently I've created my first npm [package](https://github.com/ceoro9/ioc-ts). It's a "lazy" inversion of control container. And yes, I heard about inversify, the things I don't like in it, that it's too bloated and does not support no-decorator constructor injection, like in NestJS. And btw it seems like inversify lacks activity. Anyway, for me the main goal was learning, but I'll be very glad if somebody can rate my code and leave some feedback, for now I'm just a junior dev, so for me any feedback is really valuable. May be somebody would like to help in writing more tests and docs, I welcome everybody.

Also I don't know if I can ask for this, but tomorrow I am going to the job interview and to make a good impression on a interviewer, I'll be very glad if somebody can put a star on my github repo. I'd really appreciate this.
## [6][Global typings only working with triple slash directive](https://www.reddit.com/r/typescript/comments/fj5a8u/global_typings_only_working_with_triple_slash/)
- url: https://www.reddit.com/r/typescript/comments/fj5a8u/global_typings_only_working_with_triple_slash/
---
&amp;#x200B;

https://preview.redd.it/lqyma3hfjvm41.png?width=249&amp;format=png&amp;auto=webp&amp;s=8e726199e506b541a19d79d9a53e1273d92b82a7

react-app-env.d.ts

    /// &lt;reference types="react-scripts" /&gt;
    /// &lt;reference path="types.d.ts" /&gt;

types.d.ts

    type Response = {
      token: string;
      email: string;
      firstName: string;
      lastName: string;
    }

only when bringing in types.d.ts with /// works

otherwise I keep getting error that the types are not found

What should be the correct way of mentioning global types?
## [7][Put together an illustration to clear up my own confusion as to why the Union and Intersection operators are named as such](https://www.reddit.com/r/typescript/comments/fisz4q/put_together_an_illustration_to_clear_up_my_own/)
- url: https://www.reddit.com/r/typescript/comments/fisz4q/put_together_an_illustration_to_clear_up_my_own/
---
The naming of the Union and Intersection operators had always felt curious to me, as it seemed in direct contrast to their respective functions. Why do we call `&amp;` the Intersection operator when the declared object possesses not only the mutual members, but the members of all concerning types? Why do we call `|` Union when the declared object can only reference the mutual members? Recently stumbled upon an SO question asking the same things and decided to finally sit down and reason through it. While the existing answers are great and helped a good deal, I figured it might be of value to put my own understanding down (in good old graphic form) for others who may fall for the same misinterpretive line of thinking.

SO answer:
https://stackoverflow.com/a/60686242/1107110

Illustration:
https://i.stack.imgur.com/fY4gL.png

Feel free to share your thoughts. Does it make sense? Does it clear things up? Does it confuse you more? All feedback welcome. 🙂
## [8][Correct way of storing interfaces in TypeScript](https://www.reddit.com/r/typescript/comments/fi8v85/correct_way_of_storing_interfaces_in_typescript/)
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
## [9][Designing the perfect TypeScript schema validation library](https://www.reddit.com/r/typescript/comments/fhwj5m/designing_the_perfect_typescript_schema/)
- url: https://vriad.com/blog/zod/
---

## [10][Something like "Nested" Pick](https://www.reddit.com/r/typescript/comments/fi2za5/something_like_nested_pick/)
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
## [11][Up and running in a breeze with Firebase](https://www.reddit.com/r/typescript/comments/fhxjcr/up_and_running_in_a_breeze_with_firebase/)
- url: https://blog.pragmatists.com/up-and-running-in-a-breeze-with-firebase-72ca889d22cb
---

