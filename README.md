# 🔍 Simple Search Engine with Caching in Go

This is a lightweight search engine implemented in Go. It loads documents from a compressed XML file (`.xml.gz`), tokenizes and indexes them, and supports fast keyword-based search using an inverted index. The index is cached using JSON to avoid reprocessing on every run.

## 📦 Features

- Loads and parses `.xml.gz` files
- Builds a reverse index for fast token lookup
- Tokenizes abstract text (basic analysis)
- Caches the index in `Cache/*.json` to skip rebuilding
- Searches for documents matching all query terms

## 🚀 How to Run

```bash
go run main.go
```

Make sure your XML file is in `data/` and compressed with `.gz`.

## 🧪 Sample XML Format

```xml
<entries>
  <entry>
    <title>Example Title</title>
    <url>http://example.com</url>
    <abstract>This is an example abstract.</abstract>
  </entry>
  ...
</entries>
```

## 📂 Cache Behavior

- The program checks for `Cache/<filename>.index_cache.json`
- If it exists, it loads the cached inverted index
- If not, it builds the index and writes the JSON cache

## 🛠 Requirements

- Go 1.18 or later
