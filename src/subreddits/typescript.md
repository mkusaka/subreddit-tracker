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
## [2][The Typescript Pick utility type! It constructs a type by explicitly picking props from another type.](https://www.reddit.com/r/typescript/comments/fd5j72/the_typescript_pick_utility_type_it_constructs_a/)
- url: https://i.redd.it/ld9xbryt8kk41.png
---

## [3][A Type-Level Lisp Interpreter (-ish)](https://www.reddit.com/r/typescript/comments/fda3pw/a_typelevel_lisp_interpreter_ish/)
- url: https://www.reddit.com/r/typescript/comments/fda3pw/a_typelevel_lisp_interpreter_ish/
---
The Eval type can perform basic arithmetic, comparisons, conditionals (if expresions), define local variables and lambda expressions.

e.g. this type `_` "evaluates" to `16` at compile time.

```
export type _ = Eval&lt;
  [Let, 'square', [Lambda, X, [X, '*', X]], ['square', 4]]
  &gt;;
```

[Typescript Playground](https://www.typescriptlang.org/play/#code/PTAEBUE8AcFMFoAysBusA2oCSA7ALrAOawBOoAsgIZ4AWoAljhDLAM4DGJ90eAUCKHhDhI0WPET4-YNNBzQAPlAAxAPZkAwpXTsArulU4A5KwYBbM+rzbQrVfoA0oQrranasUDl1mARqWwjM1BaRgBrRkJQVQAzWXklOSFQACVVX1U8UAARVHp0aXi5JQAichKixXlkgAUafPRuUAApADoVEkhCgRr0WEpWTwATQyMs3UGQ+tNGUGgSVSHddjx6Q1AACixQdkpjLPozSkJGTwB3GmpQSHsjIdAR2TP6WgY8Vo-QGPoAD1gh+CsegAL3+oDQK3UrAA-ABKN4mUB7UCwH5wLhmWD4BhMGiqM6gMzLOhnTy7HCyXY6fTUTxXSghFjwPpoTCzPAsDhcHitbpgADqsCMJE8IvmbCxqxwUUYBGIJFMAy8bAI90oJBIlEgrCcF1IngABgAtUiqA0Mdw0TywMw8SBIjVa2R7e4AagAjBbQGKRYN8GDfPazproNBIkimLKiAFZntMlayCyMNEYqBfJr2GFYHhWLyQLINuBpqALqo+oTFvpPPRTL4BvQqeh7cisexFuHYqAANpwOMOl40TGrdgAXQ2NDweGgrAAXCAsa1nhE4EN6JRWupCMAl-RgDV+jhVAB9Sg-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-IvJEozU45JfAWB5ElwWa6-EpAnCcXYOwZd5rXCpwi+nYNJ6QtkGvRGySK70MCIWGejiQcgsMwiwodLbdvPQC435gLuXBiutpjUJLLrMsLpy3l8LkXy5vKSyBlh4S-UWsDUYBLBHSJeZ85KQLVWBF1ZoiVsreB-PGLG+pCbpISBOAysdWT+NTgAEddDaDPoshQPbN7sR7Y58Uk57T9FYJRlKhxjE2klM6UAJ2nYKYtQwfFpsrFDvnmQPY9oozylMC7DSTh2A+yc0iRQP23so2hk7F7FtiXIFYKwIsexQokBmjt7QS6aJepQ6YGldz-tMv9Z4PmA2haJQAOIilpCQdHOBMfY92+gPHSEUdo8uMzrHOP2dPoZutojTAPDqCDH8qu4pUfgtXETIE6wlT2UYALSYsRZASbo0mnAmmDhMHmG9VYkwXigTYL+H+osr4bFGpGOy0xj70HuD-OlWRDCuaONOAGgOAhy5rI+Ol6Koc-bAFYvUjnKDsYEL4JwesrHqjpFfb3CpYRuKKIlAAiroTIql8DsWyE1R3zMnD5-l+oAJZqcgF7sMlwnqXPDcxL4XzefNG9iU3il09Eq89V-UHS29HMmW4z5qy0gQyNuMezBVUVoADTk4NKnjxiUjw0Xap1hr2X2LFVUE4LfZP3uU7tWPsAQctNWkmBkV4Io8BT5nwN+f+TPGgGX0hVfHestd-pZjJKcU6FD6+IftPJiZ-FQXeNfTvRdLfPvO9DGQfNuKnI2YSaNdAe4RKBYTIG9EjMjc4LUXXEIMBZKX8fwczSgUMZsDIPRF4doTiQgBBHVMAReL0CSSGB-JfKbRBLIFLabXzIqZKHfGgPKPfFGA-ENI-duCg-vZFS6eTUXfGexFKS4NBOsVgLMPABfa5JiNITIdiZiYLZbemZiNfLg0belSrRbWrPQuhUw19NlQAiiYAzQjiTaCcPggQ61P2NgraQwkbXPEwvLGqTKSwhZaw0fI2ZifET8LIfNS4KVCWMEe+RSUVawfARUazdEbXJRGQm0UWSkcWJgqHEVZSHPLIa7WeH4aIUXVQWxTsE3exEsLgMzIEeRVFMAVEaXBXafPWVAZGWkMyMwGTAUIUK+JQqSfGEo6wMo2yV+UwQgTUCwdUGPHjAwQYVzJUfNckB-VosgeCGcHtexYIAAHxEivkOPh02Klxr3mCKEbBaR2POOiDSPsQuKCg8UqLOUeN5iQiMFdFbkOKMHgB+ISVrQBKMAAD1gSTtgTDtkkPFqBVBghtjG1QBDjPRDiiIkTQAABmdEgAFnRIAFZ0SAA2dEgAdnRIAA50SABOdE90JtCvFpeQQ48ndEgbPkYlGaUmdAObdnGaJwGaHAFAKLUADlPYhsc4twc+B0TwVNExM5TAbwPwAIEKVgFdAQNY+wZA2wGxUrEgLISwRzPTM-RbTSdZYsFeLEewBaX7NUsABkExPAcsMSSwFdWEswPmGaNfVCeCbmH9bY0AT0lLHCH0uvIrD0tfEiEMrfMMgMtfDEorbmfwmMwM2vbEhMkAiqZMtfPE9MgaHBf0lMqlUwQk9MyaLMlLEk0sm3cs2vck9MpwvgAstfKk9Mw6NLJsoM70+QRMrEGsoskIC1dM8nPs6hNkkMsc+QcMlLLMG4VMAUlAeswUtfFZNZIrDLecrsGaEcNc4IoKf0kfd9W037L4XQC6WVRsHRfNI4FxbZfGZBYIJUBka-e7TxQ2OQfMd8s1Lsc4pwV40zXvO47cz8ldRsI8XAb4H4Kc2vLselO5L4owKAgIIOIC7s9pAOTk7QHkvxec+rDC7kirfkwU4REc0JGCiQhmJwIwIExCsgZC9M8AxwrkrCu5HC-kpigigMoiwrDs6C2Cyi-4mi9uFCjLA1R1diPC5iwioUtizCjinC7iyc2MlLMi6AuCsEwSuikMow3PCSgLFioimS-C+bQXeSkilmFSgIOCqEjS4SuhOnZjRnHnFnfncS9i+bfS6SgMtygXKS4ini-siysgOCiEmy9MrnJnZytnVy2S9y3ywyrCkyrimMg8oWT8-yMjXwIYJUEUTEJUsgeyLEImBYHATEbEDwCXRzaVMERgDSKRSgEBI88q6UnYJYiqKtTLbMDkbUxzBgleBkEUHbNeB7bEZBV+aUVUtK7BWBI0L4X4f4eAUrWURSPwRgWEvc9KvwLKqCgKowe8TaygBCnEaAmodUMjDSxYSAWy7mE62YtfMVSIHc7mcORk+CaAU690gRd69OF6tMC6+OC676l6wqj0wUwG+CMuHc4fN9ZKt9VKo49Aba6hQKlQU8wSkrEgK6hmLWSAaKoy9nZQVGzizy3S+bdG3yhSoDFK0iNKjYeHAgbebGVpWAeELo4hWkUwEoemvAEoAGJyG+SYWo1m3bAgI82AAcAITmq0HAHmhSEoNJCoOQaGRG0iowaGQ6+lXaCyQSqeHXI6gIGaENTGkm9nTWjTImpQFLYcxk7mY2zfKW3yyGrymK3kkNB2-0+AldPoVQpS6C3a7MdW8i7xQSgANTJnOqGEuvTO8TussgeutqdrxpLguqktAAADJQBQ5uxvFYAAcmBvERwPTvL2JQ7iFFscLQAy4KbFKDyYabCK5CykbdqMqsrDr7rpQY8cBI6Msm79qkMfaAr-tc1Y726Iwu66EjBQLwLfg+6G6Va1aO7+VR6F6UKG41bsN+7G6vbW7h7CAF7l74Mt717Z7zLB7R6V6ElGw+6jA3Tx5qcmJvMqV2Jo6Us27d7QBg7fKaJwBHccaV1nr5Auxs7c7QB86+Zg704y406itQozAXhxLBSnBvFCtjpEpv6I72IUh6YUg18UgpF1Ahh2JT7SdwN-74JAGIUZzOwUgC7UhAGBJwb16qbiVOJsa2owt6ZlA19-63rZi+Z6UbqyN04MgI6+HyKg4AaV1ga9ayB5yGIIaMsBHggX6d76yi7xGI7+SUM853qwtyba7SBGHYbZBjbwEr50R7zYEqQ1FtdcCTcr8p9FRlQLIWNHQgxZJcF8ZJILJwwxIeA6AFN8Y9gnEAhTT9EPpsR7TTlNQFSfB-B8qWB+iIBRiywtJwUsAjBmo5FPAbFwFXgFMHFNRRULGnybFywvHM0DRmIDRj5apbFZRbt8AUoSmXzOw+hpRXhKnzQskpb9IX57IDQ2nCBaBzRqjTFw10nmpY6wRajU1ajzH1Bgh7I-z0jyiQhSnPBBnaBbHijaoLxcF5CTdGhYGcxCRiQUR1RGhSBeRWirBH90YcBcHoBN96ZwA18WTDiBs6FXm+Z7n5JkDN8nBYIFAuw-b2maBb7BIwBNV7gAAJJVO84q13VMV5jYC4Bsfxxx588sDZOdJjbudiGFnfNff7IF52ccfoIYPmQl0AD4VoawfIPmcAeEBCJQf7Fm2vDYFdWlxoCyUR6A9GXgZl1lzu8DAVymox2QDQJSUgHSJxzNLFzwDYQKWBnlKUGUJgFsW0LqhVzZAcKHBV0ybwa6Dxq4S4UwTZugWAfneFzwcp8MJMAoXVOQU8v5+4WCOQN6b7JECLSoF1ssN13QOCT1miP1rU3QTYcUOATeI1zAT12dAQSGQgdoUNt1vCAiFpd13iGNpwGN7cipaxW9f17sdCNN-CAiTN7sbNrwfQdAHNmtut66bcxKX5-1vxFrYlzu0lk1FdZCVONfd1thCLXc9OA2ShZSz-OCbmFt-5r-fOKNghqtlrYRYdj9SuZSyds2Pmd0CeRKRAZuhkT4vazKg6+qCiAADRoiMB+CMFPc8GQCyEPf9sI1sBSlAOgGBTozsVIFtUulTWDBINMBMa2QxgeffA7qdyFBFGVgpZu2WuOblQ6JoMUSyDkUsBwEBCxEmD1gNL2w5F4Bub1MfzwvEvpnucedxqwr5IzsruFxYnCLWORBAVgSFrWv8noD6emDUPkAABJH9gDiP0IMTKLvinBsSRxKLqLQAiIRwRF1C7CaJ3Q0z8PONCPWCkJiOV0ux73KLWAccRRDqtP92nAz3AWTPATDqz38JAWTA9OhRRP8IV0sogA)
## [4][Introducing AdonisJS (v5 Preview)](https://www.reddit.com/r/typescript/comments/fdaojh/introducing_adonisjs_v5_preview/)
- url: https://blog.adonisjs.com/introducing-adonisjs-v5
---

## [5][Can tsc compile async functions to Promises instead of generators?](https://www.reddit.com/r/typescript/comments/fd978u/can_tsc_compile_async_functions_to_promises/)
- url: https://www.reddit.com/r/typescript/comments/fd978u/can_tsc_compile_async_functions_to_promises/
---
Project is compiled to target ES5, no Babel. I polyfill promises for IE11 via whatwg-fetch (and others with core-js)

The problem is that debugging async functions results in a lot of stepping through awaiter functions and gets tiresome. 

I'd like for the typescript compiler to transform the async functions to good ol promises. Is this even possible? 

I've found that __generator [can be swapped out](https://www.typescriptlang.org/v2/en/tsconfig) with my own implementation, but there's not really any examples of this.
## [6][Example serverless data pipeline for crawling PDFs from the Web and transforming their contents into structured data using AWS Textract. Built with AWS CDK + TypeScript.](https://www.reddit.com/r/typescript/comments/fcy30x/example_serverless_data_pipeline_for_crawling/)
- url: https://github.com/aeksco/aws-pdf-textract-pipeline
---

## [7][[Help Wanted] Overriding a function exported by a 3rd party library causes TS2540 error](https://www.reddit.com/r/typescript/comments/fdb9qa/help_wanted_overriding_a_function_exported_by_a/)
- url: https://www.reddit.com/r/typescript/comments/fdb9qa/help_wanted_overriding_a_function_exported_by_a/
---
As the title stated, I encountered some strange behavior when overriding 3rd party module exports.

I'm trying to wrap Redis' createClient method, so I don't have to specify connection options each time I create a connection to the database.


    import * as redis from 'redis';
    let _createClient = redis.createClient;
    // this gets underlined with red, telling me I cannot do this
    //        |
    //       \|/     
    redis.createClient = (opts?: redis.ClientOpts) =&gt; {
        return _createClient({ host: 'redis' })
    }


I tried reading up on module augmentation, but when I tried to add an implementation to a function defined in a .d.ts file I got the same error.

TSC also throws an error telling me the same thing as VSCode. Now comes the strange part, it actually seems to be working right now, even though the errors.

I spent a good 2-3 hours trying to solve this, and I can't seem to find the solution for the problem. All I need is a point in the right direction, and I'll figure it out.

Thanks
## [8][Becoming a GIT pro. Part 1: internal GIT architecture](https://www.reddit.com/r/typescript/comments/fd9za9/becoming_a_git_pro_part_1_internal_git/)
- url: https://indepth.dev/becoming-a-git-pro-part-1-internal-git-architecture/
---

## [9][import of JSON file in TS 3.8.2](https://www.reddit.com/r/typescript/comments/fd6dps/import_of_json_file_in_ts_382/)
- url: https://www.reddit.com/r/typescript/comments/fd6dps/import_of_json_file_in_ts_382/
---
I'm using a JSON file to act as a data source for my early development.  I'm importing the JSON file with the import statement and functionality added in TS 2.9.

I have the following code and data file.

# database.ts

    import * as test from './test.json';
    
    export class people {
      list : any; 
      constructor() {
        this.list=test;
      }
    }

# test.json

    [ { "ID": "me", "Alias": "D"  },  { "ID": "You", "Alias": "U2"  } ]

I even added the following config options to my tsconfig.json

    "resolveJsonModule" : true,
 "esModuleInterop": true,
 "allowSyntheticDefaultImports": true,

The problem is that list is a module and the default export has my data.  This is from the Chrome web console.

    people {list: Module}
    list: Module
    default: (2) [{…}, {…}]
    __esModule: true
    Symbol(Symbol.toStringTag): "Module"
    __proto__: Object

&amp;#x200B;

How do I pull out the default value from the module so I can assign it to list?  The default value contains my list of objects.
## [10][strong-mock: Create type safe mocks from types and interfaces that enable refactorings to not leave tests behind](https://www.reddit.com/r/typescript/comments/fcz4f4/strongmock_create_type_safe_mocks_from_types_and/)
- url: https://github.com/NiGhTTraX/strong-mock
---

## [11][Just learned about the Omit utility type! It constructs a type by picking all properties from an object and then removes any specified props.](https://www.reddit.com/r/typescript/comments/fcmz11/just_learned_about_the_omit_utility_type_it/)
- url: https://i.redd.it/9ckc6qyi2dk41.png
---

