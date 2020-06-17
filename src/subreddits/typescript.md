# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][A collection of Algorithms and Data Structures with video lectures in Typescript [Update 2]](https://www.reddit.com/r/typescript/comments/haoirx/a_collection_of_algorithms_and_data_structures/)
- url: https://github.com/jeffzh4ng/algorithms-and-data-structures
---

## [3][How To Build A Shopify Headless eCommerce Storefront](https://www.reddit.com/r/typescript/comments/haq4kv/how_to_build_a_shopify_headless_ecommerce/)
- url: http://selleo-shopify.xyz
---

## [4][Help with indexing by enum](https://www.reddit.com/r/typescript/comments/hagg2r/help_with_indexing_by_enum/)
- url: https://www.reddit.com/r/typescript/comments/hagg2r/help_with_indexing_by_enum/
---
I'm not 100% sure if what I'd like to do is possible, but here it is:

I've got an object "map" where the key is some enum, say `ActionType`, and the value is some subscriber function. I'm trying to make it so that when I do:

    subscribe(ActionType.SomeAction, (response) =&gt; {})

`response` is automagically typed correctly based on the value of the `ActionType` passed in. I've tried a couple ways without any success, but I often get tripped up with these more complex type definitions, so maybe I'm just overlooking something/not aware of some technique to do this.

Any help is appreciated!
## [5][typescript object type alias signature working!!!](https://www.reddit.com/r/typescript/comments/hapvof/typescript_object_type_alias_signature_working/)
- url: https://www.reddit.com/r/typescript/comments/hapvof/typescript_object_type_alias_signature_working/
---
 `type personType = { name: string }`

`const displayPerson = (person2: personType) =&gt; {     console.log(person2) }`

`let person1 = {     name: "aaa",     address: "bbb",     age: 30 }`

`displayPerson(person1)` 

This code should give error, because person1 type is not personType but the code is still running with out any error or warning!!!
## [6][When type inference goes wrong, use an identity function to set it right](https://www.reddit.com/r/typescript/comments/ha5rk7/when_type_inference_goes_wrong_use_an_identity/)
- url: https://effectivetypescript.com/2020/06/16/typed-identity-functions/
---

## [7][Newbie help TS2739](https://www.reddit.com/r/typescript/comments/ha53i2/newbie_help_ts2739/)
- url: https://www.reddit.com/r/typescript/comments/ha53i2/newbie_help_ts2739/
---
Hey everyone, I have the following function in React

    providers.reduce((GlobalProvider, NewProvider) =&gt; ({ children }) =&gt; (
          &lt;GlobalProvider&gt;
            &lt;NewProvider&gt;{children}&lt;/NewProvider&gt;
          &lt;/GlobalProvider&gt;
        ));

and i'm receiving the following error:

`({ children }: any) =&gt; Element' is missing the following properties from type 'LazyExoticComponent&lt;FC&lt;any&gt;&gt;': _result, $$typeof`

Basically this function takes an array of Lazy Components and nests them into one.
GlobalProvider, NewProvider are `React.LazyExoticComponent&lt;any&gt;`

Thanks for any help
## [8][Single value as a return type?](https://www.reddit.com/r/typescript/comments/ha1pog/single_value_as_a_return_type/)
- url: https://www.reddit.com/r/typescript/comments/ha1pog/single_value_as_a_return_type/
---
VSCode doesn't lint if I set the return type as a static value. I have an orchestrator method that should only ever err or return true:

       public run(): true {

Another option is to just set it to `void`. That actually sounds more sensible.

Does it ever make sense to set the return type to a static non-void value?
## [9][Anybody know how to set the type parameter of a default value?](https://www.reddit.com/r/typescript/comments/ha76sv/anybody_know_how_to_set_the_type_parameter_of_a/)
- url: https://www.reddit.com/r/typescript/comments/ha76sv/anybody_know_how_to_set_the_type_parameter_of_a/
---
    function createBatches&lt;type&gt;(
        {OMITTED}: createBatches_Interface,
        batchClass: { new (batchRecipe: batchRecipe&lt;type&gt;): Batch&lt;type&gt; } = Batch&lt;type&gt;
    )


The default parameter for `batchClass` throws an error. Does anybody know how to work around this?
## [10][TS noob needs help figuring out why TS is enforcing “key” and “ref” props on a React Class Component.](https://www.reddit.com/r/typescript/comments/ha9o4q/ts_noob_needs_help_figuring_out_why_ts_is/)
- url: https://www.reddit.com/r/typescript/comments/ha9o4q/ts_noob_needs_help_figuring_out_why_ts_is/
---
Hello everyone! I’m working on a react component library which I decided to write in TS and compile back to ES5.

Everything is working as expected except for a tiny “warning/error” that TS throws.

I’m at work rn on a 15 min break so I can’t provide detailed info and code examples which I know is a bad way to ask for help, but I’m hoping someone out there ran I to this problem and can guide me a bit.

So TS yells at me for missing the “key” and “ref” props on my component. The thing is that I never wrote those props as part of my component props, so I’m guessing something in React’s typings of the class component is attaching these “ghost” props.

When I actually include the key and ref prop the warning goes away and everything is happy. But it’s annoying that I HAVE to include these props when 1. I never wrote them into the props, and 2. I don’t have any use for them.

As of now I simply include the key and ref props and set them to “undefined” and everything works as expected. But it’s annoying that I NEED to do that.

For people that need more details or code examples, I’ll write a more detailed post if this post doesn’t yield any responses. But hopefully I’m not the first guy running into this and someone out there can point me down the right path.

Thank you guys! Have a good one!
## [11][thetutlage/japa the missing Typescript-Typescript unit test library](https://www.reddit.com/r/typescript/comments/h9p03q/thetutlagejapa_the_missing_typescripttypescript/)
- url: https://github.com/thetutlage/japa
---

