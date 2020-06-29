# CISsors

CISsors is a tool for converting CIS benchmark rules from a PDF to a YAML format.

This can be quite useful if you need further processing of these rules in other
tools and scripts.

## Installation

```bash
go get -u github.com/xornivore/cissors/...
```

## Usage

```bash
cissors --help
usage: cissors [<flags>] <file>

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose              Verbose mode.
  -o, --out=OUT              Output to file.
      --id-prefix=ID-PREFIX  ID prefix for rules.
```

## Example

```bash
cissors CIS_Kubernetes_Benchmark_v1.5.0.pdf -o cis-kubernetes-1.5.0-parsed.yaml
✂️  You are now running with CISsors

✂️  Skillfully cutting your CIS benchmark️

✂️  Found 120 rules

✂️  Done extracting rules

✂️️️  All done! Enjoy your YAML masterpiece!
```
