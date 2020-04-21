# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (17/2020)?](https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/)
- url: https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-17-2020/41247?u=llogiq)!
## [3][RFC: Transition to rust-analyzer as our official LSP implementation](https://www.reddit.com/r/rust/comments/g5ckqi/rfc_transition_to_rustanalyzer_as_our_official/)
- url: https://github.com/rust-lang/rfcs/pull/2912
---

## [4][Emulsion is a 100% Rust image viewer.](https://www.reddit.com/r/rust/comments/g5d5j9/emulsion_is_a_100_rust_image_viewer/)
- url: https://www.reddit.com/r/rust/comments/g5d5j9/emulsion_is_a_100_rust_image_viewer/
---
Emulsion is fast, small, displays transparency properly, and uses GPU acceleration. It's also minimalistic with the design philosophy of being nothing more than an image viewer but being very good at that.

Check out the website [https://arturkovacs.github.io/emulsion-website/](https://arturkovacs.github.io/emulsion-website/)

Below are some measurements that measure the time from pressing enter on the image file and having the image shown in the Emulsion window. These were made relying on my reflexes and a stopwatch. Originally I intended to make measurements for Windows' Photos but it's suprisingly difficult to get comparable values because the Photos app on my machine takes around 4-5 seconds to start up after a reboot while emulsion is around half of that. But then Photos takes around 1.3 seconds consistently to load almost any image, large and small alike. Photos loads images progressively which Emulsion doesn't do yet so Emulsion falls short in load time for very high resolutions.

    4798 * 4798 JPG    1.07 sec
    2880 * 1800 JPG    0.85 sec
    194 * 255 PNG      &lt;0.5 sec
## [5][Rust lang Tips and Tricks | Blog Post](https://www.reddit.com/r/rust/comments/g5caun/rust_lang_tips_and_tricks_blog_post/)
- url: https://mudit.blog/rust-tips-and-tricks/
---

## [6][Testing sync at Dropbox](https://www.reddit.com/r/rust/comments/g51v59/testing_sync_at_dropbox/)
- url: https://dropbox.tech/infrastructure/-testing-our-new-sync-engine
---

## [7][rlua needs maintainers!](https://www.reddit.com/r/rust/comments/g545jz/rlua_needs_maintainers/)
- url: https://www.reddit.com/r/rust/comments/g545jz/rlua_needs_maintainers/
---
Hey all, I'm the original author of the [rlua](https://github.com/kyren/rlua) crate, and I don't have a lot of time to spend on it these days, and I'm looking for additional maintainers.


I don't actually use Lua for anything currently, and I also don't have much interest in returning to Lua in the future.  As such, my Lua related crates `rlua` (and also `luster`, but that's much less important) are not receiving much attention currently.


That's a real shame, and I'd like to help at least `rlua` not become abandoned.  It needs some work done to it, none of which is terribly complex but I just don't have the time or proper motivation to be the one to do it.  Also, I don't actually even think I *should* be the one to do it, as it's very hard to make good APIs when you don't use your own API.


I'd be happy to guide people through my ideas for `rlua` and help especially with reviewing PRs for unsoundness, but there is quite a lot of stuff that needs doing in the pretty near term.  The internals of `rlua` are *pretty* hairy, which is why I don't expect anybody who wants to take up the mantle to go it completely alone, I'll happily be around to help with the hard parts.  What it needs though is better leadership and organization skills than I have or at least have to spare currently.


There seems to currently be an active fork of an old version of `rlua` called [mlua](https://github.com/khvzak/mlua), and I'd be happy to just step aside and let `mlua` be the primary Lua bindings system, however it appears that `mlua` gives up on soundness, and I think that's a real shame.  I worked very hard to try to find a *sound* bindings system for Lua, and more or less `rlua` today actually *is* a sound bindings system, and I'd like that effort not to go to waste.  (There are some caveats here but an important detail is that it is not the *bindings* system that is unsound, it's PUC-Rio Lua itself, but that's a WHOLE separate discussion.  I'll also admit that it's possible that soundness is just not something that people generally care about when dealing with Lua, and if so then that's fine but it would be a deal-breaker for me personally.  If you're the author of `mlua` then feel free also to reach out to me and we can discuss things, but I get the impression that maybe you have different priorities and that's okay!)


I'd be happy to go into (much) more detail about this with anybody who's interested, either in replies here or in PMs or in a matrix chat room or however you'd like to do it.  Reach out to me if you use or depend on `rlua` and have the skill and inclination to help.

Edit: [a pinned issue about this](https://github.com/kyren/rlua/issues/172)
## [8][Summary of Cubeb Oxidation on Mac OS](https://www.reddit.com/r/rust/comments/g57zcp/summary_of_cubeb_oxidation_on_mac_os/)
- url: https://chunminchang.github.io/blog/post/summary-of-cubeb-oxidation-on-mac-os
---

## [9][[ANN] RustCrypto: `signature` v1.0 and `ed25519` v1.0 crates](https://www.reddit.com/r/rust/comments/g4w8by/ann_rustcrypto_signature_v10_and_ed25519_v10/)
- url: https://www.reddit.com/r/rust/comments/g4w8by/ann_rustcrypto_signature_v10_and_ed25519_v10/
---
Announcing two new v1.0 releases from the https://github.com/RustCrypto/ project:

## `signature`: generic traits for producing and verifying digital signatures

- GitHub: https://github.com/RustCrypto/traits/tree/master/signature
- Crate: https://crates.io/crates/signature
- Docs: https://docs.rs/signature

## `ed25519`: multi-provider abstraction for producing and verifying Ed25519 signatures

- GitHub: https://github.com/RustCrypto/signatures/tree/master/ed25519
- Crate: https://crates.io/crates/ed25519
- Docs: https://docs.rs/ed25519

[Read more about the Ed25519 signature system here](https://ed25519.cr.yp.to/).

See also the v0.5 release of the [`ecdsa` crate](https://github.com/RustCrypto/signatures/tree/master/ecdsa), which supports the v1.0 release of the `signature` crate.

These crates enable writing code which is generic over the underlying signature implementation (and potentially, even the signature system!).

The `ed25519` crate does not contain an implementation of Ed25519, but instead contains an [`ed25519::Signature`](https://docs.rs/ed25519/latest/ed25519/struct.Signature.html) type which can be leveraged to write code in a way where consumers of that code can plug in specific "providers" for signing and verifying Ed25519 signatures.

[See the `ed25519` crate documentation for a complete usage example](https://docs.rs/ed25519/1.0.1/ed25519/#using-ed25519-generically-over-algorithm-implementationsproviders). Here is a short one:

    use ed25519::signature::{Signer, Verifier};

    pub struct HelloSigner&lt;S&gt; {
        pub signer: S
    }

    impl&lt;S&gt; HelloSigner&lt;S&gt;
    where
        S: Signer&lt;ed25519::Signature&gt;
    {
        pub fn sign(&amp;self, person: &amp;str) -&gt; ed25519::Signature {
            self.signer.sign(format!("Hello, {}!", person).as_bytes())
        }
    }

Several notable ecosystem Ed25519 libraries can be used as providers, including `ed25519-dalek`, *ring*, `sodiumoxide`, and `yubihsm`. [Please see the `ed25519` crate documentation for more information on supported providers](https://docs.rs/ed25519/1.0.1/ed25519/#available-ed25519-providers).

As an example, the above `HelloSigner` type can be instantiated with `ed25519-dalek` using the following (which uses the `signatory-dalek` wrapper crate from the [Signatory](https://github.com/iqlusioninc/signatory) project):

    use signatory_dalek::Ed25519Signer;

    /// `HelloSigner` instantiated with an `ed25519-dalek` signature provider
    pub type DalekHelloSigner = HelloSigner&lt;Ed25519Signer&gt;;

There's also an open PR to [natively support this API in `ed25519-dalek`](https://github.com/dalek-cryptography/ed25519-dalek/pull/124).

Enjoy!
## [10][A clarification reference for splitting code into multiple files](https://www.reddit.com/r/rust/comments/g5659q/a_clarification_reference_for_splitting_code_into/)
- url: https://www.reddit.com/r/rust/comments/g5659q/a_clarification_reference_for_splitting_code_into/
---
Context: I'm a zealous Rust beginner working on some personal projects. I have read a lot of the Book, the Async Book, and the Wasm Book, but I've made only 1 actual thing in Rust so far. I have previously used Java &amp; C++, as well as some functional languages like Haskell and Elm. I consider myself reasonably fluent in those languages.

I also think that clarifications are underrated, and I like to see them **all in one place, rather than scattered around various different web-pages**.

&amp;#x200B;

So, anyway, when I was reading the Book and got to chapter 7, I tried implementing my own reasonably-sized project using modules. When I tried it out, I was very confused. Rust was saying that it could not find my modules, and that I had to rename my files and move them around.

&amp;#x200B;

After re-reading it about 3 times, I think I understand it now. But for my sake and for the sake of other beginners who may be coming from these languages, I decided to make a single set of 16 clarifications that I would have found very useful from the beginning. They are all present in the Book, but they aren't listed in one place like they are here. Please correct all of the following statements which are incorrect. I will do my best to update them so that they can serve as a reference to future beginners:

&amp;#x200B;

1. Any project that you create with cargo is called a crate. (I initially thought that "crate" was a synonym for "external library", but this is NOT the case) EDIT: any Rust package is called a crate, whether or not Cargo is used. But it can be helpful to think of the entire "thing" that Cargo generates as a crate.
2. Crates can have multiple modules within them. Modules are intended to put related bits of code together.
   1. Modules can contain other modules.
3. Privacy in Rust is determined by module, not by file. You can have multiple modules in the same file hiding things from each other.
   1. Note that everything in a single file is considered as if it is within one big implicit module, unless you specifically use the "module" keyword.
4. In Rust, modules, structs, and functions are private by default.
5. Since modules can be recursive (modules within modules), a tree structure makes sense. However, this implies a root, which may not always be there. Rust solves this by adding an implicit root to the module tree, called "crate". Note that this "crate" is just a name, it's not saying that there is a crate inside your crate.
6. (For statements 7-12, assume that the whole project is in a single file, lib.rs or main.rs, depending on which type of project you're making)
7. If you have a module called "my\_module" in your file, that module's full name is really "crate::my\_module"
8. Using full names of modules can be cumbersome, so you can instead use relative names. Since the whole file is considered as a module ("crate"), you can just call it "my\_module" instead of "crate::my\_module"
9. Consider the following example (sorry for formatting but it doesn't let me use a code block):  
mod my\_module {  
pub mod my\_inner\_module {  
pub fn something ... (...) { ... }  
}  
}  
some\_function() {  
crate::my\_module::my\_inner\_module::something ... (...);  
//(alternatively)  
// my\_module::my\_inner\_module::something ... (...);  
}  
**In this case, my\_module and some\_function are both in the module crate, which means that it's okay for my\_module to not be public.** However, some\_function cannot see the module my\_inner\_module, which is why it must be declared public. Same for the function something ...
10. Sticking with the previous example, if we wanted to use the inner module without typing the entire path each time, we could use the "use" keyword (somewhat similar to C++'s "using" keyword for namespaces). For example:  
mod my\_module {  
pub mod my\_inner\_module {  
pub fn something ... (...) { ... }  
}  
}  
use my\_module::my\_inner\_module;  
// (alternatively)  
// use crate::my\_module::my\_inner\_module;some\_function() {  
my\_inner\_module::something ... (...);  
}  
NOTE that unlike C++, if you use the "use" keyword, you still have to include the inner module name when you use functions from it. "use" just lets you not have to write the path to the module all the time.
11. You can use a module "as" another name. This one seems fairly intuitive to me so I'm not going to explain it.
12. Nested paths and glob (also seem pretty intuitive to me):instead of saying  
use std::io;use std::cmp::Ordering;  
you can just say  
use std::{cmp::Ordering, io};  
If you want to use everything in std, you can just say  
use std::\*;
13. Separating your project into multiple files:  
13.1 Unlike many other languages, the file path actually changes the module name in Rust.  
13.2 There is one file called the "crate root file" which is either lib.rs or main.rs (depending on which type of project you're making)  
13.3 You know how every file has an implied module? (see statement 3.1) The name of this module is actually only "crate" if it's the crate root file, otherwise it's "crate" followed by the relative path from your source directory to that file. Yes, this means that all your sources must be in the source directory.  
13.3.1 (Reddit doesn't like multi-level lists) Example:If I have a file in src/level1/level2 called my\_mod.rs, its implied module name is crate::level1::level2::my\_mod  
13.4 (RIP Reddit formatting) If you want to use a module from another file, you have to use the "mod" keyword. Here is where the distinction between "mod" and "use" gets confusing. Remember: if you are including things from the same file, or from std, or from a different crate, you only have to use "use". But if you are including things from a different file within your own crate, you have to first write a "mod file::implied::module::name;" statement to tell Rust to load that file as a module. This is as if you had said "use" on its implied module name. You still have to use "use" for any declared modules within that file.
14. Finally, "pub use". If you have a module1 (remember, modules can either be explicitly declared, or the implicit module of the file itself) which uses module2 and module2 which itself uses module3, module1 does ***NOT*** "inherit" the use of module3 from module2. To change this, you have to change module2. Instead of "use"-ing module3, you have to "pub use". That way, the use itself becomes public.
15. You might see the term "extern crate" in some old StackOverflow answers. However, this is no longer necessary. You used to have to explicitly state your dependencies in your code, but Cargo takes care of that these days.
16. (expansion of statement 13) There is also a thing called "mod.rs" Basically the idea is, if you want the name of your implied package to be the directory, not the file name, you can put a "mod.rs" file in that directory. Example: the implied module name of "src/foo.rs" is "crate::foo", the implied name of "src/foo/bar.rs" is "crate::foo::bar", but the implied name of "src/foo/mod.rs" is "crate::foo", **NOT** "crate::foo::mod". mod.rs files are specially named modules, just like your lib/main.rs file. But for all other files, their module name contains their file name as well. See u/Lucretiel's answer for another example.

Edit: fixed statements 1 and 15. Added statement 16.
## [11][Open-sourcing dotenv-linter: lightning-fast tool to lint your .env files](https://www.reddit.com/r/rust/comments/g5eq8d/opensourcing_dotenvlinter_lightningfast_tool_to/)
- url: https://evrone.com/dotenv-linter
---

## [12][Wormhole: an easy http tunnel to localhost for development or sharing via a public url. Built on async-io using tokio.](https://www.reddit.com/r/rust/comments/g4zh7x/wormhole_an_easy_http_tunnel_to_localhost_for/)
- url: https://github.com/agrinman/wormhole
---

