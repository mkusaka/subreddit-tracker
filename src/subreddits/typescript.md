# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Announcing TypeScript 3.8 Beta](https://www.reddit.com/r/typescript/comments/emxvdn/announcing_typescript_38_beta/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-beta/
---

## [3][Diagram of every possible TypeScript type](https://www.reddit.com/r/typescript/comments/emxi8j/diagram_of_every_possible_typescript_type/)
- url: https://gist.github.com/laughinghan/31e02b3f3b79a4b1d58138beff1a2a89
---

## [4][Collection of resources related to Typescript internals](https://www.reddit.com/r/typescript/comments/emrdkq/collection_of_resources_related_to_typescript/)
- url: https://typescript.tools/
---

## [5][How can I define a field in an interface that can be one of 2 slightly different values?](https://www.reddit.com/r/typescript/comments/emy5td/how_can_i_define_a_field_in_an_interface_that_can/)
- url: https://www.reddit.com/r/typescript/comments/emy5td/how_can_i_define_a_field_in_an_interface_that_can/
---
Hello, I'm trying to define a field in an interface where values are slightly different. An example:

    export interface HelloWorld {
      aField: firstType | secondType;
      ...
    }
    
    export type firstType = 'hello' | 'world';
    export type secondType = 'hello' | 'another' | 'world';

I recently had to create \`firstType\` where I don't want its value to be a string value of "another". The definition of the 2 different types and use of a union was my attempt at getting this to work, but I get a Type error code TS2322. Example error msg:

    Types of property 'aField' are incompatible.
      Type '"hello" | "another" | "world"' is not assignable to type '"hello" | "world"'.
        Type '"another"' is not assignable to type '"hello" | "world"'.

If I were to just define the interface using type \`firstType\`, tests will pass, no error. But this is incorrect for me because this field can have a value of "another" in certain cases.

Any help is appreciated! Thank you
## [6][TS2339: Property 'ctxSession' does not exist on type 'OAuthStartOptions'.](https://www.reddit.com/r/typescript/comments/emteem/ts2339_property_ctxsession_does_not_exist_on_type/)
- url: https://www.reddit.com/r/typescript/comments/emteem/ts2339_property_ctxsession_does_not_exist_on_type/
---
Been staring at this one for 20 minutes and I guess I must not fully understand TS inside classes. Anyone know what's causing it? Setting the property to 'any' I thought would allow it to be set to anything.

    export default class OAuth {
      protected keys: Keys;
      protected req: Request;
      protected res: Response;
      
      public ctxSession: any;
      
      constructor(req: Request, res: Response, keys: Keys) {
        this.keys = keys;
        this.req = req;
        this.res = res;
      }
      
      public authenticateThenCaptureContext() {
        shopifyAuth({
          prefix : "/shopify"
          , apiKey : this.keys.SHOPIFY_API_KEY
          , secret : this.keys.SHOPIFY_API_SECRET
          , scopes : this.keys.SCOPES
          , afterAuth(ctx: any): void {
            this.ctxSession = ctx.session // title error
          }
        })
      }
## [7][Since most companies have united behind TS is possible that it will replace JS in the browsers in the future?](https://www.reddit.com/r/typescript/comments/emupty/since_most_companies_have_united_behind_ts_is/)
- url: https://www.reddit.com/r/typescript/comments/emupty/since_most_companies_have_united_behind_ts_is/
---
Also since it's a superset of js
## [8][casting a variable from a type to a subtype](https://www.reddit.com/r/typescript/comments/emnzj9/casting_a_variable_from_a_type_to_a_subtype/)
- url: https://www.reddit.com/r/typescript/comments/emnzj9/casting_a_variable_from_a_type_to_a_subtype/
---
Given some type and an extension of that type:

```
interface Item {
  type: string
}

