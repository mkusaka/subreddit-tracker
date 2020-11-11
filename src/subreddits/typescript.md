# typescript
## [1][Who's hiring Typescript developers - November](https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/)
- url: https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/
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
## [2][Does anyone know what target to use for Node 14?](https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/)
- url: https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/
---
I found [this page](https://github.com/microsoft/TypeScript/wiki/Node-Target-Mapping) which lists node target mappings for TS, but it only goes up to Node 12. Since Node 14 is the current LTS, I was wondering if it supported `ES2020` or `ESNEXT` targets. 

Really I'm asking because I want top-level-await support, but Node requires ES modules to do that, so I'd like to also know what `module` setting to use; `es2015`, `es2020`, or `esnext`?

Thanks!
## [3][How do I make this code behave synced?](https://www.reddit.com/r/typescript/comments/js4hqi/how_do_i_make_this_code_behave_synced/)
- url: https://www.reddit.com/r/typescript/comments/js4hqi/how_do_i_make_this_code_behave_synced/
---
I've found some code to extract an archive file but it's async in all kinds of ways. I need to make it run synced. Its running inside VSCode as extension code and I need to do more code when the code is done. I've tried to apply awaits but no matter what I do my calling code will never wait for the extract function to finish.

Code  references a Logger class that just writes to the console and Utils that will recursively will create target folder if not found.  


    import * as path from 'path';
    import * as fs from 'fs';
    import { Logger } from './Logger';
    import { Utils } from './Utils';
    
    export class Zip
    {
        public static async extract(zipFile:string, targetPath:string): Promise&lt;boolean&gt;
        {
            return new Promise((resolve, reject) =&gt;
            {
                try
                {
                    const JSZip = require('jszip');
    
                    fs.readFile(zipFile, async function(err, data) 
                    {
                        if (!err)
                        {
                            var zip = new JSZip();
                            await zip.loadAsync(data).then(function(contents: any)
                            {
                                Object.keys(contents.files).forEach(async function(filename)
                                {
                                    await zip.file(filename).async('nodebuffer').then(function(content: any)
                                    {
                                        let dest = path.join(targetPath, filename);
                                        let folderPath = path.dirname(dest);
        
                                        if (!fs.existsSync(folderPath))
                                        {
                                            Utils.makeDirectorySync(folderPath);
                                        }
    
                                        Logger.console(dest);
                                        fs.writeFileSync(dest, content);
                                    });
                                });
    
                                resolve(true);
                            });
                        }
                    });
                }
                catch (err)
                {
                    Logger.console('Error: ' + err);
                    reject(false);
                }
            });
        }
    }
## [4][Help breaking down typescript code](https://www.reddit.com/r/typescript/comments/js5uy6/help_breaking_down_typescript_code/)
- url: https://www.reddit.com/r/typescript/comments/js5uy6/help_breaking_down_typescript_code/
---
Hey all,

Was looking at using io-ts ([https://blog.jiayihu.net/how-to-validate-express-requests-using-io-ts/](https://blog.jiayihu.net/how-to-validate-express-requests-using-io-ts/)) as middleware validation for my expressjs app (using typeORM). I'm new to typescript and was wondering if someone could give me a hand dissecting below. 

What I understand so far:

\- We are exporting a constant variable called `validator` which is called with with type &lt;T&gt; which takes a `Decoder` as a variable. From this point on I am not sure, there a so many `=&gt;` operators that is looses me?

    import { RequestHandler } from 'express';
    import { Decoder } from 'io-ts/lib/Decoder';
    import { pipe } from 'fp-ts/lib/pipeable';
    import { fold } from 'fp-ts/lib/Either';
    
    export const validator: &lt;T&gt;(decoder: Decoder&lt;T&gt;) =&gt; RequestHandler&lt;ParamsDictionary, any, T&gt; = decoder =&gt; (req, res, next) =&gt; {
        return pipe(
            decoder.decode(req.body),
            fold(
                errors =&gt; res.status(400).send({ status: 'error', error: errors }),
                () =&gt; next(),
            ),
        );
    };
    

Much appreciated !
## [5][10 Insights from Adopting TypeScript at Scale](https://www.reddit.com/r/typescript/comments/jrgi8z/10_insights_from_adopting_typescript_at_scale/)
- url: https://www.techatbloomberg.com/blog/10-insights-adopting-typescript-at-scale/
---

## [6][Convert to Typescript Syntax](https://www.reddit.com/r/typescript/comments/jrz3dq/convert_to_typescript_syntax/)
- url: https://www.reddit.com/r/typescript/comments/jrz3dq/convert_to_typescript_syntax/
---
I've done async script load calls from typescript files. Some were formerly embedded inline on the script tag on index.html.

And then I came across this script tag which I believe initializes the fetched *some.resource.js.* Below is a simplified format:  


    &lt;script src="some.resource.js"&gt;&lt;/script&gt;
    &lt;script type='text/javascript'&gt;
        (function() {
            var voIP = VOIP({
                id: '189sfegz4',
                onSubmit: ({ step, attempt, result }) =&gt; { 
                    console.log('submit', { step, attempt, result });
                },
                init: {
                    title: 'hello world',
                    instruction: 'say hello'
                }
            });
         
           voIP.put("#voip-element");
         })();
    &lt;/script&gt;

This loads as intended, and accessed as a global window variable. Using a single-page application however,

  
how can the encapsulated function object be converted and initialized from a specific component (if that's possible)?

&amp;#x200B;

using Angular btw.
## [7][Could you peer review my PR? Specifically the types defined in the .d.ts file.](https://www.reddit.com/r/typescript/comments/jrytzc/could_you_peer_review_my_pr_specifically_the/)
- url: https://www.reddit.com/r/typescript/comments/jrytzc/could_you_peer_review_my_pr_specifically_the/
---
Hey smart people. Just wondering if any of you are free and willing to peer review a PR I've made for the deepmerge library? The library is written in JS; my PR is about improving its type definitions that are defined in a .d.ts file.

Here's the link: https://github.com/TehShrike/deepmerge/pull/211

Note: TypeScript version 4.1.1-rc required.
## [8][Can someone give me an example of constructor function in Typescript. I want to use `new` operator on function but everyone on the internet is using classes.](https://www.reddit.com/r/typescript/comments/jrl8wi/can_someone_give_me_an_example_of_constructor/)
- url: https://www.reddit.com/r/typescript/comments/jrl8wi/can_someone_give_me_an_example_of_constructor/
---
For example this doesn't work in Typescript. I want someone to fix this code. 

```ts
function Square(width: number){
  this.width = width;
}

var square = new Square(2);
```
## [9][Terraform with TypeScript](https://www.reddit.com/r/typescript/comments/jqxmoq/terraform_with_typescript/)
- url: https://medium.com/francisvitullo/terraform-with-typescript-7643defb4eb1?source=friends_link&amp;sk=975dc30cd7d48d989f2d7f0c16882c2e
---

## [10][Help pass - Quiz: "Parse nullable string"](https://www.reddit.com/r/typescript/comments/jrdaok/help_pass_quiz_parse_nullable_string/)
- url: https://www.reddit.com/r/typescript/comments/jrdaok/help_pass_quiz_parse_nullable_string/
---
Hi All, 

I am trying to learn coding in my spare time and have run into a bit of trouble.  Can anyone help?

P.S. This is not homework.  This is from a website called execute program.  

\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_

This is the prompt:

*Write a function that turns a* *string | undefined* *into a* *number.*

* *If the argument was* *undefined*  
*, the function should return* *undefined.*
* *If the argument was a string containing a valid number, it should return that number.*
* *If the argument was a string without a number, it should return undefined.*

*(You'll want to use the built-in functions* *parseInt(s: string)*  
 *and* *isNaN(n: number)**.)*

&amp;#x200B;

These are the tests you have to pass.  Currently I am failing test #1: 

&amp;#x200B;

 \&gt;maybeParseString(undefined) Expected: 

    undefined

but got: 

    type error: type error: Argument of type 'undefined' is not assignable to parameter of type 'string'.

\&gt;maybeParseString("3") Expected: 

    3

OK!

\&gt;maybeParseString("2701") Expected: 

    2701

OK!

\&gt;maybeParseString("junk") Expected: 

    undefined

OK!

\&gt;maybeParseString("not-a-number") Expected: 

    undefined

OK!

\&gt;maybeParseString(null) Expected: 

    type error

OK!

\&gt;maybeParseString(3) Expected: 

    type error

OK!

7 tests, 1 failure 

\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_

This is my code so far (which is wrong):

&amp;#x200B;

**function maybeParseString(s: string ): number | undefined {**

  **const n = parseInt(s)**

  **if (isNaN(n)) {**

**return undefined**

  **}**

  **else {return n}**

**}**

&amp;#x200B;

I have tried 1000 and 1 different ways and can't seem to crack the code. 

&amp;#x200B;

Thanks!
## [11][TypeORM: EntityMetadataNotFound: No metadata for "User" was found.](https://www.reddit.com/r/typescript/comments/jr65m4/typeorm_entitymetadatanotfound_no_metadata_for/)
- url: https://www.reddit.com/r/typescript/comments/jr65m4/typeorm_entitymetadatanotfound_no_metadata_for/
---
Hey all! 

Hoping some typeORM experts can help me out, burned wayyy to much time on this. 

Was wondering if anyone can shed some light on this error, I know there are a lot of questions and answers relating to this that are mainly around not having the correct destination in the \`entities\` section of the connection or just not including the entity at all. In my case this only happens when I try run tests (Jest + Supertest). I've tried adding the path to my entities instead of them individually but also doesn't work Wondering if it is something to do with the way I am setting up the test connection and that the rest of my code can't see the new connection. Below is my \`createConnection\` function: 

    export const createTestConnection = async () =&gt; createConnection({
        dropSchema: true,
        entities: [User], // There will be more
        logging: ["error"],
        maxQueryExecutionTime: 500,
        name: "test",
        schema: "test",
        synchronize: true,
        type: "postgres",
        url: process.env.DB_URL,
    });

I checked the DB (removed \`dropSchema\`) and the entity is there but when I run \`find()\` method in my controller it throws the error:

    import { EntityRepository, Repository } from "typeorm";
    import { User } from "../entities/User";
    
    @EntityRepository(User)
    export class UserRepository extends Repository&lt;User&gt; {
    
        findAll(): Promise&lt;User[]&gt; {
            return this.find()
        }
    
        findUserByKey(key: string): Promise&lt;User | undefined&gt; {
            return this.findOne({ key });
        }
    }
    

Stacktrace:

        EntityMetadataNotFound: No metadata for "User" was found.
    
          at new EntityMetadataNotFoundError (src/error/EntityMetadataNotFoundError.ts:9:9)
          at Connection.Object.&lt;anonymous&gt;.Connection.getMetadata (src/connection/Connection.ts:333:19)
          at EntityManager.Object.&lt;anonymous&gt;.EntityManager.getCustomRepository (src/entity-manager/EntityManager.ts:1255:86)
          at Connection.Object.&lt;anonymous&gt;.Connection.getCustomRepository (src/connection/Connection.ts:368:29)
          at Object.getCustomRepository (src/index.ts:300:55)
          at getAllUsers (controllers/User.ts:7:46)
          at Layer.handle [as handle_request] (node_modules/express/lib/router/layer.js:95:5)
          at next (node_modules/express/lib/router/route.js:137:13)
          at Route.dispatch (node_modules/express/lib/router/route.js:112:3)
          at Layer.handle [as handle_request] (node_modules/express/lib/router/layer.js:95:5)
