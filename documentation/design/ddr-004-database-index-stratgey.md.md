# Database Index Strategy

## Changelog

Nov 28th, 2020: Created the record template with draft content

## Abstract

## Design

### events Table

Because of the extensive and important use of events table on the `height` column, we have made the decision to use a B-tree index on the `height` column to speed up the system performance. The trade-off is the space requirement as B-tree is a lossless indexing strategy. In the future we may consider [BRIN](https://www.postgresql.org/docs/9.5/brin-intro.html) with a smaller `pages_per_range` after we have more data to analysis the performance and storage tradeoffs.

### view_blocks, view_transactions and view_events Table

These three projections are heavily used by our Block Explorer front-end, and we consider them as a few of the most commonly used data in a blockchain application. The way people usually access these table includes:
- accessing rows in a consecutive order
- random search on a particular record by the its unique identity

One characteristic of the data on these tables have a definite order and natural correlation. As a result, when we design the index for these table we decided to use a combination of BRIN and B-Tree index for a balance of performance and storage. For the `height` column we use BRIN while for other columns we use B-Tree (e.g. block hash, transaction hash).

Another property of the data usage is random search. While GiST and GIN are two indexing strategy, in the Postgres documentation it clearly state that they are targetted for full text search while in a Blockchain use case the searchable data are usually single words. The cloest we can get the `moniker` column of a validator but by studying the naming conventions it appears the number of words are relatively limited. As a result we decide a B-Tree idnex for these search-able columns. In the future we may re-visit GiST and GIN on its efficiency.

### view_validator and view_validator_activities Table

Validator and validator activities table are another two common tables being accessed.

Usually validator table size grows much gently. As a result we decide only to create two index on the operator address and consensus node address for facilitating projection performance.

Validator activities table is essentially a mix of `view_transactions` and `view_events`. As a result, the same strategy of BRIN and B-Tree hybrid index approach is applied.