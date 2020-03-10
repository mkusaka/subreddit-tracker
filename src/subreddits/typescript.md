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
## [2][Is the TypeScript Playground the best place to learn right now?](https://www.reddit.com/r/typescript/comments/fgd3ax/is_the_typescript_playground_the_best_place_to/)
- url: https://www.reddit.com/r/typescript/comments/fgd3ax/is_the_typescript_playground_the_best_place_to/
---
Or are there better resources? I'm a new react developer learning typescript. Suggestions are very much appreciated! TIA!
## [3][Up-to-date overview of JavaScript and TypeScript tooling](https://www.reddit.com/r/typescript/comments/ffxjhj/uptodate_overview_of_javascript_and_typescript/)
- url: https://tooling.js.org/
---

## [4][Functional React components with generic props in TypeScript](https://www.reddit.com/r/typescript/comments/ffv4c0/functional_react_components_with_generic_props_in/)
- url: https://wanago.io/2020/03/09/functional-react-components-with-generic-props-in-typescript/
---

## [5][tsParticles CodePen collection](https://www.reddit.com/r/typescript/comments/fg51rq/tsparticles_codepen_collection/)
- url: https://codepen.io/collection/DPOage
---

## [6][Understand Evolving any (Effective TypeScript sample item)](https://www.reddit.com/r/typescript/comments/ffxfo4/understand_evolving_any_effective_typescript/)
- url: https://effectivetypescript.com/2020/03/09/evolving-any/
---

## [7][Recursive Multi-Threaded Worker Processes in NodeJS](https://www.reddit.com/r/typescript/comments/ffgw9i/recursive_multithreaded_worker_processes_in_nodejs/)
- url: https://github.com/sinclairzx81/threadbox
---

## [8][Music App with electron and react](https://www.reddit.com/r/typescript/comments/ff94e4/music_app_with_electron_and_react/)
- url: https://www.reddit.com/r/typescript/comments/ff94e4/music_app_with_electron_and_react/
---
I was bored, so I made this music app with the help of `electron` and `react` in typescript.

It can play music out of any directory you choose and can even download songs into that directory. It stores metadeta about each song and shows you all artists and albums which the songs are from and stores youre last queue so you can pickup where you left off.

 

&amp;#x200B;

https://preview.redd.it/mk3zlsszeel41.png?width=1400&amp;format=png&amp;auto=webp&amp;s=eaeae0b6b5c3a05b5e9347358935c68cbb60ed74

[You can view and play all your music](https://preview.redd.it/vwmnemrzeel41.png?width=1400&amp;format=png&amp;auto=webp&amp;s=b5d37af9cf9d41a228b812b6b742e0152f1ecaa6)

[A basic controller which has basic functionality](https://preview.redd.it/mhxk4frzeel41.png?width=522&amp;format=png&amp;auto=webp&amp;s=59ffd63743579809b829cd94ecb17f1c968b90dc)

[You can customize all the functionality](https://preview.redd.it/69ydvtrzeel41.png?width=1400&amp;format=png&amp;auto=webp&amp;s=da0c3583ef7c2899daaafeca94c343a8b1887d97)

&amp;#x200B;

The project is far from perfect and any suggestions will be welcome.

[GitHub Link](https://github.com/Lutetium-Vanadium/Music)

&amp;#x200B;

PS:

 I use linux so I'm not sure if there will be any bugs in mac or windows

Also you require a Napster API key(for music metadata) and a Google API key(for downloading a video) for proper functionality
## [9][Using Typescript without Typescript](https://www.reddit.com/r/typescript/comments/fes93e/using_typescript_without_typescript/)
- url: https://www.dandoescode.com/blog/using-typescript-without-typescript/
---

## [10][Validate function parameters when passed alongside function](https://www.reddit.com/r/typescript/comments/fewodc/validate_function_parameters_when_passed/)
- url: https://www.reddit.com/r/typescript/comments/fewodc/validate_function_parameters_when_passed/
---
Hi, I have a use case where I need to pass a reference to a method and the parameters to call it with at a later point. How do I validate that the parameters are valid upon declaration?

Example of what I need:

```
interface Foo {
    method: Function
    params: /* Parameters of whatever is passed to "method" */
}

const api = {
    method: async (a: number, b: string) =&gt; {},
}

// This should compile
const foo: Foo = {
    method: api.method,
    params: [3, 'baz']
}

// This should NOT compile
const bar: Foo = {
    method: api.method,
    params: ['wrong', 3]
}
```
## [11][TypeORM relationship - relationships by id ?](https://www.reddit.com/r/typescript/comments/ff0hcr/typeorm_relationship_relationships_by_id/)
- url: https://www.reddit.com/r/typescript/comments/ff0hcr/typeorm_relationship_relationships_by_id/
---
I've been kinda confused by the relationships as I'm used to save relationship by id, while docs and examples I found suggest to get the entire object and use that instead (Isn't this strange???)

I found this on github addressing this issue ( [https://github.com/typeorm/typeorm/issues/447](https://github.com/typeorm/typeorm/issues/447) ) , where they suggest to use an object with just the id property, but it's from 2017. Is that a good way to do it ? And is it still the only way to do it ? (I find it pretty lame tbh)

    async create( @Body() product: Product) {
        product.category = &lt;any&gt;{ id: product.category };
        return { payload: await this.repository.persist(product) };
    }

Another one suggested to name the column as categoryId and it would work as expected (with id instead of object) but WHY? What does the name have to do with that ??

    @Entity()
    class Product {
    
         @Column({ type: "int", nullable: true })
         categoryId: number;
    
         @ManyToOne(type =&gt; Category)
         @JoinColumn({ name: "categoryId" })
         category: Category;
    
    }

&amp;#x200B;

I'm just confused, help \^\_\^
