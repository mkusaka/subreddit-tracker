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
## [2][Teki - An unreasonably efficient route parser in TypeScript](https://www.reddit.com/r/typescript/comments/i8doxo/teki_an_unreasonably_efficient_route_parser_in/)
- url: https://github.com/philipnilsson/teki#teki
---

## [3][Typescript node starter, with testing, dev mode, environment variable loading and github actions](https://www.reddit.com/r/typescript/comments/i84px0/typescript_node_starter_with_testing_dev_mode/)
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
## [4][Type error on call but not on definition](https://www.reddit.com/r/typescript/comments/i7si1q/type_error_on_call_but_not_on_definition/)
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
## [5][How to make cache layer redis with mongoose ?](https://www.reddit.com/r/typescript/comments/i7qxaj/how_to_make_cache_layer_redis_with_mongoose/)
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
## [6][Type error - Missing 200+ properties](https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/)
- url: https://www.reddit.com/r/typescript/comments/i7evqe/type_error_missing_200_properties/
---
    Type error: Type '{ children: Element; className: string; closeButton: true; }' is missing the following properties from type 'Pick&lt;ModalHeaderProps, "style" | "title" | "slot" | "children" | "bsPrefix" | "className" | "color" | "id" | "lang" | "role" | "tabIndex" | "dangerouslySetInnerHTML" |

     ... 246 more ... | "closeButton"&gt;': translate, onAuxClick, onAuxClickCapture  TS2739

https://imgur.com/a/XpLvrpo

Do I need to actually pass in 246+ properties to this Modal.Header component (using bootstrap)? I am pretty bad so please treat me like a huge noob! What should I do here?
## [7][Add Typescript to your Next.js project](https://www.reddit.com/r/typescript/comments/i7m2kv/add_typescript_to_your_nextjs_project/)
- url: https://blog.bhanuteja.dev/add-typescript-to-your-nextjs-project-ckdpgej1h0174kbs1h9fg86yq
---

## [8][The best Rollup config for TypeScript libraries](https://www.reddit.com/r/typescript/comments/i7635n/the_best_rollup_config_for_typescript_libraries/)
- url: https://www.reddit.com/r/typescript/comments/i7635n/the_best_rollup_config_for_typescript_libraries/
---
https://gist.github.com/aleclarson/9900ed2a9a3119d865286b218e14d226

### Features
🔥 Blazing fast builds  
😇 CommonJS bundle  
🌲 .mjs bundle  
✨ .d.ts bundle + type-checking  
🧐 Source maps  

**Please retweet this if it helps you:**  
https://twitter.com/alecdotbiz/status/1291742610374131712
## [9][Why didn't my typings match this overload](https://www.reddit.com/r/typescript/comments/i7e6pl/why_didnt_my_typings_match_this_overload/)
- url: https://www.reddit.com/r/typescript/comments/i7e6pl/why_didnt_my_typings_match_this_overload/
---
This is purely practice with working through the complex method overloads I see. In the case of passport.serializeUser

    (method) passport.Authenticator&lt;e.Handler, any, any,
     passport.AuthenticateOptions&gt;.serializeUser&lt;unknown, unknown&gt;(
    fn: (user: unknown, done: (err: any, id?: unknown) =&gt; void) =&gt; void)
    : void (+1 overload)

I thought this would work:

    passport.serializeUser(&lt;UserShape extends { id: string }&gt;
    (user: UserShape, callback: any) =&gt; {
      callback(null, user.id as string);
    });

The error lint:

    No overload matches this call.
      Overload 1 of 2, '(fn: (user: unknown, done: (err: any, 
    id?: unknown) =&gt; void) =&gt; void): void', gave the following error.
        Argument of type '&lt;UserShape extends { id: string; }&gt;
    (user: UserShape, callback: any) =&gt; void' is not assignable to 
    parameter of type '(user: unknown, done: (err: any, 
    id?: unknown) =&gt; void) =&gt; void'.
          Types of parameters 'user' and 'user' are incompatible.
    
            Type 'unknown' is not assignable to type '{ id: string; }'.
      Overload 2 of 2, '(fn: (req: Request&lt;ParamsDictionary, any, any, 
    ParsedQs&gt;, user: unknown, done: (err: any, id?: unknown) =&gt; 
    void) =&gt; void): void', gave the following error.
    
    Argument of type '&lt;UserShape extends { id: string; }&gt;(user: UserShape, 
    callback: any) =&gt; void' is not assignable to parameter of type '(req: Request&lt;ParamsDictionary, any, any, ParsedQs&gt;, user: unknown, 
    done: (err: any, id?: unknown) =&gt; void) =&gt; void'.
          Types of parameters 'user' and 'req' are incompatible.
            Property 'id' is missing in type 'Request&lt;ParamsDictionary, 
    any, any, ParsedQs&gt;' but required in type '{ id: string; }'.ts(2769)

Seems like the first overload gets close. What was my mistake?
## [10][Type definition for function that could return undefined](https://www.reddit.com/r/typescript/comments/i7d62e/type_definition_for_function_that_could_return/)
- url: https://www.reddit.com/r/typescript/comments/i7d62e/type_definition_for_function_that_could_return/
---
I need help writing a type definition for a function that could return `undefined` or anything else, but the return type has to include `undefined`. I'm wanting to avoid a generic parameter in the type as well.

```
// this doesn't work, but is kind of what I'm going for
type CouldReturnUndefined = &lt;T&gt;(...args: any[]) =&gt; T | undefined

// this should be allowed
const a: CouldReturnUndefined = () =&gt; { return undefined }

// this should cause an error
const b: CouldReturnUndefined = () =&gt; { return 10 }

// this should be allowed
const c: CouldReturnUndefined = (): number | undefined =&gt; { return 10 }
```

Any ideas?
## [11][I built YJs bindings for slate - need feedback on my code](https://www.reddit.com/r/typescript/comments/i77bqu/i_built_yjs_bindings_for_slate_need_feedback_on/)
- url: https://github.com/BitPhinix/slate-yjs
---

