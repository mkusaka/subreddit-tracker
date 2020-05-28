# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][[Need Help] Been stuck trying to implement AVL tree insertion. I'm not able to figure out if it's a problem with my code or logic since the code runs fine but output is not expected. Any help would be appreciated.](https://www.reddit.com/r/typescript/comments/gs3x8c/need_help_been_stuck_trying_to_implement_avl_tree/)
- url: https://www.reddit.com/r/typescript/comments/gs3x8c/need_help_been_stuck_trying_to_implement_avl_tree/
---
Been stuck on this problem for hours. Maybe its my bad debugging skills. So i was just trying to learn some data structures and try to implement them from scratch in Typescript so that i could learn both.

So during AVL tree insertion, everything just works fine until I balance the tree(commenting `balanceTree` and returning bare node gets correct height), whence all the `heights` are not the expected values. I'm not able to figure out if it's because of not properly passing references to `left`,`right`,`root etc.`Or somewhere my logic is wrong. Any help would be appreciated.

Here's a [pastebin link for the code](https://pastebin.com/sbczFNLq), if anyone wants a better preview.

This is the code:-

    class NodeItem {
        // key is the value assigned to a node
        key: number = 0;
        left: NodeItem = null;
        right: NodeItem = null;
        // height is the longest path from node to leaf
        height: number = 0;
        constructor(val: number) {
            this.key = val;
        }
    
        get _leftHeight(): number {
            if (this.left === null) {
                // if no node exists then let that height be 0
                return 0;
            }
            return this.left.height;
        }
    
        get _rightHeight(): number {
            if (this.right === null) {
                // if no node exists then let that height be 0
                return 0;
            }
            return this.right.height;
        }
    
        updateHeight() {
            // set height as the longest path from current node to leaf
            this.height = 1 + Math.max(this._leftHeight, this._rightHeight);
        }
    
        getBalanceFactor() {
            return this._leftHeight - this._rightHeight;
        }
    }
    
    class AVL {
        root: NodeItem = null;
    
        balanceTree(node: NodeItem, item: number) {
            /*
            &lt; -1 - left heavy
            &gt; 1  - right heavy
            0    - balanced - do nothing
            */
            const balanceFactor = node.getBalanceFactor();
            if (balanceFactor &lt; -1 &amp;&amp; item &gt; node.right.key) {
                // rightRight Case
                node = this.leftRotate(node);
            } else if (balanceFactor &lt; -1 &amp;&amp; item &lt; node.right.key) {
                // rightLeft Case
                node.right = this.rightRotate(node);
                return this.leftRotate(node);
            } else if (balanceFactor &gt; 1 &amp;&amp; item &lt; node.left.key) {
                // leftLeft Case
                node = this.rightRotate(node);
            } else if (balanceFactor &gt; 1 &amp;&amp; item &gt; node.left.key) {
                // leftRight Case
                node.left = this.leftRotate(node);
                return this.rightRotate(node);
            }
            return node;
        }
    
        leftRotate(first: NodeItem): NodeItem {
            const second = first.right;
            const secondLeft = second.left;
            second.left = first;
            first.right = secondLeft;
            first.updateHeight();
            second.updateHeight();
            return second;
        }
    
        rightRotate(first: NodeItem): NodeItem {
            const second = first.left;
            const secondRight = second.right;
            second.right = first;
            first.left = secondRight;
            first.updateHeight();
            second.updateHeight();
            return second;
        }
    
        insert(root: NodeItem, value: number): NodeItem {
            if (root === null) {
                // if no node exists, then create new node
                return new NodeItem(value);
            } else if (value &lt; root.key) {
                // go to left subtree
                root.left = this.insert(root.left, value);
            } else if (value &gt; root.key) {
                // go to right subtree
                root.right = this.insert(root.right, value);
            } else {
                // element already exists in tree
                return root;
            }
            root.updateHeight();
            return this.balanceTree(root, value);
        }
    
        inOrder(root: NodeItem): void {
            if (root !== null) {
                this.inOrder(root.left);
                console.log(`value\t:${root.key}\theight\t:${root.height}`)
                this.inOrder(root.right);
            }
            return;
        }
    }
    
    const avl = new AVL();
    const inserts: number[] = [10, 5, 15, -10, -5];
    for (const i of inserts) {
        avl.root = avl.insert(avl.root, i);
    }
    avl.inOrder(avl.root);

Results:-

    Expected:
    ---------
    value	:-10	        height	:0
    value	:-5		height	:1              10(2)
    value	:5		height	:0         -5(1)     15(0)
    value	:10		height	:2     -10(0)   5(0)
    value	:15		height	:0
    
    Current Result:
    ---------------
    value	:-10	        height	:1
    value	:-5		height	:0              10(1)
    value	:5		height	:2         -5(0)     15(0)
    value	:10		height	:1     -10(1)   5(2)
    value	:15		height	:0
