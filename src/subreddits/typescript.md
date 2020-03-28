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
## [2][Announcing TypeScript 3.9 Beta | TypeScript](https://www.reddit.com/r/typescript/comments/fq4l4r/announcing_typescript_39_beta_typescript/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9-beta/
---

## [3][Help needed! stuck on this issue](https://www.reddit.com/r/typescript/comments/fqinx4/help_needed_stuck_on_this_issue/)
- url: https://www.reddit.com/r/typescript/comments/fqinx4/help_needed_stuck_on_this_issue/
---
    this.cartService.addProduct(this.product == 
      this.db.collection('products', ref =&gt; ref.where('tagid', '==', 
      this.nfc.bytesToHexString(event.tag.id))));

So what I'm trying to do here is add a product to my cart where the tagid of the product is equal to the hexstring of an nfc tag im reading in. 

I have the correct hexstring of the tag saved into my firebase but it's not adding the product to cart when tested in app.

Can supply more code if necessary
## [4][Would you call this "lateinit"?](https://www.reddit.com/r/typescript/comments/fqhsep/would_you_call_this_lateinit/)
- url: https://www.reddit.com/r/typescript/comments/fqhsep/would_you_call_this_lateinit/
---
interface Name {  
name: string;  
}  
let a!: Name;  
function setup(): void {  
a = { name: "Mike" };  
}  
setup();

console.log(a);

I think this "a + exclamation mark"-Thing (null assertion operator is misleading in my opinion) always as 'lateinit' in the Kotlin sense of meaning  - do you agree?
## [5][How do you keep your code tight when reverting back to vanilla JS](https://www.reddit.com/r/typescript/comments/fqdxmg/how_do_you_keep_your_code_tight_when_reverting/)
- url: https://www.reddit.com/r/typescript/comments/fqdxmg/how_do_you_keep_your_code_tight_when_reverting/
---
Typescript has showed me the true path to readable, easy to modify code but I now find myself in a JS heavy environment with no option to switch to TS. It seems every refactor I do opens a slew of `undefined` errors. In general I find the code hard to extend and even reason about, since it's so easy to misread the chain of inputs and outputs between functions.

I think I need to make a habit out of JSDoc comments. I used that for a couple months before the TS switch and now it seems like it's time to go back.

Is this the best vanila JS defense against `undefined` errors and unexpected inputs or do you guys practice anything else?
## [6][Experienced Typescript Developer for my Podcast](https://www.reddit.com/r/typescript/comments/fq4c3w/experienced_typescript_developer_for_my_podcast/)
- url: https://www.reddit.com/r/typescript/comments/fq4c3w/experienced_typescript_developer_for_my_podcast/
---
Hi all!  I'm looking to interview someone on my podcast, Contentful { Creators }, on TypeScript.  If you are interested, please email me at [marcelo@themarcelolewin.com](mailto:marcelo@themarcelolewin.com) \- You can listen to the podcast here - [https://anchor.fm/contentful-creators](https://anchor.fm/contentful-creators)
## [7][Partial&lt;T&gt; is not assignable to T, help please](https://www.reddit.com/r/typescript/comments/fq7vjk/partialt_is_not_assignable_to_t_help_please/)
- url: https://www.reddit.com/r/typescript/comments/fq7vjk/partialt_is_not_assignable_to_t_help_please/
---
Hi,

I'm coding some tests to unit tests some services in my NestJS application and I encountered some errors using Partial&lt;T&gt;.

&amp;#x200B;

The error :

`error TS2322: Type 'Partial&lt;User&gt;' is not assignable to type 'User'.`

`Property 'id' is optional in type 'Partial&lt;User&gt;' but required in type 'User'.`

&amp;#x200B;

The interface :

`import { Document, Schema } from 'mongoose'`  
`export enum UserGender {`  
 `FEMALE = "FEMALE",`  
 `MALE = "MALE"`  
`}`

`export interface User extends Document {`  
`id: string;`  
`email: string;`  
`password: string;`  
`firstname: string;`  
`lastname: string;`  
`address: string;`  
`city: string;`  
`zipcode: string;`  
`country: string;`  
`gender: UserGender;`  
`}`

The function causing this error :

`const mockUserDoc: (mock?: {`  
  `id?: string;`  
  `email?: string;`  
  `password?: string;`  
  `firstname?: string;`  
  `lastname?: string;`  
  `address?: string;`  
  `city?: string;`  
  `zipcode?: string;`  
  `country?: string;`  
  `gender?: UserGender;`  
`}) =&gt; Partial&lt;User&gt; = (mock?: {`  
  `id: string;`  
  `email: string;`  
  `password: string;`  
  `firstname: string;`  
  `lastname: string;`  
  `address: string;`  
  `city: string;`  
  `zipcode: string;`  
  `country: string;`  
  `gender: UserGender;`  
`}) =&gt; {`  
 `return {`  
`_id: (mock &amp;&amp; mock.id) || "an id",`  
`email: (mock &amp;&amp; mock.email) || "john.doe@test.com",`  
`password: (mock &amp;&amp; mock.password) || "password",`  
`firstname: (mock &amp;&amp; mock.firstname) || "John",`  
`lastname: (mock &amp;&amp; mock.lastname) || "Doe",`  
`address: (mock &amp;&amp; mock.address) || "address",`  
`city: (mock &amp;&amp; mock.city) || "city",`  
`country: (mock &amp;&amp; mock.country) || "country",`  
`zipcode: (mock &amp;&amp; mock.zipcode) || "zipcode",`  
`gender: (mock &amp;&amp; mock.gender) || UserGender.MALE,`  
  `};`  
`};`

`mockUserDoc()`

I can't see where the error come from, I'm assuming it could come from the enum but still don't know why ...
## [8][Elm style JSON decoder for TypeScript, providing type safety compile- and run-time](https://www.reddit.com/r/typescript/comments/fptbvb/elm_style_json_decoder_for_typescript_providing/)
- url: https://github.com/venil7/json-decoder
---

## [9][Approximating haskell's do syntax in Typescript](https://www.reddit.com/r/typescript/comments/fpvsxu/approximating_haskells_do_syntax_in_typescript/)
- url: https://paulgray.net/do-syntax-in-typescript/
---

## [10][Any game developers here? I need to find an easy Cocos Creator Typescript documentation. Any help please?](https://www.reddit.com/r/typescript/comments/fq4bfa/any_game_developers_here_i_need_to_find_an_easy/)
- url: https://www.reddit.com/r/typescript/comments/fq4bfa/any_game_developers_here_i_need_to_find_an_easy/
---

## [11][Recursive Type Cast](https://www.reddit.com/r/typescript/comments/fq0549/recursive_type_cast/)
- url: https://www.reddit.com/r/typescript/comments/fq0549/recursive_type_cast/
---
Hello,

I have two interfaces.  One that looks like this:

    interface Group {
      name: string;
      options: Option[];
    }
    
    interface Option {
      name: string;
      optionGroups: Group[];
    }

and

    interface AnotherGroup {
      name: string;
      options: AnotherOption[];
    }
    
    interface AnotherOption {
      name: string;
      optionGroups: AnotherGroup[];
    }

I then have a react component that expects an item that looks like this

    interface Item {
      options: Group[];
    }

I want to be able to pass an item that has options of type AnotherGroup.

The type definitions basically overlap but, but the type checker complains because they are different interfaces..

Is there a way that I can recursively cast `AnotherGroup[]` to `Group[]`?
