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
## [2][Isomorphic TypeScript client for PostgreSQL](https://www.reddit.com/r/typescript/comments/j9nlal/isomorphic_typescript_client_for_postgresql/)
- url: https://www.reddit.com/r/typescript/comments/j9nlal/isomorphic_typescript_client_for_postgresql/
---
I've developed this open source library which allows safely exposing PostgreSQL APIs with minimal amount of code. Features:

\- CRUD operations

\- Subscriptions

\- Optimistic Client-Server data replication

\- Security Rules

\- Generated TS Types for database schema

WIP: Auth + Android library

&amp;#x200B;

I've recently rewritten it in TypeScript and I am now open to any feedback or comments you might have.

Thanks for reading!

https://prostgles.com/
## [3][Easy encryption of user data in their web browser with `lucertula@1.0.0`](https://www.reddit.com/r/typescript/comments/j9czri/easy_encryption_of_user_data_in_their_web_browser/)
- url: https://www.reddit.com/r/typescript/comments/j9czri/easy_encryption_of_user_data_in_their_web_browser/
---
Hi r/typescript!

I have been working on a library to simplify in-browser encryption/decryption for some time and I present you \`lucertula\`:

[https://github.com/dipasqualew/lucertula](https://github.com/dipasqualew/lucertula)

The library aims to wrap established encryption algorithms, namely symmetric and asymmetric encryptions from the folks at \`openpgpjs\` and the AES GCP implemented via the WebCrypto API, and pair them with various storage solutions (localStorage, sessionStorage, etc). I believe encryption should be easy and available to anyone, but more often than not excellent encryption libraries don't offer "plugin" encryption/decryption methods ready to be used by developers.

As this is a library "around" encryption (but doesn't implement encryption) I would be grateful if more eyes were pointed at it and any feedback is of course welcome'd

The library is written in TypeScript (tests are stil lin JS but I am in the process of migrating them as well). It also supports tree-shaking so if you don't need OpenPGP your bundle won't suffer the significant size of \`openpgpjs\`.

\--

I wrote this library to avoid the boilerplate code to encrypt to and decrypt from local storage in each of my projects, so that's my personal use case. Hopefully it can be helpful to others

Cheers!
## [4][Is transient life cycle really needed in a DIC ?](https://www.reddit.com/r/typescript/comments/j9nbtr/is_transient_life_cycle_really_needed_in_a_dic/)
- url: https://www.reddit.com/r/typescript/comments/j9nbtr/is_transient_life_cycle_really_needed_in_a_dic/
---
I have made a DIC from scratch , and I can only register factories (i.e. functions that returns a value) to it .

Every time a factory is getted from the DIC , the returned value of the factory, gets memoized so that whenever the factory getted from the DIC again , the memoized value is returned . That is the singleton life cycle . If there was no memoization then that would have been transient life cycle.

I am coming to the conclusion that transient life cycle is not really needed and in fact sometimes can make testing difficult .

Am I missing something here?
## [5][Node TS nested module import not working as expected](https://www.reddit.com/r/typescript/comments/j9io93/node_ts_nested_module_import_not_working_as/)
- url: https://www.reddit.com/r/typescript/comments/j9io93/node_ts_nested_module_import_not_working_as/
---
Hi all- hopefully someone can help me figure out what I'm doing wrong here.

I'm working on an app that communicates with multiple devices via UDP. Fairly normal so far. However, I'm looking to make it a bit more modular and run as more of an engine based on a user-provided JSON file. I already have modules built out for each of the possible devices we could be talking to and I'm looking to figure out how to dynamically instantiate new objects dynamically.

I have a gist set up with a basic structure of the files and what I'm trying to do here, along with the errors I'm getting when trying to build the code: [https://gist.github.com/ericboxer/cd31d80d9942eb1c4bc815daf54a8677](https://gist.github.com/ericboxer/cd31d80d9942eb1c4bc815daf54a8677)

Please let me know if you have any questions, or if there is a better way I might be able to achieve this.

I also originally posted this in r/Node, but I think this may be a more relevant topic here
## [6][Help with class declaration inside module declaration](https://www.reddit.com/r/typescript/comments/j9dqqq/help_with_class_declaration_inside_module/)
- url: https://www.reddit.com/r/typescript/comments/j9dqqq/help_with_class_declaration_inside_module/
---
    declare module 'node-mailgun' {
      export default class Mailgun {
        constructor(apiKey: string);
    
        sendText(
          sender: string,
          recipients: string | string[],
          subject: string,
          text: string,
          servername?: string,
          options?: object,
          callback?: (err: any) =&gt; void,
        ) =&gt; void; // : void results in same error
      }
    }

I'm getting lints on the fat arrow method return type. But when using colon syntax it shows the same error:

    '{' or ';' expected.ts(1144)

What is the fix for this?
## [7][Generic type for type conversion](https://www.reddit.com/r/typescript/comments/j93t98/generic_type_for_type_conversion/)
- url: https://www.reddit.com/r/typescript/comments/j93t98/generic_type_for_type_conversion/
---
Hi,

I have some custom types like

    type MyType = {
        id: number;
        title: string;
        checked: boolean;
    }

I am trying to make a generic type (if possible) that will convert any boolean property in the type to binary (0|1).

&amp;#x200B;

I was thinking of something like the following

    type BoToBiFormat = {[x:string]:any}
    type BoolToBinary&lt;T extends BoToBiFormat&gt; = {
        [x in keyof T]: T[x] === boolean ? 0|1 : T[x];
    }

but that obviously didn't work. Is there a way to make such a type?
## [8][How to get intuitively comfortable with TypeScript?](https://www.reddit.com/r/typescript/comments/j8vwzf/how_to_get_intuitively_comfortable_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j8vwzf/how_to_get_intuitively_comfortable_with_typescript/
---
I've seen the docs and have played around with TS in toy projects, in an attempt to add it to my repertoire. I come from a background where my projects didn't use TS at all. How does one develop an intuitive sense of using TS for production?

While I realize the significance, I find it hard to bridge the gap between my amateur knowledge and being production-ready as the rest of my skillset.

I keep having to double-check what to type something while I'm implementing, sometimes resorting to online searches especially for imported libraries. There's also the pain of not knowing for sure if I typed it the right way. Is it normal for it to be time-consuming at the start?
## [9][Vue and TypeScript](https://www.reddit.com/r/typescript/comments/j91mz2/vue_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j91mz2/vue_and_typescript/
---
I am totally new to typescript i just wanted to do some side project to learn a bit more of it with vue. But i cant even configure my first app i am having problems with vue imports. I have made shim.d.ts but nothing is fixing the import error on "import App from './App.vue'" ....

Is there any straightforward tutorial how to set up my vue app with typescript , or if anyone could give me suggestion how to fix this and avoid these kind of errors in future id be really thankful.
## [10][Bug in type definitions for HOF?](https://www.reddit.com/r/typescript/comments/j8xnfz/bug_in_type_definitions_for_hof/)
- url: https://www.reddit.com/r/typescript/comments/j8xnfz/bug_in_type_definitions_for_hof/
---
Hello, I'm currently writing a little Axios-like wrapper around the Fetch API for the size benefits. To do this, I enlist a higher order function that differentiates body-less methods like GET, DELETE, OPTIONS, etc. from POST, PUT, etc. 

Here's my code (roughly)


```
type BodylessMethods = 'get' | 'delete' | 'options' | 'head';
type BodyMethods = 'post' | 'put' | 'patch';
const isOfBodylessMethod = (method: string): method is BodylessMethods =&gt; ['get', 'delete', 'options', 'head'].includes(method);