## [3][Type safety with raw SQL](https://www.reddit.com/r/typescript/comments/grsc2m/type_safety_with_raw_sql/)
- url: https://www.reddit.com/r/typescript/comments/grsc2m/type_safety_with_raw_sql/
---
I'm a bit new to this so bear with me if this is impossible/futile but is it possible to use raw sql to query into types/interfaces? Basically if it's possible to save the result of a raw sql query(for me specifically with mssql) into a typed variable(yes I know technically it doesn't actually exist in runtime but I want the development clarity in code). Or am I just avoiding an inevitable ORM? 

Basically how do I avoid an orm while keeping type annotations, as I haven't found any example doing so, all I see are TypeORM projects. Not that I have anything against TypeORM, I just don't like making my entities live objects with decorators and much prefer to keep them as simple dumb interfaces.

Edit: I see this is getting a bit of traction, so first of all I appreciate all the replies, and if anyone else joins this thread i'll just mention I specifically have to use SQL Server(not my first choice but my hands are tied). Answers relevant only to other SQL Services are welcome, just mention it if that's the case :)
## [4][Possible to use TypeScript to check JS only? Trying to figure out incremental conversion path for huge project](https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/)
- url: https://www.reddit.com/r/typescript/comments/grj31m/possible_to_use_typescript_to_check_js_only/
---
tl;dr trying to just use TypeScript as analysis tool is that possible?

Hey we have a large project some 100,000 lines of code with open contribution within the company. The code is abused and contributing on by hundreds of devs and people don't care about standards, they are just trying to get their code in to meet their deadline.

I am one of the core maintainers and our main defense is very strict testing requirements and static analysis on our CI which will block pull requests.

In the past when we've introduced new standards, we did it incrementally by making sure the code base did not get worse (e.g. we introduced coverage requirements by only checking that only new code was covered).

We are trying to find a similar path forward for TypeScript but have so far failed. 

However, I keep hearing the benefits of using TypeScript even without converting everything and fixing errors. I noticed the "checkJS" flag and found this pretty useful and was curious if we could run tsc with no output and just use it to check code and find certain errors TypeScript is good at like unresolved variables?
## [5][Typescript Gotchas when you come from back-end world](https://www.reddit.com/r/typescript/comments/grsy3k/typescript_gotchas_when_you_come_from_backend/)
- url: https://medium.com/@admiquel/typescript-gotchas-when-you-come-from-back-end-world-cf7f6c3d0d02?source=friends_link&amp;sk=18a82d9ce141e874971e89d64f41d7dc
---

## [6][TypeScript's Language Features Documentation?](https://www.reddit.com/r/typescript/comments/gr8jt4/typescripts_language_features_documentation/)
- url: https://www.reddit.com/r/typescript/comments/gr8jt4/typescripts_language_features_documentation/
---
I'm trying to get on-board with TypeScript coming from a JavaScript background. All of the documentation only talks about the Types and how to use them. But I'm finding hidden secrets like TypeScript supporting [Optional Chaining](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-7.html) or automatically setting instance variables based on private class constructor arguments.

```
class Test {
    constructor(private name) {}

    printName() {
        console.log(this.name); // &lt;- Somehow this is magically defined by TypeScript?
    }
}
```

I totally get the Type features that TypeScript adds. That's easy to follow and well documented. But there seems to be actual language features beyond just types that it adds to and I'd like to know what all of those are. Where can I find documentation for all of the changes to the JavaScript language TypeScript adds?
## [7][I recently spoke w/a recruiter about a role requiring substantial TS experience. TS is a superset of JS that applies strict and custom types to JS code. If two developers both have very strong JS but one has two years and the other six months of TS, is there really THAT much of a difference?](https://www.reddit.com/r/typescript/comments/gr1ybe/i_recently_spoke_wa_recruiter_about_a_role/)
- url: https://www.reddit.com/r/typescript/comments/gr1ybe/i_recently_spoke_wa_recruiter_about_a_role/
---
I know that TypeScript has its idioms, its ins and outs, etc., but what I'm questioning is whether there's that substantial a difference in two strong JavaScript engineers who have different amounts of experience specifically with TypeScript. TypeScript itself doesn't seem to be a "years of experience" sort of technology; it's very easy to pick up and, with some practice and study, to implement properly. It's more like Lodash (eh, maybe that's a bit crude a comparison, but you take my meaning--it's just a tool that serves a specific purpose in JavaScript, but not a fundamentally different paradigm. It's not like going from JavaScript to, say, Java.)

Edit: A bit more information.

I picked up TypeScript about three months ago and have been using it since. I like type safety, but TypeScript itself doesn't seem to be so radical a departure from JavaScript that I'd eliminate a strong engineer from the running if they have only a few months of experience with the technology. It's really not that tough to pick up and use properly.

Let me know if I'm off-base here.
## [8][End-to-end Type Safety in Clean Architecture: a possible solution with TypeScript, GraphQL, MongoDB, React.](https://www.reddit.com/r/typescript/comments/gqx14r/endtoend_type_safety_in_clean_architecture_a/)
- url: https://charlesagile.com/end-to-end-type-safety
---

## [9][Why is it so hard to iterate over objects in TypeScript?](https://www.reddit.com/r/typescript/comments/gqxh2h/why_is_it_so_hard_to_iterate_over_objects_in/)
- url: https://effectivetypescript.com/2020/05/26/iterate-objects/
---

## [10][Is there a way to have TypeScript recompile my project when I save a file without the use of an IDE? (I'm using Vim.)](https://www.reddit.com/r/typescript/comments/gr1o6h/is_there_a_way_to_have_typescript_recompile_my/)
- url: https://www.reddit.com/r/typescript/comments/gr1o6h/is_there_a_way_to_have_typescript_recompile_my/
---

## [11][Mutex-Server, a new npm module I've developed (+ need your advise)](https://www.reddit.com/r/typescript/comments/gqso87/mutexserver_a_new_npm_module_ive_developed_need/)
- url: https://www.npmjs.com/package/mutex-server
---

