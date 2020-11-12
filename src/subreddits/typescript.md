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
## [2][Weird behaviour on VS-Code](https://www.reddit.com/r/typescript/comments/jsseoh/weird_behaviour_on_vscode/)
- url: https://www.reddit.com/r/typescript/comments/jsseoh/weird_behaviour_on_vscode/
---
Hello everyone!

Since beginning of the week something doesnt seem right.  
I dont know if its the fault of TypeScript or VS-Code (or something else) but TS seems to be the problem here.

Since the beginning of this week small "beauty" problems are marked as warnings in VS-Code such as indentations or double quotes instead of single quotes.

The warning message always points to tslint and googling the warning messages also seem to point out tslint - but nobody from the project changed the tslint file since the beginning of the project 2 months ago.

some of these messages look like this :

 `space indentation expected (indent)tslint(1)`   
or  
 `Shadowed name: 'ClassName' (no-shadowed-variable)tslint(1)` 

these werent there last week and makes the code very ugly and chaotic to look at.

I would share some code but its literally on any TS file when I press the tab key or use double quotes or instanciate a Class and on and on

EDIT :
I asked the other one working on that project who is using VSCode and he says he doesn't have the same issue. That's makes it even weirder
## [3][I have converted all my js cloud functions to typescript, what are the most basic improvements I can make to my code?](https://www.reddit.com/r/typescript/comments/jsf66p/i_have_converted_all_my_js_cloud_functions_to/)
- url: https://www.reddit.com/r/typescript/comments/jsf66p/i_have_converted_all_my_js_cloud_functions_to/
---
Everything works in Javascript, I just converted all the files from .js to .ts.  What are the most basic improvements I can make to the code to already get alot of benefits from Typescript? I have started with adding a type to all my function arguments like `req: Request, res: Response`

I was wondering if there were any other basic tips you could give me! I'm also reading online about it but reddit is always a great spot to learn some stuff aswell :) 

Cheers, thanks!!
## [4][How to create overlapping interface?](https://www.reddit.com/r/typescript/comments/jsch06/how_to_create_overlapping_interface/)
- url: https://www.reddit.com/r/typescript/comments/jsch06/how_to_create_overlapping_interface/
---
I was able to create intersection types using the following

`type IntersectingTypes&lt;T, U&gt; = { [K in Extract&lt;keyof T, keyof U&gt;]: T[K] }` 

However it doesn't work when I try to do the following

`interface IntersectingTypes&lt;T, U&gt; { [K in Extract&lt;keyof T, keyof U&gt;]: T[K] }` 

Is there a way I could create intersecting interface?
## [5][Does anyone know what target to use for Node 14?](https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/)
- url: https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/
---
I found [this page](https://github.com/microsoft/TypeScript/wiki/Node-Target-Mapping) which lists node target mappings for TS, but it only goes up to Node 12. Since Node 14 is the current LTS, I was wondering if it supported `ES2020` or `ESNEXT` targets. 

Really I'm asking because I want top-level-await support, but Node requires ES modules to do that, so I'd like to also know what `module` setting to use; `es2015`, `es2020`, or `esnext`?

Thanks!
## [6][Help breaking down typescript code](https://www.reddit.com/r/typescript/comments/js5uy6/help_breaking_down_typescript_code/)
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
## [7][How do I make this code behave synced?](https://www.reddit.com/r/typescript/comments/js4hqi/how_do_i_make_this_code_behave_synced/)
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
## [8][10 Insights from Adopting TypeScript at Scale](https://www.reddit.com/r/typescript/comments/jrgi8z/10_insights_from_adopting_typescript_at_scale/)
- url: https://www.techatbloomberg.com/blog/10-insights-adopting-typescript-at-scale/
---

## [9][Convert to Typescript Syntax](https://www.reddit.com/r/typescript/comments/jrz3dq/convert_to_typescript_syntax/)
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
## [10][Could you peer review my PR? Specifically the types defined in the .d.ts file.](https://www.reddit.com/r/typescript/comments/jrytzc/could_you_peer_review_my_pr_specifically_the/)
- url: https://www.reddit.com/r/typescript/comments/jrytzc/could_you_peer_review_my_pr_specifically_the/
---
Hey smart people. Just wondering if any of you are free and willing to peer review a PR I've made for the deepmerge library? The library is written in JS; my PR is about improving its type definitions that are defined in a .d.ts file.

Here's the link: https://github.com/TehShrike/deepmerge/pull/211

Note: TypeScript version 4.1.1-rc required.
## [11][Can someone give me an example of constructor function in Typescript. I want to use `new` operator on function but everyone on the internet is using classes.](https://www.reddit.com/r/typescript/comments/jrl8wi/can_someone_give_me_an_example_of_constructor/)
- url: https://www.reddit.com/r/typescript/comments/jrl8wi/can_someone_give_me_an_example_of_constructor/
---
For example this doesn't work in Typescript. I want someone to fix this code. 

```ts
function Square(width: number){
  this.width = width;
}

var square = new Square(2);
```
