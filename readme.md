# cubase-rename
Rename WAV files generated by "Channel Batch Export" function of cubase.

# Usage
`cubase-rename` target WAV files in the format `[prefix] - [index] - Audio - [track name].wav` in current directory.  
These are renamed to `[new prefix][track name][new suffix].wav`.

```
Usage:
  cubase-rename [flags]

Flags:
  -h, --help            help for cubase-rename
  -p, --prefix string   new prefix of file name
  -s, --suffix string   new suffix of file name
```