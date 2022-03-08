ATTACH TABLE _ UUID '885deb18-e333-4540-bfbb-fa06322144ed'
(
    `i` Int64,
    `j` Int16,
    `k` Int64
)
ENGINE = MergeTree
ORDER BY i
SETTINGS index_granularity = 8192
