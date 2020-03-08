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
## [2][Music App with electron and react](https://www.reddit.com/r/typescript/comments/ff94e4/music_app_with_electron_and_react/)
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
## [3][Using Typescript without Typescript](https://www.reddit.com/r/typescript/comments/fes93e/using_typescript_without_typescript/)
- url: https://www.dandoescode.com/blog/using-typescript-without-typescript/
---

## [4][Validate function parameters when passed alongside function](https://www.reddit.com/r/typescript/comments/fewodc/validate_function_parameters_when_passed/)
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
## [5][TypeORM relationship - relationships by id ?](https://www.reddit.com/r/typescript/comments/ff0hcr/typeorm_relationship_relationships_by_id/)
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
## [6][Optional chaining](https://www.reddit.com/r/typescript/comments/fe4r99/optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/fe4r99/optional_chaining/
---
For some reason I can't use optional chaining in my new project, other TS features like interfaces, types, ... are working fine. Also my IDE is not throwing out any errors, but upon compiling the code, I get this error for the following line of code:

Line of code: `country: location?.country`

`country: location?.country,`

`SyntaxError: Unexpected token '.'`

`at Module._compile (internal/modules/cjs/loader.js:895:18)`

What's going on?
## [7][How to replace all non-const enum value with literals? And should I do it?](https://www.reddit.com/r/typescript/comments/fe9e7a/how_to_replace_all_nonconst_enum_value_with/)
- url: https://www.reddit.com/r/typescript/comments/fe9e7a/how_to_replace_all_nonconst_enum_value_with/
---
## Problem

I want to optimize my TypeScript by replacing all literal enums with their values.

## Why not `const enum`?

* Third-party libraries may not provide them.
* Babel does not support them.
## [8][Best way to type an object with string keys?](https://www.reddit.com/r/typescript/comments/fe40sr/best_way_to_type_an_object_with_string_keys/)
- url: /r/LearnTypescript/comments/fe3zip/best_way_to_type_an_object_with_string_keys/
---

## [9][Different function return types?](https://www.reddit.com/r/typescript/comments/fe2mgz/different_function_return_types/)
- url: https://www.reddit.com/r/typescript/comments/fe2mgz/different_function_return_types/
---
Is it possible to have a function return different return types under different circumstances?

Edit: I mean without overloading.

example:

    const test = (name: string) : int | string =&gt; {
    
        if(name === 'beatdook04') return 200;
    
        return "Who are you?";    
    
    }
    
    test('beatdook04');
## [10][TypeBox - Mapping TypeScript as JSONSchema](https://www.reddit.com/r/typescript/comments/fdoolj/typebox_mapping_typescript_as_jsonschema/)
- url: https://github.com/sinclairzx81/typebox
---

## [11][Trying to store typescript code in string to use in runtime. [Guidance Appreciated]](https://www.reddit.com/r/typescript/comments/fdwf0e/trying_to_store_typescript_code_in_string_to_use/)
- url: https://www.reddit.com/r/typescript/comments/fdwf0e/trying_to_store_typescript_code_in_string_to_use/
---
As the description implies i'm trying to store actual code from typescript file in a variable so that when I console.log it i'll see the typescript implementation instead of the compiled code.  
Have any of you tried anything similar?   


What I've tried:  
1. Using decorators and Reflection,  but struggle getting the actual code (as a string) as it is in the file.

2. looking at the typescript npm package. There's just so many functions to read up on. 

3.  Tried to retrofit this code [this code](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API#user-content-re-printing-sections-of-a-typescript-file) from the typescript compiler API docs for my purposes.

So far I haven't had any luck. All pointers welcome ;)
