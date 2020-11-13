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
## [2][Having difficulty typing a function that can filter on a specific level inside of an object while preserving its original structure](https://www.reddit.com/r/typescript/comments/jt9n7j/having_difficulty_typing_a_function_that_can/)
- url: https://www.reddit.com/r/typescript/comments/jt9n7j/having_difficulty_typing_a_function_that_can/
---
I have an object called `myCached` and its type is known as 
```ts
interface MyCache {
  a: {
    b: {
      c: Target[];
      d: string;
    };
  };
}
```
Where `Target` is the type of the item in one of its property on a specific level, this type `Target` is known too. Assume `Target` is like this

```ts
type Target = "toRemove1" | "toRemove2" | "toPreserve1" | "toPreserve2";
```
In the example here, `myCache` is 
```ts
const cache: MyCache = {
  a: {
    b: { 
      c: ['toRemove1', 'toPreserve1', 'toRemove2', 'toPreserve2'], 
      d: "irrelevant" 
    },
  }
}
``` 
And I have an array of items that I need to remove. They also have the type of `Target` as in 
```ts
const thingsToRemove: Target[] = ["toRemove1", "toRemove2"];
```

And I am trying to come up with a function that can traverse this object and filter out items on a specific level. The way I designed this function is that it takes a transformer object and it travers the object and it provides the function for filtering. I want to type it properly so the users of this function can reply on the auto complete provided by the TS compiler and type checking to make sure that this function lands correctly on the correct level, in this case it is the `Target` level.

This is my code

```ts

type MappedTransform&lt;T&gt; = {
  [K in keyof T]?: MappedTransform&lt;T[K]&gt; | ((params: T[K]) =&gt; T[K]);
};

type Entries&lt;T&gt; = { [K in keyof T]-?: [K, T[K]] }[keyof T];

function traverse&lt;R&gt;(cache: R, transformObject: MappedTransform&lt;R&gt;): R {
  return (Object.entries(transformObject) as Array&lt;
    Entries&lt;MappedTransform&lt;R&gt;&gt;
  &gt;).reduce(reduceTransformNode, cache);
}

const reduceTransformNode = &lt;R, K extends keyof R&gt;(
  cacheNode: R,
  [transformKey, transformValue]: [
    K,
    MappedTransform&lt;R[K]&gt; | ((params: R[K]) =&gt; R[K]) | undefined
  ]
): R =&gt; {
  const { [transformKey]: node } = cacheNode;

  if (typeof transformValue === "undefined") return cacheNode;

  const newCacheValue =
    typeof transformValue === "function"
      ? (transformValue as (params: R[K]) =&gt; R[K])(node)
      : traverse(transformValue as R[K], node);

  return {
    ...cacheNode,
    [transformKey]: newCacheValue
  };
};

const x = traverse(cache, {
  a: {
    b: {
      // 👇Need to use Type `Target` to make sure that the transformer function lands on the type of `Target` exactly
      c: (node) =&gt; node.filter((s) =&gt; !thingsToRemove.includes(s))
    }
  }
});
```

This works fine and I have been really close to what I wanted to achieve except that **I cannot seem to find a way to add the type of `Target` to make sure that the function the user provides to do the filtering actually is on the right level i.e. `node` is of the right type `Target`**

Here is the [live demo][1]. Can someone tell me how to achieve that last bit of type safety I am seeking for here?


  [1]: https://codesandbox.io/s/superhardts-116bf?file=/src/index.ts
## [3][Dev Experience: Is it possible to create a type that accepts any string but also provides autocompletion for specific strings?](https://www.reddit.com/r/typescript/comments/jtg7kt/dev_experience_is_it_possible_to_create_a_type/)
- url: https://www.reddit.com/r/typescript/comments/jtg7kt/dev_experience_is_it_possible_to_create_a_type/
---
I have a data object that contains certain keys but it can also contain random keys, for the keys I know about I want to return their type, for not known keys I want to return `unknown`.

Like this:  


    interface IData {
        a: string[];
        b: boolean;
    }
    
    function lookup&lt;K extends string&gt;(key: K): K extends keyof IData ? IData[K] : unknown {
        //do lookup
    }
    
    let a = lookup("a"); //infers string[]
    let b = lookup("b"); //infers boolean
    let c = lookup("c"); //infers unknown

If the argument for the `lookup` function matches a key from `IData` it will set the return type accordingly, else it will return `unknown`.

