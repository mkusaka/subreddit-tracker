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
## [2][I have released a web sdk to reduce video streaming costs by 90%](https://www.reddit.com/r/typescript/comments/jqazdm/i_have_released_a_web_sdk_to_reduce_video/)
- url: https://www.reddit.com/r/typescript/comments/jqazdm/i_have_released_a_web_sdk_to_reduce_video/
---
 Hi everyone,

I am working on a javascript sdk which can reduce video streaming costs of CDN by up-to 90% using a hybrid decentralized load sharing technology. I have opened it up for beta-access for developers to try it out. Looking for feedback w.r.t the technology and any features you would want to have.

A web demo is available here [https://api.peervadoo.com/test](https://api.peervadoo.com/test) . Click on **Add new peer** to see the tech in action,

**Javascript sdk link** :- [https://github.com/vadootvpeer/sdk-javascript](https://github.com/vadootvpeer/sdk-javascript)
## [3][Constraining the shape of a dictionary object's content?](https://www.reddit.com/r/typescript/comments/jqb2qe/constraining_the_shape_of_a_dictionary_objects/)
- url: https://www.reddit.com/r/typescript/comments/jqb2qe/constraining_the_shape_of_a_dictionary_objects/
---
I have a function which consumes dictionary objs. I'd like to constrain the shape of each entry in the dictionary (which is consistent across entries), though I'm not quite sure how...

I tried using `[key: string]` as the key, but that matches all entries, even invalid ones that aren't in the dictionary (i.e. `dict['doesntExist']`)

Does anybody have any insight on how to do this?


    type constraintObject_T = {
        [allKeysWhichExistOnObject] : {
            bar: string
        }
    }

    function ConstrainedObject&lt;T extends constraintObject_T&gt;(object: T){
        ...
    }


# Dict Examples
    const valid = {
        Foo: {
            bar: 'test'
        },
        Too: {
            bar: 'test'
        }
    }

    const invalid = {
        Foo: {
            bar2: 'test'
        },
        Too: {
            bar: 123
        }
    }
## [4][Can enums have functions as values?](https://www.reddit.com/r/typescript/comments/jq355k/can_enums_have_functions_as_values/)
- url: https://www.reddit.com/r/typescript/comments/jq355k/can_enums_have_functions_as_values/
---
Hi,

New to TS here. I am learning about enums, and I was wondering if TS enums could hold functions as enums. It seems like there's no reason it shouldn't be able to, except maybe that it is not possible to serialize functions and closures, (what about expressions where variables must only be params?) and I was looking at the documentation for typescript and couldn't find any information:

 [https://www.typescriptlang.org/docs/handbook/enums.html](https://www.typescriptlang.org/docs/handbook/enums.html)

Java seems to be able to support functions in enums. Was there a specific reason for this, or is there actually a way? I know I can use an object, but it's interesting and useful to map a set of functions to a set of choices.
## [5][Is there a generic interface to get all nested keys for an interface?](https://www.reddit.com/r/typescript/comments/jq2vj5/is_there_a_generic_interface_to_get_all_nested/)
- url: https://www.reddit.com/r/typescript/comments/jq2vj5/is_there_a_generic_interface_to_get_all_nested/
---
TypeScript 4.1+ is accessible in this scenario.

I've seen a few snippets that do this:

    interface X {
      a: {
        b: boolean;
      };
      c: string;
    }
    type Example&lt;X&gt; = /* magic here */;
    // Example&lt;X&gt; = 'a.b' | 'c';

I am wanting an array version of this:

    type Example&lt;X&gt; = /* magic here */;
    // Example&lt;X&gt; = ['a', 'b'] | ['c'];

Is this possible yet in TypeScript?

Use case:

    const x: X = {
      a: {
        b: true,
      },
      c: 'yes',
    };
    const getFromX = &lt;NestedKeys extends Example&lt;X&gt;&gt;(...keys: NestedKeys): TypeOf&lt;X, NestedKeys&gt; =&gt; {
      let temp = x;
      for (const key of keys) {
        temp = temp[key];
      }
      return temp;
    };
    const bool: boolean = getFromX('a', 'b'); // true
    const str: string = getFromX('c'); // 'yes'
## [6][Purify 0.16 released! - A Functional programming library for TypeScript](https://www.reddit.com/r/typescript/comments/jpr6pl/purify_016_released_a_functional_programming/)
- url: https://www.reddit.com/r/typescript/comments/jpr6pl/purify_016_released_a_functional_programming/
---
Link to changelog: [https://gigobyte.github.io/purify/changelog/0.16](https://gigobyte.github.io/purify/changelog/0.16)

Before the usual comment asking about a comparison with fp-ts that comes up with every release post - [here](https://www.reddit.com/r/functionalprogramming/comments/ebg4pc/purify_014_released_a_functional_programming/fb5uv16/).

Purify is becoming pretty much production ready, so this will be the last 0.x release. Every release post in this subreddit has always generated some nice feedback, so I'm looking forward to that.
## [7][How to define a dictionary that is indexed with symbols in typescript ?](https://www.reddit.com/r/typescript/comments/jpwiyr/how_to_define_a_dictionary_that_is_indexed_with/)
- url: https://www.reddit.com/r/typescript/comments/jpwiyr/how_to_define_a_dictionary_that_is_indexed_with/
---
[Here](https://www.typescriptlang.org/play?#code/MYewdgzgLgBAJgS2LAXDASgU1AJzgHggE8BbAIxABsAaaHBMAcwD4YBeGAbwF8BuAKH6JkAbQDKpCpQAUASgC6A-qEiwQUABaYcAESSouIgB4pi5KvJho6DRt3Zc+QA) is what I have tried so far. Why is it wrong?
## [8][How to preserve type when plucking from object, but account for undefined keys?](https://www.reddit.com/r/typescript/comments/jq2o4h/how_to_preserve_type_when_plucking_from_object/)
- url: https://www.reddit.com/r/typescript/comments/jq2o4h/how_to_preserve_type_when_plucking_from_object/
---
Consider the below code, the code `const foundKey = classDict[key]` should have the type `string | undefined`, because it's possible to enter in a string key that's not part of the `dict` object (e.g. `dict['bar']`). 

However, I can't seem to figure out how to accomplish this. Does anybody how to accomplish this?



# Function


    function classMatcher&lt;
      K extends keyof T,
      T extends baseClassDict_T
    &gt;(keys: K[], classDict: T): T[K] | undefined {
      for (let key of keys) {
        const foundKey = classDict[key]; // &lt;---- Type = string, should equal string | undefined
        if (foundKey !== undefined) return foundKey;
      }
      return undefined;
    }

    type baseClassDict_T = {
      [key: string]: {
        component: (prop: any) =&gt; JSX.Element;
      };
    };

# Usage

    const keys = ['foo', 'bar'];
    const dict = {
        foo: {
            component: //misc...
        }
    }
    classMatcher(keys, dict)
## [9][react native how to add IOS and android model top of the input field for register?](https://www.reddit.com/r/typescript/comments/jq2j6k/react_native_how_to_add_ios_and_android_model_top/)
- url: https://www.reddit.com/r/typescript/comments/jq2j6k/react_native_how_to_add_ios_and_android_model_top/
---

## [10][Type-safe dotted paths for objects, arrays, and nested recursive objects! (and intellisense)](https://www.reddit.com/r/typescript/comments/jpf9gr/typesafe_dotted_paths_for_objects_arrays_and/)
- url: https://www.reddit.com/r/typescript/comments/jpf9gr/typesafe_dotted_paths_for_objects_arrays_and/
---
Imagine you have the following \`get\` function:

https://preview.redd.it/ylkw3bvytox51.png?width=1340&amp;format=png&amp;auto=webp&amp;s=97c7ebc93f8dd488c03ac3bfefd9bf0d4cfd9411

Do you want type safe paths and return types? With auto-completion / intellisense?

[Implementation](https://preview.redd.it/g3zu2nm16px51.png?width=1640&amp;format=png&amp;auto=webp&amp;s=ed2d3099dca13581bf1f327de9571786aa54f440)

&amp;#x200B;

[Result with errors](https://preview.redd.it/vw0x8nqf5px51.png?width=640&amp;format=png&amp;auto=webp&amp;s=53ad8e01e22e0f121aee909ae14cf0acda6eb923)

&amp;#x200B;

It handles recursive objects and arrays too - at lightning speed.

[https://gist.github.com/millsp/1eec03fbe64592c70efa4c80515f741f](https://gist.github.com/millsp/1eec03fbe64592c70efa4c80515f741f)  


\`PathDot\` will ship soon with the upcoming version \`9\`. Leave a star on:

[https://github.com/millsp/ts-toolbelt](https://github.com/millsp/ts-toolbelt)
## [11][Fire a Promise and Collect the Result Later](https://www.reddit.com/r/typescript/comments/jpbawa/fire_a_promise_and_collect_the_result_later/)
- url: https://www.reddit.com/r/typescript/comments/jpbawa/fire_a_promise_and_collect_the_result_later/
---
I want to make a call to a service (and get a response from that service), do other stuff at the same time, then take the result from that service call and use it to populate some variables.

&amp;#x200B;

My current idea is:

`private async serviceCall(): Promise&lt;ServiceResponse&gt; { return await serviceHandle(); }`

`var servicePromise = new Promise(() =&gt; this.serviceCall());`

`doOtherStuff();`

... get the ServiceResponse and store it in a variable.

\---

How do you get the serviceResponse?
