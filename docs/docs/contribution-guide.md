# Contribution guide

We are pleased to know that you want to contribute to Kiwi. This guide will
walk you through making your first contribution and general guidelines to
make changes and committing to the codebase.

## Report issues

To report an issue, open one on the [GitHub issue tracker](https://github.com/sdslabs/kiwi/issues).
An issue should include the following information:

1. Platform (Operating System used)
2. Package in which you're facing the issue
3. Go version (or Node in case of docs)
4. What is the issue?
5. How to reproduce it?

Once an issue is opened, someone will get back to you to either help you
resolve it or start working on solving it themselves.

## Make changes in codebase

There are a few important things to take care of when making changes in the
codebase.

### Tests

Make sure all your test cases pass. Modify tests if required.
Add new test cases when adding new features.

### Lint

We use [golangci-lint](https://golangci-lint.run/). Make sure your changes
pass the lint tests. You can test your changes by running:

```sh
$ golangci-lint run

# Some errors can be resolved automatically. Use:
$ golangci-lint run --fix
```

### Commits

1. Break only logical changes into multiple commits. Commits such as "fix
   typo" or "address review commits" should be squashed into the one logical
   commit.

2. Each commit should individually pass tests and lint check. No separate
   commit should be made to fix these.

3. We don't have a defined commit message style for our codebase but the
   general idea is that the commit should include a heading, a body (if it's
   required) and reference to any issue that it might resolve. A good commit
   message looks something like this:

   ```
   Short commit heading with gist of changes.

   Body of commit trying to explain the old behaviour and how this commit
   changes it for the better.

   Resolves: #123

   Signed-off-by: Contributer <example@contributor.com>
   ```
