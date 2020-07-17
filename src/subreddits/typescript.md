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
## [2][Converting the few left at my work to TS](https://www.reddit.com/r/typescript/comments/hsr3ri/converting_the_few_left_at_my_work_to_ts/)
- url: https://www.reddit.com/r/typescript/comments/hsr3ri/converting_the_few_left_at_my_work_to_ts/
---
Every other week, folks in my group are able to host a 40 minute technical talk. I’m looking at hosting one in an attempt to convert some of the late adopters in the group to TS. Some background, we primarily work in a rapid prototyping space so the old “Typescript is slower to develop with” argument is what has kept the last few from adopting. I’m trying to attack that head on while also talking about other benefits to push my argument over the edge. I’m thinking about half of the talk will be demonstration and questions so I want to keep to the most obvious wins.

The rough outline of my presentation is this:

- Intellisense: by identifying and defining your types up front you earn intellisense which speeds up development 
- Runtimes errors/refactoring: Property changes and differences are caught at compile time rather than runtime.
- Coordination/documentation: TS improves the ability to communicate through though code. Often we return to prototypes a year after the initial work was done and have to spin up and develop on a code rot ridden app.

Any advice is very welcome. Specifically, does anyone have any additional points they think I should make or angles I should address?
## [3][Use mongoose FilterQuery&lt;T&gt; + check if the keys exists](https://www.reddit.com/r/typescript/comments/hsrstk/use_mongoose_filterqueryt_check_if_the_keys_exists/)
- url: https://www.reddit.com/r/typescript/comments/hsrstk/use_mongoose_filterqueryt_check_if_the_keys_exists/
---
Right now in our project (using NestJS) we do have some methods in our service like this:

```
import { Model, CreateQuery, FilterQuery } from 'mongoose';
import { User } from './user.schema';

class UsersService {
  constructor(
    @InjectModel(User.name) private readonly usersModel: Model&lt;User&gt;,
  ) {}

  async findOne(conditions: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.findOne(conditions as FilterQuery&lt;User&gt;);
  }

  async create(fields: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.create(fields as CreateQuery&lt;User&gt;);
  }

  async updateOne(id, fields: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.updateById(id, fields as UpdateQuery&lt;User&gt;);
  }
}
```

In the above code it makes sure through TypeScript that when I do:

```
this.usersService.create({
  name: 'John',
  countryy: 'USA', // &lt;-- error because it should be 'country'
})
```

This works perfectly, but advanced query syntax is not supported through this method. For example:

```
this.usersService.findOne({
  'myObject.nestedKey.nestedKey2': 'test',
})
```

due to this I decided to just change the `conditions` parameter in UsersService.findOne() function to be `Record&lt;string, unknown&gt;`.

Is there a possibility to allow this types of advanced syntax and still check whether those keys are valid? This means that TypeScript should know, when doing `myObject.nestedKey.nestedKey2`, that `myObject` is a valid key, `nestedKey` is also a valid key inside `myObject` etc ...

Mongoose is using `FilterQuery&lt;T&gt;` but it doesn't check for valid keys I believe

Would love to know some solutions.
## [4][Foal TS - Node.JS and TypeScript framework - July Release - Management of several environments &amp; Simplified authentication](https://www.reddit.com/r/typescript/comments/hsup8w/foal_ts_nodejs_and_typescript_framework_july/)
- url: https://www.reddit.com/r/typescript/comments/hsup8w/foal_ts_nodejs_and_typescript_framework_july/
---
Foal TS version 1.11 is here!

This version facilitates the management of several environments thanks to its abstract services and it also reduces the code to produce to build an authentication.

The documentation of the new features can be found here:

\- [https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#abstract-services](https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#abstract-services)

\- [https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/session-tokens#specify-the-name-of-the-session-store-in-the-configuration](https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/session-tokens#specify-the-name-of-the-session-store-in-the-configuration)

\- [https://foalts.gitbook.io/docs/topic-guides/validation-and-sanitization#usage-with-a-hook](https://foalts.gitbook.io/docs/topic-guides/validation-and-sanitization#usage-with-a-hook)

&amp;#x200B;

In a few words, Foal TS is a [**Node.Js**](http://node.js/) and **TypeScript** framework that **provides the bricks to build a complete web application** while keeping a **simple and intuitive code**.

Backed by **thousands of tests**, it offers more than **11,000 lines of documentation**. [**#typescript**](https://www.linkedin.com/feed/hashtag/?keywords=typescript&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376) [**#javascript**](https://www.linkedin.com/feed/hashtag/?keywords=javascript&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376) [**#nodejs**](https://www.linkedin.com/feed/hashtag/?keywords=nodejs&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376)

&amp;#x200B;

https://preview.redd.it/8fpseb3rteb51.png?width=1064&amp;format=png&amp;auto=webp&amp;s=659f65321229a00a7ae887a53cc246f54709bcaf
## [5][Terminal](https://www.reddit.com/r/typescript/comments/hsu03i/terminal/)
- url: https://v.redd.it/3zarniw4keb51
---

## [6][Typing a recursive camelize function?](https://www.reddit.com/r/typescript/comments/hstaxp/typing_a_recursive_camelize_function/)
- url: https://www.reddit.com/r/typescript/comments/hstaxp/typing_a_recursive_camelize_function/
---
It would be nice to not lose type of the object but IDK if this is far beyond what TypeScript is capable of:

    export function camelize&lt;T&gt;(source: T, reverse = false) {
      // console.log('snakeCaseObject', source);
      const dest: Record&lt;string, any&gt; = {};
      const fn = reverse ? snakeCase : camelCase;

      for (let [key, value] of Object.entries(source)) {
        if (isPlainObject(value)) {
          // checks that a value is a plain object or an array - for recursive key conversion
          value = camelize(value, reverse); // recursively update keys of any values that are also objects
        }
        if (isArray(value)) {
          value = value.map(v =&gt; (isPlainObject(v) ? camelize(v, reverse) : v));
        }
        dest[fn(key)] = value;
      }

      return dest;
    }


This is what I have so far, but surely the actual return type is T but with camelCased(or uncamelcased) keys.
## [7][Currying function parameters - Convert to Numbers](https://www.reddit.com/r/typescript/comments/hshhsn/currying_function_parameters_convert_to_numbers/)
- url: https://www.reddit.com/r/typescript/comments/hshhsn/currying_function_parameters_convert_to_numbers/
---
I have function like:

*export* const isGreaterThan = (operand1: unknown, operand2: unknown) =&gt;

Number(operand1) &gt;Number(operand2)

I would like to have a curried function in the middle that takes \`operand1\` and \`operand2\` and converts them into numbers so I don't have to do that explictly in my function body.

Is there a way to do this with function currying? Thank you
## [8][Is there a way to implement a custom behaviour when casting a class to a boolean?](https://www.reddit.com/r/typescript/comments/hsn96u/is_there_a_way_to_implement_a_custom_behaviour/)
- url: https://www.reddit.com/r/typescript/comments/hsn96u/is_there_a_way_to_implement_a_custom_behaviour/
---
If you have a class in TypeScript where whether it casts to true or false depends on a custom set of conditions, is there a way to implement that like you can implement something like `toString()`, where it will also work with regular if statements?
## [9][ReacType - export a baseline React or Next.js app written in Typescript in minutes!](https://www.reddit.com/r/typescript/comments/hsdcpp/reactype_export_a_baseline_react_or_nextjs_app/)
- url: https://www.reddit.com/r/typescript/comments/hsdcpp/reactype_export_a_baseline_react_or_nextjs_app/
---
Hello Reddit! My colleagues and I would like to invite you to check out our new open source project ReacType (version 3.0)! ReacType is a prototyping tool that allows you to create the skeleton of a Next.js or classic React application in minutes. Use our drag-n-drop interface to create reusable components, add basic layout and styling, and implement routing. Export your project as a fully functional Next.js or classic React application built with Typescript and functional components.

What’s new with this release:

- Export your code as a Next.js application with routing
- Build your application with a drag-n-drop interface and watch as your code is dynamically generated
- Save your project to the cloud

Check out our [github repo](https://github.com/open-source-labs/ReacType) for more information.

This is also a [link to a medium article](https://medium.com/@tylersullberg/letters-from-the-dark-world-of-react-boilerplate-5de9b4b8e2a3) one of our developers wrote.

Thank you!
## [10][Better syntax for optional return type value](https://www.reddit.com/r/typescript/comments/hshahi/better_syntax_for_optional_return_type_value/)
- url: https://www.reddit.com/r/typescript/comments/hshahi/better_syntax_for_optional_return_type_value/
---
    export const getClients = (): Promise&lt;ClientList | undefined&gt; =&gt; api.get&lt;ClientList&gt;({
        uri: '/clients',
        errorMsg: 'Error getting all clients'
    });

So I'm hoping to find a better way to express the above type than ClientList | undefined. Something that expresses the optional nature of the result. I think my API get() function might need a modification too, it returns the ClientList if successful and undefined if not (it contains the exception thrown and internally handles logging it and showing it to the user.

Anyway, I've always been a big fan of TypeScript in concept but kept sticking with JavaScript, so I'm trying to force myself to use it more.

PS. Here is the api.get function definition:

    const get = async &lt;T&gt;(req: RequestConfig): Promise&lt;T | undefined&gt; =&gt; {
        try {
            const res = await axios.get(req.uri, req.config);
            return res.data;
        } catch (ex) {
            handleError(ex, req.errorMsg, req.suppressError);
            return undefined;
        }
    };

&amp;#x200B;
## [11][Trying to model a function declaration that returns a class instance](https://www.reddit.com/r/typescript/comments/hs8544/trying_to_model_a_function_declaration_that/)
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
