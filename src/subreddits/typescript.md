# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Is there an 'Exact&lt;T&gt;' advanced type? Something that gurantees you can only put in the exact attributes required and no more?](https://www.reddit.com/r/typescript/comments/i8vxz2/is_there_an_exactt_advanced_type_something_that/)
- url: https://www.reddit.com/r/typescript/comments/i8vxz2/is_there_an_exactt_advanced_type_something_that/
---
Eg. I have the following object:

    type User = {
        id: string
        email: string
        password: string
    }

    type UserProfile = User &amp; { name: string, dob: string }

    function insert(user: User) {
        UserDB.insert(user)
    }

So I have a User, a UserProfile, and an insert function that inserts Users into the DB.

Now my problem is that the insert function connects to a NoSQL DB. So I can technically pass in a UserProfile, and it will end up inserting the extra attributes into my NoSQL DB. Over time my NoSQL DB might end up with a lot of extra unnecessary attributes.

I was wondering if there was an Exact&lt;T&gt; type that ensures you only have User.id, email &amp; password added in and nothing more.
## [3][Validating object's type at runtime?](https://www.reddit.com/r/typescript/comments/i8yk6i/validating_objects_type_at_runtime/)
- url: https://www.reddit.com/r/typescript/comments/i8yk6i/validating_objects_type_at_runtime/
---
For a project I am currently working on, an SDK for consuming a backend API (written in Go), I would like to implement runtime type checking. As I've worked intensively with Typescript over the last couple of years, I know type safety is not guaranteed at runtime, as all type annotations will get lost during the compilation stage.

I have some projects using Typescript as a production dependency, to for example generate custom JSON decoders to allow code generation to take place based on Typescript interfaces: https://github.com/PicnicSupermarket/aegis. 

A thought that came to mind, was using Typescript as a dependency and use it to check against an interface, but as I'm not really familiar with the structure of the Typescript compiler, I don't know what's possible.

So my question is: **is it possible to take a Javascript object, together with a Typescript interface, and validate whether the Javascript object satisfies the Typescript interface**? And if that can be done by calling the Typescript compiler as a production dependency somehow.

To give some further context, this is the function where results pass through in the SDK I am working on:

```
const throwOrReturn = &lt;T&gt;(result: APIResult&lt;T&gt;): T =&gt; {
    if (result.error) {
        throw new Error(result.error);
    }
    // We can assume that data is valid (type T) if no error was found
    return result.data as T;
};

interface APIResultSuccess&lt;T&gt; { data: T; error?: undefined; }
interface APIResultFailure { data?: undefined; error: string; }
type APIResult&lt;T&gt; = APIResultSuccess&lt;T&gt; | APIResultFailure;
```

Before `return result.data as T;` I would like to check whether `result.data` is valid if assigned type `T` at runtime. So currently I am using generics, but taking the checking one layer up, where I have access to the actual interface, instead of a generic, would not be an issue.

An example of how the `throwOrReturn` is implemented:

