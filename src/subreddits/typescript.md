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
## [2][T | undefined === T ???](https://www.reddit.com/r/typescript/comments/hrza4i/t_undefined_t/)
- url: https://www.reddit.com/r/typescript/comments/hrza4i/t_undefined_t/
---
I want to create an optional type, something like `type Optional&lt;T&gt; = T | undefined` but TypeScript is telling me that such a type is equivalent to `type Optional&lt;T&gt; = T` - for some reason the undefined is lost.

I assume the cause of this is something to do with the fact the T can be of type undefined; maybe?

I want to use this type to create a type guard called `isOptional` that can be used like this: `isOptional(isNumber)(foo)`. Where  `isOptional` would essentially just run `isUndefined(foo) || isNumber(foo)`.

So from a top level approach, this is my issue:

```
if (isOptional(isNumber)(foo)) {  
  // expected: `foo` should be of type `number | undefined`  
  // actual: `foo` is of type `number`  
}
```

Is this a bug with TypeScript or am I doing something wrong?
## [3][Trying to model a function declaration that returns a class instance](https://www.reddit.com/r/typescript/comments/hs8544/trying_to_model_a_function_declaration_that/)
- url: https://www.reddit.com/r/typescript/comments/hs8544/trying_to_model_a_function_declaration_that/
---
The Error and surrounding context:

    /*
    [6:38:31 PM] File change detected. Starting incremental compilation...
    
    src/classes/sentenceBuilder/AudioMaker.ts:370:6 - error TS2349: This expression is not callable.
      Type 'typeof import("audioconcat")' has no call signatures.
    
    370      audioConcat(audiosAndPauseFiles)
             ~~~~~~~~~~~
    
    [6:38:31 PM] Found 1 error. Watching for file changes.
    */
         audioConcat(audiosAndPauseFiles)
           .concat(finalFileSavePath)
           .on('start', () =&gt; {
             console.log(`ffmpeg build process started on file at: ${finalFileSavePath}`);
           })
           .on('end', () =&gt; {
             console.log(`Sucessfully created file at: ${finalFileSavePath}`);
           })
           .on('error', (error: any, stdout: any, stderr: any) =&gt; {
             console.log('error', error);
             console.log('stdout', stdout);
             console.log('stderr', stderr);
           });
       }

My declaration file

    declare module "audioconcat" {
        class AudioConcat {
         constructor(audios: Array&lt;string&gt;);
    
         VERSION: string;
    
         ffmpeg: Function;
    
         concat(images: string[], options: any): this;
    
         on(event: 'start', callback: (command: any) =&gt; void): this;
         on(event: 'error', callback: (err: any, stdout: any, stderr: any) =&gt; void):  this;
         on(event: 'end', callback: (output: any) =&gt; void): this; 
      }
    
      export function AudioConcat(inputs: string[], opts: any): AudioConcat
    }

What exactly is the issue here? Below is the index.js containing the full library code:

    var merge = require('lodash.merge')
    var ffmpeg = require('fluent-ffmpeg')
    var version = require('./package.json').version
    
    module.exports = exports = function (inputs, opts) {
      return new Audioconcat(inputs, opts)
    }
    
    exports.VERSION = version
    exports.ffmpeg = ffmpeg
    
    function Audioconcat(inputs, opts) {
      this.inputs = inputs || []
      this.opts = opts || {}
    }
    
    Audioconcat.prototype.options = function (opts) {
      merge(this.opts, opts)
      return this
    }
    
    Audioconcat.prototype.concat = function (file) {
      if (file) {
        this.opts.output = file
      }
      return concat(this.inputs, this.opts)
    }
    
    function concat(inputs, opts) {
      var filter = 'concat:' + inputs.join('|')
    
      var renderer = ffmpeg()
        .input(filter)
        .outputOptions('-acodec copy')
    
      var output = opts.output
      if (output) {
        return renderer.save(output)
      }
    
      return renderer
    }
