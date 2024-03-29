
Example manifest could be like :

```yaml
kind: ConfigManifest
version: 0.1
# different sources needed when you want to have
# different credentials for each path
# credentials shouldn't be put directly into the url
# but rather passing it into tole arguments directly
sources:
  team01-vault: vault+https://cluster01.company.internal/v1.1/company/team01
  team02-vault: vault+https://cluster01.company.internal/v1.1/company/team02
  local-env: dotenv+rel://.env
# you could use alias if you want to short things out
aliases:
  project01-staging: team01-vault://kv/project_01/stg
#
# relative path are being defined as scheme extension rather than
# using proper URI syntax, since no library in golang actually able to
# deal with proper relative URI syntax and somehow adding it into
# this project or create a library that doing that are just so
# out of context of this project (so sed)
#
# why the source need to be in uri format too ?
#
# - the idea is I might be able to put the template source directly into sources
#   and watch the template source from sources (but still need to declare the source)
#   into `sources` section rather than directly in `mounts` section.
#
# - it able to extend more specific arguments that related to underlying source or target
#   implementation
#
#
# Special for source & target in case referencing aliases or sources, it will be expanded
# ahead of time.
#
mounts:
  - source: file+rel://template/database.yml.tmpl
    target: file+rel://config/database.yml?uid=<uid>&gid=<gid>&mode=<mode>

  - source: file+rel://template/database.yml.tmpl
    target: file+rel://config/database.yml?uid=<uid>

  - source: local-env://
    target: env://

  - source: file+rel://template/env.tmpl
    target: env://

  - source: glob+rel://template/**/*.tmpl
    target: dir+rel://output/

```
```
In the end when real lookup happens the path will expand itself beyond sources & aliases.

Example:

`project01-staging environment` will expand into `vault+https://cluster01.company.internal/v1.1/company/team01/kv/project_01/stg`.


## Lookup Pattern

```
<source>://<host>/<version>/<namespace>/<type>/<service>/<env>/<path>
```