```
interface ReturnValues {
    // Some interface code
}

const exampleCall = async (config: Config): Promise&lt;interface&gt; =&gt; {
    const res = await get&lt;interface&gt;(config, GetEndpoints.EndpointWeNeed); // Wrapper for making API calls and 4xx/5xx error handling 
    return throwOrReturn(res);
};

export default exampleCall;
```
## [4][Type-safe GraphQL APIs](https://www.reddit.com/r/typescript/comments/i8wz3u/typesafe_graphql_apis/)
- url: https://www.reddit.com/r/typescript/comments/i8wz3u/typesafe_graphql_apis/
---
Representing GraphQL, a dynamic protocol, in a statically typed language is inherently difficult. [GraphQL code generator](https://graphql-code-generator.com/) can help with this a bit, but isn't ideal.

I have devised a solution to this problem that allows all GraphQL APIs to be represented in Typescript. It is based around this interface:

    interface GraphQLObject&lt;Name extends string, T&gt; {
        // The GraphQL name of the type.
        name: Name;
        // The query string to get the type, without surrounding braces.
        query: string;
        // The conversion function from the GraphQL JSON response to the Typescript type.
        convert: (data: any) =&gt; T;
    }

It represents a subset of fields on the GraphQL type `Name`. An example usage would be:

    interface Image {
        url: string;
        width: number;
        height: number;
    }
    
    const image: GraphQLObject&lt;"Image", Image&gt; = {
        name: "Image",
        query: "url width height",
        convert: data =&gt; ({
            url: data.url,
            width: parseInt(data.width),
            height: parseInt(data.height),
        }),
    };

Fields that contain other objects can be represented via generics, and fields that take arguments can be represented in parameters.

    interface Query&lt;Image&gt; {
        imageByName: Image;
    }

    function query&lt;Image&gt;(image: GraphQLObject&lt;"Image", Image&gt;, id: string): GraphQLObject&lt;"Query", Query&lt;Image&gt;&gt; {
        return {
            name: "Query",
            query: `imageByName(id:${id}){${image.query}}`,
            convert: data =&gt; ({ imageByName: image.convert(data.imageByName) }),
        };
    }

This way you can request any set of image fields through the query function, and the id parameter can be a literal or variable.

This can be complemented with some useful helper functions:

    // Query no data.
    function no_data&lt;Name extends string&gt;(name: Name): GraphQLObject&lt;Name, void&gt; {
        return {
            name,
            query: `__typename`,
            convert: () =&gt; undefined
        };
    }

    // Combine multiple field subsets together that query the same GraphQL object.
    function combine&lt;Name extends string, A, B&gt;(
        ...objects: [GraphQLObject&lt;Name, A&gt;, GraphQLObject&lt;Name, B&gt;]
    ): GraphQLObject&lt;Name, A &amp; B&gt;;
    // And so on...
    
    function combine&lt;Name extends string&gt;(
        ...objects: GraphQLObject&lt;Name, unknown&gt;[]
    ): GraphQLObject&lt;Name, any&gt; {
        return {
            name: objects[0].name,
            query: objects.map(o =&gt; o.query).join(" "),
            convert: data =&gt; Object.assign({}, ...objects.map(o =&gt; o.convert(data)))
        };
    }

Using this I would love to make a library that automatically generates GraphQL objects for each field of each type from a GraphQL schema (similar to GraphQL code generator, but using this method). The individual fields can then be combined by the user to form more complex queries, while remaining fully typesafe. For ergonomics the library could also provide `Full*` type aliases which represent _all_ the fields of a type (e.g. `type FullImage = ImageURL &amp; ImageWidth &amp; ImageHeight`).

However, I lack the time and expertise to do it myself, so if anyone is willing to give it a shot I would be very grateful.
## [5][Type annotation erroneously overrides generic](https://www.reddit.com/r/typescript/comments/i8zupk/type_annotation_erroneously_overrides_generic/)
- url: https://www.reddit.com/r/typescript/comments/i8zupk/type_annotation_erroneously_overrides_generic/
---
I'm currently trying to refactor a function's signature to remove the need for loads of overloads. I've done this by creating a conditional generic type to handle the same cases that the overloads did.

The function works fine but for some reason adding annotations to constrain the return type from `unknown` to something else is giving me errors that didn't exist, as if the function is incorrectly inferring `T` from the annotation and I just can't figure out why. There are some times in my project where I need to declare a variable before assigning it and it just breaks.

This is the function's signature:

    const createMatrix = &lt;D extends number, T&gt;(
        dimensions: D,
        initialValues: T = null
    ): Matrix&lt;D, T&gt;
    
    type Matrix&lt;D extends number, T&gt; = D extends 1
        ? T[]
        : D extends 2
        ? T[][]
        // ...

And a TS playground link with the full code and old version [is here](https://www.staging-typescript.org/play?#code/PTAEBUE8AcFNQIYDskHsAuD0EtVNAO4BOeA5gDaSioBusRR2AJrAM6imxL3YDGh2dAAtQ9EkVYAoSSFABRAB5xe6WE1AAjWEIQ1cRAFwywAQSZNsSUqHQx4yNJhx526VKBoJGCDeTahWIVQAV3J1LVBeVABbaCxsX3gCQRFheAAzYKQVXCR2Ilh0YKIkYwDsUiQsYtgAGkiEYNZLazQbO1EGVAlpWRMVYIRyTW1dfSM+1mbKlsR8BwwsNQ8vbB8-G3c00AKikpXyYPhUdJshDKycvAam-2RO8QA6MvBzji4efnBQbHZiMkoPyQ6TEyzcIzmVBOoAA5F4iAgqARzgVQCxoMJQAAeUAAERhhBR8EEASCoSYZQilhBDDB7jwgO2GgQrHgtjg1FO224rFU6nhiMeoAASrBorRZtsFk5cgFUOQ6FJZNtoCREtFNMF0DsxbR-Oz4KwECDbISuCtvIlEFMKkholxtTpoHA8oh0qoiJECvE8M8ZAAqf2SUD+kWFYquhCgaJYRgKUAACmioRwAFoLPa8rkhogGIiAJQCTGQEKeljNArqDNcZouObqTyHNjPEPB0MAATiCI11azddToAABrjBwYIG9y9hK2jsJna67odsY+g423QJ2vAgNZZBGtyAA1IZHdgD17wRtHM5YUQIXgiaCoSzayxneDLuNA3dDX761BC3GwOkjTkOgrjuIOSChOQg4th2uwRuwJjRrG2DxtCw6DjOc65Kw9TJJisC3vej5IM+SBfuQP7qOC2yDjuOBDIeTasJhXZboU9CwcAkhRHk2q8N6qgALIofGAC82K4qICiqEgTDsJB0RaEQ9TgAAfAmwagNpvbzqwY64rUWnafRe5MceY7fAAPqAkHkMMEl2eQkj5mOIkrqhWKGRAamgGJvkAN7GZELj8cUBSkbis41rkAAyXCkJiEm6ThADcwW8byOoxjuVhRdhdbJdFfaugOACM6XadpmXatwaisKKvDFPOfnZQguWkPlMV1r5AAM6UZaFkThQ67kfhJJh5pACZNbSkXFfO8VWMI+aPOk2D2QmpmMUebD5gNVUhXxoDrVU5BjahrV1fJjXNbkwWHQA-MNc3oBdCiPDG0AJgmhb+Q9h2AwJhHCaJCYFDl5F5QtOH1NtB67aw+YA4dyOA1VY6zRFb2iQdVXwfsp1DO91qgO9XmqWp6UAL4DQaZOiV50myfJtnBEp9CU61UmwDJXCs2VwXPeAADaAC6wUGcz-PsAATELEDi+Lkt4tLcnsAAzAroti0rEuHVLvMs+wAAs2t63rKvIJAFu63bA01TYbDoGVY6Kcp4utcDSzvQmZX7aAshYgOBDdAA1ie7zcIwXw-OwrLauCEHs8pg48UNqi8rLbsp-Qete4JsC+7LAdBwOTCoGwSAwtqodEGHoADpw0d8BAccBIUmxDu79CgDZPdEOLmEDoEIRhBCg5ZGHaAEEgg71Gk+A1Qilh0mzHNEOnx2Z+gGs5xvtsFyDRdgxrpdgMHaKV6w1e1+HjdR58be-B3ifgQPffrx7YufwPevD6SMe4R4CTyQNPVAs956viXqFFe3BqLuAHtIR2O8TZHx9mDEuqVA4XwHAUUgoQvBAhpFwXgSRw6R2bk-b4L9QHgMgfUAm+B6YvynjPJAlsUHOwAKz72-p7CS3tQYeQUAmWW9Rern2xCHChZp8B0XIgxBGzFMIvwTg-KhMdn7sGThvKBTD2gchfn-O2W8so7wAGxjjYRAjhdt0HCLjGIqRl864RxOqveoLJDHwBoewGxs8-S9DAAAIVgOQCB7dtjyiYHhcOLR6iZGyM4JAQoACSMINQrkgJKdwEM9TUDoEQCJCBWbpG6IgEY6APR4gABr1HONkNkbxuAEBOpcFJiYfB6mRpIJJVwl6F3egAeTCFidSCYUouDHGVOGiizKI0epZVyisxYDX6Z0oRJ8RGjKYOMjSwUpl5DHOI4K8NzJsCWRAFylk9brI6bKLZIyxkTMOTDaZoANZGUOucxZlkbmrLudIDZjyhmiV2fszSh0jn6VACbb5VVfnMSueAAFOtbb3OSaC4+zy9mvOhe845X9OZnPmTtZF-yVm4v2ZigZXocXgpeQcglBUiUDwRSZMlyiLJt37lBK6UEAXUvUqAIKh1HZYwdF1EqS1EoiCKqy1glUqqOwhu1KGnVCXsAVd1UqoAKqDWOtdBqsBZotQkmqjq0q9KgD6njaqQ1JWkRJhNKaM0RrzVZbKlaa0NrkC2lyi5SN7VHSykTc6okrqwHqrdCQ910baWek6nGIjPoIG+r9PyakUYJqeYypg4MxTqpaNa2Gn4lFBrRgm7SVb0aYw9SmuMIaDHhvejTaQ9NhW+Qkuiu2+cbJdvFgNMxidna7NdsSweP9BFgp2WEP2Ui3HsFOrAEdTteS7OzpO-OM6GVzoLVgnBhBZErrXTvXZe9t19unfSjB+6Exn2wbIJdHjuDIIzmOsIaDd13rjLs5xT6wAvtPVwjdYReFXoEbexxqF-3iNAJIwDx767LtXmez9TArGgACXYqDeb72HufSe1eQA).

Does anyone have any ideas for what the problem might be?
## [6][Is it possible to make one function argument optional depending on another argument's value?](https://www.reddit.com/r/typescript/comments/i8wona/is_it_possible_to_make_one_function_argument/)
- url: https://www.reddit.com/r/typescript/comments/i8wona/is_it_possible_to_make_one_function_argument/
---
Hi there! Just starting with typescript and I was guessing if this is possible to achieve. Any help or alternative solutions will be appreciated, thanks in advance!

So I have this method:

    triggerInteraction({
        parent,
        action,
        custom = false,
      }: {
        parent: HTMLElement;
        action: string;
        custom?: boolean;
      }) {
        // Search for trigger and click it if found
        const trigger = custom
          ? (document.querySelector(`[data-logic="${action}"]`) as HTMLElement)
          : (parent.querySelector(`[data-logic="${action}"]`) as HTMLElement);
    
        if (trigger) {
          trigger.click();
          return true;
        } else return false;
      }

Where the ***parent*** argument is only needed if the ***custom*** argument is true. Otherwise, it won't be used by the method.

So my goal is to be able to make this a valid call:

`triggerInteraction({ action, custom: true });`

But not this (***parent*** should be required here, because ***custom*** is false and therefore will be needed in the method):

`triggerInteraction({ action, custom: false});`

Is there any way to set this conditionally optional parameter? Or is there a better approach for this kind of solution?

&amp;#x200B;

Thanks!
## [7][Teki - An unreasonably efficient route parser in TypeScript](https://www.reddit.com/r/typescript/comments/i8doxo/teki_an_unreasonably_efficient_route_parser_in/)
- url: https://github.com/philipnilsson/teki#teki
---

## [8][Typescript node starter, with testing, dev mode, environment variable loading and github actions](https://www.reddit.com/r/typescript/comments/i84px0/typescript_node_starter_with_testing_dev_mode/)
- url: https://www.reddit.com/r/typescript/comments/i84px0/typescript_node_starter_with_testing_dev_mode/
---
[https://github.com/IhsanMujdeci/node-ts](https://github.com/IhsanMujdeci/node-ts)

&amp;#x200B;

Get started with \`npm start\`

&amp;#x200B;

* Dev mode transpiling ts with nodemon like re-running.
* Testing built in with jest.
* Github actions on push and pr to master that build and test the application
* Environment variable loading from file built in.
* Aliases your project under \`@/myApp\`, you can change it in the tsconfig.json
## [9][Type error on call but not on definition](https://www.reddit.com/r/typescript/comments/i7si1q/type_error_on_call_but_not_on_definition/)
- url: https://www.reddit.com/r/typescript/comments/i7si1q/type_error_on_call_but_not_on_definition/
---
I have an interface.

    export interface ViolationDetails {
      // …snip…
      getMessage: (extra: Record&lt;string, any&gt;) =&gt; string;
    }

Then I implement it.

    const myViolation: ViolationDetails = {
      // …snip…
      getMessage: (): string =&gt; 'Missing field';
    }

No errors to be seen. But when I try to call it, I get a compiler error.

    myViolation.getMessage();
    // Expected 1 argument, but got 0.

What gives? Why is there no type error on the implementation, but only on the call?
## [10][How to make cache layer redis with mongoose ?](https://www.reddit.com/r/typescript/comments/i7qxaj/how_to_make_cache_layer_redis_with_mongoose/)
- url: https://www.reddit.com/r/typescript/comments/i7qxaj/how_to_make_cache_layer_redis_with_mongoose/
---
i have cache completed with nodejs but with typescript this not overwrite. somebody suggest me 

cache.ts


        import * as mongoose from 'mongoose';
        import * as redis from 'redis';
        import * as util from 'util';
        import { RedisClient } from 'redis';

      let redisClient: RedisClient;

      export function startRedisConnection() {
        redisClient = redis.createClient({
          host: 'redis',
          port: 6379,
          retry_strategy: () =&gt; 1000
        });

        redisClient.hgetAsync = util.promisify(redisClient.hget);
        redisClient.hsetAsync = util.promisify(redisClient.hset);

        // Monkey patch mongoose exec method
        const exec = mongoose.Query.prototype.exec;

        mongoose.Query.prototype.cache = function(options: any = {}) {
          this.useCache = true;
          this.time = options.time || 60;
          // Dynamic top-level hashKey
          this.hashKey = JSON.stringify(options.hashKey || '');
          console.log(this.hashKey);
          return this;
        };

        mongoose.Query.prototype.exec = async function() {
          // No cache
          if (!this.useCache) {
            return exec.apply(this, arguments);
          }

          // {...query, collection: name }
          const key = JSON.stringify({ ...this.getQuery(), collection: this.mongooseCollection.name });
          // Get cache
          const cacheValue = await redisClient.hgetAsync(this.hashKey, key);
          if (cacheValue) {
            const value = JSON.parse(cacheValue);
            // Mongoose exec expects us to return mongoose documents
            console.log('response from redis');
            return Array.isArray(value) ? value.map(item =&gt; new this.model(item)) : new this.model(value);
          }

          const result = await exec.apply(this, arguments);
          // Set cache
          // key: ...hashkey
          // value:  { nestedKey: ...key, nestedValue: ...result }
          if (result) {
            redisClient.hsetAsync(this.hashKey, key, JSON.stringify(result));
          }
          console.log('Helo');
          return result;
        };
      }

      export function stopRedisConnection() {
        redisClient.quit();
      }

      export function removeHashKeyFromCache(hashKey: any) {
        redisClient.del(JSON.stringify(hashKey));
      }
## [11][Type error - Missing 200+ properties](https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/)
- url: https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/
---
    Type error: Type '{ children: Element; className: string; closeButton: true; }' is missing the following properties from type 'Pick&lt;ModalHeaderProps, "style" | "title" | "slot" | "children" | "bsPrefix" | "className" | "color" | "id" | "lang" | "role" | "tabIndex" | "dangerouslySetInnerHTML" |

     ... 246 more ... | "closeButton"&gt;': translate, onAuxClick, onAuxClickCapture  TS2739

https://imgur.com/a/XpLvrpo

Do I need to actually pass in 246+ properties to this Modal.Header component (using bootstrap)? I am pretty bad so please treat me like a huge noob! What should I do here?