## [4][Setting up React Typescript Monorepo with Lerna](https://www.reddit.com/r/typescript/comments/hrrur2/setting_up_react_typescript_monorepo_with_lerna/)
- url: https://www.reddit.com/r/typescript/comments/hrrur2/setting_up_react_typescript_monorepo_with_lerna/
---
While working on complicated software solutions, many times we need to create multiple projects which may have some common building blocks, or some other sharing code.

One good approach to create such solutions is to create monorepos.

I have written this blog post to create one such typescript react monorepo using Lerna.

[http://devwithabhi.com/setting-up-react-typescript-monorepo-with-lerna/](http://devwithabhi.com/setting-up-react-typescript-monorepo-with-lerna/)

Hope you guys like it.
## [5][Active Open source server project in GitHub](https://www.reddit.com/r/typescript/comments/hs37ce/active_open_source_server_project_in_github/)
- url: https://www.reddit.com/r/typescript/comments/hs37ce/active_open_source_server_project_in_github/
---
Hey Guys,

Can you please let me know good opensource node(may be with typescript) project that are active in GitHub.
## [6][[mutex-server] mutex and semaphore in the network level](https://www.reddit.com/r/typescript/comments/hrrcyy/mutexserver_mutex_and_semaphore_in_the_network/)
- url: https://www.reddit.com/r/typescript/comments/hrrcyy/mutexserver_mutex_and_semaphore_in_the_network/
---
Critical sections in the network level.

* [https://github.com/samchon/mutex-server](https://github.com/samchon/mutex-server)
* [https://mutex.dev/api](https://mutex.dev/api)

The `mutex-server` is an npm module that can be used for building a mutex server. When you need to control a critical section on the entire system level, like distributed processing system using the network communications, this `mutex-server` can be a good solution.

Installs and opens a `mutex-server` and let clients to connect to the server. The clients who're connecting to the `mutex-server` can utilize remote critical section components like [mutex](https://mutex.dev/api/classes/msv.remotemutex.html) or [semaphore](https://mutex.dev/api/classes/msv.remotesemaphore.html).

Also, `mutex-server` has a safety device for network disconnections. When a client has been suddenly disconnected, all of the locks had acquired or tried by the client would be automatically unlocked or cancelled. Therefore, you don't worry about any network disconnection accident and just enjoy the `mutex-server` with confidence.
## [7][A library to generate the maximally efficient series of unique keys for a given alphabet.](https://www.reddit.com/r/typescript/comments/hrm45x/a_library_to_generate_the_maximally_efficient/)
- url: https://github.com/gactjs/key
---

## [8][Overload Method declaration error(s)](https://www.reddit.com/r/typescript/comments/hrqqpl/overload_method_declaration_errors/)
- url: https://www.reddit.com/r/typescript/comments/hrqqpl/overload_method_declaration_errors/
---
    declare class audioconcat {
      constructor(audios: Array&lt;string&gt;);
    
      VERSION: string;
    
      ffmpeg: Function;
    
      concat(images: string[], options: any): this;
    
      on(hook: string, callback(command: any): void): this;
      on(hook: string, callback(err: any, stdout: any, stderr: any): void): this;
      on(hook: string, callback(output: any): void): this;
    /*
    (parameter) callback: any
    'callback' is declared but its value is never read.ts(6133)
    Parsing error: ',' expected.eslint
    Parameter 'callback' implicitly has an 'any' type.ts(7006)
    */
    }

I'm not sure if I erred due to special syntax for a callback, or a method overload.

Also, I have this in `project/src/declarations.d.ts.` TS seems unable to find this file. Why isn't it being picked up? I thought it should be found recursively.

    import audioConcat from 'audioconcat';
    
    /*
    Could not find a declaration file for module 'audioconcat'.
     '/home/owner/cp/project/node_modules/audioconcat/index.js' 
    implicitly has an 'any' type.
    
    Try `npm install @types/audioconcat` if it exists or add 
    a new declaration (.d.ts) file containing `declare module 
    'audioconcat';`ts(7016)
    */

Here is the library: [https://www.npmjs.com/package/audioconcat](https://www.npmjs.com/package/audioconcat)

    audioconcat(songs)
      .concat('all.mp3')
      .on('start', function (command) {
        console.log('ffmpeg process started:', command)
      })
      .on('error', function (err, stdout, stderr) {
        console.error('Error:', err)
        console.error('ffmpeg stderr:', stderr)
      })
      .on('end', function (output) {
        console.error('Audio created in:', output)
      })
## [9][Marking Optional Parameters in Callbacks As NonOptional?](https://www.reddit.com/r/typescript/comments/hrmj9s/marking_optional_parameters_in_callbacks_as/)
- url: https://www.reddit.com/r/typescript/comments/hrmj9s/marking_optional_parameters_in_callbacks_as/
---
From the official docs:

## Optional Parameters in Callbacks [\#](https://www.typescriptlang.org/docs/handbook/declaration-files/do-s-and-don-ts.html#optional-parameters-in-callbacks)

*Don’t* use optional parameters in callbacks unless you really mean it:

    /* WRONG */ 
    interface Fetcher {    
      getObject(done: (data: any, elapsedTime?: number) =&gt; void): void; 
    } 

This has a very specific meaning: the `done` callback might be invoked with 1 argument or might be invoked with 2 arguments. The author probably intended to say that the callback might not care about the elapsedTime parameter, but **there’s no need to make the parameter optional to accomplish this – it’s always legal to provide a callback that accepts fewer arguments.**

*Do* write callback parameters as non-optional:

    /* OK */ 
    interface Fetcher {    
      getObject(done: (data: any, elapsedTime: number) =&gt; void): void; 
    }

\----------

Unless callbacks are unusual, I seem to remember a lot of lint errors thrown by TSC along the lines of `Expected 2 arguments, but received only 1`. I have a hard time squaring this with the above, can anyone clarify this?
## [10][Why does my typescript file complain about not finding the declared variable and assign its number?](https://www.reddit.com/r/typescript/comments/hrkrr3/why_does_my_typescript_file_complain_about_not/)
- url: https://www.reddit.com/r/typescript/comments/hrkrr3/why_does_my_typescript_file_complain_about_not/
---
 I'm getting an error say that it cannot find age, when trying to assign to 18, 

Does someone know why?

    declare var age: number  
    
    age = 18 // error


When I try to build binder with webpack it say variable cannot be found.
## [11][Cool helper type that will allow you to recursively replace all types in an object with another type.](https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/)
- url: https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/
---
Figured this out recently, when I had to transform large plain objects of different shapes, and replace some leaf value with a different one. `Replaced` type takes 4 generic arguments:

* `T` the object to be transformed.
* `TReplace` the type to be replaced with
* `TWith` the type to use instead of `TReplace`
* `TKeep` the type to not transform. Defaults to primitive values.  


Obviously to be used as the return type of some function that does the actual transformation.

    type Primitive = string | number | bigint | boolean | null | undefined;
    
    type Replaced&lt;T, TReplace, TWith, TKeep = Primitive&gt; = T extends TReplace | TKeep
        ? (T extends TReplace
            ? TWith | Exclude&lt;T, TReplace&gt;
            : T)
        : {
            [P in keyof T]: Replaced&lt;T[P], TReplace, TWith, TKeep&gt;
        }
    
    type Foo = symbol;
    type Bar = "3";
    
    type ToReplace = {
        x?: {
            y: Foo | number;
            z: string;
        }[]
    } | undefined;
    
    type WasReplaced = Replaced&lt;ToReplace, Foo, Bar&gt;;
    /* Output:
    {
        x?: {
            y: number | "3";
            z: string;
        }[];
    }
    */
