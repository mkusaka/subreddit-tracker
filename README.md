# subreddit-tracker

this repository track subreddit at reddit.

now track following subreddit.
- https://www.reddit.com/r/androiddev
- https://www.reddit.com/r/Kotlin
- https://www.reddit.com/r/reduxjs
- https://www.reddit.com/r/reactjs
- https://www.reddit.com/r/typescript
- https://www.reddit.com/r/javascript
- https://www.reddit.com/r/cpp
- https://www.reddit.com/r/golang
- https://www.reddit.com/r/ruby
- https://www.reddit.com/r/rust
- https://www.reddit.com/r/aws
- https://www.reddit.com/r/rails

# custom
## track list
Above list comes from my interest. So, if you want to custom subreddit track list, fork this respository & change main.go's Subreddits list, like follows.

```diff
var (
	Subreddits = []string{
-		"androiddev",
+   "awsomereddit"
		"Kotlin",
		"reduxjs",
		"reactjs",
		"typescript",
		"javascript",
		"cpp",
		"golang",
		"ruby",
		"rust",
	}
)
```
## schedule
If you want to custom PR timing, set `.github/workflows/daily.yml`'s on:schedule:cron: to any timing you want.

# ref