This works, however I don't get any autocomplete in the IDE for the known keys in `IData`. But for improved DX I would like to while still being able to enter any string as an argument, is there a way to do this currently?
## [4][A required key appears after transpiling to JS](https://www.reddit.com/r/typescript/comments/jtg2nc/a_required_key_appears_after_transpiling_to_js/)
- url: https://www.reddit.com/r/typescript/comments/jtg2nc/a_required_key_appears_after_transpiling_to_js/
---
Hi all,

I've been breaking my head over a property (`css`) which is not on my TSX component but becomes required after transpiling it to JS

    import React from 'react'
    import { default as MuiTooltip, TooltipProps as MuiTooltipProps } from '@material-ui/core/Tooltip'

    export type TooltipProps = MuiTooltipProps
    export default React.forwardRef(function Tooltip(props: TooltipProps, ref): JSX.Element {
      return &lt;MuiTooltip ref={ref} {...props} /&gt;
    })


Here I recreated the issue with the TS and JS version of the same component:  https://codesandbox.io/s/naughty-lichterman-mtx9l?file=/src/App.tsx 

Anyone any idea where `css` comes from?
## [5][[vscode] Is it possible (if so, how) to add additional documentation for builtins, vendor, etc.?](https://www.reddit.com/r/typescript/comments/jszxvp/vscode_is_it_possible_if_so_how_to_add_additional/)
- url: https://www.reddit.com/r/typescript/comments/jszxvp/vscode_is_it_possible_if_so_how_to_add_additional/
---
I don't know what this would be called, I don't think any other IDE I've used supports this.

There are cases where I'd like to add additional documentation to builtin/externals -- or maybe "personalised"? Like if I'm stupid and continue misusing an API (like "hey I know it's common sense to be able to use \`.start()\` and \`.stop()\` on this interface, once stopped it can not be started again.").

I doubt this would be a "language" feature, but is there maybe a plugin for vscode?

I could see cases for having the documentation part of the repo and out of repo. Obviously another solution would be add an interface (and note exceptions locally). This was a "hmm I wonder if this is someone else has thought of" post.
## [6][Weird behaviour on VS-Code](https://www.reddit.com/r/typescript/comments/jsseoh/weird_behaviour_on_vscode/)
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
## [7][I have converted all my js cloud functions to typescript, what are the most basic improvements I can make to my code?](https://www.reddit.com/r/typescript/comments/jsf66p/i_have_converted_all_my_js_cloud_functions_to/)
- url: https://www.reddit.com/r/typescript/comments/jsf66p/i_have_converted_all_my_js_cloud_functions_to/
---
Everything works in Javascript, I just converted all the files from .js to .ts.  What are the most basic improvements I can make to the code to already get alot of benefits from Typescript? I have started with adding a type to all my function arguments like `req: Request, res: Response`

I was wondering if there were any other basic tips you could give me! I'm also reading online about it but reddit is always a great spot to learn some stuff aswell :) 

Cheers, thanks!!
## [8][How to create overlapping interface?](https://www.reddit.com/r/typescript/comments/jsch06/how_to_create_overlapping_interface/)
- url: https://www.reddit.com/r/typescript/comments/jsch06/how_to_create_overlapping_interface/
---
I was able to create intersection types using the following

`type IntersectingTypes&lt;T, U&gt; = { [K in Extract&lt;keyof T, keyof U&gt;]: T[K] }` 

However it doesn't work when I try to do the following

`interface IntersectingTypes&lt;T, U&gt; { [K in Extract&lt;keyof T, keyof U&gt;]: T[K] }` 

Is there a way I could create intersecting interface?
## [9][Does anyone know what target to use for Node 14?](https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/)
- url: https://www.reddit.com/r/typescript/comments/js2uh5/does_anyone_know_what_target_to_use_for_node_14/
---
I found [this page](https://github.com/microsoft/TypeScript/wiki/Node-Target-Mapping) which lists node target mappings for TS, but it only goes up to Node 12. Since Node 14 is the current LTS, I was wondering if it supported `ES2020` or `ESNEXT` targets. 

Really I'm asking because I want top-level-await support, but Node requires ES modules to do that, so I'd like to also know what `module` setting to use; `es2015`, `es2020`, or `esnext`?

Thanks!
## [10][Help breaking down typescript code](https://www.reddit.com/r/typescript/comments/js5uy6/help_breaking_down_typescript_code/)
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
## [11][How do I make this code behave synced?](https://www.reddit.com/r/typescript/comments/js4hqi/how_do_i_make_this_code_behave_synced/)
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
