# Hash-based Archive and Storage Handler aka HASH

HASH is a git like toolkit for a version control system. It is fundamentally a content addressable system. As of current development status, it doesn't support VCS but it is in dev pipeline.

## Plumbing

The current version of HASH is built solely upon, what is referred to in git terminology as plumbing commands. They are low-level subcommands that can be chained together in UNIX-style or called from scripts.

### Supported Commands:
1.  `hash init`: Initializes an empty hash repository.
2. `hash hash-object -w <filepath_name>`: Creates an object and stores it as a blob with a unique sha_1 name in .hash/objects directory.
3. `hash update-index`: Adds the file or the object provided to the index or the staging area.
   a. `--fromcache`: If this flag is used you must provide the following arguments:
     * `<MODE>`: 100644 for a blob.
     * <`SHA`>: The sha_1 id of the blob object
     * `<FILENAME>`: The file name to be associated in the index with that entry.
   b. You can provide the filename directly to pick that file from the current working directory instead of the cache.
4. `hash write-tree`: Creates an index tree and returns the sha_1
5. `hash cat-file <SHA_1>`: Provide the sha_1 of the blob or the tree and it returns the content of the file.
6. `hash commit-tree <SHA_1> -p <PARENT_SHA>`: Creates a commit tree. If -p is not provided then it has no parent commit. -p should not be provided for first commits.
7. `hash update-ref <REF_PATH> <Commit_SHA>`: Creates a link between the provided ref and the provided commit SHA_1
8. `hash log`: Logs all the commits starting from latest. 

Ref: For better understanding of the commands and the working, read Git Internals available on the official Git website.

## Porcelain

In development. To be added soon