interface ItemTypeA extends Item {
  valA: number
}
```

Sometimes I need to take a variable of the type and cast it to the more specific type via the `as` keyword:
```
function fn1(item: Item) {
  if (item.type === 'A') {
    const itemA = item as ItemTypeA;
    const { valA } = itemA
    // ...
  }
}
```

Alternatively, it can be done via a type checking function:

```
const itemIsA = (item: Item): item is ItemTypeA =&gt; item.type === 'A';

function fn2(item: Item) {
  if (itemIsA(item)) {
    const { valA } = item;
    // ...
  }
}
```

The first way seems more convenient, explicit, and doesn't require an extra function. With the second way, even though the compiler knows that the item has been casted to a new type, a human reader must either infer by the checker-function name or see what the checker returns.
## [9][How to implement Hierarchical types](https://www.reddit.com/r/typescript/comments/emu4ox/how_to_implement_hierarchical_types/)
- url: https://www.reddit.com/r/typescript/comments/emu4ox/how_to_implement_hierarchical_types/
---
I'd like to do this:

    type Predicate&lt;T&gt; =
         ['ANY', ...Array&lt;Predicate&lt;T&gt;&gt;]
        | ['ALL', ...Array&lt;Predicate&lt;T&gt;&gt;]
        | [keyof T, '=', number]
        | [keyof T, '&gt;', number]
        | [keyof T, '&gt;=', number]
        | [keyof T, '&lt;', number]
        | [keyof T, '&lt;=', number]
        | [keyof T, '&lt;&gt;', number]
    ;

tsc complains that Predicate circularly references itself

Oddly vscode correctly hints the creation of instances of Predicate&lt;T&gt;

Is there some other way to do this?
## [10][Dealing with values you don't know the shape of (from being undocumented)](https://www.reddit.com/r/typescript/comments/emqpks/dealing_with_values_you_dont_know_the_shape_of/)
- url: https://www.reddit.com/r/typescript/comments/emqpks/dealing_with_values_you_dont_know_the_shape_of/
---
Current situation is I'm developing a shopify app on Node/Koa.js. It returns a context object with some important details such as the store name and other things. 

I could tell Typescript `ctx.session` is `any` to turn off type checking. But then what is the point of TS, and also it's good for me as a developer to actually learn the shape of what I'm resolving for clarity.

Documentation on ctx.session is lacking:

[https://github.com/koajs/session](https://github.com/koajs/session)

[https://www.npmjs.com/package/@shopify/koa-shopify-auth](https://www.npmjs.com/package/@shopify/koa-shopify-auth)

Faced with this situation how do you guys proceed? I am thinking I should just run a test to see what comes back, log it and write an interface based on that.

If this is a situation you guys would just mark as `any` I am open to that advice too.
## [11][[Request] Fairly new to TS, need help with a problem involving generics](https://www.reddit.com/r/typescript/comments/emrryl/request_fairly_new_to_ts_need_help_with_a_problem/)
- url: https://www.reddit.com/r/typescript/comments/emrryl/request_fairly_new_to_ts_need_help_with_a_problem/
---
EDIT: If the formatting is messed up, here is a paste of JUST the code: [https://paste.ofcode.org/MqUSEq5cyyUgX9wHednJkm](https://paste.ofcode.org/MqUSEq5cyyUgX9wHednJkm)

&amp;#x200B;

I am using the library easy-peasy for a project at work. Most of my requests follow a similar form (data, status, a thunk to call the api, and listeners to update status based on the thunk), so in js, I used a function to generate most of it. I am trying to convert this to typescript now, but I am having issues. My code looks like this (some additional properties omitted until I figure out my problem):

    import { actionOn, ActionOn, Thunk, thunk, TargetResolver, ActionTypes } from 'easy-peasy';
    import RequestState from 'utils/request_states';
    import { StoreModel } from 'store/store_root';
    import RequestResponses, {PickResponse, Stub} from 'interfaces/request_responses';
    
    
     interface RequestModel&lt;T&gt;{
        status: RequestState;
        data: T | null;
        error: null | Error;
        api_request: Thunk&lt;RequestModel&lt;T&gt;, any, any, StoreModel, Promise&lt;Stub&gt;&gt;;
        onStart: ActionOn&lt;RequestModel&lt;T&gt;, StoreModel&gt;
      };
    
    
    export function StandardRequestGenerator&lt;T&gt;(init: T, api_req: any): RequestModel&lt;T&gt; {
    	return {
    		status: RequestState.VOID,
    		data: init,
    		error: null,
    		api_request: thunk(async (actions, payloads, helpers) =&gt; {
                
                return api_req;
            }),
    		onStart: actionOn(
    			(actions) =&gt; {
            return actions.api_request.startType;
          }
    			(state) =&gt; {
    				state.status = RequestState.LOADING;
    			}
    		)
    	};
    }

The idea being that the generic would be the type of the data returned from the API request. I think that the issue is that  since the type is unknown, what easy-peasy uses to map state and actions doesn't know what to do with it. This appears to be backed up the fact that if I remove the generic and use something like number, it works fine. Anyways, I get the following error:

    Property 'api_request' does not exist on type 'ActionMapper&lt;Pick&lt;{ [P in "error" | "api_request" | "onStart" | ("status" &amp; { 1: "data"; 0: never; }[Extends&lt;T | null, object&gt;]) | ("data" &amp; { 1: "data"; 0: never; }[Extends&lt;T | null, object&gt;]) | ("error" &amp; { ...; }[Extends&lt;...&gt;]) | ("api_request" &amp; { ...; }[Extends&lt;...&gt;]) | ("onStart" &amp; { ...; }[Extends&lt;...&gt;])]: Re...'.ts(233

VSCode shows the shape of actions as:

    (parameter) actions: ActionMapper&lt;Pick&lt;{ [P in "error" | "api_request" | "onStart" | ("status" &amp; {
        1: "data";
        0: never;
    }[Extends&lt;T | null, object&gt;]) | ("data" &amp; {
        1: "data";
        0: never;
    }[Extends&lt;T | null, object&gt;]) | ("error" &amp; {
        ...;
    }[Extends&lt;...&gt;]) | ("api_request" &amp; {
        ...;
    }[Extends&lt;...&gt;]) | ("onStart" &amp; {
        ...;
    }[Extends&lt;...&gt;])]: RequestModel&lt;...&gt;[P]; }, FilterKeys&lt;...&gt;&gt;, "1"&gt;

The following is easy-peasy's action mapper, which is unfortunately gibberish to me.

    type ActionMapper&lt;ActionsModel extends object, Depth extends string&gt; = {
      [P in keyof ActionsModel]: ActionsModel[P] extends Action&lt;any, any&gt;
        ? ActionCreator&lt;ActionsModel[P]['payload']&gt;
        : ActionsModel[P] extends Thunk&lt;any, any, any, any, any&gt;
        ? ActionsModel[P]['payload'] extends void
          ? ThunkCreator&lt;void, ActionsModel[P]['result']&gt;
          : ThunkCreator&lt;ActionsModel[P]['payload'], ActionsModel[P]['result']&gt;
        : ActionsModel[P] extends object
        ? RecursiveActions&lt;
            ActionsModel[P],
            Depth extends '1'
              ? '2'
              : Depth extends '2'
              ? '3'
              : Depth extends '3'
              ? '4'
              : Depth extends '4'
              ? '5'
              : '6'
          &gt;
        : unknown;
    };
    
    type RecursiveActions&lt;
      Model extends object,
      Depth extends string
    &gt; = Depth extends '6'
      ? Model
      : ActionMapper&lt;
          O.Filter&lt;
            O.Select&lt;Model, object&gt;,
            | Array&lt;any&gt;
            | RegExp
            | Date
            | string
            | Reducer&lt;any, any&gt;
            | Computed&lt;any, any, any&gt;
            | ActionOn&lt;any, any&gt;
            | ThunkOn&lt;any, any, any&gt;
          &gt;,
          Depth
        &gt;;

I am completely stuck on what to do here to get around this and would really appreciate it if someone could help me out. Thanks.
