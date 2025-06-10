# 🔍 Simple Search Engine with Caching in Go

This is a lightweight search engine implemented in Go. It loads documents from a compressed XML file (`.xml.gz`), tokenizes and indexes them, and supports fast keyword-based search using an inverted index. The index is cached using JSON to avoid reprocessing on every run. The search runs in an interactive terminal session where you can enter queries until you type `exit`.

## 📦 Features

- Loads and parses `.xml.gz` files
- Builds a reverse index for fast token lookup
- Tokenizes abstract text using a basic analyzer
- Caches the index in `Cache/*.json` to skip rebuilding
- Searches for documents matching all query terms
- REPL-style prompt to run multiple queries in one session

## 🚀 How to Run

```bash
go run main.go -p /path/to/file.gz
```

Once started, type any search query in the terminal. Type `exit` or `quit` to stop the program.

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
- If the file exists, it loads the cached inverted index
- If not, it builds the index and writes the cache to disk
- This makes future runs significantly faster

## 💬 Sample Output

```txt
> go run main.go
Loaded 600000 documents in 2.1s
Indexed 600000 documents in 1.5s
Type your search queries below (type 'exit' to quit'):
> apollo moon landing
Found 4 results in 38ms
Title: Apollo 11
URL: http://dbpedia.org/resource/Apollo_11
Description: Apollo 11 was the spaceflight that first landed humans on the Moon...
--
...
```

## 🛠 Requirements

- Go 1.18 or later