...

private request = (method: BodyMethods | BodylessMethods) =&gt; {
  return isOfBodylessMethod(method)
    ? (resource: string, params?: Record&lt;string, unknown&gt;) =&gt; {
        ...
      }
    : (resource: string, data?: Record&lt;string, unknown&gt;, params?: Record&lt;string, unknown&gt;) =&gt; {
        ...
      };
};

public get = this.request('get');
```

The problem I'm running into is that, for some reason the bodyless methods (`.get()`, `.delete()`, etc) throw an error if I make `data` a non-optional parameter in that second function, and I don't quite know why. I'm given `TS2554: Expected 2-3 arguments, but got 1.` Why would 2 arguments be expected? The first function has 1 required, 1 optional, the second 2 required and 1 optional. Surely 1-3 arguments should be expected?

Now, this can easily be fixed by keeping `data` optional, and should likely be solved that way as the spec doesn't require a body with POST, PUT, etc., but I'd like to know why this issue is popping up. I don't understand where it is coming from.
## [11][Self casting from json?](https://www.reddit.com/r/typescript/comments/j8vjk1/self_casting_from_json/)
- url: https://www.reddit.com/r/typescript/comments/j8vjk1/self_casting_from_json/
---
I have a function that accepts a class object. If I pass a json inside, typescript seems to automatically cast the json object into the object declared in the parameter. Is this reliable? Or should I map it and pass it in?
